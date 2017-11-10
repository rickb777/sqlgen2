package code

import (
	"fmt"
	"io"
	"github.com/rickb777/sqlgen2/sqlgen/parse"
	"github.com/rickb777/sqlgen2/schema"
)

func WriteSliceFunc(w io.Writer, view View, table *schema.Table, withoutPk bool) {
	var depth int
	//var parent = tree

	for i, field := range table.Fields {
		if field.Tags.Skip || (withoutPk && field.Tags.Primary) {
			continue
		}

		// temporary variable declaration
		switch field.Type.Base {
		case parse.Map, parse.Slice:
			l1 := fmt.Sprintf("\tvar v%d %s\n", i, "[]byte")
			view.Body1 = append(view.Body1, l1)
		default:
			l1 := fmt.Sprintf("\tvar v%d %s\n", i, field.Type.Type())
			view.Body1 = append(view.Body1, l1)
		}

		// variable scanning
		l3 := fmt.Sprintf("\t\tv%d,\n", i)
		view.Body3 = append(view.Body3, l3)

		// variable setting
		//path := field.Path()[1:]

		// if the parent is a ptr struct we need to create a new
		//if parent != field.Parent && field.Parent.Type.Base == parse.Ptr {
		//	// seriously ... this works?
		//	if field.Parent != nil && field.Parent.Parent != parent {
		//		for _, p := range path {
		//			if p == parent || depth == 0 {
		//				break
		//			}
		//			l2 := fmt.Sprintf("%s}\n", tabs[:depth])
		//			view.Body2 = append(view.Body2, l2)
		//			depth--
		//		}
		//	}
		//	l2 := fmt.Sprintf("%s\tif v.%s != nil {\n", tabs[:depth], join(path[:len(path)-1], "."))
		//	view.Body2 = append(view.Body2, l2)
		//	depth++
		//}

		switch field.Type.Base {
		case parse.Map, parse.Slice, parse.Struct, parse.Ptr:
			l2 := fmt.Sprintf("%s\tv%d, _ = json.Marshal(&v.%s)\n", tabs[:depth], i, field.Path.Join("."))
			view.Body2 = append(view.Body2, l2)
		default:
			l2 := fmt.Sprintf("%s\tv%d = v.%s\n", tabs[:depth], i, field.Path.Join("."))
			view.Body2 = append(view.Body2, l2)
		}

		//parent = field.Parent
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
