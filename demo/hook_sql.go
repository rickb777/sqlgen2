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

// HookTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type HookTable struct {
	name        sqlgen2.TableName
	database    *sqlgen2.Database
	db          sqlgen2.Execer
	constraints constraint.Constraints
	ctx			context.Context
	pk          string
}

// Type conformance checks
var _ sqlgen2.TableCreator = &HookTable{}
var _ sqlgen2.TableWithCrud = &HookTable{}

// NewHookTable returns a new table instance.
// If a blank table name is supplied, the default name "hooks" will be used instead.
// The request context is initialised with the background.
func NewHookTable(name string, d *sqlgen2.Database) HookTable {
	if name == "" {
		name = "hooks"
	}
	var constraints constraint.Constraints
	return HookTable{
		name:        sqlgen2.TableName{"", name},
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
func CopyTableAsHookTable(origin sqlgen2.Table) HookTable {
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
func (tbl HookTable) Database() *sqlgen2.Database {
	return tbl.database
}

// Logger gets the trace logger.
func (tbl HookTable) Logger() *log.Logger {
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
func (tbl HookTable) Dialect() schema.Dialect {
	return tbl.database.Dialect()
}

// Name gets the table name.
func (tbl HookTable) Name() sqlgen2.TableName {
	return tbl.name
}


// PkColumn gets the column name used as a primary key.
func (tbl HookTable) PkColumn() string {
	return tbl.pk
}


// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl HookTable) DB() *sql.DB {
	return tbl.db.(*sql.DB)
}

// Execer gets the wrapped database or transaction handle.
func (tbl HookTable) Execer() sqlgen2.Execer {
	return tbl.db
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl HookTable) Tx() *sql.Tx {
	return tbl.db.(*sql.Tx)
}

// IsTx tests whether this is within a transaction.
func (tbl HookTable) IsTx() bool {
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
func (tbl HookTable) BeginTx(opts *sql.TxOptions) (HookTable, error) {
	var err error
	tbl.db, err = tbl.db.(sqlgen2.TxStarter).BeginTx(tbl.ctx, opts)
	return tbl, tbl.logIfError(err)
}

// Using returns a modified Table using the transaction supplied. This is needed
// when making multiple queries across several tables within a single transaction.
// The result is a modified copy of the table; the original is unchanged.
func (tbl HookTable) Using(tx *sql.Tx) HookTable {
	tbl.db = tx
	return tbl
}

func (tbl HookTable) logQuery(query string, args ...interface{}) {
	tbl.database.LogQuery(query, args...)
}

func (tbl HookTable) logError(err error) error {
	return tbl.database.LogError(err)
}

func (tbl HookTable) logIfError(err error) error {
	return tbl.database.LogIfError(err)
}


//--------------------------------------------------------------------------------

const NumHookColumns = 17

const NumHookDataColumns = 16

const HookColumnNames = "id,sha,after,before,category,created,deleted,forced,commit_id,message,timestamp,head_commit_author_name,head_commit_author_email,head_commit_author_username,head_commit_committer_name,head_commit_committer_email,head_commit_committer_username"

const HookDataColumnNames = "sha,after,before,category,created,deleted,forced,commit_id,message,timestamp,head_commit_author_name,head_commit_author_email,head_commit_author_username,head_commit_committer_name,head_commit_committer_email,head_commit_committer_username"

//--------------------------------------------------------------------------------

const sqlHookTableCreateColumnsSqlite = "\n"+
" `id`                             integer not null primary key autoincrement,\n"+
" `sha`                            text not null,\n"+
" `after`                          text not null,\n"+
" `before`                         text not null,\n"+
" `category`                       tinyint unsigned not null,\n"+
" `created`                        boolean not null,\n"+
" `deleted`                        boolean not null,\n"+
" `forced`                         boolean not null,\n"+
" `commit_id`                      text not null,\n"+
" `message`                        text not null,\n"+
" `timestamp`                      text not null,\n"+
" `head_commit_author_name`        text not null,\n"+
" `head_commit_author_email`       text not null,\n"+
" `head_commit_author_username`    text not null,\n"+
" `head_commit_committer_name`     text not null,\n"+
" `head_commit_committer_email`    text not null,\n"+
" `head_commit_committer_username` text not null"

const sqlHookTableCreateColumnsMysql = "\n"+
" `id`                             bigint unsigned not null primary key auto_increment,\n"+
" `sha`                            varchar(255) not null,\n"+
" `after`                          varchar(20) not null,\n"+
" `before`                         varchar(20) not null,\n"+
" `category`                       tinyint unsigned not null,\n"+
" `created`                        tinyint(1) not null,\n"+
" `deleted`                        tinyint(1) not null,\n"+
" `forced`                         tinyint(1) not null,\n"+
" `commit_id`                      varchar(255) not null,\n"+
" `message`                        varchar(255) not null,\n"+
" `timestamp`                      varchar(255) not null,\n"+
" `head_commit_author_name`        varchar(255) not null,\n"+
" `head_commit_author_email`       varchar(255) not null,\n"+
" `head_commit_author_username`    varchar(255) not null,\n"+
" `head_commit_committer_name`     varchar(255) not null,\n"+
" `head_commit_committer_email`    varchar(255) not null,\n"+
" `head_commit_committer_username` varchar(255) not null"

const sqlHookTableCreateColumnsPostgres = `
 "id"                             bigserial not null primary key,
 "sha"                            varchar(255) not null,
 "after"                          varchar(20) not null,
 "before"                         varchar(20) not null,
 "category"                       smallint not null,
 "created"                        boolean not null,
 "deleted"                        boolean not null,
 "forced"                         boolean not null,
 "commit_id"                      varchar(255) not null,
 "message"                        varchar(255) not null,
 "timestamp"                      varchar(255) not null,
 "head_commit_author_name"        varchar(255) not null,
 "head_commit_author_email"       varchar(255) not null,
 "head_commit_author_username"    varchar(255) not null,
 "head_commit_committer_name"     varchar(255) not null,
 "head_commit_committer_email"    varchar(255) not null,
 "head_commit_committer_username" varchar(255) not null`

const sqlConstrainHookTable = `
`

//--------------------------------------------------------------------------------

// CreateTable creates the table.
func (tbl HookTable) CreateTable(ifNotExists bool) (int64, error) {
	return support.Exec(tbl, nil, tbl.createTableSql(ifNotExists))
}

func (tbl HookTable) createTableSql(ifNotExists bool) string {
	var columns string
	var settings string
	switch tbl.Dialect() {
	case schema.Sqlite:
		columns = sqlHookTableCreateColumnsSqlite
		settings = ""
    case schema.Mysql:
		columns = sqlHookTableCreateColumnsMysql
		settings = " ENGINE=InnoDB DEFAULT CHARSET=utf8"
    case schema.Postgres:
		columns = sqlHookTableCreateColumnsPostgres
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
// Wrap the result in *sqlgen2.Rows if you need to access its data as a map.
func (tbl HookTable) Query(query string, args ...interface{}) (*sql.Rows, error) {
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

func scanHooks(rows *sql.Rows, firstOnly bool) (vv HookList, n int64, err error) {
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
			return vv, n, err
		}

		v := &Hook{}
		v.Id = v0
		v.Sha = v1
		v.Dates.After = v2
		v.Dates.Before = v3
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

var allHookQuotedColumnNames = []string{
	schema.Sqlite.SplitAndQuote(HookColumnNames),
	schema.Mysql.SplitAndQuote(HookColumnNames),
	schema.Postgres.SplitAndQuote(HookColumnNames),
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
	dialect := tbl.Dialect()
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s=%s",
		allHookQuotedColumnNames[dialect.Index()], tbl.name, dialect.Quote(column), dialect.Placeholder(column, 1))
	v, err := tbl.doQueryOne(req, query, arg)
	return v, err
}

func (tbl HookTable) getHooks(req require.Requirement, column string, args ...interface{}) (list HookList, err error) {
	if len(args) > 0 {
		if req == require.All {
			req = require.Exactly(len(args))
		}
		dialect := tbl.Dialect()
		pl := dialect.Placeholders(len(args))
		query := fmt.Sprintf("SELECT %s FROM %s WHERE %s IN (%s)",
			allHookQuotedColumnNames[dialect.Index()], tbl.name, dialect.Quote(column), pl)
		list, err = tbl.doQuery(req, false, query, args...)
	}

	return list, err
}

func (tbl HookTable) doQueryOne(req require.Requirement, query string, args ...interface{}) (*Hook, error) {
	list, err := tbl.doQuery(req, true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

func (tbl HookTable) doQuery(req require.Requirement, firstOnly bool, query string, args ...interface{}) (HookList, error) {
	rows, err := tbl.Query(query, args...)
	if err != nil {
		return nil, err
	}

	vv, n, err := scanHooks(rows, firstOnly)
	return vv, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(err, req, n))
}

// Fetch fetches a list of Hook based on a supplied query. This is mostly used for join queries that map its
// result columns to the fields of Hook. Other queries might be better handled by GetXxx or Select methods.
func (tbl HookTable) Fetch(req require.Requirement, query string, args ...interface{}) (HookList, error) {
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
func (tbl HookTable) SelectOneWhere(req require.Requirement, where, orderBy string, args ...interface{}) (*Hook, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT 1",
		allHookQuotedColumnNames[tbl.Dialect().Index()], tbl.name, where, orderBy)
	v, err := tbl.doQueryOne(req, query, args...)
	return v, err
}

// SelectOne allows a single Hook to be obtained from the sqlgen2.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
// If not found, *Example will be nil.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.One
// controls whether an error is generated when no result is found.
func (tbl HookTable) SelectOne(req require.Requirement, wh where.Expression, qc where.QueryConstraint) (*Hook, error) {
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
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
		allHookQuotedColumnNames[tbl.Dialect().Index()], tbl.name, where, orderBy)
	vv, err := tbl.doQuery(req, false, query, args...)
	return vv, err
}

// Select allows Hooks to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
func (tbl HookTable) Select(req require.Requirement, wh where.Expression, qc where.QueryConstraint) (HookList, error) {
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
	return tbl.SelectWhere(req, whs, orderBy, args...)
}

// CountWhere counts Hooks in the table that match a 'where' clause.
// Use a blank string for the 'where' argument if it is not needed.
//
// The args are for any placeholder parameters in the query.
func (tbl HookTable) CountWhere(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s %s", tbl.name, where)
	tbl.logQuery(query, args...)
	row := tbl.db.QueryRowContext(tbl.ctx, query, args...)
	err = row.Scan(&count)
	return count, tbl.logIfError(err)
}

// Count counts the Hooks in the table that match a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed.
func (tbl HookTable) Count(wh where.Expression) (count int64, err error) {
	whs, args := where.BuildExpression(wh, tbl.Dialect())
	return tbl.CountWhere(whs, args...)
}

//--------------------------------------------------------------------------------

// SliceId gets the id column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceId(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]uint64, error) {
	return tbl.sliceUint64List(req, tbl.pk, wh, qc)
}

// SliceSha gets the sha column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceSha(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.sliceStringList(req, "sha", wh, qc)
}

// SliceAfter gets the after column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceAfter(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.sliceStringList(req, "after", wh, qc)
}

// SliceBefore gets the before column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceBefore(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.sliceStringList(req, "before", wh, qc)
}

// SliceCategory gets the category column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceCategory(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]Category, error) {
	return tbl.sliceCategoryList(req, "category", wh, qc)
}

// SliceCommitId gets the commit_id column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceCommitId(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.sliceStringList(req, "commit_id", wh, qc)
}

