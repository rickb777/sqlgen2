package sqlgen2

import (
	"database/sql"
	"fmt"
	"strings"
	"github.com/rickb777/sqlgen2/schema"
)

// Named creates NamedArg values; it is synonymous with sql.Named().
func Named(name string, value interface{}) sql.NamedArg {
	// This method exists because the go1compat promise
	// doesn't guarantee that structs don't grow more fields,
	// so unkeyed struct literals are a vet error. Thus, we don't
	// want to allow sql.NamedArg{name, value}.
	return sql.NamedArg{Name: name, Value: value}
}

// NamedArgString converts the argument to a string of the form "name=value".
func NamedArgString(arg sql.NamedArg) string {
	return fmt.Sprintf("%s=%v", arg.Name, arg.Value)
}

//-------------------------------------------------------------------------------------------------

// NamedArgList holds a slice of NamedArgs
type NamedArgList []sql.NamedArg

// Exists verifies that one or more elements of NamedArgList return true for the passed func.
func (list NamedArgList) Exists(fn func(sql.NamedArg) bool) bool {
	for _, v := range list {
		if fn(v) {
			return true
		}
	}
	return false
}

// Contains tests whether anything in the list has a certain name.
func (list NamedArgList) Contains(name string) bool {
	return list.Exists(func(f sql.NamedArg) bool {
		return f.Name == name
	})
}

// Find returns the first sql.NamedArg that returns true for some function.
// False is returned if none match.
func (list NamedArgList) Find(fn func(sql.NamedArg) bool) (sql.NamedArg, bool) {
	for _, v := range list {
		if fn(v) {
			return v, true
		}
	}

	var empty sql.NamedArg
	return empty, false
}

// FindByName finds the first item with a particular name.
func (list NamedArgList) FindByName(name string) (sql.NamedArg, bool) {
	return list.Find(func(f sql.NamedArg) bool {
		return f.Name == name
	})
}

//-------------------------------------------------------------------------------------------------

// MkString produces a string ontainin all the values separated by sep.
func (list NamedArgList) MkString(sep string) string {
	ss := make([]string, len(list))
	for i, v := range list {
		ss[i] = NamedArgString(v)
	}
	return strings.Join(ss, sep)
}

// String produces a string ontainin all the values separated by comma.
func (list NamedArgList) String() string {
	return list.MkString(", ")
}

// Names gets all the names.
func (list NamedArgList) Names() []string {
	ss := make([]string, len(list))
	for i, v := range list {
		ss[i] = v.Name
	}
	return ss
}

// Values gets all the valules
func (list NamedArgList) Values() []interface{} {
	ss := make([]interface{}, len(list))
	for i, v := range list {
		ss[i] = v.Value
	}
	return ss
}

// Assignments gets the assignment expressions.
func (list NamedArgList) Assignments(d schema.Dialect, from int) []string {
	ss := make([]string, len(list))
	switch d {
	case schema.Postgres:
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
