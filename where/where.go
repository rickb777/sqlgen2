package where

import (
	"reflect"
	"strings"
)

const where = "WHERE "

// Dialect provides a method to convert named argument placeholders to the dialect needed
// for the database in use.
type Dialect interface {
	ReplacePlaceholders(sql string) string
}

// Expression is an element in a WHERE clause. Expressions may be nested in various ways.
type Expression interface {
	And(Expression) Clause
	Or(Expression) Clause
	Build(dialect Dialect) (string, []interface{})
	build(args []interface{}, dialect Dialect) (string, []interface{})
}

// Condition is a simple condition such as an equality test.
type Condition struct {
	Sql  string
	Args []interface{}
}

// Clause is a compound expression.
type Clause struct {
	wheres      []Expression
	conjunction string
}

type not struct {
	expression Expression
}

//-------------------------------------------------------------------------------------------------

func (not not) build(args []interface{}, dialect Dialect) (string, []interface{}) {
	sql, args := not.expression.build(args, dialect)
	return "NOT (" + sql + ")", args
}

func (not not) Build(dialect Dialect) (string, []interface{}) {
	sql, args := not.build(nil, dialect)
	sql = dialect.ReplacePlaceholders(sql)
	return where + sql, args
}

//-------------------------------------------------------------------------------------------------

func (cl Condition) build(args []interface{}, dialect Dialect) (string, []interface{}) {
	sql := cl.Sql
	for _, arg := range cl.Args {
		value := reflect.ValueOf(arg)
		switch value.Kind() {
		case reflect.Array, reflect.Slice:
			for j := 0; j < value.Len(); j++ {
				args = append(args, value.Index(j).Interface())
			}

		default:
			args = append(args, arg)
		}
	}
	return sql, args
}

func (cl Condition) Build(dialect Dialect) (string, []interface{}) {
	sql, args := cl.build(nil, dialect)
	sql = dialect.ReplacePlaceholders(sql)
	return where + sql, args
}

//-------------------------------------------------------------------------------------------------

func (wh Clause) build(args []interface{}, dialect Dialect) (string, []interface{}) {
	var sqls []string

	for _, where := range wh.wheres {
		var sql string
		sql, args = where.build(args, dialect)
		sqls = append(sqls, "("+sql+")")
	}

	sql := strings.Join(sqls, wh.conjunction)
	return sql, args
}

func (wh Clause) Build(dialect Dialect) (string, []interface{}) {
	if len(wh.wheres) == 0 {
		return "", nil
	}

	sql, args := wh.build(nil, dialect)
	sql = dialect.ReplacePlaceholders(sql)
	return where + sql, args
}
