package schema

import (
	"github.com/rickb777/sqlgen2/sqlgen/parse"
	"fmt"
)

//import (
//	_ "github.com/mattn/go-sqlite3"
//)

type sqlite struct{}

var Sqlite Dialect = sqlite{}

func (d sqlite) String() string {
	return "Sqlite"
}

// For integers, the value is a signed integer, stored in 1, 2, 3, 4, 6, or 8 bytes depending on the magnitude of the value
// For reals, the value is a floating point value, stored as an 8-byte IEEE floating point number.

func (dialect sqlite) FieldAsColumn(field *Field) string {
	if field.Tags.Auto {
		// In sqlite, "autoincrement" is less efficient than built-in "rowid"
		// and the datatype must be "integer" (https://sqlite.org/autoinc.html).
		return "integer primary key autoincrement"
	}

	switch field.Encode {
	case ENCJSON:
		return "text"
	case ENCTEXT:
		return "text"
	}

	column := "blob"

	switch field.Type.Base {
	case parse.Int, parse.Int64:
		column = "bigint"
	case parse.Int8:
		column = "tinyint"
	case parse.Int16:
		column = "smallint"
	case parse.Int32:
		column = "int"
	case parse.Uint, parse.Uint64:
		column = "bigint unsigned"
	case parse.Uint8:
		column = "tinyint unsigned"
	case parse.Uint16:
		column = "smallint unsigned"
	case parse.Uint32:
		column = "int unsigned"
	case parse.Float32:
		column = "float"
	case parse.Float64:
		column = "double"
	case parse.Bool:
		column = "boolean"
	case parse.String:
		column = "text"
	}

	if field.Tags.Primary {
		column += " primary key"
	}

	if field.Type.IsPtr {
		column += " default null"
	}

	return column
}

func (dialect sqlite) TableDDL(table *TableDescription) string {
	return baseTableDDL(table, dialect)
}

func (dialect sqlite) InsertDML(table *TableDescription) string {
	return baseInsertDML(table)
}

func (dialect sqlite) UpdateDML(table *TableDescription) string {
	return baseUpdateDML(table, table.Fields, paramIsQuery)
}

func (dialect sqlite) DeleteDML(table *TableDescription, fields FieldList) string {
	return baseDeleteDML(table, fields, paramIsQuery)
}

func (dialect sqlite) TruncateDDL(tableName string, force bool) []string {
	truncate := fmt.Sprintf("DELETE FROM %s", tableName)
	return []string{truncate}
}

// Placeholders returns a string containing the requested number of placeholders
// in the form used by MySQL and SQLite.
func (dialect sqlite) Placeholders(n int) string {
	return queryPlaceholders(n)
}

// ReplacePlaceholders converts a string containing '?' placeholders to
// the form used by MySQL and SQLite - i.e. unchanged.
func (dialect sqlite) ReplacePlaceholders(sql string) string {
	return sql
}

func (dialect sqlite) CreateTableSettings() string {
	return ""
}
