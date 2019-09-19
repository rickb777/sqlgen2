package demo

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	"github.com/rickb777/sqlapi"
	"github.com/rickb777/sqlapi/constraint"
	"github.com/rickb777/sqlapi/dialect"
	"github.com/rickb777/sqlapi/require"
	"github.com/rickb777/sqlapi/support"
	"github.com/rickb777/where"
	"github.com/rickb777/where/quote"
	"github.com/spf13/cast"
	"io"
	"log"
	"math/big"
	"os"
	"strings"
	"testing"
)

// Environment:
// GO_DRIVER  - the driver (sqlite3, mysql, postgres, pgx)
// GO_QUOTER  - the identifier quoter (ansi, mysql, none)
// GO_DSN     - the database DSN
// GO_VERBOSE - true for query logging

var verbose = false

func skipIfNoPostgresDB(t *testing.T, di dialect.Dialect) {
	if (di.Index() == dialect.PostgresIndex || di.Index() == dialect.PgxIndex) && os.Getenv("PGHOST") == "" {
		t.Skip()
	}
}

func connect(t *testing.T) (*sql.DB, dialect.Dialect) {
	dbDriver, ok := os.LookupEnv("GO_DRIVER")
	if !ok {
		dbDriver = "sqlite3"
	}

	di := dialect.PickDialect(dbDriver) //.WithQuoter(dialect.NoQuoter)
	quoter, ok := os.LookupEnv("GO_QUOTER")
	if ok {
		switch strings.ToLower(quoter) {
		case "ansi":
			di = di.WithQuoter(quote.AnsiQuoter)
		case "mysql":
			di = di.WithQuoter(quote.MySqlQuoter)
		case "none":
			di = di.WithQuoter(quote.NoQuoter)
		default:
			t.Fatalf("Warning: unrecognised quoter %q.\n", quoter)
		}
	}

	skipIfNoPostgresDB(t, di)

	dsn, ok := os.LookupEnv("GO_DSN")
	if !ok {
		dsn = "file::memory:?mode=memory&cache=shared"
	}

	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		t.Fatalf("Error: Unable to connect to %s (%v); test is only partially complete.\n\n", dbDriver, err)
	}

	err = db.Ping()
	if err != nil {
		t.Fatalf("Error: Unable to ping %s (%v); test is only partially complete.\n\n", dbDriver, err)
	}

	fmt.Printf("Successfully connected to %s.\n", dbDriver)
	return db, di
}

func newDatabase(t *testing.T) sqlapi.Database {
	db, di := connect(t)

	var lgr *log.Logger
	goVerbose, ok := os.LookupEnv("GO_VERBOSE")
	if ok && strings.ToLower(goVerbose) == "true" {
		lgr = log.New(os.Stdout, "", log.LstdFlags)
		verbose = true
	}

	return sqlapi.NewDatabase(sqlapi.WrapDB(db, di), di, lgr, nil)
}

func cleanup(db sqlapi.Execer) {
	if db != nil {
		if c, ok := db.(io.Closer); ok {
			c.Close()
		}
		os.Remove("test.db")
	}
}

func user(i int) *User {
	return &User{
		Name:         fmt.Sprintf("user%02d", i),
		EmailAddress: fmt.Sprintf("foo%d@x.z", i),
		Active:       true,
		Fave:         big.NewInt(int64(i)),
		Numbers: Numbers{
			I8:  int8(i * 5),
			U8:  uint8(i * 6),
			I16: int16(i * 10),
			U16: uint16(i * 11),
			I32: int32(i * 100),
			U32: uint32(i * 101),
			I64: int64(i * 200),
			U64: uint64(i * 201),
			F32: float32(i * 300),
			F64: float64(i * 301),
		},
	}
}

