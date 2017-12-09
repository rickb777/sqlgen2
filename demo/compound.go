package demo

// This example demonstrates
//   * no primary key column (which is unusual but permitted)
//   * an index that spans multiple columns

//go:generate sqlgen -type demo.Compound -o compound_sql.go -v -prefix Db compound.go

type Compound struct {
	Alpha  string  `sql:"unique: alpha_beta"`
	Beta   string  `sql:"unique: alpha_beta"`
}

