// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.
// sqlapi v0.47.0; sqlgen v0.66.0

package demo

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/rickb777/sqlapi"
	"github.com/rickb777/sqlapi/constraint"
	"github.com/rickb777/sqlapi/dialect"
	"github.com/rickb777/sqlapi/require"
	"github.com/rickb777/sqlapi/support"
	"github.com/rickb777/where"
	"github.com/rickb777/where/quote"
	"strings"
)

// HookTabler lists table methods provided by HookTable.
type HookTabler interface {
	sqlapi.Table

	// Constraints returns the table's constraints.
	Constraints() constraint.Constraints

	// WithConstraint returns a modified HookTabler with added data consistency constraints.
	WithConstraint(cc ...constraint.Constraint) HookTabler

	// WithPrefix returns a modified HookTabler with a given table name prefix.
	WithPrefix(pfx string) HookTabler

	// CreateTable creates the table.
	CreateTable(ctx context.Context, ifNotExists bool) (int64, error)

	// DropTable drops the table, destroying all its data.
	DropTable(ctx context.Context, ifExists bool) (int64, error)

	// Truncate drops every record from the table, if possible.
	Truncate(ctx context.Context, force bool) (err error)
}

//-------------------------------------------------------------------------------------------------

// HookQueryer lists query methods provided by HookTable.
type HookQueryer interface {
	sqlapi.Table

	// Using returns a modified HookQueryer using the Execer supplied,
	// which will typically be a transaction (i.e. SqlTx).
	Using(tx sqlapi.Execer) HookQueryer

	// Transact runs the function provided within a transaction. The transction is committed
	// unless an error occurs.
	Transact(ctx context.Context, txOptions *sql.TxOptions, fn func(HookQueryer) error) error

	// Exec executes a query without returning any rows.
	Exec(ctx context.Context, req require.Requirement, query string, args ...interface{}) (int64, error)

	// Query is the low-level request method for this table using an SQL query that must return all the columns
	// necessary for Hook values.
	Query(ctx context.Context, req require.Requirement, query string, args ...interface{}) (HookList, error)

	// QueryOneNullString is a low-level access method for one string, returning the first match.
	QueryOneNullString(ctx context.Context, req require.Requirement, query string, args ...interface{}) (result sql.NullString, err error)

	// QueryOneNullInt64 is a low-level access method for one int64, returning the first match.
	QueryOneNullInt64(ctx context.Context, req require.Requirement, query string, args ...interface{}) (result sql.NullInt64, err error)

	// QueryOneNullFloat64 is a low-level access method for one float64, returning the first match.
	QueryOneNullFloat64(ctx context.Context, req require.Requirement, query string, args ...interface{}) (result sql.NullFloat64, err error)

	// GetHookById gets the record with a given primary key value.
	GetHookById(ctx context.Context, req require.Requirement, id uint64) (*Hook, error)

	// GetHooksById gets records from the table according to a list of primary keys.
	GetHooksById(ctx context.Context, req require.Requirement, qc where.QueryConstraint, id ...uint64) (list HookList, err error)

	// SelectOneWhere allows a single Hook to be obtained from the table that matches a 'where' clause.
	SelectOneWhere(ctx context.Context, req require.Requirement, where, orderBy string, args ...interface{}) (*Hook, error)

	// SelectOne allows a single Hook to be obtained from the table that matches a 'where' clause.
	SelectOne(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) (*Hook, error)

	// SelectWhere allows Hooks to be obtained from the table that match a 'where' clause.
	SelectWhere(ctx context.Context, req require.Requirement, where, orderBy string, args ...interface{}) (HookList, error)

	// Select allows Hooks to be obtained from the table that match a 'where' clause.
	Select(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) (HookList, error)

	// CountWhere counts Hooks in the table that match a 'where' clause.
	CountWhere(ctx context.Context, where string, args ...interface{}) (count int64, err error)

	// Count counts the Hooks in the table that match a 'where' clause.
	Count(ctx context.Context, wh where.Expression) (count int64, err error)

	// SliceId gets the id column for all rows that match the 'where' condition.
	SliceId(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]uint64, error)

	// SliceSha gets the sha column for all rows that match the 'where' condition.
	SliceSha(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error)

	// SliceAfter gets the after column for all rows that match the 'where' condition.
	SliceAfter(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error)

	// SliceBefore gets the before column for all rows that match the 'where' condition.
	SliceBefore(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error)

	// SliceCommitId gets the commit_id column for all rows that match the 'where' condition.
	SliceCommitId(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error)

	// SliceMessage gets the message column for all rows that match the 'where' condition.
	SliceMessage(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error)

	// SliceTimestamp gets the timestamp column for all rows that match the 'where' condition.
	SliceTimestamp(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error)

	// SliceHeadCommitAuthorName gets the head_commit_author_name column for all rows that match the 'where' condition.
	SliceHeadCommitAuthorName(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error)

	// SliceHeadCommitAuthorUsername gets the head_commit_author_username column for all rows that match the 'where' condition.
	SliceHeadCommitAuthorUsername(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error)

	// SliceHeadCommitCommitterName gets the head_commit_committer_name column for all rows that match the 'where' condition.
	SliceHeadCommitCommitterName(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error)

	// SliceHeadCommitCommitterUsername gets the head_commit_committer_username column for all rows that match the 'where' condition.
	SliceHeadCommitCommitterUsername(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error)

	// SliceCategory gets the category column for all rows that match the 'where' condition.
	SliceCategory(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]Category, error)

	// SliceHeadCommitAuthorEmail gets the head_commit_author_email column for all rows that match the 'where' condition.
	SliceHeadCommitAuthorEmail(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]Email, error)

	// SliceHeadCommitCommitterEmail gets the head_commit_committer_email column for all rows that match the 'where' condition.
	SliceHeadCommitCommitterEmail(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]Email, error)

	// Insert adds new records for the Hooks, setting the primary key field for each one.
	Insert(ctx context.Context, req require.Requirement, vv ...*Hook) error

	// UpdateById updates one or more columns, given a id value.
	UpdateById(ctx context.Context, req require.Requirement, id uint64, fields ...sql.NamedArg) (int64, error)

	// UpdateBySha updates one or more columns, given a sha value.
	UpdateBySha(ctx context.Context, req require.Requirement, sha string, fields ...sql.NamedArg) (int64, error)

	// UpdateByAfter updates one or more columns, given a after value.
	UpdateByAfter(ctx context.Context, req require.Requirement, after string, fields ...sql.NamedArg) (int64, error)

	// UpdateByBefore updates one or more columns, given a before value.
	UpdateByBefore(ctx context.Context, req require.Requirement, before string, fields ...sql.NamedArg) (int64, error)

	// UpdateByCategory updates one or more columns, given a category value.
	UpdateByCategory(ctx context.Context, req require.Requirement, category Category, fields ...sql.NamedArg) (int64, error)

	// UpdateByCommitId updates one or more columns, given a commit_id value.
	UpdateByCommitId(ctx context.Context, req require.Requirement, commit_id string, fields ...sql.NamedArg) (int64, error)

	// UpdateByMessage updates one or more columns, given a message value.
	UpdateByMessage(ctx context.Context, req require.Requirement, message string, fields ...sql.NamedArg) (int64, error)

	// UpdateByTimestamp updates one or more columns, given a timestamp value.
	UpdateByTimestamp(ctx context.Context, req require.Requirement, timestamp string, fields ...sql.NamedArg) (int64, error)

	// UpdateByHeadCommitAuthorName updates one or more columns, given a head_commit_author_name value.
	UpdateByHeadCommitAuthorName(ctx context.Context, req require.Requirement, head_commit_author_name string, fields ...sql.NamedArg) (int64, error)

	// UpdateByHeadCommitAuthorEmail updates one or more columns, given a head_commit_author_email value.
	UpdateByHeadCommitAuthorEmail(ctx context.Context, req require.Requirement, head_commit_author_email Email, fields ...sql.NamedArg) (int64, error)

	// UpdateByHeadCommitAuthorUsername updates one or more columns, given a head_commit_author_username value.
	UpdateByHeadCommitAuthorUsername(ctx context.Context, req require.Requirement, head_commit_author_username string, fields ...sql.NamedArg) (int64, error)

	// UpdateByHeadCommitCommitterName updates one or more columns, given a head_commit_committer_name value.
	UpdateByHeadCommitCommitterName(ctx context.Context, req require.Requirement, head_commit_committer_name string, fields ...sql.NamedArg) (int64, error)

	// UpdateByHeadCommitCommitterEmail updates one or more columns, given a head_commit_committer_email value.
	UpdateByHeadCommitCommitterEmail(ctx context.Context, req require.Requirement, head_commit_committer_email Email, fields ...sql.NamedArg) (int64, error)

	// UpdateByHeadCommitCommitterUsername updates one or more columns, given a head_commit_committer_username value.
	UpdateByHeadCommitCommitterUsername(ctx context.Context, req require.Requirement, head_commit_committer_username string, fields ...sql.NamedArg) (int64, error)

	// UpdateFields updates one or more columns, given a 'where' clause.
	UpdateFields(ctx context.Context, req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error)

	// Update updates records, matching them by primary key.
	Update(ctx context.Context, req require.Requirement, vv ...*Hook) (int64, error)

	// Upsert inserts or updates a record, matching it using the expression supplied.
	// This expression is used to search for an existing record based on some specified
	// key column(s). It must match either zero or one existing record. If it matches
	// none, a new record is inserted; otherwise the matching record is updated. An
	// error results if these conditions are not met.
	Upsert(ctx context.Context, v *Hook, wh where.Expression) error

	// DeleteById deletes rows from the table, given some id values.
	// The list of ids can be arbitrarily long.
	DeleteById(ctx context.Context, req require.Requirement, id ...uint64) (int64, error)

	// DeleteBySha deletes rows from the table, given some sha values.
	// The list of ids can be arbitrarily long.
	DeleteBySha(ctx context.Context, req require.Requirement, sha ...string) (int64, error)

	// DeleteByAfter deletes rows from the table, given some after values.
	// The list of ids can be arbitrarily long.
	DeleteByAfter(ctx context.Context, req require.Requirement, after ...string) (int64, error)

	// DeleteByBefore deletes rows from the table, given some before values.
	// The list of ids can be arbitrarily long.
	DeleteByBefore(ctx context.Context, req require.Requirement, before ...string) (int64, error)

	// DeleteByCategory deletes rows from the table, given some category values.
	// The list of ids can be arbitrarily long.
	DeleteByCategory(ctx context.Context, req require.Requirement, category ...Category) (int64, error)

	// DeleteByCommitId deletes rows from the table, given some commit_id values.
	// The list of ids can be arbitrarily long.
	DeleteByCommitId(ctx context.Context, req require.Requirement, commit_id ...string) (int64, error)

	// DeleteByMessage deletes rows from the table, given some message values.
	// The list of ids can be arbitrarily long.
	DeleteByMessage(ctx context.Context, req require.Requirement, message ...string) (int64, error)

	// DeleteByTimestamp deletes rows from the table, given some timestamp values.
	// The list of ids can be arbitrarily long.
	DeleteByTimestamp(ctx context.Context, req require.Requirement, timestamp ...string) (int64, error)

	// DeleteByHeadCommitAuthorName deletes rows from the table, given some head_commit_author_name values.
	// The list of ids can be arbitrarily long.
	DeleteByHeadCommitAuthorName(ctx context.Context, req require.Requirement, head_commit_author_name ...string) (int64, error)

	// DeleteByHeadCommitAuthorEmail deletes rows from the table, given some head_commit_author_email values.
	// The list of ids can be arbitrarily long.
	DeleteByHeadCommitAuthorEmail(ctx context.Context, req require.Requirement, head_commit_author_email ...Email) (int64, error)

	// DeleteByHeadCommitAuthorUsername deletes rows from the table, given some head_commit_author_username values.
	// The list of ids can be arbitrarily long.
	DeleteByHeadCommitAuthorUsername(ctx context.Context, req require.Requirement, head_commit_author_username ...string) (int64, error)

	// DeleteByHeadCommitCommitterName deletes rows from the table, given some head_commit_committer_name values.
	// The list of ids can be arbitrarily long.
	DeleteByHeadCommitCommitterName(ctx context.Context, req require.Requirement, head_commit_committer_name ...string) (int64, error)

	// DeleteByHeadCommitCommitterEmail deletes rows from the table, given some head_commit_committer_email values.
	// The list of ids can be arbitrarily long.
	DeleteByHeadCommitCommitterEmail(ctx context.Context, req require.Requirement, head_commit_committer_email ...Email) (int64, error)

	// DeleteByHeadCommitCommitterUsername deletes rows from the table, given some head_commit_committer_username values.
	// The list of ids can be arbitrarily long.
	DeleteByHeadCommitCommitterUsername(ctx context.Context, req require.Requirement, head_commit_committer_username ...string) (int64, error)

	// Delete deletes one or more rows from the table, given a 'where' clause.
	// Use a nil value for the 'wh' argument if it is not needed (very risky!).
	Delete(ctx context.Context, req require.Requirement, wh where.Expression) (int64, error)
}

