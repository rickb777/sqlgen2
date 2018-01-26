package sqlgen2

import (
	"github.com/rickb777/sqlgen2/schema"
	"github.com/rickb777/sqlgen2/util"
	"log"
	"context"
	"database/sql"
)

type Database struct {
	db      Execer
	dialect schema.Dialect
	ctx     context.Context
	logger  *log.Logger
	wrapper interface{}
}

// NewDatabase createa a new database handler, which wraps the core *sql.DB.
func NewDatabase(db Execer, dialect schema.Dialect) *Database {
	return &Database{db, dialect, context.Background(), nil, nil}
}

func (database *Database) DB() Execer {
	return database.db
}

// Wrapper gets whatever structure is present, as needed.
func (database *Database) Dialect() schema.Dialect {
	return database.dialect
}

// SetContext sets the context for subsequent queries.
func (database *Database) SetContext(ctx context.Context) *Database {
	database.ctx = ctx
	return database
}

// Logger gets the trace logger.
func (database *Database) Logger() *log.Logger {
	return database.logger
}

// SetLogger sets the logger for subsequent queries, returning the interface.
func (database *Database) SetLogger(logger *log.Logger) *Database {
	database.logger = logger
	return database
}

// Wrapper gets whatever structure is present, as needed.
func (database *Database) Wrapper() interface{} {
	return database.wrapper
}

// SetWrapper sets a user-defined wrapper or container.
func (database *Database) SetWrapper(wrapper interface{}) *Database {
	database.wrapper = wrapper
	return database
}

//-------------------------------------------------------------------------------------------------

func (database *Database) Exec(query string, args ...interface{}) (sql.Result, error) {
	return database.db.ExecContext(database.ctx, query, args...)
}

func (database *Database) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return database.db.ExecContext(ctx, query, args...)
}

func (database *Database) Prepare(query string) (*sql.Stmt, error) {
	return database.db.PrepareContext(database.ctx, query)
}

func (database *Database) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return database.db.PrepareContext(ctx, query)
}

func (database *Database) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return database.db.QueryContext(database.ctx, query, args...)
}

func (database *Database) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return database.db.QueryContext(ctx, query, args...)
}

func (database *Database) QueryRow(query string, args ...interface{}) *sql.Row {
	return database.db.QueryRowContext(database.ctx, query, args...)
}

func (database *Database) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return database.db.QueryRowContext(ctx, query, args...)
}

//-------------------------------------------------------------------------------------------------


// DoesTableExist gets all the table names in the database/schema.
func (database *Database) TableExists(name TableName) (yes bool, err error) {
	wanted := name.String()
	rows, err := database.db.QueryContext(database.ctx, showTables(database.dialect))
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
func (database *Database) ListTables(dialect schema.Dialect) (util.StringList, error) {
	ss := make(util.StringList, 0)
	rows, err := database.db.QueryContext(database.ctx, showTables(dialect))
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
