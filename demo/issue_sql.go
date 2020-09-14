// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.
// sqlapi v0.47.1; sqlgen v0.67.0

package demo

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
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

// IssueTabler lists table methods provided by IssueTable.
type IssueTabler interface {
	sqlapi.Table

	// Constraints returns the table's constraints.
	Constraints() constraint.Constraints

	// WithConstraint returns a modified IssueTabler with added data consistency constraints.
	WithConstraint(cc ...constraint.Constraint) IssueTabler

	// WithPrefix returns a modified IssueTabler with a given table name prefix.
	WithPrefix(pfx string) IssueTabler

	// CreateTable creates the table.
	CreateTable(ctx context.Context, ifNotExists bool) (int64, error)

	// DropTable drops the table, destroying all its data.
	DropTable(ctx context.Context, ifExists bool) (int64, error)

	// CreateTableWithIndexes invokes CreateTable then CreateIndexes.
	CreateTableWithIndexes(ctx context.Context, ifNotExist bool) (err error)

	// CreateIndexes executes queries that create the indexes needed by the Issue table.
	CreateIndexes(ctx context.Context, ifNotExist bool) (err error)

	// CreateIssueAssigneeIndex creates the issue_assignee index.
	CreateIssueAssigneeIndex(ctx context.Context, ifNotExist bool) error

	// DropIssueAssigneeIndex drops the issue_assignee index.
	DropIssueAssigneeIndex(ctx context.Context, ifExists bool) error

	// Truncate drops every record from the table, if possible.
	Truncate(ctx context.Context, force bool) (err error)
}

//-------------------------------------------------------------------------------------------------

