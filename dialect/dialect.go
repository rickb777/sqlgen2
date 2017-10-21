package dialect

import (
	"strconv"
	"strings"
)

// Dialect represents a SQL dialect.
type Dialect interface {
	ReplaceNextPlaceholder(sql string, idx int) string
}

type MySQLDialect struct{}
type PostgresDialect struct{}

var (
	MySQL    MySQLDialect    // MySQL
	SQLite   MySQLDialect    // SQLite (same as MySQL)
	Postgres PostgresDialect // Postgres
)

var DefaultDialect = MySQL // Default dialect

func (dialect MySQLDialect) ReplaceNextPlaceholder(sql string, idx int) string {
	return sql
}

func (dialect PostgresDialect) ReplaceNextPlaceholder(sql string, idx int) string {
	p := "$" + strconv.Itoa(idx)
	return strings.Replace(sql, "?", p, 1)
}
