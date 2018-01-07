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

// DbCompoundTableName is the default name for this table.
const DbCompoundTableName = "compounds"

// DbCompoundTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type DbCompoundTable struct {
	Prefix, Name string
	Db           sqlgen2.Execer
	Ctx          context.Context
	Dialect      schema.Dialect
	Logger       *log.Logger
}

// Type conformance check
var _ sqlgen2.Table = &DbCompoundTable{}

// NewDbCompoundTable returns a new table instance.
// If a blank table name is supplied, the default name "compounds" will be used instead.
// The table name prefix is initially blank and the request context is the background.
func NewDbCompoundTable(name string, d sqlgen2.Execer, dialect schema.Dialect) DbCompoundTable {
	if name == "" {
		name = DbCompoundTableName
	}
	return DbCompoundTable{"", name, d, context.Background(), dialect, nil}
}

// WithPrefix sets the prefix for subsequent queries.
func (tbl DbCompoundTable) WithPrefix(pfx string) DbCompoundTable {
	tbl.Prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries.
func (tbl DbCompoundTable) WithContext(ctx context.Context) DbCompoundTable {
	tbl.Ctx = ctx
	return tbl
}

// WithLogger sets the logger for subsequent queries.
func (tbl DbCompoundTable) WithLogger(logger *log.Logger) DbCompoundTable {
	tbl.Logger = logger
	return tbl
}

// SetLogger sets the logger for subsequent queries, returning the interface.
func (tbl DbCompoundTable) SetLogger(logger *log.Logger) sqlgen2.Table {
	tbl.Logger = logger
	return tbl
}

// FullName gets the concatenated prefix and table name.
func (tbl DbCompoundTable) FullName() string {
	return tbl.Prefix + tbl.Name
}

func (tbl DbCompoundTable) prefixWithoutDot() string {
	last := len(tbl.Prefix)-1
	if last > 0 && tbl.Prefix[last] == '.' {
		return tbl.Prefix[0:last]
	}
	return tbl.Prefix
}

// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl DbCompoundTable) DB() *sql.DB {
	return tbl.Db.(*sql.DB)
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl DbCompoundTable) Tx() *sql.Tx {
	return tbl.Db.(*sql.Tx)
}

// IsTx tests whether this is within a transaction.
func (tbl DbCompoundTable) IsTx() bool {
	_, ok := tbl.Db.(*sql.Tx)
	return ok
}

// Begin starts a transaction. The default isolation level is dependent on the driver.
func (tbl DbCompoundTable) BeginTx(opts *sql.TxOptions) (DbCompoundTable, error) {
	d := tbl.Db.(*sql.DB)
	var err error
	tbl.Db, err = d.BeginTx(tbl.Ctx, opts)
	return tbl, err
}

func (tbl DbCompoundTable) logQuery(query string, args ...interface{}) {
	sqlgen2.LogQuery(tbl.Logger, query, args...)
}


//--------------------------------------------------------------------------------

const NumDbCompoundColumns = 3

const NumDbCompoundDataColumns = 3

const DbCompoundDataColumnNames = "alpha, beta, category"

//--------------------------------------------------------------------------------

// CreateTable creates the table.
func (tbl DbCompoundTable) CreateTable(ifNotExists bool) (int64, error) {
	return tbl.Exec(tbl.createTableSql(ifNotExists))
}

func (tbl DbCompoundTable) createTableSql(ifNotExists bool) string {
	var stmt string
	switch tbl.Dialect {
	case schema.Sqlite: stmt = sqlCreateDbCompoundTableSqlite
    case schema.Postgres: stmt = sqlCreateDbCompoundTablePostgres
    case schema.Mysql: stmt = sqlCreateDbCompoundTableMysql
    }
	extra := tbl.ternary(ifNotExists, "IF NOT EXISTS ", "")
	query := fmt.Sprintf(stmt, extra, tbl.Prefix, tbl.Name)
	return query
}

