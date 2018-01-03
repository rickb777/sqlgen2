package schema

import (
	"fmt"
	"github.com/rickb777/sqlgen2/sqlgen/parse"
	"bytes"
	"strconv"
)

type postgres struct{}

var Postgres Dialect = postgres{}

func (d postgres) String() string {
	return "Postgres"
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

	return column
}

func (dialect postgres) TableDDL(table *TableDescription) string {
	return baseTableDDL(table, dialect)
}

func (dialect postgres) InsertDML(table *TableDescription) string {
	return baseInsertDML(table)
}

func (dialect postgres) UpdateDML(table *TableDescription) string {
	return baseUpdateDML(table, table.Fields, postgresParam)
}

func (dialect postgres) DeleteDML(table *TableDescription, fields FieldList) string {
	return baseDeleteDML(table, fields, postgresParam)
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

const postgresPlaceholders = "$1,$2,$3,$4,$5,$6,$7,$8,$9"

// Placeholders returns a string containing the requested number of placeholders
// in the form used by PostgreSQL.
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

