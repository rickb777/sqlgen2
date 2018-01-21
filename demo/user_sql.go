// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

package demo

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/rickb777/sqlgen2"
	"github.com/rickb777/sqlgen2/require"
	"github.com/rickb777/sqlgen2/schema"
	"github.com/rickb777/sqlgen2/where"
	"log"
	"math/big"
	"strings"
)

// DbUserTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type DbUserTable struct {
	name    sqlgen2.TableName
	db      sqlgen2.Execer
	ctx     context.Context
	dialect schema.Dialect
	logger  *log.Logger
	wrapper interface{}
}

// Type conformance checks
var _ sqlgen2.TableWithIndexes = &DbUserTable{}
var _ sqlgen2.TableWithCrud = &DbUserTable{}

// NewDbUserTable returns a new table instance.
// If a blank table name is supplied, the default name "users" will be used instead.
// The request context is initialised with the background.
func NewDbUserTable(name sqlgen2.TableName, d sqlgen2.Execer, dialect schema.Dialect) DbUserTable {
	if name.Name == "" {
		name.Name = "users"
	}
	return DbUserTable{name, d, context.Background(), dialect, nil, nil}
}

// CopyTableAsDbUserTable copies a table instance, retaining the name etc but
// providing methods appropriate for 'User'.
func CopyTableAsDbUserTable(origin sqlgen2.Table) DbUserTable {
	return DbUserTable{
		name:    origin.Name(),
		db:      origin.DB(),
		ctx:     origin.Ctx(),
		dialect: origin.Dialect(),
		logger:  origin.Logger(),
	}
}

// WithPrefix sets the table name prefix for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl DbUserTable) WithPrefix(pfx string) DbUserTable {
	tbl.name.Prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl DbUserTable) WithContext(ctx context.Context) DbUserTable {
	tbl.ctx = ctx
	return tbl
}

// WithLogger sets the logger for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl DbUserTable) WithLogger(logger *log.Logger) DbUserTable {
	tbl.logger = logger
	return tbl
}

// Logger gets the trace logger.
func (tbl DbUserTable) Logger() *log.Logger {
	return tbl.logger
}

// SetLogger sets the logger for subsequent queries, returning the interface.
// The result is a modified copy of the table; the original is unchanged.
func (tbl DbUserTable) SetLogger(logger *log.Logger) sqlgen2.Table {
	tbl.logger = logger
	return tbl
}

// Ctx gets the current request context.
func (tbl DbUserTable) Ctx() context.Context {
	return tbl.ctx
}

// Dialect gets the database dialect.
func (tbl DbUserTable) Dialect() schema.Dialect {
	return tbl.dialect
}

// Wrapper gets the user-defined wrapper.
func (tbl DbUserTable) Wrapper() interface{} {
	return tbl.wrapper
}

// SetWrapper sets the user-defined wrapper.
// The result is a modified copy of the table; the original is unchanged.
func (tbl DbUserTable) SetWrapper(wrapper interface{}) sqlgen2.Table {
	tbl.wrapper = wrapper
	return tbl
}

// Name gets the table name.
func (tbl DbUserTable) Name() sqlgen2.TableName {
	return tbl.name
}

// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl DbUserTable) DB() *sql.DB {
	return tbl.db.(*sql.DB)
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl DbUserTable) Tx() *sql.Tx {
	return tbl.db.(*sql.Tx)
}

// IsTx tests whether this is within a transaction.
func (tbl DbUserTable) IsTx() bool {
	_, ok := tbl.db.(*sql.Tx)
	return ok
}

// Begin starts a transaction. The default isolation level is dependent on the driver.
// The result is a modified copy of the table; the original is unchanged.
func (tbl DbUserTable) BeginTx(opts *sql.TxOptions) (DbUserTable, error) {
	d := tbl.db.(*sql.DB)
	var err error
	tbl.db, err = d.BeginTx(tbl.ctx, opts)
	return tbl, tbl.logIfError(err)
}

// Using returns a modified Table using the transaction supplied. This is needed
// when making multiple queries across several tables within a single transaction.
// The result is a modified copy of the table; the original is unchanged.
func (tbl DbUserTable) Using(tx *sql.Tx) DbUserTable {
	tbl.db = tx
	return tbl
}

func (tbl DbUserTable) logQuery(query string, args ...interface{}) {
	sqlgen2.LogQuery(tbl.logger, query, args...)
}

