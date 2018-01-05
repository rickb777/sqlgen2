package main

import (
	"fmt"
	"go/types"
	"strings"
	. "github.com/acsellers/inflections"
	. "github.com/rickb777/sqlgen2/schema"
	"github.com/rickb777/sqlgen2/sqlgen/parse"
	"github.com/rickb777/sqlgen2/sqlgen/parse/exit"
	"github.com/rickb777/sqlgen2/sqlgen/output"
	"github.com/kortschak/utter"
)

type context struct {
	pkgStore parse.PackageStore
	indices  map[string]*Index
	table    *TableDescription
	mainPkg  string
}

func load(pkgStore parse.PackageStore, pkg, name string) (*TableDescription, error) {
	table := new(TableDescription)

	// local map of indexes, used for quick lookups and de-duping.
	indices := map[string]*Index{}

	table.Type = name
	table.Name = Pluralize(Underscore(table.Type))

	nm := pkgStore.FindNamed(pkg, name)
	tags := pkgStore.FindTags(pkg, name)
	ctx := &context{pkgStore, indices, table, pkg}
	ctx.examineStruct(nm, pkg, name, tags, nil)

	for _, idx := range ctx.indices {
		table.Index = append(table.Index, idx)
	}

	checkNoConflictingNames(pkg, name, table)

	return table, nil
}

//-------------------------------------------------------------------------------------------------

func (ctx *context) examineStruct(nm *types.Named, pkg, name string, tags map[string]parse.Tag, parent *Node) bool {
	parse.DevInfo("examineStruct %s %s %+v\n  tags %v\n", pkg, name, nm, tags)
	if nm == nil {
		return false
	}

	str := nm.Underlying().(*types.Struct)
	if str.NumFields() == 0 {
		exit.Fail(1, "%s.%s: empty structs are not supported (was there a parser warning?).\n", pkg, name)
	}

	var unexportedFields []string

	for j := 0; j < str.NumFields(); j++ {
		tField := str.Field(j)
		//parse.DevInfo("    f%d: name:%-10s pkg:%s type:%-25s f:%v, e:%v, a:%v\n", j,
		//	tField.Name(), tField.Pkg().Name(), tField.Type(), tField.IsField(), tField.Exported(), tField.Anonymous())

		if !tField.Exported() {
			tag := tags[tField.Name()]
			if !tag.Skip {
				unexportedFields = append(unexportedFields, tField.Name())
			}
		}
	}

	if len(unexportedFields) == str.NumFields() {
		output.Info("%s.%s: Info: %s contains no exported fields; it must implement sql.Scanner and driver.Valuer.\n", pkg, name, str.String())
		return true // add this struct as a field; assume the type implements Scanner/Valuer

	} else if len(unexportedFields) > 0 {
		output.Info("%s.%s: Warning: %s.%s contains unexported fields %s.\n"+
			"  (perhaps annotate with `sql:\"-\"`)\n", pkg, name, nm.Obj().Pkg().Name(), nm.Obj().Name(), strings.Join(unexportedFields, ", "))
	}

	for j := 0; j < str.NumFields(); j++ {
		tField := str.Field(j)
		parse.DevInfo("    f%d: name:%-10s pkg:%s type:%-25s f:%v, e:%v, a:%v\n", j,
			tField.Name(), tField.Pkg().Name(), tField.Type(), tField.IsField(), tField.Exported(), tField.Anonymous())

		tag := tags[tField.Name()]
		if !tag.Skip {
			if tField.Anonymous() {
				ctx.convertEmbeddedNodeToFields(tField, pkg, parent)

			} else {
				ctx.convertLeafNodeToField(tField, pkg, tags, parent)
			}
		}
	}

	return false
}

//-------------------------------------------------------------------------------------------------

func (ctx *context) convertEmbeddedNodeToFields(leaf *types.Var, pkg string, parent *Node) {
	var tags map[string]parse.Tag

	name := leaf.Name()
	parse.DevInfo("convertEmbeddedNodeToFields %s %s\n", pkg, name)
	nm := ctx.pkgStore.FindNamed(pkg, name)

	if nm == nil {
		var ok bool
		nm, ok = leaf.Type().(*types.Named)
		if !ok {
			exit.Fail(5, "Unable to find %s.%s\n", pkg, name)
		}
		pkg = nm.Obj().Pkg().Name()
		str2 := nm.Underlying().(*types.Struct)
		tags = make(map[string]parse.Tag)
		addStructTags(tags, str2)
		parse.DevInfo(" - found in other package %v %v\n", leaf.Type(), str2)
	} else {
		pkg = nm.Obj().Pkg().Name()
		tags = ctx.pkgStore.FindTags(pkg, name)
	}

	node := &Node{Name: name, Parent: parent}
	ctx.examineStruct(nm, pkg, name, tags, node)
}

//-------------------------------------------------------------------------------------------------

