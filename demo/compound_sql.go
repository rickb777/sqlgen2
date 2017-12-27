// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

package demo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/rickb777/sqlgen2"
	"github.com/rickb777/sqlgen2/where"
	"log"
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
	Dialect      sqlgen2.Dialect
	Logger       *log.Logger
}

// Type conformance check
var _ sqlgen2.Table = &DbCompoundTable{}

// NewDbCompoundTable returns a new table instance.
// If a blank table name is supplied, the default name "compounds" will be used instead.
// The table name prefix is initially blank and the request context is the background.
func NewDbCompoundTable(name string, d sqlgen2.Execer, dialect sqlgen2.Dialect) DbCompoundTable {
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
func (tbl DbCompoundTable) CreateTable(ifNotExist bool) (int64, error) {
	return tbl.Exec(tbl.createTableSql(ifNotExist))
}

func (tbl DbCompoundTable) createTableSql(ifNotExist bool) string {
	var stmt string
	switch tbl.Dialect {
	case sqlgen2.Sqlite: stmt = sqlCreateDbCompoundTableSqlite
    case sqlgen2.Postgres: stmt = sqlCreateDbCompoundTablePostgres
    case sqlgen2.Mysql: stmt = sqlCreateDbCompoundTableMysql
    }
	extra := tbl.ternary(ifNotExist, "IF NOT EXISTS ", "")
	query := fmt.Sprintf(stmt, extra, tbl.Prefix, tbl.Name)
	return query
}

func (tbl DbCompoundTable) ternary(flag bool, a, b string) string {
	if flag {
		return a
	}
	return b
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
 alpha    varchar(512),
 beta     varchar(512),
 category tinyint unsigned
)
`

const sqlCreateDbCompoundTableMysql = `
CREATE TABLE %s%s%s (
 alpha    varchar(512),
 beta     varchar(512),
 category tinyint unsigned
) ENGINE=InnoDB DEFAULT CHARSET=utf8
`

//--------------------------------------------------------------------------------

// CreateIndexes executes queries that create the indexes needed by the Compound table.
func (tbl DbCompoundTable) CreateIndexes(ifNotExist bool) (err error) {
	extra := tbl.ternary(ifNotExist, "IF NOT EXISTS ", "")
	_, err = tbl.Exec(tbl.createDbAlphaBetaIndexSql(extra))
	if err != nil {
		return err
	}

	return nil
}

func (tbl DbCompoundTable) prefixWithoutDot() string {
	last := len(tbl.Prefix)-1
	if last > 0 && tbl.Prefix[last] == '.' {
		return tbl.Prefix[0:last]
	}
	return tbl.Prefix
}

func (tbl DbCompoundTable) createDbAlphaBetaIndexSql(ifNotExist string) string {
	indexPrefix := tbl.prefixWithoutDot()
	return fmt.Sprintf(sqlCreateDbAlphaBetaIndex, ifNotExist, indexPrefix, tbl.Prefix, tbl.Name)
}

// CreateTableWithIndexes invokes CreateTable then CreateIndexes.
func (tbl DbCompoundTable) CreateTableWithIndexes(ifNotExist bool) (err error) {
	_, err = tbl.CreateTable(ifNotExist)
	if err != nil {
		return err
	}
	return tbl.CreateIndexes(ifNotExist)
}

//--------------------------------------------------------------------------------

const sqlCreateDbAlphaBetaIndex = `
CREATE UNIQUE INDEX %s%salpha_beta ON %s%s (alpha, beta)
`

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
func (tbl DbCompoundTable) QueryOne(query string, args ...interface{}) (*Compound, error) {
	tbl.logQuery(query, args...)
	row := tbl.Db.QueryRowContext(tbl.Ctx, query, args...)
	return scanDbCompound(row)
}

// Query is the low-level access function for Compounds.
func (tbl DbCompoundTable) Query(query string, args ...interface{}) ([]*Compound, error) {
	tbl.logQuery(query, args...)
	rows, err := tbl.Db.QueryContext(tbl.Ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanDbCompounds(rows)
}

//--------------------------------------------------------------------------------

// SelectOneSA allows a single Compound to be obtained from the table that match a 'where' clause and some limit.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
func (tbl DbCompoundTable) SelectOneSA(where, orderBy string, args ...interface{}) (*Compound, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s LIMIT 1", DbCompoundColumnNames, tbl.Prefix, tbl.Name, where, orderBy)
	return tbl.QueryOne(query, args...)
}

// SelectOne allows a single Compound to be obtained from the sqlgen2.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
func (tbl DbCompoundTable) SelectOne(where where.Expression, orderBy string) (*Compound, error) {
	wh, args := where.Build(tbl.Dialect)
	return tbl.SelectOneSA(wh, orderBy, args...)
}

// SelectSA allows Compounds to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
func (tbl DbCompoundTable) SelectSA(where, orderBy string, args ...interface{}) ([]*Compound, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", DbCompoundColumnNames, tbl.Prefix, tbl.Name, where, orderBy)
	return tbl.Query(query, args...)
}

// Select allows Compounds to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
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
	var stmt, params string
	switch tbl.Dialect {
	case sqlgen2.Postgres:
		stmt = sqlInsertDbCompoundPostgres
		params = sDbCompoundDataColumnParamsPostgres
	default:
		stmt = sqlInsertDbCompoundSimple
		params = sDbCompoundDataColumnParamsSimple
	}

	query := fmt.Sprintf(stmt, tbl.Prefix, tbl.Name, params)
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

const sqlInsertDbCompoundSimple = `
INSERT INTO %s%s (
	alpha, 
	beta, 
	category
) VALUES (%s)
`

const sqlInsertDbCompoundPostgres = `
INSERT INTO %s%s (
	alpha, 
	beta, 
	category
) VALUES (%s)
`

const sDbCompoundDataColumnParamsSimple = "?,?,?"

const sDbCompoundDataColumnParamsPostgres = "$1,$2,$3"

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

// scanDbCompound reads a table record into a single value.
func scanDbCompound(row *sql.Row) (*Compound, error) {
	var v0 string
	var v1 string
	var v2 Category

	err := row.Scan(
		&v0,
		&v1,
		&v2,

	)
	if err != nil {
		return nil, err
	}

	v := &Compound{}
	v.Alpha = v0
	v.Beta = v1
	v.Category = v2

	return v, nil
}

// scanDbCompounds reads table records into a slice of values.
func scanDbCompounds(rows *sql.Rows) ([]*Compound, error) {
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

		vv = append(vv, v)
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
