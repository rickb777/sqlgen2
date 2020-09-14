// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.
// sqlapi v0.47.1; sqlgen v0.67.0

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
	"github.com/rickb777/sqlapi/support"
	"github.com/rickb777/where"
	"github.com/rickb777/where/quote"
	"math/big"
	"strings"
)

// DbUserTabler lists table methods provided by DbUserTable.
type DbUserTabler interface {
	sqlapi.Table

	// Constraints returns the table's constraints.
	Constraints() constraint.Constraints

	// WithConstraint returns a modified DbUserTabler with added data consistency constraints.
	WithConstraint(cc ...constraint.Constraint) DbUserTabler

	// WithPrefix returns a modified DbUserTabler with a given table name prefix.
	WithPrefix(pfx string) DbUserTabler

	// CreateTable creates the table.
	CreateTable(ctx context.Context, ifNotExists bool) (int64, error)

	// DropTable drops the table, destroying all its data.
	DropTable(ctx context.Context, ifExists bool) (int64, error)

	// CreateTableWithIndexes invokes CreateTable then CreateIndexes.
	CreateTableWithIndexes(ctx context.Context, ifNotExist bool) (err error)

	// CreateIndexes executes queries that create the indexes needed by the User table.
	CreateIndexes(ctx context.Context, ifNotExist bool) (err error)

	// CreateEmailaddressIdxIndex creates the emailaddress_idx index.
	CreateEmailaddressIdxIndex(ctx context.Context, ifNotExist bool) error

	// DropEmailaddressIdxIndex drops the emailaddress_idx index.
	DropEmailaddressIdxIndex(ctx context.Context, ifExists bool) error

	// CreateUserLoginIndex creates the user_login index.
	CreateUserLoginIndex(ctx context.Context, ifNotExist bool) error

	// DropUserLoginIndex drops the user_login index.
	DropUserLoginIndex(ctx context.Context, ifExists bool) error

	// Truncate drops every record from the table, if possible.
	Truncate(ctx context.Context, force bool) (err error)
}

//-------------------------------------------------------------------------------------------------

