package sqlgen

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"
	"github.com/rickb777/sqlgen/schema"
)

type base struct {
	Dialect Dialect
}

// Table returns a SQL statement to create the table.
func (b *base) Table(t *schema.Table) string {

	// use a large default buffer size of so that
	// the tabbing doesn't get prematurely flushed
	// resulting in un-even lines.
	var byt = make([]byte, 0, 100000)
	var buf = bytes.NewBuffer(byt)

	// use a tab writer to evenly space the column
	// names and column types.
	var tab = tabwriter.NewWriter(buf, 0, 8, 1, ' ', 0)
	b.columnw(tab, t.Fields, true, false, false, true)

	// flush the tab writer to write to the buffer
	tab.Flush()

	return fmt.Sprintf("CREATE TABLE IF NOT EXISTS %%s (%s\n);", buf.String())
}

// Index returns a SQL statement to create the index.
func (b *base) Index(table *schema.Table, index *schema.Index) string {
	var obj = "INDEX"
	if index.Unique {
		obj = "UNIQUE INDEX"
	}
	return fmt.Sprintf("CREATE %s IF NOT EXISTS %s ON %%s (%s)", obj, index.Name,
		b.columns(index.Fields, true, true, false, false))
}

func (b *base) ColumnParams(t *schema.Table, withAuto bool) string {
	n := len(t.ColumnNames(withAuto))
	return strings.Join(b.Dialect.Params(0, n), ",")
}

func (b *base) Insert(t *schema.Table) string {
	var fields []*schema.Field
	var params []string
	var i int

	for _, field := range t.Fields {
		if !field.Auto {
			fields = append(fields, field)
			params = append(params, b.Dialect.Param(i))
			i++
		}
	}

	return fmt.Sprintf("INSERT INTO %%s (%s\n) VALUES (%s)",
		b.columns(fields, false, false, false, false),
		strings.Join(params, ","))
}

func (b *base) Update(t *schema.Table, fields []*schema.Field) string {
	return fmt.Sprintf("UPDATE %%s SET %s %s",
		b.columns(t.Fields, false, false, true, false),
		b.whereClause(fields, len(t.Fields)))
}

func (b *base) Delete(t *schema.Table, fields []*schema.Field) string {
	return fmt.Sprintf("DELETE FROM %%s%s",
		b.whereClause(fields, 0))
}

// Param returns the parameters symbol used in prepared sql statements.
func (b *base) Param(i int) string {
	return "?"
}

// Params returns the range of parameters symbols between two indices, which are half-inclusive.
func (b *base) Params(from, to int) []string {
	params := make([]string, 0, to-from)
	for i := from; i < to; i++ {
		params = append(params, b.Dialect.Param(i))
	}
	return params
}

// Column returns a SQL type for the given field.
//
// For Mysql and Postgres see:
// https://github.com/eaigner/hood/blob/master/mysql.go#L35
func (b *base) Column(f *schema.Field) string {
	switch f.Type {
	case schema.INTEGER:
		return "INTEGER"
	case schema.BOOLEAN:
		return "BOOLEAN"
	case schema.BLOB:
		return "BLOB"
	case schema.VARCHAR:
		return "TEXT"
	default:
		return "TEXT"
	}
}

// Token returns the SQL string for the requested token.
func (b *base) Token(v int) string {
	switch v {
	case schema.AUTO_INCREMENT:
		return "AUTOINCREMENT"
	case schema.PRIMARY_KEY:
		return "PRIMARY KEY"
	default:
		return ""
	}
}

// helper function to generate a block of columns. You
// can optionally generate in inline list of columns,
// include an assignment operator, and include column
// definitions.
func (b *base) columns(fields []*schema.Field, withAuto, inline, assign, ddl bool) string {
	var buf bytes.Buffer
	b.columnw(&buf, fields, withAuto, inline, assign, ddl)
	return buf.String()
}

const fieldIndentation = " "

// helper function to write a block of columns to w.
func (b *base) columnw(w io.Writer, fields []*schema.Field, withAuto, inline, assign, ddl bool) {

	comma := ""
	for i, field := range fields {
		if withAuto || !field.Auto {
			io.WriteString(w, comma)
			comma = ","

			if !inline {
				io.WriteString(w, "\n")
				io.WriteString(w, fieldIndentation)
			}

			io.WriteString(w, field.SqlName)

			if assign {
				io.WriteString(w, "=")
				io.WriteString(w, b.Dialect.Param(i))
			}

			if ddl {
				io.WriteString(w, "\t")
				io.WriteString(w, b.Dialect.Column(field))

				if field.Primary {
					io.WriteString(w, " ")
					io.WriteString(w, b.Dialect.Token(schema.PRIMARY_KEY))
				}

				if field.Auto {
					io.WriteString(w, " ")
					io.WriteString(w, b.Dialect.Token(schema.AUTO_INCREMENT))
				}
			}
		}
	}
}

// helper function to generate the Where clause
// section of a SQL statement
func (b *base) whereClause(fields []*schema.Field, pos int) string {
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
		buf.WriteString(b.Dialect.Param(i + pos))

		i++
	}
	return buf.String()
}
