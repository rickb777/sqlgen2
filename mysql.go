package sqlgen2

import (
	"github.com/rickb777/sqlgen2/schema"
	"strings"
)

const mysqlPlaceholders = "?,?,?,?,?,?,?,?,?,?"

// MySQLDialect provides specialisations needed for working with MySQL.
// This is also compatible with SQLite.
type MySQLDialect struct{}

// type conformance
var _ Dialect = MySQLDialect{}

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
		m := (n * 2) - 1
		return mysqlPlaceholders[:m]
	}
	return strings.Repeat("?,", n-1) + "?"
}

// ReplacePlaceholders converts a string containing '?' placeholders to
// the form used by MySQL and SQLite - i.e. unchanged.
func (dialect MySQLDialect) ReplacePlaceholders(sql string) string {
	return sql
}

// Param returns the i-th parameter.
func (dialect MySQLDialect) Param(i int) string {
	return "?"
}
