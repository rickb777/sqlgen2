// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

package demo

import (
	"database/sql"
	"fmt"
	"github.com/rickb777/sqlgen2/dialect"
	"github.com/rickb777/sqlgen2/where"
)

// V2UserTableName is the default name for this table.
const V2UserTableName = "users"

// V2UserTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type V2UserTable struct {
	Prefix, Name string
	Db           *sql.DB
	Dialect      dialect.Dialect
}

// NewV2UserTable returns a new table instance.
func NewV2UserTable(prefix, name string, db *sql.DB, dialect dialect.Dialect) V2UserTable {
	if name == "" {
		name = V2UserTableName
	}
	return V2UserTable{prefix, name, db, dialect}
}

// ScanV2User reads a database record into a single value.
func ScanV2User(row *sql.Row) (*User, error) {
	var v0 int64
	var v1 string
	var v2 string
	var v3 string
	var v4 bool
	var v5 bool
	var v6 string
	var v7 string
	var v8 string

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

	)
	if err != nil {
		return nil, err
	}

	v := &User{}
	v.Uid = v0
	v.Login = v1
	v.Email = v2
	v.Avatar = v3
	v.Active = v4
	v.Admin = v5
	v.token = v6
	v.secret = v7
	v.hash = v8

	return v, nil
}

// ScanV2Users reads database records into a slice of values.
func ScanV2Users(rows *sql.Rows) ([]*User, error) {
	var err error
	var vv []*User

	var v0 int64
	var v1 string
	var v2 string
	var v3 string
	var v4 bool
	var v5 bool
	var v6 string
	var v7 string
	var v8 string

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

		)
		if err != nil {
			return vv, err
		}

		v := &User{}
		v.Uid = v0
		v.Login = v1
		v.Email = v2
		v.Avatar = v3
		v.Active = v4
		v.Admin = v5
		v.token = v6
		v.secret = v7
		v.hash = v8

		vv = append(vv, v)
	}
	return vv, rows.Err()
}

func SliceV2User(v *User) []interface{} {
	var v0 int64
	var v1 string
	var v2 string
	var v3 string
	var v4 bool
	var v5 bool
	var v6 string
	var v7 string
	var v8 string

	v0 = v.Uid
	v1 = v.Login
	v2 = v.Email
	v3 = v.Avatar
	v4 = v.Active
	v5 = v.Admin
	v6 = v.token
	v7 = v.secret
	v8 = v.hash

	return []interface{}{
		v0,
		v1,
		v2,
		v3,
		v4,
		v5,
		v6,
		v7,
		v8,

	}
}

func SliceV2UserWithoutPk(v *User) []interface{} {
	var v1 string
	var v2 string
	var v3 string
	var v4 bool
	var v5 bool
	var v6 string
	var v7 string
	var v8 string

	v1 = v.Login
	v2 = v.Email
	v3 = v.Avatar
	v4 = v.Active
	v5 = v.Admin
	v6 = v.token
	v7 = v.secret
	v8 = v.hash

	return []interface{}{
		v1,
		v2,
		v3,
		v4,
		v5,
		v6,
		v7,
		v8,

	}
}

// QueryOne is the low-level access function for one User.
func (tbl V2UserTable) QueryOne(query string, args ...interface{}) (*User, error) {
	row := tbl.Db.QueryRow(query, args...)
	return ScanV2User(row)
}

// SelectOneSA allows a single User to be obtained from the database using supplied dialect-specific parameters.
func (tbl V2UserTable) SelectOneSA(where, limitClause string, args ...interface{}) (*User, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", V2UserColumnNames, tbl.Prefix, tbl.Name, where, limitClause)
	return tbl.QueryOne(query, args...)
}

// SelectOne allows a single User to be obtained from the database.
func (tbl V2UserTable) SelectOne(where where.Expression, dialect dialect.Dialect) (*User, error) {
	wh, args := where.Build(dialect)
	return tbl.SelectOneSA(wh, "LIMIT 1", args)
}

