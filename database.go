package sqlgen2

import (
	"github.com/rickb777/sqlgen2/schema"
	"github.com/rickb777/sqlgen2/util"
	"log"
	"context"
	"database/sql"
	"strings"
)

// Database wraps a *sql.DB with a dialect and (optionally) a logger.
// It's safe for concurrent use by multiple goroutines.
type Database struct {
	db      Execer
	dialect schema.Dialect
	logger  *log.Logger
	wrapper interface{}
}

// NewDatabase creates a new database handler, which wraps the core *sql.DB along with
// the appropriate dialect.
//
// You can supply the logger you need, or else nil. All queries will be logged; all database
// will be logged.
//
// The wrapper holds some associated data your application needs for this database, if any.
// Otherwise this should be nil.
func NewDatabase(db Execer, dialect schema.Dialect, logger *log.Logger, wrapper interface{}) *Database {
	return &Database{db, dialect, logger, wrapper}
}

// DB gets the Execer, which is a *sql.DB (except during testing using mocks).
func (database *Database) DB() Execer {
	return database.db
}

// BeginTx starts a transaction.
//
// The context is used until the transaction is committed or rolled back. If this
// context is cancelled, the sql package will roll back the transaction. In this
// case, Tx.Commit will then return an error.
//
// The provided TxOptions is optional and may be nil if defaults should be used.
// If a non-default isolation level is used that the driver doesn't support,
// an error will be returned.
//
// Panics if the Execer is not a TxStarter.
func (database *Database) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return database.db.(TxStarter).BeginTx(ctx, opts)
}

// Begin starts a transaction using default options. The default isolation level is
// dependent on the driver.
func (database *Database) Begin() (*sql.Tx, error) {
	return database.BeginTx(context.Background(), nil)
}

// Wrapper gets whatever structure is present, as needed.
func (database *Database) Dialect() schema.Dialect {
	return database.dialect
}

// Logger gets the trace logger.
func (database *Database) Logger() *log.Logger {
	return database.logger
}

// Wrapper gets whatever structure is present, as needed.
func (database *Database) Wrapper() interface{} {
	return database.wrapper
}

//-------------------------------------------------------------------------------------------------

// PingContext verifies a connection to the database is still alive,
// establishing a connection if necessary.
func (database *Database) PingContext(ctx context.Context) error {
	return database.db.(*sql.DB).PingContext(ctx)
}

// Ping verifies a connection to the database is still alive,
// establishing a connection if necessary.
func (database *Database) Ping() error {
	return database.PingContext(context.Background())
}

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
func (database *Database) Exec(query string, args ...interface{}) (sql.Result, error) {
	return database.ExecContext(context.Background(), query, args...)
}

// ExecContext executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
func (database *Database) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return database.db.ExecContext(ctx, query, args...)
}

// Prepare creates a prepared statement for later queries or executions.
// Multiple queries or executions may be run concurrently from the
// returned statement.
// The caller must call the statement's Close method
// when the statement is no longer needed.
func (database *Database) Prepare(query string) (*sql.Stmt, error) {
	return database.PrepareContext(context.Background(), query)
}

// PrepareContext creates a prepared statement for later queries or executions.
// Multiple queries or executions may be run concurrently from the
// returned statement.
// The caller must call the statement's Close method
// when the statement is no longer needed.
//
// The provided context is used for the preparation of the statement, not for the
// execution of the statement.
func (database *Database) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return database.db.PrepareContext(ctx, query)
}

// Query executes a query that returns rows, typically a SELECT.
// The args are for any placeholder parameters in the query.
func (database *Database) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return database.QueryContext(context.Background(), query, args...)
}

// QueryContext executes a query that returns rows, typically a SELECT.
// The args are for any placeholder parameters in the query.
func (database *Database) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return database.db.QueryContext(ctx, query, args...)
}

// QueryRow executes a query that is expected to return at most one row.
// QueryRow always returns a non-nil value. Errors are deferred until
// Row's Scan method is called.
// If the query selects no rows, the *Row's Scan will return ErrNoRows.
// Otherwise, the *Row's Scan scans the first selected row and discards
// the rest.
func (database *Database) QueryRow(query string, args ...interface{}) *sql.Row {
	return database.QueryRowContext(context.Background(), query, args...)
}

// QueryRowContext executes a query that is expected to return at most one row.
// QueryRowContext always returns a non-nil value. Errors are deferred until
// Row's Scan method is called.
// If the query selects no rows, the *Row's Scan will return ErrNoRows.
// Otherwise, the *Row's Scan scans the first selected row and discards
// the rest.
func (database *Database) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return database.db.QueryRowContext(ctx, query, args...)
}

// Stats returns database statistics.
func (database *Database) Stats() sql.DBStats {
	return database.db.(*sql.DB).Stats()
}

//-------------------------------------------------------------------------------------------------

// LogQuery writes query info to the logger, if it is not nil.
func (database *Database) LogQuery(query string, args ...interface{}) {
	//support.LogQuery(database.logger, query, args)
	if database.logger != nil {
		query = strings.TrimSpace(query)
		if len(args) > 0 {
			database.logger.Printf(query+" %v\n", args)
		} else {
			database.logger.Println(query)
		}
	}
}

// LogIfError writes error info to the logger, if both the logger and the error are non-nil.
// It returns the error.
func (database *Database) LogIfError(err error) error {
	if database.logger != nil && err != nil {
		database.logger.Printf("Error: %s\n", err)
	}
	return err
}

// LogError writes error info to the logger, if the logger is not nil. It returns the error.
func (database *Database) LogError(err error) error {
	if database.logger != nil {
		database.logger.Printf("Error: %s\n", err)
	}
	return err
}

//-------------------------------------------------------------------------------------------------

// DoesTableExist gets all the table names in the database/schema.
func (database *Database) TableExists(name TableName) (yes bool, err error) {
	wanted := name.String()
	rows, err := database.db.QueryContext(context.Background(), showTables(database.dialect))
	if err != nil {
		return false, err
	}
	defer rows.Close()

	for rows.Next() {
		var s string
		rows.Scan(&s)
		if s == wanted {
			return true, rows.Err()
		}
	}
	return false, rows.Err()
}

// ListTables gets all the table names in the database/schema.
func (database *Database) ListTables() (util.StringList, error) {
	ss := make(util.StringList, 0)
	rows, err := database.db.QueryContext(context.Background(), showTables(database.dialect))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var s string
		rows.Scan(&s)
		ss = append(ss, s)
	}
	return ss, rows.Err()
}

func showTables(dialect schema.Dialect) string {
	switch dialect.Index() {
	case schema.SqliteIndex:
		return `SELECT name FROM sqlite_master WHERE type = "table"`
	case schema.MysqlIndex:
		return `SHOW TABLES`
	case schema.PostgresIndex:
		return `SELECT tablename FROM pg_catalog.pg_tables`
	}
	panic(dialect.String())
}
