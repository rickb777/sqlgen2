// Package vanilla provides a re-usable table API.
package vanilla

// Record provides access to the primary key only; all other database columns are ignored.
// This is useful in situations where identity is the only concern.
type Record struct {
	Id       int64    `sql:"pk: true, auto: true"`
}
