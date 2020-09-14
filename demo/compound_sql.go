// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.
// sqlapi v0.47.1; sqlgen v0.67.0

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

// DbCompoundTabler lists table methods provided by DbCompoundTable.
type DbCompoundTabler interface {
	sqlapi.Table

	// Constraints returns the table's constraints.
	Constraints() constraint.Constraints

	// WithConstraint returns a modified DbCompoundTabler with added data consistency constraints.
	WithConstraint(cc ...constraint.Constraint) DbCompoundTabler

	// WithPrefix returns a modified DbCompoundTabler with a given table name prefix.
	WithPrefix(pfx string) DbCompoundTabler

	// CreateTable creates the table.
	CreateTable(ctx context.Context, ifNotExists bool) (int64, error)

	// DropTable drops the table, destroying all its data.
	DropTable(ctx context.Context, ifExists bool) (int64, error)

	// CreateTableWithIndexes invokes CreateTable then CreateIndexes.
	CreateTableWithIndexes(ctx context.Context, ifNotExist bool) (err error)

	// CreateIndexes executes queries that create the indexes needed by the Compound table.
	CreateIndexes(ctx context.Context, ifNotExist bool) (err error)

	// CreateAlphaBetaIndex creates the alpha_beta index.
	CreateAlphaBetaIndex(ctx context.Context, ifNotExist bool) error

	// DropAlphaBetaIndex drops the alpha_beta index.
	DropAlphaBetaIndex(ctx context.Context, ifExists bool) error

	// Truncate drops every record from the table, if possible.
	Truncate(ctx context.Context, force bool) (err error)
}

//-------------------------------------------------------------------------------------------------

// DbCompoundQueryer lists query methods provided by DbCompoundTable.
type DbCompoundQueryer interface {
	sqlapi.Table

	// Using returns a modified DbCompoundQueryer using the Execer supplied,
	// which will typically be a transaction (i.e. SqlTx).
	Using(tx sqlapi.Execer) DbCompoundQueryer

	// Transact runs the function provided within a transaction. The transction is committed
	// unless an error occurs.
	Transact(ctx context.Context, txOptions *sql.TxOptions, fn func(DbCompoundQueryer) error) error

	// Exec executes a query without returning any rows.
	Exec(ctx context.Context, req require.Requirement, query string, args ...interface{}) (int64, error)

	// Query is the low-level request method for this table using an SQL query that must return all the columns
	// necessary for Compound values.
	Query(ctx context.Context, req require.Requirement, query string, args ...interface{}) ([]*Compound, error)

	// QueryOneNullString is a low-level access method for one string, returning the first match.
	QueryOneNullString(ctx context.Context, req require.Requirement, query string, args ...interface{}) (result sql.NullString, err error)

	// QueryOneNullInt64 is a low-level access method for one int64, returning the first match.
	QueryOneNullInt64(ctx context.Context, req require.Requirement, query string, args ...interface{}) (result sql.NullInt64, err error)

	// QueryOneNullFloat64 is a low-level access method for one float64, returning the first match.
	QueryOneNullFloat64(ctx context.Context, req require.Requirement, query string, args ...interface{}) (result sql.NullFloat64, err error)

	// GetCompoundByAlphaAndBeta gets the record with given alpha+beta values.
	GetCompoundByAlphaAndBeta(ctx context.Context, req require.Requirement, alpha string, beta string) (*Compound, error)

	// Fetch fetches a list of Compound based on a supplied query. This is mostly used for join queries that map its
	// result columns to the fields of Compound. Other queries might be better handled by GetXxx or Select methods.
	Fetch(ctx context.Context, req require.Requirement, query string, args ...interface{}) ([]*Compound, error)

	// SelectOneWhere allows a single Compound to be obtained from the table that matches a 'where' clause.
	SelectOneWhere(ctx context.Context, req require.Requirement, where, orderBy string, args ...interface{}) (*Compound, error)

	// SelectOne allows a single Compound to be obtained from the table that matches a 'where' clause.
	SelectOne(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) (*Compound, error)

	// SelectWhere allows Compounds to be obtained from the table that match a 'where' clause.
	SelectWhere(ctx context.Context, req require.Requirement, where, orderBy string, args ...interface{}) ([]*Compound, error)

	// Select allows Compounds to be obtained from the table that match a 'where' clause.
	Select(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]*Compound, error)

	// CountWhere counts Compounds in the table that match a 'where' clause.
	CountWhere(ctx context.Context, where string, args ...interface{}) (count int64, err error)

	// Count counts the Compounds in the table that match a 'where' clause.
	Count(ctx context.Context, wh where.Expression) (count int64, err error)

	// SliceAlpha gets the alpha column for all rows that match the 'where' condition.
	SliceAlpha(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error)

	// SliceBeta gets the beta column for all rows that match the 'where' condition.
	SliceBeta(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error)

	// SliceCategory gets the category column for all rows that match the 'where' condition.
	SliceCategory(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]Category, error)

	// Insert adds new records for the Compounds.
	Insert(ctx context.Context, req require.Requirement, vv ...*Compound) error

	// UpdateByAlpha updates one or more columns, given a alpha value.
	UpdateByAlpha(ctx context.Context, req require.Requirement, alpha string, fields ...sql.NamedArg) (int64, error)

	// UpdateByBeta updates one or more columns, given a beta value.
	UpdateByBeta(ctx context.Context, req require.Requirement, beta string, fields ...sql.NamedArg) (int64, error)

	// UpdateByCategory updates one or more columns, given a category value.
	UpdateByCategory(ctx context.Context, req require.Requirement, category Category, fields ...sql.NamedArg) (int64, error)

	// UpdateFields updates one or more columns, given a 'where' clause.
	UpdateFields(ctx context.Context, req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error)

	// DeleteByAlpha deletes rows from the table, given some alpha values.
	// The list of ids can be arbitrarily long.
	DeleteByAlpha(ctx context.Context, req require.Requirement, alpha ...string) (int64, error)

	// DeleteByBeta deletes rows from the table, given some beta values.
	// The list of ids can be arbitrarily long.
	DeleteByBeta(ctx context.Context, req require.Requirement, beta ...string) (int64, error)

	// DeleteByCategory deletes rows from the table, given some category values.
	// The list of ids can be arbitrarily long.
	DeleteByCategory(ctx context.Context, req require.Requirement, category ...Category) (int64, error)

	// Delete deletes one or more rows from the table, given a 'where' clause.
	// Use a nil value for the 'wh' argument if it is not needed (very risky!).
	Delete(ctx context.Context, req require.Requirement, wh where.Expression) (int64, error)
}

