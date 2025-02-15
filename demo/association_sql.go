// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.
// sqlapi v0.62.7; sqlgen v0.79.0-1-g994aea0

package demo

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/rickb777/sqlapi"
	"github.com/rickb777/sqlapi/constraint"
	"github.com/rickb777/sqlapi/driver"
	"github.com/rickb777/sqlapi/require"
	"github.com/rickb777/sqlapi/support"
	"github.com/rickb777/where"
	"github.com/rickb777/where/dialect"
	"github.com/rickb777/where/quote"
	"io"
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

	// WithContext returns a modified AssociationTabler with a given context.
	WithContext(ctx context.Context) AssociationTabler

	// CreateTable creates the table.
	CreateTable(ifNotExists bool) (int64, error)

	// DropTable drops the table, destroying all its data.
	DropTable(ifExists bool) (int64, error)

	// Truncate drops every record from the table, if possible.
	Truncate(force bool) (err error)
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
	Transact(txOptions *pgx.TxOptions, fn func(AssociationQueryer) error) error

	// Exec executes a query without returning any rows.
	Exec(req require.Requirement, query string, args ...interface{}) (int64, error)

	// Query is the low-level request method for this table using an SQL query that must return all the columns
	// necessary for Association values.
	Query(req require.Requirement, query string, args ...interface{}) ([]*Association, error)

	// QueryOneNullString is a low-level access method for one string, returning the first match.
	QueryOneNullString(req require.Requirement, query string, args ...interface{}) (result sql.NullString, err error)

	// QueryOneNullInt64 is a low-level access method for one int64, returning the first match.
	QueryOneNullInt64(req require.Requirement, query string, args ...interface{}) (result sql.NullInt64, err error)

	// QueryOneNullFloat64 is a low-level access method for one float64, returning the first match.
	QueryOneNullFloat64(req require.Requirement, query string, args ...interface{}) (result sql.NullFloat64, err error)

	// GetAssociationById gets the record with a given primary key value.
	GetAssociationById(req require.Requirement, id int64) (*Association, error)

	// GetAssociationsById gets records from the table according to a list of primary keys.
	GetAssociationsById(req require.Requirement, qc where.QueryConstraint, id ...int64) (list []*Association, err error)

	// Fetch fetches a list of Association based on a supplied query. This is mostly used for join queries that map its
	// result columns to the fields of Association. Other queries might be better handled by GetXxx or Select methods.
	Fetch(req require.Requirement, query string, args ...interface{}) ([]*Association, error)

	// SelectOneWhere allows a single Association to be obtained from the table that matches a 'where' clause.
	SelectOneWhere(req require.Requirement, where, orderBy string, args ...interface{}) (*Association, error)

	// SelectOne allows a single Association to be obtained from the table that matches a 'where' clause.
	SelectOne(req require.Requirement, wh where.Expression, qc where.QueryConstraint) (*Association, error)

	// SelectWhere allows Associations to be obtained from the table that match a 'where' clause.
	SelectWhere(req require.Requirement, where, orderBy string, args ...interface{}) ([]*Association, error)

	// Select allows Associations to be obtained from the table that match a 'where' clause.
	Select(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]*Association, error)

	// CountWhere counts Associations in the table that match a 'where' clause.
	CountWhere(where string, args ...interface{}) (count int64, err error)

	// Count counts the Associations in the table that match a 'where' clause.
	Count(wh where.Expression) (count int64, err error)

	// SliceID gets the id column for all rows that match the 'where' condition.
	SliceID(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error)

	// SliceName gets the name column for all rows that match the 'where' condition.
	SliceName(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error)

	// SliceRef1 gets the ref1 column for all rows that match the 'where' condition.
	SliceRef1(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error)

	// SliceRef2 gets the ref2 column for all rows that match the 'where' condition.
	SliceRef2(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error)

	// SliceQuality gets the quality column for all rows that match the 'where' condition.
	SliceQuality(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]QualName, error)

	// SliceCategory gets the category column for all rows that match the 'where' condition.
	SliceCategory(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]Category, error)

	// Insert adds new records for the Associations, setting the primary key field for each one.
	Insert(req require.Requirement, vv ...*Association) error

	// UpdateByID updates one or more columns, given a id value.
	UpdateByID(req require.Requirement, id int64, fields ...sql.NamedArg) (int64, error)

	// UpdateByName updates one or more columns, given a name value.
	UpdateByName(req require.Requirement, name string, fields ...sql.NamedArg) (int64, error)

	// UpdateByQuality updates one or more columns, given a quality value.
	UpdateByQuality(req require.Requirement, quality QualName, fields ...sql.NamedArg) (int64, error)

	// UpdateByRef1 updates one or more columns, given a ref1 value.
	UpdateByRef1(req require.Requirement, ref1 int64, fields ...sql.NamedArg) (int64, error)

	// UpdateByRef2 updates one or more columns, given a ref2 value.
	UpdateByRef2(req require.Requirement, ref2 int64, fields ...sql.NamedArg) (int64, error)

	// UpdateByCategory updates one or more columns, given a category value.
	UpdateByCategory(req require.Requirement, category Category, fields ...sql.NamedArg) (int64, error)

	// UpdateFields updates one or more columns, given a 'where' clause.
	UpdateFields(req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error)

	// Update updates records, matching them by primary key.
	Update(req require.Requirement, vv ...*Association) (int64, error)

	// Upsert inserts or updates a record, matching it using the expression supplied.
	// This expression is used to search for an existing record based on some specified
	// key column(s). It must match either zero or one existing record. If it matches
	// none, a new record is inserted; otherwise the matching record is updated. An
	// error results if these conditions are not met.
	Upsert(v *Association, wh where.Expression) error

	// DeleteByID deletes rows from the table, given some id values.
	// The list of ids can be arbitrarily long.
	DeleteByID(req require.Requirement, id ...int64) (int64, error)

	// DeleteByName deletes rows from the table, given some name values.
	// The list of ids can be arbitrarily long.
	DeleteByName(req require.Requirement, name ...string) (int64, error)

	// DeleteByQuality deletes rows from the table, given some quality values.
	// The list of ids can be arbitrarily long.
	DeleteByQuality(req require.Requirement, quality ...QualName) (int64, error)

	// DeleteByRef1 deletes rows from the table, given some ref1 values.
	// The list of ids can be arbitrarily long.
	DeleteByRef1(req require.Requirement, ref1 ...int64) (int64, error)

	// DeleteByRef2 deletes rows from the table, given some ref2 values.
	// The list of ids can be arbitrarily long.
	DeleteByRef2(req require.Requirement, ref2 ...int64) (int64, error)

	// DeleteByCategory deletes rows from the table, given some category values.
	// The list of ids can be arbitrarily long.
	DeleteByCategory(req require.Requirement, category ...Category) (int64, error)

	// Delete deletes one or more rows from the table, given a 'where' clause.
	// Use a nil value for the 'wh' argument if it is not needed (very risky!).
	Delete(req require.Requirement, wh where.Expression) (int64, error)
}

