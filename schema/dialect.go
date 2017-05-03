package schema

const (
	SQLITE int = iota
	POSTGRES
	MYSQL
)

var Dialects = map[string]int{
	"sqlite":   SQLITE,
	"postgres": POSTGRES,
	"mysql":    MYSQL,
}

type Dialect interface {
	Table(*Table) string
	Index(*Table, *Index) string
	Column(*Field) string
	Insert(*Table) string
	Update(*Table, []*Field) string
	Delete(*Table, []*Field) string
	Param(int) string
	Params(int, int) []string
	ColumnParams(t *Table, withAuto bool) string
	Token(int) string
}

func New(dialect int) Dialect {
	switch dialect {
	case POSTGRES:
		return newPostgres()
	case MYSQL:
		return newMySQL()
	default:
		return newSQLite()
	}
}
