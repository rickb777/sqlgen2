package schema

import (
	"bytes"
	"io"
	"text/tabwriter"
	"strings"
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
	io.WriteString(w, "\"\t`")
	io.WriteString(w, string(field.SqlName))
	io.WriteString(w, "`\t")
	io.WriteString(w, dialect.FieldAsColumn(field))
	return ",\\n\"+\n" // for next iteration
}

func baseInsertDML(table *TableDescription, valuePlaceholders string) string {
	w := &bytes.Buffer{}
	w.WriteString(`"(`)

	table.Fields.NonAuto().SqlNames().MkString3W(w, "`", "`,`", "`")

	w.WriteString(") VALUES (")
	w.WriteString(valuePlaceholders)
	w.WriteString(`)"`)
	return w.String()
}

func baseUpdateDML(table *TableDescription, quoter func(string) string, param func(int) string) string {
	w := &bytes.Buffer{}
	w.WriteString(`"`)

	table.Fields.NonAuto().SqlNames().MkString3W(w, "`", "`=?,`", "`=? ")

	w.WriteString(baseWhereClause(FieldList{table.Primary}, 0, quoter, param))
	w.WriteString(`"`)
	return w.String()
}

func baseSplitAndQuote(csv, before, between, after string) string {
	ids := strings.Split(csv, ",")
	return baseQuoted(ids, before, between, after)
}

func backTickQuoted(identifier string) string {
	w := bytes.NewBuffer(make([]byte, 0, len(identifier) * 2))
	backTickQuotedW(w, identifier)
	return w.String()
}

func backTickQuotedW(w io.Writer, identifier string) {
	elements := strings.Split(identifier, ".")
	baseQuotedW(w, elements, "`", "`.`", "`")
}

func baseQuoted(elements []string, before, between, after string) string {
	w := bytes.NewBuffer(make([]byte, 0, len(elements) * 16))
	baseQuotedW(w, elements, before, between, after)
	return w.String()
}

func baseQuotedW(w io.Writer, elements []string, before, between, after string) {
	io.WriteString(w, before)
	for i, e := range elements {
		if i > 0 {
			io.WriteString(w, between)
		}
		io.WriteString(w, e)
	}
	io.WriteString(w, after)
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

// helper function to generate the Where clause
// section of a SQL statement
func baseWhereClause(fields FieldList, pos int, quoter func(string) string, param func(int) string) string {
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
