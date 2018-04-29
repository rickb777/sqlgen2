package schema

import (
	"strings"
	"io"
)

type Dialect interface {
	Index() int
	String() string
	Alias() string

	TableDDL(*TableDescription) string
	FieldDDL(w io.Writer, field *Field, comma string) string
	TruncateDDL(tableName string, force bool) []string
	CreateTableSettings() string
	FieldAsColumn(*Field) string
	InsertHasReturningPhrase() bool

	SplitAndQuote(csv string) string
	Quote(string) string
	QuoteW(w io.Writer, identifier string)
	QuoteWithPlaceholder(w io.Writer, identifier string, j int)
	ReplacePlaceholders(sql string, args []interface{}) string
	Placeholder(name string, j int) string
	Placeholders(n int) string
}

//-------------------------------------------------------------------------------------------------

const (
	SqliteIndex = iota
	MysqlIndex
	PostgresIndex
)

//-------------------------------------------------------------------------------------------------

// AllDialects lists all currently-supported dialects.
var AllDialects = []Dialect{Sqlite, Mysql, Postgres}

// PickDialect finds a dialect that matches by name, ignoring letter case.
// It returns nil if not found.
func PickDialect(name string) Dialect {
	for _, d := range AllDialects {
		if strings.EqualFold(name, d.String()) || strings.EqualFold(name, d.Alias()) {
			return d
		}
	}
	return nil
}