func TestCreateTable_sql_syntax(t *testing.T) {
	g := NewGomegaWithT(t)

	cases := []struct {
		dialect  dialect.Dialect
		expected string
	}{
		{dialect.Sqlite,
			`CREATE TABLE IF NOT EXISTS "prefix_users" (
 "uid" integer not null primary key autoincrement,
 "name" text not null,
 "emailaddress" text not null,
 "addressid" bigint default null,
 "avatar" text default null,
 "role" text default null,
 "active" boolean not null,
 "admin" boolean not null,
 "fave" text,
 "lastupdated" bigint not null,
 "i8" tinyint not null default -8,
 "u8" tinyint unsigned not null default 8,
 "i16" smallint not null default -16,
 "u16" smallint unsigned not null default 16,
 "i32" int not null default -32,
 "u32" int unsigned not null default 32,
 "i64" bigint not null default -64,
 "u64" bigint unsigned not null default 64,
 "f32" float not null default 3.2,
 "f64" double not null default 6.4,
 "token" text not null,
 "secret" text not null,
 CONSTRAINT "prefix_users_c1" foreign key ("addressid") references "prefix_addresses" ("id") on update restrict on delete restrict,
 CONSTRAINT "prefix_users_c2" CHECK (role < 3)
)`},

		{dialect.Mysql,
			`CREATE TABLE IF NOT EXISTS ¬prefix_users¬ (
 ¬uid¬ bigint not null primary key auto_increment,
 ¬name¬ varchar(255) not null,
 ¬emailaddress¬ varchar(255) not null,
 ¬addressid¬ bigint default null,
 ¬avatar¬ text default null,
 ¬role¬ varchar(20) default null,
 ¬active¬ boolean not null,
 ¬admin¬ boolean not null,
 ¬fave¬ json,
 ¬lastupdated¬ bigint not null,
 ¬i8¬ tinyint not null default -8,
 ¬u8¬ tinyint unsigned not null default 8,
 ¬i16¬ smallint not null default -16,
 ¬u16¬ smallint unsigned not null default 16,
 ¬i32¬ int not null default -32,
 ¬u32¬ int unsigned not null default 32,
 ¬i64¬ bigint not null default -64,
 ¬u64¬ bigint unsigned not null default 64,
 ¬f32¬ float not null default 3.2,
 ¬f64¬ double not null default 6.4,
 ¬token¬ text not null,
 ¬secret¬ text not null,
 CONSTRAINT ¬prefix_users_c1¬ foreign key (¬addressid¬) references ¬prefix_addresses¬ (¬id¬) on update restrict on delete restrict,
 CONSTRAINT ¬prefix_users_c2¬ CHECK (role < 3)
) ENGINE=InnoDB DEFAULT CHARSET=utf8`},

		{dialect.Postgres.WithQuoter(quote.NoQuoter),
			`CREATE TABLE IF NOT EXISTS prefix_users (
 uid bigserial not null primary key,
 name text not null,
 emailaddress text not null,
 addressid bigint default null,
 avatar text default null,
 role text default null,
 active boolean not null,
 admin boolean not null,
 fave json,
 lastupdated bigint not null,
 i8 int8 not null default -8,
 u8 smallint not null default 8,
 i16 smallint not null default -16,
 u16 integer not null default 16,
 i32 integer not null default -32,
 u32 bigint not null default 32,
 i64 bigint not null default -64,
 u64 bigint not null default 64,
 f32 real not null default 3.2,
 f64 double precision not null default 6.4,
 token text not null,
 secret text not null,
 CONSTRAINT prefix_users_c1 foreign key (addressid) references prefix_addresses (id) on update restrict on delete restrict,
 CONSTRAINT prefix_users_c2 CHECK (role < 3)
)`},

		{dialect.Postgres,
			`CREATE TABLE IF NOT EXISTS "prefix_users" (
 "uid" bigserial not null primary key,
 "name" text not null,
 "emailaddress" text not null,
 "addressid" bigint default null,
 "avatar" text default null,
 "role" text default null,
 "active" boolean not null,
 "admin" boolean not null,
 "fave" json,
 "lastupdated" bigint not null,
 "i8" int8 not null default -8,
 "u8" smallint not null default 8,
 "i16" smallint not null default -16,
 "u16" integer not null default 16,
 "i32" integer not null default -32,
 "u32" bigint not null default 32,
 "i64" bigint not null default -64,
 "u64" bigint not null default 64,
 "f32" real not null default 3.2,
 "f64" double precision not null default 6.4,
 "token" text not null,
 "secret" text not null,
 CONSTRAINT "prefix_users_c1" foreign key ("addressid") references "prefix_addresses" ("id") on update restrict on delete restrict,
 CONSTRAINT "prefix_users_c2" CHECK (role < 3)
)`},
	}

	for _, c := range cases {
		d := sqlapi.NewDatabase(nil, c.dialect, nil, nil)
		tbl := NewDbUserTable("users", d).
			WithPrefix("prefix_").
			WithConstraint(constraint.CheckConstraint{"role < 3"})
		s := createDbUserTableSql(tbl, true)
		expected := strings.Replace(c.expected, "¬", "`", -1)
		if s != expected {
			outputDiff(s, c.dialect.String()+".txt")
		}
		g.Expect(s).To(Equal(expected), "%s\n%s", c.dialect, s)
	}
}

