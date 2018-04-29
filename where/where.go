package where

import (
	"reflect"
	"strings"
	"fmt"
	"bytes"
)

const where = "WHERE "

// Dialect provides a method to convert named argument placeholders to the dialect needed
// for the database in use.
type Dialect interface {
	ReplacePlaceholders(sql string, args []interface{}) string
	Quote(column string) string
}

// Expression is an element in a WHERE clause. Expressions may be nested in various ways.
type Expression interface {
	fmt.Stringer
	And(Expression) Clause
	Or(Expression) Clause
	Build(dialect Dialect) (string, []interface{})
	build(args []interface{}, dialect Dialect) (string, []interface{})
}

func BuildExpression(wh Expression, dialect Dialect) (string, []interface{}) {
	if wh == nil {
		return "", nil
	}
	return wh.Build(dialect)
}

//-------------------------------------------------------------------------------------------------

// Condition is a simple condition such as an equality test.
type Condition struct {
	Column, Predicate string
	Args              []interface{}
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
	sql = dialect.ReplacePlaceholders(sql, args)
	return where + sql, args
}

func (not not) String() string {
	sql, _ := not.Build(neutralDialect{})
	return sql
}

//-------------------------------------------------------------------------------------------------

func (cl Condition) build(args []interface{}, dialect Dialect) (string, []interface{}) {
	sql := dialect.Quote(cl.Column) + cl.Predicate
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
	sql = dialect.ReplacePlaceholders(sql, args)
	return where + sql, args
}

func (cl Condition) String() string {
	sql, _ := cl.Build(neutralDialect{})
	return sql
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
	sql = dialect.ReplacePlaceholders(sql, args)
	return where + sql, args
}

func (wh Clause) String() string {
	sql, _ := wh.Build(neutralDialect{})
	return sql
}

//-------------------------------------------------------------------------------------------------

type neutralDialect struct{}

func (d neutralDialect) ReplacePlaceholders(sql string, args []interface{}) string {
	buf := &bytes.Buffer{}
	idx := 0
	for _, r := range sql {
		if r == '?' && idx < len(args) {
			v := args[idx]
			t := reflect.TypeOf(v)
			switch t.Kind() {
			case reflect.Bool,
				reflect.Int,
				reflect.Int8,
				reflect.Int16,
				reflect.Int32,
				reflect.Int64,
				reflect.Uint,
				reflect.Uint8,
				reflect.Uint16,
				reflect.Uint32,
				reflect.Uint64,
				reflect.Uintptr,
				reflect.Float32,
				reflect.Float64:
				buf.WriteString(fmt.Sprintf(`%v`, v))
			default:
				buf.WriteString(fmt.Sprintf(`'%v'`, v))
			}
			idx++
		} else {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}

func (d neutralDialect) Quote(column string) string {
	return column
}
