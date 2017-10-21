package where

import (
	"reflect"
	"testing"
)

func TestBuildWhereClause_happyCases(t *testing.T) {
	cases := []struct {
		di     Dialect
		wh     Expression
		expSql string
		args   []interface{}
	}{
		{MySQL, Clause{}, "", nil},
		{MySQL, Condition{"name not nil", nil}, "WHERE name not nil", nil},

		{MySQL,
			Condition{"name <>?", []interface{}{"Boo"}},
			"WHERE name <>?", []interface{}{"Boo"},
		},

		{MySQL,
			Eq("name", "Fred"),
			"WHERE name=?", []interface{}{"Fred"},
		},

		{MySQL,
			Eq("name", "Fred").And(Gt("age", 10)),
			"WHERE name=? AND age>?", []interface{}{"Fred", 10},
		},

		{MySQL,
			Eq("name", "Fred").Or(Gt("age", 10)),
			"WHERE name=? OR age>?", []interface{}{"Fred", 10},
		},

		{MySQL,
			Eq("name", "Fred").And(Gt("age", 10)).And(Gt("weight", 15)),
			"WHERE name=? AND age>? AND weight>?",
			[]interface{}{"Fred", 10, 15},
		},

		{MySQL,
			Eq("name", "Fred").Or(Gt("age", 10)).Or(Gt("weight", 15)),
			"WHERE name=? OR age>? OR weight>?",
			[]interface{}{"Fred", 10, 15},
		},

		{MySQL,
			Between("age", 10, 15).Or(Gt("weight", 17)),
			"WHERE age BETWEEN ? AND ? OR weight>?",
			[]interface{}{10, 15, 17},
		},

		{MySQL, GtEq("age", 10), "WHERE age>=?", []interface{}{10}},
		{MySQL, LtEq("age", 10), "WHERE age<=?", []interface{}{10}},
		{MySQL, NotEq("age", 10), "WHERE age<>?", []interface{}{10}},
		{MySQL, In("age", 10, 12, 14), "WHERE age IN (?,?,?)", []interface{}{10, 12, 14}},

		{MySQL, Not(Eq("name", "Fred")), "WHERE NOT (name=?)", []interface{}{"Fred"}},
		{MySQL, Not(Eq("name", "Fred").And(Lt("age", 10))), "WHERE NOT (name=? AND age<?)", []interface{}{"Fred", 10}},
		{MySQL, Not(Eq("name", "Fred").Or(Lt("age", 10))), "WHERE NOT (name=? OR age<?)", []interface{}{"Fred", 10}},
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
