package code

import "text/template"

// template to declare the package name.
var sPackage = `// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

package %s
`

//-------------------------------------------------------------------------------------------------

const sTable = `
// {{.Prefix}}{{.Type}}{{.Thing}} holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type {{.Prefix}}{{.Type}}{{.Thing}} struct {
	name        sqlgen2.TableName
	database    *sqlgen2.Database
	db          sqlgen2.Execer
	constraints constraint.Constraints
	ctx			context.Context
	pk          string
}

// Type conformance checks
var _ {{.Interface1}} = &{{.Prefix}}{{.Type}}{{.Thing}}{}
var _ {{.Interface2}} = &{{.Prefix}}{{.Type}}{{.Thing}}{}

// New{{.Prefix}}{{.Type}}{{.Thing}} returns a new table instance.
// If a blank table name is supplied, the default name "{{.DbName}}" will be used instead.
// The request context is initialised with the background.
func New{{.Prefix}}{{.Type}}{{.Thing}}(name string, d *sqlgen2.Database) {{.Prefix}}{{.Type}}{{.Thing}} {
	if name == "" {
		name = "{{.DbName}}"
	}
	var constraints constraint.Constraints
	{{- range .Constraints}}
	constraints = append(constraints, {{.GoString}})
	{{end}}
	return {{.Prefix}}{{.Type}}{{.Thing}}{
		name:        sqlgen2.TableName{"", name},
		database:    d,
		db:          d.DB(),
		constraints: constraints,
		ctx:         context.Background(),
		pk:          "{{.Table.SafePrimary.SqlName}}",
	}
}

// CopyTableAs{{.Prefix}}{{.Type}}{{.Thing}} copies a table instance, retaining the name etc but
// providing methods appropriate for '{{.Type}}'. It doesn't copy the constraints of the original table.
//
// It serves to provide methods appropriate for '{{.Type}}'. This is most useful when this is used to represent a
// join result. In such cases, there won't be any need for DDL methods, nor Exec, Insert, Update or Delete.
func CopyTableAs{{title .Prefix}}{{title .Type}}{{.Thing}}(origin sqlgen2.Table) {{.Prefix}}{{.Type}}{{.Thing}} {
	return {{.Prefix}}{{.Type}}{{.Thing}}{
		name:        origin.Name(),
		database:    origin.Database(),
		db:          origin.DB(),
		constraints: nil,
		ctx:         context.Background(),
		pk:          "{{.Table.SafePrimary.SqlName}}",
	}
}

{{if .Table.HasPrimaryKey}}
// SetPkColumn sets the name of the primary key column. It defaults to "{{.Table.Primary.SqlName}}".
// The result is a modified copy of the table; the original is unchanged.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) SetPkColumn(pk string) {{.Prefix}}{{.Type}}{{.Thing}} {
	tbl.pk = pk
	return tbl
}

{{end}}
// WithPrefix sets the table name prefix for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) WithPrefix(pfx string) {{.Prefix}}{{.Type}}{{.Thing}} {
	tbl.name.Prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries via this table.
// The result is a modified copy of the table; the original is unchanged.
//
// The shared context in the *Database is not altered by this method. So it
// is possible to use different contexts for different (groups of) queries.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) WithContext(ctx context.Context) {{.Prefix}}{{.Type}}{{.Thing}} {
	tbl.ctx = ctx
	return tbl
}

// Database gets the shared database information.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Database() *sqlgen2.Database {
	return tbl.database
}

// Logger gets the trace logger.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Logger() *log.Logger {
	return tbl.database.Logger()
}

// WithConstraint returns a modified Table with added data consistency constraints.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) WithConstraint(cc ...constraint.Constraint) {{.Prefix}}{{.Type}}{{.Thing}} {
	tbl.constraints = append(tbl.constraints, cc...)
	return tbl
}

// Constraints returns the table's constraints.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Constraints() constraint.Constraints {
	return tbl.constraints
}

// Ctx gets the current request context.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Ctx() context.Context {
	return tbl.ctx
}

// Dialect gets the database dialect.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Dialect() schema.Dialect {
	return tbl.database.Dialect()
}

// Name gets the table name.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Name() sqlgen2.TableName {
	return tbl.name
}

{{if .Table.HasPrimaryKey}}
// PkColumn gets the column name used as a primary key.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) PkColumn() string {
	return tbl.pk
}

{{end}}
// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) DB() *sql.DB {
	return tbl.db.(*sql.DB)
}

// Execer gets the wrapped database or transaction handle.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Execer() sqlgen2.Execer {
	return tbl.db
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Tx() *sql.Tx {
	return tbl.db.(*sql.Tx)
}

// IsTx tests whether this is within a transaction.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) IsTx() bool {
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
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) BeginTx(opts *sql.TxOptions) ({{.Prefix}}{{.Type}}{{.Thing}}, error) {
	var err error
	tbl.db, err = tbl.db.(sqlgen2.TxStarter).BeginTx(tbl.ctx, opts)
	return tbl, tbl.logIfError(err)
}

// Using returns a modified Table using the transaction supplied. This is needed
// when making multiple queries across several tables within a single transaction.
// The result is a modified copy of the table; the original is unchanged.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Using(tx *sql.Tx) {{.Prefix}}{{.Type}}{{.Thing}} {
	tbl.db = tx
	return tbl
}

func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) logQuery(query string, args ...interface{}) {
	tbl.database.LogQuery(query, args...)
}

func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) logError(err error) error {
	return tbl.database.LogError(err)
}

func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) logIfError(err error) error {
	return tbl.database.LogIfError(err)
}

`

