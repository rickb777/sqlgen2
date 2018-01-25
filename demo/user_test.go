package demo

import (
	"testing"
	. "github.com/onsi/gomega"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"context"
	"database/sql"
	"database/sql/driver"
	. "fmt"
	"math/big"
	"log"
	"os"
	"github.com/rickb777/sqlgen2"
	"github.com/rickb777/sqlgen2/where"
	"github.com/rickb777/sqlgen2/schema"
	"github.com/rickb777/sqlgen2/require"
	"github.com/rickb777/sqlgen2/model"
	"github.com/rickb777/sqlgen2/constraint"
)

var db *sql.DB
var dialect schema.Dialect

func connect() {
	dbDriver, ok := os.LookupEnv("GO_DRIVER")
	if !ok {
		dbDriver = "sqlite3"
	}
	dialect = schema.PickDialect(dbDriver)
	dsn, ok := os.LookupEnv("GO_DSN")
	if !ok {
		dsn = ":memory:"
	}
	conn, err := sql.Open(dbDriver, dsn)
	if err != nil {
		panic(err)
	}
	db = conn
}

func cleanup() {
	if db != nil {
		db.Close()
		db = nil
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

func xTestCreateTable_sql_syntax(t *testing.T) {
	RegisterTestingT(t)

	cases := []struct {
		dialect  schema.Dialect
		expected string
	}{
		{schema.Sqlite,
`CREATE TABLE IF NOT EXISTS prefix_users (
 uid          integer primary key autoincrement,
 login        text,
 emailaddress text,
 addressid    bigint default null,
 avatar       text default null,
 role         text default null,
 active       boolean,
 admin        boolean,
 fave         text,
 lastupdated  bigint,
 token        text,
 secret       text,
 CONSTRAINT prefix_users_c1 foreign key (addressid) references prefix_addresses (id) on update restrict on delete restrict,
 CONSTRAINT prefix_users_c2 CHECK (role < 3)
)`},
		{schema.Mysql,
`CREATE TABLE IF NOT EXISTS prefix_users (
 uid          bigint primary key auto_increment,
 login        varchar(255),
 emailaddress varchar(255),
 addressid    bigint default null,
 avatar       varchar(255) default null,
 role         varchar(20) default null,
 active       tinyint(1),
 admin        tinyint(1),
 fave         json,
 lastupdated  bigint,
 token        varchar(255),
 secret       varchar(255),
 CONSTRAINT prefix_users_c1 foreign key (addressid) references prefix_addresses (id) on update restrict on delete restrict,
 CONSTRAINT prefix_users_c2 CHECK (role < 3)
) ENGINE=InnoDB DEFAULT CHARSET=utf8`},
		{schema.Postgres,
`CREATE TABLE IF NOT EXISTS prefix_users (
 uid          bigserial primary key,
 login        varchar(255),
 emailaddress varchar(255),
 addressid    bigint default null,
 avatar       varchar(255) default null,
 role         varchar(20) default null,
 active       boolean,
 admin        boolean,
 fave         json,
 lastupdated  bigint,
 token        varchar(255),
 secret       varchar(255),
 CONSTRAINT prefix_users_c1 foreign key (addressid) references prefix_addresses (id) on update restrict on delete restrict,
 CONSTRAINT prefix_users_c2 CHECK (role < 3)
)`},
	}

	for _, c := range cases {
		tbl := NewDbUserTable(model.TableName{Name: "users"}, nil, c.dialect).
			WithPrefix("prefix_").
			AddConstraint(
			constraint.CheckConstraint{"role < 3"})
		s := tbl.createTableSql(true)
		Ω(s).Should(Equal(c.expected), "%s\n%s", c.dialect, s)
	}
}

func TestCreateIndexSql(t *testing.T) {
	RegisterTestingT(t)

	tbl := NewDbUserTable(model.TableName{Name: "users"}, nil, schema.Postgres).WithPrefix("prefix_")
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
		tbl := NewDbUserTable(model.TableName{Name: "users"}, nil, c.d).WithPrefix("prefix_")
		s := tbl.dropDbUserEmailIndexSql(true)
		Ω(s).Should(Equal(c.expected))
	}
}

func TestUpdateFields_ok_using_mock(t *testing.T) {
	RegisterTestingT(t)

	mockDb := mockExecer{RowsAffected: 1}

	tbl := NewDbUserTable(model.TableName{Name: "users"}, mockDb, schema.Mysql)

	n, err := tbl.UpdateFields(require.One, where.NoOp(),
		sqlgen2.Named("EmailAddress", "foo@x.com"),
		sqlgen2.Named("Hash", "abc123"))

	Ω(err).Should(BeNil())
	Ω(n).Should(Equal(int64(1)))
}

func TestUpdateFields_error_using_mock(t *testing.T) {
	RegisterTestingT(t)

	exp := Errorf("foo")
	mockDb := mockExecer{Error: exp}

	tbl := NewDbUserTable(model.TableName{Name: "users"}, mockDb, schema.Mysql)

	_, err := tbl.UpdateFields(nil, where.NoOp(),
		sqlgen2.Named("EmailAddress", "foo@x.com"),
		sqlgen2.Named("Hash", "abc123"))

	Ω(err).Should(Equal(exp))
}

func TestUpdate_ok_using_mock(t *testing.T) {
	RegisterTestingT(t)

	mockDb := mockExecer{RowsAffected: 1}

	tbl := NewDbUserTable(model.TableName{Name: "users"}, mockDb, schema.Mysql)

	n, err := tbl.Update(require.One, &User{})

	Ω(err).Should(BeNil())
	Ω(n).Should(Equal(int64(1)))
}

func TestUpdate_error_using_mock(t *testing.T) {
	RegisterTestingT(t)

	exp := Errorf("foo")
	mockDb := mockExecer{Error: exp}

	tbl := NewDbUserTable(model.TableName{Name: "users"}, mockDb, schema.Mysql)

	_, err := tbl.Update(nil, &User{})

	Ω(err).Should(Equal(exp))
}

//-------------------------------------------------------------------------------------------------

func TestCrud_using_database(t *testing.T) {
	RegisterTestingT(t)
	connect()
	defer cleanup()

	addresses := NewAddressTable(model.TableName{Name: "addresses"}, db, dialect)
	if testing.Verbose() {
		addresses = addresses.WithLogger(log.New(os.Stderr, "", log.LstdFlags))
	}

	users := NewDbUserTable(model.TableName{Name: "users"}, db, dialect)
	if testing.Verbose() {
		users = users.WithLogger(log.New(os.Stderr, "", log.LstdFlags))
	}

	_, err := users.DropTable(true)
	Ω(err).Should(BeNil())

	_, err = addresses.DropTable(true)
	Ω(err).Should(BeNil())

	err = addresses.CreateTableWithIndexes(false)
	Ω(err).Should(BeNil())

	err = users.CreateTableWithIndexes(false)
	Ω(err).Should(BeNil())

	count_empty_table_should_be_zero(t, users)

	user1 := insert_user_should_run_PreInsert(t, users)

	get_user_should_call_PostGet_and_match_expected(t, users, user1)

	get_unknown_user_should_return_nil(t, users, user1)

	must_get_unknown_user_should_return_error(t, users, user1)

	count_known_user_should_return_1(t, users)

	query_unknown_user_should_return_empty_list(t, users)

	select_unknown_user_should_return_empty_list(t, users)

	select_unknown_user_requiring_one_should_return_error(t, users)

	query_one_nullstring_for_user_should_return_valid(t, users)

	query_one_nullstring_for_unknown_should_return_invalid(t, users)

	user2 := select_known_user_requiring_one_should_return_user(t, users)

	update_user_should_call_PreUpdate(t, users, user2)

	delete_one_should_return_1(t, users)

	count_empty_table_should_be_zero(t, users)
}

func count_empty_table_should_be_zero(t *testing.T, tbl DbUserTable) {
	c1, err := tbl.Count(where.NoOp())
	Ω(err).Should(BeNil())
	Ω(c1).Should(Equal(int64(0)))
}

func insert_user_should_run_PreInsert(t *testing.T, tbl DbUserTable) *User {
	user := &User{Login: "user1", EmailAddress: "foo@x.z"}
	user = user.SetRole(UserRole)
	err := tbl.Insert(require.One, user)
	Ω(err).Should(BeNil())
	Ω(user.hash).Should(Equal("PreInsert"))
	return user
}

func get_user_should_call_PostGet_and_match_expected(t *testing.T, tbl DbUserTable, expected *User) {
	user, err := tbl.GetUser(expected.Uid)
	Ω(err).Should(BeNil())
	if user.hash != "PostGet" {
		t.Fatalf("%q", user.hash)
	}
	user.hash = expected.hash
	Ω(user).Should(Equal(expected))
}

func get_unknown_user_should_return_nil(t *testing.T, tbl DbUserTable, expected *User) {
	user, err := tbl.GetUser(expected.Uid + 100000)
	Ω(err).Should(BeNil())
	Ω(user).Should(BeNil())
}

func must_get_unknown_user_should_return_error(t *testing.T, tbl DbUserTable, expected *User) {
	_, err := tbl.MustGetUser(expected.Uid + 100000)
	Ω(err.Error()).Should(Equal("expected to fetch one but got 0"))
}

func count_known_user_should_return_1(t *testing.T, tbl DbUserTable) {
	count, err := tbl.Count(where.Eq("Login", "user1"))
	Ω(err).Should(BeNil())
	Ω(count).Should(BeEquivalentTo(1))
}

func query_unknown_user_should_return_empty_list(t *testing.T, tbl DbUserTable) {
	list, err := tbl.Query(require.None, "select * from {TABLE} where Login=?", "foo")
	Ω(err).Should(BeNil())
	Ω(len(list)).Should(Equal(0))

	u, err := tbl.QueryOne("select * from {TABLE} where Login=?", "foo")
	Ω(err).Should(BeNil())
	Ω(u).Should(BeNil())

	_, err = tbl.MustQueryOne("select * from {TABLE} where Login=?", "foo")
	Ω(err.Error()).Should(Equal("expected to fetch one but got 0"))
}

func select_unknown_user_should_return_empty_list(t *testing.T, tbl DbUserTable) {
	list, err := tbl.Select(require.None, where.Eq("Login", "unknown"), nil)
	Ω(err).Should(BeNil())
	Ω(len(list)).Should(Equal(0))
}

func select_unknown_user_requiring_one_should_return_error(t *testing.T, tbl DbUserTable) {
	list, err := tbl.Select(require.None, where.Eq("Login", "unknown"), nil)
	Ω(err).Should(BeNil())
	Ω(len(list)).Should(Equal(0))

	_, err = tbl.Select(require.One, where.Eq("Login", "unknown"), nil)
	Ω(err.Error()).Should(Equal("expected to fetch one but got 0"))
}

func query_one_nullstring_for_user_should_return_valid(t *testing.T, tbl DbUserTable) {
	s, err := tbl.QueryOneNullString("select EmailAddress from {TABLE} where Login=?", "user1")
	Ω(err).Should(BeNil())
	Ω(s.Valid).Should(BeTrue())
	Ω(s.String).Should(Equal("foo@x.z"))

	s, err = tbl.MustQueryOneNullString("select EmailAddress from {TABLE} where Login=?", "user1")
	Ω(err).Should(BeNil())
	Ω(s.Valid).Should(BeTrue())
	Ω(s.String).Should(Equal("foo@x.z"))
}

func query_one_nullstring_for_unknown_should_return_invalid(t *testing.T, tbl DbUserTable) {
	s, err := tbl.QueryOneNullString("select EmailAddress from {TABLE} where Login=?", "foo")
	Ω(err).Should(BeNil())
	Ω(s.Valid).Should(BeFalse())

	_, err = tbl.MustQueryOneNullString("select EmailAddress from {TABLE} where Login=?", "foo")
	Ω(err.Error()).Should(Equal("expected to fetch one but got 0"))
	Ω(s.Valid).Should(BeFalse())
}

func select_known_user_requiring_one_should_return_user(t *testing.T, tbl DbUserTable) *User {
	list, err := tbl.Select(require.One, where.Eq("Login", "user1"), nil)
	Ω(err).Should(BeNil())
	Ω(len(list)).Should(Equal(1))
	return list[0]
}

func update_user_should_call_PreUpdate(t *testing.T, tbl DbUserTable, user *User) {
	user.EmailAddress = "bah0@zzz.com"

	n, err := tbl.Update(require.One, user)
	Ω(err).Should(BeNil())
	Ω(n).Should(BeEquivalentTo(1))
	Ω(user.hash).Should(Equal("PreUpdate"))

	ss, err := tbl.SliceEmailaddress(require.One, where.Eq("Uid", user.Uid), nil)
	Ω(err).Should(BeNil())
	Ω(len(ss)).Should(Equal(1))
	Ω(ss[0]).Should(Equal("bah0@zzz.com"))
}

func delete_one_should_return_1(t *testing.T, tbl DbUserTable) {
	n, err := tbl.Delete(require.One, where.Eq("Login", "user1"))
	Ω(err).Should(BeNil())
	Ω(n).Should(BeEquivalentTo(1))
}

//-------------------------------------------------------------------------------------------------

func TestMultiSelect_using_database(t *testing.T) {
	RegisterTestingT(t)
	connect()
	defer cleanup()

	tbl := NewDbUserTable(model.TableName{Name: "users"}, db, dialect)
	if testing.Verbose() {
		tbl = tbl.WithLogger(log.New(os.Stderr, "", log.LstdFlags))
	}

	err := tbl.CreateTableWithIndexes(true)
	Ω(err).Should(BeNil())

	const n = 3

	var users []*User
	user0 := &User{Login: "user0", EmailAddress: "foo0@x.z"}
	// fave, avatar are null
	users = append(users, user0)

	for i := 1; i <= n; i++ {
		fave := big.NewInt(int64(i))
		user := &User{Fave: fave}
		user = user.SetRole(UserRole)
		user = user.SetLogin(Sprintf("user%d", i))
		user = user.SetEmailAddress(Sprintf("foo%d@x.z", i))
		user = user.SetAvatar(Sprintf("user%d-avatar%d", i, i))
		users = append(users, user)
	}

	err = tbl.Insert(require.All, users...)
	Ω(err).Should(BeNil())

	list, err := tbl.Select(nil, where.NotEq("Login", "nobody"), where.OrderBy("Login").Desc())
	Ω(err).Should(BeNil())
	Ω(len(list)).Should(Equal(n + 1))
	for i := 0; i <= n; i++ {
		users[n-i].hash = "PostGet"
		Ω(list[i]).Should(Equal(users[n-i]))
	}
}

func TestGetters_using_database(t *testing.T) {
	RegisterTestingT(t)
	connect()
	defer cleanup()

	tbl := NewDbUserTable(model.TableName{Name: "users"}, db, dialect)
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

	err = tbl.Insert(require.All, list...)
	Ω(err).Should(BeNil())

	logins, err := tbl.SliceLogin(require.Exactly(n), where.NoOp(), where.OrderBy("login"))
	Ω(err).Should(BeNil())
	Ω(len(logins)).Should(Equal(n))

	for i := 0; i < n; i++ {
		exp := Sprintf("user%02d", i)
		Ω(logins[i]).Should(Equal(exp))
	}
}

func TestBulk_delete_using_database(t *testing.T) {
	RegisterTestingT(t)
	connect()
	defer cleanup()

	tbl := NewDbUserTable(model.TableName{Name: "users"}, db, dialect)
	if testing.Verbose() {
		tbl = tbl.WithLogger(log.New(os.Stderr, "", log.LstdFlags))
	}

	err := tbl.CreateTableWithIndexes(true)
	Ω(err).Should(BeNil())

	err = tbl.Truncate(true)
	Ω(err).Should(BeNil())

	const n = 17

	list := make([]*User, n)
	for i := 0; i < n; i++ {
		list[i] = user(i)
	}

	err = tbl.Insert(require.All, list...)
	Ω(err).Should(BeNil())

	ids := make([]int64, n)
	for i := 0; i < n; i++ {
		ids[i] = list[i].Uid
	}

	j, err := tbl.DeleteUsers(require.All, ids...)
	Ω(err).Should(BeNil())
	Ω(j).Should(BeEquivalentTo(n))
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
