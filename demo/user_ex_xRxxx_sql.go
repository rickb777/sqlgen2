// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

package demo

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/rickb777/sqlgen2"
	"github.com/rickb777/sqlgen2/constraint"
	"github.com/rickb777/sqlgen2/model"
	"github.com/rickb777/sqlgen2/require"
	"github.com/rickb777/sqlgen2/schema"
	"github.com/rickb777/sqlgen2/support"
	"github.com/rickb777/sqlgen2/where"
	"log"
	"strings"
)

// RUserTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type RUserTable struct {
	name        model.TableName
	db          sqlgen2.Execer
	constraints constraint.Constraints
	ctx         context.Context
	dialect     schema.Dialect
	logger      *log.Logger
	wrapper     interface{}
}

// Type conformance checks
var _ sqlgen2.Table = &RUserTable{}
var _ sqlgen2.Table = &RUserTable{}

// NewRUserTable returns a new table instance.
// If a blank table name is supplied, the default name "users" will be used instead.
// The request context is initialised with the background.
func NewRUserTable(name model.TableName, d sqlgen2.Execer, dialect schema.Dialect) RUserTable {
	if name.Name == "" {
		name.Name = "users"
	}
	table := RUserTable{name, d, nil, context.Background(), dialect, nil, nil}
	table.constraints = append(table.constraints,
		constraint.FkConstraint{"addressid", constraint.Reference{"addresses", "id"}, "restrict", "restrict"})
	
	return table
}

// CopyTableAsRUserTable copies a table instance, retaining the name etc but
// providing methods appropriate for 'User'. It doesn't copy the constraints of the original table.
//
// It serves to provide methods appropriate for 'User'. This is most useulf when thie is used to represent a
// join result. In such cases, there won't be any need for DDL methods, nor Exec, Insert, Update or Delete.
func CopyTableAsRUserTable(origin sqlgen2.Table) RUserTable {
	return RUserTable{
		name:        origin.Name(),
		db:          origin.DB(),
		constraints: nil,
		ctx:         origin.Ctx(),
		dialect:     origin.Dialect(),
		logger:      origin.Logger(),
	}
}

// WithPrefix sets the table name prefix for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl RUserTable) WithPrefix(pfx string) RUserTable {
	tbl.name.Prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl RUserTable) WithContext(ctx context.Context) RUserTable {
	tbl.ctx = ctx
	return tbl
}

// WithLogger sets the logger for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl RUserTable) WithLogger(logger *log.Logger) RUserTable {
	tbl.logger = logger
	return tbl
}

// Logger gets the trace logger.
func (tbl RUserTable) Logger() *log.Logger {
	return tbl.logger
}

// SetLogger sets the logger for subsequent queries, returning the interface.
// The result is a modified copy of the table; the original is unchanged.
func (tbl RUserTable) SetLogger(logger *log.Logger) sqlgen2.Table {
	tbl.logger = logger
	return tbl
}

// WithConstraint returns a modified Table with added data consistency constraints.
func (tbl RUserTable) WithConstraint(cc ...constraint.Constraint) RUserTable {
	tbl.constraints = append(tbl.constraints, cc...)
	return tbl
}

// Ctx gets the current request context.
func (tbl RUserTable) Ctx() context.Context {
	return tbl.ctx
}

// Dialect gets the database dialect.
func (tbl RUserTable) Dialect() schema.Dialect {
	return tbl.dialect
}

// Wrapper gets the user-defined wrapper.
func (tbl RUserTable) Wrapper() interface{} {
	return tbl.wrapper
}

// SetWrapper sets the user-defined wrapper.
// The result is a modified copy of the table; the original is unchanged.
func (tbl RUserTable) SetWrapper(wrapper interface{}) sqlgen2.Table {
	tbl.wrapper = wrapper
	return tbl
}

// Name gets the table name.
func (tbl RUserTable) Name() model.TableName {
	return tbl.name
}

// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl RUserTable) DB() *sql.DB {
	return tbl.db.(*sql.DB)
}

// Execer gets the wrapped database or transaction handle.
func (tbl RUserTable) Execer() sqlgen2.Execer {
	return tbl.db
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl RUserTable) Tx() *sql.Tx {
	return tbl.db.(*sql.Tx)
}

// IsTx tests whether this is within a transaction.
func (tbl RUserTable) IsTx() bool {
	_, ok := tbl.db.(*sql.Tx)
	return ok
}

