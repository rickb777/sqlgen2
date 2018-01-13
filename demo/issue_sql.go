// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

package demo

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/rickb777/sqlgen2"
	"github.com/rickb777/sqlgen2/schema"
	"github.com/rickb777/sqlgen2/where"
	"log"
	"strings"
)

// IssueTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type IssueTable struct {
	name    sqlgen2.TableName
	db      sqlgen2.Execer
	ctx     context.Context
	dialect schema.Dialect
	logger  *log.Logger
}

// Type conformance check
var _ sqlgen2.TableWithIndexes = &IssueTable{}

// NewIssueTable returns a new table instance.
// If a blank table name is supplied, the default name "issues" will be used instead.
// The request context is initialised with the background.
func NewIssueTable(name sqlgen2.TableName, d sqlgen2.Execer, dialect schema.Dialect) IssueTable {
	if name.Name == "" {
		name.Name = "issues"
	}
	return IssueTable{name, d, context.Background(), dialect, nil}
}

// CopyTableAsIssueTable copies a table instance, retaining the name etc but
// providing methods appropriate for 'Issue'.
func CopyTableAsIssueTable(origin sqlgen2.Table) IssueTable {
	return IssueTable{
		name:    origin.Name(),
		db:      origin.DB(),
		ctx:     origin.Ctx(),
		dialect: origin.Dialect(),
		logger:  origin.Logger(),
	}
}

// WithPrefix sets the table name prefix for subsequent queries.
func (tbl IssueTable) WithPrefix(pfx string) IssueTable {
	tbl.name.Prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries.
func (tbl IssueTable) WithContext(ctx context.Context) IssueTable {
	tbl.ctx = ctx
	return tbl
}

// WithLogger sets the logger for subsequent queries.
func (tbl IssueTable) WithLogger(logger *log.Logger) IssueTable {
	tbl.logger = logger
	return tbl
}

// Ctx gets the current request context.
func (tbl IssueTable) Ctx() context.Context {
	return tbl.ctx
}

// Dialect gets the database dialect.
func (tbl IssueTable) Dialect() schema.Dialect {
	return tbl.dialect
}

// Logger gets the trace logger.
func (tbl IssueTable) Logger() *log.Logger {
	return tbl.logger
}

// SetLogger sets the logger for subsequent queries, returning the interface.
func (tbl IssueTable) SetLogger(logger *log.Logger) sqlgen2.Table {
	tbl.logger = logger
	return tbl
}

// Name gets the table name.
func (tbl IssueTable) Name() sqlgen2.TableName {
	return tbl.name
}

// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl IssueTable) DB() *sql.DB {
	return tbl.db.(*sql.DB)
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

// Begin starts a transaction. The default isolation level is dependent on the driver.
func (tbl IssueTable) BeginTx(opts *sql.TxOptions) (IssueTable, error) {
	d := tbl.db.(*sql.DB)
	var err error
	tbl.db, err = d.BeginTx(tbl.ctx, opts)
	return tbl, err
}

func (tbl IssueTable) logQuery(query string, args ...interface{}) {
	sqlgen2.LogQuery(tbl.logger, query, args...)
}


//--------------------------------------------------------------------------------

const NumIssueColumns = 8

const NumIssueDataColumns = 7

const IssuePk = "Id"

const IssueDataColumnNames = "number, date, title, bigbody, assignee, state, labels"

//--------------------------------------------------------------------------------

// CreateTable creates the table.
func (tbl IssueTable) CreateTable(ifNotExists bool) (int64, error) {
	return tbl.Exec(tbl.createTableSql(ifNotExists))
}

func (tbl IssueTable) createTableSql(ifNotExists bool) string {
	var stmt string
	switch tbl.dialect {
	case schema.Sqlite: stmt = sqlCreateIssueTableSqlite
    case schema.Postgres: stmt = sqlCreateIssueTablePostgres
    case schema.Mysql: stmt = sqlCreateIssueTableMysql
    }
	extra := tbl.ternary(ifNotExists, "IF NOT EXISTS ", "")
	query := fmt.Sprintf(stmt, extra, tbl.name)
	return query
}

func (tbl IssueTable) ternary(flag bool, a, b string) string {
	if flag {
		return a
	}
	return b
}

// DropTable drops the table, destroying all its data.
func (tbl IssueTable) DropTable(ifExists bool) (int64, error) {
	return tbl.Exec(tbl.dropTableSql(ifExists))
}

func (tbl IssueTable) dropTableSql(ifExists bool) string {
	extra := tbl.ternary(ifExists, "IF EXISTS ", "")
	query := fmt.Sprintf("DROP TABLE %s%s", extra, tbl.name)
	return query
}

const sqlCreateIssueTableSqlite = `
CREATE TABLE %s%s (
 id       integer primary key autoincrement,
 number   bigint,
 date     blob,
 title    text,
 bigbody  text,
 assignee text,
 state    text,
 labels   text
)
`

const sqlCreateIssueTablePostgres = `
CREATE TABLE %s%s (
 id       bigserial primary key,
 number   bigint,
 date     byteaa,
 title    varchar(512),
 bigbody  varchar(2048),
 assignee varchar(255),
 state    varchar(50),
 labels   json
)
`

const sqlCreateIssueTableMysql = `
CREATE TABLE %s%s (
 id       bigint primary key auto_increment,
 number   bigint,
 date     mediumblob,
 title    varchar(512),
 bigbody  varchar(2048),
 assignee varchar(255),
 state    varchar(50),
 labels   json
) ENGINE=InnoDB DEFAULT CHARSET=utf8
`

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
	ine := tbl.ternary(ifNotExist && tbl.dialect != schema.Mysql, "IF NOT EXISTS ", "")

	// Mysql does not support 'if not exists' on indexes
	// Workaround: use DropIndex first and ignore an error returned if the index didn't exist.

	if ifNotExist && tbl.dialect == schema.Mysql {
		tbl.DropIssueAssigneeIndex(false)
		ine = ""
	}

	_, err := tbl.Exec(tbl.createIssueAssigneeIndexSql(ine))
	return err
}

