package sqlgen2

import (
	"database/sql"
	"context"
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
