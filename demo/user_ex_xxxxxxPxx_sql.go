// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.
// sqlapi v0.53.0; sqlgen v0.72.0

package demo

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/rickb777/sqlapi"
	"github.com/rickb777/sqlapi/dialect"
	"github.com/rickb777/sqlapi/require"
	"github.com/rickb777/sqlapi/support"
	"github.com/rickb777/where"
	"strings"
)

// PUserTabler lists table methods provided by PUserTable.
type PUserTabler interface {
	sqlapi.Table

	// WithPrefix returns a modified PUserTabler with a given table name prefix.
	WithPrefix(pfx string) PUserTabler

	// WithContext returns a modified PUserTabler with a given context.
	WithContext(ctx context.Context) PUserTabler
}

//-------------------------------------------------------------------------------------------------

// PUserQueryer lists query methods provided by PUserTable.
type PUserQueryer interface {
	sqlapi.Table

	// Using returns a modified PUserQueryer using the Execer supplied,
	// which will typically be a transaction (i.e. SqlTx).
	Using(tx sqlapi.Execer) PUserQueryer

	// Transact runs the function provided within a transaction. The transction is committed
	// unless an error occurs.
	Transact(txOptions *sql.TxOptions, fn func(PUserQueryer) error) error

	// Exec executes a query without returning any rows.
	Exec(req require.Requirement, query string, args ...interface{}) (int64, error)

	// Insert adds new records for the Users, setting the primary key field for each one.
	Insert(req require.Requirement, vv ...*User) error

	// UpdateByUid updates one or more columns, given a uid value.
	UpdateByUid(req require.Requirement, uid int64, fields ...sql.NamedArg) (int64, error)

	// UpdateByName updates one or more columns, given a name value.
	UpdateByName(req require.Requirement, name string, fields ...sql.NamedArg) (int64, error)

	// UpdateByEmailaddress updates one or more columns, given a emailaddress value.
	UpdateByEmailaddress(req require.Requirement, emailaddress string, fields ...sql.NamedArg) (int64, error)

	// UpdateByAddressid updates one or more columns, given a addressid value.
	UpdateByAddressid(req require.Requirement, addressid int64, fields ...sql.NamedArg) (int64, error)

	// UpdateByAvatar updates one or more columns, given a avatar value.
	UpdateByAvatar(req require.Requirement, avatar string, fields ...sql.NamedArg) (int64, error)

	// UpdateByRole updates one or more columns, given a role value.
	UpdateByRole(req require.Requirement, role Role, fields ...sql.NamedArg) (int64, error)

	// UpdateByLastupdated updates one or more columns, given a lastupdated value.
	UpdateByLastupdated(req require.Requirement, lastupdated int64, fields ...sql.NamedArg) (int64, error)

	// UpdateByI8 updates one or more columns, given a i8 value.
	UpdateByI8(req require.Requirement, i8 int8, fields ...sql.NamedArg) (int64, error)

	// UpdateByU8 updates one or more columns, given a u8 value.
	UpdateByU8(req require.Requirement, u8 uint8, fields ...sql.NamedArg) (int64, error)

	// UpdateByI16 updates one or more columns, given a i16 value.
	UpdateByI16(req require.Requirement, i16 int16, fields ...sql.NamedArg) (int64, error)

	// UpdateByU16 updates one or more columns, given a u16 value.
	UpdateByU16(req require.Requirement, u16 uint16, fields ...sql.NamedArg) (int64, error)

	// UpdateByI32 updates one or more columns, given a i32 value.
	UpdateByI32(req require.Requirement, i32 int32, fields ...sql.NamedArg) (int64, error)

	// UpdateByU32 updates one or more columns, given a u32 value.
	UpdateByU32(req require.Requirement, u32 uint32, fields ...sql.NamedArg) (int64, error)

	// UpdateByI64 updates one or more columns, given a i64 value.
	UpdateByI64(req require.Requirement, i64 int64, fields ...sql.NamedArg) (int64, error)

	// UpdateByU64 updates one or more columns, given a u64 value.
	UpdateByU64(req require.Requirement, u64 uint64, fields ...sql.NamedArg) (int64, error)

	// UpdateByF32 updates one or more columns, given a f32 value.
	UpdateByF32(req require.Requirement, f32 float32, fields ...sql.NamedArg) (int64, error)

	// UpdateByF64 updates one or more columns, given a f64 value.
	UpdateByF64(req require.Requirement, f64 float64, fields ...sql.NamedArg) (int64, error)

	// UpdateFields updates one or more columns, given a 'where' clause.
	UpdateFields(req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error)

	// Update updates records, matching them by primary key.
	Update(req require.Requirement, vv ...*User) (int64, error)

	// Upsert inserts or updates a record, matching it using the expression supplied.
	// This expression is used to search for an existing record based on some specified
	// key column(s). It must match either zero or one existing record. If it matches
	// none, a new record is inserted; otherwise the matching record is updated. An
	// error results if these conditions are not met.
	Upsert(v *User, wh where.Expression) error
}

