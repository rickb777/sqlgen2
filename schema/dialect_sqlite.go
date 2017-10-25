package schema

type sqlite struct {
	base
}

func newSQLite() SDialect {
	d := &sqlite{}
	d.base.SDialect = d
	return d
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
