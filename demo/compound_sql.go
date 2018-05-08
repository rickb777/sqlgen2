// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

package demo

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"github.com/rickb777/sqlgen2"
	"github.com/rickb777/sqlgen2/constraint"
	"github.com/rickb777/sqlgen2/require"
	"github.com/rickb777/sqlgen2/schema"
	"github.com/rickb777/sqlgen2/support"
	"github.com/rickb777/sqlgen2/where"
	"io"
	"log"
)

// DbCompoundTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type DbCompoundTable struct {
	name        sqlgen2.TableName
	database    *sqlgen2.Database
	db          sqlgen2.Execer
	constraints constraint.Constraints
	ctx         context.Context
	pk          string
}

// Type conformance checks
var _ sqlgen2.TableWithIndexes = &DbCompoundTable{}
var _ sqlgen2.TableWithCrud = &DbCompoundTable{}

// NewDbCompoundTable returns a new table instance.
// If a blank table name is supplied, the default name "compounds" will be used instead.
// The request context is initialised with the background.
func NewDbCompoundTable(name string, d *sqlgen2.Database) DbCompoundTable {
	if name == "" {
		name = "compounds"
	}
	var constraints constraint.Constraints
	return DbCompoundTable{
		name:        sqlgen2.TableName{"", name},
		database:    d,
		db:          d.DB(),
		constraints: constraints,
		ctx:         context.Background(),
		pk:          "",
	}
}

// CopyTableAsDbCompoundTable copies a table instance, retaining the name etc but
// providing methods appropriate for 'Compound'. It doesn't copy the constraints of the original table.
//
// It serves to provide methods appropriate for 'Compound'. This is most useful when this is used to represent a
// join result. In such cases, there won't be any need for DDL methods, nor Exec, Insert, Update or Delete.
func CopyTableAsDbCompoundTable(origin sqlgen2.Table) DbCompoundTable {
	return DbCompoundTable{
		name:        origin.Name(),
		database:    origin.Database(),
		db:          origin.DB(),
		constraints: nil,
		ctx:         context.Background(),
		pk:          "",
	}
}

// WithPrefix sets the table name prefix for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl DbCompoundTable) WithPrefix(pfx string) DbCompoundTable {
	tbl.name.Prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries via this table.
// The result is a modified copy of the table; the original is unchanged.
//
// The shared context in the *Database is not altered by this method. So it
// is possible to use different contexts for different (groups of) queries.
func (tbl DbCompoundTable) WithContext(ctx context.Context) DbCompoundTable {
	tbl.ctx = ctx
	return tbl
}

// Database gets the shared database information.
func (tbl DbCompoundTable) Database() *sqlgen2.Database {
	return tbl.database
}

// Logger gets the trace logger.
func (tbl DbCompoundTable) Logger() *log.Logger {
	return tbl.database.Logger()
}

// WithConstraint returns a modified Table with added data consistency constraints.
func (tbl DbCompoundTable) WithConstraint(cc ...constraint.Constraint) DbCompoundTable {
	tbl.constraints = append(tbl.constraints, cc...)
	return tbl
}

// Constraints returns the table's constraints.
func (tbl DbCompoundTable) Constraints() constraint.Constraints {
	return tbl.constraints
}

// Ctx gets the current request context.
func (tbl DbCompoundTable) Ctx() context.Context {
	return tbl.ctx
}

// Dialect gets the database dialect.
func (tbl DbCompoundTable) Dialect() schema.Dialect {
	return tbl.database.Dialect()
}

// Name gets the table name.
func (tbl DbCompoundTable) Name() sqlgen2.TableName {
	return tbl.name
}

// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl DbCompoundTable) DB() *sql.DB {
	return tbl.db.(*sql.DB)
}

