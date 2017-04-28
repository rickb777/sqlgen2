// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

package demo

import (
	"database/sql"
)

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

func SelectHook(db *sql.DB, query string, args ...interface{}) (*Hook, error) {
	row := db.QueryRow(query, args...)
	return ScanHook(row)
}

func SelectHooks(db *sql.DB, query string, args ...interface{}) ([]*Hook, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanHooks(rows)
}

func InsertHook(db *sql.DB, query string, v *Hook) error {
	res, err := db.Exec(query, SliceHook(v)[1:]...)
	if err != nil {
		return err
	}

	v.Id, err = res.LastInsertId()
	return err
}

func UpdateHook(db *sql.DB, query string, v *Hook) error {
	args := SliceHook(v)[1:]
	args = append(args, v.Id)
	_, err := db.Exec(query, args...)
	return err
}

const CreateHookStmt = `
CREATE TABLE IF NOT EXISTS hooks (
 id                      INTEGER PRIMARY KEY AUTO_INCREMENT,
 sha                     VARCHAR(512),
 after                   VARCHAR(512),
 before                  VARCHAR(512),
 created                 BOOLEAN,
 deleted                 BOOLEAN,
 forced                  BOOLEAN,
 head_id                 VARCHAR(512),
 head_message            VARCHAR(512),
 head_timestamp          VARCHAR(512),
 head_author_name        VARCHAR(512),
 head_author_email       VARCHAR(512),
 head_author_username    VARCHAR(512),
 head_committer_name     VARCHAR(512),
 head_committer_email    VARCHAR(512),
 head_committer_username VARCHAR(512)
);
`

const InsertHookStmt = `
INSERT INTO hooks (
 sha,
 after,
 before,
 created,
 deleted,
 forced,
 head_id,
 head_message,
 head_timestamp,
 head_author_name,
 head_author_email,
 head_author_username,
 head_committer_name,
 head_committer_email,
 head_committer_username
) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)
`

const SelectHookStmt = `
SELECT 
 id,
 sha,
 after,
 before,
 created,
 deleted,
 forced,
 head_id,
 head_message,
 head_timestamp,
 head_author_name,
 head_author_email,
 head_author_username,
 head_committer_name,
 head_committer_email,
 head_committer_username
FROM hooks 
`

const SelectHookRangeStmt = `
SELECT 
 id,
 sha,
 after,
 before,
 created,
 deleted,
 forced,
 head_id,
 head_message,
 head_timestamp,
 head_author_name,
 head_author_email,
 head_author_username,
 head_committer_name,
 head_committer_email,
 head_committer_username
FROM hooks 
LIMIT ? OFFSET ?
`

const SelectHookCountStmt = `
SELECT count(1)
FROM hooks 
`

const SelectHookPkeyStmt = `
SELECT 
 id,
 sha,
 after,
 before,
 created,
 deleted,
 forced,
 head_id,
 head_message,
 head_timestamp,
 head_author_name,
 head_author_email,
 head_author_username,
 head_committer_name,
 head_committer_email,
 head_committer_username
FROM hooks 
WHERE id=?
`

const UpdateHookPkeyStmt = `
UPDATE hooks SET 
 id=?,
 sha=?,
 after=?,
 before=?,
 created=?,
 deleted=?,
 forced=?,
 head_id=?,
 head_message=?,
 head_timestamp=?,
 head_author_name=?,
 head_author_email=?,
 head_author_username=?,
 head_committer_name=?,
 head_committer_email=?,
 head_committer_username=? 
WHERE id=?
`

const DeleteHookPkeyStmt = `
DELETE FROM hooks 
WHERE id=?
`
