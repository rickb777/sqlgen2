package sqlgen2

// Dialect represents a SQL dialect.
type Dialect interface {
	ReplacePlaceholders(sql string) string
	Placeholders(n int) string
}

var Sqlite MySQLDialect

var DefaultDialect = Mysql