func outputDiff(a, name string) {
	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	f.WriteString(a)
	f.WriteString("\n")
	err = f.Close()
	if err != nil {
		panic(err)
	}
}

func TestCreateIndexSql(t *testing.T) {
	g := NewGomegaWithT(t)

	d := sqlapi.NewDatabase(nil, dialect.Postgres, nil, nil)
	tbl := NewDbUserTable("users", d).WithPrefix("prefix_")
	s := createDbUserTableEmailaddressIdxSql(tbl, "IF NOT EXISTS ")
	expected := `CREATE UNIQUE INDEX IF NOT EXISTS "prefix_emailaddress_idx" ON "prefix_users" ("emailaddress")`
	g.Expect(s).To(Equal(expected))
}

func TestDropIndexSql(t *testing.T) {
	g := NewGomegaWithT(t)

	cases := []struct {
		d        dialect.Dialect
		expected string
	}{
		{dialect.Sqlite, `DROP INDEX IF EXISTS "prefix_emailaddress_idx"`},
		{dialect.Mysql, "DROP INDEX `prefix_emailaddress_idx` ON `prefix_users`"},
		{dialect.Postgres, `DROP INDEX IF EXISTS "prefix_emailaddress_idx"`},
	}

	for _, c := range cases {
		d := sqlapi.NewDatabase(nil, c.d, nil, nil)
		tbl := NewDbUserTable("users", d).WithPrefix("prefix_")
		s := dropDbUserTableEmailaddressIdxSql(tbl, true)
		g.Expect(s).To(Equal(c.expected))
	}
}

func TestUpdateFields_ok_using_mock(t *testing.T) {
	g := NewGomegaWithT(t)

	mockDb := mockExecer{RowsAffected: 1}

	d := sqlapi.NewDatabase(mockDb, dialect.Mysql, nil, nil)
	tbl := NewDbUserTable("users", d)

	n, err := tbl.UpdateFields(require.One, where.NoOp(),
		sqlapi.Named("EmailAddress", "foo@x.com"),
		sqlapi.Named("Hash", "abc123"))

	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(n).To(Equal(int64(1)))
}

func TestUpdateFields_error_using_mock(t *testing.T) {
	g := NewGomegaWithT(t)

	exp := fmt.Errorf("foo")
	mockDb := mockExecer{Error: exp}

	d := sqlapi.NewDatabase(mockDb, dialect.Mysql, nil, nil)
	tbl := NewDbUserTable("users", d)

	_, err := tbl.UpdateFields(nil, where.NoOp(),
		sqlapi.Named("EmailAddress", "foo@x.com"),
		sqlapi.Named("Hash", "abc123"))

	g.Expect(errors.Cause(err)).To(Equal(exp))
}