//-------------------------------------------------------------------------------------------------

// DbCompoundTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type DbCompoundTable struct {
	name        sqlapi.TableName
	database    sqlapi.Database
	db          sqlapi.Execer
	constraints constraint.Constraints
	pk          string
}

// Type conformance checks
var _ sqlapi.TableWithIndexes = &DbCompoundTable{}

// NewDbCompoundTable returns a new table instance.
// If a blank table name is supplied, the default name "compounds" will be used instead.
// The request context is initialised with the background.
func NewDbCompoundTable(name string, d sqlapi.Database) DbCompoundTable {
	if name == "" {
		name = "compounds"
	}
	var constraints constraint.Constraints
	return DbCompoundTable{
		name:        sqlapi.TableName{Prefix: "", Name: name},
		database:    d,
		db:          d.DB(),
		constraints: constraints,
		pk:          "",
	}
}

// CopyTableAsDbCompoundTable copies a table instance, retaining the name etc but
// providing methods appropriate for 'Compound'.It doesn't copy the constraints of the original table.
//
// It serves to provide methods appropriate for 'Compound'. This is most useful when this is used to represent a
// join result. In such cases, there won't be any need for DDL methods, nor Exec, Insert, Update or Delete.
func CopyTableAsDbCompoundTable(origin sqlapi.Table) DbCompoundTable {
	return DbCompoundTable{
		name:        origin.Name(),
		database:    origin.Database(),
		db:          origin.Execer(),
		constraints: nil,
		pk:          "",
	}
}

// WithPrefix sets the table name prefix for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl DbCompoundTable) WithPrefix(pfx string) DbCompoundTabler {
	tbl.name.Prefix = pfx
	return tbl
}

// Database gets the shared database information.
func (tbl DbCompoundTable) Database() sqlapi.Database {
	return tbl.database
}

