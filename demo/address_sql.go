// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.
// sqlapi v0.25.0-11-ga42fdd5; sqlgen v0.48.0-3-g84d0e25

package demo

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/rickb777/sqlapi"
	"github.com/rickb777/sqlapi/constraint"
	"github.com/rickb777/sqlapi/dialect"
	"github.com/rickb777/sqlapi/require"
	"github.com/rickb777/sqlapi/support"
	"github.com/rickb777/where"
	"github.com/rickb777/where/quote"
	"strings"
)

// AddressTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type AddressTable struct {
	name        sqlapi.TableName
	database    sqlapi.Database
	db          sqlapi.Execer
	constraints constraint.Constraints
	ctx         context.Context
	pk          string
}

// Type conformance checks
var _ sqlapi.TableWithIndexes = &AddressTable{}
var _ sqlapi.TableWithCrud = &AddressTable{}

// NewAddressTable returns a new table instance.
// If a blank table name is supplied, the default name "addresses" will be used instead.
// The request context is initialised with the background.
func NewAddressTable(name string, d sqlapi.Database) AddressTable {
	if name == "" {
		name = "addresses"
	}
	var constraints constraint.Constraints
	return AddressTable{
		name:        sqlapi.TableName{Prefix: "", Name: name},
		database:    d,
		db:          d.DB(),
		constraints: constraints,
		ctx:         context.Background(),
		pk:          "id",
	}
}

// CopyTableAsAddressTable copies a table instance, retaining the name etc but
// providing methods appropriate for 'Address'. It doesn't copy the constraints of the original table.
//
// It serves to provide methods appropriate for 'Address'. This is most useful when this is used to represent a
// join result. In such cases, there won't be any need for DDL methods, nor Exec, Insert, Update or Delete.
func CopyTableAsAddressTable(origin sqlapi.Table) AddressTable {
	return AddressTable{
		name:        origin.Name(),
		database:    origin.Database(),
		db:          origin.DB(),
		constraints: nil,
		ctx:         context.Background(),
		pk:          "id",
	}
}

// SetPkColumn sets the name of the primary key column. It defaults to "id".
// The result is a modified copy of the table; the original is unchanged.
func (tbl AddressTable) SetPkColumn(pk string) AddressTable {
	tbl.pk = pk
	return tbl
}

// WithPrefix sets the table name prefix for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl AddressTable) WithPrefix(pfx string) AddressTable {
	tbl.name.Prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries via this table.
// The result is a modified copy of the table; the original is unchanged.
//
// The shared context in the *Database is not altered by this method. So it
// is possible to use different contexts for different (groups of) queries.
func (tbl AddressTable) WithContext(ctx context.Context) AddressTable {
	tbl.ctx = ctx
	return tbl
}

// Database gets the shared database information.
func (tbl AddressTable) Database() sqlapi.Database {
	return tbl.database
}

// Logger gets the trace logger.
func (tbl AddressTable) Logger() sqlapi.Logger {
	return tbl.database.Logger()
}

// WithConstraint returns a modified Table with added data consistency constraints.
func (tbl AddressTable) WithConstraint(cc ...constraint.Constraint) AddressTable {
	tbl.constraints = append(tbl.constraints, cc...)
	return tbl
}

// Constraints returns the table's constraints.
func (tbl AddressTable) Constraints() constraint.Constraints {
	return tbl.constraints
}

// Ctx gets the current request context.
func (tbl AddressTable) Ctx() context.Context {
	return tbl.ctx
}

// Dialect gets the database dialect.
func (tbl AddressTable) Dialect() dialect.Dialect {
	return tbl.database.Dialect()
}

// Name gets the table name.
func (tbl AddressTable) Name() sqlapi.TableName {
	return tbl.name
}

// PkColumn gets the column name used as a primary key.
func (tbl AddressTable) PkColumn() string {
	return tbl.pk
}

// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl AddressTable) DB() sqlapi.SqlDB {
	return tbl.db.(sqlapi.SqlDB)
}

