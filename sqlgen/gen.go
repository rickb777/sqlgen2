package main

import (
	"bytes"
	"flag"
	"github.com/kortschak/utter"
	"github.com/rickb777/sqlgen2/schema"
	. "github.com/rickb777/sqlgen2/sqlgen/code"
	. "github.com/rickb777/sqlgen2/sqlgen/output"
	"github.com/rickb777/sqlgen2/sqlgen/parse"
	"github.com/rickb777/sqlgen2/sqlgen/parse/exit"
	"io"
	"os"
	"strings"
	"path/filepath"
)

func main() {
	var oFile, typeName, prefix, list, kind, tagsFile, genSetters string
	var flags = funcFlags{}
	var all, sselect, insert, gofmt bool

	flag.StringVar(&oFile, "o", "", "Output file name; optional. Use '-' for stdout.\n" +
		"\tIf omitted, the first input filename is used with '_sql.go' suffix.")
	flag.StringVar(&typeName, "type", "", "The type to analyse; required.\n" +
		"\tThis is expressed in the form 'pkg.Name'")
	flag.StringVar(&prefix, "prefix", "", "Prefix for names of generated types; optional.\n" +
		"\tUse this if you need to avoid name collisions.")
	flag.StringVar(&list, "list", "", "List type for slice of model objects; optional.")
	flag.StringVar(&kind, "kind", "Table", "Kind of model: you could use 'Table', 'View', 'Join' etc as required")
	flag.StringVar(&tagsFile, "tags", "", "A YAML file containing tags that augment and override any in the Go struct(s); optional.\n" +
		"\tTags control the SQL type, size, column name, indexes etc.")
	flag.BoolVar(&Verbose, "v", false, "Show progress messages.")
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
	flag.BoolVar(&flags.sselect, "read", false, "Generate SQL select (read) methods.")
	flag.BoolVar(&flags.update, "update", false, "Generate SQL update methods.")
	flag.BoolVar(&flags.delete, "delete", false, "Generate SQL delete methods.")
	flag.BoolVar(&flags.slice, "slice", false, "Generate SQL slice (column select) methods.")
	flag.StringVar(&genSetters, "setters", "none", "Generate setters for fields of your type (see -type): none, optional, exported, all.\n" +
		"\tFields that are pointers are assumed to be optional.")

	flag.Parse()

	Require(flag.NArg() > 0, "At least one input file (or path) is required; put this after the other arguments.\n")

	if sselect {
		flags.sselect = true
	}

	if insert {
		flags.insert = true
	}

	if all {
		flags = allFuncFlags
	}

	Require(len(typeName) > 3, "-type is required. This must specify a type, qualified with its local package in the form 'pkg.Name'.\n", typeName)
	words := strings.Split(typeName, ".")
	Require(len(words) == 2, "type %q requires a package name prefix.\n", typeName)
	pkg, name := words[0], words[1]
	mainPkg := pkg

	// parse the Go source code file(s) to extract the required struct and return it as an AST.
	pkgStore, err := parse.Parse(flag.Args())
	Require(err == nil, "%v\n", err)
	//utter.Dump(pkgStore)

	if oFile == "" {
		oFile = flag.Args()[0]
		Require(strings.HasSuffix(oFile, ".go"), oFile+": must end '.go'")
		oFile = oFile[:len(oFile)-3] + "_sql.go"
		parse.DevInfo("oFile: %s\n", oFile)
	} else {
		mainPkg = lastDirName(oFile)
		parse.DevInfo("mainPkg: %s\n", mainPkg)
	}

	o := NewOutput(oFile)

	tags, err := parse.ReadTagsFile(tagsFile)
	if err != nil && err != os.ErrNotExist {
		exit.Fail(1, "Tags file %s failed: %s.\n", tagsFile, err)
	}

	// load the Tree into a schema Object
	table, err := load(pkgStore, parse.LType{pkg, name}, mainPkg, tags)
	if parse.Debug {
		utter.Dump(table)
	}

	view := NewView(name, prefix, list)
	view.Table = table
	view.Thing = kind
	view.Interface = primaryInterface(table, flags.schema)

	setters := view.FilterSetters(genSetters)

	buf := &bytes.Buffer{}

	WritePackage(buf, mainPkg)

	WriteImports(buf, table, setters, packagesToImport(flags, view.Table.HasPrimaryKey()))

	WriteType(buf, view)

	if flags.schema {
		WriteSchema(buf, view)
	}

	WriteExecFunc(buf, view)
	WriteQueryRows(buf, view)
	WriteRowsFunc(buf, view)
	WriteQueryThings(buf, view)

	if flags.sselect {
		WriteGetRow(buf, view)
		WriteSelectRowsFuncs(buf, view)
	}

	if flags.slice {
		WriteSliceColumn(buf, view)
	}

	if flags.insert {
		WriteInsertFunc(buf, view)
	}

	if flags.update {
		WriteUpdateFunc(buf, view)
	}

	if flags.insert || flags.update {
		WriteSliceFunc(buf, view, view.Table.HasLastInsertId())
	}

	if flags.delete {
		WriteDeleteFunc(buf, view)
	}

	WriteSetters(buf, view, setters)

	// formats the generated file using gofmt
	var pretty io.Reader = buf
	if gofmt {
		pretty, err = format(buf)
		Require(err == nil, "%s\n%v\n", string(buf.Bytes()), err)
	}

	o.Write(pretty, os.Stdout)
}

func packagesToImport(flags funcFlags, hasPrimaryKey bool) StringSet {
	imports := NewStringSet(
		"context",
		"database/sql",
		"log",
		"strings",
		"github.com/rickb777/sqlgen2",
		"github.com/rickb777/sqlgen2/require",
		"github.com/rickb777/sqlgen2/schema",
	)

	if flags.schema || flags.sselect || flags.slice || flags.insert || flags.update || flags.delete {
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

//-------------------------------------------------------------------------------------------------

type funcFlags struct {
	schema, sselect, insert, update, delete, slice bool
}

var allFuncFlags = funcFlags{true, true, true, true, true, true}
