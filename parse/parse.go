package parse

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
	"strings"
	"github.com/rickb777/sqlgen/parse/exit"
)

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
		file, err := parser.ParseFile(fset, p.name, p.in, parser.ParseComments)
		if err != nil {
			return err
		}
		files = append(files, file)
	}
	return nil
}

func findMatchingNodes(pkg, name string) (*Node, error) {
	// find the type of interest and parse it
	spec, pkg, err := findTypeDecl(pkg, name)
	if err != nil {
		return nil, err
	}

	return examineSpec(pkg, spec)
}

func FindImport(shortName string) string {
	case1 := fmt.Sprintf(`"%s"`, shortName)
	case2 := fmt.Sprintf(`/%s"`, shortName)

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
								DevInfo("findImport %s -> found %s\n", shortName, spec.Path.Value)
								ln := len(spec.Path.Value) - 1
								return spec.Path.Value[1:ln]
							}
						}
					}
				}
			}
		}
	}

	fmt.Fprintf(os.Stderr, "Cannot find import '%s' in the source code.", shortName)
	return ""
}

func findTypeDecl(pkg, name string) (*ast.TypeSpec, string, error) {
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
								return spec, file.Name.Name, nil
							}
						}
					}
				}
			}
		}
	}

	return nil, "", fmt.Errorf("Cannot find '%s.%s' in the source code.", pkg, name)
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
		} else {
			err = buildNode(parent, field.Type, field.Names[0].Name, tag)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func buildNode(parent *Node, expr ast.Expr, name, tag string) (err error) {
	DevInfo("buildNode: %s %s expr %s %q\n", parent.Name, parent.Type, name, tag)
	depth += 1
	defer lessDeep()

	switch ident := expr.(type) {
	case *ast.Ident:
		return buildIdentNode(parent, "", ident, name, tag)

	case *ast.SelectorExpr:
		return buildIdentNode(parent, ident.X.(*ast.Ident).Name, ident.Sel, name, tag)

	case *ast.ArrayType:
		return buildArrayNode(parent, ident, name, tag)

	case *ast.MapType:
		return buildMapNode(parent, ident, name, tag)

	case *ast.StarExpr:
		return buildPtrNode(parent, ident, name, tag)
	}

	return nil
}

func buildEmbeddedStruct(parent *Node, expr ast.Expr, tag string) error {
	DevInfo("buildEmbeddedStruct %s %s %#v %q\n", parent.Name, parent.Type, expr, tag)
	depth += 1
	defer lessDeep()

	embedded := &Node{
		Type: Type{parent.Type.Pkg, "", 0},
	}

	switch ident := expr.(type) {
	case *ast.Ident:
		DevInfo("     %T %s.%s\n", ident, parent.Type.Pkg, ident.Name)
		embedded.Name = ident.Name
		embedded.Type.Name = ident.Name

	case *ast.SelectorExpr:
		DevInfo("     %T %s.%s\n", ident, ident.X, ident.Sel.Name)
		if p2, ok := ident.X.(*ast.Ident); ok {
			DevInfo("     %#v %s\n", p2, p2.Name)
			embedded.Name = ident.Sel.Name
			embedded.Type = Type{p2.Name, ident.Sel.Name, 0}
		}

	default:
		DevInfo("     %T %#v -- unsupported\n", ident, ident)
	}

	t, _, err := findTypeDecl(embedded.Type.Pkg, embedded.Name)
	if err != nil {
		return err
	}

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

func buildIdentNode(parent *Node, pkgqual string, ident *ast.Ident, name, tag string) (err error) {
	if ident.Obj == nil {
		b, isSimple := SimpleTypes[ident.Name]
		basic := &Node{
			Name: name,
			Type: Type{pkgqual, ident.Name, b},
		}

		if !isSimple {
			t, _, err := findTypeDecl(pkgqual, ident.Name)
			if err != nil {
				return err
			}
			basic.Type.Base = SimpleTypes[t.Type.(*ast.Ident).Name]
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
