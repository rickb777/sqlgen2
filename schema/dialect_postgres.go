package schema

import (
	"fmt"
	"github.com/rickb777/sqlgen2/sqlgen/parse"
	"bytes"
	"strconv"
	"io"
	"strings"
)

type postgres struct{}

var Postgres Dialect = postgres{}

func (d postgres) Index() int {
	return PostgresIndex
}

func (d postgres) String() string {
	return "Postgres"
}

func (d postgres) Alias() string {
	return "PostgreSQL"
}

// https://www.postgresql.org/docs/9.5/static/datatype.html

func (dialect postgres) FieldAsColumn(field *Field) string {
	switch field.Encode {
	case ENCJSON:
		return "json"
	case ENCTEXT:
		return varchar(field.Tags.Size)
	}

	column := "byteaa"

	switch field.Type.Base {
	case parse.Int, parse.Int64:
		column = "bigint"
	case parse.Int8:
		column = "tinyint"
	case parse.Int16:
		column = "smallint"
	case parse.Int32:
		column = "int"
	case parse.Uint, parse.Uint64:
		column = "bigint unsigned"
	case parse.Uint8:
		column = "tinyint unsigned"
	case parse.Uint16:
		column = "smallint unsigned"
	case parse.Uint32:
		column = "int unsigned"
	case parse.Float32:
		column = "float"
	case parse.Float64:
		column = "double"
	case parse.Bool:
		column = "boolean"
	case parse.String:
		column = varchar(field.Tags.Size)
	}

	// postgres uses a special column type
	// for autoincrementing keys.
	if field.Tags.Auto {
		switch field.Type.Base {
		case parse.Int, parse.Int64, parse.Uint64:
			column = "bigserial"
		default:
			column = "serial"
		}
	}

	if field.Tags.Primary {
		column += " primary key"
	}

	if field.Type.IsPtr {
		column += " default null"
	}

	return column
}

func (dialect postgres) TableDDL(table *TableDescription) string {
	return baseTableDDL(table, dialect, " `\n", "`")
}

func (dialect postgres) FieldDDL(w io.Writer, field *Field, comma string) string {
	io.WriteString(w, comma)
	io.WriteString(w, "\t\"")
	io.WriteString(w, string(field.SqlName))
	io.WriteString(w, "\"\t")
	io.WriteString(w, dialect.FieldAsColumn(field))
	return ",\n" // for next iteration
}

func (dialect postgres) InsertDML(table *TableDescription) string {
	w := &bytes.Buffer{}
	w.WriteString("`(")

	table.Fields.NonAuto().SqlNames().MkString3W(w, `"`, `","`, `"`)

	w.WriteString(") VALUES (")
	w.WriteString(dialect.Placeholders(table.NumColumnNames(false)))
	w.WriteString(")")
	if table.Primary != nil {
		w.WriteString(" returning ")
		w.WriteString(doubleQuoter(table.Primary.SqlName))
	}
	w.WriteByte('`')
	return w.String()
}

func (dialect postgres) UpdateDML(table *TableDescription) string {
	w := &bytes.Buffer{}
	w.WriteString("`")

	comma := ""
	for j, field := range table.Fields {
		if !field.Tags.Auto {
			w.WriteString(comma)
			w.WriteString(doubleQuoter(field.SqlName))
			w.WriteString("=")
			w.WriteString(postgresParam(j))
			comma = ","
		}
	}

	w.WriteByte(' ')
	w.WriteString(baseWhereClause(FieldList{table.Primary}, 0, doubleQuoter, postgresParam))
	w.WriteByte('`')
	return w.String()
}

func (dialect postgres) TruncateDDL(tableName string, force bool) []string {
	if force {
		return []string{fmt.Sprintf("TRUNCATE %s CASCADE", tableName)}
	}

	return []string{fmt.Sprintf("TRUNCATE %s RESTRICT", tableName)}
}

func postgresParam(i int) string {
	return fmt.Sprintf("$%d", i+1)
}

func doubleQuoter(identifier string) string {
	w := bytes.NewBuffer(make([]byte, 0, len(identifier)*2))
	doubleQuoterW(w, identifier)
	return w.String()
}

func doubleQuoterW(w io.Writer, identifier string) {
	elements := strings.Split(identifier, ".")
	baseQuotedW(w, elements, `"`, `"."`, `"`)
}

func (dialect postgres) SplitAndQuote(csv string) string {
	return baseSplitAndQuote(csv, `"`, `","`, `"`)
}

func (dialect postgres) Quote(identifier string) string {
	return doubleQuoter(identifier)
}

func (dialect postgres) QuoteW(w io.Writer, identifier string) {
	doubleQuoterW(w, identifier)
}

func (dialect postgres) QuoteWithPlaceholder(w io.Writer, identifier string, idx int) {
	doubleQuoterW(w, identifier)
	io.WriteString(w, "$")
	io.WriteString(w, strconv.Itoa(idx))
}

func (dialect postgres) Quoter() func(identifier string) string {
	return doubleQuoter
}

func (dialect postgres) Placeholder(name string, j int) string {
	return fmt.Sprintf("$%d", j)
}

func (dialect postgres) Placeholders(n int) string {
	if n == 0 {
		return ""
	} else if n <= 9 {
		return postgresPlaceholders[:n*3-1]
	}
	buf := bytes.NewBufferString(postgresPlaceholders)
	for idx := 10; idx <= n; idx++ {
		if idx > 1 {
			buf.WriteByte(',')
		}
		buf.WriteByte('$')
		buf.WriteString(strconv.Itoa(idx))
	}
	return buf.String()
}

const postgresPlaceholders = "$1,$2,$3,$4,$5,$6,$7,$8,$9"

// ReplacePlaceholders converts a string containing '?' placeholders to
// the form used by PostgreSQL.
func (dialect postgres) ReplacePlaceholders(sql string) string {
	buf := &bytes.Buffer{}
	idx := 1
	for _, r := range sql {
		if r == '?' {
			buf.WriteByte('$')
			buf.WriteString(strconv.Itoa(idx))
			idx++
		} else {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}

func (dialect postgres) CreateTableSettings() string {
	return ""
}