// DbUserQueryer lists query methods provided by DbUserTable.
type DbUserQueryer interface {
	sqlapi.Table

	// Using returns a modified DbUserQueryer using the Execer supplied,
	// which will typically be a transaction (i.e. SqlTx).
	Using(tx sqlapi.Execer) DbUserQueryer

	// Transact runs the function provided within a transaction. The transction is committed
	// unless an error occurs.
	Transact(ctx context.Context, txOptions *sql.TxOptions, fn func(DbUserQueryer) error) error

	// Exec executes a query without returning any rows.
	Exec(ctx context.Context, req require.Requirement, query string, args ...interface{}) (int64, error)

	// Query is the low-level request method for this table using an SQL query that must return all the columns
	// necessary for User values.
	Query(ctx context.Context, req require.Requirement, query string, args ...interface{}) ([]*User, error)

	// QueryOneNullString is a low-level access method for one string, returning the first match.
	QueryOneNullString(ctx context.Context, req require.Requirement, query string, args ...interface{}) (result sql.NullString, err error)

	// QueryOneNullInt64 is a low-level access method for one int64, returning the first match.
	QueryOneNullInt64(ctx context.Context, req require.Requirement, query string, args ...interface{}) (result sql.NullInt64, err error)

	// QueryOneNullFloat64 is a low-level access method for one float64, returning the first match.
	QueryOneNullFloat64(ctx context.Context, req require.Requirement, query string, args ...interface{}) (result sql.NullFloat64, err error)

	// GetUserByUid gets the record with a given primary key value.
	GetUserByUid(ctx context.Context, req require.Requirement, id int64) (*User, error)

	// GetUsersByUid gets records from the table according to a list of primary keys.
	GetUsersByUid(ctx context.Context, req require.Requirement, qc where.QueryConstraint, id ...int64) (list []*User, err error)

	// GetUserByEmailAddress gets the record with a given emailaddress value.
	GetUserByEmailAddress(ctx context.Context, req require.Requirement, emailaddress string) (*User, error)

	// GetUsersByEmailAddress gets the record with a given emailaddress value.
	GetUsersByEmailAddress(ctx context.Context, req require.Requirement, qc where.QueryConstraint, emailaddress ...string) ([]*User, error)

	// GetUserByName gets the record with a given name value.
	GetUserByName(ctx context.Context, req require.Requirement, name string) (*User, error)

	// GetUsersByName gets the record with a given name value.
	GetUsersByName(ctx context.Context, req require.Requirement, qc where.QueryConstraint, name ...string) ([]*User, error)

	// Fetch fetches a list of User based on a supplied query. This is mostly used for join queries that map its
	// result columns to the fields of User. Other queries might be better handled by GetXxx or Select methods.
	Fetch(ctx context.Context, req require.Requirement, query string, args ...interface{}) ([]*User, error)

	// SelectOneWhere allows a single User to be obtained from the table that matches a 'where' clause.
	SelectOneWhere(ctx context.Context, req require.Requirement, where, orderBy string, args ...interface{}) (*User, error)

	// SelectOne allows a single User to be obtained from the table that matches a 'where' clause.
	SelectOne(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) (*User, error)

	// SelectWhere allows Users to be obtained from the table that match a 'where' clause.
	SelectWhere(ctx context.Context, req require.Requirement, where, orderBy string, args ...interface{}) ([]*User, error)

	// Select allows Users to be obtained from the table that match a 'where' clause.
	Select(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]*User, error)

	// CountWhere counts Users in the table that match a 'where' clause.
	CountWhere(ctx context.Context, where string, args ...interface{}) (count int64, err error)

	// Count counts the Users in the table that match a 'where' clause.
	Count(ctx context.Context, wh where.Expression) (count int64, err error)

	// SliceUid gets the uid column for all rows that match the 'where' condition.
	SliceUid(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error)

	// SliceName gets the name column for all rows that match the 'where' condition.
	SliceName(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error)

	// SliceEmailaddress gets the emailaddress column for all rows that match the 'where' condition.
	SliceEmailaddress(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error)

	// SliceAddressid gets the addressid column for all rows that match the 'where' condition.
	SliceAddressid(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error)

	// SliceAvatar gets the avatar column for all rows that match the 'where' condition.
	SliceAvatar(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error)

	// SliceLastupdated gets the lastupdated column for all rows that match the 'where' condition.
	SliceLastupdated(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error)

	// SliceI8 gets the i8 column for all rows that match the 'where' condition.
	SliceI8(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int8, error)

	// SliceU8 gets the u8 column for all rows that match the 'where' condition.
	SliceU8(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]uint8, error)

	// SliceI16 gets the i16 column for all rows that match the 'where' condition.
	SliceI16(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int16, error)

	// SliceU16 gets the u16 column for all rows that match the 'where' condition.
	SliceU16(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]uint16, error)

	// SliceI32 gets the i32 column for all rows that match the 'where' condition.
	SliceI32(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int32, error)

	// SliceU32 gets the u32 column for all rows that match the 'where' condition.
	SliceU32(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]uint32, error)

	// SliceI64 gets the i64 column for all rows that match the 'where' condition.
	SliceI64(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error)

	// SliceU64 gets the u64 column for all rows that match the 'where' condition.
	SliceU64(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]uint64, error)

	// SliceRole gets the role column for all rows that match the 'where' condition.
	SliceRole(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]Role, error)

	// SliceF32 gets the f32 column for all rows that match the 'where' condition.
	SliceF32(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]float32, error)

	// SliceF64 gets the f64 column for all rows that match the 'where' condition.
	SliceF64(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]float64, error)

	// Insert adds new records for the Users, setting the primary key field for each one.
	Insert(ctx context.Context, req require.Requirement, vv ...*User) error

	// UpdateByUid updates one or more columns, given a uid value.
	UpdateByUid(ctx context.Context, req require.Requirement, uid int64, fields ...sql.NamedArg) (int64, error)

	// UpdateByName updates one or more columns, given a name value.
	UpdateByName(ctx context.Context, req require.Requirement, name string, fields ...sql.NamedArg) (int64, error)

	// UpdateByEmailaddress updates one or more columns, given a emailaddress value.
	UpdateByEmailaddress(ctx context.Context, req require.Requirement, emailaddress string, fields ...sql.NamedArg) (int64, error)

	// UpdateByAddressid updates one or more columns, given a addressid value.
	UpdateByAddressid(ctx context.Context, req require.Requirement, addressid int64, fields ...sql.NamedArg) (int64, error)

	// UpdateByAvatar updates one or more columns, given a avatar value.
	UpdateByAvatar(ctx context.Context, req require.Requirement, avatar string, fields ...sql.NamedArg) (int64, error)

	// UpdateByRole updates one or more columns, given a role value.
	UpdateByRole(ctx context.Context, req require.Requirement, role Role, fields ...sql.NamedArg) (int64, error)

	// UpdateByLastupdated updates one or more columns, given a lastupdated value.
	UpdateByLastupdated(ctx context.Context, req require.Requirement, lastupdated int64, fields ...sql.NamedArg) (int64, error)

	// UpdateByI8 updates one or more columns, given a i8 value.
	UpdateByI8(ctx context.Context, req require.Requirement, i8 int8, fields ...sql.NamedArg) (int64, error)

	// UpdateByU8 updates one or more columns, given a u8 value.
	UpdateByU8(ctx context.Context, req require.Requirement, u8 uint8, fields ...sql.NamedArg) (int64, error)

	// UpdateByI16 updates one or more columns, given a i16 value.
	UpdateByI16(ctx context.Context, req require.Requirement, i16 int16, fields ...sql.NamedArg) (int64, error)

	// UpdateByU16 updates one or more columns, given a u16 value.
	UpdateByU16(ctx context.Context, req require.Requirement, u16 uint16, fields ...sql.NamedArg) (int64, error)

	// UpdateByI32 updates one or more columns, given a i32 value.
	UpdateByI32(ctx context.Context, req require.Requirement, i32 int32, fields ...sql.NamedArg) (int64, error)

	// UpdateByU32 updates one or more columns, given a u32 value.
	UpdateByU32(ctx context.Context, req require.Requirement, u32 uint32, fields ...sql.NamedArg) (int64, error)

	// UpdateByI64 updates one or more columns, given a i64 value.
	UpdateByI64(ctx context.Context, req require.Requirement, i64 int64, fields ...sql.NamedArg) (int64, error)

	// UpdateByU64 updates one or more columns, given a u64 value.
	UpdateByU64(ctx context.Context, req require.Requirement, u64 uint64, fields ...sql.NamedArg) (int64, error)

	// UpdateByF32 updates one or more columns, given a f32 value.
	UpdateByF32(ctx context.Context, req require.Requirement, f32 float32, fields ...sql.NamedArg) (int64, error)

	// UpdateByF64 updates one or more columns, given a f64 value.
	UpdateByF64(ctx context.Context, req require.Requirement, f64 float64, fields ...sql.NamedArg) (int64, error)

	// UpdateFields updates one or more columns, given a 'where' clause.
	UpdateFields(ctx context.Context, req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error)

	// Update updates records, matching them by primary key.
	Update(ctx context.Context, req require.Requirement, vv ...*User) (int64, error)

	// Upsert inserts or updates a record, matching it using the expression supplied.
	// This expression is used to search for an existing record based on some specified
	// key column(s). It must match either zero or one existing record. If it matches
	// none, a new record is inserted; otherwise the matching record is updated. An
	// error results if these conditions are not met.
	Upsert(ctx context.Context, v *User, wh where.Expression) error

	// DeleteByUid deletes rows from the table, given some uid values.
	// The list of ids can be arbitrarily long.
	DeleteByUid(ctx context.Context, req require.Requirement, uid ...int64) (int64, error)

	// DeleteByName deletes rows from the table, given some name values.
	// The list of ids can be arbitrarily long.
	DeleteByName(ctx context.Context, req require.Requirement, name ...string) (int64, error)

	// DeleteByEmailaddress deletes rows from the table, given some emailaddress values.
	// The list of ids can be arbitrarily long.
	DeleteByEmailaddress(ctx context.Context, req require.Requirement, emailaddress ...string) (int64, error)

	// DeleteByAddressid deletes rows from the table, given some addressid values.
	// The list of ids can be arbitrarily long.
	DeleteByAddressid(ctx context.Context, req require.Requirement, addressid ...int64) (int64, error)

	// DeleteByAvatar deletes rows from the table, given some avatar values.
	// The list of ids can be arbitrarily long.
	DeleteByAvatar(ctx context.Context, req require.Requirement, avatar ...string) (int64, error)

	// DeleteByRole deletes rows from the table, given some role values.
	// The list of ids can be arbitrarily long.
	DeleteByRole(ctx context.Context, req require.Requirement, role ...Role) (int64, error)

	// DeleteByLastupdated deletes rows from the table, given some lastupdated values.
	// The list of ids can be arbitrarily long.
	DeleteByLastupdated(ctx context.Context, req require.Requirement, lastupdated ...int64) (int64, error)

	// DeleteByI8 deletes rows from the table, given some i8 values.
	// The list of ids can be arbitrarily long.
	DeleteByI8(ctx context.Context, req require.Requirement, i8 ...int8) (int64, error)

	// DeleteByU8 deletes rows from the table, given some u8 values.
	// The list of ids can be arbitrarily long.
	DeleteByU8(ctx context.Context, req require.Requirement, u8 ...uint8) (int64, error)

	// DeleteByI16 deletes rows from the table, given some i16 values.
	// The list of ids can be arbitrarily long.
	DeleteByI16(ctx context.Context, req require.Requirement, i16 ...int16) (int64, error)

	// DeleteByU16 deletes rows from the table, given some u16 values.
	// The list of ids can be arbitrarily long.
	DeleteByU16(ctx context.Context, req require.Requirement, u16 ...uint16) (int64, error)

	// DeleteByI32 deletes rows from the table, given some i32 values.
	// The list of ids can be arbitrarily long.
	DeleteByI32(ctx context.Context, req require.Requirement, i32 ...int32) (int64, error)

	// DeleteByU32 deletes rows from the table, given some u32 values.
	// The list of ids can be arbitrarily long.
	DeleteByU32(ctx context.Context, req require.Requirement, u32 ...uint32) (int64, error)

	// DeleteByI64 deletes rows from the table, given some i64 values.
	// The list of ids can be arbitrarily long.
	DeleteByI64(ctx context.Context, req require.Requirement, i64 ...int64) (int64, error)

	// DeleteByU64 deletes rows from the table, given some u64 values.
	// The list of ids can be arbitrarily long.
	DeleteByU64(ctx context.Context, req require.Requirement, u64 ...uint64) (int64, error)

	// DeleteByF32 deletes rows from the table, given some f32 values.
	// The list of ids can be arbitrarily long.
	DeleteByF32(ctx context.Context, req require.Requirement, f32 ...float32) (int64, error)

	// DeleteByF64 deletes rows from the table, given some f64 values.
	// The list of ids can be arbitrarily long.
	DeleteByF64(ctx context.Context, req require.Requirement, f64 ...float64) (int64, error)

	// Delete deletes one or more rows from the table, given a 'where' clause.
	// Use a nil value for the 'wh' argument if it is not needed (very risky!).
	Delete(ctx context.Context, req require.Requirement, wh where.Expression) (int64, error)
}

//-------------------------------------------------------------------------------------------------

// DbUserTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type DbUserTable struct {
	name        sqlapi.TableName
	database    sqlapi.Database
	db          sqlapi.Execer
	constraints constraint.Constraints
	pk          string
}

// Type conformance checks
var _ sqlapi.TableWithIndexes = &DbUserTable{}

