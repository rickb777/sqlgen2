package code

import (
	"fmt"
	"io"

	"github.com/rickb777/sqlgen2/sqlgen/parse"
	"github.com/rickb777/sqlgen2/schema"
)

func WriteRowFunc(w io.Writer, view View, table *schema.Table) {

	for i, field := range table.Fields {
		if field.Tags.Skip {
			continue
		}

		// temporary variable declaration
		switch field.Type.Base {
		case parse.Map, parse.Slice, parse.Struct, parse.Ptr:
			l1 := fmt.Sprintf("\tvar v%d %s\n", i, "[]byte")
			view.Body1 = append(view.Body1, l1)
		default:
			l1 := fmt.Sprintf("\tvar v%d %s\n", i, field.Type.Type())
			view.Body1 = append(view.Body1, l1)
		}

		// variable scanning
		l2 := fmt.Sprintf("\t\t&v%d,\n", i)
		view.Body2 = append(view.Body2, l2)

		switch field.Type.Base {
		case parse.Map, parse.Slice, parse.Struct, parse.Ptr:
			l3 := fmt.Sprintf("%serr = json.Unmarshal(v%d, &v.%s)\n%sif err != nil {\n%s\treturn nil, err\n%s}\n",
				oneTab, i, field.JoinParts("."), oneTab, oneTab, oneTab)
			view.Body3 = append(view.Body3, l3)
		default:
			l3 := fmt.Sprintf("\tv.%s = v%d\n", field.JoinParts("."), i)
			view.Body3 = append(view.Body3, l3)
		}
	}

	must(tScanRow.Execute(w, view))
}

const oneTab = "\t"
const twoTabs = "\t\t"

func WriteRowsFunc(w io.Writer, view View, table *schema.Table) {

	for i, field := range table.Fields {
		if field.Tags.Skip {
			continue
		}

		// temporary variable declaration
		switch field.Type.Base {
		case parse.Map, parse.Slice, parse.Struct, parse.Ptr:
			l1 := fmt.Sprintf("\tvar v%d %s\n", i, "[]byte")
			view.Body1 = append(view.Body1, l1)
		default:
			l1 := fmt.Sprintf("\tvar v%d %s\n", i, field.Type.Type())
			view.Body1 = append(view.Body1, l1)
		}

		// variable scanning
		l2 := fmt.Sprintf("\t\t\t&v%d,\n", i)
		view.Body2 = append(view.Body2, l2)

		switch field.Type.Base {
		case parse.Map, parse.Slice, parse.Struct, parse.Ptr:
			l3 := fmt.Sprintf("%serr = json.Unmarshal(v%d, &v.%s)\n%sif err != nil {\n%s\treturn nil, err\n%s}\n",
				twoTabs, i, field.JoinParts("."), twoTabs, twoTabs, twoTabs)
			view.Body3 = append(view.Body3, l3)
		default:
			l3 := fmt.Sprintf("\t\tv.%s = v%d\n", field.JoinParts("."), i)
			view.Body3 = append(view.Body3, l3)
		}
	}

	must(tScanRows.Execute(w, view))
}
