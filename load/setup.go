package load

import (
	"bytes"
	"fmt"
	"github.com/acsellers/inflections"
	"github.com/rickb777/sqlapi/schema"
	stypes "github.com/rickb777/sqlapi/types"
	"github.com/rickb777/sqlapi/util"
	"github.com/rickb777/sqlgen2/output"
	"github.com/rickb777/sqlgen2/parse"
	"github.com/rickb777/sqlgen2/parse/exit"
	"go/types"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func PackagesToImport(flags FuncFlags, pgx bool) util.StringSet {
	imports := util.NewStringSet(
		"context",
		"database/sql",
		"strings",
		"github.com/pkg/errors",
		"github.com/rickb777/sqlapi/dialect",
	)

	base := ""

	if pgx {
		base = "github.com/rickb777/sqlapi/pgxapi"
		imports.Add("github.com/jackc/pgx")
	} else {
		base = "github.com/rickb777/sqlapi"
	}

	imports.Add(
		base,
		base+"/constraint",
	)

	if flags.Exec || flags.Query || flags.Count || flags.Update || flags.Delete || flags.Slice {
		imports.Add(base + "/support")
	}

	if flags.Exec || flags.Query || flags.Insert || flags.Update || flags.Delete || flags.Slice {
		imports.Add("github.com/rickb777/sqlapi/require")
	}

	if flags.Insert || flags.Update || flags.Schema {
		imports.Add("bytes")
	}
	if flags.Count || flags.Insert || flags.Select || flags.Slice || flags.Delete {
		imports.Add("fmt")
	}
	if flags.Count || flags.Select || flags.Slice || flags.Update || flags.Delete {
		imports.Add("github.com/rickb777/where")
	}
	if flags.Select {
		imports.Add("github.com/rickb777/where/quote")
	}
	return imports
}

func LastDirName(full string) string {
	abs, err := filepath.Abs(full)
	if err != nil {
		exit.Fail(1, "%s: %s.\n", full, err)
	}
	d1, _ := filepath.Split(abs)
	_, f2 := filepath.Split(filepath.Clean(d1))
	return f2
}

func PrimaryInterface(table *schema.TableDescription, genSchema bool) string {
	return primaryInterface(table, genSchema)
}

func primaryInterface(table *schema.TableDescription, genSchema bool) string {
	if !genSchema {
		return "Table"
	}
	if len(table.Index) == 0 {
		return "TableCreator"
	}
	return "TableWithIndexes"
}

func SecondaryInterface(flags FuncFlags) string {
	return secondaryInterface(flags)
}

func secondaryInterface(flags FuncFlags) string {
	return "Table"
}

//-------------------------------------------------------------------------------------------------

type FuncFlags struct {
	Schema, Query, Exec, Select, Count, Insert, Update, Upsert, Delete, Slice, Scan bool
}

var AllFuncFlags = FuncFlags{true, true, true, true, true, true, true, true, true, true, true}

//-------------------------------------------------------------------------------------------------

// GoFmt formats a template using gofmt.
func GoFmt(in io.Reader) (io.Reader, error) {
	var out bytes.Buffer

	gofmt := exec.Command("gofmt", "-s")
	gofmt.Stdin = in
	gofmt.Stdout = &out
	gofmt.Stderr = os.Stderr
	err := gofmt.Run()
	return &out, err
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

func (ctx *context) examineStruct(nm *types.Named, name parse.LType, tags stypes.Tags, parent *schema.Node) (addStructAsField bool) {
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
			if doNotSkip(tag) {
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
		if doNotSkip(tag) {
			if tField.Anonymous() {
				ctx.convertEmbeddedNodeToFields(tField, name.PkgName, parent)

			} else {
				ctx.convertLeafNodeToField(tField, name.PkgName, merged, parent)
			}
		}
	}

	return false
}

func doNotSkip(t *stypes.Tag) bool {
	if t == nil || !t.Skip {
		return true
	}
	return false
}

//-------------------------------------------------------------------------------------------------

func (ctx *context) convertEmbeddedNodeToFields(leaf *types.Var, pkg string, parent *schema.Node) {
	var tags stypes.Tags

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
		tags = make(stypes.Tags)
		addStructTags(tags, str2)
		parse.DevInfo(" - found in other package %v %v\n", leaf.Type(), str2)
	} else {
		nmPkg := nm.Obj().Pkg()
		pkg = nmPkg.Name()
		tags = ctx.pkgStore.FindTags(lt)
	}

	node := &schema.Node{Name: name, Type: schema.Type{PkgPath: path, PkgName: pkg, Name: name, Base: stypes.Struct}, Parent: parent}
	ctx.examineStruct(nm, lt, tags, node)
}

//-------------------------------------------------------------------------------------------------

func (ctx *context) convertLeafNodeToField(leaf *types.Var, pkg string, tags stypes.Tags, parent *schema.Node) {
	// tag may be nil
	tag := tags[leaf.Name()]
	field := &schema.Field{}
	field.Tags = tag

	// what follows needs default values
	if tag == nil {
		tag = &stypes.Tag{}
	}

	field.Encode = mapTagToEncoding[tag.Encode]

	// only recurse into the node's fields if the leaf isn't encoded
	var ok bool
	field.Node, ok = ctx.convertLeafNodeToNode(leaf, pkg, tags, parent, field.Encode == schema.ENCNONE)
	if !ok {
		return
	}

	switch leaf.Type().Underlying().(type) {
	case *types.Slice:
		field.Type.Base = stypes.Slice
	}

	field.SqlName = chooseSqlName(tag, &(field.Node), field.Name)

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
			index = &schema.Index{
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
			index = &schema.Index{
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

func chooseSqlName(tag *stypes.Tag, node *schema.Node, defaultName string) string {
	prefix := ""
	if tag.Prefixed {
		prefix = inflections.Underscore(node.JoinParts(1, "_")) + "_"
	}

	if tag.Name != "" {
		return string(prefix + strings.ToLower(tag.Name))
	}
	return string(prefix + strings.ToLower(defaultName))
}

//-------------------------------------------------------------------------------------------------

func (ctx *context) convertLeafNodeToNode(leaf *types.Var, pkg string, tags stypes.Tags, parent *schema.Node, canRecurse bool) (schema.Node, bool) {
	node := schema.Node{Name: leaf.Name(), Parent: parent}
	tp := schema.Type{}

	lt := leaf.Type()

	switch t := lt.(type) {
	case *types.Pointer:
		lt = t.Elem()
		tp.IsPtr = true
	}

	switch nm := lt.(type) {
	case *types.Basic:
		tp.Name = nm.Name()
		tp.Base = stypes.Kind(nm.Kind())

	case *types.Named:
		tObj := nm.Obj()
		setTypeName(&tp, tObj, pkg, ctx.mainPkg)
		parse.DevInfo("named %+v\n", tp)

		tp.IsScanner, tp.IsValuer = recogniseScannerValuer(nm)

		switch u := nm.Underlying().(type) {
		case *types.Basic:
			tp.Base = stypes.Kind(u.Kind())

		case *types.Slice:
			tp.Base = stypes.Slice

		case *types.Struct:
			tp.Base = stypes.Struct
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

func setTypeName(tp *schema.Type, tn *types.TypeName, pkg, mainPkg string) {
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

func addStructTags(tags stypes.Tags, str *types.Struct) {
	for i := 0; i < str.NumFields(); i++ {
		ts := str.Tag(i)
		tag, err := stypes.ParseTag(ts)
		if err != nil {
			exit.Fail(2, "%s contains unparseable tag %q (%s)", str.String(), ts, err)
		}
		tags[str.Field(i).Name()] = tag
	}
}
