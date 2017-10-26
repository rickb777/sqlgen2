package dialect

import (
	"github.com/rickb777/sqlgen2/schema"
	"fmt"
	"strings"
)

const mysqlPlaceholders = "?,?,?,?,?,?,?,?,?,?"

type MySQLDialect struct{}

var Mysql MySQLDialect

func (dialect MySQLDialect) SDialect() schema.SDialect {
	return schema.New(schema.Mysql)
}

func (dialect MySQLDialect) Placeholders(n int) string {
	if n == 0 {
		return ""
	} else if n <= 10 {
		m := (n*2)-1
		return mysqlPlaceholders[:m]
	}
	return strings.Repeat("?,", n-1) + "?"
}

func (dialect MySQLDialect) ReplacePlaceholders(sql string) string {
	return sql
}

func (dialect MySQLDialect) Column(f *schema.Field) string {
	switch f.SqlType {
	case schema.INTEGER:
		return "integer"
	case schema.BOOLEAN:
		return "boolean"
	case schema.BLOB:
		return "mediumblob"
	case schema.VARCHAR:
		// assigns an arbitrary size if
		// none is provided.
		size := f.Size
		if size == 0 {
			size = 512
		}
		return fmt.Sprintf("varchar(%d)", size)
	default:
		return ""
	}
}

func (dialect MySQLDialect) Param(i int) string {
	return "?"
}
