package code

import "text/template"

// template to declare the package name.
var sPackage = `// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

package %s
`

//-------------------------------------------------------------------------------------------------

const sTable = `
// {{.Prefix}}{{.Type}}TableName is the default name for this table.
const {{.Prefix}}{{.Type}}TableName = "{{.DbName}}"

// {{.Prefix}}{{.Type}}Table holds a given table name with the database reference, providing access methods below.
// The Prefix field is often blank but can be used to hold a table name prefix (e.g. ending in '_'). Or it can
// specify the name of the schema, in which case it should have a trailing '.'.
type {{.Prefix}}{{.Type}}Table struct {
	Prefix, Name string
	Db           sqlgen2.Execer
	Ctx          context.Context
	Dialect      schema.Dialect
	Logger       *log.Logger
}

// Type conformance check
var _ sqlgen2.Table = &{{.Prefix}}{{.Type}}Table{}

// New{{.Prefix}}{{.Type}}Table returns a new table instance.
// If a blank table name is supplied, the default name "{{.DbName}}" will be used instead.
// The table name prefix is initially blank and the request context is the background.
func New{{.Prefix}}{{.Type}}Table(name string, d sqlgen2.Execer, dialect schema.Dialect) {{.Prefix}}{{.Type}}Table {
	if name == "" {
		name = {{.Prefix}}{{.Type}}TableName
	}
	return {{.Prefix}}{{.Type}}Table{"", name, d, context.Background(), dialect, nil}
}

// WithPrefix sets the prefix for subsequent queries.
func (tbl {{.Prefix}}{{.Type}}Table) WithPrefix(pfx string) {{.Prefix}}{{.Type}}Table {
	tbl.Prefix = pfx
	return tbl
}

// WithContext sets the context for subsequent queries.
func (tbl {{.Prefix}}{{.Type}}Table) WithContext(ctx context.Context) {{.Prefix}}{{.Type}}Table {
	tbl.Ctx = ctx
	return tbl
}

// WithLogger sets the logger for subsequent queries.
func (tbl {{.Prefix}}{{.Type}}Table) WithLogger(logger *log.Logger) {{.Prefix}}{{.Type}}Table {
	tbl.Logger = logger
	return tbl
}

// SetLogger sets the logger for subsequent queries, returning the interface.
func (tbl {{.Prefix}}{{.Type}}Table) SetLogger(logger *log.Logger) sqlgen2.Table {
	tbl.Logger = logger
	return tbl
}

// FullName gets the concatenated prefix and table name.
func (tbl {{.Prefix}}{{.Type}}Table) FullName() string {
	return tbl.Prefix + tbl.Name
}

func (tbl {{.Prefix}}{{.Type}}Table) prefixWithoutDot() string {
	last := len(tbl.Prefix)-1
	if last > 0 && tbl.Prefix[last] == '.' {
		return tbl.Prefix[0:last]
	}
	return tbl.Prefix
}

// DB gets the wrapped database handle, provided this is not within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl {{.Prefix}}{{.Type}}Table) DB() *sql.DB {
	return tbl.Db.(*sql.DB)
}

// Tx gets the wrapped transaction handle, provided this is within a transaction.
// Panics if it is in the wrong state - use IsTx() if necessary.
func (tbl {{.Prefix}}{{.Type}}Table) Tx() *sql.Tx {
	return tbl.Db.(*sql.Tx)
}

// IsTx tests whether this is within a transaction.
func (tbl {{.Prefix}}{{.Type}}Table) IsTx() bool {
	_, ok := tbl.Db.(*sql.Tx)
	return ok
}

// Begin starts a transaction. The default isolation level is dependent on the driver.
func (tbl {{.Prefix}}{{.Type}}Table) BeginTx(opts *sql.TxOptions) ({{.Prefix}}{{.Type}}Table, error) {
	d := tbl.Db.(*sql.DB)
	var err error
	tbl.Db, err = d.BeginTx(tbl.Ctx, opts)
	return tbl, err
}

func (tbl {{.Prefix}}{{.Type}}Table) logQuery(query string, args ...interface{}) {
	sqlgen2.LogQuery(tbl.Logger, query, args...)
}

`

