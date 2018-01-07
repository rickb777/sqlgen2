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

// IssueTableName is the default name for this table.
const IssueTableName = "issues"

// IssueTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type IssueTable struct {
	Prefix, Name string
	Db           sqlgen2.Execer
	Ctx          context.Context
	Dialect      schema.Dialect
	Logger       *log.Logger
}

// Type conformance check
var _ sqlgen2.Table = &IssueTable{}

// NewIssueTable returns a new table instance.
// If a blank table name is supplied, the default name "issues" will be used instead.
// The table name prefix is initially blank and the request context is the background.
func NewIssueTable(name string, d sqlgen2.Execer, dialect schema.Dialect) IssueTable {
	if name == "" {
		name = IssueTableName
	}
	return IssueTable{"", name, d, context.Background(), dialect, nil}
}

// WithPrefix sets the prefix for subsequent queries.
func (tbl IssueTable) WithPrefix(pfx string) IssueTable {
	tbl.Prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries.
func (tbl IssueTable) WithContext(ctx context.Context) IssueTable {
	tbl.Ctx = ctx
	return tbl
}

// WithLogger sets the logger for subsequent queries.
func (tbl IssueTable) WithLogger(logger *log.Logger) IssueTable {
	tbl.Logger = logger
	return tbl
}

// SetLogger sets the logger for subsequent queries, returning the interface.
func (tbl IssueTable) SetLogger(logger *log.Logger) sqlgen2.Table {
	tbl.Logger = logger
	return tbl
}

// FullName gets the concatenated prefix and table name.
func (tbl IssueTable) FullName() string {
	return tbl.Prefix + tbl.Name
}

func (tbl IssueTable) prefixWithoutDot() string {
	last := len(tbl.Prefix)-1
	if last > 0 && tbl.Prefix[last] == '.' {
		return tbl.Prefix[0:last]
	}
	return tbl.Prefix
}

// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl IssueTable) DB() *sql.DB {
	return tbl.Db.(*sql.DB)
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl IssueTable) Tx() *sql.Tx {
	return tbl.Db.(*sql.Tx)
}

// IsTx tests whether this is within a transaction.
func (tbl IssueTable) IsTx() bool {
	_, ok := tbl.Db.(*sql.Tx)
	return ok
}

// Begin starts a transaction. The default isolation level is dependent on the driver.
func (tbl IssueTable) BeginTx(opts *sql.TxOptions) (IssueTable, error) {
	d := tbl.Db.(*sql.DB)
	var err error
	tbl.Db, err = d.BeginTx(tbl.Ctx, opts)
	return tbl, err
}

func (tbl IssueTable) logQuery(query string, args ...interface{}) {
	sqlgen2.LogQuery(tbl.Logger, query, args...)
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
	switch tbl.Dialect {
	case schema.Sqlite: stmt = sqlCreateIssueTableSqlite
    case schema.Postgres: stmt = sqlCreateIssueTablePostgres
    case schema.Mysql: stmt = sqlCreateIssueTableMysql
    }
	extra := tbl.ternary(ifNotExists, "IF NOT EXISTS ", "")
	query := fmt.Sprintf(stmt, extra, tbl.Prefix, tbl.Name)
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
	query := fmt.Sprintf("DROP TABLE %s%s%s", extra, tbl.Prefix, tbl.Name)
	return query
}

