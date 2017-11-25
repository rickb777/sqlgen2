package code

import (
	"fmt"
	"io"

	"github.com/rickb777/sqlgen2/sqlgen/parse"
	"github.com/rickb777/sqlgen2/schema"
)

func WriteRowFunc(w io.Writer, view View, table *schema.Table) {
	//var parent = tree

	for i, field := range table.Fields {
		if field.Tags.Skip {
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
		l2 := fmt.Sprintf("\t\t&v%d,\n", i)
		view.Body2 = append(view.Body2, l2)

		// variable setting
		//path := field.Path()[1:]
		//
		//// if the parent is a ptr struct we
		//// need to create a new
		//if parent != field.Parent && field.Parent.Type.Base == parse.Ptr {
		//	l3 := fmt.Sprintf("\tv.%s = &%s{}\n", join(path[:len(path)-1], "."), field.Parent.Type.Type())
		//	view.Body3 = append(view.Body3, l3)
		//}

		switch field.Type.Base {
		case parse.Map, parse.Slice, parse.Struct, parse.Ptr:
			l3 := fmt.Sprintf("\tjson.Unmarshal(v%d, &v.%s)\n", i, field.JoinParts("."))
			view.Body3 = append(view.Body3, l3)
		default:
			l3 := fmt.Sprintf("\tv.%s = v%d\n", field.JoinParts("."), i)
			view.Body3 = append(view.Body3, l3)
		}

		//parent = field.Parent
	}

	must(tScanRow.Execute(w, view))
}

func WriteRowsFunc(w io.Writer, view View, table *schema.Table) {
	//var parent = tree

	for i, field := range table.Fields {
		if field.Tags.Skip {
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
		l2 := fmt.Sprintf("\t\t\t&v%d,\n", i)
		view.Body2 = append(view.Body2, l2)

		// variable setting
		//path := field.Path()[1:]
		//
		//// if the parent is a ptr struct we
		//// need to create a new
		//if parent != field.Parent && field.Parent.Type.Base == parse.Ptr {
		//	l3 := fmt.Sprintf("\t\tv.%s = &%s{}\n", join(path[:len(path)-1], "."), field.Parent.Type.Type())
		//	view.Body3 = append(view.Body3, l3)
		//}

		switch field.Type.Base {
		case parse.Map, parse.Slice, parse.Struct, parse.Ptr:
			l3 := fmt.Sprintf("\t\tjson.Unmarshal(v%d, &v.%s)\n", i, field.JoinParts("."))
			view.Body3 = append(view.Body3, l3)
		default:
			l3 := fmt.Sprintf("\t\tv.%s = v%d\n", field.JoinParts("."), i)
			view.Body3 = append(view.Body3, l3)
		}

		//parent = field.Parent
	}

	must(tScanRows.Execute(w, view))
}

// join is a helper function that joins nodes
// together by name using the seperator.
//func join(nodes []*parse.Node, sep string) string {
//	var parts []string
//	for _, node := range nodes {
//		parts = append(parts, node.Name)
//	}
//	return Join(parts, sep)
//}
