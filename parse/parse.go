package parse

import (
	"fmt"
	"github.com/rickb777/sqlgen2/parse/exit"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"strings"
)

var PrintAST = false
var Debug = false
var depth = 0

func Parse(paths []string) (PackageStore, error) {
	mode := parser.ParseComments
	if PrintAST {
		mode |= parser.Trace
	}

	var gFiles []*ast.File
	fset := token.NewFileSet()
	pkgName := "undefined"

	for j, p := range paths {
		s, err := os.Stat(p)
		if err != nil {
			return nil, err
		}

		if s.IsDir() {
			DevInfo("parsing directory %s (%d of %d)\n", p, j+1, len(paths))
			pkgs, err := parser.ParseDir(fset, p, isGoFileInfo, mode)
			if err != nil {
				return nil, err
			}

			for p, pkg := range pkgs {
				DevInfo("  %s:\n", p)
				pkgName = pkg.Name
				for n, file := range pkg.Files {
					DevInfo("    %s\n", n)
					gFiles = append(gFiles, file)
				}
			}

		} else {
			DevInfo("parsing file %s (%d of %d)\n", p, j+1, len(paths))
			f, err := os.Open(p)
			if err != nil {
				return nil, err
			}

			file, err := parser.ParseFile(fset, p, f, mode)
			f.Close()
			if err != nil {
				return nil, err
			}

			pkgName = file.Name.Name
			gFiles = append(gFiles, file)
		}
	}

	DevInfo("parsing complete\n")
	return TypeCheck(fset, pkgName, gFiles)
}

func isGoFileInfo(name os.FileInfo) bool {
	return isGoFile(name.Name())
}

func isGoFile(name string) bool {
	return strings.HasSuffix(name, ".go") && !strings.HasSuffix(name, "_test.go")
}

func TypeCheck(fset *token.FileSet, pkgName string, gFiles []*ast.File) (PackageStore, error) {
	pStore := make(PackageStore)

	mode := parser.ParseComments
	if PrintAST {
		mode |= parser.Trace
	}

	DevInfo("parsing complete\n")

	sourceImporter := importer.ForCompiler(fset, "source", nil)
	var errors []error

	// A Config controls various options of the type checker.
	// The defaults work fine except for one setting:
	// we must specify how to deal with imports.
	conf := types.Config{
		Importer: &importerWrapper{
			sourceImporter.(types.ImporterFrom),
		},
		IgnoreFuncBodies:         true,
		DisableUnusedImportCheck: true,
		//FakeImportC:              true,
		Error: func(err error) {
			DevInfo("%v\n", err)
			errors = append(errors, err)
			exit.Fail(3, "Failed: %v\n", err) // type error
		},
	}

	DevInfo("type-checking package %s...\n", pkgName)

	// Type-check the package containing gFiles.
	pkg, err := conf.Check(pkgName, fset, gFiles, nil)
	if err != nil {
		if pkg == nil || !pkg.Complete() {
			fmt.Fprintln(os.Stderr, "Syntax errors occurred:")
			for i, err := range errors {
				fmt.Fprintf(os.Stderr, "%d: %v\n", i+1, err)
			}
			exit.Fail(3, "Failed.\n") // type error
		}
	}

	pStore.store(pkg, gFiles)

	DevInfo("Package  %q\n", pkg.Path())
	DevInfo("Name:    %s\n", pkg.Name())
	DevInfo("Imports: %s\n", pkg.Imports())
	for i, im := range pkg.Imports() {
		DevInfo("  %d: %s\n", i, im)
	}
	DevInfo("Scope:   %s\n", pkg.Scope())
	for i, n := range pkg.Scope().Names() {
		obj := pkg.Scope().Lookup(n)
		t, ok := obj.Type().(*types.Named)
		DevInfo("  s%d: %s %v\n", i, obj.Name(), obj.Type())
		if ok {
			o := t.Obj()
			ot := o.Type()
			otu := ot.Underlying()
			DevInfo("  %T %v\n", o, o)
			DevInfo("  %T %v\n", ot, ot)
			DevInfo("  %T %v\n", otu, otu)
			s, ok := otu.(*types.Struct)
			if ok {
				for j := 0; j < s.NumFields(); j++ {
					f := s.Field(j)
					DevInfo("    f%-2d: name:%-15s pkg:%-10s type:%-50s field:%v, exp:%v, anon:%v\n", j,
						f.Name(), f.Pkg().Name(), f.Type(), f.IsField(), f.Exported(), f.Anonymous())
				}
			}
		}
		DevInfo("\n")
	}

	DevInfo("----------\n")
	return pStore, nil
}

type importerWrapper struct {
	inner types.ImporterFrom
}

func (i *importerWrapper) Import(path string) (*types.Package, error) {
	DevInfo("Import %s\n", path)
	return i.inner.Import(path)
}

func (i *importerWrapper) ImportFrom(path, dir string, mode types.ImportMode) (*types.Package, error) {
	DevInfo("ImportFrom %s %s %v\n", path, dir, mode)
	return i.inner.ImportFrom(path, dir, mode)
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
