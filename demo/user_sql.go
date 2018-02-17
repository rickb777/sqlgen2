// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

package demo

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/rickb777/sqlgen2"
	"github.com/rickb777/sqlgen2/constraint"
	"github.com/rickb777/sqlgen2/require"
	"github.com/rickb777/sqlgen2/schema"
	"github.com/rickb777/sqlgen2/support"
	"github.com/rickb777/sqlgen2/where"
	"io"
	"log"
	"math/big"
	"strings"
)

// DbUserTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type DbUserTable struct {
	name        sqlgen2.TableName
	database    *sqlgen2.Database
	db          sqlgen2.Execer
	constraints constraint.Constraints
	ctx         context.Context
}

// Type conformance checks
var _ sqlgen2.TableWithIndexes = &DbUserTable{}
var _ sqlgen2.TableWithCrud = &DbUserTable{}

// NewDbUserTable returns a new table instance.
// If a blank table name is supplied, the default name "users" will be used instead.
// The request context is initialised with the background.
func NewDbUserTable(name sqlgen2.TableName, d *sqlgen2.Database) DbUserTable {
	if name.Name == "" {
		name.Name = "users"
	}
	table := DbUserTable{name, d, d.DB(), nil, context.Background()}
	table.constraints = append(table.constraints,
		constraint.FkConstraint{"addressid", constraint.Reference{"addresses", "id"}, "restrict", "restrict"})

	return table
}

// CopyTableAsDbUserTable copies a table instance, retaining the name etc but
// providing methods appropriate for 'User'. It doesn't copy the constraints of the original table.
//
// It serves to provide methods appropriate for 'User'. This is most useful when this is used to represent a
// join result. In such cases, there won't be any need for DDL methods, nor Exec, Insert, Update or Delete.
func CopyTableAsDbUserTable(origin sqlgen2.Table) DbUserTable {
	return DbUserTable{
		name:        origin.Name(),
		database:    origin.Database(),
		db:          origin.DB(),
		constraints: nil,
		ctx:         context.Background(),
	}
}

// WithPrefix sets the table name prefix for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl DbUserTable) WithPrefix(pfx string) DbUserTable {
	tbl.name.Prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries via this table.
// The result is a modified copy of the table; the original is unchanged.
//
// The shared context in the *Database is not altered by this method. So it
// is possible to use different contexts for different (groups of) queries.
func (tbl DbUserTable) WithContext(ctx context.Context) DbUserTable {
	tbl.ctx = ctx
	return tbl
}

// Database gets the shared database information.
func (tbl DbUserTable) Database() *sqlgen2.Database {
	return tbl.database
}

// Logger gets the trace logger.
func (tbl DbUserTable) Logger() *log.Logger {
	return tbl.database.Logger()
}

// WithConstraint returns a modified Table with added data consistency constraints.
func (tbl DbUserTable) WithConstraint(cc ...constraint.Constraint) DbUserTable {
	tbl.constraints = append(tbl.constraints, cc...)
	return tbl
}

// Constraints returns the table's constraints.
func (tbl DbUserTable) Constraints() constraint.Constraints {
	return tbl.constraints
}

