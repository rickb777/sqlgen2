package demo

import (
	"testing"
	"github.com/rickb777/sqlgen2/dialect"
	"github.com/rickb777/sqlgen2/where"
	"reflect"
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

func TestUpdateFields_postgres(t *testing.T) {
	cases := []struct{
		d dialect.Dialect
		expected string
	}{
		{dialect.Mysql, `UPDATE prefix_users SET Email=?, Hash=? WHERE Email ISNULL`},
		{dialect.Postgres, `UPDATE prefix_users SET Email=$1, Hash=$2 WHERE Email ISNULL`},
	}

	for _, c := range cases {
		tbl := DbUserTable{"prefix_", "users", nil, c.d}

		sql, args := tbl.updateFields(where.Null("Email"),
			where.Field{"Email", "foo@x.com"}, where.Field{"Hash", "abc123"})

		if sql != c.expected {
			t.Errorf("expected %s\ngot %s", c.expected, sql)
		}
		if !reflect.DeepEqual(args, []interface{}{"foo@x.com", "abc123"}) {
			t.Errorf("got %+v", args)
		}
	}
}