func (tbl IssueTable) createIssueAssigneeIndexSql(ifNotExists string) string {
	indexPrefix := tbl.name.PrefixWithoutDot()
	return fmt.Sprintf("CREATE INDEX %s%sissue_assignee ON %s (%s)", ifNotExists, indexPrefix,
		tbl.name, sqlIssueAssigneeIndexColumns)
}

// DropIssueAssigneeIndex drops the issue_assignee index.
func (tbl IssueTable) DropIssueAssigneeIndex(ifExists bool) error {
	_, err := tbl.Exec(tbl.dropIssueAssigneeIndexSql(ifExists))
	return err
}

func (tbl IssueTable) dropIssueAssigneeIndexSql(ifExists bool) string {
	// Mysql does not support 'if exists' on indexes
	ie := tbl.ternary(ifExists && tbl.dialect != schema.Mysql, "IF EXISTS ", "")
	onTbl := tbl.ternary(tbl.dialect == schema.Mysql, fmt.Sprintf(" ON %s", tbl.name), "")
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

const sqlIssueAssigneeIndexColumns = "assignee"

//--------------------------------------------------------------------------------

// Truncate drops every record from the table, if possible. It might fail if constraints exist that
// prevent some or all rows from being deleted; use the force option to override this.
//
// When 'force' is set true, be aware of the following consequences.
// When using Mysql, foreign keys in other tables can be left dangling.
// When using Postgres, a cascade happens, so all 'adjacent' tables (i.e. linked by foreign keys)
// are also truncated.
func (tbl IssueTable) Truncate(force bool) (err error) {
	for _, query := range tbl.dialect.TruncateDDL(tbl.Name().String(), force) {
		_, err = tbl.Exec(query)
		if err != nil {
			return err
		}
	}
	return nil
}

//--------------------------------------------------------------------------------

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// It returns the number of rows affected (of the database drive supports this).
func (tbl IssueTable) Exec(query string, args ...interface{}) (int64, error) {
	tbl.logQuery(query, args...)
	res, err := tbl.db.ExecContext(tbl.ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

//--------------------------------------------------------------------------------

// QueryOne is the low-level access function for one Issue.
// If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, *Issue will be nil.
func (tbl IssueTable) QueryOne(query string, args ...interface{}) (*Issue, error) {
	list, err := tbl.doQuery(true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

// Query is the low-level access function for Issues.
func (tbl IssueTable) Query(query string, args ...interface{}) ([]*Issue, error) {
	return tbl.doQuery(false, query, args...)
}

func (tbl IssueTable) doQuery(firstOnly bool, query string, args ...interface{}) ([]*Issue, error) {
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanIssues(rows, firstOnly)
}

//--------------------------------------------------------------------------------

// GetIssue gets the record with a given primary key value.
// If not found, *Issue will be nil.
func (tbl IssueTable) GetIssue(id int64) (*Issue, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE id=?", IssueColumnNames, tbl.name)
	return tbl.QueryOne(query, id)
}

// GetIssues gets records from the table according to a list of primary keys.
// Although the list of ids can be arbitrarily long, there are practical limits;
// note that Oracle DB has a limit of 1000.
func (tbl IssueTable) GetIssues(id ...int64) (list []*Issue, err error) {
	if len(id) > 0 {
		pl := tbl.dialect.Placeholders(len(id))
		query := fmt.Sprintf("SELECT %s FROM %s WHERE id IN (%s)", IssueColumnNames, tbl.name, pl)
		args := make([]interface{}, len(id))

		for i, v := range id {
			args[i] = v
		}

		list, err = tbl.Query(query, args...)
	}

	return list, err
}

//--------------------------------------------------------------------------------

// SelectOneSA allows a single Example to be obtained from the table that match a 'where' clause
// and some limit. Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
// If not found, *Example will be nil.
func (tbl IssueTable) SelectOneSA(where, orderBy string, args ...interface{}) (*Issue, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT 1", IssueColumnNames, tbl.name, where, orderBy)
	return tbl.QueryOne(query, args...)
}

// SelectOne allows a single Issue to be obtained from the sqlgen2.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
// If not found, *Example will be nil.
func (tbl IssueTable) SelectOne(wh where.Expression, qc where.QueryConstraint) (*Issue, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	return tbl.SelectOneSA(whs, orderBy, args...)
}

// SelectSA allows Issues to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
func (tbl IssueTable) SelectSA(where, orderBy string, args ...interface{}) ([]*Issue, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", IssueColumnNames, tbl.name, where, orderBy)
	return tbl.Query(query, args...)
}

// Select allows Issues to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl IssueTable) Select(wh where.Expression, qc where.QueryConstraint) ([]*Issue, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	return tbl.SelectSA(whs, orderBy, args...)
}

// CountSA counts Issues in the table that match a 'where' clause.
// Use a blank string for the 'where' argument if it is not needed.
func (tbl IssueTable) CountSA(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s %s", tbl.name, where)
	tbl.logQuery(query, args...)
	row := tbl.db.QueryRowContext(tbl.ctx, query, args...)
	err = row.Scan(&count)
	return count, err
}

// Count counts the Issues in the table that match a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed.
func (tbl IssueTable) Count(wh where.Expression) (count int64, err error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	return tbl.CountSA(whs, args...)
}

const IssueColumnNames = "id, number, date, title, bigbody, assignee, state, labels"

//--------------------------------------------------------------------------------

// SliceId gets the Id column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl IssueTable) SliceId(wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return tbl.getint64list("id", wh, qc)
}

// SliceNumber gets the Number column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl IssueTable) SliceNumber(wh where.Expression, qc where.QueryConstraint) ([]int, error) {
	return tbl.getintlist("number", wh, qc)
}

// SliceDate gets the Date column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl IssueTable) SliceDate(wh where.Expression, qc where.QueryConstraint) ([]Date, error) {
	return tbl.getDatelist("date", wh, qc)
}

// SliceTitle gets the Title column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl IssueTable) SliceTitle(wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.getstringlist("title", wh, qc)
}

