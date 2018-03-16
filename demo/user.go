package demo

import (
	"math/big"
)

// This example demonstrates
//   * indexes
//   * unexported fields

//go:generate sqlgen -type demo.User -table the_users -o user_sql.go -gofmt -v -prefix Db -all -setters=all user.go role.go

type User struct {
	Uid          int64    `sql:"pk: true, auto: true"`
	Name         string   `sql:"unique: user_login"`
	EmailAddress string   `sql:"nk: true"`
	AddressId    *int64   `sql:"fk: addresses.id, onupdate: restrict, ondelete: restrict"`
	Avatar       *string
	Role         *Role    `sql:"type: text, size: 20"`
	Active       bool
	Admin        bool
	Fave         *big.Int `sql:"encode: json"`
	LastUpdated  int64

	// something just to aid test coverage
	Numbers Numbers

	// oauth token and secret
	token  string
	secret string

	// randomly generated hash used to sign user
	// session and application tokens.
	hash string `sql:"-"`
}

type Numbers struct {
	I8  int8
	U8  uint8
	I16 int16
	U16 uint16
	I32 int32
	U32 uint32
	I64 int64
	U64 uint64
	F32 float32
	F64 float64
}

// These hooks are just used here for testing, but you could put whatever you like in them.
func (u *User) PreInsert() error {
	u.hash = "PreInsert"
	return nil
}

func (u *User) PreUpdate() error {
	u.hash = "PreUpdate"
	return nil
}

func (u *User) PostGet() error {
	u.hash = "PostGet"
	return nil
}
