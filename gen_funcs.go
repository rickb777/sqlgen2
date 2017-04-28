package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/acsellers/inflections"
	"github.com/rickb777/sqlgen/parse"
	"github.com/rickb777/sqlgen/schema"
	"bitbucket.org/pkg/inflect"
)

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
	for pkg, _ := range pmap {
		if pkg != "" {
			fmt.Fprintf(w, "\t%q\n", pkg)
		}
	}
	fmt.Fprintln(w, ")")
}

func writeType(w io.Writer, tree *parse.Node) {
	name := strings.ToLower(inflect.Pluralize(tree.Type))
	fmt.Fprintf(w, sTable, tree.Type, tree.Type, name, tree.Type, tree.Type)
}

func writeSliceFunc(w io.Writer, tree *parse.Node, withoutPk bool) {

	var buf1, buf2, buf3 bytes.Buffer

	var depth int
	var parent = tree

	for i, node := range tree.Edges() {
		if node.Tags.Skip || (withoutPk && node.Tags.Primary) {
			continue
		}

		// temporary variable declaration
		switch node.Kind {
		case parse.Map, parse.Slice:
			fmt.Fprintf(&buf1, "\tvar v%d %s\n", i, "[]byte")
		default:
			fmt.Fprintf(&buf1, "\tvar v%d %s\n", i, node.Type)
		}

		// variable scanning
		fmt.Fprintf(&buf3, "\t\tv%d,\n", i)

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
					writeN(&buf2, depth, '\t')
					fmt.Fprintln(&buf2, "}")
					depth--
				}
			}
			writeN(&buf2, depth, '\t')
			fmt.Fprintf(&buf2, "\tif v.%s != nil {\n", join(path[:len(path)-1], "."))
			depth++
		}

		writeN(&buf2, depth, '\t')
		switch node.Kind {
		case parse.Map, parse.Slice, parse.Struct, parse.Ptr:
			fmt.Fprintf(&buf2, "\tv%d, _ = json.Marshal(&v.%s)\n", i, join(path, "."))
		default:
			fmt.Fprintf(&buf2, "\tv%d = v.%s\n", i, join(path, "."))
		}

		parent = node.Parent
		i++
	}

	for depth != 0 {
		writeN(&buf2, depth, '\t')
		fmt.Fprintln(&buf2, "}")
		depth--
	}

	suffix := ""
	if withoutPk {
		suffix = "WithoutPk"
	}
	fmt.Fprintf(w,
		sSliceRow,
		tree.Type,
		suffix,
		tree.Type,
		buf1.String(),
		buf2.String(),
		buf3.String(),
	)
}

func writeRowFunc(w io.Writer, tree *parse.Node) {

	var buf1, buf2, buf3 bytes.Buffer

	var i int
	var parent = tree
	for _, node := range tree.Edges() {
		if node.Tags.Skip {
			continue
		}

		// temporary variable declaration
		switch node.Kind {
		case parse.Map, parse.Slice:
			fmt.Fprintf(&buf1, "\tvar v%d %s\n", i, "[]byte")
		default:
			fmt.Fprintf(&buf1, "\tvar v%d %s\n", i, node.Type)
		}

		// variable scanning
		fmt.Fprintf(&buf2, "\t\t&v%d,\n", i)

		// variable setting
		path := node.Path()[1:]

		// if the parent is a ptr struct we
		// need to create a new
		if parent != node.Parent && node.Parent.Kind == parse.Ptr {
			fmt.Fprintf(&buf3, "\tv.%s = &%s{}\n", join(path[:len(path)-1], "."), node.Parent.Type)
		}

		switch node.Kind {
		case parse.Map, parse.Slice, parse.Struct, parse.Ptr:
			fmt.Fprintf(&buf3, "\tjson.Unmarshal(v%d, &v.%s)\n", i, join(path, "."))
		default:
			fmt.Fprintf(&buf3, "\tv.%s = v%d\n", join(path, "."), i)
		}

		parent = node.Parent
		i++
	}

	fmt.Fprintf(w,
		sScanRow,
		tree.Type,
		tree.Type,
		tree.Type,
		buf1.String(),
		buf2.String(),
		tree.Type,
		buf3.String(),
	)
}

func writeRowsFunc(w io.Writer, tree *parse.Node) {
	var buf1, buf2, buf3 bytes.Buffer

	var i int
	var parent = tree
	for _, node := range tree.Edges() {
		if node.Tags.Skip {
			continue
		}

		// temporary variable declaration
		switch node.Kind {
		case parse.Map, parse.Slice:
			fmt.Fprintf(&buf1, "\tvar v%d %s\n", i, "[]byte")
		default:
			fmt.Fprintf(&buf1, "\tvar v%d %s\n", i, node.Type)
		}

		// variable scanning
		fmt.Fprintf(&buf2, "\t\t\t&v%d,\n", i)

		// variable setting
		path := node.Path()[1:]

		// if the parent is a ptr struct we
		// need to create a new
		if parent != node.Parent && node.Parent.Kind == parse.Ptr {
			fmt.Fprintf(&buf3, "\t\tv.%s = &%s{}\n", join(path[:len(path)-1], "."), node.Parent.Type)
		}

		switch node.Kind {
		case parse.Map, parse.Slice, parse.Struct, parse.Ptr:
			fmt.Fprintf(&buf3, "\t\tjson.Unmarshal(v%d, &v.%s)\n", i, join(path, "."))
		default:
			fmt.Fprintf(&buf3, "\t\tv.%s = v%d\n", join(path, "."), i)
		}

		parent = node.Parent
		i++
	}

	plural := inflections.Pluralize(tree.Type)
	fmt.Fprintf(w,
		sScanRows,
		plural,
		plural,
		tree.Type,
		tree.Type,
		buf1.String(),
		buf2.String(),
		tree.Type,
		buf3.String(),
	)
}

func writeSelectRow(w io.Writer, tree *parse.Node) {
	fmt.Fprintf(w, sSelectRow, tree.Type, tree.Type, tree.Type)
}

func writeSelectRows(w io.Writer, tree *parse.Node) {
	plural := inflections.Pluralize(tree.Type)
	fmt.Fprintf(w, sSelectRows, tree.Type, tree.Type, plural)
}

func writeInsertFunc(w io.Writer, tree *parse.Node, table *schema.Table) {
	if table.HasLastInsertId() {
		fmt.Fprintf(w, sInsertAndGetLastId, tree.Type, tree.Type, tree.Type, tree.Type, table.Primary.GoName)
	} else {
		fmt.Fprintf(w, sInsert, tree.Type, tree.Type, tree.Type, tree.Type)
	}
}

func writeUpdateFunc(w io.Writer, tree *parse.Node, table *schema.Table) {
	if table.HasPrimaryKey() {
		fmt.Fprintf(w, sUpdate, tree.Type, tree.Type, tree.Type, tree.Type, table.Primary.GoName)
	}
}

func writeExecFunc(w io.Writer, tree *parse.Node, table *schema.Table) {
	if table.HasPrimaryKey() {
		fmt.Fprintf(w, sExec, tree.Type)
	}
}

// join is a helper function that joins nodes
// together by name using the seperator.
func join(nodes []*parse.Node, sep string) string {
	var parts []string
	for _, node := range nodes {
		parts = append(parts, node.Name)
	}
	return strings.Join(parts, sep)
}

func writeN(w io.Writer, n int, c ...byte) {
	for i := 0; i < n; i++ {
		w.Write(c)
	}
}
