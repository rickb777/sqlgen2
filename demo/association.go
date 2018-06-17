package demo

// This example demonstrates
//   * the use of pointers for optional items

//go:generate sqlgen -type demo.Association -o association_sql.go -all -v association.go category.go

type Association struct {
	Id       int64 `sql:"pk: true, auto: true"`
	Name     *string
	Quality  *QualName
	Ref1     *int64
	Ref2     *int64
	Category *Category
}

type QualName string
