// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

package demo

import (
	"database/sql"
	"fmt"
)

// UserTableName is the default name for this table.
const UserTableName = "users"

// UserTable holds a given table name with the database reference, providing access methods below.
type UserTable struct {
	Name string
	Db   *sql.DB
}

// ScanUser reads a database record into a single value.
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
	v.Id = v0
	v.Login = v1
	v.Email = v2
	v.Avatar = v3
	v.Active = v4
	v.Admin = v5
	v.token = v6
	v.secret = v7
	v.hash = v8


	return v, nil
}

// ScanUsers reads database records into a slice of values.
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
		v.Id = v0
		v.Login = v1
		v.Email = v2
		v.Avatar = v3
		v.Active = v4
		v.Admin = v5
		v.token = v6
		v.secret = v7
		v.hash = v8

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

	v0 = v.Id
	v1 = v.Login
	v2 = v.Email
	v3 = v.Avatar
	v4 = v.Active
	v5 = v.Admin
	v6 = v.token
	v7 = v.secret
	v8 = v.hash

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

func SliceUserWithoutPk(v *User) []interface{} {
	var v1 string
	var v2 string
	var v3 string
	var v4 bool
	var v5 bool
	var v6 string
	var v7 string
	var v8 string

	v1 = v.Login
	v2 = v.Email
	v3 = v.Avatar
	v4 = v.Active
	v5 = v.Admin
	v6 = v.token
	v7 = v.secret
	v8 = v.hash

	return []interface{}{
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

func (tbl UserTable) SelectOne(query string, args ...interface{}) (*User, error) {
	row := tbl.Db.QueryRow(query, args...)
	return ScanUser(row)
}

func (tbl UserTable) Select(query string, args ...interface{}) ([]*User, error) {
	rows, err := tbl.Db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanUsers(rows)
}

func (tbl UserTable) Insert(v *User) error {
	query := fmt.Sprintf(sInsertUserStmt, tbl.Name)
	res, err := tbl.Db.Exec(query, SliceUserWithoutPk(v)...)
	if err != nil {
		return err
	}

	v.Id, err = res.LastInsertId()
	return err
}

// Update updates a record. It returns the number of rows affected.
// Not every database or database driver may support this.
func (tbl UserTable) Update(v *User) (int64, error) {
	query := fmt.Sprintf(sUpdateUserByPkStmt, tbl.Name)
	args := SliceUserWithoutPk(v)
	args = append(args, v.Id)
	return tbl.Exec(query, args...)
}

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// It returns the number of rows affected.
// Not every database or database driver may support this.
func (tbl UserTable) Exec(query string, args ...interface{}) (int64, error) {
	res, err := tbl.Db.Exec(query, args...)
	if err != nil {
		return 0, nil
	}
	return res.RowsAffected()
}

//--------------------------------------------------------------------------------

const sCreateUserStmt = `
CREATE TABLE IF NOT EXISTS %s (
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

func CreateUserStmt(tableName string) string {
	return fmt.Sprintf(sCreateUserStmt, tableName)
}

const sInsertUserStmt = `
INSERT INTO %s (
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

func InsertUserStmt(tableName string) string {
	return fmt.Sprintf(sInsertUserStmt, tableName)
}

const sSelectUserStmt = `
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
FROM %s
`

func SelectUserStmt(tableName string) string {
	return fmt.Sprintf(sSelectUserStmt, tableName)
}

const sSelectUserRangeStmt = `
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
FROM %s
LIMIT ? OFFSET ?
`

func SelectUserRangeStmt(tableName string) string {
	return fmt.Sprintf(sSelectUserRangeStmt, tableName)
}

const sSelectUserCountStmt = `
SELECT count(1)
FROM %s 
`

func SelectUserCountStmt(tableName string) string {
	return fmt.Sprintf(sSelectUserCountStmt, tableName)
}

const sSelectUserByPkStmt = `
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
FROM %s
 WHERE id=?
`

func SelectUserByPkStmt(tableName string) string {
	return fmt.Sprintf(sSelectUserByPkStmt, tableName)
}

const sUpdateUserByPkStmt = `
UPDATE %s SET 
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

func UpdateUserByPkStmt(tableName string) string {
	return fmt.Sprintf(sUpdateUserByPkStmt, tableName)
}

const sDeleteUserByPkeyStmt = `
DELETE FROM %s
 WHERE id=?
`

func DeleteUserByPkeyStmt(tableName string) string {
	return fmt.Sprintf(sDeleteUserByPkeyStmt, tableName)
}

//--------------------------------------------------------------------------------

const sCreateUserLoginStmt = `
CREATE UNIQUE INDEX IF NOT EXISTS user_login ON %s (login)
`

func CreateUserLoginStmt(tableName string) string {
	return fmt.Sprintf(sCreateUserLoginStmt, tableName)
}

const sSelectUserLoginStmt = `
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
FROM %s
 WHERE login=?
`

func SelectUserLoginStmt(tableName string) string {
	return fmt.Sprintf(sSelectUserLoginStmt, tableName)
}

const sUpdateUserLoginStmt = `
UPDATE %s SET 
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

func UpdateUserLoginStmt(tableName string) string {
	return fmt.Sprintf(sUpdateUserLoginStmt, tableName)
}

const sDeleteUserLoginStmt = `
DELETE FROM %s
 WHERE login=?
`

func DeleteUserLoginStmt(tableName string) string {
	return fmt.Sprintf(sDeleteUserLoginStmt, tableName)
}

const sCreateUserEmailStmt = `
CREATE UNIQUE INDEX IF NOT EXISTS user_email ON %s (email)
`

func CreateUserEmailStmt(tableName string) string {
	return fmt.Sprintf(sCreateUserEmailStmt, tableName)
}

const sSelectUserEmailStmt = `
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
FROM %s
 WHERE email=?
`

func SelectUserEmailStmt(tableName string) string {
	return fmt.Sprintf(sSelectUserEmailStmt, tableName)
}

const sUpdateUserEmailStmt = `
UPDATE %s SET 
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

func UpdateUserEmailStmt(tableName string) string {
	return fmt.Sprintf(sUpdateUserEmailStmt, tableName)
}

const sDeleteUserEmailStmt = `
DELETE FROM %s
 WHERE email=?
`

func DeleteUserEmailStmt(tableName string) string {
	return fmt.Sprintf(sDeleteUserEmailStmt, tableName)
}