// NewDbUserTable returns a new table instance.
// If a blank table name is supplied, the default name "the_users" will be used instead.
// The request context is initialised with the background.
func NewDbUserTable(name string, d sqlapi.Database) DbUserTable {
	if name == "" {
		name = "the_users"
	}
	var constraints constraint.Constraints
	constraints = append(constraints,
		constraint.FkConstraint{"addressid", constraint.Reference{"addresses", "id"}, "restrict", "restrict"})
	return DbUserTable{
		name:        sqlapi.TableName{Prefix: "", Name: name},
		database:    d,
		db:          d.DB(),
		constraints: constraints,
		pk:          "uid",
	}
}

// CopyTableAsDbUserTable copies a table instance, retaining the name etc but
// providing methods appropriate for 'User'.It doesn't copy the constraints of the original table.
//
// It serves to provide methods appropriate for 'User'. This is most useful when this is used to represent a
// join result. In such cases, there won't be any need for DDL methods, nor Exec, Insert, Update or Delete.
func CopyTableAsDbUserTable(origin sqlapi.Table) DbUserTable {
	return DbUserTable{
		name:        origin.Name(),
		database:    origin.Database(),
		db:          origin.Execer(),
		constraints: nil,
		pk:          "uid",
	}
}

// SetPkColumn sets the name of the primary key column. It defaults to "uid".
// The result is a modified copy of the table; the original is unchanged.
//func (tbl DbUserTable) SetPkColumn(pk string) DbUserTabler {
//	tbl.pk = pk
//	return tbl
//}

// WithPrefix sets the table name prefix for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl DbUserTable) WithPrefix(pfx string) DbUserTabler {
	tbl.name.Prefix = pfx
	return tbl
}

// Database gets the shared database information.
func (tbl DbUserTable) Database() sqlapi.Database {
	return tbl.database
}

// Logger gets the trace logger.
func (tbl DbUserTable) Logger() sqlapi.Logger {
	return tbl.database.Logger()
}

// WithConstraint returns a modified DbUserTabler with added data consistency constraints.
func (tbl DbUserTable) WithConstraint(cc ...constraint.Constraint) DbUserTabler {
	tbl.constraints = append(tbl.constraints, cc...)
	return tbl
}

// Constraints returns the table's constraints.
func (tbl DbUserTable) Constraints() constraint.Constraints {
	return tbl.constraints
}

// Dialect gets the database dialect.
func (tbl DbUserTable) Dialect() dialect.Dialect {
	return tbl.database.Dialect()
}

// Name gets the table name.
func (tbl DbUserTable) Name() sqlapi.TableName {
	return tbl.name
}

// PkColumn gets the column name used as a primary key.
func (tbl DbUserTable) PkColumn() string {
	return tbl.pk
}

// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl DbUserTable) DB() sqlapi.SqlDB {
	return tbl.db.(sqlapi.SqlDB)
}

// Execer gets the wrapped database or transaction handle.
func (tbl DbUserTable) Execer() sqlapi.Execer {
	return tbl.db
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl DbUserTable) Tx() sqlapi.SqlTx {
	return tbl.db.(sqlapi.SqlTx)
}

// IsTx tests whether this is within a transaction.
func (tbl DbUserTable) IsTx() bool {
	return tbl.db.IsTx()
}

