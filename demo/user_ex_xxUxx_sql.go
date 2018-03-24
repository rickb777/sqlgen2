// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

package demo

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"github.com/rickb777/sqlgen2"
	"github.com/rickb777/sqlgen2/constraint"
	"github.com/rickb777/sqlgen2/require"
	"github.com/rickb777/sqlgen2/schema"
	"github.com/rickb777/sqlgen2/support"
	"github.com/rickb777/sqlgen2/where"
	"io"
	"log"
)

// UUserTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type UUserTable struct {
	name        sqlgen2.TableName
	database    *sqlgen2.Database
	db          sqlgen2.Execer
	constraints constraint.Constraints
	ctx			context.Context
	pk          string
}

// Type conformance checks
var _ sqlgen2.Table = &UUserTable{}
var _ sqlgen2.Table = &UUserTable{}

// NewUUserTable returns a new table instance.
// If a blank table name is supplied, the default name "users" will be used instead.
// The request context is initialised with the background.
func NewUUserTable(name string, d *sqlgen2.Database) UUserTable {
	if name == "" {
		name = "users"
	}
	var constraints constraint.Constraints
	constraints = append(constraints, constraint.FkConstraint{"addressid", constraint.Reference{"addresses", "id"}, "restrict", "restrict"})
	
	return UUserTable{
		name:        sqlgen2.TableName{"", name},
		database:    d,
		db:          d.DB(),
		constraints: constraints,
		ctx:         context.Background(),
		pk:          "uid",
	}
}

// CopyTableAsUUserTable copies a table instance, retaining the name etc but
// providing methods appropriate for 'User'. It doesn't copy the constraints of the original table.
//
// It serves to provide methods appropriate for 'User'. This is most useful when this is used to represent a
// join result. In such cases, there won't be any need for DDL methods, nor Exec, Insert, Update or Delete.
func CopyTableAsUUserTable(origin sqlgen2.Table) UUserTable {
	return UUserTable{
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
func (tbl UUserTable) SetPkColumn(pk string) UUserTable {
	tbl.pk = pk
	return tbl
}


// WithPrefix sets the table name prefix for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl UUserTable) WithPrefix(pfx string) UUserTable {
	tbl.name.Prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries via this table.
// The result is a modified copy of the table; the original is unchanged.
//
// The shared context in the *Database is not altered by this method. So it
// is possible to use different contexts for different (groups of) queries.
func (tbl UUserTable) WithContext(ctx context.Context) UUserTable {
	tbl.ctx = ctx
	return tbl
}

// Database gets the shared database information.
func (tbl UUserTable) Database() *sqlgen2.Database {
	return tbl.database
}

// Logger gets the trace logger.
func (tbl UUserTable) Logger() *log.Logger {
	return tbl.database.Logger()
}

// WithConstraint returns a modified Table with added data consistency constraints.
func (tbl UUserTable) WithConstraint(cc ...constraint.Constraint) UUserTable {
	tbl.constraints = append(tbl.constraints, cc...)
	return tbl
}

// Constraints returns the table's constraints.
func (tbl UUserTable) Constraints() constraint.Constraints {
	return tbl.constraints
}

// Ctx gets the current request context.
func (tbl UUserTable) Ctx() context.Context {
	return tbl.ctx
}

// Dialect gets the database dialect.
func (tbl UUserTable) Dialect() schema.Dialect {
	return tbl.database.Dialect()
}

// Name gets the table name.
func (tbl UUserTable) Name() sqlgen2.TableName {
	return tbl.name
}


// PkColumn gets the column name used as a primary key.
func (tbl UUserTable) PkColumn() string {
	return tbl.pk
}


// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl UUserTable) DB() *sql.DB {
	return tbl.db.(*sql.DB)
}

// Execer gets the wrapped database or transaction handle.
func (tbl UUserTable) Execer() sqlgen2.Execer {
	return tbl.db
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl UUserTable) Tx() *sql.Tx {
	return tbl.db.(*sql.Tx)
}

// IsTx tests whether this is within a transaction.
func (tbl UUserTable) IsTx() bool {
	_, ok := tbl.db.(*sql.Tx)
	return ok
}

// BeginTx starts a transaction using the table's context.
// This context is used until the transaction is committed or rolled back.
//
// If this context is cancelled, the sql package will roll back the transaction.
// In this case, Tx.Commit will then return an error.
//
// The provided TxOptions is optional and may be nil if defaults should be used.
// If a non-default isolation level is used that the driver doesn't support,
// an error will be returned.
//
// Panics if the Execer is not TxStarter.
func (tbl UUserTable) BeginTx(opts *sql.TxOptions) (UUserTable, error) {
	var err error
	tbl.db, err = tbl.db.(sqlgen2.TxStarter).BeginTx(tbl.ctx, opts)
	return tbl, tbl.logIfError(err)
}

// Using returns a modified Table using the transaction supplied. This is needed
// when making multiple queries across several tables within a single transaction.
// The result is a modified copy of the table; the original is unchanged.
func (tbl UUserTable) Using(tx *sql.Tx) UUserTable {
	tbl.db = tx
	return tbl
}

func (tbl UUserTable) logQuery(query string, args ...interface{}) {
	tbl.database.LogQuery(query, args...)
}

func (tbl UUserTable) logError(err error) error {
	return tbl.database.LogError(err)
}

func (tbl UUserTable) logIfError(err error) error {
	return tbl.database.LogIfError(err)
}


//--------------------------------------------------------------------------------

const NumUUserColumns = 22

const NumUUserDataColumns = 21

const UUserColumnNames = "uid,name,emailaddress,addressid,avatar,role,active,admin,fave,lastupdated,i8,u8,i16,u16,i32,u32,i64,u64,f32,f64,token,secret"

const UUserDataColumnNames = "name,emailaddress,addressid,avatar,role,active,admin,fave,lastupdated,i8,u8,i16,u16,i32,u32,i64,u64,f32,f64,token,secret"

//--------------------------------------------------------------------------------

// Exec executes a query without returning any rows.
// It returns the number of rows affected (if the database driver supports this).
//
// The args are for any placeholder parameters in the query.
func (tbl UUserTable) Exec(req require.Requirement, query string, args ...interface{}) (int64, error) {
	return support.Exec(tbl, req, query, args...)
}

//--------------------------------------------------------------------------------

// Query is the low-level request method for this table. The query is logged using whatever logger is
// configured. If an error arises, this too is logged.
//
// If you need a context other than the background, use WithContext before calling Query.
//
// The args are for any placeholder parameters in the query.
//
// The caller must call rows.Close() on the result.
//
// Wrap the result in *sqlgen2.Rows if you need to access its data as a map.
func (tbl UUserTable) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return support.Query(tbl, query, args...)
}

