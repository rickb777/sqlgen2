// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.
// sqlapi v0.36.1; sqlgen v0.57.0

package demo

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/rickb777/sqlapi"
	"github.com/rickb777/sqlapi/constraint"
	"github.com/rickb777/sqlapi/dialect"
	"github.com/rickb777/sqlapi/require"
	"strings"
)

// IUserTabler lists methods provided by IUserTable.
type IUserTabler interface {
	sqlapi.Table

	// Constraints returns the table's constraints.
	Constraints() constraint.Constraints

	// WithConstraint returns a modified IUserTabler with added data consistency constraints.
	WithConstraint(cc ...constraint.Constraint) IUserTabler

	// WithPrefix returns a modified IUserTabler with a given table name prefix.
	WithPrefix(pfx string) IUserTabler

	// WithContext returns a modified IUserTabler with a given context.
	WithContext(ctx context.Context) IUserTabler

	// Using returns a modified IUserTabler using the transaction supplied.
	Using(tx sqlapi.SqlTx) IUserTabler

	// Transact runs the function provided within a transaction.
	Transact(txOptions *sql.TxOptions, fn func(IUserTabler) error) error

	// Insert adds new records for the Users, setting the primary key field for each one.
	Insert(req require.Requirement, vv ...*User) error
}

// IUserTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type IUserTable struct {
	name        sqlapi.TableName
	database    sqlapi.Database
	db          sqlapi.Execer
	constraints constraint.Constraints
	ctx         context.Context
	pk          string
}

// Type conformance checks
var _ sqlapi.Table = &IUserTable{}

// NewIUserTable returns a new table instance.
// If a blank table name is supplied, the default name "users" will be used instead.
// The request context is initialised with the background.
func NewIUserTable(name string, d sqlapi.Database) IUserTable {
	if name == "" {
		name = "users"
	}
	var constraints constraint.Constraints
	constraints = append(constraints,
		constraint.FkConstraint{"addressid", constraint.Reference{"addresses", "id"}, "restrict", "restrict"})

	return IUserTable{
		name:        sqlapi.TableName{Prefix: "", Name: name},
		database:    d,
		db:          d.DB(),
		constraints: constraints,
		ctx:         context.Background(),
		pk:          "uid",
	}
}

