package sqlgen2

import (
	"database/sql"
	"context"
	"log"
	"github.com/rickb777/sqlgen2/schema"
	"github.com/rickb777/sqlgen2/where"
	"github.com/rickb777/sqlgen2/require"
)

// Execer describes the methods of the core database API. See database/sql/DB and database/sql/Tx.
type Execer interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)

	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)

	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)

	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

// Type conformance assertions
var _ Execer = &sql.DB{}
var _ Execer = &sql.Tx{}

//-------------------------------------------------------------------------------------------------
// Hooks
// These allow values to be adjusted prior to insertion / updating or after fetching.

// CanPreInsert is implemented by value types that need a hook to run just before their data
// is inserted into the database.
type CanPreInsert interface {
	PreInsert() error
}

// CanPreUpdate is implemented by value types that need a hook to run just before their data
// is updated in the database.
type CanPreUpdate interface {
	PreUpdate() error
}

// CanPostGet is implemented by value types that need a hook to run just after their data
// is fetched from the database.
type CanPostGet interface {
	PostGet() error
}

//-------------------------------------------------------------------------------------------------

// TableName holds a two-part name. The prefix part is optional.
type TableName struct {
	// Prefix on the table name. It can be used as the schema name, in which case
	// it should include the trailing dot. Or it can be any prefix as needed.
	Prefix string

	// The principal name of the table.
	Name string
}

// String gets the full table name.
func (tn TableName) String() string {
	return tn.Prefix + tn.Name
}

// PrefixWithoutDot return the prefix; if this ends with a dot, the dot is removed.
func (tn TableName) PrefixWithoutDot() string {
	last := len(tn.Prefix)-1
	if last > 0 && tn.Prefix[last] == '.' {
		return tn.Prefix[:last]
	}
	return tn.Prefix
}

//-------------------------------------------------------------------------------------------------

// Table provides the generic features of each generated table handler.
type Table interface {
	// Name gets the table name. without prefix
	Name() TableName

	// DB gets the wrapped database handle, provided this is not within a transaction.
	// Panics if it is in the wrong state - use IsTx() if necessary.
	DB() *sql.DB

	// Tx gets the wrapped transaction handle, provided this is within a transaction.
	// Panics if it is in the wrong state - use IsTx() if necessary.
	Tx() *sql.Tx

	// IsTx tests whether this is within a transaction.
	IsTx() bool

	// Ctx gets the current request context.
	Ctx() context.Context

	// Dialect gets the database dialect.
	Dialect() schema.Dialect

	// Logger gets the trace logger.
	Logger() *log.Logger

	// SetLogger sets the trace logger.
	SetLogger(logger *log.Logger) Table

	// Wrapper gets whatever structure is present, as needed.
	Wrapper() interface{}

	// SetWrapper sets a user-defined wrapper or container.
	SetWrapper(wrapper interface{}) Table

	//---------------------------------------------------------------------------------------------
	// The following type-specific methods are also provided (but are not part of this interface).

	// WithPrefix sets the table name prefix for subsequent queries.
	// The result is a modified copy of the table; the original is unchanged.
	// WithPrefix(pfx string) SomeTypeTable

	// WithContext sets the context for subsequent queries.
	// The result is a modified copy of the table; the original is unchanged.
	//WithContext(ctx context.Context) SomeTypeTable {

	// WithLogger sets the logger for subsequent queries. An alias for SetLogger.
	// The result is a modified copy of the table; the original is unchanged.
	//WithLogger(logger *log.Logger) SomeTypeTable {

	// Begin starts a transaction. The default isolation level is dependent on the driver.
	// The result is a modified copy of the table; the original is unchanged.
	//BeginTx(opts *sql.TxOptions) (SomeTypeTable, error)

	// Using returns a modified Table using the transaction supplied. This is needed
	// when making multiple queries across several tables within a single transaction.
	// The result is a modified copy of the table; the original is unchanged.
	//Using(tx *sql.Tx) SomeTypeTable
}

type TableCreator interface {
	Table

	// CreateTable creates the database table.
	CreateTable(ifNotExists bool) (int64, error)

	// DropTable drops the database table.
	DropTable(ifExists bool) (int64, error)

	// Truncate empties the table
	Truncate(force bool) (err error)
}

