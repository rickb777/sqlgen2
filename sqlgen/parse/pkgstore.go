package parse

import (
	"github.com/rickb777/sqlgen2/sqlgen/parse/exit"
	"go/ast"
	"go/types"
)

type PackageGroup struct {
	Pkg   *types.Package
	Files []*ast.File
}

type PackageStore map[string]PackageGroup

func (st PackageStore) store(pkg *types.Package, files []*ast.File) {
	DevInfo("store %s %d\n", pkg, len(files))
	st[pkg.Name()] = PackageGroup{pkg, files}
}

func (st PackageStore) FindNamed(pkg, name string) *types.Named {
	DevInfo("Find %s %s\n", pkg, name)
	pkgGrp, exists := st[pkg]
	if !exists {
		//exit.Fail(5, "Unable to find package %q\n", pkg)
		return nil
	}

	scope := pkgGrp.Pkg.Scope()

	for i, n := range scope.Names() {
		if n == name {
			obj := scope.Lookup(n)
			t, ok := obj.Type().(*types.Named)
			DevInfo("  scope%d: %s %v\n", i, obj.Name(), obj.Type())
			if ok {
				o := t.Obj()
				ot := o.Type()
				otu := ot.Underlying()
				DevInfo("  %T %v\n", t, t)
				DevInfo("  %T %v\n", o, o)
				DevInfo("  %T %v\n", ot, ot)
				DevInfo("  %T %v\n", otu, otu)
				return t
			}
		}
	}

	return nil
}

func (st PackageStore) FindStruct(pkg, name string) (*types.Struct, map[string]Tag) {
	DevInfo("FindStruct %s %s\n", pkg, name)
	t := st.FindNamed(pkg, name)
	if t == nil {
		return nil, nil
	}

	o := t.Obj()
	ot := o.Type()
	otu := ot.Underlying()
	DevInfo("  %T %v\n", o, o)
	DevInfo("  %T %v\n", ot, ot)
	DevInfo("  %T %v\n", otu, otu)
	s, ok := otu.(*types.Struct)
	if !ok {
		return nil, nil
	}

	tags, err := findTags(st[pkg].Files, pkg, name)
	if err != nil {
		exit.Fail(4, "%s, %s: tag error: %s\n", pkg, name, err)
		return nil, nil
	}

	return s, tags
}

func (st PackageStore) FindTags(pkg, name string) map[string]Tag {
	tags, err := findTags(st[pkg].Files, pkg, name)
	if err != nil {
		exit.Fail(4, "%s, %s: tag error: %s\n", pkg, name, err)
		return nil
	}
	return tags
}

func findTags(files []*ast.File, pkg, name string) (map[string]Tag, error) {
	typeSpec, _ := findTypeDecl(files, pkg, name)
	if typeSpec == nil {
		return nil, nil
	}

	tags := make(map[string]Tag)

	switch st := typeSpec.Type.(type) {
	case *ast.StructType:
		err := findStructTags(files, pkg, name, st, tags)
		if err != nil {
			return nil, err
		}
	}

	return tags, nil
}

func findTypeDecl(files []*ast.File, pkg, name string) (*ast.TypeSpec, string) {
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
	exit.Fail(5, "Cannot find '%s%s' in the source code. Should you add more source files to be parsed?\n", pkg, name)
	return nil, ""
}

func findStructTags(files []*ast.File, pkg, name string, str *ast.StructType, tags map[string]Tag) error {
	DevInfo("findStructTags(%s %s str %d)\n", pkg, name, len(tags))

	for j, field := range str.Fields.List {
		if field.Tag != nil {
			if field.Names == nil {
				err := buildEmbeddedStruct(files, pkg, name, field.Type, "")
				if err != nil {
					return err
				}

			} else {
				for i, n := range field.Names {
					DevInfo("  tag.%d.%d %-12s = %s\n", j, i, n.Name, field.Tag.Value)
					tag, err := ParseTag(field.Tag.Value)
					if err != nil {
						return err
					}
					tags[n.Name] = *tag
				}
			}
		}
	}
	return nil
}

func buildEmbeddedStruct(files []*ast.File, pkg, name string, expr ast.Expr, tag string) (err error) {
	DevInfo("buildEmbeddedStruct expr:%#v tag:%q\n", expr, tag)
	depth += 1
	defer lessDeep()

	switch ident := expr.(type) {
	case *ast.Ident:
		DevInfo("     ident is (%T) %s\n", ident, ident.Name)

	case *ast.SelectorExpr:
		DevInfo("     ident is (%T) %s.%s\n", ident, ident.X, ident.Sel.Name)
		if p2, ok := ident.X.(*ast.Ident); ok {
			DevInfo("     %#v %s\n", p2, p2.Name)
			//embedded.Type.Base, err = baseType(p2.Name, ident.Sel.Name)
		}

	default:
		DevInfo("     ident is (%T) %#v -- unsupported\n", ident, ident)
	}

	//if err != nil {
	//	return err
	//}

	//t, _ := findTypeDecl(embedded.Type.Pkg, embedded.Name)
	//
	//str, ok := t.Type.(*ast.StructType)
	//if !ok {
	//	return fmt.Errorf("Syntax error: '%s.%s' is the wrong type %T.", embedded.Type.Pkg, embedded.Name, t.Type)
	//}
	//
	//err = buildNodesFromStructFields(embedded, str)
	//DevInfo("--parsed2:\n%s\n", embedded.String())
	//for _, n := range embedded.Nodes {
	//	parent.appendNode(n)
	//}

	return err
}
