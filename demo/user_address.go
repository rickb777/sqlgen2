package demo

import "math/big"

// This example demonstrates
//   * join result

//go:generate sqlgen -type demo.UserAddress -kind Join -o user_address_sql.go -gofmt -v -join user_address.go role.go address.go

type UserAddress struct {
	Uid          int64  `sql:"pk: true, auto: true"`
	Name         string `sql:"unique: user_login"`
	EmailAddress string `sql:"nk: true"`
	Address      *AddressFields
	Avatar       *string
	Role         *Role `sql:"type: text, size: 20"`
	Active       bool
	Admin        bool
	Fave         *big.Int `sql:"encode: json"`
	LastUpdated  int64
}
