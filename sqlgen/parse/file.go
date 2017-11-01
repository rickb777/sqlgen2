package parse

import (
	"go/ast"
	"go/token"
)

type SourceDecl struct {
	FilePath   string
	Pkg        string
	ImportList ImportList
	Types      map[QualifiedIdent]*ast.TypeSpec
	//Aliases    map[string]TypeDecl -- TODO
}

func NewSourceDecl(filePath string, content *ast.File) SourceDecl {
	il := make(ImportList, 0)
	ts := make(map[QualifiedIdent]*ast.TypeSpec, 0)
	pkg := content.Name.Name

	for _, spec := range content.Imports {
		assume(spec.Path.Kind == token.STRING, "spec.Path %#v", spec.Path)
		i := Import{}
		if spec.Name != nil {
			i.Name = spec.Name.Name
		}
		i.Path = removeQuotes(spec.Path.Value)
		il = append(il, i)
	}

	for _, decl := range content.Decls {
		gen, isGenDecl := decl.(*ast.GenDecl)
		if isGenDecl {
			for _, gs := range gen.Specs {
				switch spec := gs.(type) {
				case *ast.TypeSpec:
					ident := QualifiedIdent{pkg, spec.Name.Name}
					ts[ident] = spec
				}
			}
		}
	}

	return SourceDecl{filePath, pkg, il, ts}
}

func removeQuotes(q string) string {
	ln := len(q) - 1
	if ln < 1 {
		return q
	}
	return q[1:ln]
}
