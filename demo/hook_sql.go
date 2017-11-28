// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

package demo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/rickb777/sqlgen2"
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
	Db           sqlgen2.Execer
	Ctx          context.Context
	Dialect      sqlgen2.Dialect
}

// Type conformance check
var _ sqlgen2.Table = HookTable{}

// NewHookTable returns a new table instance.
func NewHookTable(prefix, name string, d *sql.DB, dialect sqlgen2.Dialect) HookTable {
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

// FullName gets the concatenated prefix and table name.
func (tbl HookTable) FullName() string {
	return tbl.Prefix + tbl.Name
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


// ScanHook reads a table record into a single value.
func ScanHook(row *sql.Row) (*Hook, error) {
	var v0 int64
	var v1 string
	var v2 string
	var v3 string
	var v4 Category
	var v5 bool
	var v6 bool
	var v7 bool
	var v8 string
	var v9 string
	var v10 string
	var v11 string
	var v12 string
	var v13 string
	var v14 string
	var v15 string
	var v16 string

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
		&v16,

	)
	if err != nil {
		return nil, err
	}

	v := &Hook{}
	v.Id = v0
	v.Sha = v1
	v.Dates.After = v2
	v.Dates.Before = v3
	v.Category = v4
	v.Created = v5
	v.Deleted = v6
	v.Forced = v7
	v.HeadCommit.ID = v8
	v.HeadCommit.Message = v9
	v.HeadCommit.Timestamp = v10
	v.HeadCommit.Author.Name = v11
	v.HeadCommit.Author.Email = v12
	v.HeadCommit.Author.Username = v13
	v.HeadCommit.Committer.Name = v14
	v.HeadCommit.Committer.Email = v15
	v.HeadCommit.Committer.Username = v16

	return v, nil
}

// ScanHooks reads table records into a slice of values.
func ScanHooks(rows *sql.Rows) ([]*Hook, error) {
	var err error
	var vv []*Hook

	var v0 int64
	var v1 string
	var v2 string
	var v3 string
	var v4 Category
	var v5 bool
	var v6 bool
	var v7 bool
	var v8 string
	var v9 string
	var v10 string
	var v11 string
	var v12 string
	var v13 string
	var v14 string
	var v15 string
	var v16 string

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
			&v16,

		)
		if err != nil {
			return vv, err
		}

		v := &Hook{}
		v.Id = v0
		v.Sha = v1
		v.Dates.After = v2
		v.Dates.Before = v3
		v.Category = v4
		v.Created = v5
		v.Deleted = v6
		v.Forced = v7
		v.HeadCommit.ID = v8
		v.HeadCommit.Message = v9
		v.HeadCommit.Timestamp = v10
		v.HeadCommit.Author.Name = v11
		v.HeadCommit.Author.Email = v12
		v.HeadCommit.Author.Username = v13
		v.HeadCommit.Committer.Name = v14
		v.HeadCommit.Committer.Email = v15
		v.HeadCommit.Committer.Username = v16

		vv = append(vv, v)
	}
	return vv, rows.Err()
}

func SliceHook(v *Hook) ([]interface{}, error) {
	var v0 int64
	var v1 string
	var v2 string
	var v3 string
	var v4 Category
	var v5 bool
	var v6 bool
	var v7 bool
	var v8 string
	var v9 string
	var v10 string
	var v11 string
	var v12 string
	var v13 string
	var v14 string
	var v15 string
	var v16 string

	v0 = v.Id
	v1 = v.Sha
	v2 = v.Dates.After
	v3 = v.Dates.Before
	v4 = v.Category
	v5 = v.Created
	v6 = v.Deleted
	v7 = v.Forced
	v8 = v.HeadCommit.ID
	v9 = v.HeadCommit.Message
	v10 = v.HeadCommit.Timestamp
	v11 = v.HeadCommit.Author.Name
	v12 = v.HeadCommit.Author.Email
	v13 = v.HeadCommit.Author.Username
	v14 = v.HeadCommit.Committer.Name
	v15 = v.HeadCommit.Committer.Email
	v16 = v.HeadCommit.Committer.Username

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
		v16,

	}, nil
}

func SliceHookWithoutPk(v *Hook) ([]interface{}, error) {
	var v1 string
	var v2 string
	var v3 string
	var v4 Category
	var v5 bool
	var v6 bool
	var v7 bool
	var v8 string
	var v9 string
	var v10 string
	var v11 string
	var v12 string
	var v13 string
	var v14 string
	var v15 string
	var v16 string

	v1 = v.Sha
	v2 = v.Dates.After
	v3 = v.Dates.Before
	v4 = v.Category
	v5 = v.Created
	v6 = v.Deleted
	v7 = v.Forced
	v8 = v.HeadCommit.ID
	v9 = v.HeadCommit.Message
	v10 = v.HeadCommit.Timestamp
	v11 = v.HeadCommit.Author.Name
	v12 = v.HeadCommit.Author.Email
	v13 = v.HeadCommit.Author.Username
	v14 = v.HeadCommit.Committer.Name
	v15 = v.HeadCommit.Committer.Email
	v16 = v.HeadCommit.Committer.Username

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
		v16,

	}, nil
}

