package parse

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io"
	"os"
	"strings"
)

var PrintAST = false
var Debug = false
var depth = 0

type Group struct {
	Owner   string
	Sources []Source
}

type Source struct {
	Name string
	In   io.Reader
}

func Parse(paths []string) (PackageStore, error) {
	var groups []Group
	defer func() {
		for _, g := range groups {
			for _, s := range g.Sources {
				s.In.(*os.File).Close()
			}
		}
	}()

	hasNamedFiles := false

	for j, p := range paths {
		s, err := os.Stat(p)
		if err != nil {
			return nil, err
		}

		if s.IsDir() {
			group := Group{Owner: p}
			d, err := os.Open(p)
			if err != nil {
				return nil, err
			}

			names, err := d.Readdirnames(-1)
			if err != nil {
				return nil, err
			}

			goFiles := filterGoFiles(names)

			for i, name := range goFiles {
				path := p + "/" + name
				DevInfo("Reading %s (%d of %d paths, %d of %d go files)\n", path, j, len(paths), i, len(goFiles))
				f, err := os.Open(path)
				if err != nil {
					return nil, err
				}
				group.Sources = append(group.Sources, Source{path, f})
			}

			d.Close()
			groups = append(groups, group)

		} else {
			hasNamedFiles = true
		}
	}

	if hasNamedFiles {
		group := Group{}

		for j, p := range paths {
			s, err := os.Stat(p)
			if err != nil {
				return nil, err
			}

			if !s.IsDir() {
				DevInfo("Reading %s (%d of %d paths)\n", p, j, len(paths))
				f, err := os.Open(p)
				if err != nil {
					return nil, err
				}
				group.Sources = append(group.Sources, Source{p, f})
				group.Owner = p
			}
		}

		groups = append(groups, group)
	}

	fset := token.NewFileSet()
	return ParseGroups(fset, groups...)
}

func filterGoFiles(names []string) []string {
	goFiles := make([]string, 0, len(names))
	for _, name := range names {
		if isGoFile(name) {
			goFiles = append(goFiles, name)
		}
	}
	return goFiles
}

func isGoFile(name string) bool {
	return strings.HasSuffix(name, ".go") && !strings.HasSuffix(name, "_test.go")
}

func ParseGroups(fset *token.FileSet, groups ...Group) (PackageStore, error) {
	pStore := make(PackageStore)

	for _, group := range groups {
		var gFiles []*ast.File
		for _, p := range group.Sources {
			DevInfo("parsing: %s\n", p.Name)

			mode := parser.ParseComments
			if PrintAST {
				mode |= parser.Trace
			}

			file, err := parser.ParseFile(fset, p.Name, p.In, mode)
			if err != nil {
				return nil, err
			}

			gFiles = append(gFiles, file)
		}

		//sourceImporter := importer.Default()
		sourceImporter := importer.For("source", nil)

		// A Config controls various options of the type checker.
		// The defaults work fine except for one setting:
		// we must specify how to deal with imports.
		conf := types.Config{
			Importer: &importerWrapper{
				sourceImporter.(types.ImporterFrom),
			},
			IgnoreFuncBodies:         true,
			DisableUnusedImportCheck: true,
			FakeImportC:              true,
		}

		// Type-check the package containing gFiles.
		pkg, err := conf.Check(group.Owner, fset, gFiles, nil)
		if err != nil {
			//exit.Fail(3, "%s\n", err) // type error
			fmt.Fprintf(os.Stderr, "Warning: %s\n", err)
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
