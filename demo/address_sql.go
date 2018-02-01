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
	"strings"
)

// AddressTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type AddressTable struct {
	name        sqlgen2.TableName
	database    *sqlgen2.Database
	db          sqlgen2.Execer
	constraints constraint.Constraints
	ctx			context.Context
}

// Type conformance checks
var _ sqlgen2.TableWithIndexes = &AddressTable{}
var _ sqlgen2.TableWithCrud = &AddressTable{}

// NewAddressTable returns a new table instance.
// If a blank table name is supplied, the default name "addresses" will be used instead.
// The request context is initialised with the background.
func NewAddressTable(name sqlgen2.TableName, d *sqlgen2.Database) AddressTable {
	if name.Name == "" {
		name.Name = "addresses"
	}
	table := AddressTable{name, d, d.DB(), nil, context.Background()}
	return table
}

// CopyTableAsAddressTable copies a table instance, retaining the name etc but
// providing methods appropriate for 'Address'. It doesn't copy the constraints of the original table.
//
// It serves to provide methods appropriate for 'Address'. This is most useulf when thie is used to represent a
// join result. In such cases, there won't be any need for DDL methods, nor Exec, Insert, Update or Delete.
func CopyTableAsAddressTable(origin sqlgen2.Table) AddressTable {
	return AddressTable{
		name:        origin.Name(),
		database:    origin.Database(),
		db:          origin.DB(),
		constraints: nil,
		ctx:         origin.Ctx(),
	}
}

// WithPrefix sets the table name prefix for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl AddressTable) WithPrefix(pfx string) AddressTable {
	tbl.name.Prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl AddressTable) WithContext(ctx context.Context) AddressTable {
	tbl.ctx = ctx
	return tbl
}

// Database gets the shared database information.
func (tbl AddressTable) Database() *sqlgen2.Database {
	return tbl.database
}

// Logger gets the trace logger.
func (tbl AddressTable) Logger() *log.Logger {
	return tbl.database.Logger()
}

// WithConstraint returns a modified Table with added data consistency constraints.
func (tbl AddressTable) WithConstraint(cc ...constraint.Constraint) AddressTable {
	tbl.constraints = append(tbl.constraints, cc...)
	return tbl
}

// Ctx gets the current request context.
func (tbl AddressTable) Ctx() context.Context {
	return tbl.ctx
}

// Dialect gets the database dialect.
func (tbl AddressTable) Dialect() schema.Dialect {
	return tbl.database.Dialect()
}

// Name gets the table name.
func (tbl AddressTable) Name() sqlgen2.TableName {
	return tbl.name
}

// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl AddressTable) DB() *sql.DB {
	return tbl.db.(*sql.DB)
}

// Execer gets the wrapped database or transaction handle.
func (tbl AddressTable) Execer() sqlgen2.Execer {
	return tbl.db
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl AddressTable) Tx() *sql.Tx {
	return tbl.db.(*sql.Tx)
}

// IsTx tests whether this is within a transaction.
func (tbl AddressTable) IsTx() bool {
	_, ok := tbl.db.(*sql.Tx)
	return ok
}

// Begin starts a transaction. The default isolation level is dependent on the driver.
// The result is a modified copy of the table; the original is unchanged.
func (tbl AddressTable) BeginTx(opts *sql.TxOptions) (AddressTable, error) {
	d := tbl.db.(*sql.DB)
	var err error
	tbl.db, err = d.BeginTx(tbl.ctx, opts)
	return tbl, tbl.logIfError(err)
}

// Using returns a modified Table using the transaction supplied. This is needed
// when making multiple queries across several tables within a single transaction.
// The result is a modified copy of the table; the original is unchanged.
func (tbl AddressTable) Using(tx *sql.Tx) AddressTable {
	tbl.db = tx
	return tbl
}

func (tbl AddressTable) logQuery(query string, args ...interface{}) {
	support.LogQuery(tbl.Logger(), query, args...)
}

func (tbl AddressTable) logError(err error) error {
	return support.LogError(tbl.Logger(), err)
}

func (tbl AddressTable) logIfError(err error) error {
	return support.LogIfError(tbl.Logger(), err)
}


//--------------------------------------------------------------------------------

const NumAddressColumns = 3

const NumAddressDataColumns = 2

const AddressColumnNames = "id,lines,postcode"

