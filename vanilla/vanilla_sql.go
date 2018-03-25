// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

package vanilla

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/rickb777/sqlgen2"
	"github.com/rickb777/sqlgen2/constraint"
	"github.com/rickb777/sqlgen2/require"
	"github.com/rickb777/sqlgen2/schema"
	"github.com/rickb777/sqlgen2/support"
	"github.com/rickb777/sqlgen2/where"
	"log"
)

// PrimaryKeyTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type PrimaryKeyTable struct {
	name        sqlgen2.TableName
	database    *sqlgen2.Database
	db          sqlgen2.Execer
	constraints constraint.Constraints
	ctx			context.Context
	pk          string
}

// Type conformance checks
var _ sqlgen2.Table = &PrimaryKeyTable{}
var _ sqlgen2.Table = &PrimaryKeyTable{}

// NewPrimaryKeyTable returns a new table instance.
// If a blank table name is supplied, the default name "primarykeys" will be used instead.
// The request context is initialised with the background.
func NewPrimaryKeyTable(name string, d *sqlgen2.Database) PrimaryKeyTable {
	if name == "" {
		name = "primarykeys"
	}
	var constraints constraint.Constraints
	return PrimaryKeyTable{
		name:        sqlgen2.TableName{"", name},
		database:    d,
		db:          d.DB(),
		constraints: constraints,
		ctx:         context.Background(),
		pk:          "id",
	}
}