func TestUpdate_ok_using_mock(t *testing.T) {
	g := NewGomegaWithT(t)

	mockDb := mockExecer{RowsAffected: 1}

	d := sqlapi.NewDatabase(mockDb, dialect.Mysql, nil, nil)
	tbl := NewDbUserTable("users", d)

	n, err := tbl.Update(require.One, &User{})

	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(n).To(Equal(int64(1)))
}

//-------------------------------------------------------------------------------------------------

func TestUpdate_error_using_mock(t *testing.T) {
	g := NewGomegaWithT(t)

	exp := fmt.Errorf("foo")
	mockDb := mockExecer{Error: exp}

	d := sqlapi.NewDatabase(mockDb, dialect.Mysql, nil, nil)
	tbl := NewDbUserTable("users", d)

	_, err := tbl.Update(nil, &User{})

	g.Expect(errors.Cause(err)).To(Equal(exp))
}

func TestUserCrud_using_database(t *testing.T) {
	g := NewGomegaWithT(t)

	d := newDatabase(t)
	defer cleanup(d.DB())

	addresses := NewAddressTable("addresses", d)

	users := NewDbUserTable("users", d)

	_, err := users.DropTable(true)
	g.Expect(err).NotTo(HaveOccurred())

	_, err = addresses.DropTable(true)
	g.Expect(err).NotTo(HaveOccurred())

	err = addresses.CreateTableWithIndexes(false)
	g.Expect(err).NotTo(HaveOccurred())

	err = users.CreateTableWithIndexes(false)
	g.Expect(err).NotTo(HaveOccurred())

	count_remainder_should_be(g, users, 0)

	insert_user_should_run_PreInsert(g, users, "user1")
	user1 := insert_user_should_run_PreInsert(g, users, "user2")
	insert_user_should_run_PreInsert(g, users, "user3")

	get_user_should_call_PostGet_and_match_expected(g, users, user1)

	get_unknown_user_should_return_nil(g, users, user1)

	must_get_unknown_user_should_return_error(g, users, user1)

	count_known_user_should_return_1(g, users)

	select_unknown_user_should_return_empty_list(g, users)

	select_unknown_user_requiring_one_should_return_error(g, users)

	query_one_nullstring_for_user_should_return_valid(g, users)

	query_one_nullstring_for_unknown_should_return_invalid(g, users)

	user2 := select_known_user_requiring_one_should_return_user(g, users)

	update_user_should_call_PreUpdate(g, users, user2)

	update_users_in_tx(g, users, user2)

	upsert_users(g, users, user2)

	delete_one_should_return_1(g, users)

	count_remainder_should_be(g, users, 3)
}

func count_remainder_should_be(g *GomegaWithT, tbl DbUserTable, expected int64) {
	c1, err := tbl.Count(where.NoOp())
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(c1).To(Equal(expected))
}

func insert_user_should_run_PreInsert(g *GomegaWithT, tbl DbUserTable, name string) *User {
	user := &User{Name: name, EmailAddress: name + "@x.z"}
	user = user.SetRole(UserRole)
	err := tbl.Insert(require.One, user)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(user.hash).To(Equal("PreInsert"))
	return user
}

func get_user_should_call_PostGet_and_match_expected(g *GomegaWithT, tbl DbUserTable, expected *User) {
	user, err := tbl.GetUserByUid(nil, expected.Uid)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(user.hash).To(Equal("PostGet"))
	user.hash = expected.hash
	g.Expect(user).To(Equal(expected))
}

func get_unknown_user_should_return_nil(g *GomegaWithT, tbl DbUserTable, expected *User) {
	user, err := tbl.GetUserByUid(nil, expected.Uid+100000)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(user).To(BeNil())
}

func must_get_unknown_user_should_return_error(g *GomegaWithT, tbl DbUserTable, expected *User) {
	_, err := tbl.GetUserByUid(require.One, expected.Uid+100000)
	g.Expect(err.Error()).To(Equal("expected to fetch one but got 0"))
}

func count_known_user_should_return_1(g *GomegaWithT, tbl DbUserTable) {
	count, err := tbl.Count(where.Eq("name", "user1"))
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(count).To(BeEquivalentTo(1))
}

