package demo

// This example demonstrates
//   * the use of pointers for optional items

//go:generate sqlgen -type demo.Address -o address_sql.go -all -v address.go

type Address struct {
	Id       int64    `sql:"pk: true, auto: true"`
	Lines    []string `sql:"encode: json"`
	Postcode string   `sql:"size: 20, index: postcodeIdx"`
}
