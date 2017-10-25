// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

package demo

import (
	"database/sql"
	"fmt"
	"github.com/rickb777/sqlgen2/dialect"
	"github.com/rickb777/sqlgen2/schema"
	"github.com/rickb777/sqlgen2/where"
)

// DbUserTableName is the default name for this table.
const DbUserTableName = "users"

// DbUserTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type DbUserTable struct {
	Prefix, Name string
	Db           *sql.DB
	DialectId    schema.DialectId
}

// NewDbUserTable returns a new table instance.
func NewDbUserTable(prefix, name string, db *sql.DB, dialect schema.DialectId) DbUserTable {
	if name == "" {
		name = DbUserTableName
	}
	return DbUserTable{prefix, name, db, dialect}
}

// ScanDbUser reads a database record into a single value.
func ScanDbUser(row *sql.Row) (*User, error) {
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
	v.Uid = v0
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

// ScanDbUsers reads database records into a slice of values.
func ScanDbUsers(rows *sql.Rows) ([]*User, error) {
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
		v.Uid = v0
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

	v0 = v.Uid
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
func (tbl DbUserTable) QueryOne(query string, args ...interface{}) (*User, error) {
	row := tbl.Db.QueryRow(query, args...)
	return ScanDbUser(row)
}

// SelectOneSA allows a single User to be obtained from the database using supplied dialect-specific parameters.
func (tbl DbUserTable) SelectOneSA(where, limitClause string, args ...interface{}) (*User, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", UserColumnNames, tbl.Prefix, tbl.Name, where, limitClause)
	return tbl.QueryOne(query, args...)
}

// SelectOne allows a single User to be obtained from the database.
func (tbl DbUserTable) SelectOne(where where.Expression, dialect dialect.Dialect) (*User, error) {
	wh, args := where.Build(dialect)
	return tbl.SelectOneSA(wh, "LIMIT 1", args)
}

func (tbl DbUserTable) Query(query string, args ...interface{}) ([]*User, error) {
	rows, err := tbl.Db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanDbUsers(rows)
}

// SelectSA allows Users to be obtained from the database using supplied dialect-specific parameters.
func (tbl DbUserTable) SelectSA(where string, args ...interface{}) ([]*User, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s", UserColumnNames, tbl.Prefix, tbl.Name, where)
	return tbl.Query(query, args...)
}

// Select allows Users to be obtained from the database that match a 'where' clause.
func (tbl DbUserTable) Select(where where.Expression, dialect dialect.Dialect) ([]*User, error) {
	return tbl.SelectSA(where.Build(dialect))
}

// CountSA counts Users in the database using supplied dialect-specific parameters.
func (tbl DbUserTable) CountSA(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s%s %s", tbl.Prefix, tbl.Name, where)
	row := tbl.Db.QueryRow(query, args)
	err = row.Scan(&count)
	return count, err
}

// Count counts the Users in the database that match a 'where' clause.
func (tbl DbUserTable) Count(where where.Expression, dialect dialect.Dialect) (count int64, err error) {
	return tbl.CountSA(where.Build(dialect))
}

// Insert adds new records for the Users. The Users have their primary key fields
// set to the new record identifiers.
func (tbl DbUserTable) Insert(vv ...*User) error {
	var stmt string
	switch tbl.DialectId {
	case schema.Sqlite: stmt = sqlInsertUserSqlite
    case schema.Postgres: stmt = sqlInsertUserPostgres
    case schema.Mysql: stmt = sqlInsertUserMysql
    }
	st, err := tbl.Db.Prepare(fmt.Sprintf(stmt, tbl.Prefix, tbl.Name))
	if err != nil {
		return err
	}
	defer st.Close()

	for _, v := range vv {
		res, err := st.Exec(SliceUserWithoutPk(v)...)
		if err != nil {
			return err
		}

		v.Uid, err = res.LastInsertId()
		if err != nil {
			return err
		}
	}

	return nil
}

// Update updates a record. It returns the number of rows affected.
// Not every database or database driver may support this.
func (tbl DbUserTable) Update(v *User) (int64, error) {
	var stmt string
	switch tbl.DialectId {
	case schema.Sqlite: stmt = sqlUpdateUserByPkSqlite
    case schema.Postgres: stmt = sqlUpdateUserByPkPostgres
    case schema.Mysql: stmt = sqlUpdateUserByPkMysql
    }
	query := fmt.Sprintf(stmt, tbl.Prefix, tbl.Name)
	args := SliceUserWithoutPk(v)
	args = append(args, v.Uid)
	return tbl.Exec(query, args...)
}

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// It returns the number of rows affected.
// Not every database or database driver may support this.
func (tbl DbUserTable) Exec(query string, args ...interface{}) (int64, error) {
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
func (tbl DbUserTable) CreateTable(ifNotExist bool) (int64, error) {
	return tbl.Exec(tbl.createTableSql(ifNotExist))
}

func (tbl DbUserTable) createTableSql(ifNotExist bool) string {
	var stmt string
	switch tbl.DialectId {
	case schema.Sqlite: stmt = sqlCreateUserTableSqlite
    case schema.Postgres: stmt = sqlCreateUserTablePostgres
    case schema.Mysql: stmt = sqlCreateUserTableMysql
    }
	extra := tbl.ternary(ifNotExist, "IF NOT EXISTS ", "")
	query := fmt.Sprintf(stmt, extra, tbl.Prefix, tbl.Name)
	return query
}

func (tbl DbUserTable) ternary(flag bool, a, b string) string {
	if flag {
		return a
	}
	return b
}

// CreateIndexes executes queries that create the indexes needed by the User table.
func (tbl DbUserTable) CreateIndexes(ifNotExist bool) (err error) {
	extra := tbl.ternary(ifNotExist, "IF NOT EXISTS ", "")
    
	_, err = tbl.Exec(tbl.createUserLoginIndexSql(extra))
	if err != nil {
		return err
	}
    
	_, err = tbl.Exec(tbl.createUserEmailIndexSql(extra))
	if err != nil {
		return err
	}
    
	return nil
}


func (tbl DbUserTable) createUserLoginIndexSql(ifNotExist string) string {
	var stmt string
	switch tbl.DialectId {
	case schema.Sqlite: stmt = sqlCreateUserLoginIndexSqlite
    case schema.Postgres: stmt = sqlCreateUserLoginIndexPostgres
    case schema.Mysql: stmt = sqlCreateUserLoginIndexMysql
    }
	return fmt.Sprintf(stmt, ifNotExist, tbl.Prefix, tbl.Name)
}

func (tbl DbUserTable) createUserEmailIndexSql(ifNotExist string) string {
	var stmt string
	switch tbl.DialectId {
	case schema.Sqlite: stmt = sqlCreateUserEmailIndexSqlite
    case schema.Postgres: stmt = sqlCreateUserEmailIndexPostgres
    case schema.Mysql: stmt = sqlCreateUserEmailIndexMysql
    }
	return fmt.Sprintf(stmt, ifNotExist, tbl.Prefix, tbl.Name)
}



//--------------------------------------------------------------------------------

const sqlCreateUserTableSqlite = `
CREATE TABLE %s%s%s (
 uid    integer primary key autoincrement,
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

const sqlCreateUserTablePostgres = `
CREATE TABLE %s%s%s (
 uid    serial primary key ,
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

const sqlCreateUserTableMysql = `
CREATE TABLE %s%s%s (
 uid    bigint PRIMARY KEY AUTO_INCREMENT,
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

const sqlInsertUserSqlite = sqlInsertUserMysql

const sqlUpdateUserByPkSqlite = sqlUpdateUserByPkMysql

const sqlDeleteUserByPkSqlite = sqlDeleteUserByPkMysql

//--------------------------------------------------------------------------------

const sqlInsertUserPostgres = `
INSERT INTO %s%s (
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

const sqlUpdateUserByPkPostgres = `
UPDATE %s%s SET 
	login=$2,
	email=$3,
	avatar=$4,
	active=$5,
	admin=$6,
	token=$7,
	secret=$8,
	hash=$9 
 WHERE uid=$10
`

const sqlDeleteUserByPkPostgres = `
DELETE FROM %s%s
 WHERE uid=$1
`

//--------------------------------------------------------------------------------

const sqlInsertUserMysql = `
INSERT INTO %s%s (
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

const sqlUpdateUserByPkMysql = `
UPDATE %s%s SET 
	login=?,
	email=?,
	avatar=?,
	active=?,
	admin=?,
	token=?,
	secret=?,
	hash=? 
 WHERE uid=?
`

const sqlDeleteUserByPkMysql = `
DELETE FROM %s%s
 WHERE uid=?
`

//--------------------------------------------------------------------------------

const sqlCreateUserLoginIndexSqlite = `
CREATE UNIQUE INDEX %suser_login ON %s%s (login)
`

const sqlUpdateUserLoginSqlite = `
UPDATE %s%s SET 
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

const sqlDeleteUserLoginSqlite = `
DELETE FROM %s%s
 WHERE login=?
`

//--------------------------------------------------------------------------------

const sqlCreateUserEmailIndexSqlite = `
CREATE UNIQUE INDEX %suser_email ON %s%s (email)
`

const sqlUpdateUserEmailSqlite = `
UPDATE %s%s SET 
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

const sqlDeleteUserEmailSqlite = `
DELETE FROM %s%s
 WHERE email=?
`

//--------------------------------------------------------------------------------

const sqlCreateUserLoginIndexPostgres = `
CREATE UNIQUE INDEX %suser_login ON %s%s (login)
`

const sqlUpdateUserLoginPostgres = `
UPDATE %s%s SET 
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

const sqlDeleteUserLoginPostgres = `
DELETE FROM %s%s
 WHERE login=$1
`

//--------------------------------------------------------------------------------

const sqlCreateUserEmailIndexPostgres = `
CREATE UNIQUE INDEX %suser_email ON %s%s (email)
`

const sqlUpdateUserEmailPostgres = `
UPDATE %s%s SET 
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

const sqlDeleteUserEmailPostgres = `
DELETE FROM %s%s
 WHERE email=$1
`

//--------------------------------------------------------------------------------

const sqlCreateUserLoginIndexMysql = `
CREATE UNIQUE INDEX %suser_login ON %s%s (login)
`

const sqlUpdateUserLoginMysql = `
UPDATE %s%s SET 
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

const sqlDeleteUserLoginMysql = `
DELETE FROM %s%s
 WHERE login=?
`

//--------------------------------------------------------------------------------

const sqlCreateUserEmailIndexMysql = `
CREATE UNIQUE INDEX %suser_email ON %s%s (email)
`

const sqlUpdateUserEmailMysql = `
UPDATE %s%s SET 
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

const sqlDeleteUserEmailMysql = `
DELETE FROM %s%s
 WHERE email=?
`

//--------------------------------------------------------------------------------

const NumUserColumns = 9

const UserPk = "Uid"

const UserColumnNames = "uid, login, email, avatar, active, admin, token, secret, hash"

const UserDataColumnNames = "login, email, avatar, active, admin, token, secret, hash"

//--------------------------------------------------------------------------------
