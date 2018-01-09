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
	prefix, name string
	db           sqlgen2.Execer
	ctx          context.Context
	dialect      schema.Dialect
	logger       *log.Logger
}

// Type conformance check
var _ {{.Interface}} = &{{.Prefix}}{{.Type}}{{.Thing}}{}

// New{{.Prefix}}{{.Type}}{{.Thing}} returns a new table instance.
// If a blank table name is supplied, the default name "{{.DbName}}" will be used instead.
// The table name prefix is initially blank and the request context is the background.
func New{{.Prefix}}{{.Type}}{{.Thing}}(name string, d sqlgen2.Execer, dialect schema.Dialect) {{.Prefix}}{{.Type}}{{.Thing}} {
	if name == "" {
		name = "{{.DbName}}"
	}
	return {{.Prefix}}{{.Type}}{{.Thing}}{"", name, d, context.Background(), dialect, nil}
}

{{if ne .Thing "Table" -}}
// CopyTableAs{{.Prefix}}{{.Type}}{{.Thing}} copies a table instance, retaining the name etc but
// providing methods appropriate for '{{.Type}}'.
func CopyTableAs{{.Prefix}}{{.Type}}{{.Thing}}(origin sqlgen2.Table) {{.Prefix}}{{.Type}}{{.Thing}} {
	return {{.Prefix}}{{.Type}}{{.Thing}}{
		prefix:  origin.Prefix(),
		name:    origin.Name(),
		db:      origin.DB(),
		ctx:     origin.Ctx(),
		dialect: origin.Dialect(),
		logger:  origin.Logger(),
	}
}

{{end -}}
// WithPrefix sets the table name prefix for subsequent queries.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) WithPrefix(pfx string) {{.Prefix}}{{.Type}}{{.Thing}} {
	tbl.prefix = pfx
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
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Name() string {
	return tbl.name
}

// Prefix gets the table name prefix.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Prefix() string {
	return tbl.prefix
}

// FullName gets the concatenated prefix and table name.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) FullName() string {
	return tbl.prefix + tbl.name
}

