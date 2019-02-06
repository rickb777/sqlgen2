package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/kortschak/utter"
	"github.com/rickb777/sqlapi/schema"
	"github.com/rickb777/sqlapi/types"
	. "github.com/rickb777/sqlgen2/code"
	. "github.com/rickb777/sqlgen2/load"
	"github.com/rickb777/sqlgen2/output"
	"github.com/rickb777/sqlgen2/parse"
	"github.com/rickb777/sqlgen2/parse/exit"
	"gopkg.in/yaml.v2"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	var oFile, typeName, prefix, list, kind, tableName, tagsFile, genSetters string
	var flags = FuncFlags{}
	var all, sselect, insert, gofmt, showVersion bool

	flag.StringVar(&oFile, "o", "", "Output file name; optional. Use '-' for stdout.\n"+
		"\tIf omitted, the first input filename is used with '_sql.go' suffix.")
	flag.StringVar(&typeName, "type", "", "The type to analyse; required.\n"+
		"\tThis is expressed in the form 'pkg.Name'")
	flag.StringVar(&prefix, "prefix", "", "Prefix for names of generated types; optional.\n"+
		"\tUse this if you need to avoid name collisions.")
	flag.StringVar(&list, "list", "", "List type for slice of model objects; optional.")
	flag.StringVar(&kind, "kind", "Table", "Kind of model: you could use 'Table', 'View', 'Join' etc as required.")
	flag.StringVar(&tableName, "table", "", "The name for the database table; default is based on the struct name as a plural.")
	flag.StringVar(&tagsFile, "tags", "", "A YAML file containing tags that augment and override any in the Go struct(s); optional.\n"+
		"\tTags control the SQL type, size, column name, indexes etc.")

	// filters for what gets generated
	flag.BoolVar(&all, "all", false, "Shorthand for '-schema -create -read -update -delete -slice'; recommended.\n"+
		"\tThis does not affect -setters.")
	flag.BoolVar(&sselect, "select", false, "Alias for -read")
	flag.BoolVar(&insert, "insert", false, "Alias for -create")
	flag.BoolVar(&flags.Schema, "schema", false, "Generate SQL schema create/drop methods.")
	flag.BoolVar(&flags.Insert, "create", false, "Generate SQL create (insert) methods.")
	flag.BoolVar(&flags.Exec, "exec", false, "Generate Exec method. This is also provided with -update or -delete.")
	flag.BoolVar(&flags.Select, "read", false, "Generate SQL select (read) methods.")
	flag.BoolVar(&flags.Update, "update", false, "Generate SQL update methods.")
	flag.BoolVar(&flags.Delete, "delete", false, "Generate SQL delete methods.")
	flag.BoolVar(&flags.Slice, "slice", false, "Generate SQL slice (column select) methods.")
	flag.StringVar(&genSetters, "setters", "none", "Generate setters for fields of your type (see -type): none, optional, exported, all.\n"+
		"\tFields that are pointers are assumed to be optional.")

	flag.BoolVar(&output.Verbose, "v", false, "Show progress messages.")
	flag.BoolVar(&parse.Debug, "z", false, "Show debug messages.")
	flag.BoolVar(&parse.PrintAST, "ast", false, "Trace the whole astract syntax tree (very verbose).")
	flag.BoolVar(&gofmt, "gofmt", false, "Format and simplify the generated code nicely.")
	flag.BoolVar(&showVersion, "version", false, "Show the version.")

	flag.Parse()

	if showVersion {
		fmt.Println(appVersion)
		os.Exit(0)
	}

	output.Require(flag.NArg() > 0, "At least one input file (or path) is required; put this after the other arguments.\n")

	if sselect {
		flags.Select = true
	}

	if insert {
		flags.Insert = true
	}

	if all {
		flags = AllFuncFlags
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
		mainPkg = LastDirName(oFile)
		parse.DevInfo("mainPkg: %s\n", mainPkg)
	}

	o := output.NewOutput(oFile)

	tags, err := types.ReadTagsFile(tagsFile)
	if err != nil && err != os.ErrNotExist {
		exit.Fail(1, "tags file %s failed: %s.\n", tagsFile, err)
	}

	// load the Tree into a schema Object
	table, err := Load(pkgStore, parse.LType{pkg, name}, mainPkg, tags)
	if err != nil {
		exit.Fail(1, "Go parser failed: %v.\n", err)
	}

	if parse.Debug {
		utter.Dump(table)
	}

	if len(table.Fields) < 1 {
		exit.Fail(1, "no fields found. Check earlier parser warnings.\n")
	}

	writeTableYaml(o.Derive(".yml"), table)
	writeSqlGo(o, name, prefix, tableName, kind, list, mainPkg, genSetters, table, flags, gofmt)

	output.Info("%s took %v\n", o.Path(), time.Now().Sub(start))
}

func writeTableYaml(o output.Output, table *schema.TableDescription) {
	buf := &bytes.Buffer{}
	enc := yaml.NewEncoder(buf)
	err := enc.Encode(table)
	if err != nil {
		exit.Fail(1, "YAML writer to %s failed: %v.\n", o.Path(), err)
	}
	o.Write(buf, os.Stdout)
}

func writeSqlGo(o output.Output, name, prefix, tableName, kind, list, mainPkg, genSetters string, table *schema.TableDescription, flags FuncFlags, gofmt bool) {
	view := NewView(name, prefix, tableName, list)
	view.Table = table
	view.Thing = kind
	view.Interface1 = PrimaryInterface(table, flags.Schema)
	view.Interface2 = SecondaryInterface(flags)

	setters := view.FilterSetters(genSetters)

	importSet := PackagesToImport(flags, view.Table.HasPrimaryKey())

	if flags.Select || flags.Insert || flags.Update {
		ImportsForFields(table, importSet)
	}
	ImportsForSetters(setters, importSet)

	buf := &bytes.Buffer{}

	WritePackageHeader(buf, mainPkg)

	WriteImports(buf, importSet)

	WriteType(buf, view)

	WritePrimaryDeclarations(buf, view)

	if flags.Schema {
		WriteSchemaDeclarations(buf, view)
		WriteSchemaFunctions(buf, view)
	}

	if flags.Exec || flags.Update || flags.Delete {
		WriteExecFunc(buf, view)
	}

	WriteQueryRows(buf, view)
	WriteQueryThings(buf, view)

	if flags.Select {
		WriteScanRows(buf, view)
		WriteGetRow(buf, view)
		WriteSelectRowsFuncs(buf, view)
	}

	if flags.Slice {
		WriteSliceColumn(buf, view)
	}

	if flags.Insert {
		WriteConstructInsert(buf, view)
	}

	if flags.Update {
		WriteConstructUpdate(buf, view)
	}

	if flags.Insert {
		WriteInsertFunc(buf, view)
	}

	if flags.Update {
		WriteUpdateFunc(buf, view)
	}

	if flags.Delete {
		WriteDeleteFunc(buf, view)
	}

	WriteSetters(buf, view, setters)

	// formats the generated file using gofmt
	var pretty io.Reader = buf
	if gofmt {
		var err error
		pretty, err = Format(buf)
		output.Require(err == nil, "%s\n%v\n", string(buf.Bytes()), err)
	}

	o.Write(pretty, os.Stdout)
}
