package code

import "text/template"

// template to declare the package name.
var sPackage = `// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.
// sqlapi %s; sqlgen %s

package %s
`

//-------------------------------------------------------------------------------------------------

const sTabler = `
// {{.Prefix}}{{.Type}}{{.Thinger}} lists table methods provided by {{.Prefix}}{{.Type}}{{.Thing}}.
type {{.Prefix}}{{.Type}}{{.Thinger}} interface {
	{{.Sqlapi}}.Table

	// Constraints returns the table's constraints.
	Constraints() constraint.Constraints

	// WithConstraint returns a modified {{.Prefix}}{{.Type}}{{.Thinger}} with added data consistency constraints.
	WithConstraint(cc ...constraint.Constraint) {{.Prefix}}{{.Type}}{{.Thinger}}

	// WithPrefix returns a modified {{.Prefix}}{{.Type}}{{.Thinger}} with a given table name prefix.
	WithPrefix(pfx string) {{.Prefix}}{{.Type}}{{.Thinger}}

	// WithContext returns a modified {{.Prefix}}{{.Type}}{{.Thinger}} with a given context.
	WithContext(ctx context.Context) {{.Prefix}}{{.Type}}{{.Thinger}}
`

const sQueryer = `
// {{.Prefix}}{{.Type}}Queryer lists query methods provided by {{.Prefix}}{{.Type}}{{.Thing}}.
type {{.Prefix}}{{.Type}}Queryer interface {
	// Using returns a modified {{.Prefix}}{{.Type}}{{.Thinger}} using the transaction supplied.
	Using(tx {{.Sqlapi}}.SqlTx) {{.Prefix}}{{.Type}}Queryer

	// Transact runs the function provided within a transaction.
	Transact(txOptions *{{.Sql}}.TxOptions, fn func({{.Prefix}}{{.Type}}Queryer) error) error

	// Tx gets the wrapped transaction handle, provided this is within a transaction.
	// Panics if it is in the wrong state - use IsTx() if necessary.
	Tx() {{.Sqlapi}}.SqlTx

	// IsTx tests whether this is within a transaction.
	IsTx() bool
`

var tTabler = template.Must(template.New("Tabler").Funcs(funcMap).Parse(sTabler))
var tQueryer = template.Must(template.New("Tabler").Funcs(funcMap).Parse(sQueryer))

//-------------------------------------------------------------------------------------------------

