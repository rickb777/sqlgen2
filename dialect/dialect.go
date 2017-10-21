package dialect

// Dialect represents a SQL dialect.
type Dialect interface {
	ReplaceNextPlaceholder(sql string, idx int) string
}

var SQLite MySQLDialect // SQLite (same as MySQL)

var DefaultDialect = MySQL
