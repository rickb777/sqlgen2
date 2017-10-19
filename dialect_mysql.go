package sqlgen

import (
	"fmt"
	"github.com/rickb777/sqlgen/schema"
)

type mysql struct {
	base
}

func newMySQL() Dialect {
	d := &mysql{}
	d.base.Dialect = d
	return d
}

func (d *mysql) Column(f *schema.Field) string {
	switch f.Type {
	case schema.INTEGER:
		return "INTEGER"
	case schema.BOOLEAN:
		return "BOOLEAN"
	case schema.BLOB:
		return "MEDIUMBLOB"
	case schema.VARCHAR:
		// assigns an arbitrary size if
		// none is provided.
		size := f.Size
		if size == 0 {
			size = 512
		}
		return fmt.Sprintf("VARCHAR(%d)", size)
	default:
		return ""
	}
}

func (d *mysql) Token(v int) string {
	switch v {
	case schema.AUTO_INCREMENT:
		return "AUTO_INCREMENT"
	case schema.PRIMARY_KEY:
		return "PRIMARY KEY"
	default:
		return ""
	}
}