// IssueQueryer lists query methods provided by IssueTable.
type IssueQueryer interface {
	sqlapi.Table

	// Using returns a modified IssueQueryer using the Execer supplied,
	// which will typically be a transaction (i.e. SqlTx).
	Using(tx sqlapi.Execer) IssueQueryer

	// Transact runs the function provided within a transaction. The transction is committed
	// unless an error occurs.
	Transact(ctx context.Context, txOptions *sql.TxOptions, fn func(IssueQueryer) error) error

	// Exec executes a query without returning any rows.
	Exec(ctx context.Context, req require.Requirement, query string, args ...interface{}) (int64, error)

	// Query is the low-level request method for this table using an SQL query that must return all the columns
	// necessary for Issue values.
	Query(ctx context.Context, req require.Requirement, query string, args ...interface{}) ([]*Issue, error)

	// QueryOneNullString is a low-level access method for one string, returning the first match.
	QueryOneNullString(ctx context.Context, req require.Requirement, query string, args ...interface{}) (result sql.NullString, err error)

	// QueryOneNullInt64 is a low-level access method for one int64, returning the first match.
	QueryOneNullInt64(ctx context.Context, req require.Requirement, query string, args ...interface{}) (result sql.NullInt64, err error)

	// QueryOneNullFloat64 is a low-level access method for one float64, returning the first match.
	QueryOneNullFloat64(ctx context.Context, req require.Requirement, query string, args ...interface{}) (result sql.NullFloat64, err error)

	// GetIssueById gets the record with a given primary key value.
	GetIssueById(ctx context.Context, req require.Requirement, id int64) (*Issue, error)

	// GetIssuesById gets records from the table according to a list of primary keys.
	GetIssuesById(ctx context.Context, req require.Requirement, qc where.QueryConstraint, id ...int64) (list []*Issue, err error)

	// GetIssuesByAssignee gets the records with a given assignee value.
	GetIssuesByAssignee(ctx context.Context, req require.Requirement, qc where.QueryConstraint, assignee string) ([]*Issue, error)

	// Fetch fetches a list of Issue based on a supplied query. This is mostly used for join queries that map its
	// result columns to the fields of Issue. Other queries might be better handled by GetXxx or Select methods.
	Fetch(ctx context.Context, req require.Requirement, query string, args ...interface{}) ([]*Issue, error)

	// SelectOneWhere allows a single Issue to be obtained from the table that matches a 'where' clause.
	SelectOneWhere(ctx context.Context, req require.Requirement, where, orderBy string, args ...interface{}) (*Issue, error)

	// SelectOne allows a single Issue to be obtained from the table that matches a 'where' clause.
	SelectOne(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) (*Issue, error)

	// SelectWhere allows Issues to be obtained from the table that match a 'where' clause.
	SelectWhere(ctx context.Context, req require.Requirement, where, orderBy string, args ...interface{}) ([]*Issue, error)

	// Select allows Issues to be obtained from the table that match a 'where' clause.
	Select(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]*Issue, error)

	// CountWhere counts Issues in the table that match a 'where' clause.
	CountWhere(ctx context.Context, where string, args ...interface{}) (count int64, err error)

	// Count counts the Issues in the table that match a 'where' clause.
	Count(ctx context.Context, wh where.Expression) (count int64, err error)

	// SliceId gets the id column for all rows that match the 'where' condition.
	SliceId(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error)

	// SliceNumber gets the number column for all rows that match the 'where' condition.
	SliceNumber(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int, error)

	// SliceTitle gets the title column for all rows that match the 'where' condition.
	SliceTitle(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error)

	// SliceBigbody gets the bigbody column for all rows that match the 'where' condition.
	SliceBigbody(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error)

	// SliceAssignee gets the assignee column for all rows that match the 'where' condition.
	SliceAssignee(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error)

	// SliceState gets the state column for all rows that match the 'where' condition.
	SliceState(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error)

	// Insert adds new records for the Issues, setting the primary key field for each one.
	Insert(ctx context.Context, req require.Requirement, vv ...*Issue) error

	// UpdateById updates one or more columns, given a id value.
	UpdateById(ctx context.Context, req require.Requirement, id int64, fields ...sql.NamedArg) (int64, error)

	// UpdateByNumber updates one or more columns, given a number value.
	UpdateByNumber(ctx context.Context, req require.Requirement, number int, fields ...sql.NamedArg) (int64, error)

	// UpdateByTitle updates one or more columns, given a title value.
	UpdateByTitle(ctx context.Context, req require.Requirement, title string, fields ...sql.NamedArg) (int64, error)

	// UpdateByBigbody updates one or more columns, given a bigbody value.
	UpdateByBigbody(ctx context.Context, req require.Requirement, bigbody string, fields ...sql.NamedArg) (int64, error)

	// UpdateByAssignee updates one or more columns, given a assignee value.
	UpdateByAssignee(ctx context.Context, req require.Requirement, assignee string, fields ...sql.NamedArg) (int64, error)

	// UpdateByState updates one or more columns, given a state value.
	UpdateByState(ctx context.Context, req require.Requirement, state string, fields ...sql.NamedArg) (int64, error)

	// UpdateFields updates one or more columns, given a 'where' clause.
	UpdateFields(ctx context.Context, req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error)

	// Update updates records, matching them by primary key.
	Update(ctx context.Context, req require.Requirement, vv ...*Issue) (int64, error)

	// Upsert inserts or updates a record, matching it using the expression supplied.
	// This expression is used to search for an existing record based on some specified
	// key column(s). It must match either zero or one existing record. If it matches
	// none, a new record is inserted; otherwise the matching record is updated. An
	// error results if these conditions are not met.
	Upsert(ctx context.Context, v *Issue, wh where.Expression) error

	// DeleteById deletes rows from the table, given some id values.
	// The list of ids can be arbitrarily long.
	DeleteById(ctx context.Context, req require.Requirement, id ...int64) (int64, error)

	// DeleteByNumber deletes rows from the table, given some number values.
	// The list of ids can be arbitrarily long.
	DeleteByNumber(ctx context.Context, req require.Requirement, number ...int) (int64, error)

	// DeleteByTitle deletes rows from the table, given some title values.
	// The list of ids can be arbitrarily long.
	DeleteByTitle(ctx context.Context, req require.Requirement, title ...string) (int64, error)

	// DeleteByBigbody deletes rows from the table, given some bigbody values.
	// The list of ids can be arbitrarily long.
	DeleteByBigbody(ctx context.Context, req require.Requirement, bigbody ...string) (int64, error)

	// DeleteByAssignee deletes rows from the table, given some assignee values.
	// The list of ids can be arbitrarily long.
	DeleteByAssignee(ctx context.Context, req require.Requirement, assignee ...string) (int64, error)

	// DeleteByState deletes rows from the table, given some state values.
	// The list of ids can be arbitrarily long.
	DeleteByState(ctx context.Context, req require.Requirement, state ...string) (int64, error)

	// Delete deletes one or more rows from the table, given a 'where' clause.
	// Use a nil value for the 'wh' argument if it is not needed (very risky!).
	Delete(ctx context.Context, req require.Requirement, wh where.Expression) (int64, error)
}

//-------------------------------------------------------------------------------------------------

// IssueTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type IssueTable struct {
	name        sqlapi.TableName
	database    sqlapi.Database
	db          sqlapi.Execer
	constraints constraint.Constraints
	pk          string
}

// Type conformance checks
var _ sqlapi.TableWithIndexes = &IssueTable{}

// NewIssueTable returns a new table instance.
// If a blank table name is supplied, the default name "issues" will be used instead.
// The request context is initialised with the background.
func NewIssueTable(name string, d sqlapi.Database) IssueTable {
	if name == "" {
		name = "issues"
	}
	var constraints constraint.Constraints
	return IssueTable{
		name:        sqlapi.TableName{Prefix: "", Name: name},
		database:    d,
		db:          d.DB(),
		constraints: constraints,
		pk:          "id",
	}
}

