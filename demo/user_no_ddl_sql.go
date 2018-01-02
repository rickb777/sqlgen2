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

// V2UserTableName is the default name for this table.
const V2UserTableName = "users"

// V2UserTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type V2UserTable struct {
	Prefix, Name string
	Db           sqlgen2.Execer
	Ctx          context.Context
	Dialect      schema.Dialect
	Logger       *log.Logger
}

// Type conformance check
var _ sqlgen2.Table = &V2UserTable{}

// NewV2UserTable returns a new table instance.
// If a blank table name is supplied, the default name "users" will be used instead.
// The table name prefix is initially blank and the request context is the background.
func NewV2UserTable(name string, d sqlgen2.Execer, dialect schema.Dialect) V2UserTable {
	if name == "" {
		name = V2UserTableName
	}
	return V2UserTable{"", name, d, context.Background(), dialect, nil}
}

// WithPrefix sets the prefix for subsequent queries.
func (tbl V2UserTable) WithPrefix(pfx string) V2UserTable {
	tbl.Prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries.
func (tbl V2UserTable) WithContext(ctx context.Context) V2UserTable {
	tbl.Ctx = ctx
	return tbl
}

// WithLogger sets the logger for subsequent queries.
func (tbl V2UserTable) WithLogger(logger *log.Logger) V2UserTable {
	tbl.Logger = logger
	return tbl
}

// SetLogger sets the logger for subsequent queries, returning the interface.
func (tbl V2UserTable) SetLogger(logger *log.Logger) sqlgen2.Table {
	tbl.Logger = logger
	return tbl
}

// FullName gets the concatenated prefix and table name.
func (tbl V2UserTable) FullName() string {
	return tbl.Prefix + tbl.Name
}

func (tbl V2UserTable) prefixWithoutDot() string {
	last := len(tbl.Prefix)-1
	if last > 0 && tbl.Prefix[last] == '.' {
		return tbl.Prefix[0:last]
	}
	return tbl.Prefix
}

// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl V2UserTable) DB() *sql.DB {
	return tbl.Db.(*sql.DB)
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl V2UserTable) Tx() *sql.Tx {
	return tbl.Db.(*sql.Tx)
}

// IsTx tests whether this is within a transaction.
func (tbl V2UserTable) IsTx() bool {
	_, ok := tbl.Db.(*sql.Tx)
	return ok
}

// Begin starts a transaction. The default isolation level is dependent on the driver.
func (tbl V2UserTable) BeginTx(opts *sql.TxOptions) (V2UserTable, error) {
	d := tbl.Db.(*sql.DB)
	var err error
	tbl.Db, err = d.BeginTx(tbl.Ctx, opts)
	return tbl, err
}

func (tbl V2UserTable) logQuery(query string, args ...interface{}) {
	sqlgen2.LogQuery(tbl.Logger, query, args...)
}