// Execer gets the wrapped database or transaction handle.
func (tbl AddressTable) Execer() sqlapi.Execer {
	return tbl.db
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl AddressTable) Tx() sqlapi.SqlTx {
	return tbl.db.(sqlapi.SqlTx)
}

// IsTx tests whether this is within a transaction.
func (tbl AddressTable) IsTx() bool {
	return tbl.db.IsTx()
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
func (tbl AddressTable) BeginTx(opts *sql.TxOptions) (AddressTable, error) {
	var err error
	tbl.db, err = tbl.db.(sqlapi.SqlDB).BeginTx(tbl.ctx, opts)
	return tbl, tbl.logIfError(err)
}

// Using returns a modified Table using the transaction supplied. This is needed
// when making multiple queries across several tables within a single transaction.
// The result is a modified copy of the table; the original is unchanged.
func (tbl AddressTable) Using(tx sqlapi.SqlTx) AddressTable {
	tbl.db = tx
	return tbl
}

func (tbl AddressTable) logQuery(query string, args ...interface{}) {
	tbl.database.LogQuery(query, args...)
}

func (tbl AddressTable) logError(err error) error {
	return tbl.database.LogError(err)
}

func (tbl AddressTable) logIfError(err error) error {
	return tbl.database.LogIfError(err)
}

func (tbl AddressTable) quotedName() string {
	return tbl.Dialect().Quoter().Quote(tbl.name.String())
}

func (tbl AddressTable) quotedNameW(w dialect.StringWriter) {
	tbl.Dialect().Quoter().QuoteW(w, tbl.name.String())
}

//--------------------------------------------------------------------------------

// NumAddressTableColumns is the total number of columns in AddressTable.
const NumAddressTableColumns = 4

// NumAddressTableDataColumns is the number of columns in AddressTable not including the auto-increment key.
const NumAddressTableDataColumns = 3

// AddressTableColumnNames is the list of columns in AddressTable.
const AddressTableColumnNames = "id,lines,town,postcode"

// AddressTableDataColumnNames is the list of data columns in AddressTable.
const AddressTableDataColumnNames = "lines,town,postcode"

var listOfAddressTableColumnNames = strings.Split(AddressTableColumnNames, ",")

//--------------------------------------------------------------------------------

var sqlAddressTableCreateColumnsSqlite = []string{
	"integer not null primary key autoincrement",
	"text",
	"text default null",
	"text not null",
}

var sqlAddressTableCreateColumnsMysql = []string{
	"bigint not null primary key auto_increment",
	"json",
	"varchar(80) default null",
	"varchar(20) not null",
}

var sqlAddressTableCreateColumnsPostgres = []string{
	"bigserial not null primary key",
	"json",
	"text default null",
	"text not null",
}

var sqlAddressTableCreateColumnsPgx = []string{
	"bigserial not null primary key",
	"json",
	"text default null",
	"text not null",
}

//--------------------------------------------------------------------------------

const sqlPostcodeIdxIndexColumns = "postcode"
var listOfPostcodeIdxIndexColumns = []string{"postcode"}

const sqlTownIdxIndexColumns = "town"
var listOfTownIdxIndexColumns = []string{"town"}

//--------------------------------------------------------------------------------

// CreateTable creates the table.
func (tbl AddressTable) CreateTable(ifNotExists bool) (int64, error) {
	return support.Exec(tbl, nil, tbl.createTableSql(ifNotExists))
}

func (tbl AddressTable) createTableSql(ifNotExists bool) string {
	buf := &bytes.Buffer{}
	buf.WriteString("CREATE TABLE ")
	if ifNotExists {
		buf.WriteString("IF NOT EXISTS ")
	}
	q := tbl.Dialect().Quoter()
	q.QuoteW(buf, tbl.name.String())
	buf.WriteString(" (\n ")

	var columns []string
	switch tbl.Dialect().Index() {
	case dialect.SqliteIndex:
		columns = sqlAddressTableCreateColumnsSqlite
	case dialect.MysqlIndex:
		columns = sqlAddressTableCreateColumnsMysql
	case dialect.PostgresIndex:
		columns = sqlAddressTableCreateColumnsPostgres
	case dialect.PgxIndex:
		columns = sqlAddressTableCreateColumnsPgx
	}

	comma := ""
	for i, n := range listOfAddressTableColumnNames {
		buf.WriteString(comma)
		q.QuoteW(buf, n)
		buf.WriteString(" ")
		buf.WriteString(columns[i])
		comma = ",\n "
	}

	for i, c := range tbl.constraints {
		buf.WriteString(",\n ")
		buf.WriteString(c.ConstraintSql(tbl.Dialect().Quoter(), tbl.name, i+1))
	}

	buf.WriteString("\n)")
	buf.WriteString(tbl.Dialect().CreateTableSettings())
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
	return support.Exec(tbl, nil, tbl.dropTableSql(ifExists))
}

func (tbl AddressTable) dropTableSql(ifExists bool) string {
	ie := tbl.ternary(ifExists, "IF EXISTS ", "")
	query := fmt.Sprintf("DROP TABLE %s%s", ie, tbl.quotedName())
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

	err = tbl.CreateTownIdxIndex(ifNotExist)
	if err != nil {
		return err
	}

	return nil
}

// CreatePostcodeIdxIndex creates the postcodeIdx index.
func (tbl AddressTable) CreatePostcodeIdxIndex(ifNotExist bool) error {
	ine := tbl.ternary(ifNotExist && tbl.Dialect().Index() != dialect.MysqlIndex, "IF NOT EXISTS ", "")

	// Mysql does not support 'if not exists' on indexes
	// Workaround: use DropIndex first and ignore an error returned if the index didn't exist.

	if ifNotExist && tbl.Dialect().Index() == dialect.MysqlIndex {
		// low-level no-logging Exec
		tbl.Execer().ExecContext(tbl.ctx, tbl.dropPostcodeIdxIndexSql(false))
		ine = ""
	}

	_, err := tbl.Exec(nil, tbl.createPostcodeIdxIndexSql(ine))
	return err
}

func (tbl AddressTable) createPostcodeIdxIndexSql(ifNotExists string) string {
	indexPrefix := tbl.name.PrefixWithoutDot()
	id := fmt.Sprintf("%spostcodeIdx", indexPrefix)
	q := tbl.Dialect().Quoter()
	cols := strings.Join(q.QuoteN(listOfPostcodeIdxIndexColumns), ",")
	return fmt.Sprintf("CREATE INDEX %s%s ON %s (%s)", ifNotExists,
		q.Quote(id), tbl.quotedName(), cols)
}

// DropPostcodeIdxIndex drops the postcodeIdx index.
func (tbl AddressTable) DropPostcodeIdxIndex(ifExists bool) error {
	_, err := tbl.Exec(nil, tbl.dropPostcodeIdxIndexSql(ifExists))
	return err
}

func (tbl AddressTable) dropPostcodeIdxIndexSql(ifExists bool) string {
	// Mysql does not support 'if exists' on indexes
	ie := tbl.ternary(ifExists && tbl.Dialect().Index() != dialect.MysqlIndex, "IF EXISTS ", "")
	indexPrefix := tbl.name.PrefixWithoutDot()
	id := fmt.Sprintf("%spostcodeIdx", indexPrefix)
	q := tbl.Dialect().Quoter()
	// Mysql requires extra "ON tbl" clause
	onTbl := tbl.ternary(tbl.Dialect().Index() == dialect.MysqlIndex, fmt.Sprintf(" ON %s", tbl.quotedName()), "")
	return "DROP INDEX " + ie + q.Quote(id) + onTbl
}

// CreateTownIdxIndex creates the townIdx index.
func (tbl AddressTable) CreateTownIdxIndex(ifNotExist bool) error {
	ine := tbl.ternary(ifNotExist && tbl.Dialect().Index() != dialect.MysqlIndex, "IF NOT EXISTS ", "")

	// Mysql does not support 'if not exists' on indexes
	// Workaround: use DropIndex first and ignore an error returned if the index didn't exist.

	if ifNotExist && tbl.Dialect().Index() == dialect.MysqlIndex {
		// low-level no-logging Exec
		tbl.Execer().ExecContext(tbl.ctx, tbl.dropTownIdxIndexSql(false))
		ine = ""
	}

	_, err := tbl.Exec(nil, tbl.createTownIdxIndexSql(ine))
	return err
}

func (tbl AddressTable) createTownIdxIndexSql(ifNotExists string) string {
	indexPrefix := tbl.name.PrefixWithoutDot()
	id := fmt.Sprintf("%stownIdx", indexPrefix)
	q := tbl.Dialect().Quoter()
	cols := strings.Join(q.QuoteN(listOfTownIdxIndexColumns), ",")
	return fmt.Sprintf("CREATE INDEX %s%s ON %s (%s)", ifNotExists,
		q.Quote(id), tbl.quotedName(), cols)
}

// DropTownIdxIndex drops the townIdx index.
func (tbl AddressTable) DropTownIdxIndex(ifExists bool) error {
	_, err := tbl.Exec(nil, tbl.dropTownIdxIndexSql(ifExists))
	return err
}

func (tbl AddressTable) dropTownIdxIndexSql(ifExists bool) string {
	// Mysql does not support 'if exists' on indexes
	ie := tbl.ternary(ifExists && tbl.Dialect().Index() != dialect.MysqlIndex, "IF EXISTS ", "")
	indexPrefix := tbl.name.PrefixWithoutDot()
	id := fmt.Sprintf("%stownIdx", indexPrefix)
	q := tbl.Dialect().Quoter()
	// Mysql requires extra "ON tbl" clause
	onTbl := tbl.ternary(tbl.Dialect().Index() == dialect.MysqlIndex, fmt.Sprintf(" ON %s", tbl.quotedName()), "")
	return "DROP INDEX " + ie + q.Quote(id) + onTbl
}

// DropIndexes executes queries that drop the indexes on by the Address table.
func (tbl AddressTable) DropIndexes(ifExist bool) (err error) {

	err = tbl.DropPostcodeIdxIndex(ifExist)
	if err != nil {
		return err
	}

	err = tbl.DropTownIdxIndex(ifExist)
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
func (tbl AddressTable) Exec(req require.Requirement, query string, args ...interface{}) (int64, error) {
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
// Wrap the result in *sqlapi.Rows if you need to access its data as a map.
func (tbl AddressTable) Query(query string, args ...interface{}) (sqlapi.SqlRows, error) {
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
func (tbl AddressTable) QueryOneNullString(req require.Requirement, query string, args ...interface{}) (result sql.NullString, err error) {
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
func (tbl AddressTable) QueryOneNullInt64(req require.Requirement, query string, args ...interface{}) (result sql.NullInt64, err error) {
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
func (tbl AddressTable) QueryOneNullFloat64(req require.Requirement, query string, args ...interface{}) (result sql.NullFloat64, err error) {
	err = support.QueryOneNullThing(tbl, req, &result, query, args...)
	return result, err
}

func scanAddresses(query string, rows sqlapi.SqlRows, firstOnly bool) (vv []*Address, n int64, err error) {
	for rows.Next() {
		n++

		var v0 int64
		var v1 []byte
		var v2 sql.NullString
		var v3 string

		err = rows.Scan(
			&v0,
			&v1,
			&v2,
			&v3,
		)
		if err != nil {
			return vv, n, errors.Wrap(err, query)
		}

		v := &Address{}
		v.Id = v0
		err = json.Unmarshal(v1, &v.Lines)
		if err != nil {
			return nil, n, errors.Wrap(err, query)
		}
		if v2.Valid {
			a := v2.String
			v.Town = &a
		}
		v.Postcode = v3

		var iv interface{} = v
		if hook, ok := iv.(sqlapi.CanPostGet); ok {
			err = hook.PostGet()
			if err != nil {
				return vv, n, errors.Wrap(err, query)
			}
		}

		vv = append(vv, v)

		if firstOnly {
			if rows.Next() {
				n++
			}
			return vv, n, errors.Wrap(rows.Err(), query)
		}
	}

	return vv, n, errors.Wrap(rows.Err(), query)
}

//--------------------------------------------------------------------------------

func allAddressColumnNamesQuoted(q quote.Quoter) string {
	return strings.Join(q.QuoteN(listOfAddressTableColumnNames), ",")
}

//--------------------------------------------------------------------------------

// GetAddressesById gets records from the table according to a list of primary keys.
// Although the list of ids can be arbitrarily long, there are practical limits;
// note that Oracle DB has a limit of 1000.
//
// It places a requirement, which may be nil, on the size of the expected results: in particular, require.All
// controls whether an error is generated not all the ids produce a result.
func (tbl AddressTable) GetAddressesById(req require.Requirement, id ...int64) (list []*Address, err error) {
	if len(id) > 0 {
		if req == require.All {
			req = require.Exactly(len(id))
		}
		args := make([]interface{}, len(id))

		for i, v := range id {
			args[i] = v
		}

		list, err = tbl.getAddresses(req, tbl.pk, args...)
	}

	return list, err
}

// GetAddressById gets the record with a given primary key value.
// If not found, *Address will be nil.
func (tbl AddressTable) GetAddressById(req require.Requirement, id int64) (*Address, error) {
	return tbl.getAddress(req, tbl.pk, id)
}

// GetAddressesByPostcode gets the records with a given postcode value.
// If not found, the resulting slice will be empty (nil).
func (tbl AddressTable) GetAddressesByPostcode(req require.Requirement, postcode string) ([]*Address, error) {
	return tbl.Select(req, where.And(where.Eq("postcode", postcode)), nil)
}

// GetAddressesByTown gets the records with a given town value.
// If not found, the resulting slice will be empty (nil).
func (tbl AddressTable) GetAddressesByTown(req require.Requirement, town string) ([]*Address, error) {
	return tbl.Select(req, where.And(where.Eq("town", town)), nil)
}

func (tbl AddressTable) getAddress(req require.Requirement, column string, arg interface{}) (*Address, error) {
	d := tbl.Dialect()
	q := d.Quoter()
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s=?",
		allAddressColumnNamesQuoted(q), tbl.quotedName(), q.Quote(column))
	v, err := tbl.doQueryAndScanOne(req, query, arg)
	return v, err
}

func (tbl AddressTable) getAddresses(req require.Requirement, column string, args ...interface{}) (list []*Address, err error) {
	if len(args) > 0 {
		if req == require.All {
			req = require.Exactly(len(args))
		}
		d := tbl.Dialect()
		q := d.Quoter()
		pl := d.Placeholders(len(args))
		query := fmt.Sprintf("SELECT %s FROM %s WHERE %s IN (%s)",
			allAddressColumnNamesQuoted(q), tbl.quotedName(), q.Quote(column), pl)
		list, err = tbl.doQueryAndScan(req, false, query, args...)
	}

	return list, err
}

func (tbl AddressTable) doQueryAndScanOne(req require.Requirement, query string, args ...interface{}) (*Address, error) {
	list, err := tbl.doQueryAndScan(req, true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

func (tbl AddressTable) doQueryAndScan(req require.Requirement, firstOnly bool, query string, args ...interface{}) ([]*Address, error) {
	rows, err := support.Query(tbl, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vv, n, err := scanAddresses(query, rows, firstOnly)
	return vv, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(err, req, n))
}

// Fetch fetches a list of Address based on a supplied query. This is mostly used for join queries that map its
// result columns to the fields of Address. Other queries might be better handled by GetXxx or Select methods.
func (tbl AddressTable) Fetch(req require.Requirement, query string, args ...interface{}) ([]*Address, error) {
	return tbl.doQueryAndScan(req, false, query, args...)
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
		allAddressColumnNamesQuoted(tbl.Dialect().Quoter()), tbl.quotedName(), where, orderBy)
	v, err := tbl.doQueryAndScanOne(req, query, args...)
	return v, err
}

// SelectOne allows a single Address to be obtained from the database.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
// If not found, *Example will be nil.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.One
// controls whether an error is generated when no result is found.
func (tbl AddressTable) SelectOne(req require.Requirement, wh where.Expression, qc where.QueryConstraint) (*Address, error) {
	q := tbl.Dialect().Quoter()
	whs, args := where.Where(wh, q)
	orderBy := where.Build(qc, q)
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
		allAddressColumnNamesQuoted(tbl.Dialect().Quoter()), tbl.quotedName(), where, orderBy)
	vv, err := tbl.doQueryAndScan(req, false, query, args...)
	return vv, err
}

// Select allows Addresses to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
func (tbl AddressTable) Select(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]*Address, error) {
	q := tbl.Dialect().Quoter()
	whs, args := where.Where(wh, q)
	orderBy := where.Build(qc, q)
	return tbl.SelectWhere(req, whs, orderBy, args...)
}

// CountWhere counts Addresses in the table that match a 'where' clause.
// Use a blank string for the 'where' argument if it is not needed.
//
// The args are for any placeholder parameters in the query.
func (tbl AddressTable) CountWhere(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s %s", tbl.quotedName(), where)
	rows, err := support.Query(tbl, query, args...)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&count)
	}
	return count, tbl.logIfError(err)
}

// Count counts the Addresses in the table that match a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed.
func (tbl AddressTable) Count(wh where.Expression) (count int64, err error) {
	whs, args := where.Where(wh, tbl.Dialect().Quoter())
	return tbl.CountWhere(whs, args...)
}

//--------------------------------------------------------------------------------

// SliceId gets the id column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AddressTable) SliceId(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return support.SliceInt64List(tbl, req, tbl.pk, wh, qc)
}

// SliceTown gets the town column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AddressTable) SliceTown(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return support.SliceStringPtrList(tbl, req, "town", wh, qc)
}

