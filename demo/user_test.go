package demo

import (
	"testing"
	"github.com/rickb777/sqlgen2/dialect"
)

func TestCreateTable_postgres(t *testing.T) {
	tbl := DbUserTable{"prefix_", "users", nil, dialect.Postgres}
	sql := tbl.createTableSql(true)
	expected := `
CREATE TABLE IF NOT EXISTS prefix_users (
 uid    bigserial primary key ,
 login  varchar(512),
 email  varchar(512),
 avatar varchar(512),
 active boolean,
 admin  boolean,
 token  varchar(512),
 secret varchar(512),
 hash   varchar(512)
)
`
	if sql != expected {
		t.Errorf("got %s", sql)
	}
}

func TestCreateIndex_postgres(t *testing.T) {
	tbl := DbUserTable{"prefix_", "users", nil, dialect.Postgres}
	sql := tbl.createDbUserEmailIndexSql("IF NOT EXISTS ")
	expected := `
CREATE UNIQUE INDEX IF NOT EXISTS user_email ON prefix_users (email)
`
	if sql != expected {
		t.Errorf("got %s", sql)
	}
}