// Execer gets the wrapped database or transaction handle.
func (tbl DbCompoundTable) Execer() sqlgen2.Execer {
	return tbl.db
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

// BeginTx starts a transaction using the table's context.
// This context is used until the transaction is committed or rolled back.
//
// If this context is cancelled, the sql package will roll back the transaction.
// In this case, Tx.Commit will then return an error.
//
// The provided TxOptions is optional and may be nil if defaults should be used.
// If a non-default isolation level is used that the driver doesn't support,
// an error will be returned.
//
// Panics if the Execer is not TxStarter.
func (tbl DbCompoundTable) BeginTx(opts *sql.TxOptions) (DbCompoundTable, error) {
	var err error
	tbl.db, err = tbl.db.(sqlgen2.TxStarter).BeginTx(tbl.ctx, opts)
	return tbl, tbl.logIfError(err)
}

// Using returns a modified Table using the transaction supplied. This is needed
// when making multiple queries across several tables within a single transaction.
// The result is a modified copy of the table; the original is unchanged.
func (tbl DbCompoundTable) Using(tx *sql.Tx) DbCompoundTable {
	tbl.db = tx
	return tbl
}

func (tbl DbCompoundTable) logQuery(query string, args ...interface{}) {
	tbl.database.LogQuery(query, args...)
}

func (tbl DbCompoundTable) logError(err error) error {
	return tbl.database.LogError(err)
}

func (tbl DbCompoundTable) logIfError(err error) error {
	return tbl.database.LogIfError(err)
}

//--------------------------------------------------------------------------------

const NumDbCompoundColumns = 3

const NumDbCompoundDataColumns = 3

const DbCompoundColumnNames = "alpha,beta,category"

//--------------------------------------------------------------------------------

const sqlDbCompoundTableCreateColumnsSqlite = "\n" +
	" `alpha`    text not null,\n" +
	" `beta`     text not null,\n" +
	" `category` tinyint unsigned not null"

const sqlDbCompoundTableCreateColumnsMysql = "\n" +
	" `alpha`    varchar(255) not null,\n" +
	" `beta`     varchar(255) not null,\n" +
	" `category` tinyint unsigned not null"

const sqlDbCompoundTableCreateColumnsPostgres = `
 "alpha"    varchar(255) not null,
 "beta"     varchar(255) not null,
 "category" smallint not null`

const sqlConstrainDbCompoundTable = `
`

//--------------------------------------------------------------------------------

const sqlDbAlphaBetaIndexColumns = "alpha,beta"

//--------------------------------------------------------------------------------

// CreateTable creates the table.
func (tbl DbCompoundTable) CreateTable(ifNotExists bool) (int64, error) {
	return support.Exec(tbl, nil, tbl.createTableSql(ifNotExists))
}

func (tbl DbCompoundTable) createTableSql(ifNotExists bool) string {
	var columns string
	var settings string
	switch tbl.Dialect() {
	case schema.Sqlite:
		columns = sqlDbCompoundTableCreateColumnsSqlite
		settings = ""
	case schema.Mysql:
		columns = sqlDbCompoundTableCreateColumnsMysql
		settings = " ENGINE=InnoDB DEFAULT CHARSET=utf8"
	case schema.Postgres:
		columns = sqlDbCompoundTableCreateColumnsPostgres
		settings = ""
	}
	buf := &bytes.Buffer{}
	buf.WriteString("CREATE TABLE ")
	if ifNotExists {
		buf.WriteString("IF NOT EXISTS ")
	}
	buf.WriteString(tbl.name.String())
	buf.WriteString(" (")
	buf.WriteString(columns)
	for i, c := range tbl.constraints {
		buf.WriteString(",\n ")
		buf.WriteString(c.ConstraintSql(tbl.Dialect(), tbl.name, i+1))
	}
	buf.WriteString("\n)")
	buf.WriteString(settings)
	return buf.String()
}

func (tbl DbCompoundTable) ternary(flag bool, a, b string) string {
	if flag {
		return a
	}
	return b
}

// DropTable drops the table, destroying all its data.
func (tbl DbCompoundTable) DropTable(ifExists bool) (int64, error) {
	return support.Exec(tbl, nil, tbl.dropTableSql(ifExists))
}

func (tbl DbCompoundTable) dropTableSql(ifExists bool) string {
	ie := tbl.ternary(ifExists, "IF EXISTS ", "")
	query := fmt.Sprintf("DROP TABLE %s%s", ie, tbl.name)
	return query
}

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
	ine := tbl.ternary(ifNotExist && tbl.Dialect() != schema.Mysql, "IF NOT EXISTS ", "")

	// Mysql does not support 'if not exists' on indexes
	// Workaround: use DropIndex first and ignore an error returned if the index didn't exist.

	if ifNotExist && tbl.Dialect() == schema.Mysql {
		// low-level no-logging Exec
		tbl.Execer().ExecContext(tbl.ctx, tbl.dropDbAlphaBetaIndexSql(false))
		ine = ""
	}

	_, err := tbl.Exec(nil, tbl.createDbAlphaBetaIndexSql(ine))
	return err
}