//-------------------------------------------------------------------------------------------------

// AssociationTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type AssociationTable struct {
	sqlapi.CoreTable
	constraints constraint.Constraints
	ctx         context.Context
	pk          string
}

// Type conformance checks
var _ sqlapi.TableCreator = &AssociationTable{}

// NewAssociationTable returns a new table instance.
// If a blank table name is supplied, the default name "associations" will be used instead.
// The request context is initialised with the background.
func NewAssociationTable(name string, d sqlapi.SqlDB) AssociationTable {
	if name == "" {
		name = "associations"
	}
	var constraints constraint.Constraints
	return AssociationTable{
		CoreTable: sqlapi.CoreTable{
			Nm: sqlapi.TableName{Prefix: "", Name: name},
			Ex: d,
		},
		constraints: constraints,
		ctx:         context.Background(),
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
		CoreTable: sqlapi.CoreTable{
			Nm: origin.Name(),
			Ex: origin.Execer(),
		},
		constraints: nil,
		ctx:         origin.Ctx(),
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
	tbl.Nm.Prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries via this table.
// The result is a modified copy of the table; the original is unchanged.
func (tbl AssociationTable) WithContext(ctx context.Context) AssociationTabler {
	tbl.ctx = ctx
	return tbl
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

// Ctx gets the current request context.
func (tbl AssociationTable) Ctx() context.Context {
	return tbl.ctx
}

// PkColumn gets the column name used as a primary key.
func (tbl AssociationTable) PkColumn() string {
	return tbl.pk
}

// Using returns a modified AssociationTabler using the the Execer supplied,
// which will typically be a transaction (i.e. SqlTx). This is needed when making multiple
// queries across several tables within a single transaction.
//
// The result is a modified copy of the table; the original is unchanged.
func (tbl AssociationTable) Using(tx sqlapi.Execer) AssociationQueryer {
	tbl.Ex = tx
	return tbl
}

// Transact runs the function provided within a transaction. If the function completes without error,
// the transaction is committed. If there is an error or a panic, the transaction is rolled back.
//
// The options can be nil, in which case the default behaviour is that of the underlying connection.
//
// Nested transactions (i.e. within 'fn') are permitted: they execute within the outermost transaction.
// Therefore they do not commit until the outermost transaction commits.
func (tbl AssociationTable) Transact(txOptions *pgx.TxOptions, fn func(AssociationQueryer) error) error {
	var err error
	if tbl.IsTx() {
		err = fn(tbl) // nested transactions are inlined
	} else {
		err = tbl.DB().Transact(tbl.ctx, txOptions, func(tx sqlapi.SqlTx) error {
			return fn(tbl.Using(tx))
		})
	}
	return tbl.Logger().LogIfError(tbl.Ctx(), err)
}

func (tbl AssociationTable) quotedName() string {
	return tbl.Dialect().Quoter().Quote(tbl.Nm.String())
}

func (tbl AssociationTable) quotedNameW(w driver.StringWriter) {
	tbl.Dialect().Quoter().QuoteW(w, tbl.Nm.String())
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

//-------------------------------------------------------------------------------------------------

// CreateTable creates the table.
func (tbl AssociationTable) CreateTable(ifNotExists bool) (int64, error) {
	return support.Exec(tbl, nil, createAssociationTableSql(tbl, ifNotExists))
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
	case dialect.Sqlite:
		columns = sqlAssociationTableCreateColumnsSqlite
	case dialect.Mysql:
		columns = sqlAssociationTableCreateColumnsMysql
	case dialect.Postgres:
		columns = sqlAssociationTableCreateColumnsPostgres
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
func (tbl AssociationTable) DropTable(ifExists bool) (int64, error) {
	return support.Exec(tbl, nil, dropAssociationTableSql(tbl, ifExists))
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
func (tbl AssociationTable) Truncate(force bool) (err error) {
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
func (tbl AssociationTable) Exec(req require.Requirement, query string, args ...interface{}) (int64, error) {
	return support.Exec(tbl, req, query, args...)
}

//-------------------------------------------------------------------------------------------------

// Query is the low-level request method for this table. The SQL query must return all the columns necessary for
// Association values. Placeholders should be vanilla '?' marks, which will be replaced if necessary according to
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
func (tbl AssociationTable) Query(req require.Requirement, query string, args ...interface{}) ([]*Association, error) {
	return doAssociationTableQueryAndScan(tbl, req, false, query, args)
}

func doAssociationTableQueryAndScan(tbl AssociationTabler, req require.Requirement, firstOnly bool, query string, args ...interface{}) ([]*Association, error) {
	rows, err := support.Query(tbl, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vv, n, err := ScanAssociations(query, rows, firstOnly)
	return vv, tbl.Logger().LogIfError(tbl.Ctx(), require.ChainErrorIfQueryNotSatisfiedBy(err, req, n))
}

//-------------------------------------------------------------------------------------------------

// QueryOneNullString is a low-level access method for one string. This can be used for function queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl AssociationTable) QueryOneNullString(req require.Requirement, query string, args ...interface{}) (result sql.NullString, err error) {
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
func (tbl AssociationTable) QueryOneNullInt64(req require.Requirement, query string, args ...interface{}) (result sql.NullInt64, err error) {
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
func (tbl AssociationTable) QueryOneNullFloat64(req require.Requirement, query string, args ...interface{}) (result sql.NullFloat64, err error) {
	err = support.QueryOneNullThing(tbl, req, &result, query, args...)
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
func (tbl AssociationTable) GetAssociationById(req require.Requirement, id int64) (*Association, error) {
	return tbl.SelectOne(req, where.Eq("id", id), nil)
}

// GetAssociationsById gets records from the table according to a list of primary keys.
// Although the list of ids can be arbitrarily long, there are practical limits;
// note that Oracle DB has a limit of 1000.
//
// It places a requirement, which may be nil, on the size of the expected results: in particular, require.All
// controls whether an error is generated not all the ids produce a result.
func (tbl AssociationTable) GetAssociationsById(req require.Requirement, qc where.QueryConstraint, id ...int64) (list []*Association, err error) {
	if req == require.All {
		req = require.Exactly(len(id))
	}
	return tbl.Select(req, where.InSlice("id", id), qc)
}

func doAssociationTableQueryAndScanOne(tbl AssociationTabler, req require.Requirement, query string, args ...interface{}) (*Association, error) {
	list, err := doAssociationTableQueryAndScan(tbl, req, true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

// Fetch fetches a list of Association based on a supplied query. This is mostly used for join queries that map its
// result columns to the fields of Association. Other queries might be better handled by GetXxx or Select methods.
func (tbl AssociationTable) Fetch(req require.Requirement, query string, args ...interface{}) ([]*Association, error) {
	return doAssociationTableQueryAndScan(tbl, req, false, query, args...)
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
func (tbl AssociationTable) SelectOneWhere(req require.Requirement, where, orderBy string, args ...interface{}) (*Association, error) {
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT 1",
		allAssociationColumnNamesQuoted(tbl.Dialect().Quoter()), quotedName, where, orderBy)
	v, err := doAssociationTableQueryAndScanOne(tbl, req, query, args...)
	return v, err
}

// SelectOne allows a single Association to be obtained from the table that matches a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
// If not found, *Example will be nil.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.One
// controls whether an error is generated when no result is found.
func (tbl AssociationTable) SelectOne(req require.Requirement, wh where.Expression, qc where.QueryConstraint) (*Association, error) {
	d := tbl.Dialect()
	whs, args := where.Where(wh, d.Quoter())
	orderBy := where.Build(qc, d.Index())
	return tbl.SelectOneWhere(req, whs, orderBy, args...)
}

// SelectWhere allows Associations to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
//
// The args are for any placeholder parameters in the query.
func (tbl AssociationTable) SelectWhere(req require.Requirement, where, orderBy string, args ...interface{}) ([]*Association, error) {
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("SELECT %s FROM %s %s %s",
		allAssociationColumnNamesQuoted(tbl.Dialect().Quoter()), quotedName, where, orderBy)
	vv, err := doAssociationTableQueryAndScan(tbl, req, false, query, args...)
	return vv, err
}

// Select allows Associations to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
func (tbl AssociationTable) Select(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]*Association, error) {
	d := tbl.Dialect()
	whs, args := where.Where(wh, d.Quoter())
	orderBy := where.Build(qc, d.Index())
	return tbl.SelectWhere(req, whs, orderBy, args...)
}

//-------------------------------------------------------------------------------------------------

// CountWhere counts Associations in the table that match a 'where' clause.
// Use a blank string for the 'where' argument if it is not needed.
//
// The args are for any placeholder parameters in the query.
func (tbl AssociationTable) CountWhere(where string, args ...interface{}) (count int64, err error) {
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
	return count, tbl.Logger().LogIfError(tbl.Ctx(), err)
}

// Count counts the Associations in the table that match a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed.
func (tbl AssociationTable) Count(wh where.Expression) (count int64, err error) {
	whs, args := where.Where(wh, tbl.Dialect().Quoter())
	return tbl.CountWhere(whs, args...)
}

//--------------------------------------------------------------------------------

// SliceID gets the id column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AssociationTable) SliceID(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return support.SliceInt64List(tbl, req, tbl.pk, wh, qc)
}

// SliceName gets the name column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AssociationTable) SliceName(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return support.SliceStringPtrList(tbl, req, "name", wh, qc)
}

// SliceRef1 gets the ref1 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AssociationTable) SliceRef1(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return support.SliceInt64PtrList(tbl, req, "ref1", wh, qc)
}

// SliceRef2 gets the ref2 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AssociationTable) SliceRef2(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return support.SliceInt64PtrList(tbl, req, "ref2", wh, qc)
}

// SliceQuality gets the quality column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AssociationTable) SliceQuality(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]QualName, error) {
	return sliceAssociationTableQualNamePtrList(tbl, req, "quality", wh, qc)
}

// SliceCategory gets the category column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AssociationTable) SliceCategory(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]Category, error) {
	return sliceAssociationTableCategoryPtrList(tbl, req, "category", wh, qc)
}

func sliceAssociationTableCategoryPtrList(tbl AssociationTabler, req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]Category, error) {
	d := tbl.Dialect()
	whs, args := where.Where(wh, d.Quoter())
	orderBy := where.Build(qc, d.Index())
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", d.Quoter().Quote(sqlname), quotedName, whs, orderBy)
	rows, err := support.Query(tbl, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]Category, 0, 10)

	for rows.Next() {
		var v Category
		err = rows.Scan(&v)
		if err == sql.ErrNoRows {
			return list, tbl.Logger().LogIfError(tbl.Ctx(), require.ErrorIfQueryNotSatisfiedBy(req, int64(len(list))))
		} else {
			list = append(list, v)
		}
	}
	return list, tbl.Logger().LogIfError(tbl.Ctx(), require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, int64(len(list))))
}

func sliceAssociationTableQualNamePtrList(tbl AssociationTabler, req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]QualName, error) {
	d := tbl.Dialect()
	whs, args := where.Where(wh, d.Quoter())
	orderBy := where.Build(qc, d.Index())
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", d.Quoter().Quote(sqlname), quotedName, whs, orderBy)
	rows, err := support.Query(tbl, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]QualName, 0, 10)

	for rows.Next() {
		var v QualName
		err = rows.Scan(&v)
		if err == sql.ErrNoRows {
			return list, tbl.Logger().LogIfError(tbl.Ctx(), require.ErrorIfQueryNotSatisfiedBy(req, int64(len(list))))
		} else {
			list = append(list, v)
		}
	}
	return list, tbl.Logger().LogIfError(tbl.Ctx(), require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, int64(len(list))))
}

