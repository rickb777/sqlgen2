package schema

import (
	"bytes"
	"io"
	"text/tabwriter"
	"strings"
	"fmt"
)

// Table returns a SQL statement to create the table.
func baseTableDDL(table *TableDescription, dialect Dialect, initial, final string) string {

	// use a large default buffer size of so that
	// the tabbing doesn't get prematurely flushed
	// resulting in un-even lines.
	var byt = make([]byte, 0, 100000)
	var buf = bytes.NewBuffer(byt)

	// use a tab writer to evenly space the column
	// names and column types.
	var tab = tabwriter.NewWriter(buf, 0, 8, 1, ' ', 0)
	comma := initial
	for _, field := range table.Fields {
		comma = dialect.FieldDDL(tab, field, comma)
	}
	io.WriteString(tab, final)

	// flush the tab writer to write to the buffer
	tab.Flush()

	return buf.String()
}

func backTickFieldDDL(w io.Writer, field *Field, comma string, dialect Dialect) string {
	io.WriteString(w, comma)
	comma = ",\\n\"+\n"

	io.WriteString(w, "\"\t`")
	io.WriteString(w, string(field.SqlName))
	io.WriteString(w, "`\t")
	io.WriteString(w, dialect.FieldAsColumn(field))
	return comma
}

func baseInsertDML(table *TableDescription, valuePlaceholders string, quoter func(Identifier) string) string {
	w := &bytes.Buffer{}
	w.WriteString(`"(`)

	table.Fields.NonAuto().SqlNames().Quoted(w, quoter)

	w.WriteString(") VALUES (")
	w.WriteString(valuePlaceholders)
	w.WriteString(`)"`)
	return w.String()
}

func baseUpdateDML(table *TableDescription, quoter func(Identifier) string, param func(int) string) string {
	w := &bytes.Buffer{}
	w.WriteString(`"`)

	comma := ""
	for _, id := range table.Fields.NonAuto().SqlNames() {
		io.WriteString(w, comma)
		io.WriteString(w, quoter(id))
		comma = "=?,"
	}

	w.WriteString("=? ")
	w.WriteString(baseWhereClause(FieldList{table.Primary}, 0, quoter, param))
	w.WriteString(`"`)
	return w.String()
}

func baseDeleteDML(table *TableDescription, fields FieldList, quoter func(Identifier) string, param func(int) string) string {
	w := &bytes.Buffer{}
	w.WriteString("DELETE FROM %s\n")
	w.WriteString(baseWhereClause(fields, 0, quoter, param))
	return w.String()
}

func baseSplitAndQuote(csv, before, between, after string) string {
	ids := strings.Split(csv, ",")
	w := bytes.NewBuffer(make([]byte, 0, len(ids)*10))
	sep := before
	for _, id := range ids {
		io.WriteString(w, sep)
		io.WriteString(w, string(id))
		sep = between
	}
	io.WriteString(w, after)
	return w.String()
}

func backTickQuoted(identifier Identifier) string {
	return fmt.Sprintf("`%s`", identifier)
}

const placeholders = "?,?,?,?,?,?,?,?,?,?"

func baseQueryPlaceholders(n int) string {
	if n == 0 {
		return ""
	} else if n <= 10 {
		m := (n * 2) - 1
		return placeholders[:m]
	}
	return strings.Repeat("?,", n-1) + "?"
}

func baseParamIsQuery(i int) string {
	return "?"
}

//func baseColumns(fields FieldList, withAuto, inline, assign bool, param func(int) string) string {
//	w := &bytes.Buffer{}
//	comma := ""
//	for i, field := range fields {
//		if withAuto || !field.Tags.Auto {
//			w.WriteString(comma)
//			comma = ", "
//
//			if !inline {
//				w.WriteString("\n")
//				w.WriteString(fieldIndentation)
//			}
//
//			w.WriteString(field.SqlName)
//
//			if assign {
//				w.WriteString("=")
//				w.WriteString(param(i + 1))
//			}
//		}
//	}
//	return w.String()
//}

const fieldIndentation = "\t"

// helper function to generate the Where clause
// section of a SQL statement
func baseWhereClause(fields FieldList, pos int, quoter func(Identifier) string, param func(int) string) string {
	var buf bytes.Buffer
	j := pos

	for i, field := range fields {
		switch {
		case i == 0:
			buf.WriteString("WHERE")
		default:
			buf.WriteString("\nAND")
		}

		buf.WriteString(" ")
		buf.WriteString(quoter(field.SqlName))
		buf.WriteString("=")
		buf.WriteString(param(j))

		j++
	}
	return buf.String()
}

//func (b *base) CreateTableSettings() string {
//	return ""
//}
