package where_test

import (
	"reflect"
	"testing"
	. "github.com/rickb777/sqlgen2/where"
	"github.com/rickb777/sqlgen2/schema"
)

func TestBuildWhereClause_happyCases(t *testing.T) {
	nameEqFred := Eq("name", "Fred")
	nameEqJohn := Eq("name", "John")
	ageLt10 := Lt("age", 10)
	ageGt5 := Gt("age", 5)

	cases := []struct {
		wh          Expression
		expMysql    string
		expPostgres string
		expString   string
		args        []interface{}
	}{
		{NoOp(), "", "", "", nil},

		{
			Condition{"name"," not nil", nil},
			"WHERE `name` not nil",
			`WHERE "name" not nil`,
			`WHERE name not nil`,
			nil,
		},

		{
			Condition{"p.name"," not nil", nil},
			"WHERE `p`.`name` not nil",
			`WHERE "p"."name" not nil`,
			`WHERE p.name not nil`,
			nil,
		},

		{
			Null("name"),
			"WHERE `name` IS NULL",
			`WHERE "name" IS NULL`,
			`WHERE name IS NULL`,
			nil,
		},

		{
			NotNull("name"),
			"WHERE `name` IS NOT NULL",
			`WHERE "name" IS NOT NULL`,
			`WHERE name IS NOT NULL`,
			nil,
		},

		{
			Condition{"name"," <>?", []interface{}{"Boo"}},
			"WHERE `name` <>?",
			`WHERE "name" <>$1`,
			`WHERE name <>'Boo'`,
			[]interface{}{"Boo"},
		},

		{
			nameEqFred,
			"WHERE `name`=?",
			`WHERE "name"=$1`,
			`WHERE name='Fred'`,
			[]interface{}{"Fred"},
		},

		{
			Like("name", "F%"),
			"WHERE `name` LIKE ?",
			`WHERE "name" LIKE $1`,
			`WHERE name LIKE 'F%'`,
			[]interface{}{"F%"},
		},

		{
			NoOp().And(nameEqFred),
			"WHERE (`name`=?)",
			`WHERE ("name"=$1)`,
			`WHERE (name='Fred')`,
			[]interface{}{"Fred"},
		},

		{
			nameEqFred.And(Gt("age", 10)),
			"WHERE (`name`=?) AND (`age`>?)",
			`WHERE ("name"=$1) AND ("age">$2)`,
			`WHERE (name='Fred') AND (age>10)`,
			[]interface{}{"Fred", 10},
		},

		{
			nameEqFred.Or(Gt("age", 10)),
			"WHERE (`name`=?) OR (`age`>?)",
			`WHERE ("name"=$1) OR ("age">$2)`,
			`WHERE (name='Fred') OR (age>10)`,
			[]interface{}{"Fred", 10},
		},

		{
			nameEqFred.And(ageGt5).And(Gt("weight", 15)),
			"WHERE (`name`=?) AND (`age`>?) AND (`weight`>?)",
			`WHERE ("name"=$1) AND ("age">$2) AND ("weight">$3)`,
			`WHERE (name='Fred') AND (age>5) AND (weight>15)`,
			[]interface{}{"Fred", 5, 15},
		},

		{
			nameEqFred.Or(ageGt5).Or(Gt("weight", 15)),
			"WHERE (`name`=?) OR (`age`>?) OR (`weight`>?)",
			`WHERE ("name"=$1) OR ("age">$2) OR ("weight">$3)`,
			`WHERE (name='Fred') OR (age>5) OR (weight>15)`,
			[]interface{}{"Fred", 5, 15},
		},

		{
			Between("age", 12, 18).Or(Gt("weight", 45)),
			"WHERE (`age` BETWEEN ? AND ?) OR (`weight`>?)",
			`WHERE ("age" BETWEEN $1 AND $2) OR ("weight">$3)`,
			`WHERE (age BETWEEN 12 AND 18) OR (weight>45)`,
			[]interface{}{12, 18, 45},
		},

		{
			GtEq("age", 10),
			"WHERE `age`>=?",
			`WHERE "age">=$1`,
			`WHERE age>=10`,
			[]interface{}{10},
		},

		{
			LtEq("age", 10),
			"WHERE `age`<=?",
			`WHERE "age"<=$1`,
			`WHERE age<=10`,
			[]interface{}{10},
		},

		{
			NotEq("age", 10),
			"WHERE `age`<>?",
			`WHERE "age"<>$1`,
			`WHERE age<>10`,
			[]interface{}{10},
		},

		{
			In("age", 10, 12, 14),
			"WHERE `age` IN (?,?,?)",
			`WHERE "age" IN ($1,$2,$3)`,
			"WHERE age IN (10,12,14)",
			[]interface{}{10, 12, 14},
		},

		{
			In("age", []int{10, 12, 14}),
			"WHERE `age` IN (?,?,?)",
			`WHERE "age" IN ($1,$2,$3)`,
			"WHERE age IN (10,12,14)",
			[]interface{}{10, 12, 14},
		},

		{
			Not(nameEqFred),
			"WHERE NOT (`name`=?)",
			`WHERE NOT ("name"=$1)`,
			`WHERE NOT (name='Fred')`,
			[]interface{}{"Fred"},
		},

		{
			Not(nameEqFred.And(ageLt10)),
			"WHERE NOT ((`name`=?) AND (`age`<?))",
			`WHERE NOT (("name"=$1) AND ("age"<$2))`,
			`WHERE NOT ((name='Fred') AND (age<10))`,
			[]interface{}{"Fred", 10},
		},

		{
			Not(nameEqFred.Or(ageLt10)),
			"WHERE NOT ((`name`=?) OR (`age`<?))",
			`WHERE NOT (("name"=$1) OR ("age"<$2))`,
			`WHERE NOT ((name='Fred') OR (age<10))`,
			[]interface{}{"Fred", 10},
		},

		{
			Not(nameEqFred).And(ageLt10),
			"WHERE (NOT (`name`=?)) AND (`age`<?)",
			`WHERE (NOT ("name"=$1)) AND ("age"<$2)`,
			`WHERE (NOT (name='Fred')) AND (age<10)`,
			[]interface{}{"Fred", 10},
		},

		{
			Not(nameEqFred).Or(ageLt10),
			"WHERE (NOT (`name`=?)) OR (`age`<?)",
			`WHERE (NOT ("name"=$1)) OR ("age"<$2)`,
			`WHERE (NOT (name='Fred')) OR (age<10)`,
			[]interface{}{"Fred", 10},
		},

		{
			And(nameEqFred, ageLt10),
			"WHERE (`name`=?) AND (`age`<?)",
			`WHERE ("name"=$1) AND ("age"<$2)`,
			`WHERE (name='Fred') AND (age<10)`,
			[]interface{}{"Fred", 10},
		},

		{
			And(nameEqFred).And(And(ageLt10)),
			"WHERE (`name`=?) AND (`age`<?)",
			`WHERE ("name"=$1) AND ("age"<$2)`,
			`WHERE (name='Fred') AND (age<10)`,
			[]interface{}{"Fred", 10},
		},

		{
			Or(nameEqFred, ageLt10),
			"WHERE (`name`=?) OR (`age`<?)",
			`WHERE ("name"=$1) OR ("age"<$2)`,
			`WHERE (name='Fred') OR (age<10)`,
			[]interface{}{"Fred", 10},
		},

		{
			And(nameEqFred.Or(nameEqJohn), ageLt10),
			"WHERE ((`name`=?) OR (`name`=?)) AND (`age`<?)",
			`WHERE (("name"=$1) OR ("name"=$2)) AND ("age"<$3)`,
			`WHERE ((name='Fred') OR (name='John')) AND (age<10)`,
			[]interface{}{"Fred", "John", 10},
		},

		{
			Or(nameEqFred, ageLt10.And(ageGt5)),
			"WHERE (`name`=?) OR ((`age`<?) AND (`age`>?))",
			`WHERE ("name"=$1) OR (("age"<$2) AND ("age">$3))`,
			`WHERE (name='Fred') OR ((age<10) AND (age>5))`,
			[]interface{}{"Fred", 10, 5},
		},

		{
			Or(nameEqFred, nameEqJohn).And(ageGt5),
			"WHERE ((`name`=?) OR (`name`=?)) AND (`age`>?)",
			`WHERE (("name"=$1) OR ("name"=$2)) AND ("age">$3)`,
			`WHERE ((name='Fred') OR (name='John')) AND (age>5)`,
			[]interface{}{"Fred", "John", 5},
		},

		{
			Or(nameEqFred, nameEqJohn, And(ageGt5)),
			"WHERE (`name`=?) OR (`name`=?) OR ((`age`>?))",
			`WHERE ("name"=$1) OR ("name"=$2) OR (("age">$3))`,
			`WHERE (name='Fred') OR (name='John') OR ((age>5))`,
			[]interface{}{"Fred", "John", 5},
		},

		{
			Or().Or(NoOp()).And(NoOp()),
			"",
			"",
			"",
			nil,
		},

		{
			And(Or(NoOp())),
			"",
			"",
			"",
			nil,
		},
	}

	for i, c := range cases {
		sql, args := c.wh.Build(schema.Mysql)

		if sql != c.expMysql {
			t.Errorf("%d Mysql: Wanted %s\nGot %s", i, c.expMysql, sql)
		}

		if !reflect.DeepEqual(args, c.args) {
			t.Errorf("%d Mysql: Wanted %v\nGot %v", i, c.args, args)
		}

		sql, args = c.wh.Build(schema.Postgres)

		if sql != c.expPostgres {
			t.Errorf("%d Postgres: Wanted %s\nGot %s", i, c.expPostgres, sql)
		}

		if !reflect.DeepEqual(args, c.args) {
			t.Errorf("%d Postgres: Wanted %v\nGot %v", i, c.args, args)
		}

		s := c.wh.String()

		if s != c.expString {
			t.Errorf("%d String: Wanted %s\nGot %s", i, c.expString, s)
		}
	}
}

