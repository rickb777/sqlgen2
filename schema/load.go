package schema

import (
	"fmt"
	. "github.com/acsellers/inflections"
	"github.com/rickb777/sqlgen2/sqlgen/parse"
	"github.com/rickb777/sqlgen2/sqlgen/parse/exit"
	"go/types"
	"strings"
)

type context struct {
	pkgStore parse.PackageStore
	indices  map[string]*Index
	table    *Table
}

func Load(pkgStore parse.PackageStore, pkg, name string) (*Table, error) {
	table := new(Table)

	// local map of indexes, used for quick lookups and de-duping.
	indices := map[string]*Index{}

	table.Type = name
	table.Name = Pluralize(Underscore(table.Type))

	str, tags := pkgStore.Find(pkg, name)
	ctx := &context{pkgStore, indices, table}
	ctx.examineStruct(str, pkg, name, tags, nil)
	return table, nil
}

//-------------------------------------------------------------------------------------------------

func (ctx *context) examineStruct(str *types.Struct, pkg, name string, tags map[string]parse.Tag, parent *Node) {
	parse.DevInfo("examineStruct %s %s\n  tags %v\n", pkg, name, tags)
	for j := 0; j < str.NumFields(); j++ {
		tField := str.Field(j)
		parse.DevInfo("    f%d: name:%-10s pkg:%s type:%-25s f:%v, e:%v, a:%v\n", j,
			tField.Name(), tField.Pkg().Name(), tField.Type(), tField.IsField(), tField.Exported(), tField.Anonymous())

		if !tField.Exported() {
			if tag, exists := tags[tField.Name()]; !exists || (exists && !tag.Skip) {
				fmt.Printf("Warning: %s.%s contains unexported field %q,"+
					" (perhaps it needs to be annotated with `sql:\"-\"`).\n", pkg, name, tField.Name())
			}
		}

		if tField.Anonymous() {
			ctx.convertEmbeddedNodeToFields(tField, pkg, parent)

		} else {
			ctx.convertLeafNodeToField(tField, pkg, tags, parent)
		}
	}
}

//-------------------------------------------------------------------------------------------------

func (ctx *context) convertEmbeddedNodeToFields(leaf *types.Var, pkg string, parent *Node) {
	str, tags := ctx.pkgStore.Find(pkg, leaf.Name())
	parse.DevInfo("convertEmbeddedNodeToFields %s %s\n", pkg, leaf.Name())
	node := &Node{Name: leaf.Name(), Parent: parent}
	ctx.examineStruct(str, pkg, leaf.Name(), tags, node)
}

//-------------------------------------------------------------------------------------------------

func (ctx *context) convertLeafNodeToField(leaf *types.Var, pkg string, tags map[string]parse.Tag, parent *Node) {
	field := &Field{}
	var ok bool
	field.Node, ok = ctx.convertLeafNodeToNode(leaf, pkg, tags, parent)
	if !ok {
		return
	}

	// Lookup the SQL column type
	field.SqlType = BLOB
	underlying := leaf.Type().Underlying()
	switch u := underlying.(type) {
	case *types.Basic:
		field.SqlType = mapKindToSqlType[u.Kind()]
		field.Type.Base = parse.Kind(u.Kind())

	case *types.Slice:
		field.Type.Base = parse.Slice
	}

	// substitute tag variables
	if tag, exists := tags[leaf.Name()]; exists {
		if tag.Skip {
			return
		}

		field.Tags = tag

		if tag.Primary {
			if ctx.table.Primary != nil {
				exit.Fail(1, "%s, %s: compound primary keys are not supported.\n",
					ctx.table.Primary.Type.Name, field.Type.Name)
			}
			ctx.table.Primary = field
		}

		if tag.Index != "" {
			index, ok := ctx.indices[tag.Index]
			if !ok {
				index = &Index{
					Name: tag.Index,
				}
				ctx.indices[index.Name] = index
				ctx.table.Index = append(ctx.table.Index, index)
			}
			index.Fields = append(index.Fields, field)
		}

		if tag.Unique != "" {
			index, ok := ctx.indices[tag.Index]
			if !ok {
				index = &Index{
					Name:   tag.Unique,
					Unique: true,
				}
				ctx.indices[index.Name] = index
				ctx.table.Index = append(ctx.table.Index, index)
			}
			index.Fields = append(index.Fields, field)
		}

		if tag.Type != "" {
			t, ok := mapStringToSqlType[tag.Type]
			if ok {
				field.SqlType = t
			}
		}

		field.Encode = mapTagToEncoding[tag.Encode]

		if tag.Name != "" {
			field.SqlName = strings.ToLower(tag.Name)
		}
	}

	if field.SqlName == "" {
		field.SqlName = Underscore(field.JoinParts("_"))
	}

	ctx.table.Fields = append(ctx.table.Fields, field)
}

//-------------------------------------------------------------------------------------------------

func (ctx *context) convertLeafNodeToNode(leaf *types.Var, pkg string, tags map[string]parse.Tag, parent *Node) (Node, bool) {
	node := Node{Name: leaf.Name(), Parent: parent}
	tp := Type{Pkg: "", Name: ""}

	lt := leaf.Type()
	//isPtr := false

	switch t := lt.(type) {
	case *types.Pointer:
		lt = t.Elem()
		tp.Base = parse.Ptr
		//isPtr = true
	}

	switch t := lt.(type) {
	case *types.Basic:
		tp.Name = t.Name()
		//case *types.Struct:
		//	tp.Name = t.String()
	case *types.Named:
		tObj := t.Obj()
		if tObj.Pkg().Name() != pkg {
			tp.Pkg = tObj.Pkg().Name()
		}
		tp.Name = tObj.Name()

		if str, ok := t.Underlying().(*types.Struct); ok {
			addStructTags(tags, str)
			ctx.examineStruct(str, pkg, leaf.Name(), tags, &node)
			return node, false
		}
	case *types.Array:
		tp.Name = t.String()
	case *types.Slice:
		switch el := t.Elem().(type) {
		case *types.Basic:
			tp.Name = t.String()
		case *types.Named:
			tnObj := el.Obj()
			parse.DevInfo("slice pkgname:%s pkgpath:%s name:%s\n", tnObj.Pkg().Name(), tnObj.Pkg().Path(), tnObj.Name())
			if tnObj.Pkg().Name() != pkg {
				tp.Pkg = tnObj.Pkg().Name()
			}
			tp.Name = tnObj.Name()
		}
	default:
		panic(fmt.Sprintf("%#v", lt))
	}

	node.Type = tp
	return node, true
}

func addStructTags(tags map[string]parse.Tag, str *types.Struct) {
	for i := 0; i < str.NumFields(); i++ {
		ts := str.Tag(i)
		tag, err := parse.ParseTag(ts)
		if err != nil {
			exit.Fail(2, "%s contains unparseable tag %q (%s)", str.String(), ts, err)
		}
		tags[str.Field(i).Name()] = *tag
	}

}