// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

package demo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/rickb777/sqlgen2"
	"github.com/rickb777/sqlgen2/require"
	"github.com/rickb777/sqlgen2/schema"
	"github.com/rickb777/sqlgen2/where"
	"log"
	"strings"
)

// AssociationTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type AssociationTable struct {
	name    sqlgen2.TableName
	db      sqlgen2.Execer
	ctx     context.Context
	dialect schema.Dialect
	logger  *log.Logger
	wrapper interface{}
}

// Type conformance checks
var _ sqlgen2.TableCreator = &AssociationTable{}
var _ sqlgen2.TableWithCrud = &AssociationTable{}

// NewAssociationTable returns a new table instance.
// If a blank table name is supplied, the default name "associations" will be used instead.
// The request context is initialised with the background.
func NewAssociationTable(name sqlgen2.TableName, d sqlgen2.Execer, dialect schema.Dialect) AssociationTable {
	if name.Name == "" {
		name.Name = "associations"
	}
	return AssociationTable{name, d, context.Background(), dialect, nil, nil}
}

// CopyTableAsAssociationTable copies a table instance, retaining the name etc but
// providing methods appropriate for 'Association'.
func CopyTableAsAssociationTable(origin sqlgen2.Table) AssociationTable {
	return AssociationTable{
		name:    origin.Name(),
		db:      origin.DB(),
		ctx:     origin.Ctx(),
		dialect: origin.Dialect(),
		logger:  origin.Logger(),
	}
}

// WithPrefix sets the table name prefix for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl AssociationTable) WithPrefix(pfx string) AssociationTable {
	tbl.name.Prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl AssociationTable) WithContext(ctx context.Context) AssociationTable {
	tbl.ctx = ctx
	return tbl
}

// WithLogger sets the logger for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl AssociationTable) WithLogger(logger *log.Logger) AssociationTable {
	tbl.logger = logger
	return tbl
}

// Logger gets the trace logger.
func (tbl AssociationTable) Logger() *log.Logger {
	return tbl.logger
}

// SetLogger sets the logger for subsequent queries, returning the interface.
// The result is a modified copy of the table; the original is unchanged.
func (tbl AssociationTable) SetLogger(logger *log.Logger) sqlgen2.Table {
	tbl.logger = logger
	return tbl
}

// Ctx gets the current request context.
func (tbl AssociationTable) Ctx() context.Context {
	return tbl.ctx
}

// Dialect gets the database dialect.
func (tbl AssociationTable) Dialect() schema.Dialect {
	return tbl.dialect
}

// Wrapper gets the user-defined wrapper.
func (tbl AssociationTable) Wrapper() interface{} {
	return tbl.wrapper
}

// SetWrapper sets the user-defined wrapper.
// The result is a modified copy of the table; the original is unchanged.
func (tbl AssociationTable) SetWrapper(wrapper interface{}) sqlgen2.Table {
	tbl.wrapper = wrapper
	return tbl
}

// Name gets the table name.
func (tbl AssociationTable) Name() sqlgen2.TableName {
	return tbl.name
}

// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl AssociationTable) DB() *sql.DB {
	return tbl.db.(*sql.DB)
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl AssociationTable) Tx() *sql.Tx {
	return tbl.db.(*sql.Tx)
}

// IsTx tests whether this is within a transaction.
func (tbl AssociationTable) IsTx() bool {
	_, ok := tbl.db.(*sql.Tx)
	return ok
}

// Begin starts a transaction. The default isolation level is dependent on the driver.
// The result is a modified copy of the table; the original is unchanged.
func (tbl AssociationTable) BeginTx(opts *sql.TxOptions) (AssociationTable, error) {
	d := tbl.db.(*sql.DB)
	var err error
	tbl.db, err = d.BeginTx(tbl.ctx, opts)
	return tbl, tbl.logIfError(err)
}

// Using returns a modified Table using the transaction supplied. This is needed
// when making multiple queries across several tables within a single transaction.
// The result is a modified copy of the table; the original is unchanged.
func (tbl AssociationTable) Using(tx *sql.Tx) AssociationTable {
	tbl.db = tx
	return tbl
}