func TestQueryConstraint(t *testing.T) {
	cases := []struct {
		qc          QueryConstraint
		expSqlite   string
		expMysql    string
		expPostgres string
	}{
		{nil, "", "", ""},
		{Literal("order by foo"), "order by foo", "order by foo", "order by foo"},
		{OrderBy("foo").Asc(), "ORDER BY `foo`", "ORDER BY `foo`", `ORDER BY "foo"`},
		{OrderBy("foo").Desc(), "ORDER BY `foo` DESC", "ORDER BY `foo` DESC", `ORDER BY "foo" DESC`},
		{OrderBy("foo").OrderBy("bar"), "ORDER BY `bar`", "ORDER BY `bar`", `ORDER BY "bar"`},
		{Limit(0), "", "", ""},
		{Limit(10), "LIMIT 10", "LIMIT 10", "LIMIT 10"},
		{Offset(20), "OFFSET 20", "OFFSET 20", "OFFSET 20"},
		{Limit(5).OrderBy("foo", "bar"), "ORDER BY `foo`,`bar` LIMIT 5", "ORDER BY `foo`,`bar` LIMIT 5", `ORDER BY "foo","bar" LIMIT 5`},
		{OrderBy("foo").Desc().Limit(10).Offset(20), "ORDER BY `foo` DESC LIMIT 10 OFFSET 20", "ORDER BY `foo` DESC LIMIT 10 OFFSET 20", `ORDER BY "foo" DESC LIMIT 10 OFFSET 20`},
	}

	for i, c := range cases {
		var sql string

		if c.qc != nil {
			sql = BuildQueryConstraint(c.qc, schema.Sqlite)
		}

		if sql != c.expSqlite {
			t.Errorf("%d Sqlite: Wanted %s\nGot %s", i, c.expSqlite, sql)
		}

		if c.qc != nil {
			sql = BuildQueryConstraint(c.qc, schema.Mysql)
		}

		if sql != c.expMysql {
			t.Errorf("%d Mysql: Wanted %s\nGot %s", i, c.expMysql, sql)
		}

		if c.qc != nil {
			sql = BuildQueryConstraint(c.qc, schema.Postgres)
		}

		if sql != c.expPostgres {
			t.Errorf("%d Postgres: Wanted %s\nGot %s", i, c.expPostgres, sql)
		}
	}
}