func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) prefixWithoutDot() string {
	last := len(tbl.prefix)-1
	if last > 0 && tbl.prefix[last] == '.' {
		return tbl.prefix[0:last]
	}
	return tbl.prefix
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

const sQueryRows = `
// QueryOne is the low-level access function for one {{.Type}}.
// If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, *{{.Type}} will be nil.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) QueryOne(query string, args ...interface{}) (*{{.Type}}, error) {
	list, err := tbl.doQuery(true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

// Query is the low-level access function for {{.Types}}.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Query(query string, args ...interface{}) ({{.List}}, error) {
	return tbl.doQuery(false, query, args...)
}

func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) doQuery(firstOnly bool, query string, args ...interface{}) ({{.List}}, error) {
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scan{{.Prefix}}{{.Types}}(rows, firstOnly)
}
`

var tQueryRows = template.Must(template.New("QueryRows").Funcs(funcMap).Parse(sQueryRows))

//-------------------------------------------------------------------------------------------------

const sSelectRow = `
// SelectOneSA allows a single {{.Type}} to be obtained from the table that match a 'where' clause
// and some limit.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
// If not found, *{{.Type}} will be nil.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) SelectOneSA(where, orderBy string, args ...interface{}) (*{{.Type}}, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s LIMIT 1", {{.Prefix}}{{.Type}}ColumnNames, tbl.prefix, tbl.name, where, orderBy)
	return tbl.QueryOne(query, args...)
}

// SelectOne allows a single {{.Type}} to be obtained from the sqlgen2.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
// If not found, *Example will be nil.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) SelectOne(where where.Expression, orderBy string) (*{{.Type}}, error) {
	wh, args := where.Build(tbl.dialect)
	return tbl.SelectOneSA(wh, orderBy, args...)
}
`

var tSelectRow = template.Must(template.New("SelectRow").Funcs(funcMap).Parse(sSelectRow))

//-------------------------------------------------------------------------------------------------

const sGetRow = `{{if .Table.Primary}}
//--------------------------------------------------------------------------------

// Get{{.Type}} gets the record with a given primary key value.
// If not found, *{{.Type}} will be nil.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Get{{.Type}}(id {{.Table.Primary.Type.Base.Token}}) (*{{.Type}}, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s WHERE {{.Table.Primary.SqlName}}=?", {{.Prefix}}{{.Type}}ColumnNames, tbl.prefix, tbl.name)
	return tbl.QueryOne(query, id)
}
{{end -}}
`

var tGetRow = template.Must(template.New("GetRow").Funcs(funcMap).Parse(sGetRow))

//-------------------------------------------------------------------------------------------------

const sSliceItem = `
//--------------------------------------------------------------------------------
{{range .Table.SimpleFields}}
// Slice{{.Name}} gets the {{.Name}} column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
func (tbl {{$.Prefix}}{{$.Type}}{{$.Thing}}) Slice{{.Name}}(where where.Expression, orderBy string) ([]{{.Type.Type}}, error) {
	return tbl.get{{.Type.Tag}}list("{{.SqlName}}", where, orderBy)
}
{{end}}
{{range .Table.SimpleFields.DistinctTypes}}
func (tbl {{$.Prefix}}{{$.Type}}{{$.Thing}}) get{{.Tag}}list(sqlname string, where where.Expression, orderBy string) ([]{{.Type}}, error) {
	wh, args := where.Build(tbl.dialect)
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", sqlname, tbl.prefix, tbl.name, wh, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var v {{.Type}}
	list := make([]{{.Type}}, 0, 10)
	for rows.Next() {
		err = rows.Scan(&v)
		if err != nil {
			return list, err
		}
		list = append(list, v)
	}
	return list, nil
}
{{end}}
`

var tSliceItem = template.Must(template.New("SliceItem").Funcs(funcMap).Parse(sSliceItem))

//-------------------------------------------------------------------------------------------------

// function template to select multiple rows.
const sSelectRows = `
// SelectSA allows {{.Types}} to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) SelectSA(where, orderBy string, args ...interface{}) ({{.List}}, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", {{.Prefix}}{{.Type}}ColumnNames, tbl.prefix, tbl.name, where, orderBy)
	return tbl.Query(query, args...)
}

// Select allows {{.Types}} to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'; otherwise use a blank string.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Select(where where.Expression, orderBy string) ({{.List}}, error) {
	wh, args := where.Build(tbl.dialect)
	return tbl.SelectSA(wh, orderBy, args...)
}
`

var tSelectRows = template.Must(template.New("SelectRows").Funcs(funcMap).Parse(sSelectRows))

//-------------------------------------------------------------------------------------------------

const sCountRows = `
// CountSA counts {{.Types}} in the table that match a 'where' clause.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) CountSA(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s%s %s", tbl.prefix, tbl.name, where)
	tbl.logQuery(query, args...)
	row := tbl.db.QueryRowContext(tbl.ctx, query, args...)
	err = row.Scan(&count)
	return count, err
}

// Count counts the {{.Types}} in the table that match a 'where' clause.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Count(where where.Expression) (count int64, err error) {
	wh, args := where.Build(tbl.dialect)
	return tbl.CountSA(wh, args...)
}
`

var tCountRows = template.Must(template.New("CountRows").Funcs(funcMap).Parse(sCountRows))

//-------------------------------------------------------------------------------------------------

// function template to insert a single row, updating the primary key in the struct.
const sInsertAndGetLastId = `
// Insert adds new records for the {{.Types}}. The {{.Types}} have their primary key fields
// set to the new record identifiers.
// The {{.Type}}.PreInsert(Execer) method will be called, if it exists.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Insert(vv ...*{{.Type}}) error {
	var params string
	switch tbl.dialect {
	case schema.Postgres:
		params = s{{$.Prefix}}{{$.Type}}DataColumnParamsPostgres
	default:
		params = s{{$.Prefix}}{{$.Type}}DataColumnParamsSimple
	}

	query := fmt.Sprintf(sqlInsert{{$.Prefix}}{{$.Type}}, tbl.prefix, tbl.name, params)
	st, err := tbl.db.PrepareContext(tbl.ctx, query)
	if err != nil {
		return err
	}
	defer st.Close()

	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPreInsert); ok {
			hook.PreInsert()
		}

		fields, err := slice{{.Prefix}}{{.Type}}WithoutPk(v)
		if err != nil {
			return err
		}

		tbl.logQuery(query, fields...)
		res, err := st.Exec(fields...)
		if err != nil {
			return err
		}

		{{if eq .Table.Primary.Type.Name "int64" -}}
		v.{{.Table.Primary.Name}}, err = res.LastInsertId()
		{{- else -}}
		_i64, err := res.LastInsertId()
		v.{{.Table.Primary.Name}} = {{.Table.Primary.Type.Name}}(_i64)
		{{end}}
		if err != nil {
			return err
		}
	}

	return nil
}
`

var tInsertAndGetLastId = template.Must(template.New("InsertAndGetLastId").Funcs(funcMap).Parse(sInsertAndGetLastId))

//-------------------------------------------------------------------------------------------------

// function template to insert a single row.
const sInsertSimple = `
// Insert adds new records for the {{.Types}}.
// The {{.Type}}.PreInsert(Execer) method will be called, if it exists.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Insert(vv ...*{{.Type}}) error {
	var params string
	switch tbl.dialect {
	case schema.Postgres:
		params = s{{$.Prefix}}{{$.Type}}DataColumnParamsPostgres
	default:
		params = s{{$.Prefix}}{{$.Type}}DataColumnParamsSimple
	}

	query := fmt.Sprintf(sqlInsert{{$.Prefix}}{{$.Type}}, tbl.prefix, tbl.name, params)
	st, err := tbl.db.PrepareContext(tbl.ctx, query)
	if err != nil {
		return err
	}
	defer st.Close()

	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPreInsert); ok {
			hook.PreInsert()
		}

		fields, err := slice{{.Prefix}}{{.Type}}(v)
		if err != nil {
			return err
		}

		tbl.logQuery(query, fields...)
		_, err = st.Exec(fields...)
		if err != nil {
			return err
		}
	}

	return nil
}
`

var tInsertSimple = template.Must(template.New("Insert").Funcs(funcMap).Parse(sInsertSimple))

//-------------------------------------------------------------------------------------------------

// function template to update a single row.
const sUpdateFields = `
// UpdateFields updates one or more columns, given a 'where' clause.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) UpdateFields(where where.Expression, fields ...sql.NamedArg) (int64, error) {
	query, args := tbl.updateFields(where, fields...)
	return tbl.Exec(query, args...)
}

func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) updateFields(where where.Expression, fields ...sql.NamedArg) (string, []interface{}) {
	list := sqlgen2.NamedArgList(fields)
	assignments := strings.Join(list.Assignments(tbl.dialect, 1), ", ")
	whereClause, wargs := where.Build(tbl.dialect)
	query := fmt.Sprintf("UPDATE %s%s SET %s %s", tbl.prefix, tbl.name, assignments, whereClause)
	args := append(list.Values(), wargs...)
	return query, args
}
`

var tUpdateFields = template.Must(template.New("UpdateFields").Funcs(funcMap).Parse(sUpdateFields))

//-------------------------------------------------------------------------------------------------

// function template to update rows.
const sUpdate = `{{if .Table.Primary}}
//--------------------------------------------------------------------------------

// Update updates records, matching them by primary key. It returns the number of rows affected.
// The {{.Type}}.PreUpdate(Execer) method will be called, if it exists.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Update(vv ...*{{.Type}}) (int64, error) {
	var stmt string
	switch tbl.dialect {
	case schema.Postgres:
		stmt = sqlUpdate{{$.Prefix}}{{$.Type}}ByPkPostgres
	default:
		stmt = sqlUpdate{{$.Prefix}}{{$.Type}}ByPkSimple
	}

	query := fmt.Sprintf(stmt, tbl.prefix, tbl.name)

	var count int64
	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPreUpdate); ok {
			hook.PreUpdate()
		}

		args, err := slice{{.Prefix}}{{.Type}}WithoutPk(v)
		if err != nil {
			return count, err
		}

		args = append(args, v.{{.Table.Primary.Name}})
		tbl.logQuery(query, args...)
		n, err := tbl.Exec(query, args...)
		if err != nil {
			return count, err
		}
		count += n
	}
	return count, nil
}
{{end -}}
`

var tUpdate = template.Must(template.New("Update").Funcs(funcMap).Parse(sUpdate))

//-------------------------------------------------------------------------------------------------

const sDelete = `
// Delete deletes one or more rows from the table, given a 'where' clause.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Delete(where where.Expression) (int64, error) {
	query, args := tbl.deleteRows(where)
	return tbl.Exec(query, args...)
}

func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) deleteRows(where where.Expression) (string, []interface{}) {
	whereClause, args := where.Build(tbl.dialect)
	query := fmt.Sprintf("DELETE FROM %s%s %s", tbl.prefix, tbl.name, whereClause)
	return query, args
}
`

var tDelete = template.Must(template.New("Delete").Funcs(funcMap).Parse(sDelete))

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
	for _, query := range tbl.dialect.TruncateDDL(tbl.FullName(), force) {
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

const sExec = `
// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// It returns the number of rows affected (of the database drive supports this).
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Exec(query string, args ...interface{}) (int64, error) {
	tbl.logQuery(query, args...)
	res, err := tbl.db.ExecContext(tbl.ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
`

var tExec = template.Must(template.New("Exec").Funcs(funcMap).Parse(sExec))

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
	query := fmt.Sprintf(stmt, extra, tbl.prefix, tbl.name)
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
	query := fmt.Sprintf("DROP TABLE %s%s%s", extra, tbl.prefix, tbl.name)
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
	indexPrefix := tbl.prefixWithoutDot()
	return fmt.Sprintf("CREATE {{.UniqueStr}}INDEX %s%s{{.Name}} ON %s%s (%s)", ifNotExists, indexPrefix,
		tbl.prefix, tbl.name, sql{{$.Prefix}}{{camel .Name}}IndexColumns)
}

// Drop{{camel .Name}}Index drops the {{.Name}} index.
func (tbl {{$.Prefix}}{{$.Type}}{{$.Thing}}) Drop{{camel .Name}}Index(ifExists bool) error {
	_, err := tbl.Exec(tbl.drop{{$.Prefix}}{{camel .Name}}IndexSql(ifExists))
	return err
}

func (tbl {{$.Prefix}}{{$.Type}}{{$.Thing}}) drop{{$.Prefix}}{{camel .Name}}IndexSql(ifExists bool) string {
	// Mysql does not support 'if exists' on indexes
	ie := tbl.ternary(ifExists && tbl.dialect != schema.Mysql, "IF EXISTS ", "")
	onTbl := tbl.ternary(tbl.dialect == schema.Mysql, fmt.Sprintf(" ON %s%s", tbl.prefix, tbl.name), "")
	indexPrefix := tbl.prefixWithoutDot()
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