// Dialect gets the database dialect.
func (tbl DbUserTable) Dialect() schema.Dialect {
	return tbl.database.Dialect()
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

// Execer gets the wrapped database or transaction handle.
func (tbl DbUserTable) Execer() sqlgen2.Execer {
	return tbl.db
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

// BeginTx starts a transaction using the table's context.
// This context is used until the transaction is committed or rolled back.
//
// If this context is cancelled, the sql package will roll back the transaction.
// In this case, Tx.Commit will then return an error.
//
// The provided TxOptions is optional and may be nil if defaults should be used.
// If a non-default isolation level is used that the driver doesn't support,
// an error will be returned.
//
// Panics if the Execer is not TxStarter.
func (tbl DbUserTable) BeginTx(opts *sql.TxOptions) (DbUserTable, error) {
	var err error
	tbl.db, err = tbl.db.(sqlgen2.TxStarter).BeginTx(tbl.ctx, opts)
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
	tbl.database.LogQuery(query, args...)
}

func (tbl DbUserTable) logError(err error) error {
	return tbl.database.LogError(err)
}

func (tbl DbUserTable) logIfError(err error) error {
	return tbl.database.LogIfError(err)
}

// ReplaceTableName replaces all occurrences of "{TABLE}" with the table's name.
func (tbl DbUserTable) ReplaceTableName(query string) string {
	return strings.Replace(query, "{TABLE}", tbl.name.String(), -1)
}

//--------------------------------------------------------------------------------

const NumDbUserColumns = 12

const NumDbUserDataColumns = 11

const DbUserColumnNames = "uid,name,emailaddress,addressid,avatar,role,active,admin,fave,lastupdated,token,secret"

const DbUserDataColumnNames = "name,emailaddress,addressid,avatar,role,active,admin,fave,lastupdated,token,secret"

const DbUserPk = "uid"

//--------------------------------------------------------------------------------

const sqlDbUserTableCreateColumnsSqlite = "\n" +
	" `uid`          integer primary key autoincrement,\n" +
	" `name`         text,\n" +
	" `emailaddress` text,\n" +
	" `addressid`    bigint default null,\n" +
	" `avatar`       text default null,\n" +
	" `role`         text default null,\n" +
	" `active`       boolean,\n" +
	" `admin`        boolean,\n" +
	" `fave`         text,\n" +
	" `lastupdated`  bigint,\n" +
	" `token`        text,\n" +
	" `secret`       text"

const sqlDbUserTableCreateColumnsMysql = "\n" +
	" `uid`          bigint primary key auto_increment,\n" +
	" `name`         varchar(255),\n" +
	" `emailaddress` varchar(255),\n" +
	" `addressid`    bigint default null,\n" +
	" `avatar`       varchar(255) default null,\n" +
	" `role`         varchar(20) default null,\n" +
	" `active`       tinyint(1),\n" +
	" `admin`        tinyint(1),\n" +
	" `fave`         json,\n" +
	" `lastupdated`  bigint,\n" +
	" `token`        varchar(255),\n" +
	" `secret`       varchar(255)"

const sqlDbUserTableCreateColumnsPostgres = `
 "uid"          bigserial primary key,
 "name"         varchar(255),
 "emailaddress" varchar(255),
 "addressid"    bigint default null,
 "avatar"       varchar(255) default null,
 "role"         varchar(20) default null,
 "active"       boolean,
 "admin"        boolean,
 "fave"         json,
 "lastupdated"  bigint,
 "token"        varchar(255),
 "secret"       varchar(255)`

const sqlConstrainDbUserTable = `
 CONSTRAINT DbUserc3 foreign key (addressid) references %saddresses (id) on update restrict on delete restrict
`

//--------------------------------------------------------------------------------

const sqlDbEmailaddressIdxIndexColumns = "emailaddress"

const sqlDbUserLoginIndexColumns = "name"

//--------------------------------------------------------------------------------

// CreateTable creates the table.
func (tbl DbUserTable) CreateTable(ifNotExists bool) (int64, error) {
	return support.Exec(tbl.ctx, tbl, nil, tbl.createTableSql(ifNotExists))
}

func (tbl DbUserTable) createTableSql(ifNotExists bool) string {
	var columns string
	var settings string
	switch tbl.Dialect() {
	case schema.Sqlite:
		columns = sqlDbUserTableCreateColumnsSqlite
		settings = ""
	case schema.Mysql:
		columns = sqlDbUserTableCreateColumnsMysql
		settings = " ENGINE=InnoDB DEFAULT CHARSET=utf8"
	case schema.Postgres:
		columns = sqlDbUserTableCreateColumnsPostgres
		settings = ""
	}
	buf := &bytes.Buffer{}
	buf.WriteString("CREATE TABLE ")
	if ifNotExists {
		buf.WriteString("IF NOT EXISTS ")
	}
	buf.WriteString(tbl.name.String())
	buf.WriteString(" (")
	buf.WriteString(columns)
	for i, c := range tbl.constraints {
		buf.WriteString(",\n ")
		buf.WriteString(c.ConstraintSql(tbl.Dialect(), tbl.name, i+1))
	}
	buf.WriteString("\n)")
	buf.WriteString(settings)
	return buf.String()
}

func (tbl DbUserTable) ternary(flag bool, a, b string) string {
	if flag {
		return a
	}
	return b
}

// DropTable drops the table, destroying all its data.
func (tbl DbUserTable) DropTable(ifExists bool) (int64, error) {
	return support.Exec(tbl.ctx, tbl, nil, tbl.dropTableSql(ifExists))
}

func (tbl DbUserTable) dropTableSql(ifExists bool) string {
	ie := tbl.ternary(ifExists, "IF EXISTS ", "")
	query := fmt.Sprintf("DROP TABLE %s%s", ie, tbl.name)
	return query
}

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

	err = tbl.CreateEmailaddressIdxIndex(ifNotExist)
	if err != nil {
		return err
	}

	err = tbl.CreateUserLoginIndex(ifNotExist)
	if err != nil {
		return err
	}

	return nil
}

// CreateEmailaddressIdxIndex creates the emailaddress_idx index.
func (tbl DbUserTable) CreateEmailaddressIdxIndex(ifNotExist bool) error {
	ine := tbl.ternary(ifNotExist && tbl.Dialect() != schema.Mysql, "IF NOT EXISTS ", "")

	// Mysql does not support 'if not exists' on indexes
	// Workaround: use DropIndex first and ignore an error returned if the index didn't exist.

	if ifNotExist && tbl.Dialect() == schema.Mysql {
		// low-level no-logging Exec
		tbl.Execer().ExecContext(tbl.ctx, tbl.dropDbEmailaddressIdxIndexSql(false))
		ine = ""
	}

	_, err := tbl.Exec(nil, tbl.createDbEmailaddressIdxIndexSql(ine))
	return err
}

func (tbl DbUserTable) createDbEmailaddressIdxIndexSql(ifNotExists string) string {
	indexPrefix := tbl.name.PrefixWithoutDot()
	return fmt.Sprintf("CREATE UNIQUE INDEX %s%semailaddress_idx ON %s (%s)", ifNotExists, indexPrefix,
		tbl.name, sqlDbEmailaddressIdxIndexColumns)
}

// DropEmailaddressIdxIndex drops the emailaddress_idx index.
func (tbl DbUserTable) DropEmailaddressIdxIndex(ifExists bool) error {
	_, err := tbl.Exec(nil, tbl.dropDbEmailaddressIdxIndexSql(ifExists))
	return err
}

func (tbl DbUserTable) dropDbEmailaddressIdxIndexSql(ifExists bool) string {
	// Mysql does not support 'if exists' on indexes
	ie := tbl.ternary(ifExists && tbl.Dialect() != schema.Mysql, "IF EXISTS ", "")
	onTbl := tbl.ternary(tbl.Dialect() == schema.Mysql, fmt.Sprintf(" ON %s", tbl.name), "")
	indexPrefix := tbl.name.PrefixWithoutDot()
	return fmt.Sprintf("DROP INDEX %s%semailaddress_idx%s", ie, indexPrefix, onTbl)
}

// CreateUserLoginIndex creates the user_login index.
func (tbl DbUserTable) CreateUserLoginIndex(ifNotExist bool) error {
	ine := tbl.ternary(ifNotExist && tbl.Dialect() != schema.Mysql, "IF NOT EXISTS ", "")

	// Mysql does not support 'if not exists' on indexes
	// Workaround: use DropIndex first and ignore an error returned if the index didn't exist.

	if ifNotExist && tbl.Dialect() == schema.Mysql {
		// low-level no-logging Exec
		tbl.Execer().ExecContext(tbl.ctx, tbl.dropDbUserLoginIndexSql(false))
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
	ie := tbl.ternary(ifExists && tbl.Dialect() != schema.Mysql, "IF EXISTS ", "")
	onTbl := tbl.ternary(tbl.Dialect() == schema.Mysql, fmt.Sprintf(" ON %s", tbl.name), "")
	indexPrefix := tbl.name.PrefixWithoutDot()
	return fmt.Sprintf("DROP INDEX %s%suser_login%s", ie, indexPrefix, onTbl)
}

// DropIndexes executes queries that drop the indexes on by the User table.
func (tbl DbUserTable) DropIndexes(ifExist bool) (err error) {

	err = tbl.DropEmailaddressIdxIndex(ifExist)
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

// Truncate drops every record from the table, if possible. It might fail if constraints exist that
// prevent some or all rows from being deleted; use the force option to override this.
//
// When 'force' is set true, be aware of the following consequences.
// When using Mysql, foreign keys in other tables can be left dangling.
// When using Postgres, a cascade happens, so all 'adjacent' tables (i.e. linked by foreign keys)
// are also truncated.
func (tbl DbUserTable) Truncate(force bool) (err error) {
	for _, query := range tbl.Dialect().TruncateDDL(tbl.Name().String(), force) {
		_, err = support.Exec(tbl.ctx, tbl, nil, query)
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
	return support.Exec(tbl.ctx, tbl, req, query, args...)
}

//--------------------------------------------------------------------------------

// Query is the low-level request method for this table. The query is logged using whatever logger is
// configured. If an error arises, this too is logged.
//
// If you need a context other than the background, use WithContext before calling Query.
//
// The args are for any placeholder parameters in the query.
//
// The caller must call rows.Close() on the result.
func (tbl DbUserTable) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return support.Query(tbl.ctx, tbl, query, args...)
}

//--------------------------------------------------------------------------------

// QueryOneNullString is a low-level access method for one string. This can be used for function queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl DbUserTable) QueryOneNullString(req require.Requirement, query string, args ...interface{}) (result sql.NullString, err error) {
	err = support.QueryOneNullThing(tbl, req, &result, query, args...)
	return result, err
}

// QueryOneNullInt64 is a low-level access method for one int64. This can be used for 'COUNT(1)' queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl DbUserTable) QueryOneNullInt64(req require.Requirement, query string, args ...interface{}) (result sql.NullInt64, err error) {
	err = support.QueryOneNullThing(tbl, req, &result, query, args...)
	return result, err
}

// QueryOneNullFloat64 is a low-level access method for one float64. This can be used for 'AVG(...)' queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl DbUserTable) QueryOneNullFloat64(req require.Requirement, query string, args ...interface{}) (result sql.NullFloat64, err error) {
	err = support.QueryOneNullThing(tbl, req, &result, query, args...)
	return result, err
}