//-------------------------------------------------------------------------------------------------

// HookTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type HookTable struct {
	name        sqlapi.TableName
	database    sqlapi.Database
	db          sqlapi.Execer
	constraints constraint.Constraints
	pk          string
}

// Type conformance checks
var _ sqlapi.TableCreator = &HookTable{}

// NewHookTable returns a new table instance.
// If a blank table name is supplied, the default name "hooks" will be used instead.
// The request context is initialised with the background.
func NewHookTable(name string, d sqlapi.Database) HookTable {
	if name == "" {
		name = "hooks"
	}
	var constraints constraint.Constraints
	return HookTable{
		name:        sqlapi.TableName{Prefix: "", Name: name},
		database:    d,
		db:          d.DB(),
		constraints: constraints,
		pk:          "id",
	}
}

// CopyTableAsHookTable copies a table instance, retaining the name etc but
// providing methods appropriate for 'Hook'.It doesn't copy the constraints of the original table.
//
// It serves to provide methods appropriate for 'Hook'. This is most useful when this is used to represent a
// join result. In such cases, there won't be any need for DDL methods, nor Exec, Insert, Update or Delete.
func CopyTableAsHookTable(origin sqlapi.Table) HookTable {
	return HookTable{
		name:        origin.Name(),
		database:    origin.Database(),
		db:          origin.Execer(),
		constraints: nil,
		pk:          "id",
	}
}

