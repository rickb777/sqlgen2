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
	"strings"
)

// UUserTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type UUserTable struct {
	prefix, name string
	db           sqlgen2.Execer
	ctx          context.Context
	dialect      schema.Dialect
	logger       *log.Logger
}

// Type conformance check
var _ sqlgen2.Table = &UUserTable{}

// NewUUserTable returns a new table instance.
// If a blank table name is supplied, the default name "users" will be used instead.
// The table name prefix is initially blank and the request context is the background.
func NewUUserTable(name string, d sqlgen2.Execer, dialect schema.Dialect) UUserTable {
	if name == "" {
		name = "users"
	}
	return UUserTable{"", name, d, context.Background(), dialect, nil}
}

// CopyTableAsUUserTable copies a table instance, retaining the name etc but
// providing methods appropriate for 'User'.
func CopyTableAsUUserTable(origin sqlgen2.Table) UUserTable {
	return UUserTable{
		prefix:  origin.Prefix(),
		name:    origin.Name(),
		db:      origin.DB(),
		ctx:     origin.Ctx(),
		dialect: origin.Dialect(),
		logger:  origin.Logger(),
	}
}

// WithPrefix sets the table name prefix for subsequent queries.
func (tbl UUserTable) WithPrefix(pfx string) UUserTable {
	tbl.prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries.
func (tbl UUserTable) WithContext(ctx context.Context) UUserTable {
	tbl.ctx = ctx
	return tbl
}

// WithLogger sets the logger for subsequent queries.
func (tbl UUserTable) WithLogger(logger *log.Logger) UUserTable {
	tbl.logger = logger
	return tbl
}

// Ctx gets the current request context.
func (tbl UUserTable) Ctx() context.Context {
	return tbl.ctx
}

// Dialect gets the database dialect.
func (tbl UUserTable) Dialect() schema.Dialect {
	return tbl.dialect
}

// Logger gets the trace logger.
func (tbl UUserTable) Logger() *log.Logger {
	return tbl.logger
}

// SetLogger sets the logger for subsequent queries, returning the interface.
func (tbl UUserTable) SetLogger(logger *log.Logger) sqlgen2.Table {
	tbl.logger = logger
	return tbl
}

// Name gets the table name.
func (tbl UUserTable) Name() string {
	return tbl.name
}

// Prefix gets the table name prefix.
func (tbl UUserTable) Prefix() string {
	return tbl.prefix
}

// FullName gets the concatenated prefix and table name.
func (tbl UUserTable) FullName() string {
	return tbl.prefix + tbl.name
}

func (tbl UUserTable) prefixWithoutDot() string {
	last := len(tbl.prefix)-1
	if last > 0 && tbl.prefix[last] == '.' {
		return tbl.prefix[0:last]
	}
	return tbl.prefix
}

// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl UUserTable) DB() *sql.DB {
	return tbl.db.(*sql.DB)
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl UUserTable) Tx() *sql.Tx {
	return tbl.db.(*sql.Tx)
}

// IsTx tests whether this is within a transaction.
func (tbl UUserTable) IsTx() bool {
	_, ok := tbl.db.(*sql.Tx)
	return ok
}

// Begin starts a transaction. The default isolation level is dependent on the driver.
func (tbl UUserTable) BeginTx(opts *sql.TxOptions) (UUserTable, error) {
	d := tbl.db.(*sql.DB)
	var err error
	tbl.db, err = d.BeginTx(tbl.ctx, opts)
	return tbl, err
}

func (tbl UUserTable) logQuery(query string, args ...interface{}) {
	sqlgen2.LogQuery(tbl.logger, query, args...)
}


//--------------------------------------------------------------------------------

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// It returns the number of rows affected (of the database drive supports this).
func (tbl UUserTable) Exec(query string, args ...interface{}) (int64, error) {
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
func (tbl UUserTable) QueryOne(query string, args ...interface{}) (*User, error) {
	list, err := tbl.doQuery(true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

// Query is the low-level access function for Users.
func (tbl UUserTable) Query(query string, args ...interface{}) ([]*User, error) {
	return tbl.doQuery(false, query, args...)
}

func (tbl UUserTable) doQuery(firstOnly bool, query string, args ...interface{}) ([]*User, error) {
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanUUsers(rows, firstOnly)
}

//--------------------------------------------------------------------------------

// UpdateFields updates one or more columns, given a 'where' clause.
func (tbl UUserTable) UpdateFields(where where.Expression, fields ...sql.NamedArg) (int64, error) {
	query, args := tbl.updateFields(where, fields...)
	return tbl.Exec(query, args...)
}

func (tbl UUserTable) updateFields(where where.Expression, fields ...sql.NamedArg) (string, []interface{}) {
	list := sqlgen2.NamedArgList(fields)
	assignments := strings.Join(list.Assignments(tbl.dialect, 1), ", ")
	whereClause, wargs := where.Build(tbl.dialect)
	query := fmt.Sprintf("UPDATE %s%s SET %s %s", tbl.prefix, tbl.name, assignments, whereClause)
	args := append(list.Values(), wargs...)
	return query, args
}

//--------------------------------------------------------------------------------

// Update updates records, matching them by primary key. It returns the number of rows affected.
// The User.PreUpdate(Execer) method will be called, if it exists.
func (tbl UUserTable) Update(vv ...*User) (int64, error) {
	var stmt string
	switch tbl.dialect {
	case schema.Postgres:
		stmt = sqlUpdateUUserByPkPostgres
	default:
		stmt = sqlUpdateUUserByPkSimple
	}

	query := fmt.Sprintf(stmt, tbl.prefix, tbl.name)

	var count int64
	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPreUpdate); ok {
			hook.PreUpdate()
		}

		args, err := sliceUUserWithoutPk(v)
		if err != nil {
			return count, err
		}

		args = append(args, v.Uid)
		n, err := tbl.Exec(query, args...)
		if err != nil {
			return count, err
		}
		count += n
	}
	return count, nil
}

const sqlUpdateUUserByPkSimple = `
UPDATE %s%s SET
	login=?,
	emailaddress=?,
	avatar=?,
	active=?,
	admin=?,
	fave=?,
	lastupdated=?,
	token=?,
	secret=?
WHERE uid=?
`

const sqlUpdateUUserByPkPostgres = `
UPDATE %s%s SET
	login=$2,
	emailaddress=$3,
	avatar=$4,
	active=$5,
	admin=$6,
	fave=$7,
	lastupdated=$8,
	token=$9,
	secret=$10
WHERE uid=$1
`

func sliceUUserWithoutPk(v *User) ([]interface{}, error) {

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

// scanUUsers reads table records into a slice of values.
func scanUUsers(rows *sql.Rows, firstOnly bool) ([]*User, error) {
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