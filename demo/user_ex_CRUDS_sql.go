// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

package demo

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/rickb777/sqlgen2"
	"github.com/rickb777/sqlgen2/constraint"
	"github.com/rickb777/sqlgen2/require"
	"github.com/rickb777/sqlgen2/schema"
	"github.com/rickb777/sqlgen2/support"
	"github.com/rickb777/sqlgen2/where"
	"io"
	"log"
)

// AUserTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type AUserTable struct {
	name        sqlgen2.TableName
	database    *sqlgen2.Database
	db          sqlgen2.Execer
	constraints constraint.Constraints
	ctx			context.Context
	pk          string
}

// Type conformance checks
var _ sqlgen2.TableWithIndexes = &AUserTable{}
var _ sqlgen2.TableWithCrud = &AUserTable{}

// NewAUserTable returns a new table instance.
// If a blank table name is supplied, the default name "users" will be used instead.
// The request context is initialised with the background.
func NewAUserTable(name string, d *sqlgen2.Database) AUserTable {
	if name == "" {
		name = "users"
	}
	var constraints constraint.Constraints
	constraints = append(constraints, constraint.FkConstraint{"addressid", constraint.Reference{"addresses", "id"}, "restrict", "restrict"})
	
	return AUserTable{
		name:        sqlgen2.TableName{"", name},
		database:    d,
		db:          d.DB(),
		constraints: constraints,
		ctx:         context.Background(),
		pk:          "uid",
	}
}

// CopyTableAsAUserTable copies a table instance, retaining the name etc but
// providing methods appropriate for 'User'. It doesn't copy the constraints of the original table.
//
// It serves to provide methods appropriate for 'User'. This is most useful when this is used to represent a
// join result. In such cases, there won't be any need for DDL methods, nor Exec, Insert, Update or Delete.
func CopyTableAsAUserTable(origin sqlgen2.Table) AUserTable {
	return AUserTable{
		name:        origin.Name(),
		database:    origin.Database(),
		db:          origin.DB(),
		constraints: nil,
		ctx:         context.Background(),
		pk:          "uid",
	}
}


// SetPkColumn sets the name of the primary key column. It defaults to "uid".
// The result is a modified copy of the table; the original is unchanged.
func (tbl AUserTable) SetPkColumn(pk string) AUserTable {
	tbl.pk = pk
	return tbl
}


// WithPrefix sets the table name prefix for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl AUserTable) WithPrefix(pfx string) AUserTable {
	tbl.name.Prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries via this table.
// The result is a modified copy of the table; the original is unchanged.
//
// The shared context in the *Database is not altered by this method. So it
// is possible to use different contexts for different (groups of) queries.
func (tbl AUserTable) WithContext(ctx context.Context) AUserTable {
	tbl.ctx = ctx
	return tbl
}

// Database gets the shared database information.
func (tbl AUserTable) Database() *sqlgen2.Database {
	return tbl.database
}

// Logger gets the trace logger.
func (tbl AUserTable) Logger() *log.Logger {
	return tbl.database.Logger()
}

// WithConstraint returns a modified Table with added data consistency constraints.
func (tbl AUserTable) WithConstraint(cc ...constraint.Constraint) AUserTable {
	tbl.constraints = append(tbl.constraints, cc...)
	return tbl
}

// Constraints returns the table's constraints.
func (tbl AUserTable) Constraints() constraint.Constraints {
	return tbl.constraints
}

// Ctx gets the current request context.
func (tbl AUserTable) Ctx() context.Context {
	return tbl.ctx
}

// Dialect gets the database dialect.
func (tbl AUserTable) Dialect() schema.Dialect {
	return tbl.database.Dialect()
}

// Name gets the table name.
func (tbl AUserTable) Name() sqlgen2.TableName {
	return tbl.name
}


// PkColumn gets the column name used as a primary key.
func (tbl AUserTable) PkColumn() string {
	return tbl.pk
}


// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl AUserTable) DB() *sql.DB {
	return tbl.db.(*sql.DB)
}

// Execer gets the wrapped database or transaction handle.
func (tbl AUserTable) Execer() sqlgen2.Execer {
	return tbl.db
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl AUserTable) Tx() *sql.Tx {
	return tbl.db.(*sql.Tx)
}

// IsTx tests whether this is within a transaction.
func (tbl AUserTable) IsTx() bool {
	_, ok := tbl.db.(*sql.Tx)
	return ok
}

// BeginTx starts a transaction using the table's context.
// This context is used until the transaction is committed or rolled back.
//
// If this context is cancelled, the sql package will roll back the transaction.
// In this case, Tx.Commit will then return an error.
//
// The provided TxOptions is optional and may be nil if defaults should be used.
// If a non-default isolation level is used that the driver doesn't support,
// an error will be returned.
//
// Panics if the Execer is not TxStarter.
func (tbl AUserTable) BeginTx(opts *sql.TxOptions) (AUserTable, error) {
	var err error
	tbl.db, err = tbl.db.(sqlgen2.TxStarter).BeginTx(tbl.ctx, opts)
	return tbl, tbl.logIfError(err)
}

// Using returns a modified Table using the transaction supplied. This is needed
// when making multiple queries across several tables within a single transaction.
// The result is a modified copy of the table; the original is unchanged.
func (tbl AUserTable) Using(tx *sql.Tx) AUserTable {
	tbl.db = tx
	return tbl
}

func (tbl AUserTable) logQuery(query string, args ...interface{}) {
	tbl.database.LogQuery(query, args...)
}

func (tbl AUserTable) logError(err error) error {
	return tbl.database.LogError(err)
}

func (tbl AUserTable) logIfError(err error) error {
	return tbl.database.LogIfError(err)
}


