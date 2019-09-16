package code

import (
	"fmt"
	"github.com/rickb777/collection"
	"github.com/rickb777/sqlapi/schema"
	"io"
	"sort"
)

func ImportsForFields(table *schema.TableDescription, packages collection.StringSet) {

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

func ImportsForSetters(setters schema.FieldList, packages collection.StringSet) {
	for _, field := range setters {
		if field.Type.PkgPath != "" {
			packages.Add(field.Type.PkgPath)
		}
	}
}

func WriteImports(w io.Writer, packages collection.StringSet) {
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

func sortImports(pmap collection.StringSet) []string {
	sorted := pmap.ToSlice()
	sort.Strings(sorted)
	return sorted
}
