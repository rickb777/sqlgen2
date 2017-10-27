// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

package demo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/rickb777/sqlgen2/db"
	"github.com/rickb777/sqlgen2/where"
	"strings"
)

// DbUserTableName is the default name for this table.
const DbUserTableName = "users"

// DbUserTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type DbUserTable struct {
	Prefix, Name string
	Db           db.Execer
	Ctx          context.Context
	Dialect      db.Dialect
}

// NewDbUserTable returns a new table instance.
func NewDbUserTable(prefix, name string, d *sql.DB, dialect db.Dialect) DbUserTable {
	if name == "" {
		name = DbUserTableName
	}
	return DbUserTable{prefix, name, d, context.Background(), dialect}
}

// WithContext sets the context for subsequent queries.
func (tbl DbUserTable) WithContext(ctx context.Context) DbUserTable {
	tbl.Ctx = ctx
	return tbl
}

// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl DbUserTable) DB() *sql.DB {
	return tbl.Db.(*sql.DB)
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl DbUserTable) Tx() *sql.Tx {
	return tbl.Db.(*sql.Tx)
}

// IsTx tests whether this is within a transaction.
func (tbl DbUserTable) IsTx() bool {
	_, ok := tbl.Db.(*sql.Tx)
	return ok
}

// Begin starts a transaction. The default isolation level is dependent on the driver.
func (tbl DbUserTable) BeginTx(opts *sql.TxOptions) (DbUserTable, error) {
	d := tbl.Db.(*sql.DB)
	var err error
	tbl.Db, err = d.BeginTx(tbl.Ctx, opts)
	return tbl, err
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

func SliceDbUser(v *User) []interface{} {
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

func SliceDbUserWithoutPk(v *User) []interface{} {
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

//--------------------------------------------------------------------------------

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// It returns the number of rows affected.
func (tbl DbUserTable) Exec(query string, args ...interface{}) (int64, error) {
	res, err := tbl.Db.ExecContext(tbl.Ctx, query, args...)
	if err != nil {
		return 0, nil
	}
	return res.RowsAffected()
}

//--------------------------------------------------------------------------------

// QueryOne is the low-level access function for one User.
func (tbl DbUserTable) QueryOne(query string, args ...interface{}) (*User, error) {
	row := tbl.Db.QueryRowContext(tbl.Ctx, query, args...)
	return ScanDbUser(row)
}

// Query is the low-level access function for Users.
func (tbl DbUserTable) Query(query string, args ...interface{}) ([]*User, error) {
	rows, err := tbl.Db.QueryContext(tbl.Ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanDbUsers(rows)
}

//--------------------------------------------------------------------------------

// SelectOneSA allows a single User to be obtained from the database using supplied dialect-specific parameters.
func (tbl DbUserTable) SelectOneSA(where, limitClause string, args ...interface{}) (*User, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", DbUserColumnNames, tbl.Prefix, tbl.Name, where, limitClause)
	return tbl.QueryOne(query, args...)
}

// SelectOne allows a single User to be obtained from the database.
func (tbl DbUserTable) SelectOne(where where.Expression, dialect db.Dialect) (*User, error) {
	wh, args := where.Build(dialect)
	return tbl.SelectOneSA(wh, "LIMIT 1", args)
}

// SelectSA allows Users to be obtained from the database using supplied dialect-specific parameters.
func (tbl DbUserTable) SelectSA(where string, args ...interface{}) ([]*User, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s", DbUserColumnNames, tbl.Prefix, tbl.Name, where)
	return tbl.Query(query, args...)
}

// Select allows Users to be obtained from the database that match a 'where' clause.
func (tbl DbUserTable) Select(where where.Expression, dialect db.Dialect) ([]*User, error) {
	return tbl.SelectSA(where.Build(dialect))
}

// CountSA counts Users in the database using supplied dialect-specific parameters.
func (tbl DbUserTable) CountSA(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s%s %s", tbl.Prefix, tbl.Name, where)
	row := tbl.Db.QueryRowContext(tbl.Ctx, query, args)
	err = row.Scan(&count)
	return count, err
}

// Count counts the Users in the database that match a 'where' clause.
func (tbl DbUserTable) Count(where where.Expression, dialect db.Dialect) (count int64, err error) {
	return tbl.CountSA(where.Build(dialect))
}

const DbUserColumnNames = "uid, login, email, avatar, active, admin, token, secret, hash"

//--------------------------------------------------------------------------------

// Insert adds new records for the Users. The Users have their primary key fields
// set to the new record identifiers.
func (tbl DbUserTable) Insert(vv ...*User) error {
	var stmt, params string
	switch tbl.Dialect {
	case db.Postgres:
		stmt = sqlInsertDbUserPostgres
		params = sDbUserDataColumnParamsPostgres
	default:
		stmt = sqlInsertDbUserSimple
		params = sDbUserDataColumnParamsSimple
	}
	st, err := tbl.Db.PrepareContext(tbl.Ctx, fmt.Sprintf(stmt, tbl.Prefix, tbl.Name, params))
	if err != nil {
		return err
	}
	defer st.Close()

	for _, v := range vv {
		res, err := st.Exec(SliceDbUserWithoutPk(v)...)
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

const sqlInsertDbUserSimple = `
INSERT INTO %s%s (
	login,
	email,
	avatar,
	active,
	admin,
	token,
	secret,
	hash
) VALUES (%s)
`

const sqlInsertDbUserPostgres = `
INSERT INTO %s%s (
	login,
	email,
	avatar,
	active,
	admin,
	token,
	secret,
	hash
) VALUES (%s)
`

const sDbUserDataColumnParamsSimple = "?,?,?,?,?,?,?,?"

const sDbUserDataColumnParamsPostgres = "$1,$2,$3,$4,$5,$6,$7,$8"

//--------------------------------------------------------------------------------

// Update updates a record. It returns the number of rows affected.
// Not every database or database driver may support this.
func (tbl DbUserTable) Update(v *User) (int64, error) {
	var stmt string
	switch tbl.Dialect {
	case db.Postgres:
		stmt = sqlUpdateDbUserByPkPostgres
	default:
		stmt = sqlUpdateDbUserByPkSimple
	}
	query := fmt.Sprintf(stmt, tbl.Prefix, tbl.Name)
	args := SliceDbUserWithoutPk(v)
	args = append(args, v.Uid)
	return tbl.Exec(query, args...)
}

// UpdateFields updates one or more columns, given a 'where' clause.
func (tbl DbUserTable) UpdateFields(where where.Expression, fields ...sql.NamedArg) (int64, error) {
	return tbl.Exec(tbl.updateFields(where, fields...))
}

func (tbl DbUserTable) updateFields(where where.Expression, fields ...sql.NamedArg) (string, []interface{}) {
	list := db.NamedArgList(fields)
	assignments := strings.Join(list.Assignments(tbl.Dialect, 1), ", ")
	whereClause, wargs := where.Build(tbl.Dialect)
	query := fmt.Sprintf("UPDATE %s%s SET %s %s", tbl.Prefix, tbl.Name, assignments, whereClause)
	args := append(list.Values(), wargs...)
	return query, args
}

const sqlUpdateDbUserByPkSimple = `
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

const sqlUpdateDbUserByPkPostgres = `
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

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// It returns the number of rows affected.
// Not every database or database driver may support this.
func (tbl DbUserTable) CreateTable(ifNotExist bool) (int64, error) {
	return tbl.Exec(tbl.createTableSql(ifNotExist))
}

func (tbl DbUserTable) createTableSql(ifNotExist bool) string {
	var stmt string
	switch tbl.Dialect {
	case db.Sqlite: stmt = sqlCreateDbUserTableSqlite
    case db.Postgres: stmt = sqlCreateDbUserTablePostgres
    case db.Mysql: stmt = sqlCreateDbUserTableMysql
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
    
	_, err = tbl.Exec(tbl.createDbUserLoginIndexSql(extra))
	if err != nil {
		return err
	}
    
	_, err = tbl.Exec(tbl.createDbUserEmailIndexSql(extra))
	if err != nil {
		return err
	}
    
	return nil
}


func (tbl DbUserTable) createDbUserLoginIndexSql(ifNotExist string) string {
	var stmt string
	switch tbl.Dialect {
	case db.Sqlite: stmt = sqlCreateDbUserLoginIndexSqlite
    case db.Postgres: stmt = sqlCreateDbUserLoginIndexPostgres
    case db.Mysql: stmt = sqlCreateDbUserLoginIndexMysql
    }
	return fmt.Sprintf(stmt, ifNotExist, tbl.Prefix, tbl.Name)
}

func (tbl DbUserTable) createDbUserEmailIndexSql(ifNotExist string) string {
	var stmt string
	switch tbl.Dialect {
	case db.Sqlite: stmt = sqlCreateDbUserEmailIndexSqlite
    case db.Postgres: stmt = sqlCreateDbUserEmailIndexPostgres
    case db.Mysql: stmt = sqlCreateDbUserEmailIndexMysql
    }
	return fmt.Sprintf(stmt, ifNotExist, tbl.Prefix, tbl.Name)
}



//--------------------------------------------------------------------------------

const sqlCreateDbUserTableSqlite = `
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

const sqlCreateDbUserTablePostgres = `
CREATE TABLE %s%s%s (
 uid    bigserial primary key ,
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

const sqlCreateDbUserTableMysql = `
CREATE TABLE %s%s%s (
 uid    bigint primary key auto_increment,
 login  varchar(512),
 email  varchar(512),
 avatar varchar(512),
 active tinyint(1),
 admin  tinyint(1),
 token  varchar(512),
 secret varchar(512),
 hash   varchar(512)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
`

//--------------------------------------------------------------------------------

const sqlDeleteDbUserByPkPostgres = `
DELETE FROM %s%s
 WHERE uid=$1
`

const sqlDeleteDbUserByPkSimple = `
DELETE FROM %s%s
 WHERE uid=?
`

//--------------------------------------------------------------------------------

const sqlCreateDbUserLoginIndexSqlite = `
CREATE UNIQUE INDEX %suser_login ON %s%s (login)
`

const sqlUpdateDbUserLoginSqlite = `
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

const sqlDeleteDbUserLoginSqlite = `
DELETE FROM %s%s
 WHERE login=?
`

//--------------------------------------------------------------------------------

const sqlCreateDbUserEmailIndexSqlite = `
CREATE UNIQUE INDEX %suser_email ON %s%s (email)
`

const sqlUpdateDbUserEmailSqlite = `
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

const sqlDeleteDbUserEmailSqlite = `
DELETE FROM %s%s
 WHERE email=?
`

//--------------------------------------------------------------------------------

const sqlCreateDbUserLoginIndexPostgres = `
CREATE UNIQUE INDEX %suser_login ON %s%s (login)
`

const sqlUpdateDbUserLoginPostgres = `
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

const sqlDeleteDbUserLoginPostgres = `
DELETE FROM %s%s
 WHERE login=$1
`

//--------------------------------------------------------------------------------

const sqlCreateDbUserEmailIndexPostgres = `
CREATE UNIQUE INDEX %suser_email ON %s%s (email)
`

const sqlUpdateDbUserEmailPostgres = `
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

const sqlDeleteDbUserEmailPostgres = `
DELETE FROM %s%s
 WHERE email=$1
`

//--------------------------------------------------------------------------------

const sqlCreateDbUserLoginIndexMysql = `
CREATE UNIQUE INDEX %suser_login ON %s%s (login)
`

const sqlUpdateDbUserLoginMysql = `
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

const sqlDeleteDbUserLoginMysql = `
DELETE FROM %s%s
 WHERE login=?
`

//--------------------------------------------------------------------------------

const sqlCreateDbUserEmailIndexMysql = `
CREATE UNIQUE INDEX %suser_email ON %s%s (email)
`

const sqlUpdateDbUserEmailMysql = `
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

const sqlDeleteDbUserEmailMysql = `
DELETE FROM %s%s
 WHERE email=?
`

//--------------------------------------------------------------------------------

const NumDbUserColumns = 9

const NumDbUserDataColumns = 8

const DbUserPk = "Uid"

const DbUserDataColumnNames = "login, email, avatar, active, admin, token, secret, hash"

//--------------------------------------------------------------------------------
