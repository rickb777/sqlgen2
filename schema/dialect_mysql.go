package schema

import (
	"fmt"
	"github.com/rickb777/sqlgen2/sqlgen/parse"
	"io"
)

type mysql struct{}

var Mysql Dialect = mysql{}

func (d mysql) Index() int {
	return MysqlIndex
}

func (d mysql) String() string {
	return "Mysql"
}

func (d mysql) Alias() string {
	return "MySQL"
}

// see https://dev.mysql.com/doc/refman/5.7/en/data-types.html

func (dialect mysql) FieldAsColumn(field *Field) string {
	switch field.Encode {
	case ENCJSON:
		return "json"
	case ENCTEXT:
		return varchar(field.Tags.Size)
	}

	column := "mediumblob"

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
		column = "tinyint(1)"
	case parse.String:
		column = varchar(field.Tags.Size)
	}

	if field.Tags.Primary {
		column += " primary key"
	}

	if field.Tags.Auto {
		column += " auto_increment"
	}

	if field.Type.IsPtr {
		column += " default null"
	}

	return column
}

func varchar(size int) string {
	// Assigns an arbitrary size if none is provided.
	// 255 is chosen because the max. index key length is 767 and UTF8 might use up to three bytes per character.
	if size == 0 {
		size = 255
	}
	return fmt.Sprintf("varchar(%d)", size)
}

// see https://dev.mysql.com/doc/refman/5.7/en/integer-types.html

func (dialect mysql) TableDDL(table *TableDescription) string {
	return baseTableDDL(table, dialect, "\n", `"`)
}

func (dialect mysql) FieldDDL(w io.Writer, field *Field, comma string) string {
	return backTickFieldDDL(w, field, comma, dialect)
}

func (dialect mysql) InsertDML(table *TableDescription) string {
	return baseInsertDML(table, baseQueryPlaceholders(table.NumColumnNames(false)), backTickQuoted)
}

func (dialect mysql) UpdateDML(table *TableDescription) string {
	return baseUpdateDML(table, backTickQuoted, baseParamIsQuery)
}

func (dialect mysql) DeleteDML(table *TableDescription, fields FieldList) string {
	return baseDeleteDML(table, fields, backTickQuoted, baseParamIsQuery)
}

func (dialect mysql) TruncateDDL(tableName string, force bool) []string {
	truncate := fmt.Sprintf("TRUNCATE %s", tableName)
	if !force {
		return []string{truncate}
	}

	return []string{
		"SET FOREIGN_KEY_CHECKS=0",
		truncate,
		"SET FOREIGN_KEY_CHECKS=1",
	}
}

func (dialect mysql) SplitAndQuote(csv string) string {
	return baseSplitAndQuote(csv, "`", "`,`", "`")
}

func (dialect mysql) Quoter() func (Identifier) string {
	return backTickQuoted
}

func (dialect mysql) Placeholders(n int) string {
	return baseQueryPlaceholders(n)
}

// ReplacePlaceholders converts a string containing '?' placeholders to
// the form used by MySQL and SQLite - i.e. unchanged.
func (dialect mysql) ReplacePlaceholders(sql string) string {
	return sql
}

func (dialect mysql) CreateTableSettings() string {
	return " ENGINE=InnoDB DEFAULT CHARSET=utf8"
}

