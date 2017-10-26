// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

package demo

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/rickb777/sqlgen2/dialect"
	"github.com/rickb777/sqlgen2/where"
)

// IssueTableName is the default name for this table.
const IssueTableName = "issues"

// IssueTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type IssueTable struct {
	Prefix, Name string
	Db           *sql.DB
	Dialect      dialect.Dialect
}

// NewIssueTable returns a new table instance.
func NewIssueTable(prefix, name string, db *sql.DB, dialect dialect.Dialect) IssueTable {
	if name == "" {
		name = IssueTableName
	}
	return IssueTable{prefix, name, db, dialect}
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
	json.Unmarshal(v6, &v.Labels)

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
		json.Unmarshal(v6, &v.Labels)

		vv = append(vv, v)
	}
	return vv, rows.Err()
}

func SliceIssue(v *Issue) []interface{} {
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
	v6, _ = json.Marshal(&v.Labels)

	return []interface{}{
		v0,
		v1,
		v2,
		v3,
		v4,
		v5,
		v6,

	}
}

func SliceIssueWithoutPk(v *Issue) []interface{} {
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
	v6, _ = json.Marshal(&v.Labels)

	return []interface{}{
		v1,
		v2,
		v3,
		v4,
		v5,
		v6,

	}
}

// QueryOne is the low-level access function for one Issue.
func (tbl IssueTable) QueryOne(query string, args ...interface{}) (*Issue, error) {
	row := tbl.Db.QueryRow(query, args...)
	return ScanIssue(row)
}

// SelectOneSA allows a single Issue to be obtained from the database using supplied dialect-specific parameters.
func (tbl IssueTable) SelectOneSA(where, limitClause string, args ...interface{}) (*Issue, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", IssueColumnNames, tbl.Prefix, tbl.Name, where, limitClause)
	return tbl.QueryOne(query, args...)
}

// SelectOne allows a single Issue to be obtained from the database.
func (tbl IssueTable) SelectOne(where where.Expression, dialect dialect.Dialect) (*Issue, error) {
	wh, args := where.Build(dialect)
	return tbl.SelectOneSA(wh, "LIMIT 1", args)
}

func (tbl IssueTable) Query(query string, args ...interface{}) ([]*Issue, error) {
	rows, err := tbl.Db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanIssues(rows)
}

// SelectSA allows Issues to be obtained from the database using supplied dialect-specific parameters.
func (tbl IssueTable) SelectSA(where string, args ...interface{}) ([]*Issue, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s", IssueColumnNames, tbl.Prefix, tbl.Name, where)
	return tbl.Query(query, args...)
}

// Select allows Issues to be obtained from the database that match a 'where' clause.
func (tbl IssueTable) Select(where where.Expression, dialect dialect.Dialect) ([]*Issue, error) {
	return tbl.SelectSA(where.Build(dialect))
}

// CountSA counts Issues in the database using supplied dialect-specific parameters.
func (tbl IssueTable) CountSA(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s%s %s", tbl.Prefix, tbl.Name, where)
	row := tbl.Db.QueryRow(query, args)
	err = row.Scan(&count)
	return count, err
}

// Count counts the Issues in the database that match a 'where' clause.
func (tbl IssueTable) Count(where where.Expression, dialect dialect.Dialect) (count int64, err error) {
	return tbl.CountSA(where.Build(dialect))
}

const IssueColumnNames = "id, number, title, assignee, state, labels"

