// Package require provides simple constraints to assist with detecting errors in database queries
// that arise from the wrong number of result (for example no result or too many results).
//
// The errors arising when requirements are not met are Sizer values.
package require

import "fmt"

// Requirement set an expectation on the outcome of a query.
type Requirement interface {
	errorIfNotSatisfiedBy(int64, string, string) error
	String() string
}

// ChainErrorIfQueryNotSatisfiedBy matches a requirement against the actual result size for
// a select query. The requirement may be nil in which case there will be no error.
// This function accepts an existing potential error, passig it on if not nil.
func ChainErrorIfQueryNotSatisfiedBy(err error, r Requirement, actual int64) error {
	if err != nil {
		return err
	}
	return ErrorIfQueryNotSatisfiedBy(r, actual)
}

// ErrorIfQueryNotSatisfiedBy matches a requirement against the actual result size for
// a select query. The requirement may be nil in which case there will be no error.
func ErrorIfQueryNotSatisfiedBy(r Requirement, actual int64) error {
	if r == nil {
		return nil
	}
	return r.errorIfNotSatisfiedBy(actual, "fetch", "got")
}

// ChainErrorIfExecNotSatisfiedBy matches a requirement against the actual result size for
// an exec query. The requirement may be nil in which case there will be no error.
// This function accepts an existing potential error, passig it on if not nil.
func ChainErrorIfExecNotSatisfiedBy(err error, r Requirement, actual int64) error {
	if err != nil {
		return err
	}
	return ErrorIfExecNotSatisfiedBy(r, actual)
}

// ErrorIfExecNotSatisfiedBy matches a requirement against the actual result size for
// an exec query. The requirement may be nil in which case there will be no error.
func ErrorIfExecNotSatisfiedBy(r Requirement, actual int64) error {
	if r == nil {
		return nil
	}
	return r.errorIfNotSatisfiedBy(actual, "change", "changed")
}

//-------------------------------------------------------------------------------------------------

// Exactly is a requirement that is met by a number matching exactly.
type Exactly int64

func (n Exactly) errorIfNotSatisfiedBy(actual int64, infinitive, pastpart string) error {
	if actual == int64(n) {
		return nil
	}
	return ErrWrongSize(actual, "expected to %s %d but %s %d", infinitive, n, pastpart, actual)
}

func (n Exactly) String() string {
	return fmt.Sprintf("exactly %d", n)
}

//-------------------------------------------------------------------------------------------------

// NoMoreThan is a requirement that is met by the actual results being no more than a specified value.
type NoMoreThan int64

func (n NoMoreThan) errorIfNotSatisfiedBy(actual int64, infinitive, pastpart string) error {
	if actual <= int64(n) {
		return nil
	}
	return ErrWrongSize(actual, "expected to %s no more than %d but %s %d", infinitive, n, pastpart, actual)
}

func (n NoMoreThan) String() string {
	return fmt.Sprintf("no more than %d", n)
}

// NoMoreThanOne is a requirement that is met by the actual results being no more than one.
var NoMoreThanOne = NoMoreThan(1)

//-------------------------------------------------------------------------------------------------

// AtLeast is a requirement that is met by the actual results being at least a specified value.
type AtLeast int64

func (n AtLeast) errorIfNotSatisfiedBy(actual int64, infinitive, pastpart string) error {
	if actual >= int64(n) {
		return nil
	}
	return ErrWrongSize(actual, "expected to %s at least %d but %s %d", infinitive, n, pastpart, actual)
}

func (n AtLeast) String() string {
	return fmt.Sprintf("at least %d", n)
}

// AtLeastOne is a requirement that is met by the actual results being at least one, i.e. not empty.
var AtLeastOne = AtLeast(1)

//-------------------------------------------------------------------------------------------------

// Quantifier is a requirement that is met by imprecise zero, singular or plural results. The
// value All will be automatically updated to match exactly some number known at call time.
type Quantifier int64

const (
	None Quantifier = iota
	One
	Many
)

func (q Quantifier) String() string {
	switch q {
	case None:
		return "none"
	case One:
		return "one"
	default:
		return "many"
	}
}

func (q Quantifier) errorIfNotSatisfiedBy(actual int64, infinitive, pastpart string) error {
	if actual == int64(q) {
		return nil
	}
	if actual > 1 && q == Many {
		return nil
	}
	return ErrWrongSize(actual, "expected to %s %s but %s %d", infinitive, q, pastpart, actual)
}

//-------------------------------------------------------------------------------------------------

// LateBound is a requirement that is updated to match exactly some number known at call time.
type LateBound int64

const (
	All LateBound = iota
)

func (v LateBound) errorIfNotSatisfiedBy(actual int64, infinitive, pastpart string) error {
	panic("Late-bound requirement used in an inappropriate context.")
}

func (v LateBound) String() string {
	return "all"
}