const sTable = `
// {{.Prefix}}{{.Type}}{{.Thing}} holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type {{.Prefix}}{{.Type}}{{.Thing}} struct {
	name        {{.Sqlapi}}.TableName
	database    {{.Sqlapi}}.Database
	db          {{.Sqlapi}}.Execer
	constraints constraint.Constraints
	ctx         context.Context
	pk          string
}

// Type conformance checks
var _ {{.Interface1}} = &{{.Prefix}}{{.Type}}{{.Thing}}{}

// New{{.Prefix}}{{.Type}}{{.Thing}} returns a new table instance.
// If a blank table name is supplied, the default name "{{.DbName}}" will be used instead.
// The request context is initialised with the background.
func New{{.Prefix}}{{.Type}}{{.Thing}}(name string, d {{.Sqlapi}}.Database) {{.Prefix}}{{.Type}}{{.Thing}} {
	if name == "" {
		name = "{{.DbName}}"
	}
	var constraints constraint.Constraints
	{{range .Constraints -}}
	constraints = append(constraints,
		{{.GoString}})

	{{end -}}
	return {{.Prefix}}{{.Type}}{{.Thing}}{
		name:        {{.Sqlapi}}.TableName{Prefix: "", Name: name},
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
func CopyTableAs{{title .Prefix}}{{title .Type}}{{.Thing}}(origin {{.Sqlapi}}.Table) {{.Prefix}}{{.Type}}{{.Thing}} {
	return {{.Prefix}}{{.Type}}{{.Thing}}{
		name:        origin.Name(),
		database:    origin.Database(),
		db:          origin.DB(),
		constraints: nil,
		ctx:         context.Background(),
		pk:          "{{.Table.SafePrimary.SqlName}}",
	}
}
{{- if .Table.HasPrimaryKey}}

// SetPkColumn sets the name of the primary key column. It defaults to "{{.Table.Primary.SqlName}}".
// The result is a modified copy of the table; the original is unchanged.
//func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) SetPkColumn(pk string) {{.Prefix}}{{.Type}}{{.Thinger}} {
//	tbl.pk = pk
//	return tbl
//}
{{- end}}

// WithPrefix sets the table name prefix for subsequent queries.
// The result is a modified copy of the table; the original is unchanged.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) WithPrefix(pfx string) {{.Prefix}}{{.Type}}{{.Thinger}} {
	tbl.name.Prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries via this table.
// The result is a modified copy of the table; the original is unchanged.
//
// The shared context in the *Database is not altered by this method. So it
// is possible to use different contexts for different (groups of) queries.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) WithContext(ctx context.Context) {{.Prefix}}{{.Type}}{{.Thinger}} {
	tbl.ctx = ctx
	return tbl
}

// Database gets the shared database information.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Database() {{.Sqlapi}}.Database {
	return tbl.database
}

// Logger gets the trace logger.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Logger() {{.Sqlapi}}.Logger {
	return tbl.database.Logger()
}

// WithConstraint returns a modified {{.Prefix}}{{.Type}}{{.Thinger}} with added data consistency constraints.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) WithConstraint(cc ...constraint.Constraint) {{.Prefix}}{{.Type}}{{.Thinger}} {
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
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Dialect() dialect.Dialect {
	return tbl.database.Dialect()
}

// Name gets the table name.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Name() {{.Sqlapi}}.TableName {
	return tbl.name
}
{{- if .Table.HasPrimaryKey}}

// PkColumn gets the column name used as a primary key.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) PkColumn() string {
	return tbl.pk
}
{{- end}}

// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) DB() {{.Sqlapi}}.SqlDB {
	return tbl.db.({{.Sqlapi}}.SqlDB)
}

// Execer gets the wrapped database or transaction handle.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Execer() {{.Sqlapi}}.Execer {
	return tbl.db
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Tx() {{.Sqlapi}}.SqlTx {
	return tbl.db.({{.Sqlapi}}.SqlTx)
}

// IsTx tests whether this is within a transaction.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) IsTx() bool {
	return tbl.db.IsTx()
}

// Using returns a modified {{.Prefix}}{{.Type}}{{.Thinger}} using the transaction supplied. This is
// needed when making multiple queries across several tables within a single transaction.
//
// The result is a modified copy of the table; the original is unchanged.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Using(tx {{.Sqlapi}}.SqlTx) {{.Prefix}}{{.Type}}Queryer {
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
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Transact(txOptions *{{.Sql}}.TxOptions, fn func({{.Prefix}}{{.Type}}Queryer) error) error {
	var err error
	if tbl.IsTx() {
		err = fn(tbl) // nested transactions are inlined
	} else {
		err = tbl.DB().Transact(tbl.ctx, txOptions, func(tx {{.Sqlapi}}.SqlTx) error {
			return fn(tbl.Using(tx))
		})
	}
	return tbl.Logger().LogIfError(err)
}

func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) quotedName() string {
	return tbl.Dialect().Quoter().Quote(tbl.name.String())
}

func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) quotedNameW(w dialect.StringWriter) {
	tbl.Dialect().Quoter().QuoteW(w, tbl.name.String())
}
`

var tTable = template.Must(template.New("Table").Funcs(funcMap).Parse(sTable))

//-------------------------------------------------------------------------------------------------

// function template to scan multiple rows.
const sScanRows = `
// {{.Scan}}{{.Prefix}}{{.Types}} reads rows from the database and returns a slice of corresponding values.
// It also returns a number indicating how many rows were read; this will be larger than the length of the
// slice if reading stopped after the first row.
func {{.Scan}}{{.Prefix}}{{.Types}}(query string, rows {{.Sqlapi}}.SqlRows, firstOnly bool) (vv {{.List}}, n int64, err error) {
	for rows.Next() {
		n++

{{range .Body1}}{{.}}{{- end}}
		err = rows.Scan(
{{- range .Body2}}
			{{.}},
{{- end}}
		)
		if err != nil {
			return vv, n, errors.Wrap(err, query)
		}

		v := &{{.TypePkg}}{{.Type}}{}
{{range .Body3}}{{.}}{{end}}
		var iv interface{} = v
		if hook, ok := iv.({{.Sqlapi}}.CanPostGet); ok {
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
`

