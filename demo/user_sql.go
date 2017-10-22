// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

package demo

import (
	"database/sql"
	"fmt"
	"github.com/rickb777/sqlgen/dialect"
	"github.com/rickb777/sqlgen/where"
)

// UserTableName is the default name for this table.
const UserTableName = "users"

// UserTable holds a given table name with the database reference, providing access methods below.
type UserTable struct {
	Name    string
	Db      *sql.DB
	Dialect dialect.Dialect
}

// NewUserTable returns a new table instance.
func NewUserTable(name string, db *sql.DB, dialect dialect.Dialect) UserTable {
	return UserTable{name, db, dialect}
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

// QueryOne is the low-level access function for one User.
func (tbl UserTable) QueryOne(query string, args ...interface{}) (*User, error) {
	row := tbl.Db.QueryRow(query, args...)
	return ScanUser(row)
}

// SelectOneSA allows a single User to be obtained from the database using supplied dialect-specific parameters.
func (tbl UserTable) SelectOneSA(where, limitClause string, args ...interface{}) (*User, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", sUserColumnNames, tbl.Name, where, limitClause)
	return tbl.QueryOne(query, args...)
}

// SelectOne allows a single User to be obtained from the database.
func (tbl UserTable) SelectOne(where where.Expression, dialect dialect.Dialect) (*User, error) {
	wh, args := where.Build(dialect)
	return tbl.SelectOneSA(wh, "LIMIT 1", args)
}

func (tbl UserTable) Query(query string, args ...interface{}) ([]*User, error) {
	rows, err := tbl.Db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanUsers(rows)
}

// SelectSA allows Users to be obtained from the database using supplied dialect-specific parameters.
func (tbl UserTable) SelectSA(where string, args ...interface{}) ([]*User, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s", sUserColumnNames, tbl.Name, where)
	return tbl.Query(query, args...)
}

// Select allows Users to be obtained from the database that match a 'where' clause.
func (tbl UserTable) Select(where where.Expression, dialect dialect.Dialect) ([]*User, error) {
	return tbl.SelectSA(where.Build(dialect))
}

// CountSA counts Users in the database using supplied dialect-specific parameters.
func (tbl UserTable) CountSA(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s %s", tbl.Name, where)
	row := tbl.Db.QueryRow(query, args)
	err = row.Scan(&count)
	return count, err
}

// Count counts the Users in the database that match a 'where' clause.
func (tbl UserTable) Count(where where.Expression, dialect dialect.Dialect) (count int64, err error) {
	return tbl.CountSA(where.Build(dialect))
}

// Insert adds new records for the Users.
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

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// It returns the number of rows affected.
// Not every database or database driver may support this.
func (tbl UserTable) CreateTable() (int64, error) {
//"CREATE TABLE IF NOT EXISTS %s ("
// id       INTEGER PRIMARY KEY AUTOINCREMENT,
// number   INTEGER,
// title    TEXT,
// assignee TEXT,
// state    TEXT,
// labels   BLOB
//")"
	return 0, nil
}

//--------------------------------------------------------------------------------

const NumUserColumns = 9

const sUserColumnNames = `
id, login, email, avatar, active, admin, token, secret, hash
`

const sUserDataColumnNames = `
login, email, avatar, active, admin, token, secret, hash
`

const sUserColumnParams = `
?,?,?,?,?,?,?,?,?
`

const sUserDataColumnParams = `
?,?,?,?,?,?,?,?
`

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

func CreateUserStmt(tableName string, d dialect.Dialect) string {
	return d.ReplacePlaceholders(fmt.Sprintf(sCreateUserStmt, tableName))
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

func InsertUserStmt(tableName string, d dialect.Dialect) string {
	return d.ReplacePlaceholders(fmt.Sprintf(sInsertUserStmt, tableName))
}

const sUpdateUserByPkStmt = `
UPDATE %s SET 
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

func UpdateUserByPkStmt(tableName string, d dialect.Dialect) string {
	return d.ReplacePlaceholders(fmt.Sprintf(sUpdateUserByPkStmt, tableName))
}

const sDeleteUserByPkStmt = `
DELETE FROM %s
 WHERE id=?
`

func DeleteUserByPkStmt(tableName string, d dialect.Dialect) string {
	return d.ReplacePlaceholders(fmt.Sprintf(sDeleteUserByPkStmt, tableName))
}

//--------------------------------------------------------------------------------

const sCreateUserLoginStmt = `
CREATE UNIQUE INDEX IF NOT EXISTS user_login ON %s (login)
`

func CreateUserLoginStmt(tableName string, d dialect.Dialect) string {
	return d.ReplacePlaceholders(fmt.Sprintf(sCreateUserLoginStmt, tableName))
}

const sUpdateUserLoginStmt = `
UPDATE %s SET 
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

func UpdateUserLoginStmt(tableName string, d dialect.Dialect) string {
	return d.ReplacePlaceholders(fmt.Sprintf(sUpdateUserLoginStmt, tableName))
}

const sDeleteUserLoginStmt = `
DELETE FROM %s
 WHERE login=?
`

func DeleteUserLoginStmt(tableName string, d dialect.Dialect) string {
	return d.ReplacePlaceholders(fmt.Sprintf(sDeleteUserLoginStmt, tableName))
}

const sCreateUserEmailStmt = `
CREATE UNIQUE INDEX IF NOT EXISTS user_email ON %s (email)
`

func CreateUserEmailStmt(tableName string, d dialect.Dialect) string {
	return d.ReplacePlaceholders(fmt.Sprintf(sCreateUserEmailStmt, tableName))
}

const sUpdateUserEmailStmt = `
UPDATE %s SET 
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

func UpdateUserEmailStmt(tableName string, d dialect.Dialect) string {
	return d.ReplacePlaceholders(fmt.Sprintf(sUpdateUserEmailStmt, tableName))
}

const sDeleteUserEmailStmt = `
DELETE FROM %s
 WHERE email=?
`

func DeleteUserEmailStmt(tableName string, d dialect.Dialect) string {
	return d.ReplacePlaceholders(fmt.Sprintf(sDeleteUserEmailStmt, tableName))
}
