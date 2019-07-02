// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.
// sqlapi v0.29.0; sqlgen v0.49.0-2-g18aa8b9

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

// AssociationTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type AssociationTable struct {
	name        sqlapi.TableName
	database    sqlapi.Database
	db          sqlapi.Execer
	constraints constraint.Constraints
	ctx         context.Context
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
		ctx:         context.Background(),
		pk:          "id",
	}
}

// CopyTableAsAssociationTable copies a table instance, retaining the name etc but
// providing methods appropriate for 'Association'. It doesn't copy the constraints of the original table.
//
// It serves to provide methods appropriate for 'Association'. This is most useful when this is used to represent a
// join result. In such cases, there won't be any need for DDL methods, nor Exec, Insert, Update or Delete.
func CopyTableAsAssociationTable(origin sqlapi.Table) AssociationTable {
	return AssociationTable{
		name:        origin.Name(),
		database:    origin.Database(),
		db:          origin.DB(),
		constraints: nil,
		ctx:         context.Background(),
		pk:          "id",
	}
}

// SetPkColumn sets the name of the primary key column. It defaults to "id".
// The result is a modified copy of the table; the original is unchanged.
func (tbl AssociationTable) SetPkColumn(pk string) AssociationTable {
	tbl.pk = pk
	return tbl
}

// WithPrefix sets the table name prefix for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl AssociationTable) WithPrefix(pfx string) AssociationTable {
	tbl.name.Prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries via this table.
