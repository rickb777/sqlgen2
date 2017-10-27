package where

import (
	"fmt"
	"strings"
	"github.com/rickb777/sqlgen2/dialect"
)

type Field struct {
	Name  string
	Value interface{}
}

func (f Field) String() string {
	return fmt.Sprintf("%s=%v", f.Name, f.Value)
}

//-------------------------------------------------------------------------------------------------

type FieldList []Field

// Exists verifies that one or more elements of FieldList return true for the passed func.
func (list FieldList) Exists(fn func(Field) bool) bool {
	for _, v := range list {
		if fn(v) {
			return true
		}
	}
	return false
}

func (list FieldList) Contains(name string) bool {
	return list.Exists(func(f Field) bool {
		return f.Name == name
	})
}

// Find returns the first Field that returns true for some function.
// False is returned if none match.
func (list FieldList) Find(fn func(Field) bool) (Field, bool) {
	for _, v := range list {
		if fn(v) {
			return v, true
		}
	}

	var empty Field
	return empty, false
}

func (list FieldList) FindByName(name string) (Field, bool) {
	return list.Find(func(f Field) bool {
		return f.Name == name
	})
}

func (list FieldList) MkString(sep string) string {
	ss := make([]string, len(list))
	for i, v := range list {
		ss[i] = v.String()
	}
	return strings.Join(ss, sep)
}

func (list FieldList) String() string {
	return list.MkString(", ")
}

func (list FieldList) Names() []string {
	ss := make([]string, len(list))
	for i, v := range list {
		ss[i] = v.Name
	}
	return ss
}

func (list FieldList) Assignments(d Dialect, from int) []string {
	ss := make([]string, len(list))
	switch d {
	case dialect.Postgres:
		for i, v := range list {
			ss[i] = fmt.Sprintf("%s=$%d", v.Name, i+from)
		}

	default:
		for i, v := range list {
			ss[i] = fmt.Sprintf("%s=?", v.Name)
		}
	}
	return ss
}

func (list FieldList) Values() []interface{} {
	ss := make([]interface{}, len(list))
	for i, v := range list {
		ss[i] = v.Value
	}
	return ss
}