// SetPkColumn sets the name of the primary key column. It defaults to "id".
// The result is a modified copy of the table; the original is unchanged.
//func (tbl HookTable) SetPkColumn(pk string) HookTabler {
//	tbl.pk = pk
//	return tbl
//}

// WithPrefix sets the table name prefix for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl HookTable) WithPrefix(pfx string) HookTabler {
	tbl.name.Prefix = pfx
	return tbl
}

// Database gets the shared database information.
func (tbl HookTable) Database() sqlapi.Database {
	return tbl.database
}

// Logger gets the trace logger.
func (tbl HookTable) Logger() sqlapi.Logger {
	return tbl.database.Logger()
}

// WithConstraint returns a modified HookTabler with added data consistency constraints.
func (tbl HookTable) WithConstraint(cc ...constraint.Constraint) HookTabler {
	tbl.constraints = append(tbl.constraints, cc...)
	return tbl
}

// Constraints returns the table's constraints.
func (tbl HookTable) Constraints() constraint.Constraints {
	return tbl.constraints
}

// Dialect gets the database dialect.
func (tbl HookTable) Dialect() dialect.Dialect {
	return tbl.database.Dialect()
}

// Name gets the table name.
func (tbl HookTable) Name() sqlapi.TableName {
	return tbl.name
}

// PkColumn gets the column name used as a primary key.
func (tbl HookTable) PkColumn() string {
	return tbl.pk
}

// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl HookTable) DB() sqlapi.SqlDB {
	return tbl.db.(sqlapi.SqlDB)
}

// Execer gets the wrapped database or transaction handle.
func (tbl HookTable) Execer() sqlapi.Execer {
	return tbl.db
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl HookTable) Tx() sqlapi.SqlTx {
	return tbl.db.(sqlapi.SqlTx)
}

// IsTx tests whether this is within a transaction.
func (tbl HookTable) IsTx() bool {
	return tbl.db.IsTx()
}

// Using returns a modified HookTabler using the the Execer supplied,
// which will typically be a transaction (i.e. SqlTx). This is needed when making multiple
// queries across several tables within a single transaction.
//
// The result is a modified copy of the table; the original is unchanged.
func (tbl HookTable) Using(tx sqlapi.Execer) HookQueryer {
	tbl.db = tx
	return tbl
}

// Transact runs the function provided within a transaction. If the function completes without error,
// the transaction is committed. If there is an error or a panic, the transaction is rolled back.
//
// The options can be nil, in which case the default behaviour is that of the underlying connection.
//
// Nested transactions (i.e. within 'fn') are permitted: they execute within the outermost transaction.
// Therefore they do not commit until the outermost transaction commits.
func (tbl HookTable) Transact(ctx context.Context, txOptions *sql.TxOptions, fn func(HookQueryer) error) error {
	var err error
	if tbl.IsTx() {
		err = fn(tbl) // nested transactions are inlined
	} else {
		err = tbl.DB().Transact(ctx, txOptions, func(tx sqlapi.SqlTx) error {
			return fn(tbl.Using(tx))
		})
	}
	return tbl.Logger().LogIfError(err)
}

func (tbl HookTable) quotedName() string {
	return tbl.Dialect().Quoter().Quote(tbl.name.String())
}

func (tbl HookTable) quotedNameW(w dialect.StringWriter) {
	tbl.Dialect().Quoter().QuoteW(w, tbl.name.String())
}

//-------------------------------------------------------------------------------------------------

// NumHookTableColumns is the total number of columns in HookTable.
const NumHookTableColumns = 17

// NumHookTableDataColumns is the number of columns in HookTable not including the auto-increment key.
const NumHookTableDataColumns = 16

// HookTableColumnNames is the list of columns in HookTable.
const HookTableColumnNames = "id,sha,after,before,category,created,deleted,forced,commit_id,message,timestamp,head_commit_author_name,head_commit_author_email,head_commit_author_username,head_commit_committer_name,head_commit_committer_email,head_commit_committer_username"

// HookTableDataColumnNames is the list of data columns in HookTable.
const HookTableDataColumnNames = "sha,after,before,category,created,deleted,forced,commit_id,message,timestamp,head_commit_author_name,head_commit_author_email,head_commit_author_username,head_commit_committer_name,head_commit_committer_email,head_commit_committer_username"

var listOfHookTableColumnNames = strings.Split(HookTableColumnNames, ",")

//-------------------------------------------------------------------------------------------------

var sqlHookTableCreateColumnsSqlite = []string{
	"integer not null primary key autoincrement",
	"text not null",
	"text not null",
	"text not null",
	"tinyint unsigned not null",
	"boolean not null",
	"boolean not null",
	"boolean not null",
	"text not null",
	"text not null",
	"text not null",
	"text not null",
	"text not null",
	"text not null",
	"text not null",
	"text not null",
	"text not null",
}

var sqlHookTableCreateColumnsMysql = []string{
	"bigint unsigned not null primary key auto_increment",
	"text not null",
	"varchar(20) not null",
	"varchar(20) not null",
	"tinyint unsigned not null",
	"boolean not null",
	"boolean not null",
	"boolean not null",
	"text not null",
	"text not null",
	"text not null",
	"text not null",
	"text not null",
	"text not null",
	"text not null",
	"text not null",
	"text not null",
}

var sqlHookTableCreateColumnsPostgres = []string{
	"bigserial not null primary key",
	"text not null",
	"text not null",
	"text not null",
	"smallint not null",
	"boolean not null",
	"boolean not null",
	"boolean not null",
	"text not null",
	"text not null",
	"text not null",
	"text not null",
	"text not null",
	"text not null",
	"text not null",
	"text not null",
	"text not null",
}

var sqlHookTableCreateColumnsPgx = []string{
	"bigserial not null primary key",
	"text not null",
	"text not null",
	"text not null",
	"smallint not null",
	"boolean not null",
	"boolean not null",
	"boolean not null",
	"text not null",
	"text not null",
	"text not null",
	"text not null",
	"text not null",
	"text not null",
	"text not null",
	"text not null",
	"text not null",
}

//-------------------------------------------------------------------------------------------------

// CreateTable creates the table.
func (tbl HookTable) CreateTable(ctx context.Context, ifNotExists bool) (int64, error) {
	return support.Exec(ctx, tbl, nil, createHookTableSql(tbl, ifNotExists))
}

