package parse

import (
	"fmt"
	"go/ast"
	"github.com/rickb777/sqlgen2/sqlgen/parse/exit"
	"bytes"
	"github.com/kortschak/utter"
	"os"
	"strings"
)

type Import struct {
	Name string
	Path string
}

type ImportList []Import

func (il ImportList) Find(shortPkg string) (Import, bool) {
	slashPkg := fmt.Sprintf(`/%s`, shortPkg)

	for _, i := range il {
		if i.Name == shortPkg || i.Path == shortPkg || strings.HasSuffix(i.Path, slashPkg) {
			return i, true
		}
	}

	fmt.Fprintf(os.Stderr, "Cannot find import %q in the source code.", shortPkg)
	return Import{}, false
}

func findTypeDecl(pkg, name string) (*ast.TypeSpec, string) {
	for _, file := range files {
		if file.Pkg == pkg {
			for _, spec := range file.Types {
				if spec.Name.String() == name {
					DevInfo("findTypeDecl %s.%s -> found %s %q\n", pkg, name, utter.Sdump(spec.Type), file.Pkg)
					return spec, file.Pkg
				}
			}
		}
	}

	if pkg != "" {
		pkg = pkg + "."
	}
	exit.Fail(1, "Cannot find '%s%s' in the source code. Should you add more source files to be parsed?\n%s\n", pkg, name, findAllTypeDecls())
	return nil, ""
}

func findAllTypeDecls() string {
	info := &bytes.Buffer{}

	for _, file := range files {
		fmt.Fprintf(info, "\nfile %s:\n  package %s\n", file.FilePath, file.Pkg)
		for _, spec := range file.Types {
			fmt.Fprintf(info, "    %s\n      %s\n", spec.Name, exprInfo(spec.Type))
		}
	}

	return info.String()
}

func exprInfo(expr ast.Expr) string {
	switch ident := expr.(type) {
	case *ast.Ident:
		return fmt.Sprintf("%s (%T)", ident.Name, ident)

	case *ast.SelectorExpr:
		if p2, ok := ident.X.(*ast.Ident); ok {
			return fmt.Sprintf("%s.%s %q %q (%T)", ident.X, ident.Sel.Name, p2, p2.Name, ident)
		}
		return fmt.Sprintf("%s.%s (%T)", ident.X, ident.Sel.Name, ident)
	}
	//case *ast.ArrayType:
	//	return buildArrayNode(parent, e, name, tag)
	//
	//case *ast.MapType:
	//	return buildMapNode(parent, e, name, tag)
	//
	//case *ast.StarExpr:
	//	return buildPtrNode(parent, e, name, tag)
	//}

	return fmt.Sprintf("%+v", expr)
}
