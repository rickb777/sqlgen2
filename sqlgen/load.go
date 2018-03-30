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
	"sort"
	"github.com/kortschak/utter"
)

type context struct {
	pkgStore parse.PackageStore
	indices  map[string]*Index
	table    *TableDescription
	mainPkg  string
	fileTags parse.Tags
}

func load(pkgStore parse.PackageStore, name parse.LType, mainPkg string, fileTags parse.Tags) (*TableDescription, error) {
	table := new(TableDescription)

	// local map of indexes, used for quick lookups and de-duping.
	indices := map[string]*Index{}

	table.Type = name.Name
	table.Name = Pluralize(Underscore(table.Type))

	nm := pkgStore.FindNamed(name)
	tags := pkgStore.FindTags(name)

	ctx := &context{pkgStore, indices, table, mainPkg, fileTags}
	ctx.examineStruct(nm, name, tags, nil)

	for _, idx := range ctx.indices {
		table.Index = append(table.Index, idx)
	}

	sort.Slice(table.Index, func(i, j int) bool {
		return table.Index[i].Name < table.Index[j].Name
	})

	checkNoConflictingNames(name, table)

	return table, nil
}

func mergeTags(structTags, fileTags parse.Tags) parse.Tags {
	merged := make(parse.Tags)
	for n, t := range structTags {
		merged[n] = t
	}
	for n, t := range fileTags {
		merged[n] = t
	}
	parse.DevInfo("merged tags\n-----------\n%s\n", merged.String())
	return merged
}

func checkNoConflictingNames(name parse.LType, table *TableDescription) {
	names := make(map[string]struct{})
	var duplicates Identifiers

	for _, name := range table.Fields.SqlNames() {
		_, exists := names[name]
		if exists {
			duplicates = append(duplicates, name)
		}
		names[name] = struct{}{}
	}

	if len(duplicates) > 0 {
		parse.DevInfo("checkNoConflictingNames %s %+v\n", name, utter.Sdump(table))
		exit.Fail(1, "%s: found conflicting SQL column names: %s.\nPlease set the names on these fields explicitly using tags.\n",
			name, duplicates.MkString(", "))
	}
}

//-------------------------------------------------------------------------------------------------

func isBasicInterface(t types.Type) bool {
	_, isInterface := t.(*types.Interface)
	return isInterface
}

func isTypeNamed(name string, t types.Type) bool {
	nm, isNamed := t.(*types.Named)
	if !isNamed {
		return false
	}
	return nm.Obj().Name() == name
}

func recogniseScannerValuer(nm *types.Named) (isScanner, isValuer bool) {
	for i := 0; i < nm.NumMethods(); i++ {
		method := nm.Method(i)
		name := method.Name()
		tp := method.Type().(*types.Signature)
		params := tp.Params()
		recvType := tp.Recv().Type()
		results := tp.Results()

		switch name {
		case "Scan":
			_, isPtr := recvType.(*types.Pointer)
			if isPtr && !tp.Variadic() && params.Len() == 1 && results.Len() == 1 {
				p0IsInterface := isBasicInterface(params.At(0).Type())
				r0IsError := isTypeNamed("error", results.At(0).Type())
				isScanner = p0IsInterface && r0IsError
				parse.DevInfo("recogniseScannerValuer %s %s %v\n", nm, method.FullName(), isScanner)
			}

		case "Value":
			if !tp.Variadic() && params.Len() == 0 && results.Len() == 2 {
				r0IsValue := isTypeNamed("Value", results.At(0).Type())
				r1IsError := isTypeNamed("error", results.At(1).Type())
				isValuer = r0IsValue && r1IsError
				parse.DevInfo("recogniseScannerValuer %s %s %v\n", nm, method.FullName(), isValuer)
			}
		}
	}
	return
}

//-------------------------------------------------------------------------------------------------

