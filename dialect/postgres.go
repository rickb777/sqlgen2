package dialect

import (
	"strconv"
	"strings"
	"bytes"
)

const postgresPlaceholders = "$1,$2,$3,$4,$5,$6,$7,$8,$9"

type PostgresDialect struct{}

var Postgres PostgresDialect

func (dialect PostgresDialect) Placeholders(n int) string {
	if n == 0 {
		return ""
	} else if n <= 9 {
		return postgresPlaceholders[:n*3-1]
	}
	buf := &bytes.Buffer{}
	for idx := 1; idx <= n; idx++ {
		if idx > 1 {
			buf.WriteByte(',')
		}
		buf.WriteByte('$')
		buf.WriteString(strconv.Itoa(idx))
	}
	return buf.String()
}

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