type TableWithIndexes interface {
	TableCreator

	// CreateIndexes creates the indexes for the database table.
	CreateIndexes(ifNotExist bool) (err error)

	// DropIndexes executes a query that drops all the indexes on the database table.
	DropIndexes(ifExist bool) (err error)

	// CreateTableWithIndexes creates the database table and its indexes.
	CreateTableWithIndexes(ifNotExist bool) (err error)
}

type TableWithCrud interface {
	Table

	// QueryOneNullString is a low-level access method for one string. This can be used for function queries and
	// such like. If the query selected many rows, only the first is returned; the rest are discarded.
	// If not found, the result will be invalid.
	QueryOneNullString(query string, args ...interface{}) (result sql.NullString, err error)

	// MustQueryOneNullString is a low-level access method for one string. This can be used for function queries and
	// such like.
	//
	// It places a requirement that exactly one result must be found; an error is generated when this expectation is not met.
	MustQueryOneNullString(query string, args ...interface{}) (result sql.NullString, err error)

	// QueryOneNullInt64 is a low-level access method for one int64. This can be used for 'COUNT(1)' queries and
	// such like. If the query selected many rows, only the first is returned; the rest are discarded.
	// If not found, the result will be invalid.
	QueryOneNullInt64(query string, args ...interface{}) (result sql.NullInt64, err error)

	// MustQueryOneNullInt64 is a low-level access method for one int64. This can be used for 'COUNT(1)' queries and
	// such like.
	//
	// It places a requirement that exactly one result must be found; an error is generated when this expectation is not met.
	MustQueryOneNullInt64(query string, args ...interface{}) (result sql.NullInt64, err error)

	// QueryOneNullFloat64 is a low-level access method for one float64. This can be used for 'AVG(...)' queries and
	// such like. If the query selected many rows, only the first is returned; the rest are discarded.
	// If not found, the result will be invalid.
	QueryOneNullFloat64(query string, args ...interface{}) (result sql.NullFloat64, err error)

	// MustQueryOneNullFloat64 is a low-level access method for one float64. This can be used for 'AVG(...)' queries and
	// such like.
	//
	// It places a requirement that exactly one result must be found; an error is generated when this expectation is not met.
	MustQueryOneNullFloat64(query string, args ...interface{}) (result sql.NullFloat64, err error)

	// ReplaceTableName replaces all occurrences of "{TABLE}" with the table's name.
	ReplaceTableName(query string) string

	// Exec executes a query.
	//
	// It places a requirement, which may be nil, on the number of affected rows: this
	// controls whether an error is generated when this expectation is not met.
	//
	// It returns the number of rows affected (if the DB supports that).
	Exec(req require.Requirement, query string, args ...interface{}) (int64, error)

	// CountWhere counts records that match a 'where' predicate.
	CountWhere(where string, args ...interface{}) (count int64, err error)

	// Count counts records that match a 'where' predicate.
	Count(where where.Expression) (count int64, err error)

	// UpdateFields writes new values to the specified columns for rows that match the 'where' predicate.
	// It returns the number of rows affected (if the DB supports that).
	UpdateFields(req require.Requirement, where where.Expression, fields ...sql.NamedArg) (int64, error)

	// Delete deletes rows that match the 'where' predicate.
	// It returns the number of rows affected (if the DB supports that).
	Delete(req require.Requirement, wh where.Expression) (int64, error)

	//---------------------------------------------------------------------------------------------
	// The following type-specific methods are also provided (but are not part of this interface).
	//
	//QueryOne(query string, args ...interface{}) (*User, error)
	//MustQueryOne(query string, args ...interface{}) (*User, error)
	//
	//Query(req require.Requirement, query string, args ...interface{}) ([]*User, error)
	//
	//GetUser(id int64) (*User, error)
	//MustGetUser(id int64) (*User, error)
	//
	//GetUsers(req require.Requirement, id ...int64) (list []*User, err error)
	//
	//SelectOneWhere(req require.Requirement, where, orderBy string, args ...interface{}) (*User, error)
	//SelectOne(req require.Requirement, where where.Expression, orderBy string) (*User, error)
	//SelectWhere(req require.Requirement, where, orderBy string, args ...interface{}) ([]*User, error)
	//Select(req require.Requirement, where where.Expression, orderBy string) ([]*User, error)
	//
	//Insert(req require.Requirement, vv ...*User) error
	//
	//Update(req require.Requirement, vv ...*User) (int64, error)
}