func select_unknown_user_should_return_empty_list(g *GomegaWithT, tbl DbUserTable) {
	list, err := tbl.Select(require.None, where.Eq("name", "unknown"), nil)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(list).To(HaveLen(0))
}

func select_unknown_user_requiring_one_should_return_error(g *GomegaWithT, tbl DbUserTable) {
	list, err := tbl.Select(require.None, where.Eq("name", "unknown"), nil)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(list).To(HaveLen(0))

	_, err = tbl.Select(require.One, where.Eq("name", "unknown"), nil)
	g.Expect(err.Error()).To(Equal("expected to fetch one but got 0"))
}

func query_one_nullstring_for_user_should_return_valid(g *GomegaWithT, tbl DbUserTable) {
	q := fmt.Sprintf("select emailaddress from {TABLE} where name=?")
	s, err := tbl.QueryOneNullString(nil, q, "user1")
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(s.Valid).To(BeTrue())
	g.Expect(s.String).To(Equal("user1@x.z"))

	s, err = tbl.QueryOneNullString(require.One, q, "user1")
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(s.Valid).To(BeTrue())
	g.Expect(s.String).To(Equal("user1@x.z"))
}

func query_one_nullstring_for_unknown_should_return_invalid(g *GomegaWithT, tbl DbUserTable) {
	q := fmt.Sprintf("select emailaddress from {TABLE} where name=?")
	s, err := tbl.QueryOneNullString(nil, q, "foo")
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(s.Valid).To(BeFalse())

	_, err = tbl.QueryOneNullString(require.One, q, "foo")
	g.Expect(err.Error()).To(Equal("expected to fetch one but got 0"))
	g.Expect(s.Valid).To(BeFalse())
}

func select_known_user_requiring_one_should_return_user(g *GomegaWithT, tbl DbUserTable) *User {
	list, err := tbl.Select(require.One, where.Eq("name", "user1"), nil)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(list).To(HaveLen(1))
	return list[0]
}

func update_user_should_call_PreUpdate(g *GomegaWithT, tbl DbUserTable, user *User) {
	user.EmailAddress = "bah0@zzz.com"
	//utter.Dump(user)

	n, err := tbl.Update(require.One, user)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(n).To(BeEquivalentTo(1))
	g.Expect(user.hash).To(Equal("PreUpdate"))

	ss, err := tbl.SliceEmailaddress(require.One, where.Eq("uid", user.Uid), nil)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(ss).To(HaveLen(1))
	g.Expect(ss[0]).To(Equal("bah0@zzz.com"))
}

func update_users_in_tx(g *GomegaWithT, tbl DbUserTable, user *User) {
	user.EmailAddress = "dude@zzz.com"
	//utter.Dump(user)

	err := tbl.Transact(nil, func(t2 DbUserTabler) error {
		n, err := t2.Update(require.One, user)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(n).To(BeEquivalentTo(1))
		g.Expect(user.hash).To(Equal("PreUpdate"))
		return nil
	})
	g.Expect(err).NotTo(HaveOccurred())

	ss, err := tbl.SliceEmailaddress(require.One, where.Eq("uid", user.Uid), nil)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(ss).To(HaveLen(1))
	g.Expect(ss[0]).To(Equal("dude@zzz.com"))
}

func upsert_users(g *GomegaWithT, tbl DbUserTable, user *User) {
	user.EmailAddress = "dodo@zzz.com"
	//utter.Dump(user)

	err := tbl.Upsert(user, where.Eq("name", user.Name))
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(user.hash).To(Equal("PreUpdate"))

	ss, err := tbl.SliceEmailaddress(require.One, where.Eq("uid", user.Uid), nil)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(ss).To(HaveLen(1))
	g.Expect(ss[0]).To(Equal("dodo@zzz.com"))

	u2 := &User{
		Name:         "another",
		EmailAddress: "another@z.org",
		Active:       true,
	}
	err = tbl.Upsert(u2, where.Eq("name", u2.Name))
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(user.hash).To(Equal("PreUpdate"))

	ss, err = tbl.SliceEmailaddress(require.One, where.Eq("uid", u2.Uid), nil)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(ss).To(HaveLen(1))
	g.Expect(ss[0]).To(Equal("another@z.org"))
}

