package output

import (
	"io"
	"os"
)

type OsLike interface {
	Create(name string) (io.WriteCloser, error)
	MkdirAll(path string, perm os.FileMode) error
}

type stdOs int

func (stdOs) Create(name string) (io.WriteCloser, error) {
	return os.Create(name)
}

func (stdOs) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

var Os OsLike = stdOs(0)
