// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

package demo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/rickb777/sqlgen2/db"
	"github.com/rickb777/sqlgen2/where"
	"strings"
)

// HookTableName is the default name for this table.
const HookTableName = "hooks"

// HookTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type HookTable struct {
	Prefix, Name string
	Db           db.Execer
	Ctx          context.Context
	Dialect      db.Dialect
}

// NewHookTable returns a new table instance.
func NewHookTable(prefix, name string, d *sql.DB, dialect db.Dialect) HookTable {
	if name == "" {
		name = HookTableName
	}
	return HookTable{prefix, name, d, context.Background(), dialect}
}

// WithContext sets the context for subsequent queries.
func (tbl HookTable) WithContext(ctx context.Context) HookTable {
	tbl.Ctx = ctx
	return tbl
}

// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl HookTable) DB() *sql.DB {
	return tbl.Db.(*sql.DB)
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl HookTable) Tx() *sql.Tx {
	return tbl.Db.(*sql.Tx)
}

// IsTx tests whether this is within a transaction.
func (tbl HookTable) IsTx() bool {
	_, ok := tbl.Db.(*sql.Tx)
	return ok
}

// Begin starts a transaction. The default isolation level is dependent on the driver.
func (tbl HookTable) BeginTx(opts *sql.TxOptions) (HookTable, error) {
	d := tbl.Db.(*sql.DB)
	var err error
	tbl.Db, err = d.BeginTx(tbl.Ctx, opts)
	return tbl, err
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

//--------------------------------------------------------------------------------

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// It returns the number of rows affected.
func (tbl HookTable) Exec(query string, args ...interface{}) (int64, error) {
	res, err := tbl.Db.ExecContext(tbl.Ctx, query, args...)
	if err != nil {
		return 0, nil
	}
	return res.RowsAffected()
}

//--------------------------------------------------------------------------------

// QueryOne is the low-level access function for one Hook.
func (tbl HookTable) QueryOne(query string, args ...interface{}) (*Hook, error) {
	row := tbl.Db.QueryRowContext(tbl.Ctx, query, args...)
	return ScanHook(row)
}

// Query is the low-level access function for Hooks.
func (tbl HookTable) Query(query string, args ...interface{}) ([]*Hook, error) {
	rows, err := tbl.Db.QueryContext(tbl.Ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanHooks(rows)
}

//--------------------------------------------------------------------------------

// SelectOneSA allows a single Hook to be obtained from the database using supplied dialect-specific parameters.
func (tbl HookTable) SelectOneSA(where, limitClause string, args ...interface{}) (*Hook, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", HookColumnNames, tbl.Prefix, tbl.Name, where, limitClause)
	return tbl.QueryOne(query, args...)
}

// SelectOne allows a single Hook to be obtained from the database.
func (tbl HookTable) SelectOne(where where.Expression, dialect db.Dialect) (*Hook, error) {
	wh, args := where.Build(dialect)
	return tbl.SelectOneSA(wh, "LIMIT 1", args)
}

// SelectSA allows Hooks to be obtained from the database using supplied dialect-specific parameters.
func (tbl HookTable) SelectSA(where string, args ...interface{}) ([]*Hook, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s", HookColumnNames, tbl.Prefix, tbl.Name, where)
	return tbl.Query(query, args...)
}

// Select allows Hooks to be obtained from the database that match a 'where' clause.
func (tbl HookTable) Select(where where.Expression, dialect db.Dialect) ([]*Hook, error) {
	return tbl.SelectSA(where.Build(dialect))
}

// CountSA counts Hooks in the database using supplied dialect-specific parameters.
func (tbl HookTable) CountSA(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s%s %s", tbl.Prefix, tbl.Name, where)
	row := tbl.Db.QueryRowContext(tbl.Ctx, query, args)
	err = row.Scan(&count)
	return count, err
}

// Count counts the Hooks in the database that match a 'where' clause.
func (tbl HookTable) Count(where where.Expression, dialect db.Dialect) (count int64, err error) {
	return tbl.CountSA(where.Build(dialect))
}

const HookColumnNames = "id, sha, after, before, created, deleted, forced"

//--------------------------------------------------------------------------------

// Insert adds new records for the Hooks. The Hooks have their primary key fields
// set to the new record identifiers.
func (tbl HookTable) Insert(vv ...*Hook) error {
	var stmt, params string
	switch tbl.Dialect {
	case db.Postgres:
		stmt = sqlInsertHookPostgres
		params = sHookDataColumnParamsPostgres
	default:
		stmt = sqlInsertHookSimple
		params = sHookDataColumnParamsSimple
	}
	st, err := tbl.Db.PrepareContext(tbl.Ctx, fmt.Sprintf(stmt, tbl.Prefix, tbl.Name, params))
	if err != nil {
		return err
	}
	defer st.Close()

	for _, v := range vv {
		res, err := st.Exec(SliceHookWithoutPk(v)...)
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

const sqlInsertHookSimple = `
INSERT INTO %s%s (
	sha,
	after,
	before,
	created,
	deleted,
	forced
) VALUES (%s)
`

const sqlInsertHookPostgres = `
INSERT INTO %s%s (
	sha,
	after,
	before,
	created,
	deleted,
	forced
) VALUES (%s)
`

const sHookDataColumnParamsSimple = "?,?,?,?,?,?"

const sHookDataColumnParamsPostgres = "$1,$2,$3,$4,$5,$6"

//--------------------------------------------------------------------------------

// Update updates a record. It returns the number of rows affected.
// Not every database or database driver may support this.
func (tbl HookTable) Update(v *Hook) (int64, error) {
	var stmt string
	switch tbl.Dialect {
	case db.Postgres:
		stmt = sqlUpdateHookByPkPostgres
	default:
		stmt = sqlUpdateHookByPkSimple
	}
	query := fmt.Sprintf(stmt, tbl.Prefix, tbl.Name)
	args := SliceHookWithoutPk(v)
	args = append(args, v.Id)
	return tbl.Exec(query, args...)
}

// UpdateFields updates one or more columns, given a 'where' clause.
func (tbl HookTable) UpdateFields(where where.Expression, fields ...sql.NamedArg) (int64, error) {
	return tbl.Exec(tbl.updateFields(where, fields...))
}

func (tbl HookTable) updateFields(where where.Expression, fields ...sql.NamedArg) (string, []interface{}) {
	list := db.NamedArgList(fields)
	assignments := strings.Join(list.Assignments(tbl.Dialect, 1), ", ")
	whereClause, extra := where.Build(tbl.Dialect)
	query := fmt.Sprintf("UPDATE %s%s SET %s %s", tbl.Prefix, tbl.Name, assignments, whereClause)
	args := append(list.Values(), extra...)
	return query, args
}

const sqlUpdateHookByPkSimple = `
UPDATE %s%s SET 
	sha=?,
	after=?,
	before=?,
	created=?,
	deleted=?,
	forced=? 
 WHERE id=?
`

const sqlUpdateHookByPkPostgres = `
UPDATE %s%s SET 
	sha=$2,
	after=$3,
	before=$4,
	created=$5,
	deleted=$6,
	forced=$7 
 WHERE id=$8
`

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// It returns the number of rows affected.
// Not every database or database driver may support this.
func (tbl HookTable) CreateTable(ifNotExist bool) (int64, error) {
	return tbl.Exec(tbl.createTableSql(ifNotExist))
}

func (tbl HookTable) createTableSql(ifNotExist bool) string {
	var stmt string
	switch tbl.Dialect {
	case db.Sqlite: stmt = sqlCreateHookTableSqlite
    case db.Postgres: stmt = sqlCreateHookTablePostgres
    case db.Mysql: stmt = sqlCreateHookTableMysql
    }
	extra := tbl.ternary(ifNotExist, "IF NOT EXISTS ", "")
	query := fmt.Sprintf(stmt, extra, tbl.Prefix, tbl.Name)
	return query
}

func (tbl HookTable) ternary(flag bool, a, b string) string {
	if flag {
		return a
	}
	return b
}

// CreateIndexes executes queries that create the indexes needed by the Hook table.
func (tbl HookTable) CreateIndexes(ifNotExist bool) (err error) {
	
	return nil
}




//--------------------------------------------------------------------------------

const sqlCreateHookTableSqlite = `
CREATE TABLE %s%s%s (
 id      integer primary key autoincrement,
 sha     text,
 after   text,
 before  text,
 created boolean,
 deleted boolean,
 forced  boolean
)
`

const sqlCreateHookTablePostgres = `
CREATE TABLE %s%s%s (
 id      bigserial primary key ,
 sha     varchar(512),
 after   varchar(512),
 before  varchar(512),
 created boolean,
 deleted boolean,
 forced  boolean
)
`

const sqlCreateHookTableMysql = `
CREATE TABLE %s%s%s (
 id      bigint primary key auto_increment,
 sha     varchar(512),
 after   varchar(512),
 before  varchar(512),
 created tinyint(1),
 deleted tinyint(1),
 forced  tinyint(1)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
`

//--------------------------------------------------------------------------------

const sqlDeleteHookByPkPostgres = `
DELETE FROM %s%s
 WHERE id=$1
`

const sqlDeleteHookByPkSimple = `
DELETE FROM %s%s
 WHERE id=?
`

//--------------------------------------------------------------------------------

const NumHookColumns = 7

const NumHookDataColumns = 6

const HookPk = "Id"

const HookDataColumnNames = "sha, after, before, created, deleted, forced"

//--------------------------------------------------------------------------------
