// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

package demo

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/rickb777/sqlgen2"
	"github.com/rickb777/sqlgen2/schema"
	"log"
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
	Dialect      schema.Dialect
	Logger       *log.Logger
}

// Type conformance check
var _ sqlgen2.Table = &V3UserTable{}

// NewV3UserTable returns a new table instance.
// If a blank table name is supplied, the default name "users" will be used instead.
// The table name prefix is initially blank and the request context is the background.
func NewV3UserTable(name string, d sqlgen2.Execer, dialect schema.Dialect) V3UserTable {
	if name == "" {
		name = V3UserTableName
	}
	return V3UserTable{"", name, d, context.Background(), dialect, nil}
}

// WithPrefix sets the prefix for subsequent queries.
func (tbl V3UserTable) WithPrefix(pfx string) V3UserTable {
	tbl.Prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries.
func (tbl V3UserTable) WithContext(ctx context.Context) V3UserTable {
	tbl.Ctx = ctx
	return tbl
}

// WithLogger sets the logger for subsequent queries.
func (tbl V3UserTable) WithLogger(logger *log.Logger) V3UserTable {
	tbl.Logger = logger
	return tbl
}

// SetLogger sets the logger for subsequent queries, returning the interface.
func (tbl V3UserTable) SetLogger(logger *log.Logger) sqlgen2.Table {
	tbl.Logger = logger
	return tbl
}

// FullName gets the concatenated prefix and table name.
func (tbl V3UserTable) FullName() string {
	return tbl.Prefix + tbl.Name
}

func (tbl V3UserTable) prefixWithoutDot() string {
	last := len(tbl.Prefix)-1
	if last > 0 && tbl.Prefix[last] == '.' {
		return tbl.Prefix[0:last]
	}
	return tbl.Prefix
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

func (tbl V3UserTable) logQuery(query string, args ...interface{}) {
	sqlgen2.LogQuery(tbl.Logger, query, args...)
}


//--------------------------------------------------------------------------------

const NumV3UserColumns = 11

const NumV3UserDataColumns = 10

const V3UserPk = "Uid"

const V3UserDataColumnNames = "login, emailaddress, avatar, active, admin, fave, lastupdated, token, secret, hash"

//--------------------------------------------------------------------------------

// CreateTable creates the table.
func (tbl V3UserTable) CreateTable(ifNotExist bool) (int64, error) {
	return tbl.Exec(tbl.createTableSql(ifNotExist))
}

func (tbl V3UserTable) createTableSql(ifNotExist bool) string {
	var stmt string
	switch tbl.Dialect {
	case schema.Sqlite: stmt = sqlCreateV3UserTableSqlite
    case schema.Postgres: stmt = sqlCreateV3UserTablePostgres
    case schema.Mysql: stmt = sqlCreateV3UserTableMysql
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

const sqlCreateV3UserTableSqlite = `
CREATE TABLE %s%s%s (
 uid          integer primary key autoincrement,
 login        text,
 emailaddress text,
 avatar       text,
 active       boolean,
 admin        boolean,
 fave         text,
 lastupdated  bigint,
 token        text,
 secret       text,
 hash         text
)
`

const sqlCreateV3UserTablePostgres = `
CREATE TABLE %s%s%s (
 uid          bigserial primary key,
 login        varchar(255),
 emailaddress varchar(255),
 avatar       varchar(255),
 active       boolean,
 admin        boolean,
 fave         json,
 lastupdated  bigint,
 token        varchar(255),
 secret       varchar(255),
 hash         varchar(255)
)
`

const sqlCreateV3UserTableMysql = `
CREATE TABLE %s%s%s (
 uid          bigint primary key auto_increment,
 login        varchar(255),
 emailaddress varchar(255),
 avatar       varchar(255),
 active       tinyint(1),
 admin        tinyint(1),
 fave         json,
 lastupdated  bigint,
 token        varchar(255),
 secret       varchar(255),
 hash         varchar(255)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
`

//--------------------------------------------------------------------------------

// CreateTableWithIndexes invokes CreateTable then CreateIndexes.
func (tbl V3UserTable) CreateTableWithIndexes(ifNotExist bool) (err error) {
	_, err = tbl.CreateTable(ifNotExist)
	if err != nil {
		return err
	}

	return tbl.CreateIndexes(ifNotExist)
}

// CreateIndexes executes queries that create the indexes needed by the User table.
func (tbl V3UserTable) CreateIndexes(ifNotExist bool) (err error) {
	ine := tbl.ternary(ifNotExist && tbl.Dialect != schema.Mysql, "IF NOT EXISTS ", "")

	// Mysql does not support 'if not exists' on indexes
	// Workaround: use DropIndex first and ignore an error returned if the index didn't exist.

	if ifNotExist && tbl.Dialect == schema.Mysql {
		tbl.DropIndexes(false)
		ine = ""
	}

	_, err = tbl.Exec(tbl.createV3UserLoginIndexSql(ine))
	if err != nil {
		return err
	}

	_, err = tbl.Exec(tbl.createV3UserEmailIndexSql(ine))
	if err != nil {
		return err
	}

	return nil
}

func (tbl V3UserTable) createV3UserLoginIndexSql(ifNotExists string) string {
	indexPrefix := tbl.prefixWithoutDot()
	return fmt.Sprintf("CREATE UNIQUE INDEX %s%suser_login ON %s%s (%s)", ifNotExists, indexPrefix,
		tbl.Prefix, tbl.Name, sqlV3UserLoginIndexColumns)
}

func (tbl V3UserTable) dropV3UserLoginIndexSql(ifExists, onTbl string) string {
	indexPrefix := tbl.prefixWithoutDot()
	return fmt.Sprintf("DROP INDEX %s%suser_login%s", ifExists, indexPrefix, onTbl)
}

func (tbl V3UserTable) createV3UserEmailIndexSql(ifNotExists string) string {
	indexPrefix := tbl.prefixWithoutDot()
	return fmt.Sprintf("CREATE UNIQUE INDEX %s%suser_email ON %s%s (%s)", ifNotExists, indexPrefix,
		tbl.Prefix, tbl.Name, sqlV3UserEmailIndexColumns)
}

func (tbl V3UserTable) dropV3UserEmailIndexSql(ifExists, onTbl string) string {
	indexPrefix := tbl.prefixWithoutDot()
	return fmt.Sprintf("DROP INDEX %s%suser_email%s", ifExists, indexPrefix, onTbl)
}

// DropIndexes executes queries that drop the indexes on by the User table.
func (tbl V3UserTable) DropIndexes(ifExist bool) (err error) {
	// Mysql does not support 'if exists' on indexes
	ie := tbl.ternary(ifExist && tbl.Dialect != schema.Mysql, "IF EXISTS ", "")
	onTbl := tbl.ternary(tbl.Dialect == schema.Mysql, fmt.Sprintf(" ON %s%s", tbl.Prefix, tbl.Name), "")

	_, err = tbl.Exec(tbl.dropV3UserLoginIndexSql(ie, onTbl))
	if err != nil {
		return err
	}

	_, err = tbl.Exec(tbl.dropV3UserEmailIndexSql(ie, onTbl))
	if err != nil {
		return err
	}

	return nil
}

//--------------------------------------------------------------------------------

const sqlV3UserLoginIndexColumns = "login"

const sqlV3UserEmailIndexColumns = "emailaddress"

//--------------------------------------------------------------------------------

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// It returns the number of rows affected (of the database drive supports this).
func (tbl V3UserTable) Exec(query string, args ...interface{}) (int64, error) {
	tbl.logQuery(query, args...)
	res, err := tbl.Db.ExecContext(tbl.Ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

//--------------------------------------------------------------------------------

// QueryOne is the low-level access function for one User.
func (tbl V3UserTable) QueryOne(query string, args ...interface{}) (*User, error) {
	tbl.logQuery(query, args...)
	row := tbl.Db.QueryRowContext(tbl.Ctx, query, args...)
	return scanV3User(row)
}

// Query is the low-level access function for Users.
func (tbl V3UserTable) Query(query string, args ...interface{}) ([]*User, error) {
	tbl.logQuery(query, args...)
	rows, err := tbl.Db.QueryContext(tbl.Ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanV3Users(rows)
}

// scanV3User reads a table record into a single value.
func scanV3User(row *sql.Row) (*User, error) {
	var v0 int64
	var v1 string
	var v2 string
	var v3 string
	var v4 bool
	var v5 bool
	var v6 []byte
	var v7 int64
	var v8 string
	var v9 string
	var v10 string

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
		&v10,

	)
	if err != nil {
		return nil, err
	}

	v := &User{}
	v.Uid = v0
	v.Login = v1
	v.EmailAddress = v2
	v.Avatar = v3
	v.Active = v4
	v.Admin = v5
	err = json.Unmarshal(v6, &v.Fave)
	if err != nil {
		return nil, err
	}
	v.LastUpdated = v7
	v.token = v8
	v.secret = v9
	v.hash = v10

	return v, nil
}

// scanV3Users reads table records into a slice of values.
func scanV3Users(rows *sql.Rows) ([]*User, error) {
	var err error
	var vv []*User

	var v0 int64
	var v1 string
	var v2 string
	var v3 string
	var v4 bool
	var v5 bool
	var v6 []byte
	var v7 int64
	var v8 string
	var v9 string
	var v10 string

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
			&v10,

		)
		if err != nil {
			return vv, err
		}

		v := &User{}
		v.Uid = v0
		v.Login = v1
		v.EmailAddress = v2
		v.Avatar = v3
		v.Active = v4
		v.Admin = v5
		err = json.Unmarshal(v6, &v.Fave)
		if err != nil {
			return nil, err
		}
		v.LastUpdated = v7
		v.token = v8
		v.secret = v9
		v.hash = v10

		vv = append(vv, v)
	}
	return vv, rows.Err()
}

func sliceV3UserWithoutPk(v *User) ([]interface{}, error) {

	v6, err := json.Marshal(&v.Fave)
	if err != nil {
		return nil, err
	}

	return []interface{}{
		v.Login,
		v.EmailAddress,
		v.Avatar,
		v.Active,
		v.Admin,
		v6,
		v.LastUpdated,
		v.token,
		v.secret,
		v.hash,

	}, nil
}
