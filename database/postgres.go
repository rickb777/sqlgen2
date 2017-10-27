package database

import (
	"strconv"
	"bytes"
	"fmt"
	"github.com/rickb777/sqlgen2/schema"
)

const postgresPlaceholders = "$1,$2,$3,$4,$5,$6,$7,$8,$9"

type PostgresDialect struct{}

var Postgres PostgresDialect

func (dialect PostgresDialect) SDialect() schema.SDialect {
	return schema.New(schema.Postgres)
}

func (dialect PostgresDialect) Placeholders(n int) string {
	if n == 0 {
		return ""
	} else if n <= 9 {
		return postgresPlaceholders[:n*3-1]
	}
	buf := bytes.NewBufferString(postgresPlaceholders)
	for idx := 10; idx <= n; idx++ {
		if idx > 1 {
			buf.WriteByte(',')
		}
		buf.WriteByte('$')
		buf.WriteString(strconv.Itoa(idx))
	}
	return buf.String()
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

func (dialect PostgresDialect) Param(i int) string {
	return fmt.Sprintf("$%d", i+1)
}
