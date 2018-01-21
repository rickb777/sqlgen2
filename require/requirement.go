// Package require provides simple constraints to assist with detecting errors in database queries
// that arise from the wrong number of result (for example no result or too many results).
package require

import "fmt"

// Requirement set an expectation on the outcome of a query.
type Requirement interface {
	errorIfNotSatisfiedBy(uint, string, string) error
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
	return r.errorIfNotSatisfiedBy(uint(actual), "fetch", "got")
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
	return r.errorIfNotSatisfiedBy(uint(actual), "change", "changed")
}

//-------------------------------------------------------------------------------------------------

// Exactly is a requirement that is met by a number matching exactly.
type Exactly uint

func (n Exactly) errorIfNotSatisfiedBy(actual uint, infinitive, pastpart string) error {
	if actual == uint(n) {
		return nil
	}
	return fmt.Errorf("expected to %s %d but %s %d", infinitive, n, pastpart, actual)
}

//-------------------------------------------------------------------------------------------------

// NoMoreThan is a requirement that is met by the actual results being no more than a specified value.
type NoMoreThan uint

func (n NoMoreThan) errorIfNotSatisfiedBy(actual uint, infinitive, pastpart string) error {
	if actual <= uint(n) {
		return nil
	}
	return fmt.Errorf("expected to %s no more than %d but %s %d", infinitive, n, pastpart, actual)
}

// NoMoreThanOne is a requirement that is met by the actual results being no more than one.
var NoMoreThanOne = NoMoreThan(1)

//-------------------------------------------------------------------------------------------------

// AtLeast is a requirement that is met by the actual results being at least a specified value.
type AtLeast uint

func (n AtLeast) errorIfNotSatisfiedBy(actual uint, infinitive, pastpart string) error {
	if actual >= uint(n) {
		return nil
	}
	return fmt.Errorf("expected to %s at least %d but %s %d", infinitive, n, pastpart, actual)
}

// AtLeastOne is a requirement that is met by the actual results being at least one, i.e. not empty.
var AtLeastOne = AtLeast(1)

//-------------------------------------------------------------------------------------------------

// Quantifier is a requirement that is met by imprecise zero, singular or plural results. The
// value All will be automatically updated to match exactly some number known at call time.
type Quantifier uint

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

func (q Quantifier) errorIfNotSatisfiedBy(actual uint, infinitive, pastpart string) error {
	if actual == uint(q) {
		return nil
	}
	if actual > 1 && q == Many {
		return nil
	}
	return fmt.Errorf("expected to %s %s but %s %d", infinitive, q, pastpart, actual)
}

//-------------------------------------------------------------------------------------------------

// LateBound is a requirement that is updated to match exactly some number known at call time.
type LateBound uint

const (
	All LateBound = iota
)

func (v LateBound) errorIfNotSatisfiedBy(actual uint, infinitive, pastpart string) error {
	panic("Late-bound requirement used in an inappropriate context.")
}

