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

// UserTbl holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type UserTbl struct {
	prefix, name string
	db           sqlgen2.Execer
	ctx          context.Context
	dialect      schema.Dialect
	logger       *log.Logger
}

// Type conformance check
var _ sqlgen2.TableWithIndexes = &UserTbl{}

// NewUserTbl returns a new table instance.
// If a blank table name is supplied, the default name "users" will be used instead.
// The table name prefix is initially blank and the request context is the background.
func NewUserTbl(name string, d sqlgen2.Execer, dialect schema.Dialect) UserTbl {
	if name == "" {
		name = "users"
	}
	return UserTbl{"", name, d, context.Background(), dialect, nil}
}

// CopyTableAsUserTbl copies a table instance, retaining the name etc but
// providing methods appropriate for 'User'.
func CopyTableAsUserTbl(origin sqlgen2.Table) UserTbl {
	return UserTbl{
		prefix:  origin.Prefix(),
		name:    origin.Name(),
		db:      origin.DB(),
		ctx:     origin.Ctx(),
		dialect: origin.Dialect(),
		logger:  origin.Logger(),
	}
}

// WithPrefix sets the table name prefix for subsequent queries.
func (tbl UserTbl) WithPrefix(pfx string) UserTbl {
	tbl.prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries.
func (tbl UserTbl) WithContext(ctx context.Context) UserTbl {
	tbl.ctx = ctx
	return tbl
}

// WithLogger sets the logger for subsequent queries.
func (tbl UserTbl) WithLogger(logger *log.Logger) UserTbl {
	tbl.logger = logger
	return tbl
}

// Ctx gets the current request context.
func (tbl UserTbl) Ctx() context.Context {
	return tbl.ctx
}

// Dialect gets the database dialect.
func (tbl UserTbl) Dialect() schema.Dialect {
	return tbl.dialect
}

// Logger gets the trace logger.
func (tbl UserTbl) Logger() *log.Logger {
	return tbl.logger
}

// SetLogger sets the logger for subsequent queries, returning the interface.
func (tbl UserTbl) SetLogger(logger *log.Logger) sqlgen2.Table {
	tbl.logger = logger
	return tbl
}

// Name gets the table name.
func (tbl UserTbl) Name() string {
	return tbl.name
}

// Prefix gets the table name prefix.
func (tbl UserTbl) Prefix() string {
	return tbl.prefix
}

// FullName gets the concatenated prefix and table name.
func (tbl UserTbl) FullName() string {
	return tbl.prefix + tbl.name
}

func (tbl UserTbl) prefixWithoutDot() string {
	last := len(tbl.prefix)-1
	if last > 0 && tbl.prefix[last] == '.' {
		return tbl.prefix[0:last]
	}
	return tbl.prefix
}

// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl UserTbl) DB() *sql.DB {
	return tbl.db.(*sql.DB)
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl UserTbl) Tx() *sql.Tx {
	return tbl.db.(*sql.Tx)
}

// IsTx tests whether this is within a transaction.
func (tbl UserTbl) IsTx() bool {
	_, ok := tbl.db.(*sql.Tx)
	return ok
}

// Begin starts a transaction. The default isolation level is dependent on the driver.
func (tbl UserTbl) BeginTx(opts *sql.TxOptions) (UserTbl, error) {
	d := tbl.db.(*sql.DB)
	var err error
	tbl.db, err = d.BeginTx(tbl.ctx, opts)
	return tbl, err
}

func (tbl UserTbl) logQuery(query string, args ...interface{}) {
	sqlgen2.LogQuery(tbl.logger, query, args...)
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
	switch tbl.dialect {
	case schema.Sqlite: stmt = sqlCreateUserTblSqlite
    case schema.Postgres: stmt = sqlCreateUserTblPostgres
    case schema.Mysql: stmt = sqlCreateUserTblMysql
    }
	extra := tbl.ternary(ifNotExists, "IF NOT EXISTS ", "")
	query := fmt.Sprintf(stmt, extra, tbl.prefix, tbl.name)
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
	query := fmt.Sprintf("DROP TABLE %s%s%s", extra, tbl.prefix, tbl.name)
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
	ine := tbl.ternary(ifNotExist && tbl.dialect != schema.Mysql, "IF NOT EXISTS ", "")

	// Mysql does not support 'if not exists' on indexes
	// Workaround: use DropIndex first and ignore an error returned if the index didn't exist.

	if ifNotExist && tbl.dialect == schema.Mysql {
		tbl.DropUserLoginIndex(false)
		ine = ""
	}

	_, err := tbl.Exec(tbl.createUserLoginIndexSql(ine))
	return err
}

