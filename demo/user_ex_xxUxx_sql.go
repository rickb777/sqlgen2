// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

package demo

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"github.com/rickb777/sqlgen2"
	"github.com/rickb777/sqlgen2/constraint"
	"github.com/rickb777/sqlgen2/require"
	"github.com/rickb777/sqlgen2/schema"
	"github.com/rickb777/sqlgen2/support"
	"github.com/rickb777/sqlgen2/where"
	"io"
	"log"
	"strings"
)

// UUserTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type UUserTable struct {
	name        sqlgen2.TableName
	database    *sqlgen2.Database
	db          sqlgen2.Execer
	constraints constraint.Constraints
	ctx			context.Context
}

// Type conformance checks
var _ sqlgen2.Table = &UUserTable{}
var _ sqlgen2.Table = &UUserTable{}

// NewUUserTable returns a new table instance.
// If a blank table name is supplied, the default name "users" will be used instead.
// The request context is initialised with the background.
func NewUUserTable(name sqlgen2.TableName, d *sqlgen2.Database) UUserTable {
	if name.Name == "" {
		name.Name = "users"
	}
	table := UUserTable{name, d, d.DB(), nil, context.Background()}
	table.constraints = append(table.constraints,
		constraint.FkConstraint{"addressid", constraint.Reference{"addresses", "id"}, "restrict", "restrict"})
	
	return table
}

// CopyTableAsUUserTable copies a table instance, retaining the name etc but
// providing methods appropriate for 'User'. It doesn't copy the constraints of the original table.
//
// It serves to provide methods appropriate for 'User'. This is most useful when this is used to represent a
// join result. In such cases, there won't be any need for DDL methods, nor Exec, Insert, Update or Delete.
func CopyTableAsUUserTable(origin sqlgen2.Table) UUserTable {
	return UUserTable{
		name:        origin.Name(),
		database:    origin.Database(),
		db:          origin.DB(),
		constraints: nil,
		ctx:         origin.Ctx(),
	}
}

// WithPrefix sets the table name prefix for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl UUserTable) WithPrefix(pfx string) UUserTable {
	tbl.name.Prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl UUserTable) WithContext(ctx context.Context) UUserTable {
	tbl.ctx = ctx
	return tbl
}

// Database gets the shared database information.
func (tbl UUserTable) Database() *sqlgen2.Database {
	return tbl.database
}

// Logger gets the trace logger.
func (tbl UUserTable) Logger() *log.Logger {
	return tbl.database.Logger()
}

// WithConstraint returns a modified Table with added data consistency constraints.
func (tbl UUserTable) WithConstraint(cc ...constraint.Constraint) UUserTable {
	tbl.constraints = append(tbl.constraints, cc...)
	return tbl
}

// Constraints returns the table's constraints.
func (tbl UUserTable) Constraints() constraint.Constraints {
	return tbl.constraints
}

// Ctx gets the current request context.
func (tbl UUserTable) Ctx() context.Context {
	return tbl.ctx
}

// Dialect gets the database dialect.
func (tbl UUserTable) Dialect() schema.Dialect {
	return tbl.database.Dialect()
}

// Name gets the table name.
func (tbl UUserTable) Name() sqlgen2.TableName {
	return tbl.name
}

// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl UUserTable) DB() *sql.DB {
	return tbl.db.(*sql.DB)
}

// Execer gets the wrapped database or transaction handle.
func (tbl UUserTable) Execer() sqlgen2.Execer {
	return tbl.db
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
// The result is a modified copy of the table; the original is unchanged.
func (tbl UUserTable) BeginTx(opts *sql.TxOptions) (UUserTable, error) {
	d := tbl.db.(*sql.DB)
	var err error
	tbl.db, err = d.BeginTx(tbl.ctx, opts)
	return tbl, tbl.logIfError(err)
}

// Using returns a modified Table using the transaction supplied. This is needed
// when making multiple queries across several tables within a single transaction.
// The result is a modified copy of the table; the original is unchanged.
func (tbl UUserTable) Using(tx *sql.Tx) UUserTable {
	tbl.db = tx
	return tbl
}

func (tbl UUserTable) logQuery(query string, args ...interface{}) {
	tbl.database.LogQuery(query, args...)
}

func (tbl UUserTable) logError(err error) error {
	return tbl.database.LogError(err)
}

func (tbl UUserTable) logIfError(err error) error {
	return tbl.database.LogIfError(err)
}


//--------------------------------------------------------------------------------

const NumUUserColumns = 12

const NumUUserDataColumns = 11

const UUserColumnNames = "uid,login,emailaddress,addressid,avatar,role,active,admin,fave,lastupdated,token,secret"

const UUserDataColumnNames = "login,emailaddress,addressid,avatar,role,active,admin,fave,lastupdated,token,secret"

const UUserPk = "uid"

//--------------------------------------------------------------------------------

// Exec executes a query without returning any rows.
// It returns the number of rows affected (if the database driver supports this).
//
// The args are for any placeholder parameters in the query.
func (tbl UUserTable) Exec(req require.Requirement, query string, args ...interface{}) (int64, error) {
	return support.Exec(tbl, req, query, args...)
}

//--------------------------------------------------------------------------------

// Query is the low-level access method for Users.
//
// It places a requirement, which may be nil, on the size of the expected results: this
// controls whether an error is generated when this expectation is not met.
//
// Note that this method applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl UUserTable) Query(req require.Requirement, query string, args ...interface{}) ([]*User, error) {
	query = tbl.ReplaceTableName(query)
	vv, err := tbl.doQuery(req, false, query, args...)
	return vv, err
}

