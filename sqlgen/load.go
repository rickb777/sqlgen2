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

func load(pkgStore parse.PackageStore, name parse.LType, mainPkg string) (*TableDescription, error) {
	table := new(TableDescription)

	// local map of indexes, used for quick lookups and de-duping.
	indices := map[string]*Index{}

	table.Type = name.Name
	table.Name = Pluralize(Underscore(table.Type))

	nm := pkgStore.FindNamed(name)
	tags := pkgStore.FindTags(name)
	ctx := &context{pkgStore, indices, table, mainPkg}
	ctx.examineStruct(nm, name, tags, nil)

	for _, idx := range ctx.indices {
		table.Index = append(table.Index, idx)
	}

	checkNoConflictingNames(name, table)

	return table, nil
}

//-------------------------------------------------------------------------------------------------

func (ctx *context) examineStruct(nm *types.Named, name parse.LType, tags map[string]parse.Tag, parent *Node) bool {
	parse.DevInfo("examineStruct %s %+v\n  tags %v\n", name, nm, tags)
	if nm == nil {
		return false
	}

	str := nm.Underlying().(*types.Struct)
	if str.NumFields() == 0 {
		exit.Fail(1, "%s: empty structs are not supported (was there a parser warning?).\n", name)
	}

	var unexportedFields []string

	for j := 0; j < str.NumFields(); j++ {
		tField := str.Field(j)
		if !tField.Exported() {
			tag := tags[tField.Name()]
			if !tag.Skip {
				unexportedFields = append(unexportedFields, tField.Name())
			}
		}
	}

	if len(unexportedFields) == str.NumFields() {
		output.Info("%s: Info: %s contains no exported fields; it must implement sql.Scanner and driver.Valuer.\n", name, str.String())
		return true // add this struct as a field; assume the type implements Scanner/Valuer

	} else if len(unexportedFields) > 0 {
		output.Info("%s: Warning: %s.%s contains unexported fields %s.\n"+
			"  (perhaps annotate with `sql:\"-\"`)\n", name, nm.Obj().Pkg().Name(), nm.Obj().Name(), strings.Join(unexportedFields, ", "))
	}

	for j := 0; j < str.NumFields(); j++ {
		tField := str.Field(j)
		parse.DevInfo("    f%-2d: name:%-15s pkg:%-10s type:%-50s field:%v, exp:%v, anon:%v\n", j,
			tField.Name(), tField.Pkg().Name(), tField.Type(), tField.IsField(), tField.Exported(), tField.Anonymous())

		tag := tags[tField.Name()]
		if !tag.Skip {
			if tField.Anonymous() {
				ctx.convertEmbeddedNodeToFields(tField, name.PkgName, parent)

			} else {
				ctx.convertLeafNodeToField(tField, name.PkgName, tags, parent)
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
	lt := parse.LType{pkg, name}
	nm := ctx.pkgStore.FindNamed(lt)
	path := ""

	if nm == nil {
		var ok bool
		nm, ok = leaf.Type().(*types.Named)
		if !ok {
			exit.Fail(5, "Unable to find %s\n", lt)
		}
		nmPkg := nm.Obj().Pkg()
		path = nmPkg.Path()
		pkg = nmPkg.Name()
		str2 := nm.Underlying().(*types.Struct)
		tags = make(map[string]parse.Tag)
		addStructTags(tags, str2)
		parse.DevInfo(" - found in other package %v %v\n", leaf.Type(), str2)
	} else {
		nmPkg := nm.Obj().Pkg()
		pkg = nmPkg.Name()
		tags = ctx.pkgStore.FindTags(lt)
	}

	node := &Node{Name: name, Type: Type{PkgPath: path, PkgName: pkg, Name: name, Base: parse.Struct}, Parent: parent}
	ctx.examineStruct(nm, lt, tags, node)
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

	switch leaf.Type().Underlying().(type) {
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

	switch t := lt.(type) {
	case *types.Pointer:
		lt = t.Elem()
		tp.IsPtr = true
	}

	switch nm := lt.(type) {
	case *types.Basic:
		tp.Name = nm.Name()
		tp.Base = parse.Kind(nm.Kind())

	case *types.Named:
		tObj := nm.Obj()
		setTypeName(&tp, tObj, pkg, ctx.mainPkg)
		parse.DevInfo("named %+v\n", tp)

		switch u := nm.Underlying().(type) {
		case *types.Basic:
			tp.Base = parse.Kind(u.Kind())

		case *types.Slice:
			tp.Base = parse.Slice

		case *types.Struct:
			tp.Base = parse.Struct
			if canRecurse {
				addStructTags(tags, u)
				ok := ctx.examineStruct(nm, parse.LType{pkg, leaf.Name()}, tags, &node)
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
	tp.Name = tn.Name()
	tnPkg := tn.Pkg()
	if tnPkg.Name() != mainPkg {
		tp.PkgPath = tnPkg.Path()
		tp.PkgName = tnPkg.Name()
		//parse.DevInfo("setTypeName %s %s %s\n", tp.Name, tp.PkgPath, tp.PkgName)
	} else if pkg != mainPkg {
		tp.PkgName = pkg
		//parse.DevInfo("setTypeName %s - %s\n", tp.Name, tp.PkgName)
	//} else {
	//	parse.DevInfo("setTypeName %s - -\n", tp.Name)
	}
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

func checkNoConflictingNames(name parse.LType, table *TableDescription) {
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
		parse.DevInfo("checkNoConflictingNames %s %+v\n", name, utter.Sdump(table))
		exit.Fail(1, "%s: found conflicting SQL column names: %s.\nPlease set the names on these fields explicitly using tags.\n",
			name, strings.Join(duplicates, ", "))
	}
}
