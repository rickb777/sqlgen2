package demo

import (
	"testing"
	"github.com/rickb777/sqlgen2"
	"github.com/rickb777/sqlgen2/where"
	_ "github.com/mattn/go-sqlite3"
	"reflect"
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"syscall"
	"github.com/rickb777/sqlgen2/schema"
)

const dbDriver = "sqlite3"
const dsn = "./test.db"

var db *sql.DB

func connect() *sql.DB {
	conn, err := sql.Open(dbDriver, dsn)
	if err != nil {
		panic(err)
	}
	db = conn
	return conn
}

func cleanup() {
	if db != nil {
		db.Close()
		syscall.Unlink(dsn)
	}
}

func TestCreateTable_postgres(t *testing.T) {
	tbl := NewDbUserTable("users", nil, schema.Postgres).WithPrefix("prefix_")
	sql := tbl.createTableSql(true)
	expected := `
CREATE TABLE IF NOT EXISTS prefix_users (
 uid          bigserial primary key,
 login        varchar(512),
 emailaddress varchar(512),
 avatar       varchar(512),
 active       boolean,
 admin        boolean,
 fave         json,
 lastupdated  bigint,
 token        varchar(512),
 secret       varchar(512),
 hash         varchar(512)
)
`
	if sql != expected {
		t.Errorf("got %s", sql)
	}
}

func TestCreateIndexSql(t *testing.T) {
	tbl := NewDbUserTable("users", nil, schema.Postgres).WithPrefix("prefix_")
	sql := tbl.createDbUserEmailIndexSql("IF NOT EXISTS ")
	expected := `CREATE UNIQUE INDEX IF NOT EXISTS prefix_user_email ON prefix_users (emailaddress)`
	if sql != expected {
		t.Errorf("got %s", sql)
	}
}

func TestDropIndexSql(t *testing.T) {
	tbl := NewDbUserTable("users", nil, schema.Postgres).WithPrefix("prefix_")
	sql := tbl.dropDbUserEmailIndexSql("IF EXISTS ", " ON prefix_users")
	expected := `DROP INDEX IF EXISTS prefix_user_email ON prefix_users`
	if sql != expected {
		t.Errorf("got %s", sql)
	}
}

func TestUpdateFieldsSql(t *testing.T) {
	cases := []struct {
		d        schema.Dialect
		expected string
	}{
		{schema.Sqlite, `UPDATE prefix_users SET EmailAddress=?, Hash=? WHERE EmailAddress ISNULL`},
		{schema.Mysql, `UPDATE prefix_users SET EmailAddress=?, Hash=? WHERE EmailAddress ISNULL`},
		{schema.Postgres, `UPDATE prefix_users SET EmailAddress=$1, Hash=$2 WHERE EmailAddress ISNULL`},
	}

	for _, c := range cases {
		tbl := NewDbUserTable("users", nil, c.d).WithPrefix("prefix_")

		sql, args := tbl.updateFields(where.Null("EmailAddress"),
			sqlgen2.Named("EmailAddress", "foo@x.com"), sqlgen2.Named("Hash", "abc123"))

		if sql != c.expected {
			t.Errorf("expected %s\ngot %s", c.expected, sql)
		}
		if !reflect.DeepEqual(args, []interface{}{"foo@x.com", "abc123"}) {
			t.Errorf("got %+v", args)
		}
	}
}

func TestUpdateFields_ok(t *testing.T) {
	mockDb := mockExecer{RowsAffected: 1}

	tbl := NewDbUserTable("users", mockDb, schema.Mysql)

	n, err := tbl.UpdateFields(where.NoOp(),
		sqlgen2.Named("EmailAddress", "foo@x.com"),
		sqlgen2.Named("Hash", "abc123"))

	if err != nil {
		t.Errorf("%v", err)
	} else if n != 1 {
		t.Errorf("expected 1\ngot %d", n)
	}
}

func TestUpdateFields_error(t *testing.T) {
	exp := fmt.Errorf("foo")
	mockDb := mockExecer{Error: exp}

	tbl := NewDbUserTable("users", mockDb, schema.Mysql)

	_, err := tbl.UpdateFields(where.NoOp(),
		sqlgen2.Named("EmailAddress", "foo@x.com"),
		sqlgen2.Named("Hash", "abc123"))

	if err != exp {
		t.Errorf("expected an error, got %v", err)
	}
}

func TestUpdate_ok(t *testing.T) {
	mockDb := mockExecer{RowsAffected: 1}

	tbl := NewDbUserTable("users", mockDb, schema.Mysql)

	n, err := tbl.Update(&User{})

	if err != nil {
		t.Errorf("%v", err)
	} else if n != 1 {
		t.Errorf("expected 1\ngot %d", n)
	}
}

func TestUpdate_error(t *testing.T) {
	exp := fmt.Errorf("foo")
	mockDb := mockExecer{Error: exp}

	tbl := NewDbUserTable("users", mockDb, schema.Mysql)

	_, err := tbl.Update(&User{})

	if err != exp {
		t.Errorf("expected an error, got %v", err)
	}
}

// inserts are harder to test because they use prepared statements

func TestCrudUsingSqlite(t *testing.T) {
	defer cleanup()

	tbl := NewDbUserTable("users", connect(), schema.Sqlite)

	err := tbl.CreateTableWithIndexes(false)
	if err != nil {
		t.Fatalf("%v", err)
	}

	c1, err := tbl.Count(where.NoOp())
	if err != nil {
		t.Fatalf("%v", err)
	}
	if c1 != 0 {
		t.Errorf("expected 0, got %d", c1)
	}

	err = tbl.Insert(&User{Login: "user1", EmailAddress: "foo@x.z"})
	if err != nil {
		t.Fatalf("%v", err)
	}

	c1, err = tbl.Count(where.Eq("Login", "user1"))
	if err != nil {
		t.Fatalf("%v", err)
	}
	if c1 != 1 {
		t.Errorf("expected 1, got %d", c1)
	}

	ul, err := tbl.Select(where.Eq("Login", "user1"), "")
	if err != nil {
		t.Fatalf("%v", err)
	}
	if len(ul) != 1 || ul[0].Login != "user1" {
		t.Errorf("expected 1, got %v", ul)
	}

	u1, err := tbl.SelectOne(where.Eq("Login", "user1"), "")
	if err != nil {
		t.Fatalf("%v", err)
	}
	if u1 == nil || u1.Login != "user1" {
		t.Errorf("expected 1, got %v", ul)
	}

	n, err := tbl.Delete(where.Eq("Login", "user1"))
	if err != nil {
		t.Fatalf("%v", err)
	}
	if n != 1 {
		t.Errorf("expected 1, got %d", n)
	}
}

//-------------------------------------------------------------------------------------------------

type mockExecer struct {
	RowsAffected driver.RowsAffected
	Error        error
}

func (m mockExecer) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return m.RowsAffected, m.Error
}

func (m mockExecer) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return nil, nil
}

func (m mockExecer) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return nil, nil
}

func (m mockExecer) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return nil
}
