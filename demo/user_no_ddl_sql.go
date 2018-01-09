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

// V2UserJoin holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type V2UserJoin struct {
	prefix, name string
	db           sqlgen2.Execer
	ctx          context.Context
	dialect      schema.Dialect
	logger       *log.Logger
}

// Type conformance check
var _ sqlgen2.Table = &V2UserJoin{}

// NewV2UserJoin returns a new table instance.
// If a blank table name is supplied, the default name "users" will be used instead.
// The table name prefix is initially blank and the request context is the background.
func NewV2UserJoin(name string, d sqlgen2.Execer, dialect schema.Dialect) V2UserJoin {
	if name == "" {
		name = "users"
	}
	return V2UserJoin{"", name, d, context.Background(), dialect, nil}
}

// CopyTableAsV2UserJoin copies a table instance, retaining the name etc but
// providing methods appropriate for 'User'.
func CopyTableAsV2UserJoin(origin sqlgen2.Table) V2UserJoin {
	return V2UserJoin{
		prefix:  origin.Prefix(),
		name:    origin.Name(),
		db:      origin.DB(),
		ctx:     origin.Ctx(),
		dialect: origin.Dialect(),
		logger:  origin.Logger(),
	}
}

// WithPrefix sets the table name prefix for subsequent queries.
func (tbl V2UserJoin) WithPrefix(pfx string) V2UserJoin {
	tbl.prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries.
func (tbl V2UserJoin) WithContext(ctx context.Context) V2UserJoin {
	tbl.ctx = ctx
	return tbl
}

// WithLogger sets the logger for subsequent queries.
func (tbl V2UserJoin) WithLogger(logger *log.Logger) V2UserJoin {
	tbl.logger = logger
	return tbl
}

// Ctx gets the current request context.
func (tbl V2UserJoin) Ctx() context.Context {
	return tbl.ctx
}

// Dialect gets the database dialect.
func (tbl V2UserJoin) Dialect() schema.Dialect {
	return tbl.dialect
}

// Logger gets the trace logger.
func (tbl V2UserJoin) Logger() *log.Logger {
	return tbl.logger
}

// SetLogger sets the logger for subsequent queries, returning the interface.
func (tbl V2UserJoin) SetLogger(logger *log.Logger) sqlgen2.Table {
	tbl.logger = logger
	return tbl
}

// Name gets the table name.
func (tbl V2UserJoin) Name() string {
	return tbl.name
}

// Prefix gets the table name prefix.
func (tbl V2UserJoin) Prefix() string {
	return tbl.prefix
}

// FullName gets the concatenated prefix and table name.
func (tbl V2UserJoin) FullName() string {
	return tbl.prefix + tbl.name
}

func (tbl V2UserJoin) prefixWithoutDot() string {
	last := len(tbl.prefix)-1
	if last > 0 && tbl.prefix[last] == '.' {
		return tbl.prefix[0:last]
	}
	return tbl.prefix
}

// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl V2UserJoin) DB() *sql.DB {
	return tbl.db.(*sql.DB)
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl V2UserJoin) Tx() *sql.Tx {
	return tbl.db.(*sql.Tx)
}

// IsTx tests whether this is within a transaction.
func (tbl V2UserJoin) IsTx() bool {
	_, ok := tbl.db.(*sql.Tx)
	return ok
}

// Begin starts a transaction. The default isolation level is dependent on the driver.
func (tbl V2UserJoin) BeginTx(opts *sql.TxOptions) (V2UserJoin, error) {
	d := tbl.db.(*sql.DB)
	var err error
	tbl.db, err = d.BeginTx(tbl.ctx, opts)
	return tbl, err
}

func (tbl V2UserJoin) logQuery(query string, args ...interface{}) {
	sqlgen2.LogQuery(tbl.logger, query, args...)
}


