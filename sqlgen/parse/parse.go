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

var PrintAST = false
var Debug = false
var depth = 0
var fset = token.NewFileSet()
var files []SourceDecl

type Source struct {
	Name string
	In   io.Reader
}

func Parse(pkg, name string, path []string) (*Node, error) {
	var files []Source
	for _, p := range path {
		f, err := os.Open(p)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		files = append(files, Source{p, f})
	}
	return DoParse(pkg, name, files)
}

func DoParse(pkg, name string, files []Source) (*Node, error) {
	err := parseAllFiles(files)
	if err != nil {
		return nil, err
	}

	return findMatchingNodes(pkg, name)
}

func parseAllFiles(path []Source) (error) {
	files = make([]SourceDecl, 0, len(path))

	for _, p := range path {
		f, err := parse(p)
		if err != nil {
			return err
		}
		files = append(files, f)
	}

	return nil
}

func parse(path Source) (SourceDecl, error) {
	DevInfo("parsing: %s\n", path.Name)
	mode := parser.ParseComments
	if PrintAST {
		mode |= parser.Trace
	}
	f, err := parser.ParseFile(fset, path.Name, path.In, mode)
	if err != nil {
		return SourceDecl{}, err
	}
	return NewSourceDecl(path.Name, f), nil
}

func findMatchingNodes(pkg, name string) (*Node, error) {
	// find the type of interest and parse it
	spec, pkg := findTypeDecl(pkg, name)
	return examineSpec(pkg, spec)
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
		TypeRef: NewTypeRef(pkg, spec.Name.String(), Struct),
	}

	err := buildNodesFromStructFields(node, str)
	DevInfo("parsed1:\n%s\n", node.String())
	return node, err
}

func buildNodesFromStructFields(parent *Node, str *ast.StructType) (err error) {
	DevInfo("buildNodesFromStructFields: %s %s\n", parent.Name, parent.TypeRef)
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
				if !name.IsExported() {
					DevInfo("  %s is not exported\n", name)
					basic := &Node{Name: name.Name, TypeRef: NewTypeRef("foo", "", Scanner)}
					basic.Tags, err = parseTag(tag)
					if err != nil {
						return err
					}
					parent.appendNode(basic)
					return nil
				}
			}

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
	DevInfo("buildNode: %s %s expr:%+v %s %q\n", parent.Name, parent.TypeRef, expr, name, tag)
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
	DevInfo("buildEmbeddedStruct %s %s expr:%#v tag:%q\n", parent.Name, parent.TypeRef, expr, tag)
	depth += 1
	defer lessDeep()

	embedded := &Node{
		TypeRef: NewTypeRef(parent.TypeRef.Name.Pkg, "", 0),
	}

	switch ident := expr.(type) {
	case *ast.Ident:
		DevInfo("     ident is (%T) %s.%s\n", ident, parent.TypeRef.Name.Pkg, ident.Name)
		embedded.Name = ident.Name
		embedded.TypeRef.Name.Identifier = ident.Name
		//embedded.TypeRef.Base, err = baseType(parent.TypeRef.Pkg, embedded.Name)

	case *ast.SelectorExpr:
		DevInfo("     ident is (%T) %s.%s\n", ident, ident.X, ident.Sel.Name)
		if p2, ok := ident.X.(*ast.Ident); ok {
			DevInfo("     %#v %s\n", p2, p2.Name)
			embedded.Name = ident.Sel.Name
			embedded.TypeRef = NewTypeRef(p2.Name, ident.Sel.Name, 0)
			//embedded.TypeRef.Base, err = baseType(p2.Name, ident.Sel.Name)
		}

	default:
		DevInfo("     ident is (%T) %#v -- unsupported\n", ident, ident)
	}

	//if err != nil {
	//	return err
	//}

	t, _ := findTypeDecl(embedded.TypeRef.Name.Pkg, embedded.Name)

	str, ok := t.Type.(*ast.StructType)
	if !ok {
		return fmt.Errorf("Syntax error: '%s.%s' is the wrong type %T.", embedded.TypeRef.Name.Pkg, embedded.Name, t.Type)
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
		node := &Node{Name: name, TypeRef: NewTypeRef(typePkg, typeName, b)}
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
		node := &Node{Name: name, TypeRef: NewTypeRef(typePkg, typeName, b)}
		return node, nil
		//case *ast.StructType:
		//tt.Fields.
	}

	return nil, fmt.Errorf("Cannot find match for %s.%s from %+v", typePkg, typeName, t)
}

