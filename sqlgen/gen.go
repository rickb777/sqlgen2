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
	var genSchema, genFuncs, gofmt bool

	flag.StringVar(&oFile, "o", "", "output file name (or file path); if omitted, the first input filename is used with _sql.go suffix")
	flag.StringVar(&typeName, "type", "", "type to analyse; required")
	flag.StringVar(&prefix, "prefix", "", "prefix for names of generated types; optional")
	flag.StringVar(&list, "list", "", "list type for slice of model objects; optional")
	flag.StringVar(&kind, "kind", "Table", "kind of model: default is Table but you could use View, Join etc as required")
	flag.StringVar(&tagsFile, "tags", "", "a YAML file containing tags that augment and override any in the Go struct(s); optional")
	flag.BoolVar(&Verbose, "v", false, "progress messages")
	flag.BoolVar(&parse.Debug, "z", false, "debug messages")
	flag.BoolVar(&parse.PrintAST, "ast", false, "trace the whole astract syntax tree (very verbose)")
	flag.BoolVar(&genSchema, "schema", true, "generate sql schema and queries")
	flag.BoolVar(&genFuncs, "funcs", true, "generate sql crud functions")
	flag.StringVar(&genSetters, "setters", "none", "generate setters for fields: none, optional, exported, all")
	flag.BoolVar(&gofmt, "gofmt", false, "format and simplify the generated code nicely")

	flag.Parse()

	Require(flag.NArg() > 0, "at least one input file is required; put this after the other args\n")

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
	if err != nil && err != os.ErrNotExist{
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
	view.Interface = primaryInterface(table, genSchema)

	setters := view.FilterSetters(genSetters)

	buf := &bytes.Buffer{}

	WritePackage(buf, mainPkg)

	WriteImports(buf, table, setters, packagesToImport(genFuncs, genSchema, view.Table.HasPrimaryKey()))

	WriteType(buf, view)

	if genSchema {
		WriteSchema(buf, view)
	}

	WriteExecFunc(buf, view)
	WriteQueryFuncs(buf, view)

	if genFuncs {
		WriteGetRow(buf, view)
		WriteSelectItem(buf, view)
		WriteSelectRow(buf, view)
		WriteInsertFunc(buf, view)
		WriteUpdateFunc(buf, view)
		WriteDeleteFunc(buf, view)
	}

	WriteRowsFunc(buf, view)
	WriteSliceFunc(buf, view, view.Table.HasLastInsertId())
	WriteSetters(buf, view, setters)

	// formats the generated file using gofmt
	var pretty io.Reader = buf
	if gofmt {
		pretty, err = format(buf)
		Require(err == nil, "%s\n%v\n", string(buf.Bytes()), err)
	}

	o.Write(pretty, os.Stdout)
}

func packagesToImport(genFuncs, genSchema, hasPrimaryKey bool) StringSet {
	imports := NewStringSet(
		"context",
		"database/sql",
		"log",
		"github.com/rickb777/sqlgen2",
		"github.com/rickb777/sqlgen2/schema",
	)

	if genFuncs || genSchema {
		imports.Add("fmt")
	}
	if genFuncs {
		imports.Add("github.com/rickb777/sqlgen2/where")
	}
	if genFuncs {
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
