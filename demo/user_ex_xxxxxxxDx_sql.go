// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.
// sqlapi v0.59.0; sqlgen v0.77.0

package demo

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/rickb777/sqlapi"
	"github.com/rickb777/sqlapi/driver"
	"github.com/rickb777/sqlapi/require"
	"github.com/rickb777/sqlapi/support"
	"github.com/rickb777/where"
	"strings"
)

// DUserTabler lists table methods provided by DUserTable.
type DUserTabler interface {
	sqlapi.Table

	// WithPrefix returns a modified DUserTabler with a given table name prefix.
	WithPrefix(pfx string) DUserTabler

	// WithContext returns a modified DUserTabler with a given context.
	WithContext(ctx context.Context) DUserTabler
}

//-------------------------------------------------------------------------------------------------

// DUserQueryer lists query methods provided by DUserTable.
type DUserQueryer interface {
	sqlapi.Table

	// Using returns a modified DUserQueryer using the Execer supplied,
	// which will typically be a transaction (i.e. SqlTx).
	Using(tx sqlapi.Execer) DUserQueryer

	// Transact runs the function provided within a transaction. The transction is committed
	// unless an error occurs.
	Transact(txOptions *pgx.TxOptions, fn func(DUserQueryer) error) error

	// Exec executes a query without returning any rows.
	Exec(req require.Requirement, query string, args ...interface{}) (int64, error)

	// DeleteByUid deletes rows from the table, given some uid values.
	// The list of ids can be arbitrarily long.
	DeleteByUid(req require.Requirement, uid ...int64) (int64, error)

	// DeleteByName deletes rows from the table, given some name values.
	// The list of ids can be arbitrarily long.
	DeleteByName(req require.Requirement, name ...string) (int64, error)

	// DeleteByEmailaddress deletes rows from the table, given some emailaddress values.
	// The list of ids can be arbitrarily long.
	DeleteByEmailaddress(req require.Requirement, emailaddress ...string) (int64, error)

	// DeleteByAddressid deletes rows from the table, given some addressid values.
	// The list of ids can be arbitrarily long.
	DeleteByAddressid(req require.Requirement, addressid ...int64) (int64, error)

	// DeleteByAvatar deletes rows from the table, given some avatar values.
	// The list of ids can be arbitrarily long.
	DeleteByAvatar(req require.Requirement, avatar ...string) (int64, error)

	// DeleteByRole deletes rows from the table, given some role values.
	// The list of ids can be arbitrarily long.
	DeleteByRole(req require.Requirement, role ...Role) (int64, error)

	// DeleteByLastupdated deletes rows from the table, given some lastupdated values.
	// The list of ids can be arbitrarily long.
	DeleteByLastupdated(req require.Requirement, lastupdated ...int64) (int64, error)

	// DeleteByI8 deletes rows from the table, given some i8 values.
	// The list of ids can be arbitrarily long.
	DeleteByI8(req require.Requirement, i8 ...int8) (int64, error)

	// DeleteByU8 deletes rows from the table, given some u8 values.
	// The list of ids can be arbitrarily long.
	DeleteByU8(req require.Requirement, u8 ...uint8) (int64, error)

	// DeleteByI16 deletes rows from the table, given some i16 values.
	// The list of ids can be arbitrarily long.
	DeleteByI16(req require.Requirement, i16 ...int16) (int64, error)

	// DeleteByU16 deletes rows from the table, given some u16 values.
	// The list of ids can be arbitrarily long.
	DeleteByU16(req require.Requirement, u16 ...uint16) (int64, error)

	// DeleteByI32 deletes rows from the table, given some i32 values.
	// The list of ids can be arbitrarily long.
	DeleteByI32(req require.Requirement, i32 ...int32) (int64, error)

	// DeleteByU32 deletes rows from the table, given some u32 values.
	// The list of ids can be arbitrarily long.
	DeleteByU32(req require.Requirement, u32 ...uint32) (int64, error)

	// DeleteByI64 deletes rows from the table, given some i64 values.
	// The list of ids can be arbitrarily long.
	DeleteByI64(req require.Requirement, i64 ...int64) (int64, error)

	// DeleteByU64 deletes rows from the table, given some u64 values.
	// The list of ids can be arbitrarily long.
	DeleteByU64(req require.Requirement, u64 ...uint64) (int64, error)

	// DeleteByF32 deletes rows from the table, given some f32 values.
	// The list of ids can be arbitrarily long.
	DeleteByF32(req require.Requirement, f32 ...float32) (int64, error)

	// DeleteByF64 deletes rows from the table, given some f64 values.
	// The list of ids can be arbitrarily long.
	DeleteByF64(req require.Requirement, f64 ...float64) (int64, error)

	// Delete deletes one or more rows from the table, given a 'where' clause.
	// Use a nil value for the 'wh' argument if it is not needed (very risky!).
	Delete(req require.Requirement, wh where.Expression) (int64, error)
}

