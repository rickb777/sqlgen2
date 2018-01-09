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

// AssociationTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type AssociationTable struct {
	prefix, name string
	db           sqlgen2.Execer
	ctx          context.Context
	dialect      schema.Dialect
	logger       *log.Logger
}

// Type conformance check
var _ sqlgen2.TableCreator = &AssociationTable{}

// NewAssociationTable returns a new table instance.
// If a blank table name is supplied, the default name "associations" will be used instead.
// The table name prefix is initially blank and the request context is the background.
func NewAssociationTable(name string, d sqlgen2.Execer, dialect schema.Dialect) AssociationTable {
	if name == "" {
		name = "associations"
	}
	return AssociationTable{"", name, d, context.Background(), dialect, nil}
}

// WithPrefix sets the table name prefix for subsequent queries.
func (tbl AssociationTable) WithPrefix(pfx string) AssociationTable {
	tbl.prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries.
func (tbl AssociationTable) WithContext(ctx context.Context) AssociationTable {
	tbl.ctx = ctx
	return tbl
}

// WithLogger sets the logger for subsequent queries.
func (tbl AssociationTable) WithLogger(logger *log.Logger) AssociationTable {
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

// Logger gets the trace logger.
func (tbl AssociationTable) Logger() *log.Logger {
	return tbl.logger
}

// SetLogger sets the logger for subsequent queries, returning the interface.
func (tbl AssociationTable) SetLogger(logger *log.Logger) sqlgen2.Table {
	tbl.logger = logger
	return tbl
}

// Name gets the table name.
func (tbl AssociationTable) Name() string {
	return tbl.name
}

// Prefix gets the table name prefix.
func (tbl AssociationTable) Prefix() string {
	return tbl.prefix
}

// FullName gets the concatenated prefix and table name.
func (tbl AssociationTable) FullName() string {
	return tbl.prefix + tbl.name
}

func (tbl AssociationTable) prefixWithoutDot() string {
	last := len(tbl.prefix)-1
	if last > 0 && tbl.prefix[last] == '.' {
		return tbl.prefix[0:last]
	}
	return tbl.prefix
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
func (tbl AssociationTable) BeginTx(opts *sql.TxOptions) (AssociationTable, error) {
	d := tbl.db.(*sql.DB)
	var err error
	tbl.db, err = d.BeginTx(tbl.ctx, opts)
	return tbl, err
}

func (tbl AssociationTable) logQuery(query string, args ...interface{}) {
	sqlgen2.LogQuery(tbl.logger, query, args...)
}


//--------------------------------------------------------------------------------

const NumAssociationColumns = 6

const NumAssociationDataColumns = 5

const AssociationPk = "Id"

const AssociationDataColumnNames = "name, quality, ref1, ref2, category"

//--------------------------------------------------------------------------------

// CreateTable creates the table.
func (tbl AssociationTable) CreateTable(ifNotExists bool) (int64, error) {
	return tbl.Exec(tbl.createTableSql(ifNotExists))
}

func (tbl AssociationTable) createTableSql(ifNotExists bool) string {
	var stmt string
	switch tbl.dialect {
	case schema.Sqlite: stmt = sqlCreateAssociationTableSqlite
    case schema.Postgres: stmt = sqlCreateAssociationTablePostgres
    case schema.Mysql: stmt = sqlCreateAssociationTableMysql
    }
	extra := tbl.ternary(ifNotExists, "IF NOT EXISTS ", "")
	query := fmt.Sprintf(stmt, extra, tbl.prefix, tbl.name)
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
	return tbl.Exec(tbl.dropTableSql(ifExists))
}

func (tbl AssociationTable) dropTableSql(ifExists bool) string {
	extra := tbl.ternary(ifExists, "IF EXISTS ", "")
	query := fmt.Sprintf("DROP TABLE %s%s%s", extra, tbl.prefix, tbl.name)
	return query
}

const sqlCreateAssociationTableSqlite = `
CREATE TABLE %s%s%s (
 id       integer primary key autoincrement,
 name     text default null,
 quality  text default null,
 ref1     bigint default null,
 ref2     bigint default null,
 category tinyint unsigned default null
)
`

const sqlCreateAssociationTablePostgres = `
CREATE TABLE %s%s%s (
 id       bigserial primary key,
 name     varchar(255) default null,
 quality  varchar(255) default null,
 ref1     bigint default null,
 ref2     bigint default null,
 category tinyint unsigned default null
)
`

const sqlCreateAssociationTableMysql = `
CREATE TABLE %s%s%s (
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
func (tbl AssociationTable) Exec(query string, args ...interface{}) (int64, error) {
	tbl.logQuery(query, args...)
	res, err := tbl.db.ExecContext(tbl.ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

//--------------------------------------------------------------------------------

// QueryOne is the low-level access function for one Association.
// If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, *Association will be nil.
func (tbl AssociationTable) QueryOne(query string, args ...interface{}) (*Association, error) {
	list, err := tbl.doQuery(true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

// Query is the low-level access function for Associations.
func (tbl AssociationTable) Query(query string, args ...interface{}) ([]*Association, error) {
	return tbl.doQuery(false, query, args...)
}

func (tbl AssociationTable) doQuery(firstOnly bool, query string, args ...interface{}) ([]*Association, error) {
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanAssociations(rows, firstOnly)
}

//--------------------------------------------------------------------------------

// GetAssociation gets the record with a given primary key value.
// If not found, *Association will be nil.
func (tbl AssociationTable) GetAssociation(id int64) (*Association, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s WHERE id=?", AssociationColumnNames, tbl.prefix, tbl.name)
	return tbl.QueryOne(query, id)
}

//--------------------------------------------------------------------------------

// SliceId gets the Id column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
func (tbl AssociationTable) SliceId(where where.Expression, orderBy string) ([]int64, error) {
	return tbl.getint64list("id", where, orderBy)
}

// SliceName gets the Name column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
func (tbl AssociationTable) SliceName(where where.Expression, orderBy string) ([]string, error) {
	return tbl.getstringPtrlist("name", where, orderBy)
}

// SliceQuality gets the Quality column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
func (tbl AssociationTable) SliceQuality(where where.Expression, orderBy string) ([]string, error) {
	return tbl.getstringPtrlist("quality", where, orderBy)
}

// SliceRef1 gets the Ref1 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
func (tbl AssociationTable) SliceRef1(where where.Expression, orderBy string) ([]int64, error) {
	return tbl.getint64Ptrlist("ref1", where, orderBy)
}

// SliceRef2 gets the Ref2 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
func (tbl AssociationTable) SliceRef2(where where.Expression, orderBy string) ([]int64, error) {
	return tbl.getint64Ptrlist("ref2", where, orderBy)
}

// SliceCategory gets the Category column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
func (tbl AssociationTable) SliceCategory(where where.Expression, orderBy string) ([]Category, error) {
	return tbl.getCategoryPtrlist("category", where, orderBy)
}


func (tbl AssociationTable) getCategoryPtrlist(sqlname string, where where.Expression, orderBy string) ([]Category, error) {
	wh, args := where.Build(tbl.dialect)
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", sqlname, tbl.prefix, tbl.name, wh, orderBy)
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

func (tbl AssociationTable) getint64list(sqlname string, where where.Expression, orderBy string) ([]int64, error) {
	wh, args := where.Build(tbl.dialect)
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", sqlname, tbl.prefix, tbl.name, wh, orderBy)
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

func (tbl AssociationTable) getint64Ptrlist(sqlname string, where where.Expression, orderBy string) ([]int64, error) {
	wh, args := where.Build(tbl.dialect)
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", sqlname, tbl.prefix, tbl.name, wh, orderBy)
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

func (tbl AssociationTable) getstringPtrlist(sqlname string, where where.Expression, orderBy string) ([]string, error) {
	wh, args := where.Build(tbl.dialect)
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", sqlname, tbl.prefix, tbl.name, wh, orderBy)
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

// SelectOneSA allows a single Association to be obtained from the table that match a 'where' clause
// and some limit.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
// If not found, *Association will be nil.
func (tbl AssociationTable) SelectOneSA(where, orderBy string, args ...interface{}) (*Association, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s LIMIT 1", AssociationColumnNames, tbl.prefix, tbl.name, where, orderBy)
	return tbl.QueryOne(query, args...)
}

// SelectOne allows a single Association to be obtained from the sqlgen2.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
// If not found, *Example will be nil.
func (tbl AssociationTable) SelectOne(where where.Expression, orderBy string) (*Association, error) {
	wh, args := where.Build(tbl.dialect)
	return tbl.SelectOneSA(wh, orderBy, args...)
}

// SelectSA allows Associations to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
func (tbl AssociationTable) SelectSA(where, orderBy string, args ...interface{}) ([]*Association, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", AssociationColumnNames, tbl.prefix, tbl.name, where, orderBy)
	return tbl.Query(query, args...)
}

// Select allows Associations to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
func (tbl AssociationTable) Select(where where.Expression, orderBy string) ([]*Association, error) {
	wh, args := where.Build(tbl.dialect)
	return tbl.SelectSA(wh, orderBy, args...)
}

// CountSA counts Associations in the table that match a 'where' clause.
func (tbl AssociationTable) CountSA(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s%s %s", tbl.prefix, tbl.name, where)
	tbl.logQuery(query, args...)
	row := tbl.db.QueryRowContext(tbl.ctx, query, args...)
	err = row.Scan(&count)
	return count, err
}

// Count counts the Associations in the table that match a 'where' clause.
func (tbl AssociationTable) Count(where where.Expression) (count int64, err error) {
	wh, args := where.Build(tbl.dialect)
	return tbl.CountSA(wh, args...)
}

const AssociationColumnNames = "id, name, quality, ref1, ref2, category"

//--------------------------------------------------------------------------------

// Insert adds new records for the Associations. The Associations have their primary key fields
// set to the new record identifiers.
// The Association.PreInsert(Execer) method will be called, if it exists.
func (tbl AssociationTable) Insert(vv ...*Association) error {
	var params string
	switch tbl.dialect {
	case schema.Postgres:
		params = sAssociationDataColumnParamsPostgres
	default:
		params = sAssociationDataColumnParamsSimple
	}

	query := fmt.Sprintf(sqlInsertAssociation, tbl.prefix, tbl.name, params)
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

		fields, err := sliceAssociationWithoutPk(v)
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

const sqlInsertAssociation = `
INSERT INTO %s%s (
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
func (tbl AssociationTable) UpdateFields(where where.Expression, fields ...sql.NamedArg) (int64, error) {
	query, args := tbl.updateFields(where, fields...)
	return tbl.Exec(query, args...)
}

func (tbl AssociationTable) updateFields(where where.Expression, fields ...sql.NamedArg) (string, []interface{}) {
	list := sqlgen2.NamedArgList(fields)
	assignments := strings.Join(list.Assignments(tbl.dialect, 1), ", ")
	whereClause, wargs := where.Build(tbl.dialect)
	query := fmt.Sprintf("UPDATE %s%s SET %s %s", tbl.prefix, tbl.name, assignments, whereClause)
	args := append(list.Values(), wargs...)
	return query, args
}

//--------------------------------------------------------------------------------

// Update updates records, matching them by primary key. It returns the number of rows affected.
// The Association.PreUpdate(Execer) method will be called, if it exists.
func (tbl AssociationTable) Update(vv ...*Association) (int64, error) {
	var stmt string
	switch tbl.dialect {
	case schema.Postgres:
		stmt = sqlUpdateAssociationByPkPostgres
	default:
		stmt = sqlUpdateAssociationByPkSimple
	}

	query := fmt.Sprintf(stmt, tbl.prefix, tbl.name)

	var count int64
	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPreUpdate); ok {
			hook.PreUpdate()
		}

		args, err := sliceAssociationWithoutPk(v)
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

const sqlUpdateAssociationByPkSimple = `
UPDATE %s%s SET
	name=?,
	quality=?,
	ref1=?,
	ref2=?,
	category=?
WHERE id=?
`

const sqlUpdateAssociationByPkPostgres = `
UPDATE %s%s SET
	name=$2,
	quality=$3,
	ref1=$4,
	ref2=$5,
	category=$6
WHERE id=$1
`

//--------------------------------------------------------------------------------

// Delete deletes one or more rows from the table, given a 'where' clause.
func (tbl AssociationTable) Delete(where where.Expression) (int64, error) {
	query, args := tbl.deleteRows(where)
	return tbl.Exec(query, args...)
}

func (tbl AssociationTable) deleteRows(where where.Expression) (string, []interface{}) {
	whereClause, args := where.Build(tbl.dialect)
	query := fmt.Sprintf("DELETE FROM %s%s %s", tbl.prefix, tbl.name, whereClause)
	return query, args
}

//--------------------------------------------------------------------------------

// scanAssociations reads table records into a slice of values.
func scanAssociations(rows *sql.Rows, firstOnly bool) ([]*Association, error) {
	var err error
	var vv []*Association

	var v0 int64
	var v1 sql.NullString
	var v2 sql.NullString
	var v3 sql.NullInt64
	var v4 sql.NullInt64
	var v5 sql.NullInt64

	for rows.Next() {
		err = rows.Scan(
			&v0,
			&v1,
			&v2,
			&v3,
			&v4,
			&v5,
		)
		if err != nil {
			return vv, err
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

func sliceAssociationWithoutPk(v *Association) ([]interface{}, error) {


	return []interface{}{
		v.Name,
		v.Quality,
		v.Ref1,
		v.Ref2,
		v.Category,

	}, nil
}
