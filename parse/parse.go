package parse

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
	"strings"
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
	return parseAllFiles(pkg, name, files)
}

func parseAllFiles(pkg, name string, path []file) (*Node, error) {
	files = make([]*ast.File, 0, len(path))

	for _, p := range path {
		DevInfo("parsing: %s\n", p.name)
		file, err := parser.ParseFile(fset, p.name, p.in, parser.ParseComments)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}

	// find the type of interest and parse it
	if spec, pkg := findTypeDecl(pkg, name); spec != nil {
		return examineSpec(pkg, spec)
	}

	return nil, fmt.Errorf("Cannot find '%s.%s' in the source code.", pkg, name)
}

func findTypeDecl(pkg, name string) (*ast.TypeSpec, string) {
	for _, file := range files {
		if file.Name.Name == pkg {
			for _, decl := range file.Decls {
				if gen, ok := decl.(*ast.GenDecl); ok {
					if spec, ok := gen.Specs[0].(*ast.TypeSpec); ok {
						if spec.Name.String() == name {
							DevInfo("findTypeDecl %s.%s -> found %#v %s\n", pkg, name, spec.Type, file.Name.Name)
							return spec, file.Name.Name
						}
					}
				}
			}
		}
	}
	DevInfo("findTypeDecl %s.%s -> not found\n", pkg, name)
	return nil, ""
}

func examineSpec(pkg string, spec *ast.TypeSpec) (*Node, error) {
	DevInfo("examineSpec: spec %s\n", pkg)
	depth += 1
	defer lessDeep()

	node := &Node{
		Name: spec.Name.String(),
		Type: spec.Name.String(),
		Pkg:  pkg,
	}

	str, ok := spec.Type.(*ast.StructType)
	if !ok {
		return nil, fmt.Errorf("%q is not a struct type", spec.Name.Name)
	}

	err := buildNodesFromStructFields(node, str)
	DevInfo("parsed1:\n%s\n", node.String())
	return node, err
}

func buildNodesFromStructFields(parent *Node, str *ast.StructType) (err error) {
	DevInfo("buildNodesFromStructFields: %s %s.%s\n", parent.Name, parent.Pkg, parent.Type)
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
	DevInfo("buildNode: %s %s.%s expr %s %q\n", parent.Name, parent.Pkg, parent.Type, name, tag)
	depth += 1
	defer lessDeep()

	switch ident := expr.(type) {
	case *ast.Ident:
		return buildIdentNode(parent, ident, name, tag)

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
	DevInfo("buildEmbeddedStruct %s %s.%s %#v %q\n", parent.Name, parent.Pkg, parent.Type, expr, tag)
	depth += 1
	defer lessDeep()

	embedded := &Node{
		Name: "",
		Type: "",
		Pkg:  parent.Pkg,
	}

	switch ident := expr.(type) {
	case *ast.Ident:
		DevInfo("     %T %s.%s\n", ident, parent.Pkg, ident.Name)
		embedded.Name = ident.Name
		embedded.Type = ident.Name

	case *ast.SelectorExpr:
		DevInfo("     %T %s.%s\n", ident, ident.X, ident.Sel.Name)
		if p2, ok := ident.X.(*ast.Ident); ok {
			DevInfo("     %#v %s\n", p2, p2.Name)
			embedded.Pkg = p2.Name
			embedded.Name = ident.Sel.Name
			embedded.Type = ident.Sel.Name
		}

	default:
		DevInfo("     %T %#v -- unsupported\n", ident, ident)
	}

	t, _ := findTypeDecl(embedded.Pkg, embedded.Name)
	if t == nil {
		return fmt.Errorf("Cannot find '%s.%s' in the source code.", embedded.Pkg, embedded.Name)
	}

	str, ok := t.Type.(*ast.StructType)
	if !ok {
		return fmt.Errorf("Syntax error: '%s.%s' is the wrong type %T.", embedded.Pkg, embedded.Name, t.Type)
	}

	err := buildNodesFromStructFields(embedded, str)
	DevInfo("--parsed2:\n%s\n", embedded.String())
	for _, n := range embedded.Nodes {
		parent.appendNode(n)
	}

	return err
}

func buildIdentNode(parent *Node, ident *ast.Ident, name, tag string) (err error) {
	if ident.Obj == nil {
		node := &Node{
			Name: name,
			Type: ident.Name,
			Kind: Types[ident.Name],
		}
		node.Tags, err = parseTag(tag)
		if err != nil {
			return err
		}
		parent.appendNode(node)
		return nil
	}

	spec, ok := ident.Obj.Decl.(*ast.TypeSpec)
	if !ok {
		return invalidType(name)
	}

	node := &Node{
		Name: name,
		Type: ident.Name,
		Kind: Struct,
	}

	node.Tags, err = parseTag(tag)
	if err != nil {
		return err
	}

	parent.appendNode(node)
	return buildNodesFromStructFields(node, spec.Type.(*ast.StructType))
}

func buildArrayNode(parent *Node, ident *ast.ArrayType, name, tag string) (err error) {
	if ident.Len != nil {
		return invalidType(name)
	}

	node := &Node{
		Name: name,
		Kind: Slice,
		Type: fmt.Sprintf("[]%s", ident.Elt),
	}

	node.Tags, err = parseTag(tag)
	if err != nil {
		return err
	}

	if node.Type == "[]byte" {
		node.Kind = Bytes
	}

	parent.appendNode(node)
	return nil
}

func buildMapNode(parent *Node, ident *ast.MapType, name, tag string) (err error) {
	type_ := fmt.Sprintf("map[%s]%s", ident.Key, ident.Value)
	node := &Node{Name: name, Type: type_, Kind: Map}
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

	node := &Node{Name: name, Type: innerIdent.Name, Kind: Ptr}
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
		fmt.Fprintf(os.Stderr, in + format, args...)
	}
}

func lessDeep() {
	depth -= 1
}
