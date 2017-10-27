package db

import "github.com/rickb777/sqlgen2/schema"

// Dialect represents a SQL dialect.
type Dialect interface {
	ReplacePlaceholders(sql string) string
	Placeholders(n int) string
	SDialect() schema.SDialect
}

var Sqlite MySQLDialect

var DefaultDialect = Mysql
