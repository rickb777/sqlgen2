// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.
// sqlapi v0.47.0; sqlgen v0.66.0

package demopgx

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx"
	"github.com/pkg/errors"
	"github.com/rickb777/sqlapi/dialect"
	"github.com/rickb777/sqlapi/pgxapi"
	"github.com/rickb777/sqlapi/pgxapi/constraint"
	"github.com/rickb777/sqlapi/pgxapi/support"
	"github.com/rickb777/sqlapi/require"
	"github.com/rickb777/where"
	"github.com/rickb777/where/quote"
	"strings"
)

// AddressTabler lists table methods provided by AddressTable.
type AddressTabler interface {
	pgxapi.Table

	// Constraints returns the table's constraints.
	Constraints() constraint.Constraints

	// WithConstraint returns a modified AddressTabler with added data consistency constraints.
	WithConstraint(cc ...constraint.Constraint) AddressTabler

	// WithPrefix returns a modified AddressTabler with a given table name prefix.
	WithPrefix(pfx string) AddressTabler

	// CreateTable creates the table.
	CreateTable(ctx context.Context, ifNotExists bool) (int64, error)

	// DropTable drops the table, destroying all its data.
	DropTable(ctx context.Context, ifExists bool) (int64, error)

	// CreateTableWithIndexes invokes CreateTable then CreateIndexes.
	CreateTableWithIndexes(ctx context.Context, ifNotExist bool) (err error)

	// CreateIndexes executes queries that create the indexes needed by the Address table.
	CreateIndexes(ctx context.Context, ifNotExist bool) (err error)

	// CreatePostcodeIdxIndex creates the postcodeIdx index.
	CreatePostcodeIdxIndex(ctx context.Context, ifNotExist bool) error

	// DropPostcodeIdxIndex drops the postcodeIdx index.
	DropPostcodeIdxIndex(ctx context.Context, ifExists bool) error

	// CreateTownIdxIndex creates the townIdx index.
	CreateTownIdxIndex(ctx context.Context, ifNotExist bool) error

	// DropTownIdxIndex drops the townIdx index.
	DropTownIdxIndex(ctx context.Context, ifExists bool) error

	// CreateUprnIdxIndex creates the uprn_idx index.
	CreateUprnIdxIndex(ctx context.Context, ifNotExist bool) error

	// DropUprnIdxIndex drops the uprn_idx index.
	DropUprnIdxIndex(ctx context.Context, ifExists bool) error

	// Truncate drops every record from the table, if possible.
	Truncate(ctx context.Context, force bool) (err error)
}

//-------------------------------------------------------------------------------------------------

// AddressQueryer lists query methods provided by AddressTable.
type AddressQueryer interface {
	pgxapi.Table

	// Using returns a modified AddressQueryer using the Execer supplied,
	// which will typically be a transaction (i.e. SqlTx).
	Using(tx pgxapi.Execer) AddressQueryer

	// Transact runs the function provided within a transaction. The transction is committed
	// unless an error occurs.
	Transact(ctx context.Context, txOptions *pgx.TxOptions, fn func(AddressQueryer) error) error

	// Exec executes a query without returning any rows.
	Exec(ctx context.Context, req require.Requirement, query string, args ...interface{}) (int64, error)

	// Query is the low-level request method for this table using an SQL query that must return all the columns
	// necessary for Address values.
	Query(ctx context.Context, req require.Requirement, query string, args ...interface{}) ([]*Address, error)

	// QueryOneNullString is a low-level access method for one string, returning the first match.
	QueryOneNullString(ctx context.Context, req require.Requirement, query string, args ...interface{}) (result sql.NullString, err error)

	// QueryOneNullInt64 is a low-level access method for one int64, returning the first match.
	QueryOneNullInt64(ctx context.Context, req require.Requirement, query string, args ...interface{}) (result sql.NullInt64, err error)

	// QueryOneNullFloat64 is a low-level access method for one float64, returning the first match.
	QueryOneNullFloat64(ctx context.Context, req require.Requirement, query string, args ...interface{}) (result sql.NullFloat64, err error)

	// GetAddressById gets the record with a given primary key value.
	GetAddressById(ctx context.Context, req require.Requirement, id int64) (*Address, error)

	// GetAddressesById gets records from the table according to a list of primary keys.
	GetAddressesById(ctx context.Context, req require.Requirement, qc where.QueryConstraint, id ...int64) (list []*Address, err error)

	// GetAddressesByPostcode gets the records with a given postcode value.
	GetAddressesByPostcode(ctx context.Context, req require.Requirement, qc where.QueryConstraint, postcode string) ([]*Address, error)

	// GetAddressesByTown gets the records with a given town value.
	GetAddressesByTown(ctx context.Context, req require.Requirement, qc where.QueryConstraint, town string) ([]*Address, error)

	// GetAddressByUPRN gets the record with a given uprn value.
	GetAddressByUPRN(ctx context.Context, req require.Requirement, uprn string) (*Address, error)

	// GetAddressesByUPRN gets the record with a given uprn value.
	GetAddressesByUPRN(ctx context.Context, req require.Requirement, qc where.QueryConstraint, uprn ...string) ([]*Address, error)

	// SelectOneWhere allows a single Address to be obtained from the table that matches a 'where' clause.
	SelectOneWhere(ctx context.Context, req require.Requirement, where, orderBy string, args ...interface{}) (*Address, error)

	// SelectOne allows a single Address to be obtained from the table that matches a 'where' clause.
	SelectOne(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) (*Address, error)

	// SelectWhere allows Addresses to be obtained from the table that match a 'where' clause.
	SelectWhere(ctx context.Context, req require.Requirement, where, orderBy string, args ...interface{}) ([]*Address, error)

	// Select allows Addresses to be obtained from the table that match a 'where' clause.
	Select(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]*Address, error)

	// CountWhere counts Addresses in the table that match a 'where' clause.
	CountWhere(ctx context.Context, where string, args ...interface{}) (count int64, err error)

	// Count counts the Addresses in the table that match a 'where' clause.
	Count(ctx context.Context, wh where.Expression) (count int64, err error)

	// SliceId gets the id column for all rows that match the 'where' condition.
	SliceId(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error)

	// SliceTown gets the town column for all rows that match the 'where' condition.
	SliceTown(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error)

	// SlicePostcode gets the postcode column for all rows that match the 'where' condition.
	SlicePostcode(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error)

	// SliceUprn gets the uprn column for all rows that match the 'where' condition.
	SliceUprn(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error)

	// Insert adds new records for the Addresses, setting the primary key field for each one.
	Insert(ctx context.Context, req require.Requirement, vv ...*Address) error

	// UpdateById updates one or more columns, given a id value.
	UpdateById(ctx context.Context, req require.Requirement, id int64, fields ...sql.NamedArg) (int64, error)

	// UpdateByTown updates one or more columns, given a town value.
	UpdateByTown(ctx context.Context, req require.Requirement, town string, fields ...sql.NamedArg) (int64, error)

	// UpdateByPostcode updates one or more columns, given a postcode value.
	UpdateByPostcode(ctx context.Context, req require.Requirement, postcode string, fields ...sql.NamedArg) (int64, error)

	// UpdateByUprn updates one or more columns, given a uprn value.
	UpdateByUprn(ctx context.Context, req require.Requirement, uprn string, fields ...sql.NamedArg) (int64, error)

	// UpdateFields updates one or more columns, given a 'where' clause.
	UpdateFields(ctx context.Context, req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error)

	// Update updates records, matching them by primary key.
	Update(ctx context.Context, req require.Requirement, vv ...*Address) (int64, error)

	// Upsert inserts or updates a record, matching it using the expression supplied.
	// This expression is used to search for an existing record based on some specified
	// key column(s). It must match either zero or one existing record. If it matches
	// none, a new record is inserted; otherwise the matching record is updated. An
	// error results if these conditions are not met.
	Upsert(ctx context.Context, v *Address, wh where.Expression) error

	// DeleteById deletes rows from the table, given some id values.
	// The list of ids can be arbitrarily long.
	DeleteById(ctx context.Context, req require.Requirement, id ...int64) (int64, error)

	// DeleteByTown deletes rows from the table, given some town values.
	// The list of ids can be arbitrarily long.
	DeleteByTown(ctx context.Context, req require.Requirement, town ...string) (int64, error)

	// DeleteByPostcode deletes rows from the table, given some postcode values.
	// The list of ids can be arbitrarily long.
	DeleteByPostcode(ctx context.Context, req require.Requirement, postcode ...string) (int64, error)

	// DeleteByUprn deletes rows from the table, given some uprn values.
	// The list of ids can be arbitrarily long.
	DeleteByUprn(ctx context.Context, req require.Requirement, uprn ...string) (int64, error)

	// Delete deletes one or more rows from the table, given a 'where' clause.
	// Use a nil value for the 'wh' argument if it is not needed (very risky!).
	Delete(ctx context.Context, req require.Requirement, wh where.Expression) (int64, error)
}

