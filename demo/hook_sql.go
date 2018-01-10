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

// HookTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type HookTable struct {
	prefix, name string
	db           sqlgen2.Execer
	ctx          context.Context
	dialect      schema.Dialect
	logger       *log.Logger
}

// Type conformance check
var _ sqlgen2.TableCreator = &HookTable{}

// NewHookTable returns a new table instance.
// If a blank table name is supplied, the default name "hooks" will be used instead.
// The table name prefix is initially blank and the request context is the background.
func NewHookTable(name string, d sqlgen2.Execer, dialect schema.Dialect) HookTable {
	if name == "" {
		name = "hooks"
	}
	return HookTable{"", name, d, context.Background(), dialect, nil}
}

// CopyTableAsHookTable copies a table instance, retaining the name etc but
// providing methods appropriate for 'Hook'.
func CopyTableAsHookTable(origin sqlgen2.Table) HookTable {
	return HookTable{
		prefix:  origin.Prefix(),
		name:    origin.Name(),
		db:      origin.DB(),
		ctx:     origin.Ctx(),
		dialect: origin.Dialect(),
		logger:  origin.Logger(),
	}
}

// WithPrefix sets the table name prefix for subsequent queries.
func (tbl HookTable) WithPrefix(pfx string) HookTable {
	tbl.prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries.
func (tbl HookTable) WithContext(ctx context.Context) HookTable {
	tbl.ctx = ctx
	return tbl
}

// WithLogger sets the logger for subsequent queries.
func (tbl HookTable) WithLogger(logger *log.Logger) HookTable {
	tbl.logger = logger
	return tbl
}

// Ctx gets the current request context.
func (tbl HookTable) Ctx() context.Context {
	return tbl.ctx
}

// Dialect gets the database dialect.
func (tbl HookTable) Dialect() schema.Dialect {
	return tbl.dialect
}

// Logger gets the trace logger.
func (tbl HookTable) Logger() *log.Logger {
	return tbl.logger
}

// SetLogger sets the logger for subsequent queries, returning the interface.
func (tbl HookTable) SetLogger(logger *log.Logger) sqlgen2.Table {
	tbl.logger = logger
	return tbl
}

// Name gets the table name.
func (tbl HookTable) Name() string {
	return tbl.name
}

// Prefix gets the table name prefix.
func (tbl HookTable) Prefix() string {
	return tbl.prefix
}

// FullName gets the concatenated prefix and table name.
func (tbl HookTable) FullName() string {
	return tbl.prefix + tbl.name
}

func (tbl HookTable) prefixWithoutDot() string {
	last := len(tbl.prefix)-1
	if last > 0 && tbl.prefix[last] == '.' {
		return tbl.prefix[0:last]
	}
	return tbl.prefix
}

// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl HookTable) DB() *sql.DB {
	return tbl.db.(*sql.DB)
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl HookTable) Tx() *sql.Tx {
	return tbl.db.(*sql.Tx)
}

// IsTx tests whether this is within a transaction.
func (tbl HookTable) IsTx() bool {
	_, ok := tbl.db.(*sql.Tx)
	return ok
}

// Begin starts a transaction. The default isolation level is dependent on the driver.
func (tbl HookTable) BeginTx(opts *sql.TxOptions) (HookTable, error) {
	d := tbl.db.(*sql.DB)
	var err error
	tbl.db, err = d.BeginTx(tbl.ctx, opts)
	return tbl, err
}

func (tbl HookTable) logQuery(query string, args ...interface{}) {
	sqlgen2.LogQuery(tbl.logger, query, args...)
}


//--------------------------------------------------------------------------------

const NumHookColumns = 17

const NumHookDataColumns = 16

const HookPk = "Id"

const HookDataColumnNames = "sha, after, before, category, created, deleted, forced, commit_id, message, timestamp, head_commit_author_name, head_commit_author_email, head_commit_author_username, head_commit_committer_name, head_commit_committer_email, head_commit_committer_username"

//--------------------------------------------------------------------------------

// CreateTable creates the table.
func (tbl HookTable) CreateTable(ifNotExists bool) (int64, error) {
	return tbl.Exec(tbl.createTableSql(ifNotExists))
}

