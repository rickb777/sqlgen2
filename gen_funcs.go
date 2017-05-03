package main

import (
	"fmt"
	"io"
	"sort"
	. "strings"

	//. "github.com/acsellers/inflections"
	"github.com/rickb777/sqlgen/parse"
	"github.com/rickb777/sqlgen/schema"
	"github.com/rickb777/sqlgen/output"
)

const tabs = "\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t"

func writeImports(w io.Writer, tree *parse.Node, pkgs ...string) {
	var pmap = map[string]struct{}{}

	// add default packages
	for _, pkg := range pkgs {
		pmap[pkg] = struct{}{}
	}

	// check each edge node to see if it is
	// encoded, which might require us to import
	// other packages
	for _, node := range tree.Edges() {
		if node.Tags == nil || len(node.Tags.Encode) == 0 {
			continue
		}
		switch node.Tags.Encode {
		case "json":
			pmap["encoding/json"] = struct{}{}
			// case "gzip":
			// 	pmap["compress/gzip"] = struct{}{}
			// case "snappy":
			// 	pmap["github.com/golang/snappy"] = struct{}{}
		}
	}

	if len(pmap) == 0 {
		return
	}

	// write the import block, including each
	// encoder package that was specified.
	fmt.Fprintln(w, "\nimport (")
	sorted := make([]string, 0, len(pmap))
	for pkg, _ := range pmap {
		if pkg != "" {
			sorted = append(sorted, pkg)
		}
	}
	sort.Strings(sorted)
	for _, pkg := range sorted {
		fmt.Fprintf(w, "\t%q\n", pkg)
	}
	fmt.Fprintln(w, ")")
}

func writeType(w io.Writer, view View) {
	must(tTable.Execute(w, view))
}

func writeSliceFunc(w io.Writer, tree *parse.Node, view View, withoutPk bool) {
	var depth int
	var parent = tree

	for i, node := range tree.Edges() {
		if node.Tags.Skip || (withoutPk && node.Tags.Primary) {
			continue
		}

		// temporary variable declaration
		switch node.Kind {
		case parse.Map, parse.Slice:
			l1 := fmt.Sprintf("\tvar v%d %s\n", i, "[]byte")
			view.Body1 = append(view.Body1, l1)
		default:
			l1 := fmt.Sprintf("\tvar v%d %s\n", i, node.Type)
			view.Body1 = append(view.Body1, l1)
		}

		// variable scanning
		l3 := fmt.Sprintf("\t\tv%d,\n", i)
		view.Body3 = append(view.Body3, l3)

		// variable setting
		path := node.Path()[1:]

		// if the parent is a ptr struct we
		// need to create a new
		if parent != node.Parent && node.Parent.Kind == parse.Ptr {
			// seriously ... this works?
			if node.Parent != nil && node.Parent.Parent != parent {
				for _, p := range path {
					if p == parent || depth == 0 {
						break
					}
					l2 := fmt.Sprintf("%s}\n", tabs[:depth])
					view.Body2 = append(view.Body2, l2)
					depth--
				}
			}
			l2 := fmt.Sprintf("%s\tif v.%s != nil {\n", tabs[:depth], join(path[:len(path)-1], "."))
			view.Body2 = append(view.Body2, l2)
			depth++
		}

		switch node.Kind {
		case parse.Map, parse.Slice, parse.Struct, parse.Ptr:
			l2 := fmt.Sprintf("%s\tv%d, _ = json.Marshal(&v.%s)\n", tabs[:depth], i, join(path, "."))
			view.Body2 = append(view.Body2, l2)
		default:
			l2 := fmt.Sprintf("%s\tv%d = v.%s\n", tabs[:depth], i, join(path, "."))
			view.Body2 = append(view.Body2, l2)
		}

		parent = node.Parent
		i++
	}

	for depth != 0 {
		l2 := fmt.Sprintf("%s}\n", tabs[:depth])
		view.Body2 = append(view.Body2, l2)
		depth--
	}

	if withoutPk {
		view.Suffix = "WithoutPk"
	}
	must(tSliceRow.Execute(w, view))
}

