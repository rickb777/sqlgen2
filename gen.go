package main

import (
	"bytes"
	"flag"
	"io"
	"os"

	"github.com/rickb777/sqlgen/parse"
	"github.com/rickb777/sqlgen/schema"
	. "github.com/rickb777/sqlgen/output"
	"fmt"
	"strings"
)

var (
	oFile      = flag.String("o", "", "output file name (or file path); required")
	typeName   = flag.String("type", "", "type to generate; required")
	database   = flag.String("db", "sqlite", "sql dialect; optional")
	prefix     = flag.String("prefix", "", "prefix for names of generated types; optional")
	genSchema  = flag.Bool("schema", true, "generate sql schema and queries")
	genFuncs   = flag.Bool("funcs", true, "generate sql helper functions")
	extraFuncs = flag.Bool("extras", true, "generate extra sql helper functions")
	gofmt      = flag.Bool("gofmt", false, "format and simplify the generated code nicely")
)

func main() {
	flag.BoolVar(&Verbose, "v", false, "progress messages")
	flag.BoolVar(&parse.Debug, "z", false, "debug messages")

	flag.Parse()

	Require(flag.NArg() > 0, "at least one input file is required.\n")
	Require(*oFile != "", "-o is required.\n")

	words := strings.Split(*typeName, ".")
	Require(len(words) == 2, "type %q requires a package name prefix.\n", *typeName)

	// parse the Go source code file(s) to extract the required struct and return it as an AST.
	tree, err := parse.Parse(words[0], words[1], flag.Args())
	Require(err == nil, "%v\n", err)

	o := NewOutput(*oFile)

	// if the code is generated in a different folder
	// that the struct we need to set the package name and import the struct
	pkg := o.Pkg()
	imports := tree.Pkg
	if pkg == "" {
		pkg = tree.Pkg
		imports = ""
	} else {
		fmt.Fprintf(os.Stderr, "%s: sub-directories are not yet supported.\n", *oFile)
		os.Exit(1)
	}

	// load the Tree into a schema Object
	table := schema.Load(tree)
	dialect := schema.New(schema.Dialects[*database])

	buf := &bytes.Buffer{}

	writePackage(buf, pkg)

	if *genFuncs {
		view := newView(tree, *prefix)
		view.Table = table

		writeImports(buf, tree, "database/sql", "fmt", imports)
		writeType(buf, view)
		writeRowFunc(buf, tree, view)
		writeRowsFunc(buf, tree, view)
		writeSliceFunc(buf, tree, view, false)
		writeSliceFunc(buf, tree, view, true)

		if *extraFuncs {
			writeSelectRow(buf, view)
			writeSelectRows(buf, view)
			writeCountRows(buf, view)
			writeInsertFunc(buf, view, table)
			writeUpdateFunc(buf, view, table)
			writeExecFunc(buf, view, table)
		}
	}

	// write the sql functions
	if *genSchema {
		writeSchema(buf, dialect, tree, table)
	}

	// formats the generated file using gofmt
	var pretty io.Reader = buf
	if *gofmt {
		pretty, err = format(buf)
		Require(err == nil, "%s\n%v\n", string(buf.Bytes()), err)
	}

	o.Write(pretty, os.Stdout)
}
