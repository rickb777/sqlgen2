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

// UserTblName is the default name for this table.
const UserTblName = "users"

// UserTbl holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type UserTbl struct {
	Prefix, Name string
	Db           sqlgen2.Execer
	Ctx          context.Context
	Dialect      schema.Dialect
	Logger       *log.Logger
}

// Type conformance check
var _ sqlgen2.Table = &UserTbl{}

// NewUserTbl returns a new table instance.
// If a blank table name is supplied, the default name "users" will be used instead.
// The table name prefix is initially blank and the request context is the background.
func NewUserTbl(name string, d sqlgen2.Execer, dialect schema.Dialect) UserTbl {
	if name == "" {
		name = UserTblName
	}
	return UserTbl{"", name, d, context.Background(), dialect, nil}
}

// WithPrefix sets the prefix for subsequent queries.
func (tbl UserTbl) WithPrefix(pfx string) UserTbl {
	tbl.Prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries.
func (tbl UserTbl) WithContext(ctx context.Context) UserTbl {
	tbl.Ctx = ctx
	return tbl
}

// WithLogger sets the logger for subsequent queries.
func (tbl UserTbl) WithLogger(logger *log.Logger) UserTbl {
	tbl.Logger = logger
	return tbl
}

// SetLogger sets the logger for subsequent queries, returning the interface.
func (tbl UserTbl) SetLogger(logger *log.Logger) sqlgen2.Table {
	tbl.Logger = logger
	return tbl
}

// FullName gets the concatenated prefix and table name.
func (tbl UserTbl) FullName() string {
	return tbl.Prefix + tbl.Name
}

func (tbl UserTbl) prefixWithoutDot() string {
	last := len(tbl.Prefix)-1
	if last > 0 && tbl.Prefix[last] == '.' {
		return tbl.Prefix[0:last]
	}
	return tbl.Prefix
}

// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl UserTbl) DB() *sql.DB {
	return tbl.Db.(*sql.DB)
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl UserTbl) Tx() *sql.Tx {
	return tbl.Db.(*sql.Tx)
}

// IsTx tests whether this is within a transaction.
func (tbl UserTbl) IsTx() bool {
	_, ok := tbl.Db.(*sql.Tx)
	return ok
}

// Begin starts a transaction. The default isolation level is dependent on the driver.
func (tbl UserTbl) BeginTx(opts *sql.TxOptions) (UserTbl, error) {
	d := tbl.Db.(*sql.DB)
	var err error
	tbl.Db, err = d.BeginTx(tbl.Ctx, opts)
	return tbl, err
}

func (tbl UserTbl) logQuery(query string, args ...interface{}) {
	sqlgen2.LogQuery(tbl.Logger, query, args...)
}


//--------------------------------------------------------------------------------

const NumUserColumns = 10

const NumUserDataColumns = 9

const UserPk = "Uid"

const UserDataColumnNames = "login, emailaddress, avatar, active, admin, fave, lastupdated, token, secret"

//--------------------------------------------------------------------------------

// CreateTable creates the table.
func (tbl UserTbl) CreateTable(ifNotExists bool) (int64, error) {
	return tbl.Exec(tbl.createTableSql(ifNotExists))
}

func (tbl UserTbl) createTableSql(ifNotExists bool) string {
	var stmt string
	switch tbl.Dialect {
	case schema.Sqlite: stmt = sqlCreateUserTblSqlite
    case schema.Postgres: stmt = sqlCreateUserTblPostgres
    case schema.Mysql: stmt = sqlCreateUserTblMysql
    }
	extra := tbl.ternary(ifNotExists, "IF NOT EXISTS ", "")
	query := fmt.Sprintf(stmt, extra, tbl.Prefix, tbl.Name)
	return query
}

func (tbl UserTbl) ternary(flag bool, a, b string) string {
	if flag {
		return a
	}
	return b
}

// DropTable drops the table, destroying all its data.
func (tbl UserTbl) DropTable(ifExists bool) (int64, error) {
	return tbl.Exec(tbl.dropTableSql(ifExists))
}

func (tbl UserTbl) dropTableSql(ifExists bool) string {
	extra := tbl.ternary(ifExists, "IF EXISTS ", "")
	query := fmt.Sprintf("DROP TABLE %s%s%s", extra, tbl.Prefix, tbl.Name)
	return query
}

const sqlCreateUserTblSqlite = `
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
 secret       text
)
`

const sqlCreateUserTblPostgres = `
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
 secret       varchar(255)
)
`

const sqlCreateUserTblMysql = `
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
 secret       varchar(255)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
`

//--------------------------------------------------------------------------------

// CreateTableWithIndexes invokes CreateTable then CreateIndexes.
func (tbl UserTbl) CreateTableWithIndexes(ifNotExist bool) (err error) {
	_, err = tbl.CreateTable(ifNotExist)
	if err != nil {
		return err
	}

	return tbl.CreateIndexes(ifNotExist)
}

// CreateIndexes executes queries that create the indexes needed by the User table.
func (tbl UserTbl) CreateIndexes(ifNotExist bool) (err error) {

	err = tbl.CreateUserLoginIndex(ifNotExist)
	if err != nil {
		return err
	}

	err = tbl.CreateUserEmailIndex(ifNotExist)
	if err != nil {
		return err
	}

	return nil
}

