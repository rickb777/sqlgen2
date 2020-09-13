// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.
// sqlapi v0.47.0; sqlgen v0.66.0

package demo

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/rickb777/sqlapi"
	"github.com/rickb777/sqlapi/dialect"
	"strings"
)

// XUserTabler lists table methods provided by XUserTable.
type XUserTabler interface {
	sqlapi.Table

	// WithPrefix returns a modified XUserTabler with a given table name prefix.
	WithPrefix(pfx string) XUserTabler
}

//-------------------------------------------------------------------------------------------------

// XUserQueryer lists query methods provided by XUserTable.
type XUserQueryer interface {
	sqlapi.Table

	// Using returns a modified XUserQueryer using the Execer supplied,
	// which will typically be a transaction (i.e. SqlTx).
	Using(tx sqlapi.Execer) XUserQueryer

	// Transact runs the function provided within a transaction. The transction is committed
	// unless an error occurs.
	Transact(ctx context.Context, txOptions *sql.TxOptions, fn func(XUserQueryer) error) error
}

//-------------------------------------------------------------------------------------------------

// XUserTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type XUserTable struct {
	name     sqlapi.TableName
	database sqlapi.Database
	db       sqlapi.Execer
	pk       string
}

// Type conformance checks
var _ sqlapi.Table = &XUserTable{}

// NewXUserTable returns a new table instance.
// If a blank table name is supplied, the default name "users" will be used instead.
// The request context is initialised with the background.
func NewXUserTable(name string, d sqlapi.Database) XUserTable {
	if name == "" {
		name = "users"
	}
	return XUserTable{
		name:     sqlapi.TableName{Prefix: "", Name: name},
		database: d,
		db:       d.DB(),
		pk:       "uid",
	}
}

// CopyTableAsXUserTable copies a table instance, retaining the name etc but
// providing methods appropriate for 'User'.
//
// It serves to provide methods appropriate for 'User'. This is most useful when this is used to represent a
// join result. In such cases, there won't be any need for DDL methods, nor Exec, Insert, Update or Delete.
func CopyTableAsXUserTable(origin sqlapi.Table) XUserTable {
	return XUserTable{
		name:     origin.Name(),
		database: origin.Database(),
		db:       origin.Execer(),
		pk:       "uid",
	}
}

// SetPkColumn sets the name of the primary key column. It defaults to "uid".
// The result is a modified copy of the table; the original is unchanged.
//func (tbl XUserTable) SetPkColumn(pk string) XUserTabler {
//	tbl.pk = pk
//	return tbl
//}

// WithPrefix sets the table name prefix for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl XUserTable) WithPrefix(pfx string) XUserTabler {
	tbl.name.Prefix = pfx
	return tbl
}

// Database gets the shared database information.
func (tbl XUserTable) Database() sqlapi.Database {
	return tbl.database
}

// Logger gets the trace logger.
func (tbl XUserTable) Logger() sqlapi.Logger {
	return tbl.database.Logger()
}

// Dialect gets the database dialect.
func (tbl XUserTable) Dialect() dialect.Dialect {
	return tbl.database.Dialect()
}

// Name gets the table name.
func (tbl XUserTable) Name() sqlapi.TableName {
	return tbl.name
}

// PkColumn gets the column name used as a primary key.
func (tbl XUserTable) PkColumn() string {
	return tbl.pk
}

// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl XUserTable) DB() sqlapi.SqlDB {
	return tbl.db.(sqlapi.SqlDB)
}

// Execer gets the wrapped database or transaction handle.
func (tbl XUserTable) Execer() sqlapi.Execer {
	return tbl.db
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl XUserTable) Tx() sqlapi.SqlTx {
	return tbl.db.(sqlapi.SqlTx)
}

// IsTx tests whether this is within a transaction.
func (tbl XUserTable) IsTx() bool {
	return tbl.db.IsTx()
}

// Using returns a modified XUserTabler using the the Execer supplied,
// which will typically be a transaction (i.e. SqlTx). This is needed when making multiple
// queries across several tables within a single transaction.
//
// The result is a modified copy of the table; the original is unchanged.
func (tbl XUserTable) Using(tx sqlapi.Execer) XUserQueryer {
	tbl.db = tx
	return tbl
}

// Transact runs the function provided within a transaction. If the function completes without error,
// the transaction is committed. If there is an error or a panic, the transaction is rolled back.
//
// The options can be nil, in which case the default behaviour is that of the underlying connection.
//
// Nested transactions (i.e. within 'fn') are permitted: they execute within the outermost transaction.
// Therefore they do not commit until the outermost transaction commits.
func (tbl XUserTable) Transact(ctx context.Context, txOptions *sql.TxOptions, fn func(XUserQueryer) error) error {
	var err error
	if tbl.IsTx() {
		err = fn(tbl) // nested transactions are inlined
	} else {
		err = tbl.DB().Transact(ctx, txOptions, func(tx sqlapi.SqlTx) error {
			return fn(tbl.Using(tx))
		})
	}
	return tbl.Logger().LogIfError(err)
}

func (tbl XUserTable) quotedName() string {
	return tbl.Dialect().Quoter().Quote(tbl.name.String())
}

func (tbl XUserTable) quotedNameW(w dialect.StringWriter) {
	tbl.Dialect().Quoter().QuoteW(w, tbl.name.String())
}

//-------------------------------------------------------------------------------------------------

// NumXUserTableColumns is the total number of columns in XUserTable.
const NumXUserTableColumns = 22

// NumXUserTableDataColumns is the number of columns in XUserTable not including the auto-increment key.
const NumXUserTableDataColumns = 21

// XUserTableColumnNames is the list of columns in XUserTable.
const XUserTableColumnNames = "uid,name,emailaddress,addressid,avatar,role,active,admin,fave,lastupdated,i8,u8,i16,u16,i32,u32,i64,u64,f32,f64,token,secret"

// XUserTableDataColumnNames is the list of data columns in XUserTable.
const XUserTableDataColumnNames = "name,emailaddress,addressid,avatar,role,active,admin,fave,lastupdated,i8,u8,i16,u16,i32,u32,i64,u64,f32,f64,token,secret"

var listOfXUserTableColumnNames = strings.Split(XUserTableColumnNames, ",")

// scanXUsers reads rows from the database and returns a slice of corresponding values.
// It also returns a number indicating how many rows were read; this will be larger than the length of the
// slice if reading stopped after the first row.
func scanXUsers(query string, rows sqlapi.SqlRows, firstOnly bool) (vv []*User, n int64, err error) {
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