// Begin starts a transaction. The default isolation level is dependent on the driver.
// The result is a modified copy of the table; the original is unchanged.
func (tbl RUserTable) BeginTx(opts *sql.TxOptions) (RUserTable, error) {
	d := tbl.db.(*sql.DB)
	var err error
	tbl.db, err = d.BeginTx(tbl.ctx, opts)
	return tbl, tbl.logIfError(err)
}

// Using returns a modified Table using the transaction supplied. This is needed
// when making multiple queries across several tables within a single transaction.
// The result is a modified copy of the table; the original is unchanged.
func (tbl RUserTable) Using(tx *sql.Tx) RUserTable {
	tbl.db = tx
	return tbl
}

func (tbl RUserTable) logQuery(query string, args ...interface{}) {
	support.LogQuery(tbl.logger, query, args...)
}

func (tbl RUserTable) logError(err error) error {
	return support.LogError(tbl.logger, err)
}

func (tbl RUserTable) logIfError(err error) error {
	return support.LogIfError(tbl.logger, err)
}


//--------------------------------------------------------------------------------

const NumRUserColumns = 12

const NumRUserDataColumns = 11

const RUserColumnNames = "uid,login,emailaddress,addressid,avatar,role,active,admin,fave,lastupdated,token,secret"

const RUserDataColumnNames = "login,emailaddress,addressid,avatar,role,active,admin,fave,lastupdated,token,secret"

const RUserPk = "uid"

//--------------------------------------------------------------------------------

// Query is the low-level access method for Users.
//
// It places a requirement, which may be nil, on the size of the expected results: this
// controls whether an error is generated when this expectation is not met.
//
// Note that this method applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl RUserTable) Query(req require.Requirement, query string, args ...interface{}) ([]*User, error) {
	query = tbl.ReplaceTableName(query)
	vv, err := tbl.doQuery(req, false, query, args...)
	return vv, err
}

// QueryOne is the low-level access method for one User.
// If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, *User will be nil.
//
// Note that this method applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl RUserTable) QueryOne(query string, args ...interface{}) (*User, error) {
	query = tbl.ReplaceTableName(query)
	return tbl.doQueryOne(nil, query, args...)
}

// MustQueryOne is the low-level access method for one User.
//
// It places a requirement that exactly one result must be found; an error is generated when this expectation is not met.
//
// Note that this method applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl RUserTable) MustQueryOne(query string, args ...interface{}) (*User, error) {
	query = tbl.ReplaceTableName(query)
	return tbl.doQueryOne(require.One, query, args...)
}

func (tbl RUserTable) doQueryOne(req require.Requirement, query string, args ...interface{}) (*User, error) {
	list, err := tbl.doQuery(req, true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

func (tbl RUserTable) doQuery(req require.Requirement, firstOnly bool, query string, args ...interface{}) ([]*User, error) {
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, tbl.logError(err)
	}
	defer rows.Close()

	vv, n, err := scanRUsers(rows, firstOnly)
	return vv, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(err, req, n))
}

