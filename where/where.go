package where

import (
	"reflect"
	"strings"
)

const where = "WHERE "

type Expression interface {
	Build(dialect Dialect) (string, []interface{})
	build(args []interface{}, idx int, dialect Dialect) (string, []interface{}, int)
}

type Condition struct {
	Sql  string
	Args []interface{}
}

type Clause struct {
	wheres      []Condition
	subclause   *Clause
	conjunction string
}

type not struct {
	expression Expression
}

//-------------------------------------------------------------------------------------------------

func (not not) build(args []interface{}, idx int, dialect Dialect) (string, []interface{}, int) {
	s, a, n := not.expression.build(args, idx, dialect)
	return "NOT (" + s + ")", a, n
}

func (not not) Build(dialect Dialect) (string, []interface{}) {
	s, a, _ := not.build(nil, 0, dialect)
	return where + s, a
}

//-------------------------------------------------------------------------------------------------

func (cl Condition) build(args []interface{}, idx int, dialect Dialect) (string, []interface{}, int) {
	sql := cl.Sql
	for _, arg := range cl.Args {
		value := reflect.ValueOf(arg)
		switch value.Kind() {
		case reflect.Array, reflect.Slice:
			for j := 0; j < value.Len(); j++ {
				idx++
				sql = dialect.ReplaceNextPlaceholder(sql, idx)
				args = append(args, value.Index(j).Interface())
			}

		default:
			idx++
			sql = dialect.ReplaceNextPlaceholder(sql, idx)
			args = append(args, arg)
		}
	}
	return sql, args, idx
}

func (cl Condition) Build(dialect Dialect) (string, []interface{}) {
	wh := Clause{[]Condition{cl}, nil, and}
	s, a := wh.Build(dialect)
	return s, a
}

//-------------------------------------------------------------------------------------------------

func (wh Clause) build(args []interface{}, idx int, dialect Dialect) (string, []interface{}, int) {
	if len(wh.wheres) == 0 {
		return "", args, idx
	}

	var sqls []string

	for _, where := range wh.wheres {
		var sql string
		sql, args, idx = where.build(args, idx, dialect)
		sqls = append(sqls, sql)
	}

	sql := strings.Join(sqls, wh.conjunction)
	return sql, args, idx
}

func (wh Clause) Build(dialect Dialect) (string, []interface{}) {
	if len(wh.wheres) == 0 {
		return "", nil
	}

	sql, args, _ := wh.build(nil, 0, dialect)

	return where + sql, args
}
