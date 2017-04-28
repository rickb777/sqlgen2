
package demo

// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

import (
	"database/sql"
)

func ScanUser(row *sql.Row) (*User, error) {
	var v0 int64
var v1 string
var v2 string
var v3 string
var v4 bool
var v5 bool
var v6 string
var v7 string
var v8 string


	err := row.Scan(
		&v0,
&v1,
&v2,
&v3,
&v4,
&v5,
&v6,
&v7,
&v8,

	)
	if err != nil {
		return nil, err
	}

	v := &User{}
	v.Id=v0
v.Login=v1
v.Email=v2
v.Avatar=v3
v.Active=v4
v.Admin=v5
v.token=v6
v.secret=v7
v.hash=v8


	return v, nil
}

func ScanUsers(rows *sql.Rows) ([]*User, error) {
	var err error
	var vv []*User

	var v0 int64
var v1 string
var v2 string
var v3 string
var v4 bool
var v5 bool
var v6 string
var v7 string
var v8 string

	for rows.Next() {
		err = rows.Scan(
			&v0,
&v1,
&v2,
&v3,
&v4,
&v5,
&v6,
&v7,
&v8,

		)
		if err != nil {
			return vv, err
		}

		v := &User{}
		v.Id=v0
v.Login=v1
v.Email=v2
v.Avatar=v3
v.Active=v4
v.Admin=v5
v.token=v6
v.secret=v7
v.hash=v8

		vv = append(vv, v)
	}
	return vv, rows.Err()
}

func SliceUser(v *User) []interface{} {
	var v0 int64
var v1 string
var v2 string
var v3 string
var v4 bool
var v5 bool
var v6 string
var v7 string
var v8 string

	v0=v.Id
v1=v.Login
v2=v.Email
v3=v.Avatar
v4=v.Active
v5=v.Admin
v6=v.token
v7=v.secret
v8=v.hash


	return []interface{}{
		v0,
v1,
v2,
v3,
v4,
v5,
v6,
v7,
v8,

	}
}

func SelectUser(db *sql.DB, query string, args ...interface{}) (*User, error) {
	row := db.QueryRow(query, args...)
	return ScanUser(row)
}

func SelectUsers(db *sql.DB, query string, args ...interface{}) ([]*User, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanUsers(rows)
}

func InsertUser(db *sql.DB, query string, v *User) error {
	res, err := db.Exec(query, SliceUser(v)[1:]...)
	if err != nil {
		return err
	}

	v.Id, err = res.LastInsertId()
	return err
}

func UpdateUser(db *sql.DB, query string, v *User) error {
	args := SliceUser(v)[1:]
	args = append(args, v.Id)
	_, err := db.Exec(query, args...)
	return err 
}

const CreateUserStmt = `
CREATE TABLE IF NOT EXISTS users (
 id     INTEGER PRIMARY KEY AUTOINCREMENT,
 login  TEXT,
 email  TEXT,
 avatar TEXT,
 active BOOLEAN,
 admin  BOOLEAN,
 token  TEXT,
 secret TEXT,
 hash   TEXT
);
`

const InsertUserStmt = `
INSERT INTO users (
 login,
 email,
 avatar,
 active,
 admin,
 token,
 secret,
 hash
) VALUES (?,?,?,?,?,?,?,?)
`

const SelectUserStmt = `
SELECT 
 id,
 login,
 email,
 avatar,
 active,
 admin,
 token,
 secret,
 hash
FROM users 
`

const SelectUserRangeStmt = `
SELECT 
 id,
 login,
 email,
 avatar,
 active,
 admin,
 token,
 secret,
 hash
FROM users 
LIMIT ? OFFSET ?
`

const SelectUserCountStmt = `
SELECT count(1)
FROM users 
`

const SelectUserPkeyStmt = `
SELECT 
 id,
 login,
 email,
 avatar,
 active,
 admin,
 token,
 secret,
 hash
FROM users 
WHERE id=?
`

const UpdateUserPkeyStmt = `
UPDATE users SET 
 id=?,
 login=?,
 email=?,
 avatar=?,
 active=?,
 admin=?,
 token=?,
 secret=?,
 hash=? 
WHERE id=?
`

const DeleteUserPkeyStmt = `
DELETE FROM users 
WHERE id=?
`

const CreateUserLoginStmt = `
CREATE UNIQUE INDEX IF NOT EXISTS user_login ON users ( login)
`

const SelectUserLoginStmt = `
SELECT 
 id,
 login,
 email,
 avatar,
 active,
 admin,
 token,
 secret,
 hash
FROM users 
WHERE login=?
`

const UpdateUserLoginStmt = `
UPDATE users SET 
 id=?,
 login=?,
 email=?,
 avatar=?,
 active=?,
 admin=?,
 token=?,
 secret=?,
 hash=? 
WHERE login=?
`

const DeleteUserLoginStmt = `
DELETE FROM users 
WHERE login=?
`

const CreateUserEmailStmt = `
CREATE UNIQUE INDEX IF NOT EXISTS user_email ON users ( email)
`

const SelectUserEmailStmt = `
SELECT 
 id,
 login,
 email,
 avatar,
 active,
 admin,
 token,
 secret,
 hash
FROM users 
WHERE email=?
`

const UpdateUserEmailStmt = `
UPDATE users SET 
 id=?,
 login=?,
 email=?,
 avatar=?,
 active=?,
 admin=?,
 token=?,
 secret=?,
 hash=? 
WHERE email=?
`

const DeleteUserEmailStmt = `
DELETE FROM users 
WHERE email=?
`