// The result is a modified copy of the table; the original is unchanged.
//
// The shared context in the *Database is not altered by this method. So it
// is possible to use different contexts for different (groups of) queries.
func (tbl AssociationTable) WithContext(ctx context.Context) AssociationTable {
	tbl.ctx = ctx
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

// WithConstraint returns a modified Table with added data consistency constraints.
func (tbl AssociationTable) WithConstraint(cc ...constraint.Constraint) AssociationTable {
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
func (tbl AssociationTable) BeginTx(opts *sql.TxOptions) (AssociationTable, error) {
	var err error
	tbl.db, err = tbl.db.(sqlapi.SqlDB).BeginTx(tbl.ctx, opts)
	return tbl, tbl.Logger().LogIfError(err)
}

// Using returns a modified Table using the transaction supplied. This is needed
// when making multiple queries across several tables within a single transaction.
// The result is a modified copy of the table; the original is unchanged.
func (tbl AssociationTable) Using(tx sqlapi.SqlTx) AssociationTable {
	tbl.db = tx
	return tbl
}

func (tbl AssociationTable) quotedName() string {
	return tbl.Dialect().Quoter().Quote(tbl.name.String())
}

func (tbl AssociationTable) quotedNameW(w dialect.StringWriter) {
	tbl.Dialect().Quoter().QuoteW(w, tbl.name.String())
}

//--------------------------------------------------------------------------------

// NumAssociationTableColumns is the total number of columns in AssociationTable.
const NumAssociationTableColumns = 6

// NumAssociationTableDataColumns is the number of columns in AssociationTable not including the auto-increment key.
const NumAssociationTableDataColumns = 5

// AssociationTableColumnNames is the list of columns in AssociationTable.
const AssociationTableColumnNames = "id,name,quality,ref1,ref2,category"

// AssociationTableDataColumnNames is the list of data columns in AssociationTable.
const AssociationTableDataColumnNames = "name,quality,ref1,ref2,category"

var listOfAssociationTableColumnNames = strings.Split(AssociationTableColumnNames, ",")

//--------------------------------------------------------------------------------

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

//--------------------------------------------------------------------------------

// CreateTable creates the table.
func (tbl AssociationTable) CreateTable(ifNotExists bool) (int64, error) {
	return support.Exec(tbl, nil, tbl.createTableSql(ifNotExists))
}

func (tbl AssociationTable) createTableSql(ifNotExists bool) string {
	buf := &bytes.Buffer{}
	buf.WriteString("CREATE TABLE ")
	if ifNotExists {
		buf.WriteString("IF NOT EXISTS ")
	}
	q := tbl.Dialect().Quoter()
	q.QuoteW(buf, tbl.name.String())
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

	for i, c := range tbl.constraints {
		buf.WriteString(",\n ")
		buf.WriteString(c.ConstraintSql(tbl.Dialect().Quoter(), tbl.name, i+1))
	}

	buf.WriteString("\n)")
	buf.WriteString(tbl.Dialect().CreateTableSettings())
	return buf.String()
}

func (tbl AssociationTable) ternary(flag bool, a, b string) string {
	if flag {
		return a
	}
	return b
}

// DropTable drops the table, destroying all its data.
func (tbl AssociationTable) DropTable(ifExists bool) (int64, error) {
	return support.Exec(tbl, nil, tbl.dropTableSql(ifExists))
}

func (tbl AssociationTable) dropTableSql(ifExists bool) string {
	ie := tbl.ternary(ifExists, "IF EXISTS ", "")
	query := fmt.Sprintf("DROP TABLE %s%s", ie, tbl.quotedName())
	return query
}

//--------------------------------------------------------------------------------

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

//--------------------------------------------------------------------------------

// Exec executes a query without returning any rows.
// It returns the number of rows affected (if the database driver supports this).
//
// The args are for any placeholder parameters in the query.
func (tbl AssociationTable) Exec(req require.Requirement, query string, args ...interface{}) (int64, error) {
	return support.Exec(tbl, req, query, args...)
}

//--------------------------------------------------------------------------------

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
	return tbl.doQueryAndScan(req, false, query, args)
}

func (tbl AssociationTable) doQueryAndScan(req require.Requirement, firstOnly bool, query string, args ...interface{}) ([]*Association, error) {
	rows, err := support.Query(tbl, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vv, n, err := ScanAssociations(query, rows, firstOnly)
	return vv, tbl.Logger().LogIfError(require.ChainErrorIfQueryNotSatisfiedBy(err, req, n))
}

//--------------------------------------------------------------------------------

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

// GetAssociationsById gets records from the table according to a list of primary keys.
// Although the list of ids can be arbitrarily long, there are practical limits;
// note that Oracle DB has a limit of 1000.
//
// It places a requirement, which may be nil, on the size of the expected results: in particular, require.All
// controls whether an error is generated not all the ids produce a result.
func (tbl AssociationTable) GetAssociationsById(req require.Requirement, id ...int64) (list []*Association, err error) {
	if len(id) > 0 {
		if req == require.All {
			req = require.Exactly(len(id))
		}
		args := make([]interface{}, len(id))

		for i, v := range id {
			args[i] = v
		}

		list, err = tbl.getAssociations(req, tbl.pk, args...)
	}

	return list, err
}

// GetAssociationById gets the record with a given primary key value.
// If not found, *Association will be nil.
func (tbl AssociationTable) GetAssociationById(req require.Requirement, id int64) (*Association, error) {
	return tbl.getAssociation(req, tbl.pk, id)
}

func (tbl AssociationTable) getAssociation(req require.Requirement, column string, arg interface{}) (*Association, error) {
	d := tbl.Dialect()
	q := d.Quoter()
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s=?",
		allAssociationColumnNamesQuoted(q), tbl.quotedName(), q.Quote(column))
	v, err := tbl.doQueryAndScanOne(req, query, arg)
	return v, err
}

func (tbl AssociationTable) getAssociations(req require.Requirement, column string, args ...interface{}) (list []*Association, err error) {
	if len(args) > 0 {
		if req == require.All {
			req = require.Exactly(len(args))
		}
		d := tbl.Dialect()
		q := d.Quoter()
		pl := d.Placeholders(len(args))
		query := fmt.Sprintf("SELECT %s FROM %s WHERE %s IN (%s)",
			allAssociationColumnNamesQuoted(q), tbl.quotedName(), q.Quote(column), pl)
		list, err = tbl.doQueryAndScan(req, false, query, args...)
	}

	return list, err
}

func (tbl AssociationTable) doQueryAndScanOne(req require.Requirement, query string, args ...interface{}) (*Association, error) {
	list, err := tbl.doQueryAndScan(req, true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

// Fetch fetches a list of Association based on a supplied query. This is mostly used for join queries that map its
// result columns to the fields of Association. Other queries might be better handled by GetXxx or Select methods.
func (tbl AssociationTable) Fetch(req require.Requirement, query string, args ...interface{}) ([]*Association, error) {
	return tbl.doQueryAndScan(req, false, query, args...)
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
func (tbl AssociationTable) SelectOneWhere(req require.Requirement, where, orderBy string, args ...interface{}) (*Association, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT 1",
		allAssociationColumnNamesQuoted(tbl.Dialect().Quoter()), tbl.quotedName(), where, orderBy)
	v, err := tbl.doQueryAndScanOne(req, query, args...)
	return v, err
}

// SelectOne allows a single Association to be obtained from the database.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
// If not found, *Example will be nil.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.One
// controls whether an error is generated when no result is found.
func (tbl AssociationTable) SelectOne(req require.Requirement, wh where.Expression, qc where.QueryConstraint) (*Association, error) {
	q := tbl.Dialect().Quoter()
	whs, args := where.Where(wh, q)
	orderBy := where.Build(qc, q)
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
	query := fmt.Sprintf("SELECT %s FROM %s %s %s",
		allAssociationColumnNamesQuoted(tbl.Dialect().Quoter()), tbl.quotedName(), where, orderBy)
	vv, err := tbl.doQueryAndScan(req, false, query, args...)
	return vv, err
}

// Select allows Associations to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
func (tbl AssociationTable) Select(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]*Association, error) {
	q := tbl.Dialect().Quoter()
	whs, args := where.Where(wh, q)
	orderBy := where.Build(qc, q)
	return tbl.SelectWhere(req, whs, orderBy, args...)
}

//--------------------------------------------------------------------------------

// CountWhere counts Associations in the table that match a 'where' clause.
// Use a blank string for the 'where' argument if it is not needed.
//
// The args are for any placeholder parameters in the query.
func (tbl AssociationTable) CountWhere(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s %s", tbl.quotedName(), where)
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

// Count counts the Associations in the table that match a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed.
func (tbl AssociationTable) Count(wh where.Expression) (count int64, err error) {
	whs, args := where.Where(wh, tbl.Dialect().Quoter())
	return tbl.CountWhere(whs, args...)
}

//--------------------------------------------------------------------------------

// SliceId gets the id column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AssociationTable) SliceId(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
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
	return tbl.sliceQualNamePtrList(req, "quality", wh, qc)
}

// SliceCategory gets the category column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AssociationTable) SliceCategory(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]Category, error) {
	return tbl.sliceCategoryPtrList(req, "category", wh, qc)
}

func (tbl AssociationTable) sliceCategoryPtrList(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]Category, error) {
	q := tbl.Dialect().Quoter()
	whs, args := where.Where(wh, q)
	orderBy := where.Build(qc, q)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", q.Quote(sqlname), tbl.quotedName(), whs, orderBy)
	rows, err := support.Query(tbl, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var v Category
	list := make([]Category, 0, 10)

	for rows.Next() {
		err = rows.Scan(&v)
		if err == sql.ErrNoRows {
			return list, tbl.Logger().LogIfError(require.ErrorIfQueryNotSatisfiedBy(req, int64(len(list))))
		} else {
			list = append(list, v)
		}
	}
	return list, tbl.Logger().LogIfError(require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, int64(len(list))))
}