func (tbl HookTable) createTableSql(ifNotExists bool) string {
	var stmt string
	switch tbl.dialect {
	case schema.Sqlite: stmt = sqlCreateHookTableSqlite
    case schema.Postgres: stmt = sqlCreateHookTablePostgres
    case schema.Mysql: stmt = sqlCreateHookTableMysql
    }
	extra := tbl.ternary(ifNotExists, "IF NOT EXISTS ", "")
	query := fmt.Sprintf(stmt, extra, tbl.prefix, tbl.name)
	return query
}

func (tbl HookTable) ternary(flag bool, a, b string) string {
	if flag {
		return a
	}
	return b
}

// DropTable drops the table, destroying all its data.
func (tbl HookTable) DropTable(ifExists bool) (int64, error) {
	return tbl.Exec(tbl.dropTableSql(ifExists))
}

func (tbl HookTable) dropTableSql(ifExists bool) string {
	extra := tbl.ternary(ifExists, "IF EXISTS ", "")
	query := fmt.Sprintf("DROP TABLE %s%s%s", extra, tbl.prefix, tbl.name)
	return query
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

// Truncate drops every record from the table, if possible. It might fail if constraints exist that
// prevent some or all rows from being deleted; use the force option to override this.
//
// When 'force' is set true, be aware of the following consequences.
// When using Mysql, foreign keys in other tables can be left dangling.
// When using Postgres, a cascade happens, so all 'adjacent' tables (i.e. linked by foreign keys)
// are also truncated.
func (tbl HookTable) Truncate(force bool) (err error) {
	for _, query := range tbl.dialect.TruncateDDL(tbl.FullName(), force) {
		_, err = tbl.Exec(query)
		if err != nil {
			return err
		}
	}
	return nil
}

//--------------------------------------------------------------------------------

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// It returns the number of rows affected (of the database drive supports this).
func (tbl HookTable) Exec(query string, args ...interface{}) (int64, error) {
	tbl.logQuery(query, args...)
	res, err := tbl.db.ExecContext(tbl.ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

//--------------------------------------------------------------------------------

// QueryOne is the low-level access function for one Hook.
// If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, *Hook will be nil.
func (tbl HookTable) QueryOne(query string, args ...interface{}) (*Hook, error) {
	list, err := tbl.doQuery(true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

// Query is the low-level access function for Hooks.
func (tbl HookTable) Query(query string, args ...interface{}) (HookList, error) {
	return tbl.doQuery(false, query, args...)
}

func (tbl HookTable) doQuery(firstOnly bool, query string, args ...interface{}) (HookList, error) {
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanHooks(rows, firstOnly)
}

//--------------------------------------------------------------------------------

// GetHook gets the record with a given primary key value.
// If not found, *Hook will be nil.
func (tbl HookTable) GetHook(id int64) (*Hook, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s WHERE id=?", HookColumnNames, tbl.prefix, tbl.name)
	return tbl.QueryOne(query, id)
}

// GetHooks gets records from the table according to a list of primary keys.
// Although the list of ids can be arbitrarily long, there are practical limits;
// note that Oracle DB has a limit of 1000.
func (tbl HookTable) GetHooks(id ...int64) (list HookList, err error) {
	if len(id) > 0 {
		pl := tbl.dialect.Placeholders(len(id))
		query := fmt.Sprintf("SELECT %s FROM %s%s WHERE id IN (%s)", HookColumnNames, tbl.prefix, tbl.name, pl)
		args := make([]interface{}, len(id))

		for i, v := range id {
			args[i] = v
		}

		list, err = tbl.Query(query, args...)
	}

	return list, err
}

//--------------------------------------------------------------------------------

// SelectOneSA allows a single Example to be obtained from the table that match a 'where' clause
// and some limit. Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
// If not found, *Example will be nil.
func (tbl HookTable) SelectOneSA(where, orderBy string, args ...interface{}) (*Hook, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s LIMIT 1", HookColumnNames, tbl.prefix, tbl.name, where, orderBy)
	return tbl.QueryOne(query, args...)
}

// SelectOne allows a single Hook to be obtained from the sqlgen2.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
// If not found, *Example will be nil.
func (tbl HookTable) SelectOne(wh where.Expression, qc where.QueryConstraint) (*Hook, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	return tbl.SelectOneSA(whs, orderBy, args...)
}

// SelectSA allows Hooks to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
func (tbl HookTable) SelectSA(where, orderBy string, args ...interface{}) (HookList, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", HookColumnNames, tbl.prefix, tbl.name, where, orderBy)
	return tbl.Query(query, args...)
}

// Select allows Hooks to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) Select(wh where.Expression, qc where.QueryConstraint) (HookList, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	return tbl.SelectSA(whs, orderBy, args...)
}

// CountSA counts Hooks in the table that match a 'where' clause.
// Use a blank string for the 'where' argument if it is not needed.
func (tbl HookTable) CountSA(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s%s %s", tbl.prefix, tbl.name, where)
	tbl.logQuery(query, args...)
	row := tbl.db.QueryRowContext(tbl.ctx, query, args...)
	err = row.Scan(&count)
	return count, err
}

// Count counts the Hooks in the table that match a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed.
func (tbl HookTable) Count(wh where.Expression) (count int64, err error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	return tbl.CountSA(whs, args...)
}

const HookColumnNames = "id, sha, after, before, category, created, deleted, forced, commit_id, message, timestamp, head_commit_author_name, head_commit_author_email, head_commit_author_username, head_commit_committer_name, head_commit_committer_email, head_commit_committer_username"

//--------------------------------------------------------------------------------

// SliceId gets the Id column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceId(wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return tbl.getint64list("id", wh, qc)
}

// SliceSha gets the Sha column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceSha(wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.getstringlist("sha", wh, qc)
}

// SliceAfter gets the After column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceAfter(wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.getstringlist("after", wh, qc)
}

// SliceBefore gets the Before column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceBefore(wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.getstringlist("before", wh, qc)
}

