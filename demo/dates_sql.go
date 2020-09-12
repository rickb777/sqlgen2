// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.
// sqlapi v0.45.0; sqlgen v0.64.0

package demo

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/rickb777/date"
	"github.com/rickb777/sqlapi"
	"github.com/rickb777/sqlapi/constraint"
	"github.com/rickb777/sqlapi/dialect"
	"github.com/rickb777/sqlapi/require"
	"github.com/rickb777/sqlapi/support"
	"github.com/rickb777/where"
	"github.com/rickb777/where/quote"
	"strings"
)

// DatesTabler lists table methods provided by DatesTable.
type DatesTabler interface {
	sqlapi.Table

	// Constraints returns the table's constraints.
	Constraints() constraint.Constraints

	// WithConstraint returns a modified DatesTabler with added data consistency constraints.
	WithConstraint(cc ...constraint.Constraint) DatesTabler

	// WithPrefix returns a modified DatesTabler with a given table name prefix.
	WithPrefix(pfx string) DatesTabler

	// WithContext returns a modified DatesTabler with a given context.
	WithContext(ctx context.Context) DatesTabler

	// CreateTable creates the table.
	CreateTable(ifNotExists bool) (int64, error)

	// DropTable drops the table, destroying all its data.
	DropTable(ifExists bool) (int64, error)

	// Truncate drops every record from the table, if possible.
	Truncate(force bool) (err error)
}

//-------------------------------------------------------------------------------------------------

// DatesQueryer lists query methods provided by DatesTable.
type DatesQueryer interface {
	// Name gets the table name. without prefix
	Name() sqlapi.TableName

	// Dialect gets the database dialect.
	Dialect() dialect.Dialect

	// Logger gets the trace logger.
	Logger() sqlapi.Logger

	// Using returns a modified DatesTabler using the transaction supplied.
	Using(tx sqlapi.SqlTx) DatesQueryer

	// Transact runs the function provided within a transaction.
	Transact(txOptions *sql.TxOptions, fn func(DatesQueryer) error) error

	// Tx gets the wrapped transaction handle, provided this is within a transaction.
	// Panics if it is in the wrong state - use IsTx() if necessary.
	Tx() sqlapi.SqlTx

	// IsTx tests whether this is within a transaction.
	IsTx() bool

	// Exec executes a query without returning any rows.

	// Query is the low-level request method for this table using an SQL query that must return all the columns
	// necessary for Dates values.
	Query(req require.Requirement, query string, args ...interface{}) ([]*Dates, error)

	// QueryOneNullString is a low-level access method for one string, returning the first match.
	QueryOneNullString(req require.Requirement, query string, args ...interface{}) (result sql.NullString, err error)

	// QueryOneNullInt64 is a low-level access method for one int64, returning the first match.
	QueryOneNullInt64(req require.Requirement, query string, args ...interface{}) (result sql.NullInt64, err error)

	// QueryOneNullFloat64 is a low-level access method for one float64, returning the first match.
	QueryOneNullFloat64(req require.Requirement, query string, args ...interface{}) (result sql.NullFloat64, err error)

	// GetDatesById gets the record with a given primary key value.
	GetDatesById(req require.Requirement, id uint64) (*Dates, error)

	// GetDatessById gets records from the table according to a list of primary keys.
	GetDatessById(req require.Requirement, qc where.QueryConstraint, id ...uint64) (list []*Dates, err error)

	// SelectOneWhere allows a single Dates to be obtained from the table that matches a 'where' clause.
	SelectOneWhere(req require.Requirement, where, orderBy string, args ...interface{}) (*Dates, error)

	// SelectOne allows a single Dates to be obtained from the table that matches a 'where' clause.
	SelectOne(req require.Requirement, wh where.Expression, qc where.QueryConstraint) (*Dates, error)

	// SelectWhere allows Datess to be obtained from the table that match a 'where' clause.
	SelectWhere(req require.Requirement, where, orderBy string, args ...interface{}) ([]*Dates, error)

	// Select allows Datess to be obtained from the table that match a 'where' clause.
	Select(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]*Dates, error)

	// CountWhere counts Datess in the table that match a 'where' clause.
	CountWhere(where string, args ...interface{}) (count int64, err error)

	// Count counts the Datess in the table that match a 'where' clause.
	Count(wh where.Expression) (count int64, err error)

	// SliceId gets the id column for all rows that match the 'where' condition.
	SliceId(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]uint64, error)

	// SliceInteger gets the integer column for all rows that match the 'where' condition.
	SliceInteger(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]date.Date, error)

	// SliceString gets the string column for all rows that match the 'where' condition.
	SliceString(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]date.DateString, error)

	// Insert adds new records for the Datess, setting the primary key field for each one.
	Insert(req require.Requirement, vv ...*Dates) error

	// UpdateById updates one or more columns, given a id value.
	UpdateById(req require.Requirement, id uint64, fields ...sql.NamedArg) (int64, error)

	// UpdateByInteger updates one or more columns, given a integer value.
	UpdateByInteger(req require.Requirement, integer date.Date, fields ...sql.NamedArg) (int64, error)

	// UpdateByString updates one or more columns, given a string value.
	UpdateByString(req require.Requirement, string date.DateString, fields ...sql.NamedArg) (int64, error)

	// UpdateFields updates one or more columns, given a 'where' clause.
	UpdateFields(req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error)

	// Update updates records, matching them by primary key.
	Update(req require.Requirement, vv ...*Dates) (int64, error)

	// Upsert inserts or updates a record, matching it using the expression supplied.
	// This expression is used to search for an existing record based on some specified
	// key column(s). It must match either zero or one existing record. If it matches
	// none, a new record is inserted; otherwise the matching record is updated. An
	// error results if these conditions are not met.
	Upsert(v *Dates, wh where.Expression) error

	// DeleteById deletes rows from the table, given some id values.
	// The list of ids can be arbitrarily long.
	DeleteById(req require.Requirement, id ...uint64) (int64, error)

	// DeleteByInteger deletes rows from the table, given some integer values.
	// The list of ids can be arbitrarily long.
	DeleteByInteger(req require.Requirement, integer ...date.Date) (int64, error)

	// DeleteByString deletes rows from the table, given some string values.
	// The list of ids can be arbitrarily long.
	DeleteByString(req require.Requirement, string ...date.DateString) (int64, error)

	// Delete deletes one or more rows from the table, given a 'where' clause.
	// Use a nil value for the 'wh' argument if it is not needed (very risky!).
	Delete(req require.Requirement, wh where.Expression) (int64, error)
}

