package code

import (
	. "fmt"
	"io"

	"github.com/rickb777/sqlgen2/schema"
)

func WriteRowsFunc(w io.Writer, view View) {

	for i, field := range view.Table.Fields {
		if field.Tags.Skip {
			continue
		}

		nullable := field.Type.NullableValue()

		// temporary variable declaration
		switch field.Encode {
		case schema.ENCJSON, schema.ENCTEXT:
			l1 := Sprintf("\tvar v%d %s\n", i, "[]byte")
			view.Body1 = append(view.Body1, l1)
		default:
			l1 := Sprintf("\tvar v%d %s\n", i, field.Type.Type())
			if nullable != "" {
				l1 = Sprintf("\tvar v%d sql.Null%s\n", i, nullable)
			}
			view.Body1 = append(view.Body1, l1)
		}

		// variable scanning
		l2 := Sprintf("&v%d", i)
		view.Body2 = append(view.Body2, l2)

		switch field.Encode {
		case schema.ENCJSON:
			l3 := Sprintf("\t\terr = json.Unmarshal(v%d, &v.%s)\n\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n",
				i, field.JoinParts(0, "."))
			view.Body3 = append(view.Body3, l3)
		case schema.ENCTEXT:
			l3 := Sprintf("\t\terr = encoding.UnmarshalText(v%d, &v.%s)\n\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n",
				i, field.JoinParts(0, "."))
			view.Body3 = append(view.Body3, l3)
		default:
			if nullable != "" {
				l3a := Sprintf("\t\tif v%d.Valid {\n", i)
				l3b := Sprintf("\t\t\tv.%s = &(%s(v%d.%s))\n", field.JoinParts(0, "."), field.Type.Name, i, nullable)
				l3c := "\t\t}\n"
				view.Body3 = append(view.Body3, l3a)
				view.Body3 = append(view.Body3, l3b)
				view.Body3 = append(view.Body3, l3c)
			} else {
				amp := ""
				if field.Type.IsPtr {
					amp = "&"
				}
				l3 := Sprintf("\t\tv.%s = %sv%d\n", field.JoinParts(0, "."), amp, i)
				view.Body3 = append(view.Body3, l3)
			}
		}
	}

	must(tScanRows.Execute(w, view))
}
