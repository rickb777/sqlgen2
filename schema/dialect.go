package schema

import "fmt"

type DialectId int

const (
	Sqlite   DialectId = iota
	Postgres
	Mysql
)

var Dialects []string

func init() {
	Dialects = make([]string, len(AllDialectIds))
	for i, d := range AllDialectIds {
		Dialects[i] = d.String()
	}
}

func New(dialect DialectId) SDialect {
	switch dialect {
	case Postgres:
		return newPostgres()
	case Mysql:
		return newMySQL()
	default:
		return newSQLite()
	}
}

func (f *Field) AsColumn(dialect DialectId) string {
	switch dialect {
	case Mysql:
		return mysqlColumn(f)
	case Postgres:
		return postgresColumn(f)
	case Sqlite:
		return sqliteColumn(f)
	}
	panic(fmt.Errorf("%s is not implemented", dialect))
}
