package dialect

import (
	"strconv"
	"strings"
	"bytes"
)

type PostgresDialect struct{}

var Postgres PostgresDialect // Postgres

func (dialect PostgresDialect) ReplaceNextPlaceholder(sql string, idx int) string {
	p := "$" + strconv.Itoa(idx)
	return strings.Replace(sql, "?", p, 1)
}

func (dialect PostgresDialect) ReplacePlaceholders(sql string) string {
	buf := &bytes.Buffer{}
	idx := 1
	for _, r := range sql {
		if r == '?' {
			buf.WriteByte('$')
			buf.WriteString(strconv.Itoa(idx))
			idx++
		} else {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}