func (tbl DbUserTable) logError(err error) error {
	return sqlgen2.LogError(tbl.logger, err)
}

func (tbl DbUserTable) logIfError(err error) error {
	return sqlgen2.LogIfError(tbl.logger, err)
}

//--------------------------------------------------------------------------------

const NumDbUserColumns = 11

const NumDbUserDataColumns = 10

const DbUserPk = "Uid"

const DbUserDataColumnNames = "login, emailaddress, avatar, role, active, admin, fave, lastupdated, token, secret"

//--------------------------------------------------------------------------------

// CreateTable creates the table.
func (tbl DbUserTable) CreateTable(ifNotExists bool) (int64, error) {
	return tbl.Exec(nil, tbl.createTableSql(ifNotExists))
}

func (tbl DbUserTable) createTableSql(ifNotExists bool) string {
	var stmt string
	switch tbl.dialect {
	case schema.Sqlite:
		stmt = sqlCreateDbUserTableSqlite
	case schema.Postgres:
		stmt = sqlCreateDbUserTablePostgres
	case schema.Mysql:
		stmt = sqlCreateDbUserTableMysql
	}
	extra := tbl.ternary(ifNotExists, "IF NOT EXISTS ", "")
	query := fmt.Sprintf(stmt, extra, tbl.name)
	return query
}

func (tbl DbUserTable) ternary(flag bool, a, b string) string {
	if flag {
		return a
	}
	return b
}

// DropTable drops the table, destroying all its data.
func (tbl DbUserTable) DropTable(ifExists bool) (int64, error) {
	return tbl.Exec(nil, tbl.dropTableSql(ifExists))
}

func (tbl DbUserTable) dropTableSql(ifExists bool) string {
	extra := tbl.ternary(ifExists, "IF EXISTS ", "")
	query := fmt.Sprintf("DROP TABLE %s%s", extra, tbl.name)
	return query
}

const sqlCreateDbUserTableSqlite = `
CREATE TABLE %s%s (
 uid          integer primary key autoincrement,
 login        text,
 emailaddress text,
 avatar       text default null,
 role         tinyint default null,
 active       boolean,
 admin        boolean,
 fave         text,
 lastupdated  bigint,
 token        text,
 secret       text
)
`

const sqlCreateDbUserTablePostgres = `
CREATE TABLE %s%s (
 uid          bigserial primary key,
 login        varchar(255),
 emailaddress varchar(255),
 avatar       varchar(255) default null,
 role         tinyint default null,
 active       boolean,
 admin        boolean,
 fave         json,
 lastupdated  bigint,
 token        varchar(255),
 secret       varchar(255)
)
`

