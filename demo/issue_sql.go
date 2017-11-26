// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

package demo

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/rickb777/sqlgen2"
	"github.com/rickb777/sqlgen2/where"
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
	Dialect      sqlgen2.Dialect
}

// NewIssueTable returns a new table instance.
func NewIssueTable(prefix, name string, d *sql.DB, dialect sqlgen2.Dialect) IssueTable {
	if name == "" {
		name = IssueTableName
	}
	return IssueTable{prefix, name, d, context.Background(), dialect}
}

// WithContext sets the context for subsequent queries.
func (tbl IssueTable) WithContext(ctx context.Context) IssueTable {
	tbl.Ctx = ctx
	return tbl
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


// ScanIssue reads a database record into a single value.
func ScanIssue(row *sql.Row) (*Issue, error) {
	var v0 int64
	var v1 int
	var v2 string
	var v3 string
	var v4 string
	var v5 string
	var v6 []byte

	err := row.Scan(
		&v0,
		&v1,
		&v2,
		&v3,
		&v4,
		&v5,
		&v6,

	)
	if err != nil {
		return nil, err
	}

	v := &Issue{}
	v.Id = v0
	v.Number = v1
	v.Title = v2
	v.Body = v3
	v.Assignee = v4
	v.State = v5
	err = json.Unmarshal(v6, &v.Labels)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// ScanIssues reads database records into a slice of values.
func ScanIssues(rows *sql.Rows) ([]*Issue, error) {
	var err error
	var vv []*Issue

	var v0 int64
	var v1 int
	var v2 string
	var v3 string
	var v4 string
	var v5 string
	var v6 []byte

	for rows.Next() {
		err = rows.Scan(
			&v0,
			&v1,
			&v2,
			&v3,
			&v4,
			&v5,
			&v6,

		)
		if err != nil {
			return vv, err
		}

		v := &Issue{}
		v.Id = v0
		v.Number = v1
		v.Title = v2
		v.Body = v3
		v.Assignee = v4
		v.State = v5
		err = json.Unmarshal(v6, &v.Labels)
		if err != nil {
			return nil, err
		}

		vv = append(vv, v)
	}
	return vv, rows.Err()
}

func SliceIssue(v *Issue) ([]interface{}, error) {
	var v0 int64
	var v1 int
	var v2 string
	var v3 string
	var v4 string
	var v5 string
	var v6 []byte

	v0 = v.Id
	v1 = v.Number
	v2 = v.Title
	v3 = v.Body
	v4 = v.Assignee
	v5 = v.State
	v6, err := json.Marshal(&v.Labels)
	if err != nil {
		return nil, err
	}

	return []interface{}{
		v0,
		v1,
		v2,
		v3,
		v4,
		v5,
		v6,

	}, nil
}

func SliceIssueWithoutPk(v *Issue) ([]interface{}, error) {
	var v1 int
	var v2 string
	var v3 string
	var v4 string
	var v5 string
	var v6 []byte

	v1 = v.Number
	v2 = v.Title
	v3 = v.Body
	v4 = v.Assignee
	v5 = v.State
	v6, err := json.Marshal(&v.Labels)
	if err != nil {
		return nil, err
	}

	return []interface{}{
		v1,
		v2,
		v3,
		v4,
		v5,
		v6,

	}, nil
}

//--------------------------------------------------------------------------------

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// It returns the number of rows affected.
func (tbl IssueTable) Exec(query string, args ...interface{}) (int64, error) {
	res, err := tbl.Db.ExecContext(tbl.Ctx, query, args...)
	if err != nil {
		return 0, nil
	}
	return res.RowsAffected()
}

//--------------------------------------------------------------------------------

// QueryOne is the low-level access function for one Issue.
func (tbl IssueTable) QueryOne(query string, args ...interface{}) (*Issue, error) {
	row := tbl.Db.QueryRowContext(tbl.Ctx, query, args...)
	return ScanIssue(row)
}

// Query is the low-level access function for Issues.
func (tbl IssueTable) Query(query string, args ...interface{}) ([]*Issue, error) {
	rows, err := tbl.Db.QueryContext(tbl.Ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanIssues(rows)
}

//--------------------------------------------------------------------------------

// SelectOneSA allows a single Issue to be obtained from the database that match a 'where' clause and some limit.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
func (tbl IssueTable) SelectOneSA(where, orderBy string, args ...interface{}) (*Issue, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s LIMIT 1", IssueColumnNames, tbl.Prefix, tbl.Name, where, orderBy)
	return tbl.QueryOne(query, args...)
}

// SelectOne allows a single Issue to be obtained from the sqlgen2.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
func (tbl IssueTable) SelectOne(where where.Expression, orderBy string) (*Issue, error) {
	wh, args := where.Build(tbl.Dialect)
	return tbl.SelectOneSA(wh, orderBy, args)
}

// SelectSA allows Issues to be obtained from the database that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
func (tbl IssueTable) SelectSA(where, orderBy string, args ...interface{}) ([]*Issue, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", IssueColumnNames, tbl.Prefix, tbl.Name, where, orderBy)
	return tbl.Query(query, args...)
}

// Select allows Issues to be obtained from the database that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
func (tbl IssueTable) Select(where where.Expression, orderBy string) ([]*Issue, error) {
	wh, args := where.Build(tbl.Dialect)
	return tbl.SelectSA(wh, orderBy, args)
}

// CountSA counts Issues in the database that match a 'where' clause.
func (tbl IssueTable) CountSA(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s%s %s", tbl.Prefix, tbl.Name, where)
	row := tbl.Db.QueryRowContext(tbl.Ctx, query, args)
	err = row.Scan(&count)
	return count, err
}

// Count counts the Issues in the database that match a 'where' clause.
func (tbl IssueTable) Count(where where.Expression) (count int64, err error) {
	return tbl.CountSA(where.Build(tbl.Dialect))
}

const IssueColumnNames = "id, number, title, bigbody, assignee, state, labels"

//--------------------------------------------------------------------------------

// Insert adds new records for the Issues. The Issues have their primary key fields
// set to the new record identifiers.
// The Issue.PreInsert(Execer) method will be called, if it exists.
func (tbl IssueTable) Insert(vv ...*Issue) error {
	var stmt, params string
	switch tbl.Dialect {
	case sqlgen2.Postgres:
		stmt = sqlInsertIssuePostgres
		params = sIssueDataColumnParamsPostgres
	default:
		stmt = sqlInsertIssueSimple
		params = sIssueDataColumnParamsSimple
	}

	st, err := tbl.Db.PrepareContext(tbl.Ctx, fmt.Sprintf(stmt, tbl.Prefix, tbl.Name, params))
	if err != nil {
		return err
	}
	defer st.Close()

	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(interface{PreInsert(sqlgen2.Execer)}); ok {
			hook.PreInsert(tbl.Db)
		}

		fields, err := SliceIssueWithoutPk(v)
		if err != nil {
			return err
		}

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

const sqlInsertIssueSimple = `
INSERT INTO %s%s (
	number,
	title,
	bigbody,
	assignee,
	state,
	labels
) VALUES (%s)
`

const sqlInsertIssuePostgres = `
INSERT INTO %s%s (
	number,
	title,
	bigbody,
	assignee,
	state,
	labels
) VALUES (%s)
`

const sIssueDataColumnParamsSimple = "?,?,?,?,?,?"

const sIssueDataColumnParamsPostgres = "$1,$2,$3,$4,$5,$6"

//--------------------------------------------------------------------------------

// UpdateFields updates one or more columns, given a 'where' clause.
func (tbl IssueTable) UpdateFields(where where.Expression, fields ...sql.NamedArg) (int64, error) {
	return tbl.Exec(tbl.updateFields(where, fields...))
}

func (tbl IssueTable) updateFields(where where.Expression, fields ...sql.NamedArg) (string, []interface{}) {
	list := sqlgen2.NamedArgList(fields)
	assignments := strings.Join(list.Assignments(tbl.Dialect, 1), ", ")
	whereClause, wargs := where.Build(tbl.Dialect)
	query := fmt.Sprintf("UPDATE %s%s SET %s %s", tbl.Prefix, tbl.Name, assignments, whereClause)
	args := append(list.Values(), wargs...)
	return query, args
}

// Update updates records, matching them by primary key. It returns the number of rows affected.
// The Issue.PreUpdate(Execer) method will be called, if it exists.
func (tbl IssueTable) Update(vv ...*Issue) (int64, error) {
	var stmt string
	switch tbl.Dialect {
	case sqlgen2.Postgres:
		stmt = sqlUpdateIssueByPkPostgres
	default:
		stmt = sqlUpdateIssueByPkSimple
	}

	var count int64
	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(interface{PreUpdate(sqlgen2.Execer)}); ok {
			hook.PreUpdate(tbl.Db)
		}

		query := fmt.Sprintf(stmt, tbl.Prefix, tbl.Name)

		args, err := SliceIssueWithoutPk(v)
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
UPDATE %s%s SET 
	number=?,
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
	title=$3,
	bigbody=$4,
	assignee=$5,
	state=$6,
	labels=$7 
 WHERE id=$8
`

//--------------------------------------------------------------------------------

// DeleteFields deleted one or more rows, given a 'where' clause.
func (tbl IssueTable) Delete(where where.Expression) (int64, error) {
	return tbl.Exec(tbl.deleteRows(where))
}

func (tbl IssueTable) deleteRows(where where.Expression) (string, []interface{}) {
	whereClause, args := where.Build(tbl.Dialect)
	query := fmt.Sprintf("DELETE FROM %s%s %s", tbl.Prefix, tbl.Name, whereClause)
	return query, args
}

//--------------------------------------------------------------------------------

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// It returns the number of rows affected.
// Not every database or database driver may support this.
func (tbl IssueTable) CreateTable(ifNotExist bool) (int64, error) {
	return tbl.Exec(tbl.createTableSql(ifNotExist))
}

func (tbl IssueTable) createTableSql(ifNotExist bool) string {
	var stmt string
	switch tbl.Dialect {
	case sqlgen2.Sqlite: stmt = sqlCreateIssueTableSqlite
    case sqlgen2.Postgres: stmt = sqlCreateIssueTablePostgres
    case sqlgen2.Mysql: stmt = sqlCreateIssueTableMysql
    }
	extra := tbl.ternary(ifNotExist, "IF NOT EXISTS ", "")
	query := fmt.Sprintf(stmt, extra, tbl.Prefix, tbl.Name)
	return query
}

func (tbl IssueTable) ternary(flag bool, a, b string) string {
	if flag {
		return a
	}
	return b
}

//--------------------------------------------------------------------------------

// CreateIndexes executes queries that create the indexes needed by the Issue table.
func (tbl IssueTable) CreateIndexes(ifNotExist bool) (err error) {
	extra := tbl.ternary(ifNotExist, "IF NOT EXISTS ", "")
    
	_, err = tbl.Exec(tbl.createIssueAssigneeIndexSql(extra))
	if err != nil {
		return err
	}
    
	return nil
}


func (tbl IssueTable) createIssueAssigneeIndexSql(ifNotExist string) string {
	return fmt.Sprintf(sqlCreateIssueAssigneeIndex, ifNotExist, tbl.Prefix, tbl.Name)
}


//--------------------------------------------------------------------------------

const sqlCreateIssueTableSqlite = `
CREATE TABLE %s%s%s (
 id       integer primary key autoincrement,
 number   integer,
 title    text,
 bigbody  text,
 assignee text,
 state    text,
 labels   blob
)
`

const sqlCreateIssueTablePostgres = `
CREATE TABLE %s%s%s (
 id       bigserial primary key ,
 number   integer,
 title    varchar(512),
 bigbody  varchar(2048),
 assignee varchar(512),
 state    varchar(50),
 labels   byteaa
)
`

const sqlCreateIssueTableMysql = `
CREATE TABLE %s%s%s (
 id       bigint primary key auto_increment,
 number   bigint,
 title    varchar(512),
 bigbody  varchar(2048),
 assignee varchar(512),
 state    varchar(50),
 labels   mediumblob
) ENGINE=InnoDB DEFAULT CHARSET=utf8
`

//--------------------------------------------------------------------------------

const sqlCreateIssueAssigneeIndex = `
CREATE INDEX %sissue_assignee ON %s%s (assignee)
`

//--------------------------------------------------------------------------------

const NumIssueColumns = 7

const NumIssueDataColumns = 6

const IssuePk = "Id"

const IssueDataColumnNames = "number, title, bigbody, assignee, state, labels"

//--------------------------------------------------------------------------------
