package code

import (
	"fmt"
	"github.com/rickb777/sqlapi/schema"
	"io"
)

func WriteConstructInsert(w io.Writer, view View) {
	commaNeededAgain := true

	list := view.Table.Fields.NoSkips()

	view.Body1 = appendf(view.Body1, "\tq := tbl.Dialect().Quoter()")
	view.Body1 = appendf(view.Body1, nltab+"s = make([]interface{}, 0, %d)", len(list))

	view.Body2 = appendf(view.Body2, nltab+"comma := \"\"")
	view.Body2 = appendf(view.Body2, nltab+"w.WriteString(\" (\")")

	for _, field := range list {
		joinedName := field.JoinParts(0, ".")
		view.Body2 = appendf(view.Body2, "\n")

		if field.GetTags().Primary {
			view.Body2 = appendf(view.Body2, nltab+"if withPk {")
			view.Body2 = appendf(view.Body2, nltab+"\tq.QuoteW(w, %q)", field.SqlName)
			view.Body2 = appendf(view.Body2, nltab+"\tcomma = \",\"")
			view.Body2 = appendf(view.Body2, nltab+"\ts = append(s, v.%s)", joinedName)
			view.Body2 = appendf(view.Body2, nltab+"}")

		} else {
			switch field.Encode {
			case schema.ENCJSON:
				view.Body2 = appendf(view.Body2, nltab+"w.WriteString(comma)")
				view.Body2 = appendf(view.Body2, nltab+"q.QuoteW(w, %q)", field.SqlName)
				if commaNeededAgain {
					view.Body2 = appendf(view.Body2, nltab+"comma = \",\"")
					commaNeededAgain = false
				}

				view.Body2 = appendf(view.Body2, nltab+"x, err := json.Marshal(&v.%s)", joinedName)
				view.Body2 = appendf(view.Body2, nltab+"if err != nil {")
				view.Body2 = appendf(view.Body2, nltab+"\treturn nil, tbl.database.LogError(errors.WithStack(err))")
				view.Body2 = appendf(view.Body2, nltab+"}")
				view.Body2 = appendf(view.Body2, nltab+"s = append(s, x)")

			case schema.ENCTEXT:
				view.Body2 = appendf(view.Body2, nltab+"w.WriteString(comma)")
				view.Body2 = appendf(view.Body2, nltab+"q.QuoteW(w, %q)", field.SqlName)
				if commaNeededAgain {
					view.Body2 = appendf(view.Body2, nltab+"comma = \",\"")
					commaNeededAgain = false
				}

				view.Body2 = appendf(view.Body2, nltab+"x, err := encoding.MarshalText(&v.%s)", joinedName)
				view.Body2 = appendf(view.Body2, nltab+"if err != nil {")
				view.Body2 = appendf(view.Body2, nltab+"\treturn nil, tbl.database.LogError(errors.WithStack(err))")
				view.Body2 = appendf(view.Body2, nltab+"}")
				view.Body2 = appendf(view.Body2, nltab+"s = append(s, x)")

			default:
				if field.Type.IsPtr {
					view.Body2 = appendf(view.Body2, nltab+"if v.%s != nil {", joinedName)
					view.Body2 = appendf(view.Body2, nltab+"\tw.WriteString(comma)")
					view.Body2 = appendf(view.Body2, nltab+"\tq.QuoteW(w, %q)", field.SqlName)
					view.Body2 = appendf(view.Body2, nltab+"\ts = append(s, v.%s)", joinedName)
					if commaNeededAgain {
						view.Body2 = appendf(view.Body2, nltab+"\tcomma = \",\"")
					}
					view.Body2 = appendf(view.Body2, nltab+"}")

				} else {
					view.Body2 = appendf(view.Body2, nltab+"w.WriteString(comma)")
					view.Body2 = appendf(view.Body2, nltab+"q.QuoteW(w, %q)", field.SqlName)
					view.Body2 = appendf(view.Body2, nltab+"s = append(s, v.%s)", joinedName)
					if commaNeededAgain {
						view.Body2 = appendf(view.Body2, nltab+"comma = \",\"")
						commaNeededAgain = false
					}
				}
			}
		}
	}

	view.Body2 = appendf(view.Body2, "\n")
	view.Body2 = appendf(view.Body2, nltab+"w.WriteString(\")\")")
	must(tConstructInsert.Execute(w, view))
}