const sqlCreateDbUserTableMysql = `
CREATE TABLE %s%s (
 uid          bigint primary key auto_increment,
 login        varchar(255),
 emailaddress varchar(255),
 avatar       varchar(255) default null,
 role         tinyint default null,
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
func (tbl DbUserTable) CreateTableWithIndexes(ifNotExist bool) (err error) {
	_, err = tbl.CreateTable(ifNotExist)
	if err != nil {
		return err
	}

	return tbl.CreateIndexes(ifNotExist)
}

// CreateIndexes executes queries that create the indexes needed by the User table.
func (tbl DbUserTable) CreateIndexes(ifNotExist bool) (err error) {

	err = tbl.CreateUserEmailIndex(ifNotExist)
	if err != nil {
		return err
	}

	err = tbl.CreateUserLoginIndex(ifNotExist)
	if err != nil {
		return err
	}

	return nil
}

// CreateUserEmailIndex creates the user_email index.
func (tbl DbUserTable) CreateUserEmailIndex(ifNotExist bool) error {
	ine := tbl.ternary(ifNotExist && tbl.dialect != schema.Mysql, "IF NOT EXISTS ", "")

	// Mysql does not support 'if not exists' on indexes
	// Workaround: use DropIndex first and ignore an error returned if the index didn't exist.

	if ifNotExist && tbl.dialect == schema.Mysql {
		tbl.DropUserEmailIndex(false)
		ine = ""
	}

	_, err := tbl.Exec(nil, tbl.createDbUserEmailIndexSql(ine))
	return err
}

func (tbl DbUserTable) createDbUserEmailIndexSql(ifNotExists string) string {
	indexPrefix := tbl.name.PrefixWithoutDot()
	return fmt.Sprintf("CREATE UNIQUE INDEX %s%suser_email ON %s (%s)", ifNotExists, indexPrefix,
		tbl.name, sqlDbUserEmailIndexColumns)
}

// DropUserEmailIndex drops the user_email index.
func (tbl DbUserTable) DropUserEmailIndex(ifExists bool) error {
	_, err := tbl.Exec(nil, tbl.dropDbUserEmailIndexSql(ifExists))
	return err
}

func (tbl DbUserTable) dropDbUserEmailIndexSql(ifExists bool) string {
	// Mysql does not support 'if exists' on indexes
	ie := tbl.ternary(ifExists && tbl.dialect != schema.Mysql, "IF EXISTS ", "")
	onTbl := tbl.ternary(tbl.dialect == schema.Mysql, fmt.Sprintf(" ON %s", tbl.name), "")
	indexPrefix := tbl.name.PrefixWithoutDot()
	return fmt.Sprintf("DROP INDEX %s%suser_email%s", ie, indexPrefix, onTbl)
}

// CreateUserLoginIndex creates the user_login index.
func (tbl DbUserTable) CreateUserLoginIndex(ifNotExist bool) error {
	ine := tbl.ternary(ifNotExist && tbl.dialect != schema.Mysql, "IF NOT EXISTS ", "")

	// Mysql does not support 'if not exists' on indexes
	// Workaround: use DropIndex first and ignore an error returned if the index didn't exist.

	if ifNotExist && tbl.dialect == schema.Mysql {
		tbl.DropUserLoginIndex(false)
		ine = ""
	}

	_, err := tbl.Exec(nil, tbl.createDbUserLoginIndexSql(ine))
	return err
}

func (tbl DbUserTable) createDbUserLoginIndexSql(ifNotExists string) string {
	indexPrefix := tbl.name.PrefixWithoutDot()
	return fmt.Sprintf("CREATE UNIQUE INDEX %s%suser_login ON %s (%s)", ifNotExists, indexPrefix,
		tbl.name, sqlDbUserLoginIndexColumns)
}

// DropUserLoginIndex drops the user_login index.
func (tbl DbUserTable) DropUserLoginIndex(ifExists bool) error {
	_, err := tbl.Exec(nil, tbl.dropDbUserLoginIndexSql(ifExists))
	return err
}

func (tbl DbUserTable) dropDbUserLoginIndexSql(ifExists bool) string {
	// Mysql does not support 'if exists' on indexes
	ie := tbl.ternary(ifExists && tbl.dialect != schema.Mysql, "IF EXISTS ", "")
	onTbl := tbl.ternary(tbl.dialect == schema.Mysql, fmt.Sprintf(" ON %s", tbl.name), "")
	indexPrefix := tbl.name.PrefixWithoutDot()
	return fmt.Sprintf("DROP INDEX %s%suser_login%s", ie, indexPrefix, onTbl)
}

// DropIndexes executes queries that drop the indexes on by the User table.
func (tbl DbUserTable) DropIndexes(ifExist bool) (err error) {

	err = tbl.DropUserEmailIndex(ifExist)
	if err != nil {
		return err
	}

	err = tbl.DropUserLoginIndex(ifExist)
	if err != nil {
		return err
	}

	return nil
}

//--------------------------------------------------------------------------------

const sqlDbUserEmailIndexColumns = "emailaddress"

const sqlDbUserLoginIndexColumns = "login"

//--------------------------------------------------------------------------------

// Truncate drops every record from the table, if possible. It might fail if constraints exist that
// prevent some or all rows from being deleted; use the force option to override this.
//
// When 'force' is set true, be aware of the following consequences.
// When using Mysql, foreign keys in other tables can be left dangling.
// When using Postgres, a cascade happens, so all 'adjacent' tables (i.e. linked by foreign keys)
// are also truncated.
func (tbl DbUserTable) Truncate(force bool) (err error) {
	for _, query := range tbl.dialect.TruncateDDL(tbl.Name().String(), force) {
		_, err = tbl.Exec(nil, query)
		if err != nil {
			return err
		}
	}
	return nil
}

//--------------------------------------------------------------------------------

// Exec executes a query without returning any rows.
// It returns the number of rows affected (if the database driver supports this).
//
// The args are for any placeholder parameters in the query.
func (tbl DbUserTable) Exec(req require.Requirement, query string, args ...interface{}) (int64, error) {
	tbl.logQuery(query, args...)
	res, err := tbl.db.ExecContext(tbl.ctx, query, args...)
	if err != nil {
		return 0, tbl.logError(err)
	}
	n, err := res.RowsAffected()
	return n, tbl.logIfError(require.ChainErrorIfExecNotSatisfiedBy(err, req, n))
}

//--------------------------------------------------------------------------------

// Query is the low-level access method for Users.
//
// It places a requirement, which may be nil, on the size of the expected results: this
// controls whether an error is generated when this expectation is not met.
//
// Note that this method applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl DbUserTable) Query(req require.Requirement, query string, args ...interface{}) ([]*User, error) {
	query = tbl.ReplaceTableName(query)
	vv, err := tbl.doQuery(req, false, query, args...)
	return vv, err
}

// QueryOne is the low-level access method for one User.
// If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, *User will be nil.
//
// Note that this method applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl DbUserTable) QueryOne(query string, args ...interface{}) (*User, error) {
	query = tbl.ReplaceTableName(query)
	return tbl.doQueryOne(nil, query, args...)
}

// MustQueryOne is the low-level access method for one User.
//
// It places a requirement that exactly one result must be found; an error is generated when this expectation is not met.
//
// Note that this method applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl DbUserTable) MustQueryOne(query string, args ...interface{}) (*User, error) {
	query = tbl.ReplaceTableName(query)
	return tbl.doQueryOne(require.One, query, args...)
}

func (tbl DbUserTable) doQueryOne(req require.Requirement, query string, args ...interface{}) (*User, error) {
	list, err := tbl.doQuery(req, true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

func (tbl DbUserTable) doQuery(req require.Requirement, firstOnly bool, query string, args ...interface{}) ([]*User, error) {
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, tbl.logError(err)
	}
	defer rows.Close()

	vv, n, err := scanDbUsers(rows, firstOnly)
	return vv, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(err, req, n))
}

func scanDbUsers(rows *sql.Rows, firstOnly bool) (vv []*User, n int64, err error) {
	for rows.Next() {
		n++

		var v0 int64
		var v1 string
		var v2 string
		var v3 sql.NullString
		var v4 *Role
		var v5 bool
		var v6 bool
		var v7 []byte
		var v8 int64
		var v9 string
		var v10 string

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
			return vv, n, err
		}

		v := &User{}
		v.Uid = v0
		v.Login = v1
		v.EmailAddress = v2
		if v3.Valid {
			a := v3.String
			v.Avatar = &a
		}
		v.Role = v4
		v.Active = v5
		v.Admin = v6
		err = json.Unmarshal(v7, &v.Fave)
		if err != nil {
			return nil, n, err
		}
		v.LastUpdated = v8
		v.token = v9
		v.secret = v10

		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPostGet); ok {
			err = hook.PostGet()
			if err != nil {
				return vv, n, err
			}
		}

		vv = append(vv, v)

		if firstOnly {
			if rows.Next() {
				n++
			}
			return vv, n, rows.Err()
		}
	}

	return vv, n, rows.Err()
}

//--------------------------------------------------------------------------------

// QueryOneNullString is a low-level access method for one string. This can be used for function queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl DbUserTable) QueryOneNullString(query string, args ...interface{}) (result sql.NullString, err error) {
	err = tbl.doQueryOneNullThing(nil, &result, query, args...)
	return result, err
}

// MustQueryOneNullString is a low-level access method for one string. This can be used for function queries and
// such like.
//
// It places a requirement that exactly one result must be found; an error is generated when this expectation is not met.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl DbUserTable) MustQueryOneNullString(query string, args ...interface{}) (result sql.NullString, err error) {
	err = tbl.doQueryOneNullThing(require.One, &result, query, args...)
	return result, err
}

// QueryOneNullInt64 is a low-level access method for one int64. This can be used for 'COUNT(1)' queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl DbUserTable) QueryOneNullInt64(query string, args ...interface{}) (result sql.NullInt64, err error) {
	err = tbl.doQueryOneNullThing(nil, &result, query, args...)
	return result, err
}

// MustQueryOneNullInt64 is a low-level access method for one int64. This can be used for 'COUNT(1)' queries and
// such like.
//
// It places a requirement that exactly one result must be found; an error is generated when this expectation is not met.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl DbUserTable) MustQueryOneNullInt64(query string, args ...interface{}) (result sql.NullInt64, err error) {
	err = tbl.doQueryOneNullThing(require.One, &result, query, args...)
	return result, err
}

// QueryOneNullFloat64 is a low-level access method for one float64. This can be used for 'AVG(...)' queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl DbUserTable) QueryOneNullFloat64(query string, args ...interface{}) (result sql.NullFloat64, err error) {
	err = tbl.doQueryOneNullThing(nil, &result, query, args...)
	return result, err
}

// MustQueryOneNullFloat64 is a low-level access method for one float64. This can be used for 'AVG(...)' queries and
// such like.
//
// It places a requirement that exactly one result must be found; an error is generated when this expectation is not met.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl DbUserTable) MustQueryOneNullFloat64(query string, args ...interface{}) (result sql.NullFloat64, err error) {
	err = tbl.doQueryOneNullThing(require.One, &result, query, args...)
	return result, err
}

func (tbl DbUserTable) doQueryOneNullThing(req require.Requirement, holder interface{}, query string, args ...interface{}) error {
	var n int64 = 0
	query = tbl.ReplaceTableName(query)
	tbl.logQuery(query, args...)

	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return tbl.logError(err)
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(holder)

		if err == sql.ErrNoRows {
			return tbl.logIfError(require.ErrorIfQueryNotSatisfiedBy(req, 0))
		} else {
			n++
		}

		if rows.Next() {
			n++ // not singular
		}
	}

	return tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, n))
}

// ReplaceTableName replaces all occurrences of "{TABLE}" with the table's name.
func (tbl DbUserTable) ReplaceTableName(query string) string {
	return strings.Replace(query, "{TABLE}", tbl.name.String(), -1)
}

//--------------------------------------------------------------------------------

// GetUser gets the record with a given primary key value.
// If not found, *User will be nil.
func (tbl DbUserTable) GetUser(id int64) (*User, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE uid=?", DbUserColumnNames, tbl.name)
	v, err := tbl.doQueryOne(nil, query, id)
	return v, err
}

// MustGetUser gets the record with a given primary key value.
//
// It places a requirement that exactly one result must be found; an error is generated when this expectation is not met.
func (tbl DbUserTable) MustGetUser(id int64) (*User, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE uid=?", DbUserColumnNames, tbl.name)
	v, err := tbl.doQueryOne(require.One, query, id)
	return v, err
}

// GetUsers gets records from the table according to a list of primary keys.
// Although the list of ids can be arbitrarily long, there are practical limits;
// note that Oracle DB has a limit of 1000.
//
// It places a requirement, which may be nil, on the size of the expected results: in particular, require.All
// controls whether an error is generated not all the ids produce a result.
func (tbl DbUserTable) GetUsers(req require.Requirement, id ...int64) (list []*User, err error) {
	if len(id) > 0 {
		if req == require.All {
			req = require.Exactly(len(id))
		}
		pl := tbl.dialect.Placeholders(len(id))
		query := fmt.Sprintf("SELECT %s FROM %s WHERE uid IN (%s)", DbUserColumnNames, tbl.name, pl)
		args := make([]interface{}, len(id))

		for i, v := range id {
			args[i] = v
		}

		list, err = tbl.doQuery(req, false, query, args...)
	}

	return list, err
}

//--------------------------------------------------------------------------------

// SelectOneWhere allows a single Example to be obtained from the table that match a 'where' clause
// and some limit. Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
// If not found, *Example will be nil.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.One
// controls whether an error is generated when no result is found.
//
// The args are for any placeholder parameters in the query.
func (tbl DbUserTable) SelectOneWhere(req require.Requirement, where, orderBy string, args ...interface{}) (*User, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT 1", DbUserColumnNames, tbl.name, where, orderBy)
	v, err := tbl.doQueryOne(req, query, args...)
	return v, err
}

// SelectOne allows a single User to be obtained from the sqlgen2.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
// If not found, *Example will be nil.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.One
// controls whether an error is generated when no result is found.
func (tbl DbUserTable) SelectOne(req require.Requirement, wh where.Expression, qc where.QueryConstraint) (*User, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	return tbl.SelectOneWhere(req, whs, orderBy, args...)
}

// SelectWhere allows Users to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
//
// The args are for any placeholder parameters in the query.
func (tbl DbUserTable) SelectWhere(req require.Requirement, where, orderBy string, args ...interface{}) ([]*User, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", DbUserColumnNames, tbl.name, where, orderBy)
	vv, err := tbl.doQuery(req, false, query, args...)
	return vv, err
}

// Select allows Users to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
func (tbl DbUserTable) Select(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]*User, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	return tbl.SelectWhere(req, whs, orderBy, args...)
}

// CountWhere counts Users in the table that match a 'where' clause.
// Use a blank string for the 'where' argument if it is not needed.
//
// The args are for any placeholder parameters in the query.
func (tbl DbUserTable) CountWhere(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s %s", tbl.name, where)
	tbl.logQuery(query, args...)
	row := tbl.db.QueryRowContext(tbl.ctx, query, args...)
	err = row.Scan(&count)
	return count, tbl.logIfError(err)
}

// Count counts the Users in the table that match a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed.
func (tbl DbUserTable) Count(wh where.Expression) (count int64, err error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	return tbl.CountWhere(whs, args...)
}

const DbUserColumnNames = "uid, login, emailaddress, avatar, role, active, admin, fave, lastupdated, token, secret"

//--------------------------------------------------------------------------------

// SliceUid gets the Uid column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DbUserTable) SliceUid(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return tbl.getint64list(req, "uid", wh, qc)
}

// SliceLogin gets the Login column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DbUserTable) SliceLogin(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.getstringlist(req, "login", wh, qc)
}

// SliceEmailAddress gets the EmailAddress column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DbUserTable) SliceEmailaddress(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.getstringlist(req, "emailaddress", wh, qc)
}

// SliceAvatar gets the Avatar column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DbUserTable) SliceAvatar(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.getstringPtrlist(req, "avatar", wh, qc)
}

// SliceRole gets the Role column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DbUserTable) SliceRole(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]Role, error) {
	return tbl.getRolePtrlist(req, "role", wh, qc)
}

// SliceActive gets the Active column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DbUserTable) SliceActive(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]bool, error) {
	return tbl.getboollist(req, "active", wh, qc)
}

// SliceAdmin gets the Admin column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DbUserTable) SliceAdmin(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]bool, error) {
	return tbl.getboollist(req, "admin", wh, qc)
}

// SliceLastUpdated gets the LastUpdated column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DbUserTable) SliceLastupdated(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return tbl.getint64list(req, "lastupdated", wh, qc)
}

func (tbl DbUserTable) getRolePtrlist(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]Role, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", sqlname, tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, tbl.logError(err)
	}
	defer rows.Close()

	var v Role
	list := make([]Role, 0, 10)

	for rows.Next() {
		err = rows.Scan(&v)
		if err == sql.ErrNoRows {
			return list, tbl.logIfError(require.ErrorIfQueryNotSatisfiedBy(req, int64(len(list))))
		} else {
			list = append(list, v)
		}
	}
	return list, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, int64(len(list))))
}

func (tbl DbUserTable) getboollist(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]bool, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", sqlname, tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, tbl.logError(err)
	}
	defer rows.Close()

	var v bool
	list := make([]bool, 0, 10)

	for rows.Next() {
		err = rows.Scan(&v)
		if err == sql.ErrNoRows {
			return list, tbl.logIfError(require.ErrorIfQueryNotSatisfiedBy(req, int64(len(list))))
		} else {
			list = append(list, v)
		}
	}
	return list, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, int64(len(list))))
}

func (tbl DbUserTable) getint64list(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", sqlname, tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, tbl.logError(err)
	}
	defer rows.Close()

	var v int64
	list := make([]int64, 0, 10)

	for rows.Next() {
		err = rows.Scan(&v)
		if err == sql.ErrNoRows {
			return list, tbl.logIfError(require.ErrorIfQueryNotSatisfiedBy(req, int64(len(list))))
		} else {
			list = append(list, v)
		}
	}
	return list, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, int64(len(list))))
}

func (tbl DbUserTable) getstringlist(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", sqlname, tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, tbl.logError(err)
	}
	defer rows.Close()

	var v string
	list := make([]string, 0, 10)

	for rows.Next() {
		err = rows.Scan(&v)
		if err == sql.ErrNoRows {
			return list, tbl.logIfError(require.ErrorIfQueryNotSatisfiedBy(req, int64(len(list))))
		} else {
			list = append(list, v)
		}
	}
	return list, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, int64(len(list))))
}

func (tbl DbUserTable) getstringPtrlist(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", sqlname, tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, tbl.logError(err)
	}
	defer rows.Close()

	var v string
	list := make([]string, 0, 10)

	for rows.Next() {
		err = rows.Scan(&v)
		if err == sql.ErrNoRows {
			return list, tbl.logIfError(require.ErrorIfQueryNotSatisfiedBy(req, int64(len(list))))
		} else {
			list = append(list, v)
		}
	}
	return list, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, int64(len(list))))
}

//--------------------------------------------------------------------------------

// Insert adds new records for the Users.
// The Users have their primary key fields set to the new record identifiers.
// The User.PreInsert(Execer) method will be called, if it exists.
func (tbl DbUserTable) Insert(req require.Requirement, vv ...*User) error {
	var params string
	switch tbl.dialect {
	case schema.Postgres:
		params = sDbUserDataColumnParamsPostgres
	default:
		params = sDbUserDataColumnParamsSimple
	}

	if req == require.All {
		req = require.Exactly(len(vv))
	}

	var count int64
	query := fmt.Sprintf(sqlInsertDbUser, tbl.name, params)
	st, err := tbl.db.PrepareContext(tbl.ctx, query)
	if err != nil {
		return err
	}
	defer st.Close()

	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPreInsert); ok {
			err := hook.PreInsert()
			if err != nil {
				return tbl.logError(err)
			}
		}

		fields, err := sliceDbUserWithoutPk(v)
		if err != nil {
			return tbl.logError(err)
		}

		tbl.logQuery(query, fields...)
		res, err := st.ExecContext(tbl.ctx, fields...)
		if err != nil {
			return tbl.logError(err)
		}

		v.Uid, err = res.LastInsertId()
		if err != nil {
			return tbl.logError(err)
		}

		n, err := res.RowsAffected()
		if err != nil {
			return tbl.logError(err)
		}
		count += n
	}

	return tbl.logIfError(require.ErrorIfExecNotSatisfiedBy(req, count))
}

const sqlInsertDbUser = `
INSERT INTO %s (
	login,
	emailaddress,
	avatar,
	role,
	active,
	admin,
	fave,
	lastupdated,
	token,
	secret
) VALUES (%s)
`

const sDbUserDataColumnParamsSimple = "?,?,?,?,?,?,?,?,?,?"

const sDbUserDataColumnParamsPostgres = "$1,$2,$3,$4,$5,$6,$7,$8,$9,$10"

//--------------------------------------------------------------------------------

// UpdateFields updates one or more columns, given a 'where' clause.
//
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl DbUserTable) UpdateFields(req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error) {
	query, args := tbl.updateFields(wh, fields...)
	return tbl.Exec(req, query, args...)
}

func (tbl DbUserTable) updateFields(wh where.Expression, fields ...sql.NamedArg) (string, []interface{}) {
	list := sqlgen2.NamedArgList(fields)
	assignments := strings.Join(list.Assignments(tbl.dialect, 1), ", ")
	whs, wargs := where.BuildExpression(wh, tbl.dialect)
	query := fmt.Sprintf("UPDATE %s SET %s %s", tbl.name, assignments, whs)
	args := append(list.Values(), wargs...)
	return query, args
}

//--------------------------------------------------------------------------------

// Update updates records, matching them by primary key. It returns the number of rows affected.
// The User.PreUpdate(Execer) method will be called, if it exists.
func (tbl DbUserTable) Update(req require.Requirement, vv ...*User) (int64, error) {
	var stmt string
	switch tbl.dialect {
	case schema.Postgres:
		stmt = sqlUpdateDbUserByPkPostgres
	default:
		stmt = sqlUpdateDbUserByPkSimple
	}

	if req == require.All {
		req = require.Exactly(len(vv))
	}

	var count int64
	query := fmt.Sprintf(stmt, tbl.name)

	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPreUpdate); ok {
			err := hook.PreUpdate()
			if err != nil {
				return count, tbl.logError(err)
			}
		}

		args, err := sliceDbUserWithoutPk(v)
		args = append(args, v.Uid)
		if err != nil {
			return count, tbl.logError(err)
		}

		n, err := tbl.Exec(nil, query, args...)
		if err != nil {
			return count, err
		}
		count += n
	}

	return count, tbl.logIfError(require.ErrorIfExecNotSatisfiedBy(req, count))
}

const sqlUpdateDbUserByPkSimple = `
UPDATE %s SET
	login=?,
	emailaddress=?,
	avatar=?,
	role=?,
	active=?,
	admin=?,
	fave=?,
	lastupdated=?,
	token=?,
	secret=?
