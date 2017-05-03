package schema

type SqlType int

// List of basic types
const (
	INTEGER SqlType = iota
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
	Type string
	Name string

	Fields  []*Field
	Index   []*Index
	Primary *Field // compound primaries are not supported
}

type Field struct {
	Name    string
	SqlName string
	Type    SqlType
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

func (t *Table) ColumnNames(withAuto bool) []string {
	names := make([]string, 0, len(t.Fields))
	for _, f := range t.Fields {
		if withAuto || !f.Auto {
			names = append(names, f.SqlName)
		}
	}
	return names
}