func (tbl AssociationTable) logQuery(query string, args ...interface{}) {
	sqlgen2.LogQuery(tbl.logger, query, args...)
}

func (tbl AssociationTable) logError(err error) error {
	return sqlgen2.LogError(tbl.logger, err)
}

func (tbl AssociationTable) logIfError(err error) error {
	return sqlgen2.LogIfError(tbl.logger, err)
}


//--------------------------------------------------------------------------------

const NumAssociationColumns = 6

const NumAssociationDataColumns = 5

const AssociationPk = "Id"

const AssociationDataColumnNames = "name, quality, ref1, ref2, category"

//--------------------------------------------------------------------------------

// CreateTable creates the table.
func (tbl AssociationTable) CreateTable(ifNotExists bool) (int64, error) {
	return tbl.Exec(nil, tbl.createTableSql(ifNotExists))
}

func (tbl AssociationTable) createTableSql(ifNotExists bool) string {
	var stmt string
	switch tbl.dialect {
	case schema.Sqlite: stmt = sqlCreateAssociationTableSqlite
    case schema.Postgres: stmt = sqlCreateAssociationTablePostgres
    case schema.Mysql: stmt = sqlCreateAssociationTableMysql
    }
	extra := tbl.ternary(ifNotExists, "IF NOT EXISTS ", "")
	query := fmt.Sprintf(stmt, extra, tbl.name)
	return query
}

func (tbl AssociationTable) ternary(flag bool, a, b string) string {
	if flag {
		return a
	}
	return b
}

// DropTable drops the table, destroying all its data.
func (tbl AssociationTable) DropTable(ifExists bool) (int64, error) {
	return tbl.Exec(nil, tbl.dropTableSql(ifExists))
}

func (tbl AssociationTable) dropTableSql(ifExists bool) string {
	extra := tbl.ternary(ifExists, "IF EXISTS ", "")
	query := fmt.Sprintf("DROP TABLE %s%s", extra, tbl.name)
	return query
}

const sqlCreateAssociationTableSqlite = `
CREATE TABLE %s%s (
 id       integer primary key autoincrement,
 name     text default null,
 quality  text default null,
 ref1     bigint default null,
 ref2     bigint default null,
 category tinyint unsigned default null
)
`

const sqlCreateAssociationTablePostgres = `
CREATE TABLE %s%s (
 id       bigserial primary key,
 name     varchar(255) default null,
 quality  varchar(255) default null,
 ref1     bigint default null,
 ref2     bigint default null,
 category tinyint unsigned default null
)
`