func (ctx *context) convertLeafNodeToField(leaf *types.Var, pkg string, tags map[string]parse.Tag, parent *Node) {
	tag := tags[leaf.Name()]
	field := &Field{}
	field.Tags = tag
	field.Encode = mapTagToEncoding[tag.Encode]

	// only recurse into the node's fields if the leaf isn't encoded
	var ok bool
	field.Node, ok = ctx.convertLeafNodeToNode(leaf, pkg, tags, parent, field.Encode == ENCNONE)
	if !ok {
		return
	}

	// Lookup the SQL column type
	underlying := leaf.Type().Underlying()
	switch u := underlying.(type) {
	case *types.Basic:
		field.Type.Base = parse.Kind(u.Kind())

	case *types.Slice:
		field.Type.Base = parse.Slice
	}

	if tag.Primary {
		if ctx.table.Primary != nil {
			exit.Fail(1, "%s, %s: compound primary keys are not supported.\n",
				ctx.table.Primary.Type.Name, field.Type.Name)
		}
		ctx.table.Primary = field
	}

	if tag.Index != "" {
		index, ok := ctx.indices[tag.Index]
		if !ok {
			index = &Index{
				Name: tag.Index,
			}
			ctx.indices[index.Name] = index
		}
		index.Fields = append(index.Fields, field)
	}

	if tag.Unique != "" {
		index, ok := ctx.indices[tag.Unique]
		if !ok {
			index = &Index{
				Name: tag.Unique,
			}
			ctx.indices[index.Name] = index
		}
		index.Fields = append(index.Fields, field)
		index.Unique = true
	}

	if tag.Type != "" {
		base, ok := mapStringToSqlType[tag.Type]
		if ok {
			field.Type.Base = base
		} else {
			output.Info("%s.%s: Warning: unrecognised type %q\n  (allowed: %s)\n", pkg, leaf.Name(), tag.Type, allowedSqlTypeStrings())
		}
	}

	prefix := ""
	if tag.Prefixed {
		prefix = Underscore(field.JoinParts(1, "_")) + "_"
	}

	if tag.Name != "" {
		field.SqlName = prefix + strings.ToLower(tag.Name)
	} else {
		field.SqlName = prefix + strings.ToLower(field.Name)
	}

	ctx.table.Fields = append(ctx.table.Fields, field)
}

//-------------------------------------------------------------------------------------------------

func (ctx *context) convertLeafNodeToNode(leaf *types.Var, pkg string, tags map[string]parse.Tag, parent *Node, canRecurse bool) (Node, bool) {
	node := Node{Name: leaf.Name(), Parent: parent}
	tp := Type{}

	lt := leaf.Type()
	//isPtr := false

	switch t := lt.(type) {
	case *types.Pointer:
		lt = t.Elem()
		tp.IsPtr = true
	}

	switch nm := lt.(type) {
	case *types.Basic:
		tp.Name = nm.Name()
		//case *types.Struct:
		//	tp.Name = nm.String()
	case *types.Named:
		tObj := nm.Obj()
		setTypeName(&tp, tObj, pkg, ctx.mainPkg)
		parse.DevInfo("named %+v\n", tp)

		if str, ok := nm.Underlying().(*types.Struct); ok {
			tp.Base = parse.Struct
			if canRecurse {
				addStructTags(tags, str)
				ok = ctx.examineStruct(nm, pkg, leaf.Name(), tags, &node)
				node.Type = tp
				return node, ok
			}
		}
	case *types.Array:
		tp.Name = nm.String()
	case *types.Slice:
		switch el := nm.Elem().(type) {
		case *types.Basic:
			tp.Name = nm.String()
		case *types.Named:
			tnObj := el.Obj()
			setTypeName(&tp, tnObj, pkg, ctx.mainPkg)
			parse.DevInfo("slice %+v\n", tp)
		}
	default:
		panic(fmt.Sprintf("%#v", lt))
	}

	node.Type = tp
	return node, true
}

func setTypeName(tp *Type, tn *types.TypeName, pkg, mainPkg string) {
	if tn.Pkg().Name() != pkg {
		tp.PkgPath = tn.Pkg().Path()
		tp.PkgName = tn.Pkg().Name()
	} else if pkg != mainPkg {
		tp.PkgName = pkg
	}
	tp.Name = tn.Name()
}

func addStructTags(tags map[string]parse.Tag, str *types.Struct) {
	for i := 0; i < str.NumFields(); i++ {
		ts := str.Tag(i)
		tag, err := parse.ParseTag(ts)
		if err != nil {
			exit.Fail(2, "%s contains unparseable tag %q (%s)", str.String(), ts, err)
		}
		tags[str.Field(i).Name()] = *tag
	}
}

func checkNoConflictingNames(pkg, name string, table *TableDescription) {
	names := make(map[string]struct{})
	var duplicates []string

	for _, field := range table.Fields {
		name := strings.ToLower(field.SqlName)
		_, exists := names[name]
		if exists {
			duplicates = append(duplicates, name)
		}
		names[name] = struct{}{}
	}

	if len(duplicates) > 0 {
		parse.DevInfo("checkNoConflictingNames %s %s %+v\n", pkg, name, utter.Sdump(table))
		exit.Fail(1, "%s.%s: found conflicting SQL column names: %s.\nPlease set the names on these fields explicitly using tags.\n",
			pkg, name, strings.Join(duplicates, ", "))
	}
}
