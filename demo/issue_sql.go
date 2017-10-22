// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

package demo

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/rickb777/sqlgen/dialect"
	"github.com/rickb777/sqlgen/where"
)

// IssueTableName is the default name for this table.
const IssueTableName = "issues"

// IssueTable holds a given table name with the database reference, providing access methods below.
type IssueTable struct {
	Name    string
	Db      *sql.DB
	Dialect dialect.Dialect
}

// NewIssueTable returns a new table instance.
func NewIssueTable(name string, db *sql.DB, dialect dialect.Dialect) IssueTable {
	return IssueTable{name, db, dialect}
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
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", sIssueColumnNames, tbl.Name, where, limitClause)
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
	query := fmt.Sprintf("SELECT %s FROM %s %s", sIssueColumnNames, tbl.Name, where)
	return tbl.Query(query, args...)
}

// Select allows Issues to be obtained from the database that match a 'where' clause.
func (tbl IssueTable) Select(where where.Expression, dialect dialect.Dialect) ([]*Issue, error) {
	return tbl.SelectSA(where.Build(dialect))
}

// CountSA counts Issues in the database using supplied dialect-specific parameters.
func (tbl IssueTable) CountSA(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s %s", tbl.Name, where)
	row := tbl.Db.QueryRow(query, args)
	err = row.Scan(&count)
	return count, err
}

// Count counts the Issues in the database that match a 'where' clause.
func (tbl IssueTable) Count(where where.Expression, dialect dialect.Dialect) (count int64, err error) {
	return tbl.CountSA(where.Build(dialect))
}

// Insert adds new records for the Issues.
func (tbl IssueTable) Insert(v *Issue) error {
	query := fmt.Sprintf(sInsertIssueStmt, tbl.Name)
	res, err := tbl.Db.Exec(query, SliceIssueWithoutPk(v)...)
	if err != nil {
		return err
	}

	v.Id, err = res.LastInsertId()
	return err
}

// Update updates a record. It returns the number of rows affected.
// Not every database or database driver may support this.
func (tbl IssueTable) Update(v *Issue) (int64, error) {
	query := fmt.Sprintf(sUpdateIssueByPkStmt, tbl.Name)
	args := SliceIssueWithoutPk(v)
	args = append(args, v.Id)
	return tbl.Exec(query, args...)
}

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
func (tbl IssueTable) CreateTable() (int64, error) {
//"CREATE TABLE IF NOT EXISTS %s ("
// id       INTEGER PRIMARY KEY AUTOINCREMENT,
// number   INTEGER,
// title    TEXT,
// assignee TEXT,
// state    TEXT,
// labels   BLOB
//")"
	return 0, nil
}

//--------------------------------------------------------------------------------

const NumIssueColumns = 6

const sIssueColumnNames = `
id, number, title, assignee, state, labels
`

const sIssueDataColumnNames = `
number, title, assignee, state, labels
`

const sIssueColumnParams = `
?,?,?,?,?,?
`

const sIssueDataColumnParams = `
?,?,?,?,?
`

const sCreateIssueStmt = `
CREATE TABLE IF NOT EXISTS %s (
 id       INTEGER PRIMARY KEY AUTOINCREMENT,
 number   INTEGER,
 title    TEXT,
 assignee TEXT,
 state    TEXT,
 labels   BLOB
);
`

func CreateIssueStmt(tableName string, d dialect.Dialect) string {
	return d.ReplacePlaceholders(fmt.Sprintf(sCreateIssueStmt, tableName))
}

const sInsertIssueStmt = `
INSERT INTO %s (
 number,
 title,
 assignee,
 state,
 labels
) VALUES (?,?,?,?,?)
`

func InsertIssueStmt(tableName string, d dialect.Dialect) string {
	return d.ReplacePlaceholders(fmt.Sprintf(sInsertIssueStmt, tableName))
}

const sUpdateIssueByPkStmt = `
UPDATE %s SET 
 number=?,
 title=?,
 assignee=?,
 state=?,
 labels=? 
 WHERE id=?
`

func UpdateIssueByPkStmt(tableName string, d dialect.Dialect) string {
	return d.ReplacePlaceholders(fmt.Sprintf(sUpdateIssueByPkStmt, tableName))
}

const sDeleteIssueByPkStmt = `
DELETE FROM %s
 WHERE id=?
`

func DeleteIssueByPkStmt(tableName string, d dialect.Dialect) string {
	return d.ReplacePlaceholders(fmt.Sprintf(sDeleteIssueByPkStmt, tableName))
}

//--------------------------------------------------------------------------------

const sCreateIssueAssigneeStmt = `
CREATE INDEX IF NOT EXISTS issue_assignee ON %s (assignee)
`

func CreateIssueAssigneeStmt(tableName string, d dialect.Dialect) string {
	return d.ReplacePlaceholders(fmt.Sprintf(sCreateIssueAssigneeStmt, tableName))
}
