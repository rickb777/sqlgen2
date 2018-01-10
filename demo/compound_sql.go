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

// DbCompoundTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type DbCompoundTable struct {
	prefix, name string
	db           sqlgen2.Execer
	ctx          context.Context
	dialect      schema.Dialect
	logger       *log.Logger
}

// Type conformance check
var _ sqlgen2.TableWithIndexes = &DbCompoundTable{}

// NewDbCompoundTable returns a new table instance.
// If a blank table name is supplied, the default name "compounds" will be used instead.
// The table name prefix is initially blank and the request context is the background.
func NewDbCompoundTable(name string, d sqlgen2.Execer, dialect schema.Dialect) DbCompoundTable {
	if name == "" {
		name = "compounds"
	}
	return DbCompoundTable{"", name, d, context.Background(), dialect, nil}
}

// CopyTableAsDbCompoundTable copies a table instance, retaining the name etc but
// providing methods appropriate for 'Compound'.
func CopyTableAsDbCompoundTable(origin sqlgen2.Table) DbCompoundTable {
	return DbCompoundTable{
		prefix:  origin.Prefix(),
		name:    origin.Name(),
		db:      origin.DB(),
		ctx:     origin.Ctx(),
		dialect: origin.Dialect(),
		logger:  origin.Logger(),
	}
}

// WithPrefix sets the table name prefix for subsequent queries.
func (tbl DbCompoundTable) WithPrefix(pfx string) DbCompoundTable {
	tbl.prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries.
func (tbl DbCompoundTable) WithContext(ctx context.Context) DbCompoundTable {
	tbl.ctx = ctx
	return tbl
}

// WithLogger sets the logger for subsequent queries.
func (tbl DbCompoundTable) WithLogger(logger *log.Logger) DbCompoundTable {
	tbl.logger = logger
	return tbl
}

// Ctx gets the current request context.
func (tbl DbCompoundTable) Ctx() context.Context {
	return tbl.ctx
}

// Dialect gets the database dialect.
func (tbl DbCompoundTable) Dialect() schema.Dialect {
	return tbl.dialect
}

// Logger gets the trace logger.
func (tbl DbCompoundTable) Logger() *log.Logger {
	return tbl.logger
}

// SetLogger sets the logger for subsequent queries, returning the interface.
func (tbl DbCompoundTable) SetLogger(logger *log.Logger) sqlgen2.Table {
	tbl.logger = logger
	return tbl
}

// Name gets the table name.
func (tbl DbCompoundTable) Name() string {
	return tbl.name
}

// Prefix gets the table name prefix.
func (tbl DbCompoundTable) Prefix() string {
	return tbl.prefix
}

// FullName gets the concatenated prefix and table name.
func (tbl DbCompoundTable) FullName() string {
	return tbl.prefix + tbl.name
}

func (tbl DbCompoundTable) prefixWithoutDot() string {
	last := len(tbl.prefix) - 1
	if last > 0 && tbl.prefix[last] == '.' {
		return tbl.prefix[0:last]
	}
	return tbl.prefix
}

// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl DbCompoundTable) DB() *sql.DB {
	return tbl.db.(*sql.DB)
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl DbCompoundTable) Tx() *sql.Tx {
	return tbl.db.(*sql.Tx)
}

// IsTx tests whether this is within a transaction.
func (tbl DbCompoundTable) IsTx() bool {
	_, ok := tbl.db.(*sql.Tx)
	return ok
}

// Begin starts a transaction. The default isolation level is dependent on the driver.
func (tbl DbCompoundTable) BeginTx(opts *sql.TxOptions) (DbCompoundTable, error) {
	d := tbl.db.(*sql.DB)
	var err error
	tbl.db, err = d.BeginTx(tbl.ctx, opts)
	return tbl, err
}

func (tbl DbCompoundTable) logQuery(query string, args ...interface{}) {
	sqlgen2.LogQuery(tbl.logger, query, args...)
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
	switch tbl.dialect {
	case schema.Sqlite:
		stmt = sqlCreateDbCompoundTableSqlite
	case schema.Postgres:
		stmt = sqlCreateDbCompoundTablePostgres
	case schema.Mysql:
		stmt = sqlCreateDbCompoundTableMysql
	}
	extra := tbl.ternary(ifNotExists, "IF NOT EXISTS ", "")
	query := fmt.Sprintf(stmt, extra, tbl.prefix, tbl.name)
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
	query := fmt.Sprintf("DROP TABLE %s%s%s", extra, tbl.prefix, tbl.name)
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
	ine := tbl.ternary(ifNotExist && tbl.dialect != schema.Mysql, "IF NOT EXISTS ", "")

	// Mysql does not support 'if not exists' on indexes
	// Workaround: use DropIndex first and ignore an error returned if the index didn't exist.

	if ifNotExist && tbl.dialect == schema.Mysql {
		tbl.DropAlphaBetaIndex(false)
		ine = ""
	}

	_, err := tbl.Exec(tbl.createDbAlphaBetaIndexSql(ine))
	return err
}

func (tbl DbCompoundTable) createDbAlphaBetaIndexSql(ifNotExists string) string {
	indexPrefix := tbl.prefixWithoutDot()
	return fmt.Sprintf("CREATE UNIQUE INDEX %s%salpha_beta ON %s%s (%s)", ifNotExists, indexPrefix,
		tbl.prefix, tbl.name, sqlDbAlphaBetaIndexColumns)
}

// DropAlphaBetaIndex drops the alpha_beta index.
func (tbl DbCompoundTable) DropAlphaBetaIndex(ifExists bool) error {
	_, err := tbl.Exec(tbl.dropDbAlphaBetaIndexSql(ifExists))
	return err
}

