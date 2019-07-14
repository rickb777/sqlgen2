// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.
// sqlapi v0.32.0; sqlgen v0.55.0

package demo

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/rickb777/sqlapi"
	"github.com/rickb777/sqlapi/constraint"
	"github.com/rickb777/sqlapi/dialect"
	"github.com/rickb777/sqlapi/require"
	"github.com/rickb777/sqlapi/support"
	"strings"
)

// UserAddressJoiner lists methods provided by UserAddressJoin.
type UserAddressJoiner interface {
	sqlapi.Table

	// Constraints returns the table's constraints.
	Constraints() constraint.Constraints

	// WithConstraint returns a modified UserAddressJoiner with added data consistency constraints.
	WithConstraint(cc ...constraint.Constraint) UserAddressJoiner

	// WithPrefix returns a modified UserAddressJoiner with a given table name prefix.
	WithPrefix(pfx string) UserAddressJoiner

	// WithContext returns a modified UserAddressJoiner with a given context.
	WithContext(ctx context.Context) UserAddressJoiner

	// Using returns a modified UserAddressJoiner using the transaction supplied.
	Using(tx sqlapi.SqlTx) UserAddressJoiner

	// Transact runs the function provided within a transaction.
	Transact(txOptions *sql.TxOptions, fn func(UserAddressJoiner) error) error

	// Query is the low-level request method for this table using an SQL query that must return all the columns
	// necessary for UserAddress values.
	Query(req require.Requirement, query string, args ...interface{}) ([]*UserAddress, error)

	// QueryOneNullString is a low-level access method for one string, returning the first match.
	QueryOneNullString(req require.Requirement, query string, args ...interface{}) (result sql.NullString, err error)

	// QueryOneNullInt64 is a low-level access method for one int64, returning the first match.
	QueryOneNullInt64(req require.Requirement, query string, args ...interface{}) (result sql.NullInt64, err error)

	// QueryOneNullFloat64 is a low-level access method for one float64, returning the first match.
	QueryOneNullFloat64(req require.Requirement, query string, args ...interface{}) (result sql.NullFloat64, err error)
}

// UserAddressJoin holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type UserAddressJoin struct {
	name        sqlapi.TableName
	database    sqlapi.Database
	db          sqlapi.Execer
	constraints constraint.Constraints
	ctx         context.Context
	pk          string
}

// Type conformance checks
var _ sqlapi.Table = &UserAddressJoin{}

// NewUserAddressJoin returns a new table instance.
// If a blank table name is supplied, the default name "useraddresses" will be used instead.
// The request context is initialised with the background.
func NewUserAddressJoin(name string, d sqlapi.Database) UserAddressJoin {
	if name == "" {
		name = "useraddresses"
	}
	var constraints constraint.Constraints
	return UserAddressJoin{
		name:        sqlapi.TableName{Prefix: "", Name: name},
		database:    d,
		db:          d.DB(),
		constraints: constraints,
		ctx:         context.Background(),
		pk:          "uid",
	}
}

// CopyTableAsUserAddressJoin copies a table instance, retaining the name etc but
// providing methods appropriate for 'UserAddress'. It doesn't copy the constraints of the original table.
//
// It serves to provide methods appropriate for 'UserAddress'. This is most useful when this is used to represent a
// join result. In such cases, there won't be any need for DDL methods, nor Exec, Insert, Update or Delete.
func CopyTableAsUserAddressJoin(origin sqlapi.Table) UserAddressJoin {
	return UserAddressJoin{
		name:        origin.Name(),
		database:    origin.Database(),
		db:          origin.DB(),
		constraints: nil,
		ctx:         context.Background(),
		pk:          "uid",
	}
}

// SetPkColumn sets the name of the primary key column. It defaults to "uid".
// The result is a modified copy of the table; the original is unchanged.
//func (tbl UserAddressJoin) SetPkColumn(pk string) UserAddressJoiner {
//	tbl.pk = pk
//	return tbl
//}

// WithPrefix sets the table name prefix for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl UserAddressJoin) WithPrefix(pfx string) UserAddressJoiner {
	tbl.name.Prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries via this table.
