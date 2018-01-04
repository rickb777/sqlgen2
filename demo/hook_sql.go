// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

package demo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/rickb777/sqlgen2"
	"github.com/rickb777/sqlgen2/schema"
	"github.com/rickb777/sqlgen2/where"
	"log"
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
	Dialect      schema.Dialect
	Logger       *log.Logger
}

// Type conformance check
var _ sqlgen2.Table = &HookTable{}

// NewHookTable returns a new table instance.
// If a blank table name is supplied, the default name "hooks" will be used instead.
// The table name prefix is initially blank and the request context is the background.
func NewHookTable(name string, d sqlgen2.Execer, dialect schema.Dialect) HookTable {
	if name == "" {
		name = HookTableName
	}
	return HookTable{"", name, d, context.Background(), dialect, nil}
}

// WithPrefix sets the prefix for subsequent queries.
func (tbl HookTable) WithPrefix(pfx string) HookTable {
	tbl.Prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries.
func (tbl HookTable) WithContext(ctx context.Context) HookTable {
	tbl.Ctx = ctx
	return tbl
}

// WithLogger sets the logger for subsequent queries.
func (tbl HookTable) WithLogger(logger *log.Logger) HookTable {
	tbl.Logger = logger
	return tbl
}

// SetLogger sets the logger for subsequent queries, returning the interface.
func (tbl HookTable) SetLogger(logger *log.Logger) sqlgen2.Table {
	tbl.Logger = logger
	return tbl
}

// FullName gets the concatenated prefix and table name.
func (tbl HookTable) FullName() string {
	return tbl.Prefix + tbl.Name
}

func (tbl HookTable) prefixWithoutDot() string {
	last := len(tbl.Prefix)-1
	if last > 0 && tbl.Prefix[last] == '.' {
		return tbl.Prefix[0:last]
	}
	return tbl.Prefix
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

func (tbl HookTable) logQuery(query string, args ...interface{}) {
	sqlgen2.LogQuery(tbl.Logger, query, args...)
}


//--------------------------------------------------------------------------------

const NumHookColumns = 17

const NumHookDataColumns = 16

const HookPk = "Id"

const HookDataColumnNames = "sha, after, before, category, created, deleted, forced, commit_id, message, timestamp, head_commit_author_name, head_commit_author_email, head_commit_author_username, head_commit_committer_name, head_commit_committer_email, head_commit_committer_username"

//--------------------------------------------------------------------------------

// CreateTable creates the table.
func (tbl HookTable) CreateTable(ifNotExist bool) (int64, error) {
	return tbl.Exec(tbl.createTableSql(ifNotExist))
}

func (tbl HookTable) createTableSql(ifNotExist bool) string {
	var stmt string
	switch tbl.Dialect {
	case schema.Sqlite: stmt = sqlCreateHookTableSqlite
    case schema.Postgres: stmt = sqlCreateHookTablePostgres
    case schema.Mysql: stmt = sqlCreateHookTableMysql
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

const sqlCreateHookTableSqlite = `
CREATE TABLE %s%s%s (
 id                             integer primary key autoincrement,
 sha                            text,
 after                          text,
 before                         text,
 category                       tinyint unsigned,
 created                        boolean,
 deleted                        boolean,
 forced                         boolean,
 commit_id                      text,
 message                        text,
 timestamp                      text,
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
 id                             bigserial primary key,
 sha                            varchar(255),
 after                          varchar(20),
 before                         varchar(20),
 category                       tinyint unsigned,
 created                        boolean,
 deleted                        boolean,
 forced                         boolean,
 commit_id                      varchar(255),
 message                        varchar(255),
 timestamp                      varchar(255),
 head_commit_author_name        varchar(255),
 head_commit_author_email       varchar(255),
 head_commit_author_username    varchar(255),
 head_commit_committer_name     varchar(255),
 head_commit_committer_email    varchar(255),
 head_commit_committer_username varchar(255)
)
`

const sqlCreateHookTableMysql = `
CREATE TABLE %s%s%s (
 id                             bigint primary key auto_increment,
 sha                            varchar(255),
 after                          varchar(20),
 before                         varchar(20),
 category                       tinyint unsigned,
 created                        tinyint(1),
 deleted                        tinyint(1),
 forced                         tinyint(1),
 commit_id                      varchar(255),
 message                        varchar(255),
 timestamp                      varchar(255),
 head_commit_author_name        varchar(255),
 head_commit_author_email       varchar(255),
 head_commit_author_username    varchar(255),
 head_commit_committer_name     varchar(255),
 head_commit_committer_email    varchar(255),
 head_commit_committer_username varchar(255)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
`

//--------------------------------------------------------------------------------

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// It returns the number of rows affected (of the database drive supports this).
func (tbl HookTable) Exec(query string, args ...interface{}) (int64, error) {
	tbl.logQuery(query, args...)
	res, err := tbl.Db.ExecContext(tbl.Ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

//--------------------------------------------------------------------------------

// QueryOne is the low-level access function for one Hook.
func (tbl HookTable) QueryOne(query string, args ...interface{}) (*Hook, error) {
	tbl.logQuery(query, args...)
	row := tbl.Db.QueryRowContext(tbl.Ctx, query, args...)
	return scanHook(row)
}

// Query is the low-level access function for Hooks.
func (tbl HookTable) Query(query string, args ...interface{}) (HookList, error) {
	tbl.logQuery(query, args...)
	rows, err := tbl.Db.QueryContext(tbl.Ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanHooks(rows)
}

//--------------------------------------------------------------------------------

// GetHook gets the record with a given primary key value.
func (tbl HookTable) GetHook(id int64) (*Hook, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s WHERE id=?", HookColumnNames, tbl.Prefix, tbl.Name)
	return tbl.QueryOne(query, id)
}

//--------------------------------------------------------------------------------

// SliceId gets the Id column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
func (tbl HookTable) SliceId(where where.Expression, orderBy string) ([]int64, error) {
	return tbl.getint64list("id", where, orderBy)
}

// SliceSha gets the Sha column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
func (tbl HookTable) SliceSha(where where.Expression, orderBy string) ([]string, error) {
	return tbl.getstringlist("sha", where, orderBy)
}

// SliceCategory gets the Category column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
func (tbl HookTable) SliceCategory(where where.Expression, orderBy string) ([]Category, error) {
	return tbl.getCategorylist("category", where, orderBy)
}

// SliceCreated gets the Created column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
func (tbl HookTable) SliceCreated(where where.Expression, orderBy string) ([]bool, error) {
	return tbl.getboollist("created", where, orderBy)
}

// SliceDeleted gets the Deleted column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
func (tbl HookTable) SliceDeleted(where where.Expression, orderBy string) ([]bool, error) {
	return tbl.getboollist("deleted", where, orderBy)
}

// SliceForced gets the Forced column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
func (tbl HookTable) SliceForced(where where.Expression, orderBy string) ([]bool, error) {
	return tbl.getboollist("forced", where, orderBy)
}


func (tbl HookTable) getint64list(sqlname string, where where.Expression, orderBy string) ([]int64, error) {
	wh, args := where.Build(tbl.Dialect)
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", sqlname, tbl.Prefix, tbl.Name, wh, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.Db.QueryContext(tbl.Ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var v int64
	list := make([]int64, 0, 10)
	for rows.Next() {
		err = rows.Scan(&v)
		if err != nil {
			return list, err
		}
		list = append(list, v)
	}
	return list, nil
}

func (tbl HookTable) getstringlist(sqlname string, where where.Expression, orderBy string) ([]string, error) {
	wh, args := where.Build(tbl.Dialect)
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", sqlname, tbl.Prefix, tbl.Name, wh, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.Db.QueryContext(tbl.Ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var v string
	list := make([]string, 0, 10)
	for rows.Next() {
		err = rows.Scan(&v)
		if err != nil {
			return list, err
		}
		list = append(list, v)
	}
	return list, nil
}

func (tbl HookTable) getCategorylist(sqlname string, where where.Expression, orderBy string) ([]Category, error) {
	wh, args := where.Build(tbl.Dialect)
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", sqlname, tbl.Prefix, tbl.Name, wh, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.Db.QueryContext(tbl.Ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var v Category
	list := make([]Category, 0, 10)
	for rows.Next() {
		err = rows.Scan(&v)
		if err != nil {
			return list, err
		}
		list = append(list, v)
	}
	return list, nil
}

func (tbl HookTable) getboollist(sqlname string, where where.Expression, orderBy string) ([]bool, error) {
	wh, args := where.Build(tbl.Dialect)
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", sqlname, tbl.Prefix, tbl.Name, wh, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.Db.QueryContext(tbl.Ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var v bool
	list := make([]bool, 0, 10)
	for rows.Next() {
		err = rows.Scan(&v)
		if err != nil {
			return list, err
		}
		list = append(list, v)
	}
	return list, nil
}


//--------------------------------------------------------------------------------

// SelectOneSA allows a single Hook to be obtained from the table that match a 'where' clause
// and some limit.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
func (tbl HookTable) SelectOneSA(where, orderBy string, args ...interface{}) (*Hook, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s LIMIT 1", HookColumnNames, tbl.Prefix, tbl.Name, where, orderBy)
	return tbl.QueryOne(query, args...)
}

// SelectOne allows a single Hook to be obtained from the sqlgen2.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
func (tbl HookTable) SelectOne(where where.Expression, orderBy string) (*Hook, error) {
	wh, args := where.Build(tbl.Dialect)
	return tbl.SelectOneSA(wh, orderBy, args...)
}

// SelectSA allows Hooks to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
func (tbl HookTable) SelectSA(where, orderBy string, args ...interface{}) (HookList, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", HookColumnNames, tbl.Prefix, tbl.Name, where, orderBy)
	return tbl.Query(query, args...)
}

// Select allows Hooks to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
func (tbl HookTable) Select(where where.Expression, orderBy string) (HookList, error) {
	wh, args := where.Build(tbl.Dialect)
	return tbl.SelectSA(wh, orderBy, args...)
}

// CountSA counts Hooks in the table that match a 'where' clause.
func (tbl HookTable) CountSA(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s%s %s", tbl.Prefix, tbl.Name, where)
	tbl.logQuery(query, args...)
	row := tbl.Db.QueryRowContext(tbl.Ctx, query, args...)
	err = row.Scan(&count)
	return count, err
}

// Count counts the Hooks in the table that match a 'where' clause.
func (tbl HookTable) Count(where where.Expression) (count int64, err error) {
	wh, args := where.Build(tbl.Dialect)
	return tbl.CountSA(wh, args...)
}

const HookColumnNames = "id, sha, after, before, category, created, deleted, forced, commit_id, message, timestamp, head_commit_author_name, head_commit_author_email, head_commit_author_username, head_commit_committer_name, head_commit_committer_email, head_commit_committer_username"

//--------------------------------------------------------------------------------

// Insert adds new records for the Hooks. The Hooks have their primary key fields
// set to the new record identifiers.
// The Hook.PreInsert(Execer) method will be called, if it exists.
func (tbl HookTable) Insert(vv ...*Hook) error {
	var params string
	switch tbl.Dialect {
	case schema.Postgres:
		params = sHookDataColumnParamsPostgres
	default:
		params = sHookDataColumnParamsSimple
	}

	query := fmt.Sprintf(sqlInsertHook, tbl.Prefix, tbl.Name, params)
	st, err := tbl.Db.PrepareContext(tbl.Ctx, query)
	if err != nil {
		return err
	}
	defer st.Close()

	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPreInsert); ok {
			hook.PreInsert(tbl.Db)
		}

		fields, err := sliceHookWithoutPk(v)
		if err != nil {
			return err
		}

		tbl.logQuery(query, fields...)
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

const sqlInsertHook = `
INSERT INTO %s%s (
	sha,
	after,
	before,
	category,
	created,
	deleted,
	forced,
	commit_id,
	message,
	timestamp,
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
	query, args := tbl.updateFields(where, fields...)
	return tbl.Exec(query, args...)
}

func (tbl HookTable) updateFields(where where.Expression, fields ...sql.NamedArg) (string, []interface{}) {
	list := sqlgen2.NamedArgList(fields)
	assignments := strings.Join(list.Assignments(tbl.Dialect, 1), ", ")
	whereClause, wargs := where.Build(tbl.Dialect)
	query := fmt.Sprintf("UPDATE %s%s SET %s %s", tbl.Prefix, tbl.Name, assignments, whereClause)
	args := append(list.Values(), wargs...)
	return query, args
}

//--------------------------------------------------------------------------------

// Update updates records, matching them by primary key. It returns the number of rows affected.
// The Hook.PreUpdate(Execer) method will be called, if it exists.
func (tbl HookTable) Update(vv ...*Hook) (int64, error) {
	var stmt string
	switch tbl.Dialect {
	case schema.Postgres:
		stmt = sqlUpdateHookByPkPostgres
	default:
		stmt = sqlUpdateHookByPkSimple
	}

	query := fmt.Sprintf(stmt, tbl.Prefix, tbl.Name)

	var count int64
	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPreUpdate); ok {
			hook.PreUpdate(tbl.Db)
		}

		args, err := sliceHookWithoutPk(v)
		if err != nil {
			return count, err
		}

		args = append(args, v.Id)
		tbl.logQuery(query, args...)
		n, err := tbl.Exec(query, args...)
		if err != nil {
			return count, err
		}
		count += n
	}
	return count, nil
}

const sqlUpdateHookByPkSimple = `
UPDATE %%s%%s SET
	sha=?,
	after=?,
	before=?,
	category=?,
	created=?,
	deleted=?,
	forced=?,
	commit_id=?,
	message=?,
	timestamp=?,
	head_commit_author_name=?,
	head_commit_author_email=?,
	head_commit_author_username=?,
	head_commit_committer_name=?,
	head_commit_committer_email=?,
	head_commit_committer_username=?
WHERE id=?
`

const sqlUpdateHookByPkPostgres = `
UPDATE %%s%%s SET
	sha=$2,
	after=$3,
	before=$4,
	category=$5,
	created=$6,
	deleted=$7,
	forced=$8,
	commit_id=$9,
	message=$10,
	timestamp=$11,
	head_commit_author_name=$12,
	head_commit_author_email=$13,
	head_commit_author_username=$14,
	head_commit_committer_name=$15,
	head_commit_committer_email=$16,
	head_commit_committer_username=$17
WHERE id=$1
`

//--------------------------------------------------------------------------------

// Delete deletes one or more rows from the table, given a 'where' clause.
func (tbl HookTable) Delete(where where.Expression) (int64, error) {
	query, args := tbl.deleteRows(where)
	return tbl.Exec(query, args...)
}

func (tbl HookTable) deleteRows(where where.Expression) (string, []interface{}) {
	whereClause, args := where.Build(tbl.Dialect)
	query := fmt.Sprintf("DELETE FROM %s%s %s", tbl.Prefix, tbl.Name, whereClause)
	return query, args
}

// Truncate drops every record from the table, if possible. It might fail if constraints exist that
// prevent some or all rows from being deleted; use the force option to override this.
//
// When 'force' is set true, be aware of the following consequences.
// When using Mysql, foreign keys in other tables can be left dangling.
// When using Postgres, a cascade happens, so all 'adjacent' tables (i.e. linked by foreign keys)
// are also truncated.
func (tbl HookTable) Truncate(force bool) (err error) {
	for _, query := range tbl.Dialect.TruncateDDL(tbl.FullName(), force) {
		_, err = tbl.Exec(query)
		if err != nil {
			return err
		}
	}
	return nil
}

//--------------------------------------------------------------------------------

// scanHook reads a table record into a single value.
func scanHook(row *sql.Row) (*Hook, error) {
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
	var v12 Email
	var v13 string
	var v14 string
	var v15 Email
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

// scanHooks reads table records into a slice of values.
func scanHooks(rows *sql.Rows) (HookList, error) {
	var err error
	var vv HookList

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
	var v12 Email
	var v13 string
	var v14 string
	var v15 Email
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

func sliceHookWithoutPk(v *Hook) ([]interface{}, error) {


	return []interface{}{
		v.Sha,
		v.Dates.After,
		v.Dates.Before,
		v.Category,
		v.Created,
		v.Deleted,
		v.Forced,
		v.HeadCommit.ID,
		v.HeadCommit.Message,
		v.HeadCommit.Timestamp,
		v.HeadCommit.Author.Name,
		v.HeadCommit.Author.Email,
		v.HeadCommit.Author.Username,
		v.HeadCommit.Committer.Name,
		v.HeadCommit.Committer.Email,
		v.HeadCommit.Committer.Username,

	}, nil
}
