package main

import (
	"bytes"
	"flag"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
	"github.com/kortschak/utter"
	"github.com/rickb777/sqlapi/schema"
	stypes "github.com/rickb777/sqlapi/types"
	"github.com/rickb777/sqlapi/util"
	. "github.com/rickb777/sqlgen2/code"
	"github.com/rickb777/sqlgen2/output"
	"github.com/rickb777/sqlgen2/parse"
	"github.com/rickb777/sqlgen2/parse/exit"
	"fmt"
	"os/exec"
	"github.com/acsellers/inflections"
	"sort"
	"go/types"
)

func main() {
	start := time.Now()

	var oFile, typeName, prefix, list, kind, tableName, tagsFile, genSetters string
	var flags = funcFlags{}
	var all, sselect, insert, gofmt bool

	flag.StringVar(&oFile, "o", "", "Output file name; optional. Use '-' for stdout.\n" +
		"\tIf omitted, the first input filename is used with '_sql.go' suffix.")
	flag.StringVar(&typeName, "type", "", "The type to analyse; required.\n" +
		"\tThis is expressed in the form 'pkg.Name'")
	flag.StringVar(&prefix, "prefix", "", "Prefix for names of generated types; optional.\n" +
		"\tUse this if you need to avoid name collisions.")
	flag.StringVar(&list, "list", "", "List type for slice of model objects; optional.")
	flag.StringVar(&kind, "kind", "Table", "Kind of model: you could use 'Table', 'View', 'Join' etc as required.")
	flag.StringVar(&tableName, "table", "", "The name for the database table; default is based on the struct name as a plural.")
	flag.StringVar(&tagsFile, "tags", "", "A YAML file containing tags that augment and override any in the Go struct(s); optional.\n" +
		"\tTags control the SQL type, size, column name, indexes etc.")
	flag.BoolVar(&output.Verbose, "v", false, "Show progress messages.")
	flag.BoolVar(&parse.Debug, "z", false, "Show debug messages.")
	flag.BoolVar(&parse.PrintAST, "ast", false, "Trace the whole astract syntax tree (very verbose).")
	flag.BoolVar(&gofmt, "gofmt", false, "Format and simplify the generated code nicely.")

	// filters for what gets generated
	flag.BoolVar(&all, "all", false, "Shorthand for '-schema -create -read -update -delete -slice'; recommended.\n" +
		"\tThis does not affect -setters.")
	flag.BoolVar(&sselect, "select", false, "Alias for -read")
	flag.BoolVar(&insert, "insert", false, "Alias for -create")
	flag.BoolVar(&flags.schema, "schema", false, "Generate SQL schema create/drop methods.")
	flag.BoolVar(&flags.insert, "create", false, "Generate SQL create (insert) methods.")
	flag.BoolVar(&flags.exec, "exec", false, "Generate Exec method. This is also provided with -update or -delete.")
	flag.BoolVar(&flags.sselect, "read", false, "Generate SQL select (read) methods.")
	flag.BoolVar(&flags.update, "update", false, "Generate SQL update methods.")
	flag.BoolVar(&flags.delete, "delete", false, "Generate SQL delete methods.")
	flag.BoolVar(&flags.slice, "slice", false, "Generate SQL slice (column select) methods.")
	flag.StringVar(&genSetters, "setters", "none", "Generate setters for fields of your type (see -type): none, optional, exported, all.\n" +
		"\tFields that are pointers are assumed to be optional.")

	flag.Parse()

	output.Require(flag.NArg() > 0, "At least one input file (or path) is required; put this after the other arguments.\n"+
		"  version   - prints the current version then exits.\n")

	if flag.Args()[0] == "version" {
		fmt.Printf("sqlgen %s\n       branch %s built on %s\n       origin %s\n", util.Version, util.GitBranch, util.BuildDate, util.GitOrigin)
		os.Exit(0)
	}

	if sselect {
		flags.sselect = true
	}

	if insert {
		flags.insert = true
	}

	if all {
		flags = allFuncFlags
	}

	output.Require(len(typeName) > 3, "-type is required. This must specify a type, qualified with its local package in the form 'pkg.Name'.\n", typeName)
	words := strings.Split(typeName, ".")
	output.Require(len(words) == 2, "type %q requires a package name prefix.\n", typeName)
	pkg, name := words[0], words[1]
	mainPkg := pkg

	// parse the Go source code file(s) to extract the required struct and return it as an AST.
	pkgStore, err := parse.Parse(flag.Args())
	output.Require(err == nil, "%v\n", err)
	//utter.Dump(pkgStore)

	if oFile == "" {
		oFile = flag.Args()[0]
		output.Require(strings.HasSuffix(oFile, ".go"), oFile+": must end '.go'")
		oFile = oFile[:len(oFile)-3] + "_sql.go"
		parse.DevInfo("oFile: %s\n", oFile)
	} else {
		mainPkg = lastDirName(oFile)
		parse.DevInfo("mainPkg: %s\n", mainPkg)
	}

	o := output.NewOutput(oFile)

	tags, err := stypes.ReadTagsFile(tagsFile)
	if err != nil && err != os.ErrNotExist {
		exit.Fail(1, "tags file %s failed: %s.\n", tagsFile, err)
	}

	// load the Tree into a schema Object
	table, err := load(pkgStore, parse.LType{pkg, name}, mainPkg, tags)
	if parse.Debug {
		utter.Dump(table)
	}

	if len(table.Fields) < 1 {
		exit.Fail(1, "no fields found. Check earlier parser warnings.\n")
	}

	view := NewView(name, prefix, tableName, list)
	view.Table = table
	view.Thing = kind
	view.Interface1 = primaryInterface(table, flags.schema)
	view.Interface2 = secondaryInterface(flags)

	setters := view.FilterSetters(genSetters)

	importSet := packagesToImport(flags, view.Table.HasPrimaryKey())

	if flags.sselect || flags.insert || flags.update {
		ImportsForFields(table, importSet)
	}
	ImportsForSetters(setters, importSet)

	buf := &bytes.Buffer{}

	WritePackageHeader(buf, mainPkg)

	WriteImports(buf, importSet)

	WriteType(buf, view)

	WritePrimaryDeclarations(buf, view)

	if flags.schema {
		WriteSchemaDeclarations(buf, view)
		WriteSchemaFunctions(buf, view)
	}

	if flags.exec || flags.update || flags.delete {
		WriteExecFunc(buf, view)
	}

	WriteQueryRows(buf, view)
	WriteQueryThings(buf, view)

	if flags.sselect {
		WriteScanRows(buf, view)
		WriteGetRow(buf, view)
		WriteSelectRowsFuncs(buf, view)
	}

	if flags.slice {
		WriteSliceColumn(buf, view)
	}

	if flags.insert {
		WriteConstructInsert(buf, view)
	}

	if flags.update {
		WriteConstructUpdate(buf, view)
	}

	if flags.insert {
		WriteInsertFunc(buf, view)
	}

	if flags.update {
		WriteUpdateFunc(buf, view)
	}

	if flags.delete {
		WriteDeleteFunc(buf, view)
	}

	WriteSetters(buf, view, setters)

	// formats the generated file using gofmt
	var pretty io.Reader = buf
	if gofmt {
		pretty, err = format(buf)
		output.Require(err == nil, "%s\n%v\n", string(buf.Bytes()), err)
	}

	o.Write(pretty, os.Stdout)

	output.Info("%s took %v\n", o.Path(), time.Now().Sub(start))
}

