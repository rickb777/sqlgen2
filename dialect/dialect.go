package dialect

// Dialect represents a SQL dialect.
type Dialect interface {
	ReplacePlaceholders(sql string) string
	Placeholders(n int) string
}

var SQLite MySQLDialect // SQLite (same as MySQL)

var DefaultDialect = MySQL