// CopyTableAsIssueTable copies a table instance, retaining the name etc but
// providing methods appropriate for 'Issue'.It doesn't copy the constraints of the original table.
//
// It serves to provide methods appropriate for 'Issue'. This is most useful when this is used to represent a
// join result. In such cases, there won't be any need for DDL methods, nor Exec, Insert, Update or Delete.
func CopyTableAsIssueTable(origin sqlapi.Table) IssueTable {
	return IssueTable{
		name:        origin.Name(),
		database:    origin.Database(),
		db:          origin.Execer(),
		constraints: nil,
		pk:          "id",
	}
}

// SetPkColumn sets the name of the primary key column. It defaults to "id".
// The result is a modified copy of the table; the original is unchanged.
//func (tbl IssueTable) SetPkColumn(pk string) IssueTabler {
//	tbl.pk = pk
//	return tbl
//}

// WithPrefix sets the table name prefix for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl IssueTable) WithPrefix(pfx string) IssueTabler {
	tbl.name.Prefix = pfx
	return tbl
}

// Database gets the shared database information.
func (tbl IssueTable) Database() sqlapi.Database {
	return tbl.database
}

// Logger gets the trace logger.
func (tbl IssueTable) Logger() sqlapi.Logger {
	return tbl.database.Logger()
}

// WithConstraint returns a modified IssueTabler with added data consistency constraints.
func (tbl IssueTable) WithConstraint(cc ...constraint.Constraint) IssueTabler {
	tbl.constraints = append(tbl.constraints, cc...)
	return tbl
}

// Constraints returns the table's constraints.
func (tbl IssueTable) Constraints() constraint.Constraints {
	return tbl.constraints
}

// Dialect gets the database dialect.
func (tbl IssueTable) Dialect() dialect.Dialect {
	return tbl.database.Dialect()
}

// Name gets the table name.
func (tbl IssueTable) Name() sqlapi.TableName {
	return tbl.name
}

// PkColumn gets the column name used as a primary key.
func (tbl IssueTable) PkColumn() string {
	return tbl.pk
}

// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl IssueTable) DB() sqlapi.SqlDB {
	return tbl.db.(sqlapi.SqlDB)
}

// Execer gets the wrapped database or transaction handle.
func (tbl IssueTable) Execer() sqlapi.Execer {
	return tbl.db
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl IssueTable) Tx() sqlapi.SqlTx {
	return tbl.db.(sqlapi.SqlTx)
}

// IsTx tests whether this is within a transaction.
func (tbl IssueTable) IsTx() bool {
	return tbl.db.IsTx()
}

