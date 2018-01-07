// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

package demo

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/rickb777/sqlgen2"
	"github.com/rickb777/sqlgen2/schema"
	"log"
)

// V4UserTableName is the default name for this table.
const V4UserTableName = "users"

// V4UserTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type V4UserTable struct {
	Prefix, Name string
	Db           sqlgen2.Execer
	Ctx          context.Context
	Dialect      schema.Dialect
	Logger       *log.Logger
}

// Type conformance check
var _ sqlgen2.Table = &V4UserTable{}

// NewV4UserTable returns a new table instance.
// If a blank table name is supplied, the default name "users" will be used instead.
// The table name prefix is initially blank and the request context is the background.
func NewV4UserTable(name string, d sqlgen2.Execer, dialect schema.Dialect) V4UserTable {
	if name == "" {
		name = V4UserTableName
	}
	return V4UserTable{"", name, d, context.Background(), dialect, nil}
}

// WithPrefix sets the prefix for subsequent queries.
func (tbl V4UserTable) WithPrefix(pfx string) V4UserTable {
	tbl.Prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries.
func (tbl V4UserTable) WithContext(ctx context.Context) V4UserTable {
	tbl.Ctx = ctx
	return tbl
}

// WithLogger sets the logger for subsequent queries.
func (tbl V4UserTable) WithLogger(logger *log.Logger) V4UserTable {
	tbl.Logger = logger
	return tbl
}

// SetLogger sets the logger for subsequent queries, returning the interface.
func (tbl V4UserTable) SetLogger(logger *log.Logger) sqlgen2.Table {
	tbl.Logger = logger
	return tbl
}

// FullName gets the concatenated prefix and table name.
func (tbl V4UserTable) FullName() string {
	return tbl.Prefix + tbl.Name
}

func (tbl V4UserTable) prefixWithoutDot() string {
	last := len(tbl.Prefix)-1
	if last > 0 && tbl.Prefix[last] == '.' {
		return tbl.Prefix[0:last]
	}
	return tbl.Prefix
}

// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl V4UserTable) DB() *sql.DB {
	return tbl.Db.(*sql.DB)
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl V4UserTable) Tx() *sql.Tx {
	return tbl.Db.(*sql.Tx)
}

// IsTx tests whether this is within a transaction.
func (tbl V4UserTable) IsTx() bool {
	_, ok := tbl.Db.(*sql.Tx)
	return ok
}

// Begin starts a transaction. The default isolation level is dependent on the driver.
func (tbl V4UserTable) BeginTx(opts *sql.TxOptions) (V4UserTable, error) {
	d := tbl.Db.(*sql.DB)
	var err error
	tbl.Db, err = d.BeginTx(tbl.Ctx, opts)
	return tbl, err
}

func (tbl V4UserTable) logQuery(query string, args ...interface{}) {
	sqlgen2.LogQuery(tbl.Logger, query, args...)
}


//--------------------------------------------------------------------------------

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// It returns the number of rows affected (of the database drive supports this).
func (tbl V4UserTable) Exec(query string, args ...interface{}) (int64, error) {
	tbl.logQuery(query, args...)
	res, err := tbl.Db.ExecContext(tbl.Ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

//--------------------------------------------------------------------------------

// QueryOne is the low-level access function for one User.
// If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, *User will be nil.
func (tbl V4UserTable) QueryOne(query string, args ...interface{}) (*User, error) {
	list, err := tbl.doQuery(true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

// Query is the low-level access function for Users.
func (tbl V4UserTable) Query(query string, args ...interface{}) ([]*User, error) {
	return tbl.doQuery(false, query, args...)
}

func (tbl V4UserTable) doQuery(firstOnly bool, query string, args ...interface{}) ([]*User, error) {
	tbl.logQuery(query, args...)
	rows, err := tbl.Db.QueryContext(tbl.Ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanV4Users(rows, firstOnly)
}

// scanV4Users reads table records into a slice of values.
func scanV4Users(rows *sql.Rows, firstOnly bool) ([]*User, error) {
	var err error
	var vv []*User

	var v0 int64
	var v1 string
	var v2 string
	var v3 string
	var v4 bool
	var v5 bool
	var v6 []byte
	var v7 int64
	var v8 string
	var v9 string

	for rows.Next() {
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
		)
		if err != nil {
			return vv, err
		}

		v := &User{}
		v.Uid = v0
		v.Login = v1
		v.EmailAddress = v2
		v.Avatar = v3
		v.Active = v4
		v.Admin = v5
		err = json.Unmarshal(v6, &v.Fave)
		if err != nil {
			return nil, err
		}
		v.LastUpdated = v7
		v.token = v8
		v.secret = v9

		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPostGet); ok {
			err = hook.PostGet()
			if err != nil {
				return vv, err
			}
		}

		vv = append(vv, v)

		if firstOnly {
			return vv, rows.Err()
		}
	}

	return vv, rows.Err()
}

func sliceV4UserWithoutPk(v *User) ([]interface{}, error) {

	v6, err := json.Marshal(&v.Fave)
	if err != nil {
		return nil, err
	}

	return []interface{}{
		v.Login,
		v.EmailAddress,
		v.Avatar,
		v.Active,
		v.Admin,
		v6,
		v.LastUpdated,
		v.token,
		v.secret,

	}, nil
}
