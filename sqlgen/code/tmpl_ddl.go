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
	name    sqlgen2.TableName
	db      sqlgen2.Execer
	ctx     context.Context
	dialect schema.Dialect
	logger  *log.Logger
}

// Type conformance check
var _ {{.Interface}} = &{{.Prefix}}{{.Type}}{{.Thing}}{}

// New{{.Prefix}}{{.Type}}{{.Thing}} returns a new table instance.
// If a blank table name is supplied, the default name "{{.DbName}}" will be used instead.
// The request context is initialised with the background.
func New{{.Prefix}}{{.Type}}{{.Thing}}(name sqlgen2.TableName, d sqlgen2.Execer, dialect schema.Dialect) {{.Prefix}}{{.Type}}{{.Thing}} {
	if name.Name == "" {
		name.Name = "{{.DbName}}"
	}
	return {{.Prefix}}{{.Type}}{{.Thing}}{name, d, context.Background(), dialect, nil}
}

// CopyTableAs{{.Prefix}}{{.Type}}{{.Thing}} copies a table instance, retaining the name etc but
// providing methods appropriate for '{{.Type}}'.
func CopyTableAs{{.Prefix}}{{.Type}}{{.Thing}}(origin sqlgen2.Table) {{.Prefix}}{{.Type}}{{.Thing}} {
	return {{.Prefix}}{{.Type}}{{.Thing}}{
		name:    origin.Name(),
		db:      origin.DB(),
		ctx:     origin.Ctx(),
		dialect: origin.Dialect(),
		logger:  origin.Logger(),
	}
}

// WithPrefix sets the table name prefix for subsequent queries.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) WithPrefix(pfx string) {{.Prefix}}{{.Type}}{{.Thing}} {
	tbl.name.Prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) WithContext(ctx context.Context) {{.Prefix}}{{.Type}}{{.Thing}} {
	tbl.ctx = ctx
	return tbl
}

// WithLogger sets the logger for subsequent queries.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) WithLogger(logger *log.Logger) {{.Prefix}}{{.Type}}{{.Thing}} {
	tbl.logger = logger
	return tbl
}

// Ctx gets the current request context.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Ctx() context.Context {
	return tbl.ctx
}

// Dialect gets the database dialect.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Dialect() schema.Dialect {
	return tbl.dialect
}

// Logger gets the trace logger.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Logger() *log.Logger {
	return tbl.logger
}

// SetLogger sets the logger for subsequent queries, returning the interface.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) SetLogger(logger *log.Logger) sqlgen2.Table {
	tbl.logger = logger
	return tbl
}

// Name gets the table name.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Name() sqlgen2.TableName {
	return tbl.name
}

// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) DB() *sql.DB {
	return tbl.db.(*sql.DB)
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

// Begin starts a transaction. The default isolation level is dependent on the driver.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) BeginTx(opts *sql.TxOptions) ({{.Prefix}}{{.Type}}{{.Thing}}, error) {
	d := tbl.db.(*sql.DB)
	var err error
	tbl.db, err = d.BeginTx(tbl.ctx, opts)
	return tbl, err
}

func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) logQuery(query string, args ...interface{}) {
	sqlgen2.LogQuery(tbl.logger, query, args...)
}

