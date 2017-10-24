// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

package demo

import (
	"database/sql"
	"fmt"
	"github.com/rickb777/sqlgen2/dialect"
	"github.com/rickb777/sqlgen2/schema"
	"github.com/rickb777/sqlgen2/where"
)

// HookTableName is the default name for this table.
const HookTableName = "hooks"

// HookTable holds a given table name with the database reference, providing access methods below.
type HookTable struct {
	Name      string
	Db        *sql.DB
	DialectId schema.DialectId
}

// NewHookTable returns a new table instance.
func NewHookTable(name string, db *sql.DB, dialect schema.DialectId) HookTable {
	if name == "" {
		name = HookTableName
	}
	return HookTable{name, db, dialect}
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

// QueryOne is the low-level access function for one Hook.
func (tbl HookTable) QueryOne(query string, args ...interface{}) (*Hook, error) {
	row := tbl.Db.QueryRow(query, args...)
	return ScanHook(row)
}

// SelectOneSA allows a single Hook to be obtained from the database using supplied dialect-specific parameters.
func (tbl HookTable) SelectOneSA(where, limitClause string, args ...interface{}) (*Hook, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", sHookColumnNames, tbl.Name, where, limitClause)
	return tbl.QueryOne(query, args...)
}

// SelectOne allows a single Hook to be obtained from the database.
func (tbl HookTable) SelectOne(where where.Expression, dialect dialect.Dialect) (*Hook, error) {
	wh, args := where.Build(dialect)
	return tbl.SelectOneSA(wh, "LIMIT 1", args)
}

func (tbl HookTable) Query(query string, args ...interface{}) ([]*Hook, error) {
	rows, err := tbl.Db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanHooks(rows)
}

// SelectSA allows Hooks to be obtained from the database using supplied dialect-specific parameters.
func (tbl HookTable) SelectSA(where string, args ...interface{}) ([]*Hook, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s", sHookColumnNames, tbl.Name, where)
	return tbl.Query(query, args...)
}

// Select allows Hooks to be obtained from the database that match a 'where' clause.
func (tbl HookTable) Select(where where.Expression, dialect dialect.Dialect) ([]*Hook, error) {
	return tbl.SelectSA(where.Build(dialect))
}

// CountSA counts Hooks in the database using supplied dialect-specific parameters.
func (tbl HookTable) CountSA(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s %s", tbl.Name, where)
	row := tbl.Db.QueryRow(query, args)
	err = row.Scan(&count)
	return count, err
}

// Count counts the Hooks in the database that match a 'where' clause.
func (tbl HookTable) Count(where where.Expression, dialect dialect.Dialect) (count int64, err error) {
	return tbl.CountSA(where.Build(dialect))
}

// Insert adds new records for the Hooks.
func (tbl HookTable) Insert(v *Hook) error {
	var stmt string
	switch tbl.DialectId {
	case schema.Sqlite: stmt = sInsertHookStmtSqlite
    case schema.Postgres: stmt = sInsertHookStmtPostgres
    case schema.Mysql: stmt = sInsertHookStmtMysql
    }
	query := fmt.Sprintf(stmt, tbl.Name)
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
	var stmt string
	switch tbl.DialectId {
	case schema.Sqlite: stmt = sUpdateHookByPkStmtSqlite
    case schema.Postgres: stmt = sUpdateHookByPkStmtPostgres
    case schema.Mysql: stmt = sUpdateHookByPkStmtMysql
    }
	query := fmt.Sprintf(stmt, tbl.Name)
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

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// It returns the number of rows affected.
// Not every database or database driver may support this.
func (tbl HookTable) CreateTable() (int64, error) {
	var stmt string
	switch tbl.DialectId {
	case schema.Sqlite: stmt = sCreateHookStmtSqlite
    case schema.Postgres: stmt = sCreateHookStmtPostgres
    case schema.Mysql: stmt = sCreateHookStmtMysql
    }
	query := fmt.Sprintf(stmt, tbl.Name)
	return tbl.Exec(query)
}

//--------------------------------------------------------------------------------

const NumHookColumns = 7

const sHookColumnNames = `
id, sha, after, before, created, deleted, forced
`

const sHookDataColumnNames = `
sha, after, before, created, deleted, forced
`

const sHookColumnParamsSqlite = `
?,?,?,?,?,?,?
`

const sHookDataColumnParamsSqlite = `
?,?,?,?,?,?
`

const sCreateHookStmtSqlite = `
CREATE TABLE IF NOT EXISTS %s (
 id      INTEGER PRIMARY KEY AUTOINCREMENT,
 sha     TEXT,
 after   TEXT,
 before  TEXT,
 created BOOLEAN,
 deleted BOOLEAN,
 forced  BOOLEAN
)
`

const sInsertHookStmtSqlite = `
INSERT INTO %s (
 sha,
 after,
 before,
 created,
 deleted,
 forced
) VALUES (?,?,?,?,?,?)
`

const sUpdateHookByPkStmtSqlite = `
UPDATE %s SET 
 sha=?,
 after=?,
 before=?,
 created=?,
 deleted=?,
 forced=? 
 WHERE id=?
`

const sDeleteHookByPkStmtSqlite = `
DELETE FROM %s
 WHERE id=?
`

//--------------------------------------------------------------------------------

const sHookColumnParamsPostgres = `
$1,$2,$3,$4,$5,$6,$7
`

const sHookDataColumnParamsPostgres = `
$1,$2,$3,$4,$5,$6
`

const sCreateHookStmtPostgres = `
CREATE TABLE IF NOT EXISTS %s (
 id      SERIAL PRIMARY KEY ,
 sha     VARCHAR(512),
 after   VARCHAR(512),
 before  VARCHAR(512),
 created BOOLEAN,
 deleted BOOLEAN,
 forced  BOOLEAN
)
`

const sInsertHookStmtPostgres = `
INSERT INTO %s (
 sha,
 after,
 before,
 created,
 deleted,
 forced
) VALUES ($1,$2,$3,$4,$5,$6)
`

const sUpdateHookByPkStmtPostgres = `
UPDATE %s SET 
 sha=$2,
 after=$3,
 before=$4,
 created=$5,
 deleted=$6,
 forced=$7 
 WHERE id=$8
`

const sDeleteHookByPkStmtPostgres = `
DELETE FROM %s
 WHERE id=$1
`

//--------------------------------------------------------------------------------

const sHookColumnParamsMysql = `
?,?,?,?,?,?,?
`

const sHookDataColumnParamsMysql = `
?,?,?,?,?,?
`

const sCreateHookStmtMysql = `
CREATE TABLE IF NOT EXISTS %s (
 id      BIGINT PRIMARY KEY AUTO_INCREMENT,
 sha     VARCHAR(512),
 after   VARCHAR(512),
 before  VARCHAR(512),
 created TINYINT(1),
 deleted TINYINT(1),
 forced  TINYINT(1)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
`

const sInsertHookStmtMysql = `
INSERT INTO %s (
 sha,
 after,
 before,
 created,
 deleted,
 forced
) VALUES (?,?,?,?,?,?)
`

const sUpdateHookByPkStmtMysql = `
UPDATE %s SET 
 sha=?,
 after=?,
 before=?,
 created=?,
 deleted=?,
 forced=? 
 WHERE id=?
`

const sDeleteHookByPkStmtMysql = `
DELETE FROM %s
 WHERE id=?
`

//--------------------------------------------------------------------------------