// Using returns a modified DbUserTabler using the the Execer supplied,
// which will typically be a transaction (i.e. SqlTx). This is needed when making multiple
// queries across several tables within a single transaction.
//
// The result is a modified copy of the table; the original is unchanged.
func (tbl DbUserTable) Using(tx sqlapi.Execer) DbUserQueryer {
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
func (tbl DbUserTable) Transact(ctx context.Context, txOptions *sql.TxOptions, fn func(DbUserQueryer) error) error {
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

func (tbl DbUserTable) quotedName() string {
	return tbl.Dialect().Quoter().Quote(tbl.name.String())
}

func (tbl DbUserTable) quotedNameW(w dialect.StringWriter) {
	tbl.Dialect().Quoter().QuoteW(w, tbl.name.String())
}

//-------------------------------------------------------------------------------------------------

// NumDbUserTableColumns is the total number of columns in DbUserTable.
const NumDbUserTableColumns = 22

// NumDbUserTableDataColumns is the number of columns in DbUserTable not including the auto-increment key.
const NumDbUserTableDataColumns = 21

// DbUserTableColumnNames is the list of columns in DbUserTable.
const DbUserTableColumnNames = "uid,name,emailaddress,addressid,avatar,role,active,admin,fave,lastupdated,i8,u8,i16,u16,i32,u32,i64,u64,f32,f64,token,secret"

// DbUserTableDataColumnNames is the list of data columns in DbUserTable.
const DbUserTableDataColumnNames = "name,emailaddress,addressid,avatar,role,active,admin,fave,lastupdated,i8,u8,i16,u16,i32,u32,i64,u64,f32,f64,token,secret"

var listOfDbUserTableColumnNames = strings.Split(DbUserTableColumnNames, ",")

//-------------------------------------------------------------------------------------------------

var sqlDbUserTableCreateColumnsSqlite = []string{
	"integer not null primary key autoincrement",
	"text not null",
	"text not null",
	"bigint default null",
	"text default null",
	"text default null",
	"boolean not null",
	"boolean not null",
	"text",
	"bigint not null",
	"tinyint not null default -8",
	"tinyint unsigned not null default 8",
	"smallint not null default -16",
	"smallint unsigned not null default 16",
	"int not null default -32",
	"int unsigned not null default 32",
	"bigint not null default -64",
	"bigint unsigned not null default 64",
	"float not null default 3.2",
	"double not null default 6.4",
	"text not null",
	"text not null",
}

var sqlDbUserTableCreateColumnsMysql = []string{
	"bigint not null primary key auto_increment",
	"varchar(255) not null",
	"varchar(255) not null",
	"bigint default null",
	"text default null",
	"varchar(20) default null",
	"boolean not null",
	"boolean not null",
	"json",
	"bigint not null",
	"tinyint not null default -8",
	"tinyint unsigned not null default 8",
	"smallint not null default -16",
	"smallint unsigned not null default 16",
	"int not null default -32",
	"int unsigned not null default 32",
	"bigint not null default -64",
	"bigint unsigned not null default 64",
	"float not null default 3.2",
	"double not null default 6.4",
	"text not null",
	"text not null",
}

var sqlDbUserTableCreateColumnsPostgres = []string{
	"bigserial not null primary key",
	"text not null",
	"text not null",
	"bigint default null",
	"text default null",
	"text default null",
	"boolean not null",
	"boolean not null",
	"json",
	"bigint not null",
	"int8 not null default -8",
	"smallint not null default 8",
	"smallint not null default -16",
	"integer not null default 16",
	"integer not null default -32",
	"bigint not null default 32",
	"bigint not null default -64",
	"bigint not null default 64",
	"real not null default 3.2",
	"double precision not null default 6.4",
	"text not null",
	"text not null",
}

var sqlDbUserTableCreateColumnsPgx = []string{
	"bigserial not null primary key",
	"text not null",
	"text not null",
	"bigint default null",
	"text default null",
	"text default null",
	"boolean not null",
	"boolean not null",
	"json",
	"bigint not null",
	"int8 not null default -8",
	"smallint not null default 8",
	"smallint not null default -16",
	"integer not null default 16",
	"integer not null default -32",
	"bigint not null default 32",
	"bigint not null default -64",
	"bigint not null default 64",
	"real not null default 3.2",
	"double precision not null default 6.4",
	"text not null",
	"text not null",
}

//-------------------------------------------------------------------------------------------------

const sqlDbEmailaddressIdxIndexColumns = "emailaddress"

var listOfDbEmailaddressIdxIndexColumns = []string{"emailaddress"}

const sqlDbUserLoginIndexColumns = "name"

var listOfDbUserLoginIndexColumns = []string{"name"}

//-------------------------------------------------------------------------------------------------

// CreateTable creates the table.
func (tbl DbUserTable) CreateTable(ctx context.Context, ifNotExists bool) (int64, error) {
	return support.Exec(ctx, tbl, nil, createDbUserTableSql(tbl, ifNotExists))
}

func createDbUserTableSql(tbl DbUserTabler, ifNotExists bool) string {
	buf := &bytes.Buffer{}
	buf.WriteString("CREATE TABLE ")
	if ifNotExists {
		buf.WriteString("IF NOT EXISTS ")
	}
	q := tbl.Dialect().Quoter()
	q.QuoteW(buf, tbl.Name().String())
	buf.WriteString(" (\n ")

	var columns []string
	switch tbl.Dialect().Index() {
	case dialect.SqliteIndex:
		columns = sqlDbUserTableCreateColumnsSqlite
	case dialect.MysqlIndex:
		columns = sqlDbUserTableCreateColumnsMysql
	case dialect.PostgresIndex:
		columns = sqlDbUserTableCreateColumnsPostgres
	case dialect.PgxIndex:
		columns = sqlDbUserTableCreateColumnsPgx
	}

	comma := ""
	for i, n := range listOfDbUserTableColumnNames {
		buf.WriteString(comma)
		q.QuoteW(buf, n)
		buf.WriteString(" ")
		buf.WriteString(columns[i])
		comma = ",\n "
	}

	for i, c := range tbl.(DbUserTable).Constraints() {
		buf.WriteString(",\n ")
		buf.WriteString(c.ConstraintSql(tbl.Dialect().Quoter(), tbl.Name(), i+1))
	}

	buf.WriteString("\n)")
	buf.WriteString(tbl.Dialect().CreateTableSettings())
	return buf.String()
}

func ternaryDbUserTable(flag bool, a, b string) string {
	if flag {
		return a
	}
	return b
}

// DropTable drops the table, destroying all its data.
func (tbl DbUserTable) DropTable(ctx context.Context, ifExists bool) (int64, error) {
	return support.Exec(ctx, tbl, nil, dropDbUserTableSql(tbl, ifExists))
}

func dropDbUserTableSql(tbl DbUserTabler, ifExists bool) string {
	ie := ternaryDbUserTable(ifExists, "IF EXISTS ", "")
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("DROP TABLE %s%s", ie, quotedName)
	return query
}

//-------------------------------------------------------------------------------------------------

// CreateTableWithIndexes invokes CreateTable then CreateIndexes.
func (tbl DbUserTable) CreateTableWithIndexes(ctx context.Context, ifNotExist bool) (err error) {
	_, err = tbl.CreateTable(ctx, ifNotExist)
	if err != nil {
		return err
	}

	return tbl.CreateIndexes(ctx, ifNotExist)
}

// CreateIndexes executes queries that create the indexes needed by the User table.
func (tbl DbUserTable) CreateIndexes(ctx context.Context, ifNotExist bool) (err error) {

	err = tbl.CreateEmailaddressIdxIndex(ctx, ifNotExist)
	if err != nil {
		return err
	}

	err = tbl.CreateUserLoginIndex(ctx, ifNotExist)
	if err != nil {
		return err
	}

	return nil
}

// CreateEmailaddressIdxIndex creates the emailaddress_idx index.
func (tbl DbUserTable) CreateEmailaddressIdxIndex(ctx context.Context, ifNotExist bool) error {
	ine := ternaryDbUserTable(ifNotExist && tbl.Dialect().Index() != dialect.MysqlIndex, "IF NOT EXISTS ", "")

	// Mysql does not support 'if not exists' on indexes
	// Workaround: use DropIndex first and ignore an error returned if the index didn't exist.

	if ifNotExist && tbl.Dialect().Index() == dialect.MysqlIndex {
		// low-level no-logging Exec
		tbl.Execer().ExecContext(ctx, dropDbUserTableEmailaddressIdxSql(tbl, false))
		ine = ""
	}

	_, err := tbl.Exec(ctx, nil, createDbUserTableEmailaddressIdxSql(tbl, ine))
	return err
}

func createDbUserTableEmailaddressIdxSql(tbl DbUserTabler, ifNotExists string) string {
	indexPrefix := tbl.Name().PrefixWithoutDot()
	id := fmt.Sprintf("%s%s_emailaddress_idx", indexPrefix, tbl.Name().Name)
	q := tbl.Dialect().Quoter()
	cols := strings.Join(q.QuoteN(listOfDbEmailaddressIdxIndexColumns), ",")
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	return fmt.Sprintf("CREATE UNIQUE INDEX %s%s ON %s (%s)", ifNotExists,
		q.Quote(id), quotedName, cols)
}

// DropEmailaddressIdxIndex drops the emailaddress_idx index.
func (tbl DbUserTable) DropEmailaddressIdxIndex(ctx context.Context, ifExists bool) error {
	_, err := tbl.Exec(ctx, nil, dropDbUserTableEmailaddressIdxSql(tbl, ifExists))
	return err
}

func dropDbUserTableEmailaddressIdxSql(tbl DbUserTabler, ifExists bool) string {
	// Mysql does not support 'if exists' on indexes
	ie := ternaryDbUserTable(ifExists && tbl.Dialect().Index() != dialect.MysqlIndex, "IF EXISTS ", "")
	indexPrefix := tbl.Name().PrefixWithoutDot()
	id := fmt.Sprintf("%s%s_emailaddress_idx", indexPrefix, tbl.Name().Name)
	q := tbl.Dialect().Quoter()
	// Mysql requires extra "ON tbl" clause
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	onTbl := ternaryDbUserTable(tbl.Dialect().Index() == dialect.MysqlIndex, fmt.Sprintf(" ON %s", quotedName), "")
	return "DROP INDEX " + ie + q.Quote(id) + onTbl
}

// CreateUserLoginIndex creates the user_login index.
func (tbl DbUserTable) CreateUserLoginIndex(ctx context.Context, ifNotExist bool) error {
	ine := ternaryDbUserTable(ifNotExist && tbl.Dialect().Index() != dialect.MysqlIndex, "IF NOT EXISTS ", "")

	// Mysql does not support 'if not exists' on indexes
	// Workaround: use DropIndex first and ignore an error returned if the index didn't exist.

	if ifNotExist && tbl.Dialect().Index() == dialect.MysqlIndex {
		// low-level no-logging Exec
		tbl.Execer().ExecContext(ctx, dropDbUserTableUserLoginSql(tbl, false))
		ine = ""
	}

	_, err := tbl.Exec(ctx, nil, createDbUserTableUserLoginSql(tbl, ine))
	return err
}

func createDbUserTableUserLoginSql(tbl DbUserTabler, ifNotExists string) string {
	indexPrefix := tbl.Name().PrefixWithoutDot()
	id := fmt.Sprintf("%s%s_user_login", indexPrefix, tbl.Name().Name)
	q := tbl.Dialect().Quoter()
	cols := strings.Join(q.QuoteN(listOfDbUserLoginIndexColumns), ",")
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	return fmt.Sprintf("CREATE UNIQUE INDEX %s%s ON %s (%s)", ifNotExists,
		q.Quote(id), quotedName, cols)
}

// DropUserLoginIndex drops the user_login index.
func (tbl DbUserTable) DropUserLoginIndex(ctx context.Context, ifExists bool) error {
	_, err := tbl.Exec(ctx, nil, dropDbUserTableUserLoginSql(tbl, ifExists))
	return err
}

func dropDbUserTableUserLoginSql(tbl DbUserTabler, ifExists bool) string {
	// Mysql does not support 'if exists' on indexes
	ie := ternaryDbUserTable(ifExists && tbl.Dialect().Index() != dialect.MysqlIndex, "IF EXISTS ", "")
	indexPrefix := tbl.Name().PrefixWithoutDot()
	id := fmt.Sprintf("%s%s_user_login", indexPrefix, tbl.Name().Name)
	q := tbl.Dialect().Quoter()
	// Mysql requires extra "ON tbl" clause
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	onTbl := ternaryDbUserTable(tbl.Dialect().Index() == dialect.MysqlIndex, fmt.Sprintf(" ON %s", quotedName), "")
	return "DROP INDEX " + ie + q.Quote(id) + onTbl
}

// DropIndexes executes queries that drop the indexes on by the User table.
func (tbl DbUserTable) DropIndexes(ctx context.Context, ifExist bool) (err error) {

	err = tbl.DropEmailaddressIdxIndex(ctx, ifExist)
	if err != nil {
		return err
	}

	err = tbl.DropUserLoginIndex(ctx, ifExist)
	if err != nil {
		return err
	}

	return nil
}

//-------------------------------------------------------------------------------------------------

// Truncate drops every record from the table, if possible. It might fail if constraints exist that
// prevent some or all rows from being deleted; use the force option to override this.
//
// When 'force' is set true, be aware of the following consequences.
// When using Mysql, foreign keys in other tables can be left dangling.
// When using Postgres, a cascade happens, so all 'adjacent' tables (i.e. linked by foreign keys)
// are also truncated.
func (tbl DbUserTable) Truncate(ctx context.Context, force bool) (err error) {
	for _, query := range tbl.Dialect().TruncateDDL(tbl.Name().String(), force) {
		_, err = support.Exec(ctx, tbl, nil, query)
		if err != nil {
			return err
		}
	}
	return nil
}

//-------------------------------------------------------------------------------------------------

// Exec executes a query without returning any rows.
// It returns the number of rows affected (if the database driver supports this).
//
// The args are for any placeholder parameters in the query.
//
// If the context ctx is nil, it defaults to context.Background().
func (tbl DbUserTable) Exec(ctx context.Context, req require.Requirement, query string, args ...interface{}) (int64, error) {
	return support.Exec(ctx, tbl, req, query, args...)
}

//-------------------------------------------------------------------------------------------------

// Query is the low-level request method for this table. The SQL query must return all the columns necessary for
// User values. Placeholders should be vanilla '?' marks, which will be replaced if necessary according to
// the chosen dialect.
//
// The query is logged using whatever logger is configured. If an error arises, this too is logged.
//
// The args are for any placeholder parameters in the query.
//
// The support API provides a core 'support.Query' function, on which this method depends. If appropriate,
// use that function directly; wrap the result in *sqlapi.Rows if you need to access its data as a map.
//
// If the context ctx is nil, it defaults to context.Background().
func (tbl DbUserTable) Query(ctx context.Context, req require.Requirement, query string, args ...interface{}) ([]*User, error) {
	return doDbUserTableQueryAndScan(ctx, tbl, req, false, query, args)
}

func doDbUserTableQueryAndScan(ctx context.Context, tbl DbUserTabler, req require.Requirement, firstOnly bool, query string, args ...interface{}) ([]*User, error) {
	rows, err := support.Query(ctx, tbl, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vv, n, err := ScanDbUsers(query, rows, firstOnly)
	return vv, tbl.Logger().LogIfError(require.ChainErrorIfQueryNotSatisfiedBy(err, req, n))
}

//-------------------------------------------------------------------------------------------------

// QueryOneNullString is a low-level access method for one string. This can be used for function queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl DbUserTable) QueryOneNullString(ctx context.Context, req require.Requirement, query string, args ...interface{}) (result sql.NullString, err error) {
	err = support.QueryOneNullThing(ctx, tbl, req, &result, query, args...)
	return result, err
}

// QueryOneNullInt64 is a low-level access method for one int64. This can be used for 'COUNT(1)' queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl DbUserTable) QueryOneNullInt64(ctx context.Context, req require.Requirement, query string, args ...interface{}) (result sql.NullInt64, err error) {
	err = support.QueryOneNullThing(ctx, tbl, req, &result, query, args...)
	return result, err
}

// QueryOneNullFloat64 is a low-level access method for one float64. This can be used for 'AVG(...)' queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl DbUserTable) QueryOneNullFloat64(ctx context.Context, req require.Requirement, query string, args ...interface{}) (result sql.NullFloat64, err error) {
	err = support.QueryOneNullThing(ctx, tbl, req, &result, query, args...)
	return result, err
}

