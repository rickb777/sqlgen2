package schema

import "github.com/rickb777/sqlgen2/sqlgen/parse"

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

type Field struct {
	Name    string
	Type    parse.TypeRef
	Path    []string
	SqlName string
	SqlType SqlType
	Encode  SqlEncode
	Primary bool
	Auto    bool
	Size    int
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
		if withAuto || !f.Auto {
			num++
		}
	}
	return num
}

func (t *Table) ColumnNames(withAuto bool) []string {
	names := make([]string, 0, len(t.Fields))
	for _, f := range t.Fields {
		if withAuto || !f.Auto {
			names = append(names, f.SqlName)
		}
	}
	return names
}
