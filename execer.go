package sqlgen2

import (
	"database/sql"
	"context"
)

// Execer describes the methods of the core database API. See database/sql/DB and database/sql/Tx.
type Execer interface {
	// ExecContext executes a query without returning any rows.
	// The args are for any placeholder parameters in the query.
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)

	// PrepareContext creates a prepared statement for later queries or executions.
	// Multiple queries or executions may be run concurrently from the
	// returned statement.
	// The caller must call the statement's Close method
	// when the statement is no longer needed.
	//
	// The provided context is used for the preparation of the statement, not for the
	// execution of the statement.
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)

	// QueryContext executes a query that returns rows, typically a SELECT.
	// The args are for any placeholder parameters in the query.
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)

	// QueryRowContext executes a query that is expected to return at most one row.
	// QueryRowContext always returns a non-nil value. Errors are deferred until
	// Row's Scan method is called.
	// If the query selects no rows, the *Row's Scan will return ErrNoRows.
	// Otherwise, the *Row's Scan scans the first selected row and discards
	// the rest.
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

// TxStarter is able to begin transactions. See database/sql/DB.
type TxStarter interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
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