//--------------------------------------------------------------------------------

const NumAUserColumns = 22

const NumAUserDataColumns = 21

const AUserColumnNames = "uid,name,emailaddress,addressid,avatar,role,active,admin,fave,lastupdated,i8,u8,i16,u16,i32,u32,i64,u64,f32,f64,token,secret"

const AUserDataColumnNames = "name,emailaddress,addressid,avatar,role,active,admin,fave,lastupdated,i8,u8,i16,u16,i32,u32,i64,u64,f32,f64,token,secret"

//--------------------------------------------------------------------------------

const sqlAUserTableCreateColumnsSqlite = "\n"+
" `uid`          integer not null primary key autoincrement,\n"+
" `name`         text not null,\n"+
" `emailaddress` text not null,\n"+
" `addressid`    bigint default null,\n"+
" `avatar`       text default null,\n"+
" `role`         text default null,\n"+
" `active`       boolean not null,\n"+
" `admin`        boolean not null,\n"+
" `fave`         text,\n"+
" `lastupdated`  bigint not null,\n"+
" `i8`           tinyint not null default -8,\n"+
" `u8`           tinyint unsigned not null default 8,\n"+
" `i16`          smallint not null default -16,\n"+
" `u16`          smallint unsigned not null default 16,\n"+
" `i32`          int not null default -32,\n"+
" `u32`          int unsigned not null default 32,\n"+
" `i64`          bigint not null default -64,\n"+
" `u64`          bigint unsigned not null default 64,\n"+
" `f32`          float not null default 3.2,\n"+
" `f64`          double not null default 6.4,\n"+
" `token`        text not null,\n"+
" `secret`       text not null"

const sqlAUserTableCreateColumnsMysql = "\n"+
" `uid`          bigint not null primary key auto_increment,\n"+
" `name`         varchar(255) not null,\n"+
" `emailaddress` varchar(255) not null,\n"+
" `addressid`    bigint default null,\n"+
" `avatar`       varchar(255) default null,\n"+
" `role`         varchar(20) default null,\n"+
" `active`       tinyint(1) not null,\n"+
" `admin`        tinyint(1) not null,\n"+
" `fave`         json,\n"+
" `lastupdated`  bigint not null,\n"+
" `i8`           tinyint not null default -8,\n"+
" `u8`           tinyint unsigned not null default 8,\n"+
" `i16`          smallint not null default -16,\n"+
" `u16`          smallint unsigned not null default 16,\n"+
" `i32`          int not null default -32,\n"+
" `u32`          int unsigned not null default 32,\n"+
" `i64`          bigint not null default -64,\n"+
" `u64`          bigint unsigned not null default 64,\n"+
" `f32`          float not null default 3.2,\n"+
" `f64`          double not null default 6.4,\n"+
" `token`        varchar(255) not null,\n"+
" `secret`       varchar(255) not null"

const sqlAUserTableCreateColumnsPostgres = `
 "uid"          bigserial not null primary key,
 "name"         varchar(255) not null,
 "emailaddress" varchar(255) not null,
 "addressid"    bigint default null,
 "avatar"       varchar(255) default null,
 "role"         varchar(20) default null,
 "active"       boolean not null,
 "admin"        boolean not null,
 "fave"         json,
 "lastupdated"  bigint not null,
 "i8"           int8 not null default -8,
 "u8"           smallint not null default 8,
 "i16"          smallint not null default -16,
 "u16"          integer not null default 16,
 "i32"          integer not null default -32,
 "u32"          bigint not null default 32,
 "i64"          bigint not null default -64,
 "u64"          bigint not null default 64,
 "f32"          real not null default 3.2,
 "f64"          double precision not null default 6.4,
 "token"        varchar(255) not null,
 "secret"       varchar(255) not null`

const sqlConstrainAUserTable = `
 CONSTRAINT AUserc3 foreign key (addressid) references %saddresses (id) on update restrict on delete restrict
`

//--------------------------------------------------------------------------------

const sqlAEmailaddressIdxIndexColumns = "emailaddress"

const sqlAUserLoginIndexColumns = "name"

//--------------------------------------------------------------------------------

// CreateTable creates the table.
func (tbl AUserTable) CreateTable(ifNotExists bool) (int64, error) {
	return support.Exec(tbl, nil, tbl.createTableSql(ifNotExists))
}

func (tbl AUserTable) createTableSql(ifNotExists bool) string {
	var columns string
	var settings string
	switch tbl.Dialect() {
	case schema.Sqlite:
		columns = sqlAUserTableCreateColumnsSqlite
		settings = ""
    case schema.Mysql:
		columns = sqlAUserTableCreateColumnsMysql
		settings = " ENGINE=InnoDB DEFAULT CHARSET=utf8"
    case schema.Postgres:
		columns = sqlAUserTableCreateColumnsPostgres
		settings = ""
    }
	buf := &bytes.Buffer{}
	buf.WriteString("CREATE TABLE ")
	if ifNotExists {
		buf.WriteString("IF NOT EXISTS ")
	}
	buf.WriteString(tbl.name.String())
	buf.WriteString(" (")
	buf.WriteString(columns)
	for i, c := range tbl.constraints {
		buf.WriteString(",\n ")
		buf.WriteString(c.ConstraintSql(tbl.Dialect(), tbl.name, i+1))
	}
	buf.WriteString("\n)")
	buf.WriteString(settings)
	return buf.String()
}

func (tbl AUserTable) ternary(flag bool, a, b string) string {
	if flag {
		return a
	}
	return b
}

