// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

package demo

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/rickb777/sqlgen2"
	"github.com/rickb777/sqlgen2/schema"
	"github.com/rickb777/sqlgen2/where"
	"log"
)

// SUserTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type SUserTable struct {
	prefix, name string
	db           sqlgen2.Execer
	ctx          context.Context
	dialect      schema.Dialect
	logger       *log.Logger
}

// Type conformance check
var _ sqlgen2.Table = &SUserTable{}

// NewSUserTable returns a new table instance.
// If a blank table name is supplied, the default name "users" will be used instead.
// The table name prefix is initially blank and the request context is the background.
func NewSUserTable(name string, d sqlgen2.Execer, dialect schema.Dialect) SUserTable {
	if name == "" {
		name = "users"
	}
	return SUserTable{"", name, d, context.Background(), dialect, nil}
}

// CopyTableAsSUserTable copies a table instance, retaining the name etc but
// providing methods appropriate for 'User'.
func CopyTableAsSUserTable(origin sqlgen2.Table) SUserTable {
	return SUserTable{
		prefix:  origin.Prefix(),
		name:    origin.Name(),
		db:      origin.DB(),
		ctx:     origin.Ctx(),
		dialect: origin.Dialect(),
		logger:  origin.Logger(),
	}
}

// WithPrefix sets the table name prefix for subsequent queries.
func (tbl SUserTable) WithPrefix(pfx string) SUserTable {
	tbl.prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries.
func (tbl SUserTable) WithContext(ctx context.Context) SUserTable {
	tbl.ctx = ctx
	return tbl
}

// WithLogger sets the logger for subsequent queries.
func (tbl SUserTable) WithLogger(logger *log.Logger) SUserTable {
	tbl.logger = logger
	return tbl
}

// Ctx gets the current request context.
func (tbl SUserTable) Ctx() context.Context {
	return tbl.ctx
}

// Dialect gets the database dialect.
func (tbl SUserTable) Dialect() schema.Dialect {
	return tbl.dialect
}

// Logger gets the trace logger.
func (tbl SUserTable) Logger() *log.Logger {
	return tbl.logger
}

// SetLogger sets the logger for subsequent queries, returning the interface.
func (tbl SUserTable) SetLogger(logger *log.Logger) sqlgen2.Table {
	tbl.logger = logger
	return tbl
}

// Name gets the table name.
func (tbl SUserTable) Name() string {
	return tbl.name
}

// Prefix gets the table name prefix.
func (tbl SUserTable) Prefix() string {
	return tbl.prefix
}

// FullName gets the concatenated prefix and table name.
func (tbl SUserTable) FullName() string {
	return tbl.prefix + tbl.name
}

func (tbl SUserTable) prefixWithoutDot() string {
	last := len(tbl.prefix)-1
	if last > 0 && tbl.prefix[last] == '.' {
		return tbl.prefix[0:last]
	}
	return tbl.prefix
}

// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl SUserTable) DB() *sql.DB {
	return tbl.db.(*sql.DB)
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl SUserTable) Tx() *sql.Tx {
	return tbl.db.(*sql.Tx)
}

// IsTx tests whether this is within a transaction.
func (tbl SUserTable) IsTx() bool {
	_, ok := tbl.db.(*sql.Tx)
	return ok
}

// Begin starts a transaction. The default isolation level is dependent on the driver.
func (tbl SUserTable) BeginTx(opts *sql.TxOptions) (SUserTable, error) {
	d := tbl.db.(*sql.DB)
	var err error
	tbl.db, err = d.BeginTx(tbl.ctx, opts)
	return tbl, err
}

func (tbl SUserTable) logQuery(query string, args ...interface{}) {
	sqlgen2.LogQuery(tbl.logger, query, args...)
}


//--------------------------------------------------------------------------------

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// It returns the number of rows affected (of the database drive supports this).
func (tbl SUserTable) Exec(query string, args ...interface{}) (int64, error) {
	tbl.logQuery(query, args...)
	res, err := tbl.db.ExecContext(tbl.ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

//--------------------------------------------------------------------------------

// QueryOne is the low-level access function for one User.
// If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, *User will be nil.
func (tbl SUserTable) QueryOne(query string, args ...interface{}) (*User, error) {
	list, err := tbl.doQuery(true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

// Query is the low-level access function for Users.
func (tbl SUserTable) Query(query string, args ...interface{}) ([]*User, error) {
	return tbl.doQuery(false, query, args...)
}

func (tbl SUserTable) doQuery(firstOnly bool, query string, args ...interface{}) ([]*User, error) {
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanSUsers(rows, firstOnly)
}

//--------------------------------------------------------------------------------

// SliceUid gets the Uid column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl SUserTable) SliceUid(wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return tbl.getint64list("uid", wh, qc)
}

// SliceLogin gets the Login column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl SUserTable) SliceLogin(wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.getstringlist("login", wh, qc)
}

// SliceEmailAddress gets the EmailAddress column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl SUserTable) SliceEmailaddress(wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.getstringlist("emailaddress", wh, qc)
}

// SliceAvatar gets the Avatar column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl SUserTable) SliceAvatar(wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.getstringlist("avatar", wh, qc)
}

// SliceActive gets the Active column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl SUserTable) SliceActive(wh where.Expression, qc where.QueryConstraint) ([]bool, error) {
	return tbl.getboollist("active", wh, qc)
}

// SliceAdmin gets the Admin column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl SUserTable) SliceAdmin(wh where.Expression, qc where.QueryConstraint) ([]bool, error) {
	return tbl.getboollist("admin", wh, qc)
}

// SliceLastUpdated gets the LastUpdated column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl SUserTable) SliceLastupdated(wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return tbl.getint64list("lastupdated", wh, qc)
}


func (tbl SUserTable) getboollist(sqlname string, wh where.Expression, qc where.QueryConstraint) ([]bool, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", sqlname, tbl.prefix, tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var v bool
	list := make([]bool, 0, 10)
	for rows.Next() {
		err = rows.Scan(&v)
		if err != nil {
			return list, err
		}
		list = append(list, v)
	}
	return list, nil
}

func (tbl SUserTable) getint64list(sqlname string, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", sqlname, tbl.prefix, tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var v int64
	list := make([]int64, 0, 10)
	for rows.Next() {
		err = rows.Scan(&v)
		if err != nil {
			return list, err
		}
		list = append(list, v)
	}
	return list, nil
}

func (tbl SUserTable) getstringlist(sqlname string, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", sqlname, tbl.prefix, tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var v string
	list := make([]string, 0, 10)
	for rows.Next() {
		err = rows.Scan(&v)
		if err != nil {
			return list, err
		}
		list = append(list, v)
	}
	return list, nil
}


// scanSUsers reads table records into a slice of values.
func scanSUsers(rows *sql.Rows, firstOnly bool) ([]*User, error) {
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
