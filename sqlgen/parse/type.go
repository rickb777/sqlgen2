package parse

import "fmt"

type Type struct {
	Pkg  string // package name (short variant)
	Name string // name of source code type.
	Base Kind   // underlying source code kind.
}

func (t Type) Type() string {
	if len(t.Pkg) > 0 {
		return fmt.Sprintf("%s.%s", t.Pkg, t.Name)
	} else {
		return t.Name
	}
}

func (t Type) String() string {
	return fmt.Sprintf("%s (%s)", t.Type(), t.Base)
}