func (tbl DbCompoundTable) createDbAlphaBetaIndexSql(ifNotExists string) string {
	indexPrefix := tbl.name.PrefixWithoutDot()
	return fmt.Sprintf("CREATE UNIQUE INDEX %s%salpha_beta ON %s (%s)", ifNotExists, indexPrefix,
		tbl.name, sqlDbAlphaBetaIndexColumns)
}

// DropAlphaBetaIndex drops the alpha_beta index.
func (tbl DbCompoundTable) DropAlphaBetaIndex(ifExists bool) error {
	_, err := tbl.Exec(nil, tbl.dropDbAlphaBetaIndexSql(ifExists))
	return err
}

func (tbl DbCompoundTable) dropDbAlphaBetaIndexSql(ifExists bool) string {
	// Mysql does not support 'if exists' on indexes
	ie := tbl.ternary(ifExists && tbl.Dialect() != schema.Mysql, "IF EXISTS ", "")
	onTbl := tbl.ternary(tbl.Dialect() == schema.Mysql, fmt.Sprintf(" ON %s", tbl.name), "")
	indexPrefix := tbl.name.PrefixWithoutDot()
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

// Truncate drops every record from the table, if possible. It might fail if constraints exist that
// prevent some or all rows from being deleted; use the force option to override this.
//
// When 'force' is set true, be aware of the following consequences.
// When using Mysql, foreign keys in other tables can be left dangling.
// When using Postgres, a cascade happens, so all 'adjacent' tables (i.e. linked by foreign keys)
// are also truncated.
func (tbl DbCompoundTable) Truncate(force bool) (err error) {
	for _, query := range tbl.Dialect().TruncateDDL(tbl.Name().String(), force) {
		_, err = support.Exec(tbl, nil, query)
		if err != nil {
			return err
		}
	}
	return nil
}

//--------------------------------------------------------------------------------

// Exec executes a query without returning any rows.
// It returns the number of rows affected (if the database driver supports this).
//
// The args are for any placeholder parameters in the query.
func (tbl DbCompoundTable) Exec(req require.Requirement, query string, args ...interface{}) (int64, error) {
	return support.Exec(tbl, req, query, args...)
}

//--------------------------------------------------------------------------------

// Query is the low-level request method for this table. The query is logged using whatever logger is
// configured. If an error arises, this too is logged.
//
// If you need a context other than the background, use WithContext before calling Query.
//
// The args are for any placeholder parameters in the query.
//
// The caller must call rows.Close() on the result.
//
// Wrap the result in *sqlgen2.Rows if you need to access its data as a map.
func (tbl DbCompoundTable) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return support.Query(tbl, query, args...)
}

//--------------------------------------------------------------------------------

// QueryOneNullString is a low-level access method for one string. This can be used for function queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl DbCompoundTable) QueryOneNullString(req require.Requirement, query string, args ...interface{}) (result sql.NullString, err error) {
	err = support.QueryOneNullThing(tbl, req, &result, query, args...)
	return result, err
}