// ScanDbUsers reads rows from the database and returns a slice of corresponding values.
// It also returns a number indicating how many rows were read; this will be larger than the length of the
// slice if reading stopped after the first row.
func ScanDbUsers(query string, rows sqlapi.SqlRows, firstOnly bool) (vv []*User, n int64, err error) {
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

//--------------------------------------------------------------------------------

func allDbUserColumnNamesQuoted(q quote.Quoter) string {
	return strings.Join(q.QuoteN(listOfDbUserTableColumnNames), ",")
}

//--------------------------------------------------------------------------------

// GetUserByUid gets the record with a given primary key value.
// If not found, *User will be nil.
func (tbl DbUserTable) GetUserByUid(ctx context.Context, req require.Requirement, id int64) (*User, error) {
	return tbl.SelectOne(ctx, req, where.Eq("uid", id), nil)
}

// GetUsersByUid gets records from the table according to a list of primary keys.
// Although the list of ids can be arbitrarily long, there are practical limits;
// note that Oracle DB has a limit of 1000.
//
// It places a requirement, which may be nil, on the size of the expected results: in particular, require.All
// controls whether an error is generated not all the ids produce a result.
func (tbl DbUserTable) GetUsersByUid(ctx context.Context, req require.Requirement, qc where.QueryConstraint, uid ...int64) (list []*User, err error) {
	if req == require.All {
		req = require.Exactly(len(uid))
	}
	return tbl.Select(ctx, req, where.In("uid", uid), qc)
}

// GetUserByEmailAddress gets the record with a given emailaddress value.
// If not found, *User will be nil.
func (tbl DbUserTable) GetUserByEmailAddress(ctx context.Context, req require.Requirement, emailaddress string) (*User, error) {
	return tbl.SelectOne(ctx, req, where.And(where.Eq("emailaddress", emailaddress)), nil)
}

// GetUsersByEmailAddress gets the record with a given emailaddress value.
func (tbl DbUserTable) GetUsersByEmailAddress(ctx context.Context, req require.Requirement, qc where.QueryConstraint, emailaddress ...string) ([]*User, error) {
	if req == require.All {
		req = require.Exactly(len(emailaddress))
	}
	return tbl.Select(ctx, req, where.In("emailaddress", emailaddress), qc)
}

// GetUserByName gets the record with a given name value.
// If not found, *User will be nil.
func (tbl DbUserTable) GetUserByName(ctx context.Context, req require.Requirement, name string) (*User, error) {
	return tbl.SelectOne(ctx, req, where.And(where.Eq("name", name)), nil)
}

// GetUsersByName gets the record with a given name value.
func (tbl DbUserTable) GetUsersByName(ctx context.Context, req require.Requirement, qc where.QueryConstraint, name ...string) ([]*User, error) {
	if req == require.All {
		req = require.Exactly(len(name))
	}
	return tbl.Select(ctx, req, where.In("name", name), qc)
}

func doDbUserTableQueryAndScanOne(ctx context.Context, tbl DbUserTabler, req require.Requirement, query string, args ...interface{}) (*User, error) {
	list, err := doDbUserTableQueryAndScan(ctx, tbl, req, true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

// Fetch fetches a list of User based on a supplied query. This is mostly used for join queries that map its
// result columns to the fields of User. Other queries might be better handled by GetXxx or Select methods.
func (tbl DbUserTable) Fetch(ctx context.Context, req require.Requirement, query string, args ...interface{}) ([]*User, error) {
	return doDbUserTableQueryAndScan(ctx, tbl, req, false, query, args...)
}

//-------------------------------------------------------------------------------------------------

// SelectOneWhere allows a single User to be obtained from the table that matches a 'where' clause
// and some limit. Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
// If not found, *Example will be nil.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.One
// controls whether an error is generated when no result is found.
//
// The args are for any placeholder parameters in the query.
func (tbl DbUserTable) SelectOneWhere(ctx context.Context, req require.Requirement, where, orderBy string, args ...interface{}) (*User, error) {
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT 1",
		allDbUserColumnNamesQuoted(tbl.Dialect().Quoter()), quotedName, where, orderBy)
	v, err := doDbUserTableQueryAndScanOne(ctx, tbl, req, query, args...)
	return v, err
}

// SelectOne allows a single User to be obtained from the table that matches a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
// If not found, *Example will be nil.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.One
// controls whether an error is generated when no result is found.
func (tbl DbUserTable) SelectOne(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) (*User, error) {
	q := tbl.Dialect().Quoter()
	whs, args := where.Where(wh, q)
	orderBy := where.Build(qc, q)
	return tbl.SelectOneWhere(ctx, req, whs, orderBy, args...)
}

// SelectWhere allows Users to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
//
// The args are for any placeholder parameters in the query.
func (tbl DbUserTable) SelectWhere(ctx context.Context, req require.Requirement, where, orderBy string, args ...interface{}) ([]*User, error) {
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("SELECT %s FROM %s %s %s",
		allDbUserColumnNamesQuoted(tbl.Dialect().Quoter()), quotedName, where, orderBy)
	vv, err := doDbUserTableQueryAndScan(ctx, tbl, req, false, query, args...)
	return vv, err
}

// Select allows Users to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
func (tbl DbUserTable) Select(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]*User, error) {
	q := tbl.Dialect().Quoter()
	whs, args := where.Where(wh, q)
	orderBy := where.Build(qc, q)
	return tbl.SelectWhere(ctx, req, whs, orderBy, args...)
}

//-------------------------------------------------------------------------------------------------

// CountWhere counts Users in the table that match a 'where' clause.
// Use a blank string for the 'where' argument if it is not needed.
//
// The args are for any placeholder parameters in the query.
func (tbl DbUserTable) CountWhere(ctx context.Context, where string, args ...interface{}) (count int64, err error) {
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s %s", quotedName, where)
	rows, err := support.Query(ctx, tbl, query, args...)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&count)
	}
	return count, tbl.Logger().LogIfError(err)
}