var tTable = template.Must(template.New("Table").Funcs(funcMap).Parse(sTable))

//-------------------------------------------------------------------------------------------------

// function template to scan a single row.
const sScanRow = `
// scan{{.Prefix}}{{.Type}} reads a table record into a single value.
func scan{{.Prefix}}{{.Type}}(row *sql.Row) (*{{.Type}}, error) {
{{range .Body1}}{{.}}{{- end}}
	err := row.Scan(
{{range .Body2}}{{.}}{{- end}}
	)
	if err != nil {
		return nil, err
	}

	v := &{{.Type}}{}
{{range .Body3}}{{.}}{{- end}}
	return v, nil
}
`

var tScanRow = template.Must(template.New("ScanRow").Funcs(funcMap).Parse(sScanRow))

//-------------------------------------------------------------------------------------------------

// function template to scan multiple rows.
const sScanRows = `
// scan{{.Prefix}}{{.Types}} reads table records into a slice of values.
func scan{{.Prefix}}{{.Types}}(rows *sql.Rows) ({{.List}}, error) {
	var err error
	var vv {{.List}}

{{range .Body1}}{{.}}{{- end}}
	for rows.Next() {
		err = rows.Scan(
{{range .Body2}}{{.}}{{- end}}
		)
		if err != nil {
			return vv, err
		}

		v := &{{.Type}}{}
{{range .Body3}}{{.}}{{end}}
		vv = append(vv, v)
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

const sQueryRow = `
// QueryOne is the low-level access function for one {{.Type}}.
func (tbl {{.Prefix}}{{.Type}}Table) QueryOne(query string, args ...interface{}) (*{{.Type}}, error) {
	tbl.logQuery(query, args...)
	row := tbl.Db.QueryRowContext(tbl.Ctx, query, args...)
	return scan{{.Prefix}}{{.Type}}(row)
}
`

var tQueryRow = template.Must(template.New("SelectRow").Funcs(funcMap).Parse(sQueryRow))

//-------------------------------------------------------------------------------------------------

const sQueryRows = `
// Query is the low-level access function for {{.Types}}.
func (tbl {{.Prefix}}{{.Type}}Table) Query(query string, args ...interface{}) ({{.List}}, error) {
	tbl.logQuery(query, args...)
	rows, err := tbl.Db.QueryContext(tbl.Ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scan{{.Prefix}}{{.Types}}(rows)
}
`

var tQueryRows = template.Must(template.New("SelectRows").Funcs(funcMap).Parse(sQueryRows))

//-------------------------------------------------------------------------------------------------

const sSelectRow = `
// SelectOneSA allows a single {{.Type}} to be obtained from the table that match a 'where' clause and some limit.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
func (tbl {{.Prefix}}{{.Type}}Table) SelectOneSA(where, orderBy string, args ...interface{}) (*{{.Type}}, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s LIMIT 1", {{.Prefix}}{{.Type}}ColumnNames, tbl.Prefix, tbl.Name, where, orderBy)
	return tbl.QueryOne(query, args...)
}

// SelectOne allows a single {{.Type}} to be obtained from the sqlgen2.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
func (tbl {{.Prefix}}{{.Type}}Table) SelectOne(where where.Expression, orderBy string) (*{{.Type}}, error) {
	wh, args := where.Build(tbl.Dialect)
	return tbl.SelectOneSA(wh, orderBy, args...)
}
`

var tSelectRow = template.Must(template.New("SelectRow").Funcs(funcMap).Parse(sSelectRow))

//-------------------------------------------------------------------------------------------------

const sGetRow = `{{if .Table.Primary}}
//--------------------------------------------------------------------------------

// Get gets the record with a given primary key value.
func (tbl {{.Prefix}}{{.Type}}Table) Get(id {{.Table.Primary.Type.Base.Token}}) (*{{.Type}}, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s WHERE {{.Table.Primary.SqlName}}=?", {{.Prefix}}{{.Type}}ColumnNames, tbl.Prefix, tbl.Name)
	return tbl.QueryOne(query, id)
}
{{end -}}
`

var tGetRow = template.Must(template.New("SelectRow").Funcs(funcMap).Parse(sGetRow))

//-------------------------------------------------------------------------------------------------

// function template to select multiple rows.
const sSelectRows = `
// SelectSA allows {{.Types}} to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
func (tbl {{.Prefix}}{{.Type}}Table) SelectSA(where, orderBy string, args ...interface{}) ({{.List}}, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", {{.Prefix}}{{.Type}}ColumnNames, tbl.Prefix, tbl.Name, where, orderBy)
	return tbl.Query(query, args...)
}

// Select allows {{.Types}} to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
func (tbl {{.Prefix}}{{.Type}}Table) Select(where where.Expression, orderBy string) ({{.List}}, error) {
	wh, args := where.Build(tbl.Dialect)
	return tbl.SelectSA(wh, orderBy, args...)
}
`

var tSelectRows = template.Must(template.New("SelectRows").Funcs(funcMap).Parse(sSelectRows))

//-------------------------------------------------------------------------------------------------

const sCountRows = `
// CountSA counts {{.Types}} in the table that match a 'where' clause.
func (tbl {{.Prefix}}{{.Type}}Table) CountSA(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s%s %s", tbl.Prefix, tbl.Name, where)
	tbl.logQuery(query, args...)
	row := tbl.Db.QueryRowContext(tbl.Ctx, query, args...)
	err = row.Scan(&count)
	return count, err
}

// Count counts the {{.Types}} in the table that match a 'where' clause.
func (tbl {{.Prefix}}{{.Type}}Table) Count(where where.Expression) (count int64, err error) {
	wh, args := where.Build(tbl.Dialect)
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
func (tbl {{.Prefix}}{{.Type}}Table) Insert(vv ...*{{.Type}}) error {
	var stmt, params string
	switch tbl.Dialect {
	case schema.Postgres:
		stmt = sqlInsert{{$.Prefix}}{{$.Type}}Postgres
		params = s{{$.Prefix}}{{$.Type}}DataColumnParamsPostgres
	default:
		stmt = sqlInsert{{$.Prefix}}{{$.Type}}Simple
		params = s{{$.Prefix}}{{$.Type}}DataColumnParamsSimple
	}

	query := fmt.Sprintf(stmt, tbl.Prefix, tbl.Name, params)
	st, err := tbl.Db.PrepareContext(tbl.Ctx, query)
	if err != nil {
		return err
	}
	defer st.Close()

	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPreInsert); ok {
			hook.PreInsert(tbl.Db)
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
func (tbl {{.Prefix}}{{.Type}}Table) Insert(vv ...*{{.Type}}) error {
	var stmt, params string
	switch tbl.Dialect {
	case schema.Postgres:
		stmt = sqlInsert{{$.Prefix}}{{$.Type}}Postgres
		params = s{{$.Prefix}}{{$.Type}}DataColumnParamsPostgres
	default:
		stmt = sqlInsert{{$.Prefix}}{{$.Type}}Simple
		params = s{{$.Prefix}}{{$.Type}}DataColumnParamsSimple
	}

	query := fmt.Sprintf(stmt, tbl.Prefix, tbl.Name, params)
	st, err := tbl.Db.PrepareContext(tbl.Ctx, query)
	if err != nil {
		return err
	}
	defer st.Close()

	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPreInsert); ok {
			hook.PreInsert(tbl.Db)
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
const sUpdate = `
// UpdateFields updates one or more columns, given a 'where' clause.
func (tbl {{.Prefix}}{{.Type}}Table) UpdateFields(where where.Expression, fields ...sql.NamedArg) (int64, error) {
	query, args := tbl.updateFields(where, fields...)
	return tbl.Exec(query, args...)
}

func (tbl {{.Prefix}}{{.Type}}Table) updateFields(where where.Expression, fields ...sql.NamedArg) (string, []interface{}) {
	list := sqlgen2.NamedArgList(fields)
	assignments := strings.Join(list.Assignments(tbl.Dialect, 1), ", ")
	whereClause, wargs := where.Build(tbl.Dialect)
	query := fmt.Sprintf("UPDATE %s%s SET %s %s", tbl.Prefix, tbl.Name, assignments, whereClause)
	args := append(list.Values(), wargs...)
	return query, args
}

// Update updates records, matching them by primary key. It returns the number of rows affected.
// The {{.Type}}.PreUpdate(Execer) method will be called, if it exists.
func (tbl {{.Prefix}}{{.Type}}Table) Update(vv ...*{{.Type}}) (int64, error) {
	var stmt string
	switch tbl.Dialect {
	case schema.Postgres:
		stmt = sqlUpdate{{$.Prefix}}{{$.Type}}ByPkPostgres
	default:
		stmt = sqlUpdate{{$.Prefix}}{{$.Type}}ByPkSimple
	}

	var count int64
	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPreUpdate); ok {
			hook.PreUpdate(tbl.Db)
		}

		query := fmt.Sprintf(stmt, tbl.Prefix, tbl.Name)

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
`

var tUpdate = template.Must(template.New("Update").Funcs(funcMap).Parse(sUpdate))

//-------------------------------------------------------------------------------------------------

const sDelete = `
// Delete deletes one or more rows from the table, given a 'where' clause.
func (tbl {{.Prefix}}{{.Type}}Table) Delete(where where.Expression) (int64, error) {
	query, args := tbl.deleteRows(where)
	return tbl.Exec(query, args...)
}

func (tbl {{.Prefix}}{{.Type}}Table) deleteRows(where where.Expression) (string, []interface{}) {
	whereClause, args := where.Build(tbl.Dialect)
	query := fmt.Sprintf("DELETE FROM %s%s %s", tbl.Prefix, tbl.Name, whereClause)
	return query, args
}

// Truncate drops every record from the table, if possible. It might fail if constraints exist that
// prevent some or all rows from being deleted; use the force option to override this.
//
// When 'force' is set true, be aware of the following consequences.
// When using Mysql, foreign keys in other tables can be left dangling.
// When using Postgres, a cascade happens, so all 'adjacent' tables (i.e. linked by foreign keys)
// are also truncated.
func (tbl {{.Prefix}}{{.Type}}Table) Truncate(force bool) (err error) {
	for _, query := range tbl.Dialect.TruncateDDL(tbl.FullName(), force) {
		_, err = tbl.Exec(query)
		if err != nil {
			return err
		}
	}
	return nil
}
`

var tDelete = template.Must(template.New("Delete").Funcs(funcMap).Parse(sDelete))

//-------------------------------------------------------------------------------------------------

// function template to call sql exec
const sExec = `
// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// It returns the number of rows affected (of the database drive supports this).
func (tbl {{.Prefix}}{{.Type}}Table) Exec(query string, args ...interface{}) (int64, error) {
	tbl.logQuery(query, args...)
	res, err := tbl.Db.ExecContext(tbl.Ctx, query, args...)
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
func (tbl {{.Prefix}}{{.Type}}Table) CreateTable(ifNotExist bool) (int64, error) {
	return tbl.Exec(tbl.createTableSql(ifNotExist))
}

func (tbl {{.Prefix}}{{.Type}}Table) createTableSql(ifNotExist bool) string {
	var stmt string
	switch tbl.Dialect {
	{{range .Dialects -}}
	case schema.{{.}}: stmt = sqlCreate{{$.Prefix}}{{$.Type}}Table{{.}}
    {{end -}}
	}
	extra := tbl.ternary(ifNotExist, "IF NOT EXISTS ", "")
	query := fmt.Sprintf(stmt, extra, tbl.Prefix, tbl.Name)
	return query
}

func (tbl {{.Prefix}}{{.Type}}Table) ternary(flag bool, a, b string) string {
	if flag {
		return a
	}
	return b
}
`
var tCreateTableFunc = template.Must(template.New("CreateTable").Funcs(funcMap).Parse(sCreateTableFunc))

// function template to create DDL for indexes
const sCreateIndexesFunc = `
// CreateTableWithIndexes invokes CreateTable then CreateIndexes.
func (tbl {{.Prefix}}{{.Type}}Table) CreateTableWithIndexes(ifNotExist bool) (err error) {
	_, err = tbl.CreateTable(ifNotExist)
	if err != nil {
		return err
	}

	return tbl.CreateIndexes(ifNotExist)
}

// CreateIndexes executes queries that create the indexes needed by the {{.Type}} table.
func (tbl {{.Prefix}}{{.Type}}Table) CreateIndexes(ifNotExist bool) (err error) {
	{{if gt (len .Table.Index) 0 -}}
	ine := tbl.ternary(ifNotExist && tbl.Dialect != schema.Mysql, "IF NOT EXISTS ", "")

	// Mysql does not support 'if not exists' on indexes
	// Workaround: use DropIndex first and ignore an error returned if the index didn't exist.

	if ifNotExist && tbl.Dialect == schema.Mysql {
		tbl.DropIndexes(false)
		ine = ""
	}
{{end -}}
{{range .Table.Index}}
	_, err = tbl.Exec(tbl.create{{$.Prefix}}{{camel .Name}}IndexSql(ine))
	if err != nil {
		return err
	}
{{end}}
	return nil
}
{{range .Table.Index}}
func (tbl {{$.Prefix}}{{$.Type}}Table) create{{$.Prefix}}{{camel .Name}}IndexSql(ifNotExists string) string {
	indexPrefix := tbl.prefixWithoutDot()
	return fmt.Sprintf("CREATE {{.UniqueStr}}INDEX %s%s{{.Name}} ON %s%s (%s)", ifNotExists, indexPrefix,
		tbl.Prefix, tbl.Name, sql{{$.Prefix}}{{camel .Name}}IndexColumns)
}

func (tbl {{$.Prefix}}{{$.Type}}Table) drop{{$.Prefix}}{{camel .Name}}IndexSql(ifExists, onTbl string) string {
	indexPrefix := tbl.prefixWithoutDot()
	return fmt.Sprintf("DROP INDEX %s%s{{.Name}}%s", ifExists, indexPrefix, onTbl)
}
{{end}}
// DropIndexes executes queries that drop the indexes on by the {{.Type}} table.
func (tbl {{.Prefix}}{{.Type}}Table) DropIndexes(ifExist bool) (err error) {
	{{if gt (len .Table.Index) 0 -}}
	// Mysql does not support 'if exists' on indexes
	ie := tbl.ternary(ifExist && tbl.Dialect != schema.Mysql, "IF EXISTS ", "")
	onTbl := tbl.ternary(tbl.Dialect == schema.Mysql, fmt.Sprintf(" ON %s%s", tbl.Prefix, tbl.Name), "")
{{- end}}
{{range .Table.Index}}
	_, err = tbl.Exec(tbl.drop{{$.Prefix}}{{camel .Name}}IndexSql(ie, onTbl))
	if err != nil {
		return err
	}
{{end}}
	return nil
}
`

var tCreateIndexesFunc = template.Must(template.New("CreateIndex").Funcs(funcMap).Parse(sCreateIndexesFunc))
