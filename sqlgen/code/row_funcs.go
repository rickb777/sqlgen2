package code

import (
	"fmt"
	"io"

	"github.com/rickb777/sqlgen2/schema"
)

func WriteRowsFunc(w io.Writer, view View) {

	for i, field := range view.Table.Fields {
		if field.Tags.Skip {
			continue
		}

		// temporary variable declaration
		switch field.Encode {
		case schema.ENCJSON, schema.ENCTEXT:
			l1 := fmt.Sprintf("\tvar v%d %s\n", i, "[]byte")
			view.Body1 = append(view.Body1, l1)
		default:
			l1 := fmt.Sprintf("\tvar v%d %s\n", i, field.Type.Type())
			view.Body1 = append(view.Body1, l1)
		}

		// variable scanning
		l2 := fmt.Sprintf("&v%d", i)
		view.Body2 = append(view.Body2, l2)

		switch field.Encode {
		case schema.ENCJSON:
			l3 := fmt.Sprintf("\t\terr = json.Unmarshal(v%d, &v.%s)\n\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n",
				i, field.JoinParts(0, "."))
			view.Body3 = append(view.Body3, l3)
		case schema.ENCTEXT:
			l3 := fmt.Sprintf("\t\terr = encoding.UnmarshalText(v%d, &v.%s)\n\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n",
				i, field.JoinParts(0, "."))
			view.Body3 = append(view.Body3, l3)
		default:
			l3 := fmt.Sprintf("\t\tv.%s = v%d\n", field.JoinParts(0, "."), i)
			view.Body3 = append(view.Body3, l3)
		}
	}

	must(tScanRows.Execute(w, view))
}