func constructAssociationTableInsert(tbl AssociationTable, w io.StringWriter, v *Association, withPk bool) (s []interface{}, err error) {
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

func constructAssociationTableUpdate(tbl AssociationTable, w io.StringWriter, v *Association) (s []interface{}, err error) {
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
func (tbl AssociationTable) Insert(req require.Requirement, vv ...*Association) error {
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
				return tbl.Logger().LogError(tbl.Ctx(), err)
			}
		}

		b := driver.Adapt(&bytes.Buffer{})
		b.WriteString("INSERT INTO ")
		tbl.quotedNameW(b)

		fields, err := constructAssociationTableInsert(tbl, b, v, false)
		if err != nil {
			return tbl.Logger().LogError(tbl.Ctx(), err)
		}

		b.WriteString(" VALUES (")
		b.WriteString(tbl.Dialect().Placeholders(len(fields)))
		b.WriteString(")")
		b.WriteString(returning)

		query := b.String()
		tbl.Logger().LogQuery(tbl.Ctx(), query, fields...)

		var n int64 = 1
		if insertHasReturningPhrase {
			row := tbl.Execer().QueryRow(tbl.ctx, query, fields...)
			var i64 int64
			err = row.Scan(&i64)
			v.Id = i64

		} else {
			i64, e2 := tbl.Execer().Insert(tbl.ctx, tbl.pk, query, fields...)
			if e2 != nil {
				return tbl.Logger().LogError(tbl.Ctx(), e2)
			}
			v.Id = i64
		}

		if err != nil {
			return tbl.Logger().LogError(tbl.Ctx(), err)
		}
		count += n
	}

	return tbl.Logger().LogIfError(tbl.Ctx(), require.ErrorIfExecNotSatisfiedBy(req, count))
}