const AddressDataColumnNames = "lines,postcode"

const AddressPk = "id"

//--------------------------------------------------------------------------------

const sqlAddressTableCreateColumnsSqlite = "\n"+
" `id`       integer primary key autoincrement,\n"+
" `lines`    text,\n"+
" `postcode` text"

const sqlAddressTableCreateColumnsMysql = "\n"+
" `id`       bigint primary key auto_increment,\n"+
" `lines`    json,\n"+
" `postcode` varchar(20)"

const sqlAddressTableCreateColumnsPostgres = `
 "id"       bigserial primary key,
 "lines"    json,
 "postcode" varchar(20)`

const sqlConstrainAddressTable = `
`

//--------------------------------------------------------------------------------

const sqlPostcodeIdxIndexColumns = "postcode"

//--------------------------------------------------------------------------------

// CreateTable creates the table.
func (tbl AddressTable) CreateTable(ifNotExists bool) (int64, error) {
	return tbl.Exec(nil, tbl.createTableSql(ifNotExists))
}

func (tbl AddressTable) createTableSql(ifNotExists bool) string {
	var columns string
	var settings string
	switch tbl.Dialect() {
	case schema.Sqlite:
		columns = sqlAddressTableCreateColumnsSqlite
		settings = ""
    case schema.Mysql:
		columns = sqlAddressTableCreateColumnsMysql
		settings = " ENGINE=InnoDB DEFAULT CHARSET=utf8"
    case schema.Postgres:
		columns = sqlAddressTableCreateColumnsPostgres
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

func (tbl AddressTable) ternary(flag bool, a, b string) string {
	if flag {
		return a
	}
	return b
}

// DropTable drops the table, destroying all its data.
func (tbl AddressTable) DropTable(ifExists bool) (int64, error) {
	return tbl.Exec(nil, tbl.dropTableSql(ifExists))
}

func (tbl AddressTable) dropTableSql(ifExists bool) string {
	ie := tbl.ternary(ifExists, "IF EXISTS ", "")
	query := fmt.Sprintf("DROP TABLE %s%s", ie, tbl.name)
	return query
}

//--------------------------------------------------------------------------------

// CreateTableWithIndexes invokes CreateTable then CreateIndexes.
func (tbl AddressTable) CreateTableWithIndexes(ifNotExist bool) (err error) {
	_, err = tbl.CreateTable(ifNotExist)
	if err != nil {
		return err
	}

	return tbl.CreateIndexes(ifNotExist)
}

// CreateIndexes executes queries that create the indexes needed by the Address table.
func (tbl AddressTable) CreateIndexes(ifNotExist bool) (err error) {

	err = tbl.CreatePostcodeIdxIndex(ifNotExist)
	if err != nil {
		return err
	}

	return nil
}

// CreatePostcodeIdxIndex creates the postcodeIdx index.
func (tbl AddressTable) CreatePostcodeIdxIndex(ifNotExist bool) error {
	ine := tbl.ternary(ifNotExist && tbl.Dialect() != schema.Mysql, "IF NOT EXISTS ", "")

	// Mysql does not support 'if not exists' on indexes
	// Workaround: use DropIndex first and ignore an error returned if the index didn't exist.

	if ifNotExist && tbl.Dialect() == schema.Mysql {
		tbl.Execer().ExecContext(tbl.Ctx(), tbl.dropPostcodeIdxIndexSql(false))
		ine = ""
	}

	_, err := tbl.Exec(nil, tbl.createPostcodeIdxIndexSql(ine))
	return err
}

func (tbl AddressTable) createPostcodeIdxIndexSql(ifNotExists string) string {
	indexPrefix := tbl.name.PrefixWithoutDot()
	return fmt.Sprintf("CREATE INDEX %s%spostcodeIdx ON %s (%s)", ifNotExists, indexPrefix,
		tbl.name, sqlPostcodeIdxIndexColumns)
}

// DropPostcodeIdxIndex drops the postcodeIdx index.
func (tbl AddressTable) DropPostcodeIdxIndex(ifExists bool) error {
	_, err := tbl.Exec(nil, tbl.dropPostcodeIdxIndexSql(ifExists))
	return err
}

func (tbl AddressTable) dropPostcodeIdxIndexSql(ifExists bool) string {
	// Mysql does not support 'if exists' on indexes
	ie := tbl.ternary(ifExists && tbl.Dialect() != schema.Mysql, "IF EXISTS ", "")
	onTbl := tbl.ternary(tbl.Dialect() == schema.Mysql, fmt.Sprintf(" ON %s", tbl.name), "")
	indexPrefix := tbl.name.PrefixWithoutDot()
	return fmt.Sprintf("DROP INDEX %s%spostcodeIdx%s", ie, indexPrefix, onTbl)
}

// DropIndexes executes queries that drop the indexes on by the Address table.
func (tbl AddressTable) DropIndexes(ifExist bool) (err error) {

	err = tbl.DropPostcodeIdxIndex(ifExist)
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
func (tbl AddressTable) Truncate(force bool) (err error) {
	for _, query := range tbl.Dialect().TruncateDDL(tbl.Name().String(), force) {
		_, err = tbl.Exec(nil, query)
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
func (tbl AddressTable) Exec(req require.Requirement, query string, args ...interface{}) (int64, error) {
	return support.Exec(tbl, req, query, args...)
}

//--------------------------------------------------------------------------------

// Query is the low-level access method for Addresses.
//
// It places a requirement, which may be nil, on the size of the expected results: this
// controls whether an error is generated when this expectation is not met.
//
// Note that this method applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl AddressTable) Query(req require.Requirement, query string, args ...interface{}) ([]*Address, error) {
	query = tbl.ReplaceTableName(query)
	vv, err := tbl.doQuery(req, false, query, args...)
	return vv, err
}

// QueryOne is the low-level access method for one Address.
// If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, *Address will be nil.
//
// Note that this method applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl AddressTable) QueryOne(query string, args ...interface{}) (*Address, error) {
	query = tbl.ReplaceTableName(query)
	return tbl.doQueryOne(nil, query, args...)
}

// MustQueryOne is the low-level access method for one Address.
//
// It places a requirement that exactly one result must be found; an error is generated when this expectation is not met.
//
// Note that this method applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl AddressTable) MustQueryOne(query string, args ...interface{}) (*Address, error) {
	query = tbl.ReplaceTableName(query)
	return tbl.doQueryOne(require.One, query, args...)
}

func (tbl AddressTable) doQueryOne(req require.Requirement, query string, args ...interface{}) (*Address, error) {
	list, err := tbl.doQuery(req, true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

func (tbl AddressTable) doQuery(req require.Requirement, firstOnly bool, query string, args ...interface{}) ([]*Address, error) {
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, tbl.logError(err)
	}
	defer rows.Close()

	vv, n, err := scanAddresses(rows, firstOnly)
	return vv, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(err, req, n))
}

func scanAddresses(rows *sql.Rows, firstOnly bool) (vv []*Address, n int64, err error) {
	for rows.Next() {
		n++

		var v0 int64
		var v1 []byte
		var v2 string

		err = rows.Scan(
			&v0,
			&v1,
			&v2,
		)
		if err != nil {
			return vv, n, err
		}

		v := &Address{}
		v.Id = v0
		err = json.Unmarshal(v1, &v.Lines)
		if err != nil {
			return nil, n, err
		}
		v.Postcode = v2

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
func (tbl AddressTable) QueryOneNullString(query string, args ...interface{}) (result sql.NullString, err error) {
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
func (tbl AddressTable) MustQueryOneNullString(query string, args ...interface{}) (result sql.NullString, err error) {
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
func (tbl AddressTable) QueryOneNullInt64(query string, args ...interface{}) (result sql.NullInt64, err error) {
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
func (tbl AddressTable) MustQueryOneNullInt64(query string, args ...interface{}) (result sql.NullInt64, err error) {
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
func (tbl AddressTable) QueryOneNullFloat64(query string, args ...interface{}) (result sql.NullFloat64, err error) {
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
func (tbl AddressTable) MustQueryOneNullFloat64(query string, args ...interface{}) (result sql.NullFloat64, err error) {
	err = support.QueryOneNullThing(tbl, require.One, &result, query, args...)
	return result, err
}

// ReplaceTableName replaces all occurrences of "{TABLE}" with the table's name.
func (tbl AddressTable) ReplaceTableName(query string) string {
	return strings.Replace(query, "{TABLE}", tbl.name.String(), -1)
}

//--------------------------------------------------------------------------------

var allAddressQuotedColumnNames = []string{
	schema.Sqlite.SplitAndQuote(AddressColumnNames),
	schema.Mysql.SplitAndQuote(AddressColumnNames),
	schema.Postgres.SplitAndQuote(AddressColumnNames),
}

//--------------------------------------------------------------------------------

// GetAddress gets the record with a given primary key value.
// If not found, *Address will be nil.
func (tbl AddressTable) GetAddress(id int64) (*Address, error) {
	dialect := tbl.Dialect()
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s=?",
		allAddressQuotedColumnNames[dialect.Index()], tbl.name, dialect.Quote("id"))
	v, err := tbl.doQueryOne(nil, query, id)
	return v, err
}

// MustGetAddress gets the record with a given primary key value.
//
// It places a requirement that exactly one result must be found; an error is generated when this expectation is not met.
func (tbl AddressTable) MustGetAddress(id int64) (*Address, error) {
	dialect := tbl.Dialect()
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s=?",
		allAddressQuotedColumnNames[dialect.Index()], tbl.name, dialect.Quote("id"))
	v, err := tbl.doQueryOne(require.One, query, id)
	return v, err
}

// GetAddresses gets records from the table according to a list of primary keys.
// Although the list of ids can be arbitrarily long, there are practical limits;
// note that Oracle DB has a limit of 1000.
//
// It places a requirement, which may be nil, on the size of the expected results: in particular, require.All
// controls whether an error is generated not all the ids produce a result.
func (tbl AddressTable) GetAddresses(req require.Requirement, id ...int64) (list []*Address, err error) {
	if len(id) > 0 {
		if req == require.All {
			req = require.Exactly(len(id))
		}
		dialect := tbl.Dialect()
		pl := dialect.Placeholders(len(id))
		query := fmt.Sprintf("SELECT %s FROM %s WHERE %s IN (%s)",
			allAddressQuotedColumnNames[dialect.Index()], tbl.name, dialect.Quote("id"), pl)
		args := make([]interface{}, len(id))

		for i, v := range id {
			args[i] = v
		}

		list, err = tbl.doQuery(req, false, query, args...)
	}

	return list, err
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
func (tbl AddressTable) SelectOneWhere(req require.Requirement, where, orderBy string, args ...interface{}) (*Address, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT 1",
		allAddressQuotedColumnNames[tbl.Dialect().Index()], tbl.name, where, orderBy)
	v, err := tbl.doQueryOne(req, query, args...)
	return v, err
}

// SelectOne allows a single Address to be obtained from the sqlgen2.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
// If not found, *Example will be nil.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.One
// controls whether an error is generated when no result is found.
func (tbl AddressTable) SelectOne(req require.Requirement, wh where.Expression, qc where.QueryConstraint) (*Address, error) {
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
	return tbl.SelectOneWhere(req, whs, orderBy, args...)
}

// SelectWhere allows Addresses to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
//
// The args are for any placeholder parameters in the query.
func (tbl AddressTable) SelectWhere(req require.Requirement, where, orderBy string, args ...interface{}) ([]*Address, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s %s",
		allAddressQuotedColumnNames[tbl.Dialect().Index()], tbl.name, where, orderBy)
	vv, err := tbl.doQuery(req, false, query, args...)
	return vv, err
}

// Select allows Addresses to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
func (tbl AddressTable) Select(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]*Address, error) {
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
	return tbl.SelectWhere(req, whs, orderBy, args...)
}

// CountWhere counts Addresses in the table that match a 'where' clause.
// Use a blank string for the 'where' argument if it is not needed.
//
// The args are for any placeholder parameters in the query.
func (tbl AddressTable) CountWhere(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s %s", tbl.name, where)
	tbl.logQuery(query, args...)
	row := tbl.db.QueryRowContext(tbl.ctx, query, args...)
	err = row.Scan(&count)
	return count, tbl.logIfError(err)
}

// Count counts the Addresses in the table that match a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed.
func (tbl AddressTable) Count(wh where.Expression) (count int64, err error) {
	whs, args := where.BuildExpression(wh, tbl.Dialect())
	return tbl.CountWhere(whs, args...)
}

//--------------------------------------------------------------------------------

// SliceId gets the Id column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AddressTable) SliceId(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return tbl.getint64list(req, "id", wh, qc)
}

// SlicePostcode gets the Postcode column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AddressTable) SlicePostcode(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.getstringlist(req, "postcode", wh, qc)
}


func (tbl AddressTable) getint64list(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
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

func (tbl AddressTable) getstringlist(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
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


func constructAddressInsert(w io.Writer, v *Address, dialect schema.Dialect, withPk bool) (s []interface{}, err error) {
	s = make([]interface{}, 0, 3)

	comma := ""
	io.WriteString(w, " (")

	if withPk {
		dialect.QuoteW(w, "id")
		comma = ","
		s = append(s, v.Id)
	}

	io.WriteString(w, comma)

	dialect.QuoteW(w, "lines")
	comma = ","
	x, err := json.Marshal(&v.Lines)
	if err != nil {
		return nil, err
	}
	s = append(s, x)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "postcode")
	s = append(s, v.Postcode)
	comma = ","
	io.WriteString(w, ")")
	return s, nil
}

func constructAddressUpdate(w io.Writer, v *Address, dialect schema.Dialect) (s []interface{}, err error) {
	j := 1
	s = make([]interface{}, 0, 2)

	comma := ""

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "lines", j)
	comma = ", "
		j++
	x, err := json.Marshal(&v.Lines)
	if err != nil {
		return nil, err
	}
	s = append(s, x)

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "postcode", j)
	s = append(s, v.Postcode)
	comma = ", "
		j++

	return s, nil
}

//--------------------------------------------------------------------------------

var allAddressQuotedInserts = []string{
	// Sqlite
	"(`lines`,`postcode`) VALUES (?,?)",
	// Mysql
	"(`lines`,`postcode`) VALUES (?,?)",
	// Postgres
	`("lines","postcode") VALUES ($1,$2) returning "id"`,
}

//--------------------------------------------------------------------------------

// Insert adds new records for the Addresses.
// The Addresses have their primary key fields set to the new record identifiers.
// The Address.PreInsert() method will be called, if it exists.
func (tbl AddressTable) Insert(req require.Requirement, vv ...*Address) error {
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

		fields, err := constructAddressInsert(b, v, tbl.Dialect(), false)
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

		v.Id, err = res.LastInsertId()
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

// UpdateFields updates one or more columns, given a 'where' clause.
//
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl AddressTable) UpdateFields(req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error) {
	return support.UpdateFields(tbl, req, wh, fields...)
}

//--------------------------------------------------------------------------------

var allAddressQuotedUpdates = []string{
	// Sqlite
	"`lines`=?,`postcode`=? WHERE `id`=?",
	// Mysql
	"`lines`=?,`postcode`=? WHERE `id`=?",
	// Postgres
	`"lines"=$2,"postcode"=$3 WHERE "id"=$1`,
}

//--------------------------------------------------------------------------------

// Update updates records, matching them by primary key. It returns the number of rows affected.
// The Address.PreUpdate(Execer) method will be called, if it exists.
func (tbl AddressTable) Update(req require.Requirement, vv ...*Address) (int64, error) {
	if req == require.All {
		req = require.Exactly(len(vv))
	}

	var count int64
	dialect := tbl.Dialect()
	//columns := allAddressQuotedUpdates[dialect.Index()]
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

		args, err := constructAddressUpdate(b, v, dialect)
		k := len(args)
		args = append(args, v.Id)
		if err != nil {
			return count, tbl.logError(err)
		}

		io.WriteString(b, " WHERE ")
		dialect.QuoteWithPlaceholder(b, "id", k)

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

// DeleteAddresses deletes rows from the table, given some primary keys.
// The list of ids can be arbitrarily long.
func (tbl AddressTable) DeleteAddresses(req require.Requirement, id ...int64) (int64, error) {
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
	col := dialect.Quote("id")
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
func (tbl AddressTable) Delete(req require.Requirement, wh where.Expression) (int64, error) {
	query, args := tbl.deleteRows(wh)
	return tbl.Exec(req, query, args...)
}

func (tbl AddressTable) deleteRows(wh where.Expression) (string, []interface{}) {
	whs, args := where.BuildExpression(wh, tbl.Dialect())
	query := fmt.Sprintf("DELETE FROM %s %s", tbl.name, whs)
	return query, args
}

//--------------------------------------------------------------------------------
