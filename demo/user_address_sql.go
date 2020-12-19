// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.
// sqlapi v0.57.0-2-gdefb875; sqlgen v0.75.0

package demo

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/rickb777/sqlapi"
	"github.com/rickb777/sqlapi/driver"
	"github.com/rickb777/sqlapi/require"
	"github.com/rickb777/sqlapi/support"
	"strings"
)

// UserAddressJoiner lists table methods provided by UserAddressJoin.
type UserAddressJoiner interface {
	sqlapi.Table

	// WithPrefix returns a modified UserAddressJoiner with a given table name prefix.
	WithPrefix(pfx string) UserAddressJoiner

	// WithContext returns a modified UserAddressJoiner with a given context.
	WithContext(ctx context.Context) UserAddressJoiner
}

//-------------------------------------------------------------------------------------------------

// UserAddressQueryer lists query methods provided by UserAddressJoin.
type UserAddressQueryer interface {
	sqlapi.Table

	// Using returns a modified UserAddressQueryer using the Execer supplied,
	// which will typically be a transaction (i.e. SqlTx).
	Using(tx sqlapi.Execer) UserAddressQueryer

	// Transact runs the function provided within a transaction. The transction is committed
	// unless an error occurs.
	Transact(txOptions *pgx.TxOptions, fn func(UserAddressQueryer) error) error

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

//-------------------------------------------------------------------------------------------------

// UserAddressJoin holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type UserAddressJoin struct {
	sqlapi.CoreTable
	ctx context.Context
	pk  string
}

// Type conformance checks
var _ sqlapi.Table = &UserAddressJoin{}

// NewUserAddressJoin returns a new table instance.
// If a blank table name is supplied, the default name "useraddresses" will be used instead.
// The request context is initialised with the background.
func NewUserAddressJoin(name string, d sqlapi.SqlDB) UserAddressJoin {
	if name == "" {
		name = "useraddresses"
	}
	return UserAddressJoin{
		CoreTable: sqlapi.CoreTable{
			Nm: sqlapi.TableName{Prefix: "", Name: name},
			Ex: d,
		},
		ctx: context.Background(),
		pk:  "uid",
	}
}

// CopyTableAsUserAddressJoin copies a table instance, retaining the name etc but
// providing methods appropriate for 'UserAddress'.
//
// It serves to provide methods appropriate for 'UserAddress'. This is most useful when this is used to represent a
// join result. In such cases, there won't be any need for DDL methods, nor Exec, Insert, Update or Delete.
func CopyTableAsUserAddressJoin(origin sqlapi.Table) UserAddressJoin {
	return UserAddressJoin{
		CoreTable: sqlapi.CoreTable{
			Nm: origin.Name(),
			Ex: origin.Execer(),
		},
		ctx: origin.Ctx(),
		pk:  "uid",
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
	tbl.Nm.Prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries via this table.
// The result is a modified copy of the table; the original is unchanged.
func (tbl UserAddressJoin) WithContext(ctx context.Context) UserAddressJoiner {
	tbl.ctx = ctx
	return tbl
}

// Ctx gets the current request context.
func (tbl UserAddressJoin) Ctx() context.Context {
	return tbl.ctx
}

// PkColumn gets the column name used as a primary key.
func (tbl UserAddressJoin) PkColumn() string {
	return tbl.pk
}

// Using returns a modified UserAddressJoiner using the the Execer supplied,
// which will typically be a transaction (i.e. SqlTx). This is needed when making multiple
// queries across several tables within a single transaction.
//
// The result is a modified copy of the table; the original is unchanged.
func (tbl UserAddressJoin) Using(tx sqlapi.Execer) UserAddressQueryer {
	tbl.Ex = tx
	return tbl
}

// Transact runs the function provided within a transaction. If the function completes without error,
// the transaction is committed. If there is an error or a panic, the transaction is rolled back.
//
// The options can be nil, in which case the default behaviour is that of the underlying connection.
//
// Nested transactions (i.e. within 'fn') are permitted: they execute within the outermost transaction.
// Therefore they do not commit until the outermost transaction commits.
func (tbl UserAddressJoin) Transact(txOptions *pgx.TxOptions, fn func(UserAddressQueryer) error) error {
	var err error
	if tbl.IsTx() {
		err = fn(tbl) // nested transactions are inlined
	} else {
		err = tbl.DB().Transact(tbl.ctx, txOptions, func(tx sqlapi.SqlTx) error {
			return fn(tbl.Using(tx))
		})
	}
	return tbl.Logger().LogIfError(tbl.Ctx(), err)
}

func (tbl UserAddressJoin) quotedName() string {
	return tbl.Dialect().Quoter().Quote(tbl.Nm.String())
}

func (tbl UserAddressJoin) quotedNameW(w driver.StringWriter) {
	tbl.Dialect().Quoter().QuoteW(w, tbl.Nm.String())
}

//-------------------------------------------------------------------------------------------------

// NumUserAddressJoinColumns is the total number of columns in UserAddressJoin.
const NumUserAddressJoinColumns = 13

// NumUserAddressJoinDataColumns is the number of columns in UserAddressJoin not including the auto-increment key.
const NumUserAddressJoinDataColumns = 12

// UserAddressJoinColumnNames is the list of columns in UserAddressJoin.
const UserAddressJoinColumnNames = "uid,name,emailaddress,lines,town,postcode,uprn,avatar,role,active,admin,fave,lastupdated"

// UserAddressJoinDataColumnNames is the list of data columns in UserAddressJoin.
const UserAddressJoinDataColumnNames = "name,emailaddress,lines,town,postcode,uprn,avatar,role,active,admin,fave,lastupdated"

var listOfUserAddressJoinColumnNames = strings.Split(UserAddressJoinColumnNames, ",")

//-------------------------------------------------------------------------------------------------

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
	return vv, tbl.Logger().LogIfError(tbl.Ctx(), require.ChainErrorIfQueryNotSatisfiedBy(err, req, n))
}

//-------------------------------------------------------------------------------------------------

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
		var v6 string
		var v7 sql.NullString
		var v8 sql.NullString
		var v9 bool
		var v10 bool
		var v11 []byte
		var v12 int64

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
		v.Address.UPRN = v6
		if v7.Valid {
			a := v7.String
			v.Avatar = &a
		}
		if v8.Valid {
			v.Role = new(Role)
			err = v.Role.Scan(v8.String)
			if err != nil {
				return nil, n, errors.Wrap(err, query)
			}
		}
		v.Active = v9
		v.Admin = v10
		err = json.Unmarshal(v11, &v.Fave)
		if err != nil {
			return nil, n, errors.Wrap(err, query)
		}
		v.LastUpdated = v12

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