// Using returns a modified IssueTabler using the the Execer supplied,
// which will typically be a transaction (i.e. SqlTx). This is needed when making multiple
// queries across several tables within a single transaction.
//
// The result is a modified copy of the table; the original is unchanged.
func (tbl IssueTable) Using(tx sqlapi.Execer) IssueQueryer {
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
func (tbl IssueTable) Transact(ctx context.Context, txOptions *sql.TxOptions, fn func(IssueQueryer) error) error {
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

func (tbl IssueTable) quotedName() string {
	return tbl.Dialect().Quoter().Quote(tbl.name.String())
}

func (tbl IssueTable) quotedNameW(w dialect.StringWriter) {
	tbl.Dialect().Quoter().QuoteW(w, tbl.name.String())
}

//-------------------------------------------------------------------------------------------------

// NumIssueTableColumns is the total number of columns in IssueTable.
const NumIssueTableColumns = 8

// NumIssueTableDataColumns is the number of columns in IssueTable not including the auto-increment key.
const NumIssueTableDataColumns = 7

// IssueTableColumnNames is the list of columns in IssueTable.
const IssueTableColumnNames = "id,number,date,title,bigbody,assignee,state,labels"

// IssueTableDataColumnNames is the list of data columns in IssueTable.
const IssueTableDataColumnNames = "number,date,title,bigbody,assignee,state,labels"

var listOfIssueTableColumnNames = strings.Split(IssueTableColumnNames, ",")

//-------------------------------------------------------------------------------------------------

var sqlIssueTableCreateColumnsSqlite = []string{
	"integer not null primary key autoincrement",
	"bigint not null",
	"blob not null",
	"text not null",
	"text not null",
	"text not null",
	"text not null",
	"text",
}

var sqlIssueTableCreateColumnsMysql = []string{
	"bigint not null primary key auto_increment",
	"bigint not null",
	"mediumblob not null",
	"varchar(512) not null",
	"varchar(2048) not null",
	"varchar(255) not null",
	"varchar(50) not null",
	"json",
}

var sqlIssueTableCreateColumnsPostgres = []string{
	"bigserial not null primary key",
	"bigint not null",
	"bytea not null",
	"text not null",
	"text not null",
	"text not null",
	"text not null",
	"json",
}

var sqlIssueTableCreateColumnsPgx = []string{
	"bigserial not null primary key",
	"bigint not null",
	"bytea not null",
	"text not null",
	"text not null",
	"text not null",
	"text not null",
	"json",
}

//-------------------------------------------------------------------------------------------------

const sqlIssueAssigneeIndexColumns = "assignee"

var listOfIssueAssigneeIndexColumns = []string{"assignee"}

//-------------------------------------------------------------------------------------------------

// CreateTable creates the table.
func (tbl IssueTable) CreateTable(ctx context.Context, ifNotExists bool) (int64, error) {
	return support.Exec(ctx, tbl, nil, createIssueTableSql(tbl, ifNotExists))
}

func createIssueTableSql(tbl IssueTabler, ifNotExists bool) string {
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
		columns = sqlIssueTableCreateColumnsSqlite
	case dialect.MysqlIndex:
		columns = sqlIssueTableCreateColumnsMysql
	case dialect.PostgresIndex:
		columns = sqlIssueTableCreateColumnsPostgres
	case dialect.PgxIndex:
		columns = sqlIssueTableCreateColumnsPgx
	}

	comma := ""
	for i, n := range listOfIssueTableColumnNames {
		buf.WriteString(comma)
		q.QuoteW(buf, n)
		buf.WriteString(" ")
		buf.WriteString(columns[i])
		comma = ",\n "
	}

	for i, c := range tbl.(IssueTable).Constraints() {
		buf.WriteString(",\n ")
		buf.WriteString(c.ConstraintSql(tbl.Dialect().Quoter(), tbl.Name(), i+1))
	}

	buf.WriteString("\n)")
	buf.WriteString(tbl.Dialect().CreateTableSettings())
	return buf.String()
}

func ternaryIssueTable(flag bool, a, b string) string {
	if flag {
		return a
	}
	return b
}

// DropTable drops the table, destroying all its data.
func (tbl IssueTable) DropTable(ctx context.Context, ifExists bool) (int64, error) {
	return support.Exec(ctx, tbl, nil, dropIssueTableSql(tbl, ifExists))
}

func dropIssueTableSql(tbl IssueTabler, ifExists bool) string {
	ie := ternaryIssueTable(ifExists, "IF EXISTS ", "")
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("DROP TABLE %s%s", ie, quotedName)
	return query
}

//-------------------------------------------------------------------------------------------------

// CreateTableWithIndexes invokes CreateTable then CreateIndexes.
func (tbl IssueTable) CreateTableWithIndexes(ctx context.Context, ifNotExist bool) (err error) {
	_, err = tbl.CreateTable(ctx, ifNotExist)
	if err != nil {
		return err
	}

	return tbl.CreateIndexes(ctx, ifNotExist)
}

// CreateIndexes executes queries that create the indexes needed by the Issue table.
func (tbl IssueTable) CreateIndexes(ctx context.Context, ifNotExist bool) (err error) {

	err = tbl.CreateIssueAssigneeIndex(ctx, ifNotExist)
	if err != nil {
		return err
	}

	return nil
}

// CreateIssueAssigneeIndex creates the issue_assignee index.
func (tbl IssueTable) CreateIssueAssigneeIndex(ctx context.Context, ifNotExist bool) error {
	ine := ternaryIssueTable(ifNotExist && tbl.Dialect().Index() != dialect.MysqlIndex, "IF NOT EXISTS ", "")

	// Mysql does not support 'if not exists' on indexes
	// Workaround: use DropIndex first and ignore an error returned if the index didn't exist.

	if ifNotExist && tbl.Dialect().Index() == dialect.MysqlIndex {
		// low-level no-logging Exec
		tbl.Execer().ExecContext(ctx, dropIssueTableIssueAssigneeSql(tbl, false))
		ine = ""
	}

	_, err := tbl.Exec(ctx, nil, createIssueTableIssueAssigneeSql(tbl, ine))
	return err
}

func createIssueTableIssueAssigneeSql(tbl IssueTabler, ifNotExists string) string {
	indexPrefix := tbl.Name().PrefixWithoutDot()
	id := fmt.Sprintf("%s%s_issue_assignee", indexPrefix, tbl.Name().Name)
	q := tbl.Dialect().Quoter()
	cols := strings.Join(q.QuoteN(listOfIssueAssigneeIndexColumns), ",")
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	return fmt.Sprintf("CREATE INDEX %s%s ON %s (%s)", ifNotExists,
		q.Quote(id), quotedName, cols)
}

// DropIssueAssigneeIndex drops the issue_assignee index.
func (tbl IssueTable) DropIssueAssigneeIndex(ctx context.Context, ifExists bool) error {
	_, err := tbl.Exec(ctx, nil, dropIssueTableIssueAssigneeSql(tbl, ifExists))
	return err
}

func dropIssueTableIssueAssigneeSql(tbl IssueTabler, ifExists bool) string {
	// Mysql does not support 'if exists' on indexes
	ie := ternaryIssueTable(ifExists && tbl.Dialect().Index() != dialect.MysqlIndex, "IF EXISTS ", "")
	indexPrefix := tbl.Name().PrefixWithoutDot()
	id := fmt.Sprintf("%s%s_issue_assignee", indexPrefix, tbl.Name().Name)
	q := tbl.Dialect().Quoter()
	// Mysql requires extra "ON tbl" clause
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	onTbl := ternaryIssueTable(tbl.Dialect().Index() == dialect.MysqlIndex, fmt.Sprintf(" ON %s", quotedName), "")
	return "DROP INDEX " + ie + q.Quote(id) + onTbl
}

// DropIndexes executes queries that drop the indexes on by the Issue table.
func (tbl IssueTable) DropIndexes(ctx context.Context, ifExist bool) (err error) {

	err = tbl.DropIssueAssigneeIndex(ctx, ifExist)
	if err != nil {
		return err
	}

	return nil
}

//-------------------------------------------------------------------------------------------------

// Truncate drops every record from the table, if possible. It might fail if constraints exist that
// prevent some or all rows from being deleted; use the force option to override this.
//
// When 'force' is set true, be aware of the following consequences.
// When using Mysql, foreign keys in other tables can be left dangling.
// When using Postgres, a cascade happens, so all 'adjacent' tables (i.e. linked by foreign keys)
// are also truncated.
func (tbl IssueTable) Truncate(ctx context.Context, force bool) (err error) {
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
func (tbl IssueTable) Exec(ctx context.Context, req require.Requirement, query string, args ...interface{}) (int64, error) {
	return support.Exec(ctx, tbl, req, query, args...)
}

//-------------------------------------------------------------------------------------------------

// Query is the low-level request method for this table. The SQL query must return all the columns necessary for
// Issue values. Placeholders should be vanilla '?' marks, which will be replaced if necessary according to
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
func (tbl IssueTable) Query(ctx context.Context, req require.Requirement, query string, args ...interface{}) ([]*Issue, error) {
	return doIssueTableQueryAndScan(ctx, tbl, req, false, query, args)
}

func doIssueTableQueryAndScan(ctx context.Context, tbl IssueTabler, req require.Requirement, firstOnly bool, query string, args ...interface{}) ([]*Issue, error) {
	rows, err := support.Query(ctx, tbl, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vv, n, err := ScanIssues(query, rows, firstOnly)
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
func (tbl IssueTable) QueryOneNullString(ctx context.Context, req require.Requirement, query string, args ...interface{}) (result sql.NullString, err error) {
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
func (tbl IssueTable) QueryOneNullInt64(ctx context.Context, req require.Requirement, query string, args ...interface{}) (result sql.NullInt64, err error) {
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
func (tbl IssueTable) QueryOneNullFloat64(ctx context.Context, req require.Requirement, query string, args ...interface{}) (result sql.NullFloat64, err error) {
	err = support.QueryOneNullThing(ctx, tbl, req, &result, query, args...)
	return result, err
}

// ScanIssues reads rows from the database and returns a slice of corresponding values.
// It also returns a number indicating how many rows were read; this will be larger than the length of the
// slice if reading stopped after the first row.
func ScanIssues(query string, rows sqlapi.SqlRows, firstOnly bool) (vv []*Issue, n int64, err error) {
	for rows.Next() {
		n++

		var v0 int64
		var v1 int
		var v2 Date
		var v3 string
		var v4 string
		var v5 string
		var v6 string
		var v7 []byte

		err = rows.Scan(
			&v0,
			&v1,
			&v2,
			&v3,
			&v4,
			&v5,
			&v6,
			&v7,
		)
		if err != nil {
			return vv, n, errors.Wrap(err, query)
		}

		v := &Issue{}
		v.Id = v0
		v.Number = v1
		v.Date = v2
		v.Title = v3
		v.Body = v4
		v.Assignee = v5
		v.State = v6
		err = json.Unmarshal(v7, &v.Labels)
		if err != nil {
			return nil, n, errors.Wrap(err, query)
		}

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

func allIssueColumnNamesQuoted(q quote.Quoter) string {
	return strings.Join(q.QuoteN(listOfIssueTableColumnNames), ",")
}

//--------------------------------------------------------------------------------

// GetIssueById gets the record with a given primary key value.
// If not found, *Issue will be nil.
func (tbl IssueTable) GetIssueById(ctx context.Context, req require.Requirement, id int64) (*Issue, error) {
	return tbl.SelectOne(ctx, req, where.Eq("id", id), nil)
}

// GetIssuesById gets records from the table according to a list of primary keys.
// Although the list of ids can be arbitrarily long, there are practical limits;
// note that Oracle DB has a limit of 1000.
//
// It places a requirement, which may be nil, on the size of the expected results: in particular, require.All
// controls whether an error is generated not all the ids produce a result.
func (tbl IssueTable) GetIssuesById(ctx context.Context, req require.Requirement, qc where.QueryConstraint, id ...int64) (list []*Issue, err error) {
	if req == require.All {
		req = require.Exactly(len(id))
	}
	return tbl.Select(ctx, req, where.In("id", id), qc)
}

// GetIssuesByAssignee gets the records with a given assignee value.
// If not found, the resulting slice will be empty (nil).
func (tbl IssueTable) GetIssuesByAssignee(ctx context.Context, req require.Requirement, qc where.QueryConstraint, assignee string) ([]*Issue, error) {
	return tbl.Select(ctx, req, where.And(where.Eq("assignee", assignee)), qc)
}

func doIssueTableQueryAndScanOne(ctx context.Context, tbl IssueTabler, req require.Requirement, query string, args ...interface{}) (*Issue, error) {
	list, err := doIssueTableQueryAndScan(ctx, tbl, req, true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

// Fetch fetches a list of Issue based on a supplied query. This is mostly used for join queries that map its
// result columns to the fields of Issue. Other queries might be better handled by GetXxx or Select methods.
func (tbl IssueTable) Fetch(ctx context.Context, req require.Requirement, query string, args ...interface{}) ([]*Issue, error) {
	return doIssueTableQueryAndScan(ctx, tbl, req, false, query, args...)
}

//-------------------------------------------------------------------------------------------------

// SelectOneWhere allows a single Issue to be obtained from the table that matches a 'where' clause
// and some limit. Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
// If not found, *Example will be nil.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.One
// controls whether an error is generated when no result is found.
//
// The args are for any placeholder parameters in the query.
func (tbl IssueTable) SelectOneWhere(ctx context.Context, req require.Requirement, where, orderBy string, args ...interface{}) (*Issue, error) {
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT 1",
		allIssueColumnNamesQuoted(tbl.Dialect().Quoter()), quotedName, where, orderBy)
	v, err := doIssueTableQueryAndScanOne(ctx, tbl, req, query, args...)
	return v, err
}

// SelectOne allows a single Issue to be obtained from the table that matches a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
// If not found, *Example will be nil.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.One
// controls whether an error is generated when no result is found.
func (tbl IssueTable) SelectOne(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) (*Issue, error) {
	q := tbl.Dialect().Quoter()
	whs, args := where.Where(wh, q)
	orderBy := where.Build(qc, q)
	return tbl.SelectOneWhere(ctx, req, whs, orderBy, args...)
}

// SelectWhere allows Issues to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
//
// The args are for any placeholder parameters in the query.
func (tbl IssueTable) SelectWhere(ctx context.Context, req require.Requirement, where, orderBy string, args ...interface{}) ([]*Issue, error) {
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("SELECT %s FROM %s %s %s",
		allIssueColumnNamesQuoted(tbl.Dialect().Quoter()), quotedName, where, orderBy)
	vv, err := doIssueTableQueryAndScan(ctx, tbl, req, false, query, args...)
	return vv, err
}

// Select allows Issues to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
func (tbl IssueTable) Select(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]*Issue, error) {
	q := tbl.Dialect().Quoter()
	whs, args := where.Where(wh, q)
	orderBy := where.Build(qc, q)
	return tbl.SelectWhere(ctx, req, whs, orderBy, args...)
}

//-------------------------------------------------------------------------------------------------

// CountWhere counts Issues in the table that match a 'where' clause.
// Use a blank string for the 'where' argument if it is not needed.
//
// The args are for any placeholder parameters in the query.
func (tbl IssueTable) CountWhere(ctx context.Context, where string, args ...interface{}) (count int64, err error) {
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

// Count counts the Issues in the table that match a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed.
func (tbl IssueTable) Count(ctx context.Context, wh where.Expression) (count int64, err error) {
	whs, args := where.Where(wh, tbl.Dialect().Quoter())
	return tbl.CountWhere(ctx, whs, args...)
}

//--------------------------------------------------------------------------------

// SliceId gets the id column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl IssueTable) SliceId(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return support.SliceInt64List(ctx, tbl, req, tbl.pk, wh, qc)
}

// SliceNumber gets the number column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl IssueTable) SliceNumber(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int, error) {
	return support.SliceIntList(ctx, tbl, req, "number", wh, qc)
}

// SliceTitle gets the title column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl IssueTable) SliceTitle(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return support.SliceStringList(ctx, tbl, req, "title", wh, qc)
}

// SliceBigbody gets the bigbody column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl IssueTable) SliceBigbody(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return support.SliceStringList(ctx, tbl, req, "bigbody", wh, qc)
}

// SliceAssignee gets the assignee column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl IssueTable) SliceAssignee(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return support.SliceStringList(ctx, tbl, req, "assignee", wh, qc)
}

// SliceState gets the state column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl IssueTable) SliceState(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return support.SliceStringList(ctx, tbl, req, "state", wh, qc)
}

func constructIssueTableInsert(tbl IssueTable, w dialect.StringWriter, v *Issue, withPk bool) (s []interface{}, err error) {
	q := tbl.Dialect().Quoter()
	s = make([]interface{}, 0, 8)

	comma := ""
	w.WriteString(" (")

	if withPk {
		q.QuoteW(w, "id")
		comma = ","
		s = append(s, v.Id)
	}

	w.WriteString(comma)
	q.QuoteW(w, "number")
	s = append(s, v.Number)
	comma = ","

	w.WriteString(comma)
	q.QuoteW(w, "date")
	s = append(s, v.Date)

	w.WriteString(comma)
	q.QuoteW(w, "title")
	s = append(s, v.Title)

	w.WriteString(comma)
	q.QuoteW(w, "bigbody")
	s = append(s, v.Body)

	w.WriteString(comma)
	q.QuoteW(w, "assignee")
	s = append(s, v.Assignee)

	w.WriteString(comma)
	q.QuoteW(w, "state")
	s = append(s, v.State)

	w.WriteString(comma)
	q.QuoteW(w, "labels")
	x, err := json.Marshal(&v.Labels)
	if err != nil {
		return nil, tbl.Logger().LogError(errors.WithStack(err))
	}
	s = append(s, x)

	w.WriteString(")")
	return s, nil
}

func constructIssueTableUpdate(tbl IssueTable, w dialect.StringWriter, v *Issue) (s []interface{}, err error) {
	q := tbl.Dialect().Quoter()
	j := 1
	s = make([]interface{}, 0, 7)

	comma := ""

	w.WriteString(comma)
	q.QuoteW(w, "number")
	w.WriteString("=?")
	s = append(s, v.Number)
	j++
	comma = ", "

	w.WriteString(comma)
	q.QuoteW(w, "date")
	w.WriteString("=?")
	s = append(s, v.Date)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "title")
	w.WriteString("=?")
	s = append(s, v.Title)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "bigbody")
	w.WriteString("=?")
	s = append(s, v.Body)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "assignee")
	w.WriteString("=?")
	s = append(s, v.Assignee)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "state")
	w.WriteString("=?")
	s = append(s, v.State)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "labels")
	w.WriteString("=?")
	j++

	x, err := json.Marshal(&v.Labels)
	if err != nil {
		return nil, tbl.Logger().LogError(errors.WithStack(err))
	}
	s = append(s, x)
	return s, nil
}