//-------------------------------------------------------------------------------------------------

// PUserTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type PUserTable struct {
	sqlapi.CoreTable
	ctx context.Context
	pk  string
}

// Type conformance checks
var _ sqlapi.Table = &PUserTable{}

// NewPUserTable returns a new table instance.
// If a blank table name is supplied, the default name "users" will be used instead.
// The request context is initialised with the background.
func NewPUserTable(name string, d sqlapi.SqlDB) PUserTable {
	if name == "" {
		name = "users"
	}
	return PUserTable{
		CoreTable: sqlapi.CoreTable{
			Nm: sqlapi.TableName{Prefix: "", Name: name},
			Ex: d,
		},
		ctx: context.Background(),
		pk:  "uid",
	}
}

// CopyTableAsPUserTable copies a table instance, retaining the name etc but
// providing methods appropriate for 'User'.
//
// It serves to provide methods appropriate for 'User'. This is most useful when this is used to represent a
// join result. In such cases, there won't be any need for DDL methods, nor Exec, Insert, Update or Delete.
func CopyTableAsPUserTable(origin sqlapi.Table) PUserTable {
	return PUserTable{
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
//func (tbl PUserTable) SetPkColumn(pk string) PUserTabler {
//	tbl.pk = pk
//	return tbl
//}

// WithPrefix sets the table name prefix for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl PUserTable) WithPrefix(pfx string) PUserTabler {
	tbl.Nm.Prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries via this table.
// The result is a modified copy of the table; the original is unchanged.
func (tbl PUserTable) WithContext(ctx context.Context) PUserTabler {
	tbl.ctx = ctx
	return tbl
}

// Ctx gets the current request context.
func (tbl PUserTable) Ctx() context.Context {
	return tbl.ctx
}

// PkColumn gets the column name used as a primary key.
func (tbl PUserTable) PkColumn() string {
	return tbl.pk
}

// Using returns a modified PUserTabler using the the Execer supplied,
// which will typically be a transaction (i.e. SqlTx). This is needed when making multiple
// queries across several tables within a single transaction.
//
// The result is a modified copy of the table; the original is unchanged.
func (tbl PUserTable) Using(tx sqlapi.Execer) PUserQueryer {
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
func (tbl PUserTable) Transact(txOptions *sql.TxOptions, fn func(PUserQueryer) error) error {
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

func (tbl PUserTable) quotedName() string {
	return tbl.Dialect().Quoter().Quote(tbl.Nm.String())
}

func (tbl PUserTable) quotedNameW(w dialect.StringWriter) {
	tbl.Dialect().Quoter().QuoteW(w, tbl.Nm.String())
}

//-------------------------------------------------------------------------------------------------

// NumPUserTableColumns is the total number of columns in PUserTable.
const NumPUserTableColumns = 22

// NumPUserTableDataColumns is the number of columns in PUserTable not including the auto-increment key.
const NumPUserTableDataColumns = 21

// PUserTableColumnNames is the list of columns in PUserTable.
const PUserTableColumnNames = "uid,name,emailaddress,addressid,avatar,role,active,admin,fave,lastupdated,i8,u8,i16,u16,i32,u32,i64,u64,f32,f64,token,secret"

// PUserTableDataColumnNames is the list of data columns in PUserTable.
const PUserTableDataColumnNames = "name,emailaddress,addressid,avatar,role,active,admin,fave,lastupdated,i8,u8,i16,u16,i32,u32,i64,u64,f32,f64,token,secret"

var listOfPUserTableColumnNames = strings.Split(PUserTableColumnNames, ",")

//-------------------------------------------------------------------------------------------------

// Exec executes a query without returning any rows.
// It returns the number of rows affected (if the database driver supports this).
//
// The args are for any placeholder parameters in the query.
func (tbl PUserTable) Exec(req require.Requirement, query string, args ...interface{}) (int64, error) {
	return support.Exec(tbl, req, query, args...)
}

// scanPUsers reads rows from the database and returns a slice of corresponding values.
// It also returns a number indicating how many rows were read; this will be larger than the length of the
// slice if reading stopped after the first row.
func scanPUsers(query string, rows sqlapi.SqlRows, firstOnly bool) (vv []*User, n int64, err error) {
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

func constructPUserTableInsert(tbl PUserTable, w dialect.StringWriter, v *User, withPk bool) (s []interface{}, err error) {
	q := tbl.Dialect().Quoter()
	s = make([]interface{}, 0, 22)

	comma := ""
	w.WriteString(" (")

	if withPk {
		q.QuoteW(w, "uid")
		comma = ","
		s = append(s, v.Uid)
	}

	w.WriteString(comma)
	q.QuoteW(w, "name")
	s = append(s, v.Name)
	comma = ","

	w.WriteString(comma)
	q.QuoteW(w, "emailaddress")
	s = append(s, v.EmailAddress)

	if v.AddressId != nil {
		w.WriteString(comma)
		q.QuoteW(w, "addressid")
		s = append(s, v.AddressId)
	}

	if v.Avatar != nil {
		w.WriteString(comma)
		q.QuoteW(w, "avatar")
		s = append(s, v.Avatar)
	}

	if v.Role != nil {
		w.WriteString(comma)
		q.QuoteW(w, "role")
		s = append(s, v.Role)
	}

	w.WriteString(comma)
	q.QuoteW(w, "active")
	s = append(s, v.Active)

	w.WriteString(comma)
	q.QuoteW(w, "admin")
	s = append(s, v.Admin)

	w.WriteString(comma)
	q.QuoteW(w, "fave")
	x, err := json.Marshal(&v.Fave)
	if err != nil {
		return nil, tbl.Logger().LogError(errors.WithStack(err))
	}
	s = append(s, x)

	w.WriteString(comma)
	q.QuoteW(w, "lastupdated")
	s = append(s, v.LastUpdated)

	w.WriteString(comma)
	q.QuoteW(w, "i8")
	s = append(s, v.Numbers.I8)

	w.WriteString(comma)
	q.QuoteW(w, "u8")
	s = append(s, v.Numbers.U8)

	w.WriteString(comma)
	q.QuoteW(w, "i16")
	s = append(s, v.Numbers.I16)

	w.WriteString(comma)
	q.QuoteW(w, "u16")
	s = append(s, v.Numbers.U16)

	w.WriteString(comma)
	q.QuoteW(w, "i32")
	s = append(s, v.Numbers.I32)

	w.WriteString(comma)
	q.QuoteW(w, "u32")
	s = append(s, v.Numbers.U32)

	w.WriteString(comma)
	q.QuoteW(w, "i64")
	s = append(s, v.Numbers.I64)

	w.WriteString(comma)
	q.QuoteW(w, "u64")
	s = append(s, v.Numbers.U64)

	w.WriteString(comma)
	q.QuoteW(w, "f32")
	s = append(s, v.Numbers.F32)

	w.WriteString(comma)
	q.QuoteW(w, "f64")
	s = append(s, v.Numbers.F64)

	w.WriteString(comma)
	q.QuoteW(w, "token")
	s = append(s, v.token)

	w.WriteString(comma)
	q.QuoteW(w, "secret")
	s = append(s, v.secret)

	w.WriteString(")")
	return s, nil
}

func constructPUserTableUpdate(tbl PUserTable, w dialect.StringWriter, v *User) (s []interface{}, err error) {
	q := tbl.Dialect().Quoter()
	j := 1
	s = make([]interface{}, 0, 21)

	comma := ""

	w.WriteString(comma)
	q.QuoteW(w, "name")
	w.WriteString("=?")
	s = append(s, v.Name)
	j++
	comma = ", "

	w.WriteString(comma)
	q.QuoteW(w, "emailaddress")
	w.WriteString("=?")
	s = append(s, v.EmailAddress)
	j++

	w.WriteString(comma)
	if v.AddressId != nil {
		q.QuoteW(w, "addressid")
		w.WriteString("=?")
		s = append(s, v.AddressId)
		j++
	} else {
		q.QuoteW(w, "addressid")
		w.WriteString("=NULL")
	}

	w.WriteString(comma)
	if v.Avatar != nil {
		q.QuoteW(w, "avatar")
		w.WriteString("=?")
		s = append(s, v.Avatar)
		j++
	} else {
		q.QuoteW(w, "avatar")
		w.WriteString("=NULL")
	}

	w.WriteString(comma)
	if v.Role != nil {
		q.QuoteW(w, "role")
		w.WriteString("=?")
		s = append(s, v.Role)
		j++
	} else {
		q.QuoteW(w, "role")
		w.WriteString("=NULL")
	}

	w.WriteString(comma)
	q.QuoteW(w, "active")
	w.WriteString("=?")
	s = append(s, v.Active)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "admin")
	w.WriteString("=?")
	s = append(s, v.Admin)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "fave")
	w.WriteString("=?")
	j++

	x, err := json.Marshal(&v.Fave)
	if err != nil {
		return nil, tbl.Logger().LogError(errors.WithStack(err))
	}
	s = append(s, x)

	w.WriteString(comma)
	q.QuoteW(w, "lastupdated")
	w.WriteString("=?")
	s = append(s, v.LastUpdated)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "i8")
	w.WriteString("=?")
	s = append(s, v.Numbers.I8)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "u8")
	w.WriteString("=?")
	s = append(s, v.Numbers.U8)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "i16")
	w.WriteString("=?")
	s = append(s, v.Numbers.I16)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "u16")
	w.WriteString("=?")
	s = append(s, v.Numbers.U16)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "i32")
	w.WriteString("=?")
	s = append(s, v.Numbers.I32)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "u32")
	w.WriteString("=?")
	s = append(s, v.Numbers.U32)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "i64")
	w.WriteString("=?")
	s = append(s, v.Numbers.I64)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "u64")
	w.WriteString("=?")
	s = append(s, v.Numbers.U64)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "f32")
	w.WriteString("=?")
	s = append(s, v.Numbers.F32)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "f64")
	w.WriteString("=?")
	s = append(s, v.Numbers.F64)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "token")
	w.WriteString("=?")
	s = append(s, v.token)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "secret")
	w.WriteString("=?")
	s = append(s, v.secret)
	j++
	return s, nil
}

