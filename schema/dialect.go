package schema

type Dialect interface {
	TableDDL(*TableDescription) string
	InsertDML(*TableDescription) string
	UpdateDML(*TableDescription) string
	DeleteDML(*TableDescription, FieldList) string
	TruncateDDL(tableName string, force bool) []string
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