func createHookTableSql(tbl HookTabler, ifNotExists bool) string {
	buf := &bytes.Buffer{}
	buf.WriteString("CREATE TABLE ")
	if ifNotExists {
		buf.WriteString("IF NOT EXISTS ")
	}
	q := tbl.Dialect().Quoter()
	q.QuoteW(buf, tbl.Name().String())
	buf.WriteString(" (\n ")

	var columns []string
	switch tbl.Dialect().Index() {
	case dialect.SqliteIndex:
		columns = sqlHookTableCreateColumnsSqlite
	case dialect.MysqlIndex:
		columns = sqlHookTableCreateColumnsMysql
	case dialect.PostgresIndex:
		columns = sqlHookTableCreateColumnsPostgres
	case dialect.PgxIndex:
		columns = sqlHookTableCreateColumnsPgx
	}

	comma := ""
	for i, n := range listOfHookTableColumnNames {
		buf.WriteString(comma)
		q.QuoteW(buf, n)
		buf.WriteString(" ")
		buf.WriteString(columns[i])
		comma = ",\n "
	}

	for i, c := range tbl.(HookTable).Constraints() {
		buf.WriteString(",\n ")
		buf.WriteString(c.ConstraintSql(tbl.Dialect().Quoter(), tbl.Name(), i+1))
	}

	buf.WriteString("\n)")
	buf.WriteString(tbl.Dialect().CreateTableSettings())
	return buf.String()
}

func ternaryHookTable(flag bool, a, b string) string {
	if flag {
		return a
	}
	return b
}

// DropTable drops the table, destroying all its data.
func (tbl HookTable) DropTable(ctx context.Context, ifExists bool) (int64, error) {
	return support.Exec(ctx, tbl, nil, dropHookTableSql(tbl, ifExists))
}

func dropHookTableSql(tbl HookTabler, ifExists bool) string {
	ie := ternaryHookTable(ifExists, "IF EXISTS ", "")
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("DROP TABLE %s%s", ie, quotedName)
	return query
}

//-------------------------------------------------------------------------------------------------

// Truncate drops every record from the table, if possible. It might fail if constraints exist that
// prevent some or all rows from being deleted; use the force option to override this.
//
// When 'force' is set true, be aware of the following consequences.
// When using Mysql, foreign keys in other tables can be left dangling.
// When using Postgres, a cascade happens, so all 'adjacent' tables (i.e. linked by foreign keys)
// are also truncated.
func (tbl HookTable) Truncate(ctx context.Context, force bool) (err error) {
	for _, query := range tbl.Dialect().TruncateDDL(tbl.Name().String(), force) {
		_, err = support.Exec(ctx, tbl, nil, query)
		if err != nil {
			return err
		}
	}
	return nil
}

//-------------------------------------------------------------------------------------------------

// Exec executes a query without returning any rows.
// It returns the number of rows affected (if the database driver supports this).
//
// The args are for any placeholder parameters in the query.
//
// If the context ctx is nil, it defaults to context.Background().
func (tbl HookTable) Exec(ctx context.Context, req require.Requirement, query string, args ...interface{}) (int64, error) {
	return support.Exec(ctx, tbl, req, query, args...)
}

//-------------------------------------------------------------------------------------------------

// Query is the low-level request method for this table. The SQL query must return all the columns necessary for
// Hook values. Placeholders should be vanilla '?' marks, which will be replaced if necessary according to
// the chosen dialect.
//
// The query is logged using whatever logger is configured. If an error arises, this too is logged.
//
// The args are for any placeholder parameters in the query.
//
// The support API provides a core 'support.Query' function, on which this method depends. If appropriate,
// use that function directly; wrap the result in *sqlapi.Rows if you need to access its data as a map.
//
// If the context ctx is nil, it defaults to context.Background().
func (tbl HookTable) Query(ctx context.Context, req require.Requirement, query string, args ...interface{}) (HookList, error) {
	return doHookTableQueryAndScan(ctx, tbl, req, false, query, args)
}

func doHookTableQueryAndScan(ctx context.Context, tbl HookTabler, req require.Requirement, firstOnly bool, query string, args ...interface{}) (HookList, error) {
	rows, err := support.Query(ctx, tbl, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vv, n, err := ScanHooks(query, rows, firstOnly)
	return vv, tbl.Logger().LogIfError(require.ChainErrorIfQueryNotSatisfiedBy(err, req, n))
}

//-------------------------------------------------------------------------------------------------

// QueryOneNullString is a low-level access method for one string. This can be used for function queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl HookTable) QueryOneNullString(ctx context.Context, req require.Requirement, query string, args ...interface{}) (result sql.NullString, err error) {
	err = support.QueryOneNullThing(ctx, tbl, req, &result, query, args...)
	return result, err
}

// QueryOneNullInt64 is a low-level access method for one int64. This can be used for 'COUNT(1)' queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl HookTable) QueryOneNullInt64(ctx context.Context, req require.Requirement, query string, args ...interface{}) (result sql.NullInt64, err error) {
	err = support.QueryOneNullThing(ctx, tbl, req, &result, query, args...)
	return result, err
}

// QueryOneNullFloat64 is a low-level access method for one float64. This can be used for 'AVG(...)' queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl HookTable) QueryOneNullFloat64(ctx context.Context, req require.Requirement, query string, args ...interface{}) (result sql.NullFloat64, err error) {
	err = support.QueryOneNullThing(ctx, tbl, req, &result, query, args...)
	return result, err
}

// ScanHooks reads rows from the database and returns a slice of corresponding values.
// It also returns a number indicating how many rows were read; this will be larger than the length of the
// slice if reading stopped after the first row.
func ScanHooks(query string, rows sqlapi.SqlRows, firstOnly bool) (vv HookList, n int64, err error) {
	for rows.Next() {
		n++

		var v0 uint64
		var v1 string
		var v2 string
		var v3 string
		var v4 Category
		var v5 bool
		var v6 bool
		var v7 bool
		var v8 string
		var v9 string
		var v10 string
		var v11 string
		var v12 Email
		var v13 string
		var v14 string
		var v15 Email
		var v16 string

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
			&v12,
			&v13,
			&v14,
			&v15,
			&v16,
		)
		if err != nil {
			return vv, n, errors.Wrap(err, query)
		}

		v := &Hook{}
		v.Id = v0
		v.Sha = v1
		v.Bounds.After = v2
		v.Bounds.Before = v3
		v.Category = v4
		v.Created = v5
		v.Deleted = v6
		v.Forced = v7
		v.HeadCommit.ID = v8
		v.HeadCommit.Message = v9
		v.HeadCommit.Timestamp = v10
		v.HeadCommit.Author.Name = v11
		v.HeadCommit.Author.Email = v12
		v.HeadCommit.Author.Username = v13
		v.HeadCommit.Committer.Name = v14
		v.HeadCommit.Committer.Email = v15
		v.HeadCommit.Committer.Username = v16

		var iv interface{} = v
		if hook, ok := iv.(sqlapi.CanPostGet); ok {
			err = hook.PostGet()
			if err != nil {
				return vv, n, errors.Wrap(err, query)
			}
		}

		vv = append(vv, v)

		if firstOnly {
			if rows.Next() {
				n++
			}
			return vv, n, errors.Wrap(rows.Err(), query)
		}
	}

	return vv, n, errors.Wrap(rows.Err(), query)
}