//-------------------------------------------------------------------------------------------------

// DatesTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type DatesTable struct {
	name        sqlapi.TableName
	database    sqlapi.Database
	db          sqlapi.Execer
	constraints constraint.Constraints
	ctx         context.Context
	pk          string
}

// Type conformance checks
var _ sqlapi.TableCreator = &DatesTable{}

// NewDatesTable returns a new table instance.
// If a blank table name is supplied, the default name "datess" will be used instead.
// The request context is initialised with the background.
func NewDatesTable(name string, d sqlapi.Database) DatesTable {
	if name == "" {
		name = "datess"
	}
	var constraints constraint.Constraints
	return DatesTable{
		name:        sqlapi.TableName{Prefix: "", Name: name},
		database:    d,
		db:          d.DB(),
		constraints: constraints,
		ctx:         context.Background(),
		pk:          "id",
	}
}

// CopyTableAsDatesTable copies a table instance, retaining the name etc but
// providing methods appropriate for 'Dates'. It doesn't copy the constraints of the original table.
//
// It serves to provide methods appropriate for 'Dates'. This is most useful when this is used to represent a
// join result. In such cases, there won't be any need for DDL methods, nor Exec, Insert, Update or Delete.
func CopyTableAsDatesTable(origin sqlapi.Table) DatesTable {
	return DatesTable{
		name:        origin.Name(),
		database:    origin.Database(),
		db:          origin.Execer(),
		constraints: nil,
		ctx:         context.Background(),
		pk:          "id",
	}
}

// SetPkColumn sets the name of the primary key column. It defaults to "id".
// The result is a modified copy of the table; the original is unchanged.
//func (tbl DatesTable) SetPkColumn(pk string) DatesTabler {
//	tbl.pk = pk
//	return tbl
//}

