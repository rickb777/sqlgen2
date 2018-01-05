package demo

// This example demonstrates
//   * the use of pointers for optional items
//   * field types that are repeated several times

//go:generate sqlgen -type demo.Association -o association_sql.go -v association.go category.go

// TODO FINISH THIS

type Association struct {
	Id       int64 `sql:"pk: true, auto: true"`
	Name     *string
	Quality  *string
	Ref1     *int64
	Ref2     *int64
	Category *Category
}
