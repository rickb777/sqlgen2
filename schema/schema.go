package schema

// List of basic types
const (
	INTEGER int = iota
	VARCHAR
	BOOLEAN
	REAL
	BLOB
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
	Primary []*Field
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
	return len(t.Primary) == 1 && t.Primary[0].Type == INTEGER
}

func (t *Table) PrimaryKeyFieldName() string {
	if len(t.Primary) != 1 {
		return ""
	}
	return t.Primary[0].GoName
}