//-------------------------------------------------------------------------------------------------

// AddressTable holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type AddressTable struct {
	name        pgxapi.TableName
	database    pgxapi.Database
	db          pgxapi.Execer
	constraints constraint.Constraints
	pk          string
}

// Type conformance checks
var _ pgxapi.TableWithIndexes = &AddressTable{}

// NewAddressTable returns a new table instance.
// If a blank table name is supplied, the default name "addresses" will be used instead.
// The request context is initialised with the background.
func NewAddressTable(name string, d pgxapi.Database) AddressTable {
	if name == "" {
		name = "addresses"
	}
	var constraints constraint.Constraints
	return AddressTable{
		name:        pgxapi.TableName{Prefix: "", Name: name},
		database:    d,
		db:          d.DB(),
		constraints: constraints,
		pk:          "id",
	}
}

// CopyTableAsAddressTable copies a table instance, retaining the name etc but
// providing methods appropriate for 'Address'.It doesn't copy the constraints of the original table.
//
// It serves to provide methods appropriate for 'Address'. This is most useful when this is used to represent a
// join result. In such cases, there won't be any need for DDL methods, nor Exec, Insert, Update or Delete.
func CopyTableAsAddressTable(origin pgxapi.Table) AddressTable {
	return AddressTable{
		name:        origin.Name(),
		database:    origin.Database(),
		db:          origin.Execer(),
		constraints: nil,
		pk:          "id",
	}
}

// SetPkColumn sets the name of the primary key column. It defaults to "id".
// The result is a modified copy of the table; the original is unchanged.
//func (tbl AddressTable) SetPkColumn(pk string) AddressTabler {
//	tbl.pk = pk
//	return tbl
//}

// WithPrefix sets the table name prefix for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl AddressTable) WithPrefix(pfx string) AddressTabler {
	tbl.name.Prefix = pfx
	return tbl
}

// Database gets the shared database information.
func (tbl AddressTable) Database() pgxapi.Database {
	return tbl.database
}

// Logger gets the trace logger.
func (tbl AddressTable) Logger() pgxapi.Logger {
	return tbl.database.Logger()
}

// WithConstraint returns a modified AddressTabler with added data consistency constraints.
func (tbl AddressTable) WithConstraint(cc ...constraint.Constraint) AddressTabler {
	tbl.constraints = append(tbl.constraints, cc...)
	return tbl
}

// Constraints returns the table's constraints.
func (tbl AddressTable) Constraints() constraint.Constraints {
	return tbl.constraints
}

// Dialect gets the database dialect.
func (tbl AddressTable) Dialect() dialect.Dialect {
	return tbl.database.Dialect()
}

// Name gets the table name.
func (tbl AddressTable) Name() pgxapi.TableName {
	return tbl.name
}

