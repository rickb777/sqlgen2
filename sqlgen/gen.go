package main

import (
	"bytes"
	"flag"
	"github.com/kortschak/utter"
	. "github.com/rickb777/sqlgen2/sqlgen/code"
	. "github.com/rickb777/sqlgen2/sqlgen/output"
	"github.com/rickb777/sqlgen2/sqlgen/parse"
	"io"
	"os"
	"strings"
)

func main() {
	var oFile, typeName, prefix, list string
	var genSchema, genFuncs, gofmt bool

	flag.StringVar(&oFile, "o", "", "output file name (or file path); if omitted, the first input filename is used with _sql.go suffix")
	flag.StringVar(&typeName, "type", "", "type to generate; required")
	flag.StringVar(&prefix, "prefix", "", "prefix for names of generated types; optional")
	flag.StringVar(&list, "list", "", "list type for slice of model objects; optional")
	flag.BoolVar(&Verbose, "v", false, "progress messages")
	flag.BoolVar(&parse.Debug, "z", false, "debug messages")
	flag.BoolVar(&parse.PrintAST, "ast", false, "trace the whole astract syntax tree (very verbose)")
	flag.BoolVar(&genSchema, "schema", true, "generate sql schema and queries")
	flag.BoolVar(&genFuncs, "funcs", true, "generate sql crud functions")
	flag.BoolVar(&gofmt, "gofmt", false, "format and simplify the generated code nicely")

	flag.Parse()

	Require(flag.NArg() > 0, "at least one input file is required; put this after the other args\n")

	words := strings.Split(typeName, ".")
	Require(len(words) == 2, "type %q requires a package name prefix.\n", typeName)
	pkg, name := words[0], words[1]

	// parse the Go source code file(s) to extract the required struct and return it as an AST.
	pkgStore, err := parse.Parse(flag.Args())
	Require(err == nil, "%v\n", err)
	//utter.Dump(pkgStore)

	if oFile == "" {
		oFile = flag.Args()[0]
		Require(strings.HasSuffix(oFile, ".go"), oFile+": must end '.go'")
		oFile = oFile[:len(oFile)-3] + "_sql.go"
	}

	o := NewOutput(oFile)

	// load the Tree into a schema Object
	table, err := load(pkgStore, pkg, name)
	if parse.Debug {
		utter.Dump(table)
	}

	view := NewView(name, prefix, list)
	view.Table = table

	buf := &bytes.Buffer{}

	WritePackage(buf, pkg)

	WriteImports(buf, table, packagesToImport(genFuncs, genSchema, view.Table.HasPrimaryKey()))

	WriteType(buf, view)

	if genSchema {
		WriteSchema(buf, view)
	}

	WriteExecFunc(buf, view)
	WriteQueryFuncs(buf, view)

	if genFuncs {
		WriteGetRow(buf, view)
		WriteSelectRow(buf, view)
		WriteInsertFunc(buf, view)
		WriteUpdateFunc(buf, view)
		WriteDeleteFunc(buf, view)
	}

	WriteRowFunc(buf, view)
	WriteRowsFunc(buf, view)
	WriteSliceFunc(buf, view, view.Table.HasLastInsertId())

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
	if genFuncs && hasPrimaryKey {
		imports.Add("strings")
	}
	return imports
}