// SliceBody gets the Body column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl IssueTable) SliceBigbody(wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.getstringlist("bigbody", wh, qc)
}

// SliceAssignee gets the Assignee column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl IssueTable) SliceAssignee(wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.getstringlist("assignee", wh, qc)
}

// SliceState gets the State column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl IssueTable) SliceState(wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.getstringlist("state", wh, qc)
}


func (tbl IssueTable) getDatelist(sqlname string, wh where.Expression, qc where.QueryConstraint) ([]Date, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", sqlname, tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var v Date
	list := make([]Date, 0, 10)
	for rows.Next() {
		err = rows.Scan(&v)
		if err != nil {
			return list, err
		}
		list = append(list, v)
	}
	return list, nil
}

func (tbl IssueTable) getintlist(sqlname string, wh where.Expression, qc where.QueryConstraint) ([]int, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", sqlname, tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var v int
	list := make([]int, 0, 10)
	for rows.Next() {
		err = rows.Scan(&v)
		if err != nil {
			return list, err
		}
		list = append(list, v)
	}
	return list, nil
}

func (tbl IssueTable) getint64list(sqlname string, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", sqlname, tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var v int64
	list := make([]int64, 0, 10)
	for rows.Next() {
		err = rows.Scan(&v)
		if err != nil {
			return list, err
		}
		list = append(list, v)
	}
	return list, nil
}

