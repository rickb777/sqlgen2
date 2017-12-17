package code

import (
	"fmt"
	"io"
	"github.com/rickb777/sqlgen2/schema"
)

func WriteSliceFunc(w io.Writer, view View, withoutPk bool) {
	var depth int

	needsErr := false
	for _, field := range view.Table.Fields {
		if field.Tags.Skip || (withoutPk && field.Tags.Primary) {
			continue
		}

		switch field.Encode {
		case schema.ENCJSON:
			needsErr = true
		}
	}

	if needsErr {
		view.Body1 = append(view.Body1, "\tvar err error\n\n")
	}

	for i, field := range view.Table.Fields {
		if field.Tags.Skip || (withoutPk && field.Tags.Primary) {
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
		l3 := fmt.Sprintf("\t\tv%d,\n", i)
		view.Body3 = append(view.Body3, l3)

		switch field.Encode {
		case schema.ENCJSON:
			l2 := fmt.Sprintf("%s\tv%d, err = json.Marshal(&v.%s)\n%s\tif err != nil {\n%s\t\treturn nil, err\n%s\t}\n",
				tabs[:depth], i, field.JoinParts(0, "."), tabs[:depth], tabs[:depth], tabs[:depth])
			view.Body2 = append(view.Body2, l2)
		case schema.ENCTEXT:
			l2 := fmt.Sprintf("%s\tv%d, err = encoding.MarshalText(&v.%s)\n%s\tif err != nil {\n%s\t\treturn nil, err\n%s\t}\n",
				tabs[:depth], i, field.JoinParts(0, "."), tabs[:depth], tabs[:depth], tabs[:depth])
			view.Body2 = append(view.Body2, l2)
		default:
			l2 := fmt.Sprintf("%s\tv%d = v.%s\n", tabs[:depth], i, field.JoinParts(0, "."))
			view.Body2 = append(view.Body2, l2)
		}
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
