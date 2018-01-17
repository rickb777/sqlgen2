package demo

import (
	"testing"
	. "github.com/onsi/gomega"
	"github.com/rickb777/sqlgen2"
	"github.com/rickb777/sqlgen2/where"
	_ "github.com/mattn/go-sqlite3"
	"reflect"
	"context"
	"database/sql"
	"database/sql/driver"
	. "fmt"
	"syscall"
	"github.com/rickb777/sqlgen2/schema"
	"math/big"
	"log"
	"os"
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

func user(i int) *User {
	return &User{
		Login:        Sprintf("user%02d", i),
		EmailAddress: Sprintf("foo%d@x.z", i),
		Active:       true,
		Fave:         big.NewInt(int64(i)),
	}
}

func TestCreateTable_postgres(t *testing.T) {
	RegisterTestingT(t)

	tbl := NewDbUserTable(sqlgen2.TableName{Name: "users"}, nil, schema.Postgres).WithPrefix("prefix_")
	sql := tbl.createTableSql(true)
	expected := `
CREATE TABLE IF NOT EXISTS prefix_users (
 uid          bigserial primary key,
 login        varchar(255),
 emailaddress varchar(255),
 avatar       varchar(255) default null,
 active       boolean,
 admin        boolean,
 fave         json,
 lastupdated  bigint,
 token        varchar(255),
 secret       varchar(255)
)
`
	if sql != expected {
		t.Errorf("got %s", sql)
	}
}

func TestCreateIndexSql(t *testing.T) {
	RegisterTestingT(t)

	tbl := NewDbUserTable(sqlgen2.TableName{Name: "users"}, nil, schema.Postgres).WithPrefix("prefix_")
	s := tbl.createDbUserEmailIndexSql("IF NOT EXISTS ")
	expected := `CREATE UNIQUE INDEX IF NOT EXISTS prefix_user_email ON prefix_users (emailaddress)`
	Ω(s).Should(Equal(expected))
}

func TestDropIndexSql(t *testing.T) {
	RegisterTestingT(t)

	cases := []struct {
		d        schema.Dialect
		expected string
	}{
		{schema.Sqlite, `DROP INDEX IF EXISTS prefix_user_email`},
		{schema.Mysql, `DROP INDEX prefix_user_email ON prefix_users`},
		{schema.Postgres, `DROP INDEX IF EXISTS prefix_user_email`},
	}

	for _, c := range cases {
		tbl := NewDbUserTable(sqlgen2.TableName{Name: "users"}, nil, c.d).WithPrefix("prefix_")
		s := tbl.dropDbUserEmailIndexSql(true)
		Ω(s).Should(Equal(c.expected))
	}
}

func TestUpdateFieldsSql(t *testing.T) {
	RegisterTestingT(t)

	cases := []struct {
		d        schema.Dialect
		expected string
	}{
		{schema.Sqlite, `UPDATE prefix_users SET EmailAddress=?, Hash=? WHERE EmailAddress IS NULL`},
		{schema.Mysql, `UPDATE prefix_users SET EmailAddress=?, Hash=? WHERE EmailAddress IS NULL`},
		{schema.Postgres, `UPDATE prefix_users SET EmailAddress=$1, Hash=$2 WHERE EmailAddress IS NULL`},
	}

	for _, c := range cases {
		tbl := NewDbUserTable(sqlgen2.TableName{Name: "users"}, nil, c.d).WithPrefix("prefix_")

		s, args := tbl.updateFields(where.Null("EmailAddress"),
			sqlgen2.Named("EmailAddress", "foo@x.com"), sqlgen2.Named("Hash", "abc123"))

		Ω(s).Should(Equal(c.expected))
		Ω(args).Should(Equal([]interface{}{"foo@x.com", "abc123"}))
	}
}

func TestUpdateFields_ok(t *testing.T) {
	RegisterTestingT(t)

	mockDb := mockExecer{RowsAffected: 1}

	tbl := NewDbUserTable(sqlgen2.TableName{Name: "users"}, mockDb, schema.Mysql)

	n, err := tbl.UpdateFields(where.NoOp(),
		sqlgen2.Named("EmailAddress", "foo@x.com"),
		sqlgen2.Named("Hash", "abc123"))

	Ω(err).Should(BeNil())
	Ω(n).Should(Equal(int64(1)))
}

func TestUpdateFields_error(t *testing.T) {
	RegisterTestingT(t)

	exp := Errorf("foo")
	mockDb := mockExecer{Error: exp}

	tbl := NewDbUserTable(sqlgen2.TableName{Name: "users"}, mockDb, schema.Mysql)

	_, err := tbl.UpdateFields(where.NoOp(),
		sqlgen2.Named("EmailAddress", "foo@x.com"),
		sqlgen2.Named("Hash", "abc123"))

	Ω(err).Should(Equal(exp))
}

func TestUpdate_ok(t *testing.T) {
	RegisterTestingT(t)

	mockDb := mockExecer{RowsAffected: 1}

	tbl := NewDbUserTable(sqlgen2.TableName{Name: "users"}, mockDb, schema.Mysql)

	n, err := tbl.Update(&User{})

	Ω(err).Should(BeNil())
	Ω(n).Should(Equal(int64(1)))
}

func TestUpdate_error(t *testing.T) {
	RegisterTestingT(t)

	exp := Errorf("foo")
	mockDb := mockExecer{Error: exp}

	tbl := NewDbUserTable(sqlgen2.TableName{Name: "users"}, mockDb, schema.Mysql)

	_, err := tbl.Update(&User{})

	Ω(err).Should(Equal(exp))
}

// inserts are harder to test because they use prepared statements

func TestCrudUsingSqlite(t *testing.T) {
	RegisterTestingT(t)
	defer cleanup()

	tbl := NewDbUserTable(sqlgen2.TableName{Name: "users"}, connect(), schema.Sqlite)
	if testing.Verbose() {
		tbl = tbl.WithLogger(log.New(os.Stderr, "", log.LstdFlags))
	}

	err := tbl.CreateTableWithIndexes(false)
	Ω(err).Should(BeNil())

	c1, err := tbl.Count(where.NoOp())
	Ω(err).Should(BeNil())
	if c1 != 0 {
		t.Errorf("expected 0, got %d", c1)
	}

	user1 := &User{Login: "user1", EmailAddress: "foo@x.z"}
	err = tbl.Insert(user1)
	Ω(err).Should(BeNil())
	if user1.hash != "PreInsert" {
		t.Fatalf("%q", user1.hash)
	}

	user2, err := tbl.GetUser(user1.Uid)
	Ω(err).Should(BeNil())
	if user2.hash != "PostGet" {
		t.Fatalf("%q", user2.hash)
	}
	user1.hash = user2.hash
	if !reflect.DeepEqual(user1, user2) {
		t.Errorf("expected %#v, got %#v", user1, user2)
	}

	user3, err := tbl.GetUser(user1.Uid + 100000)
	Ω(err).Should(BeNil())
	if user3 != nil {
		t.Fatalf("%v", user3)
	}

	c1, err = tbl.Count(where.Eq("Login", "user1"))
	Ω(err).Should(BeNil())
	if c1 != 1 {
		t.Errorf("expected 1, got %d", c1)
	}

	ul1, err := tbl.Select(where.Eq("Login", "unknown"), nil)
	Ω(err).Should(BeNil())
	if len(ul1) != 0 {
		t.Errorf("expected 0, got %v", ul1)
	}

	ul2, err := tbl.Select(where.Eq("Login", "user1"), nil)
	Ω(err).Should(BeNil())
	if len(ul2) != 1 || ul2[0].Login != "user1" {
		t.Errorf("expected 1, got %v", ul2)
	}

	ul2[0].EmailAddress = "bah0@zzz.com"

	n, err := tbl.Update(ul2[0])
	Ω(err).Should(BeNil())
	if n != 1 {
		t.Errorf("expected 1, got %v", n)
	}
	if ul2[0].hash != "PreUpdate" {
		t.Fatalf("%q", ul2[0].hash)
	}

	u1, err := tbl.SelectOne(where.Eq("Login", "user1"), nil)
	Ω(err).Should(BeNil())
	if u1 == nil || u1.Login != "user1" {
		t.Errorf("expected 1, got %v", ul2)
	}

	n, err = tbl.Delete(where.Eq("Login", "user1"))
	Ω(err).Should(BeNil())
	if n != 1 {
		t.Errorf("expected 1, got %d", n)
	}
}

func TestMultiSelectUsingSqlite(t *testing.T) {
	RegisterTestingT(t)
	defer cleanup()

	tbl := NewDbUserTable(sqlgen2.TableName{Name: "users"}, connect(), schema.Sqlite)
	if testing.Verbose() {
		tbl = tbl.WithLogger(log.New(os.Stderr, "", log.LstdFlags))
	}

	err := tbl.CreateTableWithIndexes(false)
	Ω(err).Should(BeNil())

	const n = 3

	var users []*User
	user0 := &User{Login: "user0", EmailAddress: "foo0@x.z"}
	// fave, avatar are null
	users = append(users, user0)

	for i := 1; i <= n; i++ {
		fave := big.NewInt(int64(i))
		user := &User{Fave: fave}
		user = user.SetLogin(Sprintf("user%d", i))
		user = user.SetEmailAddress(Sprintf("foo%d@x.z", i))
		user = user.SetAvatar(Sprintf("user%d-avatar%d", i, i))
		users = append(users, user)
	}

	err = tbl.Insert(users...)
	Ω(err).Should(BeNil())

	list, err := tbl.Select(where.NotEq("Login", "nobody"), where.OrderBy("Login").Desc())
	Ω(err).Should(BeNil())
	Ω(len(list)).Should(Equal(n + 1))
	for i := 0; i <= n; i++ {
		users[n-i].hash = "PostGet"
		Ω(list[i]).Should(Equal(users[n-i]))
	}
}

func TestGettersUsingSqlite(t *testing.T) {
	RegisterTestingT(t)
	defer cleanup()

	tbl := NewDbUserTable(sqlgen2.TableName{Name: "users"}, connect(), schema.Sqlite)
	if testing.Verbose() {
		tbl = tbl.WithLogger(log.New(os.Stderr, "", log.LstdFlags))
	}

	err := tbl.CreateTableWithIndexes(false)
	Ω(err).Should(BeNil())

	err = tbl.Truncate(true)
	Ω(err).Should(BeNil())

	const n = 20

	list := make([]*User, n)
	for i := 0; i < n; i++ {
		list[i] = user(i)
	}

	err = tbl.Insert(list...)
	Ω(err).Should(BeNil())

	logins, err := tbl.SliceLogin(where.NoOp(), where.OrderBy("login"))
	Ω(err).Should(BeNil())
	if len(logins) != n {
		t.Errorf("expected 20, got %d", len(logins))
	}
	for i := 0; i < n; i++ {
		exp := Sprintf("user%02d", i)
		if logins[i] != exp {
			t.Errorf("expected %s, got %s", exp, logins[i])
		}
	}
}

func TestBulkDeleteUsingSqlite(t *testing.T) {
	RegisterTestingT(t)
	defer cleanup()

	tbl := NewDbUserTable(sqlgen2.TableName{Name: "users"}, connect(), schema.Sqlite)
	if testing.Verbose() {
		tbl = tbl.WithLogger(log.New(os.Stderr, "", log.LstdFlags))
	}

	err := tbl.CreateTableWithIndexes(false)
	Ω(err).Should(BeNil())

	err = tbl.Truncate(true)
	Ω(err).Should(BeNil())

	const n = 17

	list := make([]*User, n)
	for i := 0; i < n; i++ {
		list[i] = user(i)
	}

	err = tbl.Insert(list...)
	Ω(err).Should(BeNil())

	ids := make([]int64, n)
	for i := 0; i < n; i++ {
		ids[i] = list[i].Uid
	}

	j, err := tbl.DeleteUsers(ids...)
	Ω(err).Should(BeNil())
	if j != n {
		t.Errorf("Got %d", j)
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