// CreateUserLoginIndex creates the user_login index.
func (tbl UserTbl) CreateUserLoginIndex(ifNotExist bool) error {
	ine := tbl.ternary(ifNotExist && tbl.Dialect != schema.Mysql, "IF NOT EXISTS ", "")

	// Mysql does not support 'if not exists' on indexes
	// Workaround: use DropIndex first and ignore an error returned if the index didn't exist.

	if ifNotExist && tbl.Dialect == schema.Mysql {
		tbl.DropUserLoginIndex(false)
		ine = ""
	}

	_, err := tbl.Exec(tbl.createUserLoginIndexSql(ine))
	return err
}

func (tbl UserTbl) createUserLoginIndexSql(ifNotExists string) string {
	indexPrefix := tbl.prefixWithoutDot()
	return fmt.Sprintf("CREATE UNIQUE INDEX %s%suser_login ON %s%s (%s)", ifNotExists, indexPrefix,
		tbl.Prefix, tbl.Name, sqlUserLoginIndexColumns)
}

// DropUserLoginIndex drops the user_login index.
func (tbl UserTbl) DropUserLoginIndex(ifExists bool) error {
	_, err := tbl.Exec(tbl.dropUserLoginIndexSql(ifExists))
	return err
}

func (tbl UserTbl) dropUserLoginIndexSql(ifExists bool) string {
	// Mysql does not support 'if exists' on indexes
	ie := tbl.ternary(ifExists && tbl.Dialect != schema.Mysql, "IF EXISTS ", "")
	onTbl := tbl.ternary(tbl.Dialect == schema.Mysql, fmt.Sprintf(" ON %s%s", tbl.Prefix, tbl.Name), "")
	indexPrefix := tbl.prefixWithoutDot()
	return fmt.Sprintf("DROP INDEX %s%suser_login%s", ie, indexPrefix, onTbl)
}

// CreateUserEmailIndex creates the user_email index.
func (tbl UserTbl) CreateUserEmailIndex(ifNotExist bool) error {
	ine := tbl.ternary(ifNotExist && tbl.Dialect != schema.Mysql, "IF NOT EXISTS ", "")

	// Mysql does not support 'if not exists' on indexes
	// Workaround: use DropIndex first and ignore an error returned if the index didn't exist.

	if ifNotExist && tbl.Dialect == schema.Mysql {
		tbl.DropUserEmailIndex(false)
		ine = ""
	}

	_, err := tbl.Exec(tbl.createUserEmailIndexSql(ine))
	return err
}

func (tbl UserTbl) createUserEmailIndexSql(ifNotExists string) string {
	indexPrefix := tbl.prefixWithoutDot()
	return fmt.Sprintf("CREATE UNIQUE INDEX %s%suser_email ON %s%s (%s)", ifNotExists, indexPrefix,
		tbl.Prefix, tbl.Name, sqlUserEmailIndexColumns)
}

// DropUserEmailIndex drops the user_email index.
func (tbl UserTbl) DropUserEmailIndex(ifExists bool) error {
	_, err := tbl.Exec(tbl.dropUserEmailIndexSql(ifExists))
	return err
}

func (tbl UserTbl) dropUserEmailIndexSql(ifExists bool) string {
	// Mysql does not support 'if exists' on indexes
	ie := tbl.ternary(ifExists && tbl.Dialect != schema.Mysql, "IF EXISTS ", "")
	onTbl := tbl.ternary(tbl.Dialect == schema.Mysql, fmt.Sprintf(" ON %s%s", tbl.Prefix, tbl.Name), "")
	indexPrefix := tbl.prefixWithoutDot()
	return fmt.Sprintf("DROP INDEX %s%suser_email%s", ie, indexPrefix, onTbl)
}

// DropIndexes executes queries that drop the indexes on by the User table.
func (tbl UserTbl) DropIndexes(ifExist bool) (err error) {

	err = tbl.DropUserLoginIndex(ifExist)
	if err != nil {
		return err
	}

	err = tbl.DropUserEmailIndex(ifExist)
	if err != nil {
		return err
	}

	return nil
}

//--------------------------------------------------------------------------------

const sqlUserLoginIndexColumns = "login"

const sqlUserEmailIndexColumns = "emailaddress"

//--------------------------------------------------------------------------------

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// It returns the number of rows affected (of the database drive supports this).
func (tbl UserTbl) Exec(query string, args ...interface{}) (int64, error) {
	tbl.logQuery(query, args...)
	res, err := tbl.Db.ExecContext(tbl.Ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

//--------------------------------------------------------------------------------

// QueryOne is the low-level access function for one User.
// If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, *User will be nil.
func (tbl UserTbl) QueryOne(query string, args ...interface{}) (*User, error) {
	list, err := tbl.doQuery(true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

// Query is the low-level access function for Users.
func (tbl UserTbl) Query(query string, args ...interface{}) ([]*User, error) {
	return tbl.doQuery(false, query, args...)
}

func (tbl UserTbl) doQuery(firstOnly bool, query string, args ...interface{}) ([]*User, error) {
	tbl.logQuery(query, args...)
	rows, err := tbl.Db.QueryContext(tbl.Ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanUsers(rows, firstOnly)
}

// scanUsers reads table records into a slice of values.
func scanUsers(rows *sql.Rows, firstOnly bool) ([]*User, error) {
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

		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPostGet); ok {
			err = hook.PostGet()
			if err != nil {
				return vv, err
			}
		}

		vv = append(vv, v)

		if firstOnly {
			return vv, rows.Err()
		}
	}

	return vv, rows.Err()
}

func sliceUserWithoutPk(v *User) ([]interface{}, error) {

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

	}, nil
}
