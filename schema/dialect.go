package schema

type Dialect interface {
	TableDDL(*TableDescription) string
	IndexDDL(*TableDescription, *Index) string
	InsertDML(*TableDescription) string
	UpdateDML(*TableDescription, []*Field) string
	DeleteDML(*TableDescription, []*Field) string
	//Param(int) string
	CreateTableSettings() string
	FieldAsColumn(*Field) string

	ReplacePlaceholders(sql string) string
	Placeholders(n int) string
	String() string
}

//-------------------------------------------------------------------------------------------------

var AllDialects = []Dialect{Sqlite, Postgres, Mysql}

var Dialects []string

func init() {
	Dialects = make([]string, len(AllDialects))
	for i, d := range AllDialects {
		Dialects[i] = d.String()
	}
}

func (f *Field) AsColumn(dialect Dialect) string {
	return dialect.FieldAsColumn(f)
}