func (tbl AssociationTable) sliceQualNamePtrList(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]QualName, error) {
	q := tbl.Dialect().Quoter()
	whs, args := where.Where(wh, q)
	orderBy := where.Build(qc, q)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", q.Quote(sqlname), tbl.quotedName(), whs, orderBy)
	rows, err := support.Query(tbl, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var v QualName
	list := make([]QualName, 0, 10)

	for rows.Next() {
		err = rows.Scan(&v)
		if err == sql.ErrNoRows {
			return list, tbl.Logger().LogIfError(require.ErrorIfQueryNotSatisfiedBy(req, int64(len(list))))
		} else {
			list = append(list, v)
		}
	}
	return list, tbl.Logger().LogIfError(require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, int64(len(list))))
}

func (tbl AssociationTable) constructAssociationInsert(w dialect.StringWriter, v *Association, withPk bool) (s []interface{}, err error) {
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

func (tbl AssociationTable) constructAssociationUpdate(w dialect.StringWriter, v *Association) (s []interface{}, err error) {
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

// Insert adds new records for the Associations.
// The Associations have their primary key fields set to the new record identifiers.
// The Association.PreInsert() method will be called, if it exists.
func (tbl AssociationTable) Insert(req require.Requirement, vv ...*Association) error {
	if req == require.All {
		req = require.Exactly(len(vv))
	}

	var count int64
	insertHasReturningPhrase := tbl.Dialect().InsertHasReturningPhrase()
	returning := ""
	if tbl.Dialect().InsertHasReturningPhrase() {
		returning = fmt.Sprintf(" returning %q", tbl.pk)
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

		fields, err := tbl.constructAssociationInsert(b, v, false)
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
			err = row.Scan(&v.Id)

		} else {
			i64, e2 := tbl.db.InsertContext(tbl.ctx, query, fields...)
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

// UpdateFields updates one or more columns, given a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
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
				return count, tbl.Logger().LogError(err)
			}
		}

		b := dialect.Adapt(&bytes.Buffer{})
		b.WriteString("UPDATE ")
		tbl.quotedNameW(b)
		b.WriteString(" SET ")

		args, err := tbl.constructAssociationUpdate(b, v)
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
		return tbl.Logger().LogIfError(err)
	}

	if rows.Next() {
		return require.ErrWrongSize(2, "expected to find no more than 1 but got at least 2 using %q", wh)
	}

	v.Id = id
	_, err = tbl.Update(require.One, v)
	return err
}

