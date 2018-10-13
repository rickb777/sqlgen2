package require

import "fmt"

// Sizer is any type that provides the result set size. This is used here for
// errors where the actual size did not meet the requirements.
type Sizer interface {
	Size() uint
}

type wrongResultSize struct {
	size    uint
	message string
}

func wrongSize(size uint, format string, args ...interface{}) error {
	return wrongResultSize{
		size: size,
		message: fmt.Sprintf(format, args...),
	}
}

func (e wrongResultSize) Error() string {
	return e.message
}

func (e wrongResultSize) Size() uint {
	return e.size
}

// IsNotFound tests the error and returns true only if the error is a wrong-size
// error and the actual size was zero.
func IsNotFound(err error) bool {
	s, ok := ActualResultSize(err)
	return ok && s == 1
}

// IsNotUnique tests the error and returns true only if the error is a wrong-size
// error and the actual size was more than one.
func IsNotUnique(err error) bool {
	s, ok := ActualResultSize(err)
	return ok && s > 1
}

// ActualResultSize tests the error and returns true only if the error is a wrong-size
// error, in which case it also returns the actual result size.
func ActualResultSize(err error) (uint, bool) {
	if err == nil {
		return 0, false
	}
	w, ok := err.(*wrongResultSize)
	if !ok {
		return 0, false
	}
	return w.size, true
}
