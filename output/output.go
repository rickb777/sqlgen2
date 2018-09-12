package output

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Output struct {
	Dirs, Name string
}

func NewOutput(path string) Output {
	slash := strings.LastIndexByte(path, '/')

	if slash < 0 {
		return Output{
			Dirs: ".",
			Name: path,
		}
	}

	return Output{
		Dirs: path[:slash],
		Name: path[slash+1:],
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
		out = file
	}

	io.Copy(out, content)
}
