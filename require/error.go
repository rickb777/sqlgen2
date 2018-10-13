package require

import "fmt"

// Sizer is any type that provides the result set size. This is used here for
// errors where the actual size did not meet the requirements.
type Sizer interface {
	Size() int64
}

type wrongResultSize struct {
	size    int64
	message string
}

// ErrWrongSize returns an error based on the actual size received and a message
// describing the unsatisfied requirement. The returned value is both an error
// and a Sizer.
func ErrWrongSize(actualSize int64, format string, args ...interface{}) error {
	return wrongResultSize{
		size:    actualSize,
		message: fmt.Sprintf(format, args...),
	}
}

func (e wrongResultSize) Error() string {
	return e.message
}

// Size gets the actual size. Typically this is the number of records received.
func (e wrongResultSize) Size() int64 {
	return e.size
}

// IsNotFound tests the error and returns true only if the error is a wrong-size
// error and the actual size less than one.
func IsNotFound(err error) bool {
	s, ok := ActualResultSize(err)
	return ok && s < 1
}

// IsNotUnique tests the error and returns true only if the error is a wrong-size
// error and the actual size was more than one.
func IsNotUnique(err error) bool {
	s, ok := ActualResultSize(err)
	return ok && s > 1
}

// ActualResultSize tests the error and returns true only if the error is a wrong-size
// error, in which case it also returns the actual result size.
func ActualResultSize(err error) (int64, bool) {
	if err == nil {
		return 0, false
	}
	w, ok := err.(wrongResultSize)
	if !ok {
		return 0, false
	}
	return w.size, true
}
