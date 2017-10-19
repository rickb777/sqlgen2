package sqlgen

import "github.com/rickb777/sqlgen/schema"

type Dialect interface {
	Table(*schema.Table) string
	Index(*schema.Table, *schema.Index) string
	Column(*schema.Field) string
	Insert(*schema.Table) string
	Update(*schema.Table, []*schema.Field) string
	Delete(*schema.Table, []*schema.Field) string
	Param(int) string
	Params(int, int) []string
	ColumnParams(t *schema.Table, withAuto bool) string
	Token(int) string
}

var (
	SQLite   Dialect = newSQLite()
	Postgres Dialect = newPostgres()
	Mysql    Dialect = newMySQL()
)