// UpdateByID updates one or more columns, given a id value.
func (tbl AssociationTable) UpdateByID(req require.Requirement, id int64, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(req, where.Eq("id", id), fields...)
}

// UpdateByName updates one or more columns, given a name value.
func (tbl AssociationTable) UpdateByName(req require.Requirement, name string, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(req, where.Eq("name", name), fields...)
}

// UpdateByQuality updates one or more columns, given a quality value.
func (tbl AssociationTable) UpdateByQuality(req require.Requirement, quality QualName, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(req, where.Eq("quality", quality), fields...)
}

// UpdateByRef1 updates one or more columns, given a ref1 value.
func (tbl AssociationTable) UpdateByRef1(req require.Requirement, ref1 int64, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(req, where.Eq("ref1", ref1), fields...)
}

// UpdateByRef2 updates one or more columns, given a ref2 value.
func (tbl AssociationTable) UpdateByRef2(req require.Requirement, ref2 int64, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(req, where.Eq("ref2", ref2), fields...)
}

// UpdateByCategory updates one or more columns, given a category value.
func (tbl AssociationTable) UpdateByCategory(req require.Requirement, category Category, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(req, where.Eq("category", category), fields...)
}

// UpdateFields updates one or more columns, given a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed (but note that this is risky!).
func (tbl AssociationTable) UpdateFields(req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error) {
	return support.UpdateFields(tbl, req, wh, fields...)
}

