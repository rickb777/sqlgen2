package sqlgen2

import (
	"database/sql"
	"context"
	"github.com/rickb777/sqlgen2/where"
	"log"
	"github.com/rickb777/sqlgen2/schema"
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
		return tn.Prefix[0:last]
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

	// Exec executes a query.
	// It returns the number of rows affected (if the DB supports that).
	Exec(query string, args ...interface{}) (int64, error)

	// CountSA counts records that match a 'where' predicate.
	CountSA(where string, args ...interface{}) (count int64, err error)

	// Count counts records that match a 'where' predicate.
	Count(where where.Expression) (count int64, err error)

	// UpdateFields writes new values to the specified columns for rows that match the 'where' predicate.
	// It returns the number of rows affected (if the DB supports that).
	UpdateFields(where where.Expression, fields ...sql.NamedArg) (int64, error)

	// Delete deletes rows that match the 'where' predicate.
	// It returns the number of rows affected (if the DB supports that).
	Delete(where where.Expression) (int64, error)

	// These methods are provided but have a specific record type so do not conform to a general interface.
	//QueryOne(query string, args ...interface{}) (*User, error)
	//Query(query string, args ...interface{}) ([]*User, error)
	//SelectOneSA(where, orderBy string, args ...interface{}) (*User, error)
	//SelectOne(where where.Expression, orderBy string) (*User, error)
	//SelectSA(where, orderBy string, args ...interface{}) ([]*User, error)
	//Select(where where.Expression, orderBy string) ([]*User, error)
	//Insert(vv ...*User) error
	//Update(vv ...*User) (int64, error)
}