//--------------------------------------------------------------------------------

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// It returns the number of rows affected (of the database drive supports this).
func (tbl V2UserJoin) Exec(query string, args ...interface{}) (int64, error) {
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
func (tbl V2UserJoin) QueryOne(query string, args ...interface{}) (*User, error) {
	list, err := tbl.doQuery(true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

// Query is the low-level access function for Users.
func (tbl V2UserJoin) Query(query string, args ...interface{}) ([]*User, error) {
	return tbl.doQuery(false, query, args...)
}

func (tbl V2UserJoin) doQuery(firstOnly bool, query string, args ...interface{}) ([]*User, error) {
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanV2Users(rows, firstOnly)
}

//--------------------------------------------------------------------------------

// GetUser gets the record with a given primary key value.
// If not found, *User will be nil.
func (tbl V2UserJoin) GetUser(id int64) (*User, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s WHERE uid=?", V2UserColumnNames, tbl.prefix, tbl.name)
	return tbl.QueryOne(query, id)
}

// GetUsers gets records from the table according to a list of primary keys.
// Although the list of ids can be arbitrarily long, there are practical limits;
// note that Oracle DB has a limit of 1000.
func (tbl V2UserJoin) GetUsers(id ...int64) (list []*User, err error) {
	if len(id) > 0 {
		pl := tbl.dialect.Placeholders(len(id))
		query := fmt.Sprintf("SELECT %s FROM %s%s WHERE uid IN (%s)", V2UserColumnNames, tbl.prefix, tbl.name, pl)
		args := make([]interface{}, len(id))

		for i, v := range id {
			args[i] = v
		}

		list, err = tbl.Query(query, args...)
	}

	return list, err
}

//--------------------------------------------------------------------------------

// SliceUid gets the Uid column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'; otherwise use nil.
func (tbl V2UserJoin) SliceUid(wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return tbl.getint64list("uid", wh, qc)
}

// SliceLogin gets the Login column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'; otherwise use nil.
func (tbl V2UserJoin) SliceLogin(wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.getstringlist("login", wh, qc)
}

// SliceEmailAddress gets the EmailAddress column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'; otherwise use nil.
func (tbl V2UserJoin) SliceEmailaddress(wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.getstringlist("emailaddress", wh, qc)
}

// SliceAvatar gets the Avatar column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'; otherwise use nil.
func (tbl V2UserJoin) SliceAvatar(wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.getstringlist("avatar", wh, qc)
}

// SliceActive gets the Active column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'; otherwise use nil.
func (tbl V2UserJoin) SliceActive(wh where.Expression, qc where.QueryConstraint) ([]bool, error) {
	return tbl.getboollist("active", wh, qc)
}

// SliceAdmin gets the Admin column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'; otherwise use nil.
func (tbl V2UserJoin) SliceAdmin(wh where.Expression, qc where.QueryConstraint) ([]bool, error) {
	return tbl.getboollist("admin", wh, qc)
}

// SliceLastUpdated gets the LastUpdated column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'; otherwise use nil.
func (tbl V2UserJoin) SliceLastupdated(wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return tbl.getint64list("lastupdated", wh, qc)
}


func (tbl V2UserJoin) getboollist(sqlname string, wh where.Expression, qc where.QueryConstraint) ([]bool, error) {
	whs, args := wh.Build(tbl.dialect)
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

func (tbl V2UserJoin) getint64list(sqlname string, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	whs, args := wh.Build(tbl.dialect)
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

func (tbl V2UserJoin) getstringlist(sqlname string, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	whs, args := wh.Build(tbl.dialect)
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


//--------------------------------------------------------------------------------

// SelectOneSA allows a single User to be obtained from the table that match a 'where' clause
// and some limit.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
// If not found, *User will be nil.
func (tbl V2UserJoin) SelectOneSA(where, orderBy string, args ...interface{}) (*User, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s LIMIT 1", V2UserColumnNames, tbl.prefix, tbl.name, where, orderBy)
	return tbl.QueryOne(query, args...)
}

// SelectOne allows a single User to be obtained from the sqlgen2.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'; otherwise use nil.
// If not found, *Example will be nil.
func (tbl V2UserJoin) SelectOne(wh where.Expression, qc where.QueryConstraint) (*User, error) {
	whs, args := wh.Build(tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	return tbl.SelectOneSA(whs, orderBy, args...)
}

// SelectSA allows Users to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
func (tbl V2UserJoin) SelectSA(where, orderBy string, args ...interface{}) ([]*User, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", V2UserColumnNames, tbl.prefix, tbl.name, where, orderBy)
	return tbl.Query(query, args...)
}

// Select allows Users to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'; otherwise use nil.
func (tbl V2UserJoin) Select(wh where.Expression, qc where.QueryConstraint) ([]*User, error) {
	whs, args := wh.Build(tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	return tbl.SelectSA(whs, orderBy, args...)
}

// CountSA counts Users in the table that match a 'where' clause.
func (tbl V2UserJoin) CountSA(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s%s %s", tbl.prefix, tbl.name, where)
	tbl.logQuery(query, args...)
	row := tbl.db.QueryRowContext(tbl.ctx, query, args...)
	err = row.Scan(&count)
	return count, err
}

// Count counts the Users in the table that match a 'where' clause.
func (tbl V2UserJoin) Count(where where.Expression) (count int64, err error) {
	wh, args := where.Build(tbl.dialect)
	return tbl.CountSA(wh, args...)
}

const V2UserColumnNames = "uid, login, emailaddress, avatar, active, admin, fave, lastupdated, token, secret"

//--------------------------------------------------------------------------------

// Insert adds new records for the Users. The Users have their primary key fields
// set to the new record identifiers.
// The User.PreInsert(Execer) method will be called, if it exists.
func (tbl V2UserJoin) Insert(vv ...*User) error {
	var params string
	switch tbl.dialect {
	case schema.Postgres:
		params = sV2UserDataColumnParamsPostgres
	default:
		params = sV2UserDataColumnParamsSimple
	}

	query := fmt.Sprintf(sqlInsertV2User, tbl.prefix, tbl.name, params)
	st, err := tbl.db.PrepareContext(tbl.ctx, query)
	if err != nil {
		return err
	}
	defer st.Close()

	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPreInsert); ok {
			hook.PreInsert()
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

const sqlInsertV2User = `
INSERT INTO %s%s (
	login,
	emailaddress,
	avatar,
	active,
	admin,
	fave,
	lastupdated,
	token,
	secret
) VALUES (%s)
`

const sV2UserDataColumnParamsSimple = "?,?,?,?,?,?,?,?,?"

const sV2UserDataColumnParamsPostgres = "$1,$2,$3,$4,$5,$6,$7,$8,$9"

//--------------------------------------------------------------------------------

// UpdateFields updates one or more columns, given a 'where' clause.
func (tbl V2UserJoin) UpdateFields(where where.Expression, fields ...sql.NamedArg) (int64, error) {
	query, args := tbl.updateFields(where, fields...)
	return tbl.Exec(query, args...)
}

func (tbl V2UserJoin) updateFields(where where.Expression, fields ...sql.NamedArg) (string, []interface{}) {
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
func (tbl V2UserJoin) Update(vv ...*User) (int64, error) {
	var stmt string
	switch tbl.dialect {
	case schema.Postgres:
		stmt = sqlUpdateV2UserByPkPostgres
	default:
		stmt = sqlUpdateV2UserByPkSimple
	}

	query := fmt.Sprintf(stmt, tbl.prefix, tbl.name)

	var count int64
	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPreUpdate); ok {
			hook.PreUpdate()
		}

		args, err := sliceV2UserWithoutPk(v)
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

const sqlUpdateV2UserByPkSimple = `
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

const sqlUpdateV2UserByPkPostgres = `
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

//--------------------------------------------------------------------------------

// DeleteUsers deletes rows from the table, given some primary keys.
// The list of ids can be arbitrarily long.
func (tbl V2UserJoin) DeleteUsers(id ...int64) (int64, error) {
	const batch = 1000 // limited by Oracle DB
	const qt = "DELETE FROM %s%s WHERE uid IN (%s)"

	var count, n int64
	var err error
	var max = batch
	if len(id) < batch {
		max = len(id)
	}
	args := make([]interface{}, max)

	if len(id) > batch {
		pl := tbl.dialect.Placeholders(batch)
		query := fmt.Sprintf(qt, tbl.prefix, tbl.name, pl)

		for len(id) > batch {
			for i := 0; i < batch; i++ {
				args[i] = id[i]
			}

			n, err = tbl.Exec(query, args...)
			count += n
			if err != nil {
				return count, err
			}

			id = id[batch:]
		}
	}

	if len(id) > 0 {
		pl := tbl.dialect.Placeholders(len(id))
		query := fmt.Sprintf(qt, tbl.prefix, tbl.name, pl)

		for i := 0; i < batch; i++ {
			args[i] = id[i]
		}

		n, err = tbl.Exec(query, args...)
		count += n
	}

	return count, err
}

// Delete deletes one or more rows from the table, given a 'where' clause.
func (tbl V2UserJoin) Delete(where where.Expression) (int64, error) {
	query, args := tbl.deleteRows(where)
	return tbl.Exec(query, args...)
}

func (tbl V2UserJoin) deleteRows(where where.Expression) (string, []interface{}) {
	whereClause, args := where.Build(tbl.dialect)
	query := fmt.Sprintf("DELETE FROM %s%s %s", tbl.prefix, tbl.name, whereClause)
	return query, args
}

//--------------------------------------------------------------------------------

// scanV2Users reads table records into a slice of values.
func scanV2Users(rows *sql.Rows, firstOnly bool) ([]*User, error) {
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

	}, nil
}
