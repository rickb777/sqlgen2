// Package vanilla provides a re-usable table API.
package vanilla

// The generated SQL is then hand-modified so this normally remains commented out.
//xx:generate sqlgen -type vanilla.PrimaryKey -o vanilla_sql.go -read -delete -slice -v vanilla.go

// PrimaryKey provides access to the primary key only; all other database columns are ignored.
// This is useful in situations where identity is the only concern.
type PrimaryKey struct {
	Id       int64    `sql:"pk: true, auto: true"`
}