func (tbl V2UserTable) Query(query string, args ...interface{}) ([]*User, error) {
	rows, err := tbl.Db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanV2Users(rows)
}

// SelectSA allows Users to be obtained from the database using supplied dialect-specific parameters.
func (tbl V2UserTable) SelectSA(where string, args ...interface{}) ([]*User, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s", V2UserColumnNames, tbl.Prefix, tbl.Name, where)
	return tbl.Query(query, args...)
}

// Select allows Users to be obtained from the database that match a 'where' clause.
func (tbl V2UserTable) Select(where where.Expression, dialect dialect.Dialect) ([]*User, error) {
	return tbl.SelectSA(where.Build(dialect))
}

// CountSA counts Users in the database using supplied dialect-specific parameters.
func (tbl V2UserTable) CountSA(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s%s %s", tbl.Prefix, tbl.Name, where)
	row := tbl.Db.QueryRow(query, args)
	err = row.Scan(&count)
	return count, err
}

// Count counts the Users in the database that match a 'where' clause.
func (tbl V2UserTable) Count(where where.Expression, dialect dialect.Dialect) (count int64, err error) {
	return tbl.CountSA(where.Build(dialect))
}

const V2UserColumnNames = "uid, login, email, avatar, active, admin, token, secret, hash"

// Insert adds new records for the Users. The Users have their primary key fields
// set to the new record identifiers.
func (tbl V2UserTable) Insert(vv ...*User) error {
	var stmt, params string
	switch tbl.Dialect {
	case dialect.Postgres:
		stmt = sqlInsertV2UserPostgres
		params = sV2UserDataColumnParamsPostgres
	default:
		stmt = sqlInsertV2UserSimple
		params = sV2UserDataColumnParamsSimple
	}
	st, err := tbl.Db.Prepare(fmt.Sprintf(stmt, tbl.Prefix, tbl.Name, params))
	if err != nil {
		return err
	}
	defer st.Close()

	for _, v := range vv {
		res, err := st.Exec(SliceV2UserWithoutPk(v)...)
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
	email,
	avatar,
	active,
	admin,
	token,
	secret,
	hash
) VALUES (%s)
`

const sqlInsertV2UserPostgres = `
INSERT INTO %s%s (
	login,
	email,
	avatar,
	active,
	admin,
	token,
	secret,
	hash
) VALUES (%s)
`

const sV2UserDataColumnParamsSimple = "?,?,?,?,?,?,?,?"

const sV2UserDataColumnParamsPostgres = "$1,$2,$3,$4,$5,$6,$7,$8"

// Update updates a record. It returns the number of rows affected.
// Not every database or database driver may support this.
func (tbl V2UserTable) Update(v *User) (int64, error) {
	var stmt string
	switch tbl.Dialect {
	case dialect.Postgres:
		stmt = sqlUpdateV2UserByPkPostgres
	default:
		stmt = sqlUpdateV2UserByPkSimple
	}
	query := fmt.Sprintf(stmt, tbl.Prefix, tbl.Name)
	args := SliceV2UserWithoutPk(v)
	args = append(args, v.Uid)
	return tbl.Exec(query, args...)
}

const sqlUpdateV2UserByPkSimple = `
UPDATE %s%s SET 
	login=?,
	email=?,
	avatar=?,
	active=?,
	admin=?,
	token=?,
	secret=?,
	hash=? 
 WHERE uid=?
`

const sqlUpdateV2UserByPkPostgres = `
UPDATE %s%s SET 
	login=$2,
	email=$3,
	avatar=$4,
	active=$5,
	admin=$6,
	token=$7,
	secret=$8,
	hash=$9 
 WHERE uid=$10
`

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// It returns the number of rows affected.
// Not every database or database driver may support this.
func (tbl V2UserTable) Exec(query string, args ...interface{}) (int64, error) {
	res, err := tbl.Db.Exec(query, args...)
	if err != nil {
		return 0, nil
	}
	return res.RowsAffected()
}
