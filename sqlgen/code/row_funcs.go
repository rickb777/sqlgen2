package code

import (
	"fmt"
	"io"

	"github.com/rickb777/sqlgen2/schema"
	"text/template"
)

func writeRowFunc(w io.Writer, view View, table *schema.TableDescription, tabs string, tmpl *template.Template) {

	for i, field := range table.Fields {
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
		l2 := fmt.Sprintf("%s\t&v%d,\n", tabs, i)
		view.Body2 = append(view.Body2, l2)

		switch field.Encode {
		case schema.ENCJSON:
			l3 := fmt.Sprintf("%serr = json.Unmarshal(v%d, &v.%s)\n%sif err != nil {\n%s\treturn nil, err\n%s}\n",
				tabs, i, field.JoinParts(0, "."), tabs, tabs, tabs)
			view.Body3 = append(view.Body3, l3)
		case schema.ENCTEXT:
			l3 := fmt.Sprintf("%serr = encoding.UnmarshalText(v%d, &v.%s)\n%sif err != nil {\n%s\treturn nil, err\n%s}\n",
				tabs, i, field.JoinParts(0, "."), tabs, tabs, tabs)
			view.Body3 = append(view.Body3, l3)
		default:
			l3 := fmt.Sprintf("%sv.%s = v%d\n", tabs, field.JoinParts(0, "."), i)
			view.Body3 = append(view.Body3, l3)
		}
	}

	must(tmpl.Execute(w, view))
}

func WriteRowFunc(w io.Writer, view View, table *schema.TableDescription) {
	writeRowFunc(w, view, table, oneTab, tScanRow)
}

func WriteRowsFunc(w io.Writer, view View, table *schema.TableDescription) {
	writeRowFunc(w, view, table, twoTabs, tScanRows)
}

const oneTab = "\t"
const twoTabs = "\t\t"