func WriteConstructUpdate(w io.Writer, view View) {
	commaNeededAgain := true

	list := view.Table.Fields.NoPrimary().NoSkips()

	view.Body1 = appendf(view.Body1, "\tq := tbl.Dialect().Quoter()")
	view.Body1 = appendf(view.Body1, nltab+"j := 1")
	view.Body1 = appendf(view.Body1, nltab+"s = make([]interface{}, 0, %d)", len(list))

	view.Body2 = appendf(view.Body2, nltab+"comma := \"\"")

	for _, field := range list {
		joinedName := field.JoinParts(0, ".")

		view.Body2 = appendf(view.Body2, "\n")
		view.Body2 = appendf(view.Body2, nltab+"w.WriteString(comma)")
		switch field.Encode {
		case schema.ENCJSON:
			view.Body2 = appendf(view.Body2, nltab+"q.QuoteW(w, %q)", field.SqlName)
			view.Body2 = appendf(view.Body2, nltab+"w.WriteString(\"=?\")")
			if commaNeededAgain {
				view.Body2 = appendf(view.Body2, nltab+"comma = \", \"")
				commaNeededAgain = false
			}
			view.Body2 = appendf(view.Body2, nltab+"j++")
			view.Body2 = appendf(view.Body2, "\n")

			view.Body2 = appendf(view.Body2, nltab+"x, err := json.Marshal(&v.%s)", joinedName)
			view.Body2 = appendf(view.Body2, nltab+"if err != nil {")
			view.Body2 = appendf(view.Body2, nltab+"\treturn nil, tbl.database.LogError(errors.WithStack(err))")
			view.Body2 = appendf(view.Body2, nltab+"}")
			view.Body2 = appendf(view.Body2, nltab+"s = append(s, x)")

		case schema.ENCTEXT:
			view.Body2 = appendf(view.Body2, nltab+"q.QuoteW(w, %q)", field.SqlName)
			view.Body2 = appendf(view.Body2, nltab+"w.WriteString(\"=?\")")
			view.Body2 = appendf(view.Body2, nltab+"j++")
			if commaNeededAgain {
				view.Body2 = appendf(view.Body2, nltab+"comma = \", \"")
				commaNeededAgain = false
			}

			view.Body2 = appendf(view.Body2, nltab+"x, err := encoding.MarshalText(&v.%s)", joinedName)
			view.Body2 = appendf(view.Body2, nltab+"if err != nil {")
			view.Body2 = appendf(view.Body2, nltab+"\treturn nil, tbl.database.LogError(errors.WithStack(err))")
			view.Body2 = appendf(view.Body2, nltab+"}")
			view.Body2 = appendf(view.Body2, nltab+"s = append(s, x)")
			view.Body2 = appendf(view.Body2, "\n")

		default:
			if field.Type.IsPtr {
				view.Body2 = appendf(view.Body2, nltab+"if v.%s != nil {", joinedName)
				view.Body2 = appendf(view.Body2, nltab+"\tq.QuoteW(w, %q)", field.SqlName)
				view.Body2 = appendf(view.Body2, nltab+"\tw.WriteString(\"=?\")")
				view.Body2 = appendf(view.Body2, nltab+"\ts = append(s, v.%s)", joinedName)
				view.Body2 = appendf(view.Body2, nltab+"\tj++")
				if commaNeededAgain {
					view.Body2 = appendf(view.Body2, nltab+"\tcomma = \", \"")
				}
				view.Body2 = appendf(view.Body2, nltab+"} else {")
				view.Body2 = appendf(view.Body2, nltab+"\tq.QuoteW(w, %q)", field.SqlName)
				view.Body2 = appendf(view.Body2, nltab+"\tw.WriteString(\"=NULL\")")
				view.Body2 = appendf(view.Body2, nltab+"}")

			} else {
				view.Body2 = appendf(view.Body2, nltab+"q.QuoteW(w, %q)", field.SqlName)
				view.Body2 = appendf(view.Body2, nltab+"w.WriteString(\"=?\")")
				view.Body2 = appendf(view.Body2, nltab+"s = append(s, v.%s)", joinedName)
				view.Body2 = appendf(view.Body2, nltab+"j++")
				if commaNeededAgain {
					view.Body2 = appendf(view.Body2, nltab+"comma = \", \"")
					commaNeededAgain = false
				}
			}
		}
	}

	must(tConstructUpdate.Execute(w, view))
}

const nltab = "\n\t"

func appendf(list []string, value string, args ...interface{}) []string {
	return append(list, fmt.Sprintf(value, args...))
}
