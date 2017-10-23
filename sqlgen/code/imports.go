package code

import (
	"fmt"
	"io"
	"sort"
	"github.com/rickb777/sqlgen2/schema"
	"github.com/rickb777/sqlgen2/sqlgen/parse"
)

const tabs = "\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t"

func WriteImports(w io.Writer, table *schema.Table, pkgs ...string) {
	var pmap = map[string]struct{}{}

	// add default packages
	for _, pkg := range pkgs {
		pmap[pkg] = struct{}{}
	}

	// check each edge field to see if it is
	// encoded, which might require us to import
	// other packages
	for _, field := range table.Fields {
		if field.Type.Pkg != "" {
			longName := parse.FindImport(field.Type)
			pmap[longName] = struct{}{}
		}

		switch field.Encode {
		case schema.ENCJSON:
			pmap["encoding/json"] = struct{}{}
			// case "gzip":
			// 	pmap["compress/gzip"] = struct{}{}
			// case "snappy":
			// 	pmap["github.com/golang/snappy"] = struct{}{}
		}
	}

	if len(pmap) > 0 {
		doWriteImports(w, pmap)
	}
}

func sortImports(pmap map[string]struct{}) []string {
	sorted := make([]string, 0, len(pmap))
	for pkg := range pmap {
		if pkg != "" {
			sorted = append(sorted, pkg)
		}
	}
	sort.Strings(sorted)
	return sorted
}

func doWriteImports(w io.Writer, pmap map[string]struct{}) {
	// write the import block, including each
	// encoder package that was specified.
	fmt.Fprintln(w, "\nimport (")
	imports := sortImports(pmap)
	for _, pkg := range imports {
		fmt.Fprintf(w, "\t%q\n", pkg)
	}
	fmt.Fprintln(w, ")")
}
