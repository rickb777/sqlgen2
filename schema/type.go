package schema

import (
	"fmt"
	. "github.com/rickb777/sqlgen2/sqlgen/parse"
)

type Type struct {
	PkgPath string // package name (full path)
	PkgName string // package name (short name)
	Name    string // name of source code type.
	IsPtr   bool
	Base    Kind // underlying source code kind.
}

func (t Type) Tag() string {
	if t.IsPtr {
		return t.Name + "Ptr"
	}
	return t.Name
}

func (t Type) Type() string {
	if len(t.PkgName) > 0 {
		return fmt.Sprintf("%s.%s", t.PkgName, t.Name)
	} else {
		return t.Name
	}
}

//func (t Type) NullableStorage() string {
//	if t.Base == parse.Ptr && t.PkgName == "" {
//		switch t.Name {
//		case "string":
//			return "sql.NullString"
//		case "int", "int8", "int16", "int32", "int64":
//			return "sql.NullInt64"
//		case "float32", "float54":
//			return "sql.NullFloat64"
//		case "bool":
//			return "sql.NullBool"
//		}
//	}
//	return t.Type()
//}

func (t Type) IsNullable() bool {
	return t.IsPtr && t.PkgName == ""
}

func (t Type) NullableValue() string {
	if t.IsNullable() {
		switch t.Base {
		case String:
			return "String"
		case Int, Int8, Int16, Int32, Int64:
			return "Int64"
		case Float32, Float64:
			return "Float64"
		case Bool:
			return "Bool"
		}
	}
	return ""
}

func (t Type) String() string {
	return fmt.Sprintf("%s (%v)", t.Type(), t.Base)
}
