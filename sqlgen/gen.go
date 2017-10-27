package main

import (
	"bytes"
	"flag"
	"io"
	"os"

	"fmt"
	"github.com/rickb777/sqlgen2/schema"
	. "github.com/rickb777/sqlgen2/sqlgen/code"
	. "github.com/rickb777/sqlgen2/sqlgen/output"
	"github.com/rickb777/sqlgen2/sqlgen/parse"
	"strings"
)

func main() {
	var oFile, typeName, prefix string
	var genSchema, genFuncs, gofmt bool

	flag.StringVar(&oFile, "o", "", "output file name (or file path); if omitted, the first input filename is used with _sql.go suffix")
	flag.StringVar(&typeName, "type", "", "type to generate; required")
	flag.StringVar(&prefix, "prefix", "", "prefix for names of generated types; optional")
	flag.BoolVar(&Verbose, "v", false, "progress messages")
	flag.BoolVar(&parse.Debug, "z", false, "debug messages")
	flag.BoolVar(&genSchema, "schema", true, "generate sql schema and queries")
	flag.BoolVar(&genFuncs, "funcs", true, "generate sql crud functions")
	flag.BoolVar(&gofmt, "gofmt", false, "format and simplify the generated code nicely")

	flag.Parse()

	Require(flag.NArg() > 0, "at least one input file is required; put this after the other args\n")

	words := strings.Split(typeName, ".")
	Require(len(words) == 2, "type %q requires a package name prefix.\n", typeName)

	// parse the Go source code file(s) to extract the required struct and return it as an AST.
	tree, err := parse.Parse(words[0], words[1], flag.Args())
	Require(err == nil, "%v\n", err)

	if oFile == "" {
		oFile = flag.Args()[0]
		Require(strings.HasSuffix(oFile, ".go"), oFile+": must end '.go'")
		oFile = oFile[:len(oFile)-3] + "_sql.go"
	}

	o := NewOutput(oFile)

	packagesToImport := NewStringSet("context", "database/sql", "github.com/rickb777/sqlgen2/database")

	if genFuncs || genSchema {
		packagesToImport.Add("fmt")
	}
	if genFuncs {
		packagesToImport.Add("github.com/rickb777/sqlgen2/where", "strings")
	}

	// if the code is generated in a different folder
	// that the struct we need to set the package name and import the struct
	pkg := o.Pkg()
	if pkg == "" {
		pkg = tree.Type.Pkg
	} else {
		fmt.Fprintf(os.Stderr, "%s: sub-directories are not yet supported.\n", oFile)
		os.Exit(1)
	}

	// load the Tree into a schema Object
	table := schema.Load(tree)

	buf := &bytes.Buffer{}

	WritePackage(buf, pkg)

	view := NewView(tree, prefix)
	view.Table = table
	view.Dialects = schema.Dialects

	WriteImports(buf, table, packagesToImport)

	WriteType(buf, view)
	WriteRowFunc(buf, tree, view)
	WriteRowsFunc(buf, tree, view)
	WriteSliceFunc(buf, tree, view, false)
	WriteSliceFunc(buf, tree, view, true)
	WriteExecFunc(buf, view, table)
	WriteQueryFuncs(buf, view, table)

	if genFuncs {
		WriteSelectRow(buf, view, table)
		WriteInsertFunc(buf, view, table)
		WriteUpdateFunc(buf, view, table)
	}

	if genSchema {
		WriteCreateTableFunc(buf, view, table)
		WriteCreateIndexFunc(buf, view, table)
		WriteSchema(buf, view, table)
	}

	// formats the generated file using gofmt
	var pretty io.Reader = buf
	if gofmt {
		pretty, err = format(buf)
		Require(err == nil, "%s\n%v\n", string(buf.Bytes()), err)
	}

	o.Write(pretty, os.Stdout)
}