// CopyTableAsPrimaryKeyTable copies a table instance, retaining the name etc but
// providing methods appropriate for 'PrimaryKey'. It doesn't copy the constraints of the original table.
//
// It serves to provide methods appropriate for 'PrimaryKey'. This is most useful when this is used to represent a
// join result. In such cases, there won't be any need for DDL methods, nor Exec, Insert, Update or Delete.
func CopyTableAsPrimaryKeyTable(origin sqlgen2.Table) PrimaryKeyTable {
	return PrimaryKeyTable{
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
func (tbl PrimaryKeyTable) SetPkColumn(pk string) PrimaryKeyTable {
	tbl.pk = pk
	return tbl
}


// WithPrefix sets the table name prefix for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl PrimaryKeyTable) WithPrefix(pfx string) PrimaryKeyTable {
	tbl.name.Prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries via this table.
// The result is a modified copy of the table; the original is unchanged.
//
// The shared context in the *Database is not altered by this method. So it
// is possible to use different contexts for different (groups of) queries.
func (tbl PrimaryKeyTable) WithContext(ctx context.Context) PrimaryKeyTable {
	tbl.ctx = ctx
	return tbl
}

// Database gets the shared database information.
func (tbl PrimaryKeyTable) Database() *sqlgen2.Database {
	return tbl.database
}

// Logger gets the trace logger.
func (tbl PrimaryKeyTable) Logger() *log.Logger {
	return tbl.database.Logger()
}

// WithConstraint returns a modified Table with added data consistency constraints.
func (tbl PrimaryKeyTable) WithConstraint(cc ...constraint.Constraint) PrimaryKeyTable {
	tbl.constraints = append(tbl.constraints, cc...)
	return tbl
}

// Constraints returns the table's constraints.
func (tbl PrimaryKeyTable) Constraints() constraint.Constraints {
	return tbl.constraints
}

// Ctx gets the current request context.
func (tbl PrimaryKeyTable) Ctx() context.Context {
	return tbl.ctx
}

// Dialect gets the database dialect.
func (tbl PrimaryKeyTable) Dialect() schema.Dialect {
	return tbl.database.Dialect()
}

// Name gets the table name.
func (tbl PrimaryKeyTable) Name() sqlgen2.TableName {
	return tbl.name
}


// PkColumn gets the column name used as a primary key.
func (tbl PrimaryKeyTable) PkColumn() string {
	return tbl.pk
}


// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl PrimaryKeyTable) DB() *sql.DB {
	return tbl.db.(*sql.DB)
}

// Execer gets the wrapped database or transaction handle.
func (tbl PrimaryKeyTable) Execer() sqlgen2.Execer {
	return tbl.db
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl PrimaryKeyTable) Tx() *sql.Tx {
	return tbl.db.(*sql.Tx)
}

// IsTx tests whether this is within a transaction.
func (tbl PrimaryKeyTable) IsTx() bool {
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
func (tbl PrimaryKeyTable) BeginTx(opts *sql.TxOptions) (PrimaryKeyTable, error) {
	var err error
	tbl.db, err = tbl.db.(sqlgen2.TxStarter).BeginTx(tbl.ctx, opts)
	return tbl, tbl.logIfError(err)
}

// Using returns a modified Table using the transaction supplied. This is needed
// when making multiple queries across several tables within a single transaction.
// The result is a modified copy of the table; the original is unchanged.
func (tbl PrimaryKeyTable) Using(tx *sql.Tx) PrimaryKeyTable {
	tbl.db = tx
	return tbl
}

func (tbl PrimaryKeyTable) logQuery(query string, args ...interface{}) {
	tbl.database.LogQuery(query, args...)
}

func (tbl PrimaryKeyTable) logError(err error) error {
	return tbl.database.LogError(err)
}

func (tbl PrimaryKeyTable) logIfError(err error) error {
	return tbl.database.LogIfError(err)
}


//--------------------------------------------------------------------------------

const NumPrimaryKeyColumns = 1

const NumPrimaryKeyDataColumns = 0

const PrimaryKeyColumnNames = "id"

const PrimaryKeyDataColumnNames = ""

//--------------------------------------------------------------------------------

// Exec executes a query without returning any rows.
// It returns the number of rows affected (if the database driver supports this).
//
// The args are for any placeholder parameters in the query.
func (tbl PrimaryKeyTable) Exec(req require.Requirement, query string, args ...interface{}) (int64, error) {
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
func (tbl PrimaryKeyTable) Query(query string, args ...interface{}) (*sql.Rows, error) {
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
func (tbl PrimaryKeyTable) QueryOneNullString(req require.Requirement, query string, args ...interface{}) (result sql.NullString, err error) {
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
func (tbl PrimaryKeyTable) QueryOneNullInt64(req require.Requirement, query string, args ...interface{}) (result sql.NullInt64, err error) {
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
func (tbl PrimaryKeyTable) QueryOneNullFloat64(req require.Requirement, query string, args ...interface{}) (result sql.NullFloat64, err error) {
	err = support.QueryOneNullThing(tbl, req, &result, query, args...)
	return result, err
}

func scanPrimaryKeys(rows *sql.Rows, firstOnly bool) (vv []*PrimaryKey, n int64, err error) {
	for rows.Next() {
		n++

		var v0 int64

		err = rows.Scan(
			&v0,
		)
		if err != nil {
			return vv, n, err
		}

		v := &PrimaryKey{}
		v.Id = v0

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

var allPrimaryKeyQuotedColumnNames = []string{
	schema.Sqlite.SplitAndQuote(PrimaryKeyColumnNames),
	schema.Mysql.SplitAndQuote(PrimaryKeyColumnNames),
	schema.Postgres.SplitAndQuote(PrimaryKeyColumnNames),
}

//--------------------------------------------------------------------------------

// GetPrimaryKeysById gets records from the table according to a list of primary keys.
// Although the list of ids can be arbitrarily long, there are practical limits;
// note that Oracle DB has a limit of 1000.
//
// It places a requirement, which may be nil, on the size of the expected results: in particular, require.All
// controls whether an error is generated not all the ids produce a result.
func (tbl PrimaryKeyTable) GetPrimaryKeysById(req require.Requirement, id ...int64) (list []*PrimaryKey, err error) {
	if len(id) > 0 {
		if req == require.All {
			req = require.Exactly(len(id))
		}
		args := make([]interface{}, len(id))

		for i, v := range id {
			args[i] = v
		}

		list, err = tbl.getPrimaryKeys(req, tbl.pk, args...)
	}

	return list, err
}

// GetPrimaryKeyById gets the record with a given primary key value.
// If not found, *PrimaryKey will be nil.
func (tbl PrimaryKeyTable) GetPrimaryKeyById(req require.Requirement, id int64) (*PrimaryKey, error) {
	return tbl.getPrimaryKey(req, tbl.pk, id)
}

func (tbl PrimaryKeyTable) getPrimaryKey(req require.Requirement, column string, arg interface{}) (*PrimaryKey, error) {
	dialect := tbl.Dialect()
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s=%s",
		allPrimaryKeyQuotedColumnNames[dialect.Index()], tbl.name, dialect.Quote(column), dialect.Placeholder(column, 1))
	v, err := tbl.doQueryOne(req, query, arg)
	return v, err
}

func (tbl PrimaryKeyTable) getPrimaryKeys(req require.Requirement, column string, args ...interface{}) (list []*PrimaryKey, err error) {
	if len(args) > 0 {
		if req == require.All {
			req = require.Exactly(len(args))
		}
		dialect := tbl.Dialect()
		pl := dialect.Placeholders(len(args))
		query := fmt.Sprintf("SELECT %s FROM %s WHERE %s IN (%s)",
			allPrimaryKeyQuotedColumnNames[dialect.Index()], tbl.name, dialect.Quote(column), pl)
		list, err = tbl.doQuery(req, false, query, args...)
	}

	return list, err
}

func (tbl PrimaryKeyTable) doQueryOne(req require.Requirement, query string, args ...interface{}) (*PrimaryKey, error) {
	list, err := tbl.doQuery(req, true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

func (tbl PrimaryKeyTable) doQuery(req require.Requirement, firstOnly bool, query string, args ...interface{}) ([]*PrimaryKey, error) {
	rows, err := tbl.Query(query, args...)
	if err != nil {
		return nil, err
	}

	vv, n, err := scanPrimaryKeys(rows, firstOnly)
	return vv, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(err, req, n))
}

// Fetch fetches a list of PrimaryKey based on a supplied query. This is mostly used for join queries that map its
// result columns to the fields of PrimaryKey. Other queries might be better handled by GetXxx or Select methods.
func (tbl PrimaryKeyTable) Fetch(req require.Requirement, query string, args ...interface{}) ([]*PrimaryKey, error) {
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
func (tbl PrimaryKeyTable) SelectOneWhere(req require.Requirement, where, orderBy string, args ...interface{}) (*PrimaryKey, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT 1",
		allPrimaryKeyQuotedColumnNames[tbl.Dialect().Index()], tbl.name, where, orderBy)
	v, err := tbl.doQueryOne(req, query, args...)
	return v, err
}

// SelectOne allows a single PrimaryKey to be obtained from the sqlgen2.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
// If not found, *Example will be nil.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.One
// controls whether an error is generated when no result is found.
func (tbl PrimaryKeyTable) SelectOne(req require.Requirement, wh where.Expression, qc where.QueryConstraint) (*PrimaryKey, error) {
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
	return tbl.SelectOneWhere(req, whs, orderBy, args...)
}

// SelectWhere allows PrimaryKeys to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
//
// The args are for any placeholder parameters in the query.
func (tbl PrimaryKeyTable) SelectWhere(req require.Requirement, where, orderBy string, args ...interface{}) ([]*PrimaryKey, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s %s",
		allPrimaryKeyQuotedColumnNames[tbl.Dialect().Index()], tbl.name, where, orderBy)
	vv, err := tbl.doQuery(req, false, query, args...)
	return vv, err
}

// Select allows PrimaryKeys to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
func (tbl PrimaryKeyTable) Select(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]*PrimaryKey, error) {
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
	return tbl.SelectWhere(req, whs, orderBy, args...)
}

// CountWhere counts PrimaryKeys in the table that match a 'where' clause.
// Use a blank string for the 'where' argument if it is not needed.
//
// The args are for any placeholder parameters in the query.
func (tbl PrimaryKeyTable) CountWhere(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s %s", tbl.name, where)
	tbl.logQuery(query, args...)
	row := tbl.db.QueryRowContext(tbl.ctx, query, args...)
	err = row.Scan(&count)
	return count, tbl.logIfError(err)
}

// Count counts the PrimaryKeys in the table that match a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed.
func (tbl PrimaryKeyTable) Count(wh where.Expression) (count int64, err error) {
	whs, args := where.BuildExpression(wh, tbl.Dialect())
	return tbl.CountWhere(whs, args...)
}

//--------------------------------------------------------------------------------

// SliceId gets the id column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl PrimaryKeyTable) SliceId(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return tbl.getint64list(req, tbl.pk, wh, qc)
}

func (tbl PrimaryKeyTable) getint64list(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
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


//--------------------------------------------------------------------------------

// DeletePrimaryKeys deletes rows from the table, given some primary keys.
// The list of ids can be arbitrarily long.
func (tbl PrimaryKeyTable) DeletePrimaryKeys(req require.Requirement, id ...int64) (int64, error) {
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
func (tbl PrimaryKeyTable) Delete(req require.Requirement, wh where.Expression) (int64, error) {
	query, args := tbl.deleteRows(wh)
	return tbl.Exec(req, query, args...)
}

func (tbl PrimaryKeyTable) deleteRows(wh where.Expression) (string, []interface{}) {
	whs, args := where.BuildExpression(wh, tbl.Dialect())
	query := fmt.Sprintf("DELETE FROM %s %s", tbl.name, whs)
	return query, args
}

//--------------------------------------------------------------------------------
