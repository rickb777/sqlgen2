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

// https://www.postgresql.org/docs/9.6/static/datatype.html
// https://www.convert-in.com/mysql-to-postgres-types-mapping.htm

func (dialect postgres) FieldAsColumn(field *Field) string {
	switch field.Encode {
	case ENCJSON:
		return "json"
	case ENCTEXT:
		return varchar(field.Tags.Size)
	}

	column := "bytea"
	dflt := field.Tags.Default

	switch field.Type.Base {
	case parse.Int, parse.Int64:
		column = "bigint"
	case parse.Int8:
		column = "int8"
	case parse.Int16:
		column = "smallint"
	case parse.Int32:
		column = "integer"
	case parse.Uint, parse.Uint64:
		// Some DBs (including postgresql) do not support unsigned integers. Rejecting
		// uint64 >= 1<<63 prevents them becoming indistinguishable from int64s < 0. If you need
		// to insert such int64 anyway, you have to explicitly convert in on input to int64 and
		// convert it back to uint64 on output - and take care to never insert a signed integer
		// into the same column.
		column = "bigint" // incomplete number range but more storage efficiency
	case parse.Uint8:
		column = "smallint"
	case parse.Uint16:
		column = "integer"
	case parse.Uint32:
		column = "bigint"
	case parse.Float32:
		column = "real"
	case parse.Float64:
		column = "double precision"
	case parse.Bool:
		column = "boolean"
	case parse.String:
		column = varchar(field.Tags.Size)
		dflt = fmt.Sprintf("'%s'", field.Tags.Default)
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

	return fieldTags(field, column, dflt)
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

func (dialect postgres) InsertHasReturningPhrase() bool {
	return true
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
	elements := strings.Split(strings.ToLower(identifier), ".")
	baseQuotedW(w, elements, `"`, `"."`, `"`)
}

func (dialect postgres) SplitAndQuote(csv string) string {
	return baseSplitAndQuote(strings.ToLower(csv), `"`, `","`, `"`)
}

func (dialect postgres) Quote(identifier string) string {
	return doubleQuoter(identifier)
}

func (dialect postgres) QuoteW(w io.Writer, identifier string) {
	doubleQuoterW(w, identifier)
}

func (dialect postgres) QuoteWithPlaceholder(w io.Writer, identifier string, idx int) {
	doubleQuoterW(w, identifier)
	io.WriteString(w, "=$")
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
func (dialect postgres) ReplacePlaceholders(sql string, _ []interface{}) string {
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
