package sqlgen2

import (
	"github.com/rickb777/sqlgen2/schema"
	"fmt"
	"strings"
)

const mysqlPlaceholders = "?,?,?,?,?,?,?,?,?,?"

// MySQLDialect provides specialisations needed for working with MySQL.
// This is also compatible with SQLite.
type MySQLDialect struct{}

// Mysql implements specialisations needed for working with MySQL.
// This is also compatible with SQLite.
var Mysql MySQLDialect

func (dialect MySQLDialect) SDialect() schema.SDialect {
	return schema.New(schema.Mysql)
}

// Placeholders returns a string containing the requested number of placeholders
// in the form used by MySQL and SQLite.
func (dialect MySQLDialect) Placeholders(n int) string {
	if n == 0 {
		return ""
	} else if n <= 10 {
		m := (n*2)-1
		return mysqlPlaceholders[:m]
	}
	return strings.Repeat("?,", n-1) + "?"
}

// ReplacePlaceholders converts a string containing '?' placeholders to
// the form used by MySQL and SQLite - i.e. unchanged.
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
		size := f.Tags.Size
		if size == 0 {
			size = 512
		}
		return fmt.Sprintf("varchar(%d)", size)
	default:
		return ""
	}
}

// Param returns the i-th parameter.
func (dialect MySQLDialect) Param(i int) string {
	return "?"
}