// QueryOneNullInt64 is a low-level access method for one int64. This can be used for 'COUNT(1)' queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl DbCompoundTable) QueryOneNullInt64(req require.Requirement, query string, args ...interface{}) (result sql.NullInt64, err error) {
	err = support.QueryOneNullThing(tbl, req, &result, query, args...)
	return result, err
}

// QueryOneNullFloat64 is a low-level access method for one float64. This can be used for 'AVG(...)' queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl DbCompoundTable) QueryOneNullFloat64(req require.Requirement, query string, args ...interface{}) (result sql.NullFloat64, err error) {
	err = support.QueryOneNullThing(tbl, req, &result, query, args...)
	return result, err
}

func scanDbCompounds(rows *sql.Rows, firstOnly bool) (vv []*Compound, n int64, err error) {
	for rows.Next() {
		n++

		var v0 string
		var v1 string
		var v2 Category

		err = rows.Scan(
			&v0,
			&v1,
			&v2,
		)
		if err != nil {
			return vv, n, err
		}

		v := &Compound{}
		v.Alpha = v0
		v.Beta = v1
		v.Category = v2

		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPostGet); ok {
			err = hook.PostGet()
			if err != nil {
				return vv, n, err
			}
		}

		vv = append(vv, v)

		if firstOnly {
			if rows.Next() {
				n++
			}
			return vv, n, rows.Err()
		}
	}

	return vv, n, rows.Err()
}

//--------------------------------------------------------------------------------

var allDbCompoundQuotedColumnNames = []string{
	schema.Sqlite.SplitAndQuote(DbCompoundColumnNames),
	schema.Mysql.SplitAndQuote(DbCompoundColumnNames),
	schema.Postgres.SplitAndQuote(DbCompoundColumnNames),
}

// GetCompoundByAlphaAndBeta gets the record with given alpha+beta values.
// If not found, *Compound will be nil.
func (tbl DbCompoundTable) GetCompoundByAlphaAndBeta(req require.Requirement, alpha string, beta string) (*Compound, error) {
	return tbl.SelectOne(req, where.And(where.Eq("alpha", alpha), where.Eq("beta", beta)), nil)
}

func (tbl DbCompoundTable) getCompound(req require.Requirement, column string, arg interface{}) (*Compound, error) {
	dialect := tbl.Dialect()
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s=%s",
		allDbCompoundQuotedColumnNames[dialect.Index()], tbl.name, dialect.Quote(column), dialect.Placeholder(column, 1))
	v, err := tbl.doQueryOne(req, query, arg)
	return v, err
}

func (tbl DbCompoundTable) getCompounds(req require.Requirement, column string, args ...interface{}) (list []*Compound, err error) {
	if len(args) > 0 {
		if req == require.All {
			req = require.Exactly(len(args))
		}
		dialect := tbl.Dialect()
		pl := dialect.Placeholders(len(args))
		query := fmt.Sprintf("SELECT %s FROM %s WHERE %s IN (%s)",
			allDbCompoundQuotedColumnNames[dialect.Index()], tbl.name, dialect.Quote(column), pl)
		list, err = tbl.doQuery(req, false, query, args...)
	}

	return list, err
}