var tScanRows = template.Must(template.New("ScanRows").Funcs(funcMap).Parse(sScanRows))

//-------------------------------------------------------------------------------------------------

const sSetterDecl = ``

const sSetterFunc = `
// Set{{.Setter.Name}} sets the {{.Setter.Name}} field and returns the modified {{.TypePkg}}{{.Type}}.
func (v *{{.Type}}) Set{{.Setter.Name}}(x {{.Setter.Type.Type}}) *{{.TypePkg}}{{.Type}} {
	{{if .Setter.Type.IsPtr -}}
	v.{{.Setter.JoinParts 0 "."}} = &x
{{- else -}}
	v.{{.Setter.JoinParts 0 "."}} = x
{{- end}}
	return v
}
`

var tSetterDecl = template.Must(template.New("SetterDecl").Funcs(funcMap).Parse(sSetterDecl))
var tSetterFunc = template.Must(template.New("SetterFunc").Funcs(funcMap).Parse(sSetterFunc))

//-------------------------------------------------------------------------------------------------

const sTruncateDecl = `
	// Truncate drops every record from the table, if possible.
	Truncate(force bool) (err error)
`

const sTruncateFunc = `
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

var tTruncateDecl = template.Must(template.New("TruncateDecl").Funcs(funcMap).Parse(sTruncateDecl))
var tTruncateFunc = template.Must(template.New("TruncateFunc").Funcs(funcMap).Parse(sTruncateFunc))

//-------------------------------------------------------------------------------------------------

const sCreateTableDecl = `
	// CreateTable creates the table.
	CreateTable(ifNotExists bool) (int64, error)

	// DropTable drops the table, destroying all its data.
	DropTable(ifExists bool) (int64, error)
`

// function template to create a table
const sCreateTableFunc = `
// CreateTable creates the table.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) CreateTable(ifNotExists bool) (int64, error) {
	return support.Exec(tbl, nil, create{{.Prefix}}{{.Type}}{{.Thing}}Sql(tbl, ifNotExists))
}

func create{{.Prefix}}{{.Type}}{{.Thing}}Sql(tbl {{.Prefix}}{{.Type}}{{.Thinger}}, ifNotExists bool) string {
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
	{{- range .Dialects}}
	case dialect.{{.String}}Index:
		columns = sql{{$.Prefix}}{{$.Type}}{{$.Thing}}CreateColumns{{.}}
    {{- end}}
	}

	comma := ""
	for i, n := range listOf{{$.Prefix}}{{$.Type}}{{$.Thing}}ColumnNames {
		buf.WriteString(comma)
		q.QuoteW(buf, n)
		buf.WriteString(" ")
		buf.WriteString(columns[i])
		comma = ",\n "
	}

	for i, c := range tbl.({{.Prefix}}{{.Type}}{{.Thing}}).Constraints() {
		buf.WriteString(",\n ")
		buf.WriteString(c.ConstraintSql(tbl.Dialect().Quoter(), tbl.Name(), i+1))
	}

	buf.WriteString("\n)")
	buf.WriteString(tbl.Dialect().CreateTableSettings())
	return buf.String()
}

func ternary{{.Prefix}}{{.Type}}{{.Thing}}(flag bool, a, b string) string {
	if flag {
		return a
	}
	return b
}

// DropTable drops the table, destroying all its data.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) DropTable(ifExists bool) (int64, error) {
	return support.Exec(tbl, nil, drop{{.Prefix}}{{.Type}}{{.Thing}}Sql(tbl, ifExists))
}