//--------------------------------------------------------------------------------

// Update updates records, matching them by primary key. It returns the number of rows affected.
// The Association.PreUpdate(Execer) method will be called, if it exists.
func (tbl AssociationTable) Update(req require.Requirement, vv ...*Association) (int64, error) {
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
				return count, tbl.Logger().LogError(tbl.Ctx(), err)
			}
		}

		b := driver.Adapt(&bytes.Buffer{})
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
		n, err := tbl.Exec(nil, query, args...)
		if err != nil {
			return count, err
		}
		count += n
	}

	return count, tbl.Logger().LogIfError(tbl.Ctx(), require.ErrorIfExecNotSatisfiedBy(req, count))
}

//--------------------------------------------------------------------------------

// Upsert inserts or updates a record, matching it using the expression supplied.
// This expression is used to search for an existing record based on some specified
// key column(s). It must match either zero or one existing record. If it matches
// none, a new record is inserted; otherwise the matching record is updated. An
// error results if these conditions are not met.
func (tbl AssociationTable) Upsert(v *Association, wh where.Expression) error {
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

	var id int64
	err = rows.Scan(&id)
	if err != nil {
		return tbl.Logger().LogIfError(tbl.Ctx(), err)
	}

	if rows.Next() {
		return require.ErrWrongSize(2, "expected to find no more than 1 but got at least 2 using %q", wh)
	}

	v.Id = id
	_, err = tbl.Update(require.One, v)
	return err
}

