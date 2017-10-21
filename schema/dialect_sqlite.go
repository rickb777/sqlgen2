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
	return SQLITE
}

func sqliteColumn(f *Field) string {
	switch f.Type {
	case INTEGER:
		return "INTEGER"
	case BOOLEAN:
		return "BOOLEAN"
	case BLOB:
		return "BLOB"
	case VARCHAR:
		return "TEXT"
	default:
		return "TEXT"
	}
}