const sqlCreateAssociationTableMysql = `
CREATE TABLE %s%s (
 id       bigint primary key auto_increment,
 name     varchar(255) default null,
 quality  varchar(255) default null,
 ref1     bigint default null,
 ref2     bigint default null,
 category tinyint unsigned default null
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
func (tbl AssociationTable) Truncate(force bool) (err error) {
	for _, query := range tbl.dialect.TruncateDDL(tbl.Name().String(), force) {
		_, err = tbl.Exec(nil, query)
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
func (tbl AssociationTable) Exec(req require.Requirement, query string, args ...interface{}) (int64, error) {
	tbl.logQuery(query, args...)
	res, err := tbl.db.ExecContext(tbl.ctx, query, args...)
	if err != nil {
		return 0, tbl.logError(err)
	}
	n, err := res.RowsAffected()
	return n, tbl.logIfError(require.ChainErrorIfExecNotSatisfiedBy(err, req, n))
}

//--------------------------------------------------------------------------------

// Query is the low-level access method for Associations.
//
// It places a requirement, which may be nil, on the size of the expected results: this
// controls whether an error is generated when this expectation is not met.
//
// Note that this method applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl AssociationTable) Query(req require.Requirement, query string, args ...interface{}) ([]*Association, error) {
	query = tbl.ReplaceTableName(query)
	vv, err := tbl.doQuery(req, false, query, args...)
	return vv, err
}

// QueryOne is the low-level access method for one Association.
// If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, *Association will be nil.
//
// Note that this method applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl AssociationTable) QueryOne(query string, args ...interface{}) (*Association, error) {
	query = tbl.ReplaceTableName(query)
	return tbl.doQueryOne(nil, query, args...)
}

// MustQueryOne is the low-level access method for one Association.
//
// It places a requirement that exactly one result must be found; an error is generated when this expectation is not met.
//
// Note that this method applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl AssociationTable) MustQueryOne(query string, args ...interface{}) (*Association, error) {
	query = tbl.ReplaceTableName(query)
	return tbl.doQueryOne(require.One, query, args...)
}

func (tbl AssociationTable) doQueryOne(req require.Requirement, query string, args ...interface{}) (*Association, error) {
	list, err := tbl.doQuery(req, true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

func (tbl AssociationTable) doQuery(req require.Requirement, firstOnly bool, query string, args ...interface{}) ([]*Association, error) {
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, tbl.logError(err)
	}
	defer rows.Close()

	vv, n, err := scanAssociations(rows, firstOnly)
	return vv, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(err, req, n))
}

func scanAssociations(rows *sql.Rows, firstOnly bool) (vv []*Association, n int64, err error) {
	for rows.Next() {
		n++

		var v0 int64
		var v1 sql.NullString
		var v2 sql.NullString
		var v3 sql.NullInt64
		var v4 sql.NullInt64
		var v5 sql.NullInt64

		err = rows.Scan(
			&v0,
			&v1,
			&v2,
			&v3,
			&v4,
			&v5,
		)
		if err != nil {
			return vv, n, err
		}

		v := &Association{}
		v.Id = v0
		if v1.Valid {
			a := v1.String
			v.Name = &a
		}
		if v2.Valid {
			a := v2.String
			v.Quality = &a
		}
		if v3.Valid {
			a := v3.Int64
			v.Ref1 = &a
		}
		if v4.Valid {
			a := v4.Int64
			v.Ref2 = &a
		}
		if v5.Valid {
			a := Category(v5.Int64)
			v.Category = &a
		}

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

// QueryOneNullString is a low-level access method for one string. This can be used for function queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl AssociationTable) QueryOneNullString(query string, args ...interface{}) (result sql.NullString, err error) {
	err = tbl.doQueryOneNullThing(nil, &result, query, args...)
	return result, err
}

// MustQueryOneNullString is a low-level access method for one string. This can be used for function queries and
// such like.
//
// It places a requirement that exactly one result must be found; an error is generated when this expectation is not met.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl AssociationTable) MustQueryOneNullString(query string, args ...interface{}) (result sql.NullString, err error) {
	err = tbl.doQueryOneNullThing(require.One, &result, query, args...)
	return result, err
}

// QueryOneNullInt64 is a low-level access method for one int64. This can be used for 'COUNT(1)' queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl AssociationTable) QueryOneNullInt64(query string, args ...interface{}) (result sql.NullInt64, err error) {
	err = tbl.doQueryOneNullThing(nil, &result, query, args...)
	return result, err
}

// MustQueryOneNullInt64 is a low-level access method for one int64. This can be used for 'COUNT(1)' queries and
// such like.
//
// It places a requirement that exactly one result must be found; an error is generated when this expectation is not met.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl AssociationTable) MustQueryOneNullInt64(query string, args ...interface{}) (result sql.NullInt64, err error) {
	err = tbl.doQueryOneNullThing(require.One, &result, query, args...)
	return result, err
}

// QueryOneNullFloat64 is a low-level access method for one float64. This can be used for 'AVG(...)' queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl AssociationTable) QueryOneNullFloat64(query string, args ...interface{}) (result sql.NullFloat64, err error) {
	err = tbl.doQueryOneNullThing(nil, &result, query, args...)
	return result, err
}

// MustQueryOneNullFloat64 is a low-level access method for one float64. This can be used for 'AVG(...)' queries and
// such like.
//
// It places a requirement that exactly one result must be found; an error is generated when this expectation is not met.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl AssociationTable) MustQueryOneNullFloat64(query string, args ...interface{}) (result sql.NullFloat64, err error) {
	err = tbl.doQueryOneNullThing(require.One, &result, query, args...)
	return result, err
}

func (tbl AssociationTable) doQueryOneNullThing(req require.Requirement, holder interface{}, query string, args ...interface{}) error {
	var n int64 = 0
	query = tbl.ReplaceTableName(query)
	tbl.logQuery(query, args...)

	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return tbl.logError(err)
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(holder)

		if err == sql.ErrNoRows {
			return tbl.logIfError(require.ErrorIfQueryNotSatisfiedBy(req, 0))
		} else {
			n++
		}

		if rows.Next() {
			n++ // not singular
		}
	}

	return tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, n))
}

// ReplaceTableName replaces all occurrences of "{TABLE}" with the table's name.
func (tbl AssociationTable) ReplaceTableName(query string) string {
	return strings.Replace(query, "{TABLE}", tbl.name.String(), -1)
}

//--------------------------------------------------------------------------------

// GetAssociation gets the record with a given primary key value.
// If not found, *Association will be nil.
func (tbl AssociationTable) GetAssociation(id int64) (*Association, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE id=?", AssociationColumnNames, tbl.name)
	v, err := tbl.doQueryOne(nil, query, id)
	return v, err
}

// MustGetAssociation gets the record with a given primary key value.
//
// It places a requirement that exactly one result must be found; an error is generated when this expectation is not met.
func (tbl AssociationTable) MustGetAssociation(id int64) (*Association, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE id=?", AssociationColumnNames, tbl.name)
	v, err := tbl.doQueryOne(require.One, query, id)
	return v, err
}

// GetAssociations gets records from the table according to a list of primary keys.
// Although the list of ids can be arbitrarily long, there are practical limits;
// note that Oracle DB has a limit of 1000.
//
// It places a requirement, which may be nil, on the size of the expected results: in particular, require.All
// controls whether an error is generated not all the ids produce a result.
func (tbl AssociationTable) GetAssociations(req require.Requirement, id ...int64) (list []*Association, err error) {
	if len(id) > 0 {
		if req == require.All {
			req = require.Exactly(len(id))
		}
		pl := tbl.dialect.Placeholders(len(id))
		query := fmt.Sprintf("SELECT %s FROM %s WHERE id IN (%s)", AssociationColumnNames, tbl.name, pl)
		args := make([]interface{}, len(id))

		for i, v := range id {
			args[i] = v
		}

		list, err = tbl.doQuery(req, false, query, args...)
	}

	return list, err
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
func (tbl AssociationTable) SelectOneWhere(req require.Requirement, where, orderBy string, args ...interface{}) (*Association, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT 1", AssociationColumnNames, tbl.name, where, orderBy)
	v, err := tbl.doQueryOne(req, query, args...)
	return v, err
}

// SelectOne allows a single Association to be obtained from the sqlgen2.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
// If not found, *Example will be nil.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.One
// controls whether an error is generated when no result is found.
func (tbl AssociationTable) SelectOne(req require.Requirement, wh where.Expression, qc where.QueryConstraint) (*Association, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	return tbl.SelectOneWhere(req, whs, orderBy, args...)
}

// SelectWhere allows Associations to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
//
// The args are for any placeholder parameters in the query.
func (tbl AssociationTable) SelectWhere(req require.Requirement, where, orderBy string, args ...interface{}) ([]*Association, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", AssociationColumnNames, tbl.name, where, orderBy)
	vv, err := tbl.doQuery(req, false, query, args...)
	return vv, err
}

// Select allows Associations to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
func (tbl AssociationTable) Select(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]*Association, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	return tbl.SelectWhere(req, whs, orderBy, args...)
}

// CountWhere counts Associations in the table that match a 'where' clause.
// Use a blank string for the 'where' argument if it is not needed.
//
// The args are for any placeholder parameters in the query.
func (tbl AssociationTable) CountWhere(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s %s", tbl.name, where)
	tbl.logQuery(query, args...)
	row := tbl.db.QueryRowContext(tbl.ctx, query, args...)
	err = row.Scan(&count)
	return count, tbl.logIfError(err)
}

// Count counts the Associations in the table that match a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed.
func (tbl AssociationTable) Count(wh where.Expression) (count int64, err error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	return tbl.CountWhere(whs, args...)
}

const AssociationColumnNames = "id, name, quality, ref1, ref2, category"

//--------------------------------------------------------------------------------

// SliceId gets the Id column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AssociationTable) SliceId(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return tbl.getint64list(req, "id", wh, qc)
}

// SliceName gets the Name column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AssociationTable) SliceName(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.getstringPtrlist(req, "name", wh, qc)
}

// SliceQuality gets the Quality column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AssociationTable) SliceQuality(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.getstringPtrlist(req, "quality", wh, qc)
}

// SliceRef1 gets the Ref1 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AssociationTable) SliceRef1(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return tbl.getint64Ptrlist(req, "ref1", wh, qc)
}

// SliceRef2 gets the Ref2 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AssociationTable) SliceRef2(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return tbl.getint64Ptrlist(req, "ref2", wh, qc)
}

// SliceCategory gets the Category column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AssociationTable) SliceCategory(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]Category, error) {
	return tbl.getCategoryPtrlist(req, "category", wh, qc)
}


func (tbl AssociationTable) getCategoryPtrlist(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]Category, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", sqlname, tbl.name, whs, orderBy)
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

func (tbl AssociationTable) getint64list(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", sqlname, tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, tbl.logError(err)
	}
	defer rows.Close()

	var v int64
	list := make([]int64, 0, 10)

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

func (tbl AssociationTable) getint64Ptrlist(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", sqlname, tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, tbl.logError(err)
	}
	defer rows.Close()

	var v int64
	list := make([]int64, 0, 10)

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

func (tbl AssociationTable) getstringPtrlist(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", sqlname, tbl.name, whs, orderBy)
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


//--------------------------------------------------------------------------------

// Insert adds new records for the Associations.
// The Associations have their primary key fields set to the new record identifiers.
// The Association.PreInsert(Execer) method will be called, if it exists.
func (tbl AssociationTable) Insert(req require.Requirement, vv ...*Association) error {
	var params string
	switch tbl.dialect {
	case schema.Postgres:
		params = sAssociationDataColumnParamsPostgres
	default:
		params = sAssociationDataColumnParamsSimple
	}

	if req == require.All {
		req = require.Exactly(len(vv))
	}

	var count int64
	query := fmt.Sprintf(sqlInsertAssociation, tbl.name, params)
	st, err := tbl.db.PrepareContext(tbl.ctx, query)
	if err != nil {
		return err
	}
	defer st.Close()

	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPreInsert); ok {
			err := hook.PreInsert()
			if err != nil {
				return tbl.logError(err)
			}
		}

		fields, err := sliceAssociationWithoutPk(v)
		if err != nil {
			return tbl.logError(err)
		}

		tbl.logQuery(query, fields...)
		res, err := st.ExecContext(tbl.ctx, fields...)
		if err != nil {
			return tbl.logError(err)
		}

		v.Id, err = res.LastInsertId()
		if err != nil {
			return tbl.logError(err)
		}

		n, err := res.RowsAffected()
		if err != nil {
			return tbl.logError(err)
		}
		count += n
	}

	return tbl.logIfError(require.ErrorIfExecNotSatisfiedBy(req, count))
}

const sqlInsertAssociation = `
INSERT INTO %s (
	name,
	quality,
	ref1,
	ref2,
	category
) VALUES (%s)
`

const sAssociationDataColumnParamsSimple = "?,?,?,?,?"

const sAssociationDataColumnParamsPostgres = "$1,$2,$3,$4,$5"

//--------------------------------------------------------------------------------

// UpdateFields updates one or more columns, given a 'where' clause.
//
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl AssociationTable) UpdateFields(req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error) {
	query, args := tbl.updateFields(wh, fields...)
	return tbl.Exec(req, query, args...)
}

func (tbl AssociationTable) updateFields(wh where.Expression, fields ...sql.NamedArg) (string, []interface{}) {
	list := sqlgen2.NamedArgList(fields)
	assignments := strings.Join(list.Assignments(tbl.dialect, 1), ", ")
	whs, wargs := where.BuildExpression(wh, tbl.dialect)
	query := fmt.Sprintf("UPDATE %s SET %s %s", tbl.name, assignments, whs)
	args := append(list.Values(), wargs...)
	return query, args
}

//--------------------------------------------------------------------------------

// Update updates records, matching them by primary key. It returns the number of rows affected.
// The Association.PreUpdate(Execer) method will be called, if it exists.
func (tbl AssociationTable) Update(req require.Requirement, vv ...*Association) (int64, error) {
	var stmt string
	switch tbl.dialect {
	case schema.Postgres:
		stmt = sqlUpdateAssociationByPkPostgres
	default:
		stmt = sqlUpdateAssociationByPkSimple
	}

	if req == require.All {
		req = require.Exactly(len(vv))
	}

	var count int64
	query := fmt.Sprintf(stmt, tbl.name)

	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPreUpdate); ok {
			err := hook.PreUpdate()
			if err != nil {
				return count, tbl.logError(err)
			}
		}

		args, err := sliceAssociationWithoutPk(v)
		args = append(args, v.Id)
		if err != nil {
			return count, tbl.logError(err)
		}

		n, err := tbl.Exec(nil, query, args...)
		if err != nil {
			return count, err
		}
		count += n
	}

	return count, tbl.logIfError(require.ErrorIfExecNotSatisfiedBy(req, count))
}

const sqlUpdateAssociationByPkSimple = `
UPDATE %s SET
	name=?,
	quality=?,
	ref1=?,
	ref2=?,
	category=?
