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

type CanPreInsert interface {
	PreInsert(Execer) error
}

//-------------------------------------------------------------------------------------------------

// Table provides the generic features of each generated table handler.
type Table interface {
	// FullName gets the concatenated prefix and table name.
	FullName() string

	// DB gets the wrapped database handle, provided this is not within a transaction.
	// Panics if it is in the wrong state - use IsTx() if necessary.
	DB() *sql.DB

	// Tx gets the wrapped transaction handle, provided this is within a transaction.
	// Panics if it is in the wrong state - use IsTx() if necessary.
	Tx() *sql.Tx

	// IsTx tests whether this is within a transaction.
	IsTx() bool
}

type TableWithDDL interface {
	Table
	CreateTable(ifNotExist bool) (int64, error)
	CreateIndexes(ifNotExist bool) (err error)
	CreateTableWithIndexes(ifNotExist bool) (err error)
}

type TableWithCrud interface {
	Table
	Exec(query string, args ...interface{}) (int64, error)
	//QueryOne(query string, args ...interface{}) (*User, error)
	//Query(query string, args ...interface{}) ([]*User, error)
	//SelectOneSA(where, orderBy string, args ...interface{}) (*User, error)
	//SelectOne(where where.Expression, orderBy string) (*User, error)
	//SelectSA(where, orderBy string, args ...interface{}) ([]*User, error)
	//Select(where where.Expression, orderBy string) ([]*User, error)
	CountSA(where string, args ...interface{}) (count int64, err error)
	//Count(where where.Expression) (count int64, err error)
	//Insert(vv ...*User) error
	//UpdateFields(where where.Expression, fields ...sql.NamedArg) (int64, error)
	//Update(vv ...*User) (int64, error)
	//Delete(where where.Expression) (int64, error)
}