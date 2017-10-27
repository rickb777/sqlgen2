package where

import "bytes"

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
	for i := range values {
		if i > 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte('?')
	}
	buf.WriteByte(')')
	return Condition{buf.String(), values}
}

//-------------------------------------------------------------------------------------------------

// Not negates an expression.
func Not(el Expression) not {
	return not{el}
}

//-------------------------------------------------------------------------------------------------

// And combines two conditions into a clause that requires they are both true.
// Implementation restriction: it is not yet possible to mix this with And().
func (cl Condition) And(c2 Condition) Clause {
	return Clause{[]Condition{cl, c2}, nil, and}
}

// Or combines two conditions into a clause that requires either is true.
// Implementation restriction: it is not yet possible to mix this with Or().
func (cl Condition) Or(c2 Condition) Clause {
	return Clause{[]Condition{cl, c2}, nil, or}
}

// And combines two clauses into a clause that requires they are both true.
// Implementation restriction: it is not yet possible to mix this with Or().
func (wh Clause) And(c2 Condition) Clause {
	return Clause{append(wh.wheres, c2), wh.subclause, and}
}

// Or combines two clauses into a clause that requires either is true.
// Implementation restriction: it is not yet possible to mix this with And().
func (wh Clause) Or(c2 Condition) Clause {
	return Clause{append(wh.wheres, c2), wh.subclause, or}
}
