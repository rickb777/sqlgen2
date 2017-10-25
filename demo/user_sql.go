// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

package demo

import (
	"database/sql"
	"fmt"
	"github.com/rickb777/sqlgen2/dialect"
	"github.com/rickb777/sqlgen2/schema"
	"github.com/rickb777/sqlgen2/where"
)

// UserTableName is the default name for this table.
const UserTableName = "users"

// UserTable holds a given table name with the database reference, providing access methods below.
type UserTable struct {
	Name      string
	Db        *sql.DB
	DialectId schema.DialectId
}

// NewUserTable returns a new table instance.
func NewUserTable(name string, db *sql.DB, dialect schema.DialectId) UserTable {
	if name == "" {
		name = UserTableName
	}
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
	var stmt string
	switch tbl.DialectId {
	case schema.Sqlite: stmt = sInsertUserStmtSqlite
    case schema.Postgres: stmt = sInsertUserStmtPostgres
    case schema.Mysql: stmt = sInsertUserStmtMysql
    }
	query := fmt.Sprintf(stmt, tbl.Name)
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
	var stmt string
	switch tbl.DialectId {
	case schema.Sqlite: stmt = sUpdateUserByPkStmtSqlite
    case schema.Postgres: stmt = sUpdateUserByPkStmtPostgres
    case schema.Mysql: stmt = sUpdateUserByPkStmtMysql
    }
	query := fmt.Sprintf(stmt, tbl.Name)
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
	var stmt string
	switch tbl.DialectId {
	case schema.Sqlite: stmt = sCreateUserStmtSqlite
    case schema.Postgres: stmt = sCreateUserStmtPostgres
    case schema.Mysql: stmt = sCreateUserStmtMysql
    }
	query := fmt.Sprintf(stmt, tbl.Name)
	return tbl.Exec(query)
}

//--------------------------------------------------------------------------------

const NumUserColumns = 9

const sUserColumnNames = `
id, login, email, avatar, active, admin, token, secret, hash
`

const sUserDataColumnNames = `
login, email, avatar, active, admin, token, secret, hash
`

const sCreateUserStmtSqlite = `
CREATE TABLE IF NOT EXISTS %s (
 id     integer primary key autoincrement,
 login  text,
 email  text,
 avatar text,
 active boolean,
 admin  boolean,
 token  text,
 secret text,
 hash   text
)
`

const sCreateUserStmtPostgres = `
CREATE TABLE IF NOT EXISTS %s (
 id     serial PRIMARY KEY ,
 login  varchar(512),
 email  varchar(512),
 avatar varchar(512),
 active boolean,
 admin  boolean,
 token  varchar(512),
 secret varchar(512),
 hash   varchar(512)
)
`

const sCreateUserStmtMysql = `
CREATE TABLE IF NOT EXISTS %s (
 id     bigint PRIMARY KEY AUTO_INCREMENT,
 login  VARCHAR(512),
 email  VARCHAR(512),
 avatar VARCHAR(512),
 active TINYINT(1),
 admin  TINYINT(1),
 token  VARCHAR(512),
 secret VARCHAR(512),
 hash   VARCHAR(512)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
`

//--------------------------------------------------------------------------------

const sInsertUserStmtSqlite = sInsertUserStmtMysql

const sUpdateUserByPkStmtSqlite = sUpdateUserByPkStmtMysql

const sDeleteUserByPkStmtSqlite = sDeleteUserByPkStmtMysql

//--------------------------------------------------------------------------------

const sInsertUserStmtPostgres = `
INSERT INTO %s (
 login,
 email,
 avatar,
 active,
 admin,
 token,
 secret,
 hash
) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
`

const sUpdateUserByPkStmtPostgres = `
UPDATE %s SET 
 login=$2,
 email=$3,
 avatar=$4,
 active=$5,
 admin=$6,
 token=$7,
 secret=$8,
 hash=$9 
 WHERE id=$10
`

const sDeleteUserByPkStmtPostgres = `
DELETE FROM %s
 WHERE id=$1
`

//--------------------------------------------------------------------------------

const sInsertUserStmtMysql = `
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

const sUpdateUserByPkStmtMysql = `
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

const sDeleteUserByPkStmtMysql = `
DELETE FROM %s
 WHERE id=?
`

//--------------------------------------------------------------------------------

const sCreateUserLoginStmtSqlite = `
CREATE UNIQUE INDEX IF NOT EXISTS user_login ON %s (login)
`

const sUpdateUserLoginStmtSqlite = `
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

const sDeleteUserLoginStmtSqlite = `
DELETE FROM %s
 WHERE login=?
`

//--------------------------------------------------------------------------------

const sCreateUserEmailStmtSqlite = `
CREATE UNIQUE INDEX IF NOT EXISTS user_email ON %s (email)
`

const sUpdateUserEmailStmtSqlite = `
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

const sDeleteUserEmailStmtSqlite = `
DELETE FROM %s
 WHERE email=?
`

//--------------------------------------------------------------------------------

const sCreateUserLoginStmtPostgres = `
CREATE UNIQUE INDEX IF NOT EXISTS user_login ON %s (login)
`

const sUpdateUserLoginStmtPostgres = `
UPDATE %s SET 
 login=$2,
 email=$3,
 avatar=$4,
 active=$5,
 admin=$6,
 token=$7,
 secret=$8,
 hash=$9 
 WHERE login=$10
`

const sDeleteUserLoginStmtPostgres = `
DELETE FROM %s
 WHERE login=$1
`

//--------------------------------------------------------------------------------

const sCreateUserEmailStmtPostgres = `
CREATE UNIQUE INDEX IF NOT EXISTS user_email ON %s (email)
`

const sUpdateUserEmailStmtPostgres = `
UPDATE %s SET 
 login=$2,
 email=$3,
 avatar=$4,
 active=$5,
 admin=$6,
 token=$7,
 secret=$8,
 hash=$9 
 WHERE email=$10
`

const sDeleteUserEmailStmtPostgres = `
DELETE FROM %s
 WHERE email=$1
`

//--------------------------------------------------------------------------------

const sCreateUserLoginStmtMysql = `
CREATE UNIQUE INDEX IF NOT EXISTS user_login ON %s (login)
`

const sUpdateUserLoginStmtMysql = `
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

const sDeleteUserLoginStmtMysql = `
DELETE FROM %s
 WHERE login=?
`

//--------------------------------------------------------------------------------

const sCreateUserEmailStmtMysql = `
CREATE UNIQUE INDEX IF NOT EXISTS user_email ON %s (email)
`

const sUpdateUserEmailStmtMysql = `
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

const sDeleteUserEmailStmtMysql = `
DELETE FROM %s
 WHERE email=?
`

//--------------------------------------------------------------------------------