func (tbl DbCompoundTable) dropDbAlphaBetaIndexSql(ifExists bool) string {
	// Mysql does not support 'if exists' on indexes
	ie := tbl.ternary(ifExists && tbl.dialect != schema.Mysql, "IF EXISTS ", "")
	onTbl := tbl.ternary(tbl.dialect == schema.Mysql, fmt.Sprintf(" ON %s%s", tbl.prefix, tbl.name), "")
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

// Truncate drops every record from the table, if possible. It might fail if constraints exist that
// prevent some or all rows from being deleted; use the force option to override this.
//
// When 'force' is set true, be aware of the following consequences.
// When using Mysql, foreign keys in other tables can be left dangling.
// When using Postgres, a cascade happens, so all 'adjacent' tables (i.e. linked by foreign keys)
// are also truncated.
func (tbl DbCompoundTable) Truncate(force bool) (err error) {
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
func (tbl DbCompoundTable) Exec(query string, args ...interface{}) (int64, error) {
	tbl.logQuery(query, args...)
	res, err := tbl.db.ExecContext(tbl.ctx, query, args...)
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
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanDbCompounds(rows, firstOnly)
}

//--------------------------------------------------------------------------------

// SelectOneSA allows a single Compound to be obtained from the table that match a 'where' clause
// and some limit.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
// If not found, *Compound will be nil.
func (tbl DbCompoundTable) SelectOneSA(where, orderBy string, args ...interface{}) (*Compound, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s LIMIT 1", DbCompoundColumnNames, tbl.prefix, tbl.name, where, orderBy)
	return tbl.QueryOne(query, args...)
}

// SelectOne allows a single Compound to be obtained from the sqlgen2.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'; otherwise use nil.
// If not found, *Example will be nil.
func (tbl DbCompoundTable) SelectOne(wh where.Expression, qc where.QueryConstraint) (*Compound, error) {
	whs, args := wh.Build(tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	return tbl.SelectOneSA(whs, orderBy, args...)
}

// SelectSA allows Compounds to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
func (tbl DbCompoundTable) SelectSA(where, orderBy string, args ...interface{}) ([]*Compound, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", DbCompoundColumnNames, tbl.prefix, tbl.name, where, orderBy)
	return tbl.Query(query, args...)
}

// Select allows Compounds to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'; otherwise use nil.
func (tbl DbCompoundTable) Select(wh where.Expression, qc where.QueryConstraint) ([]*Compound, error) {
	whs, args := wh.Build(tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	return tbl.SelectSA(whs, orderBy, args...)
}

// CountSA counts Compounds in the table that match a 'where' clause.
func (tbl DbCompoundTable) CountSA(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s%s %s", tbl.prefix, tbl.name, where)
	tbl.logQuery(query, args...)
	row := tbl.db.QueryRowContext(tbl.ctx, query, args...)
	err = row.Scan(&count)
	return count, err
}

// Count counts the Compounds in the table that match a 'where' clause.
func (tbl DbCompoundTable) Count(where where.Expression) (count int64, err error) {
	wh, args := where.Build(tbl.dialect)
	return tbl.CountSA(wh, args...)
}

const DbCompoundColumnNames = "alpha, beta, category"

//--------------------------------------------------------------------------------

// SliceAlpha gets the Alpha column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'; otherwise use nil.
func (tbl DbCompoundTable) SliceAlpha(wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.getstringlist("alpha", wh, qc)
}

// SliceBeta gets the Beta column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'; otherwise use nil.
func (tbl DbCompoundTable) SliceBeta(wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.getstringlist("beta", wh, qc)
}

// SliceCategory gets the Category column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'; otherwise use nil.
func (tbl DbCompoundTable) SliceCategory(wh where.Expression, qc where.QueryConstraint) ([]Category, error) {
	return tbl.getCategorylist("category", wh, qc)
}

func (tbl DbCompoundTable) getCategorylist(sqlname string, wh where.Expression, qc where.QueryConstraint) ([]Category, error) {
	whs, args := wh.Build(tbl.dialect)
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

func (tbl DbCompoundTable) getstringlist(sqlname string, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	whs, args := wh.Build(tbl.dialect)
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

// Insert adds new records for the Compounds.
// The Compound.PreInsert(Execer) method will be called, if it exists.
func (tbl DbCompoundTable) Insert(vv ...*Compound) error {
	var params string
	switch tbl.dialect {
	case schema.Postgres:
		params = sDbCompoundDataColumnParamsPostgres
	default:
		params = sDbCompoundDataColumnParamsSimple
	}

	query := fmt.Sprintf(sqlInsertDbCompound, tbl.prefix, tbl.name, params)
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
	assignments := strings.Join(list.Assignments(tbl.dialect, 1), ", ")
	whereClause, wargs := where.Build(tbl.dialect)
	query := fmt.Sprintf("UPDATE %s%s SET %s %s", tbl.prefix, tbl.name, assignments, whereClause)
	args := append(list.Values(), wargs...)
	return query, args
}

func sliceDbCompound(v *Compound) ([]interface{}, error) {

	return []interface{}{
		v.Alpha,
		v.Beta,
		v.Category,
	}, nil
}

//--------------------------------------------------------------------------------

// Delete deletes one or more rows from the table, given a 'where' clause.
func (tbl DbCompoundTable) Delete(where where.Expression) (int64, error) {
	query, args := tbl.deleteRows(where)
	return tbl.Exec(query, args...)
}

func (tbl DbCompoundTable) deleteRows(where where.Expression) (string, []interface{}) {
	whereClause, args := where.Build(tbl.dialect)
	query := fmt.Sprintf("DELETE FROM %s%s %s", tbl.prefix, tbl.name, whereClause)
	return query, args
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
