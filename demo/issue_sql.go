
package demo

// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

import (
	"encoding/json"
	"database/sql"
)

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
	v.Id=v0
v.Number=v1
v.Title=v2
v.Body=v3
v.Assignee=v4
v.State=v5
json.Unmarshal(v6, &v.Labels)


	return v, nil
}

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
		v.Id=v0
v.Number=v1
v.Title=v2
v.Body=v3
v.Assignee=v4
v.State=v5
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

	v0=v.Id
v1=v.Number
v2=v.Title
v3=v.Body
v4=v.Assignee
v5=v.State
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

func SelectIssue(db *sql.DB, query string, args ...interface{}) (*Issue, error) {
	row := db.QueryRow(query, args...)
	return ScanIssue(row)
}

func SelectIssues(db *sql.DB, query string, args ...interface{}) ([]*Issue, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanIssues(rows)
}

func InsertIssue(db *sql.DB, query string, v *Issue) error {
	res, err := db.Exec(query, SliceIssue(v)[1:]...)
	if err != nil {
		return err
	}

	v.Id, err = res.LastInsertId()
	return err
}

func UpdateIssue(db *sql.DB, query string, v *Issue) error {
	args := SliceIssue(v)[1:]
	args = append(args, v.Id)
	_, err := db.Exec(query, args...)
	return err 
}

const CreateIssueStmt = `
CREATE TABLE IF NOT EXISTS issues (
 id       SERIAL PRIMARY KEY ,
 number   INTEGER,
 title    VARCHAR(512),
 body     VARCHAR(2048),
 assignee VARCHAR(512),
 state    VARCHAR(50),
 labels   BYTEA
);
`

const InsertIssueStmt = `
INSERT INTO issues (
 number,
 title,
 body,
 assignee,
 state,
 labels
) VALUES ($1,$2,$3,$4,$5,$6)
`

const SelectIssueStmt = `
SELECT 
 id,
 number,
 title,
 body,
 assignee,
 state,
 labels
FROM issues 
`

const SelectIssueRangeStmt = `
SELECT 
 id,
 number,
 title,
 body,
 assignee,
 state,
 labels
FROM issues 
LIMIT $1 OFFSET $2
`

const SelectIssueCountStmt = `
SELECT count(1)
FROM issues 
`

const SelectIssuePkeyStmt = `
SELECT 
 id,
 number,
 title,
 body,
 assignee,
 state,
 labels
FROM issues 
WHERE id=$1
`

const UpdateIssuePkeyStmt = `
UPDATE issues SET 
 id=$1,
 number=$2,
 title=$3,
 body=$4,
 assignee=$5,
 state=$6,
 labels=$7 
WHERE id=$8
`

const DeleteIssuePkeyStmt = `
DELETE FROM issues 
WHERE id=$1
`

const CreateIssueAssigneeStmt = `
CREATE INDEX IF NOT EXISTS issue_assignee ON issues ( assignee)
`

const SelectIssueAssigneeStmt = `
SELECT 
 id,
 number,
 title,
 body,
 assignee,
 state,
 labels
FROM issues 
WHERE assignee=$1
`

const SelectIssueAssigneeRangeStmt = `
SELECT 
 id,
 number,
 title,
 body,
 assignee,
 state,
 labels
FROM issues 
WHERE assignee=$1
LIMIT $2 OFFSET $3
`

const SelectIssueAssigneeCountStmt = `
SELECT count(1)
FROM issues 
WHERE assignee=$1
`
