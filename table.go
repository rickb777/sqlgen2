package sqlgen2

import (
	"database/sql"
	"github.com/rickb777/sqlgen2/require"
	"github.com/rickb777/sqlgen2/schema"
	"github.com/rickb777/sqlgen2/where"
	"log"
	"context"
)

// Table provides the generic features of each generated table handler.
type Table interface {
	// Name gets the table name. without prefix
	Name() TableName

	// Database gets the shared database information.
	Database() *Database

	// Execer gets the wrapped database or transaction handle.
	Execer() Execer

	// DB gets the wrapped database handle, provided this is not within a transaction.
	// Panics if it is in the wrong state - use IsTx() if necessary.
	DB() *sql.DB

	// Tx gets the wrapped transaction handle, provided this is within a transaction.
	// Panics if it is in the wrong state - use IsTx() if necessary.
	Tx() *sql.Tx

	// IsTx tests whether this is within a transaction.
	IsTx() bool

	// Constraints returns the table's constraints.
	// (not included here because of package inter-dependencies)
	//Constraints() constraint.Constraints

	// Ctx gets the current request context.
	Ctx() context.Context

	// Dialect gets the database dialect.
	Dialect() schema.Dialect

	// Logger gets the trace logger.
	Logger() *log.Logger

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

	// WithConstraint returns a modified Table with added data consistency constraints.
	//WithConstraint(cc ...sqlgen2.Constraint) SomeTypeTable {
	//---------------------------------------------------------------------------------------------

	// Query is the low-level request method for this table. The query is logged using whatever logger is
	// configured. If an error arises, this too is logged.
	//
	// If you need a context other than the background, use WithContext before calling Query.
	//
	// The args are for any placeholder parameters in the query.
	//
	// The caller must call rows.Close() on the result.
	Query(query string, args ...interface{}) (*sql.Rows, error)
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
	QueryOneNullString(req require.Requirement, query string, args ...interface{}) (result sql.NullString, err error)

	// QueryOneNullInt64 is a low-level access method for one int64. This can be used for 'COUNT(1)' queries and
	// such like. If the query selected many rows, only the first is returned; the rest are discarded.
	// If not found, the result will be invalid.
	QueryOneNullInt64(req require.Requirement, query string, args ...interface{}) (result sql.NullInt64, err error)

	// QueryOneNullFloat64 is a low-level access method for one float64. This can be used for 'AVG(...)' queries and
	// such like. If the query selected many rows, only the first is returned; the rest are discarded.
	// If not found, the result will be invalid.
	QueryOneNullFloat64(req require.Requirement, query string, args ...interface{}) (result sql.NullFloat64, err error)

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
	//GetSomeType(id int64) (*SomeType, error)
	//MustGetSomeType(id int64) (*SomeType, error)
	//
	//GetSomeTypes(req require.Requirement, id ...int64) (list []*SomeType, err error)
	//
	//SelectOneWhere(req require.Requirement, where, orderBy string, args ...interface{}) (*SomeType, error)
	//SelectOne(req require.Requirement, where where.Expression, orderBy string) (*SomeType, error)
	//SelectWhere(req require.Requirement, where, orderBy string, args ...interface{}) ([]*SomeType, error)
	//Select(req require.Requirement, where where.Expression, orderBy string) ([]*SomeType, error)
	//
	//Insert(req require.Requirement, vv ...*SomeType) error
	//
	//Update(req require.Requirement, vv ...*SomeType) (int64, error)
}
