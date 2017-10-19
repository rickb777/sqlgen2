package sqlgen

import (
	"fmt"
	"github.com/rickb777/sqlgen/schema"
)

type posgres struct {
	base
}

func newPostgres() Dialect {
	d := &posgres{}
	d.base.Dialect = d
	return d
}

func (d *posgres) Column(f *schema.Field) string {
	// posgres uses a special column type
	// for autoincrementing keys.
	if f.Auto {
		return "SERIAL"
	}

	switch f.Type {
	case schema.INTEGER:
		return "INTEGER"
	case schema.BOOLEAN:
		return "BOOLEAN"
	case schema.BLOB:
		return "BYTEA"
	case schema.VARCHAR:
		// assigns an arbitrary size if
		// none is provided.
		size := f.Size
		if size == 0 {
			size = 512
		}
		return fmt.Sprintf("VARCHAR(%d)", size)
	case schema.JSON:
		return "JSON"
	default:
		return ""
	}
}

func (d *posgres) Token(v int) string {
	switch v {
	case schema.AUTO_INCREMENT:
		// postgres does not support the
		// auto-increment keyword.
		return ""
	case schema.PRIMARY_KEY:
		return "PRIMARY KEY"
	default:
		return ""
	}
}

func (d *posgres) Param(i int) string {
	return fmt.Sprintf("$%d", i+1)
}
