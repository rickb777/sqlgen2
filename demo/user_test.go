package demo

import (
	. "fmt"
	"context"
	"testing"
	"database/sql"
	"database/sql/driver"
	"log"
	"math/big"
	"os"
	"strings"
	. "github.com/onsi/gomega"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/spf13/cast"
	"github.com/rickb777/sqlgen2"
	"github.com/rickb777/sqlgen2/where"
	"github.com/rickb777/sqlgen2/schema"
	"github.com/rickb777/sqlgen2/require"
	"github.com/rickb777/sqlgen2/constraint"
	"github.com/kortschak/utter"
)

var db *sql.DB
var dialect schema.Dialect

func connect(t *testing.T) {
	db = nil
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
		t.Logf("\n***Warning: unable to connect to %s (%v); test is only partially complete.\n\n", dbDriver, err)
		return
	}
	err = conn.Ping()
	if err != nil {
		t.Logf("\n***Warning: unable to connect to %s (%v); test is only partially complete.\n\n", dbDriver, err)
		return
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
		Name:         Sprintf("user%02d", i),
		EmailAddress: Sprintf("foo%d@x.z", i),
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
	RegisterTestingT(t)

	cases := []struct {
		dialect  schema.Dialect
		expected string
	}{
		{schema.Sqlite,
`CREATE TABLE IF NOT EXISTS prefix_users (
 ¬uid¬          integer not null primary key autoincrement,
 ¬name¬         text not null,
 ¬emailaddress¬ text not null,
 ¬addressid¬    bigint default null,
 ¬avatar¬       text default null,
 ¬role¬         text default null,
 ¬active¬       boolean not null,
 ¬admin¬        boolean not null,
 ¬fave¬         text,
 ¬lastupdated¬  bigint not null,
 ¬i8¬           tinyint not null default -8,
 ¬u8¬           tinyint unsigned not null default 8,
 ¬i16¬          smallint not null default -16,
 ¬u16¬          smallint unsigned not null default 16,
 ¬i32¬          int not null default -32,
 ¬u32¬          int unsigned not null default 32,
 ¬i64¬          bigint not null default -64,
 ¬u64¬          bigint unsigned not null default 64,
 ¬f32¬          float not null default 3.2,
 ¬f64¬          double not null default 6.4,
 ¬token¬        text not null,
 ¬secret¬       text not null,
 CONSTRAINT prefix_users_c1 foreign key (¬addressid¬) references prefix_addresses (¬id¬) on update restrict on delete restrict,
 CONSTRAINT prefix_users_c2 CHECK (role < 3)
)`},
		{schema.Mysql,
`CREATE TABLE IF NOT EXISTS prefix_users (
 ¬uid¬          bigint not null primary key auto_increment,
 ¬name¬         varchar(255) not null,
 ¬emailaddress¬ varchar(255) not null,
 ¬addressid¬    bigint default null,
 ¬avatar¬       varchar(255) default null,
 ¬role¬         varchar(20) default null,
 ¬active¬       tinyint(1) not null,
 ¬admin¬        tinyint(1) not null,
 ¬fave¬         json,
 ¬lastupdated¬  bigint not null,
 ¬i8¬           tinyint not null default -8,
 ¬u8¬           tinyint unsigned not null default 8,
 ¬i16¬          smallint not null default -16,
 ¬u16¬          smallint unsigned not null default 16,
 ¬i32¬          int not null default -32,
 ¬u32¬          int unsigned not null default 32,
 ¬i64¬          bigint not null default -64,
 ¬u64¬          bigint unsigned not null default 64,
 ¬f32¬          float not null default 3.2,
 ¬f64¬          double not null default 6.4,
 ¬token¬        varchar(255) not null,
 ¬secret¬       varchar(255) not null,
 CONSTRAINT prefix_users_c1 foreign key (¬addressid¬) references prefix_addresses (¬id¬) on update restrict on delete restrict,
 CONSTRAINT prefix_users_c2 CHECK (role < 3)
) ENGINE=InnoDB DEFAULT CHARSET=utf8`},
		{schema.Postgres,
`CREATE TABLE IF NOT EXISTS prefix_users (
 "uid"          bigserial not null primary key,
 "name"         varchar(255) not null,
 "emailaddress" varchar(255) not null,
 "addressid"    bigint default null,
 "avatar"       varchar(255) default null,
 "role"         varchar(20) default null,
 "active"       boolean not null,
 "admin"        boolean not null,
 "fave"         json,
 "lastupdated"  bigint not null,
 "i8"           int8 not null default -8,
 "u8"           smallint not null default 8,
 "i16"          smallint not null default -16,
 "u16"          integer not null default 16,
 "i32"          integer not null default -32,
 "u32"          bigint not null default 32,
 "i64"          bigint not null default -64,
 "u64"          bigint not null default 64,
 "f32"          real not null default 3.2,
 "f64"          double precision not null default 6.4,
 "token"        varchar(255) not null,
 "secret"       varchar(255) not null,
 CONSTRAINT prefix_users_c1 foreign key ("addressid") references prefix_addresses ("id") on update restrict on delete restrict,
 CONSTRAINT prefix_users_c2 CHECK (role < 3)
)`},
	}

	for _, c := range cases {
		d := sqlgen2.NewDatabase(nil, c.dialect, nil, nil)
		tbl := NewDbUserTable("users", d).
			WithPrefix("prefix_").
			WithConstraint(
			constraint.CheckConstraint{"role < 3"})
		s := tbl.createTableSql(true)
		expected := strings.Replace(c.expected, "¬", "`", -1)
		if s != expected {
			outputDiff(s, c.dialect.String()+".txt")
		}
		Ω(s).Should(Equal(expected), "%s\n%s", c.dialect, s)
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
	RegisterTestingT(t)

	d := sqlgen2.NewDatabase(nil, schema.Postgres, nil, nil)
	tbl := NewDbUserTable("users", d).WithPrefix("prefix_")
	s := tbl.createDbEmailaddressIdxIndexSql("IF NOT EXISTS ")
	expected := `CREATE UNIQUE INDEX IF NOT EXISTS prefix_emailaddress_idx ON prefix_users (emailaddress)`
	Ω(s).Should(Equal(expected))
}

func TestDropIndexSql(t *testing.T) {
	RegisterTestingT(t)

	cases := []struct {
		d        schema.Dialect
		expected string
	}{
		{schema.Sqlite, `DROP INDEX IF EXISTS prefix_emailaddress_idx`},
		{schema.Mysql, `DROP INDEX prefix_emailaddress_idx ON prefix_users`},
		{schema.Postgres, `DROP INDEX IF EXISTS prefix_emailaddress_idx`},
	}

	for _, c := range cases {
		d := sqlgen2.NewDatabase(nil, c.d, nil, nil)
		tbl := NewDbUserTable("users", d).WithPrefix("prefix_")
		s := tbl.dropDbEmailaddressIdxIndexSql(true)
		Ω(s).Should(Equal(c.expected))
	}
}

func TestUpdateFields_ok_using_mock(t *testing.T) {
	RegisterTestingT(t)

	mockDb := mockExecer{RowsAffected: 1}

	d := sqlgen2.NewDatabase(mockDb, schema.Mysql, nil, nil)
	tbl := NewDbUserTable("users", d)

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

	d := sqlgen2.NewDatabase(mockDb, schema.Mysql, nil, nil)
	tbl := NewDbUserTable("users", d)

	_, err := tbl.UpdateFields(nil, where.NoOp(),
		sqlgen2.Named("EmailAddress", "foo@x.com"),
		sqlgen2.Named("Hash", "abc123"))

	Ω(err).Should(Equal(exp))
}

func TestUpdate_ok_using_mock(t *testing.T) {
	RegisterTestingT(t)

	mockDb := mockExecer{RowsAffected: 1}

	d := sqlgen2.NewDatabase(mockDb, schema.Mysql, nil, nil)
	tbl := NewDbUserTable("users", d)

	n, err := tbl.Update(require.One, &User{})

	Ω(err).Should(BeNil())
	Ω(n).Should(Equal(int64(1)))
}

func TestUpdate_error_using_mock(t *testing.T) {
	RegisterTestingT(t)

	exp := Errorf("foo")
	mockDb := mockExecer{Error: exp}

	d := sqlgen2.NewDatabase(mockDb, schema.Mysql, nil, nil)
	tbl := NewDbUserTable("users", d)

	_, err := tbl.Update(nil, &User{})

	Ω(err).Should(Equal(exp))
}

//-------------------------------------------------------------------------------------------------

func TestCrud_using_database(t *testing.T) {
	RegisterTestingT(t)
	connect(t)
	if db == nil {
		return
	}
	defer cleanup()

	d := sqlgen2.NewDatabase(db, dialect, nil, nil)
	if testing.Verbose() {
		lgr := log.New(os.Stderr, "", log.LstdFlags)
		d = sqlgen2.NewDatabase(db, dialect, lgr, nil)
	}
	addresses := NewAddressTable("addresses", d)

	users := NewDbUserTable("users", d)

	_, err := users.DropTable(true)
	Ω(err).Should(BeNil())

	_, err = addresses.DropTable(true)
	Ω(err).Should(BeNil())

	err = addresses.CreateTableWithIndexes(false)
	Ω(err).Should(BeNil())

	err = users.CreateTableWithIndexes(false)
	Ω(err).Should(BeNil())

	count_remainder_should_be(t, users, 0)

	insert_user_should_run_PreInsert(t, users, "user1")
	user1 := insert_user_should_run_PreInsert(t, users, "user2")
	insert_user_should_run_PreInsert(t, users, "user3")

	get_user_should_call_PostGet_and_match_expected(t, users, user1)

	get_unknown_user_should_return_nil(t, users, user1)

	must_get_unknown_user_should_return_error(t, users, user1)

	count_known_user_should_return_1(t, users)

	select_unknown_user_should_return_empty_list(t, users)

	select_unknown_user_requiring_one_should_return_error(t, users)

	query_one_nullstring_for_user_should_return_valid(t, users)

	query_one_nullstring_for_unknown_should_return_invalid(t, users)

	user2 := select_known_user_requiring_one_should_return_user(t, users)

	update_user_should_call_PreUpdate(t, users, user2)

	update_users_in_tx(t, users, user2)

	delete_one_should_return_1(t, users)

	count_remainder_should_be(t, users, 2)
}

func count_remainder_should_be(t *testing.T, tbl DbUserTable, expected int64) {
	c1, err := tbl.Count(where.NoOp())
	Ω(err).Should(BeNil())
	Ω(c1).Should(Equal(expected))
}

func insert_user_should_run_PreInsert(t *testing.T, tbl DbUserTable, name string) *User {
	user := &User{Name: name, EmailAddress: name + "@x.z"}
	user = user.SetRole(UserRole)
	err := tbl.Insert(require.One, user)
	Ω(err).Should(BeNil())
	Ω(user.hash).Should(Equal("PreInsert"))
	return user
}

func get_user_should_call_PostGet_and_match_expected(t *testing.T, tbl DbUserTable, expected *User) {
	user, err := tbl.GetUserByUid(nil, expected.Uid)
	Ω(err).Should(BeNil())
	if user.hash != "PostGet" {
		t.Fatalf("%q", user.hash)
	}
	user.hash = expected.hash
	Ω(user).Should(Equal(expected))
}

func get_unknown_user_should_return_nil(t *testing.T, tbl DbUserTable, expected *User) {
	user, err := tbl.GetUserByUid(nil, expected.Uid+100000)
	Ω(err).Should(BeNil())
	Ω(user).Should(BeNil())
}

func must_get_unknown_user_should_return_error(t *testing.T, tbl DbUserTable, expected *User) {
	_, err := tbl.GetUserByUid(require.One, expected.Uid+100000)
	Ω(err.Error()).Should(Equal("expected to fetch one but got 0"))
}

func count_known_user_should_return_1(t *testing.T, tbl DbUserTable) {
	count, err := tbl.Count(where.Eq("name", "user1"))
	Ω(err).Should(BeNil())
	Ω(count).Should(BeEquivalentTo(1))
}

func select_unknown_user_should_return_empty_list(t *testing.T, tbl DbUserTable) {
	list, err := tbl.Select(require.None, where.Eq("name", "unknown"), nil)
	Ω(err).Should(BeNil())
	Ω(list).Should(HaveLen(0))
}

func select_unknown_user_requiring_one_should_return_error(t *testing.T, tbl DbUserTable) {
	list, err := tbl.Select(require.None, where.Eq("name", "unknown"), nil)
	Ω(err).Should(BeNil())
	Ω(list).Should(HaveLen(0))

	_, err = tbl.Select(require.One, where.Eq("name", "unknown"), nil)
	Ω(err.Error()).Should(Equal("expected to fetch one but got 0"))
}

func query_one_nullstring_for_user_should_return_valid(t *testing.T, tbl DbUserTable) {
	p := dialect.Placeholder("name", 1)
	q := Sprintf("select %s from {TABLE} where %s=%s", dialect.Quote("emailaddress"), dialect.Quote("name"), p)
	s, err := tbl.QueryOneNullString(nil, q, "user1")
	Ω(err).Should(BeNil())
	Ω(s.Valid).Should(BeTrue())
	Ω(s.String).Should(Equal("user1@x.z"))

	s, err = tbl.QueryOneNullString(require.One, q, "user1")
	Ω(err).Should(BeNil())
	Ω(s.Valid).Should(BeTrue())
	Ω(s.String).Should(Equal("user1@x.z"))
}

func query_one_nullstring_for_unknown_should_return_invalid(t *testing.T, tbl DbUserTable) {
	p := tbl.Dialect().Placeholder("name", 1)
	q := Sprintf("select %s from {TABLE} where %s=%s", dialect.Quote("emailaddress"), dialect.Quote("name"), p)
	s, err := tbl.QueryOneNullString(nil, q, "foo")
	Ω(err).Should(BeNil())
	Ω(s.Valid).Should(BeFalse())

	_, err = tbl.QueryOneNullString(require.One, q, "foo")
	Ω(err.Error()).Should(Equal("expected to fetch one but got 0"))
	Ω(s.Valid).Should(BeFalse())
}

func select_known_user_requiring_one_should_return_user(t *testing.T, tbl DbUserTable) *User {
	list, err := tbl.Select(require.One, where.Eq("name", "user1"), nil)
	Ω(err).Should(BeNil())
	Ω(list).Should(HaveLen(1))
	return list[0]
}

func update_user_should_call_PreUpdate(t *testing.T, tbl DbUserTable, user *User) {
	user.EmailAddress = "bah0@zzz.com"
	utter.Dump(user)

	n, err := tbl.Update(require.One, user)
	Ω(err).Should(BeNil())
	Ω(n).Should(BeEquivalentTo(1))
	Ω(user.hash).Should(Equal("PreUpdate"))

	ss, err := tbl.SliceEmailaddress(require.One, where.Eq("uid", user.Uid), nil)
	Ω(err).Should(BeNil())
	Ω(ss).Should(HaveLen(1))
	Ω(ss[0]).Should(Equal("bah0@zzz.com"))
}

func update_users_in_tx(t *testing.T, tbl DbUserTable, user *User) {
	user.EmailAddress = "dude@zzz.com"
	utter.Dump(user)

	t2, err := tbl.BeginTx(nil)
	Ω(err).Should(BeNil())

	n, err := t2.Update(require.One, user)
	Ω(err).Should(BeNil())
	Ω(n).Should(BeEquivalentTo(1))
	Ω(user.hash).Should(Equal("PreUpdate"))

	err = t2.Tx().Commit()
	Ω(err).Should(BeNil())

	ss, err := tbl.SliceEmailaddress(require.One, where.Eq("uid", user.Uid), nil)
	Ω(err).Should(BeNil())
	Ω(ss).Should(HaveLen(1))
	Ω(ss[0]).Should(Equal("dude@zzz.com"))
}

func delete_one_should_return_1(t *testing.T, tbl DbUserTable) {
	n, err := tbl.Delete(require.One, where.Eq("name", "user1"))
	Ω(err).Should(BeNil())
	Ω(n).Should(BeEquivalentTo(1))
}

//-------------------------------------------------------------------------------------------------

func TestMultiSelect_using_database(t *testing.T) {
	RegisterTestingT(t)
	connect(t)
	if db == nil {
		return
	}
	defer cleanup()

	d := sqlgen2.NewDatabase(db, dialect, nil, nil)
	if testing.Verbose() {
		lgr := log.New(os.Stderr, "", log.LstdFlags)
		d = sqlgen2.NewDatabase(db, dialect, lgr, nil)
	}
	tbl := NewDbUserTable("users", d)

	_, err := tbl.DropTable(true)
	Ω(err).Should(BeNil())

	err = tbl.CreateTableWithIndexes(true)
	Ω(err).Should(BeNil())

	const n = 3

	var users []*User
	user0 := &User{Name: "user0", EmailAddress: "foo0@x.z"}
	// fave, avatar are null
	users = append(users, user0)

	for i := 1; i <= n; i++ {
		fave := big.NewInt(int64(i))
		user := &User{Fave: fave}
		user = user.SetRole(UserRole)
		user = user.SetName(Sprintf("user%d", i))
		user = user.SetEmailAddress(Sprintf("foo%d@x.z", i))
		user = user.SetAvatar(Sprintf("user%d-avatar%d", i, i))
		users = append(users, user)
	}

	err = tbl.Insert(require.All, users...)
	Ω(err).Should(BeNil())

	list, err := tbl.Select(nil, where.NotEq("name", "nobody"), where.OrderBy("name").Desc())
	Ω(err).Should(BeNil())
	Ω(list).Should(HaveLen(n + 1))
	for i := 0; i <= n; i++ {
		users[n-i].hash = "PostGet"
		Ω(list[i]).Should(Equal(users[n-i]))
	}
}

func TestGetters_using_database(t *testing.T) {
	RegisterTestingT(t)
	connect(t)
	if db == nil {
		return
	}
	defer cleanup()

	d := sqlgen2.NewDatabase(db, dialect, nil, nil)
	if testing.Verbose() {
		lgr := log.New(os.Stderr, "", log.LstdFlags)
		d = sqlgen2.NewDatabase(db, dialect, lgr, nil)
	}
	tbl := NewDbUserTable("users", d)

	err := tbl.CreateTableWithIndexes(true)
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

	names, err := tbl.SliceName(require.Exactly(n), where.NoOp(), where.OrderBy("name"))
	Ω(err).Should(BeNil())
	Ω(names).Should(HaveLen(n))

	for i := 0; i < n; i++ {
		exp := Sprintf("user%02d", i)
		Ω(names[i]).Should(Equal(exp))
	}
}

func TestRowsAsMaps_using_database(t *testing.T) {
	RegisterTestingT(t)
	connect(t)
	if db == nil {
		return
	}
	defer cleanup()

	d := sqlgen2.NewDatabase(db, dialect, nil, nil)
	if testing.Verbose() {
		lgr := log.New(os.Stderr, "", log.LstdFlags)
		d = sqlgen2.NewDatabase(db, dialect, lgr, nil)
	}
	tbl := NewDbUserTable("users", d)

	err := tbl.CreateTableWithIndexes(true)
	Ω(err).Should(BeNil())

	err = tbl.Truncate(true)
	Ω(err).Should(BeNil())

	const n = 5

	list := make([]*User, n)
	for i := 0; i < n; i++ {
		list[i] = user(i)
	}

	err = tbl.Insert(require.All, list...)
	Ω(err).Should(BeNil())

	rows, err := tbl.Query("SELECT * from users")
	Ω(err).Should(BeNil())

	ram, err := sqlgen2.WrapRows(rows)
	Ω(err).Should(BeNil())

	i := 0
	for ram.Next() {
		m, err := ram.ScanToMap()
		Ω(err).Should(BeNil())

		Ω(m.Columns).Should(HaveLen(22))
		Ω(m.ColumnTypes).Should(HaveLen(22))
		Ω(m.Data).Should(HaveLen(22))

		Ω(m.Data["name"]).Should(BeEquivalentTo(Sprintf("user%02d", i)), Sprintf("%d %+v %#v", i, m.ColumnTypes[1], m.Data["name"]))
		Ω(cast.ToBool(m.Data["admin"])).Should(Equal(false), Sprintf("%d %+v %#v", i, m.ColumnTypes[7], m.Data["admin"]))
		Ω(cast.ToInt(cast.ToString(m.Data["i8"]))).Should(Equal(i*5), Sprintf("%d %+v %#v", i, m.ColumnTypes[10], m.Data["i8"]))
		Ω(cast.ToInt(cast.ToString(m.Data["u8"]))).Should(Equal(i*6), Sprintf("%d %+v %#v", i, m.ColumnTypes[11], m.Data["u8"]))
		Ω(cast.ToInt(cast.ToString(m.Data["i16"]))).Should(Equal(i*10), Sprintf("%d %+v %#v", i, m.ColumnTypes[12], m.Data["i16"]))
		Ω(cast.ToInt(cast.ToString(m.Data["u16"]))).Should(Equal(i*11), Sprintf("%d %+v %#v", i, m.ColumnTypes[13], m.Data["u16"]))
		Ω(cast.ToInt(cast.ToString(m.Data["i32"]))).Should(Equal(i*100), Sprintf("%d %+v %#v", i, m.ColumnTypes[14], m.Data["i32"]))
		Ω(cast.ToInt(cast.ToString(m.Data["u32"]))).Should(Equal(i*101), Sprintf("%d %+v %#v", i, m.ColumnTypes[15], m.Data["u32"]))
		Ω(cast.ToInt(cast.ToString(m.Data["i64"]))).Should(Equal(i*200), Sprintf("%d %+v %#v", i, m.ColumnTypes[16], m.Data["i64"]))
		Ω(cast.ToInt(cast.ToString(m.Data["u64"]))).Should(Equal(i*201), Sprintf("%d %+v %#v", i, m.ColumnTypes[17], m.Data["u64"]))
		Ω(cast.ToFloat32(cast.ToString(m.Data["f32"]))).Should(BeEquivalentTo(i*300), Sprintf("%d %+v %#v", i, m.ColumnTypes[18], m.Data["f32"]))
		Ω(cast.ToFloat32(cast.ToString(m.Data["f64"]))).Should(BeEquivalentTo(i*301), Sprintf("%d %+v %#v", i, m.ColumnTypes[19], m.Data["f64"]))

		fave := big.NewInt(-1)
		err = fave.UnmarshalJSON(m.Data["fave"].([]byte))
		Ω(err).Should(BeNil())
		Ω(fave.Cmp(big.NewInt(int64(i)))).Should(Equal(0), Sprintf("%d %+v %#v", i, m.ColumnTypes[8], m.Data["fave"]))
		i++
	}
	Ω(i).Should(Equal(n))
}

func TestBulk_delete_using_database(t *testing.T) {
	RegisterTestingT(t)
	connect(t)
	if db == nil {
		return
	}
	defer cleanup()

	d := sqlgen2.NewDatabase(db, dialect, nil, nil)
	if testing.Verbose() {
		lgr := log.New(os.Stderr, "", log.LstdFlags)
		d = sqlgen2.NewDatabase(db, dialect, lgr, nil)
	}
	tbl := NewDbUserTable("users", d)

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

func TestNumericRanges_using_database(t *testing.T) {
	RegisterTestingT(t)
	connect(t)
	if db == nil {
		return
	}
	defer cleanup()

	d := sqlgen2.NewDatabase(db, dialect, nil, nil)
	if testing.Verbose() {
		lgr := log.New(os.Stderr, "", log.LstdFlags)
		d = sqlgen2.NewDatabase(db, dialect, lgr, nil)
	}
	tbl := NewDbUserTable("users", d)

	err := tbl.CreateTableWithIndexes(true)
	Ω(err).Should(BeNil())

	err = tbl.Truncate(true)
	Ω(err).Should(BeNil())

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
	Ω(err).Should(BeNil())

	for i := 0; i < n; i++ {
		j := uint64(1) << uint(i)
		name := Sprintf("user%02d", i)
		u, err := tbl.GetUserByName(require.One, name)
		Ω(err).Should(BeNil())
		Ω(u.Numbers.I8).Should(Equal(int8(j)), name)
		Ω(u.Numbers.U8).Should(Equal(uint8(j)), name)
		Ω(u.Numbers.I16).Should(Equal(int16(j)), name)
		Ω(u.Numbers.U16).Should(Equal(uint16(j)), name)
		Ω(u.Numbers.I32).Should(Equal(int32(j)), name)
		Ω(u.Numbers.U32).Should(Equal(uint32(j)), name)
		Ω(u.Numbers.I64).Should(Equal(int64(j)), name)
		Ω(u.Numbers.U64).Should(Equal(j), name)
		Ω(u.Numbers.F32).Should(Equal(float32(j)), name)
		Ω(u.Numbers.F64).Should(Equal(float64(j)), name)
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