// WithPrefix sets the table name prefix for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl DatesTable) WithPrefix(pfx string) DatesTabler {
	tbl.name.Prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries via this table.
// The result is a modified copy of the table; the original is unchanged.
//
// The shared context in the *Database is not altered by this method. So it
// is possible to use different contexts for different (groups of) queries.
func (tbl DatesTable) WithContext(ctx context.Context) DatesTabler {
	tbl.ctx = ctx
	return tbl
}

// Database gets the shared database information.
func (tbl DatesTable) Database() sqlapi.Database {
	return tbl.database
}

// Logger gets the trace logger.
func (tbl DatesTable) Logger() sqlapi.Logger {
	return tbl.database.Logger()
}

// WithConstraint returns a modified DatesTabler with added data consistency constraints.
func (tbl DatesTable) WithConstraint(cc ...constraint.Constraint) DatesTabler {
	tbl.constraints = append(tbl.constraints, cc...)
	return tbl
}

// Constraints returns the table's constraints.
func (tbl DatesTable) Constraints() constraint.Constraints {
	return tbl.constraints
}

// Ctx gets the current request context.
func (tbl DatesTable) Ctx() context.Context {
	return tbl.ctx
}

// Dialect gets the database dialect.
func (tbl DatesTable) Dialect() dialect.Dialect {
	return tbl.database.Dialect()
}

// Name gets the table name.
func (tbl DatesTable) Name() sqlapi.TableName {
	return tbl.name
}

// PkColumn gets the column name used as a primary key.
func (tbl DatesTable) PkColumn() string {
	return tbl.pk
}

// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl DatesTable) DB() sqlapi.SqlDB {
	return tbl.db.(sqlapi.SqlDB)
}

// Execer gets the wrapped database or transaction handle.
func (tbl DatesTable) Execer() sqlapi.Execer {
	return tbl.db
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl DatesTable) Tx() sqlapi.SqlTx {
	return tbl.db.(sqlapi.SqlTx)
}

// IsTx tests whether this is within a transaction.
func (tbl DatesTable) IsTx() bool {
	return tbl.db.IsTx()
}

// Using returns a modified DatesTabler using the transaction supplied. This is
// needed when making multiple queries across several tables within a single transaction.
//
// The result is a modified copy of the table; the original is unchanged.
func (tbl DatesTable) Using(tx sqlapi.SqlTx) DatesQueryer {
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
func (tbl DatesTable) Transact(txOptions *sql.TxOptions, fn func(DatesQueryer) error) error {
	var err error
	if tbl.IsTx() {
		err = fn(tbl) // nested transactions are inlined
	} else {
		err = tbl.DB().Transact(tbl.ctx, txOptions, func(tx sqlapi.SqlTx) error {
			return fn(tbl.Using(tx))
		})
	}
	return tbl.Logger().LogIfError(err)
}

func (tbl DatesTable) quotedName() string {
	return tbl.Dialect().Quoter().Quote(tbl.name.String())
}

func (tbl DatesTable) quotedNameW(w dialect.StringWriter) {
	tbl.Dialect().Quoter().QuoteW(w, tbl.name.String())
}

//-------------------------------------------------------------------------------------------------

// NumDatesTableColumns is the total number of columns in DatesTable.
const NumDatesTableColumns = 3

// NumDatesTableDataColumns is the number of columns in DatesTable not including the auto-increment key.
const NumDatesTableDataColumns = 2

// DatesTableColumnNames is the list of columns in DatesTable.
const DatesTableColumnNames = "id,integer,string"

// DatesTableDataColumnNames is the list of data columns in DatesTable.
const DatesTableDataColumnNames = "integer,string"

var listOfDatesTableColumnNames = strings.Split(DatesTableColumnNames, ",")

//-------------------------------------------------------------------------------------------------

var sqlDatesTableCreateColumnsSqlite = []string{
	"integer not null primary key autoincrement",
	"bigint not null",
	"text not null",
}

var sqlDatesTableCreateColumnsMysql = []string{
	"bigint unsigned not null primary key auto_increment",
	"bigint not null",
	"text not null",
}

var sqlDatesTableCreateColumnsPostgres = []string{
	"bigserial not null primary key",
	"bigint not null",
	"text not null",
}

var sqlDatesTableCreateColumnsPgx = []string{
	"bigserial not null primary key",
	"bigint not null",
	"text not null",
}

//-------------------------------------------------------------------------------------------------

// CreateTable creates the table.
func (tbl DatesTable) CreateTable(ifNotExists bool) (int64, error) {
	return support.Exec(tbl, nil, createDatesTableSql(tbl, ifNotExists))
}

func createDatesTableSql(tbl DatesTabler, ifNotExists bool) string {
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
		columns = sqlDatesTableCreateColumnsSqlite
	case dialect.MysqlIndex:
		columns = sqlDatesTableCreateColumnsMysql
	case dialect.PostgresIndex:
		columns = sqlDatesTableCreateColumnsPostgres
	case dialect.PgxIndex:
		columns = sqlDatesTableCreateColumnsPgx
	}

	comma := ""
	for i, n := range listOfDatesTableColumnNames {
		buf.WriteString(comma)
		q.QuoteW(buf, n)
		buf.WriteString(" ")
		buf.WriteString(columns[i])
		comma = ",\n "
	}

	for i, c := range tbl.(DatesTable).Constraints() {
		buf.WriteString(",\n ")
		buf.WriteString(c.ConstraintSql(tbl.Dialect().Quoter(), tbl.Name(), i+1))
	}

	buf.WriteString("\n)")
	buf.WriteString(tbl.Dialect().CreateTableSettings())
	return buf.String()
}

