package demo

import (
	"testing"
	"github.com/rickb777/sqlgen2/db"
	"github.com/rickb777/sqlgen2/where"
	"reflect"
	"context"
)

func TestCreateTable_postgres(t *testing.T) {
	tbl := DbUserTable{"prefix_", "users", nil, context.Background(), db.Postgres}
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
	tbl := DbUserTable{"prefix_", "users", nil, context.Background(), db.Postgres}
	sql := tbl.createDbUserEmailIndexSql("IF NOT EXISTS ")
	expected := `
CREATE UNIQUE INDEX IF NOT EXISTS user_email ON prefix_users (email)
`
	if sql != expected {
		t.Errorf("got %s", sql)
	}
}

func TestUpdateFields_postgres(t *testing.T) {
	cases := []struct{
		d db.Dialect
		expected string
	}{
		{db.Mysql, `UPDATE prefix_users SET Email=?, Hash=? WHERE Email ISNULL`},
		{db.Postgres, `UPDATE prefix_users SET Email=$1, Hash=$2 WHERE Email ISNULL`},
	}

	for _, c := range cases {
		tbl := DbUserTable{"prefix_", "users", nil, context.Background(), c.d}

		sql, args := tbl.updateFields(where.Null("Email"),
			db.Named("Email", "foo@x.com"), db.Named("Hash", "abc123"))

		if sql != c.expected {
			t.Errorf("expected %s\ngot %s", c.expected, sql)
		}
		if !reflect.DeepEqual(args, []interface{}{"foo@x.com", "abc123"}) {
			t.Errorf("got %+v", args)
		}
	}
}
