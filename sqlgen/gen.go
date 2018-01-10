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
	var oFile, typeName, prefix, list, kind, tagsFile, genSetters, genFuncs string
	var flags = funcFlags{}
	var gofmt bool

	flag.StringVar(&oFile, "o", "", "output file name (or file path); if omitted, the first input filename is used with _sql.go suffix")
	flag.StringVar(&typeName, "type", "", "type to analyse; required")
	flag.StringVar(&prefix, "prefix", "", "prefix for names of generated types; optional")
	flag.StringVar(&list, "list", "", "list type for slice of model objects; optional")
	flag.StringVar(&kind, "kind", "Table", "kind of model: default is Table but you could use View, Join etc as required")
	flag.StringVar(&tagsFile, "tags", "", "a YAML file containing tags that augment and override any in the Go struct(s); optional")
	flag.BoolVar(&Verbose, "v", false, "progress messages")
	flag.BoolVar(&parse.Debug, "z", false, "debug messages")
	flag.BoolVar(&parse.PrintAST, "ast", false, "trace the whole astract syntax tree (very verbose)")
	flag.BoolVar(&flags.schema, "schema", true, "generate sql schema and queries")
	flag.BoolVar(&flags.insert, "create", true, "generate sql create (insert) functions")
	flag.BoolVar(&flags.sselect, "read", true, "generate sql select functions")
	flag.BoolVar(&flags.update, "update", true, "generate sql update functions")
	flag.BoolVar(&flags.delete, "delete", true, "generate sql delete functions")
	flag.BoolVar(&flags.slice, "slice", true, "generate sql slice (column select) functions")
	flag.StringVar(&genFuncs, "funcs", "", "shorthand for generate crud functions: none, all")
	flag.StringVar(&genSetters, "setters", "none", "generate setters for fields: none, optional, exported, all")
	flag.BoolVar(&gofmt, "gofmt", false, "format and simplify the generated code nicely")

	flag.Parse()

	Require(flag.NArg() > 0, "at least one input file is required; put this after the other args\n")

	switch genFuncs {
	case "all":
		flags = allFuncFlags
	case "none":
		flags = noFuncFlags
	case "":
	default:
		exit.Fail(1, "-funcs: value must be 'all' or 'none'.\n")
	}

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
	WriteQueryFuncs(buf, view)

	if flags.sselect {
		WriteGetRow(buf, view)
		WriteSelectRow(buf, view)
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

	WriteRowsFunc(buf, view)
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
		"github.com/rickb777/sqlgen2",
		"github.com/rickb777/sqlgen2/schema",
	)

	if flags.schema || flags.sselect || flags.slice || flags.insert || flags.update || flags.delete {
		imports.Add("fmt")
	}
	if flags.sselect || flags.slice || flags.update || flags.delete {
		imports.Add("github.com/rickb777/sqlgen2/where")
	}
	if flags.update {
		imports.Add("strings")
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

type funcFlags struct {
	schema, sselect, insert, update, delete, slice bool
}

var allFuncFlags = funcFlags{true, true, true, true, true, true}
var noFuncFlags = funcFlags{false, false, false, false, false, false}