// DropTable drops the table, destroying all its data.
func (tbl AUserTable) DropTable(ifExists bool) (int64, error) {
	return support.Exec(tbl, nil, tbl.dropTableSql(ifExists))
}

func (tbl AUserTable) dropTableSql(ifExists bool) string {
	ie := tbl.ternary(ifExists, "IF EXISTS ", "")
	query := fmt.Sprintf("DROP TABLE %s%s", ie, tbl.name)
	return query
}

//--------------------------------------------------------------------------------

// CreateTableWithIndexes invokes CreateTable then CreateIndexes.
func (tbl AUserTable) CreateTableWithIndexes(ifNotExist bool) (err error) {
	_, err = tbl.CreateTable(ifNotExist)
	if err != nil {
		return err
	}

	return tbl.CreateIndexes(ifNotExist)
}

// CreateIndexes executes queries that create the indexes needed by the User table.
func (tbl AUserTable) CreateIndexes(ifNotExist bool) (err error) {

	err = tbl.CreateEmailaddressIdxIndex(ifNotExist)
	if err != nil {
		return err
	}

	err = tbl.CreateUserLoginIndex(ifNotExist)
	if err != nil {
		return err
	}

	return nil
}

// CreateEmailaddressIdxIndex creates the emailaddress_idx index.
func (tbl AUserTable) CreateEmailaddressIdxIndex(ifNotExist bool) error {
	ine := tbl.ternary(ifNotExist && tbl.Dialect() != schema.Mysql, "IF NOT EXISTS ", "")

	// Mysql does not support 'if not exists' on indexes
	// Workaround: use DropIndex first and ignore an error returned if the index didn't exist.

	if ifNotExist && tbl.Dialect() == schema.Mysql {
		// low-level no-logging Exec
		tbl.Execer().ExecContext(tbl.ctx, tbl.dropAEmailaddressIdxIndexSql(false))
		ine = ""
	}

	_, err := tbl.Exec(nil, tbl.createAEmailaddressIdxIndexSql(ine))
	return err
}

func (tbl AUserTable) createAEmailaddressIdxIndexSql(ifNotExists string) string {
	indexPrefix := tbl.name.PrefixWithoutDot()
	return fmt.Sprintf("CREATE UNIQUE INDEX %s%semailaddress_idx ON %s (%s)", ifNotExists, indexPrefix,
		tbl.name, sqlAEmailaddressIdxIndexColumns)
}

// DropEmailaddressIdxIndex drops the emailaddress_idx index.
func (tbl AUserTable) DropEmailaddressIdxIndex(ifExists bool) error {
	_, err := tbl.Exec(nil, tbl.dropAEmailaddressIdxIndexSql(ifExists))
	return err
}

func (tbl AUserTable) dropAEmailaddressIdxIndexSql(ifExists bool) string {
	// Mysql does not support 'if exists' on indexes
	ie := tbl.ternary(ifExists && tbl.Dialect() != schema.Mysql, "IF EXISTS ", "")
	onTbl := tbl.ternary(tbl.Dialect() == schema.Mysql, fmt.Sprintf(" ON %s", tbl.name), "")
	indexPrefix := tbl.name.PrefixWithoutDot()
	return fmt.Sprintf("DROP INDEX %s%semailaddress_idx%s", ie, indexPrefix, onTbl)
}

// CreateUserLoginIndex creates the user_login index.
func (tbl AUserTable) CreateUserLoginIndex(ifNotExist bool) error {
	ine := tbl.ternary(ifNotExist && tbl.Dialect() != schema.Mysql, "IF NOT EXISTS ", "")

	// Mysql does not support 'if not exists' on indexes
	// Workaround: use DropIndex first and ignore an error returned if the index didn't exist.

	if ifNotExist && tbl.Dialect() == schema.Mysql {
		// low-level no-logging Exec
		tbl.Execer().ExecContext(tbl.ctx, tbl.dropAUserLoginIndexSql(false))
		ine = ""
	}

	_, err := tbl.Exec(nil, tbl.createAUserLoginIndexSql(ine))
	return err
}

func (tbl AUserTable) createAUserLoginIndexSql(ifNotExists string) string {
	indexPrefix := tbl.name.PrefixWithoutDot()
	return fmt.Sprintf("CREATE UNIQUE INDEX %s%suser_login ON %s (%s)", ifNotExists, indexPrefix,
		tbl.name, sqlAUserLoginIndexColumns)
}

// DropUserLoginIndex drops the user_login index.
func (tbl AUserTable) DropUserLoginIndex(ifExists bool) error {
	_, err := tbl.Exec(nil, tbl.dropAUserLoginIndexSql(ifExists))
	return err
}

func (tbl AUserTable) dropAUserLoginIndexSql(ifExists bool) string {
	// Mysql does not support 'if exists' on indexes
	ie := tbl.ternary(ifExists && tbl.Dialect() != schema.Mysql, "IF EXISTS ", "")
	onTbl := tbl.ternary(tbl.Dialect() == schema.Mysql, fmt.Sprintf(" ON %s", tbl.name), "")
	indexPrefix := tbl.name.PrefixWithoutDot()
	return fmt.Sprintf("DROP INDEX %s%suser_login%s", ie, indexPrefix, onTbl)
}

// DropIndexes executes queries that drop the indexes on by the User table.
func (tbl AUserTable) DropIndexes(ifExist bool) (err error) {

	err = tbl.DropEmailaddressIdxIndex(ifExist)
	if err != nil {
		return err
	}

	err = tbl.DropUserLoginIndex(ifExist)
	if err != nil {
		return err
	}

	return nil
}

//--------------------------------------------------------------------------------

