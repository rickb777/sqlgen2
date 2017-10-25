package where

import (
	"reflect"
	"testing"
	"github.com/rickb777/sqlgen2/dialect"
)

func TestBuildWhereClause_happyCases(t *testing.T) {
	cases := []struct {
		di     Dialect
		wh     Expression
		expSql string
		args   []interface{}
	}{
		{dialect.MySQL, Clause{}, "", nil},
		{dialect.MySQL, Condition{"name not nil", nil}, "WHERE name not nil", nil},

		{dialect.MySQL,
			Condition{"name <>?", []interface{}{"Boo"}},
			"WHERE name <>?", []interface{}{"Boo"},
		},

		{dialect.MySQL,
			Eq("name", "Fred"),
			"WHERE name=?", []interface{}{"Fred"},
		},

		{dialect.MySQL,
			Eq("name", "Fred").And(Gt("age", 10)),
			"WHERE name=? AND age>?", []interface{}{"Fred", 10},
		},

		{dialect.MySQL,
			Eq("name", "Fred").Or(Gt("age", 10)),
			"WHERE name=? OR age>?", []interface{}{"Fred", 10},
		},

		{dialect.MySQL,
			Eq("name", "Fred").And(Gt("age", 10)).And(Gt("weight", 15)),
			"WHERE name=? AND age>? AND weight>?",
			[]interface{}{"Fred", 10, 15},
		},

		{dialect.MySQL,
			Eq("name", "Fred").Or(Gt("age", 10)).Or(Gt("weight", 15)),
			"WHERE name=? OR age>? OR weight>?",
			[]interface{}{"Fred", 10, 15},
		},

		{dialect.MySQL,
			Between("age", 10, 15).Or(Gt("weight", 17)),
			"WHERE age BETWEEN ? AND ? OR weight>?",
			[]interface{}{10, 15, 17},
		},

		{dialect.MySQL, GtEq("age", 10), "WHERE age>=?", []interface{}{10}},
		{dialect.MySQL, LtEq("age", 10), "WHERE age<=?", []interface{}{10}},
		{dialect.MySQL, NotEq("age", 10), "WHERE age<>?", []interface{}{10}},
		{dialect.MySQL, In("age", 10, 12, 14), "WHERE age IN (?,?,?)", []interface{}{10, 12, 14}},

		{dialect.MySQL, Not(Eq("name", "Fred")), "WHERE NOT (name=?)", []interface{}{"Fred"}},
		{dialect.MySQL, Not(Eq("name", "Fred").And(Lt("age", 10))), "WHERE NOT (name=? AND age<?)", []interface{}{"Fred", 10}},
		{dialect.MySQL, Not(Eq("name", "Fred").Or(Lt("age", 10))), "WHERE NOT (name=? OR age<?)", []interface{}{"Fred", 10}},

		//-----------------------------------------------------------------------------------------

		{dialect.Postgres,
			Condition{"name <>?", []interface{}{"Boo"}},
			"WHERE name <>$1", []interface{}{"Boo"},
		},

		{dialect.Postgres,
			Eq("name", "Fred"),
			"WHERE name=$1", []interface{}{"Fred"},
		},

		{dialect.Postgres,
			Eq("name", "Fred").And(Gt("age", 10)),
			"WHERE name=$1 AND age>$2", []interface{}{"Fred", 10},
		},

		{dialect.Postgres,
			Eq("name", "Fred").Or(Gt("age", 10)),
			"WHERE name=$1 OR age>$2", []interface{}{"Fred", 10},
		},

		{dialect.Postgres,
			Eq("name", "Fred").And(Gt("age", 10)).And(Gt("weight", 15)),
			"WHERE name=$1 AND age>$2 AND weight>$3",
			[]interface{}{"Fred", 10, 15},
		},

		{dialect.Postgres,
			Eq("name", "Fred").Or(Gt("age", 10)).Or(Gt("weight", 15)),
			"WHERE name=$1 OR age>$2 OR weight>$3",
			[]interface{}{"Fred", 10, 15},
		},

		{dialect.Postgres,
			Between("age", 10, 15).Or(Gt("weight", 17)),
			"WHERE age BETWEEN $1 AND $2 OR weight>$3",
			[]interface{}{10, 15, 17},
		},

		{dialect.Postgres, GtEq("age", 10), "WHERE age>=$1", []interface{}{10}},
		{dialect.Postgres, LtEq("age", 10), "WHERE age<=$1", []interface{}{10}},
		{dialect.Postgres, NotEq("age", 10), "WHERE age<>$1", []interface{}{10}},
		{dialect.Postgres, In("age", 10, 12, 14), "WHERE age IN ($1,$2,$3)", []interface{}{10, 12, 14}},

		{dialect.Postgres, Not(Eq("name", "Fred")), "WHERE NOT (name=$1)", []interface{}{"Fred"}},
		{dialect.Postgres, Not(Eq("name", "Fred").And(Lt("age", 10))), "WHERE NOT (name=$1 AND age<$2)", []interface{}{"Fred", 10}},
		{dialect.Postgres, Not(Eq("name", "Fred").Or(Lt("age", 10))), "WHERE NOT (name=$1 OR age<$2)", []interface{}{"Fred", 10}},
	}

	for i, c := range cases {
		sql, args := c.wh.Build(c.di)

		if sql != c.expSql {
			t.Errorf("%d: Wanted %s\nGot %s", i, c.expSql, sql)
		}

		if !reflect.DeepEqual(args, c.args) {
			t.Errorf("%d: Wanted %v\nGot %v", i, c.args, args)
		}
	}
}