//--------------------------------------------------------------------------------

// Insert adds new records for the Issues.// The Issues have their primary key fields set to the new record identifiers.
// The Issue.PreInsert() method will be called, if it exists.
func (tbl IssueTable) Insert(ctx context.Context, req require.Requirement, vv ...*Issue) error {
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

		fields, err := constructIssueTableInsert(tbl, b, v, false)
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
			v.Id = i64

		} else {
			i64, e2 := tbl.db.InsertContext(ctx, tbl.pk, query, fields...)
			if e2 != nil {
				return tbl.Logger().LogError(e2)
			}
			v.Id = i64
		}

		if err != nil {
			return tbl.Logger().LogError(err)
		}
		count += n
	}

	return tbl.Logger().LogIfError(require.ErrorIfExecNotSatisfiedBy(req, count))
}

// UpdateById updates one or more columns, given a id value.
func (tbl IssueTable) UpdateById(ctx context.Context, req require.Requirement, id int64, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("id", id), fields...)
}

// UpdateByNumber updates one or more columns, given a number value.
func (tbl IssueTable) UpdateByNumber(ctx context.Context, req require.Requirement, number int, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("number", number), fields...)
}

// UpdateByTitle updates one or more columns, given a title value.
func (tbl IssueTable) UpdateByTitle(ctx context.Context, req require.Requirement, title string, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("title", title), fields...)
}