// QueryOne is the low-level access method for one User.
// If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, *User will be nil.
//
// Note that this method applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl UUserTable) QueryOne(query string, args ...interface{}) (*User, error) {
	query = tbl.ReplaceTableName(query)
	return tbl.doQueryOne(nil, query, args...)
}

// MustQueryOne is the low-level access method for one User.
//
// It places a requirement that exactly one result must be found; an error is generated when this expectation is not met.
//
// Note that this method applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl UUserTable) MustQueryOne(query string, args ...interface{}) (*User, error) {
	query = tbl.ReplaceTableName(query)
	return tbl.doQueryOne(require.One, query, args...)
}

func (tbl UUserTable) doQueryOne(req require.Requirement, query string, args ...interface{}) (*User, error) {
	list, err := tbl.doQuery(req, true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

func (tbl UUserTable) doQuery(req require.Requirement, firstOnly bool, query string, args ...interface{}) ([]*User, error) {
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, tbl.logError(err)
	}
	defer rows.Close()

	vv, n, err := scanUUsers(rows, firstOnly)
	return vv, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(err, req, n))
}

func scanUUsers(rows *sql.Rows, firstOnly bool) (vv []*User, n int64, err error) {
	for rows.Next() {
		n++

		var v0 int64
		var v1 string
		var v2 string
		var v3 sql.NullInt64
		var v4 sql.NullString
		var v5 sql.NullString
		var v6 bool
		var v7 bool
		var v8 []byte
		var v9 int64
		var v10 string
		var v11 string

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
			&v11,
		)
		if err != nil {
			return vv, n, err
		}

		v := &User{}
		v.Uid = v0
		v.Login = v1
		v.EmailAddress = v2
		if v3.Valid {
			a := v3.Int64
			v.AddressId = &a
		}
		if v4.Valid {
			a := v4.String
			v.Avatar = &a
		}
		if v5.Valid {
			v.Role = new(Role)
			err = v.Role.Scan(v5.String)
			if err != nil {
				return nil, n, err
			}
		}
		v.Active = v6
		v.Admin = v7
		err = json.Unmarshal(v8, &v.Fave)
		if err != nil {
			return nil, n, err
		}
		v.LastUpdated = v9
		v.token = v10
		v.secret = v11

		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPostGet); ok {
			err = hook.PostGet()
			if err != nil {
				return vv, n, err
			}
		}

		vv = append(vv, v)

		if firstOnly {
			if rows.Next() {
				n++
			}
			return vv, n, rows.Err()
		}
	}

	return vv, n, rows.Err()
}

//--------------------------------------------------------------------------------

// QueryOneNullString is a low-level access method for one string. This can be used for function queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl UUserTable) QueryOneNullString(query string, args ...interface{}) (result sql.NullString, err error) {
	err = support.QueryOneNullThing(tbl, nil, &result, query, args...)
	return result, err
}

// MustQueryOneNullString is a low-level access method for one string. This can be used for function queries and
// such like.
//
// It places a requirement that exactly one result must be found; an error is generated when this expectation is not met.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl UUserTable) MustQueryOneNullString(query string, args ...interface{}) (result sql.NullString, err error) {
	err = support.QueryOneNullThing(tbl, require.One, &result, query, args...)
	return result, err
}

// QueryOneNullInt64 is a low-level access method for one int64. This can be used for 'COUNT(1)' queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl UUserTable) QueryOneNullInt64(query string, args ...interface{}) (result sql.NullInt64, err error) {
	err = support.QueryOneNullThing(tbl, nil, &result, query, args...)
	return result, err
}

// MustQueryOneNullInt64 is a low-level access method for one int64. This can be used for 'COUNT(1)' queries and
// such like.
//
// It places a requirement that exactly one result must be found; an error is generated when this expectation is not met.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl UUserTable) MustQueryOneNullInt64(query string, args ...interface{}) (result sql.NullInt64, err error) {
	err = support.QueryOneNullThing(tbl, require.One, &result, query, args...)
	return result, err
}

// QueryOneNullFloat64 is a low-level access method for one float64. This can be used for 'AVG(...)' queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl UUserTable) QueryOneNullFloat64(query string, args ...interface{}) (result sql.NullFloat64, err error) {
	err = support.QueryOneNullThing(tbl, nil, &result, query, args...)
	return result, err
}