// The result is a modified copy of the table; the original is unchanged.
//
// The shared context in the *Database is not altered by this method. So it
// is possible to use different contexts for different (groups of) queries.
func (tbl UserAddressJoin) WithContext(ctx context.Context) UserAddressJoiner {
	tbl.ctx = ctx
	return tbl
}

// Database gets the shared database information.
func (tbl UserAddressJoin) Database() sqlapi.Database {
	return tbl.database
}

// Logger gets the trace logger.
func (tbl UserAddressJoin) Logger() sqlapi.Logger {
	return tbl.database.Logger()
}

// WithConstraint returns a modified UserAddressJoiner with added data consistency constraints.
func (tbl UserAddressJoin) WithConstraint(cc ...constraint.Constraint) UserAddressJoiner {
	tbl.constraints = append(tbl.constraints, cc...)
	return tbl
}

// Constraints returns the table's constraints.
func (tbl UserAddressJoin) Constraints() constraint.Constraints {
	return tbl.constraints
}

// Ctx gets the current request context.
func (tbl UserAddressJoin) Ctx() context.Context {
	return tbl.ctx
}

// Dialect gets the database dialect.
func (tbl UserAddressJoin) Dialect() dialect.Dialect {
	return tbl.database.Dialect()
}

// Name gets the table name.
func (tbl UserAddressJoin) Name() sqlapi.TableName {
	return tbl.name
}

// PkColumn gets the column name used as a primary key.
func (tbl UserAddressJoin) PkColumn() string {
	return tbl.pk
}

// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl UserAddressJoin) DB() sqlapi.SqlDB {
	return tbl.db.(sqlapi.SqlDB)
}

// Execer gets the wrapped database or transaction handle.
func (tbl UserAddressJoin) Execer() sqlapi.Execer {
	return tbl.db
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl UserAddressJoin) Tx() sqlapi.SqlTx {
	return tbl.db.(sqlapi.SqlTx)
}

// IsTx tests whether this is within a transaction.
func (tbl UserAddressJoin) IsTx() bool {
	return tbl.db.IsTx()
}

// Using returns a modified UserAddressJoiner using the transaction supplied. This is needed
// when making multiple queries across several tables within a single transaction.
// The result is a modified copy of the table; the original is unchanged.
func (tbl UserAddressJoin) Using(tx sqlapi.SqlTx) UserAddressJoiner {
	tbl.db = tx
	return tbl
}

// Transact runs the function provided within a transaction. If the function completes without error,
// the transaction is committed. If there is an error or a panic, the transaction is rolled back.
//
// Nested transactions (i.e. within 'fn') are permitted: they execute within the outermost transaction.
// Therefore they do not commit until the outermost transaction commits.
func (tbl UserAddressJoin) Transact(txOptions *sql.TxOptions, fn func(UserAddressJoiner) error) error {
	var err error
	if tbl.IsTx() {
		err = fn(tbl) // nested transactions are inlined
	} else {
		err = tbl.DB().Transact(tbl.ctx, txOptions, func(tx sqlapi.SqlTx) error {
			return fn(tbl.Using(tx))
		})
	}
	return tbl.Logger().LogIfError(err)
}

func (tbl UserAddressJoin) quotedName() string {
	return tbl.Dialect().Quoter().Quote(tbl.name.String())
}

func (tbl UserAddressJoin) quotedNameW(w dialect.StringWriter) {
	tbl.Dialect().Quoter().QuoteW(w, tbl.name.String())
}

//--------------------------------------------------------------------------------

// NumUserAddressJoinColumns is the total number of columns in UserAddressJoin.
const NumUserAddressJoinColumns = 12

// NumUserAddressJoinDataColumns is the number of columns in UserAddressJoin not including the auto-increment key.
const NumUserAddressJoinDataColumns = 11

// UserAddressJoinColumnNames is the list of columns in UserAddressJoin.
const UserAddressJoinColumnNames = "uid,name,emailaddress,lines,town,postcode,avatar,role,active,admin,fave,lastupdated"

