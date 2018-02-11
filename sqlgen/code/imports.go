package code

import (
	"fmt"
	"io"
	"sort"
	"github.com/rickb777/sqlgen2/schema"
	"github.com/rickb777/sqlgen2/util"
)

func ImportsForFields(table *schema.TableDescription, packages util.StringSet) {

	// check each edge field to see if it is
	// encoded, which might require us to import
	// other packages
	for _, field := range table.Fields {

		switch field.Encode {
		case schema.ENCJSON:
			packages["encoding/json"] = struct{}{}
		case schema.ENCTEXT:
			packages["encoding"] = struct{}{}
			// case "gzip":
			// 	packages["compress/gzip"] = struct{}{}
		default:
			if field.Type.PkgPath != "" {
				packages.Add(field.Type.PkgPath)
			}
		}
	}
}

func ImportsForSetters(setters schema.FieldList, packages util.StringSet) {
	for _, field := range setters {
		if field.Type.PkgPath != "" {
			packages.Add(field.Type.PkgPath)
		}
	}
}

func WriteImports(w io.Writer, packages util.StringSet) {
	if packages.NonEmpty() {
		// write the import block, including each
		// encoder package that was specified.
		fmt.Fprintln(w, "\nimport (")
		imports := sortImports(packages)
		for _, pkg := range imports {
			fmt.Fprintf(w, "\t%q\n", pkg)
		}
		fmt.Fprintln(w, ")")
	}
}

func sortImports(pmap util.StringSet) []string {
	sorted := pmap.ToSlice()
	sort.Strings(sorted)
	return sorted
}