// PkColumn gets the column name used as a primary key.
func (tbl AddressTable) PkColumn() string {
	return tbl.pk
}

// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl AddressTable) DB() pgxapi.SqlDB {
	return tbl.db.(pgxapi.SqlDB)
}

// Execer gets the wrapped database or transaction handle.
func (tbl AddressTable) Execer() pgxapi.Execer {
	return tbl.db
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl AddressTable) Tx() pgxapi.SqlTx {
	return tbl.db.(pgxapi.SqlTx)
}

// IsTx tests whether this is within a transaction.
func (tbl AddressTable) IsTx() bool {
	return tbl.db.IsTx()
}

// Using returns a modified AddressTabler using the the Execer supplied,
// which will typically be a transaction (i.e. SqlTx). This is needed when making multiple
// queries across several tables within a single transaction.
//
// The result is a modified copy of the table; the original is unchanged.
func (tbl AddressTable) Using(tx pgxapi.Execer) AddressQueryer {
	tbl.db = tx
	return tbl
}

// Transact runs the function provided within a transaction. If the function completes without error,
// the transaction is committed. If there is an error or a panic, the transaction is rolled back.
//
// The options can be nil, in which case the default behaviour is that of the underlying connection.
//
// Nested transactions (i.e. within 'fn') are permitted: they execute within the outermost transaction.
// Therefore they do not commit until the outermost transaction commits.
func (tbl AddressTable) Transact(ctx context.Context, txOptions *pgx.TxOptions, fn func(AddressQueryer) error) error {
	var err error
	if tbl.IsTx() {
		err = fn(tbl) // nested transactions are inlined
	} else {
		err = tbl.DB().Transact(ctx, txOptions, func(tx pgxapi.SqlTx) error {
			return fn(tbl.Using(tx))
		})
	}
	return tbl.Logger().LogIfError(err)
}

func (tbl AddressTable) quotedName() string {
	return tbl.Dialect().Quoter().Quote(tbl.name.String())
}

func (tbl AddressTable) quotedNameW(w dialect.StringWriter) {
	tbl.Dialect().Quoter().QuoteW(w, tbl.name.String())
}

//-------------------------------------------------------------------------------------------------

// NumAddressTableColumns is the total number of columns in AddressTable.
const NumAddressTableColumns = 5

// NumAddressTableDataColumns is the number of columns in AddressTable not including the auto-increment key.
const NumAddressTableDataColumns = 4

// AddressTableColumnNames is the list of columns in AddressTable.
const AddressTableColumnNames = "id,lines,town,postcode,uprn"

// AddressTableDataColumnNames is the list of data columns in AddressTable.
const AddressTableDataColumnNames = "lines,town,postcode,uprn"

var listOfAddressTableColumnNames = strings.Split(AddressTableColumnNames, ",")

//-------------------------------------------------------------------------------------------------

var sqlAddressTableCreateColumnsSqlite = []string{
	"integer not null primary key autoincrement",
	"text",
	"text default null",
	"text not null",
	"text not null",
}

var sqlAddressTableCreateColumnsMysql = []string{
	"bigint not null primary key auto_increment",
	"json",
	"varchar(80) default null",
	"varchar(20) not null",
	"varchar(20) not null",
}

var sqlAddressTableCreateColumnsPostgres = []string{
	"bigserial not null primary key",
	"json",
	"text default null",
	"text not null",
	"text not null",
}

var sqlAddressTableCreateColumnsPgx = []string{
	"bigserial not null primary key",
	"json",
	"text default null",
	"text not null",
	"text not null",
}

//-------------------------------------------------------------------------------------------------

const sqlPostcodeIdxIndexColumns = "postcode"

var listOfPostcodeIdxIndexColumns = []string{"postcode"}

const sqlTownIdxIndexColumns = "town"

var listOfTownIdxIndexColumns = []string{"town"}

const sqlUprnIdxIndexColumns = "uprn"

var listOfUprnIdxIndexColumns = []string{"uprn"}

//-------------------------------------------------------------------------------------------------

// CreateTable creates the table.
func (tbl AddressTable) CreateTable(ctx context.Context, ifNotExists bool) (int64, error) {
	return support.Exec(ctx, tbl, nil, createAddressTableSql(tbl, ifNotExists))
}

func createAddressTableSql(tbl AddressTabler, ifNotExists bool) string {
	buf := &bytes.Buffer{}
	buf.WriteString("CREATE TABLE ")
	if ifNotExists {
		buf.WriteString("IF NOT EXISTS ")
	}
	q := tbl.Dialect().Quoter()
	q.QuoteW(buf, tbl.Name().String())
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

	for i, c := range tbl.(AddressTable).Constraints() {
		buf.WriteString(",\n ")
		buf.WriteString(c.ConstraintSql(tbl.Dialect().Quoter(), tbl.Name(), i+1))
	}

	buf.WriteString("\n)")
	buf.WriteString(tbl.Dialect().CreateTableSettings())
	return buf.String()
}

func ternaryAddressTable(flag bool, a, b string) string {
	if flag {
		return a
	}
	return b
}

// DropTable drops the table, destroying all its data.
func (tbl AddressTable) DropTable(ctx context.Context, ifExists bool) (int64, error) {
	return support.Exec(ctx, tbl, nil, dropAddressTableSql(tbl, ifExists))
}

func dropAddressTableSql(tbl AddressTabler, ifExists bool) string {
	ie := ternaryAddressTable(ifExists, "IF EXISTS ", "")
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("DROP TABLE %s%s", ie, quotedName)
	return query
}

//-------------------------------------------------------------------------------------------------

// CreateTableWithIndexes invokes CreateTable then CreateIndexes.
func (tbl AddressTable) CreateTableWithIndexes(ctx context.Context, ifNotExist bool) (err error) {
	_, err = tbl.CreateTable(ctx, ifNotExist)
	if err != nil {
		return err
	}

	return tbl.CreateIndexes(ctx, ifNotExist)
}