// SlicePostcode gets the postcode column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AddressTable) SlicePostcode(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return support.SliceStringList(tbl, req, "postcode", wh, qc)
}

func (tbl AddressTable) constructAddressInsert(w dialect.StringWriter, v *Address, withPk bool) (s []interface{}, err error) {
	q := tbl.Dialect().Quoter()
	s = make([]interface{}, 0, 4)

	comma := ""
	w.WriteString(" (")

	if withPk {
		q.QuoteW(w, "id")
		comma = ","
		s = append(s, v.Id)
	}

	w.WriteString(comma)
	q.QuoteW(w, "lines")
	comma = ","
	x, err := json.Marshal(&v.Lines)
	if err != nil {
		return nil, tbl.database.LogError(errors.WithStack(err))
	}
	s = append(s, x)

	if v.Town != nil {
		w.WriteString(comma)
		q.QuoteW(w, "town")
		s = append(s, v.Town)
	}

	w.WriteString(comma)
	q.QuoteW(w, "postcode")
	s = append(s, v.Postcode)

	w.WriteString(")")
	return s, nil
}

func (tbl AddressTable) constructAddressUpdate(w dialect.StringWriter, v *Address) (s []interface{}, err error) {
	q := tbl.Dialect().Quoter()
	j := 1
	s = make([]interface{}, 0, 3)

	comma := ""

	w.WriteString(comma)
	q.QuoteW(w, "lines")
	w.WriteString("=?")
	comma = ", "
	j++

	x, err := json.Marshal(&v.Lines)
	if err != nil {
		return nil, tbl.database.LogError(errors.WithStack(err))
	}
	s = append(s, x)

	w.WriteString(comma)
	if v.Town != nil {
		q.QuoteW(w, "town")
		w.WriteString("=?")
		s = append(s, v.Town)
		j++
	} else {
		q.QuoteW(w, "town")
		w.WriteString("=NULL")
	}

	w.WriteString(comma)
	q.QuoteW(w, "postcode")
	w.WriteString("=?")
	s = append(s, v.Postcode)
	j++
	return s, nil
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
	insertHasReturningPhrase := tbl.Dialect().InsertHasReturningPhrase()
	returning := ""
	if tbl.Dialect().InsertHasReturningPhrase() {
		returning = fmt.Sprintf(" returning %q", tbl.pk)
	}

	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlapi.CanPreInsert); ok {
			err := hook.PreInsert()
			if err != nil {
				return tbl.logError(err)
			}
		}

		b := dialect.Adapt(&bytes.Buffer{})
		b.WriteString("INSERT INTO ")
		tbl.quotedNameW(b)

		fields, err := tbl.constructAddressInsert(b, v, false)
		if err != nil {
			return tbl.logError(err)
		}

		b.WriteString(" VALUES (")
		b.WriteString(tbl.Dialect().Placeholders(len(fields)))
		b.WriteString(")")
		b.WriteString(returning)

		query := b.String()
		tbl.logQuery(query, fields...)

		var n int64 = 1
		if insertHasReturningPhrase {
			row := tbl.db.QueryRowContext(tbl.ctx, query, fields...)
			err = row.Scan(&v.Id)

		} else {
			i64, e2 := tbl.db.InsertContext(tbl.ctx, query, fields...)
			if e2 != nil {
				return tbl.logError(e2)
			}

			v.Id = i64
			}

		if err != nil {
			return tbl.logError(err)
		}
		count += n
	}

	return tbl.logIfError(require.ErrorIfExecNotSatisfiedBy(req, count))
}