func scanRUsers(rows *sql.Rows, firstOnly bool) (vv []*User, n int64, err error) {
	for rows.Next() {
		n++

		var v0 int64
		var v1 string
		var v2 string
		var v3 sql.NullInt64
		var v4 sql.NullString
		var v5 *Role
		var v6 bool
		var v7 bool
		var v8 []byte
		var v9 int64
		var v10 string
		var v11 string

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
		)
		if err != nil {
			return vv, n, err
		}

		v := &User{}
		v.Uid = v0
		v.Login = v1
		v.EmailAddress = v2
		if v3.Valid {
			a := v3.Int64
			v.AddressId = &a
		}
		if v4.Valid {
			a := v4.String
			v.Avatar = &a
		}
		v.Role = v5
		v.Active = v6
		v.Admin = v7
		err = json.Unmarshal(v8, &v.Fave)
		if err != nil {
			return nil, n, err
		}
		v.LastUpdated = v9
		v.token = v10
		v.secret = v11

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
func (tbl RUserTable) QueryOneNullString(query string, args ...interface{}) (result sql.NullString, err error) {
	err = support.QueryOneNullThing(tbl, nil, &result, query, args...)
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
func (tbl RUserTable) MustQueryOneNullString(query string, args ...interface{}) (result sql.NullString, err error) {
	err = support.QueryOneNullThing(tbl, require.One, &result, query, args...)
	return result, err
}

// QueryOneNullInt64 is a low-level access method for one int64. This can be used for 'COUNT(1)' queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl RUserTable) QueryOneNullInt64(query string, args ...interface{}) (result sql.NullInt64, err error) {
	err = support.QueryOneNullThing(tbl, nil, &result, query, args...)
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
func (tbl RUserTable) MustQueryOneNullInt64(query string, args ...interface{}) (result sql.NullInt64, err error) {
	err = support.QueryOneNullThing(tbl, require.One, &result, query, args...)
	return result, err
}

// QueryOneNullFloat64 is a low-level access method for one float64. This can be used for 'AVG(...)' queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl RUserTable) QueryOneNullFloat64(query string, args ...interface{}) (result sql.NullFloat64, err error) {
	err = support.QueryOneNullThing(tbl, nil, &result, query, args...)
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
func (tbl RUserTable) MustQueryOneNullFloat64(query string, args ...interface{}) (result sql.NullFloat64, err error) {
	err = support.QueryOneNullThing(tbl, require.One, &result, query, args...)
	return result, err
}

// ReplaceTableName replaces all occurrences of "{TABLE}" with the table's name.
func (tbl RUserTable) ReplaceTableName(query string) string {
	return strings.Replace(query, "{TABLE}", tbl.name.String(), -1)
}

//--------------------------------------------------------------------------------

var allRUserQuotedColumnNames = []string{
	schema.Sqlite.SplitAndQuote(RUserColumnNames),
	schema.Mysql.SplitAndQuote(RUserColumnNames),
	schema.Postgres.SplitAndQuote(RUserColumnNames),
}

//--------------------------------------------------------------------------------

// GetUser gets the record with a given primary key value.
// If not found, *User will be nil.
func (tbl RUserTable) GetUser(id int64) (*User, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s=?",
		allRUserQuotedColumnNames[tbl.dialect.Index()], tbl.name, tbl.dialect.Quote("uid"))
	v, err := tbl.doQueryOne(nil, query, id)
	return v, err
}

// MustGetUser gets the record with a given primary key value.
//
// It places a requirement that exactly one result must be found; an error is generated when this expectation is not met.
func (tbl RUserTable) MustGetUser(id int64) (*User, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s=?",
		allRUserQuotedColumnNames[tbl.dialect.Index()], tbl.name, tbl.dialect.Quote("uid"))
	v, err := tbl.doQueryOne(require.One, query, id)
	return v, err
}

// GetUsers gets records from the table according to a list of primary keys.
// Although the list of ids can be arbitrarily long, there are practical limits;
// note that Oracle DB has a limit of 1000.
//
// It places a requirement, which may be nil, on the size of the expected results: in particular, require.All
// controls whether an error is generated not all the ids produce a result.
func (tbl RUserTable) GetUsers(req require.Requirement, id ...int64) (list []*User, err error) {
	if len(id) > 0 {
		if req == require.All {
			req = require.Exactly(len(id))
		}
		pl := tbl.dialect.Placeholders(len(id))
		query := fmt.Sprintf("SELECT %s FROM %s WHERE %s IN (%s)",
			allRUserQuotedColumnNames[tbl.dialect.Index()], tbl.name, tbl.dialect.Quote("uid"), pl)
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
func (tbl RUserTable) SelectOneWhere(req require.Requirement, where, orderBy string, args ...interface{}) (*User, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT 1", RUserColumnNames, tbl.name, where, orderBy)
	v, err := tbl.doQueryOne(req, query, args...)
	return v, err
}

// SelectOne allows a single User to be obtained from the sqlgen2.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
// If not found, *Example will be nil.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.One
// controls whether an error is generated when no result is found.
func (tbl RUserTable) SelectOne(req require.Requirement, wh where.Expression, qc where.QueryConstraint) (*User, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	return tbl.SelectOneWhere(req, whs, orderBy, args...)
}

// SelectWhere allows Users to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
//
// The args are for any placeholder parameters in the query.
func (tbl RUserTable) SelectWhere(req require.Requirement, where, orderBy string, args ...interface{}) ([]*User, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", RUserColumnNames, tbl.name, where, orderBy)
	vv, err := tbl.doQuery(req, false, query, args...)
	return vv, err
}

// Select allows Users to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
func (tbl RUserTable) Select(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]*User, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	return tbl.SelectWhere(req, whs, orderBy, args...)
}

// CountWhere counts Users in the table that match a 'where' clause.
// Use a blank string for the 'where' argument if it is not needed.
//
// The args are for any placeholder parameters in the query.
func (tbl RUserTable) CountWhere(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s %s", tbl.name, where)
	tbl.logQuery(query, args...)
	row := tbl.db.QueryRowContext(tbl.ctx, query, args...)
	err = row.Scan(&count)
	return count, tbl.logIfError(err)
}

// Count counts the Users in the table that match a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed.
func (tbl RUserTable) Count(wh where.Expression) (count int64, err error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	return tbl.CountWhere(whs, args...)
}