// SliceCategory gets the Category column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceCategory(wh where.Expression, qc where.QueryConstraint) ([]Category, error) {
	return tbl.getCategorylist("category", wh, qc)
}

// SliceCreated gets the Created column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceCreated(wh where.Expression, qc where.QueryConstraint) ([]bool, error) {
	return tbl.getboollist("created", wh, qc)
}

// SliceDeleted gets the Deleted column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceDeleted(wh where.Expression, qc where.QueryConstraint) ([]bool, error) {
	return tbl.getboollist("deleted", wh, qc)
}

// SliceForced gets the Forced column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceForced(wh where.Expression, qc where.QueryConstraint) ([]bool, error) {
	return tbl.getboollist("forced", wh, qc)
}

// SliceID gets the ID column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceCommitId(wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.getstringlist("commit_id", wh, qc)
}

// SliceMessage gets the Message column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceMessage(wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.getstringlist("message", wh, qc)
}

// SliceTimestamp gets the Timestamp column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceTimestamp(wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.getstringlist("timestamp", wh, qc)
}

// SliceName gets the Name column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceHeadCommitAuthorName(wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.getstringlist("head_commit_author_name", wh, qc)
}

// SliceEmail gets the Email column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceHeadCommitAuthorEmail(wh where.Expression, qc where.QueryConstraint) ([]Email, error) {
	return tbl.getEmaillist("head_commit_author_email", wh, qc)
}

// SliceUsername gets the Username column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceHeadCommitAuthorUsername(wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.getstringlist("head_commit_author_username", wh, qc)
}

// SliceName gets the Name column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceHeadCommitCommitterName(wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.getstringlist("head_commit_committer_name", wh, qc)
}

// SliceEmail gets the Email column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceHeadCommitCommitterEmail(wh where.Expression, qc where.QueryConstraint) ([]Email, error) {
	return tbl.getEmaillist("head_commit_committer_email", wh, qc)
}

// SliceUsername gets the Username column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl HookTable) SliceHeadCommitCommitterUsername(wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.getstringlist("head_commit_committer_username", wh, qc)
}