func ternaryDatesTable(flag bool, a, b string) string {
	if flag {
		return a
	}
	return b
}

// DropTable drops the table, destroying all its data.
func (tbl DatesTable) DropTable(ifExists bool) (int64, error) {
	return support.Exec(tbl, nil, dropDatesTableSql(tbl, ifExists))
}

func dropDatesTableSql(tbl DatesTabler, ifExists bool) string {
	ie := ternaryDatesTable(ifExists, "IF EXISTS ", "")
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
func (tbl DatesTable) Truncate(force bool) (err error) {
	for _, query := range tbl.Dialect().TruncateDDL(tbl.Name().String(), force) {
		_, err = support.Exec(tbl, nil, query)
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
func (tbl DatesTable) Exec(req require.Requirement, query string, args ...interface{}) (int64, error) {
	return support.Exec(tbl, req, query, args...)
}

//-------------------------------------------------------------------------------------------------

// Query is the low-level request method for this table. The SQL query must return all the columns necessary for
// Dates values. Placeholders should be vanilla '?' marks, which will be replaced if necessary according to
// the chosen dialect.
//
// The query is logged using whatever logger is configured. If an error arises, this too is logged.
//
// If you need a context other than the background, use WithContext before calling Query.
//
// The args are for any placeholder parameters in the query.
//
// The support API provides a core 'support.Query' function, on which this method depends. If appropriate,
// use that function directly; wrap the result in *sqlapi.Rows if you need to access its data as a map.
func (tbl DatesTable) Query(req require.Requirement, query string, args ...interface{}) ([]*Dates, error) {
	return doDatesTableQueryAndScan(tbl, req, false, query, args)
}

func doDatesTableQueryAndScan(tbl DatesTabler, req require.Requirement, firstOnly bool, query string, args ...interface{}) ([]*Dates, error) {
	rows, err := support.Query(tbl.(sqlapi.Table), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vv, n, err := ScanDatess(query, rows, firstOnly)
	return vv, tbl.(sqlapi.Table).Logger().LogIfError(require.ChainErrorIfQueryNotSatisfiedBy(err, req, n))
}

//-------------------------------------------------------------------------------------------------

// QueryOneNullString is a low-level access method for one string. This can be used for function queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl DatesTable) QueryOneNullString(req require.Requirement, query string, args ...interface{}) (result sql.NullString, err error) {
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
func (tbl DatesTable) QueryOneNullInt64(req require.Requirement, query string, args ...interface{}) (result sql.NullInt64, err error) {
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
func (tbl DatesTable) QueryOneNullFloat64(req require.Requirement, query string, args ...interface{}) (result sql.NullFloat64, err error) {
	err = support.QueryOneNullThing(tbl, req, &result, query, args...)
	return result, err
}

// ScanDatess reads rows from the database and returns a slice of corresponding values.
// It also returns a number indicating how many rows were read; this will be larger than the length of the
// slice if reading stopped after the first row.
func ScanDatess(query string, rows sqlapi.SqlRows, firstOnly bool) (vv []*Dates, n int64, err error) {
	for rows.Next() {
		n++

		var v0 uint64
		var v1 date.Date
		var v2 date.DateString

		err = rows.Scan(
			&v0,
			&v1,
			&v2,
		)
		if err != nil {
			return vv, n, errors.Wrap(err, query)
		}

		v := &Dates{}
		v.Id = v0
		v.Integer = v1
		v.String = v2

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

func allDatesColumnNamesQuoted(q quote.Quoter) string {
	return strings.Join(q.QuoteN(listOfDatesTableColumnNames), ",")
}

//--------------------------------------------------------------------------------

// GetDatesById gets the record with a given primary key value.
// If not found, *Dates will be nil.
func (tbl DatesTable) GetDatesById(req require.Requirement, id uint64) (*Dates, error) {
	return tbl.SelectOne(req, where.Eq("id", id), nil)
}

// GetDatessById gets records from the table according to a list of primary keys.
// Although the list of ids can be arbitrarily long, there are practical limits;
// note that Oracle DB has a limit of 1000.
//
// It places a requirement, which may be nil, on the size of the expected results: in particular, require.All
// controls whether an error is generated not all the ids produce a result.
func (tbl DatesTable) GetDatessById(req require.Requirement, qc where.QueryConstraint, id ...uint64) (list []*Dates, err error) {
	if req == require.All {
		req = require.Exactly(len(id))
	}
	return tbl.Select(req, where.In("id", id), qc)
}

func doDatesTableQueryAndScanOne(tbl DatesTabler, req require.Requirement, query string, args ...interface{}) (*Dates, error) {
	list, err := doDatesTableQueryAndScan(tbl, req, true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

// Fetch fetches a list of Dates based on a supplied query. This is mostly used for join queries that map its
// result columns to the fields of Dates. Other queries might be better handled by GetXxx or Select methods.
func (tbl DatesTable) Fetch(req require.Requirement, query string, args ...interface{}) ([]*Dates, error) {
	return doDatesTableQueryAndScan(tbl, req, false, query, args...)
}

//-------------------------------------------------------------------------------------------------

// SelectOneWhere allows a single Dates to be obtained from the table that matches a 'where' clause
// and some limit. Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
// If not found, *Example will be nil.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.One
// controls whether an error is generated when no result is found.
//
// The args are for any placeholder parameters in the query.
func (tbl DatesTable) SelectOneWhere(req require.Requirement, where, orderBy string, args ...interface{}) (*Dates, error) {
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT 1",
		allDatesColumnNamesQuoted(tbl.Dialect().Quoter()), quotedName, where, orderBy)
	v, err := doDatesTableQueryAndScanOne(tbl, req, query, args...)
	return v, err
}

// SelectOne allows a single Dates to be obtained from the table that matches a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
// If not found, *Example will be nil.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.One
// controls whether an error is generated when no result is found.
func (tbl DatesTable) SelectOne(req require.Requirement, wh where.Expression, qc where.QueryConstraint) (*Dates, error) {
	q := tbl.Dialect().Quoter()
	whs, args := where.Where(wh, q)
	orderBy := where.Build(qc, q)
	return tbl.SelectOneWhere(req, whs, orderBy, args...)
}

// SelectWhere allows Datess to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
//
// The args are for any placeholder parameters in the query.
func (tbl DatesTable) SelectWhere(req require.Requirement, where, orderBy string, args ...interface{}) ([]*Dates, error) {
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("SELECT %s FROM %s %s %s",
		allDatesColumnNamesQuoted(tbl.Dialect().Quoter()), quotedName, where, orderBy)
	vv, err := doDatesTableQueryAndScan(tbl, req, false, query, args...)
	return vv, err
}

// Select allows Datess to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
func (tbl DatesTable) Select(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]*Dates, error) {
	q := tbl.Dialect().Quoter()
	whs, args := where.Where(wh, q)
	orderBy := where.Build(qc, q)
	return tbl.SelectWhere(req, whs, orderBy, args...)
}

//-------------------------------------------------------------------------------------------------

// CountWhere counts Datess in the table that match a 'where' clause.
// Use a blank string for the 'where' argument if it is not needed.
//
// The args are for any placeholder parameters in the query.
func (tbl DatesTable) CountWhere(where string, args ...interface{}) (count int64, err error) {
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s %s", quotedName, where)
	rows, err := support.Query(tbl, query, args...)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&count)
	}
	return count, tbl.Logger().LogIfError(err)
}

// Count counts the Datess in the table that match a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed.
func (tbl DatesTable) Count(wh where.Expression) (count int64, err error) {
	whs, args := where.Where(wh, tbl.Dialect().Quoter())
	return tbl.CountWhere(whs, args...)
}

//--------------------------------------------------------------------------------

// SliceId gets the id column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DatesTable) SliceId(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]uint64, error) {
	return support.SliceUint64List(tbl, req, tbl.pk, wh, qc)
}

// SliceInteger gets the integer column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DatesTable) SliceInteger(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]date.Date, error) {
	return sliceDatesTableDateList(tbl, req, "integer", wh, qc)
}

// SliceString gets the string column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DatesTable) SliceString(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]date.DateString, error) {
	return sliceDatesTableDateStringList(tbl, req, "string", wh, qc)
}

func sliceDatesTableDateList(tbl DatesTabler, req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]date.Date, error) {
	q := tbl.Dialect().Quoter()
	whs, args := where.Where(wh, q)
	orderBy := where.Build(qc, q)
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", q.Quote(sqlname), quotedName, whs, orderBy)
	rows, err := support.Query(tbl.(sqlapi.Table), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]date.Date, 0, 10)

	for rows.Next() {
		var v date.Date
		err = rows.Scan(&v)
		if err == sql.ErrNoRows {
			return list, tbl.(sqlapi.Table).Logger().LogIfError(require.ErrorIfQueryNotSatisfiedBy(req, int64(len(list))))
		} else {
			list = append(list, v)
		}
	}
	return list, tbl.(sqlapi.Table).Logger().LogIfError(require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, int64(len(list))))
}