// Logger gets the trace logger.
func (tbl DbCompoundTable) Logger() sqlapi.Logger {
	return tbl.database.Logger()
}

// WithConstraint returns a modified DbCompoundTabler with added data consistency constraints.
func (tbl DbCompoundTable) WithConstraint(cc ...constraint.Constraint) DbCompoundTabler {
	tbl.constraints = append(tbl.constraints, cc...)
	return tbl
}

// Constraints returns the table's constraints.
func (tbl DbCompoundTable) Constraints() constraint.Constraints {
	return tbl.constraints
}

// Dialect gets the database dialect.
func (tbl DbCompoundTable) Dialect() dialect.Dialect {
	return tbl.database.Dialect()
}

// Name gets the table name.
func (tbl DbCompoundTable) Name() sqlapi.TableName {
	return tbl.name
}

// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl DbCompoundTable) DB() sqlapi.SqlDB {
	return tbl.db.(sqlapi.SqlDB)
}

// Execer gets the wrapped database or transaction handle.
func (tbl DbCompoundTable) Execer() sqlapi.Execer {
	return tbl.db
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl DbCompoundTable) Tx() sqlapi.SqlTx {
	return tbl.db.(sqlapi.SqlTx)
}

// IsTx tests whether this is within a transaction.
func (tbl DbCompoundTable) IsTx() bool {
	return tbl.db.IsTx()
}