func (tbl HookTable) getCategorylist(sqlname string, wh where.Expression, qc where.QueryConstraint) ([]Category, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", sqlname, tbl.prefix, tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
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

func (tbl HookTable) getEmaillist(sqlname string, wh where.Expression, qc where.QueryConstraint) ([]Email, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", sqlname, tbl.prefix, tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var v Email
	list := make([]Email, 0, 10)
	for rows.Next() {
		err = rows.Scan(&v)
		if err != nil {
			return list, err
		}
		list = append(list, v)
	}
	return list, nil
}

func (tbl HookTable) getboollist(sqlname string, wh where.Expression, qc where.QueryConstraint) ([]bool, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", sqlname, tbl.prefix, tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
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

func (tbl HookTable) getint64list(sqlname string, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", sqlname, tbl.prefix, tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
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

func (tbl HookTable) getstringlist(sqlname string, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", sqlname, tbl.prefix, tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
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


//--------------------------------------------------------------------------------

// Insert adds new records for the Hooks. The Hooks have their primary key fields
// set to the new record identifiers.
// The Hook.PreInsert(Execer) method will be called, if it exists.
func (tbl HookTable) Insert(vv ...*Hook) error {
	var params string
	switch tbl.dialect {
	case schema.Postgres:
		params = sHookDataColumnParamsPostgres
	default:
		params = sHookDataColumnParamsSimple
	}

	query := fmt.Sprintf(sqlInsertHook, tbl.prefix, tbl.name, params)
	st, err := tbl.db.PrepareContext(tbl.ctx, query)
	if err != nil {
		return err
	}
	defer st.Close()

	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPreInsert); ok {
			hook.PreInsert()
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
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl HookTable) UpdateFields(wh where.Expression, fields ...sql.NamedArg) (int64, error) {
	query, args := tbl.updateFields(wh, fields...)
	return tbl.Exec(query, args...)
}

func (tbl HookTable) updateFields(wh where.Expression, fields ...sql.NamedArg) (string, []interface{}) {
	list := sqlgen2.NamedArgList(fields)
	assignments := strings.Join(list.Assignments(tbl.dialect, 1), ", ")
	whs, wargs := where.BuildExpression(wh, tbl.dialect)
	query := fmt.Sprintf("UPDATE %s%s SET %s %s", tbl.prefix, tbl.name, assignments, whs)
	args := append(list.Values(), wargs...)
	return query, args
}

//--------------------------------------------------------------------------------

// Update updates records, matching them by primary key. It returns the number of rows affected.
// The Hook.PreUpdate(Execer) method will be called, if it exists.
func (tbl HookTable) Update(vv ...*Hook) (int64, error) {
	var stmt string
	switch tbl.dialect {
	case schema.Postgres:
		stmt = sqlUpdateHookByPkPostgres
	default:
		stmt = sqlUpdateHookByPkSimple
	}

	query := fmt.Sprintf(stmt, tbl.prefix, tbl.name)

	var count int64
	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPreUpdate); ok {
			hook.PreUpdate()
		}

		args, err := sliceHookWithoutPk(v)
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
UPDATE %s%s SET
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

//--------------------------------------------------------------------------------

// DeleteHooks deletes rows from the table, given some primary keys.
// The list of ids can be arbitrarily long.
func (tbl HookTable) DeleteHooks(id ...int64) (int64, error) {
	const batch = 1000 // limited by Oracle DB
	const qt = "DELETE FROM %s%s WHERE id IN (%s)"

	var count, n int64
	var err error
	var max = batch
	if len(id) < batch {
		max = len(id)
	}
	args := make([]interface{}, max)

	if len(id) > batch {
		pl := tbl.dialect.Placeholders(batch)
		query := fmt.Sprintf(qt, tbl.prefix, tbl.name, pl)

		for len(id) > batch {
			for i := 0; i < batch; i++ {
				args[i] = id[i]
			}

			n, err = tbl.Exec(query, args...)
			count += n
			if err != nil {
				return count, err
			}

			id = id[batch:]
		}
	}

	if len(id) > 0 {
		pl := tbl.dialect.Placeholders(len(id))
		query := fmt.Sprintf(qt, tbl.prefix, tbl.name, pl)

		for i := 0; i < batch; i++ {
			args[i] = id[i]
		}

		n, err = tbl.Exec(query, args...)
		count += n
	}

	return count, err
}

// Delete deletes one or more rows from the table, given a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl HookTable) Delete(wh where.Expression) (int64, error) {
	query, args := tbl.deleteRows(wh)
	return tbl.Exec(query, args...)
}

func (tbl HookTable) deleteRows(wh where.Expression) (string, []interface{}) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	query := fmt.Sprintf("DELETE FROM %s%s %s", tbl.prefix, tbl.name, whs)
	return query, args
}

//--------------------------------------------------------------------------------

// scanHooks reads table records into a slice of values.
func scanHooks(rows *sql.Rows, firstOnly bool) (HookList, error) {
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

		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPostGet); ok {
			err = hook.PostGet()
			if err != nil {
				return vv, err
			}
		}

		vv = append(vv, v)

		if firstOnly {
			return vv, rows.Err()
		}
	}

	return vv, rows.Err()
}