//--------------------------------------------------------------------------------

func allHookColumnNamesQuoted(q quote.Quoter) string {
	return strings.Join(q.QuoteN(listOfHookTableColumnNames), ",")
}

//--------------------------------------------------------------------------------

// GetHookById gets the record with a given primary key value.
// If not found, *Hook will be nil.
func (tbl HookTable) GetHookById(ctx context.Context, req require.Requirement, id uint64) (*Hook, error) {
	return tbl.SelectOne(ctx, req, where.Eq("id", id), nil)
}

// GetHooksById gets records from the table according to a list of primary keys.
// Although the list of ids can be arbitrarily long, there are practical limits;
// note that Oracle DB has a limit of 1000.
//
// It places a requirement, which may be nil, on the size of the expected results: in particular, require.All
// controls whether an error is generated not all the ids produce a result.
func (tbl HookTable) GetHooksById(ctx context.Context, req require.Requirement, qc where.QueryConstraint, id ...uint64) (list HookList, err error) {
	if req == require.All {
		req = require.Exactly(len(id))
	}
	return tbl.Select(ctx, req, where.In("id", id), qc)
}

func doHookTableQueryAndScanOne(ctx context.Context, tbl HookTabler, req require.Requirement, query string, args ...interface{}) (*Hook, error) {
	list, err := doHookTableQueryAndScan(ctx, tbl, req, true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

// Fetch fetches a list of Hook based on a supplied query. This is mostly used for join queries that map its
// result columns to the fields of Hook. Other queries might be better handled by GetXxx or Select methods.
func (tbl HookTable) Fetch(ctx context.Context, req require.Requirement, query string, args ...interface{}) (HookList, error) {
	return doHookTableQueryAndScan(ctx, tbl, req, false, query, args...)
}

//-------------------------------------------------------------------------------------------------

// SelectOneWhere allows a single Hook to be obtained from the table that matches a 'where' clause
// and some limit. Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
// If not found, *Example will be nil.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.One
// controls whether an error is generated when no result is found.
//
// The args are for any placeholder parameters in the query.
func (tbl HookTable) SelectOneWhere(ctx context.Context, req require.Requirement, where, orderBy string, args ...interface{}) (*Hook, error) {
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT 1",
		allHookColumnNamesQuoted(tbl.Dialect().Quoter()), quotedName, where, orderBy)
	v, err := doHookTableQueryAndScanOne(ctx, tbl, req, query, args...)
	return v, err
}

// SelectOne allows a single Hook to be obtained from the table that matches a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
// If not found, *Example will be nil.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.One
// controls whether an error is generated when no result is found.
func (tbl HookTable) SelectOne(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) (*Hook, error) {
	q := tbl.Dialect().Quoter()
	whs, args := where.Where(wh, q)
	orderBy := where.Build(qc, q)
	return tbl.SelectOneWhere(ctx, req, whs, orderBy, args...)
}

// SelectWhere allows Hooks to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
//
// The args are for any placeholder parameters in the query.
func (tbl HookTable) SelectWhere(ctx context.Context, req require.Requirement, where, orderBy string, args ...interface{}) (HookList, error) {
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("SELECT %s FROM %s %s %s",
		allHookColumnNamesQuoted(tbl.Dialect().Quoter()), quotedName, where, orderBy)
	vv, err := doHookTableQueryAndScan(ctx, tbl, req, false, query, args...)
	return vv, err
}

// Select allows Hooks to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
func (tbl HookTable) Select(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) (HookList, error) {
	q := tbl.Dialect().Quoter()
	whs, args := where.Where(wh, q)
	orderBy := where.Build(qc, q)
	return tbl.SelectWhere(ctx, req, whs, orderBy, args...)
}

//-------------------------------------------------------------------------------------------------

// CountWhere counts Hooks in the table that match a 'where' clause.
// Use a blank string for the 'where' argument if it is not needed.
//
// The args are for any placeholder parameters in the query.
func (tbl HookTable) CountWhere(ctx context.Context, where string, args ...interface{}) (count int64, err error) {
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s %s", quotedName, where)
	rows, err := support.Query(ctx, tbl, query, args...)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&count)
	}
	return count, tbl.Logger().LogIfError(err)
}

// Count counts the Hooks in the table that match a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed.
func (tbl HookTable) Count(ctx context.Context, wh where.Expression) (count int64, err error) {
	whs, args := where.Where(wh, tbl.Dialect().Quoter())
	return tbl.CountWhere(ctx, whs, args...)
}

//--------------------------------------------------------------------------------

// SliceId gets the id column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceId(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]uint64, error) {
	return support.SliceUint64List(ctx, tbl, req, tbl.pk, wh, qc)
}

// SliceSha gets the sha column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceSha(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return support.SliceStringList(ctx, tbl, req, "sha", wh, qc)
}

// SliceAfter gets the after column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceAfter(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return support.SliceStringList(ctx, tbl, req, "after", wh, qc)
}

// SliceBefore gets the before column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceBefore(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return support.SliceStringList(ctx, tbl, req, "before", wh, qc)
}

// SliceCommitId gets the commit_id column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceCommitId(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return support.SliceStringList(ctx, tbl, req, "commit_id", wh, qc)
}

// SliceMessage gets the message column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceMessage(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return support.SliceStringList(ctx, tbl, req, "message", wh, qc)
}

// SliceTimestamp gets the timestamp column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceTimestamp(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return support.SliceStringList(ctx, tbl, req, "timestamp", wh, qc)
}

// SliceHeadCommitAuthorName gets the head_commit_author_name column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceHeadCommitAuthorName(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return support.SliceStringList(ctx, tbl, req, "head_commit_author_name", wh, qc)
}

// SliceHeadCommitAuthorUsername gets the head_commit_author_username column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceHeadCommitAuthorUsername(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return support.SliceStringList(ctx, tbl, req, "head_commit_author_username", wh, qc)
}

// SliceHeadCommitCommitterName gets the head_commit_committer_name column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceHeadCommitCommitterName(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return support.SliceStringList(ctx, tbl, req, "head_commit_committer_name", wh, qc)
}

// SliceHeadCommitCommitterUsername gets the head_commit_committer_username column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceHeadCommitCommitterUsername(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return support.SliceStringList(ctx, tbl, req, "head_commit_committer_username", wh, qc)
}

// SliceCategory gets the category column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceCategory(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]Category, error) {
	return sliceHookTableCategoryList(ctx, tbl, req, "category", wh, qc)
}

// SliceHeadCommitAuthorEmail gets the head_commit_author_email column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceHeadCommitAuthorEmail(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]Email, error) {
	return sliceHookTableEmailList(ctx, tbl, req, "head_commit_author_email", wh, qc)
}