`

var tTable = template.Must(template.New("Table").Funcs(funcMap).Parse(sTable))

//-------------------------------------------------------------------------------------------------

// function template to scan multiple rows.
const sScanRows = `
// scan{{.Prefix}}{{.Types}} reads table records into a slice of values.
func scan{{.Prefix}}{{.Types}}(rows *sql.Rows, firstOnly bool) ({{.List}}, error) {
	var err error
	var vv {{.List}}

{{range .Body1}}{{.}}{{- end}}
	for rows.Next() {
		err = rows.Scan(
{{- range .Body2}}
			{{.}},
{{- end}}
		)
		if err != nil {
			return vv, err
		}

		v := &{{.Type}}{}
{{range .Body3}}{{.}}{{end}}
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
`

var tScanRows = template.Must(template.New("ScanRows").Funcs(funcMap).Parse(sScanRows))

//-------------------------------------------------------------------------------------------------

const sSliceRow = `
func slice{{.Prefix}}{{.Type}}{{.Suffix}}(v *{{.Type}}) ([]interface{}, error) {
{{range .Body1}}{{.}}{{- end}}
{{range .Body2}}{{.}}{{- end}}
	return []interface{}{
{{range .Body3}}{{.}}{{- end}}
	}, nil
}
`

var tSliceRow = template.Must(template.New("SliceRow").Funcs(funcMap).Parse(sSliceRow))

//-------------------------------------------------------------------------------------------------

const sSetter = `
// Set{{.Setter.Name}} sets the {{.Setter.Name}} field and returns the modified {{.Type}}.
func (v *{{.Type}}) Set{{.Setter.Name}}(x {{.Setter.Type.Type}}) *{{.Type}} {
	{{if .Setter.Type.IsNullable -}}
	v.{{.Setter.Name}} = &x
{{- else -}}
	v.{{.Setter.Name}} = x
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
	for _, query := range tbl.dialect.TruncateDDL(tbl.Name().String(), force) {
		_, err = tbl.Exec(query)
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
	return tbl.Exec(tbl.createTableSql(ifNotExists))
}

func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) createTableSql(ifNotExists bool) string {
	var stmt string
	switch tbl.dialect {
	{{range .Dialects -}}
	case schema.{{.}}: stmt = sqlCreate{{$.Prefix}}{{$.Type}}{{$.Thing}}{{.}}
    {{end -}}
	}
	extra := tbl.ternary(ifNotExists, "IF NOT EXISTS ", "")
	query := fmt.Sprintf(stmt, extra, tbl.name)
	return query
}

func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) ternary(flag bool, a, b string) string {
	if flag {
		return a
	}
	return b
}

// DropTable drops the table, destroying all its data.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) DropTable(ifExists bool) (int64, error) {
	return tbl.Exec(tbl.dropTableSql(ifExists))
}

func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) dropTableSql(ifExists bool) string {
	extra := tbl.ternary(ifExists, "IF EXISTS ", "")
	query := fmt.Sprintf("DROP TABLE %s%s", extra, tbl.name)
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
	ine := tbl.ternary(ifNotExist && tbl.dialect != schema.Mysql, "IF NOT EXISTS ", "")

	// Mysql does not support 'if not exists' on indexes
	// Workaround: use DropIndex first and ignore an error returned if the index didn't exist.

	if ifNotExist && tbl.dialect == schema.Mysql {
		tbl.Drop{{camel .Name}}Index(false)
		ine = ""
	}

	_, err := tbl.Exec(tbl.create{{$.Prefix}}{{camel .Name}}IndexSql(ine))
	return err
}

func (tbl {{$.Prefix}}{{$.Type}}{{$.Thing}}) create{{$.Prefix}}{{camel .Name}}IndexSql(ifNotExists string) string {
	indexPrefix := tbl.name.PrefixWithoutDot()
	return fmt.Sprintf("CREATE {{.UniqueStr}}INDEX %s%s{{.Name}} ON %s (%s)", ifNotExists, indexPrefix,
		tbl.name, sql{{$.Prefix}}{{camel .Name}}IndexColumns)
}

// Drop{{camel .Name}}Index drops the {{.Name}} index.
func (tbl {{$.Prefix}}{{$.Type}}{{$.Thing}}) Drop{{camel .Name}}Index(ifExists bool) error {
	_, err := tbl.Exec(tbl.drop{{$.Prefix}}{{camel .Name}}IndexSql(ifExists))
	return err
}

func (tbl {{$.Prefix}}{{$.Type}}{{$.Thing}}) drop{{$.Prefix}}{{camel .Name}}IndexSql(ifExists bool) string {
	// Mysql does not support 'if exists' on indexes
	ie := tbl.ternary(ifExists && tbl.dialect != schema.Mysql, "IF EXISTS ", "")
	onTbl := tbl.ternary(tbl.dialect == schema.Mysql, fmt.Sprintf(" ON %s", tbl.name), "")
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
