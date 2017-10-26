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

// see https://dev.mysql.com/doc/refman/5.7/en/integer-types.html

func mysqlIntegerColumn(f *Field) string {
	switch f.Type.Base {
	case parse.Int, parse.Int64:
		return "bigint"
	case parse.Int8:
		return "tinyint"
	case parse.Int16:
		return "smallint"
	case parse.Int32:
		return "int"
	case parse.Uint, parse.Uint64:
		return "bigint unsigned"
	case parse.Uint8:
		return "tinyint unsigned"
	case parse.Uint16:
		return "smallint unsigned"
	case parse.Uint32:
		return "int unsigned"
	case parse.Float32:
		return "float"
	case parse.Float64:
		return "double"
	}
	return ""
}

func mysqlRealColumn(f *Field) string {
	switch f.Type.Base {
	case parse.Float32:
		return "float"
	case parse.Float64:
		return "double"
	}
	return ""
}

func mysqlColumn(f *Field) string {
	switch f.SqlType {
	case INTEGER:
		return mysqlIntegerColumn(f)
	case REAL:
		return mysqlRealColumn(f)
	case BOOLEAN:
		return "tinyint(1)"
	case BLOB:
		return "mediumblob"
	case VARCHAR:
		// assigns an arbitrary size if
		// none is provided.
		size := f.Size
		if size == 0 {
			size = 512
		}
		return fmt.Sprintf("varchar(%d)", size)
	}
	return ""
}

func mysqlToken(v SqlToken) string {
	switch v {
	case AUTO_INCREMENT:
		return "auto_increment"
	case PRIMARY_KEY:
		return "primary key"
	default:
		return ""
	}
}