func delete_one_should_return_1(g *GomegaWithT, tbl DbUserTable) {
	n, err := tbl.Delete(require.One, where.Eq("name", "user1"))
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(n).To(BeEquivalentTo(1))
}

//-------------------------------------------------------------------------------------------------

func xTestMultiSelect_using_database(t *testing.T) {
	g := NewGomegaWithT(t)
	d := newDatabase(t)
	defer cleanup(d.DB())

	tbl := NewDbUserTable("users", d)

	_, err := tbl.DropTable(true)
	g.Expect(err).NotTo(HaveOccurred())

	err = tbl.CreateTableWithIndexes(true)
	g.Expect(err).NotTo(HaveOccurred())

	const n = 3

	var users []*User
	user0 := &User{Name: "user0", EmailAddress: "foo0@x.z"}
	// fave, avatar are null
	users = append(users, user0)

	for i := 1; i <= n; i++ {
		fave := big.NewInt(int64(i))
		user := &User{Fave: fave}
		user = user.SetRole(UserRole)
		user = user.SetName(fmt.Sprintf("user%d", i))
		user = user.SetEmailAddress(fmt.Sprintf("foo%d@x.z", i))
		user = user.SetAvatar(fmt.Sprintf("user%d-avatar%d", i, i))
		users = append(users, user)
	}

	err = tbl.Insert(require.All, users...)
	g.Expect(err).NotTo(HaveOccurred())

	list, err := tbl.Select(nil, where.NotEq("name", "nobody"), where.OrderBy("name").Desc())
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(list).To(HaveLen(n + 1))
	for i := 0; i <= n; i++ {
		users[n-i].hash = "PostGet"
		g.Expect(list[i]).To(Equal(users[n-i]))
	}
}

func xTestGetters_using_database(t *testing.T) {
	g := NewGomegaWithT(t)
	d := newDatabase(t)
	defer cleanup(d.DB())

	tbl := NewDbUserTable("users", d)

	err := tbl.CreateTableWithIndexes(true)
	g.Expect(err).NotTo(HaveOccurred())

	err = tbl.Truncate(true)
	g.Expect(err).NotTo(HaveOccurred())

	const n = 20

	list := make([]*User, n)
	for i := 0; i < n; i++ {
		list[i] = user(i)
	}

	err = tbl.Insert(require.All, list...)
	g.Expect(err).NotTo(HaveOccurred())

	names, err := tbl.SliceName(require.Exactly(n), where.NoOp(), where.OrderBy("name"))
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(names).To(HaveLen(n))

	for i := 0; i < n; i++ {
		exp := fmt.Sprintf("user%02d", i)
		g.Expect(names[i]).To(Equal(exp))
	}
}