//--------------------------------------------------------------------------------

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// It returns the number of rows affected (of the database drive supports this).
func (tbl V2UserTable) Exec(query string, args ...interface{}) (int64, error) {
	tbl.logQuery(query, args...)
	res, err := tbl.Db.ExecContext(tbl.Ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

//--------------------------------------------------------------------------------

// QueryOne is the low-level access function for one User.
func (tbl V2UserTable) QueryOne(query string, args ...interface{}) (*User, error) {
	tbl.logQuery(query, args...)
	row := tbl.Db.QueryRowContext(tbl.Ctx, query, args...)
	return scanV2User(row)
}

// Query is the low-level access function for Users.
func (tbl V2UserTable) Query(query string, args ...interface{}) ([]*User, error) {
	tbl.logQuery(query, args...)
	rows, err := tbl.Db.QueryContext(tbl.Ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanV2Users(rows)
}

//--------------------------------------------------------------------------------

// SelectOneSA allows a single User to be obtained from the table that match a 'where' clause and some limit.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
func (tbl V2UserTable) SelectOneSA(where, orderBy string, args ...interface{}) (*User, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s LIMIT 1", V2UserColumnNames, tbl.Prefix, tbl.Name, where, orderBy)
	return tbl.QueryOne(query, args...)
}

// SelectOne allows a single User to be obtained from the sqlgen2.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
func (tbl V2UserTable) SelectOne(where where.Expression, orderBy string) (*User, error) {
	wh, args := where.Build(tbl.Dialect)
	return tbl.SelectOneSA(wh, orderBy, args...)
}

// SelectSA allows Users to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
func (tbl V2UserTable) SelectSA(where, orderBy string, args ...interface{}) ([]*User, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", V2UserColumnNames, tbl.Prefix, tbl.Name, where, orderBy)
	return tbl.Query(query, args...)
}

// Select allows Users to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
func (tbl V2UserTable) Select(where where.Expression, orderBy string) ([]*User, error) {
	wh, args := where.Build(tbl.Dialect)
	return tbl.SelectSA(wh, orderBy, args...)
}

// CountSA counts Users in the table that match a 'where' clause.
func (tbl V2UserTable) CountSA(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s%s %s", tbl.Prefix, tbl.Name, where)
	tbl.logQuery(query, args...)
	row := tbl.Db.QueryRowContext(tbl.Ctx, query, args...)
	err = row.Scan(&count)
	return count, err
}

// Count counts the Users in the table that match a 'where' clause.
func (tbl V2UserTable) Count(where where.Expression) (count int64, err error) {
	wh, args := where.Build(tbl.Dialect)
	return tbl.CountSA(wh, args...)
}

const V2UserColumnNames = "uid, login, emailaddress, avatar, active, admin, fave, lastupdated, token, secret, hash"

//--------------------------------------------------------------------------------

// Insert adds new records for the Users. The Users have their primary key fields
// set to the new record identifiers.
// The User.PreInsert(Execer) method will be called, if it exists.
func (tbl V2UserTable) Insert(vv ...*User) error {
	var stmt, params string
	switch tbl.Dialect {
	case schema.Postgres:
		stmt = sqlInsertV2UserPostgres
		params = sV2UserDataColumnParamsPostgres
	default:
		stmt = sqlInsertV2UserSimple
		params = sV2UserDataColumnParamsSimple
	}

	query := fmt.Sprintf(stmt, tbl.Prefix, tbl.Name, params)
	st, err := tbl.Db.PrepareContext(tbl.Ctx, query)
	if err != nil {
		return err
	}
	defer st.Close()

	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPreInsert); ok {
			hook.PreInsert(tbl.Db)
		}

		fields, err := sliceV2UserWithoutPk(v)
		if err != nil {
			return err
		}

		tbl.logQuery(query, fields...)
		res, err := st.Exec(fields...)
		if err != nil {
			return err
		}

		v.Uid, err = res.LastInsertId()
		if err != nil {
			return err
		}
	}

	return nil
}

const sqlInsertV2UserSimple = `
INSERT INTO %s%s (
	login, 
	emailaddress, 
	avatar, 
	active, 
	admin, 
	fave, 
	lastupdated, 
	token, 
	secret, 
	hash
) VALUES (%s)
`

const sqlInsertV2UserPostgres = `
INSERT INTO %s%s (
	login, 
	emailaddress, 
	avatar, 
	active, 
	admin, 
	fave, 
	lastupdated, 
	token, 
	secret, 
	hash
) VALUES (%s)
`

const sV2UserDataColumnParamsSimple = "?,?,?,?,?,?,?,?,?,?"

const sV2UserDataColumnParamsPostgres = "$1,$2,$3,$4,$5,$6,$7,$8,$9,$10"

//--------------------------------------------------------------------------------

// UpdateFields updates one or more columns, given a 'where' clause.
func (tbl V2UserTable) UpdateFields(where where.Expression, fields ...sql.NamedArg) (int64, error) {
	query, args := tbl.updateFields(where, fields...)
	return tbl.Exec(query, args...)
}

func (tbl V2UserTable) updateFields(where where.Expression, fields ...sql.NamedArg) (string, []interface{}) {
	list := sqlgen2.NamedArgList(fields)
	assignments := strings.Join(list.Assignments(tbl.Dialect, 1), ", ")
	whereClause, wargs := where.Build(tbl.Dialect)
	query := fmt.Sprintf("UPDATE %s%s SET %s %s", tbl.Prefix, tbl.Name, assignments, whereClause)
	args := append(list.Values(), wargs...)
	return query, args
}

// Update updates records, matching them by primary key. It returns the number of rows affected.
// The User.PreUpdate(Execer) method will be called, if it exists.
func (tbl V2UserTable) Update(vv ...*User) (int64, error) {
	var stmt string
	switch tbl.Dialect {
	case schema.Postgres:
		stmt = sqlUpdateV2UserByPkPostgres
	default:
		stmt = sqlUpdateV2UserByPkSimple
	}

	var count int64
	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPreUpdate); ok {
			hook.PreUpdate(tbl.Db)
		}

		query := fmt.Sprintf(stmt, tbl.Prefix, tbl.Name)

		args, err := sliceV2UserWithoutPk(v)
		if err != nil {
			return count, err
		}

		args = append(args, v.Uid)
		tbl.logQuery(query, args...)
		n, err := tbl.Exec(query, args...)
		if err != nil {
			return count, err
		}
		count += n
	}
	return count, nil
}

const sqlUpdateV2UserByPkSimple = `
UPDATE %%s%%s SET 
 WHERE uid=?
`

const sqlUpdateV2UserByPkPostgres = `
UPDATE %%s%%s SET 
 WHERE uid=$12
`

//--------------------------------------------------------------------------------

// Delete deletes one or more rows from the table, given a 'where' clause.
func (tbl V2UserTable) Delete(where where.Expression) (int64, error) {
	query, args := tbl.deleteRows(where)
	return tbl.Exec(query, args...)
}

func (tbl V2UserTable) deleteRows(where where.Expression) (string, []interface{}) {
	whereClause, args := where.Build(tbl.Dialect)
	query := fmt.Sprintf("DELETE FROM %s%s %s", tbl.Prefix, tbl.Name, whereClause)
	return query, args
}

// scanV2User reads a table record into a single value.
func scanV2User(row *sql.Row) (*User, error) {
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
	var v10 string

	err := row.Scan(
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
		&v10,

	)
	if err != nil {
		return nil, err
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
	v.hash = v10

	return v, nil
}

// scanV2Users reads table records into a slice of values.
func scanV2Users(rows *sql.Rows) ([]*User, error) {
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
	var v10 string

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
			&v10,

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
		v.hash = v10

		vv = append(vv, v)
	}
	return vv, rows.Err()
}

func sliceV2UserWithoutPk(v *User) ([]interface{}, error) {

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
		v.hash,

	}, nil
}