// CreateIndexes executes queries that create the indexes needed by the Address table.
func (tbl AddressTable) CreateIndexes(ctx context.Context, ifNotExist bool) (err error) {

	err = tbl.CreatePostcodeIdxIndex(ctx, ifNotExist)
	if err != nil {
		return err
	}

	err = tbl.CreateTownIdxIndex(ctx, ifNotExist)
	if err != nil {
		return err
	}

	err = tbl.CreateUprnIdxIndex(ctx, ifNotExist)
	if err != nil {
		return err
	}

	return nil
}

// CreatePostcodeIdxIndex creates the postcodeIdx index.
func (tbl AddressTable) CreatePostcodeIdxIndex(ctx context.Context, ifNotExist bool) error {
	ine := ternaryAddressTable(ifNotExist && tbl.Dialect().Index() != dialect.MysqlIndex, "IF NOT EXISTS ", "")

	// Mysql does not support 'if not exists' on indexes
	// Workaround: use DropIndex first and ignore an error returned if the index didn't exist.

	if ifNotExist && tbl.Dialect().Index() == dialect.MysqlIndex {
		// low-level no-logging Exec
		tbl.Execer().ExecContext(ctx, dropAddressTablePostcodeIdxSql(tbl, false))
		ine = ""
	}

	_, err := tbl.Exec(ctx, nil, createAddressTablePostcodeIdxSql(tbl, ine))
	return err
}

func createAddressTablePostcodeIdxSql(tbl AddressTabler, ifNotExists string) string {
	indexPrefix := tbl.Name().PrefixWithoutDot()
	id := fmt.Sprintf("%s%s_postcodeIdx", indexPrefix, tbl.Name().Name)
	q := tbl.Dialect().Quoter()
	cols := strings.Join(q.QuoteN(listOfPostcodeIdxIndexColumns), ",")
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	return fmt.Sprintf("CREATE INDEX %s%s ON %s (%s)", ifNotExists,
		q.Quote(id), quotedName, cols)
}

// DropPostcodeIdxIndex drops the postcodeIdx index.
func (tbl AddressTable) DropPostcodeIdxIndex(ctx context.Context, ifExists bool) error {
	_, err := tbl.Exec(ctx, nil, dropAddressTablePostcodeIdxSql(tbl, ifExists))
	return err
}

func dropAddressTablePostcodeIdxSql(tbl AddressTabler, ifExists bool) string {
	// Mysql does not support 'if exists' on indexes
	ie := ternaryAddressTable(ifExists && tbl.Dialect().Index() != dialect.MysqlIndex, "IF EXISTS ", "")
	indexPrefix := tbl.Name().PrefixWithoutDot()
	id := fmt.Sprintf("%s%s_postcodeIdx", indexPrefix, tbl.Name().Name)
	q := tbl.Dialect().Quoter()
	// Mysql requires extra "ON tbl" clause
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	onTbl := ternaryAddressTable(tbl.Dialect().Index() == dialect.MysqlIndex, fmt.Sprintf(" ON %s", quotedName), "")
	return "DROP INDEX " + ie + q.Quote(id) + onTbl
}

// CreateTownIdxIndex creates the townIdx index.
func (tbl AddressTable) CreateTownIdxIndex(ctx context.Context, ifNotExist bool) error {
	ine := ternaryAddressTable(ifNotExist && tbl.Dialect().Index() != dialect.MysqlIndex, "IF NOT EXISTS ", "")

	// Mysql does not support 'if not exists' on indexes
	// Workaround: use DropIndex first and ignore an error returned if the index didn't exist.

	if ifNotExist && tbl.Dialect().Index() == dialect.MysqlIndex {
		// low-level no-logging Exec
		tbl.Execer().ExecContext(ctx, dropAddressTableTownIdxSql(tbl, false))
		ine = ""
	}

	_, err := tbl.Exec(ctx, nil, createAddressTableTownIdxSql(tbl, ine))
	return err
}

func createAddressTableTownIdxSql(tbl AddressTabler, ifNotExists string) string {
	indexPrefix := tbl.Name().PrefixWithoutDot()
	id := fmt.Sprintf("%s%s_townIdx", indexPrefix, tbl.Name().Name)
	q := tbl.Dialect().Quoter()
	cols := strings.Join(q.QuoteN(listOfTownIdxIndexColumns), ",")
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	return fmt.Sprintf("CREATE INDEX %s%s ON %s (%s)", ifNotExists,
		q.Quote(id), quotedName, cols)
}

// DropTownIdxIndex drops the townIdx index.
func (tbl AddressTable) DropTownIdxIndex(ctx context.Context, ifExists bool) error {
	_, err := tbl.Exec(ctx, nil, dropAddressTableTownIdxSql(tbl, ifExists))
	return err
}

func dropAddressTableTownIdxSql(tbl AddressTabler, ifExists bool) string {
	// Mysql does not support 'if exists' on indexes
	ie := ternaryAddressTable(ifExists && tbl.Dialect().Index() != dialect.MysqlIndex, "IF EXISTS ", "")
	indexPrefix := tbl.Name().PrefixWithoutDot()
	id := fmt.Sprintf("%s%s_townIdx", indexPrefix, tbl.Name().Name)
	q := tbl.Dialect().Quoter()
	// Mysql requires extra "ON tbl" clause
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	onTbl := ternaryAddressTable(tbl.Dialect().Index() == dialect.MysqlIndex, fmt.Sprintf(" ON %s", quotedName), "")
	return "DROP INDEX " + ie + q.Quote(id) + onTbl
}

