// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

package demo

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
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

// IssueTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type IssueTable struct {
	name        sqlgen2.TableName
	database    *sqlgen2.Database
	db          sqlgen2.Execer
	constraints constraint.Constraints
	ctx			context.Context
	pk          string
}

// Type conformance checks
var _ sqlgen2.TableWithIndexes = &IssueTable{}
var _ sqlgen2.TableWithCrud = &IssueTable{}

// NewIssueTable returns a new table instance.
// If a blank table name is supplied, the default name "issues" will be used instead.
// The request context is initialised with the background.
func NewIssueTable(name string, d *sqlgen2.Database) IssueTable {
	if name == "" {
		name = "issues"
	}
	var constraints constraint.Constraints
	return IssueTable{
		name:        sqlgen2.TableName{"", name},
		database:    d,
		db:          d.DB(),
		constraints: constraints,
		ctx:         context.Background(),
		pk:          "id",
	}
}

// CopyTableAsIssueTable copies a table instance, retaining the name etc but
// providing methods appropriate for 'Issue'. It doesn't copy the constraints of the original table.
//
// It serves to provide methods appropriate for 'Issue'. This is most useful when this is used to represent a
// join result. In such cases, there won't be any need for DDL methods, nor Exec, Insert, Update or Delete.
func CopyTableAsIssueTable(origin sqlgen2.Table) IssueTable {
	return IssueTable{
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
func (tbl IssueTable) SetPkColumn(pk string) IssueTable {
	tbl.pk = pk
	return tbl
}


// WithPrefix sets the table name prefix for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl IssueTable) WithPrefix(pfx string) IssueTable {
	tbl.name.Prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries via this table.
// The result is a modified copy of the table; the original is unchanged.
//
// The shared context in the *Database is not altered by this method. So it
// is possible to use different contexts for different (groups of) queries.
func (tbl IssueTable) WithContext(ctx context.Context) IssueTable {
	tbl.ctx = ctx
	return tbl
}

// Database gets the shared database information.
func (tbl IssueTable) Database() *sqlgen2.Database {
	return tbl.database
}

// Logger gets the trace logger.
func (tbl IssueTable) Logger() *log.Logger {
	return tbl.database.Logger()
}

// WithConstraint returns a modified Table with added data consistency constraints.
func (tbl IssueTable) WithConstraint(cc ...constraint.Constraint) IssueTable {
	tbl.constraints = append(tbl.constraints, cc...)
	return tbl
}

// Constraints returns the table's constraints.
func (tbl IssueTable) Constraints() constraint.Constraints {
	return tbl.constraints
}

// Ctx gets the current request context.
func (tbl IssueTable) Ctx() context.Context {
	return tbl.ctx
}

// Dialect gets the database dialect.
func (tbl IssueTable) Dialect() schema.Dialect {
	return tbl.database.Dialect()
}

// Name gets the table name.
func (tbl IssueTable) Name() sqlgen2.TableName {
	return tbl.name
}


// PkColumn gets the column name used as a primary key.
func (tbl IssueTable) PkColumn() string {
	return tbl.pk
}


// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl IssueTable) DB() *sql.DB {
	return tbl.db.(*sql.DB)
}

// Execer gets the wrapped database or transaction handle.
func (tbl IssueTable) Execer() sqlgen2.Execer {
	return tbl.db
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl IssueTable) Tx() *sql.Tx {
	return tbl.db.(*sql.Tx)
}

// IsTx tests whether this is within a transaction.
func (tbl IssueTable) IsTx() bool {
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
func (tbl IssueTable) BeginTx(opts *sql.TxOptions) (IssueTable, error) {
	var err error
	tbl.db, err = tbl.db.(sqlgen2.TxStarter).BeginTx(tbl.ctx, opts)
	return tbl, tbl.logIfError(err)
}

// Using returns a modified Table using the transaction supplied. This is needed
// when making multiple queries across several tables within a single transaction.
// The result is a modified copy of the table; the original is unchanged.
func (tbl IssueTable) Using(tx *sql.Tx) IssueTable {
	tbl.db = tx
	return tbl
}

func (tbl IssueTable) logQuery(query string, args ...interface{}) {
	tbl.database.LogQuery(query, args...)
}

func (tbl IssueTable) logError(err error) error {
	return tbl.database.LogError(err)
}

func (tbl IssueTable) logIfError(err error) error {
	return tbl.database.LogIfError(err)
}


//--------------------------------------------------------------------------------

const NumIssueColumns = 8

const NumIssueDataColumns = 7

const IssueColumnNames = "id,number,date,title,bigbody,assignee,state,labels"

const IssueDataColumnNames = "number,date,title,bigbody,assignee,state,labels"

//--------------------------------------------------------------------------------

const sqlIssueTableCreateColumnsSqlite = "\n"+
" `id`       integer not null primary key autoincrement,\n"+
" `number`   bigint not null,\n"+
" `date`     blob not null,\n"+
" `title`    text not null,\n"+
" `bigbody`  text not null,\n"+
" `assignee` text not null,\n"+
" `state`    text not null,\n"+
" `labels`   text"

const sqlIssueTableCreateColumnsMysql = "\n"+
" `id`       bigint not null primary key auto_increment,\n"+
" `number`   bigint not null,\n"+
" `date`     mediumblob not null,\n"+
" `title`    varchar(512) not null,\n"+
" `bigbody`  varchar(2048) not null,\n"+
" `assignee` varchar(255) not null,\n"+
" `state`    varchar(50) not null,\n"+
" `labels`   json"

const sqlIssueTableCreateColumnsPostgres = `
 "id"       bigserial not null primary key,
 "number"   bigint not null,
 "date"     bytea not null,
 "title"    varchar(512) not null,
 "bigbody"  varchar(2048) not null,
 "assignee" varchar(255) not null,
 "state"    varchar(50) not null,
 "labels"   json`

const sqlConstrainIssueTable = `
`

//--------------------------------------------------------------------------------

const sqlIssueAssigneeIndexColumns = "assignee"

//--------------------------------------------------------------------------------

// CreateTable creates the table.
func (tbl IssueTable) CreateTable(ifNotExists bool) (int64, error) {
	return support.Exec(tbl, nil, tbl.createTableSql(ifNotExists))
}

func (tbl IssueTable) createTableSql(ifNotExists bool) string {
	var columns string
	var settings string
	switch tbl.Dialect() {
	case schema.Sqlite:
		columns = sqlIssueTableCreateColumnsSqlite
		settings = ""
    case schema.Mysql:
		columns = sqlIssueTableCreateColumnsMysql
		settings = " ENGINE=InnoDB DEFAULT CHARSET=utf8"
    case schema.Postgres:
		columns = sqlIssueTableCreateColumnsPostgres
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

func (tbl IssueTable) ternary(flag bool, a, b string) string {
	if flag {
		return a
	}
	return b
}

// DropTable drops the table, destroying all its data.
func (tbl IssueTable) DropTable(ifExists bool) (int64, error) {
	return support.Exec(tbl, nil, tbl.dropTableSql(ifExists))
}

func (tbl IssueTable) dropTableSql(ifExists bool) string {
	ie := tbl.ternary(ifExists, "IF EXISTS ", "")
	query := fmt.Sprintf("DROP TABLE %s%s", ie, tbl.name)
	return query
}

//--------------------------------------------------------------------------------

// CreateTableWithIndexes invokes CreateTable then CreateIndexes.
func (tbl IssueTable) CreateTableWithIndexes(ifNotExist bool) (err error) {
	_, err = tbl.CreateTable(ifNotExist)
	if err != nil {
		return err
	}

	return tbl.CreateIndexes(ifNotExist)
}

// CreateIndexes executes queries that create the indexes needed by the Issue table.
func (tbl IssueTable) CreateIndexes(ifNotExist bool) (err error) {

	err = tbl.CreateIssueAssigneeIndex(ifNotExist)
	if err != nil {
		return err
	}

	return nil
}

// CreateIssueAssigneeIndex creates the issue_assignee index.
func (tbl IssueTable) CreateIssueAssigneeIndex(ifNotExist bool) error {
	ine := tbl.ternary(ifNotExist && tbl.Dialect() != schema.Mysql, "IF NOT EXISTS ", "")

	// Mysql does not support 'if not exists' on indexes
	// Workaround: use DropIndex first and ignore an error returned if the index didn't exist.

	if ifNotExist && tbl.Dialect() == schema.Mysql {
		// low-level no-logging Exec
		tbl.Execer().ExecContext(tbl.ctx, tbl.dropIssueAssigneeIndexSql(false))
		ine = ""
	}

	_, err := tbl.Exec(nil, tbl.createIssueAssigneeIndexSql(ine))
	return err
}

func (tbl IssueTable) createIssueAssigneeIndexSql(ifNotExists string) string {
	indexPrefix := tbl.name.PrefixWithoutDot()
	return fmt.Sprintf("CREATE INDEX %s%sissue_assignee ON %s (%s)", ifNotExists, indexPrefix,
		tbl.name, sqlIssueAssigneeIndexColumns)
}

// DropIssueAssigneeIndex drops the issue_assignee index.
func (tbl IssueTable) DropIssueAssigneeIndex(ifExists bool) error {
	_, err := tbl.Exec(nil, tbl.dropIssueAssigneeIndexSql(ifExists))
	return err
}

func (tbl IssueTable) dropIssueAssigneeIndexSql(ifExists bool) string {
	// Mysql does not support 'if exists' on indexes
	ie := tbl.ternary(ifExists && tbl.Dialect() != schema.Mysql, "IF EXISTS ", "")
	onTbl := tbl.ternary(tbl.Dialect() == schema.Mysql, fmt.Sprintf(" ON %s", tbl.name), "")
	indexPrefix := tbl.name.PrefixWithoutDot()
	return fmt.Sprintf("DROP INDEX %s%sissue_assignee%s", ie, indexPrefix, onTbl)
}

// DropIndexes executes queries that drop the indexes on by the Issue table.
func (tbl IssueTable) DropIndexes(ifExist bool) (err error) {

	err = tbl.DropIssueAssigneeIndex(ifExist)
	if err != nil {
		return err
	}

	return nil
}

//--------------------------------------------------------------------------------

// Truncate drops every record from the table, if possible. It might fail if constraints exist that
// prevent some or all rows from being deleted; use the force option to override this.
//
// When 'force' is set true, be aware of the following consequences.
// When using Mysql, foreign keys in other tables can be left dangling.
// When using Postgres, a cascade happens, so all 'adjacent' tables (i.e. linked by foreign keys)
// are also truncated.
func (tbl IssueTable) Truncate(force bool) (err error) {
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
func (tbl IssueTable) Exec(req require.Requirement, query string, args ...interface{}) (int64, error) {
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
func (tbl IssueTable) Query(query string, args ...interface{}) (*sql.Rows, error) {
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
func (tbl IssueTable) QueryOneNullString(req require.Requirement, query string, args ...interface{}) (result sql.NullString, err error) {
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
func (tbl IssueTable) QueryOneNullInt64(req require.Requirement, query string, args ...interface{}) (result sql.NullInt64, err error) {
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
func (tbl IssueTable) QueryOneNullFloat64(req require.Requirement, query string, args ...interface{}) (result sql.NullFloat64, err error) {
	err = support.QueryOneNullThing(tbl, req, &result, query, args...)
	return result, err
}

func scanIssues(rows *sql.Rows, firstOnly bool) (vv []*Issue, n int64, err error) {
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
			return vv, n, err
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
			return nil, n, err
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

var allIssueQuotedColumnNames = []string{
	schema.Sqlite.SplitAndQuote(IssueColumnNames),
	schema.Mysql.SplitAndQuote(IssueColumnNames),
	schema.Postgres.SplitAndQuote(IssueColumnNames),
}

//--------------------------------------------------------------------------------

// GetIssuesById gets records from the table according to a list of primary keys.
// Although the list of ids can be arbitrarily long, there are practical limits;
// note that Oracle DB has a limit of 1000.
//
// It places a requirement, which may be nil, on the size of the expected results: in particular, require.All
// controls whether an error is generated not all the ids produce a result.
func (tbl IssueTable) GetIssuesById(req require.Requirement, id ...int64) (list []*Issue, err error) {
	if len(id) > 0 {
		if req == require.All {
			req = require.Exactly(len(id))
		}
		args := make([]interface{}, len(id))

		for i, v := range id {
			args[i] = v
		}

		list, err = tbl.getIssues(req, tbl.pk, args...)
	}

	return list, err
}

// GetIssueById gets the record with a given primary key value.
// If not found, *Issue will be nil.
func (tbl IssueTable) GetIssueById(req require.Requirement, id int64) (*Issue, error) {
	return tbl.getIssue(req, tbl.pk, id)
}

// GetIssuesByAssignee gets the records with a given assignee value.
// If not found, the resulting slice will be empty (nil).
func (tbl IssueTable) GetIssuesByAssignee(req require.Requirement, assignee string) ([]*Issue, error) {
	return tbl.Select(req, where.And(where.Eq("assignee", assignee)), nil)
}

func (tbl IssueTable) getIssue(req require.Requirement, column string, arg interface{}) (*Issue, error) {
	dialect := tbl.Dialect()
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s=%s",
		allIssueQuotedColumnNames[dialect.Index()], tbl.name, dialect.Quote(column), dialect.Placeholder(column, 1))
	v, err := tbl.doQueryOne(req, query, arg)
	return v, err
}

func (tbl IssueTable) getIssues(req require.Requirement, column string, args ...interface{}) (list []*Issue, err error) {
	if len(args) > 0 {
		if req == require.All {
			req = require.Exactly(len(args))
		}
		dialect := tbl.Dialect()
		pl := dialect.Placeholders(len(args))
		query := fmt.Sprintf("SELECT %s FROM %s WHERE %s IN (%s)",
			allIssueQuotedColumnNames[dialect.Index()], tbl.name, dialect.Quote(column), pl)
		list, err = tbl.doQuery(req, false, query, args...)
	}

	return list, err
}

func (tbl IssueTable) doQueryOne(req require.Requirement, query string, args ...interface{}) (*Issue, error) {
	list, err := tbl.doQuery(req, true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

func (tbl IssueTable) doQuery(req require.Requirement, firstOnly bool, query string, args ...interface{}) ([]*Issue, error) {
	rows, err := tbl.Query(query, args...)
	if err != nil {
		return nil, err
	}

	vv, n, err := scanIssues(rows, firstOnly)
	return vv, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(err, req, n))
}

// Fetch fetches a list of Issue based on a supplied query. This is mostly used for join queries that map its
// result columns to the fields of Issue. Other queries might be better handled by GetXxx or Select methods.
func (tbl IssueTable) Fetch(req require.Requirement, query string, args ...interface{}) ([]*Issue, error) {
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
func (tbl IssueTable) SelectOneWhere(req require.Requirement, where, orderBy string, args ...interface{}) (*Issue, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT 1",
		allIssueQuotedColumnNames[tbl.Dialect().Index()], tbl.name, where, orderBy)
	v, err := tbl.doQueryOne(req, query, args...)
	return v, err
}

// SelectOne allows a single Issue to be obtained from the sqlgen2.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
// If not found, *Example will be nil.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.One
// controls whether an error is generated when no result is found.
func (tbl IssueTable) SelectOne(req require.Requirement, wh where.Expression, qc where.QueryConstraint) (*Issue, error) {
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
	return tbl.SelectOneWhere(req, whs, orderBy, args...)
}

// SelectWhere allows Issues to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
//
// The args are for any placeholder parameters in the query.
func (tbl IssueTable) SelectWhere(req require.Requirement, where, orderBy string, args ...interface{}) ([]*Issue, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s %s",
		allIssueQuotedColumnNames[tbl.Dialect().Index()], tbl.name, where, orderBy)
	vv, err := tbl.doQuery(req, false, query, args...)
	return vv, err
}

// Select allows Issues to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
func (tbl IssueTable) Select(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]*Issue, error) {
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
	return tbl.SelectWhere(req, whs, orderBy, args...)
}

// CountWhere counts Issues in the table that match a 'where' clause.
// Use a blank string for the 'where' argument if it is not needed.
//
// The args are for any placeholder parameters in the query.
func (tbl IssueTable) CountWhere(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s %s", tbl.name, where)
	tbl.logQuery(query, args...)
	row := tbl.db.QueryRowContext(tbl.ctx, query, args...)
	err = row.Scan(&count)
	return count, tbl.logIfError(err)
}

// Count counts the Issues in the table that match a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed.
func (tbl IssueTable) Count(wh where.Expression) (count int64, err error) {
	whs, args := where.BuildExpression(wh, tbl.Dialect())
	return tbl.CountWhere(whs, args...)
}

//--------------------------------------------------------------------------------

// SliceId gets the id column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl IssueTable) SliceId(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return tbl.sliceInt64List(req, tbl.pk, wh, qc)
}

// SliceNumber gets the number column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl IssueTable) SliceNumber(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int, error) {
	return tbl.sliceIntList(req, "number", wh, qc)
}

// SliceTitle gets the title column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl IssueTable) SliceTitle(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.sliceStringList(req, "title", wh, qc)
}

// SliceBigbody gets the bigbody column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl IssueTable) SliceBigbody(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.sliceStringList(req, "bigbody", wh, qc)
}

// SliceAssignee gets the assignee column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl IssueTable) SliceAssignee(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.sliceStringList(req, "assignee", wh, qc)
}

// SliceState gets the state column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl IssueTable) SliceState(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.sliceStringList(req, "state", wh, qc)
}

func (tbl IssueTable) sliceIntList(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]int, error) {
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

	var v int
	list := make([]int, 0, 10)

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

func (tbl IssueTable) sliceInt64List(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
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

func (tbl IssueTable) sliceStringList(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
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


func constructIssueInsert(w io.Writer, v *Issue, dialect schema.Dialect, withPk bool) (s []interface{}, err error) {
	s = make([]interface{}, 0, 8)

	comma := ""
	io.WriteString(w, " (")

	if withPk {
		dialect.QuoteW(w, "id")
		comma = ","
		s = append(s, v.Id)
	}

	io.WriteString(w, comma)

	dialect.QuoteW(w, "number")
	s = append(s, v.Number)
	comma = ","
	io.WriteString(w, comma)

	dialect.QuoteW(w, "date")
	s = append(s, v.Date)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "title")
	s = append(s, v.Title)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "bigbody")
	s = append(s, v.Body)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "assignee")
	s = append(s, v.Assignee)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "state")
	s = append(s, v.State)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "labels")
	x, err := json.Marshal(&v.Labels)
	if err != nil {
		return nil, err
	}
	s = append(s, x)
	io.WriteString(w, ")")
	return s, nil
}