// UpdateFields updates one or more columns, given a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl AddressTable) UpdateFields(req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error) {
	return support.UpdateFields(tbl, req, wh, fields...)
}

//--------------------------------------------------------------------------------

// Update updates records, matching them by primary key. It returns the number of rows affected.
// The Address.PreUpdate(Execer) method will be called, if it exists.
func (tbl AddressTable) Update(req require.Requirement, vv ...*Address) (int64, error) {
	if req == require.All {
		req = require.Exactly(len(vv))
	}

	var count int64
	d := tbl.Dialect()
	q := d.Quoter()

	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlapi.CanPreUpdate); ok {
			err := hook.PreUpdate()
			if err != nil {
				return count, tbl.logError(err)
			}
		}

		b := dialect.Adapt(&bytes.Buffer{})
		b.WriteString("UPDATE ")
		tbl.quotedNameW(b)
		b.WriteString(" SET ")

		args, err := tbl.constructAddressUpdate(b, v)
		if err != nil {
			return count, err
		}
		args = append(args, v.Id)

		b.WriteString(" WHERE ")
		q.QuoteW(b, tbl.pk)
		b.WriteString("=?")

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
	qName := tbl.quotedName()

	if req == require.All {
		req = require.Exactly(len(id))
	}

	var count, n int64
	var err error
	var max = batch
	if len(id) < batch {
		max = len(id)
	}
	d := tbl.Dialect()
	col := d.Quoter().Quote(tbl.pk)
	args := make([]interface{}, max)

	if len(id) > batch {
		pl := d.Placeholders(batch)
		query := fmt.Sprintf(qt, qName, col, pl)

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
		pl := d.Placeholders(len(id))
		query := fmt.Sprintf(qt, qName, col, pl)

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
	whs, args := where.Where(wh, tbl.Dialect().Quoter())
	query := fmt.Sprintf("DELETE FROM %s %s", tbl.quotedName(), whs)
	return query, args
}

//--------------------------------------------------------------------------------