// Truncate drops every record from the table, if possible. It might fail if constraints exist that
// prevent some or all rows from being deleted; use the force option to override this.
//
// When 'force' is set true, be aware of the following consequences.
// When using Mysql, foreign keys in other tables can be left dangling.
// When using Postgres, a cascade happens, so all 'adjacent' tables (i.e. linked by foreign keys)
// are also truncated.
func (tbl AUserTable) Truncate(force bool) (err error) {
	for _, query := range tbl.Dialect().TruncateDDL(tbl.Name().String(), force) {
		_, err = support.Exec(tbl, nil, query)
		if err != nil {
			return err
		}
	}
	return nil
}

//--------------------------------------------------------------------------------

// Exec executes a query without returning any rows.
// It returns the number of rows affected (if the database driver supports this).
//
// The args are for any placeholder parameters in the query.
func (tbl AUserTable) Exec(req require.Requirement, query string, args ...interface{}) (int64, error) {
	return support.Exec(tbl, req, query, args...)
}

//--------------------------------------------------------------------------------

// Query is the low-level request method for this table. The query is logged using whatever logger is
// configured. If an error arises, this too is logged.
//
// If you need a context other than the background, use WithContext before calling Query.
//
// The args are for any placeholder parameters in the query.
//
// The caller must call rows.Close() on the result.
//
// Wrap the result in *sqlgen2.Rows if you need to access its data as a map.
func (tbl AUserTable) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return support.Query(tbl, query, args...)
}

//--------------------------------------------------------------------------------

// QueryOneNullString is a low-level access method for one string. This can be used for function queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl AUserTable) QueryOneNullString(req require.Requirement, query string, args ...interface{}) (result sql.NullString, err error) {
	err = support.QueryOneNullThing(tbl, req, &result, query, args...)
	return result, err
}

// QueryOneNullInt64 is a low-level access method for one int64. This can be used for 'COUNT(1)' queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl AUserTable) QueryOneNullInt64(req require.Requirement, query string, args ...interface{}) (result sql.NullInt64, err error) {
	err = support.QueryOneNullThing(tbl, req, &result, query, args...)
	return result, err
}

// QueryOneNullFloat64 is a low-level access method for one float64. This can be used for 'AVG(...)' queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl AUserTable) QueryOneNullFloat64(req require.Requirement, query string, args ...interface{}) (result sql.NullFloat64, err error) {
	err = support.QueryOneNullThing(tbl, req, &result, query, args...)
	return result, err
}

func scanAUsers(rows *sql.Rows, firstOnly bool) (vv []*User, n int64, err error) {
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
		var v10 int8
		var v11 uint8
		var v12 int16
		var v13 uint16
		var v14 int32
		var v15 uint32
		var v16 int64
		var v17 uint64
		var v18 float32
		var v19 float64
		var v20 string
		var v21 string

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
			&v12,
			&v13,
			&v14,
			&v15,
			&v16,
			&v17,
			&v18,
			&v19,
			&v20,
			&v21,
		)
		if err != nil {
			return vv, n, err
		}

		v := &User{}
		v.Uid = v0
		v.Name = v1
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
		v.Numbers.I8 = v10
		v.Numbers.U8 = v11
		v.Numbers.I16 = v12
		v.Numbers.U16 = v13
		v.Numbers.I32 = v14
		v.Numbers.U32 = v15
		v.Numbers.I64 = v16
		v.Numbers.U64 = v17
		v.Numbers.F32 = v18
		v.Numbers.F64 = v19
		v.token = v20
		v.secret = v21

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

var allAUserQuotedColumnNames = []string{
	schema.Sqlite.SplitAndQuote(AUserColumnNames),
	schema.Mysql.SplitAndQuote(AUserColumnNames),
	schema.Postgres.SplitAndQuote(AUserColumnNames),
}

//--------------------------------------------------------------------------------

// GetUsersByUid gets records from the table according to a list of primary keys.
// Although the list of ids can be arbitrarily long, there are practical limits;
// note that Oracle DB has a limit of 1000.
//
// It places a requirement, which may be nil, on the size of the expected results: in particular, require.All
// controls whether an error is generated not all the ids produce a result.
func (tbl AUserTable) GetUsersByUid(req require.Requirement, id ...int64) (list []*User, err error) {
	if len(id) > 0 {
		if req == require.All {
			req = require.Exactly(len(id))
		}
		args := make([]interface{}, len(id))

		for i, v := range id {
			args[i] = v
		}

		list, err = tbl.getUsers(req, tbl.pk, args...)
	}

	return list, err
}

// GetUserByUid gets the record with a given primary key value.
// If not found, *User will be nil.
func (tbl AUserTable) GetUserByUid(req require.Requirement, id int64) (*User, error) {
	return tbl.getUser(req, tbl.pk, id)
}

// GetUserByEmailAddress gets the record with a given emailaddress value.
// If not found, *User will be nil.
func (tbl AUserTable) GetUserByEmailAddress(req require.Requirement, emailaddress string) (*User, error) {
	return tbl.SelectOne(req, where.And(where.Eq("emailaddress", emailaddress)), nil)
}

// GetUserByName gets the record with a given name value.
// If not found, *User will be nil.
func (tbl AUserTable) GetUserByName(req require.Requirement, name string) (*User, error) {
	return tbl.SelectOne(req, where.And(where.Eq("name", name)), nil)
}

func (tbl AUserTable) getUser(req require.Requirement, column string, arg interface{}) (*User, error) {
	dialect := tbl.Dialect()
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s=%s",
		allAUserQuotedColumnNames[dialect.Index()], tbl.name, dialect.Quote(column), dialect.Placeholder(column, 1))
	v, err := tbl.doQueryOne(req, query, arg)
	return v, err
}

