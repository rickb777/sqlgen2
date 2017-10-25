package schema

type sqlite struct {
	base
}

func newSQLite() Dialect {
	d := &sqlite{}
	d.base.Dialect = d
	return d
}

func (d *sqlite) Id() DialectId {
	return Sqlite
}

func sqliteColumn(f *Field) string {
	switch f.SqlType {
	case INTEGER:
		return "integer"
	case BOOLEAN:
		return "boolean"
	case BLOB:
		return "blob"
	case VARCHAR:
		return "text"
	default:
		return "text"
	}
}

func sqliteToken(v SqlToken) string {
	switch v {
	case AUTO_INCREMENT:
		return "autoincrement"
	case PRIMARY_KEY:
		return "primary key"
	default:
		return ""
	}
}