//--------------------------------------------------------------------------------

// QueryOneNullString is a low-level access method for one string. This can be used for function queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl UUserTable) QueryOneNullString(req require.Requirement, query string, args ...interface{}) (result sql.NullString, err error) {
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
func (tbl UUserTable) QueryOneNullInt64(req require.Requirement, query string, args ...interface{}) (result sql.NullInt64, err error) {
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
func (tbl UUserTable) QueryOneNullFloat64(req require.Requirement, query string, args ...interface{}) (result sql.NullFloat64, err error) {
	err = support.QueryOneNullThing(tbl, req, &result, query, args...)
	return result, err
}

func constructUUserUpdate(w io.Writer, v *User, dialect schema.Dialect) (s []interface{}, err error) {
	j := 1
	s = make([]interface{}, 0, 21)

	comma := ""

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "name", j)
	s = append(s, v.Name)
	comma = ", "
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "emailaddress", j)
	s = append(s, v.EmailAddress)
		j++

	io.WriteString(w, comma)
	if v.AddressId != nil {
		dialect.QuoteWithPlaceholder(w, "addressid", j)
		s = append(s, v.AddressId)
		j++
	} else {
		dialect.QuoteW(w, "addressid")
		io.WriteString(w, "=NULL")
	}

	io.WriteString(w, comma)
	if v.Avatar != nil {
		dialect.QuoteWithPlaceholder(w, "avatar", j)
		s = append(s, v.Avatar)
		j++
	} else {
		dialect.QuoteW(w, "avatar")
		io.WriteString(w, "=NULL")
	}

	io.WriteString(w, comma)
	if v.Role != nil {
		dialect.QuoteWithPlaceholder(w, "role", j)
		s = append(s, v.Role)
		j++
	} else {
		dialect.QuoteW(w, "role")
		io.WriteString(w, "=NULL")
	}

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "active", j)
	s = append(s, v.Active)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "admin", j)
	s = append(s, v.Admin)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "fave", j)
		j++
	x, err := json.Marshal(&v.Fave)
	if err != nil {
		return nil, err
	}
	s = append(s, x)

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "lastupdated", j)
	s = append(s, v.LastUpdated)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "i8", j)
	s = append(s, v.Numbers.I8)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "u8", j)
	s = append(s, v.Numbers.U8)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "i16", j)
	s = append(s, v.Numbers.I16)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "u16", j)
	s = append(s, v.Numbers.U16)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "i32", j)
	s = append(s, v.Numbers.I32)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "u32", j)
	s = append(s, v.Numbers.U32)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "i64", j)
	s = append(s, v.Numbers.I64)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "u64", j)
	s = append(s, v.Numbers.U64)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "f32", j)
	s = append(s, v.Numbers.F32)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "f64", j)
	s = append(s, v.Numbers.F64)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "token", j)
	s = append(s, v.token)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "secret", j)
	s = append(s, v.secret)
		j++

	return s, nil
}