const sqlCreateIssueTableSqlite = `
CREATE TABLE %s%s%s (
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
CREATE TABLE %s%s%s (
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
CREATE TABLE %s%s%s (
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
	ine := tbl.ternary(ifNotExist && tbl.Dialect != schema.Mysql, "IF NOT EXISTS ", "")

	// Mysql does not support 'if not exists' on indexes
	// Workaround: use DropIndex first and ignore an error returned if the index didn't exist.

	if ifNotExist && tbl.Dialect == schema.Mysql {
		tbl.DropIssueAssigneeIndex(false)
		ine = ""
	}

	_, err := tbl.Exec(tbl.createIssueAssigneeIndexSql(ine))
	return err
}

func (tbl IssueTable) createIssueAssigneeIndexSql(ifNotExists string) string {
	indexPrefix := tbl.prefixWithoutDot()
	return fmt.Sprintf("CREATE INDEX %s%sissue_assignee ON %s%s (%s)", ifNotExists, indexPrefix,
		tbl.Prefix, tbl.Name, sqlIssueAssigneeIndexColumns)
}

// DropIssueAssigneeIndex drops the issue_assignee index.
func (tbl IssueTable) DropIssueAssigneeIndex(ifExists bool) error {
	_, err := tbl.Exec(tbl.dropIssueAssigneeIndexSql(ifExists))
	return err
}

func (tbl IssueTable) dropIssueAssigneeIndexSql(ifExists bool) string {
	// Mysql does not support 'if exists' on indexes
	ie := tbl.ternary(ifExists && tbl.Dialect != schema.Mysql, "IF EXISTS ", "")
	onTbl := tbl.ternary(tbl.Dialect == schema.Mysql, fmt.Sprintf(" ON %s%s", tbl.Prefix, tbl.Name), "")
	indexPrefix := tbl.prefixWithoutDot()
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

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// It returns the number of rows affected (of the database drive supports this).
func (tbl IssueTable) Exec(query string, args ...interface{}) (int64, error) {
	tbl.logQuery(query, args...)
	res, err := tbl.Db.ExecContext(tbl.Ctx, query, args...)
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
	rows, err := tbl.Db.QueryContext(tbl.Ctx, query, args...)
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
	query := fmt.Sprintf("SELECT %s FROM %s%s WHERE id=?", IssueColumnNames, tbl.Prefix, tbl.Name)
	return tbl.QueryOne(query, id)
}

//--------------------------------------------------------------------------------

// SliceId gets the Id column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
func (tbl IssueTable) SliceId(where where.Expression, orderBy string) ([]int64, error) {
	return tbl.getint64list("id", where, orderBy)
}

// SliceNumber gets the Number column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
func (tbl IssueTable) SliceNumber(where where.Expression, orderBy string) ([]int, error) {
	return tbl.getintlist("number", where, orderBy)
}

// SliceDate gets the Date column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
func (tbl IssueTable) SliceDate(where where.Expression, orderBy string) ([]Date, error) {
	return tbl.getDatelist("date", where, orderBy)
}

// SliceTitle gets the Title column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
func (tbl IssueTable) SliceTitle(where where.Expression, orderBy string) ([]string, error) {
	return tbl.getstringlist("title", where, orderBy)
}

// SliceBody gets the Body column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
func (tbl IssueTable) SliceBody(where where.Expression, orderBy string) ([]string, error) {
	return tbl.getstringlist("bigbody", where, orderBy)
}

// SliceAssignee gets the Assignee column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
func (tbl IssueTable) SliceAssignee(where where.Expression, orderBy string) ([]string, error) {
	return tbl.getstringlist("assignee", where, orderBy)
}

// SliceState gets the State column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
func (tbl IssueTable) SliceState(where where.Expression, orderBy string) ([]string, error) {
	return tbl.getstringlist("state", where, orderBy)
}


func (tbl IssueTable) getDatelist(sqlname string, where where.Expression, orderBy string) ([]Date, error) {
	wh, args := where.Build(tbl.Dialect)
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", sqlname, tbl.Prefix, tbl.Name, wh, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.Db.QueryContext(tbl.Ctx, query, args...)
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

func (tbl IssueTable) getintlist(sqlname string, where where.Expression, orderBy string) ([]int, error) {
	wh, args := where.Build(tbl.Dialect)
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", sqlname, tbl.Prefix, tbl.Name, wh, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.Db.QueryContext(tbl.Ctx, query, args...)
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

func (tbl IssueTable) getint64list(sqlname string, where where.Expression, orderBy string) ([]int64, error) {
	wh, args := where.Build(tbl.Dialect)
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", sqlname, tbl.Prefix, tbl.Name, wh, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.Db.QueryContext(tbl.Ctx, query, args...)
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

func (tbl IssueTable) getstringlist(sqlname string, where where.Expression, orderBy string) ([]string, error) {
	wh, args := where.Build(tbl.Dialect)
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", sqlname, tbl.Prefix, tbl.Name, wh, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.Db.QueryContext(tbl.Ctx, query, args...)
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

// SelectOneSA allows a single Issue to be obtained from the table that match a 'where' clause
// and some limit.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
// If not found, *Issue will be nil.
func (tbl IssueTable) SelectOneSA(where, orderBy string, args ...interface{}) (*Issue, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s LIMIT 1", IssueColumnNames, tbl.Prefix, tbl.Name, where, orderBy)
	return tbl.QueryOne(query, args...)
}

// SelectOne allows a single Issue to be obtained from the sqlgen2.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
// If not found, *Example will be nil.
func (tbl IssueTable) SelectOne(where where.Expression, orderBy string) (*Issue, error) {
	wh, args := where.Build(tbl.Dialect)
	return tbl.SelectOneSA(wh, orderBy, args...)
}

// SelectSA allows Issues to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
func (tbl IssueTable) SelectSA(where, orderBy string, args ...interface{}) ([]*Issue, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", IssueColumnNames, tbl.Prefix, tbl.Name, where, orderBy)
	return tbl.Query(query, args...)
}

// Select allows Issues to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
func (tbl IssueTable) Select(where where.Expression, orderBy string) ([]*Issue, error) {
	wh, args := where.Build(tbl.Dialect)
	return tbl.SelectSA(wh, orderBy, args...)
}

// CountSA counts Issues in the table that match a 'where' clause.
func (tbl IssueTable) CountSA(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s%s %s", tbl.Prefix, tbl.Name, where)
	tbl.logQuery(query, args...)
	row := tbl.Db.QueryRowContext(tbl.Ctx, query, args...)
	err = row.Scan(&count)
	return count, err
}

// Count counts the Issues in the table that match a 'where' clause.
func (tbl IssueTable) Count(where where.Expression) (count int64, err error) {
	wh, args := where.Build(tbl.Dialect)
	return tbl.CountSA(wh, args...)
}

const IssueColumnNames = "id, number, date, title, bigbody, assignee, state, labels"

//--------------------------------------------------------------------------------

// Insert adds new records for the Issues. The Issues have their primary key fields
// set to the new record identifiers.
// The Issue.PreInsert(Execer) method will be called, if it exists.
func (tbl IssueTable) Insert(vv ...*Issue) error {
	var params string
	switch tbl.Dialect {
	case schema.Postgres:
		params = sIssueDataColumnParamsPostgres
	default:
		params = sIssueDataColumnParamsSimple
	}

	query := fmt.Sprintf(sqlInsertIssue, tbl.Prefix, tbl.Name, params)
	st, err := tbl.Db.PrepareContext(tbl.Ctx, query)
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
INSERT INTO %s%s (
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
func (tbl IssueTable) UpdateFields(where where.Expression, fields ...sql.NamedArg) (int64, error) {
	query, args := tbl.updateFields(where, fields...)
	return tbl.Exec(query, args...)
}

func (tbl IssueTable) updateFields(where where.Expression, fields ...sql.NamedArg) (string, []interface{}) {
	list := sqlgen2.NamedArgList(fields)
	assignments := strings.Join(list.Assignments(tbl.Dialect, 1), ", ")
	whereClause, wargs := where.Build(tbl.Dialect)
	query := fmt.Sprintf("UPDATE %s%s SET %s %s", tbl.Prefix, tbl.Name, assignments, whereClause)
	args := append(list.Values(), wargs...)
	return query, args
}

//--------------------------------------------------------------------------------

// Update updates records, matching them by primary key. It returns the number of rows affected.
// The Issue.PreUpdate(Execer) method will be called, if it exists.
func (tbl IssueTable) Update(vv ...*Issue) (int64, error) {
	var stmt string
	switch tbl.Dialect {
	case schema.Postgres:
		stmt = sqlUpdateIssueByPkPostgres
	default:
		stmt = sqlUpdateIssueByPkSimple
	}

	query := fmt.Sprintf(stmt, tbl.Prefix, tbl.Name)

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
		tbl.logQuery(query, args...)
		n, err := tbl.Exec(query, args...)
		if err != nil {
			return count, err
		}
		count += n
	}
	return count, nil
}

const sqlUpdateIssueByPkSimple = `
UPDATE %s%s SET
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
UPDATE %s%s SET
	number=$2,
	date=$3,
	title=$4,
	bigbody=$5,
	assignee=$6,
	state=$7,
	labels=$8