func buildIdentNode(parent *Node, pkgqual string, ident *ast.Ident, name, tag string) (err error) {
	if ident.Obj == nil {
		return buildIdentNodeWithNonDenotedType(parent, pkgqual, ident, name, tag)
	}

	return buildIdentNodeWithDenotedType(parent, pkgqual, ident, name, tag)
}

func buildIdentNodeWithNonDenotedType(parent *Node, pkgqual string, ident *ast.Ident, name, tag string) (err error) {
	// this case happens for an identifier with a non-denoted type, e.g.
	//    Foo MyType

	DevInfo("buildIdentNode: %s %s pkg:%q ident:%v (obj:nil) name:%s tag:%q\n", parent.Name, parent.TypeRef, pkgqual, ident, name, tag)
	depth += 1
	defer lessDeep()

	enclosingPkg := parent.TypeRef.Name.Pkg

	b, isSimple := SimpleTypes[ident.Name]
	if isSimple {
		DevInfo("baseType for %s is simple %s\n", ident.Name, b)
		basic := &Node{Name: name, TypeRef: NewTypeRef(pkgqual, ident.Name, b)}
		basic.Tags, err = parseTag(tag)
		if err != nil {
			return err
		}
		parent.appendNode(basic)
		return nil
	}

	if pkgqual == "" {
		// type is not qualified so we need the containing package instead
		pkgqual = enclosingPkg
	}

	DevInfo("baseType for %s.%s\n", pkgqual, ident.Name)
	t, _ := findTypeDecl(pkgqual, ident.Name)

	switch tt := t.Type.(type) {
	case *ast.Ident:
		b = SimpleTypes[tt.Name]
		basic := &Node{Name: name, TypeRef: NewTypeRef(pkgqual, ident.Name, b)}
		basic.Tags, err = parseTag(tag)
		if err != nil {
			return err
		}
		parent.appendNode(basic)
		return nil

	case *ast.StructType:
		structNode := &Node{
			Name: name,
			TypeRef: NewTypeRef(pkgqual, ident.Name, Struct),
		}

		structNode.Tags, err = parseTag(tag)
		if err != nil {
			return err
		}

		//parent.appendNode(structNode)
		return buildNodesFromStructFields(structNode, tt)
	}

	//if err != nil {
	//	return err
	//}

	//basic.Tags, err = parseTag(tag)
	//if err != nil {
	//	return err
	//}
	//parent.appendNode(basic)
	//return nil
	return fmt.Errorf("Cannot find match for %s.%s from %+v", pkgqual, ident.Name, t)
}

func buildIdentNodeWithDenotedType(parent *Node, pkgqual string, ident *ast.Ident, name, tag string) (err error) {
	// this case happens for an identifier with a denoted type, e.g.
	//    Foo mypkg.MyType

	DevInfo("buildIdentNode: %s %s pkg:%q ident:%v obj:%+v name:%s tag:%q\n", parent.Name, parent.TypeRef, pkgqual, ident, ident.Obj, name, tag)
	depth += 1
	defer lessDeep()

	structNode := &Node{
		Name: name,
		TypeRef: NewTypeRef(parent.TypeRef.Name.Pkg, ident.Name, Struct),
	}

	structNode.Tags, err = parseTag(tag)
	if err != nil {
		return err
	}

	parent.appendNode(structNode)

	spec, ok := ident.Obj.Decl.(*ast.TypeSpec)
	if !ok {
		return invalidType(name)
	}

	switch x := spec.Type.(type) {
	case *ast.Ident:
		structNode.TypeRef.Base = SimpleTypes[x.Name]
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
		TypeRef: NewTypeRef("", fmt.Sprintf("[]%s", ident.Elt), Slice),
	}

	node.Tags, err = parseTag(tag)
	if err != nil {
		return err
	}

	if node.TypeRef.Name.Identifier == "[]byte" {
		node.TypeRef.Base = Bytes
	}

	parent.appendNode(node)
	return nil
}

func buildMapNode(parent *Node, ident *ast.MapType, name, tag string) (err error) {
	type_ := fmt.Sprintf("map[%s]%s", ident.Key, ident.Value)
	node := &Node{Name: name, TypeRef: NewTypeRef("", type_, Map)}
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

	node := &Node{Name: name, TypeRef: NewTypeRef("", innerIdent.Name, Ptr)}
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