func (ctx *context) examineStruct(nm *types.Named, name parse.LType, tags parse.Tags, parent *Node) (addStructAsField bool) {
	merged := mergeTags(tags, ctx.fileTags)
	parse.DevInfo("examineStruct %s %+v\n -- tags\n%v\n", name, nm, merged)
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
			tag := merged[tField.Name()]
			if !tag.Skip {
				unexportedFields = append(unexportedFields, tField.Name())
			}
		}
	}

	if len(unexportedFields) == str.NumFields() {
		output.Info("%s: Info: %s contains no exported fields; it must implement sql.Scanner and driver.Valuer.\n", name, str.String())
		return true

	} else if len(unexportedFields) > 0 {
		output.Info("%s: Warning: %s.%s contains unexported fields %s.\n"+
			"  (perhaps annotate with `sql:\"-\"`)\n", name, nm.Obj().Pkg().Name(), nm.Obj().Name(), strings.Join(unexportedFields, ", "))
	}

	for j := 0; j < str.NumFields(); j++ {
		tField := str.Field(j)
		parse.DevInfo("    f%-2d: name:%-15s pkg:%-10s type:%-50s field:%v, exp:%v, anon:%v\n", j,
			tField.Name(), tField.Pkg().Name(), tField.Type(), tField.IsField(), tField.Exported(), tField.Anonymous())

		tag := merged[tField.Name()]
		if !tag.Skip {
			if tField.Anonymous() {
				ctx.convertEmbeddedNodeToFields(tField, name.PkgName, parent)

			} else {
				ctx.convertLeafNodeToField(tField, name.PkgName, merged, parent)
			}
		}
	}

	return false
}

//-------------------------------------------------------------------------------------------------

func (ctx *context) convertEmbeddedNodeToFields(leaf *types.Var, pkg string, parent *Node) {
	var tags parse.Tags

	name := leaf.Name()
	parse.DevInfo("convertEmbeddedNodeToFields %s %s\n", pkg, name)
	lt := parse.LType{pkg, name}
	nm := ctx.pkgStore.FindNamed(lt)
	path := ""

	if nm == nil {
		var ok bool
		nm, ok = leaf.Type().(*types.Named)
		if !ok {
			exit.Fail(5, "unable to find %s\n", lt)
		}
		nmPkg := nm.Obj().Pkg()
		path = nmPkg.Path()
		pkg = nmPkg.Name()
		str2 := nm.Underlying().(*types.Struct)
		tags = make(parse.Tags)
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

func (ctx *context) convertLeafNodeToField(leaf *types.Var, pkg string, tags parse.Tags, parent *Node) {
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

	prefix := ""
	if tag.Prefixed {
		prefix = Underscore(field.JoinParts(1, "_")) + "_"
	}

	if tag.Name != "" {
		field.SqlName = string(prefix + strings.ToLower(tag.Name))
	} else {
		field.SqlName = string(prefix + strings.ToLower(field.Name))
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

	if tag.Natural {
		tag.Unique = field.SqlName + "_idx"
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

	ctx.table.Fields = append(ctx.table.Fields, field)
}

//-------------------------------------------------------------------------------------------------

func (ctx *context) convertLeafNodeToNode(leaf *types.Var, pkg string, tags parse.Tags, parent *Node, canRecurse bool) (Node, bool) {
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

		tp.IsScanner, tp.IsValuer = recogniseScannerValuer(nm)

		switch u := nm.Underlying().(type) {
		case *types.Basic:
			tp.Base = parse.Kind(u.Kind())

		case *types.Slice:
			tp.Base = parse.Slice

		case *types.Struct:
			tp.Base = parse.Struct
			if canRecurse && !tp.IsScanner && !tp.IsValuer {
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

func addStructTags(tags parse.Tags, str *types.Struct) {
	for i := 0; i < str.NumFields(); i++ {
		ts := str.Tag(i)
		tag, err := parse.ParseTag(ts)
		if err != nil {
			exit.Fail(2, "%s contains unparseable tag %q (%s)", str.String(), ts, err)
		}
		tags[str.Field(i).Name()] = *tag
	}
}