func (tbl DbCompoundTable) ternary(flag bool, a, b string) string {
	if flag {
		return a
	}
	return b
}

// DropTable drops the table, destroying all its data.
func (tbl DbCompoundTable) DropTable(ifExists bool) (int64, error) {
	return tbl.Exec(tbl.dropTableSql(ifExists))
}

func (tbl DbCompoundTable) dropTableSql(ifExists bool) string {
	extra := tbl.ternary(ifExists, "IF EXISTS ", "")
	query := fmt.Sprintf("DROP TABLE %s%s%s", extra, tbl.Prefix, tbl.Name)
	return query
}

const sqlCreateDbCompoundTableSqlite = `
CREATE TABLE %s%s%s (
 alpha    text,
 beta     text,
 category tinyint unsigned
)
`

const sqlCreateDbCompoundTablePostgres = `
CREATE TABLE %s%s%s (
 alpha    varchar(255),
 beta     varchar(255),
 category tinyint unsigned
)
`

const sqlCreateDbCompoundTableMysql = `
CREATE TABLE %s%s%s (
 alpha    varchar(255),
 beta     varchar(255),
 category tinyint unsigned
) ENGINE=InnoDB DEFAULT CHARSET=utf8
`

//--------------------------------------------------------------------------------

// CreateTableWithIndexes invokes CreateTable then CreateIndexes.
func (tbl DbCompoundTable) CreateTableWithIndexes(ifNotExist bool) (err error) {
	_, err = tbl.CreateTable(ifNotExist)
	if err != nil {
		return err
	}

	return tbl.CreateIndexes(ifNotExist)
}

// CreateIndexes executes queries that create the indexes needed by the Compound table.
func (tbl DbCompoundTable) CreateIndexes(ifNotExist bool) (err error) {

	err = tbl.CreateAlphaBetaIndex(ifNotExist)
	if err != nil {
		return err
	}

	return nil
}

// CreateAlphaBetaIndex creates the alpha_beta index.
func (tbl DbCompoundTable) CreateAlphaBetaIndex(ifNotExist bool) error {
	ine := tbl.ternary(ifNotExist && tbl.Dialect != schema.Mysql, "IF NOT EXISTS ", "")

	// Mysql does not support 'if not exists' on indexes
	// Workaround: use DropIndex first and ignore an error returned if the index didn't exist.

	if ifNotExist && tbl.Dialect == schema.Mysql {
		tbl.DropAlphaBetaIndex(false)
		ine = ""
	}

	_, err := tbl.Exec(tbl.createDbAlphaBetaIndexSql(ine))
	return err
}

func (tbl DbCompoundTable) createDbAlphaBetaIndexSql(ifNotExists string) string {
	indexPrefix := tbl.prefixWithoutDot()
	return fmt.Sprintf("CREATE UNIQUE INDEX %s%salpha_beta ON %s%s (%s)", ifNotExists, indexPrefix,
		tbl.Prefix, tbl.Name, sqlDbAlphaBetaIndexColumns)
}

// DropAlphaBetaIndex drops the alpha_beta index.
func (tbl DbCompoundTable) DropAlphaBetaIndex(ifExists bool) error {
	_, err := tbl.Exec(tbl.dropDbAlphaBetaIndexSql(ifExists))
	return err
}

func (tbl DbCompoundTable) dropDbAlphaBetaIndexSql(ifExists bool) string {
	// Mysql does not support 'if exists' on indexes
	ie := tbl.ternary(ifExists && tbl.Dialect != schema.Mysql, "IF EXISTS ", "")
	onTbl := tbl.ternary(tbl.Dialect == schema.Mysql, fmt.Sprintf(" ON %s%s", tbl.Prefix, tbl.Name), "")
	indexPrefix := tbl.prefixWithoutDot()
	return fmt.Sprintf("DROP INDEX %s%salpha_beta%s", ie, indexPrefix, onTbl)
}

