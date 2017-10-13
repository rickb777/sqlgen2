package main

import (
	"bytes"
	"flag"
	"io"
	"os"

	"fmt"
	. "github.com/rickb777/sqlgen/output"
	"github.com/rickb777/sqlgen/parse"
	"github.com/rickb777/sqlgen/schema"
	"strings"
)

func main() {
	var oFile, typeName, database, prefix string
	var genSchema, genFuncs, extraFuncs, gofmt bool

	flag.StringVar(&oFile, "o", "", "output file name (or file path); if omitted, the first input filename is used with _sql.go suffix")
	flag.StringVar(&typeName, "type", "", "type to generate; required")
	flag.StringVar(&database, "db", "sqlite", "sql dialect; optional")
	flag.StringVar(&prefix, "prefix", "", "prefix for names of generated types; optional")
	flag.BoolVar(&Verbose, "v", false, "progress messages")
	flag.BoolVar(&parse.Debug, "z", false, "debug messages")
	flag.BoolVar(&genSchema, "schema", true, "generate sql schema and queries")
	flag.BoolVar(&genFuncs, "funcs", true, "generate sql helper functions")
	flag.BoolVar(&extraFuncs, "extras", true, "generate extra sql helper functions")
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

	// if the code is generated in a different folder
	// that the struct we need to set the package name and import the struct
	pkg := o.Pkg()
	imports := tree.Pkg
	if pkg == "" {
		pkg = tree.Pkg
		imports = ""
	} else {
		fmt.Fprintf(os.Stderr, "%s: sub-directories are not yet supported.\n", oFile)
		os.Exit(1)
	}

	// load the Tree into a schema Object
	table := schema.Load(tree)
	dialect := schema.New(schema.Dialects[database])

	buf := &bytes.Buffer{}

	writePackage(buf, pkg)

	if genFuncs {
		view := newView(tree, prefix)
		view.Table = table

		writeImports(buf, tree, "database/sql", "fmt", imports)
		writeType(buf, view)
		writeRowFunc(buf, tree, view)
		writeRowsFunc(buf, tree, view)
		writeSliceFunc(buf, tree, view, false)
		writeSliceFunc(buf, tree, view, true)

		if extraFuncs {
			writeSelectRow(buf, view)
			writeSelectRows(buf, view)
			writeCountRows(buf, view)
			writeInsertFunc(buf, view, table)
			writeUpdateFunc(buf, view, table)
			writeExecFunc(buf, view, table)
		}
	}

	// write the sql functions
	if genSchema {
		writeSchema(buf, dialect, tree, table)
	}

	// formats the generated file using gofmt
	var pretty io.Reader = buf
	if gofmt {
		pretty, err = format(buf)
		Require(err == nil, "%s\n%v\n", string(buf.Bytes()), err)
	}

	o.Write(pretty, os.Stdout)
}
