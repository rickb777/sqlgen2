// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

package demo

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

// IssueTableName is the default name for this table.
const IssueTableName = "issues"

// IssueTable holds a given table name with the database reference, providing access methods below.
type IssueTable struct {
	Name string
	Db   *sql.DB
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

func (tbl IssueTable) QueryOne(query string, args ...interface{}) (*Issue, error) {
	row := tbl.Db.QueryRow(query, args...)
	return ScanIssue(row)
}

func (tbl IssueTable) SelectOne(where string, args ...interface{}) (*Issue, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s LIMIT 1", sIssueColumnNames, tbl.Name, where)
	return tbl.QueryOne(query, args...)
}

func (tbl IssueTable) Query(query string, args ...interface{}) ([]*Issue, error) {
	rows, err := tbl.Db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanIssues(rows)
}

func (tbl IssueTable) Select(where string, args ...interface{}) ([]*Issue, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s", sIssueColumnNames, tbl.Name, where)
	return tbl.Query(query, args...)
}

func (tbl IssueTable) Count(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s %s", tbl.Name, where)
	row := tbl.Db.QueryRow(query, args)
	err = row.Scan(&count)
	return count, err
}

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

//--------------------------------------------------------------------------------

const sIssueColumnNames = `
id, number, title, assignee, state, labels
`

const sIssueDataColumnNames = `
number, title, assignee, state, labels
`

const sIssueColumnParams = `
$1,$2,$3,$4,$5,$6
`

const sIssueDataColumnParams = `
$1,$2,$3,$4,$5
`

const sCreateIssueStmt = `
CREATE TABLE IF NOT EXISTS %s (
 id       SERIAL PRIMARY KEY ,
 number   INTEGER,
 title    VARCHAR(512),
 assignee VARCHAR(512),
 state    VARCHAR(50),
 labels   BYTEA
);
`

func CreateIssueStmt(tableName string) string {
	return fmt.Sprintf(sCreateIssueStmt, tableName)
}

const sInsertIssueStmt = `
INSERT INTO %s (
 number,
 title,
 assignee,
 state,
 labels
) VALUES ($1,$2,$3,$4,$5)
`

func InsertIssueStmt(tableName string) string {
	return fmt.Sprintf(sInsertIssueStmt, tableName)
}

const sUpdateIssueByPkStmt = `
UPDATE %s SET 
 number=$2,
 title=$3,
 assignee=$4,
 state=$5,
 labels=$6 
 WHERE id=$7
`

func UpdateIssueByPkStmt(tableName string) string {
	return fmt.Sprintf(sUpdateIssueByPkStmt, tableName)
}

const sDeleteIssueByPkStmt = `
DELETE FROM %s
 WHERE id=$1
`

func DeleteIssueByPkStmt(tableName string) string {
	return fmt.Sprintf(sDeleteIssueByPkStmt, tableName)
}

//--------------------------------------------------------------------------------

const sCreateIssueAssigneeStmt = `
CREATE INDEX IF NOT EXISTS issue_assignee ON %s (assignee)
`

func CreateIssueAssigneeStmt(tableName string) string {
	return fmt.Sprintf(sCreateIssueAssigneeStmt, tableName)
}
