package parse

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io"
	"log"
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
			group := Group{Owner: p}
			DevInfo("Reading %s (%d of %d paths)\n", p, j, len(paths))
			f, err := os.Open(p)
			if err != nil {
				return nil, err
			}
			group.Sources = append(group.Sources, Source{p, f})
			groups = append(groups, group)
		}
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

		// A Config controls various options of the type checker.
		// The defaults work fine except for one setting:
		// we must specify how to deal with imports.
		conf := types.Config{
			Importer:                 importer.Default(),
			DisableUnusedImportCheck: true,
		}

		// Type-check the package containing gFiles.
		pkg, err := conf.Check(group.Owner, fset, gFiles, nil)
		if err != nil {
			log.Fatal(err) // type error
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
						DevInfo("    f%d: name:%-10s pkg:%s type:%-25s f:%v, e:%v, a:%v\n", j,
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

//func findMatchingNodes(pkg, name string) (PackageStore, error) {
//	// find the type of interest and parse it
//	spec, pkg := findTypeDecl(pkg, name)
//	return examineSpec(pkg, spec)
//}

//func examineSpec(pkg string, spec *ast.TypeSpec) (*Node, error) {
//	DevInfo("examineSpec: spec %q\n", pkg)
//	depth += 1
//	defer lessDeep()
//
//	str, ok := spec.Type.(*ast.StructType)
//	if !ok {
//		return nil, fmt.Errorf("%q is not a struct type", spec.Name.Name)
//	}
//
//	node := &Node{
//		Name: spec.Name.String(),
//		Type: Type{pkg, spec.Name.String(), Struct},
//	}
//
//	err := buildNodesFromStructFields(node, str)
//	DevInfo("parsed1:\n%s\n", node.String())
//	return node, err
//}

//func buildNode(parent *Node, expr ast.Expr, name, tag string) (err error) {
//	DevInfo("buildNode: %s %s expr:%+v %s %q\n", parent.Name, parent.Type, expr, name, tag)
//	depth += 1
//	defer lessDeep()
//
//	switch e := expr.(type) {
//	case *ast.Ident:
//		return buildIdentNode(parent, "", e, name, tag)
//
//	case *ast.SelectorExpr:
//		return buildIdentNode(parent, e.X.(*ast.Ident).Name, e.Sel, name, tag)
//
//	case *ast.ArrayType:
//		return buildArrayNode(parent, e, name, tag)
//
//	case *ast.MapType:
//		return buildMapNode(parent, e, name, tag)
//
//	case *ast.StarExpr:
//		return buildPtrNode(parent, e, name, tag)
//	}
//
//	return nil
//}

//func baseType(enclosingPkg, name, typePkg, typeName string) (*Node, error) {
//
//	b, isSimple := SimpleTypes[typeName]
//	if isSimple {
//		DevInfo("baseType for %s is simple %s\n", typeName, b)
//		node := &Node{Name: name, Type: Type{typePkg, typeName, b}}
//		return node, nil
//	}
//
//	if typePkg == "" {
//		// type is not qualified so we need the containing package instead
//		typePkg = enclosingPkg
//	}
//
//	DevInfo("baseType for %s.%s\n", typePkg, typeName)
//	t, _ := findTypeDecl(typePkg, typeName)
//
//	switch tt := t.Type.(type) {
//	case *ast.Ident:
//		b = SimpleTypes[tt.Name]
//		node := &Node{Name: name, Type: Type{typePkg, typeName, b}}
//		return node, nil
//		//case *ast.StructType:
//		//tt.Fields.
//	}
//
//	return nil, fmt.Errorf("Cannot find match for %s.%s from %+v", typePkg, typeName, t)
//}

//func buildIdentNode(parent *Node, pkgqual string, ident *ast.Ident, name, tag string) (err error) {
//	if ident.Obj == nil {
//		// this case happens for an identifier with a non-denoted type, e.g.
//		//    Foo MyType
//
//		DevInfo("buildIdentNode: %s %s pkg:%q ident:%v (obj:nil) name:%s tag:%q\n", parent.Name, parent.Type, pkgqual, ident, name, tag)
//		depth += 1
//		defer lessDeep()
//
//		basic, err := baseType(parent.Type.Pkg, name, pkgqual, ident.Name)
//		if err != nil {
//			return err
//		}
//
//		basic.Tags, err = parseTag(tag)
//		if err != nil {
//			return err
//		}
//		parent.appendNode(basic)
//		return nil
//	}
//
//	// this case happens for an identifier with a denoted type, e.g.
//	//    Foo mypkg.MyType
//
//	DevInfo("buildIdentNode: %s %s pkg:%q ident:%v obj:%+v name:%s tag:%q\n", parent.Name, parent.Type, pkgqual, ident, *ident.Obj, name, tag)
//	depth += 1
//	defer lessDeep()
//
//	spec, ok := ident.Obj.Decl.(*ast.TypeSpec)
//	if !ok {
//		return invalidType(name)
//	}
//
//	structNode := &Node{
//		Name: name,
//		Type: Type{parent.Type.Pkg, ident.Name, Struct},
//	}
//
//	structNode.Tags, err = parseTag(tag)
//	if err != nil {
//		return err
//	}
//
//	parent.appendNode(structNode)
//
//	switch x := spec.Type.(type) {
//	case *ast.Ident:
//		structNode.Type.Base = SimpleTypes[x.Name]
//		return nil
//
//	case *ast.StructType:
//		return buildNodesFromStructFields(structNode, x)
//	}
//
//	return fmt.Errorf("unsupported '%#v'", spec)
//}

//func buildArrayNode(parent *Node, ident *ast.ArrayType, name, tag string) (err error) {
//	if ident.Len != nil {
//		return invalidType(name)
//	}
//
//	node := &Node{
//		Name: name,
//		Type: Type{"", fmt.Sprintf("[]%s", ident.Elt), Slice},
//	}
//
//	node.Tags, err = parseTag(tag)
//	if err != nil {
//		return err
//	}
//
//	if node.Type.Name == "[]byte" {
//		node.Type.Base = Bytes
//	}
//
//	parent.appendNode(node)
//	return nil
//}

//func buildMapNode(parent *Node, ident *ast.MapType, name, tag string) (err error) {
//	type_ := fmt.Sprintf("map[%s]%s", ident.Key, ident.Value)
//	node := &Node{Name: name, Type: Type{"", type_, Map}}
//	node.Tags, err = parseTag(tag)
//	if err != nil {
//		return err
//	}
//	parent.appendNode(node)
//	return nil
//}

//func buildPtrNode(parent *Node, ident *ast.StarExpr, name, tag string) (err error) {
//	innerIdent, ok := ident.X.(*ast.Ident)
//	if !ok {
//		return invalidType(name)
//	}
//
//	if innerIdent.Obj == nil || innerIdent.Obj.Decl == nil {
//		return invalidType(name)
//	}
//
//	spec, ok := innerIdent.Obj.Decl.(*ast.TypeSpec)
//	if !ok {
//		return invalidType(name)
//	}
//
//	node := &Node{Name: name, Type: Type{"", innerIdent.Name, Ptr}}
//	node.Tags, err = parseTag(tag)
//	if err != nil {
//		return err
//	}
//
//	if node.Tags.Skip {
//		return nil
//	}
//
//	parent.appendNode(node)
//	return buildNodesFromStructFields(node, spec.Type.(*ast.StructType))
//}

//func invalidType(name string) error {
//	return fmt.Errorf("%s is not a valid type", name)
//}

func DevInfo(format string, args ...interface{}) {
	if Debug {
		in := strings.Repeat(" ", depth*2)
		fmt.Fprintf(os.Stdout, in+format, args...)
	}
}

func lessDeep() {
	depth -= 1
}
