package where_test

import (
	"reflect"
	"testing"
	"github.com/rickb777/sqlgen2"
	. "github.com/rickb777/sqlgen2/where"
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
		args        []interface{}
	}{
		{NoOp(), "", "", nil},

		{
			Condition{"name not nil", nil},
			"WHERE name not nil",
			"WHERE name not nil",
			nil,
		},

		{
			Null("name"),
			"WHERE name ISNULL",
			"WHERE name ISNULL",
			nil,
		},

		{
			Condition{"name <>?", []interface{}{"Boo"}},
			"WHERE name <>?",
			"WHERE name <>$1",
			[]interface{}{"Boo"},
		},

		{
			nameEqFred,
			"WHERE name=?",
			"WHERE name=$1",
			[]interface{}{"Fred"},
		},

		{
			nameEqFred.And(NoOp()),
			"WHERE (name=?)",
			"WHERE (name=$1)",
			[]interface{}{"Fred"},
		},

		{
			nameEqFred.And(Gt("age", 10)),
			"WHERE (name=?) AND (age>?)",
			"WHERE (name=$1) AND (age>$2)",
			[]interface{}{"Fred", 10},
		},

		{
			nameEqFred.Or(Gt("age", 10)),
			"WHERE (name=?) OR (age>?)",
			"WHERE (name=$1) OR (age>$2)",
			[]interface{}{"Fred", 10},
		},

		{
			nameEqFred.And(ageGt5).And(Gt("weight", 15)),
			"WHERE (name=?) AND (age>?) AND (weight>?)",
			"WHERE (name=$1) AND (age>$2) AND (weight>$3)",
			[]interface{}{"Fred", 5, 15},
		},

		{
			nameEqFred.Or(ageGt5).Or(Gt("weight", 15)),
			"WHERE (name=?) OR (age>?) OR (weight>?)",
			"WHERE (name=$1) OR (age>$2) OR (weight>$3)",
			[]interface{}{"Fred", 5, 15},
		},

		{
			Between("age", 12, 18).Or(Gt("weight", 45)),
			"WHERE (age BETWEEN ? AND ?) OR (weight>?)",
			"WHERE (age BETWEEN $1 AND $2) OR (weight>$3)",
			[]interface{}{12, 18, 45},
		},

		{
			GtEq("age", 10),
			"WHERE age>=?",
			"WHERE age>=$1",
			[]interface{}{10},
		},

		{
			LtEq("age", 10),
			"WHERE age<=?",
			"WHERE age<=$1",
			[]interface{}{10},
		},

		{
			NotEq("age", 10),
			"WHERE age<>?",
			"WHERE age<>$1",
			[]interface{}{10},
		},

		{
			In("age", 10, 12, 14),
			"WHERE age IN (?,?,?)",
			"WHERE age IN ($1,$2,$3)",
			[]interface{}{10, 12, 14},
		},

		{
			In("age", []int{10, 12, 14}),
			"WHERE age IN (?,?,?)",
			"WHERE age IN ($1,$2,$3)",
			[]interface{}{10, 12, 14},
		},

		{
			Not(nameEqFred),
			"WHERE NOT (name=?)",
			"WHERE NOT (name=$1)",
			[]interface{}{"Fred"},
		},

		{
			Not(nameEqFred.And(ageLt10)),
			"WHERE NOT ((name=?) AND (age<?))",
			"WHERE NOT ((name=$1) AND (age<$2))",
			[]interface{}{"Fred", 10},
		},

		{
			Not(nameEqFred.Or(ageLt10)),
			"WHERE NOT ((name=?) OR (age<?))",
			"WHERE NOT ((name=$1) OR (age<$2))",
			[]interface{}{"Fred", 10},
		},

		{
			And(nameEqFred, ageLt10),
			"WHERE (name=?) AND (age<?)",
			"WHERE (name=$1) AND (age<$2)",
			[]interface{}{"Fred", 10},
		},

		{
			And(nameEqFred).And(And(ageLt10)),
			"WHERE (name=?) AND (age<?)",
			"WHERE (name=$1) AND (age<$2)",
			[]interface{}{"Fred", 10},
		},

		{
			Or(nameEqFred, ageLt10),
			"WHERE (name=?) OR (age<?)",
			"WHERE (name=$1) OR (age<$2)",
			[]interface{}{"Fred", 10},
		},

		{
			And(nameEqFred.Or(nameEqJohn), ageLt10),
			"WHERE ((name=?) OR (name=?)) AND (age<?)",
			"WHERE ((name=$1) OR (name=$2)) AND (age<$3)",
			[]interface{}{"Fred", "John", 10},
		},

		{
			Or(nameEqFred, ageLt10.And(ageGt5)),
			"WHERE (name=?) OR ((age<?) AND (age>?))",
			"WHERE (name=$1) OR ((age<$2) AND (age>$3))",
			[]interface{}{"Fred", 10, 5},
		},

		{
			Or(nameEqFred, nameEqJohn).And(ageGt5),
			"WHERE ((name=?) OR (name=?)) AND (age>?)",
			"WHERE ((name=$1) OR (name=$2)) AND (age>$3)",
			[]interface{}{"Fred", "John", 5},
		},

		{
			Or(nameEqFred, nameEqJohn, And(ageGt5)),
			"WHERE (name=?) OR (name=?) OR ((age>?))",
			"WHERE (name=$1) OR (name=$2) OR ((age>$3))",
			[]interface{}{"Fred", "John", 5},
		},

		{
			Or().Or(NoOp()).And(NoOp()),
			"",
			"",
			nil,
		},

		{
			And(Or(NoOp())),
			"",
			"",
			nil,
		},
	}

	for i, c := range cases {
		sql, args := c.wh.Build(sqlgen2.Mysql)

		if sql != c.expMysql {
			t.Errorf("%d Mysql: Wanted %s\nGot %s", i, c.expMysql, sql)
		}

		if !reflect.DeepEqual(args, c.args) {
			t.Errorf("%d Mysql: Wanted %v\nGot %v", i, c.args, args)
		}

		sql, args = c.wh.Build(sqlgen2.Postgres)

		if sql != c.expPostgres {
			t.Errorf("%d Postgres: Wanted %s\nGot %s", i, c.expPostgres, sql)
		}

		if !reflect.DeepEqual(args, c.args) {
			t.Errorf("%d Postgres: Wanted %v\nGot %v", i, c.args, args)
		}
	}
}
