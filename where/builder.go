package where

import "bytes"

const and = " AND "
const or = " OR "

func Null(column string) Condition {
	return Condition{column + " ISNULL", []interface{}{}}
}

func Eq(column string, value interface{}) Condition {
	return Condition{column + "=?", []interface{}{value}}
}

func Gt(column string, value interface{}) Condition {
	return Condition{column + ">?", []interface{}{value}}
}

func GtEq(column string, value interface{}) Condition {
	return Condition{column + ">=?", []interface{}{value}}
}

func Lt(column string, value interface{}) Condition {
	return Condition{column + "<?", []interface{}{value}}
}

func LtEq(column string, value interface{}) Condition {
	return Condition{column + "<=?", []interface{}{value}}
}

func NotEq(column string, value interface{}) Condition {
	return Condition{column + "<>?", []interface{}{value}}
}

func Between(column string, a, b interface{}) Condition {
	return Condition{column + " BETWEEN ? AND ?", []interface{}{a, b}}
}

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

func Not(el Element) not {
	return not{el}
}

//-------------------------------------------------------------------------------------------------

func (cl Condition) And(c2 Condition) Clause {
	return Clause{[]Condition{cl, c2}, nil, and}
}

func (cl Condition) Or(c2 Condition) Clause {
	return Clause{[]Condition{cl, c2}, nil, or}
}

func (wh Clause) And(c2 Condition) Clause {
	return Clause{append(wh.wheres, c2), wh.subclause, and}
}

func (wh Clause) Or(c2 Condition) Clause {
	return Clause{append(wh.wheres, c2), wh.subclause, or}
}
