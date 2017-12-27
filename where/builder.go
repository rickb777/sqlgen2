package where

import (
	"bytes"
	"reflect"
)

const and = " AND "
const or = " OR "

// Null returns an 'ISNULL' condition on a column.
func Null(column string) Condition {
	return Condition{column + " ISNULL", []interface{}{}}
}

// Eq returns an equality condition on a column.
func Eq(column string, value interface{}) Condition {
	return Condition{column + "=?", []interface{}{value}}
}

// NotEq returns a not equal condition on a column.
func NotEq(column string, value interface{}) Condition {
	return Condition{column + "<>?", []interface{}{value}}
}

// Gt returns a greater than condition on a column.
func Gt(column string, value interface{}) Condition {
	return Condition{column + ">?", []interface{}{value}}
}

// GtEq returns a greater than or equal condition on a column.
func GtEq(column string, value interface{}) Condition {
	return Condition{column + ">=?", []interface{}{value}}
}

// Lt returns a less than condition on a column.
func Lt(column string, value interface{}) Condition {
	return Condition{column + "<?", []interface{}{value}}
}

// LtEq returns a less than or equal than condition on a column.
func LtEq(column string, value interface{}) Condition {
	return Condition{column + "<=?", []interface{}{value}}
}

// Between returns a between condition on a column.
func Between(column string, a, b interface{}) Condition {
	return Condition{column + " BETWEEN ? AND ?", []interface{}{a, b}}
}

// In returns an in condition on a column.
func In(column string, values ...interface{}) Condition {
	buf := &bytes.Buffer{}
	buf.WriteString(column)
	buf.WriteString(" IN (")
	i := 0
	for _, arg := range values {
		value := reflect.ValueOf(arg)
		switch value.Kind() {
		case reflect.Array, reflect.Slice:
			for j := 0; j < value.Len(); j++ {
				if i > 0 {
					buf.WriteByte(',')
				}
				buf.WriteByte('?')
				i++
			}

		default:
			if i > 0 {
				buf.WriteByte(',')
			}
			buf.WriteByte('?')
			i++
		}
	}
	buf.WriteByte(')')
	return Condition{buf.String(), values}
}

//-------------------------------------------------------------------------------------------------

// Not negates an expression.
func Not(el Expression) Expression {
	return not{el}
}

//-------------------------------------------------------------------------------------------------

// And combines two conditions into a clause that requires they are both true.
func (cl Condition) And(c2 Condition) Clause {
	return Clause{[]Expression{cl, c2}, and}
}

// Or combines two conditions into a clause that requires either is true.
func (cl Condition) Or(c2 Condition) Clause {
	return Clause{[]Expression{cl, c2}, or}
}

// And combines two clauses into a clause that requires they are both true.
// SQL implementation note: AND has higher precedence than OR.
func (wh Clause) And(exp Expression) Clause {
	cl, isClause := exp.(Clause)
	if isClause && wh.conjunction == and && cl.conjunction == and {
		return Clause{append(wh.wheres, cl.wheres...), and}
	} else if !isClause && wh.conjunction == and {
		return Clause{append(wh.wheres, exp), and}
	}
	return Clause{[]Expression{wh, exp}, and}
}

// Or combines two clauses into a clause that requires either is true.
// SQL implementation note: AND has higher precedence than OR.
func (wh Clause) Or(exp Expression) Clause {
	cl, isClause := exp.(Clause)
	if isClause && wh.conjunction == or && cl.conjunction == or {
		return Clause{append(wh.wheres, cl.wheres...), or}
	} else if !isClause && wh.conjunction == or {
		return Clause{append(wh.wheres, exp), or}
	}
	return Clause{[]Expression{wh, exp}, or}
}

// And combines some expressions into a clause that requires they are all true.
func And(exp ...Expression) Clause {
	return Clause{exp, and}
}

// Or combines some expressions into a clause that requires that any is true.
func Or(exp ...Expression) Clause {
	return Clause{exp, or}
}