//-------------------------------------------------------------------------------------------------

// DeleteByID deletes rows from the table, given some id values.
// The list of ids can be arbitrarily long.
func (tbl AssociationTable) DeleteByID(req require.Requirement, id ...int64) (int64, error) {
	ii := support.Int64AsInterfaceSlice(id)
	return support.DeleteByColumn(tbl, req, "id", ii...)
}

// DeleteByName deletes rows from the table, given some name values.
// The list of ids can be arbitrarily long.
func (tbl AssociationTable) DeleteByName(req require.Requirement, name ...string) (int64, error) {
	ii := support.StringAsInterfaceSlice(name)
	return support.DeleteByColumn(tbl, req, "name", ii...)
}

// DeleteByQuality deletes rows from the table, given some quality values.
// The list of ids can be arbitrarily long.
func (tbl AssociationTable) DeleteByQuality(req require.Requirement, quality ...QualName) (int64, error) {
	ii := make([]interface{}, len(quality))
	for i, v := range quality {
		ii[i] = v
	}
	return support.DeleteByColumn(tbl, req, "quality", ii...)
}

// DeleteByRef1 deletes rows from the table, given some ref1 values.
// The list of ids can be arbitrarily long.
func (tbl AssociationTable) DeleteByRef1(req require.Requirement, ref1 ...int64) (int64, error) {
	ii := support.Int64AsInterfaceSlice(ref1)
	return support.DeleteByColumn(tbl, req, "ref1", ii...)
}

// DeleteByRef2 deletes rows from the table, given some ref2 values.
// The list of ids can be arbitrarily long.
func (tbl AssociationTable) DeleteByRef2(req require.Requirement, ref2 ...int64) (int64, error) {
	ii := support.Int64AsInterfaceSlice(ref2)
	return support.DeleteByColumn(tbl, req, "ref2", ii...)
}

// DeleteByCategory deletes rows from the table, given some category values.
// The list of ids can be arbitrarily long.
func (tbl AssociationTable) DeleteByCategory(req require.Requirement, category ...Category) (int64, error) {
	ii := make([]interface{}, len(category))
	for i, v := range category {
		ii[i] = v
	}
	return support.DeleteByColumn(tbl, req, "category", ii...)
}

// Delete deletes one or more rows from the table, given a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl AssociationTable) Delete(req require.Requirement, wh where.Expression) (int64, error) {
	query, args := deleteRowsAssociationTableSql(tbl, wh)
	return tbl.Exec(req, query, args...)
}

func deleteRowsAssociationTableSql(tbl AssociationTabler, wh where.Expression) (string, []interface{}) {
	whs, args := where.Where(wh, tbl.Dialect().Quoter())
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("DELETE FROM %s %s", quotedName, whs)
	return query, args
}

//-------------------------------------------------------------------------------------------------
