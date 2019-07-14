package parse

import (
	"fmt"
	"github.com/pkg/errors"
	"go/token"
	"golang.org/x/tools/go/packages"
	"os"
	"strings"
)

var PrintAST = false
var Debug = false
var depth = 0

func Parse(paths []string) (PackageStore, error) {
	fset := token.NewFileSet()
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedImports | packages.NeedDeps | packages.NeedSyntax | packages.NeedTypes,
		Fset: fset,
	}
	pkgs, err := packages.Load(cfg, paths...)
	if err != nil {
		return nil, err
	}

	pStore := make(PackageStore)
	for _, pkg := range pkgs {
		if len(pkg.Errors) > 0 {
			for _, e := range pkg.Errors {
				fmt.Fprintf(os.Stderr, "%v\n", e)
			}
			return nil, errors.New("Failed due syntax errors.")
		}

		pStore[pkg.Name] = PackageGroup{
			Pkg:   pkg.Types,
			Files: pkg.Syntax,
		}
	}
	return pStore, nil
}

func DevInfo(format string, args ...interface{}) {
	if Debug {
		in := strings.Repeat(" ", depth*2)
		fmt.Fprintf(os.Stdout, in+format, args...)
	}
}

func lessDeep() {
	depth -= 1
}
