package parse

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
	"strings"
	"github.com/rickb777/sqlgen/sqlgen/parse/exit"
)

var printAST = false
var Debug = false
var depth = 0
var fset = token.NewFileSet()
var files []*ast.File

type file struct {
	name string
	in   io.Reader
}

func Parse(pkg, name string, path []string) (*Node, error) {
	var files []file
	for _, p := range path {
		f, err := os.Open(p)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		files = append(files, file{p, f})
	}

	err := parseAllFiles(files)
	if err != nil {
		return nil, err
	}

	return findMatchingNodes(pkg, name)
}

func parseAllFiles(path []file) (error) {
	files = make([]*ast.File, 0, len(path))

	for _, p := range path {
		DevInfo("parsing: %s\n", p.name)
		mode := parser.ParseComments
		if Debug && printAST {
			mode |= parser.Trace
		}
		file, err := parser.ParseFile(fset, p.name, p.in, mode)
		if err != nil {
			return err
		}
		files = append(files, file)
	}
	return nil
}

func findMatchingNodes(pkg, name string) (*Node, error) {
	// find the type of interest and parse it
	spec, pkg := findTypeDecl(pkg, name)
	return examineSpec(pkg, spec)
}

func findTypeDecl(pkg, name string) (*ast.TypeSpec, string) {
	for _, file := range files {
		if file.Name.Name == pkg {
			for _, decl := range file.Decls {
				gen, isGenDecl := decl.(*ast.GenDecl)
				if isGenDecl {
					for _, gs := range gen.Specs {
						spec, isTypeSpec := gs.(*ast.TypeSpec)
						if isTypeSpec {
							if spec.Name.String() == name {
								DevInfo("findTypeDecl %s.%s -> found %#v %s\n", pkg, name, spec.Type, file.Name.Name)
								return spec, file.Name.Name
							}
						}
					}
				}
			}
		}
	}

	if pkg != "" {
		pkg = pkg + "."
	}
	exit.Fail(1, "Cannot find '%s%s' in the source code. Should you add more source files to be parsed?\n", pkg, name)
	return nil, ""
}

func examineSpec(pkg string, spec *ast.TypeSpec) (*Node, error) {
	DevInfo("examineSpec: spec %q\n", pkg)
	depth += 1
	defer lessDeep()

	str, ok := spec.Type.(*ast.StructType)
	if !ok {
		return nil, fmt.Errorf("%q is not a struct type", spec.Name.Name)
	}

	node := &Node{
		Name: spec.Name.String(),
		Type: Type{pkg, spec.Name.String(), Struct},
	}

	err := buildNodesFromStructFields(node, str)
	DevInfo("parsed1:\n%s\n", node.String())
	return node, err
}

