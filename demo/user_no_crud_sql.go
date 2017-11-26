// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

package demo

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/rickb777/sqlgen2"
)

// V3UserTableName is the default name for this table.
const V3UserTableName = "users"

// V3UserTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type V3UserTable struct {
	Prefix, Name string
	Db           sqlgen2.Execer
	Ctx          context.Context
	Dialect      sqlgen2.Dialect
}

// NewV3UserTable returns a new table instance.
func NewV3UserTable(prefix, name string, d *sql.DB, dialect sqlgen2.Dialect) V3UserTable {
	if name == "" {
		name = V3UserTableName
	}
	return V3UserTable{prefix, name, d, context.Background(), dialect}
}

// WithContext sets the context for subsequent queries.
func (tbl V3UserTable) WithContext(ctx context.Context) V3UserTable {
	tbl.Ctx = ctx
	return tbl
}

// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl V3UserTable) DB() *sql.DB {
	return tbl.Db.(*sql.DB)
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl V3UserTable) Tx() *sql.Tx {
	return tbl.Db.(*sql.Tx)
}

// IsTx tests whether this is within a transaction.
func (tbl V3UserTable) IsTx() bool {
	_, ok := tbl.Db.(*sql.Tx)
	return ok
}

// Begin starts a transaction. The default isolation level is dependent on the driver.
func (tbl V3UserTable) BeginTx(opts *sql.TxOptions) (V3UserTable, error) {
	d := tbl.Db.(*sql.DB)
	var err error
	tbl.Db, err = d.BeginTx(tbl.Ctx, opts)
	return tbl, err
}


// ScanV3User reads a database record into a single value.
func ScanV3User(row *sql.Row) (*User, error) {
	var v0 int64
	var v1 string
	var v2 string
	var v3 string
	var v4 bool
	var v5 bool
	var v6 big.Int
	var v7 string
	var v8 string
	var v9 string

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
		&v9,

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
	json.Unmarshal(v6, &v.Fave)
	v.token = v7
	v.secret = v8
	v.hash = v9

	return v, nil
}

// ScanV3Users reads database records into a slice of values.
func ScanV3Users(rows *sql.Rows) ([]*User, error) {
	var err error
	var vv []*User

	var v0 int64
	var v1 string
	var v2 string
	var v3 string
	var v4 bool
	var v5 bool
	var v6 big.Int
	var v7 string
	var v8 string
	var v9 string

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
			&v9,

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
		json.Unmarshal(v6, &v.Fave)
		v.token = v7
		v.secret = v8
		v.hash = v9

		vv = append(vv, v)
	}
	return vv, rows.Err()
}

func SliceV3User(v *User) []interface{} {
	var v0 int64
	var v1 string
	var v2 string
	var v3 string
	var v4 bool
	var v5 bool
	var v6 big.Int
	var v7 string
	var v8 string
	var v9 string

	v0 = v.Uid
	v1 = v.Login
	v2 = v.Email
	v3 = v.Avatar
	v4 = v.Active
	v5 = v.Admin
	v6, _ = json.Marshal(&v.Fave)
	v7 = v.token
	v8 = v.secret
	v9 = v.hash

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
		v9,

	}
}

func SliceV3UserWithoutPk(v *User) []interface{} {
	var v1 string
	var v2 string
	var v3 string
	var v4 bool
	var v5 bool
	var v6 big.Int
	var v7 string
	var v8 string
	var v9 string

	v1 = v.Login
	v2 = v.Email
	v3 = v.Avatar
	v4 = v.Active
	v5 = v.Admin
	v6, _ = json.Marshal(&v.Fave)
	v7 = v.token
	v8 = v.secret
	v9 = v.hash

	return []interface{}{
		v1,
		v2,
		v3,
		v4,
		v5,
		v6,
		v7,
		v8,
		v9,

	}
}

//--------------------------------------------------------------------------------

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// It returns the number of rows affected.
func (tbl V3UserTable) Exec(query string, args ...interface{}) (int64, error) {
	res, err := tbl.Db.ExecContext(tbl.Ctx, query, args...)
	if err != nil {
		return 0, nil
	}
	return res.RowsAffected()
}

