// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.
// sqlapi v0.29.0; sqlgen v0.49.0

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

// HookTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type HookTable struct {
	name        sqlapi.TableName
	database    sqlapi.Database
	db          sqlapi.Execer
	constraints constraint.Constraints
	ctx         context.Context
	pk          string
}

// Type conformance checks
var _ sqlapi.TableCreator = &HookTable{}
var _ sqlapi.TableWithCrud = &HookTable{}

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
		ctx:         context.Background(),
		pk:          "id",
	}
}

// CopyTableAsHookTable copies a table instance, retaining the name etc but
// providing methods appropriate for 'Hook'. It doesn't copy the constraints of the original table.
//
// It serves to provide methods appropriate for 'Hook'. This is most useful when this is used to represent a
// join result. In such cases, there won't be any need for DDL methods, nor Exec, Insert, Update or Delete.
func CopyTableAsHookTable(origin sqlapi.Table) HookTable {
	return HookTable{
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
func (tbl HookTable) SetPkColumn(pk string) HookTable {
	tbl.pk = pk
	return tbl
}

// WithPrefix sets the table name prefix for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl HookTable) WithPrefix(pfx string) HookTable {
	tbl.name.Prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries via this table.
// The result is a modified copy of the table; the original is unchanged.
//
// The shared context in the *Database is not altered by this method. So it
// is possible to use different contexts for different (groups of) queries.
func (tbl HookTable) WithContext(ctx context.Context) HookTable {
	tbl.ctx = ctx
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

// WithConstraint returns a modified Table with added data consistency constraints.
func (tbl HookTable) WithConstraint(cc ...constraint.Constraint) HookTable {
	tbl.constraints = append(tbl.constraints, cc...)
	return tbl
}

// Constraints returns the table's constraints.
func (tbl HookTable) Constraints() constraint.Constraints {
	return tbl.constraints
}

// Ctx gets the current request context.
func (tbl HookTable) Ctx() context.Context {
	return tbl.ctx
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
func (tbl HookTable) BeginTx(opts *sql.TxOptions) (HookTable, error) {
	var err error
	tbl.db, err = tbl.db.(sqlapi.SqlDB).BeginTx(tbl.ctx, opts)
	return tbl, tbl.Logger().LogIfError(err)
}

// Using returns a modified Table using the transaction supplied. This is needed
// when making multiple queries across several tables within a single transaction.
// The result is a modified copy of the table; the original is unchanged.
func (tbl HookTable) Using(tx sqlapi.SqlTx) HookTable {
	tbl.db = tx
	return tbl
}

func (tbl HookTable) quotedName() string {
	return tbl.Dialect().Quoter().Quote(tbl.name.String())
}

func (tbl HookTable) quotedNameW(w dialect.StringWriter) {
	tbl.Dialect().Quoter().QuoteW(w, tbl.name.String())
}

//--------------------------------------------------------------------------------

// NumHookTableColumns is the total number of columns in HookTable.
const NumHookTableColumns = 17

// NumHookTableDataColumns is the number of columns in HookTable not including the auto-increment key.
const NumHookTableDataColumns = 16

// HookTableColumnNames is the list of columns in HookTable.
const HookTableColumnNames = "id,sha,after,before,category,created,deleted,forced,commit_id,message,timestamp,head_commit_author_name,head_commit_author_email,head_commit_author_username,head_commit_committer_name,head_commit_committer_email,head_commit_committer_username"

// HookTableDataColumnNames is the list of data columns in HookTable.
const HookTableDataColumnNames = "sha,after,before,category,created,deleted,forced,commit_id,message,timestamp,head_commit_author_name,head_commit_author_email,head_commit_author_username,head_commit_committer_name,head_commit_committer_email,head_commit_committer_username"

var listOfHookTableColumnNames = strings.Split(HookTableColumnNames, ",")

//--------------------------------------------------------------------------------

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

//--------------------------------------------------------------------------------

// CreateTable creates the table.
func (tbl HookTable) CreateTable(ifNotExists bool) (int64, error) {
	return support.Exec(tbl, nil, tbl.createTableSql(ifNotExists))
}

func (tbl HookTable) createTableSql(ifNotExists bool) string {
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

	for i, c := range tbl.constraints {
		buf.WriteString(",\n ")
		buf.WriteString(c.ConstraintSql(tbl.Dialect().Quoter(), tbl.name, i+1))
	}

	buf.WriteString("\n)")
	buf.WriteString(tbl.Dialect().CreateTableSettings())
	return buf.String()
}

func (tbl HookTable) ternary(flag bool, a, b string) string {
	if flag {
		return a
	}
	return b
}

// DropTable drops the table, destroying all its data.
func (tbl HookTable) DropTable(ifExists bool) (int64, error) {
	return support.Exec(tbl, nil, tbl.dropTableSql(ifExists))
}

func (tbl HookTable) dropTableSql(ifExists bool) string {
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
func (tbl HookTable) Truncate(force bool) (err error) {
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
func (tbl HookTable) Exec(req require.Requirement, query string, args ...interface{}) (int64, error) {
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
// Wrap the result in *sqlapi.Rows if you need to access its data as a map.
func (tbl HookTable) Query(query string, args ...interface{}) (sqlapi.SqlRows, error) {
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
func (tbl HookTable) QueryOneNullString(req require.Requirement, query string, args ...interface{}) (result sql.NullString, err error) {
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
func (tbl HookTable) QueryOneNullInt64(req require.Requirement, query string, args ...interface{}) (result sql.NullInt64, err error) {
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
func (tbl HookTable) QueryOneNullFloat64(req require.Requirement, query string, args ...interface{}) (result sql.NullFloat64, err error) {
	err = support.QueryOneNullThing(tbl, req, &result, query, args...)
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

// GetHooksById gets records from the table according to a list of primary keys.
// Although the list of ids can be arbitrarily long, there are practical limits;
// note that Oracle DB has a limit of 1000.
//
// It places a requirement, which may be nil, on the size of the expected results: in particular, require.All
// controls whether an error is generated not all the ids produce a result.
func (tbl HookTable) GetHooksById(req require.Requirement, id ...uint64) (list HookList, err error) {
	if len(id) > 0 {
		if req == require.All {
			req = require.Exactly(len(id))
		}
		args := make([]interface{}, len(id))

		for i, v := range id {
			args[i] = v
		}

		list, err = tbl.getHooks(req, tbl.pk, args...)
	}

	return list, err
}

// GetHookById gets the record with a given primary key value.
// If not found, *Hook will be nil.
func (tbl HookTable) GetHookById(req require.Requirement, id uint64) (*Hook, error) {
	return tbl.getHook(req, tbl.pk, id)
}

func (tbl HookTable) getHook(req require.Requirement, column string, arg interface{}) (*Hook, error) {
	d := tbl.Dialect()
	q := d.Quoter()
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s=?",
		allHookColumnNamesQuoted(q), tbl.quotedName(), q.Quote(column))
	v, err := tbl.doQueryAndScanOne(req, query, arg)
	return v, err
}

func (tbl HookTable) getHooks(req require.Requirement, column string, args ...interface{}) (list HookList, err error) {
	if len(args) > 0 {
		if req == require.All {
			req = require.Exactly(len(args))
		}
		d := tbl.Dialect()
		q := d.Quoter()
		pl := d.Placeholders(len(args))
		query := fmt.Sprintf("SELECT %s FROM %s WHERE %s IN (%s)",
			allHookColumnNamesQuoted(q), tbl.quotedName(), q.Quote(column), pl)
		list, err = tbl.doQueryAndScan(req, false, query, args...)
	}

	return list, err
}

func (tbl HookTable) doQueryAndScanOne(req require.Requirement, query string, args ...interface{}) (*Hook, error) {
	list, err := tbl.doQueryAndScan(req, true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

func (tbl HookTable) doQueryAndScan(req require.Requirement, firstOnly bool, query string, args ...interface{}) (HookList, error) {
	rows, err := support.Query(tbl, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vv, n, err := ScanHooks(query, rows, firstOnly)
	return vv, tbl.Logger().LogIfError(require.ChainErrorIfQueryNotSatisfiedBy(err, req, n))
}

// Fetch fetches a list of Hook based on a supplied query. This is mostly used for join queries that map its
// result columns to the fields of Hook. Other queries might be better handled by GetXxx or Select methods.
func (tbl HookTable) Fetch(req require.Requirement, query string, args ...interface{}) (HookList, error) {
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
func (tbl HookTable) SelectOneWhere(req require.Requirement, where, orderBy string, args ...interface{}) (*Hook, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT 1",
		allHookColumnNamesQuoted(tbl.Dialect().Quoter()), tbl.quotedName(), where, orderBy)
	v, err := tbl.doQueryAndScanOne(req, query, args...)
	return v, err
}

// SelectOne allows a single Hook to be obtained from the database.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
// If not found, *Example will be nil.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.One
// controls whether an error is generated when no result is found.
func (tbl HookTable) SelectOne(req require.Requirement, wh where.Expression, qc where.QueryConstraint) (*Hook, error) {
	q := tbl.Dialect().Quoter()
	whs, args := where.Where(wh, q)
	orderBy := where.Build(qc, q)
	return tbl.SelectOneWhere(req, whs, orderBy, args...)
}

// SelectWhere allows Hooks to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
//
// The args are for any placeholder parameters in the query.
func (tbl HookTable) SelectWhere(req require.Requirement, where, orderBy string, args ...interface{}) (HookList, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s %s",
		allHookColumnNamesQuoted(tbl.Dialect().Quoter()), tbl.quotedName(), where, orderBy)
	vv, err := tbl.doQueryAndScan(req, false, query, args...)
	return vv, err
}

// Select allows Hooks to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
func (tbl HookTable) Select(req require.Requirement, wh where.Expression, qc where.QueryConstraint) (HookList, error) {
	q := tbl.Dialect().Quoter()
	whs, args := where.Where(wh, q)
	orderBy := where.Build(qc, q)
	return tbl.SelectWhere(req, whs, orderBy, args...)
}

//--------------------------------------------------------------------------------

// CountWhere counts Hooks in the table that match a 'where' clause.
// Use a blank string for the 'where' argument if it is not needed.
//
// The args are for any placeholder parameters in the query.
func (tbl HookTable) CountWhere(where string, args ...interface{}) (count int64, err error) {
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

// Count counts the Hooks in the table that match a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed.
func (tbl HookTable) Count(wh where.Expression) (count int64, err error) {
	whs, args := where.Where(wh, tbl.Dialect().Quoter())
	return tbl.CountWhere(whs, args...)
}

//--------------------------------------------------------------------------------

// SliceId gets the id column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceId(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]uint64, error) {
	return support.SliceUint64List(tbl, req, tbl.pk, wh, qc)
}

// SliceSha gets the sha column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceSha(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return support.SliceStringList(tbl, req, "sha", wh, qc)
}

// SliceAfter gets the after column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceAfter(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return support.SliceStringList(tbl, req, "after", wh, qc)
}

// SliceBefore gets the before column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceBefore(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return support.SliceStringList(tbl, req, "before", wh, qc)
}

// SliceCommitId gets the commit_id column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceCommitId(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return support.SliceStringList(tbl, req, "commit_id", wh, qc)
}

// SliceMessage gets the message column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceMessage(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return support.SliceStringList(tbl, req, "message", wh, qc)
}

// SliceTimestamp gets the timestamp column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceTimestamp(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return support.SliceStringList(tbl, req, "timestamp", wh, qc)
}

// SliceHeadCommitAuthorName gets the head_commit_author_name column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceHeadCommitAuthorName(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return support.SliceStringList(tbl, req, "head_commit_author_name", wh, qc)
}

// SliceHeadCommitAuthorUsername gets the head_commit_author_username column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceHeadCommitAuthorUsername(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return support.SliceStringList(tbl, req, "head_commit_author_username", wh, qc)
}

// SliceHeadCommitCommitterName gets the head_commit_committer_name column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceHeadCommitCommitterName(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return support.SliceStringList(tbl, req, "head_commit_committer_name", wh, qc)
}

// SliceHeadCommitCommitterUsername gets the head_commit_committer_username column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceHeadCommitCommitterUsername(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return support.SliceStringList(tbl, req, "head_commit_committer_username", wh, qc)
}

// SliceCategory gets the category column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceCategory(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]Category, error) {
	return tbl.sliceCategoryList(req, "category", wh, qc)
}

// SliceHeadCommitAuthorEmail gets the head_commit_author_email column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceHeadCommitAuthorEmail(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]Email, error) {
	return tbl.sliceEmailList(req, "head_commit_author_email", wh, qc)
}

