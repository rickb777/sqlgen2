package parse

import (
	"fmt"
	"go/ast"
)

type QualifiedIdent struct {
	Pkg        string // short variant; may be blank
	Identifier string
}

func NewQualifiedIdent(ident *ast.Ident) QualifiedIdent {
	if ident.Obj == nil {
		return QualifiedIdent{"", ident.Name}
	}
	//assume(ident.Obj.Kind == ast.Pkg, "ident.Obj %#v", ident.Obj)
	return QualifiedIdent{ident.Obj.Name, ident.Name}
}

func (t QualifiedIdent) String() string {
	if len(t.Pkg) > 0 {
		return fmt.Sprintf("%s.%s", t.Pkg, t.Identifier)
	} else {
		return t.Identifier
	}
}

//-------------------------------------------------------------------------------------------------

type TypeRef struct {
	Name QualifiedIdent
	Base Kind   // underlying source code kind.
}

func NewTypeRef(pkg, name string, base Kind) TypeRef {
	return TypeRef{QualifiedIdent{pkg, name}, base}
}

func (t TypeRef) String() string {
	return fmt.Sprintf("%s (%s)", t.Name, t.Base)
}

//-------------------------------------------------------------------------------------------------

func assume(predicate bool, info string, args ...interface{}) {
	if !predicate {
		panic(fmt.Sprintf("Incorrect assumption: " + info, args...))
	}
}