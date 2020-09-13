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

// AssociationTabler lists table methods provided by AssociationTable.
type AssociationTabler interface {
	sqlapi.Table

	// Constraints returns the table's constraints.
	Constraints() constraint.Constraints

	// WithConstraint returns a modified AssociationTabler with added data consistency constraints.
	WithConstraint(cc ...constraint.Constraint) AssociationTabler

	// WithPrefix returns a modified AssociationTabler with a given table name prefix.
	WithPrefix(pfx string) AssociationTabler

	// CreateTable creates the table.
	CreateTable(ctx context.Context, ifNotExists bool) (int64, error)

	// DropTable drops the table, destroying all its data.
	DropTable(ctx context.Context, ifExists bool) (int64, error)

	// Truncate drops every record from the table, if possible.
	Truncate(ctx context.Context, force bool) (err error)
}

//-------------------------------------------------------------------------------------------------

// AssociationQueryer lists query methods provided by AssociationTable.
type AssociationQueryer interface {
	sqlapi.Table

	// Using returns a modified AssociationQueryer using the Execer supplied,
	// which will typically be a transaction (i.e. SqlTx).
	Using(tx sqlapi.Execer) AssociationQueryer

	// Transact runs the function provided within a transaction. The transction is committed
	// unless an error occurs.
	Transact(ctx context.Context, txOptions *sql.TxOptions, fn func(AssociationQueryer) error) error

	// Exec executes a query without returning any rows.
	Exec(ctx context.Context, req require.Requirement, query string, args ...interface{}) (int64, error)

	// Query is the low-level request method for this table using an SQL query that must return all the columns
	// necessary for Association values.
	Query(ctx context.Context, req require.Requirement, query string, args ...interface{}) ([]*Association, error)

	// QueryOneNullString is a low-level access method for one string, returning the first match.
	QueryOneNullString(ctx context.Context, req require.Requirement, query string, args ...interface{}) (result sql.NullString, err error)

	// QueryOneNullInt64 is a low-level access method for one int64, returning the first match.
	QueryOneNullInt64(ctx context.Context, req require.Requirement, query string, args ...interface{}) (result sql.NullInt64, err error)

	// QueryOneNullFloat64 is a low-level access method for one float64, returning the first match.
	QueryOneNullFloat64(ctx context.Context, req require.Requirement, query string, args ...interface{}) (result sql.NullFloat64, err error)

	// GetAssociationById gets the record with a given primary key value.
	GetAssociationById(ctx context.Context, req require.Requirement, id int64) (*Association, error)

	// GetAssociationsById gets records from the table according to a list of primary keys.
	GetAssociationsById(ctx context.Context, req require.Requirement, qc where.QueryConstraint, id ...int64) (list []*Association, err error)

	// SelectOneWhere allows a single Association to be obtained from the table that matches a 'where' clause.
	SelectOneWhere(ctx context.Context, req require.Requirement, where, orderBy string, args ...interface{}) (*Association, error)

	// SelectOne allows a single Association to be obtained from the table that matches a 'where' clause.
	SelectOne(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) (*Association, error)

	// SelectWhere allows Associations to be obtained from the table that match a 'where' clause.
	SelectWhere(ctx context.Context, req require.Requirement, where, orderBy string, args ...interface{}) ([]*Association, error)

	// Select allows Associations to be obtained from the table that match a 'where' clause.
	Select(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]*Association, error)

	// CountWhere counts Associations in the table that match a 'where' clause.
	CountWhere(ctx context.Context, where string, args ...interface{}) (count int64, err error)

	// Count counts the Associations in the table that match a 'where' clause.
	Count(ctx context.Context, wh where.Expression) (count int64, err error)

	// SliceId gets the id column for all rows that match the 'where' condition.
	SliceId(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error)

	// SliceName gets the name column for all rows that match the 'where' condition.
	SliceName(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error)

	// SliceRef1 gets the ref1 column for all rows that match the 'where' condition.
	SliceRef1(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error)

	// SliceRef2 gets the ref2 column for all rows that match the 'where' condition.
	SliceRef2(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error)

	// SliceQuality gets the quality column for all rows that match the 'where' condition.
	SliceQuality(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]QualName, error)

	// SliceCategory gets the category column for all rows that match the 'where' condition.
	SliceCategory(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]Category, error)

	// Insert adds new records for the Associations, setting the primary key field for each one.
	Insert(ctx context.Context, req require.Requirement, vv ...*Association) error

	// UpdateById updates one or more columns, given a id value.
	UpdateById(ctx context.Context, req require.Requirement, id int64, fields ...sql.NamedArg) (int64, error)

	// UpdateByName updates one or more columns, given a name value.
	UpdateByName(ctx context.Context, req require.Requirement, name string, fields ...sql.NamedArg) (int64, error)

	// UpdateByQuality updates one or more columns, given a quality value.
	UpdateByQuality(ctx context.Context, req require.Requirement, quality QualName, fields ...sql.NamedArg) (int64, error)

	// UpdateByRef1 updates one or more columns, given a ref1 value.
	UpdateByRef1(ctx context.Context, req require.Requirement, ref1 int64, fields ...sql.NamedArg) (int64, error)

	// UpdateByRef2 updates one or more columns, given a ref2 value.
	UpdateByRef2(ctx context.Context, req require.Requirement, ref2 int64, fields ...sql.NamedArg) (int64, error)

	// UpdateByCategory updates one or more columns, given a category value.
	UpdateByCategory(ctx context.Context, req require.Requirement, category Category, fields ...sql.NamedArg) (int64, error)

	// UpdateFields updates one or more columns, given a 'where' clause.
	UpdateFields(ctx context.Context, req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error)

	// Update updates records, matching them by primary key.
	Update(ctx context.Context, req require.Requirement, vv ...*Association) (int64, error)

	// Upsert inserts or updates a record, matching it using the expression supplied.
	// This expression is used to search for an existing record based on some specified
	// key column(s). It must match either zero or one existing record. If it matches
	// none, a new record is inserted; otherwise the matching record is updated. An
	// error results if these conditions are not met.
	Upsert(ctx context.Context, v *Association, wh where.Expression) error

	// DeleteById deletes rows from the table, given some id values.
	// The list of ids can be arbitrarily long.
	DeleteById(ctx context.Context, req require.Requirement, id ...int64) (int64, error)

	// DeleteByName deletes rows from the table, given some name values.
	// The list of ids can be arbitrarily long.
	DeleteByName(ctx context.Context, req require.Requirement, name ...string) (int64, error)

	// DeleteByQuality deletes rows from the table, given some quality values.
	// The list of ids can be arbitrarily long.
	DeleteByQuality(ctx context.Context, req require.Requirement, quality ...QualName) (int64, error)

	// DeleteByRef1 deletes rows from the table, given some ref1 values.
	// The list of ids can be arbitrarily long.
	DeleteByRef1(ctx context.Context, req require.Requirement, ref1 ...int64) (int64, error)

	// DeleteByRef2 deletes rows from the table, given some ref2 values.
	// The list of ids can be arbitrarily long.
	DeleteByRef2(ctx context.Context, req require.Requirement, ref2 ...int64) (int64, error)

	// DeleteByCategory deletes rows from the table, given some category values.
	// The list of ids can be arbitrarily long.
	DeleteByCategory(ctx context.Context, req require.Requirement, category ...Category) (int64, error)

	// Delete deletes one or more rows from the table, given a 'where' clause.
	// Use a nil value for the 'wh' argument if it is not needed (very risky!).
	Delete(ctx context.Context, req require.Requirement, wh where.Expression) (int64, error)
}