// CopyTableAsIUserTable copies a table instance, retaining the name etc but
// providing methods appropriate for 'User'. It doesn't copy the constraints of the original table.
//
// It serves to provide methods appropriate for 'User'. This is most useful when this is used to represent a
// join result. In such cases, there won't be any need for DDL methods, nor Exec, Insert, Update or Delete.
func CopyTableAsIUserTable(origin sqlapi.Table) IUserTable {
	return IUserTable{
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
//func (tbl IUserTable) SetPkColumn(pk string) IUserTabler {
//	tbl.pk = pk
//	return tbl
//}

// WithPrefix sets the table name prefix for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl IUserTable) WithPrefix(pfx string) IUserTabler {
	tbl.name.Prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries via this table.
// The result is a modified copy of the table; the original is unchanged.
//
// The shared context in the *Database is not altered by this method. So it
// is possible to use different contexts for different (groups of) queries.
func (tbl IUserTable) WithContext(ctx context.Context) IUserTabler {
	tbl.ctx = ctx
	return tbl
}

// Database gets the shared database information.
func (tbl IUserTable) Database() sqlapi.Database {
	return tbl.database
}

// Logger gets the trace logger.
func (tbl IUserTable) Logger() sqlapi.Logger {
	return tbl.database.Logger()
}

// WithConstraint returns a modified IUserTabler with added data consistency constraints.
func (tbl IUserTable) WithConstraint(cc ...constraint.Constraint) IUserTabler {
	tbl.constraints = append(tbl.constraints, cc...)
	return tbl
}

// Constraints returns the table's constraints.
func (tbl IUserTable) Constraints() constraint.Constraints {
	return tbl.constraints
}

// Ctx gets the current request context.
func (tbl IUserTable) Ctx() context.Context {
	return tbl.ctx
}

// Dialect gets the database dialect.
func (tbl IUserTable) Dialect() dialect.Dialect {
	return tbl.database.Dialect()
}

// Name gets the table name.
func (tbl IUserTable) Name() sqlapi.TableName {
	return tbl.name
}

// PkColumn gets the column name used as a primary key.
func (tbl IUserTable) PkColumn() string {
	return tbl.pk
}

// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl IUserTable) DB() sqlapi.SqlDB {
	return tbl.db.(sqlapi.SqlDB)
}

// Execer gets the wrapped database or transaction handle.
func (tbl IUserTable) Execer() sqlapi.Execer {
	return tbl.db
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl IUserTable) Tx() sqlapi.SqlTx {
	return tbl.db.(sqlapi.SqlTx)
}

// IsTx tests whether this is within a transaction.
func (tbl IUserTable) IsTx() bool {
	return tbl.db.IsTx()
}

// Using returns a modified IUserTabler using the transaction supplied. This is needed
// when making multiple queries across several tables within a single transaction.
// The result is a modified copy of the table; the original is unchanged.
func (tbl IUserTable) Using(tx sqlapi.SqlTx) IUserTabler {
	tbl.db = tx
	return tbl
}

// Transact runs the function provided within a transaction. If the function completes without error,
// the transaction is committed. If there is an error or a panic, the transaction is rolled back.
//
// Nested transactions (i.e. within 'fn') are permitted: they execute within the outermost transaction.
// Therefore they do not commit until the outermost transaction commits.
func (tbl IUserTable) Transact(txOptions *sql.TxOptions, fn func(IUserTabler) error) error {
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

func (tbl IUserTable) quotedName() string {
	return tbl.Dialect().Quoter().Quote(tbl.name.String())
}

func (tbl IUserTable) quotedNameW(w dialect.StringWriter) {
	tbl.Dialect().Quoter().QuoteW(w, tbl.name.String())
}

//--------------------------------------------------------------------------------

// NumIUserTableColumns is the total number of columns in IUserTable.
const NumIUserTableColumns = 22

// NumIUserTableDataColumns is the number of columns in IUserTable not including the auto-increment key.
const NumIUserTableDataColumns = 21

// IUserTableColumnNames is the list of columns in IUserTable.
const IUserTableColumnNames = "uid,name,emailaddress,addressid,avatar,role,active,admin,fave,lastupdated,i8,u8,i16,u16,i32,u32,i64,u64,f32,f64,token,secret"

// IUserTableDataColumnNames is the list of data columns in IUserTable.
const IUserTableDataColumnNames = "name,emailaddress,addressid,avatar,role,active,admin,fave,lastupdated,i8,u8,i16,u16,i32,u32,i64,u64,f32,f64,token,secret"

var listOfIUserTableColumnNames = strings.Split(IUserTableColumnNames, ",")

// scanIUsers reads rows from the database and returns a slice of corresponding values.
// It also returns a number indicating how many rows were read; this will be larger than the length of the
// slice if reading stopped after the first row.
func scanIUsers(query string, rows sqlapi.SqlRows, firstOnly bool) (vv []*User, n int64, err error) {
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

func constructIUserTableInsert(tbl IUserTable, w dialect.StringWriter, v *User, withPk bool) (s []interface{}, err error) {
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

//--------------------------------------------------------------------------------

// Insert adds new records for the Users.// The Users have their primary key fields set to the new record identifiers.
// The User.PreInsert() method will be called, if it exists.
func (tbl IUserTable) Insert(req require.Requirement, vv ...*User) error {
	if req == require.All {
		req = require.Exactly(len(vv))
	}

	var count int64
	insertHasReturningPhrase := tbl.Dialect().InsertHasReturningPhrase()
	returning := ""
	if tbl.Dialect().InsertHasReturningPhrase() {
		returning = fmt.Sprintf(" returning %q", tbl.pk)
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

		fields, err := constructIUserTableInsert(tbl, b, v, false)
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
			row := tbl.db.QueryRowContext(tbl.ctx, query, fields...)
			err = row.Scan(&v.Uid)

		} else {
			i64, e2 := tbl.db.InsertContext(tbl.ctx, query, fields...)
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