WHERE id=?
`

const sqlUpdateAssociationByPkPostgres = `
UPDATE %s SET
	name=$2,
	quality=$3,
	ref1=$4,
	ref2=$5,
	category=$6
WHERE id=$1
`

func sliceAssociationWithoutPk(v *Association) ([]interface{}, error) {


	return []interface{}{
		v.Name,
		v.Quality,
		v.Ref1,
		v.Ref2,
		v.Category,

	}, nil
}

//--------------------------------------------------------------------------------

// DeleteAssociations deletes rows from the table, given some primary keys.
// The list of ids can be arbitrarily long.
func (tbl AssociationTable) DeleteAssociations(req require.Requirement, id ...int64) (int64, error) {
	const batch = 1000 // limited by Oracle DB
	const qt = "DELETE FROM %s WHERE id IN (%s)"

	if req == require.All {
		req = require.Exactly(len(id))
	}

	var count, n int64
	var err error
	var max = batch
	if len(id) < batch {
		max = len(id)
	}
	args := make([]interface{}, max)

	if len(id) > batch {
		pl := tbl.dialect.Placeholders(batch)
		query := fmt.Sprintf(qt, tbl.name, pl)

		for len(id) > batch {
			for i := 0; i < batch; i++ {
				args[i] = id[i]
			}

			n, err = tbl.Exec(nil, query, args...)
			count += n
			if err != nil {
				return count, err
			}

			id = id[batch:]
		}
	}

	if len(id) > 0 {
		pl := tbl.dialect.Placeholders(len(id))
		query := fmt.Sprintf(qt, tbl.name, pl)

		for i := 0; i < len(id); i++ {
			args[i] = id[i]
		}

		n, err = tbl.Exec(nil, query, args...)
		count += n
	}

	return count, tbl.logIfError(require.ChainErrorIfExecNotSatisfiedBy(err, req, n))
}

// Delete deletes one or more rows from the table, given a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl AssociationTable) Delete(req require.Requirement, wh where.Expression) (int64, error) {
	query, args := tbl.deleteRows(wh)
	return tbl.Exec(req, query, args...)
}

func (tbl AssociationTable) deleteRows(wh where.Expression) (string, []interface{}) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	query := fmt.Sprintf("DELETE FROM %s %s", tbl.name, whs)
	return query, args
}

//--------------------------------------------------------------------------------
