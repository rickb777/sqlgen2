// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.
// sqlapi v0.62.7; sqlgen v0.79.0-1-g994aea0

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

// EUserTabler lists table methods provided by EUserTable.
type EUserTabler interface {
	sqlapi.Table

	// WithPrefix returns a modified EUserTabler with a given table name prefix.
	WithPrefix(pfx string) EUserTabler

	// WithContext returns a modified EUserTabler with a given context.
	WithContext(ctx context.Context) EUserTabler
}

//-------------------------------------------------------------------------------------------------

// EUserQueryer lists query methods provided by EUserTable.
type EUserQueryer interface {
	sqlapi.Table

	// Using returns a modified EUserQueryer using the Execer supplied,
	// which will typically be a transaction (i.e. SqlTx).
	Using(tx sqlapi.Execer) EUserQueryer

	// Transact runs the function provided within a transaction. The transction is committed
	// unless an error occurs.
	Transact(txOptions *pgx.TxOptions, fn func(EUserQueryer) error) error

	// Exec executes a query without returning any rows.
	Exec(req require.Requirement, query string, args ...interface{}) (int64, error)
}

//-------------------------------------------------------------------------------------------------

// EUserTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type EUserTable struct {
	sqlapi.CoreTable
	ctx context.Context
	pk  string
}

// Type conformance checks
var _ sqlapi.Table = &EUserTable{}

// NewEUserTable returns a new table instance.
// If a blank table name is supplied, the default name "users" will be used instead.
// The request context is initialised with the background.
func NewEUserTable(name string, d sqlapi.SqlDB) EUserTable {
	if name == "" {
		name = "users"
	}
	return EUserTable{
		CoreTable: sqlapi.CoreTable{
			Nm: sqlapi.TableName{Prefix: "", Name: name},
			Ex: d,
		},
		ctx: context.Background(),
		pk:  "uid",
	}
}

// CopyTableAsEUserTable copies a table instance, retaining the name etc but
// providing methods appropriate for 'User'.
//
// It serves to provide methods appropriate for 'User'. This is most useful when this is used to represent a
// join result. In such cases, there won't be any need for DDL methods, nor Exec, Insert, Update or Delete.
func CopyTableAsEUserTable(origin sqlapi.Table) EUserTable {
	return EUserTable{
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
//func (tbl EUserTable) SetPkColumn(pk string) EUserTabler {
//	tbl.pk = pk
//	return tbl
//}

// WithPrefix sets the table name prefix for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl EUserTable) WithPrefix(pfx string) EUserTabler {
	tbl.Nm.Prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries via this table.
// The result is a modified copy of the table; the original is unchanged.
func (tbl EUserTable) WithContext(ctx context.Context) EUserTabler {
	tbl.ctx = ctx
	return tbl
}

// Ctx gets the current request context.
func (tbl EUserTable) Ctx() context.Context {
	return tbl.ctx
}

// PkColumn gets the column name used as a primary key.
func (tbl EUserTable) PkColumn() string {
	return tbl.pk
}

// Using returns a modified EUserTabler using the the Execer supplied,
// which will typically be a transaction (i.e. SqlTx). This is needed when making multiple
// queries across several tables within a single transaction.
//
// The result is a modified copy of the table; the original is unchanged.
func (tbl EUserTable) Using(tx sqlapi.Execer) EUserQueryer {
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
func (tbl EUserTable) Transact(txOptions *pgx.TxOptions, fn func(EUserQueryer) error) error {
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

func (tbl EUserTable) quotedName() string {
	return tbl.Dialect().Quoter().Quote(tbl.Nm.String())
}

func (tbl EUserTable) quotedNameW(w driver.StringWriter) {
	tbl.Dialect().Quoter().QuoteW(w, tbl.Nm.String())
}

//-------------------------------------------------------------------------------------------------

// NumEUserTableColumns is the total number of columns in EUserTable.
const NumEUserTableColumns = 22

// NumEUserTableDataColumns is the number of columns in EUserTable not including the auto-increment key.
const NumEUserTableDataColumns = 21

// EUserTableColumnNames is the list of columns in EUserTable.
const EUserTableColumnNames = "uid,name,emailaddress,addressid,avatar,role,active,admin,fave,lastupdated,i8,u8,i16,u16,i32,u32,i64,u64,f32,f64,token,secret"

// EUserTableDataColumnNames is the list of data columns in EUserTable.
const EUserTableDataColumnNames = "name,emailaddress,addressid,avatar,role,active,admin,fave,lastupdated,i8,u8,i16,u16,i32,u32,i64,u64,f32,f64,token,secret"

var listOfEUserTableColumnNames = strings.Split(EUserTableColumnNames, ",")

//-------------------------------------------------------------------------------------------------

// Exec executes a query without returning any rows.
// It returns the number of rows affected (if the database driver supports this).
//
// The args are for any placeholder parameters in the query.
func (tbl EUserTable) Exec(req require.Requirement, query string, args ...interface{}) (int64, error) {
	return support.Exec(tbl, req, query, args...)
}

// scanEUsers reads rows from the database and returns a slice of corresponding values.
// It also returns a number indicating how many rows were read; this will be larger than the length of the
// slice if reading stopped after the first row.
func scanEUsers(query string, rows sqlapi.SqlRows, firstOnly bool) (vv []*User, n int64, err error) {
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