// SliceHeadCommitCommitterEmail gets the head_commit_committer_email column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceHeadCommitCommitterEmail(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]Email, error) {
	return sliceHookTableEmailList(ctx, tbl, req, "head_commit_committer_email", wh, qc)
}

func sliceHookTableCategoryList(ctx context.Context, tbl HookTabler, req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]Category, error) {
	q := tbl.Dialect().Quoter()
	whs, args := where.Where(wh, q)
	orderBy := where.Build(qc, q)
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", q.Quote(sqlname), quotedName, whs, orderBy)
	rows, err := support.Query(ctx, tbl, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]Category, 0, 10)

	for rows.Next() {
		var v Category
		err = rows.Scan(&v)
		if err == sql.ErrNoRows {
			return list, tbl.Logger().LogIfError(require.ErrorIfQueryNotSatisfiedBy(req, int64(len(list))))
		} else {
			list = append(list, v)
		}
	}
	return list, tbl.Logger().LogIfError(require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, int64(len(list))))
}

func sliceHookTableEmailList(ctx context.Context, tbl HookTabler, req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]Email, error) {
	q := tbl.Dialect().Quoter()
	whs, args := where.Where(wh, q)
	orderBy := where.Build(qc, q)
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", q.Quote(sqlname), quotedName, whs, orderBy)
	rows, err := support.Query(ctx, tbl, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]Email, 0, 10)

	for rows.Next() {
		var v Email
		err = rows.Scan(&v)
		if err == sql.ErrNoRows {
			return list, tbl.Logger().LogIfError(require.ErrorIfQueryNotSatisfiedBy(req, int64(len(list))))
		} else {
			list = append(list, v)
		}
	}
	return list, tbl.Logger().LogIfError(require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, int64(len(list))))
}

func constructHookTableInsert(tbl HookTable, w dialect.StringWriter, v *Hook, withPk bool) (s []interface{}, err error) {
	q := tbl.Dialect().Quoter()
	s = make([]interface{}, 0, 17)

	comma := ""
	w.WriteString(" (")

	if withPk {
		q.QuoteW(w, "id")
		comma = ","
		s = append(s, v.Id)
	}

	w.WriteString(comma)
	q.QuoteW(w, "sha")
	s = append(s, v.Sha)
	comma = ","

	w.WriteString(comma)
	q.QuoteW(w, "after")
	s = append(s, v.Bounds.After)

	w.WriteString(comma)
	q.QuoteW(w, "before")
	s = append(s, v.Bounds.Before)

	w.WriteString(comma)
	q.QuoteW(w, "category")
	s = append(s, v.Category)

	w.WriteString(comma)
	q.QuoteW(w, "created")
	s = append(s, v.Created)

	w.WriteString(comma)
	q.QuoteW(w, "deleted")
	s = append(s, v.Deleted)

	w.WriteString(comma)
	q.QuoteW(w, "forced")
	s = append(s, v.Forced)

	w.WriteString(comma)
	q.QuoteW(w, "commit_id")
	s = append(s, v.HeadCommit.ID)

	w.WriteString(comma)
	q.QuoteW(w, "message")
	s = append(s, v.HeadCommit.Message)

	w.WriteString(comma)
	q.QuoteW(w, "timestamp")
	s = append(s, v.HeadCommit.Timestamp)

	w.WriteString(comma)
	q.QuoteW(w, "head_commit_author_name")
	s = append(s, v.HeadCommit.Author.Name)

	w.WriteString(comma)
	q.QuoteW(w, "head_commit_author_email")
	s = append(s, v.HeadCommit.Author.Email)

	w.WriteString(comma)
	q.QuoteW(w, "head_commit_author_username")
	s = append(s, v.HeadCommit.Author.Username)

	w.WriteString(comma)
	q.QuoteW(w, "head_commit_committer_name")
	s = append(s, v.HeadCommit.Committer.Name)

	w.WriteString(comma)
	q.QuoteW(w, "head_commit_committer_email")
	s = append(s, v.HeadCommit.Committer.Email)

	w.WriteString(comma)
	q.QuoteW(w, "head_commit_committer_username")
	s = append(s, v.HeadCommit.Committer.Username)

	w.WriteString(")")
	return s, nil
}

func constructHookTableUpdate(tbl HookTable, w dialect.StringWriter, v *Hook) (s []interface{}, err error) {
	q := tbl.Dialect().Quoter()
	j := 1
	s = make([]interface{}, 0, 16)

	comma := ""

	w.WriteString(comma)
	q.QuoteW(w, "sha")
	w.WriteString("=?")
	s = append(s, v.Sha)
	j++
	comma = ", "

	w.WriteString(comma)
	q.QuoteW(w, "after")
	w.WriteString("=?")
	s = append(s, v.Bounds.After)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "before")
	w.WriteString("=?")
	s = append(s, v.Bounds.Before)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "category")
	w.WriteString("=?")
	s = append(s, v.Category)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "created")
	w.WriteString("=?")
	s = append(s, v.Created)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "deleted")
	w.WriteString("=?")
	s = append(s, v.Deleted)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "forced")
	w.WriteString("=?")
	s = append(s, v.Forced)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "commit_id")
	w.WriteString("=?")
	s = append(s, v.HeadCommit.ID)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "message")
	w.WriteString("=?")
	s = append(s, v.HeadCommit.Message)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "timestamp")
	w.WriteString("=?")
	s = append(s, v.HeadCommit.Timestamp)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "head_commit_author_name")
	w.WriteString("=?")
	s = append(s, v.HeadCommit.Author.Name)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "head_commit_author_email")
	w.WriteString("=?")
	s = append(s, v.HeadCommit.Author.Email)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "head_commit_author_username")
	w.WriteString("=?")
	s = append(s, v.HeadCommit.Author.Username)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "head_commit_committer_name")
	w.WriteString("=?")
	s = append(s, v.HeadCommit.Committer.Name)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "head_commit_committer_email")
	w.WriteString("=?")
	s = append(s, v.HeadCommit.Committer.Email)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "head_commit_committer_username")
	w.WriteString("=?")
	s = append(s, v.HeadCommit.Committer.Username)
	j++
	return s, nil
}

//--------------------------------------------------------------------------------