// SliceMessage gets the message column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceMessage(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.sliceStringList(req, "message", wh, qc)
}

// SliceTimestamp gets the timestamp column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceTimestamp(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.sliceStringList(req, "timestamp", wh, qc)
}

// SliceHeadCommitAuthorName gets the head_commit_author_name column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceHeadCommitAuthorName(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.sliceStringList(req, "head_commit_author_name", wh, qc)
}

// SliceHeadCommitAuthorEmail gets the head_commit_author_email column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceHeadCommitAuthorEmail(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]Email, error) {
	return tbl.sliceEmailList(req, "head_commit_author_email", wh, qc)
}

// SliceHeadCommitAuthorUsername gets the head_commit_author_username column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceHeadCommitAuthorUsername(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.sliceStringList(req, "head_commit_author_username", wh, qc)
}

// SliceHeadCommitCommitterName gets the head_commit_committer_name column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceHeadCommitCommitterName(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.sliceStringList(req, "head_commit_committer_name", wh, qc)
}

// SliceHeadCommitCommitterEmail gets the head_commit_committer_email column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceHeadCommitCommitterEmail(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]Email, error) {
	return tbl.sliceEmailList(req, "head_commit_committer_email", wh, qc)
}

// SliceHeadCommitCommitterUsername gets the head_commit_committer_username column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceHeadCommitCommitterUsername(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.sliceStringList(req, "head_commit_committer_username", wh, qc)
}