func (tbl AUserTable) getUsers(req require.Requirement, column string, args ...interface{}) (list []*User, err error) {
	if len(args) > 0 {
		if req == require.All {
			req = require.Exactly(len(args))
		}
		dialect := tbl.Dialect()
		pl := dialect.Placeholders(len(args))
		query := fmt.Sprintf("SELECT %s FROM %s WHERE %s IN (%s)",
			allAUserQuotedColumnNames[dialect.Index()], tbl.name, dialect.Quote(column), pl)
		list, err = tbl.doQuery(req, false, query, args...)
	}

	return list, err
}

func (tbl AUserTable) doQueryOne(req require.Requirement, query string, args ...interface{}) (*User, error) {
	list, err := tbl.doQuery(req, true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

func (tbl AUserTable) doQuery(req require.Requirement, firstOnly bool, query string, args ...interface{}) ([]*User, error) {
	rows, err := tbl.Query(query, args...)
	if err != nil {
		return nil, err
	}

	vv, n, err := scanAUsers(rows, firstOnly)
	return vv, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(err, req, n))
}

// Fetch fetches a list of User based on a supplied query. This is mostly used for join queries that map its
// result columns to the fields of User. Other queries might be better handled by GetXxx or Select methods.
func (tbl AUserTable) Fetch(req require.Requirement, query string, args ...interface{}) ([]*User, error) {
	return tbl.doQuery(req, false, query, args...)
}

//--------------------------------------------------------------------------------

// SelectOneWhere allows a single Example to be obtained from the table that match a 'where' clause
// and some limit. Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
// If not found, *Example will be nil.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.One
// controls whether an error is generated when no result is found.
//
// The args are for any placeholder parameters in the query.
func (tbl AUserTable) SelectOneWhere(req require.Requirement, where, orderBy string, args ...interface{}) (*User, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT 1",
		allAUserQuotedColumnNames[tbl.Dialect().Index()], tbl.name, where, orderBy)
	v, err := tbl.doQueryOne(req, query, args...)
	return v, err
}

// SelectOne allows a single User to be obtained from the sqlgen2.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
// If not found, *Example will be nil.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.One
// controls whether an error is generated when no result is found.
func (tbl AUserTable) SelectOne(req require.Requirement, wh where.Expression, qc where.QueryConstraint) (*User, error) {
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
	return tbl.SelectOneWhere(req, whs, orderBy, args...)
}

// SelectWhere allows Users to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
//
// The args are for any placeholder parameters in the query.
func (tbl AUserTable) SelectWhere(req require.Requirement, where, orderBy string, args ...interface{}) ([]*User, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s %s",
		allAUserQuotedColumnNames[tbl.Dialect().Index()], tbl.name, where, orderBy)
	vv, err := tbl.doQuery(req, false, query, args...)
	return vv, err
}

// Select allows Users to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
func (tbl AUserTable) Select(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]*User, error) {
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
	return tbl.SelectWhere(req, whs, orderBy, args...)
}

// CountWhere counts Users in the table that match a 'where' clause.
// Use a blank string for the 'where' argument if it is not needed.
//
// The args are for any placeholder parameters in the query.
func (tbl AUserTable) CountWhere(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s %s", tbl.name, where)
	tbl.logQuery(query, args...)
	row := tbl.db.QueryRowContext(tbl.ctx, query, args...)
	err = row.Scan(&count)
	return count, tbl.logIfError(err)
}

// Count counts the Users in the table that match a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed.
func (tbl AUserTable) Count(wh where.Expression) (count int64, err error) {
	whs, args := where.BuildExpression(wh, tbl.Dialect())
	return tbl.CountWhere(whs, args...)
}

//--------------------------------------------------------------------------------

// SliceUid gets the uid column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AUserTable) SliceUid(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return tbl.sliceInt64List(req, tbl.pk, wh, qc)
}

// SliceName gets the name column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AUserTable) SliceName(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.sliceStringList(req, "name", wh, qc)
}

// SliceEmailaddress gets the emailaddress column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AUserTable) SliceEmailaddress(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.sliceStringList(req, "emailaddress", wh, qc)
}

// SliceAddressid gets the addressid column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AUserTable) SliceAddressid(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return tbl.sliceInt64PtrList(req, "addressid", wh, qc)
}

// SliceAvatar gets the avatar column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AUserTable) SliceAvatar(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.sliceStringPtrList(req, "avatar", wh, qc)
}

// SliceRole gets the role column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AUserTable) SliceRole(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]Role, error) {
	return tbl.sliceRolePtrList(req, "role", wh, qc)
}

// SliceLastupdated gets the lastupdated column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AUserTable) SliceLastupdated(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return tbl.sliceInt64List(req, "lastupdated", wh, qc)
}

// SliceI8 gets the i8 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AUserTable) SliceI8(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int8, error) {
	return tbl.sliceInt8List(req, "i8", wh, qc)
}

// SliceU8 gets the u8 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AUserTable) SliceU8(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]uint8, error) {
	return tbl.sliceUint8List(req, "u8", wh, qc)
}

// SliceI16 gets the i16 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AUserTable) SliceI16(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int16, error) {
	return tbl.sliceInt16List(req, "i16", wh, qc)
}

// SliceU16 gets the u16 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AUserTable) SliceU16(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]uint16, error) {
	return tbl.sliceUint16List(req, "u16", wh, qc)
}

// SliceI32 gets the i32 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AUserTable) SliceI32(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int32, error) {
	return tbl.sliceInt32List(req, "i32", wh, qc)
}

// SliceU32 gets the u32 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AUserTable) SliceU32(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]uint32, error) {
	return tbl.sliceUint32List(req, "u32", wh, qc)
}

