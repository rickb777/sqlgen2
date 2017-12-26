package schema

import (
	"fmt"
	"github.com/rickb777/sqlgen2/sqlgen/parse"
)

type mysql struct {
	base
}

func newMySQL() SDialect {
	d := &mysql{}
	d.base.SDialect = d
	return d
}

func (d *mysql) CreateTableSettings() string {
	return " ENGINE=InnoDB DEFAULT CHARSET=utf8"
}

// see https://dev.mysql.com/doc/refman/5.7/en/data-types.html

func mysqlColumn(field *Field) string {
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

	return column
}

func varchar(size int) string {
	// assigns an arbitrary size if
	// none is provided.
	if size == 0 {
		size = 512
	}
	return fmt.Sprintf("varchar(%d)", size)
}

// see https://dev.mysql.com/doc/refman/5.7/en/integer-types.html
