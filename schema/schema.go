package schema

// List of basic types
const (
	INTEGER int = iota
	VARCHAR
	BOOLEAN
	REAL
	BLOB
	JSON
)

// List of vendor-specific keywords
const (
	AUTO_INCREMENT = iota
	PRIMARY_KEY
)

type Table struct {
	Name string

	Fields  []*Field
	Index   []*Index
	Primary *Field // compound primaries are not supported
}

type Field struct {
	GoName  string
	SqlName string
	Type    int
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
	return t.Primary != nil && t.Primary.Type == INTEGER
}

func (t *Table) HasPrimaryKey() bool {
	return t.Primary != nil
}