// Insert adds new records for the Issues. The Issues have their primary key fields
// set to the new record identifiers.
func (tbl IssueTable) Insert(vv ...*Issue) error {
	var stmt, params string
	switch tbl.Dialect {
	case dialect.Postgres:
		stmt = sqlInsertIssuePostgres
		params = sIssueDataColumnParamsPostgres
	default:
		stmt = sqlInsertIssueSimple
		params = sIssueDataColumnParamsSimple
	}
	st, err := tbl.Db.Prepare(fmt.Sprintf(stmt, tbl.Prefix, tbl.Name, params))
	if err != nil {
		return err
	}
	defer st.Close()

	for _, v := range vv {
		res, err := st.Exec(SliceIssueWithoutPk(v)...)
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
	assignee,
	state,
	labels
) VALUES (%s)
`

const sqlInsertIssuePostgres = `
INSERT INTO %s%s (
	number,
	title,
	assignee,
	state,
	labels
) VALUES (%s)
`

const sIssueDataColumnParamsSimple = "?,?,?,?,?"

const sIssueDataColumnParamsPostgres = "$1,$2,$3,$4,$5"

// Update updates a record. It returns the number of rows affected.
// Not every database or database driver may support this.
func (tbl IssueTable) Update(v *Issue) (int64, error) {
	var stmt string
	switch tbl.Dialect {
	case dialect.Postgres:
		stmt = sqlUpdateIssueByPkPostgres
	default:
		stmt = sqlUpdateIssueByPkSimple
	}
	query := fmt.Sprintf(stmt, tbl.Prefix, tbl.Name)
	args := SliceIssueWithoutPk(v)
	args = append(args, v.Id)
	return tbl.Exec(query, args...)
}

const sqlUpdateIssueByPkSimple = `
UPDATE %s%s SET 
	number=?,
	title=?,
	assignee=?,
	state=?,
	labels=? 
 WHERE id=?
`

const sqlUpdateIssueByPkPostgres = `
UPDATE %s%s SET 
	number=$2,
	title=$3,
	assignee=$4,
	state=$5,
	labels=$6 
 WHERE id=$7
`

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// It returns the number of rows affected.
// Not every database or database driver may support this.
func (tbl IssueTable) Exec(query string, args ...interface{}) (int64, error) {
	res, err := tbl.Db.Exec(query, args...)
	if err != nil {
		return 0, nil
	}
	return res.RowsAffected()
}

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
	case dialect.Sqlite: stmt = sqlCreateIssueTableSqlite
    case dialect.Postgres: stmt = sqlCreateIssueTablePostgres
    case dialect.Mysql: stmt = sqlCreateIssueTableMysql
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
	var stmt string
	switch tbl.Dialect {
	case dialect.Sqlite: stmt = sqlCreateIssueAssigneeIndexSqlite
    case dialect.Postgres: stmt = sqlCreateIssueAssigneeIndexPostgres
    case dialect.Mysql: stmt = sqlCreateIssueAssigneeIndexMysql
    }
	return fmt.Sprintf(stmt, ifNotExist, tbl.Prefix, tbl.Name)
}



//--------------------------------------------------------------------------------

const sqlCreateIssueTableSqlite = `
CREATE TABLE %s%s%s (
 id       integer primary key autoincrement,
 number   integer,
 title    text,
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
 assignee varchar(512),
 state    varchar(50),
 labels   mediumblob
) ENGINE=InnoDB DEFAULT CHARSET=utf8
`

//--------------------------------------------------------------------------------

const sqlDeleteIssueByPkPostgres = `
DELETE FROM %s%s
 WHERE id=$1
`

const sqlDeleteIssueByPkSimple = `
DELETE FROM %s%s
 WHERE id=?
`

//--------------------------------------------------------------------------------

const sqlCreateIssueAssigneeIndexSqlite = `
CREATE INDEX %sissue_assignee ON %s%s (assignee)
`

//--------------------------------------------------------------------------------

const sqlCreateIssueAssigneeIndexPostgres = `
CREATE INDEX %sissue_assignee ON %s%s (assignee)
`

//--------------------------------------------------------------------------------

const sqlCreateIssueAssigneeIndexMysql = `
CREATE INDEX %sissue_assignee ON %s%s (assignee)
`

//--------------------------------------------------------------------------------

const NumIssueColumns = 6

const NumIssueDataColumns = 5

const IssuePk = "Id"

const IssueDataColumnNames = "number, title, assignee, state, labels"

//--------------------------------------------------------------------------------