// CreateUprnIdxIndex creates the uprn_idx index.
func (tbl AddressTable) CreateUprnIdxIndex(ctx context.Context, ifNotExist bool) error {
	ine := ternaryAddressTable(ifNotExist && tbl.Dialect().Index() != dialect.MysqlIndex, "IF NOT EXISTS ", "")

	// Mysql does not support 'if not exists' on indexes
	// Workaround: use DropIndex first and ignore an error returned if the index didn't exist.

	if ifNotExist && tbl.Dialect().Index() == dialect.MysqlIndex {
		// low-level no-logging Exec
		tbl.Execer().ExecContext(ctx, dropAddressTableUprnIdxSql(tbl, false))
		ine = ""
	}

	_, err := tbl.Exec(ctx, nil, createAddressTableUprnIdxSql(tbl, ine))
	return err
}

func createAddressTableUprnIdxSql(tbl AddressTabler, ifNotExists string) string {
	indexPrefix := tbl.Name().PrefixWithoutDot()
	id := fmt.Sprintf("%s%s_uprn_idx", indexPrefix, tbl.Name().Name)
	q := tbl.Dialect().Quoter()
	cols := strings.Join(q.QuoteN(listOfUprnIdxIndexColumns), ",")
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	return fmt.Sprintf("CREATE UNIQUE INDEX %s%s ON %s (%s)", ifNotExists,
		q.Quote(id), quotedName, cols)
}

// DropUprnIdxIndex drops the uprn_idx index.
func (tbl AddressTable) DropUprnIdxIndex(ctx context.Context, ifExists bool) error {
	_, err := tbl.Exec(ctx, nil, dropAddressTableUprnIdxSql(tbl, ifExists))
	return err
}

func dropAddressTableUprnIdxSql(tbl AddressTabler, ifExists bool) string {
	// Mysql does not support 'if exists' on indexes
	ie := ternaryAddressTable(ifExists && tbl.Dialect().Index() != dialect.MysqlIndex, "IF EXISTS ", "")
	indexPrefix := tbl.Name().PrefixWithoutDot()
	id := fmt.Sprintf("%s%s_uprn_idx", indexPrefix, tbl.Name().Name)
	q := tbl.Dialect().Quoter()
	// Mysql requires extra "ON tbl" clause
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	onTbl := ternaryAddressTable(tbl.Dialect().Index() == dialect.MysqlIndex, fmt.Sprintf(" ON %s", quotedName), "")
	return "DROP INDEX " + ie + q.Quote(id) + onTbl
}

// DropIndexes executes queries that drop the indexes on by the Address table.
func (tbl AddressTable) DropIndexes(ctx context.Context, ifExist bool) (err error) {

	err = tbl.DropPostcodeIdxIndex(ctx, ifExist)
	if err != nil {
		return err
	}

	err = tbl.DropTownIdxIndex(ctx, ifExist)
	if err != nil {
		return err
	}

	err = tbl.DropUprnIdxIndex(ctx, ifExist)
	if err != nil {
		return err
	}

	return nil
}

//-------------------------------------------------------------------------------------------------

// Truncate drops every record from the table, if possible. It might fail if constraints exist that
// prevent some or all rows from being deleted; use the force option to override this.
//
// When 'force' is set true, be aware of the following consequences.
// When using Mysql, foreign keys in other tables can be left dangling.
// When using Postgres, a cascade happens, so all 'adjacent' tables (i.e. linked by foreign keys)
// are also truncated.
func (tbl AddressTable) Truncate(ctx context.Context, force bool) (err error) {
	for _, query := range tbl.Dialect().TruncateDDL(tbl.Name().String(), force) {
		_, err = support.Exec(ctx, tbl, nil, query)
		if err != nil {
			return err
		}
	}
	return nil
}

//-------------------------------------------------------------------------------------------------

// Exec executes a query without returning any rows.
// It returns the number of rows affected (if the database driver supports this).
//
// The args are for any placeholder parameters in the query.
//
// If the context ctx is nil, it defaults to context.Background().
func (tbl AddressTable) Exec(ctx context.Context, req require.Requirement, query string, args ...interface{}) (int64, error) {
	return support.Exec(ctx, tbl, req, query, args...)
}

//-------------------------------------------------------------------------------------------------

// Query is the low-level request method for this table. The SQL query must return all the columns necessary for
// Address values. Placeholders should be vanilla '?' marks, which will be replaced if necessary according to
// the chosen dialect.
//
// The query is logged using whatever logger is configured. If an error arises, this too is logged.
//
// The args are for any placeholder parameters in the query.
//
// The support API provides a core 'support.Query' function, on which this method depends. If appropriate,
// use that function directly; wrap the result in *pgxapi.Rows if you need to access its data as a map.
//
// If the context ctx is nil, it defaults to context.Background().
func (tbl AddressTable) Query(ctx context.Context, req require.Requirement, query string, args ...interface{}) ([]*Address, error) {
	return doAddressTableQueryAndScan(ctx, tbl, req, false, query, args)
}

func doAddressTableQueryAndScan(ctx context.Context, tbl AddressTabler, req require.Requirement, firstOnly bool, query string, args ...interface{}) ([]*Address, error) {
	rows, err := support.Query(ctx, tbl, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vv, n, err := ScanAddresses(query, rows, firstOnly)
	return vv, tbl.Logger().LogIfError(require.ChainErrorIfQueryNotSatisfiedBy(err, req, n))
}

//-------------------------------------------------------------------------------------------------

// QueryOneNullString is a low-level access method for one string. This can be used for function queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl AddressTable) QueryOneNullString(ctx context.Context, req require.Requirement, query string, args ...interface{}) (result sql.NullString, err error) {
	err = support.QueryOneNullThing(ctx, tbl, req, &result, query, args...)
	return result, err
}

// QueryOneNullInt64 is a low-level access method for one int64. This can be used for 'COUNT(1)' queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl AddressTable) QueryOneNullInt64(ctx context.Context, req require.Requirement, query string, args ...interface{}) (result sql.NullInt64, err error) {
	err = support.QueryOneNullThing(ctx, tbl, req, &result, query, args...)
	return result, err
}