func packagesToImport(flags funcFlags, hasPrimaryKey bool) util.StringSet {
	imports := util.NewStringSet(
		"context",
		"database/sql",
		"log",
		"github.com/rickb777/sqlgen2",
		"github.com/rickb777/sqlgen2/constraint",
		"github.com/rickb777/sqlgen2/require",
		"github.com/rickb777/sqlgen2/schema",
		"github.com/rickb777/sqlgen2/support",
	)

	if flags.insert || flags.update || flags.schema {
		imports.Add("bytes")
	}
	if flags.insert || flags.update {
		imports.Add("io")
	}
	if flags.insert || flags.sselect || flags.slice || flags.delete {
		imports.Add("fmt")
	}
	if flags.sselect || flags.slice || flags.update || flags.delete {
		imports.Add("github.com/rickb777/sqlgen2/where")
	}
	return imports
}

func lastDirName(full string) string {
	abs, err := filepath.Abs(full)
	if err != nil {
		exit.Fail(1, "%s: %s.\n", full, err)
	}
	d1, _ := filepath.Split(abs)
	_, f2 := filepath.Split(filepath.Clean(d1))
	return f2
}

func primaryInterface(table *schema.TableDescription, genSchema bool) string {
	if !genSchema {
		return "sqlgen2.Table"
	}
	if len(table.Index) == 0 {
		return "sqlgen2.TableCreator"
	}
	return "sqlgen2.TableWithIndexes"
}