// UpdateFields updates one or more columns, given a 'where' clause.
//
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl UUserTable) UpdateFields(req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error) {
	return support.UpdateFields(tbl, req, wh, fields...)
}

//--------------------------------------------------------------------------------

var allUUserQuotedUpdates = []string{
	// Sqlite
	"`name`=?,`emailaddress`=?,`addressid`=?,`avatar`=?,`role`=?,`active`=?,`admin`=?,`fave`=?,`lastupdated`=?,`i8`=?,`u8`=?,`i16`=?,`u16`=?,`i32`=?,`u32`=?,`i64`=?,`u64`=?,`f32`=?,`f64`=?,`token`=?,`secret`=? WHERE `uid`=?",
	// Mysql
	"`name`=?,`emailaddress`=?,`addressid`=?,`avatar`=?,`role`=?,`active`=?,`admin`=?,`fave`=?,`lastupdated`=?,`i8`=?,`u8`=?,`i16`=?,`u16`=?,`i32`=?,`u32`=?,`i64`=?,`u64`=?,`f32`=?,`f64`=?,`token`=?,`secret`=? WHERE `uid`=?",
	// Postgres
	`"name"=$2,"emailaddress"=$3,"addressid"=$4,"avatar"=$5,"role"=$6,"active"=$7,"admin"=$8,"fave"=$9,"lastupdated"=$10,"i8"=$11,"u8"=$12,"i16"=$13,"u16"=$14,"i32"=$15,"u32"=$16,"i64"=$17,"u64"=$18,"f32"=$19,"f64"=$20,"token"=$21,"secret"=$22 WHERE "uid"=$1`,
}

//--------------------------------------------------------------------------------

// Update updates records, matching them by primary key. It returns the number of rows affected.
// The User.PreUpdate(Execer) method will be called, if it exists.
func (tbl UUserTable) Update(req require.Requirement, vv ...*User) (int64, error) {
	if req == require.All {
		req = require.Exactly(len(vv))
	}

	var count int64
	dialect := tbl.Dialect()
	//columns := allUUserQuotedUpdates[dialect.Index()]
	//query := fmt.Sprintf("UPDATE %s SET %s", tbl.name, columns)

	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPreUpdate); ok {
			err := hook.PreUpdate()
			if err != nil {
				return count, tbl.logError(err)
			}
		}

		b := &bytes.Buffer{}
		io.WriteString(b, "UPDATE ")
		io.WriteString(b, tbl.name.String())
		io.WriteString(b, " SET ")

		args, err := constructUUserUpdate(b, v, dialect)
		k := len(args) + 1
		args = append(args, v.Uid)
		if err != nil {
			return count, tbl.logError(err)
		}

		io.WriteString(b, " WHERE ")
		dialect.QuoteWithPlaceholder(b, tbl.pk, k)

		query := b.String()
		n, err := tbl.Exec(nil, query, args...)
		if err != nil {
			return count, err
		}
		count += n
	}

	return count, tbl.logIfError(require.ErrorIfExecNotSatisfiedBy(req, count))
}