//--------------------------------------------------------------------------------

// DeleteAssociations deletes rows from the table, given some primary keys.
// The list of ids can be arbitrarily long.
func (tbl AssociationTable) DeleteAssociations(req require.Requirement, id ...int64) (int64, error) {
	const batch = 1000 // limited by Oracle DB
	const qt = "DELETE FROM %s WHERE %s IN (%s)"
	qName := tbl.quotedName()

	if req == require.All {
		req = require.Exactly(len(id))
	}

	var count, n int64
	var err error
	var max = batch
	if len(id) < batch {
		max = len(id)
	}
	d := tbl.Dialect()
	col := d.Quoter().Quote(tbl.pk)
	args := make([]interface{}, max)

	if len(id) > batch {
		pl := d.Placeholders(batch)
		query := fmt.Sprintf(qt, qName, col, pl)

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
		pl := d.Placeholders(len(id))
		query := fmt.Sprintf(qt, qName, col, pl)

		for i := 0; i < len(id); i++ {
			args[i] = id[i]
		}

		n, err = tbl.Exec(nil, query, args...)
		count += n
	}

	return count, tbl.Logger().LogIfError(require.ChainErrorIfExecNotSatisfiedBy(err, req, n))
}

// Delete deletes one or more rows from the table, given a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl AssociationTable) Delete(req require.Requirement, wh where.Expression) (int64, error) {
	query, args := tbl.deleteRows(wh)
	return tbl.Exec(req, query, args...)
}

func (tbl AssociationTable) deleteRows(wh where.Expression) (string, []interface{}) {
	whs, args := where.Where(wh, tbl.Dialect().Quoter())
	query := fmt.Sprintf("DELETE FROM %s %s", tbl.quotedName(), whs)
	return query, args
}

//--------------------------------------------------------------------------------