// MustQueryOneNullFloat64 is a low-level access method for one float64. This can be used for 'AVG(...)' queries and
// such like.
//
// It places a requirement that exactly one result must be found; an error is generated when this expectation is not met.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl UUserTable) MustQueryOneNullFloat64(query string, args ...interface{}) (result sql.NullFloat64, err error) {
	err = support.QueryOneNullThing(tbl, require.One, &result, query, args...)
	return result, err
}

// ReplaceTableName replaces all occurrences of "{TABLE}" with the table's name.
func (tbl UUserTable) ReplaceTableName(query string) string {
	return strings.Replace(query, "{TABLE}", tbl.name.String(), -1)
}

func constructUUserUpdate(w io.Writer, v *User, dialect schema.Dialect) (s []interface{}, err error) {
	j := 1
	s = make([]interface{}, 0, 11)

	comma := ""

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "login", j)
	s = append(s, v.Login)
	comma = ", "
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "emailaddress", j)
	s = append(s, v.EmailAddress)
		j++

	io.WriteString(w, comma)
	if v.AddressId != nil {
		dialect.QuoteWithPlaceholder(w, "addressid", j)
		s = append(s, v.AddressId)
		j++
	} else {
		dialect.QuoteW(w, "addressid")
		io.WriteString(w, "=NULL")
	}

	io.WriteString(w, comma)
	if v.Avatar != nil {
		dialect.QuoteWithPlaceholder(w, "avatar", j)
		s = append(s, v.Avatar)
		j++
	} else {
		dialect.QuoteW(w, "avatar")
		io.WriteString(w, "=NULL")
	}

	io.WriteString(w, comma)
	if v.Role != nil {
		dialect.QuoteWithPlaceholder(w, "role", j)
		s = append(s, v.Role)
		j++
	} else {
		dialect.QuoteW(w, "role")
		io.WriteString(w, "=NULL")
	}

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "active", j)
	s = append(s, v.Active)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "admin", j)
	s = append(s, v.Admin)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "fave", j)
		j++
	x, err := json.Marshal(&v.Fave)
	if err != nil {
		return nil, err
	}
	s = append(s, x)

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "lastupdated", j)
	s = append(s, v.LastUpdated)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "token", j)
	s = append(s, v.token)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "secret", j)
	s = append(s, v.secret)
		j++

	return s, nil
}

// UpdateFields updates one or more columns, given a 'where' clause.
//
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl UUserTable) UpdateFields(req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error) {
	return support.UpdateFields(tbl, req, wh, fields...)
}

//--------------------------------------------------------------------------------

var allUUserQuotedUpdates = []string{
	// Sqlite
	"`login`=?,`emailaddress`=?,`addressid`=?,`avatar`=?,`role`=?,`active`=?,`admin`=?,`fave`=?,`lastupdated`=?,`token`=?,`secret`=? WHERE `uid`=?",
	// Mysql
	"`login`=?,`emailaddress`=?,`addressid`=?,`avatar`=?,`role`=?,`active`=?,`admin`=?,`fave`=?,`lastupdated`=?,`token`=?,`secret`=? WHERE `uid`=?",
	// Postgres
	`"login"=$2,"emailaddress"=$3,"addressid"=$4,"avatar"=$5,"role"=$6,"active"=$7,"admin"=$8,"fave"=$9,"lastupdated"=$10,"token"=$11,"secret"=$12 WHERE "uid"=$1`,
}

//--------------------------------------------------------------------------------

// Update updates records, matching them by primary key. It returns the number of rows affected.
// The User.PreUpdate(Execer) method will be called, if it exists.
func (tbl UUserTable) Update(req require.Requirement, vv ...*User) (int64, error) {
	if req == require.All {
		req = require.Exactly(len(vv))
	}

	var count int64
	dialect := tbl.Dialect()
	//columns := allUUserQuotedUpdates[dialect.Index()]
	//query := fmt.Sprintf("UPDATE %s SET %s", tbl.name, columns)

	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPreUpdate); ok {
			err := hook.PreUpdate()
			if err != nil {
				return count, tbl.logError(err)
			}
		}

		b := &bytes.Buffer{}
		io.WriteString(b, "UPDATE ")
		io.WriteString(b, tbl.name.String())
		io.WriteString(b, " SET ")

		args, err := constructUUserUpdate(b, v, dialect)
		k := len(args)
		args = append(args, v.Uid)
		if err != nil {
			return count, tbl.logError(err)
		}

		io.WriteString(b, " WHERE ")
		dialect.QuoteWithPlaceholder(b, "uid", k)

		query := b.String()
		n, err := tbl.Exec(nil, query, args...)
		if err != nil {
			return count, err
		}
		count += n
	}

	return count, tbl.logIfError(require.ErrorIfExecNotSatisfiedBy(req, count))
}