//--------------------------------------------------------------------------------

// Insert adds new records for the Users.// The Users have their primary key fields set to the new record identifiers.
// The User.PreInsert() method will be called, if it exists.
func (tbl PUserTable) Insert(req require.Requirement, vv ...*User) error {
	if req == require.All {
		req = require.Exactly(len(vv))
	}

	var count int64
	returning := ""
	insertHasReturningPhrase := tbl.Dialect().InsertHasReturningPhrase()
	if insertHasReturningPhrase {
		returning = fmt.Sprintf(" RETURNING %q", tbl.pk)
	}

	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlapi.CanPreInsert); ok {
			err := hook.PreInsert()
			if err != nil {
				return tbl.Logger().LogError(err)
			}
		}

		b := dialect.Adapt(&bytes.Buffer{})
		b.WriteString("INSERT INTO ")
		tbl.quotedNameW(b)

		fields, err := constructPUserTableInsert(tbl, b, v, false)
		if err != nil {
			return tbl.Logger().LogError(err)
		}

		b.WriteString(" VALUES (")
		b.WriteString(tbl.Dialect().Placeholders(len(fields)))
		b.WriteString(")")
		b.WriteString(returning)

		query := b.String()
		tbl.Logger().LogQuery(query, fields...)

		var n int64 = 1
		if insertHasReturningPhrase {
			row := tbl.Execer().QueryRowContext(tbl.ctx, query, fields...)
			var i64 int64
			err = row.Scan(&i64)
			v.Uid = i64

		} else {
			i64, e2 := tbl.Execer().InsertContext(tbl.ctx, tbl.pk, query, fields...)
			if e2 != nil {
				return tbl.Logger().LogError(e2)
			}
			v.Uid = i64
		}

		if err != nil {
			return tbl.Logger().LogError(err)
		}
		count += n
	}

	return tbl.Logger().LogIfError(require.ErrorIfExecNotSatisfiedBy(req, count))
}

