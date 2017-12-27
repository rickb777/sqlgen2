package demo

import (
	"math/big"
	"github.com/rickb777/sqlgen2"
)

// This example demonstrates
//   * indexes
//   * unexported fields

//go:generate sqlgen -type demo.User -o user_sql.go -v -prefix Db user.go

type User struct {
	Uid          int64   `sql:"pk: true, auto: true"`
	Login        string  `sql:"unique: user_login"`
	EmailAddress string  `sql:"unique: user_email"`
	Avatar       string
	Active       bool
	Admin        bool
	Fave         big.Int `sql:"encode: json"`
	LastUpdated  int64

	// oauth token and secret
	token  string
	secret string

	// randomly generated hash used to sign user
	// session and application tokens.
	hash string
}

func (u *User) PreInsert(sqlgen2.Execer) error {
	return nil
}

func (u *User) PreUpdate(sqlgen2.Execer) error {
	return nil
}