//--------------------------------------------------------------------------------

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// It returns the number of rows affected (of the database drive supports this).
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

// SelectOneSA allows a single Hook to be obtained from the table that match a 'where' clause and some limit.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
func (tbl HookTable) SelectOneSA(where, orderBy string, args ...interface{}) (*Hook, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s LIMIT 1", HookColumnNames, tbl.Prefix, tbl.Name, where, orderBy)
	return tbl.QueryOne(query, args...)
}

// SelectOne allows a single Hook to be obtained from the sqlgen2.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
func (tbl HookTable) SelectOne(where where.Expression, orderBy string) (*Hook, error) {
	wh, args := where.Build(tbl.Dialect)
	return tbl.SelectOneSA(wh, orderBy, args)
}

// SelectSA allows Hooks to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
func (tbl HookTable) SelectSA(where, orderBy string, args ...interface{}) ([]*Hook, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", HookColumnNames, tbl.Prefix, tbl.Name, where, orderBy)
	return tbl.Query(query, args...)
}

// Select allows Hooks to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
func (tbl HookTable) Select(where where.Expression, orderBy string) ([]*Hook, error) {
	wh, args := where.Build(tbl.Dialect)
	return tbl.SelectSA(wh, orderBy, args)
}

// CountSA counts Hooks in the table that match a 'where' clause.
func (tbl HookTable) CountSA(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s%s %s", tbl.Prefix, tbl.Name, where)
	row := tbl.Db.QueryRowContext(tbl.Ctx, query, args)
	err = row.Scan(&count)
	return count, err
}

// Count counts the Hooks in the table that match a 'where' clause.
func (tbl HookTable) Count(where where.Expression) (count int64, err error) {
	return tbl.CountSA(where.Build(tbl.Dialect))
}

const HookColumnNames = "id, sha, dates_after, dates_before, category, created, deleted, forced, head_commit_id, head_commit_message, head_commit_timestamp, head_commit_author_name, head_commit_author_email, head_commit_author_username, head_commit_committer_name, head_commit_committer_email, head_commit_committer_username"

//--------------------------------------------------------------------------------