//-------------------------------------------------------------------------------------------------

// AssociationTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type AssociationTable struct {
	name        sqlapi.TableName
	database    sqlapi.Database
	db          sqlapi.Execer
	constraints constraint.Constraints
	pk          string
}

// Type conformance checks
var _ sqlapi.TableCreator = &AssociationTable{}

// NewAssociationTable returns a new table instance.
// If a blank table name is supplied, the default name "associations" will be used instead.
// The request context is initialised with the background.
func NewAssociationTable(name string, d sqlapi.Database) AssociationTable {
	if name == "" {
		name = "associations"
	}
	var constraints constraint.Constraints
	return AssociationTable{
		name:        sqlapi.TableName{Prefix: "", Name: name},
		database:    d,
		db:          d.DB(),
		constraints: constraints,
		pk:          "id",
	}
}

// CopyTableAsAssociationTable copies a table instance, retaining the name etc but
// providing methods appropriate for 'Association'.It doesn't copy the constraints of the original table.
//
// It serves to provide methods appropriate for 'Association'. This is most useful when this is used to represent a
// join result. In such cases, there won't be any need for DDL methods, nor Exec, Insert, Update or Delete.
func CopyTableAsAssociationTable(origin sqlapi.Table) AssociationTable {
	return AssociationTable{
		name:        origin.Name(),
		database:    origin.Database(),
		db:          origin.Execer(),
		constraints: nil,
		pk:          "id",
	}
}