// Count counts the Users in the table that match a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed.
func (tbl DbUserTable) Count(ctx context.Context, wh where.Expression) (count int64, err error) {
	whs, args := where.Where(wh, tbl.Dialect().Quoter())
	return tbl.CountWhere(ctx, whs, args...)
}

//--------------------------------------------------------------------------------

// SliceUid gets the uid column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DbUserTable) SliceUid(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return support.SliceInt64List(ctx, tbl, req, tbl.pk, wh, qc)
}

// SliceName gets the name column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DbUserTable) SliceName(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return support.SliceStringList(ctx, tbl, req, "name", wh, qc)
}

// SliceEmailaddress gets the emailaddress column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DbUserTable) SliceEmailaddress(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return support.SliceStringList(ctx, tbl, req, "emailaddress", wh, qc)
}

// SliceAddressid gets the addressid column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DbUserTable) SliceAddressid(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return support.SliceInt64PtrList(ctx, tbl, req, "addressid", wh, qc)
}

// SliceAvatar gets the avatar column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DbUserTable) SliceAvatar(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return support.SliceStringPtrList(ctx, tbl, req, "avatar", wh, qc)
}

// SliceLastupdated gets the lastupdated column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DbUserTable) SliceLastupdated(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return support.SliceInt64List(ctx, tbl, req, "lastupdated", wh, qc)
}

// SliceI8 gets the i8 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DbUserTable) SliceI8(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int8, error) {
	return support.SliceInt8List(ctx, tbl, req, "i8", wh, qc)
}

// SliceU8 gets the u8 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DbUserTable) SliceU8(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]uint8, error) {
	return support.SliceUint8List(ctx, tbl, req, "u8", wh, qc)
}

// SliceI16 gets the i16 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DbUserTable) SliceI16(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int16, error) {
	return support.SliceInt16List(ctx, tbl, req, "i16", wh, qc)
}

// SliceU16 gets the u16 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DbUserTable) SliceU16(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]uint16, error) {
	return support.SliceUint16List(ctx, tbl, req, "u16", wh, qc)
}

// SliceI32 gets the i32 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DbUserTable) SliceI32(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int32, error) {
	return support.SliceInt32List(ctx, tbl, req, "i32", wh, qc)
}

// SliceU32 gets the u32 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DbUserTable) SliceU32(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]uint32, error) {
	return support.SliceUint32List(ctx, tbl, req, "u32", wh, qc)
}

// SliceI64 gets the i64 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DbUserTable) SliceI64(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return support.SliceInt64List(ctx, tbl, req, "i64", wh, qc)
}

// SliceU64 gets the u64 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DbUserTable) SliceU64(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]uint64, error) {
	return support.SliceUint64List(ctx, tbl, req, "u64", wh, qc)
}

// SliceRole gets the role column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DbUserTable) SliceRole(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]Role, error) {
	return sliceDbUserTableRolePtrList(ctx, tbl, req, "role", wh, qc)
}

// SliceF32 gets the f32 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DbUserTable) SliceF32(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]float32, error) {
	return sliceDbUserTableFloat32List(ctx, tbl, req, "f32", wh, qc)
}

// SliceF64 gets the f64 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl DbUserTable) SliceF64(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]float64, error) {
	return sliceDbUserTableFloat64List(ctx, tbl, req, "f64", wh, qc)
}

func sliceDbUserTableRolePtrList(ctx context.Context, tbl DbUserTabler, req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]Role, error) {
	q := tbl.Dialect().Quoter()
	whs, args := where.Where(wh, q)
	orderBy := where.Build(qc, q)
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", q.Quote(sqlname), quotedName, whs, orderBy)
	rows, err := support.Query(ctx, tbl, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]Role, 0, 10)

	for rows.Next() {
		var v Role
		err = rows.Scan(&v)
		if err == sql.ErrNoRows {
			return list, tbl.Logger().LogIfError(require.ErrorIfQueryNotSatisfiedBy(req, int64(len(list))))
		} else {
			list = append(list, v)
		}
	}
	return list, tbl.Logger().LogIfError(require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, int64(len(list))))
}

func sliceDbUserTableFloat32List(ctx context.Context, tbl DbUserTabler, req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]float32, error) {
	q := tbl.Dialect().Quoter()
	whs, args := where.Where(wh, q)
	orderBy := where.Build(qc, q)
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", q.Quote(sqlname), quotedName, whs, orderBy)
	rows, err := support.Query(ctx, tbl, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]float32, 0, 10)

	for rows.Next() {
		var v float32
		err = rows.Scan(&v)
		if err == sql.ErrNoRows {
			return list, tbl.Logger().LogIfError(require.ErrorIfQueryNotSatisfiedBy(req, int64(len(list))))
		} else {
			list = append(list, v)
		}
	}
	return list, tbl.Logger().LogIfError(require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, int64(len(list))))
}

func sliceDbUserTableFloat64List(ctx context.Context, tbl DbUserTabler, req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]float64, error) {
	q := tbl.Dialect().Quoter()
	whs, args := where.Where(wh, q)
	orderBy := where.Build(qc, q)
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", q.Quote(sqlname), quotedName, whs, orderBy)
	rows, err := support.Query(ctx, tbl, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]float64, 0, 10)

	for rows.Next() {
		var v float64
		err = rows.Scan(&v)
		if err == sql.ErrNoRows {
			return list, tbl.Logger().LogIfError(require.ErrorIfQueryNotSatisfiedBy(req, int64(len(list))))
		} else {
			list = append(list, v)
		}
	}
	return list, tbl.Logger().LogIfError(require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, int64(len(list))))
}