// Insert adds new records for the Hooks. The Hooks have their primary key fields
// set to the new record identifiers.
// The Hook.PreInsert(Execer) method will be called, if it exists.
func (tbl HookTable) Insert(vv ...*Hook) error {
	var stmt, params string
	switch tbl.Dialect {
	case sqlgen2.Postgres:
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
		var iv interface{} = v
		if hook, ok := iv.(interface{PreInsert(sqlgen2.Execer)}); ok {
			hook.PreInsert(tbl.Db)
		}

		fields, err := SliceHookWithoutPk(v)
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

const sqlInsertHookSimple = `
INSERT INTO %s%s (
	sha,
	dates_after,
	dates_before,
	category,
	created,
	deleted,
	forced,
	head_commit_id,
	head_commit_message,
	head_commit_timestamp,
	head_commit_author_name,
	head_commit_author_email,
	head_commit_author_username,
	head_commit_committer_name,
	head_commit_committer_email,
	head_commit_committer_username
) VALUES (%s)
`

const sqlInsertHookPostgres = `
INSERT INTO %s%s (
	sha,
	dates_after,
	dates_before,
	category,
	created,
	deleted,
	forced,
	head_commit_id,
	head_commit_message,
	head_commit_timestamp,
	head_commit_author_name,
	head_commit_author_email,
	head_commit_author_username,
	head_commit_committer_name,
	head_commit_committer_email,
	head_commit_committer_username
) VALUES (%s)
`

const sHookDataColumnParamsSimple = "?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?"

const sHookDataColumnParamsPostgres = "$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16"

//--------------------------------------------------------------------------------

// UpdateFields updates one or more columns, given a 'where' clause.
func (tbl HookTable) UpdateFields(where where.Expression, fields ...sql.NamedArg) (int64, error) {
	return tbl.Exec(tbl.updateFields(where, fields...))
}

func (tbl HookTable) updateFields(where where.Expression, fields ...sql.NamedArg) (string, []interface{}) {
	list := sqlgen2.NamedArgList(fields)
	assignments := strings.Join(list.Assignments(tbl.Dialect, 1), ", ")
	whereClause, wargs := where.Build(tbl.Dialect)
	query := fmt.Sprintf("UPDATE %s%s SET %s %s", tbl.Prefix, tbl.Name, assignments, whereClause)
	args := append(list.Values(), wargs...)
	return query, args
}

// Update updates records, matching them by primary key. It returns the number of rows affected.
// The Hook.PreUpdate(Execer) method will be called, if it exists.
func (tbl HookTable) Update(vv ...*Hook) (int64, error) {
	var stmt string
	switch tbl.Dialect {
	case sqlgen2.Postgres:
		stmt = sqlUpdateHookByPkPostgres
	default:
		stmt = sqlUpdateHookByPkSimple
	}

	var count int64
	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(interface{PreUpdate(sqlgen2.Execer)}); ok {
			hook.PreUpdate(tbl.Db)
		}

		query := fmt.Sprintf(stmt, tbl.Prefix, tbl.Name)

		args, err := SliceHookWithoutPk(v)
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

const sqlUpdateHookByPkSimple = `
UPDATE %s%s SET 
	sha=?,
	dates_after=?,
	dates_before=?,
	category=?,
	created=?,
	deleted=?,
	forced=?,
	head_commit_id=?,
	head_commit_message=?,
	head_commit_timestamp=?,
	head_commit_author_name=?,
	head_commit_author_email=?,
	head_commit_author_username=?,
	head_commit_committer_name=?,
	head_commit_committer_email=?,
	head_commit_committer_username=? 
 WHERE id=?
`

const sqlUpdateHookByPkPostgres = `
UPDATE %s%s SET 
	sha=$2,
	dates_after=$3,
	dates_before=$4,
	category=$5,
	created=$6,
	deleted=$7,
	forced=$8,
	head_commit_id=$9,
	head_commit_message=$10,
	head_commit_timestamp=$11,
	head_commit_author_name=$12,
	head_commit_author_email=$13,
	head_commit_author_username=$14,
	head_commit_committer_name=$15,
	head_commit_committer_email=$16,
	head_commit_committer_username=$17 
 WHERE id=$18
`

//--------------------------------------------------------------------------------

// Delete deletes one or more rows from the table, given a 'where' clause.
func (tbl HookTable) Delete(where where.Expression) (int64, error) {
	return tbl.Exec(tbl.deleteRows(where))
}

func (tbl HookTable) deleteRows(where where.Expression) (string, []interface{}) {
	whereClause, args := where.Build(tbl.Dialect)
	query := fmt.Sprintf("DELETE FROM %s%s %s", tbl.Prefix, tbl.Name, whereClause)
	return query, args
}

//--------------------------------------------------------------------------------

// CreateTable creates the table.
func (tbl HookTable) CreateTable(ifNotExist bool) (int64, error) {
	return tbl.Exec(tbl.createTableSql(ifNotExist))
}

func (tbl HookTable) createTableSql(ifNotExist bool) string {
	var stmt string
	switch tbl.Dialect {
	case sqlgen2.Sqlite: stmt = sqlCreateHookTableSqlite
    case sqlgen2.Postgres: stmt = sqlCreateHookTablePostgres
    case sqlgen2.Mysql: stmt = sqlCreateHookTableMysql
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

//--------------------------------------------------------------------------------

// CreateIndexes executes queries that create the indexes needed by the Hook table.
func (tbl HookTable) CreateIndexes(ifNotExist bool) (err error) {
	
	return nil
}



//--------------------------------------------------------------------------------

const sqlCreateHookTableSqlite = `
CREATE TABLE %s%s%s (
 id                             integer primary key autoincrement,
 sha                            text,
 dates_after                    text,
 dates_before                   text,
 category                       integer,
 created                        boolean,
 deleted                        boolean,
 forced                         boolean,
 head_commit_id                 text,
 head_commit_message            text,
 head_commit_timestamp          text,
 head_commit_author_name        text,
 head_commit_author_email       text,
 head_commit_author_username    text,
 head_commit_committer_name     text,
 head_commit_committer_email    text,
 head_commit_committer_username text
)
`

const sqlCreateHookTablePostgres = `
CREATE TABLE %s%s%s (
 id                             bigserial primary key ,
 sha                            varchar(512),
 dates_after                    varchar(512),
 dates_before                   varchar(512),
 category                       integer,
 created                        boolean,
 deleted                        boolean,
 forced                         boolean,
 head_commit_id                 varchar(512),
 head_commit_message            varchar(512),
 head_commit_timestamp          varchar(512),
 head_commit_author_name        varchar(512),
 head_commit_author_email       varchar(512),
 head_commit_author_username    varchar(512),
 head_commit_committer_name     varchar(512),
 head_commit_committer_email    varchar(512),
 head_commit_committer_username varchar(512)
)
`

const sqlCreateHookTableMysql = `
CREATE TABLE %s%s%s (
 id                             bigint primary key auto_increment,
 sha                            varchar(512),
 dates_after                    varchar(512),
 dates_before                   varchar(512),
 category                       tinyint unsigned,
 created                        tinyint(1),
 deleted                        tinyint(1),
 forced                         tinyint(1),
 head_commit_id                 varchar(512),
 head_commit_message            varchar(512),
 head_commit_timestamp          varchar(512),
 head_commit_author_name        varchar(512),
 head_commit_author_email       varchar(512),
 head_commit_author_username    varchar(512),
 head_commit_committer_name     varchar(512),
 head_commit_committer_email    varchar(512),
 head_commit_committer_username varchar(512)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
`

//--------------------------------------------------------------------------------

//--------------------------------------------------------------------------------

const NumHookColumns = 17

const NumHookDataColumns = 16

const HookPk = "Id"

const HookDataColumnNames = "sha, dates_after, dates_before, category, created, deleted, forced, head_commit_id, head_commit_message, head_commit_timestamp, head_commit_author_name, head_commit_author_email, head_commit_author_username, head_commit_committer_name, head_commit_committer_email, head_commit_committer_username"

//--------------------------------------------------------------------------------