// UserAddressJoinDataColumnNames is the list of data columns in UserAddressJoin.
const UserAddressJoinDataColumnNames = "name,emailaddress,lines,town,postcode,avatar,role,active,admin,fave,lastupdated"

var listOfUserAddressJoinColumnNames = strings.Split(UserAddressJoinColumnNames, ",")

//--------------------------------------------------------------------------------

// Query is the low-level request method for this table. The SQL query must return all the columns necessary for
// UserAddress values. Placeholders should be vanilla '?' marks, which will be replaced if necessary according to
// the chosen dialect.
//
// The query is logged using whatever logger is configured. If an error arises, this too is logged.
//
// If you need a context other than the background, use WithContext before calling Query.
//
// The args are for any placeholder parameters in the query.
//
// The support API provides a core 'support.Query' function, on which this method depends. If appropriate,
// use that function directly; wrap the result in *sqlapi.Rows if you need to access its data as a map.
func (tbl UserAddressJoin) Query(req require.Requirement, query string, args ...interface{}) ([]*UserAddress, error) {
	return doUserAddressJoinQueryAndScan(tbl, req, false, query, args)
}

func doUserAddressJoinQueryAndScan(tbl UserAddressJoiner, req require.Requirement, firstOnly bool, query string, args ...interface{}) ([]*UserAddress, error) {
	rows, err := support.Query(tbl, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vv, n, err := scanUserAddresses(query, rows, firstOnly)
	return vv, tbl.Logger().LogIfError(require.ChainErrorIfQueryNotSatisfiedBy(err, req, n))
}

//--------------------------------------------------------------------------------

// QueryOneNullString is a low-level access method for one string. This can be used for function queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl UserAddressJoin) QueryOneNullString(req require.Requirement, query string, args ...interface{}) (result sql.NullString, err error) {
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
func (tbl UserAddressJoin) QueryOneNullInt64(req require.Requirement, query string, args ...interface{}) (result sql.NullInt64, err error) {
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
func (tbl UserAddressJoin) QueryOneNullFloat64(req require.Requirement, query string, args ...interface{}) (result sql.NullFloat64, err error) {
	err = support.QueryOneNullThing(tbl, req, &result, query, args...)
	return result, err
}

// scanUserAddresses reads rows from the database and returns a slice of corresponding values.
// It also returns a number indicating how many rows were read; this will be larger than the length of the
// slice if reading stopped after the first row.
func scanUserAddresses(query string, rows sqlapi.SqlRows, firstOnly bool) (vv []*UserAddress, n int64, err error) {
	for rows.Next() {
		n++

		var v0 int64
		var v1 string
		var v2 string
		var v3 []byte
		var v4 sql.NullString
		var v5 string
		var v6 sql.NullString
		var v7 sql.NullString
		var v8 bool
		var v9 bool
		var v10 []byte
		var v11 int64

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
			return vv, n, errors.Wrap(err, query)
		}

		v := &UserAddress{}
		v.Uid = v0
		v.Name = v1
		v.EmailAddress = v2
		err = json.Unmarshal(v3, &v.Address.Lines)
		if err != nil {
			return nil, n, errors.Wrap(err, query)
		}
		if v4.Valid {
			a := v4.String
			v.Address.Town = &a
		}
		v.Address.Postcode = v5
		if v6.Valid {
			a := v6.String
			v.Avatar = &a
		}
		if v7.Valid {
			v.Role = new(Role)
			err = v.Role.Scan(v7.String)
			if err != nil {
				return nil, n, errors.Wrap(err, query)
			}
		}
		v.Active = v8
		v.Admin = v9
		err = json.Unmarshal(v10, &v.Fave)
		if err != nil {
			return nil, n, errors.Wrap(err, query)
		}
		v.LastUpdated = v11

		var iv interface{} = v
		if hook, ok := iv.(sqlapi.CanPostGet); ok {
			err = hook.PostGet()
			if err != nil {
				return vv, n, errors.Wrap(err, query)
			}
		}

		vv = append(vv, v)

		if firstOnly {
			if rows.Next() {
				n++
			}
			return vv, n, errors.Wrap(rows.Err(), query)
		}
	}

	return vv, n, errors.Wrap(rows.Err(), query)
}