// SetPkColumn sets the name of the primary key column. It defaults to "id".
// The result is a modified copy of the table; the original is unchanged.
//func (tbl AssociationTable) SetPkColumn(pk string) AssociationTabler {
//	tbl.pk = pk
//	return tbl
//}

// WithPrefix sets the table name prefix for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl AssociationTable) WithPrefix(pfx string) AssociationTabler {
	tbl.name.Prefix = pfx
	return tbl
}

// Database gets the shared database information.
func (tbl AssociationTable) Database() sqlapi.Database {
	return tbl.database
}

// Logger gets the trace logger.
func (tbl AssociationTable) Logger() sqlapi.Logger {
	return tbl.database.Logger()
}

// WithConstraint returns a modified AssociationTabler with added data consistency constraints.
func (tbl AssociationTable) WithConstraint(cc ...constraint.Constraint) AssociationTabler {
	tbl.constraints = append(tbl.constraints, cc...)
	return tbl
}

// Constraints returns the table's constraints.
func (tbl AssociationTable) Constraints() constraint.Constraints {
	return tbl.constraints
}

// Dialect gets the database dialect.
func (tbl AssociationTable) Dialect() dialect.Dialect {
	return tbl.database.Dialect()
}

// Name gets the table name.
func (tbl AssociationTable) Name() sqlapi.TableName {
	return tbl.name
}

// PkColumn gets the column name used as a primary key.
func (tbl AssociationTable) PkColumn() string {
	return tbl.pk
}

// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl AssociationTable) DB() sqlapi.SqlDB {
	return tbl.db.(sqlapi.SqlDB)
}

// Execer gets the wrapped database or transaction handle.
func (tbl AssociationTable) Execer() sqlapi.Execer {
	return tbl.db
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl AssociationTable) Tx() sqlapi.SqlTx {
	return tbl.db.(sqlapi.SqlTx)
}

// IsTx tests whether this is within a transaction.
func (tbl AssociationTable) IsTx() bool {
	return tbl.db.IsTx()
}

