package schema

import (
	"bytes"
	"fmt"
	"io"
	"text/tabwriter"
)

type SDialect interface {
	Table(*TableDescription, DialectId) string
	Index(*TableDescription, *Index) string
	Insert(*TableDescription) string
	Update(*TableDescription, []*Field) string
	Delete(*TableDescription, []*Field) string
	Param(int) string
	CreateTableSettings() string
}

type base struct {
	SDialect SDialect
}

// Table returns a SQL statement to create the table.
func (b *base) Table(t *TableDescription, did DialectId) string {

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

		if field.Tags.Primary {
			io.WriteString(tab, PRIMARY_KEY.AsToken(did))
		}

		if field.Tags.Auto {
			io.WriteString(tab, AUTO_INCREMENT.AsToken(did))
		}
	}

	// flush the tab writer to write to the buffer
	tab.Flush()

	return buf.String()
}

// Index returns a SQL statement to create the index.
func (b *base) Index(table *TableDescription, index *Index) string {
	var obj = "INDEX"
	if index.Unique {
		obj = "UNIQUE INDEX"
	}
	return fmt.Sprintf("CREATE %s %%s%%s%s ON %%s%%s (%s)", obj, index.Name,
		b.columns(index.Fields, true, true, false))
}

func (b *base) Insert(t *TableDescription) string {
	var fields []*Field
	var params []string
	var i int

	for _, field := range t.Fields {
		if !field.Tags.Auto {
			fields = append(fields, field)
			params = append(params, b.SDialect.Param(i))
			i++
		}
	}

	return fmt.Sprintf("INSERT INTO %%s%%s (%s\n) VALUES (%%s)",
		b.columns(fields, false, false, false))
}

func (b *base) Update(t *TableDescription, fields []*Field) string {
	return fmt.Sprintf("UPDATE %%s%%s SET %s %s",
		b.columns(t.Fields, false, false, true),
		b.whereClause(fields, len(t.Fields)))
}

func (b *base) Delete(t *TableDescription, fields []*Field) string {
	return fmt.Sprintf("DELETE FROM %%s%%s%s",
		b.whereClause(fields, 0))
}

// Param returns the parameters symbol used in prepared sql statements.
func (b *base) Param(i int) string {
	return "?"
}

// helper function to generate a block of columns. You
// can optionally generate in inline list of columns,
// include an assignment operator, and include column
// definitions.
func (b *base) columns(fields []*Field, withAuto, inline, assign bool) string {
	w := &bytes.Buffer{}
	comma := ""
	for i, field := range fields {
		if withAuto || !field.Tags.Auto {
			io.WriteString(w, comma)
			comma = ", "

			if !inline {
				io.WriteString(w, "\n")
				io.WriteString(w, fieldIndentation)
			}

			io.WriteString(w, field.SqlName)

			if assign {
				io.WriteString(w, "=")
				io.WriteString(w, b.SDialect.Param(i))
			}
		}
	}
	return w.String()
}

const fieldIndentation = "\t"

// helper function to generate the Where clause
// section of a SQL statement
func (b *base) whereClause(fields []*Field, pos int) string {
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
		buf.WriteString(b.SDialect.Param(i + pos))

		i++
	}
	return buf.String()
}

func (b *base) CreateTableSettings() string {
	return ""
}
