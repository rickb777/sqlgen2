package demo

//go:generate sqlgen -type demo.User -o user_sql.go -v -z -prefix Db user.go

type User struct {
	Id     int64  `sql:"pk: true, auto: true"`
	Login  string `sql:"unique: user_login"`
	Email  string `sql:"unique: user_email"`
	Avatar string
	Active bool
	Admin  bool

	// oauth token and secret
	token  string
	secret string

	// randomly generated hash used to sign user
	// session and application tokens.
	hash string
}