// QueryOneNullFloat64 is a low-level access method for one float64. This can be used for 'AVG(...)' queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl AddressTable) QueryOneNullFloat64(ctx context.Context, req require.Requirement, query string, args ...interface{}) (result sql.NullFloat64, err error) {
	err = support.QueryOneNullThing(ctx, tbl, req, &result, query, args...)
	return result, err
}

// ScanAddresses reads rows from the database and returns a slice of corresponding values.
// It also returns a number indicating how many rows were read; this will be larger than the length of the
// slice if reading stopped after the first row.
func ScanAddresses(query string, rows pgxapi.SqlRows, firstOnly bool) (vv []*Address, n int64, err error) {
	for rows.Next() {
		n++

		var v0 int64
		var v1 []byte
		var v2 sql.NullString
		var v3 string
		var v4 string

		err = rows.Scan(
			&v0,
			&v1,
			&v2,
			&v3,
			&v4,
		)
		if err != nil {
			return vv, n, errors.Wrap(err, query)
		}

		v := &Address{}
		v.Id = v0
		err = json.Unmarshal(v1, &v.AddressFields.Lines)
		if err != nil {
			return nil, n, errors.Wrap(err, query)
		}
		if v2.Valid {
			a := v2.String
			v.AddressFields.Town = &a
		}
		v.AddressFields.Postcode = v3
		v.AddressFields.UPRN = v4

		var iv interface{} = v
		if hook, ok := iv.(pgxapi.CanPostGet); ok {
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

// GetAddressById gets the record with a given primary key value.
// If not found, *Address will be nil.
func (tbl AddressTable) GetAddressById(ctx context.Context, req require.Requirement, id int64) (*Address, error) {
	return tbl.SelectOne(ctx, req, where.Eq("id", id), nil)
}

// GetAddressesById gets records from the table according to a list of primary keys.
// Although the list of ids can be arbitrarily long, there are practical limits;
// note that Oracle DB has a limit of 1000.
//
// It places a requirement, which may be nil, on the size of the expected results: in particular, require.All
// controls whether an error is generated not all the ids produce a result.
func (tbl AddressTable) GetAddressesById(ctx context.Context, req require.Requirement, qc where.QueryConstraint, id ...int64) (list []*Address, err error) {
	if req == require.All {
		req = require.Exactly(len(id))
	}
	return tbl.Select(ctx, req, where.In("id", id), qc)
}

// GetAddressesByPostcode gets the records with a given postcode value.
// If not found, the resulting slice will be empty (nil).
func (tbl AddressTable) GetAddressesByPostcode(ctx context.Context, req require.Requirement, qc where.QueryConstraint, postcode string) ([]*Address, error) {
	return tbl.Select(ctx, req, where.And(where.Eq("postcode", postcode)), qc)
}

// GetAddressesByTown gets the records with a given town value.
// If not found, the resulting slice will be empty (nil).
func (tbl AddressTable) GetAddressesByTown(ctx context.Context, req require.Requirement, qc where.QueryConstraint, town string) ([]*Address, error) {
	return tbl.Select(ctx, req, where.And(where.Eq("town", town)), qc)
}

// GetAddressByUPRN gets the record with a given uprn value.
// If not found, *Address will be nil.
func (tbl AddressTable) GetAddressByUPRN(ctx context.Context, req require.Requirement, uprn string) (*Address, error) {
	return tbl.SelectOne(ctx, req, where.And(where.Eq("uprn", uprn)), nil)
}

// GetAddressesByUPRN gets the record with a given uprn value.
func (tbl AddressTable) GetAddressesByUPRN(ctx context.Context, req require.Requirement, qc where.QueryConstraint, uprn ...string) ([]*Address, error) {
	if req == require.All {
		req = require.Exactly(len(uprn))
	}
	return tbl.Select(ctx, req, where.In("uprn", uprn), qc)
}

func doAddressTableQueryAndScanOne(ctx context.Context, tbl AddressTabler, req require.Requirement, query string, args ...interface{}) (*Address, error) {
	list, err := doAddressTableQueryAndScan(ctx, tbl, req, true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

// Fetch fetches a list of Address based on a supplied query. This is mostly used for join queries that map its
// result columns to the fields of Address. Other queries might be better handled by GetXxx or Select methods.
func (tbl AddressTable) Fetch(ctx context.Context, req require.Requirement, query string, args ...interface{}) ([]*Address, error) {
	return doAddressTableQueryAndScan(ctx, tbl, req, false, query, args...)
}

//-------------------------------------------------------------------------------------------------

// SelectOneWhere allows a single Address to be obtained from the table that matches a 'where' clause
// and some limit. Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
// If not found, *Example will be nil.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.One
// controls whether an error is generated when no result is found.
//
// The args are for any placeholder parameters in the query.
func (tbl AddressTable) SelectOneWhere(ctx context.Context, req require.Requirement, where, orderBy string, args ...interface{}) (*Address, error) {
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT 1",
		allAddressColumnNamesQuoted(tbl.Dialect().Quoter()), quotedName, where, orderBy)
	v, err := doAddressTableQueryAndScanOne(ctx, tbl, req, query, args...)
	return v, err
}

// SelectOne allows a single Address to be obtained from the table that matches a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
// If not found, *Example will be nil.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.One
// controls whether an error is generated when no result is found.
func (tbl AddressTable) SelectOne(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) (*Address, error) {
	q := tbl.Dialect().Quoter()
	whs, args := where.Where(wh, q)
	orderBy := where.Build(qc, q)
	return tbl.SelectOneWhere(ctx, req, whs, orderBy, args...)
}

// SelectWhere allows Addresses to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
//
// The args are for any placeholder parameters in the query.
func (tbl AddressTable) SelectWhere(ctx context.Context, req require.Requirement, where, orderBy string, args ...interface{}) ([]*Address, error) {
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("SELECT %s FROM %s %s %s",
		allAddressColumnNamesQuoted(tbl.Dialect().Quoter()), quotedName, where, orderBy)
	vv, err := doAddressTableQueryAndScan(ctx, tbl, req, false, query, args...)
	return vv, err
}

// Select allows Addresses to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
func (tbl AddressTable) Select(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]*Address, error) {
	q := tbl.Dialect().Quoter()
	whs, args := where.Where(wh, q)
	orderBy := where.Build(qc, q)
	return tbl.SelectWhere(ctx, req, whs, orderBy, args...)
}

//-------------------------------------------------------------------------------------------------

// CountWhere counts Addresses in the table that match a 'where' clause.
// Use a blank string for the 'where' argument if it is not needed.
//
// The args are for any placeholder parameters in the query.
func (tbl AddressTable) CountWhere(ctx context.Context, where string, args ...interface{}) (count int64, err error) {
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s %s", quotedName, where)
	rows, err := support.Query(ctx, tbl, query, args...)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&count)
	}
	return count, tbl.Logger().LogIfError(err)
}

// Count counts the Addresses in the table that match a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed.
func (tbl AddressTable) Count(ctx context.Context, wh where.Expression) (count int64, err error) {
	whs, args := where.Where(wh, tbl.Dialect().Quoter())
	return tbl.CountWhere(ctx, whs, args...)
}

//--------------------------------------------------------------------------------

// SliceId gets the id column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AddressTable) SliceId(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return support.SliceInt64List(ctx, tbl, req, tbl.pk, wh, qc)
}