// UpdateByUid updates one or more columns, given a uid value.
func (tbl PUserTable) UpdateByUid(req require.Requirement, uid int64, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(req, where.Eq("uid", uid), fields...)
}

// UpdateByName updates one or more columns, given a name value.
func (tbl PUserTable) UpdateByName(req require.Requirement, name string, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(req, where.Eq("name", name), fields...)
}

// UpdateByEmailaddress updates one or more columns, given a emailaddress value.
func (tbl PUserTable) UpdateByEmailaddress(req require.Requirement, emailaddress string, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(req, where.Eq("emailaddress", emailaddress), fields...)
}

// UpdateByAddressid updates one or more columns, given a addressid value.
func (tbl PUserTable) UpdateByAddressid(req require.Requirement, addressid int64, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(req, where.Eq("addressid", addressid), fields...)
}

// UpdateByAvatar updates one or more columns, given a avatar value.
func (tbl PUserTable) UpdateByAvatar(req require.Requirement, avatar string, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(req, where.Eq("avatar", avatar), fields...)
}

// UpdateByRole updates one or more columns, given a role value.
func (tbl PUserTable) UpdateByRole(req require.Requirement, role Role, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(req, where.Eq("role", role), fields...)
}

// UpdateByLastupdated updates one or more columns, given a lastupdated value.
func (tbl PUserTable) UpdateByLastupdated(req require.Requirement, lastupdated int64, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(req, where.Eq("lastupdated", lastupdated), fields...)
}