WHERE id=$1
`

//--------------------------------------------------------------------------------

// Delete deletes one or more rows from the table, given a 'where' clause.
func (tbl IssueTable) Delete(where where.Expression) (int64, error) {
	query, args := tbl.deleteRows(where)
	return tbl.Exec(query, args...)
}

func (tbl IssueTable) deleteRows(where where.Expression) (string, []interface{}) {
	whereClause, args := where.Build(tbl.Dialect)
	query := fmt.Sprintf("DELETE FROM %s%s %s", tbl.Prefix, tbl.Name, whereClause)
	return query, args
}

// Truncate drops every record from the table, if possible. It might fail if constraints exist that
// prevent some or all rows from being deleted; use the force option to override this.
//
// When 'force' is set true, be aware of the following consequences.
// When using Mysql, foreign keys in other tables can be left dangling.
// When using Postgres, a cascade happens, so all 'adjacent' tables (i.e. linked by foreign keys)
// are also truncated.
func (tbl IssueTable) Truncate(force bool) (err error) {
	for _, query := range tbl.Dialect.TruncateDDL(tbl.FullName(), force) {
		_, err = tbl.Exec(query)
		if err != nil {
			return err
		}
	}
	return nil
}

//--------------------------------------------------------------------------------

// scanIssues reads table records into a slice of values.
func scanIssues(rows *sql.Rows, firstOnly bool) ([]*Issue, error) {
	var err error
	var vv []*Issue

	var v0 int64
	var v1 int
	var v2 Date
	var v3 string
	var v4 string
	var v5 string
	var v6 string
	var v7 []byte

	for rows.Next() {
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
