// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

package demo

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"github.com/rickb777/sqlgen2"
	"github.com/rickb777/sqlgen2/constraint"
	"github.com/rickb777/sqlgen2/require"
	"github.com/rickb777/sqlgen2/schema"
	"github.com/rickb777/sqlgen2/support"
	"github.com/rickb777/sqlgen2/where"
	"io"
	"log"
)

// AssociationTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type AssociationTable struct {
	name        sqlgen2.TableName
	database    *sqlgen2.Database
	db          sqlgen2.Execer
	constraints constraint.Constraints
	ctx			context.Context
	pk          string
}

// Type conformance checks
var _ sqlgen2.TableCreator = &AssociationTable{}
var _ sqlgen2.TableWithCrud = &AssociationTable{}

// NewAssociationTable returns a new table instance.
// If a blank table name is supplied, the default name "associations" will be used instead.
// The request context is initialised with the background.
func NewAssociationTable(name string, d *sqlgen2.Database) AssociationTable {
	if name == "" {
		name = "associations"
	}
	var constraints constraint.Constraints
	return AssociationTable{
		name:        sqlgen2.TableName{"", name},
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
func CopyTableAsAssociationTable(origin sqlgen2.Table) AssociationTable {
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
func (tbl AssociationTable) Database() *sqlgen2.Database {
	return tbl.database
}

// Logger gets the trace logger.
func (tbl AssociationTable) Logger() *log.Logger {
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
func (tbl AssociationTable) Dialect() schema.Dialect {
	return tbl.database.Dialect()
}

// Name gets the table name.
func (tbl AssociationTable) Name() sqlgen2.TableName {
	return tbl.name
}


// PkColumn gets the column name used as a primary key.
func (tbl AssociationTable) PkColumn() string {
	return tbl.pk
}


// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl AssociationTable) DB() *sql.DB {
	return tbl.db.(*sql.DB)
}

// Execer gets the wrapped database or transaction handle.
func (tbl AssociationTable) Execer() sqlgen2.Execer {
	return tbl.db
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl AssociationTable) Tx() *sql.Tx {
	return tbl.db.(*sql.Tx)
}

// IsTx tests whether this is within a transaction.
func (tbl AssociationTable) IsTx() bool {
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
func (tbl AssociationTable) BeginTx(opts *sql.TxOptions) (AssociationTable, error) {
	var err error
	tbl.db, err = tbl.db.(sqlgen2.TxStarter).BeginTx(tbl.ctx, opts)
	return tbl, tbl.logIfError(err)
}

// Using returns a modified Table using the transaction supplied. This is needed
// when making multiple queries across several tables within a single transaction.
// The result is a modified copy of the table; the original is unchanged.
func (tbl AssociationTable) Using(tx *sql.Tx) AssociationTable {
	tbl.db = tx
	return tbl
}

func (tbl AssociationTable) logQuery(query string, args ...interface{}) {
	tbl.database.LogQuery(query, args...)
}

func (tbl AssociationTable) logError(err error) error {
	return tbl.database.LogError(err)
}

func (tbl AssociationTable) logIfError(err error) error {
	return tbl.database.LogIfError(err)
}


//--------------------------------------------------------------------------------

const NumAssociationColumns = 6

const NumAssociationDataColumns = 5

const AssociationColumnNames = "id,name,quality,ref1,ref2,category"

const AssociationDataColumnNames = "name,quality,ref1,ref2,category"

//--------------------------------------------------------------------------------

const sqlAssociationTableCreateColumnsSqlite = "\n"+
" `id`       integer not null primary key autoincrement,\n"+
" `name`     text default null,\n"+
" `quality`  text default null,\n"+
" `ref1`     bigint default null,\n"+
" `ref2`     bigint default null,\n"+
" `category` tinyint unsigned default null"

const sqlAssociationTableCreateColumnsMysql = "\n"+
" `id`       bigint not null primary key auto_increment,\n"+
" `name`     varchar(255) default null,\n"+
" `quality`  varchar(255) default null,\n"+
" `ref1`     bigint default null,\n"+
" `ref2`     bigint default null,\n"+
" `category` tinyint unsigned default null"

const sqlAssociationTableCreateColumnsPostgres = `
 "id"       bigserial not null primary key,
 "name"     varchar(255) default null,
 "quality"  varchar(255) default null,
 "ref1"     bigint default null,
 "ref2"     bigint default null,
 "category" smallint default null`

const sqlConstrainAssociationTable = `
`

//--------------------------------------------------------------------------------

// CreateTable creates the table.
func (tbl AssociationTable) CreateTable(ifNotExists bool) (int64, error) {
	return support.Exec(tbl, nil, tbl.createTableSql(ifNotExists))
}

func (tbl AssociationTable) createTableSql(ifNotExists bool) string {
	var columns string
	var settings string
	switch tbl.Dialect() {
	case schema.Sqlite:
		columns = sqlAssociationTableCreateColumnsSqlite
		settings = ""
    case schema.Mysql:
		columns = sqlAssociationTableCreateColumnsMysql
		settings = " ENGINE=InnoDB DEFAULT CHARSET=utf8"
    case schema.Postgres:
		columns = sqlAssociationTableCreateColumnsPostgres
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
	query := fmt.Sprintf("DROP TABLE %s%s", ie, tbl.name)
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

// Query is the low-level request method for this table. The query is logged using whatever logger is
// configured. If an error arises, this too is logged.
//
// If you need a context other than the background, use WithContext before calling Query.
//
// The args are for any placeholder parameters in the query.
//
// The caller must call rows.Close() on the result.
//
// Wrap the result in *sqlgen2.Rows if you need to access its data as a map.
func (tbl AssociationTable) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return support.Query(tbl, query, args...)
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

func scanAssociations(rows *sql.Rows, firstOnly bool) (vv []*Association, n int64, err error) {
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
			return vv, n, err
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

var allAssociationQuotedColumnNames = []string{
	schema.Sqlite.SplitAndQuote(AssociationColumnNames),
	schema.Mysql.SplitAndQuote(AssociationColumnNames),
	schema.Postgres.SplitAndQuote(AssociationColumnNames),
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
	dialect := tbl.Dialect()
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s=%s",
		allAssociationQuotedColumnNames[dialect.Index()], tbl.name, dialect.Quote(column), dialect.Placeholder(column, 1))
	v, err := tbl.doQueryOne(req, query, arg)
	return v, err
}

func (tbl AssociationTable) getAssociations(req require.Requirement, column string, args ...interface{}) (list []*Association, err error) {
	if len(args) > 0 {
		if req == require.All {
			req = require.Exactly(len(args))
		}
		dialect := tbl.Dialect()
		pl := dialect.Placeholders(len(args))
		query := fmt.Sprintf("SELECT %s FROM %s WHERE %s IN (%s)",
			allAssociationQuotedColumnNames[dialect.Index()], tbl.name, dialect.Quote(column), pl)
		list, err = tbl.doQuery(req, false, query, args...)
	}

	return list, err
}

func (tbl AssociationTable) doQueryOne(req require.Requirement, query string, args ...interface{}) (*Association, error) {
	list, err := tbl.doQuery(req, true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

func (tbl AssociationTable) doQuery(req require.Requirement, firstOnly bool, query string, args ...interface{}) ([]*Association, error) {
	rows, err := tbl.Query(query, args...)
	if err != nil {
		return nil, err
	}

	vv, n, err := scanAssociations(rows, firstOnly)
	return vv, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(err, req, n))
}

// Fetch fetches a list of Association based on a supplied query. This is mostly used for join queries that map its
// result columns to the fields of Association. Other queries might be better handled by GetXxx or Select methods.
func (tbl AssociationTable) Fetch(req require.Requirement, query string, args ...interface{}) ([]*Association, error) {
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
func (tbl AssociationTable) SelectOneWhere(req require.Requirement, where, orderBy string, args ...interface{}) (*Association, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT 1",
		allAssociationQuotedColumnNames[tbl.Dialect().Index()], tbl.name, where, orderBy)
	v, err := tbl.doQueryOne(req, query, args...)
	return v, err
}

// SelectOne allows a single Association to be obtained from the sqlgen2.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
// If not found, *Example will be nil.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.One
// controls whether an error is generated when no result is found.
func (tbl AssociationTable) SelectOne(req require.Requirement, wh where.Expression, qc where.QueryConstraint) (*Association, error) {
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
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
		allAssociationQuotedColumnNames[tbl.Dialect().Index()], tbl.name, where, orderBy)
	vv, err := tbl.doQuery(req, false, query, args...)
	return vv, err
}

// Select allows Associations to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
func (tbl AssociationTable) Select(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]*Association, error) {
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
	return tbl.SelectWhere(req, whs, orderBy, args...)
}

// CountWhere counts Associations in the table that match a 'where' clause.
// Use a blank string for the 'where' argument if it is not needed.
//
// The args are for any placeholder parameters in the query.
func (tbl AssociationTable) CountWhere(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s %s", tbl.name, where)
	tbl.logQuery(query, args...)
	row := tbl.db.QueryRowContext(tbl.ctx, query, args...)
	err = row.Scan(&count)
	return count, tbl.logIfError(err)
}

// Count counts the Associations in the table that match a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed.
func (tbl AssociationTable) Count(wh where.Expression) (count int64, err error) {
	whs, args := where.BuildExpression(wh, tbl.Dialect())
	return tbl.CountWhere(whs, args...)
}

//--------------------------------------------------------------------------------

// SliceId gets the id column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AssociationTable) SliceId(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return tbl.sliceInt64List(req, tbl.pk, wh, qc)
}

// SliceName gets the name column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AssociationTable) SliceName(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.sliceStringPtrList(req, "name", wh, qc)
}

// SliceQuality gets the quality column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AssociationTable) SliceQuality(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]QualName, error) {
	return tbl.sliceQualNamePtrList(req, "quality", wh, qc)
}

// SliceRef1 gets the ref1 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AssociationTable) SliceRef1(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return tbl.sliceInt64PtrList(req, "ref1", wh, qc)
}

// SliceRef2 gets the ref2 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AssociationTable) SliceRef2(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return tbl.sliceInt64PtrList(req, "ref2", wh, qc)
}