WHERE uid=?
`

const sqlUpdateDbUserByPkPostgres = `
UPDATE %s SET
	login=$2,
	emailaddress=$3,
	avatar=$4,
	role=$5,
	active=$6,
	admin=$7,
	fave=$8,
	lastupdated=$9,
	token=$10,
	secret=$11
WHERE uid=$1
`

func sliceDbUserWithoutPk(v *User) ([]interface{}, error) {

	v7, err := json.Marshal(&v.Fave)
	if err != nil {
		return nil, err
	}

	return []interface{}{
		v.Login,
		v.EmailAddress,
		v.Avatar,
		v.Role,
		v.Active,
		v.Admin,
		v7,
		v.LastUpdated,
		v.token,
		v.secret,
	}, nil
}

//--------------------------------------------------------------------------------

// DeleteUsers deletes rows from the table, given some primary keys.
// The list of ids can be arbitrarily long.
func (tbl DbUserTable) DeleteUsers(req require.Requirement, id ...int64) (int64, error) {
	const batch = 1000 // limited by Oracle DB
	const qt = "DELETE FROM %s WHERE uid IN (%s)"

	if req == require.All {
		req = require.Exactly(len(id))
	}

	var count, n int64
	var err error
	var max = batch
	if len(id) < batch {
		max = len(id)
	}
	args := make([]interface{}, max)

	if len(id) > batch {
		pl := tbl.dialect.Placeholders(batch)
		query := fmt.Sprintf(qt, tbl.name, pl)

		for len(id) > batch {
			for i := 0; i < batch; i++ {
				args[i] = id[i]
			}

			n, err = tbl.Exec(nil, query, args...)
			count += n
			if err != nil {
				return count, err
			}

			id = id[batch:]
		}
	}

	if len(id) > 0 {
		pl := tbl.dialect.Placeholders(len(id))
		query := fmt.Sprintf(qt, tbl.name, pl)

		for i := 0; i < len(id); i++ {
			args[i] = id[i]
		}

		n, err = tbl.Exec(nil, query, args...)
		count += n
	}

	return count, tbl.logIfError(require.ChainErrorIfExecNotSatisfiedBy(err, req, n))
}

// Delete deletes one or more rows from the table, given a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl DbUserTable) Delete(req require.Requirement, wh where.Expression) (int64, error) {
	query, args := tbl.deleteRows(wh)
	return tbl.Exec(req, query, args...)
}

func (tbl DbUserTable) deleteRows(wh where.Expression) (string, []interface{}) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	query := fmt.Sprintf("DELETE FROM %s %s", tbl.name, whs)
	return query, args
}

//--------------------------------------------------------------------------------

//--------------------------------------------------------------------------------

// SetUid sets the Uid field and returns the modified User.
func (v *User) SetUid(x int64) *User {
	v.Uid = x
	return v
}

// SetLogin sets the Login field and returns the modified User.
func (v *User) SetLogin(x string) *User {
	v.Login = x
	return v
}

// SetEmailAddress sets the EmailAddress field and returns the modified User.
func (v *User) SetEmailAddress(x string) *User {
	v.EmailAddress = x
	return v
}

// SetAvatar sets the Avatar field and returns the modified User.
func (v *User) SetAvatar(x string) *User {
	v.Avatar = &x
	return v
}

// SetRole sets the Role field and returns the modified User.
func (v *User) SetRole(x Role) *User {
	v.Role = &x
	return v
}

// SetActive sets the Active field and returns the modified User.
func (v *User) SetActive(x bool) *User {
	v.Active = x
	return v
}

// SetAdmin sets the Admin field and returns the modified User.
func (v *User) SetAdmin(x bool) *User {
	v.Admin = x
	return v
}

// SetFave sets the Fave field and returns the modified User.
func (v *User) SetFave(x big.Int) *User {
	v.Fave = &x
	return v
}

// SetLastUpdated sets the LastUpdated field and returns the modified User.
func (v *User) SetLastUpdated(x int64) *User {
	v.LastUpdated = x
	return v
}

// Settoken sets the token field and returns the modified User.
func (v *User) Settoken(x string) *User {
	v.token = x
	return v
}

// Setsecret sets the secret field and returns the modified User.
func (v *User) Setsecret(x string) *User {
	v.secret = x
	return v
}
