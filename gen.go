package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/rickb777/sqlgen/parse"
	"github.com/rickb777/sqlgen/schema"
)

var (
	input      = flag.String("file", "", "input file name; required")
	output     = flag.String("o", "", "output file name; required")
	typeName   = flag.String("type", "", "type to generate; required")
	database   = flag.String("db", "sqlite", "sql dialect; required")
	genSchema  = flag.Bool("schema", true, "generate sql schema and queries")
	genFuncs   = flag.Bool("funcs", true, "generate sql helper functions")
	extraFuncs = flag.Bool("extras", true, "generate extra sql helper functions")
	gofmt      = flag.Bool("gofmt", false, "format and simplify the generated code nicely")
)

func main() {
	flag.Parse()

	// parses the syntax tree into something a bit
	// easier to work with.
	tree, err := parse.Parse(*input, *typeName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	// load the Tree into a schema Object
	table := schema.Load(tree)
	dialect := schema.New(schema.Dialects[*database])

	buf := &bytes.Buffer{}

	writePackage(buf, tree.Pkg)

	if *genFuncs {
		writeImports(buf, tree, "database/sql", "fmt")
		writeType(buf, tree)
		writeRowFunc(buf, tree)
		writeRowsFunc(buf, tree)
		writeSliceFunc(buf, tree, false)
		writeSliceFunc(buf, tree, true)

		if *extraFuncs {
			writeSelectRow(buf, tree)
			writeSelectRows(buf, tree)
			writeInsertFunc(buf, tree, table)
			writeUpdateFunc(buf, tree, table)
			writeExecFunc(buf, tree, table)
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
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n%v\n", string(buf.Bytes()), err)
			return
		}
	}

	// create output source for file. defaults to
	// stdout but may be file.
	var out io.WriteCloser = os.Stdout
	if *output != "" {
		out, err = os.Create(*output)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return
		}
		defer out.Close()
	}

	io.Copy(out, pretty)
}