// SliceHeadCommitCommitterEmail gets the head_commit_committer_email column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceHeadCommitCommitterEmail(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]Email, error) {
	return tbl.sliceEmailList(req, "head_commit_committer_email", wh, qc)
}

func (tbl HookTable) sliceCategoryList(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]Category, error) {
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

func (tbl HookTable) sliceEmailList(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]Email, error) {
	q := tbl.Dialect().Quoter()
	whs, args := where.Where(wh, q)
	orderBy := where.Build(qc, q)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", q.Quote(sqlname), tbl.quotedName(), whs, orderBy)
	rows, err := support.Query(tbl, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var v Email
	list := make([]Email, 0, 10)

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

func (tbl HookTable) constructHookInsert(w dialect.StringWriter, v *Hook, withPk bool) (s []interface{}, err error) {
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

func (tbl HookTable) constructHookUpdate(w dialect.StringWriter, v *Hook) (s []interface{}, err error) {
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

// Insert adds new records for the Hooks.
// The Hooks have their primary key fields set to the new record identifiers.
// The Hook.PreInsert() method will be called, if it exists.
func (tbl HookTable) Insert(req require.Requirement, vv ...*Hook) error {
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

		fields, err := tbl.constructHookInsert(b, v, false)
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
			i64, e2 := tbl.db.InsertContext(tbl.ctx, query, fields...)
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

// UpdateFields updates one or more columns, given a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl HookTable) UpdateFields(req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error) {
	return support.UpdateFields(tbl, req, wh, fields...)
}

//--------------------------------------------------------------------------------

// Update updates records, matching them by primary key. It returns the number of rows affected.
// The Hook.PreUpdate(Execer) method will be called, if it exists.
func (tbl HookTable) Update(req require.Requirement, vv ...*Hook) (int64, error) {
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

		args, err := tbl.constructHookUpdate(b, v)
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
func (tbl HookTable) Upsert(v *Hook, wh where.Expression) error {
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

//--------------------------------------------------------------------------------

// DeleteHooks deletes rows from the table, given some primary keys.
// The list of ids can be arbitrarily long.
func (tbl HookTable) DeleteHooks(req require.Requirement, id ...uint64) (int64, error) {
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
func (tbl HookTable) Delete(req require.Requirement, wh where.Expression) (int64, error) {
	query, args := tbl.deleteRows(wh)
	return tbl.Exec(req, query, args...)
}

func (tbl HookTable) deleteRows(wh where.Expression) (string, []interface{}) {
	whs, args := where.Where(wh, tbl.Dialect().Quoter())
	query := fmt.Sprintf("DELETE FROM %s %s", tbl.quotedName(), whs)
	return query, args
}

//--------------------------------------------------------------------------------