func (tbl DbCompoundTable) doQueryOne(req require.Requirement, query string, args ...interface{}) (*Compound, error) {
	list, err := tbl.doQuery(req, true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

func (tbl DbCompoundTable) doQuery(req require.Requirement, firstOnly bool, query string, args ...interface{}) ([]*Compound, error) {
	rows, err := tbl.Query(query, args...)
	if err != nil {
		return nil, err
	}

	vv, n, err := scanDbCompounds(rows, firstOnly)
	return vv, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(err, req, n))
}

// Fetch fetches a list of Compound based on a supplied query. This is mostly used for join queries that map its
// result columns to the fields of Compound. Other queries might be better handled by GetXxx or Select methods.
func (tbl DbCompoundTable) Fetch(req require.Requirement, query string, args ...interface{}) ([]*Compound, error) {
	return tbl.doQuery(req, false, query, args...)
}

//--------------------------------------------------------------------------------

// SelectOneWhere allows a single Example to be obtained from the table that match a 'where' clause
// and some limit. Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
// If not found, *Example will be nil.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.One
// controls whether an error is generated when no result is found.
//
// The args are for any placeholder parameters in the query.
func (tbl DbCompoundTable) SelectOneWhere(req require.Requirement, where, orderBy string, args ...interface{}) (*Compound, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT 1",
		allDbCompoundQuotedColumnNames[tbl.Dialect().Index()], tbl.name, where, orderBy)
	v, err := tbl.doQueryOne(req, query, args...)
	return v, err
}

// SelectOne allows a single Compound to be obtained from the sqlgen2.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
// If not found, *Example will be nil.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.One
// controls whether an error is generated when no result is found.
func (tbl DbCompoundTable) SelectOne(req require.Requirement, wh where.Expression, qc where.QueryConstraint) (*Compound, error) {
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
	return tbl.SelectOneWhere(req, whs, orderBy, args...)
}

// SelectWhere allows Compounds to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
//
// The args are for any placeholder parameters in the query.
func (tbl DbCompoundTable) SelectWhere(req require.Requirement, where, orderBy string, args ...interface{}) ([]*Compound, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s %s",
		allDbCompoundQuotedColumnNames[tbl.Dialect().Index()], tbl.name, where, orderBy)
	vv, err := tbl.doQuery(req, false, query, args...)
	return vv, err
}

// Select allows Compounds to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
func (tbl DbCompoundTable) Select(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]*Compound, error) {
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
	return tbl.SelectWhere(req, whs, orderBy, args...)
}

// CountWhere counts Compounds in the table that match a 'where' clause.
// Use a blank string for the 'where' argument if it is not needed.
//
// The args are for any placeholder parameters in the query.
func (tbl DbCompoundTable) CountWhere(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s %s", tbl.name, where)
	tbl.logQuery(query, args...)
	row := tbl.db.QueryRowContext(tbl.ctx, query, args...)
	err = row.Scan(&count)
	return count, tbl.logIfError(err)
}

// Count counts the Compounds in the table that match a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed.
func (tbl DbCompoundTable) Count(wh where.Expression) (count int64, err error) {
	whs, args := where.BuildExpression(wh, tbl.Dialect())
	return tbl.CountWhere(whs, args...)
}

//--------------------------------------------------------------------------------

// SliceAlpha gets the alpha column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DbCompoundTable) SliceAlpha(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.sliceStringList(req, "alpha", wh, qc)
}

// SliceBeta gets the beta column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DbCompoundTable) SliceBeta(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.sliceStringList(req, "beta", wh, qc)
}

// SliceCategory gets the category column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DbCompoundTable) SliceCategory(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]Category, error) {
	return tbl.sliceCategoryList(req, "category", wh, qc)
}

func (tbl DbCompoundTable) sliceCategoryList(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]Category, error) {
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", dialect.Quote(sqlname), tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, tbl.logError(err)
	}
	defer rows.Close()

	var v Category
	list := make([]Category, 0, 10)

	for rows.Next() {
		err = rows.Scan(&v)
		if err == sql.ErrNoRows {
			return list, tbl.logIfError(require.ErrorIfQueryNotSatisfiedBy(req, int64(len(list))))
		} else {
			list = append(list, v)
		}
	}
	return list, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, int64(len(list))))
}

func (tbl DbCompoundTable) sliceStringList(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", dialect.Quote(sqlname), tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, tbl.logError(err)
	}
	defer rows.Close()

	var v string
	list := make([]string, 0, 10)

	for rows.Next() {
		err = rows.Scan(&v)
		if err == sql.ErrNoRows {
			return list, tbl.logIfError(require.ErrorIfQueryNotSatisfiedBy(req, int64(len(list))))
		} else {
			list = append(list, v)
		}
	}
	return list, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, int64(len(list))))
}