func sliceDatesTableDateStringList(tbl DatesTabler, req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]date.DateString, error) {
	q := tbl.Dialect().Quoter()
	whs, args := where.Where(wh, q)
	orderBy := where.Build(qc, q)
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", q.Quote(sqlname), quotedName, whs, orderBy)
	rows, err := support.Query(tbl.(sqlapi.Table), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]date.DateString, 0, 10)

	for rows.Next() {
		var v date.DateString
		err = rows.Scan(&v)
		if err == sql.ErrNoRows {
			return list, tbl.(sqlapi.Table).Logger().LogIfError(require.ErrorIfQueryNotSatisfiedBy(req, int64(len(list))))
		} else {
			list = append(list, v)
		}
	}
	return list, tbl.(sqlapi.Table).Logger().LogIfError(require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, int64(len(list))))
}

func constructDatesTableInsert(tbl DatesTable, w dialect.StringWriter, v *Dates, withPk bool) (s []interface{}, err error) {
	q := tbl.Dialect().Quoter()
	s = make([]interface{}, 0, 3)

	comma := ""
	w.WriteString(" (")

	if withPk {
		q.QuoteW(w, "id")
		comma = ","
		s = append(s, v.Id)
	}

	w.WriteString(comma)
	q.QuoteW(w, "integer")
	s = append(s, v.Integer)
	comma = ","

	w.WriteString(comma)
	q.QuoteW(w, "string")
	s = append(s, v.String)

	w.WriteString(")")
	return s, nil
}