func buildNodesFromStructFields(parent *Node, str *ast.StructType) (err error) {
	DevInfo("buildNodesFromStructFields: %s %s\n", parent.Name, parent.Type)
	depth += 1
	defer lessDeep()

	for _, field := range str.Fields.List {
		var tag string
		if field.Tag != nil {
			tag = field.Tag.Value
		}

		if field.Names == nil {
			err = buildEmbeddedStruct(parent, field.Type, tag)
			if err != nil {
				return err
			}

		} else {
			for _, name := range field.Names {
				err = buildNode(parent, field.Type, name.Name, tag)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func buildNode(parent *Node, expr ast.Expr, name, tag string) (err error) {
	DevInfo("buildNode: %s %s expr:%+v %s %q\n", parent.Name, parent.Type, expr, name, tag)
	depth += 1
	defer lessDeep()

	switch e := expr.(type) {
	case *ast.Ident:
		return buildIdentNode(parent, "", e, name, tag)

	case *ast.SelectorExpr:
		return buildIdentNode(parent, e.X.(*ast.Ident).Name, e.Sel, name, tag)

	case *ast.ArrayType:
		return buildArrayNode(parent, e, name, tag)

	case *ast.MapType:
		return buildMapNode(parent, e, name, tag)

	case *ast.StarExpr:
		return buildPtrNode(parent, e, name, tag)
	}

	return nil
}

func buildEmbeddedStruct(parent *Node, expr ast.Expr, tag string) (err error) {
	DevInfo("buildEmbeddedStruct %s %s expr:%#v tag:%q\n", parent.Name, parent.Type, expr, tag)
	depth += 1
	defer lessDeep()

	embedded := &Node{
		Type: Type{parent.Type.Pkg, "", 0},
	}

	switch ident := expr.(type) {
	case *ast.Ident:
		DevInfo("     ident is (%T) %s.%s\n", ident, parent.Type.Pkg, ident.Name)
		embedded.Name = ident.Name
		embedded.Type.Name = ident.Name
		//embedded.Type.Base, err = baseType(parent.Type.Pkg, embedded.Name)

	case *ast.SelectorExpr:
		DevInfo("     ident is (%T) %s.%s\n", ident, ident.X, ident.Sel.Name)
		if p2, ok := ident.X.(*ast.Ident); ok {
			DevInfo("     %#v %s\n", p2, p2.Name)
			embedded.Name = ident.Sel.Name
			embedded.Type = Type{p2.Name, ident.Sel.Name, 0}
			//embedded.Type.Base, err = baseType(p2.Name, ident.Sel.Name)
		}

	default:
		DevInfo("     ident is (%T) %#v -- unsupported\n", ident, ident)
	}

	//if err != nil {
	//	return err
	//}

	t, _ := findTypeDecl(embedded.Type.Pkg, embedded.Name)

	str, ok := t.Type.(*ast.StructType)
	if !ok {
		return fmt.Errorf("Syntax error: '%s.%s' is the wrong type %T.", embedded.Type.Pkg, embedded.Name, t.Type)
	}

	err = buildNodesFromStructFields(embedded, str)
	DevInfo("--parsed2:\n%s\n", embedded.String())
	for _, n := range embedded.Nodes {
		parent.appendNode(n)
	}

	return err
}

func baseType(enclosingPkg, name, typePkg, typeName string) (*Node, error) {

	b, isSimple := SimpleTypes[typeName]
	if isSimple {
		DevInfo("baseType for %s is simple %s\n", typeName, b)
		node := &Node{Name: name, Type: Type{typePkg, typeName, b}}
		return node, nil
	}

	if typePkg == "" {
		// type is not qualified so we need the containing package instead
		typePkg = enclosingPkg
	}

	DevInfo("baseType for %s.%s\n", typePkg, typeName)
	t, _ := findTypeDecl(typePkg, typeName)

	switch tt := t.Type.(type) {
	case *ast.Ident:
		b = SimpleTypes[tt.Name]
		node := &Node{Name: name, Type: Type{typePkg, typeName, b}}
		return node, nil
		//case *ast.StructType:
		//tt.Fields.
	}

	return nil, fmt.Errorf("Cannot find match for %s.%s from %+v", typePkg, typeName, t)
}

func buildIdentNode(parent *Node, pkgqual string, ident *ast.Ident, name, tag string) (err error) {
	if ident.Obj == nil {
		// this case happens for an identifier with a non-denoted type, e.g.
		//    Foo MyType

		DevInfo("buildIdentNode: %s %s pkg:%q ident:%v (obj:nil) name:%s tag:%q\n", parent.Name, parent.Type, pkgqual, ident, name, tag)
		depth += 1
		defer lessDeep()

		basic, err := baseType(parent.Type.Pkg, name, pkgqual, ident.Name)
		if err != nil {
			return err
		}

		basic.Tags, err = parseTag(tag)
		if err != nil {
			return err
		}
		parent.appendNode(basic)
		return nil
	}

	// this case happens for an identifier with a denoted type, e.g.
	//    Foo mypkg.MyType

	DevInfo("buildIdentNode: %s %s pkg:%q ident:%v obj:%+v name:%s tag:%q\n", parent.Name, parent.Type, pkgqual, ident, *ident.Obj, name, tag)
	depth += 1
	defer lessDeep()

	spec, ok := ident.Obj.Decl.(*ast.TypeSpec)
	if !ok {
		return invalidType(name)
	}

	structNode := &Node{
		Name: name,
		Type: Type{parent.Type.Pkg, ident.Name, Struct},
	}

	structNode.Tags, err = parseTag(tag)
	if err != nil {
		return err
	}

	parent.appendNode(structNode)

	switch x := spec.Type.(type) {
	case *ast.Ident:
		structNode.Type.Base = SimpleTypes[x.Name]
		return nil

	case *ast.StructType:
		return buildNodesFromStructFields(structNode, x)
	}

	return fmt.Errorf("unsupported '%#v'", spec)
}

func buildArrayNode(parent *Node, ident *ast.ArrayType, name, tag string) (err error) {
	if ident.Len != nil {
		return invalidType(name)
	}

	node := &Node{
		Name: name,
		Type: Type{"", fmt.Sprintf("[]%s", ident.Elt), Slice},
	}

	node.Tags, err = parseTag(tag)
	if err != nil {
		return err
	}

	if node.Type.Name == "[]byte" {
		node.Type.Base = Bytes
	}

	parent.appendNode(node)
	return nil
}

func buildMapNode(parent *Node, ident *ast.MapType, name, tag string) (err error) {
	type_ := fmt.Sprintf("map[%s]%s", ident.Key, ident.Value)
	node := &Node{Name: name, Type: Type{"", type_, Map}}
	node.Tags, err = parseTag(tag)
	if err != nil {
		return err
	}
	parent.appendNode(node)
	return nil
}

func buildPtrNode(parent *Node, ident *ast.StarExpr, name, tag string) (err error) {
	innerIdent, ok := ident.X.(*ast.Ident)
	if !ok {
		return invalidType(name)
	}

	if innerIdent.Obj == nil || innerIdent.Obj.Decl == nil {
		return invalidType(name)
	}

	spec, ok := innerIdent.Obj.Decl.(*ast.TypeSpec)
	if !ok {
		return invalidType(name)
	}

	node := &Node{Name: name, Type: Type{"", innerIdent.Name, Ptr}}
	node.Tags, err = parseTag(tag)
	if err != nil {
		return err
	}

	if node.Tags.Skip {
		return nil
	}

	parent.appendNode(node)
	return buildNodesFromStructFields(node, spec.Type.(*ast.StructType))
}

func invalidType(name string) error {
	return fmt.Errorf("%s is not a valid type", name)
}

func DevInfo(format string, args ...interface{}) {
	if Debug {
		in := strings.Repeat(" ", depth*2)
		fmt.Fprintf(os.Stderr, in+format, args...)
	}
}

func lessDeep() {
	depth -= 1
}