// Insert adds new records for the Hooks.// The Hooks have their primary key fields set to the new record identifiers.
// The Hook.PreInsert() method will be called, if it exists.
func (tbl HookTable) Insert(ctx context.Context, req require.Requirement, vv ...*Hook) error {
	if req == require.All {
		req = require.Exactly(len(vv))
	}

	if ctx == nil {
		ctx = context.Background()
	}

	var count int64
	returning := ""
	insertHasReturningPhrase := tbl.Dialect().InsertHasReturningPhrase()
	if insertHasReturningPhrase {
		returning = fmt.Sprintf(" RETURNING %q", tbl.pk)
	}

	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlapi.CanPreInsert); ok {
			err := hook.PreInsert()
			if err != nil {
				return tbl.Logger().LogError(err)
			}
		}

		b := dialect.Adapt(&bytes.Buffer{})
		b.WriteString("INSERT INTO ")
		tbl.quotedNameW(b)

		fields, err := constructHookTableInsert(tbl, b, v, false)
		if err != nil {
			return tbl.Logger().LogError(err)
		}

		b.WriteString(" VALUES (")
		b.WriteString(tbl.Dialect().Placeholders(len(fields)))
		b.WriteString(")")
		b.WriteString(returning)

		query := b.String()
		tbl.Logger().LogQuery(query, fields...)

		var n int64 = 1
		if insertHasReturningPhrase {
			row := tbl.db.QueryRowContext(ctx, query, fields...)
			var i64 int64
			err = row.Scan(&i64)
			v.Id = uint64(i64)

		} else {
			i64, e2 := tbl.db.InsertContext(ctx, tbl.pk, query, fields...)
			if e2 != nil {
				return tbl.Logger().LogError(e2)
			}
			v.Id = uint64(i64)
		}

		if err != nil {
			return tbl.Logger().LogError(err)
		}
		count += n
	}

	return tbl.Logger().LogIfError(require.ErrorIfExecNotSatisfiedBy(req, count))
}

// UpdateById updates one or more columns, given a id value.
func (tbl HookTable) UpdateById(ctx context.Context, req require.Requirement, id uint64, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("id", id), fields...)
}

// UpdateBySha updates one or more columns, given a sha value.
func (tbl HookTable) UpdateBySha(ctx context.Context, req require.Requirement, sha string, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("sha", sha), fields...)
}

// UpdateByAfter updates one or more columns, given a after value.
func (tbl HookTable) UpdateByAfter(ctx context.Context, req require.Requirement, after string, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("after", after), fields...)
}

// UpdateByBefore updates one or more columns, given a before value.
func (tbl HookTable) UpdateByBefore(ctx context.Context, req require.Requirement, before string, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("before", before), fields...)
}

// UpdateByCategory updates one or more columns, given a category value.
func (tbl HookTable) UpdateByCategory(ctx context.Context, req require.Requirement, category Category, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("category", category), fields...)
}

// UpdateByCommitId updates one or more columns, given a commit_id value.
func (tbl HookTable) UpdateByCommitId(ctx context.Context, req require.Requirement, commit_id string, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("commit_id", commit_id), fields...)
}

// UpdateByMessage updates one or more columns, given a message value.
func (tbl HookTable) UpdateByMessage(ctx context.Context, req require.Requirement, message string, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("message", message), fields...)
}

// UpdateByTimestamp updates one or more columns, given a timestamp value.
func (tbl HookTable) UpdateByTimestamp(ctx context.Context, req require.Requirement, timestamp string, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("timestamp", timestamp), fields...)
}

// UpdateByHeadCommitAuthorName updates one or more columns, given a head_commit_author_name value.
func (tbl HookTable) UpdateByHeadCommitAuthorName(ctx context.Context, req require.Requirement, head_commit_author_name string, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("head_commit_author_name", head_commit_author_name), fields...)
}

// UpdateByHeadCommitAuthorEmail updates one or more columns, given a head_commit_author_email value.
func (tbl HookTable) UpdateByHeadCommitAuthorEmail(ctx context.Context, req require.Requirement, head_commit_author_email Email, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("head_commit_author_email", head_commit_author_email), fields...)
}

// UpdateByHeadCommitAuthorUsername updates one or more columns, given a head_commit_author_username value.
func (tbl HookTable) UpdateByHeadCommitAuthorUsername(ctx context.Context, req require.Requirement, head_commit_author_username string, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("head_commit_author_username", head_commit_author_username), fields...)
}

// UpdateByHeadCommitCommitterName updates one or more columns, given a head_commit_committer_name value.
func (tbl HookTable) UpdateByHeadCommitCommitterName(ctx context.Context, req require.Requirement, head_commit_committer_name string, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("head_commit_committer_name", head_commit_committer_name), fields...)
}

// UpdateByHeadCommitCommitterEmail updates one or more columns, given a head_commit_committer_email value.
func (tbl HookTable) UpdateByHeadCommitCommitterEmail(ctx context.Context, req require.Requirement, head_commit_committer_email Email, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("head_commit_committer_email", head_commit_committer_email), fields...)
}

// UpdateByHeadCommitCommitterUsername updates one or more columns, given a head_commit_committer_username value.
func (tbl HookTable) UpdateByHeadCommitCommitterUsername(ctx context.Context, req require.Requirement, head_commit_committer_username string, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("head_commit_committer_username", head_commit_committer_username), fields...)
}

// UpdateFields updates one or more columns, given a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed (but note that this is risky!).
func (tbl HookTable) UpdateFields(ctx context.Context, req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error) {
	return support.UpdateFields(ctx, tbl, req, wh, fields...)
}

//--------------------------------------------------------------------------------

// Update updates records, matching them by primary key. It returns the number of rows affected.
// The Hook.PreUpdate(Execer) method will be called, if it exists.
func (tbl HookTable) Update(ctx context.Context, req require.Requirement, vv ...*Hook) (int64, error) {
	if req == require.All {
		req = require.Exactly(len(vv))
	}

	var count int64
	d := tbl.Dialect()
	q := d.Quoter()

	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlapi.CanPreUpdate); ok {
			err := hook.PreUpdate()
			if err != nil {
				return count, tbl.Logger().LogError(err)
			}
		}

		b := dialect.Adapt(&bytes.Buffer{})
		b.WriteString("UPDATE ")
		tbl.quotedNameW(b)
		b.WriteString(" SET ")

		args, err := constructHookTableUpdate(tbl, b, v)
		if err != nil {
			return count, err
		}
		args = append(args, v.Id)

		b.WriteString(" WHERE ")
		q.QuoteW(b, tbl.pk)
		b.WriteString("=?")

		query := b.String()
		n, err := tbl.Exec(ctx, nil, query, args...)
		if err != nil {
			return count, err
		}
		count += n
	}

	return count, tbl.Logger().LogIfError(require.ErrorIfExecNotSatisfiedBy(req, count))
}

//--------------------------------------------------------------------------------