func constructDatesTableUpdate(tbl DatesTable, w dialect.StringWriter, v *Dates) (s []interface{}, err error) {
	q := tbl.Dialect().Quoter()
	j := 1
	s = make([]interface{}, 0, 2)

	comma := ""

	w.WriteString(comma)
	q.QuoteW(w, "integer")
	w.WriteString("=?")
	s = append(s, v.Integer)
	j++
	comma = ", "

	w.WriteString(comma)
	q.QuoteW(w, "string")
	w.WriteString("=?")
	s = append(s, v.String)
	j++
	return s, nil
}

//--------------------------------------------------------------------------------

// Insert adds new records for the Datess.// The Datess have their primary key fields set to the new record identifiers.
// The Dates.PreInsert() method will be called, if it exists.
func (tbl DatesTable) Insert(req require.Requirement, vv ...*Dates) error {
	if req == require.All {
		req = require.Exactly(len(vv))
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

		fields, err := constructDatesTableInsert(tbl, b, v, false)
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
			row := tbl.db.QueryRowContext(tbl.ctx, query, fields...)
			var i64 int64
			err = row.Scan(&i64)
			v.Id = uint64(i64)

		} else {
			i64, e2 := tbl.db.InsertContext(tbl.ctx, tbl.pk, query, fields...)
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
func (tbl DatesTable) UpdateById(req require.Requirement, id uint64, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(req, where.Eq("id", id), fields...)
}

// UpdateByInteger updates one or more columns, given a integer value.
func (tbl DatesTable) UpdateByInteger(req require.Requirement, integer date.Date, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(req, where.Eq("integer", integer), fields...)
}

// UpdateByString updates one or more columns, given a string value.
func (tbl DatesTable) UpdateByString(req require.Requirement, string date.DateString, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(req, where.Eq("string", string), fields...)
}

// UpdateFields updates one or more columns, given a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed (but note that this is risky!).
func (tbl DatesTable) UpdateFields(req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error) {
	return support.UpdateFields(tbl, req, wh, fields...)
}

//--------------------------------------------------------------------------------

// Update updates records, matching them by primary key. It returns the number of rows affected.
// The Dates.PreUpdate(Execer) method will be called, if it exists.
func (tbl DatesTable) Update(req require.Requirement, vv ...*Dates) (int64, error) {
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

		args, err := constructDatesTableUpdate(tbl, b, v)
		if err != nil {
			return count, err
		}
		args = append(args, v.Id)

		b.WriteString(" WHERE ")
		q.QuoteW(b, tbl.pk)
		b.WriteString("=?")

		query := b.String()
		n, err := tbl.Exec(nil, query, args...)
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
func (tbl DatesTable) Upsert(v *Dates, wh where.Expression) error {
	col := tbl.Dialect().Quoter().Quote(tbl.pk)
	qName := tbl.quotedName()
	whs, args := where.Where(wh, tbl.Dialect().Quoter())

	query := fmt.Sprintf("SELECT %s FROM %s %s", col, qName, whs)
	rows, err := support.Query(tbl, query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	if !rows.Next() {
		return tbl.Insert(require.One, v)
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
	_, err = tbl.Update(require.One, v)
	return err
}

//-------------------------------------------------------------------------------------------------

// DeleteById deletes rows from the table, given some id values.
// The list of ids can be arbitrarily long.
func (tbl DatesTable) DeleteById(req require.Requirement, id ...uint64) (int64, error) {
	ii := support.Uint64AsInterfaceSlice(id)
	return support.DeleteByColumn(tbl, req, "id", ii...)
}

// DeleteByInteger deletes rows from the table, given some integer values.
// The list of ids can be arbitrarily long.
func (tbl DatesTable) DeleteByInteger(req require.Requirement, integer ...date.Date) (int64, error) {
	ii := make([]interface{}, len(integer))
	for i, v := range integer {
		ii[i] = v
	}
	return support.DeleteByColumn(tbl, req, "integer", ii...)
}

// DeleteByString deletes rows from the table, given some string values.
// The list of ids can be arbitrarily long.
func (tbl DatesTable) DeleteByString(req require.Requirement, string ...date.DateString) (int64, error) {
	ii := make([]interface{}, len(string))
	for i, v := range string {
		ii[i] = v
	}
	return support.DeleteByColumn(tbl, req, "string", ii...)
}

// Delete deletes one or more rows from the table, given a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl DatesTable) Delete(req require.Requirement, wh where.Expression) (int64, error) {
	query, args := deleteRowsDatesTableSql(tbl, wh)
	return tbl.Exec(req, query, args...)
}

func deleteRowsDatesTableSql(tbl DatesTabler, wh where.Expression) (string, []interface{}) {
	whs, args := where.Where(wh, tbl.Dialect().Quoter())
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("DELETE FROM %s %s", quotedName, whs)
	return query, args
}

//-------------------------------------------------------------------------------------------------