// Using returns a modified DbCompoundTabler using the the Execer supplied,
// which will typically be a transaction (i.e. SqlTx). This is needed when making multiple
// queries across several tables within a single transaction.
//
// The result is a modified copy of the table; the original is unchanged.
func (tbl DbCompoundTable) Using(tx sqlapi.Execer) DbCompoundQueryer {
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
func (tbl DbCompoundTable) Transact(ctx context.Context, txOptions *sql.TxOptions, fn func(DbCompoundQueryer) error) error {
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

func (tbl DbCompoundTable) quotedName() string {
	return tbl.Dialect().Quoter().Quote(tbl.name.String())
}

func (tbl DbCompoundTable) quotedNameW(w dialect.StringWriter) {
	tbl.Dialect().Quoter().QuoteW(w, tbl.name.String())
}

//-------------------------------------------------------------------------------------------------

// NumDbCompoundTableColumns is the total number of columns in DbCompoundTable.
const NumDbCompoundTableColumns = 3

// NumDbCompoundTableDataColumns is the number of columns in DbCompoundTable not including the auto-increment key.
const NumDbCompoundTableDataColumns = 3

// DbCompoundTableColumnNames is the list of columns in DbCompoundTable.
const DbCompoundTableColumnNames = "alpha,beta,category"

var listOfDbCompoundTableColumnNames = strings.Split(DbCompoundTableColumnNames, ",")

//-------------------------------------------------------------------------------------------------

var sqlDbCompoundTableCreateColumnsSqlite = []string{
	"text not null",
	"text not null",
	"tinyint unsigned not null",
}

var sqlDbCompoundTableCreateColumnsMysql = []string{
	"varchar(255) not null",
	"varchar(255) not null",
	"tinyint unsigned not null",
}

var sqlDbCompoundTableCreateColumnsPostgres = []string{
	"text not null",
	"text not null",
	"smallint not null",
}

var sqlDbCompoundTableCreateColumnsPgx = []string{
	"text not null",
	"text not null",
	"smallint not null",
}

//-------------------------------------------------------------------------------------------------

const sqlDbAlphaBetaIndexColumns = "alpha,beta"

var listOfDbAlphaBetaIndexColumns = strings.Split(sqlDbAlphaBetaIndexColumns, ",")

//-------------------------------------------------------------------------------------------------

// CreateTable creates the table.
func (tbl DbCompoundTable) CreateTable(ctx context.Context, ifNotExists bool) (int64, error) {
	return support.Exec(ctx, tbl, nil, createDbCompoundTableSql(tbl, ifNotExists))
}

func createDbCompoundTableSql(tbl DbCompoundTabler, ifNotExists bool) string {
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
		columns = sqlDbCompoundTableCreateColumnsSqlite
	case dialect.MysqlIndex:
		columns = sqlDbCompoundTableCreateColumnsMysql
	case dialect.PostgresIndex:
		columns = sqlDbCompoundTableCreateColumnsPostgres
	case dialect.PgxIndex:
		columns = sqlDbCompoundTableCreateColumnsPgx
	}

	comma := ""
	for i, n := range listOfDbCompoundTableColumnNames {
		buf.WriteString(comma)
		q.QuoteW(buf, n)
		buf.WriteString(" ")
		buf.WriteString(columns[i])
		comma = ",\n "
	}

	for i, c := range tbl.(DbCompoundTable).Constraints() {
		buf.WriteString(",\n ")
		buf.WriteString(c.ConstraintSql(tbl.Dialect().Quoter(), tbl.Name(), i+1))
	}

	buf.WriteString("\n)")
	buf.WriteString(tbl.Dialect().CreateTableSettings())
	return buf.String()
}

func ternaryDbCompoundTable(flag bool, a, b string) string {
	if flag {
		return a
	}
	return b
}

// DropTable drops the table, destroying all its data.
func (tbl DbCompoundTable) DropTable(ctx context.Context, ifExists bool) (int64, error) {
	return support.Exec(ctx, tbl, nil, dropDbCompoundTableSql(tbl, ifExists))
}

func dropDbCompoundTableSql(tbl DbCompoundTabler, ifExists bool) string {
	ie := ternaryDbCompoundTable(ifExists, "IF EXISTS ", "")
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("DROP TABLE %s%s", ie, quotedName)
	return query
}

//-------------------------------------------------------------------------------------------------

// CreateTableWithIndexes invokes CreateTable then CreateIndexes.
func (tbl DbCompoundTable) CreateTableWithIndexes(ctx context.Context, ifNotExist bool) (err error) {
	_, err = tbl.CreateTable(ctx, ifNotExist)
	if err != nil {
		return err
	}

	return tbl.CreateIndexes(ctx, ifNotExist)
}

// CreateIndexes executes queries that create the indexes needed by the Compound table.
func (tbl DbCompoundTable) CreateIndexes(ctx context.Context, ifNotExist bool) (err error) {

	err = tbl.CreateAlphaBetaIndex(ctx, ifNotExist)
	if err != nil {
		return err
	}

	return nil
}

// CreateAlphaBetaIndex creates the alpha_beta index.
func (tbl DbCompoundTable) CreateAlphaBetaIndex(ctx context.Context, ifNotExist bool) error {
	ine := ternaryDbCompoundTable(ifNotExist && tbl.Dialect().Index() != dialect.MysqlIndex, "IF NOT EXISTS ", "")

	// Mysql does not support 'if not exists' on indexes
	// Workaround: use DropIndex first and ignore an error returned if the index didn't exist.

	if ifNotExist && tbl.Dialect().Index() == dialect.MysqlIndex {
		// low-level no-logging Exec
		tbl.Execer().ExecContext(ctx, dropDbCompoundTableAlphaBetaSql(tbl, false))
		ine = ""
	}

	_, err := tbl.Exec(ctx, nil, createDbCompoundTableAlphaBetaSql(tbl, ine))
	return err
}

func createDbCompoundTableAlphaBetaSql(tbl DbCompoundTabler, ifNotExists string) string {
	indexPrefix := tbl.Name().PrefixWithoutDot()
	id := fmt.Sprintf("%s%s_alpha_beta", indexPrefix, tbl.Name().Name)
	q := tbl.Dialect().Quoter()
	cols := strings.Join(q.QuoteN(listOfDbAlphaBetaIndexColumns), ",")
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	return fmt.Sprintf("CREATE UNIQUE INDEX %s%s ON %s (%s)", ifNotExists,
		q.Quote(id), quotedName, cols)
}

// DropAlphaBetaIndex drops the alpha_beta index.
func (tbl DbCompoundTable) DropAlphaBetaIndex(ctx context.Context, ifExists bool) error {
	_, err := tbl.Exec(ctx, nil, dropDbCompoundTableAlphaBetaSql(tbl, ifExists))
	return err
}

func dropDbCompoundTableAlphaBetaSql(tbl DbCompoundTabler, ifExists bool) string {
	// Mysql does not support 'if exists' on indexes
	ie := ternaryDbCompoundTable(ifExists && tbl.Dialect().Index() != dialect.MysqlIndex, "IF EXISTS ", "")
	indexPrefix := tbl.Name().PrefixWithoutDot()
	id := fmt.Sprintf("%s%s_alpha_beta", indexPrefix, tbl.Name().Name)
	q := tbl.Dialect().Quoter()
	// Mysql requires extra "ON tbl" clause
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	onTbl := ternaryDbCompoundTable(tbl.Dialect().Index() == dialect.MysqlIndex, fmt.Sprintf(" ON %s", quotedName), "")
	return "DROP INDEX " + ie + q.Quote(id) + onTbl
}

// DropIndexes executes queries that drop the indexes on by the Compound table.
func (tbl DbCompoundTable) DropIndexes(ctx context.Context, ifExist bool) (err error) {

	err = tbl.DropAlphaBetaIndex(ctx, ifExist)
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
func (tbl DbCompoundTable) Truncate(ctx context.Context, force bool) (err error) {
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
func (tbl DbCompoundTable) Exec(ctx context.Context, req require.Requirement, query string, args ...interface{}) (int64, error) {
	return support.Exec(ctx, tbl, req, query, args...)
}

//-------------------------------------------------------------------------------------------------

// Query is the low-level request method for this table. The SQL query must return all the columns necessary for
// Compound values. Placeholders should be vanilla '?' marks, which will be replaced if necessary according to
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
func (tbl DbCompoundTable) Query(ctx context.Context, req require.Requirement, query string, args ...interface{}) ([]*Compound, error) {
	return doDbCompoundTableQueryAndScan(ctx, tbl, req, false, query, args)
}

func doDbCompoundTableQueryAndScan(ctx context.Context, tbl DbCompoundTabler, req require.Requirement, firstOnly bool, query string, args ...interface{}) ([]*Compound, error) {
	rows, err := support.Query(ctx, tbl, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vv, n, err := ScanDbCompounds(query, rows, firstOnly)
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
func (tbl DbCompoundTable) QueryOneNullString(ctx context.Context, req require.Requirement, query string, args ...interface{}) (result sql.NullString, err error) {
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
func (tbl DbCompoundTable) QueryOneNullInt64(ctx context.Context, req require.Requirement, query string, args ...interface{}) (result sql.NullInt64, err error) {
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
func (tbl DbCompoundTable) QueryOneNullFloat64(ctx context.Context, req require.Requirement, query string, args ...interface{}) (result sql.NullFloat64, err error) {
	err = support.QueryOneNullThing(ctx, tbl, req, &result, query, args...)
	return result, err
}

// ScanDbCompounds reads rows from the database and returns a slice of corresponding values.
// It also returns a number indicating how many rows were read; this will be larger than the length of the
// slice if reading stopped after the first row.
func ScanDbCompounds(query string, rows sqlapi.SqlRows, firstOnly bool) (vv []*Compound, n int64, err error) {
	for rows.Next() {
		n++

		var v0 string
		var v1 string
		var v2 Category

		err = rows.Scan(
			&v0,
			&v1,
			&v2,
		)
		if err != nil {
			return vv, n, errors.Wrap(err, query)
		}

		v := &Compound{}
		v.Alpha = v0
		v.Beta = v1
		v.Category = v2

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

func allDbCompoundColumnNamesQuoted(q quote.Quoter) string {
	return strings.Join(q.QuoteN(listOfDbCompoundTableColumnNames), ",")
}

// GetCompoundByAlphaAndBeta gets the record with given alpha+beta values.
// If not found, *Compound will be nil.
func (tbl DbCompoundTable) GetCompoundByAlphaAndBeta(ctx context.Context, req require.Requirement, alpha string, beta string) (*Compound, error) {
	return tbl.SelectOne(ctx, req, where.And(where.Eq("alpha", alpha), where.Eq("beta", beta)), nil)
}

func doDbCompoundTableQueryAndScanOne(ctx context.Context, tbl DbCompoundTabler, req require.Requirement, query string, args ...interface{}) (*Compound, error) {
	list, err := doDbCompoundTableQueryAndScan(ctx, tbl, req, true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

// Fetch fetches a list of Compound based on a supplied query. This is mostly used for join queries that map its
// result columns to the fields of Compound. Other queries might be better handled by GetXxx or Select methods.
func (tbl DbCompoundTable) Fetch(ctx context.Context, req require.Requirement, query string, args ...interface{}) ([]*Compound, error) {
	return doDbCompoundTableQueryAndScan(ctx, tbl, req, false, query, args...)
}

//-------------------------------------------------------------------------------------------------

// SelectOneWhere allows a single Compound to be obtained from the table that matches a 'where' clause
// and some limit. Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
// If not found, *Example will be nil.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.One
// controls whether an error is generated when no result is found.
//
// The args are for any placeholder parameters in the query.
func (tbl DbCompoundTable) SelectOneWhere(ctx context.Context, req require.Requirement, where, orderBy string, args ...interface{}) (*Compound, error) {
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT 1",
		allDbCompoundColumnNamesQuoted(tbl.Dialect().Quoter()), quotedName, where, orderBy)
	v, err := doDbCompoundTableQueryAndScanOne(ctx, tbl, req, query, args...)
	return v, err
}

// SelectOne allows a single Compound to be obtained from the table that matches a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
// If not found, *Example will be nil.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.One
// controls whether an error is generated when no result is found.
func (tbl DbCompoundTable) SelectOne(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) (*Compound, error) {
	q := tbl.Dialect().Quoter()
	whs, args := where.Where(wh, q)
	orderBy := where.Build(qc, q)
	return tbl.SelectOneWhere(ctx, req, whs, orderBy, args...)
}

// SelectWhere allows Compounds to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
//
// The args are for any placeholder parameters in the query.
func (tbl DbCompoundTable) SelectWhere(ctx context.Context, req require.Requirement, where, orderBy string, args ...interface{}) ([]*Compound, error) {
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("SELECT %s FROM %s %s %s",
		allDbCompoundColumnNamesQuoted(tbl.Dialect().Quoter()), quotedName, where, orderBy)
	vv, err := doDbCompoundTableQueryAndScan(ctx, tbl, req, false, query, args...)
	return vv, err
}

// Select allows Compounds to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
func (tbl DbCompoundTable) Select(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]*Compound, error) {
	q := tbl.Dialect().Quoter()
	whs, args := where.Where(wh, q)
	orderBy := where.Build(qc, q)
	return tbl.SelectWhere(ctx, req, whs, orderBy, args...)
}

//-------------------------------------------------------------------------------------------------

// CountWhere counts Compounds in the table that match a 'where' clause.
// Use a blank string for the 'where' argument if it is not needed.
//
// The args are for any placeholder parameters in the query.
func (tbl DbCompoundTable) CountWhere(ctx context.Context, where string, args ...interface{}) (count int64, err error) {
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

// Count counts the Compounds in the table that match a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed.
func (tbl DbCompoundTable) Count(ctx context.Context, wh where.Expression) (count int64, err error) {
	whs, args := where.Where(wh, tbl.Dialect().Quoter())
	return tbl.CountWhere(ctx, whs, args...)
}

//--------------------------------------------------------------------------------

// SliceAlpha gets the alpha column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DbCompoundTable) SliceAlpha(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return support.SliceStringList(ctx, tbl, req, "alpha", wh, qc)
}

// SliceBeta gets the beta column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DbCompoundTable) SliceBeta(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return support.SliceStringList(ctx, tbl, req, "beta", wh, qc)
}

// SliceCategory gets the category column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DbCompoundTable) SliceCategory(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]Category, error) {
	return sliceDbCompoundTableCategoryList(ctx, tbl, req, "category", wh, qc)
}

func sliceDbCompoundTableCategoryList(ctx context.Context, tbl DbCompoundTabler, req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]Category, error) {
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

func constructDbCompoundTableInsert(tbl DbCompoundTable, w dialect.StringWriter, v *Compound, withPk bool) (s []interface{}, err error) {
	q := tbl.Dialect().Quoter()
	s = make([]interface{}, 0, 3)

	comma := ""
	w.WriteString(" (")

	w.WriteString(comma)
	q.QuoteW(w, "alpha")
	s = append(s, v.Alpha)
	comma = ","

	w.WriteString(comma)
	q.QuoteW(w, "beta")
	s = append(s, v.Beta)

	w.WriteString(comma)
	q.QuoteW(w, "category")
	s = append(s, v.Category)

	w.WriteString(")")
	return s, nil
}

func constructDbCompoundTableUpdate(tbl DbCompoundTable, w dialect.StringWriter, v *Compound) (s []interface{}, err error) {
	q := tbl.Dialect().Quoter()
	j := 1
	s = make([]interface{}, 0, 3)

	comma := ""

	w.WriteString(comma)
	q.QuoteW(w, "alpha")
	w.WriteString("=?")
	s = append(s, v.Alpha)
	j++
	comma = ", "

	w.WriteString(comma)
	q.QuoteW(w, "beta")
	w.WriteString("=?")
	s = append(s, v.Beta)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "category")
	w.WriteString("=?")
	s = append(s, v.Category)
	j++
	return s, nil
}

//--------------------------------------------------------------------------------

// Insert adds new records for the Compounds.
// The Compound.PreInsert() method will be called, if it exists.
func (tbl DbCompoundTable) Insert(ctx context.Context, req require.Requirement, vv ...*Compound) error {
	if req == require.All {
		req = require.Exactly(len(vv))
	}

	if ctx == nil {
		ctx = context.Background()
	}

	var count int64
	returning := ""
	insertHasReturningPhrase := false

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

		fields, err := constructDbCompoundTableInsert(tbl, b, v, true)
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

		} else {
			_, e3 := tbl.db.ExecContext(ctx, query, fields...)
			if e3 != nil {
				return tbl.Logger().LogError(e3)
			}
		}

		if err != nil {
			return tbl.Logger().LogError(err)
		}
		count += n
	}

	return tbl.Logger().LogIfError(require.ErrorIfExecNotSatisfiedBy(req, count))
}

// UpdateByAlpha updates one or more columns, given a alpha value.
func (tbl DbCompoundTable) UpdateByAlpha(ctx context.Context, req require.Requirement, alpha string, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("alpha", alpha), fields...)
}

// UpdateByBeta updates one or more columns, given a beta value.
func (tbl DbCompoundTable) UpdateByBeta(ctx context.Context, req require.Requirement, beta string, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("beta", beta), fields...)
}

// UpdateByCategory updates one or more columns, given a category value.
func (tbl DbCompoundTable) UpdateByCategory(ctx context.Context, req require.Requirement, category Category, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("category", category), fields...)
}

// UpdateFields updates one or more columns, given a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed (but note that this is risky!).
func (tbl DbCompoundTable) UpdateFields(ctx context.Context, req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error) {
	return support.UpdateFields(ctx, tbl, req, wh, fields...)
}

//-------------------------------------------------------------------------------------------------

// DeleteByAlpha deletes rows from the table, given some alpha values.
// The list of ids can be arbitrarily long.
func (tbl DbCompoundTable) DeleteByAlpha(ctx context.Context, req require.Requirement, alpha ...string) (int64, error) {
	ii := support.StringAsInterfaceSlice(alpha)
	return support.DeleteByColumn(ctx, tbl, req, "alpha", ii...)
}

// DeleteByBeta deletes rows from the table, given some beta values.
// The list of ids can be arbitrarily long.
func (tbl DbCompoundTable) DeleteByBeta(ctx context.Context, req require.Requirement, beta ...string) (int64, error) {
	ii := support.StringAsInterfaceSlice(beta)
	return support.DeleteByColumn(ctx, tbl, req, "beta", ii...)
}

// DeleteByCategory deletes rows from the table, given some category values.
// The list of ids can be arbitrarily long.
func (tbl DbCompoundTable) DeleteByCategory(ctx context.Context, req require.Requirement, category ...Category) (int64, error) {
	ii := make([]interface{}, len(category))
	for i, v := range category {
		ii[i] = v
	}
	return support.DeleteByColumn(ctx, tbl, req, "category", ii...)
}

// Delete deletes one or more rows from the table, given a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl DbCompoundTable) Delete(ctx context.Context, req require.Requirement, wh where.Expression) (int64, error) {
	query, args := deleteRowsDbCompoundTableSql(tbl, wh)
	return tbl.Exec(ctx, req, query, args...)
}

func deleteRowsDbCompoundTableSql(tbl DbCompoundTabler, wh where.Expression) (string, []interface{}) {
	whs, args := where.Where(wh, tbl.Dialect().Quoter())
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("DELETE FROM %s %s", quotedName, whs)
	return query, args
}

//-------------------------------------------------------------------------------------------------