func (tbl UserTbl) createUserLoginIndexSql(ifNotExists string) string {
	indexPrefix := tbl.prefixWithoutDot()
	return fmt.Sprintf("CREATE UNIQUE INDEX %s%suser_login ON %s%s (%s)", ifNotExists, indexPrefix,
		tbl.prefix, tbl.name, sqlUserLoginIndexColumns)
}

// DropUserLoginIndex drops the user_login index.
func (tbl UserTbl) DropUserLoginIndex(ifExists bool) error {
	_, err := tbl.Exec(tbl.dropUserLoginIndexSql(ifExists))
	return err
}

func (tbl UserTbl) dropUserLoginIndexSql(ifExists bool) string {
	// Mysql does not support 'if exists' on indexes
	ie := tbl.ternary(ifExists && tbl.dialect != schema.Mysql, "IF EXISTS ", "")
	onTbl := tbl.ternary(tbl.dialect == schema.Mysql, fmt.Sprintf(" ON %s%s", tbl.prefix, tbl.name), "")
	indexPrefix := tbl.prefixWithoutDot()
	return fmt.Sprintf("DROP INDEX %s%suser_login%s", ie, indexPrefix, onTbl)
}

// CreateUserEmailIndex creates the user_email index.
func (tbl UserTbl) CreateUserEmailIndex(ifNotExist bool) error {
	ine := tbl.ternary(ifNotExist && tbl.dialect != schema.Mysql, "IF NOT EXISTS ", "")

	// Mysql does not support 'if not exists' on indexes
	// Workaround: use DropIndex first and ignore an error returned if the index didn't exist.

	if ifNotExist && tbl.dialect == schema.Mysql {
		tbl.DropUserEmailIndex(false)
		ine = ""
	}

	_, err := tbl.Exec(tbl.createUserEmailIndexSql(ine))
	return err
}

func (tbl UserTbl) createUserEmailIndexSql(ifNotExists string) string {
	indexPrefix := tbl.prefixWithoutDot()
	return fmt.Sprintf("CREATE UNIQUE INDEX %s%suser_email ON %s%s (%s)", ifNotExists, indexPrefix,
		tbl.prefix, tbl.name, sqlUserEmailIndexColumns)
}

// DropUserEmailIndex drops the user_email index.
func (tbl UserTbl) DropUserEmailIndex(ifExists bool) error {
	_, err := tbl.Exec(tbl.dropUserEmailIndexSql(ifExists))
	return err
}

func (tbl UserTbl) dropUserEmailIndexSql(ifExists bool) string {
	// Mysql does not support 'if exists' on indexes
	ie := tbl.ternary(ifExists && tbl.dialect != schema.Mysql, "IF EXISTS ", "")
	onTbl := tbl.ternary(tbl.dialect == schema.Mysql, fmt.Sprintf(" ON %s%s", tbl.prefix, tbl.name), "")
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

// Truncate drops every record from the table, if possible. It might fail if constraints exist that
// prevent some or all rows from being deleted; use the force option to override this.
//
// When 'force' is set true, be aware of the following consequences.
// When using Mysql, foreign keys in other tables can be left dangling.
// When using Postgres, a cascade happens, so all 'adjacent' tables (i.e. linked by foreign keys)
// are also truncated.
func (tbl UserTbl) Truncate(force bool) (err error) {
	for _, query := range tbl.dialect.TruncateDDL(tbl.FullName(), force) {
		_, err = tbl.Exec(query)
		if err != nil {
			return err
		}
	}
	return nil
}

//--------------------------------------------------------------------------------

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// It returns the number of rows affected (of the database drive supports this).
func (tbl UserTbl) Exec(query string, args ...interface{}) (int64, error) {
	tbl.logQuery(query, args...)
	res, err := tbl.db.ExecContext(tbl.ctx, query, args...)
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
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
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
