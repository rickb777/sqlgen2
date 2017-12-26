package schema

import (
	"fmt"
	"github.com/rickb777/sqlgen2/sqlgen/parse"
)

type posgres struct {
	base
}

func newPostgres() SDialect {
	d := &posgres{}
	d.base.SDialect = d
	return d
}

// https://www.postgresql.org/docs/9.5/static/datatype.html

func postgresColumn(field *Field) string {
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

	// posgres uses a special column type
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

func (d *posgres) Param(i int) string {
	return fmt.Sprintf("$%d", i+1)
}
