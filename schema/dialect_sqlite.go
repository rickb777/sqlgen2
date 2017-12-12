package schema

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

func sqliteColumn(f *Field) string {
	switch f.SqlType {
	case INTEGER:
		// The value is a signed integer, stored in 1, 2, 3, 4, 6, or 8 bytes depending on the magnitude of the value
		return mysqlIntegerColumn(f)

	case REAL:
		// The value is a floating point value, stored as an 8-byte IEEE floating point number.
		return mysqlRealColumn(f)

	case BOOLEAN:
		return "boolean"

	case VARCHAR:
		// The value is a text string, stored using the database encoding (UTF-8, UTF-16BE or UTF-16LE).
		return "text"

	case BLOB:
		// The value is a blob of data, stored exactly as it was input.
		return "blob"
	}

	return "text"
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
