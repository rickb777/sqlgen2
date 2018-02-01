package code

import (
	"fmt"
	"io"
	"github.com/rickb777/sqlgen2/schema"
)

func WriteConstructInsert(w io.Writer, view View) {
	commaNeeded := true

	list := view.Table.Fields.NoSkips()

	view.Body1 = append(view.Body1, fmt.Sprintf("\ts = make([]interface{}, 0, %d)\n", len(list)))

	view.Body2 = append(view.Body2, "\tcomma := \"\"\n")
	view.Body2 = append(view.Body2, "\tio.WriteString(w, \" (\")\n\n")

	for _, field := range list {
		joinedName := field.JoinParts(0, ".")

		if field.Tags.Primary {
			view.Body2 = append(view.Body2, "\tif withPk {\n")
			view.Body2 = append(view.Body2, fmt.Sprintf("\t\tdialect.QuoteW(w, %q)\n", field.SqlName))
			view.Body2 = append(view.Body2, "\t\tcomma = \",\"\n")
			view.Body2 = append(view.Body2, fmt.Sprintf("\t\ts = append(s, v.%s)\n", joinedName))
			view.Body2 = append(view.Body2, "\t}\n\n")

		} else {
			switch field.Encode {
			case schema.ENCJSON:
				view.Body2 = append(view.Body2, "\tio.WriteString(w, comma)\n\n")
				view.Body2 = append(view.Body2, fmt.Sprintf("\tdialect.QuoteW(w, %q)\n", field.SqlName))
				if commaNeeded {
					view.Body2 = append(view.Body2, "\tcomma = \",\"\n")
				}

				view.Body2 = append(view.Body2, fmt.Sprintf("\tx, err := json.Marshal(&v.%s)\n", joinedName))
				view.Body2 = append(view.Body2, "\tif err != nil {\n\t\treturn nil, err\n\t}\n")
				view.Body2 = append(view.Body2, "\ts = append(s, x)\n")

			case schema.ENCTEXT:
				view.Body2 = append(view.Body2, "\tio.WriteString(w, comma)\n\n")
				view.Body2 = append(view.Body2, fmt.Sprintf("\tdialect.QuoteW(w, %q)\n", field.SqlName))
				if commaNeeded {
					view.Body2 = append(view.Body2, "\tcomma = \",\"\n")
				}

				view.Body2 = append(view.Body2, fmt.Sprintf("\tx, err := encoding.MarshalText(&v.%s)\n", joinedName))
				view.Body2 = append(view.Body2, "\tif err != nil {\n\t\treturn nil, err\n\t}\n")
				view.Body2 = append(view.Body2, "\ts = append(s, x)\n")

			default:
				if field.Type.IsPtr {
					view.Body2 = append(view.Body2, fmt.Sprintf("\tif v.%s != nil {\n", joinedName))
					view.Body2 = append(view.Body2, "\t\tio.WriteString(w, comma)\n\n")
					view.Body2 = append(view.Body2, fmt.Sprintf("\t\tdialect.QuoteW(w, %q)\n", field.SqlName))
					view.Body2 = append(view.Body2, fmt.Sprintf("\t\ts = append(s, v.%s)\n", joinedName))
					if commaNeeded {
						view.Body2 = append(view.Body2, "\t\tcomma = \",\"\n")
					}
					view.Body2 = append(view.Body2, "\t}\n")

				} else {
					view.Body2 = append(view.Body2, "\tio.WriteString(w, comma)\n\n")
					view.Body2 = append(view.Body2, fmt.Sprintf("\tdialect.QuoteW(w, %q)\n", field.SqlName))
					view.Body2 = append(view.Body2, fmt.Sprintf("\ts = append(s, v.%s)\n", joinedName))
					if commaNeeded {
						view.Body2 = append(view.Body2, "\tcomma = \",\"\n")
					}
					commaNeeded = false
				}
			}
		}
	}

	view.Body2 = append(view.Body2, "\tio.WriteString(w, \")\")")
	must(tConstructInsert.Execute(w, view))
}

func WriteConstructUpdate(w io.Writer, view View) {
	commaNeeded := true

	list := view.Table.Fields.NoSkipOrPrimary()

	view.Body1 = append(view.Body1, "\tj := 1\n")
	view.Body1 = append(view.Body1, fmt.Sprintf("\ts = make([]interface{}, 0, %d)\n", len(list)))

	view.Body2 = append(view.Body2, "\tcomma := \"\"\n")

	for _, field := range list {
		joinedName := field.JoinParts(0, ".")

		view.Body2 = append(view.Body2, "\n\tio.WriteString(w, comma)\n")
		switch field.Encode {
		case schema.ENCJSON:
			view.Body2 = append(view.Body2, fmt.Sprintf("\tdialect.QuoteWithPlaceholder(w, %q, j)\n", field.SqlName))
			if commaNeeded {
				view.Body2 = append(view.Body2, "\tcomma = \", \"\n")
			}
			view.Body2 = append(view.Body2, "\t\tj++\n")

			view.Body2 = append(view.Body2, fmt.Sprintf("\tx, err := json.Marshal(&v.%s)\n", joinedName))
			view.Body2 = append(view.Body2, "\tif err != nil {\n\t\treturn nil, err\n\t}\n")
			view.Body2 = append(view.Body2, "\ts = append(s, x)\n")

		case schema.ENCTEXT:
			view.Body2 = append(view.Body2, fmt.Sprintf("\tdialect.QuoteWithPlaceholder(w, %q, j)\n", field.SqlName))
			if commaNeeded {
				view.Body2 = append(view.Body2, "\tcomma = \", \"\n")
			}
			view.Body2 = append(view.Body2, "\t\tj++\n")

			view.Body2 = append(view.Body2, fmt.Sprintf("\tx, err := encoding.MarshalText(&v.%s)\n", joinedName))
			view.Body2 = append(view.Body2, "\tif err != nil {\n\t\treturn nil, err\n\t}\n")
			view.Body2 = append(view.Body2, "\ts = append(s, x)\n")

		default:
			if field.Type.IsPtr {
				view.Body2 = append(view.Body2, fmt.Sprintf("\tif v.%s != nil {\n", joinedName))
				view.Body2 = append(view.Body2, fmt.Sprintf("\t\tdialect.QuoteWithPlaceholder(w, %q, j)\n", field.SqlName))
				view.Body2 = append(view.Body2, fmt.Sprintf("\t\ts = append(s, v.%s)\n", joinedName))
				if commaNeeded {
					view.Body2 = append(view.Body2, "\t\tcomma = \", \"\n")
				}
				view.Body2 = append(view.Body2, "\t\tj++\n")
				view.Body2 = append(view.Body2, "\t} else {\n")
				view.Body2 = append(view.Body2, fmt.Sprintf("\t\tdialect.QuoteW(w, %q)\n", field.SqlName))
				view.Body2 = append(view.Body2, "\t\tio.WriteString(w, \"=NULL\")\n")
				view.Body2 = append(view.Body2, "\t}\n")

			} else {
				view.Body2 = append(view.Body2, fmt.Sprintf("\tdialect.QuoteWithPlaceholder(w, %q, j)\n", field.SqlName))
				view.Body2 = append(view.Body2, fmt.Sprintf("\ts = append(s, v.%s)\n", joinedName))
				if commaNeeded {
					view.Body2 = append(view.Body2, "\tcomma = \", \"\n")
				}
				view.Body2 = append(view.Body2, "\t\tj++\n")
				commaNeeded = false
			}
		}
	}

	must(tConstructUpdate.Execute(w, view))
}