func constructDbCompoundInsert(w io.Writer, v *Compound, dialect schema.Dialect, withPk bool) (s []interface{}, err error) {
	s = make([]interface{}, 0, 3)

	comma := ""
	io.WriteString(w, " (")

	io.WriteString(w, comma)

	dialect.QuoteW(w, "alpha")
	s = append(s, v.Alpha)
	comma = ","
	io.WriteString(w, comma)

	dialect.QuoteW(w, "beta")
	s = append(s, v.Beta)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "category")
	s = append(s, v.Category)
	io.WriteString(w, ")")
	return s, nil
}

func constructDbCompoundUpdate(w io.Writer, v *Compound, dialect schema.Dialect) (s []interface{}, err error) {
	j := 1
	s = make([]interface{}, 0, 3)

	comma := ""

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "alpha", j)
	s = append(s, v.Alpha)
	comma = ", "
	j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "beta", j)
	s = append(s, v.Beta)
	j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "category", j)
	s = append(s, v.Category)
	j++

	return s, nil
}

//--------------------------------------------------------------------------------

// Insert adds new records for the Compounds.

// The Compound.PreInsert() method will be called, if it exists.
func (tbl DbCompoundTable) Insert(req require.Requirement, vv ...*Compound) error {
	if req == require.All {
		req = require.Exactly(len(vv))
	}

	var count int64
	//columns := allXExampleQuotedInserts[tbl.Dialect().Index()]
	//query := fmt.Sprintf("INSERT INTO %s %s", tbl.name, columns)
	//st, err := tbl.db.PrepareContext(tbl.ctx, query)
	//if err != nil {
	//	return err
	//}
	//defer st.Close()

	insertHasReturningPhrase := false
	returning := ""
	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPreInsert); ok {
			err := hook.PreInsert()
			if err != nil {
				return tbl.logError(err)
			}
		}

		b := &bytes.Buffer{}
		io.WriteString(b, "INSERT INTO ")
		io.WriteString(b, tbl.name.String())

		fields, err := constructDbCompoundInsert(b, v, tbl.Dialect(), true)
		if err != nil {
			return tbl.logError(err)
		}

		io.WriteString(b, " VALUES (")
		io.WriteString(b, tbl.Dialect().Placeholders(len(fields)))
		io.WriteString(b, ")")
		io.WriteString(b, returning)

		query := b.String()
		tbl.logQuery(query, fields...)

		var n int64 = 1
		if insertHasReturningPhrase {
			row := tbl.db.QueryRowContext(tbl.ctx, query, fields...)
			var i64 int64
			err = row.Scan(&i64)

		} else {
			res, e2 := tbl.db.ExecContext(tbl.ctx, query, fields...)
			if e2 != nil {
				return tbl.logError(e2)
			}

			if e2 != nil {
				return tbl.logError(e2)
			}

			n, err = res.RowsAffected()
		}

		if err != nil {
			return tbl.logError(err)
		}
		count += n
	}

	return tbl.logIfError(require.ErrorIfExecNotSatisfiedBy(req, count))
}

// UpdateFields updates one or more columns, given a 'where' clause.
//
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl DbCompoundTable) UpdateFields(req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error) {
	return support.UpdateFields(tbl, req, wh, fields...)
}

//--------------------------------------------------------------------------------

// Delete deletes one or more rows from the table, given a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl DbCompoundTable) Delete(req require.Requirement, wh where.Expression) (int64, error) {
	query, args := tbl.deleteRows(wh)
	return tbl.Exec(req, query, args...)
}

func (tbl DbCompoundTable) deleteRows(wh where.Expression) (string, []interface{}) {
	whs, args := where.BuildExpression(wh, tbl.Dialect())
	query := fmt.Sprintf("DELETE FROM %s %s", tbl.name, whs)
	return query, args
}

//--------------------------------------------------------------------------------