// UpdateByI8 updates one or more columns, given a i8 value.
func (tbl PUserTable) UpdateByI8(req require.Requirement, i8 int8, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(req, where.Eq("i8", i8), fields...)
}

// UpdateByU8 updates one or more columns, given a u8 value.
func (tbl PUserTable) UpdateByU8(req require.Requirement, u8 uint8, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(req, where.Eq("u8", u8), fields...)
}

// UpdateByI16 updates one or more columns, given a i16 value.
func (tbl PUserTable) UpdateByI16(req require.Requirement, i16 int16, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(req, where.Eq("i16", i16), fields...)
}

// UpdateByU16 updates one or more columns, given a u16 value.
func (tbl PUserTable) UpdateByU16(req require.Requirement, u16 uint16, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(req, where.Eq("u16", u16), fields...)
}

// UpdateByI32 updates one or more columns, given a i32 value.
func (tbl PUserTable) UpdateByI32(req require.Requirement, i32 int32, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(req, where.Eq("i32", i32), fields...)
}

// UpdateByU32 updates one or more columns, given a u32 value.
func (tbl PUserTable) UpdateByU32(req require.Requirement, u32 uint32, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(req, where.Eq("u32", u32), fields...)
}

// UpdateByI64 updates one or more columns, given a i64 value.
func (tbl PUserTable) UpdateByI64(req require.Requirement, i64 int64, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(req, where.Eq("i64", i64), fields...)
}