var tTable = template.Must(template.New("Table").Funcs(funcMap).Parse(sTable))

//-------------------------------------------------------------------------------------------------

// function template to scan multiple rows.
const sScanRows = `
func scan{{.Prefix}}{{.Types}}(rows *sql.Rows, firstOnly bool) (vv {{.List}}, n int64, err error) {
	for rows.Next() {
		n++

{{range .Body1}}{{.}}{{- end}}
		err = rows.Scan(
{{- range .Body2}}
			{{.}},
{{- end}}
		)
		if err != nil {
			return vv, n, err
		}

		v := &{{.Type}}{}
{{range .Body3}}{{.}}{{end}}
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
`

var tScanRows = template.Must(template.New("ScanRows").Funcs(funcMap).Parse(sScanRows))

//-------------------------------------------------------------------------------------------------

const sSetter = `
// Set{{.Setter.Name}} sets the {{.Setter.Name}} field and returns the modified {{.Type}}.
func (v *{{.Type}}) Set{{.Setter.Name}}(x {{.Setter.Type.Type}}) *{{.Type}} {
	{{if .Setter.Type.IsPtr -}}
	v.{{.Setter.JoinParts 0 "."}} = &x
{{- else -}}
	v.{{.Setter.JoinParts 0 "."}} = x
{{- end}}
	return v
}
`

var tSetter = template.Must(template.New("SliceRow").Funcs(funcMap).Parse(sSetter))

//-------------------------------------------------------------------------------------------------

const sTruncate = `
// Truncate drops every record from the table, if possible. It might fail if constraints exist that
// prevent some or all rows from being deleted; use the force option to override this.
//
// When 'force' is set true, be aware of the following consequences.
// When using Mysql, foreign keys in other tables can be left dangling.
// When using Postgres, a cascade happens, so all 'adjacent' tables (i.e. linked by foreign keys)
// are also truncated.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Truncate(force bool) (err error) {
	for _, query := range tbl.Dialect().TruncateDDL(tbl.Name().String(), force) {
		_, err = support.Exec(tbl, nil, query)
		if err != nil {
			return err
		}
	}
	return nil
}
`

var tTruncate = template.Must(template.New("Truncate").Funcs(funcMap).Parse(sTruncate))

//-------------------------------------------------------------------------------------------------

