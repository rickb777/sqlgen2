package code

import (
	"fmt"
	"io"
	"github.com/rickb777/sqlgen2/sqlgen/parse"
)

func WriteSliceFunc(w io.Writer, tree *parse.Node, view View, withoutPk bool) {
	var depth int
	var parent = tree

	for i, node := range tree.Leaves() {
		if node.Tags.Skip || (withoutPk && node.Tags.Primary) {
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
		l3 := fmt.Sprintf("\t\tv%d,\n", i)
		view.Body3 = append(view.Body3, l3)

		// variable setting
		path := node.Path()[1:]

		// if the parent is a ptr struct we
		// need to create a new
		if parent != node.Parent && node.Parent.Type.Base == parse.Ptr {
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

		switch node.Type.Base {
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