func constructDbUserTableInsert(tbl DbUserTable, w dialect.StringWriter, v *User, withPk bool) (s []interface{}, err error) {
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

func constructDbUserTableUpdate(tbl DbUserTable, w dialect.StringWriter, v *User) (s []interface{}, err error) {
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
func (tbl DbUserTable) Insert(ctx context.Context, req require.Requirement, vv ...*User) error {
	if req == require.All {
		req = require.Exactly(len(vv))
	}

	if ctx == nil {
		ctx = context.Background()
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

		fields, err := constructDbUserTableInsert(tbl, b, v, false)
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
			row := tbl.db.QueryRowContext(ctx, query, fields...)
			var i64 int64
			err = row.Scan(&i64)
			v.Uid = i64

		} else {
			i64, e2 := tbl.db.InsertContext(ctx, tbl.pk, query, fields...)
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
func (tbl DbUserTable) UpdateByUid(ctx context.Context, req require.Requirement, uid int64, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("uid", uid), fields...)
}

// UpdateByName updates one or more columns, given a name value.
func (tbl DbUserTable) UpdateByName(ctx context.Context, req require.Requirement, name string, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("name", name), fields...)
}

// UpdateByEmailaddress updates one or more columns, given a emailaddress value.
func (tbl DbUserTable) UpdateByEmailaddress(ctx context.Context, req require.Requirement, emailaddress string, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("emailaddress", emailaddress), fields...)
}

// UpdateByAddressid updates one or more columns, given a addressid value.
func (tbl DbUserTable) UpdateByAddressid(ctx context.Context, req require.Requirement, addressid int64, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("addressid", addressid), fields...)
}

// UpdateByAvatar updates one or more columns, given a avatar value.
func (tbl DbUserTable) UpdateByAvatar(ctx context.Context, req require.Requirement, avatar string, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("avatar", avatar), fields...)
}

// UpdateByRole updates one or more columns, given a role value.
func (tbl DbUserTable) UpdateByRole(ctx context.Context, req require.Requirement, role Role, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("role", role), fields...)
}

// UpdateByLastupdated updates one or more columns, given a lastupdated value.
func (tbl DbUserTable) UpdateByLastupdated(ctx context.Context, req require.Requirement, lastupdated int64, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("lastupdated", lastupdated), fields...)
}

// UpdateByI8 updates one or more columns, given a i8 value.
func (tbl DbUserTable) UpdateByI8(ctx context.Context, req require.Requirement, i8 int8, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("i8", i8), fields...)
}

// UpdateByU8 updates one or more columns, given a u8 value.
func (tbl DbUserTable) UpdateByU8(ctx context.Context, req require.Requirement, u8 uint8, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("u8", u8), fields...)
}

// UpdateByI16 updates one or more columns, given a i16 value.
func (tbl DbUserTable) UpdateByI16(ctx context.Context, req require.Requirement, i16 int16, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("i16", i16), fields...)
}

// UpdateByU16 updates one or more columns, given a u16 value.
func (tbl DbUserTable) UpdateByU16(ctx context.Context, req require.Requirement, u16 uint16, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("u16", u16), fields...)
}

// UpdateByI32 updates one or more columns, given a i32 value.
func (tbl DbUserTable) UpdateByI32(ctx context.Context, req require.Requirement, i32 int32, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("i32", i32), fields...)
}

// UpdateByU32 updates one or more columns, given a u32 value.
func (tbl DbUserTable) UpdateByU32(ctx context.Context, req require.Requirement, u32 uint32, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("u32", u32), fields...)
}

// UpdateByI64 updates one or more columns, given a i64 value.
func (tbl DbUserTable) UpdateByI64(ctx context.Context, req require.Requirement, i64 int64, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("i64", i64), fields...)
}

// UpdateByU64 updates one or more columns, given a u64 value.
func (tbl DbUserTable) UpdateByU64(ctx context.Context, req require.Requirement, u64 uint64, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("u64", u64), fields...)
}

// UpdateByF32 updates one or more columns, given a f32 value.
func (tbl DbUserTable) UpdateByF32(ctx context.Context, req require.Requirement, f32 float32, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("f32", f32), fields...)
}

// UpdateByF64 updates one or more columns, given a f64 value.
func (tbl DbUserTable) UpdateByF64(ctx context.Context, req require.Requirement, f64 float64, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("f64", f64), fields...)
}

// UpdateFields updates one or more columns, given a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed (but note that this is risky!).
func (tbl DbUserTable) UpdateFields(ctx context.Context, req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error) {
	return support.UpdateFields(ctx, tbl, req, wh, fields...)
}

//--------------------------------------------------------------------------------

