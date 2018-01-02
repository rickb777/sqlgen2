package schema

import (
	"github.com/rickb777/sqlgen2/sqlgen/parse"
	"strings"
)

type SqlEncode int

// List of vendor-specific keywords
const (
	ENCNONE SqlEncode = iota
	ENCJSON
	ENCTEXT
)

type SqlToken int

// List of vendor-specific keywords
const (
	AUTO_INCREMENT SqlToken = iota
	PRIMARY_KEY
)

type TableDescription struct {
	Type string
	Name string

	Fields  []*Field
	Index   []*Index
	Primary *Field // compound primaries are not supported
}

type Node struct {
	Name   string
	Type   Type
	Parent *Node
}

type Field struct {
	Node
	SqlName string
	Encode  SqlEncode
	Tags    parse.Tag
}

type Index struct {
	Name   string
	Unique bool

	Fields []*Field
}

func (t *TableDescription) HasLastInsertId() bool {
	return t.Primary != nil && t.Primary.Type.Base.IsInteger()
}

func (t *TableDescription) HasPrimaryKey() bool {
	return t.Primary != nil
}

func (t *TableDescription) NumColumnNames(withAuto bool) int {
	num := 0
	for _, f := range t.Fields {
		if withAuto || !f.Tags.Auto {
			num++
		}
	}
	return num
}

func (t *TableDescription) ColumnNames(withAuto bool) []string {
	names := make([]string, 0, len(t.Fields))
	for _, f := range t.Fields {
		if withAuto || !f.Tags.Auto {
			names = append(names, f.SqlName)
		}
	}
	return names
}

//-------------------------------------------------------------------------------------------------

func (i *Index) UniqueStr() string {
	if i.Unique {
		return "UNIQUE "
	}
	return ""
}

//-------------------------------------------------------------------------------------------------

// Parts gets the node containment chain as a sequence of names of parts.
func (node *Node) Parts() []string {
	d := 0
	for n := node; n != nil; n = n.Parent {
		d++
	}

	p := make([]string, d)
	d--
	for n := node; n != nil; n = n.Parent {
		p[d] = n.Name
		d--
	}
	return p
}

func (node *Node) JoinParts(delta int, sep string) string {
	parts := node.Parts()
	if delta > 0 {
		parts = parts[:len(parts)-delta]
	}
	return strings.Join(parts, sep)
}