// SliceTown gets the town column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AddressTable) SliceTown(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return support.SliceStringPtrList(ctx, tbl, req, "town", wh, qc)
}

// SlicePostcode gets the postcode column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AddressTable) SlicePostcode(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return support.SliceStringList(ctx, tbl, req, "postcode", wh, qc)
}

// SliceUprn gets the uprn column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl AddressTable) SliceUprn(ctx context.Context, req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return support.SliceStringList(ctx, tbl, req, "uprn", wh, qc)
}

func constructAddressTableInsert(tbl AddressTable, w dialect.StringWriter, v *Address, withPk bool) (s []interface{}, err error) {
	q := tbl.Dialect().Quoter()
	s = make([]interface{}, 0, 5)

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
	x, err := json.Marshal(&v.AddressFields.Lines)
	if err != nil {
		return nil, tbl.Logger().LogError(errors.WithStack(err))
	}
	s = append(s, x)

	if v.AddressFields.Town != nil {
		w.WriteString(comma)
		q.QuoteW(w, "town")
		s = append(s, v.AddressFields.Town)
	}

	w.WriteString(comma)
	q.QuoteW(w, "postcode")
	s = append(s, v.AddressFields.Postcode)

	w.WriteString(comma)
	q.QuoteW(w, "uprn")
	s = append(s, v.AddressFields.UPRN)

	w.WriteString(")")
	return s, nil
}

func constructAddressTableUpdate(tbl AddressTable, w dialect.StringWriter, v *Address) (s []interface{}, err error) {
	q := tbl.Dialect().Quoter()
	j := 1
	s = make([]interface{}, 0, 4)

	comma := ""

	w.WriteString(comma)
	q.QuoteW(w, "lines")
	w.WriteString("=?")
	comma = ", "
	j++

	x, err := json.Marshal(&v.AddressFields.Lines)
	if err != nil {
		return nil, tbl.Logger().LogError(errors.WithStack(err))
	}
	s = append(s, x)

	w.WriteString(comma)
	if v.AddressFields.Town != nil {
		q.QuoteW(w, "town")
		w.WriteString("=?")
		s = append(s, v.AddressFields.Town)
		j++
	} else {
		q.QuoteW(w, "town")
		w.WriteString("=NULL")
	}

	w.WriteString(comma)
	q.QuoteW(w, "postcode")
	w.WriteString("=?")
	s = append(s, v.AddressFields.Postcode)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "uprn")
	w.WriteString("=?")
	s = append(s, v.AddressFields.UPRN)
	j++
	return s, nil
}

//--------------------------------------------------------------------------------

// Insert adds new records for the Addresses.// The Addresses have their primary key fields set to the new record identifiers.
// The Address.PreInsert() method will be called, if it exists.
func (tbl AddressTable) Insert(ctx context.Context, req require.Requirement, vv ...*Address) error {
	if req == require.All {
		req = require.Exactly(len(vv))
	}

	if ctx == nil {
		ctx = context.Background()
	}

	var count int64
	returning := ""
	insertHasReturningPhrase := tbl.Dialect().InsertHasReturningPhrase()
	if insertHasReturningPhrase {
		returning = fmt.Sprintf(" RETURNING %q", tbl.pk)
	}

	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(pgxapi.CanPreInsert); ok {
			err := hook.PreInsert()
			if err != nil {
				return tbl.Logger().LogError(err)
			}
		}

		b := dialect.Adapt(&bytes.Buffer{})
		b.WriteString("INSERT INTO ")
		tbl.quotedNameW(b)

		fields, err := constructAddressTableInsert(tbl, b, v, false)
		if err != nil {
			return tbl.Logger().LogError(err)
		}

		b.WriteString(" VALUES (")
		b.WriteString(tbl.Dialect().Placeholders(len(fields)))
		b.WriteString(")")
		b.WriteString(returning)

		query := b.String()
		tbl.Logger().LogQuery(query, fields...)

		var n int64 = 1
		if insertHasReturningPhrase {
			row := tbl.db.QueryRowContext(ctx, query, fields...)
			var i64 int64
			err = row.Scan(&i64)
			v.Id = i64

		} else {
			i64, e2 := tbl.db.InsertContext(ctx, tbl.pk, query, fields...)
			if e2 != nil {
				return tbl.Logger().LogError(e2)
			}
			v.Id = i64
		}

		if err != nil {
			return tbl.Logger().LogError(err)
		}
		count += n
	}

	return tbl.Logger().LogIfError(require.ErrorIfExecNotSatisfiedBy(req, count))
}

