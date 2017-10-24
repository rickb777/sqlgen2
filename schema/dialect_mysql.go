package schema

import (
	"fmt"
	"github.com/rickb777/sqlgen2/sqlgen/parse"
)

type mysql struct {
	base
}

func newMySQL() Dialect {
	d := &mysql{}
	d.base.Dialect = d
	return d
}

func (d *mysql) Id() DialectId {
	return Mysql
}

func (d *mysql) CreateTableSettings() string {
	return " ENGINE=InnoDB DEFAULT CHARSET=utf8"
}

// see https://dev.mysql.com/doc/refman/5.7/en/integer-types.html

func mysqlIntegerColumn(f *Field) string {
	switch f.Type.Base {
	case parse.Int, parse.Int64:
		return "BIGINT"
	case parse.Int8:
		return "TINYINT"
	case parse.Int16:
		return "SMALLINT"
	case parse.Int32:
		return "INT"
	case parse.Uint, parse.Uint64:
		return "BIGINT UNSIGNED"
	case parse.Uint8:
		return "TINYINT UNSIGNED"
	case parse.Uint16:
		return "SMALLINT UNSIGNED"
	case parse.Uint32:
		return "INT UNSIGNED"
	case parse.Float32:
		return "FLOAT"
	case parse.Float64:
		return "DOUBLE"
	}
	return ""
}

func mysqlRealColumn(f *Field) string {
	switch f.Type.Base {
	case parse.Float32:
		return "FLOAT"
	case parse.Float64:
		return "DOUBLE"
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
		return "TINYINT(1)"
	case BLOB:
		return "MEDIUMBLOB"
	case VARCHAR:
		// assigns an arbitrary size if
		// none is provided.
		size := f.Size
		if size == 0 {
			size = 512
		}
		return fmt.Sprintf("VARCHAR(%d)", size)
	}
	return ""
}

func mysqlToken(v SqlToken) string {
	switch v {
	case AUTO_INCREMENT:
		return "AUTO_INCREMENT"
	case PRIMARY_KEY:
		return "PRIMARY KEY"
	default:
		return ""
	}
}