// UpdateByU64 updates one or more columns, given a u64 value.
func (tbl PUserTable) UpdateByU64(req require.Requirement, u64 uint64, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(req, where.Eq("u64", u64), fields...)
}

// UpdateByF32 updates one or more columns, given a f32 value.
func (tbl PUserTable) UpdateByF32(req require.Requirement, f32 float32, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(req, where.Eq("f32", f32), fields...)
}

// UpdateByF64 updates one or more columns, given a f64 value.
func (tbl PUserTable) UpdateByF64(req require.Requirement, f64 float64, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(req, where.Eq("f64", f64), fields...)
}

// UpdateFields updates one or more columns, given a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed (but note that this is risky!).
func (tbl PUserTable) UpdateFields(req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error) {
	return support.UpdateFields(tbl, req, wh, fields...)
}

//--------------------------------------------------------------------------------

// Update updates records, matching them by primary key. It returns the number of rows affected.
// The User.PreUpdate(Execer) method will be called, if it exists.
func (tbl PUserTable) Update(req require.Requirement, vv ...*User) (int64, error) {
	if req == require.All {
		req = require.Exactly(len(vv))
	}

	var count int64
	d := tbl.Dialect()
	q := d.Quoter()

	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlapi.CanPreUpdate); ok {
			err := hook.PreUpdate()
			if err != nil {
				return count, tbl.Logger().LogError(err)
			}
		}

		b := dialect.Adapt(&bytes.Buffer{})
		b.WriteString("UPDATE ")
		tbl.quotedNameW(b)
		b.WriteString(" SET ")

		args, err := constructPUserTableUpdate(tbl, b, v)
		if err != nil {
			return count, err
		}
		args = append(args, v.Uid)

		b.WriteString(" WHERE ")
		q.QuoteW(b, tbl.pk)
		b.WriteString("=?")

		query := b.String()
		n, err := tbl.Exec(nil, query, args...)
		if err != nil {
			return count, err
		}
		count += n
	}

	return count, tbl.Logger().LogIfError(require.ErrorIfExecNotSatisfiedBy(req, count))
}

//--------------------------------------------------------------------------------

// Upsert inserts or updates a record, matching it using the expression supplied.
// This expression is used to search for an existing record based on some specified
// key column(s). It must match either zero or one existing record. If it matches
// none, a new record is inserted; otherwise the matching record is updated. An
// error results if these conditions are not met.
func (tbl PUserTable) Upsert(v *User, wh where.Expression) error {
	col := tbl.Dialect().Quoter().Quote(tbl.pk)
	qName := tbl.quotedName()
	whs, args := where.Where(wh, tbl.Dialect().Quoter())

	query := fmt.Sprintf("SELECT %s FROM %s %s", col, qName, whs)
	rows, err := support.Query(tbl, query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	if !rows.Next() {
		return tbl.Insert(require.One, v)
	}

	var id int64
	err = rows.Scan(&id)
	if err != nil {
		return tbl.Logger().LogIfError(err)
	}

	if rows.Next() {
		return require.ErrWrongSize(2, "expected to find no more than 1 but got at least 2 using %q", wh)
	}

	v.Uid = id
	_, err = tbl.Update(require.One, v)
	return err
}