// SliceI64 gets the i64 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AUserTable) SliceI64(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return tbl.sliceInt64List(req, "i64", wh, qc)
}

// SliceU64 gets the u64 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AUserTable) SliceU64(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]uint64, error) {
	return tbl.sliceUint64List(req, "u64", wh, qc)
}

// SliceF32 gets the f32 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AUserTable) SliceF32(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]float32, error) {
	return tbl.sliceFloat32List(req, "f32", wh, qc)
}

// SliceF64 gets the f64 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AUserTable) SliceF64(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]float64, error) {
	return tbl.sliceFloat64List(req, "f64", wh, qc)
}

func (tbl AUserTable) sliceRolePtrList(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]Role, error) {
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", dialect.Quote(sqlname), tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, tbl.logError(err)
	}
	defer rows.Close()

	var v Role
	list := make([]Role, 0, 10)

	for rows.Next() {
		err = rows.Scan(&v)
		if err == sql.ErrNoRows {
			return list, tbl.logIfError(require.ErrorIfQueryNotSatisfiedBy(req, int64(len(list))))
		} else {
			list = append(list, v)
		}
	}
	return list, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, int64(len(list))))
}

func (tbl AUserTable) sliceFloat32List(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]float32, error) {
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", dialect.Quote(sqlname), tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, tbl.logError(err)
	}
	defer rows.Close()

	var v float32
	list := make([]float32, 0, 10)

	for rows.Next() {
		err = rows.Scan(&v)
		if err == sql.ErrNoRows {
			return list, tbl.logIfError(require.ErrorIfQueryNotSatisfiedBy(req, int64(len(list))))
		} else {
			list = append(list, v)
		}
	}
	return list, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, int64(len(list))))
}

func (tbl AUserTable) sliceFloat64List(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]float64, error) {
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", dialect.Quote(sqlname), tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, tbl.logError(err)
	}
	defer rows.Close()

	var v float64
	list := make([]float64, 0, 10)

	for rows.Next() {
		err = rows.Scan(&v)
		if err == sql.ErrNoRows {
			return list, tbl.logIfError(require.ErrorIfQueryNotSatisfiedBy(req, int64(len(list))))
		} else {
			list = append(list, v)
		}
	}
	return list, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, int64(len(list))))
}

func (tbl AUserTable) sliceInt16List(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]int16, error) {
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", dialect.Quote(sqlname), tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, tbl.logError(err)
	}
	defer rows.Close()

	var v int16
	list := make([]int16, 0, 10)

	for rows.Next() {
		err = rows.Scan(&v)
		if err == sql.ErrNoRows {
			return list, tbl.logIfError(require.ErrorIfQueryNotSatisfiedBy(req, int64(len(list))))
		} else {
			list = append(list, v)
		}
	}
	return list, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, int64(len(list))))
}

func (tbl AUserTable) sliceInt32List(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]int32, error) {
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", dialect.Quote(sqlname), tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, tbl.logError(err)
	}
	defer rows.Close()

	var v int32
	list := make([]int32, 0, 10)

	for rows.Next() {
		err = rows.Scan(&v)
		if err == sql.ErrNoRows {
			return list, tbl.logIfError(require.ErrorIfQueryNotSatisfiedBy(req, int64(len(list))))
		} else {
			list = append(list, v)
		}
	}
	return list, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, int64(len(list))))
}

func (tbl AUserTable) sliceInt64List(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", dialect.Quote(sqlname), tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, tbl.logError(err)
	}
	defer rows.Close()

	var v int64
	list := make([]int64, 0, 10)

	for rows.Next() {
		err = rows.Scan(&v)
		if err == sql.ErrNoRows {
			return list, tbl.logIfError(require.ErrorIfQueryNotSatisfiedBy(req, int64(len(list))))
		} else {
			list = append(list, v)
		}
	}
	return list, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, int64(len(list))))
}

func (tbl AUserTable) sliceInt64PtrList(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", dialect.Quote(sqlname), tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, tbl.logError(err)
	}
	defer rows.Close()

	var v int64
	list := make([]int64, 0, 10)

	for rows.Next() {
		err = rows.Scan(&v)
		if err == sql.ErrNoRows {
			return list, tbl.logIfError(require.ErrorIfQueryNotSatisfiedBy(req, int64(len(list))))
		} else {
			list = append(list, v)
		}
	}
	return list, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, int64(len(list))))
}

func (tbl AUserTable) sliceInt8List(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]int8, error) {
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", dialect.Quote(sqlname), tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, tbl.logError(err)
	}
	defer rows.Close()

	var v int8
	list := make([]int8, 0, 10)

	for rows.Next() {
		err = rows.Scan(&v)
		if err == sql.ErrNoRows {
			return list, tbl.logIfError(require.ErrorIfQueryNotSatisfiedBy(req, int64(len(list))))
		} else {
			list = append(list, v)
		}
	}
	return list, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, int64(len(list))))
}

func (tbl AUserTable) sliceStringList(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", dialect.Quote(sqlname), tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, tbl.logError(err)
	}
	defer rows.Close()

	var v string
	list := make([]string, 0, 10)

	for rows.Next() {
		err = rows.Scan(&v)
		if err == sql.ErrNoRows {
			return list, tbl.logIfError(require.ErrorIfQueryNotSatisfiedBy(req, int64(len(list))))
		} else {
			list = append(list, v)
		}
	}
	return list, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, int64(len(list))))
}

