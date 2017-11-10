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
	var v2 hook.go.Dates
	var v3 hook.go.Category
	var v4 bool
	var v5 bool
	var v6 bool
	var v7 *hook.go.Commit

	err := row.Scan(
		&v0,
		&v1,
		&v2,
		&v3,
		&v4,
		&v5,
		&v6,
		&v7,

	)
	if err != nil {
		return nil, err
	}

	v := &Hook{}

	return v, nil
}

// ScanHooks reads database records into a slice of values.
func ScanHooks(rows *sql.Rows) ([]*Hook, error) {
	var err error
	var vv []*Hook

	var v0 int64
	var v1 string
	var v2 hook.go.Dates
	var v3 hook.go.Category
	var v4 bool
	var v5 bool
	var v6 bool
	var v7 *hook.go.Commit

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

		)
		if err != nil {
			return vv, err
		}

		v := &Hook{}

		vv = append(vv, v)
	}
	return vv, rows.Err()
}

func SliceHook(v *Hook) []interface{} {
	var v0 int64
	var v1 string
	var v2 hook.go.Dates
	var v3 hook.go.Category
	var v4 bool
	var v5 bool
	var v6 bool
	var v7 *hook.go.Commit


	return []interface{}{
		v0,
		v1,
		v2,
		v3,
		v4,
		v5,
		v6,
		v7,

	}
}

func SliceHookWithoutPk(v *Hook) []interface{} {
	var v1 string
	var v2 hook.go.Dates
	var v3 hook.go.Category
	var v4 bool
	var v5 bool
	var v6 bool
	var v7 *hook.go.Commit


	return []interface{}{
		v1,
		v2,
		v3,
		v4,
		v5,
		v6,
		v7,

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

// SelectOneSA allows a single Hook to be obtained from the database that match a 'where' clause and some limit.
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

// SelectSA allows Hooks to be obtained from the database that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
func (tbl HookTable) SelectSA(where, orderBy string, args ...interface{}) ([]*Hook, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", HookColumnNames, tbl.Prefix, tbl.Name, where, orderBy)
	return tbl.Query(query, args...)
}

// Select allows Hooks to be obtained from the database that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
func (tbl HookTable) Select(where where.Expression, orderBy string) ([]*Hook, error) {
	wh, args := where.Build(tbl.Dialect)
	return tbl.SelectSA(wh, orderBy, args)
}

// CountSA counts Hooks in the database that match a 'where' clause.
func (tbl HookTable) CountSA(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s%s %s", tbl.Prefix, tbl.Name, where)
	row := tbl.Db.QueryRowContext(tbl.Ctx, query, args)
	err = row.Scan(&count)
	return count, err
}

// Count counts the Hooks in the database that match a 'where' clause.
func (tbl HookTable) Count(where where.Expression) (count int64, err error) {
	return tbl.CountSA(where.Build(tbl.Dialect))
}

const HookColumnNames = "id, sha, dates, category, created, deleted, forced, head_commit"

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
	dates,
	category,
	created,
	deleted,
	forced,
	head_commit
) VALUES (%s)
`

const sqlInsertHookPostgres = `
INSERT INTO %s%s (
	sha,
	dates,
	category,
	created,
	deleted,
	forced,
	head_commit
) VALUES (%s)
`

const sHookDataColumnParamsSimple = "?,?,?,?,?,?,?"

const sHookDataColumnParamsPostgres = "$1,$2,$3,$4,$5,$6,$7"

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
		args := SliceHookWithoutPk(v)
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
	dates=?,
	category=?,
	created=?,
	deleted=?,
	forced=?,
	head_commit=? 
 WHERE id=?
`

const sqlUpdateHookByPkPostgres = `
UPDATE %s%s SET 
	sha=$2,
	dates=$3,
	category=$4,
	created=$5,
	deleted=$6,
	forced=$7,
	head_commit=$8 
 WHERE id=$9
`

//--------------------------------------------------------------------------------

// DeleteFields deleted one or more rows, given a 'where' clause.
func (tbl HookTable) Delete(where where.Expression) (int64, error) {
	return tbl.Exec(tbl.deleteRows(where))
}

func (tbl HookTable) deleteRows(where where.Expression) (string, []interface{}) {
	whereClause, args := where.Build(tbl.Dialect)
	query := fmt.Sprintf("DELETE FROM %s%s %s", tbl.Prefix, tbl.Name, whereClause)
	return query, args
}

//--------------------------------------------------------------------------------

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
 id          integer primary key autoincrement,
 sha         text,
 dates       blob,
 category    integer,
 created     boolean,
 deleted     boolean,
 forced      boolean,
 head_commit blob
)
`

const sqlCreateHookTablePostgres = `
CREATE TABLE %s%s%s (
 id          bigserial primary key ,
 sha         varchar(512),
 dates       byteaa,
 category    integer,
 created     boolean,
 deleted     boolean,
 forced      boolean,
 head_commit byteaa
)
`

const sqlCreateHookTableMysql = `
CREATE TABLE %s%s%s (
 id          bigint primary key auto_increment,
 sha         varchar(512),
 dates       mediumblob,
 category    tinyint unsigned,
 created     tinyint(1),
 deleted     tinyint(1),
 forced      tinyint(1),
 head_commit mediumblob
) ENGINE=InnoDB DEFAULT CHARSET=utf8
`

//--------------------------------------------------------------------------------

//--------------------------------------------------------------------------------

const NumHookColumns = 8

const NumHookDataColumns = 7

const HookPk = "Id"

const HookDataColumnNames = "sha, dates, category, created, deleted, forced, head_commit"

//--------------------------------------------------------------------------------