func TestRowsAsMaps_using_database(t *testing.T) {
	g := NewGomegaWithT(t)
	d := newDatabase(t)
	defer cleanup(d.DB())

	tbl := NewDbUserTable("users", d)

	err := tbl.CreateTableWithIndexes(true)
	g.Expect(err).NotTo(HaveOccurred())

	err = tbl.Truncate(true)
	g.Expect(err).NotTo(HaveOccurred())

	const n = 5

	list := make([]*User, n)
	for i := 0; i < n; i++ {
		list[i] = user(i)
	}

	err = tbl.Insert(require.All, list...)
	g.Expect(err).NotTo(HaveOccurred())

	rows, err := support.Query(tbl, "SELECT * from users")
	g.Expect(err).NotTo(HaveOccurred())

	ram, err := sqlapi.WrapRows(rows)
	g.Expect(err).NotTo(HaveOccurred())

	i := 0
	for ram.Next() {
		m, e2 := ram.ScanToMap()
		g.Expect(e2).NotTo(HaveOccurred())

		g.Expect(m.Columns).To(HaveLen(22))
		g.Expect(m.ColumnTypes).To(HaveLen(22))
		g.Expect(m.Data).To(HaveLen(22))

		g.Expect(m.Data["name"]).To(BeEquivalentTo(fmt.Sprintf("user%02d", i)), fmt.Sprintf("%d %+v %#v", i, m.ColumnTypes[1], m.Data["name"]))
		g.Expect(cast.ToBool(m.Data["admin"])).To(Equal(false), fmt.Sprintf("%d %+v %#v", i, m.ColumnTypes[7], m.Data["admin"]))
		g.Expect(cast.ToInt(cast.ToString(m.Data["i8"]))).To(Equal(i*5), fmt.Sprintf("%d %+v %#v", i, m.ColumnTypes[10], m.Data["i8"]))
		g.Expect(cast.ToInt(cast.ToString(m.Data["u8"]))).To(Equal(i*6), fmt.Sprintf("%d %+v %#v", i, m.ColumnTypes[11], m.Data["u8"]))
		g.Expect(cast.ToInt(cast.ToString(m.Data["i16"]))).To(Equal(i*10), fmt.Sprintf("%d %+v %#v", i, m.ColumnTypes[12], m.Data["i16"]))
		g.Expect(cast.ToInt(cast.ToString(m.Data["u16"]))).To(Equal(i*11), fmt.Sprintf("%d %+v %#v", i, m.ColumnTypes[13], m.Data["u16"]))
		g.Expect(cast.ToInt(cast.ToString(m.Data["i32"]))).To(Equal(i*100), fmt.Sprintf("%d %+v %#v", i, m.ColumnTypes[14], m.Data["i32"]))
		g.Expect(cast.ToInt(cast.ToString(m.Data["u32"]))).To(Equal(i*101), fmt.Sprintf("%d %+v %#v", i, m.ColumnTypes[15], m.Data["u32"]))
		g.Expect(cast.ToInt(cast.ToString(m.Data["i64"]))).To(Equal(i*200), fmt.Sprintf("%d %+v %#v", i, m.ColumnTypes[16], m.Data["i64"]))
		g.Expect(cast.ToInt(cast.ToString(m.Data["u64"]))).To(Equal(i*201), fmt.Sprintf("%d %+v %#v", i, m.ColumnTypes[17], m.Data["u64"]))
		g.Expect(cast.ToFloat32(cast.ToString(m.Data["f32"]))).To(BeEquivalentTo(i*300), fmt.Sprintf("%d %+v %#v", i, m.ColumnTypes[18], m.Data["f32"]))
		g.Expect(cast.ToFloat32(cast.ToString(m.Data["f64"]))).To(BeEquivalentTo(i*301), fmt.Sprintf("%d %+v %#v", i, m.ColumnTypes[19], m.Data["f64"]))

		fave := big.NewInt(-1)
		err = fave.UnmarshalJSON(m.Data["fave"].([]byte))
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(fave.Cmp(big.NewInt(int64(i)))).To(Equal(0), fmt.Sprintf("%d %+v %#v", i, m.ColumnTypes[8], m.Data["fave"]))
		i++
	}
	g.Expect(i).To(Equal(n))
}

func xTestBulk_delete_using_database(t *testing.T) {
	g := NewGomegaWithT(t)
	d := newDatabase(t)
	defer cleanup(d.DB())

	tbl := NewDbUserTable("users", d)

	err := tbl.CreateTableWithIndexes(true)
	g.Expect(err).NotTo(HaveOccurred())

	err = tbl.Truncate(true)
	g.Expect(err).NotTo(HaveOccurred())

	const n = 17

	list := make([]*User, n)
	for i := 0; i < n; i++ {
		list[i] = user(i)
	}

	err = tbl.Insert(require.All, list...)
	g.Expect(err).NotTo(HaveOccurred())

	ids := make([]int64, n)
	for i := 0; i < n; i++ {
		ids[i] = list[i].Uid
	}

	j, err := tbl.DeleteUsers(require.All, ids...)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(j).To(BeEquivalentTo(n))
}

