package schema

import (
	"github.com/rickb777/sqlgen2/sqlgen/parse"
	"strings"
)

type SqlType int

// List of basic types
const (
	INTEGER SqlType = iota
	REAL
	VARCHAR
	BOOLEAN
	BLOB
	JSON
)

type SqlEncode int

// List of vendor-specific keywords
const (
	ENCNONE SqlEncode = iota
	ENCJSON
)

type SqlToken int

// List of vendor-specific keywords
const (
	AUTO_INCREMENT SqlToken = iota
	PRIMARY_KEY
)

type Table struct {
	Type string
	Name string

	Fields  []*Field
	Index   []*Index
	Primary *Field // compound primaries are not supported
}

type PathItem struct {
	Name string
	Ptr  bool
}

type Field struct {
	Name    string
	Type    parse.Type
	Path    Path
	SqlName string
	SqlType SqlType
	Encode  SqlEncode
	Tags    parse.Tag
}

type Index struct {
	Name   string
	Unique bool

	Fields []*Field
}

func (t *Table) HasLastInsertId() bool {
	return t.Primary != nil && t.Primary.SqlType == INTEGER
}

func (t *Table) HasPrimaryKey() bool {
	return t.Primary != nil
}

func (t *Table) NumColumnNames(withAuto bool) int {
	num := 0
	for _, f := range t.Fields {
		if withAuto || !f.Tags.Auto {
			num++
		}
	}
	return num
}

func (t *Table) ColumnNames(withAuto bool) []string {
	names := make([]string, 0, len(t.Fields))
	for _, f := range t.Fields {
		if withAuto || !f.Tags.Auto {
			names = append(names, f.SqlName)
		}
	}
	return names
}

//-------------------------------------------------------------------------------------------------

type Path []PathItem

func PathOf(items ...string) Path {
	r := make([]PathItem, 0, len(items))
	for _, s := range items {
		r = append(r, PathItem{s, false})
	}
	return Path(r)
}

func (p Path) Strings() []string {
	ss := make([]string, 0, len(p))
	for _, s := range p {
		ss = append(ss, s.Name)
	}
	return ss
}

func (p Path) Join(sep string) string {
	return strings.Join(p.Strings(), sep)
}
