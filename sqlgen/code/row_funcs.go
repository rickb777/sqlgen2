package code

import (
	. "fmt"
	"io"

	"github.com/rickb777/sqlgen2/schema"
	"strings"
)

func WriteRowsFunc(w io.Writer, view View) {

	for i, field := range view.Table.Fields {
		if !field.Tags.Skip {
			nullable := ""
			if !field.Type.IsScanner && field.Encode != schema.ENCDRIVER {
				// Scanner types should not have a NullString proxy
				nullable = field.Type.NullableValue()
			}

			l1 := writeRowDecl(i, field, nullable)
			view.Body1 = append(view.Body1, l1)

			l2 := writeRowRef(i, field, nullable)
			view.Body2 = append(view.Body2, l2)

			l3 := writeRowAssignment(i, field, nullable)
			view.Body3 = append(view.Body3, l3...)
		}
	}

	must(tScanRows.Execute(w, view))
}

const indent = "\t\t"

func writeRowDecl(i int, field *schema.Field, nullable string) string {
	t := ""
	switch field.Encode {
	case schema.ENCJSON, schema.ENCTEXT:
		t = "[]byte"

	case schema.ENCDRIVER:
		t = field.Type.Star() + field.Type.Type() //+ " // driver"

	default:
		if nullable != "" {
			t = "sql.Null"+nullable
		} else {
			t = field.Type.Star() + field.Type.Type() //+ " // plain"
		}
	}
	return indent+Sprintf("var v%d %s\n", i, t)
}

func writeRowRef(i int, field *schema.Field, nullable string) string {
	amp := "&"
	switch field.Encode {
	case schema.ENCJSON, schema.ENCTEXT:
	case schema.ENCDRIVER:
		if field.Type.IsPtr {
			amp = ""
		}
	default:
	}
	return Sprintf("%sv%d", amp, i)
}

func writeRowAssignment(i int, field *schema.Field, nullable string) []string {
	var lines []string

	switch field.Encode {
	case schema.ENCJSON:
		l1 := Sprintf(indent+"err = json.Unmarshal(v%d, &v.%s)\n", i, field.JoinParts(0, "."))
		lines = append(lines, l1)
		lines = append(lines, indent+"if err != nil {\n")
		lines = append(lines, indent+"\treturn nil, n, err\n")
		lines = append(lines, indent+"}\n")

	case schema.ENCTEXT:
		l1 := Sprintf(indent+"err = encoding.UnmarshalText(v%d, &v.%s)\n",
			i, field.JoinParts(0, "."))
		lines = append(lines, l1)
		lines = append(lines, indent+"if err != nil {\n")
		lines = append(lines, indent+"\treturn nil, n, err\n")
		lines = append(lines, indent+"}\n")

	case schema.ENCDRIVER:
		l1 := Sprintf(indent+"v.%s = v%d\n", field.JoinParts(0, "."), i)
		lines = append(lines, l1)

	default:
		if nullable != "" {
			l1 := Sprintf(indent+"if v%d.Valid {\n", i)
			l2 := Sprintf(indent+"\ta := %s(v%d.%s)\n", field.Type.Type(), i, nullable)
			if field.Type.Name == strings.ToLower(nullable) {
				l2 = Sprintf(indent+"\ta := v%d.%s\n", i, nullable)
			}
			l3 := Sprintf(indent+"\tv.%s = &a\n", field.JoinParts(0, "."))

			lines = append(lines, l1)
			lines = append(lines, l2)
			lines = append(lines, l3)
			lines = append(lines, indent+"}\n")

		} else {
			l1 := Sprintf(indent+"v.%s = v%d\n", field.JoinParts(0, "."), i)
			lines = append(lines, l1)
		}
	}

	return lines
}
