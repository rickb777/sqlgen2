package schema

import (
	"fmt"
	"github.com/rickb777/sqlgen2/sqlgen/parse"
)

type Type struct {
	PkgPath string     // package name (full path)
	PkgName string     // package name (short name)
	Name    string     // name of source code type.
	Base    parse.Kind // underlying source code kind.
}

func (t Type) Type() string {
	if len(t.PkgName) > 0 {
		return fmt.Sprintf("%s.%s", t.PkgName, t.Name)
	} else {
		return t.Name
	}
}

func (t Type) String() string {
	return fmt.Sprintf("%s (%s)", t.Type(), t.Base)
}