func (tbl HookTable) sliceCategoryList(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]Category, error) {
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

func (tbl HookTable) sliceEmailList(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]Email, error) {
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

	var v Email
	list := make([]Email, 0, 10)

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

func (tbl HookTable) sliceStringList(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
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

func (tbl HookTable) sliceUint64List(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]uint64, error) {
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

	var v uint64
	list := make([]uint64, 0, 10)

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


func constructHookInsert(w io.Writer, v *Hook, dialect schema.Dialect, withPk bool) (s []interface{}, err error) {
	s = make([]interface{}, 0, 17)

	comma := ""
	io.WriteString(w, " (")

	if withPk {
		dialect.QuoteW(w, "id")
		comma = ","
		s = append(s, v.Id)
	}

	io.WriteString(w, comma)

	dialect.QuoteW(w, "sha")
	s = append(s, v.Sha)
	comma = ","
	io.WriteString(w, comma)

	dialect.QuoteW(w, "after")
	s = append(s, v.Dates.After)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "before")
	s = append(s, v.Dates.Before)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "category")
	s = append(s, v.Category)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "created")
	s = append(s, v.Created)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "deleted")
	s = append(s, v.Deleted)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "forced")
	s = append(s, v.Forced)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "commit_id")
	s = append(s, v.HeadCommit.ID)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "message")
	s = append(s, v.HeadCommit.Message)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "timestamp")
	s = append(s, v.HeadCommit.Timestamp)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "head_commit_author_name")
	s = append(s, v.HeadCommit.Author.Name)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "head_commit_author_email")
	s = append(s, v.HeadCommit.Author.Email)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "head_commit_author_username")
	s = append(s, v.HeadCommit.Author.Username)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "head_commit_committer_name")
	s = append(s, v.HeadCommit.Committer.Name)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "head_commit_committer_email")
	s = append(s, v.HeadCommit.Committer.Email)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "head_commit_committer_username")
	s = append(s, v.HeadCommit.Committer.Username)
	io.WriteString(w, ")")
	return s, nil
}

