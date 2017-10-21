package dialect

import (
	"strconv"
	"strings"
)

type PostgresDialect struct{}

var Postgres PostgresDialect // Postgres

func (dialect PostgresDialect) ReplaceNextPlaceholder(sql string, idx int) string {
	p := "$" + strconv.Itoa(idx)
	return strings.Replace(sql, "?", p, 1)
}
