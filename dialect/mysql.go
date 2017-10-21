package dialect

import (
	"github.com/rickb777/sqlgen/schema"
	"fmt"
)

type MySQLDialect struct{}

var MySQL MySQLDialect // MySQL

func (dialect MySQLDialect) ReplaceNextPlaceholder(sql string, idx int) string {
	return sql
}

func (dialect MySQLDialect) ReplacePlaceholders(sql string) string {
	return sql
}

func (dialect MySQLDialect) Column(f *schema.Field) string {
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