func secondaryInterface(flags funcFlags) string {
	if flags.exec && flags.sselect && flags.insert && flags.update && flags.delete && flags.slice {
		return "sqlgen2.TableWithCrud"
	}
	return "sqlgen2.Table"
}

//-------------------------------------------------------------------------------------------------

type funcFlags struct {
	schema, exec, sselect, insert, update, delete, slice bool
}

var allFuncFlags = funcFlags{true, true, true, true, true, true, true}

//-------------------------------------------------------------------------------------------------

// format formats a template using gofmt.
func format(in io.Reader) (io.Reader, error) {
	var out bytes.Buffer

	gofmt := exec.Command("gofmt", "-s")
	gofmt.Stdin = in
	gofmt.Stdout = &out
	gofmt.Stderr = os.Stderr
	err := gofmt.Run()
	return &out, err
}

//-------------------------------------------------------------------------------------------------

type context struct {
	pkgStore parse.PackageStore
	indices  map[string]*schema.Index
	table    *schema.TableDescription
	mainPkg  string
	fileTags stypes.Tags
}

func load(pkgStore parse.PackageStore, name parse.LType, mainPkg string, fileTags stypes.Tags) (*schema.TableDescription, error) {
	table := new(schema.TableDescription)

	// local map of indexes, used for quick lookups and de-duping.
	indices := map[string]*schema.Index{}

	table.Type = name.Name
	table.Name = inflections.Pluralize(inflections.Underscore(table.Type))

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

func mergeTags(structTags, fileTags stypes.Tags) stypes.Tags {
	merged := make(stypes.Tags)
	for n, t := range structTags {
		merged[n] = t
	}
	for n, t := range fileTags {
		merged[n] = t
	}
	parse.DevInfo("merged tags\n-----------\n%s\n", merged.String())
	return merged
}

func checkNoConflictingNames(name parse.LType, table *schema.TableDescription) {
	names := make(map[string]struct{})
	var duplicates schema.Identifiers

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
	tag := tags[leaf.Name()]
	field := &schema.Field{}
	field.Tags = tag
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

	prefix := ""
	if tag.Prefixed {
		prefix = inflections.Underscore(field.JoinParts(1, "_")) + "_"
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
		tags[str.Field(i).Name()] = *tag
	}
}

//-------------------------------------------------------------------------------------------------

var mapTagToEncoding = map[string]schema.SqlEncode{
	"":       schema.ENCNONE,
	"json":   schema.ENCJSON,
	"text":   schema.ENCTEXT,
	"driver": schema.ENCDRIVER,
}

var mapStringToSqlType = map[string]stypes.Kind{
	// Go-flavour names
	"bool":    stypes.Bool,
	"int":     stypes.Int,
	"int8":    stypes.Int8,
	"int16":   stypes.Int16,
	"int32":   stypes.Int32,
	"int64":   stypes.Int64,
	"uint":    stypes.Uint,
	"uint8":   stypes.Uint8,
	"uint16":  stypes.Uint16,
	"uint32":  stypes.Uint32,
	"uint64":  stypes.Uint64,
	"float32": stypes.Float32,
	"float64": stypes.Float64,
	"string":  stypes.String,

	// SQL-flavour names
	"text":     stypes.String,
	"json":     stypes.String,
	"varchar":  stypes.String,
	"varchar2": stypes.String,
	"number":   stypes.Int,
	"tinyint":  stypes.Int8,
	"smallint": stypes.Int16,
	"integer":  stypes.Int,
	"bigint":   stypes.Int64,
	"blob":     stypes.Struct,
	"bytea":    stypes.Struct,
}

func allowedSqlTypeStrings() string {
	var s []string
	for k := range mapStringToSqlType {
		s = append(s, k)
	}
	sort.Strings(s)
	return strings.Join(s, ", ")
}