func constructHookUpdate(w io.Writer, v *Hook, dialect schema.Dialect) (s []interface{}, err error) {
	j := 1
	s = make([]interface{}, 0, 16)

	comma := ""

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "sha", j)
	s = append(s, v.Sha)
	comma = ", "
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "after", j)
	s = append(s, v.Dates.After)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "before", j)
	s = append(s, v.Dates.Before)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "category", j)
	s = append(s, v.Category)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "created", j)
	s = append(s, v.Created)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "deleted", j)
	s = append(s, v.Deleted)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "forced", j)
	s = append(s, v.Forced)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "commit_id", j)
	s = append(s, v.HeadCommit.ID)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "message", j)
	s = append(s, v.HeadCommit.Message)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "timestamp", j)
	s = append(s, v.HeadCommit.Timestamp)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "head_commit_author_name", j)
	s = append(s, v.HeadCommit.Author.Name)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "head_commit_author_email", j)
	s = append(s, v.HeadCommit.Author.Email)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "head_commit_author_username", j)
	s = append(s, v.HeadCommit.Author.Username)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "head_commit_committer_name", j)
	s = append(s, v.HeadCommit.Committer.Name)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "head_commit_committer_email", j)
	s = append(s, v.HeadCommit.Committer.Email)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "head_commit_committer_username", j)
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

		fields, err := constructHookInsert(b, v, tbl.Dialect(), false)
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
			var i64 int64
			err = row.Scan(&i64)
			v.Id = uint64(i64)

		} else {
			res, e2 := tbl.db.ExecContext(tbl.ctx, query, fields...)
			if e2 != nil {
				return tbl.logError(e2)
			}

			i64, e2 := res.LastInsertId()
			v.Id = uint64(i64)
			
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
func (tbl HookTable) UpdateFields(req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error) {
	return support.UpdateFields(tbl, req, wh, fields...)
}

//--------------------------------------------------------------------------------

var allHookQuotedUpdates = []string{
	// Sqlite
	"`sha`=?,`after`=?,`before`=?,`category`=?,`created`=?,`deleted`=?,`forced`=?,`commit_id`=?,`message`=?,`timestamp`=?,`head_commit_author_name`=?,`head_commit_author_email`=?,`head_commit_author_username`=?,`head_commit_committer_name`=?,`head_commit_committer_email`=?,`head_commit_committer_username`=? WHERE `id`=?",
	// Mysql
	"`sha`=?,`after`=?,`before`=?,`category`=?,`created`=?,`deleted`=?,`forced`=?,`commit_id`=?,`message`=?,`timestamp`=?,`head_commit_author_name`=?,`head_commit_author_email`=?,`head_commit_author_username`=?,`head_commit_committer_name`=?,`head_commit_committer_email`=?,`head_commit_committer_username`=? WHERE `id`=?",
	// Postgres
	`"sha"=$2,"after"=$3,"before"=$4,"category"=$5,"created"=$6,"deleted"=$7,"forced"=$8,"commit_id"=$9,"message"=$10,"timestamp"=$11,"head_commit_author_name"=$12,"head_commit_author_email"=$13,"head_commit_author_username"=$14,"head_commit_committer_name"=$15,"head_commit_committer_email"=$16,"head_commit_committer_username"=$17 WHERE "id"=$1`,
}

//--------------------------------------------------------------------------------

// Update updates records, matching them by primary key. It returns the number of rows affected.
// The Hook.PreUpdate(Execer) method will be called, if it exists.
func (tbl HookTable) Update(req require.Requirement, vv ...*Hook) (int64, error) {
	if req == require.All {
		req = require.Exactly(len(vv))
	}

	var count int64
	dialect := tbl.Dialect()
	//columns := allHookQuotedUpdates[dialect.Index()]
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

		args, err := constructHookUpdate(b, v, dialect)
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

// DeleteHooks deletes rows from the table, given some primary keys.
// The list of ids can be arbitrarily long.
func (tbl HookTable) DeleteHooks(req require.Requirement, id ...uint64) (int64, error) {
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
func (tbl HookTable) Delete(req require.Requirement, wh where.Expression) (int64, error) {
	query, args := tbl.deleteRows(wh)
	return tbl.Exec(req, query, args...)
}

func (tbl HookTable) deleteRows(wh where.Expression) (string, []interface{}) {
	whs, args := where.BuildExpression(wh, tbl.Dialect())
	query := fmt.Sprintf("DELETE FROM %s %s", tbl.name, whs)
	return query, args
}

//--------------------------------------------------------------------------------