//--------------------------------------------------------------------------------

// QueryOne is the low-level access function for one User.
func (tbl V3UserTable) QueryOne(query string, args ...interface{}) (*User, error) {
	row := tbl.Db.QueryRowContext(tbl.Ctx, query, args...)
	return ScanV3User(row)
}

// Query is the low-level access function for Users.
func (tbl V3UserTable) Query(query string, args ...interface{}) ([]*User, error) {
	rows, err := tbl.Db.QueryContext(tbl.Ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanV3Users(rows)
}

//--------------------------------------------------------------------------------

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// It returns the number of rows affected.
// Not every database or database driver may support this.
func (tbl V3UserTable) CreateTable(ifNotExist bool) (int64, error) {
	return tbl.Exec(tbl.createTableSql(ifNotExist))
}

func (tbl V3UserTable) createTableSql(ifNotExist bool) string {
	var stmt string
	switch tbl.Dialect {
	case sqlgen2.Sqlite: stmt = sqlCreateV3UserTableSqlite
    case sqlgen2.Postgres: stmt = sqlCreateV3UserTablePostgres
    case sqlgen2.Mysql: stmt = sqlCreateV3UserTableMysql
    }
	extra := tbl.ternary(ifNotExist, "IF NOT EXISTS ", "")
	query := fmt.Sprintf(stmt, extra, tbl.Prefix, tbl.Name)
	return query
}

func (tbl V3UserTable) ternary(flag bool, a, b string) string {
	if flag {
		return a
	}
	return b
}

//--------------------------------------------------------------------------------

// CreateIndexes executes queries that create the indexes needed by the User table.
func (tbl V3UserTable) CreateIndexes(ifNotExist bool) (err error) {
	extra := tbl.ternary(ifNotExist, "IF NOT EXISTS ", "")
    
	_, err = tbl.Exec(tbl.createV3UserLoginIndexSql(extra))
	if err != nil {
		return err
	}
    
	_, err = tbl.Exec(tbl.createV3UserEmailIndexSql(extra))
	if err != nil {
		return err
	}
    
	return nil
}


func (tbl V3UserTable) createV3UserLoginIndexSql(ifNotExist string) string {
	return fmt.Sprintf(sqlCreateV3UserLoginIndex, ifNotExist, tbl.Prefix, tbl.Name)
}

func (tbl V3UserTable) createV3UserEmailIndexSql(ifNotExist string) string {
	return fmt.Sprintf(sqlCreateV3UserEmailIndex, ifNotExist, tbl.Prefix, tbl.Name)
}


//--------------------------------------------------------------------------------

const sqlCreateV3UserTableSqlite = `
CREATE TABLE %s%s%s (
 uid    integer primary key autoincrement,
 login  text,
 email  text,
 avatar text,
 active boolean,
 admin  boolean,
 fave   blob,
 token  text,
 secret text,
 hash   text
)
`

const sqlCreateV3UserTablePostgres = `
CREATE TABLE %s%s%s (
 uid    bigserial primary key ,
 login  varchar(512),
 email  varchar(512),
 avatar varchar(512),
 active boolean,
 admin  boolean,
 fave   byteaa,
 token  varchar(512),
 secret varchar(512),
 hash   varchar(512)
)
`

const sqlCreateV3UserTableMysql = `
CREATE TABLE %s%s%s (
 uid    bigint primary key auto_increment,
 login  varchar(512),
 email  varchar(512),
 avatar varchar(512),
 active tinyint(1),
 admin  tinyint(1),
 fave   mediumblob,
 token  varchar(512),
 secret varchar(512),
 hash   varchar(512)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
`

//--------------------------------------------------------------------------------

const sqlCreateV3UserLoginIndex = `
CREATE UNIQUE INDEX %suser_login ON %s%s (login)
`

const sqlCreateV3UserEmailIndex = `
CREATE UNIQUE INDEX %suser_email ON %s%s (email)
`

//--------------------------------------------------------------------------------

const NumV3UserColumns = 10

const NumV3UserDataColumns = 9

const V3UserPk = "Uid"

const V3UserDataColumnNames = "login, email, avatar, active, admin, fave, token, secret, hash"

//--------------------------------------------------------------------------------