func (tbl AUserTable) sliceStringPtrList(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", dialect.Quote(sqlname), tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, tbl.logError(err)
	}
	defer rows.Close()

	var v string
	list := make([]string, 0, 10)

	for rows.Next() {
		err = rows.Scan(&v)
		if err == sql.ErrNoRows {
			return list, tbl.logIfError(require.ErrorIfQueryNotSatisfiedBy(req, int64(len(list))))
		} else {
			list = append(list, v)
		}
	}
	return list, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, int64(len(list))))
}

func (tbl AUserTable) sliceUint16List(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]uint16, error) {
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", dialect.Quote(sqlname), tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, tbl.logError(err)
	}
	defer rows.Close()

	var v uint16
	list := make([]uint16, 0, 10)

	for rows.Next() {
		err = rows.Scan(&v)
		if err == sql.ErrNoRows {
			return list, tbl.logIfError(require.ErrorIfQueryNotSatisfiedBy(req, int64(len(list))))
		} else {
			list = append(list, v)
		}
	}
	return list, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, int64(len(list))))
}

func (tbl AUserTable) sliceUint32List(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]uint32, error) {
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", dialect.Quote(sqlname), tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, tbl.logError(err)
	}
	defer rows.Close()

	var v uint32
	list := make([]uint32, 0, 10)

	for rows.Next() {
		err = rows.Scan(&v)
		if err == sql.ErrNoRows {
			return list, tbl.logIfError(require.ErrorIfQueryNotSatisfiedBy(req, int64(len(list))))
		} else {
			list = append(list, v)
		}
	}
	return list, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, int64(len(list))))
}

func (tbl AUserTable) sliceUint64List(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]uint64, error) {
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", dialect.Quote(sqlname), tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, tbl.logError(err)
	}
	defer rows.Close()

	var v uint64
	list := make([]uint64, 0, 10)

	for rows.Next() {
		err = rows.Scan(&v)
		if err == sql.ErrNoRows {
			return list, tbl.logIfError(require.ErrorIfQueryNotSatisfiedBy(req, int64(len(list))))
		} else {
			list = append(list, v)
		}
	}
	return list, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, int64(len(list))))
}

func (tbl AUserTable) sliceUint8List(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]uint8, error) {
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", dialect.Quote(sqlname), tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, tbl.logError(err)
	}
	defer rows.Close()

	var v uint8
	list := make([]uint8, 0, 10)

	for rows.Next() {
		err = rows.Scan(&v)
		if err == sql.ErrNoRows {
			return list, tbl.logIfError(require.ErrorIfQueryNotSatisfiedBy(req, int64(len(list))))
		} else {
			list = append(list, v)
		}
	}
	return list, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, int64(len(list))))
}


func constructAUserInsert(w io.Writer, v *User, dialect schema.Dialect, withPk bool) (s []interface{}, err error) {
	s = make([]interface{}, 0, 22)

	comma := ""
	io.WriteString(w, " (")

	if withPk {
		dialect.QuoteW(w, "uid")
		comma = ","
		s = append(s, v.Uid)
	}

	io.WriteString(w, comma)

	dialect.QuoteW(w, "name")
	s = append(s, v.Name)
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

	dialect.QuoteW(w, "i8")
	s = append(s, v.Numbers.I8)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "u8")
	s = append(s, v.Numbers.U8)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "i16")
	s = append(s, v.Numbers.I16)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "u16")
	s = append(s, v.Numbers.U16)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "i32")
	s = append(s, v.Numbers.I32)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "u32")
	s = append(s, v.Numbers.U32)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "i64")
	s = append(s, v.Numbers.I64)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "u64")
	s = append(s, v.Numbers.U64)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "f32")
	s = append(s, v.Numbers.F32)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "f64")
	s = append(s, v.Numbers.F64)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "token")
	s = append(s, v.token)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "secret")
	s = append(s, v.secret)
	io.WriteString(w, ")")
	return s, nil
}

func constructAUserUpdate(w io.Writer, v *User, dialect schema.Dialect) (s []interface{}, err error) {
	j := 1
	s = make([]interface{}, 0, 21)

	comma := ""

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "name", j)
	s = append(s, v.Name)
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
	dialect.QuoteWithPlaceholder(w, "i8", j)
	s = append(s, v.Numbers.I8)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "u8", j)
	s = append(s, v.Numbers.U8)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "i16", j)
	s = append(s, v.Numbers.I16)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "u16", j)
	s = append(s, v.Numbers.U16)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "i32", j)
	s = append(s, v.Numbers.I32)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "u32", j)
	s = append(s, v.Numbers.U32)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "i64", j)
	s = append(s, v.Numbers.I64)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "u64", j)
	s = append(s, v.Numbers.U64)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "f32", j)
	s = append(s, v.Numbers.F32)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "f64", j)
	s = append(s, v.Numbers.F64)
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

//--------------------------------------------------------------------------------

