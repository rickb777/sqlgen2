package schema

import (
	"fmt"
)

type posgres struct {
	base
}

func newPostgres() SDialect {
	d := &posgres{}
	d.base.SDialect = d
	return d
}

// see https://github.com/eaigner/hood/blob/master/mysql.go#L35

func postgresColumn(f *Field) string {
	// posgres uses a special column type
	// for autoincrementing keys.
	if f.Auto {
		return "serial"
	}

	switch f.SqlType {
	case INTEGER:
		return "integer"
	case BOOLEAN:
		return "boolean"
	case BLOB:
		return "byteaa"
	case VARCHAR:
		// assigns an arbitrary size if
		// none is provided.
		size := f.Size
		if size == 0 {
			size = 512
		}
		return fmt.Sprintf("varchar(%d)", size)
	case JSON:
		return "json"
	default:
		return ""
	}
}

func postgresToken(v SqlToken) string {
	switch v {
	case AUTO_INCREMENT:
		// postgres does not support the
		// auto-increment keyword.
		return ""
	case PRIMARY_KEY:
		return "primary key"
	default:
		return ""
	}
}

func (d *posgres) Param(i int) string {
	return fmt.Sprintf("$%d", i+1)
}