func scanDbUsers(rows *sql.Rows, firstOnly bool) (vv []*User, n int64, err error) {
	for rows.Next() {
		n++

		var v0 int64
		var v1 string
		var v2 string
		var v3 sql.NullInt64
		var v4 sql.NullString
		var v5 sql.NullString
		var v6 bool
		var v7 bool
		var v8 []byte
		var v9 int64
		var v10 string
		var v11 string

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
			&v11,
		)
		if err != nil {
			return vv, n, err
		}

		v := &User{}
		v.Uid = v0
		v.Name = v1
		v.EmailAddress = v2
		if v3.Valid {
			a := v3.Int64
			v.AddressId = &a
		}
		if v4.Valid {
			a := v4.String
			v.Avatar = &a
		}
		if v5.Valid {
			v.Role = new(Role)
			err = v.Role.Scan(v5.String)
			if err != nil {
				return nil, n, err
			}
		}
		v.Active = v6
		v.Admin = v7
		err = json.Unmarshal(v8, &v.Fave)
		if err != nil {
			return nil, n, err
		}
		v.LastUpdated = v9
		v.token = v10
		v.secret = v11

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

var allDbUserQuotedColumnNames = []string{
	schema.Sqlite.SplitAndQuote(DbUserColumnNames),
	schema.Mysql.SplitAndQuote(DbUserColumnNames),
	schema.Postgres.SplitAndQuote(DbUserColumnNames),
}

//--------------------------------------------------------------------------------

// GetUsersByUid gets records from the table according to a list of primary keys.
// Although the list of ids can be arbitrarily long, there are practical limits;
// note that Oracle DB has a limit of 1000.
//
// It places a requirement, which may be nil, on the size of the expected results: in particular, require.All
// controls whether an error is generated not all the ids produce a result.
func (tbl DbUserTable) GetUsersByUid(req require.Requirement, id ...int64) (list []*User, err error) {
	if len(id) > 0 {
		if req == require.All {
			req = require.Exactly(len(id))
		}
		args := make([]interface{}, len(id))

		for i, v := range id {
			args[i] = v
		}

		list, err = tbl.getUsers(req, "uid", args...)
	}

	return list, err
}

// GetUserByUid gets the record with a given primary key value.
// If not found, *User will be nil.
func (tbl DbUserTable) GetUserByUid(req require.Requirement, id int64) (*User, error) {
	return tbl.getUser(req, "uid", id)
}

// GetUserByEmailAddress gets the record with a given emailaddress value.
// If not found, *User will be nil.
func (tbl DbUserTable) GetUserByEmailAddress(req require.Requirement, value string) (*User, error) {
	return tbl.getUser(req, "emailaddress", value)
}

// GetUserByName gets the record with a given name value.
// If not found, *User will be nil.
func (tbl DbUserTable) GetUserByName(req require.Requirement, value string) (*User, error) {
	return tbl.getUser(req, "name", value)
}

func (tbl DbUserTable) getUser(req require.Requirement, column string, arg interface{}) (*User, error) {
	dialect := tbl.Dialect()
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s=%s",
		allDbUserQuotedColumnNames[dialect.Index()], tbl.name, dialect.Quote(column), dialect.Placeholder(column, 1))
	v, err := tbl.doQueryOne(req, query, arg)
	return v, err
}

