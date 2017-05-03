package output

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestNewOutput(t *testing.T) {
	cases := []struct {
		input, dirs, name, pkg string
	}{
		{"", ".", "", ""},
		{"-", ".", "-", ""},
		{"foo.go", ".", "foo.go", ""},
		{"bar/foo.go", "bar", "foo.go", "bar"},
		{"zob/bar/foo.go", "zob/bar", "foo.go", "bar"},
	}
	for i, c := range cases {
		o := NewOutput(c.input)
		if o.Dirs != c.dirs {
			t.Errorf("%d: Expected dirs %q, got %q", i, c.dirs, o.Dirs)
		}
		if o.Name != c.name {
			t.Errorf("%d: Expected name %q, got %q", i, c.name, o.Name)
		}
		if o.Pkg() != c.pkg {
			t.Errorf("%d: Expected pkg %q, got %q", i, c.pkg, o.Pkg())
		}
	}
}

func TestWriteNoAction(t *testing.T) {
	o := NewOutput("")
	content := bytes.NewBufferString("some content")
	result := &bytes.Buffer{}

	o.Write(content, result)

	if result.String() != "" {
		t.Errorf("Got %q", result.String())
	}
}

func TestWriteStdout(t *testing.T) {
	o := NewOutput("-")
	content := bytes.NewBufferString("some content")
	result := &bytes.Buffer{}

	o.Write(content, result)

	if result.String() != "some content" {
		t.Errorf("Got %q", result.String())
	}
}

func TestWriteFile_createError(t *testing.T) {
	Os = &stubOs{
		createErr: os.ErrInvalid,
	}
	o := NewOutput("foo.go")

	_, err := o.create()

	if err == nil {
		t.Errorf("Expected an error")
	}
}

func TestWriteFile_simpleFileWrite(t *testing.T) {
	result := &bytes.Buffer{}
	stub := &stubOs{
		createFile: &nopCloser{result},
	}
	Os = stub

	o := NewOutput("foo.go")
	content := bytes.NewBufferString("some content")

	o.Write(content, nil)

	if result.String() != "some content" {
		t.Errorf("Got %q", result.String())
	}
	if stub.createName != "./foo.go" {
		t.Errorf("Got %q", stub.createName)
	}
}

func TestWriteFile_createDirectoryAndFile(t *testing.T) {
	result := &bytes.Buffer{}
	stub := &stubOs{
		createFile: &nopCloser{result},
	}
	Os = stub

	o := NewOutput("bar/foo.go")
	content := bytes.NewBufferString("some content")

	o.Write(content, nil)

	if result.String() != "some content" {
		t.Errorf("Got %q", result.String())
	}
	if stub.createName != "bar/foo.go" {
		t.Errorf("Got %q", stub.createName)
	}
	if stub.mkdirAllPath != "bar" {
		t.Errorf("Got %q", stub.mkdirAllPath)
	}
}

//-------------------------------------------------------------------------------------------------

type stubOs struct {
	createFile io.WriteCloser
	createErr  error
	createName string

	mkdirAllPath string
	mkdirAllErr  error
}

func (s *stubOs) Create(name string) (io.WriteCloser, error) {
	s.createName = name
	return s.createFile, s.createErr
}

func (s *stubOs) MkdirAll(path string, perm os.FileMode) error {
	s.mkdirAllPath = path
	return s.mkdirAllErr
}

type nopCloser struct {
	io.Writer
}

func (nopCloser) Close() error { return nil }