// DropIndexes executes queries that drop the indexes on by the Compound table.
func (tbl DbCompoundTable) DropIndexes(ifExist bool) (err error) {

	err = tbl.DropAlphaBetaIndex(ifExist)
	if err != nil {
		return err
	}

	return nil
}

//--------------------------------------------------------------------------------

const sqlDbAlphaBetaIndexColumns = "alpha, beta"

//--------------------------------------------------------------------------------

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// It returns the number of rows affected (of the database drive supports this).
func (tbl DbCompoundTable) Exec(query string, args ...interface{}) (int64, error) {
	tbl.logQuery(query, args...)
	res, err := tbl.Db.ExecContext(tbl.Ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

//--------------------------------------------------------------------------------

// QueryOne is the low-level access function for one Compound.
// If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, *Compound will be nil.
func (tbl DbCompoundTable) QueryOne(query string, args ...interface{}) (*Compound, error) {
	list, err := tbl.doQuery(true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

// Query is the low-level access function for Compounds.
func (tbl DbCompoundTable) Query(query string, args ...interface{}) ([]*Compound, error) {
	return tbl.doQuery(false, query, args...)
}

func (tbl DbCompoundTable) doQuery(firstOnly bool, query string, args ...interface{}) ([]*Compound, error) {
	tbl.logQuery(query, args...)
	rows, err := tbl.Db.QueryContext(tbl.Ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanDbCompounds(rows, firstOnly)
}

//--------------------------------------------------------------------------------

// SliceAlpha gets the Alpha column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
func (tbl DbCompoundTable) SliceAlpha(where where.Expression, orderBy string) ([]string, error) {
	return tbl.getstringlist("alpha", where, orderBy)
}

// SliceBeta gets the Beta column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
func (tbl DbCompoundTable) SliceBeta(where where.Expression, orderBy string) ([]string, error) {
	return tbl.getstringlist("beta", where, orderBy)
}

// SliceCategory gets the Category column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
func (tbl DbCompoundTable) SliceCategory(where where.Expression, orderBy string) ([]Category, error) {
	return tbl.getCategorylist("category", where, orderBy)
}


func (tbl DbCompoundTable) getCategorylist(sqlname string, where where.Expression, orderBy string) ([]Category, error) {
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

func (tbl DbCompoundTable) getstringlist(sqlname string, where where.Expression, orderBy string) ([]string, error) {
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


//--------------------------------------------------------------------------------

// SelectOneSA allows a single Compound to be obtained from the table that match a 'where' clause
// and some limit.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
// If not found, *Compound will be nil.
func (tbl DbCompoundTable) SelectOneSA(where, orderBy string, args ...interface{}) (*Compound, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s LIMIT 1", DbCompoundColumnNames, tbl.Prefix, tbl.Name, where, orderBy)
	return tbl.QueryOne(query, args...)
}

// SelectOne allows a single Compound to be obtained from the sqlgen2.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
// If not found, *Example will be nil.
func (tbl DbCompoundTable) SelectOne(where where.Expression, orderBy string) (*Compound, error) {
	wh, args := where.Build(tbl.Dialect)
	return tbl.SelectOneSA(wh, orderBy, args...)
}

// SelectSA allows Compounds to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
func (tbl DbCompoundTable) SelectSA(where, orderBy string, args ...interface{}) ([]*Compound, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", DbCompoundColumnNames, tbl.Prefix, tbl.Name, where, orderBy)
	return tbl.Query(query, args...)
}

// Select allows Compounds to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
func (tbl DbCompoundTable) Select(where where.Expression, orderBy string) ([]*Compound, error) {
	wh, args := where.Build(tbl.Dialect)
	return tbl.SelectSA(wh, orderBy, args...)
}

// CountSA counts Compounds in the table that match a 'where' clause.
func (tbl DbCompoundTable) CountSA(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s%s %s", tbl.Prefix, tbl.Name, where)
	tbl.logQuery(query, args...)
	row := tbl.Db.QueryRowContext(tbl.Ctx, query, args...)
	err = row.Scan(&count)
	return count, err
}

// Count counts the Compounds in the table that match a 'where' clause.
func (tbl DbCompoundTable) Count(where where.Expression) (count int64, err error) {
	wh, args := where.Build(tbl.Dialect)
	return tbl.CountSA(wh, args...)
}

const DbCompoundColumnNames = "alpha, beta, category"

//--------------------------------------------------------------------------------

// Insert adds new records for the Compounds.
// The Compound.PreInsert(Execer) method will be called, if it exists.
func (tbl DbCompoundTable) Insert(vv ...*Compound) error {
	var params string
	switch tbl.Dialect {
	case schema.Postgres:
		params = sDbCompoundDataColumnParamsPostgres
	default:
		params = sDbCompoundDataColumnParamsSimple
	}

	query := fmt.Sprintf(sqlInsertDbCompound, tbl.Prefix, tbl.Name, params)
	st, err := tbl.Db.PrepareContext(tbl.Ctx, query)
	if err != nil {
		return err
	}
	defer st.Close()

	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPreInsert); ok {
			hook.PreInsert()
		}

		fields, err := sliceDbCompound(v)
		if err != nil {
			return err
		}

		tbl.logQuery(query, fields...)
		_, err = st.Exec(fields...)
		if err != nil {
			return err
		}
	}

	return nil
}

const sqlInsertDbCompound = `
INSERT INTO %s%s (
	alpha,
	beta,
	category
) VALUES (%s)
`

const sDbCompoundDataColumnParamsSimple = "?,?,?"

const sDbCompoundDataColumnParamsPostgres = "$1,$2,$3"

//--------------------------------------------------------------------------------

// UpdateFields updates one or more columns, given a 'where' clause.
func (tbl DbCompoundTable) UpdateFields(where where.Expression, fields ...sql.NamedArg) (int64, error) {
	query, args := tbl.updateFields(where, fields...)
	return tbl.Exec(query, args...)
}

func (tbl DbCompoundTable) updateFields(where where.Expression, fields ...sql.NamedArg) (string, []interface{}) {
	list := sqlgen2.NamedArgList(fields)
	assignments := strings.Join(list.Assignments(tbl.Dialect, 1), ", ")
	whereClause, wargs := where.Build(tbl.Dialect)
	query := fmt.Sprintf("UPDATE %s%s SET %s %s", tbl.Prefix, tbl.Name, assignments, whereClause)
	args := append(list.Values(), wargs...)
	return query, args
}

//--------------------------------------------------------------------------------

// Delete deletes one or more rows from the table, given a 'where' clause.
func (tbl DbCompoundTable) Delete(where where.Expression) (int64, error) {
	query, args := tbl.deleteRows(where)
	return tbl.Exec(query, args...)
}

func (tbl DbCompoundTable) deleteRows(where where.Expression) (string, []interface{}) {
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
func (tbl DbCompoundTable) Truncate(force bool) (err error) {
	for _, query := range tbl.Dialect.TruncateDDL(tbl.FullName(), force) {
		_, err = tbl.Exec(query)
		if err != nil {
			return err
		}
	}
	return nil
}

//--------------------------------------------------------------------------------

// scanDbCompounds reads table records into a slice of values.
func scanDbCompounds(rows *sql.Rows, firstOnly bool) ([]*Compound, error) {
	var err error
	var vv []*Compound

	var v0 string
	var v1 string
	var v2 Category

	for rows.Next() {
		err = rows.Scan(
			&v0,
			&v1,
			&v2,
		)
		if err != nil {
			return vv, err
		}

		v := &Compound{}
		v.Alpha = v0
		v.Beta = v1
		v.Category = v2

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

func sliceDbCompound(v *Compound) ([]interface{}, error) {


	return []interface{}{
		v.Alpha,
		v.Beta,
		v.Category,

	}, nil
}

//--------------------------------------------------------------------------------
