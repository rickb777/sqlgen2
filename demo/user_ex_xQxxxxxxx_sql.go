// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.
// sqlapi v0.60.1; sqlgen v0.78.0

package demo

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/rickb777/sqlapi"
	"github.com/rickb777/sqlapi/driver"
	"github.com/rickb777/sqlapi/require"
	"github.com/rickb777/sqlapi/support"
	"strings"
)

// QUserTabler lists table methods provided by QUserTable.
type QUserTabler interface {
	sqlapi.Table

	// WithPrefix returns a modified QUserTabler with a given table name prefix.
	WithPrefix(pfx string) QUserTabler

	// WithContext returns a modified QUserTabler with a given context.
	WithContext(ctx context.Context) QUserTabler
}

//-------------------------------------------------------------------------------------------------

// QUserQueryer lists query methods provided by QUserTable.
type QUserQueryer interface {
	sqlapi.Table

	// Using returns a modified QUserQueryer using the Execer supplied,
	// which will typically be a transaction (i.e. SqlTx).
	Using(tx sqlapi.Execer) QUserQueryer

	// Transact runs the function provided within a transaction. The transction is committed
	// unless an error occurs.
	Transact(txOptions *pgx.TxOptions, fn func(QUserQueryer) error) error

	// Query is the low-level request method for this table using an SQL query that must return all the columns
	// necessary for User values.
	Query(req require.Requirement, query string, args ...interface{}) ([]*User, error)

	// QueryOneNullString is a low-level access method for one string, returning the first match.
	QueryOneNullString(req require.Requirement, query string, args ...interface{}) (result sql.NullString, err error)

	// QueryOneNullInt64 is a low-level access method for one int64, returning the first match.
	QueryOneNullInt64(req require.Requirement, query string, args ...interface{}) (result sql.NullInt64, err error)

	// QueryOneNullFloat64 is a low-level access method for one float64, returning the first match.
	QueryOneNullFloat64(req require.Requirement, query string, args ...interface{}) (result sql.NullFloat64, err error)
}

//-------------------------------------------------------------------------------------------------

// QUserTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type QUserTable struct {
	sqlapi.CoreTable
	ctx context.Context
	pk  string
}

// Type conformance checks
var _ sqlapi.Table = &QUserTable{}

// NewQUserTable returns a new table instance.
// If a blank table name is supplied, the default name "users" will be used instead.
// The request context is initialised with the background.
func NewQUserTable(name string, d sqlapi.SqlDB) QUserTable {
	if name == "" {
		name = "users"
	}
	return QUserTable{
		CoreTable: sqlapi.CoreTable{
			Nm: sqlapi.TableName{Prefix: "", Name: name},
			Ex: d,
		},
		ctx: context.Background(),
		pk:  "uid",
	}
}