func (tbl DbUserTable) getUsers(req require.Requirement, column string, args ...interface{}) (list []*User, err error) {
	if len(args) > 0 {
		if req == require.All {
			req = require.Exactly(len(args))
		}
		dialect := tbl.Dialect()
		pl := dialect.Placeholders(len(args))
		query := fmt.Sprintf("SELECT %s FROM %s WHERE %s IN (%s)",
			allDbUserQuotedColumnNames[dialect.Index()], tbl.name, dialect.Quote(column), pl)
		list, err = tbl.doQuery(req, false, query, args...)
	}

	return list, err
}

func (tbl DbUserTable) doQueryOne(req require.Requirement, query string, args ...interface{}) (*User, error) {
	list, err := tbl.doQuery(req, true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

func (tbl DbUserTable) doQuery(req require.Requirement, firstOnly bool, query string, args ...interface{}) ([]*User, error) {
	rows, err := tbl.Query(query, args...)
	if err != nil {
		return nil, err
	}

	vv, n, err := scanDbUsers(rows, firstOnly)
	return vv, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(err, req, n))
}

// Fetch fetches a list of User based on a supplied query. This is mostly used for join queries that map its
// result columns to the fields of User. Other queries might be better handled by GetXxx or Select methods.
func (tbl DbUserTable) Fetch(req require.Requirement, query string, args ...interface{}) ([]*User, error) {
	return tbl.doQuery(req, false, query, args...)
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
	query := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT 1",
		allDbUserQuotedColumnNames[tbl.Dialect().Index()], tbl.name, where, orderBy)
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
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
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
	query := fmt.Sprintf("SELECT %s FROM %s %s %s",
		allDbUserQuotedColumnNames[tbl.Dialect().Index()], tbl.name, where, orderBy)
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
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
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
	whs, args := where.BuildExpression(wh, tbl.Dialect())
	return tbl.CountWhere(whs, args...)
}