// function template to create a table
const sCreateTableFunc = `
// CreateTable creates the table.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) CreateTable(ifNotExists bool) (int64, error) {
	return support.Exec(tbl, nil, tbl.createTableSql(ifNotExists))
}

func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) createTableSql(ifNotExists bool) string {
	var columns string
	var settings string
	switch tbl.Dialect() {
	{{range .Dialects -}}
	case schema.{{.String}}:
		columns = sql{{$.Prefix}}{{$.Type}}{{$.Thing}}CreateColumns{{.}}
		settings = "{{.CreateTableSettings}}"
    {{end -}}
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

func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) ternary(flag bool, a, b string) string {
	if flag {
		return a
	}
	return b
}

// DropTable drops the table, destroying all its data.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) DropTable(ifExists bool) (int64, error) {
	return support.Exec(tbl, nil, tbl.dropTableSql(ifExists))
}

func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) dropTableSql(ifExists bool) string {
	ie := tbl.ternary(ifExists, "IF EXISTS ", "")
	query := fmt.Sprintf("DROP TABLE %s%s", ie, tbl.name)
	return query
}
`
var tCreateTableFunc = template.Must(template.New("CreateTable").Funcs(funcMap).Parse(sCreateTableFunc))

// function template to create DDL for indexes
const sCreateIndexesFunc = `
// CreateTableWithIndexes invokes CreateTable then CreateIndexes.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) CreateTableWithIndexes(ifNotExist bool) (err error) {
	_, err = tbl.CreateTable(ifNotExist)
	if err != nil {
		return err
	}

	return tbl.CreateIndexes(ifNotExist)
}

// CreateIndexes executes queries that create the indexes needed by the {{.Type}} table.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) CreateIndexes(ifNotExist bool) (err error) {
{{range .Table.Index}}
	err = tbl.Create{{camel .Name}}Index(ifNotExist)
	if err != nil {
		return err
	}
{{end}}
	return nil
}
{{range .Table.Index}}
// Create{{camel .Name}}Index creates the {{.Name}} index.
func (tbl {{$.Prefix}}{{$.Type}}{{$.Thing}}) Create{{camel .Name}}Index(ifNotExist bool) error {
	ine := tbl.ternary(ifNotExist && tbl.Dialect() != schema.Mysql, "IF NOT EXISTS ", "")

	// Mysql does not support 'if not exists' on indexes
	// Workaround: use DropIndex first and ignore an error returned if the index didn't exist.

	if ifNotExist && tbl.Dialect() == schema.Mysql {
		// low-level no-logging Exec
		tbl.Execer().ExecContext(tbl.ctx, tbl.drop{{$.Prefix}}{{camel .Name}}IndexSql(false))
		ine = ""
	}

	_, err := tbl.Exec(nil, tbl.create{{$.Prefix}}{{camel .Name}}IndexSql(ine))
	return err
}

func (tbl {{$.Prefix}}{{$.Type}}{{$.Thing}}) create{{$.Prefix}}{{camel .Name}}IndexSql(ifNotExists string) string {
	indexPrefix := tbl.name.PrefixWithoutDot()
	return fmt.Sprintf("CREATE {{.UniqueStr}}INDEX %s%s{{.Name}} ON %s (%s)", ifNotExists, indexPrefix,
		tbl.name, sql{{$.Prefix}}{{camel .Name}}IndexColumns)
}

// Drop{{camel .Name}}Index drops the {{.Name}} index.
func (tbl {{$.Prefix}}{{$.Type}}{{$.Thing}}) Drop{{camel .Name}}Index(ifExists bool) error {
	_, err := tbl.Exec(nil, tbl.drop{{$.Prefix}}{{camel .Name}}IndexSql(ifExists))
	return err
}

func (tbl {{$.Prefix}}{{$.Type}}{{$.Thing}}) drop{{$.Prefix}}{{camel .Name}}IndexSql(ifExists bool) string {
	// Mysql does not support 'if exists' on indexes
	ie := tbl.ternary(ifExists && tbl.Dialect() != schema.Mysql, "IF EXISTS ", "")
	onTbl := tbl.ternary(tbl.Dialect() == schema.Mysql, fmt.Sprintf(" ON %s", tbl.name), "")
	indexPrefix := tbl.name.PrefixWithoutDot()
	return fmt.Sprintf("DROP INDEX %s%s{{.Name}}%s", ie, indexPrefix, onTbl)
}
{{end}}
// DropIndexes executes queries that drop the indexes on by the {{.Type}} table.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) DropIndexes(ifExist bool) (err error) {
{{range .Table.Index}}
	err = tbl.Drop{{camel .Name}}Index(ifExist)
	if err != nil {
		return err
	}
{{end}}
	return nil
}
`

var tCreateIndexesFunc = template.Must(template.New("CreateIndex").Funcs(funcMap).Parse(sCreateIndexesFunc))