// Using returns a modified AssociationTabler using the the Execer supplied,
// which will typically be a transaction (i.e. SqlTx). This is needed when making multiple
// queries across several tables within a single transaction.
//
// The result is a modified copy of the table; the original is unchanged.
func (tbl AssociationTable) Using(tx sqlapi.Execer) AssociationQueryer {
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
func (tbl AssociationTable) Transact(ctx context.Context, txOptions *sql.TxOptions, fn func(AssociationQueryer) error) error {
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

func (tbl AssociationTable) quotedName() string {
	return tbl.Dialect().Quoter().Quote(tbl.name.String())
}

func (tbl AssociationTable) quotedNameW(w dialect.StringWriter) {
	tbl.Dialect().Quoter().QuoteW(w, tbl.name.String())
}

//-------------------------------------------------------------------------------------------------

// NumAssociationTableColumns is the total number of columns in AssociationTable.
const NumAssociationTableColumns = 6

// NumAssociationTableDataColumns is the number of columns in AssociationTable not including the auto-increment key.
const NumAssociationTableDataColumns = 5

// AssociationTableColumnNames is the list of columns in AssociationTable.
const AssociationTableColumnNames = "id,name,quality,ref1,ref2,category"

// AssociationTableDataColumnNames is the list of data columns in AssociationTable.
const AssociationTableDataColumnNames = "name,quality,ref1,ref2,category"

var listOfAssociationTableColumnNames = strings.Split(AssociationTableColumnNames, ",")

//-------------------------------------------------------------------------------------------------

var sqlAssociationTableCreateColumnsSqlite = []string{
	"integer not null primary key autoincrement",
	"text default null",
	"text default null",
	"bigint default null",
	"bigint default null",
	"tinyint unsigned default null",
}

var sqlAssociationTableCreateColumnsMysql = []string{
	"bigint not null primary key auto_increment",
	"text default null",
	"text default null",
	"bigint default null",
	"bigint default null",
	"tinyint unsigned default null",
}

var sqlAssociationTableCreateColumnsPostgres = []string{
	"bigserial not null primary key",
	"text default null",
	"text default null",
	"bigint default null",
	"bigint default null",
	"smallint default null",
}

var sqlAssociationTableCreateColumnsPgx = []string{
	"bigserial not null primary key",
	"text default null",
	"text default null",
	"bigint default null",
	"bigint default null",
	"smallint default null",
}

//-------------------------------------------------------------------------------------------------

// CreateTable creates the table.
func (tbl AssociationTable) CreateTable(ctx context.Context, ifNotExists bool) (int64, error) {
	return support.Exec(ctx, tbl, nil, createAssociationTableSql(tbl, ifNotExists))
}

func createAssociationTableSql(tbl AssociationTabler, ifNotExists bool) string {
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
		columns = sqlAssociationTableCreateColumnsSqlite
	case dialect.MysqlIndex:
		columns = sqlAssociationTableCreateColumnsMysql
	case dialect.PostgresIndex:
		columns = sqlAssociationTableCreateColumnsPostgres
	case dialect.PgxIndex:
		columns = sqlAssociationTableCreateColumnsPgx
	}

	comma := ""
	for i, n := range listOfAssociationTableColumnNames {
		buf.WriteString(comma)
		q.QuoteW(buf, n)
		buf.WriteString(" ")
		buf.WriteString(columns[i])
		comma = ",\n "
	}

	for i, c := range tbl.(AssociationTable).Constraints() {
		buf.WriteString(",\n ")
		buf.WriteString(c.ConstraintSql(tbl.Dialect().Quoter(), tbl.Name(), i+1))
	}

	buf.WriteString("\n)")
	buf.WriteString(tbl.Dialect().CreateTableSettings())
	return buf.String()
}

func ternaryAssociationTable(flag bool, a, b string) string {
	if flag {
		return a
	}
	return b
}

// DropTable drops the table, destroying all its data.
func (tbl AssociationTable) DropTable(ctx context.Context, ifExists bool) (int64, error) {
	return support.Exec(ctx, tbl, nil, dropAssociationTableSql(tbl, ifExists))
}

func dropAssociationTableSql(tbl AssociationTabler, ifExists bool) string {
	ie := ternaryAssociationTable(ifExists, "IF EXISTS ", "")
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
func (tbl AssociationTable) Truncate(ctx context.Context, force bool) (err error) {
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
func (tbl AssociationTable) Exec(ctx context.Context, req require.Requirement, query string, args ...interface{}) (int64, error) {
	return support.Exec(ctx, tbl, req, query, args...)
}

//-------------------------------------------------------------------------------------------------

// Query is the low-level request method for this table. The SQL query must return all the columns necessary for
// Association values. Placeholders should be vanilla '?' marks, which will be replaced if necessary according to
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
func (tbl AssociationTable) Query(ctx context.Context, req require.Requirement, query string, args ...interface{}) ([]*Association, error) {
	return doAssociationTableQueryAndScan(ctx, tbl, req, false, query, args)
}

func doAssociationTableQueryAndScan(ctx context.Context, tbl AssociationTabler, req require.Requirement, firstOnly bool, query string, args ...interface{}) ([]*Association, error) {
	rows, err := support.Query(ctx, tbl, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vv, n, err := ScanAssociations(query, rows, firstOnly)
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
func (tbl AssociationTable) QueryOneNullString(ctx context.Context, req require.Requirement, query string, args ...interface{}) (result sql.NullString, err error) {
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
func (tbl AssociationTable) QueryOneNullInt64(ctx context.Context, req require.Requirement, query string, args ...interface{}) (result sql.NullInt64, err error) {
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
func (tbl AssociationTable) QueryOneNullFloat64(ctx context.Context, req require.Requirement, query string, args ...interface{}) (result sql.NullFloat64, err error) {
	err = support.QueryOneNullThing(ctx, tbl, req, &result, query, args...)
	return result, err
}

// ScanAssociations reads rows from the database and returns a slice of corresponding values.
// It also returns a number indicating how many rows were read; this will be larger than the length of the
// slice if reading stopped after the first row.
func ScanAssociations(query string, rows sqlapi.SqlRows, firstOnly bool) (vv []*Association, n int64, err error) {
	for rows.Next() {
		n++

		var v0 int64
		var v1 sql.NullString
		var v2 sql.NullString
		var v3 sql.NullInt64
		var v4 sql.NullInt64
		var v5 sql.NullInt64

		err = rows.Scan(
			&v0,
			&v1,
			&v2,
			&v3,
			&v4,
			&v5,
		)
		if err != nil {
			return vv, n, errors.Wrap(err, query)
		}

		v := &Association{}
		v.Id = v0
		if v1.Valid {
			a := v1.String
			v.Name = &a
		}
		if v2.Valid {
			a := QualName(v2.String)
			v.Quality = &a
		}
		if v3.Valid {
			a := v3.Int64
			v.Ref1 = &a
		}
		if v4.Valid {
			a := v4.Int64
			v.Ref2 = &a
		}
		if v5.Valid {
			a := Category(v5.Int64)
			v.Category = &a
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

func allAssociationColumnNamesQuoted(q quote.Quoter) string {
	return strings.Join(q.QuoteN(listOfAssociationTableColumnNames), ",")
}

//--------------------------------------------------------------------------------

// GetAssociationById gets the record with a given primary key value.
// If not found, *Association will be nil.
func (tbl AssociationTable) GetAssociationById(ctx context.Context, req require.Requirement, id int64) (*Association, error) {
	return tbl.SelectOne(ctx, req, where.Eq("id", id), nil)
}

// GetAssociationsById gets records from the table according to a list of primary keys.
// Although the list of ids can be arbitrarily long, there are practical limits;
// note that Oracle DB has a limit of 1000.
//
// It places a requirement, which may be nil, on the size of the expected results: in particular, require.All
// controls whether an error is generated not all the ids produce a result.
func (tbl AssociationTable) GetAssociationsById(ctx context.Context, req require.Requirement, qc where.QueryConstraint, id ...int64) (list []*Association, err error) {
	if req == require.All {
		req = require.Exactly(len(id))
	}
	return tbl.Select(ctx, req, where.In("id", id), qc)
}

func doAssociationTableQueryAndScanOne(ctx context.Context, tbl AssociationTabler, req require.Requirement, query string, args ...interface{}) (*Association, error) {
	list, err := doAssociationTableQueryAndScan(ctx, tbl, req, true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

// Fetch fetches a list of Association based on a supplied query. This is mostly used for join queries that map its
// result columns to the fields of Association. Other queries might be better handled by GetXxx or Select methods.
func (tbl AssociationTable) Fetch(ctx context.Context, req require.Requirement, query string, args ...interface{}) ([]*Association, error) {
	return doAssociationTableQueryAndScan(ctx, tbl, req, false, query, args...)
}

//-------------------------------------------------------------------------------------------------

// SelectOneWhere allows a single Association to be obtained from the table that matches a 'where' clause
// and some limit. Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
// If not found, *Example will be nil.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.One
// controls whether an error is generated when no result is found.
//
// The args are for any placeholder parameters in the query.
func (tbl AssociationTable) SelectOneWhere(ctx context.Context, req require.Requirement, where, orderBy string, args ...interface{}) (*Association, error) {
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT 1",
		allAssociationColumnNamesQuoted(tbl.Dialect().Quoter()), quotedName, where, orderBy)
	v, err := doAssociationTableQueryAndScanOne(ctx, tbl, req, query, args...)
	return v, err
}

// SelectOne allows a single Association to be obtained from the table that matches a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
// If not found, *Example will be nil.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.One
// controls whether an error is generated when no result is found.
func (tbl AssociationTable) SelectOne(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) (*Association, error) {
	q := tbl.Dialect().Quoter()
	whs, args := where.Where(wh, q)
	orderBy := where.Build(qc, q)
	return tbl.SelectOneWhere(ctx, req, whs, orderBy, args...)
}

// SelectWhere allows Associations to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
//
// The args are for any placeholder parameters in the query.
func (tbl AssociationTable) SelectWhere(ctx context.Context, req require.Requirement, where, orderBy string, args ...interface{}) ([]*Association, error) {
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("SELECT %s FROM %s %s %s",
		allAssociationColumnNamesQuoted(tbl.Dialect().Quoter()), quotedName, where, orderBy)
	vv, err := doAssociationTableQueryAndScan(ctx, tbl, req, false, query, args...)
	return vv, err
}

// Select allows Associations to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
func (tbl AssociationTable) Select(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]*Association, error) {
	q := tbl.Dialect().Quoter()
	whs, args := where.Where(wh, q)
	orderBy := where.Build(qc, q)
	return tbl.SelectWhere(ctx, req, whs, orderBy, args...)
}

//-------------------------------------------------------------------------------------------------

// CountWhere counts Associations in the table that match a 'where' clause.
// Use a blank string for the 'where' argument if it is not needed.
//
// The args are for any placeholder parameters in the query.
func (tbl AssociationTable) CountWhere(ctx context.Context, where string, args ...interface{}) (count int64, err error) {
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

// Count counts the Associations in the table that match a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed.
func (tbl AssociationTable) Count(ctx context.Context, wh where.Expression) (count int64, err error) {
	whs, args := where.Where(wh, tbl.Dialect().Quoter())
	return tbl.CountWhere(ctx, whs, args...)
}

//--------------------------------------------------------------------------------

// SliceId gets the id column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AssociationTable) SliceId(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return support.SliceInt64List(ctx, tbl, req, tbl.pk, wh, qc)
}

// SliceName gets the name column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AssociationTable) SliceName(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return support.SliceStringPtrList(ctx, tbl, req, "name", wh, qc)
}

// SliceRef1 gets the ref1 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AssociationTable) SliceRef1(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return support.SliceInt64PtrList(ctx, tbl, req, "ref1", wh, qc)
}

// SliceRef2 gets the ref2 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AssociationTable) SliceRef2(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return support.SliceInt64PtrList(ctx, tbl, req, "ref2", wh, qc)
}

// SliceQuality gets the quality column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AssociationTable) SliceQuality(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]QualName, error) {
	return sliceAssociationTableQualNamePtrList(ctx, tbl, req, "quality", wh, qc)
}

// SliceCategory gets the category column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AssociationTable) SliceCategory(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]Category, error) {
	return sliceAssociationTableCategoryPtrList(ctx, tbl, req, "category", wh, qc)
}

func sliceAssociationTableCategoryPtrList(ctx context.Context, tbl AssociationTabler, req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]Category, error) {
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

func sliceAssociationTableQualNamePtrList(ctx context.Context, tbl AssociationTabler, req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]QualName, error) {
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

	list := make([]QualName, 0, 10)

	for rows.Next() {
		var v QualName
		err = rows.Scan(&v)
		if err == sql.ErrNoRows {
			return list, tbl.Logger().LogIfError(require.ErrorIfQueryNotSatisfiedBy(req, int64(len(list))))
		} else {
			list = append(list, v)
		}
	}
	return list, tbl.Logger().LogIfError(require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, int64(len(list))))
}

func constructAssociationTableInsert(tbl AssociationTable, w dialect.StringWriter, v *Association, withPk bool) (s []interface{}, err error) {
	q := tbl.Dialect().Quoter()
	s = make([]interface{}, 0, 6)

	comma := ""
	w.WriteString(" (")

	if withPk {
		q.QuoteW(w, "id")
		comma = ","
		s = append(s, v.Id)
	}

	if v.Name != nil {
		w.WriteString(comma)
		q.QuoteW(w, "name")
		s = append(s, v.Name)
		comma = ","
	}

	if v.Quality != nil {
		w.WriteString(comma)
		q.QuoteW(w, "quality")
		s = append(s, v.Quality)
		comma = ","
	}

	if v.Ref1 != nil {
		w.WriteString(comma)
		q.QuoteW(w, "ref1")
		s = append(s, v.Ref1)
		comma = ","
	}

	if v.Ref2 != nil {
		w.WriteString(comma)
		q.QuoteW(w, "ref2")
		s = append(s, v.Ref2)
		comma = ","
	}

	if v.Category != nil {
		w.WriteString(comma)
		q.QuoteW(w, "category")
		s = append(s, v.Category)
		comma = ","
	}

	w.WriteString(")")
	return s, nil
}

func constructAssociationTableUpdate(tbl AssociationTable, w dialect.StringWriter, v *Association) (s []interface{}, err error) {
	q := tbl.Dialect().Quoter()
	j := 1
	s = make([]interface{}, 0, 5)

	comma := ""

	w.WriteString(comma)
	if v.Name != nil {
		q.QuoteW(w, "name")
		w.WriteString("=?")
		s = append(s, v.Name)
		j++
		comma = ", "
	} else {
		q.QuoteW(w, "name")
		w.WriteString("=NULL")
	}

	w.WriteString(comma)
	if v.Quality != nil {
		q.QuoteW(w, "quality")
		w.WriteString("=?")
		s = append(s, v.Quality)
		j++
		comma = ", "
	} else {
		q.QuoteW(w, "quality")
		w.WriteString("=NULL")
	}

	w.WriteString(comma)
	if v.Ref1 != nil {
		q.QuoteW(w, "ref1")
		w.WriteString("=?")
		s = append(s, v.Ref1)
		j++
		comma = ", "
	} else {
		q.QuoteW(w, "ref1")
		w.WriteString("=NULL")
	}

	w.WriteString(comma)
	if v.Ref2 != nil {
		q.QuoteW(w, "ref2")
		w.WriteString("=?")
		s = append(s, v.Ref2)
		j++
		comma = ", "
	} else {
		q.QuoteW(w, "ref2")
		w.WriteString("=NULL")
	}

	w.WriteString(comma)
	if v.Category != nil {
		q.QuoteW(w, "category")
		w.WriteString("=?")
		s = append(s, v.Category)
		j++
		comma = ", "
	} else {
		q.QuoteW(w, "category")
		w.WriteString("=NULL")
	}
	return s, nil
}

//--------------------------------------------------------------------------------

// Insert adds new records for the Associations.// The Associations have their primary key fields set to the new record identifiers.
// The Association.PreInsert() method will be called, if it exists.
func (tbl AssociationTable) Insert(ctx context.Context, req require.Requirement, vv ...*Association) error {
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

		fields, err := constructAssociationTableInsert(tbl, b, v, false)
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
func (tbl AssociationTable) UpdateById(ctx context.Context, req require.Requirement, id int64, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("id", id), fields...)
}

// UpdateByName updates one or more columns, given a name value.
func (tbl AssociationTable) UpdateByName(ctx context.Context, req require.Requirement, name string, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("name", name), fields...)
}

// UpdateByQuality updates one or more columns, given a quality value.
func (tbl AssociationTable) UpdateByQuality(ctx context.Context, req require.Requirement, quality QualName, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("quality", quality), fields...)
}

// UpdateByRef1 updates one or more columns, given a ref1 value.
func (tbl AssociationTable) UpdateByRef1(ctx context.Context, req require.Requirement, ref1 int64, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("ref1", ref1), fields...)
}

