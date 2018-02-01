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
	"io"
	"log"
	"strings"
)

// CUserTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type CUserTable struct {
	name        sqlgen2.TableName
	database    *sqlgen2.Database
	db          sqlgen2.Execer
	constraints constraint.Constraints
	ctx			context.Context
}

// Type conformance checks
var _ sqlgen2.Table = &CUserTable{}
var _ sqlgen2.Table = &CUserTable{}

// NewCUserTable returns a new table instance.
// If a blank table name is supplied, the default name "users" will be used instead.
// The request context is initialised with the background.
func NewCUserTable(name sqlgen2.TableName, d *sqlgen2.Database) CUserTable {
	if name.Name == "" {
		name.Name = "users"
	}
	table := CUserTable{name, d, d.DB(), nil, context.Background()}
	table.constraints = append(table.constraints,
		constraint.FkConstraint{"addressid", constraint.Reference{"addresses", "id"}, "restrict", "restrict"})
	
	return table
}

// CopyTableAsCUserTable copies a table instance, retaining the name etc but
// providing methods appropriate for 'User'. It doesn't copy the constraints of the original table.
//
// It serves to provide methods appropriate for 'User'. This is most useulf when thie is used to represent a
// join result. In such cases, there won't be any need for DDL methods, nor Exec, Insert, Update or Delete.
func CopyTableAsCUserTable(origin sqlgen2.Table) CUserTable {
	return CUserTable{
		name:        origin.Name(),
		database:    origin.Database(),
		db:          origin.DB(),
		constraints: nil,
		ctx:         origin.Ctx(),
	}
}

// WithPrefix sets the table name prefix for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl CUserTable) WithPrefix(pfx string) CUserTable {
	tbl.name.Prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl CUserTable) WithContext(ctx context.Context) CUserTable {
	tbl.ctx = ctx
	return tbl
}

// Database gets the shared database information.
func (tbl CUserTable) Database() *sqlgen2.Database {
	return tbl.database
}

// Logger gets the trace logger.
func (tbl CUserTable) Logger() *log.Logger {
	return tbl.database.Logger()
}

// WithConstraint returns a modified Table with added data consistency constraints.
func (tbl CUserTable) WithConstraint(cc ...constraint.Constraint) CUserTable {
	tbl.constraints = append(tbl.constraints, cc...)
	return tbl
}

// Ctx gets the current request context.
func (tbl CUserTable) Ctx() context.Context {
	return tbl.ctx
}

// Dialect gets the database dialect.
func (tbl CUserTable) Dialect() schema.Dialect {
	return tbl.database.Dialect()
}

// Name gets the table name.
func (tbl CUserTable) Name() sqlgen2.TableName {
	return tbl.name
}

// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl CUserTable) DB() *sql.DB {
	return tbl.db.(*sql.DB)
}

// Execer gets the wrapped database or transaction handle.
func (tbl CUserTable) Execer() sqlgen2.Execer {
	return tbl.db
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl CUserTable) Tx() *sql.Tx {
	return tbl.db.(*sql.Tx)
}

// IsTx tests whether this is within a transaction.
func (tbl CUserTable) IsTx() bool {
	_, ok := tbl.db.(*sql.Tx)
	return ok
}

// Begin starts a transaction. The default isolation level is dependent on the driver.
// The result is a modified copy of the table; the original is unchanged.
func (tbl CUserTable) BeginTx(opts *sql.TxOptions) (CUserTable, error) {
	d := tbl.db.(*sql.DB)
	var err error
	tbl.db, err = d.BeginTx(tbl.ctx, opts)
	return tbl, tbl.logIfError(err)
}

// Using returns a modified Table using the transaction supplied. This is needed
// when making multiple queries across several tables within a single transaction.
// The result is a modified copy of the table; the original is unchanged.
func (tbl CUserTable) Using(tx *sql.Tx) CUserTable {
	tbl.db = tx
	return tbl
}

func (tbl CUserTable) logQuery(query string, args ...interface{}) {
	support.LogQuery(tbl.Logger(), query, args...)
}

func (tbl CUserTable) logError(err error) error {
	return support.LogError(tbl.Logger(), err)
}

func (tbl CUserTable) logIfError(err error) error {
	return support.LogIfError(tbl.Logger(), err)
}


//--------------------------------------------------------------------------------

const NumCUserColumns = 12

const NumCUserDataColumns = 11

const CUserColumnNames = "uid,login,emailaddress,addressid,avatar,role,active,admin,fave,lastupdated,token,secret"

const CUserDataColumnNames = "login,emailaddress,addressid,avatar,role,active,admin,fave,lastupdated,token,secret"

const CUserPk = "uid"

//--------------------------------------------------------------------------------