func constructIssueUpdate(w io.Writer, v *Issue, dialect schema.Dialect) (s []interface{}, err error) {
	j := 1
	s = make([]interface{}, 0, 7)

	comma := ""

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "number", j)
	s = append(s, v.Number)
	comma = ", "
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "date", j)
	s = append(s, v.Date)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "title", j)
	s = append(s, v.Title)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "bigbody", j)
	s = append(s, v.Body)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "assignee", j)
	s = append(s, v.Assignee)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "state", j)
	s = append(s, v.State)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "labels", j)
		j++
	x, err := json.Marshal(&v.Labels)
	if err != nil {
		return nil, err
	}
	s = append(s, x)

	return s, nil
}

//--------------------------------------------------------------------------------

// Insert adds new records for the Issues.
// The Issues have their primary key fields set to the new record identifiers.
// The Issue.PreInsert() method will be called, if it exists.
func (tbl IssueTable) Insert(req require.Requirement, vv ...*Issue) error {
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

		fields, err := constructIssueInsert(b, v, tbl.Dialect(), false)
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
func (tbl IssueTable) UpdateFields(req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error) {
	return support.UpdateFields(tbl, req, wh, fields...)
}

//--------------------------------------------------------------------------------

var allIssueQuotedUpdates = []string{
	// Sqlite
	"`number`=?,`date`=?,`title`=?,`bigbody`=?,`assignee`=?,`state`=?,`labels`=? WHERE `id`=?",
	// Mysql
	"`number`=?,`date`=?,`title`=?,`bigbody`=?,`assignee`=?,`state`=?,`labels`=? WHERE `id`=?",
	// Postgres
	`"number"=$2,"date"=$3,"title"=$4,"bigbody"=$5,"assignee"=$6,"state"=$7,"labels"=$8 WHERE "id"=$1`,
}

//--------------------------------------------------------------------------------

// Update updates records, matching them by primary key. It returns the number of rows affected.
// The Issue.PreUpdate(Execer) method will be called, if it exists.
func (tbl IssueTable) Update(req require.Requirement, vv ...*Issue) (int64, error) {
	if req == require.All {
		req = require.Exactly(len(vv))
	}

	var count int64
	dialect := tbl.Dialect()
	//columns := allIssueQuotedUpdates[dialect.Index()]
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

		args, err := constructIssueUpdate(b, v, dialect)
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

// DeleteIssues deletes rows from the table, given some primary keys.
// The list of ids can be arbitrarily long.
func (tbl IssueTable) DeleteIssues(req require.Requirement, id ...int64) (int64, error) {
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
func (tbl IssueTable) Delete(req require.Requirement, wh where.Expression) (int64, error) {
	query, args := tbl.deleteRows(wh)
	return tbl.Exec(req, query, args...)
}

func (tbl IssueTable) deleteRows(wh where.Expression) (string, []interface{}) {
	whs, args := where.BuildExpression(wh, tbl.Dialect())
	query := fmt.Sprintf("DELETE FROM %s %s", tbl.name, whs)
	return query, args
}

//--------------------------------------------------------------------------------

//--------------------------------------------------------------------------------

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