//-------------------------------------------------------------------------------------------------

func xTestNumericRanges_using_database(t *testing.T) {
	g := NewGomegaWithT(t)
	d := newDatabase(t)
	defer cleanup(d.DB())

	tbl := NewDbUserTable("users", d)

	err := tbl.CreateTableWithIndexes(true)
	g.Expect(err).NotTo(HaveOccurred())

	err = tbl.Truncate(true)
	g.Expect(err).NotTo(HaveOccurred())

	const n = 63 // note: cannot support 64 bits unsigned

	list := make([]*User, n)
	for i := 0; i < n; i++ {
		j := uint64(1) << uint(i)
		u := user(i)
		u.Numbers.I8 = int8(j)
		u.Numbers.U8 = uint8(j)
		u.Numbers.I16 = int16(j)
		u.Numbers.U16 = uint16(j)
		u.Numbers.I32 = int32(j)
		u.Numbers.U32 = uint32(j)
		u.Numbers.I64 = int64(j)
		u.Numbers.U64 = j
		u.Numbers.F32 = float32(j)
		u.Numbers.F64 = float64(j)
		list[i] = u
	}

	err = tbl.Insert(require.All, list...)
	g.Expect(err).NotTo(HaveOccurred())

	for i := 0; i < n; i++ {
		j := uint64(1) << uint(i)
		name := fmt.Sprintf("user%02d", i)
		u, e2 := tbl.GetUserByName(require.One, name)
		g.Expect(e2).NotTo(HaveOccurred())
		g.Expect(u.Numbers.I8).To(Equal(int8(j)), name)
		g.Expect(u.Numbers.U8).To(Equal(uint8(j)), name)
		g.Expect(u.Numbers.I16).To(Equal(int16(j)), name)
		g.Expect(u.Numbers.U16).To(Equal(uint16(j)), name)
		g.Expect(u.Numbers.I32).To(Equal(int32(j)), name)
		g.Expect(u.Numbers.U32).To(Equal(uint32(j)), name)
		g.Expect(u.Numbers.I64).To(Equal(int64(j)), name)
		g.Expect(u.Numbers.U64).To(Equal(j), name)
		g.Expect(u.Numbers.F32).To(Equal(float32(j)), name)
		g.Expect(u.Numbers.F64).To(Equal(float64(j)), name)
	}
}

//-------------------------------------------------------------------------------------------------

type mockExecer struct {
	RowsAffected int64
	Error        error
}

func (m mockExecer) QueryContext(ctx context.Context, query string, args ...interface{}) (sqlapi.SqlRows, error) {
	return nil, nil
}

func (m mockExecer) QueryRowContext(ctx context.Context, query string, args ...interface{}) sqlapi.SqlRow {
	return nil
}

func (m mockExecer) ExecContext(ctx context.Context, query string, args ...interface{}) (int64, error) {
	return m.RowsAffected, m.Error
}

func (m mockExecer) InsertContext(ctx context.Context, query string, args ...interface{}) (int64, error) {
	panic("implement me")
}

func (m mockExecer) PrepareContext(ctx context.Context, name, query string) (sqlapi.SqlStmt, error) {
	return nil, nil
}

func (m mockExecer) IsTx() bool {
	panic("implement me")
}

func (m mockExecer) SingleConn(ctx context.Context, fn func(conn *sql.Conn) error) error {
	panic("implement me")
}

func (m mockExecer) Transact(ctx context.Context, txOptions *sql.TxOptions, fn func(sqlapi.SqlTx) error) error {
	panic("implement me")
}

func (m mockExecer) PingContext(ctx context.Context) error {
	panic("implement me")
}

func (m mockExecer) Stats() sql.DBStats {
	panic("implement me")
}

func (m mockExecer) Close() error {
	panic("implement me")
}