//--------------------------------------------------------------------------------

// SliceUid gets the Uid column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DbUserTable) SliceUid(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return tbl.getint64list(req, "uid", wh, qc)
}

// SliceName gets the Name column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DbUserTable) SliceName(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.getstringlist(req, "name", wh, qc)
}

// SliceEmailAddress gets the EmailAddress column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DbUserTable) SliceEmailaddress(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.getstringlist(req, "emailaddress", wh, qc)
}

// SliceAddressId gets the AddressId column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DbUserTable) SliceAddressid(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return tbl.getint64Ptrlist(req, "addressid", wh, qc)
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
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", dialect.Quote(sqlname), tbl.name, whs, orderBy)
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
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", dialect.Quote(sqlname), tbl.name, whs, orderBy)
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
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", dialect.Quote(sqlname), tbl.name, whs, orderBy)
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

func (tbl DbUserTable) getint64Ptrlist(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", dialect.Quote(sqlname), tbl.name, whs, orderBy)
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
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", dialect.Quote(sqlname), tbl.name, whs, orderBy)
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
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", dialect.Quote(sqlname), tbl.name, whs, orderBy)
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

func constructDbUserInsert(w io.Writer, v *User, dialect schema.Dialect, withPk bool) (s []interface{}, err error) {
	s = make([]interface{}, 0, 12)

	comma := ""
	io.WriteString(w, " (")

	if withPk {
		dialect.QuoteW(w, "uid")
		comma = ","
		s = append(s, v.Uid)
	}

	io.WriteString(w, comma)

	dialect.QuoteW(w, "name")
	s = append(s, v.Name)
	comma = ","
	io.WriteString(w, comma)

	dialect.QuoteW(w, "emailaddress")
	s = append(s, v.EmailAddress)
	if v.AddressId != nil {
		io.WriteString(w, comma)

		dialect.QuoteW(w, "addressid")
		s = append(s, v.AddressId)
	}
	if v.Avatar != nil {
		io.WriteString(w, comma)

		dialect.QuoteW(w, "avatar")
		s = append(s, v.Avatar)
	}
	if v.Role != nil {
		io.WriteString(w, comma)

		dialect.QuoteW(w, "role")
		s = append(s, v.Role)
	}
	io.WriteString(w, comma)

	dialect.QuoteW(w, "active")
	s = append(s, v.Active)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "admin")
	s = append(s, v.Admin)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "fave")
	x, err := json.Marshal(&v.Fave)
	if err != nil {
		return nil, err
	}
	s = append(s, x)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "lastupdated")
	s = append(s, v.LastUpdated)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "token")
	s = append(s, v.token)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "secret")
	s = append(s, v.secret)
	io.WriteString(w, ")")
	return s, nil
}