// Update updates records, matching them by primary key. It returns the number of rows affected.
// The User.PreUpdate(Execer) method will be called, if it exists.
func (tbl DbUserTable) Update(ctx context.Context, req require.Requirement, vv ...*User) (int64, error) {
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

		args, err := constructDbUserTableUpdate(tbl, b, v)
		if err != nil {
			return count, err
		}
		args = append(args, v.Uid)

		b.WriteString(" WHERE ")
		q.QuoteW(b, tbl.pk)
		b.WriteString("=?")

		query := b.String()
		n, err := tbl.Exec(ctx, nil, query, args...)
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
func (tbl DbUserTable) Upsert(ctx context.Context, v *User, wh where.Expression) error {
	col := tbl.Dialect().Quoter().Quote(tbl.pk)
	qName := tbl.quotedName()
	whs, args := where.Where(wh, tbl.Dialect().Quoter())

	query := fmt.Sprintf("SELECT %s FROM %s %s", col, qName, whs)
	rows, err := support.Query(ctx, tbl, query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	if !rows.Next() {
		return tbl.Insert(ctx, require.One, v)
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
	_, err = tbl.Update(ctx, require.One, v)
	return err
}

//-------------------------------------------------------------------------------------------------

// DeleteByUid deletes rows from the table, given some uid values.
// The list of ids can be arbitrarily long.
func (tbl DbUserTable) DeleteByUid(ctx context.Context, req require.Requirement, uid ...int64) (int64, error) {
	ii := support.Int64AsInterfaceSlice(uid)
	return support.DeleteByColumn(ctx, tbl, req, "uid", ii...)
}

// DeleteByName deletes rows from the table, given some name values.
// The list of ids can be arbitrarily long.
func (tbl DbUserTable) DeleteByName(ctx context.Context, req require.Requirement, name ...string) (int64, error) {
	ii := support.StringAsInterfaceSlice(name)
	return support.DeleteByColumn(ctx, tbl, req, "name", ii...)
}

// DeleteByEmailaddress deletes rows from the table, given some emailaddress values.
// The list of ids can be arbitrarily long.
func (tbl DbUserTable) DeleteByEmailaddress(ctx context.Context, req require.Requirement, emailaddress ...string) (int64, error) {
	ii := support.StringAsInterfaceSlice(emailaddress)
	return support.DeleteByColumn(ctx, tbl, req, "emailaddress", ii...)
}

// DeleteByAddressid deletes rows from the table, given some addressid values.
// The list of ids can be arbitrarily long.
func (tbl DbUserTable) DeleteByAddressid(ctx context.Context, req require.Requirement, addressid ...int64) (int64, error) {
	ii := support.Int64AsInterfaceSlice(addressid)
	return support.DeleteByColumn(ctx, tbl, req, "addressid", ii...)
}

// DeleteByAvatar deletes rows from the table, given some avatar values.
// The list of ids can be arbitrarily long.
func (tbl DbUserTable) DeleteByAvatar(ctx context.Context, req require.Requirement, avatar ...string) (int64, error) {
	ii := support.StringAsInterfaceSlice(avatar)
	return support.DeleteByColumn(ctx, tbl, req, "avatar", ii...)
}

// DeleteByRole deletes rows from the table, given some role values.
// The list of ids can be arbitrarily long.
func (tbl DbUserTable) DeleteByRole(ctx context.Context, req require.Requirement, role ...Role) (int64, error) {
	ii := make([]interface{}, len(role))
	for i, v := range role {
		ii[i] = v
	}
	return support.DeleteByColumn(ctx, tbl, req, "role", ii...)
}

// DeleteByLastupdated deletes rows from the table, given some lastupdated values.
// The list of ids can be arbitrarily long.
func (tbl DbUserTable) DeleteByLastupdated(ctx context.Context, req require.Requirement, lastupdated ...int64) (int64, error) {
	ii := support.Int64AsInterfaceSlice(lastupdated)
	return support.DeleteByColumn(ctx, tbl, req, "lastupdated", ii...)
}

// DeleteByI8 deletes rows from the table, given some i8 values.
// The list of ids can be arbitrarily long.
func (tbl DbUserTable) DeleteByI8(ctx context.Context, req require.Requirement, i8 ...int8) (int64, error) {
	ii := support.Int8AsInterfaceSlice(i8)
	return support.DeleteByColumn(ctx, tbl, req, "i8", ii...)
}

// DeleteByU8 deletes rows from the table, given some u8 values.
// The list of ids can be arbitrarily long.
func (tbl DbUserTable) DeleteByU8(ctx context.Context, req require.Requirement, u8 ...uint8) (int64, error) {
	ii := support.Uint8AsInterfaceSlice(u8)
	return support.DeleteByColumn(ctx, tbl, req, "u8", ii...)
}

// DeleteByI16 deletes rows from the table, given some i16 values.
// The list of ids can be arbitrarily long.
func (tbl DbUserTable) DeleteByI16(ctx context.Context, req require.Requirement, i16 ...int16) (int64, error) {
	ii := support.Int16AsInterfaceSlice(i16)
	return support.DeleteByColumn(ctx, tbl, req, "i16", ii...)
}

// DeleteByU16 deletes rows from the table, given some u16 values.
// The list of ids can be arbitrarily long.
func (tbl DbUserTable) DeleteByU16(ctx context.Context, req require.Requirement, u16 ...uint16) (int64, error) {
	ii := support.Uint16AsInterfaceSlice(u16)
	return support.DeleteByColumn(ctx, tbl, req, "u16", ii...)
}

// DeleteByI32 deletes rows from the table, given some i32 values.
// The list of ids can be arbitrarily long.
func (tbl DbUserTable) DeleteByI32(ctx context.Context, req require.Requirement, i32 ...int32) (int64, error) {
	ii := support.Int32AsInterfaceSlice(i32)
	return support.DeleteByColumn(ctx, tbl, req, "i32", ii...)
}

// DeleteByU32 deletes rows from the table, given some u32 values.
// The list of ids can be arbitrarily long.
func (tbl DbUserTable) DeleteByU32(ctx context.Context, req require.Requirement, u32 ...uint32) (int64, error) {
	ii := support.Uint32AsInterfaceSlice(u32)
	return support.DeleteByColumn(ctx, tbl, req, "u32", ii...)
}

// DeleteByI64 deletes rows from the table, given some i64 values.
// The list of ids can be arbitrarily long.
func (tbl DbUserTable) DeleteByI64(ctx context.Context, req require.Requirement, i64 ...int64) (int64, error) {
	ii := support.Int64AsInterfaceSlice(i64)
	return support.DeleteByColumn(ctx, tbl, req, "i64", ii...)
}

// DeleteByU64 deletes rows from the table, given some u64 values.
// The list of ids can be arbitrarily long.
func (tbl DbUserTable) DeleteByU64(ctx context.Context, req require.Requirement, u64 ...uint64) (int64, error) {
	ii := support.Uint64AsInterfaceSlice(u64)
	return support.DeleteByColumn(ctx, tbl, req, "u64", ii...)
}

// DeleteByF32 deletes rows from the table, given some f32 values.
// The list of ids can be arbitrarily long.
func (tbl DbUserTable) DeleteByF32(ctx context.Context, req require.Requirement, f32 ...float32) (int64, error) {
	ii := make([]interface{}, len(f32))
	for i, v := range f32 {
		ii[i] = v
	}
	return support.DeleteByColumn(ctx, tbl, req, "f32", ii...)
}

// DeleteByF64 deletes rows from the table, given some f64 values.
// The list of ids can be arbitrarily long.
func (tbl DbUserTable) DeleteByF64(ctx context.Context, req require.Requirement, f64 ...float64) (int64, error) {
	ii := make([]interface{}, len(f64))
	for i, v := range f64 {
		ii[i] = v
	}
	return support.DeleteByColumn(ctx, tbl, req, "f64", ii...)
}

// Delete deletes one or more rows from the table, given a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl DbUserTable) Delete(ctx context.Context, req require.Requirement, wh where.Expression) (int64, error) {
	query, args := deleteRowsDbUserTableSql(tbl, wh)
	return tbl.Exec(ctx, req, query, args...)
}

func deleteRowsDbUserTableSql(tbl DbUserTabler, wh where.Expression) (string, []interface{}) {
	whs, args := where.Where(wh, tbl.Dialect().Quoter())
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("DELETE FROM %s %s", quotedName, whs)
	return query, args
}

//-------------------------------------------------------------------------------------------------

//-------------------------------------------------------------------------------------------------

// SetUid sets the Uid field and returns the modified User.
func (v *User) SetUid(x int64) *User {
	v.Uid = x
	return v
}

// SetName sets the Name field and returns the modified User.
func (v *User) SetName(x string) *User {
	v.Name = x
	return v
}

// SetEmailAddress sets the EmailAddress field and returns the modified User.
func (v *User) SetEmailAddress(x string) *User {
	v.EmailAddress = x
	return v
}

// SetAddressId sets the AddressId field and returns the modified User.
func (v *User) SetAddressId(x int64) *User {
	v.AddressId = &x
	return v
}

// SetAvatar sets the Avatar field and returns the modified User.
func (v *User) SetAvatar(x string) *User {
	v.Avatar = &x
	return v
}

// SetRole sets the Role field and returns the modified User.
func (v *User) SetRole(x Role) *User {
	v.Role = &x
	return v
}

// SetActive sets the Active field and returns the modified User.
func (v *User) SetActive(x bool) *User {
	v.Active = x
	return v
}

// SetAdmin sets the Admin field and returns the modified User.
func (v *User) SetAdmin(x bool) *User {
	v.Admin = x
	return v
}

// SetFave sets the Fave field and returns the modified User.
func (v *User) SetFave(x big.Int) *User {
	v.Fave = &x
	return v
}

// SetLastUpdated sets the LastUpdated field and returns the modified User.
func (v *User) SetLastUpdated(x int64) *User {
	v.LastUpdated = x
	return v
}

// SetI8 sets the I8 field and returns the modified User.
func (v *User) SetI8(x int8) *User {
	v.Numbers.I8 = x
	return v
}

// SetU8 sets the U8 field and returns the modified User.
func (v *User) SetU8(x uint8) *User {
	v.Numbers.U8 = x
	return v
}

// SetI16 sets the I16 field and returns the modified User.
func (v *User) SetI16(x int16) *User {
	v.Numbers.I16 = x
	return v
}

// SetU16 sets the U16 field and returns the modified User.
func (v *User) SetU16(x uint16) *User {
	v.Numbers.U16 = x
	return v
}

// SetI32 sets the I32 field and returns the modified User.
func (v *User) SetI32(x int32) *User {
	v.Numbers.I32 = x
	return v
}

// SetU32 sets the U32 field and returns the modified User.
func (v *User) SetU32(x uint32) *User {
	v.Numbers.U32 = x
	return v
}

// SetI64 sets the I64 field and returns the modified User.
func (v *User) SetI64(x int64) *User {
	v.Numbers.I64 = x
	return v
}

// SetU64 sets the U64 field and returns the modified User.
func (v *User) SetU64(x uint64) *User {
	v.Numbers.U64 = x
	return v
}

// SetF32 sets the F32 field and returns the modified User.
func (v *User) SetF32(x float32) *User {
	v.Numbers.F32 = x
	return v
}

// SetF64 sets the F64 field and returns the modified User.
func (v *User) SetF64(x float64) *User {
	v.Numbers.F64 = x
	return v
}

// Settoken sets the token field and returns the modified User.
func (v *User) Settoken(x string) *User {
	v.token = x
	return v
}

// Setsecret sets the secret field and returns the modified User.
func (v *User) Setsecret(x string) *User {
	v.secret = x
	return v
}