// Upsert inserts or updates a record, matching it using the expression supplied.
// This expression is used to search for an existing record based on some specified
// key column(s). It must match either zero or one existing record. If it matches
// none, a new record is inserted; otherwise the matching record is updated. An
// error results if these conditions are not met.
func (tbl HookTable) Upsert(ctx context.Context, v *Hook, wh where.Expression) error {
	col := tbl.Dialect().Quoter().Quote(tbl.pk)
	qName := tbl.quotedName()
	whs, args := where.Where(wh, tbl.Dialect().Quoter())

	query := fmt.Sprintf("SELECT %s FROM %s %s", col, qName, whs)
	rows, err := support.Query(ctx, tbl, query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	if !rows.Next() {
		return tbl.Insert(ctx, require.One, v)
	}

	var id uint64
	err = rows.Scan(&id)
	if err != nil {
		return tbl.Logger().LogIfError(err)
	}

	if rows.Next() {
		return require.ErrWrongSize(2, "expected to find no more than 1 but got at least 2 using %q", wh)
	}

	v.Id = id
	_, err = tbl.Update(ctx, require.One, v)
	return err
}

//-------------------------------------------------------------------------------------------------

// DeleteById deletes rows from the table, given some id values.
// The list of ids can be arbitrarily long.
func (tbl HookTable) DeleteById(ctx context.Context, req require.Requirement, id ...uint64) (int64, error) {
	ii := support.Uint64AsInterfaceSlice(id)
	return support.DeleteByColumn(ctx, tbl, req, "id", ii...)
}

// DeleteBySha deletes rows from the table, given some sha values.
// The list of ids can be arbitrarily long.
func (tbl HookTable) DeleteBySha(ctx context.Context, req require.Requirement, sha ...string) (int64, error) {
	ii := support.StringAsInterfaceSlice(sha)
	return support.DeleteByColumn(ctx, tbl, req, "sha", ii...)
}

// DeleteByAfter deletes rows from the table, given some after values.
// The list of ids can be arbitrarily long.
func (tbl HookTable) DeleteByAfter(ctx context.Context, req require.Requirement, after ...string) (int64, error) {
	ii := support.StringAsInterfaceSlice(after)
	return support.DeleteByColumn(ctx, tbl, req, "after", ii...)
}

// DeleteByBefore deletes rows from the table, given some before values.
// The list of ids can be arbitrarily long.
func (tbl HookTable) DeleteByBefore(ctx context.Context, req require.Requirement, before ...string) (int64, error) {
	ii := support.StringAsInterfaceSlice(before)
	return support.DeleteByColumn(ctx, tbl, req, "before", ii...)
}

// DeleteByCategory deletes rows from the table, given some category values.
// The list of ids can be arbitrarily long.
func (tbl HookTable) DeleteByCategory(ctx context.Context, req require.Requirement, category ...Category) (int64, error) {
	ii := make([]interface{}, len(category))
	for i, v := range category {
		ii[i] = v
	}
	return support.DeleteByColumn(ctx, tbl, req, "category", ii...)
}

// DeleteByCommitId deletes rows from the table, given some commit_id values.
// The list of ids can be arbitrarily long.
func (tbl HookTable) DeleteByCommitId(ctx context.Context, req require.Requirement, commit_id ...string) (int64, error) {
	ii := support.StringAsInterfaceSlice(commit_id)
	return support.DeleteByColumn(ctx, tbl, req, "commit_id", ii...)
}

// DeleteByMessage deletes rows from the table, given some message values.
// The list of ids can be arbitrarily long.
func (tbl HookTable) DeleteByMessage(ctx context.Context, req require.Requirement, message ...string) (int64, error) {
	ii := support.StringAsInterfaceSlice(message)
	return support.DeleteByColumn(ctx, tbl, req, "message", ii...)
}

// DeleteByTimestamp deletes rows from the table, given some timestamp values.
// The list of ids can be arbitrarily long.
func (tbl HookTable) DeleteByTimestamp(ctx context.Context, req require.Requirement, timestamp ...string) (int64, error) {
	ii := support.StringAsInterfaceSlice(timestamp)
	return support.DeleteByColumn(ctx, tbl, req, "timestamp", ii...)
}

// DeleteByHeadCommitAuthorName deletes rows from the table, given some head_commit_author_name values.
// The list of ids can be arbitrarily long.
func (tbl HookTable) DeleteByHeadCommitAuthorName(ctx context.Context, req require.Requirement, head_commit_author_name ...string) (int64, error) {
	ii := support.StringAsInterfaceSlice(head_commit_author_name)
	return support.DeleteByColumn(ctx, tbl, req, "head_commit_author_name", ii...)
}

// DeleteByHeadCommitAuthorEmail deletes rows from the table, given some head_commit_author_email values.
// The list of ids can be arbitrarily long.
func (tbl HookTable) DeleteByHeadCommitAuthorEmail(ctx context.Context, req require.Requirement, head_commit_author_email ...Email) (int64, error) {
	ii := make([]interface{}, len(head_commit_author_email))
	for i, v := range head_commit_author_email {
		ii[i] = v
	}
	return support.DeleteByColumn(ctx, tbl, req, "head_commit_author_email", ii...)
}

// DeleteByHeadCommitAuthorUsername deletes rows from the table, given some head_commit_author_username values.
// The list of ids can be arbitrarily long.
func (tbl HookTable) DeleteByHeadCommitAuthorUsername(ctx context.Context, req require.Requirement, head_commit_author_username ...string) (int64, error) {
	ii := support.StringAsInterfaceSlice(head_commit_author_username)
	return support.DeleteByColumn(ctx, tbl, req, "head_commit_author_username", ii...)
}

// DeleteByHeadCommitCommitterName deletes rows from the table, given some head_commit_committer_name values.
// The list of ids can be arbitrarily long.
func (tbl HookTable) DeleteByHeadCommitCommitterName(ctx context.Context, req require.Requirement, head_commit_committer_name ...string) (int64, error) {
	ii := support.StringAsInterfaceSlice(head_commit_committer_name)
	return support.DeleteByColumn(ctx, tbl, req, "head_commit_committer_name", ii...)
}

// DeleteByHeadCommitCommitterEmail deletes rows from the table, given some head_commit_committer_email values.
// The list of ids can be arbitrarily long.
func (tbl HookTable) DeleteByHeadCommitCommitterEmail(ctx context.Context, req require.Requirement, head_commit_committer_email ...Email) (int64, error) {
	ii := make([]interface{}, len(head_commit_committer_email))
	for i, v := range head_commit_committer_email {
		ii[i] = v
	}
	return support.DeleteByColumn(ctx, tbl, req, "head_commit_committer_email", ii...)
}

// DeleteByHeadCommitCommitterUsername deletes rows from the table, given some head_commit_committer_username values.
// The list of ids can be arbitrarily long.
func (tbl HookTable) DeleteByHeadCommitCommitterUsername(ctx context.Context, req require.Requirement, head_commit_committer_username ...string) (int64, error) {
	ii := support.StringAsInterfaceSlice(head_commit_committer_username)
	return support.DeleteByColumn(ctx, tbl, req, "head_commit_committer_username", ii...)
}

// Delete deletes one or more rows from the table, given a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl HookTable) Delete(ctx context.Context, req require.Requirement, wh where.Expression) (int64, error) {
	query, args := deleteRowsHookTableSql(tbl, wh)
	return tbl.Exec(ctx, req, query, args...)
}

func deleteRowsHookTableSql(tbl HookTabler, wh where.Expression) (string, []interface{}) {
	whs, args := where.Where(wh, tbl.Dialect().Quoter())
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("DELETE FROM %s %s", quotedName, whs)
	return query, args
}

//-------------------------------------------------------------------------------------------------