// UpdateByRef2 updates one or more columns, given a ref2 value.
func (tbl AssociationTable) UpdateByRef2(ctx context.Context, req require.Requirement, ref2 int64, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("ref2", ref2), fields...)
}

// UpdateByCategory updates one or more columns, given a category value.
func (tbl AssociationTable) UpdateByCategory(ctx context.Context, req require.Requirement, category Category, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("category", category), fields...)
}

// UpdateFields updates one or more columns, given a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed (but note that this is risky!).
func (tbl AssociationTable) UpdateFields(ctx context.Context, req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error) {
	return support.UpdateFields(ctx, tbl, req, wh, fields...)
}

//--------------------------------------------------------------------------------

// Update updates records, matching them by primary key. It returns the number of rows affected.
// The Association.PreUpdate(Execer) method will be called, if it exists.
func (tbl AssociationTable) Update(ctx context.Context, req require.Requirement, vv ...*Association) (int64, error) {
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

		args, err := constructAssociationTableUpdate(tbl, b, v)
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
func (tbl AssociationTable) Upsert(ctx context.Context, v *Association, wh where.Expression) error {
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
func (tbl AssociationTable) DeleteById(ctx context.Context, req require.Requirement, id ...int64) (int64, error) {
	ii := support.Int64AsInterfaceSlice(id)
	return support.DeleteByColumn(ctx, tbl, req, "id", ii...)
}

// DeleteByName deletes rows from the table, given some name values.
// The list of ids can be arbitrarily long.
func (tbl AssociationTable) DeleteByName(ctx context.Context, req require.Requirement, name ...string) (int64, error) {
	ii := support.StringAsInterfaceSlice(name)
	return support.DeleteByColumn(ctx, tbl, req, "name", ii...)
}

// DeleteByQuality deletes rows from the table, given some quality values.
// The list of ids can be arbitrarily long.
func (tbl AssociationTable) DeleteByQuality(ctx context.Context, req require.Requirement, quality ...QualName) (int64, error) {
	ii := make([]interface{}, len(quality))
	for i, v := range quality {
		ii[i] = v
	}
	return support.DeleteByColumn(ctx, tbl, req, "quality", ii...)
}

// DeleteByRef1 deletes rows from the table, given some ref1 values.
// The list of ids can be arbitrarily long.
func (tbl AssociationTable) DeleteByRef1(ctx context.Context, req require.Requirement, ref1 ...int64) (int64, error) {
	ii := support.Int64AsInterfaceSlice(ref1)
	return support.DeleteByColumn(ctx, tbl, req, "ref1", ii...)
}

// DeleteByRef2 deletes rows from the table, given some ref2 values.
// The list of ids can be arbitrarily long.
func (tbl AssociationTable) DeleteByRef2(ctx context.Context, req require.Requirement, ref2 ...int64) (int64, error) {
	ii := support.Int64AsInterfaceSlice(ref2)
	return support.DeleteByColumn(ctx, tbl, req, "ref2", ii...)
}

// DeleteByCategory deletes rows from the table, given some category values.
// The list of ids can be arbitrarily long.
func (tbl AssociationTable) DeleteByCategory(ctx context.Context, req require.Requirement, category ...Category) (int64, error) {
	ii := make([]interface{}, len(category))
	for i, v := range category {
		ii[i] = v
	}
	return support.DeleteByColumn(ctx, tbl, req, "category", ii...)
}

// Delete deletes one or more rows from the table, given a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl AssociationTable) Delete(ctx context.Context, req require.Requirement, wh where.Expression) (int64, error) {
	query, args := deleteRowsAssociationTableSql(tbl, wh)
	return tbl.Exec(ctx, req, query, args...)
}

func deleteRowsAssociationTableSql(tbl AssociationTabler, wh where.Expression) (string, []interface{}) {
	whs, args := where.Where(wh, tbl.Dialect().Quoter())
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("DELETE FROM %s %s", quotedName, whs)
	return query, args
}

//-------------------------------------------------------------------------------------------------