//-------------------------------------------------------------------------------------------------

// DUserTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type DUserTable struct {
	sqlapi.CoreTable
	ctx context.Context
	pk  string
}

// Type conformance checks
var _ sqlapi.Table = &DUserTable{}

// NewDUserTable returns a new table instance.
// If a blank table name is supplied, the default name "users" will be used instead.
// The request context is initialised with the background.
func NewDUserTable(name string, d sqlapi.SqlDB) DUserTable {
	if name == "" {
		name = "users"
	}
	return DUserTable{
		CoreTable: sqlapi.CoreTable{
			Nm: sqlapi.TableName{Prefix: "", Name: name},
			Ex: d,
		},
		ctx: context.Background(),
		pk:  "uid",
	}
}

// CopyTableAsDUserTable copies a table instance, retaining the name etc but
// providing methods appropriate for 'User'.
//
// It serves to provide methods appropriate for 'User'. This is most useful when this is used to represent a
// join result. In such cases, there won't be any need for DDL methods, nor Exec, Insert, Update or Delete.
func CopyTableAsDUserTable(origin sqlapi.Table) DUserTable {
	return DUserTable{
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
//func (tbl DUserTable) SetPkColumn(pk string) DUserTabler {
//	tbl.pk = pk
//	return tbl
//}

// WithPrefix sets the table name prefix for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl DUserTable) WithPrefix(pfx string) DUserTabler {
	tbl.Nm.Prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries via this table.
// The result is a modified copy of the table; the original is unchanged.
func (tbl DUserTable) WithContext(ctx context.Context) DUserTabler {
	tbl.ctx = ctx
	return tbl
}

// Ctx gets the current request context.
func (tbl DUserTable) Ctx() context.Context {
	return tbl.ctx
}

// PkColumn gets the column name used as a primary key.
func (tbl DUserTable) PkColumn() string {
	return tbl.pk
}

// Using returns a modified DUserTabler using the the Execer supplied,
// which will typically be a transaction (i.e. SqlTx). This is needed when making multiple
// queries across several tables within a single transaction.
//
// The result is a modified copy of the table; the original is unchanged.
func (tbl DUserTable) Using(tx sqlapi.Execer) DUserQueryer {
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
func (tbl DUserTable) Transact(txOptions *pgx.TxOptions, fn func(DUserQueryer) error) error {
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

func (tbl DUserTable) quotedName() string {
	return tbl.Dialect().Quoter().Quote(tbl.Nm.String())
}

func (tbl DUserTable) quotedNameW(w driver.StringWriter) {
	tbl.Dialect().Quoter().QuoteW(w, tbl.Nm.String())
}

//-------------------------------------------------------------------------------------------------

// NumDUserTableColumns is the total number of columns in DUserTable.
const NumDUserTableColumns = 22

// NumDUserTableDataColumns is the number of columns in DUserTable not including the auto-increment key.
const NumDUserTableDataColumns = 21

// DUserTableColumnNames is the list of columns in DUserTable.
const DUserTableColumnNames = "uid,name,emailaddress,addressid,avatar,role,active,admin,fave,lastupdated,i8,u8,i16,u16,i32,u32,i64,u64,f32,f64,token,secret"

// DUserTableDataColumnNames is the list of data columns in DUserTable.
const DUserTableDataColumnNames = "name,emailaddress,addressid,avatar,role,active,admin,fave,lastupdated,i8,u8,i16,u16,i32,u32,i64,u64,f32,f64,token,secret"

var listOfDUserTableColumnNames = strings.Split(DUserTableColumnNames, ",")

//-------------------------------------------------------------------------------------------------

// Exec executes a query without returning any rows.
// It returns the number of rows affected (if the database driver supports this).
//
// The args are for any placeholder parameters in the query.
func (tbl DUserTable) Exec(req require.Requirement, query string, args ...interface{}) (int64, error) {
	return support.Exec(tbl, req, query, args...)
}

// scanDUsers reads rows from the database and returns a slice of corresponding values.
// It also returns a number indicating how many rows were read; this will be larger than the length of the
// slice if reading stopped after the first row.
func scanDUsers(query string, rows sqlapi.SqlRows, firstOnly bool) (vv []*User, n int64, err error) {
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

//-------------------------------------------------------------------------------------------------

// DeleteByUid deletes rows from the table, given some uid values.
// The list of ids can be arbitrarily long.
func (tbl DUserTable) DeleteByUid(req require.Requirement, uid ...int64) (int64, error) {
	ii := support.Int64AsInterfaceSlice(uid)
	return support.DeleteByColumn(tbl, req, "uid", ii...)
}

// DeleteByName deletes rows from the table, given some name values.
// The list of ids can be arbitrarily long.
func (tbl DUserTable) DeleteByName(req require.Requirement, name ...string) (int64, error) {
	ii := support.StringAsInterfaceSlice(name)
	return support.DeleteByColumn(tbl, req, "name", ii...)
}

// DeleteByEmailaddress deletes rows from the table, given some emailaddress values.
// The list of ids can be arbitrarily long.
func (tbl DUserTable) DeleteByEmailaddress(req require.Requirement, emailaddress ...string) (int64, error) {
	ii := support.StringAsInterfaceSlice(emailaddress)
	return support.DeleteByColumn(tbl, req, "emailaddress", ii...)
}

// DeleteByAddressid deletes rows from the table, given some addressid values.
// The list of ids can be arbitrarily long.
func (tbl DUserTable) DeleteByAddressid(req require.Requirement, addressid ...int64) (int64, error) {
	ii := support.Int64AsInterfaceSlice(addressid)
	return support.DeleteByColumn(tbl, req, "addressid", ii...)
}

// DeleteByAvatar deletes rows from the table, given some avatar values.
// The list of ids can be arbitrarily long.
func (tbl DUserTable) DeleteByAvatar(req require.Requirement, avatar ...string) (int64, error) {
	ii := support.StringAsInterfaceSlice(avatar)
	return support.DeleteByColumn(tbl, req, "avatar", ii...)
}

// DeleteByRole deletes rows from the table, given some role values.
// The list of ids can be arbitrarily long.
func (tbl DUserTable) DeleteByRole(req require.Requirement, role ...Role) (int64, error) {
	ii := make([]interface{}, len(role))
	for i, v := range role {
		ii[i] = v
	}
	return support.DeleteByColumn(tbl, req, "role", ii...)
}

// DeleteByLastupdated deletes rows from the table, given some lastupdated values.
// The list of ids can be arbitrarily long.
func (tbl DUserTable) DeleteByLastupdated(req require.Requirement, lastupdated ...int64) (int64, error) {
	ii := support.Int64AsInterfaceSlice(lastupdated)
	return support.DeleteByColumn(tbl, req, "lastupdated", ii...)
}

// DeleteByI8 deletes rows from the table, given some i8 values.
// The list of ids can be arbitrarily long.
func (tbl DUserTable) DeleteByI8(req require.Requirement, i8 ...int8) (int64, error) {
	ii := support.Int8AsInterfaceSlice(i8)
	return support.DeleteByColumn(tbl, req, "i8", ii...)
}

// DeleteByU8 deletes rows from the table, given some u8 values.
// The list of ids can be arbitrarily long.
func (tbl DUserTable) DeleteByU8(req require.Requirement, u8 ...uint8) (int64, error) {
	ii := support.Uint8AsInterfaceSlice(u8)
	return support.DeleteByColumn(tbl, req, "u8", ii...)
}

// DeleteByI16 deletes rows from the table, given some i16 values.
// The list of ids can be arbitrarily long.
func (tbl DUserTable) DeleteByI16(req require.Requirement, i16 ...int16) (int64, error) {
	ii := support.Int16AsInterfaceSlice(i16)
	return support.DeleteByColumn(tbl, req, "i16", ii...)
}

// DeleteByU16 deletes rows from the table, given some u16 values.
// The list of ids can be arbitrarily long.
func (tbl DUserTable) DeleteByU16(req require.Requirement, u16 ...uint16) (int64, error) {
	ii := support.Uint16AsInterfaceSlice(u16)
	return support.DeleteByColumn(tbl, req, "u16", ii...)
}

// DeleteByI32 deletes rows from the table, given some i32 values.
// The list of ids can be arbitrarily long.
func (tbl DUserTable) DeleteByI32(req require.Requirement, i32 ...int32) (int64, error) {
	ii := support.Int32AsInterfaceSlice(i32)
	return support.DeleteByColumn(tbl, req, "i32", ii...)
}

// DeleteByU32 deletes rows from the table, given some u32 values.
// The list of ids can be arbitrarily long.
func (tbl DUserTable) DeleteByU32(req require.Requirement, u32 ...uint32) (int64, error) {
	ii := support.Uint32AsInterfaceSlice(u32)
	return support.DeleteByColumn(tbl, req, "u32", ii...)
}

// DeleteByI64 deletes rows from the table, given some i64 values.
// The list of ids can be arbitrarily long.
func (tbl DUserTable) DeleteByI64(req require.Requirement, i64 ...int64) (int64, error) {
	ii := support.Int64AsInterfaceSlice(i64)
	return support.DeleteByColumn(tbl, req, "i64", ii...)
}

// DeleteByU64 deletes rows from the table, given some u64 values.
// The list of ids can be arbitrarily long.
func (tbl DUserTable) DeleteByU64(req require.Requirement, u64 ...uint64) (int64, error) {
	ii := support.Uint64AsInterfaceSlice(u64)
	return support.DeleteByColumn(tbl, req, "u64", ii...)
}

// DeleteByF32 deletes rows from the table, given some f32 values.
// The list of ids can be arbitrarily long.
func (tbl DUserTable) DeleteByF32(req require.Requirement, f32 ...float32) (int64, error) {
	ii := make([]interface{}, len(f32))
	for i, v := range f32 {
		ii[i] = v
	}
	return support.DeleteByColumn(tbl, req, "f32", ii...)
}

// DeleteByF64 deletes rows from the table, given some f64 values.
// The list of ids can be arbitrarily long.
func (tbl DUserTable) DeleteByF64(req require.Requirement, f64 ...float64) (int64, error) {
	ii := make([]interface{}, len(f64))
	for i, v := range f64 {
		ii[i] = v
	}
	return support.DeleteByColumn(tbl, req, "f64", ii...)
}

// Delete deletes one or more rows from the table, given a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl DUserTable) Delete(req require.Requirement, wh where.Expression) (int64, error) {
	query, args := deleteRowsDUserTableSql(tbl, wh)
	return tbl.Exec(req, query, args...)
}

func deleteRowsDUserTableSql(tbl DUserTabler, wh where.Expression) (string, []interface{}) {
	whs, args := where.Where(wh, tbl.Dialect().Quoter())
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("DELETE FROM %s %s", quotedName, whs)
	return query, args
}

//-------------------------------------------------------------------------------------------------