// CopyTableAsQUserTable copies a table instance, retaining the name etc but
// providing methods appropriate for 'User'.
//
// It serves to provide methods appropriate for 'User'. This is most useful when this is used to represent a
// join result. In such cases, there won't be any need for DDL methods, nor Exec, Insert, Update or Delete.
func CopyTableAsQUserTable(origin sqlapi.Table) QUserTable {
	return QUserTable{
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
//func (tbl QUserTable) SetPkColumn(pk string) QUserTabler {
//	tbl.pk = pk
//	return tbl
//}

// WithPrefix sets the table name prefix for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl QUserTable) WithPrefix(pfx string) QUserTabler {
	tbl.Nm.Prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries via this table.
// The result is a modified copy of the table; the original is unchanged.
func (tbl QUserTable) WithContext(ctx context.Context) QUserTabler {
	tbl.ctx = ctx
	return tbl
}

// Ctx gets the current request context.
func (tbl QUserTable) Ctx() context.Context {
	return tbl.ctx
}

// PkColumn gets the column name used as a primary key.
func (tbl QUserTable) PkColumn() string {
	return tbl.pk
}

// Using returns a modified QUserTabler using the the Execer supplied,
// which will typically be a transaction (i.e. SqlTx). This is needed when making multiple
// queries across several tables within a single transaction.
//
// The result is a modified copy of the table; the original is unchanged.
func (tbl QUserTable) Using(tx sqlapi.Execer) QUserQueryer {
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
func (tbl QUserTable) Transact(txOptions *pgx.TxOptions, fn func(QUserQueryer) error) error {
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

func (tbl QUserTable) quotedName() string {
	return tbl.Dialect().Quoter().Quote(tbl.Nm.String())
}

func (tbl QUserTable) quotedNameW(w driver.StringWriter) {
	tbl.Dialect().Quoter().QuoteW(w, tbl.Nm.String())
}

//-------------------------------------------------------------------------------------------------

// NumQUserTableColumns is the total number of columns in QUserTable.
const NumQUserTableColumns = 22

// NumQUserTableDataColumns is the number of columns in QUserTable not including the auto-increment key.
const NumQUserTableDataColumns = 21

// QUserTableColumnNames is the list of columns in QUserTable.
const QUserTableColumnNames = "uid,name,emailaddress,addressid,avatar,role,active,admin,fave,lastupdated,i8,u8,i16,u16,i32,u32,i64,u64,f32,f64,token,secret"

// QUserTableDataColumnNames is the list of data columns in QUserTable.
const QUserTableDataColumnNames = "name,emailaddress,addressid,avatar,role,active,admin,fave,lastupdated,i8,u8,i16,u16,i32,u32,i64,u64,f32,f64,token,secret"

var listOfQUserTableColumnNames = strings.Split(QUserTableColumnNames, ",")

//-------------------------------------------------------------------------------------------------

// Query is the low-level request method for this table. The SQL query must return all the columns necessary for
// User values. Placeholders should be vanilla '?' marks, which will be replaced if necessary according to
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
func (tbl QUserTable) Query(req require.Requirement, query string, args ...interface{}) ([]*User, error) {
	return doQUserTableQueryAndScan(tbl, req, false, query, args)
}

func doQUserTableQueryAndScan(tbl QUserTabler, req require.Requirement, firstOnly bool, query string, args ...interface{}) ([]*User, error) {
	rows, err := support.Query(tbl, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vv, n, err := scanQUsers(query, rows, firstOnly)
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
func (tbl QUserTable) QueryOneNullString(req require.Requirement, query string, args ...interface{}) (result sql.NullString, err error) {
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
func (tbl QUserTable) QueryOneNullInt64(req require.Requirement, query string, args ...interface{}) (result sql.NullInt64, err error) {
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
func (tbl QUserTable) QueryOneNullFloat64(req require.Requirement, query string, args ...interface{}) (result sql.NullFloat64, err error) {
	err = support.QueryOneNullThing(tbl, req, &result, query, args...)
	return result, err
}

// scanQUsers reads rows from the database and returns a slice of corresponding values.
// It also returns a number indicating how many rows were read; this will be larger than the length of the
// slice if reading stopped after the first row.
func scanQUsers(query string, rows sqlapi.SqlRows, firstOnly bool) (vv []*User, n int64, err error) {
	for rows.Next() {
		n++

		var v0 int64
		var v1 string
		var v2 string
		var v3 sql.NullInt64
		var v4 sql.NullString
		var v5 sql.NullString
		var v6 bool
		var v7 bool
		var v8 []byte
		var v9 int64
		var v10 int8
		var v11 uint8
		var v12 int16
		var v13 uint16
		var v14 int32
		var v15 uint32
		var v16 int64
		var v17 uint64
		var v18 float32
		var v19 float64
		var v20 string
		var v21 string

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
			&v17,
			&v18,
			&v19,
			&v20,
			&v21,
		)
		if err != nil {
			return vv, n, errors.Wrap(err, query)
		}

		v := &User{}
		v.Uid = v0
		v.Name = v1
		v.EmailAddress = v2
		if v3.Valid {
			a := v3.Int64
			v.AddressId = &a
		}
		if v4.Valid {
			a := v4.String
			v.Avatar = &a
		}
		if v5.Valid {
			v.Role = new(Role)
			err = v.Role.Scan(v5.String)
			if err != nil {
				return nil, n, errors.Wrap(err, query)
			}
		}
		v.Active = v6
		v.Admin = v7
		err = json.Unmarshal(v8, &v.Fave)
		if err != nil {
			return nil, n, errors.Wrap(err, query)
		}
		v.LastUpdated = v9
		v.Numbers.I8 = v10
		v.Numbers.U8 = v11
		v.Numbers.I16 = v12
		v.Numbers.U16 = v13
		v.Numbers.I32 = v14
		v.Numbers.U32 = v15
		v.Numbers.I64 = v16
		v.Numbers.U64 = v17
		v.Numbers.F32 = v18
		v.Numbers.F64 = v19
		v.token = v20
		v.secret = v21

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