// Insert adds new records for the Users.
// The Users have their primary key fields set to the new record identifiers.
// The User.PreInsert() method will be called, if it exists.
func (tbl AUserTable) Insert(req require.Requirement, vv ...*User) error {
	if req == require.All {
		req = require.Exactly(len(vv))
	}

	var count int64
	//columns := allXExampleQuotedInserts[tbl.Dialect().Index()]
	//query := fmt.Sprintf("INSERT INTO %s %s", tbl.name, columns)
	//st, err := tbl.db.PrepareContext(tbl.ctx, query)
	//if err != nil {
	//	return err
	//}
	//defer st.Close()

	insertHasReturningPhrase := tbl.Dialect().InsertHasReturningPhrase()
	returning := ""
	if tbl.Dialect().InsertHasReturningPhrase() {
		returning = fmt.Sprintf(" returning %q", tbl.pk)
	}

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

		fields, err := constructAUserInsert(b, v, tbl.Dialect(), false)
		if err != nil {
			return tbl.logError(err)
		}

		io.WriteString(b, " VALUES (")
		io.WriteString(b, tbl.Dialect().Placeholders(len(fields)))
		io.WriteString(b, ")")
		io.WriteString(b, returning)

		query := b.String()
		tbl.logQuery(query, fields...)

		var n int64 = 1
		if insertHasReturningPhrase {
			row := tbl.db.QueryRowContext(tbl.ctx, query, fields...)
			err = row.Scan(&v.Uid)

		} else {
			res, e2 := tbl.db.ExecContext(tbl.ctx, query, fields...)
			if e2 != nil {
				return tbl.logError(e2)
			}

			v.Uid, err = res.LastInsertId()
			if e2 != nil {
				return tbl.logError(e2)
			}
	
			n, err = res.RowsAffected()
		}

		if err != nil {
			return tbl.logError(err)
		}
		count += n
	}

	return tbl.logIfError(require.ErrorIfExecNotSatisfiedBy(req, count))
}

// UpdateFields updates one or more columns, given a 'where' clause.
//
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl AUserTable) UpdateFields(req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error) {
	return support.UpdateFields(tbl, req, wh, fields...)
}

//--------------------------------------------------------------------------------

var allAUserQuotedUpdates = []string{
	// Sqlite
	"`name`=?,`emailaddress`=?,`addressid`=?,`avatar`=?,`role`=?,`active`=?,`admin`=?,`fave`=?,`lastupdated`=?,`i8`=?,`u8`=?,`i16`=?,`u16`=?,`i32`=?,`u32`=?,`i64`=?,`u64`=?,`f32`=?,`f64`=?,`token`=?,`secret`=? WHERE `uid`=?",
	// Mysql
	"`name`=?,`emailaddress`=?,`addressid`=?,`avatar`=?,`role`=?,`active`=?,`admin`=?,`fave`=?,`lastupdated`=?,`i8`=?,`u8`=?,`i16`=?,`u16`=?,`i32`=?,`u32`=?,`i64`=?,`u64`=?,`f32`=?,`f64`=?,`token`=?,`secret`=? WHERE `uid`=?",
	// Postgres
	`"name"=$2,"emailaddress"=$3,"addressid"=$4,"avatar"=$5,"role"=$6,"active"=$7,"admin"=$8,"fave"=$9,"lastupdated"=$10,"i8"=$11,"u8"=$12,"i16"=$13,"u16"=$14,"i32"=$15,"u32"=$16,"i64"=$17,"u64"=$18,"f32"=$19,"f64"=$20,"token"=$21,"secret"=$22 WHERE "uid"=$1`,
}

//--------------------------------------------------------------------------------

// Update updates records, matching them by primary key. It returns the number of rows affected.
// The User.PreUpdate(Execer) method will be called, if it exists.
func (tbl AUserTable) Update(req require.Requirement, vv ...*User) (int64, error) {
	if req == require.All {
		req = require.Exactly(len(vv))
	}

	var count int64
	dialect := tbl.Dialect()
	//columns := allAUserQuotedUpdates[dialect.Index()]
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

		args, err := constructAUserUpdate(b, v, dialect)
		k := len(args) + 1
		args = append(args, v.Uid)
		if err != nil {
			return count, tbl.logError(err)
		}

		io.WriteString(b, " WHERE ")
		dialect.QuoteWithPlaceholder(b, tbl.pk, k)

		query := b.String()
		n, err := tbl.Exec(nil, query, args...)
		if err != nil {
			return count, err
		}
		count += n
	}

	return count, tbl.logIfError(require.ErrorIfExecNotSatisfiedBy(req, count))
}

//--------------------------------------------------------------------------------

// DeleteUsers deletes rows from the table, given some primary keys.
// The list of ids can be arbitrarily long.
func (tbl AUserTable) DeleteUsers(req require.Requirement, id ...int64) (int64, error) {
	const batch = 1000 // limited by Oracle DB
	const qt = "DELETE FROM %s WHERE %s IN (%s)"

	if req == require.All {
		req = require.Exactly(len(id))
	}

	var count, n int64
	var err error
	var max = batch
	if len(id) < batch {
		max = len(id)
	}
	dialect := tbl.Dialect()
	col := dialect.Quote(tbl.pk)
	args := make([]interface{}, max)

	if len(id) > batch {
		pl := dialect.Placeholders(batch)
		query := fmt.Sprintf(qt, tbl.name, col, pl)

		for len(id) > batch {
			for i := 0; i < batch; i++ {
				args[i] = id[i]
			}

			n, err = tbl.Exec(nil, query, args...)
			count += n
			if err != nil {
				return count, err
			}

			id = id[batch:]
		}
	}

	if len(id) > 0 {
		pl := dialect.Placeholders(len(id))
		query := fmt.Sprintf(qt, tbl.name, col, pl)

		for i := 0; i < len(id); i++ {
			args[i] = id[i]
		}

		n, err = tbl.Exec(nil, query, args...)
		count += n
	}

	return count, tbl.logIfError(require.ChainErrorIfExecNotSatisfiedBy(err, req, n))
}

// Delete deletes one or more rows from the table, given a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl AUserTable) Delete(req require.Requirement, wh where.Expression) (int64, error) {
	query, args := tbl.deleteRows(wh)
	return tbl.Exec(req, query, args...)
}

func (tbl AUserTable) deleteRows(wh where.Expression) (string, []interface{}) {
	whs, args := where.BuildExpression(wh, tbl.Dialect())
	query := fmt.Sprintf("DELETE FROM %s %s", tbl.name, whs)
	return query, args
}

//--------------------------------------------------------------------------------