func (tbl IssueTable) getstringlist(sqlname string, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", sqlname, tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var v string
	list := make([]string, 0, 10)
	for rows.Next() {
		err = rows.Scan(&v)
		if err != nil {
			return list, err
		}
		list = append(list, v)
	}
	return list, nil
}


//--------------------------------------------------------------------------------

// Insert adds new records for the Issues. The Issues have their primary key fields
// set to the new record identifiers.
// The Issue.PreInsert(Execer) method will be called, if it exists.
func (tbl IssueTable) Insert(vv ...*Issue) error {
	var params string
	switch tbl.dialect {
	case schema.Postgres:
		params = sIssueDataColumnParamsPostgres
	default:
		params = sIssueDataColumnParamsSimple
	}

	query := fmt.Sprintf(sqlInsertIssue, tbl.name, params)
	st, err := tbl.db.PrepareContext(tbl.ctx, query)
	if err != nil {
		return err
	}
	defer st.Close()

	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPreInsert); ok {
			hook.PreInsert()
		}

		fields, err := sliceIssueWithoutPk(v)
		if err != nil {
			return err
		}

		tbl.logQuery(query, fields...)
		res, err := st.Exec(fields...)
		if err != nil {
			return err
		}

		v.Id, err = res.LastInsertId()
		if err != nil {
			return err
		}
	}

	return nil
}

const sqlInsertIssue = `
INSERT INTO %s (
	number,
	date,
	title,
	bigbody,
	assignee,
	state,
	labels
) VALUES (%s)
`

const sIssueDataColumnParamsSimple = "?,?,?,?,?,?,?"

const sIssueDataColumnParamsPostgres = "$1,$2,$3,$4,$5,$6,$7"

//--------------------------------------------------------------------------------

// UpdateFields updates one or more columns, given a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl IssueTable) UpdateFields(wh where.Expression, fields ...sql.NamedArg) (int64, error) {
	query, args := tbl.updateFields(wh, fields...)
	return tbl.Exec(query, args...)
}

func (tbl IssueTable) updateFields(wh where.Expression, fields ...sql.NamedArg) (string, []interface{}) {
	list := sqlgen2.NamedArgList(fields)
	assignments := strings.Join(list.Assignments(tbl.dialect, 1), ", ")
	whs, wargs := where.BuildExpression(wh, tbl.dialect)
	query := fmt.Sprintf("UPDATE %s SET %s %s", tbl.name, assignments, whs)
	args := append(list.Values(), wargs...)
	return query, args
}

