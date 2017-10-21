// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

package demo

import (
	"database/sql"
	"fmt"
)

// HookTableName is the default name for this table.
const HookTableName = "hooks"

// HookTable holds a given table name with the database reference, providing access methods below.
type HookTable struct {
	Name string
	Db   *sql.DB
}

// ScanHook reads a database record into a single value.
func ScanHook(row *sql.Row) (*Hook, error) {
	var v0 int64
	var v1 string
	var v2 string
	var v3 string
	var v4 bool
	var v5 bool
	var v6 bool
	var v7 string
	var v8 string
	var v9 string
	var v10 string
	var v11 string
	var v12 string
	var v13 string
	var v14 string
	var v15 string

	err := row.Scan(
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

	)
	if err != nil {
		return nil, err
	}

	v := &Hook{}
	v.Id = v0
	v.Sha = v1
	v.After = v2
	v.Before = v3
	v.Created = v4
	v.Deleted = v5
	v.Forced = v6
	v.HeadCommit = &Commit{}
	v.HeadCommit.ID = v7
	v.HeadCommit.Message = v8
	v.HeadCommit.Timestamp = v9
	v.HeadCommit.Author = &Author{}
	v.HeadCommit.Author.Name = v10
	v.HeadCommit.Author.Email = v11
	v.HeadCommit.Author.Username = v12
	v.HeadCommit.Committer = &Author{}
	v.HeadCommit.Committer.Name = v13
	v.HeadCommit.Committer.Email = v14
	v.HeadCommit.Committer.Username = v15

	return v, nil
}

// ScanHooks reads database records into a slice of values.
func ScanHooks(rows *sql.Rows) ([]*Hook, error) {
	var err error
	var vv []*Hook

	var v0 int64
	var v1 string
	var v2 string
	var v3 string
	var v4 bool
	var v5 bool
	var v6 bool
	var v7 string
	var v8 string
	var v9 string
	var v10 string
	var v11 string
	var v12 string
	var v13 string
	var v14 string
	var v15 string

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
			&v8,
			&v9,
			&v10,
			&v11,
			&v12,
			&v13,
			&v14,
			&v15,

		)
		if err != nil {
			return vv, err
		}

		v := &Hook{}
		v.Id = v0
		v.Sha = v1
		v.After = v2
		v.Before = v3
		v.Created = v4
		v.Deleted = v5
		v.Forced = v6
		v.HeadCommit = &Commit{}
		v.HeadCommit.ID = v7
		v.HeadCommit.Message = v8
		v.HeadCommit.Timestamp = v9
		v.HeadCommit.Author = &Author{}
		v.HeadCommit.Author.Name = v10
		v.HeadCommit.Author.Email = v11
		v.HeadCommit.Author.Username = v12
		v.HeadCommit.Committer = &Author{}
		v.HeadCommit.Committer.Name = v13
		v.HeadCommit.Committer.Email = v14
		v.HeadCommit.Committer.Username = v15

		vv = append(vv, v)
	}
	return vv, rows.Err()
}

func SliceHook(v *Hook) []interface{} {
	var v0 int64
	var v1 string
	var v2 string
	var v3 string
	var v4 bool
	var v5 bool
	var v6 bool
	var v7 string
	var v8 string
	var v9 string
	var v10 string
	var v11 string
	var v12 string
	var v13 string
	var v14 string
	var v15 string

	v0 = v.Id
	v1 = v.Sha
	v2 = v.After
	v3 = v.Before
	v4 = v.Created
	v5 = v.Deleted
	v6 = v.Forced
	if v.HeadCommit != nil {
		v7 = v.HeadCommit.ID
		v8 = v.HeadCommit.Message
		v9 = v.HeadCommit.Timestamp
		if v.HeadCommit.Author != nil {
			v10 = v.HeadCommit.Author.Name
			v11 = v.HeadCommit.Author.Email
			v12 = v.HeadCommit.Author.Username
		}
	}
	if v.HeadCommit.Committer != nil {
		v13 = v.HeadCommit.Committer.Name
		v14 = v.HeadCommit.Committer.Email
		v15 = v.HeadCommit.Committer.Username
	}

	return []interface{}{
		v0,
		v1,
		v2,
		v3,
		v4,
		v5,
		v6,
		v7,
		v8,
		v9,
		v10,
		v11,
		v12,
		v13,
		v14,
		v15,

	}
}

