package where

import (
	"bytes"
	"fmt"
)

// QueryConstraint is a value that is appended to a SELECT statement.
type QueryConstraint interface {
	Build(dialect Dialect) string
}

func BuildQueryConstraint(qc QueryConstraint, dialect Dialect) string {
	if qc == nil {
		return ""
	}
	return qc.Build(dialect)
}

//-------------------------------------------------------------------------------------------------

type literal string

// Literal returns the literal string supplied, converting it to a QueryConstraint.
func Literal(sqlPart string) QueryConstraint {
	return literal(sqlPart)
}

func (qc literal) Build(dialect Dialect) string {
	return string(qc)
}

//-------------------------------------------------------------------------------------------------

type queryConstraint struct {
	orderBy       []string
	desc          bool
	limit, offset int
}

var _ QueryConstraint = &queryConstraint{}

func OrderBy(column ...string) *queryConstraint {
	return &queryConstraint{orderBy: column}
}

func Limit(n int) *queryConstraint {
	return &queryConstraint{limit: n}
}

func Offset(n int) *queryConstraint {
	return &queryConstraint{offset: n}
}

func (qc *queryConstraint) OrderBy(column ...string) *queryConstraint {
	qc.orderBy = column
	return qc
}

func (qc *queryConstraint) Asc() *queryConstraint {
	qc.desc = false
	return qc
}

func (qc *queryConstraint) Desc() *queryConstraint {
	qc.desc = true
	return qc
}

func (qc *queryConstraint) Limit(n int) *queryConstraint {
	qc.limit = n
	return qc
}

func (qc *queryConstraint) Offset(n int) *queryConstraint {
	qc.offset = n
	return qc
}

func (qc *queryConstraint) Build(dialect Dialect) string {
	b := &bytes.Buffer{}
	spacer := ""
	if len(qc.orderBy) > 0 {
		b.WriteString("ORDER BY ")
		comma := ""
		for _, col := range qc.orderBy {
			b.WriteString(comma)
			b.WriteString(col)
			comma = ","
		}
		if qc.desc {
			b.WriteString(" DESC")
		}
		spacer = " "
	}
	if qc.limit > 0 {
		fmt.Fprintf(b, "%sLIMIT %d", spacer, qc.limit)
		spacer = " "
	}
	if qc.offset > 0 {
		fmt.Fprintf(b, "%sOFFSET %d", spacer, qc.offset)
	}
	return b.String()
}