// SliceCategory gets the category column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AssociationTable) SliceCategory(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]Category, error) {
	return tbl.sliceCategoryPtrList(req, "category", wh, qc)
}

func (tbl AssociationTable) sliceCategoryPtrList(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]Category, error) {
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

	var v Category
	list := make([]Category, 0, 10)

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

func (tbl AssociationTable) sliceQualNamePtrList(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]QualName, error) {
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

	var v QualName
	list := make([]QualName, 0, 10)

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

func (tbl AssociationTable) sliceInt64List(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
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

func (tbl AssociationTable) sliceInt64PtrList(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
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

func (tbl AssociationTable) sliceStringPtrList(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
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


func constructAssociationInsert(w io.Writer, v *Association, dialect schema.Dialect, withPk bool) (s []interface{}, err error) {
	s = make([]interface{}, 0, 6)

	comma := ""
	io.WriteString(w, " (")

	if withPk {
		dialect.QuoteW(w, "id")
		comma = ","
		s = append(s, v.Id)
	}

	if v.Name != nil {
		io.WriteString(w, comma)

		dialect.QuoteW(w, "name")
		s = append(s, v.Name)
		comma = ","
	}
	if v.Quality != nil {
		io.WriteString(w, comma)

		dialect.QuoteW(w, "quality")
		s = append(s, v.Quality)
		comma = ","
	}
	if v.Ref1 != nil {
		io.WriteString(w, comma)

		dialect.QuoteW(w, "ref1")
		s = append(s, v.Ref1)
		comma = ","
	}
	if v.Ref2 != nil {
		io.WriteString(w, comma)

		dialect.QuoteW(w, "ref2")
		s = append(s, v.Ref2)
		comma = ","
	}
	if v.Category != nil {
		io.WriteString(w, comma)

		dialect.QuoteW(w, "category")
		s = append(s, v.Category)
		comma = ","
	}
	io.WriteString(w, ")")
	return s, nil
}

func constructAssociationUpdate(w io.Writer, v *Association, dialect schema.Dialect) (s []interface{}, err error) {
	j := 1
	s = make([]interface{}, 0, 5)

	comma := ""

	io.WriteString(w, comma)
	if v.Name != nil {
		dialect.QuoteWithPlaceholder(w, "name", j)
		s = append(s, v.Name)
		comma = ", "
		j++
	} else {
		dialect.QuoteW(w, "name")
		io.WriteString(w, "=NULL")
	}

	io.WriteString(w, comma)
	if v.Quality != nil {
		dialect.QuoteWithPlaceholder(w, "quality", j)
		s = append(s, v.Quality)
		comma = ", "
		j++
	} else {
		dialect.QuoteW(w, "quality")
		io.WriteString(w, "=NULL")
	}

	io.WriteString(w, comma)
	if v.Ref1 != nil {
		dialect.QuoteWithPlaceholder(w, "ref1", j)
		s = append(s, v.Ref1)
		comma = ", "
		j++
	} else {
		dialect.QuoteW(w, "ref1")
		io.WriteString(w, "=NULL")
	}

	io.WriteString(w, comma)
	if v.Ref2 != nil {
		dialect.QuoteWithPlaceholder(w, "ref2", j)
		s = append(s, v.Ref2)
		comma = ", "
		j++
	} else {
		dialect.QuoteW(w, "ref2")
		io.WriteString(w, "=NULL")
	}

	io.WriteString(w, comma)
	if v.Category != nil {
		dialect.QuoteWithPlaceholder(w, "category", j)
		s = append(s, v.Category)
		comma = ", "
		j++
	} else {
		dialect.QuoteW(w, "category")
		io.WriteString(w, "=NULL")
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
		returning = fmt.Sprintf(" returning %q", tbl.pk)
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

		fields, err := constructAssociationInsert(b, v, tbl.Dialect(), false)
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
			err = row.Scan(&v.Id)

		} else {
			res, e2 := tbl.db.ExecContext(tbl.ctx, query, fields...)
			if e2 != nil {
				return tbl.logError(e2)
			}

			v.Id, err = res.LastInsertId()
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
func (tbl AssociationTable) UpdateFields(req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error) {
	return support.UpdateFields(tbl, req, wh, fields...)
}

//--------------------------------------------------------------------------------

var allAssociationQuotedUpdates = []string{
	// Sqlite
	"`name`=?,`quality`=?,`ref1`=?,`ref2`=?,`category`=? WHERE `id`=?",
	// Mysql
	"`name`=?,`quality`=?,`ref1`=?,`ref2`=?,`category`=? WHERE `id`=?",
	// Postgres
	`"name"=$2,"quality"=$3,"ref1"=$4,"ref2"=$5,"category"=$6 WHERE "id"=$1`,
}

//--------------------------------------------------------------------------------

// Update updates records, matching them by primary key. It returns the number of rows affected.
// The Association.PreUpdate(Execer) method will be called, if it exists.
func (tbl AssociationTable) Update(req require.Requirement, vv ...*Association) (int64, error) {
	if req == require.All {
		req = require.Exactly(len(vv))
	}

	var count int64
	dialect := tbl.Dialect()
	//columns := allAssociationQuotedUpdates[dialect.Index()]
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

		args, err := constructAssociationUpdate(b, v, dialect)
		k := len(args) + 1
		args = append(args, v.Id)
		if err != nil {
			return count, tbl.logError(err)
		}

		io.WriteString(b, " WHERE ")
		dialect.QuoteWithPlaceholder(b, tbl.pk, k)

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

// DeleteAssociations deletes rows from the table, given some primary keys.
// The list of ids can be arbitrarily long.
func (tbl AssociationTable) DeleteAssociations(req require.Requirement, id ...int64) (int64, error) {
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
	col := dialect.Quote(tbl.pk)
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
func (tbl AssociationTable) Delete(req require.Requirement, wh where.Expression) (int64, error) {
	query, args := tbl.deleteRows(wh)
	return tbl.Exec(req, query, args...)
}

func (tbl AssociationTable) deleteRows(wh where.Expression) (string, []interface{}) {
	whs, args := where.BuildExpression(wh, tbl.Dialect())
	query := fmt.Sprintf("DELETE FROM %s %s", tbl.name, whs)
	return query, args
}

//--------------------------------------------------------------------------------
