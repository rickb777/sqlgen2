package schema

import (
	"fmt"
)

type posgres struct {
	base
}

func newPostgres() Dialect {
	d := &posgres{}
	d.base.Dialect = d
	return d
}

func (d *posgres) Id() DialectId {
	return POSTGRES
}

// see https://github.com/eaigner/hood/blob/master/mysql.go#L35

func postgresColumn(f *Field) string {
	// posgres uses a special column type
	// for autoincrementing keys.
	if f.Auto {
		return "SERIAL"
	}

	switch f.Type {
	case INTEGER:
		return "INTEGER"
	case BOOLEAN:
		return "BOOLEAN"
	case BLOB:
		return "BYTEA"
	case VARCHAR:
		// assigns an arbitrary size if
		// none is provided.
		size := f.Size
		if size == 0 {
			size = 512
		}
		return fmt.Sprintf("VARCHAR(%d)", size)
	case JSON:
		return "JSON"
	default:
		return ""
	}
}

func (d *posgres) Token(v int) string {
	switch v {
	case AUTO_INCREMENT:
		// postgres does not support the
		// auto-increment keyword.
		return ""
	case PRIMARY_KEY:
		return "PRIMARY KEY"
	default:
		return ""
	}
}

func (d *posgres) Param(i int) string {
	return fmt.Sprintf("$%d", i+1)
}
