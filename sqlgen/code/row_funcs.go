package code

import (
	"fmt"
	"io"
	. "strings"

	"github.com/rickb777/sqlgen2/sqlgen/parse"
)

func WriteRowFunc(w io.Writer, tree *parse.Node, view View) {
	var i int
	var parent = tree
	for _, node := range tree.Leaves() {
		if node.Tags.Skip {
			continue
		}

		// temporary variable declaration
		switch node.Type.Base {
		case parse.Map, parse.Slice:
			l1 := fmt.Sprintf("\tvar v%d %s\n", i, "[]byte")
			view.Body1 = append(view.Body1, l1)
		default:
			l1 := fmt.Sprintf("\tvar v%d %s\n", i, node.Type.Type())
			view.Body1 = append(view.Body1, l1)
		}

		// variable scanning
		l2 := fmt.Sprintf("\t\t&v%d,\n", i)
		view.Body2 = append(view.Body2, l2)

		// variable setting
		path := node.Path()[1:]

		// if the parent is a ptr struct we
		// need to create a new
		if parent != node.Parent && node.Parent.Type.Base == parse.Ptr {
			l3 := fmt.Sprintf("\tv.%s = &%s{}\n", join(path[:len(path)-1], "."), node.Parent.Type.Type())
			view.Body3 = append(view.Body3, l3)
		}

		switch node.Type.Base {
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

func WriteRowsFunc(w io.Writer, tree *parse.Node, view View) {
	var i int
	var parent = tree
	for _, node := range tree.Leaves() {
		if node.Tags.Skip {
			continue
		}

		// temporary variable declaration
		switch node.Type.Base {
		case parse.Map, parse.Slice:
			l1 := fmt.Sprintf("\tvar v%d %s\n", i, "[]byte")
			view.Body1 = append(view.Body1, l1)
		default:
			l1 := fmt.Sprintf("\tvar v%d %s\n", i, node.Type.Type())
			view.Body1 = append(view.Body1, l1)
		}

		// variable scanning
		l2 := fmt.Sprintf("\t\t\t&v%d,\n", i)
		view.Body2 = append(view.Body2, l2)

		// variable setting
		path := node.Path()[1:]

		// if the parent is a ptr struct we
		// need to create a new
		if parent != node.Parent && node.Parent.Type.Base == parse.Ptr {
			l3 := fmt.Sprintf("\t\tv.%s = &%s{}\n", join(path[:len(path)-1], "."), node.Parent.Type.Type())
			view.Body3 = append(view.Body3, l3)
		}

		switch node.Type.Base {
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

// join is a helper function that joins nodes
// together by name using the seperator.
func join(nodes []*parse.Node, sep string) string {
	var parts []string
	for _, node := range nodes {
		parts = append(parts, node.Name)
	}
	return Join(parts, sep)
}
