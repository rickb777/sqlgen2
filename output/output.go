package output

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Output struct {
	Dirs, Name string
	Derived    bool // no stdout when true
}

func NewOutput(path string) Output {
	slash := strings.LastIndexByte(path, '/')

	if slash < 0 {
		return Output{
			Dirs:    ".",
			Name:    path,
			Derived: false,
		}
	}

	return Output{
		Dirs:    path[:slash],
		Name:    path[slash+1:],
		Derived: false,
	}
}

func (o Output) Derive(extn string) Output {
	if o.Name == "" || o.Name == "-" {
		d := o
		d.Derived = true
		return d
	}

	dot := strings.LastIndexByte(o.Name, '.')
	if dot < 0 {
		return Output{
			Dirs:    o.Dirs,
			Name:    o.Name + extn,
			Derived: true,
		}
	}

	return Output{
		Dirs:    o.Dirs,
		Name:    o.Name[:dot] + extn,
		Derived: true,
	}
}

func (o Output) Path() string {
	return fmt.Sprintf("%s/%s", o.Dirs, o.Name)
}

func (o Output) Pkg() string {
	if o.Dirs == "." {
		return ""
	}

	slash := strings.LastIndexByte(o.Dirs, '/')
	if slash < 0 {
		return o.Dirs
	}

	return o.Dirs[slash+1:]
}

func (o Output) create() (io.WriteCloser, error) {
	if o.Dirs != "." {
		err := Os.MkdirAll(o.Dirs, 0755)
		if err != nil && !os.IsExist(err) {
			return nil, err
		}
		if err == nil {
			Info("mkdir %s/\n", o.Dirs)
		}
	}

	out, err := Os.Create(o.Path())
	if err != nil {
		return nil, err
	}

	Info("writing %s/%s\n", o.Dirs, o.Name)

	return out, nil
}

func (o Output) Write(content io.Reader, out io.Writer) {
	if o.Name == "" {
		return
	}

	if o.Name != "-" {
		file, err := o.create()
		Require(err == nil, "%v\n", err)
		defer file.Close()
		io.Copy(file, content)
	} else if !o.Derived {
		io.Copy(out, content)
	}
}
