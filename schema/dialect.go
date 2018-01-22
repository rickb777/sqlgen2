package schema

import "strings"

type Dialect interface {
	TableDDL(*TableDescription) string
	InsertDML(*TableDescription) string
	UpdateDML(*TableDescription) string
	DeleteDML(*TableDescription, FieldList) string
	TruncateDDL(tableName string, force bool) []string
	CreateTableSettings() string
	FieldAsColumn(*Field) string

	ReplacePlaceholders(sql string) string
	Placeholders(n int) string
	String() string
	Alias() string
}

//-------------------------------------------------------------------------------------------------

// AllDialects lists all currently-supported dialects.
var AllDialects = []Dialect{Sqlite, Postgres, Mysql}

// DialectNames gets all the currently-supported dialect names. These names are
// typically mixed-case.
func DialectNames() []string {
	dialects := make([]string, len(AllDialects))
	for i, d := range AllDialects {
		dialects[i] = d.String()
	}
	return dialects
}

// PickDialect finds a dialect that matches by name, ignoring letter case.
// It returns nil if not found.
func PickDialect(name string) Dialect {
	for _, d := range AllDialects {
		if strings.EqualFold(name, d.String()) || strings.EqualFold(name, d.Alias()) {
			return d
		}
	}
	return nil
}
