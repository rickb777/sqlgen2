package schema

import "github.com/rickb777/sqlgen2/sqlgen/parse"

//import (
//	_ "github.com/mattn/go-sqlite3"
//)

type sqlite struct {
	base
}

func newSQLite() SDialect {
	d := &sqlite{}
	d.base.SDialect = d
	return d
}

// For integers, the value is a signed integer, stored in 1, 2, 3, 4, 6, or 8 bytes depending on the magnitude of the value
// For reals, the value is a floating point value, stored as an 8-byte IEEE floating point number.

func sqliteColumn(f *Field) string {
	switch f.Encode {
	case ENCJSON:
		return "text"
	case ENCTEXT:
		return "text"
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
		return "text"
	}

	// The value is a blob of data, stored exactly as it was input.
	return "blob"
}

func sqliteToken(v SqlToken) string {
	switch v {
	case AUTO_INCREMENT:
		// in sqlite, "autoincrement" is best avoided
		// https://sqlite.org/autoinc.html
		return ""
	case PRIMARY_KEY:
		return " primary key"
	default:
		return ""
	}
}