// UpdateByBigbody updates one or more columns, given a bigbody value.
func (tbl IssueTable) UpdateByBigbody(ctx context.Context, req require.Requirement, bigbody string, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("bigbody", bigbody), fields...)
}

// UpdateByAssignee updates one or more columns, given a assignee value.
func (tbl IssueTable) UpdateByAssignee(ctx context.Context, req require.Requirement, assignee string, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("assignee", assignee), fields...)
}

// UpdateByState updates one or more columns, given a state value.
func (tbl IssueTable) UpdateByState(ctx context.Context, req require.Requirement, state string, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("state", state), fields...)
}

// UpdateFields updates one or more columns, given a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed (but note that this is risky!).
func (tbl IssueTable) UpdateFields(ctx context.Context, req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error) {
	return support.UpdateFields(ctx, tbl, req, wh, fields...)
}

//--------------------------------------------------------------------------------

// Update updates records, matching them by primary key. It returns the number of rows affected.
// The Issue.PreUpdate(Execer) method will be called, if it exists.
func (tbl IssueTable) Update(ctx context.Context, req require.Requirement, vv ...*Issue) (int64, error) {
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

		args, err := constructIssueTableUpdate(tbl, b, v)
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
func (tbl IssueTable) Upsert(ctx context.Context, v *Issue, wh where.Expression) error {
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

	var id int64
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
func (tbl IssueTable) DeleteById(ctx context.Context, req require.Requirement, id ...int64) (int64, error) {
	ii := support.Int64AsInterfaceSlice(id)
	return support.DeleteByColumn(ctx, tbl, req, "id", ii...)
}

// DeleteByNumber deletes rows from the table, given some number values.
// The list of ids can be arbitrarily long.
func (tbl IssueTable) DeleteByNumber(ctx context.Context, req require.Requirement, number ...int) (int64, error) {
	ii := support.IntAsInterfaceSlice(number)
	return support.DeleteByColumn(ctx, tbl, req, "number", ii...)
}

// DeleteByTitle deletes rows from the table, given some title values.
// The list of ids can be arbitrarily long.
func (tbl IssueTable) DeleteByTitle(ctx context.Context, req require.Requirement, title ...string) (int64, error) {
	ii := support.StringAsInterfaceSlice(title)
	return support.DeleteByColumn(ctx, tbl, req, "title", ii...)
}

// DeleteByBigbody deletes rows from the table, given some bigbody values.
// The list of ids can be arbitrarily long.
func (tbl IssueTable) DeleteByBigbody(ctx context.Context, req require.Requirement, bigbody ...string) (int64, error) {
	ii := support.StringAsInterfaceSlice(bigbody)
	return support.DeleteByColumn(ctx, tbl, req, "bigbody", ii...)
}

// DeleteByAssignee deletes rows from the table, given some assignee values.
// The list of ids can be arbitrarily long.
func (tbl IssueTable) DeleteByAssignee(ctx context.Context, req require.Requirement, assignee ...string) (int64, error) {
	ii := support.StringAsInterfaceSlice(assignee)
	return support.DeleteByColumn(ctx, tbl, req, "assignee", ii...)
}

// DeleteByState deletes rows from the table, given some state values.
// The list of ids can be arbitrarily long.
func (tbl IssueTable) DeleteByState(ctx context.Context, req require.Requirement, state ...string) (int64, error) {
	ii := support.StringAsInterfaceSlice(state)
	return support.DeleteByColumn(ctx, tbl, req, "state", ii...)
}

// Delete deletes one or more rows from the table, given a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl IssueTable) Delete(ctx context.Context, req require.Requirement, wh where.Expression) (int64, error) {
	query, args := deleteRowsIssueTableSql(tbl, wh)
	return tbl.Exec(ctx, req, query, args...)
}

func deleteRowsIssueTableSql(tbl IssueTabler, wh where.Expression) (string, []interface{}) {
	whs, args := where.Where(wh, tbl.Dialect().Quoter())
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("DELETE FROM %s %s", quotedName, whs)
	return query, args
}

//-------------------------------------------------------------------------------------------------

//-------------------------------------------------------------------------------------------------

// SetId sets the Id field and returns the modified Issue.
func (v *Issue) SetId(x int64) *Issue {
	v.Id = x
	return v
}

// SetNumber sets the Number field and returns the modified Issue.
func (v *Issue) SetNumber(x int) *Issue {
	v.Number = x
	return v
}

// SetDate sets the Date field and returns the modified Issue.
func (v *Issue) SetDate(x Date) *Issue {
	v.Date = x
	return v
}

// SetTitle sets the Title field and returns the modified Issue.
func (v *Issue) SetTitle(x string) *Issue {
	v.Title = x
	return v
}

// SetBody sets the Body field and returns the modified Issue.
func (v *Issue) SetBody(x string) *Issue {
	v.Body = x
	return v
}

// SetAssignee sets the Assignee field and returns the modified Issue.
func (v *Issue) SetAssignee(x string) *Issue {
	v.Assignee = x
	return v
}

// SetState sets the State field and returns the modified Issue.
func (v *Issue) SetState(x string) *Issue {
	v.State = x
	return v
}

// SetLabels sets the Labels field and returns the modified Issue.
func (v *Issue) SetLabels(x []string) *Issue {
	v.Labels = x
	return v
}