func drop{{.Prefix}}{{.Type}}{{.Thing}}Sql(tbl {{.Prefix}}{{.Type}}{{.Thinger}}, ifExists bool) string {
	ie := ternary{{.Prefix}}{{.Type}}{{.Thing}}(ifExists, "IF EXISTS ", "")
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	query := fmt.Sprintf("DROP TABLE %s%s", ie, quotedName)
	return query
}
`

var tCreateTableDecl = template.Must(template.New("CreateTableDecl").Funcs(funcMap).Parse(sCreateTableDecl))
var tCreateTableFunc = template.Must(template.New("CreateTableFunc").Funcs(funcMap).Parse(sCreateTableFunc))

const sCreateIndexesDecl = `
	// CreateTableWithIndexes invokes CreateTable then CreateIndexes.
	CreateTableWithIndexes(ifNotExist bool) (err error)

	// CreateIndexes executes queries that create the indexes needed by the {{.Type}} table.
	CreateIndexes(ifNotExist bool) (err error)
{{- range .Table.Index}}

	// Create{{camel .Name}}Index creates the {{.Name}} index.
	Create{{camel .Name}}Index(ifNotExist bool) error

	// Drop{{camel .Name}}Index drops the {{.Name}} index.
	Drop{{camel .Name}}Index(ifExists bool) error
{{- end}}
`

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
	ine := ternary{{$.Prefix}}{{$.Type}}{{$.Thing}}(ifNotExist && tbl.Dialect().Index() != dialect.MysqlIndex, "IF NOT EXISTS ", "")

	// Mysql does not support 'if not exists' on indexes
	// Workaround: use DropIndex first and ignore an error returned if the index didn't exist.

	if ifNotExist && tbl.Dialect().Index() == dialect.MysqlIndex {
		// low-level no-logging Exec
		tbl.Execer().ExecContext(tbl.ctx, drop{{$.Prefix}}{{$.Type}}{{$.Thing}}{{camel .Name}}Sql(tbl, false))
		ine = ""
	}

	_, err := tbl.Exec(nil, create{{$.Prefix}}{{$.Type}}{{$.Thing}}{{camel .Name}}Sql(tbl, ine))
	return err
}

func create{{$.Prefix}}{{$.Type}}{{$.Thing}}{{camel .Name}}Sql(tbl {{$.Prefix}}{{$.Type}}{{$.Thinger}}, ifNotExists string) string {
	indexPrefix := tbl.Name().PrefixWithoutDot()
	id := fmt.Sprintf("%s%s_{{.Name}}", indexPrefix, tbl.Name().Name)
	q := tbl.Dialect().Quoter()
	cols := strings.Join(q.QuoteN(listOf{{$.Prefix}}{{camel .Name}}IndexColumns), ",")
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	return fmt.Sprintf("CREATE {{.UniqueStr}}INDEX %s%s ON %s (%s)", ifNotExists,
		q.Quote(id), quotedName, cols)
}

// Drop{{camel .Name}}Index drops the {{.Name}} index.
func (tbl {{$.Prefix}}{{$.Type}}{{$.Thing}}) Drop{{camel .Name}}Index(ifExists bool) error {
	_, err := tbl.Exec(nil, drop{{$.Prefix}}{{$.Type}}{{$.Thing}}{{camel .Name}}Sql(tbl, ifExists))
	return err
}

func drop{{$.Prefix}}{{$.Type}}{{$.Thing}}{{camel .Name}}Sql(tbl {{$.Prefix}}{{$.Type}}{{$.Thinger}}, ifExists bool) string {
	// Mysql does not support 'if exists' on indexes
	ie := ternary{{$.Prefix}}{{$.Type}}{{$.Thing}}(ifExists && tbl.Dialect().Index() != dialect.MysqlIndex, "IF EXISTS ", "")
	indexPrefix := tbl.Name().PrefixWithoutDot()
	id := fmt.Sprintf("%s%s_{{.Name}}", indexPrefix, tbl.Name().Name)
	q := tbl.Dialect().Quoter()
	// Mysql requires extra "ON tbl" clause
	quotedName := tbl.Dialect().Quoter().Quote(tbl.Name().String())
	onTbl := ternary{{$.Prefix}}{{$.Type}}{{$.Thing}}(tbl.Dialect().Index() == dialect.MysqlIndex, fmt.Sprintf(" ON %s", quotedName), "")
	return "DROP INDEX " + ie + q.Quote(id) + onTbl
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

var tCreateIndexesDecl = template.Must(template.New("CreateIndexDecl").Funcs(funcMap).Parse(sCreateIndexesDecl))
var tCreateIndexesFunc = template.Must(template.New("CreateIndexFunc").Funcs(funcMap).Parse(sCreateIndexesFunc))