//--------------------------------------------------------------------------------

// Update updates records, matching them by primary key. It returns the number of rows affected.
// The Issue.PreUpdate(Execer) method will be called, if it exists.
func (tbl IssueTable) Update(vv ...*Issue) (int64, error) {
	var stmt string
	switch tbl.dialect {
	case schema.Postgres:
		stmt = sqlUpdateIssueByPkPostgres
	default:
		stmt = sqlUpdateIssueByPkSimple
	}

	query := fmt.Sprintf(stmt, tbl.name)

	var count int64
	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPreUpdate); ok {
			hook.PreUpdate()
		}

		args, err := sliceIssueWithoutPk(v)
		if err != nil {
			return count, err
		}

		args = append(args, v.Id)
		n, err := tbl.Exec(query, args...)
		if err != nil {
			return count, err
		}
		count += n
	}
	return count, nil
}

const sqlUpdateIssueByPkSimple = `
UPDATE %s SET
	number=?,
	date=?,
	title=?,
	bigbody=?,
	assignee=?,
	state=?,
	labels=?
WHERE id=?
`

const sqlUpdateIssueByPkPostgres = `
UPDATE %s SET
	number=$2,
	date=$3,
	title=$4,
	bigbody=$5,
	assignee=$6,
	state=$7,
	labels=$8
WHERE id=$1
`

func sliceIssueWithoutPk(v *Issue) ([]interface{}, error) {

	v7, err := json.Marshal(&v.Labels)
	if err != nil {
		return nil, err
	}

	return []interface{}{
		v.Number,
		v.Date,
		v.Title,
		v.Body,
		v.Assignee,
		v.State,
		v7,

	}, nil
}

//--------------------------------------------------------------------------------

// DeleteIssues deletes rows from the table, given some primary keys.
// The list of ids can be arbitrarily long.
func (tbl IssueTable) DeleteIssues(id ...int64) (int64, error) {
	const batch = 1000 // limited by Oracle DB
	const qt = "DELETE FROM %s WHERE id IN (%s)"

	var count, n int64
	var err error
	var max = batch
	if len(id) < batch {
		max = len(id)
	}
	args := make([]interface{}, max)

	if len(id) > batch {
		pl := tbl.dialect.Placeholders(batch)
		query := fmt.Sprintf(qt, tbl.name, pl)

		for len(id) > batch {
			for i := 0; i < batch; i++ {
				args[i] = id[i]
			}

			n, err = tbl.Exec(query, args...)
			count += n
			if err != nil {
				return count, err
			}

			id = id[batch:]
		}
	}

	if len(id) > 0 {
		pl := tbl.dialect.Placeholders(len(id))
		query := fmt.Sprintf(qt, tbl.name, pl)

		for i := 0; i < len(id); i++ {
			args[i] = id[i]
		}

		n, err = tbl.Exec(query, args...)
		count += n
	}

	return count, err
}

// Delete deletes one or more rows from the table, given a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl IssueTable) Delete(wh where.Expression) (int64, error) {
	query, args := tbl.deleteRows(wh)
	return tbl.Exec(query, args...)
}

func (tbl IssueTable) deleteRows(wh where.Expression) (string, []interface{}) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	query := fmt.Sprintf("DELETE FROM %s %s", tbl.name, whs)
	return query, args
}

//--------------------------------------------------------------------------------

// scanIssues reads table records into a slice of values.
func scanIssues(rows *sql.Rows, firstOnly bool) ([]*Issue, error) {
	var err error
	var vv []*Issue

	for rows.Next() {
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
			return vv, err
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
			return nil, err
		}

		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPostGet); ok {
			err = hook.PostGet()
			if err != nil {
				return vv, err
			}
		}

		vv = append(vv, v)

		if firstOnly {
			return vv, rows.Err()
		}
	}

	return vv, rows.Err()
}

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
