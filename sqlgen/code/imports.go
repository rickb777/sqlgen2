package code

import (
	"fmt"
	"io"
	"sort"
	"github.com/rickb777/sqlgen2/schema"
	"github.com/rickb777/sqlgen2/sqlgen/parse"
)

const tabs = "\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t"

func WriteImports(w io.Writer, table *schema.Table, packages StringSet) {

	// check each edge field to see if it is
	// encoded, which might require us to import
	// other packages
	for _, field := range table.Fields {
		if field.Type.Pkg != "" {
			longName := parse.FindImport(field.Type)
			packages[longName] = struct{}{}
		}

		switch field.Encode {
		case schema.ENCJSON:
			packages["encoding/json"] = struct{}{}
			// case "gzip":
			// 	packages["compress/gzip"] = struct{}{}
			// case "snappy":
			// 	packages["github.com/golang/snappy"] = struct{}{}
		}
	}

	if packages.NonEmpty() {
		doWriteImports(w, packages)
	}
}

func sortImports(pmap StringSet) []string {
	sorted := pmap.ToSlice()
	sort.Strings(sorted)
	return sorted
}

func doWriteImports(w io.Writer, pmap StringSet) {
	// write the import block, including each
	// encoder package that was specified.
	fmt.Fprintln(w, "\nimport (")
	imports := sortImports(pmap)
	for _, pkg := range imports {
		fmt.Fprintf(w, "\t%q\n", pkg)
	}
	fmt.Fprintln(w, ")")
}
