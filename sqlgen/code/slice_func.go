package code

import (
	"fmt"
	"io"
	"github.com/rickb777/sqlgen2/schema"
)

func WriteSliceFunc(w io.Writer, view View, withoutPk bool) {
	for i, field := range view.Table.Fields {
		if field.Tags.Skip || (withoutPk && field.Tags.Primary) {
			continue
		}

		switch field.Encode {
		case schema.ENCJSON:
			l2 := fmt.Sprintf("\tv%d, err := json.Marshal(&v.%s)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n",
				i, field.JoinParts(0, "."))
			view.Body2 = append(view.Body2, l2)

			l3 := fmt.Sprintf("\t\tv%d,\n", i)
			view.Body3 = append(view.Body3, l3)

		case schema.ENCTEXT:
			l2 := fmt.Sprintf("\tv%d, err := encoding.MarshalText(&v.%s)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n",
				i, field.JoinParts(0, "."))
			view.Body2 = append(view.Body2, l2)

			l3 := fmt.Sprintf("\t\tv%d,\n", i)
			view.Body3 = append(view.Body3, l3)

		default:
			l3 := fmt.Sprintf("\t\tv.%s,\n", field.JoinParts(0, "."))
			view.Body3 = append(view.Body3, l3)
		}
	}

	if withoutPk {
		view.Suffix = "WithoutPk"
	}
	must(tSliceRow.Execute(w, view))
}