func constructDbUserUpdate(w io.Writer, v *User, dialect schema.Dialect) (s []interface{}, err error) {
	j := 1
	s = make([]interface{}, 0, 11)

	comma := ""

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "name", j)
	s = append(s, v.Name)
	comma = ", "
	j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "emailaddress", j)
	s = append(s, v.EmailAddress)
	j++

	io.WriteString(w, comma)
	if v.AddressId != nil {
		dialect.QuoteWithPlaceholder(w, "addressid", j)
		s = append(s, v.AddressId)
		j++
	} else {
		dialect.QuoteW(w, "addressid")
		io.WriteString(w, "=NULL")
	}

	io.WriteString(w, comma)
	if v.Avatar != nil {
		dialect.QuoteWithPlaceholder(w, "avatar", j)
		s = append(s, v.Avatar)
		j++
	} else {
		dialect.QuoteW(w, "avatar")
		io.WriteString(w, "=NULL")
	}

	io.WriteString(w, comma)
	if v.Role != nil {
		dialect.QuoteWithPlaceholder(w, "role", j)
		s = append(s, v.Role)
		j++
	} else {
		dialect.QuoteW(w, "role")
		io.WriteString(w, "=NULL")
	}

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "active", j)
	s = append(s, v.Active)
	j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "admin", j)
	s = append(s, v.Admin)
	j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "fave", j)
	j++
	x, err := json.Marshal(&v.Fave)
	if err != nil {
		return nil, err
	}
	s = append(s, x)

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "lastupdated", j)
	s = append(s, v.LastUpdated)
	j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "token", j)
	s = append(s, v.token)
	j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "secret", j)
	s = append(s, v.secret)
	j++

	return s, nil
}

//--------------------------------------------------------------------------------

// Insert adds new records for the Users.
// The Users have their primary key fields set to the new record identifiers.
// The User.PreInsert() method will be called, if it exists.
func (tbl DbUserTable) Insert(req require.Requirement, vv ...*User) error {
	if req == require.All {
		req = require.Exactly(len(vv))
	}

	var count int64
	//columns := allXExampleQuotedInserts[tbl.Dialect().Index()]
	//query := fmt.Sprintf("INSERT INTO %s %s", tbl.name, columns)
	//st, err := tbl.db.PrepareContext(tbl.ctx, query)
	//if err != nil {
	//	return err
	//}
	//defer st.Close()

	insertHasReturningPhrase := tbl.Dialect().InsertHasReturningPhrase()
	returning := ""
	if tbl.Dialect().InsertHasReturningPhrase() {
		returning = fmt.Sprintf(" returning %q", DbUserPk)
	}

	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPreInsert); ok {
			err := hook.PreInsert()
			if err != nil {
				return tbl.logError(err)
			}
		}

		b := &bytes.Buffer{}
		io.WriteString(b, "INSERT INTO ")
		io.WriteString(b, tbl.name.String())

		fields, err := constructDbUserInsert(b, v, tbl.Dialect(), false)
		if err != nil {
			return tbl.logError(err)
		}

		io.WriteString(b, " VALUES (")
		io.WriteString(b, tbl.Dialect().Placeholders(len(fields)))
		io.WriteString(b, ")")
		io.WriteString(b, returning)

		query := b.String()
		tbl.logQuery(query, fields...)

		var n int64 = 1
		if insertHasReturningPhrase {
			row := tbl.db.QueryRowContext(tbl.ctx, query, fields...)
			err = row.Scan(&v.Uid)

		} else {
			res, e2 := tbl.db.ExecContext(tbl.ctx, query, fields...)
			if e2 != nil {
				return tbl.logError(e2)
			}

			v.Uid, err = res.LastInsertId()
			if e2 != nil {
				return tbl.logError(e2)
			}

			n, err = res.RowsAffected()
		}

		if err != nil {
			return tbl.logError(err)
		}
		count += n
	}

	return tbl.logIfError(require.ErrorIfExecNotSatisfiedBy(req, count))
}

