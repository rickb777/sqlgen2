package parse

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
)

var Debug = false
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
	return parseAll(pkg, name, files)
}

func parseAll(pkg, name string, path []file) (*Node, error) {
	files = make([]*ast.File, 0, len(path))

	for _, p := range path {
		file, err := parser.ParseFile(fset, p.name, p.in, parser.ParseComments)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}

	if spec, pkg := findTypeDecl(pkg, name); spec != nil {
		return parse(spec, pkg)
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
							return spec, file.Name.Name
						}
					}
				}
			}
		}
	}
	return nil, ""
}

func parse(spec *ast.TypeSpec, pkg string) (*Node, error) {
	node := &Node{
		Name: spec.Name.String(),
		Type: spec.Name.String(),
		Pkg:  pkg,
	}

	str, ok := spec.Type.(*ast.StructType)
	if !ok {
		return nil, fmt.Errorf("%q is not a struct type", spec.Name.Name)
	}

	err := buildNodes(node, str)
	if Debug {
		fmt.Print("parsed: ", node.String())
	}
	return node, err
}

func buildNodes(parent *Node, ident *ast.StructType) (err error) {
	for _, field := range ident.Fields.List {
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

func buildEmbeddedStruct(parent *Node, expr ast.Expr, tag string) error {
	pkg := parent.Pkg
	name := ""
	//fmt.Printf("anon %#v %q\n", expr, tag)
	switch ident := expr.(type) {
	case *ast.Ident:
		//fmt.Printf("     %T %s.%s\n", ident, parent.Pkg, ident.Name)
		name = ident.Name
	case *ast.SelectorExpr:
		//fmt.Printf("     %T %s.%s\n", ident, ident.X, ident.Sel.Name)
		if p2, ok := ident.X.(*ast.Ident); ok {
			//fmt.Printf("     %#v %s\n", p2, p2.Name)
			pkg = p2.Name
			name = ident.Sel.Name
		}
	}

	t, s := findTypeDecl(pkg, name)
	if Debug {
		fmt.Printf("%#v %s\n", t.Type, s)
	}
	if str, ok := t.Type.(*ast.StructType); ok {
		return buildNodes(parent, str)
	}
	return nil // TODO
}

func buildNode(parent *Node, expr ast.Expr, name, tag string) (err error) {
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
	return buildNodes(node, spec.Type.(*ast.StructType))
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
	return buildNodes(node, spec.Type.(*ast.StructType))
}

func invalidType(name string) error {
	return fmt.Errorf("%s is not a valid type", name)
}
