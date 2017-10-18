package parse

import (
	"fmt"
	"go/ast"
	"go/token"
	"os"
	"strings"
	"github.com/rickb777/sqlgen/parse/exit"
)

func FindImport(tp Type) string {
	case1 := fmt.Sprintf(`"%s"`, tp.Pkg)
	case2 := fmt.Sprintf(`/%s"`, tp.Pkg)

	for _, file := range files {
		for _, decl := range file.Decls {
			gen, ok := decl.(*ast.GenDecl)
			if ok && gen.Tok == token.IMPORT {
				for _, gs := range gen.Specs {
					spec, isImportSpec := gs.(*ast.ImportSpec)
					if isImportSpec {
						if spec.Path.Kind == token.STRING {
							// spec.Name must be nil because . or renamed imports are explicitly not supported
							if spec.Name != nil {
								exit.Fail(1,
									"import %s: implementation limitation: renamed imports are not supported.\n",
									spec.Path.Value)
							}
							if spec.Path.Value == case1 || strings.HasSuffix(spec.Path.Value, case2) {
								DevInfo("findImport %s -> found %s\n", tp.Pkg, spec.Path.Value)
								ln := len(spec.Path.Value) - 1
								return spec.Path.Value[1:ln]
							}
						}
					}
				}
			}
		}
	}

	fmt.Fprintf(os.Stderr, "Cannot find import '%s' in the source code (%+v).", tp.Pkg, tp)
	return ""
}