// UpdateFields updates one or more columns, given a 'where' clause.
//
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl DbUserTable) UpdateFields(req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error) {
	return support.UpdateFields(tbl.ctx, tbl, req, wh, fields...)
}

//--------------------------------------------------------------------------------

var allDbUserQuotedUpdates = []string{
	// Sqlite
	"`name`=?,`emailaddress`=?,`addressid`=?,`avatar`=?,`role`=?,`active`=?,`admin`=?,`fave`=?,`lastupdated`=?,`token`=?,`secret`=? WHERE `uid`=?",
	// Mysql
	"`name`=?,`emailaddress`=?,`addressid`=?,`avatar`=?,`role`=?,`active`=?,`admin`=?,`fave`=?,`lastupdated`=?,`token`=?,`secret`=? WHERE `uid`=?",
	// Postgres
	`"name"=$2,"emailaddress"=$3,"addressid"=$4,"avatar"=$5,"role"=$6,"active"=$7,"admin"=$8,"fave"=$9,"lastupdated"=$10,"token"=$11,"secret"=$12 WHERE "uid"=$1`,
}

//--------------------------------------------------------------------------------

// Update updates records, matching them by primary key. It returns the number of rows affected.
// The User.PreUpdate(Execer) method will be called, if it exists.
func (tbl DbUserTable) Update(req require.Requirement, vv ...*User) (int64, error) {
	if req == require.All {
		req = require.Exactly(len(vv))
	}

	var count int64
	dialect := tbl.Dialect()
	//columns := allDbUserQuotedUpdates[dialect.Index()]
	//query := fmt.Sprintf("UPDATE %s SET %s", tbl.name, columns)

	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPreUpdate); ok {
			err := hook.PreUpdate()
			if err != nil {
				return count, tbl.logError(err)
			}
		}

		b := &bytes.Buffer{}
		io.WriteString(b, "UPDATE ")
		io.WriteString(b, tbl.name.String())
		io.WriteString(b, " SET ")

		args, err := constructDbUserUpdate(b, v, dialect)
		k := len(args) + 1
		args = append(args, v.Uid)
		if err != nil {
			return count, tbl.logError(err)
		}

		io.WriteString(b, " WHERE ")
		dialect.QuoteWithPlaceholder(b, "uid", k)

		query := b.String()
		n, err := tbl.Exec(nil, query, args...)
		if err != nil {
			return count, err
		}
		count += n
	}

	return count, tbl.logIfError(require.ErrorIfExecNotSatisfiedBy(req, count))
}

//--------------------------------------------------------------------------------

// DeleteUsers deletes rows from the table, given some primary keys.
// The list of ids can be arbitrarily long.
func (tbl DbUserTable) DeleteUsers(req require.Requirement, id ...int64) (int64, error) {
	const batch = 1000 // limited by Oracle DB
	const qt = "DELETE FROM %s WHERE %s IN (%s)"

	if req == require.All {
		req = require.Exactly(len(id))
	}

	var count, n int64
	var err error
	var max = batch
	if len(id) < batch {
		max = len(id)
	}
	dialect := tbl.Dialect()
	col := dialect.Quote("uid")
	args := make([]interface{}, max)

	if len(id) > batch {
		pl := dialect.Placeholders(batch)
		query := fmt.Sprintf(qt, tbl.name, col, pl)

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
		pl := dialect.Placeholders(len(id))
		query := fmt.Sprintf(qt, tbl.name, col, pl)

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
	whs, args := where.BuildExpression(wh, tbl.Dialect())
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

// SetName sets the Name field and returns the modified User.
func (v *User) SetName(x string) *User {
	v.Name = x
	return v
}

// SetEmailAddress sets the EmailAddress field and returns the modified User.
func (v *User) SetEmailAddress(x string) *User {
	v.EmailAddress = x
	return v
}

// SetAddressId sets the AddressId field and returns the modified User.
func (v *User) SetAddressId(x int64) *User {
	v.AddressId = &x
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