// Query is the low-level access method for Users.
//
// It places a requirement, which may be nil, on the size of the expected results: this
// controls whether an error is generated when this expectation is not met.
//
// Note that this method applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl CUserTable) Query(req require.Requirement, query string, args ...interface{}) ([]*User, error) {
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
func (tbl CUserTable) QueryOne(query string, args ...interface{}) (*User, error) {
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
func (tbl CUserTable) MustQueryOne(query string, args ...interface{}) (*User, error) {
	query = tbl.ReplaceTableName(query)
	return tbl.doQueryOne(require.One, query, args...)
}

func (tbl CUserTable) doQueryOne(req require.Requirement, query string, args ...interface{}) (*User, error) {
	list, err := tbl.doQuery(req, true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

func (tbl CUserTable) doQuery(req require.Requirement, firstOnly bool, query string, args ...interface{}) ([]*User, error) {
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, tbl.logError(err)
	}
	defer rows.Close()

	vv, n, err := scanCUsers(rows, firstOnly)
	return vv, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(err, req, n))
}

func scanCUsers(rows *sql.Rows, firstOnly bool) (vv []*User, n int64, err error) {
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
func (tbl CUserTable) QueryOneNullString(query string, args ...interface{}) (result sql.NullString, err error) {
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
func (tbl CUserTable) MustQueryOneNullString(query string, args ...interface{}) (result sql.NullString, err error) {
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
func (tbl CUserTable) QueryOneNullInt64(query string, args ...interface{}) (result sql.NullInt64, err error) {
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
func (tbl CUserTable) MustQueryOneNullInt64(query string, args ...interface{}) (result sql.NullInt64, err error) {
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
func (tbl CUserTable) QueryOneNullFloat64(query string, args ...interface{}) (result sql.NullFloat64, err error) {
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
func (tbl CUserTable) MustQueryOneNullFloat64(query string, args ...interface{}) (result sql.NullFloat64, err error) {
	err = support.QueryOneNullThing(tbl, require.One, &result, query, args...)
	return result, err
}

// ReplaceTableName replaces all occurrences of "{TABLE}" with the table's name.
func (tbl CUserTable) ReplaceTableName(query string) string {
	return strings.Replace(query, "{TABLE}", tbl.name.String(), -1)
}

func constructCUserInsert(w io.Writer, v *User, dialect schema.Dialect, withPk bool) (s []interface{}, err error) {
	s = make([]interface{}, 0, 12)

	comma := ""
	io.WriteString(w, " (")

	if withPk {
		dialect.QuoteW(w, "uid")
		comma = ","
		s = append(s, v.Uid)
	}

	io.WriteString(w, comma)

	dialect.QuoteW(w, "login")
	s = append(s, v.Login)
	comma = ","
	io.WriteString(w, comma)

	dialect.QuoteW(w, "emailaddress")
	s = append(s, v.EmailAddress)
	if v.AddressId != nil {
		io.WriteString(w, comma)

		dialect.QuoteW(w, "addressid")
		s = append(s, v.AddressId)
	}
	if v.Avatar != nil {
		io.WriteString(w, comma)

		dialect.QuoteW(w, "avatar")
		s = append(s, v.Avatar)
	}
	if v.Role != nil {
		io.WriteString(w, comma)

		dialect.QuoteW(w, "role")
		s = append(s, v.Role)
	}
	io.WriteString(w, comma)

	dialect.QuoteW(w, "active")
	s = append(s, v.Active)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "admin")
	s = append(s, v.Admin)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "fave")
	x, err := json.Marshal(&v.Fave)
	if err != nil {
		return nil, err
	}
	s = append(s, x)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "lastupdated")
	s = append(s, v.LastUpdated)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "token")
	s = append(s, v.token)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "secret")
	s = append(s, v.secret)
	io.WriteString(w, ")")
	return s, nil
}

//--------------------------------------------------------------------------------

var allCUserQuotedInserts = []string{
	// Sqlite
	"(`login`,`emailaddress`,`addressid`,`avatar`,`role`,`active`,`admin`,`fave`,`lastupdated`,`token`,`secret`) VALUES (?,?,?,?,?,?,?,?,?,?,?)",
	// Mysql
	"(`login`,`emailaddress`,`addressid`,`avatar`,`role`,`active`,`admin`,`fave`,`lastupdated`,`token`,`secret`) VALUES (?,?,?,?,?,?,?,?,?,?,?)",
	// Postgres
	`("login","emailaddress","addressid","avatar","role","active","admin","fave","lastupdated","token","secret") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) returning "uid"`,
}

//--------------------------------------------------------------------------------

// Insert adds new records for the Users.
// The Users have their primary key fields set to the new record identifiers.
// The User.PreInsert() method will be called, if it exists.
func (tbl CUserTable) Insert(req require.Requirement, vv ...*User) error {
	if req == require.All {
		req = require.Exactly(len(vv))
	}

	var count int64

	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPreInsert); ok {
			err := hook.PreInsert()
			if err != nil {
				return tbl.logError(err)
			}
		}

		b := &bytes.Buffer{}
		io.WriteString(b, "INSERT INTO ")
		io.WriteString(b, tbl.name.String())

		fields, err := constructCUserInsert(b, v, tbl.Dialect(), false)
		if err != nil {
			return tbl.logError(err)
		}

		io.WriteString(b, " VALUES (")
		io.WriteString(b, tbl.Dialect().Placeholders(len(fields)))
		io.WriteString(b, ")")

		query := b.String()
		tbl.logQuery(query, fields...)
		res, err := tbl.db.ExecContext(tbl.ctx, query, fields...)
		if err != nil {
			return tbl.logError(err)
		}

		v.Uid, err = res.LastInsertId()
		if err != nil {
			return tbl.logError(err)
		}

		n, err := res.RowsAffected()
		if err != nil {
			return tbl.logError(err)
		}
		count += n
	}

	return tbl.logIfError(require.ErrorIfExecNotSatisfiedBy(req, count))
}