func SliceHookWithoutPk(v *Hook) []interface{} {
	var v1 string
	var v2 string
	var v3 string
	var v4 bool
	var v5 bool
	var v6 bool
	var v7 string
	var v8 string
	var v9 string
	var v10 string
	var v11 string
	var v12 string
	var v13 string
	var v14 string
	var v15 string

	v1 = v.Sha
	v2 = v.After
	v3 = v.Before
	v4 = v.Created
	v5 = v.Deleted
	v6 = v.Forced
	if v.HeadCommit != nil {
		v7 = v.HeadCommit.ID
		v8 = v.HeadCommit.Message
		v9 = v.HeadCommit.Timestamp
		if v.HeadCommit.Author != nil {
			v10 = v.HeadCommit.Author.Name
			v11 = v.HeadCommit.Author.Email
			v12 = v.HeadCommit.Author.Username
		}
	}
	if v.HeadCommit.Committer != nil {
		v13 = v.HeadCommit.Committer.Name
		v14 = v.HeadCommit.Committer.Email
		v15 = v.HeadCommit.Committer.Username
	}

	return []interface{}{
		v1,
		v2,
		v3,
		v4,
		v5,
		v6,
		v7,
		v8,
		v9,
		v10,
		v11,
		v12,
		v13,
		v14,
		v15,

	}
}

func (tbl HookTable) QueryOne(query string, args ...interface{}) (*Hook, error) {
	row := tbl.Db.QueryRow(query, args...)
	return ScanHook(row)
}

func (tbl HookTable) SelectOne(where string, args ...interface{}) (*Hook, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s LIMIT 1", sHookColumnNames, tbl.Name, where)
	return tbl.QueryOne(query, args...)
}

func (tbl HookTable) Query(query string, args ...interface{}) ([]*Hook, error) {
	rows, err := tbl.Db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanHooks(rows)
}

func (tbl HookTable) Select(where string, args ...interface{}) ([]*Hook, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s", sHookColumnNames, tbl.Name, where)
	return tbl.Query(query, args...)
}

func (tbl HookTable) Count(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s %s", tbl.Name, where)
	row := tbl.Db.QueryRow(query, args)
	err = row.Scan(&count)
	return count, err
}

func (tbl HookTable) Insert(v *Hook) error {
	query := fmt.Sprintf(sInsertHookStmt, tbl.Name)
	res, err := tbl.Db.Exec(query, SliceHookWithoutPk(v)...)
	if err != nil {
		return err
	}

	v.Id, err = res.LastInsertId()
	return err
}

// Update updates a record. It returns the number of rows affected.
// Not every database or database driver may support this.
func (tbl HookTable) Update(v *Hook) (int64, error) {
	query := fmt.Sprintf(sUpdateHookByPkStmt, tbl.Name)
	args := SliceHookWithoutPk(v)
	args = append(args, v.Id)
	return tbl.Exec(query, args...)
}

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// It returns the number of rows affected.
// Not every database or database driver may support this.
func (tbl HookTable) Exec(query string, args ...interface{}) (int64, error) {
	res, err := tbl.Db.Exec(query, args...)
	if err != nil {
		return 0, nil
	}
	return res.RowsAffected()
}

//--------------------------------------------------------------------------------

const sHookColumnNames = `
id, sha, after, before, created, deleted, forced
`

const sHookDataColumnNames = `
sha, after, before, created, deleted, forced
`

const sHookColumnParams = `
?,?,?,?,?,?,?
`

const sHookDataColumnParams = `
?,?,?,?,?,?
`

const sCreateHookStmt = `
CREATE TABLE IF NOT EXISTS %s (
 id      INTEGER PRIMARY KEY AUTOINCREMENT,
 sha     TEXT,
 after   TEXT,
 before  TEXT,
 created BOOLEAN,
 deleted BOOLEAN,
 forced  BOOLEAN
);
`

func CreateHookStmt(tableName string) string {
	return fmt.Sprintf(sCreateHookStmt, tableName)
}

const sInsertHookStmt = `
INSERT INTO %s (
 sha,
 after,
 before,
 created,
 deleted,
 forced
) VALUES (?,?,?,?,?,?)
`

func InsertHookStmt(tableName string) string {
	return fmt.Sprintf(sInsertHookStmt, tableName)
}

const sUpdateHookByPkStmt = `
UPDATE %s SET 
 sha=?,
 after=?,
 before=?,
 created=?,
 deleted=?,
 forced=? 
 WHERE id=?
`

func UpdateHookByPkStmt(tableName string) string {
	return fmt.Sprintf(sUpdateHookByPkStmt, tableName)
}

const sDeleteHookByPkStmt = `
DELETE FROM %s
 WHERE id=?
`

func DeleteHookByPkStmt(tableName string) string {
	return fmt.Sprintf(sDeleteHookByPkStmt, tableName)
}

//--------------------------------------------------------------------------------
