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

func postgresColumn(f *Field) string {
	// posgres uses a special column type
	// for autoincrementing keys.
	if f.Tags.Auto {
		switch f.Type.Base {
		case parse.Int64, parse.Uint64:
			return "bigserial"
		}
		return "serial"
	}

	switch f.Encode {
	case ENCJSON:
		return "json"
	case ENCTEXT:
		return varchar(f.Tags.Size)
	}

	switch f.Type.Base {
	case parse.Int, parse.Int64:
		return "bigint"
	case parse.Int8:
		return "tinyint"
	case parse.Int16:
		return "smallint"
	case parse.Int32:
		return "int"
	case parse.Uint, parse.Uint64:
		return "bigint unsigned"
	case parse.Uint8:
		return "tinyint unsigned"
	case parse.Uint16:
		return "smallint unsigned"
	case parse.Uint32:
		return "int unsigned"
	case parse.Float32:
		return "float"
	case parse.Float64:
		return "double"
	case parse.Bool:
		return "boolean"
	case parse.String:
		return varchar(f.Tags.Size)
	}

	return "byteaa"
}

func postgresToken(v SqlToken) string {
	switch v {
	case AUTO_INCREMENT:
		// postgres does not support the auto-increment keyword.
		return ""
	case PRIMARY_KEY:
		return " primary key"
	default:
		return ""
	}
}

func (d *posgres) Param(i int) string {
	return fmt.Sprintf("$%d", i+1)
}