// UpdateById updates one or more columns, given a id value.
func (tbl AddressTable) UpdateById(ctx context.Context, req require.Requirement, id int64, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("id", id), fields...)
}

// UpdateByTown updates one or more columns, given a town value.
func (tbl AddressTable) UpdateByTown(ctx context.Context, req require.Requirement, town string, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("town", town), fields...)
}

// UpdateByPostcode updates one or more columns, given a postcode value.
func (tbl AddressTable) UpdateByPostcode(ctx context.Context, req require.Requirement, postcode string, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("postcode", postcode), fields...)
}

// UpdateByUprn updates one or more columns, given a uprn value.
func (tbl AddressTable) UpdateByUprn(ctx context.Context, req require.Requirement, uprn string, fields ...sql.NamedArg) (int64, error) {
	return tbl.UpdateFields(ctx, req, where.Eq("uprn", uprn), fields...)
}

// UpdateFields updates one or more columns, given a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed (but note that this is risky!).
func (tbl AddressTable) UpdateFields(ctx context.Context, req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error) {
	return support.UpdateFields(ctx, tbl, req, wh, fields...)
}

//--------------------------------------------------------------------------------

// Update updates records, matching them by primary key. It returns the number of rows affected.
// The Address.PreUpdate(Execer) method will be called, if it exists.
func (tbl AddressTable) Update(ctx context.Context, req require.Requirement, vv ...*Address) (int64, error) {
	if req == require.All {
		req = require.Exactly(len(vv))
	}

	var count int64
	d := tbl.Dialect()
	q := d.Quoter()

	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(pgxapi.CanPreUpdate); ok {
			err := hook.PreUpdate()
			if err != nil {
				return count, tbl.Logger().LogError(err)
			}
		}

		b := dialect.Adapt(&bytes.Buffer{})
		b.WriteString("UPDATE ")
		tbl.quotedNameW(b)
		b.WriteString(" SET ")

		args, err := constructAddressTableUpdate(tbl, b, v)
		if err != nil {
			return count, err
		}
		args = append(args, v.Id)

		b.WriteString(" WHERE ")
		q.QuoteW(b, tbl.pk)
		b.WriteString("=?")

		query := b.String()
		n, err := tbl.Exec(ctx, nil, query, args...)
		if err != nil {
			return count, err
		}
		count += n
	}

	return count, tbl.Logger().LogIfError(require.ErrorIfExecNotSatisfiedBy(req, count))
}

//--------------------------------------------------------------------------------

// Upsert inserts or updates a record, matching it using the expression supplied.
// This expression is used to search for an existing record based on some specified
// key column(s). It must match either zero or one existing record. If it matches
// none, a new record is inserted; otherwise the matching record is updated. An
// error results if these conditions are not met.
func (tbl AddressTable) Upsert(ctx context.Context, v *Address, wh where.Expression) error {
	col := tbl.Dialect().Quoter().Quote(tbl.pk)
	qName := tbl.quotedName()
	whs, args := where.Where(wh, tbl.Dialect().Quoter())

	query := fmt.Sprintf("SELECT %s FROM %s %s", col, qName, whs)
	rows, err := support.Query(ctx, tbl, query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	if !rows.Next() {
		return tbl.Insert(ctx, require.One, v)
	}

	var id int64
	err = rows.Scan(&id)
	if err != nil {
		return tbl.Logger().LogIfError(err)
	}

	if rows.Next() {
		return require.ErrWrongSize(2, "expected to find no more than 1 but got at least 2 using %q", wh)
	}

	v.Id = id
	_, err = tbl.Update(ctx, require.One, v)
	return err
}

//-------------------------------------------------------------------------------------------------

// DeleteById deletes rows from the table, given some id values.
// The list of ids can be arbitrarily long.
func (tbl AddressTable) DeleteById(ctx context.Context, req require.Requirement, id ...int64) (int64, error) {
	ii := support.Int64AsInterfaceSlice(id)
	return support.DeleteByColumn(ctx, tbl, req, "id", ii...)
}

// DeleteByTown deletes rows from the table, given some town values.
// The list of ids can be arbitrarily long.
func (tbl AddressTable) DeleteByTown(ctx context.Context, req require.Requirement, town ...string) (int64, error) {
	ii := support.StringAsInterfaceSlice(town)
	return support.DeleteByColumn(ctx, tbl, req, "town", ii...)
}

// DeleteByPostcode deletes rows from the table, given some postcode values.
// The list of ids can be arbitrarily long.
func (tbl AddressTable) DeleteByPostcode(ctx context.Context, req require.Requirement, postcode ...string) (int64, error) {
	ii := support.StringAsInterfaceSlice(postcode)
	return support.DeleteByColumn(ctx, tbl, req, "postcode", ii...)
}

// DeleteByUprn deletes rows from the table, given some uprn values.
// The list of ids can be arbitrarily long.
func (tbl AddressTable) DeleteByUprn(ctx context.Context, req require.Requirement, uprn ...string) (int64, error) {
	ii := support.StringAsInterfaceSlice(uprn)
	return support.DeleteByColumn(ctx, tbl, req, "uprn", ii...)
}

// Delete deletes one or more rows from the table, given a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl AddressTable) Delete(ctx context.Context, req require.Requirement, wh where.Expression) (int64, error) {
	query, args := deleteRowsAddressTableSql(tbl, wh)
	return tbl.Exec(ctx, req, query, args...)
}

func deleteRowsAddressTableSql(tbl AddressTabler, wh where.Expression) (string, []interface{}) {
	whs, args := where.Where(wh, tbl.Dialect().Quoter())
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("DELETE FROM %s %s", quotedName, whs)
	return query, args
}

//-------------------------------------------------------------------------------------------------
