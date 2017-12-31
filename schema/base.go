package schema

import (
	"bytes"
	"io"
	"text/tabwriter"
)

// Table returns a SQL statement to create the table.
func baseTableDDL(t *TableDescription, did Dialect) string {

	// use a large default buffer size of so that
	// the tabbing doesn't get prematurely flushed
	// resulting in un-even lines.
	var byt = make([]byte, 0, 100000)
	var buf = bytes.NewBuffer(byt)

	// use a tab writer to evenly space the column
	// names and column types.
	var tab = tabwriter.NewWriter(buf, 0, 8, 1, ' ', 0)
	comma := ""
	for _, field := range t.Fields {
		io.WriteString(tab, comma)
		comma = ","

		io.WriteString(tab, "\n")
		io.WriteString(tab, fieldIndentation)

		io.WriteString(tab, field.SqlName)

		io.WriteString(tab, "\t")
		io.WriteString(tab, field.AsColumn(did))
	}

	// flush the tab writer to write to the buffer
	tab.Flush()

	return buf.String()
}

// IndexDDL returns a SQL statement to create the index.
func baseIndexDDL(table *TableDescription, index *Index) string {
	w := &bytes.Buffer{}

	w.WriteString("CREATE ")
	if index.Unique {
		w.WriteString("UNIQUE ")
	}
	w.WriteString("INDEX %s%s")
	w.WriteString(index.Name)
	w.WriteString(" ON %s%s (")

	comma := ""
	for _, field := range index.Fields {
		w.WriteString(comma)
		w.WriteString(field.SqlName)
		comma = ", "
	}

	w.WriteString(")")
	return w.String()
}

func baseInsertDML(t *TableDescription) string {
	w := &bytes.Buffer{}
	w.WriteString("INSERT INTO %s%s (")

	comma := ""
	for _, field := range t.Fields {
		if !field.Tags.Auto {
			w.WriteString(comma)
			w.WriteString("\n")
			w.WriteString(fieldIndentation)
			w.WriteString(field.SqlName)
			comma = ", "
		}
	}

	w.WriteString("\n) VALUES (%s)")
	return w.String()
}

func baseUpdateDML(t *TableDescription, fields []*Field, param func(int) string) string {
	w := &bytes.Buffer{}
	w.WriteString("UPDATE %%s%%s SET ")

	comma := ""
	for i, field := range fields {
		if !field.Tags.Auto {
			w.WriteString(comma)
			w.WriteString("\n")
			w.WriteString(fieldIndentation)
			w.WriteString(field.SqlName)
			w.WriteString("=")
			w.WriteString(param(i + 1))
			comma = ", "
		}
	}

	w.WriteString(baseWhereClause(fields, len(t.Fields), param))
	return w.String()
}

func baseDeleteDML(t *TableDescription, fields []*Field, param func(int) string) string {
	w := &bytes.Buffer{}
	w.WriteString("DELETE FROM %s%s")
	w.WriteString(baseWhereClause(fields, 0, param))
	return w.String()
}

// Param returns the parameters symbol used in prepared sql statements.
func baseParam(i int) string {
	return "?"
}

func baseColumns(fields []*Field, withAuto, inline, assign bool, param func(int) string) string {
	w := &bytes.Buffer{}
	comma := ""
	for i, field := range fields {
		if withAuto || !field.Tags.Auto {
			w.WriteString(comma)
			comma = ", "

			if !inline {
				w.WriteString("\n")
				w.WriteString(fieldIndentation)
			}

			w.WriteString(field.SqlName)

			if assign {
				w.WriteString("=")
				w.WriteString(param(i + 1))
			}
		}
	}
	return w.String()
}

const fieldIndentation = "\t"

// helper function to generate the Where clause
// section of a SQL statement
func baseWhereClause(fields []*Field, pos int, param func(int) string) string {
	var buf bytes.Buffer

	var i int
	for _, field := range fields {
		buf.WriteString("\n ")
		switch {
		case i == 0:
			buf.WriteString("WHERE")
		default:
			buf.WriteString("AND")
		}

		buf.WriteString(" ")
		buf.WriteString(field.SqlName)
		buf.WriteString("=")
		buf.WriteString(param(i + pos))

		i++
	}
	return buf.String()
}

//func (b *base) CreateTableSettings() string {
//	return ""
//}