func writeRowFunc(w io.Writer, tree *parse.Node, view View) {
	var i int
	var parent = tree
	for _, node := range tree.Edges() {
		if node.Tags.Skip {
			continue
		}

		// temporary variable declaration
		switch node.Kind {
		case parse.Map, parse.Slice:
			l1 := fmt.Sprintf("\tvar v%d %s\n", i, "[]byte")
			view.Body1 = append(view.Body1, l1)
		default:
			l1 := fmt.Sprintf("\tvar v%d %s\n", i, node.Type)
			view.Body1 = append(view.Body1, l1)
		}

		// variable scanning
		l2 := fmt.Sprintf("\t\t&v%d,\n", i)
		view.Body2 = append(view.Body2, l2)

		// variable setting
		path := node.Path()[1:]

		// if the parent is a ptr struct we
		// need to create a new
		if parent != node.Parent && node.Parent.Kind == parse.Ptr {
			l3 := fmt.Sprintf("\tv.%s = &%s{}\n", join(path[:len(path)-1], "."), node.Parent.Type)
			view.Body3 = append(view.Body3, l3)
		}

		switch node.Kind {
		case parse.Map, parse.Slice, parse.Struct, parse.Ptr:
			l3 := fmt.Sprintf("\tjson.Unmarshal(v%d, &v.%s)\n", i, join(path, "."))
			view.Body3 = append(view.Body3, l3)
		default:
			l3 := fmt.Sprintf("\tv.%s = v%d\n", join(path, "."), i)
			view.Body3 = append(view.Body3, l3)
		}

		parent = node.Parent
		i++
	}

	must(tScanRow.Execute(w, view))
}

func writeRowsFunc(w io.Writer, tree *parse.Node, view View) {
	var i int
	var parent = tree
	for _, node := range tree.Edges() {
		if node.Tags.Skip {
			continue
		}

		// temporary variable declaration
		switch node.Kind {
		case parse.Map, parse.Slice:
			l1 := fmt.Sprintf("\tvar v%d %s\n", i, "[]byte")
			view.Body1 = append(view.Body1, l1)
		default:
			l1 := fmt.Sprintf("\tvar v%d %s\n", i, node.Type)
			view.Body1 = append(view.Body1, l1)
		}

		// variable scanning
		l2 := fmt.Sprintf("\t\t\t&v%d,\n", i)
		view.Body2 = append(view.Body2, l2)

		// variable setting
		path := node.Path()[1:]

		// if the parent is a ptr struct we
		// need to create a new
		if parent != node.Parent && node.Parent.Kind == parse.Ptr {
			l3 := fmt.Sprintf("\t\tv.%s = &%s{}\n", join(path[:len(path)-1], "."), node.Parent.Type)
			view.Body3 = append(view.Body3, l3)
		}

		switch node.Kind {
		case parse.Map, parse.Slice, parse.Struct, parse.Ptr:
			l3 := fmt.Sprintf("\t\tjson.Unmarshal(v%d, &v.%s)\n", i, join(path, "."))
			view.Body3 = append(view.Body3, l3)
		default:
			l3 := fmt.Sprintf("\t\tv.%s = v%d\n", join(path, "."), i)
			view.Body3 = append(view.Body3, l3)
		}

		parent = node.Parent
		i++
	}

	must(tScanRows.Execute(w, view))
}

func writeSelectRow(w io.Writer, view View) {
	must(tSelectRow.Execute(w, view))
}

func writeSelectRows(w io.Writer, view View) {
	must(tSelectRows.Execute(w, view))
}

func writeCountRows(w io.Writer, view View) {
	must(tCountRows.Execute(w, view))
}

func writeInsertFunc(w io.Writer, view View, table *schema.Table) {
	if table.HasLastInsertId() {
		must(tInsertAndGetLastId.Execute(w, view))
	} else {
		must(tInsert.Execute(w, view))
	}
}

func writeUpdateFunc(w io.Writer, view View, table *schema.Table) {
	if table.HasPrimaryKey() {
		must(tUpdate.Execute(w, view))
	}
}

func writeExecFunc(w io.Writer, view View, table *schema.Table) {
	if table.HasPrimaryKey() {
		must(tExec.Execute(w, view))
	}
}

// join is a helper function that joins nodes
// together by name using the seperator.
func join(nodes []*parse.Node, sep string) string {
	var parts []string
	for _, node := range nodes {
		parts = append(parts, node.Name)
	}
	return Join(parts, sep)
}

func writeN(w io.Writer, n int, c ...byte) {
	for i := 0; i < n; i++ {
		w.Write(c)
	}
}

func must(err error) {
	output.Require(err == nil, "%v\n", err)
}
