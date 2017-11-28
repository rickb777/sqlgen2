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
	Dialect      sqlgen2.Dialect
}

// Type conformance check
var _ sqlgen2.Table = {{.Prefix}}{{.Type}}Table{}

// New{{.Prefix}}{{.Type}}Table returns a new table instance.
// If a blank table name is supplied, the default name "{{.DbName}}" will be used instead.
// The table name prefix is initially blank and the request context is the background.
func New{{.Prefix}}{{.Type}}Table(name string, d *sql.DB, dialect sqlgen2.Dialect) {{.Prefix}}{{.Type}}Table {
	if name == "" {
		name = {{.Prefix}}{{.Type}}TableName
	}
	return {{.Prefix}}{{.Type}}Table{"", name, d, context.Background(), dialect}
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

// FullName gets the concatenated prefix and table name.
func (tbl {{.Prefix}}{{.Type}}Table) FullName() string {
	return tbl.Prefix + tbl.Name
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

`

var tTable = template.Must(template.New("Table").Funcs(funcMap).Parse(sTable))

//-------------------------------------------------------------------------------------------------

// function template to scan a single row.
const sScanRow = `
// Scan{{.Prefix}}{{.Type}} reads a table record into a single value.
func Scan{{.Prefix}}{{.Type}}(row *sql.Row) (*{{.Type}}, error) {
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
// Scan{{.Prefix}}{{.Types}} reads table records into a slice of values.
func Scan{{.Prefix}}{{.Types}}(rows *sql.Rows) ([]*{{.Type}}, error) {
	var err error
	var vv []*{{.Type}}

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
func Slice{{.Prefix}}{{.Type}}{{.Suffix}}(v *{{.Type}}) ([]interface{}, error) {
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
	row := tbl.Db.QueryRowContext(tbl.Ctx, query, args...)
	return Scan{{.Prefix}}{{.Type}}(row)
}
`

var tQueryRow = template.Must(template.New("SelectRow").Funcs(funcMap).Parse(sQueryRow))

//-------------------------------------------------------------------------------------------------

const sQueryRows = `
// Query is the low-level access function for {{.Types}}.
func (tbl {{.Prefix}}{{.Type}}Table) Query(query string, args ...interface{}) ([]*{{.Type}}, error) {
	rows, err := tbl.Db.QueryContext(tbl.Ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return Scan{{.Prefix}}{{.Types}}(rows)
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
	return tbl.SelectOneSA(wh, orderBy, args)
}
`

var tSelectRow = template.Must(template.New("SelectRow").Funcs(funcMap).Parse(sSelectRow))

//-------------------------------------------------------------------------------------------------

// function template to select multiple rows.
const sSelectRows = `
// SelectSA allows {{.Types}} to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
func (tbl {{.Prefix}}{{.Type}}Table) SelectSA(where, orderBy string, args ...interface{}) ([]*{{.Type}}, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", {{.Prefix}}{{.Type}}ColumnNames, tbl.Prefix, tbl.Name, where, orderBy)
	return tbl.Query(query, args...)
}

// Select allows {{.Types}} to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
func (tbl {{.Prefix}}{{.Type}}Table) Select(where where.Expression, orderBy string) ([]*{{.Type}}, error) {
	wh, args := where.Build(tbl.Dialect)
	return tbl.SelectSA(wh, orderBy, args)
}
`

var tSelectRows = template.Must(template.New("SelectRows").Funcs(funcMap).Parse(sSelectRows))

//-------------------------------------------------------------------------------------------------

const sCountRows = `
// CountSA counts {{.Types}} in the table that match a 'where' clause.
func (tbl {{.Prefix}}{{.Type}}Table) CountSA(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s%s %s", tbl.Prefix, tbl.Name, where)
	row := tbl.Db.QueryRowContext(tbl.Ctx, query, args)
	err = row.Scan(&count)
	return count, err
}

// Count counts the {{.Types}} in the table that match a 'where' clause.
func (tbl {{.Prefix}}{{.Type}}Table) Count(where where.Expression) (count int64, err error) {
	return tbl.CountSA(where.Build(tbl.Dialect))
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
	case sqlgen2.Postgres:
		stmt = sqlInsert{{$.Prefix}}{{$.Type}}Postgres
		params = s{{$.Prefix}}{{$.Type}}DataColumnParamsPostgres
	default:
		stmt = sqlInsert{{$.Prefix}}{{$.Type}}Simple
		params = s{{$.Prefix}}{{$.Type}}DataColumnParamsSimple
	}

	st, err := tbl.Db.PrepareContext(tbl.Ctx, fmt.Sprintf(stmt, tbl.Prefix, tbl.Name, params))
	if err != nil {
		return err
	}
	defer st.Close()

	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(interface{PreInsert(sqlgen2.Execer)}); ok {
			hook.PreInsert(tbl.Db)
		}

		fields, err := Slice{{.Prefix}}{{.Type}}WithoutPk(v)
		if err != nil {
			return err
		}

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
	case sqlgen2.Postgres:
		stmt = sqlInsert{{$.Prefix}}{{$.Type}}Postgres
		params = s{{$.Prefix}}{{$.Type}}DataColumnParamsPostgres
	default:
		stmt = sqlInsert{{$.Prefix}}{{$.Type}}Simple
		params = s{{$.Prefix}}{{$.Type}}DataColumnParamsSimple
	}

	st, err := tbl.Db.PrepareContext(tbl.Ctx, fmt.Sprintf(stmt, tbl.Prefix, tbl.Name, params))
	if err != nil {
		return err
	}
	defer st.Close()

	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(interface{PreInsert(sqlgen2.Execer)}); ok {
			hook.PreInsert(tbl.Db)
		}

		fields, err := Slice{{.Prefix}}{{.Type}}Stmt(v)
		if err != nil {
			return err
		}

		res, err := st.Exec(fields...)
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
	return tbl.Exec(tbl.updateFields(where, fields...))
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
	case sqlgen2.Postgres:
		stmt = sqlUpdate{{$.Prefix}}{{$.Type}}ByPkPostgres
	default:
		stmt = sqlUpdate{{$.Prefix}}{{$.Type}}ByPkSimple
	}

	var count int64
	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(interface{PreUpdate(sqlgen2.Execer)}); ok {
			hook.PreUpdate(tbl.Db)
		}

		query := fmt.Sprintf(stmt, tbl.Prefix, tbl.Name)

		args, err := Slice{{.Prefix}}{{.Type}}WithoutPk(v)
		if err != nil {
			return count, err
		}

		args = append(args, v.{{.Table.Primary.Name}})
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
	return tbl.Exec(tbl.deleteRows(where))
}

func (tbl {{.Prefix}}{{.Type}}Table) deleteRows(where where.Expression) (string, []interface{}) {
	whereClause, args := where.Build(tbl.Dialect)
	query := fmt.Sprintf("DELETE FROM %s%s %s", tbl.Prefix, tbl.Name, whereClause)
	return query, args
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
	res, err := tbl.Db.ExecContext(tbl.Ctx, query, args...)
	if err != nil {
		return 0, nil
	}
	return res.RowsAffected()
}
`

var tExec = template.Must(template.New("Exec").Funcs(funcMap).Parse(sExec))

//-------------------------------------------------------------------------------------------------

// function template to create a table
const sCreateTable = `
// CreateTable creates the table.
func (tbl {{.Prefix}}{{.Type}}Table) CreateTable(ifNotExist bool) (int64, error) {
	return tbl.Exec(tbl.createTableSql(ifNotExist))
}

func (tbl {{.Prefix}}{{.Type}}Table) createTableSql(ifNotExist bool) string {
	var stmt string
	switch tbl.Dialect {
	{{range .Dialects -}}
	case sqlgen2.{{.}}: stmt = sqlCreate{{$.Prefix}}{{$.Type}}Table{{.}}
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

var tCreateTable = template.Must(template.New("CreateTable").Funcs(funcMap).Parse(sCreateTable))

//-------------------------------------------------------------------------------------------------

// function template to create an index
const sCreateIndex = `
// CreateIndexes executes queries that create the indexes needed by the {{.Type}} table.
func (tbl {{.Prefix}}{{.Type}}Table) CreateIndexes(ifNotExist bool) (err error) {
	{{if gt (len .Body1) 0 -}}
	extra := tbl.ternary(ifNotExist, "IF NOT EXISTS ", "")
    {{end -}}
	{{range .Body1}}{{$n := .}}
	_, err = tbl.Exec(tbl.create{{$.Prefix}}{{$n}}IndexSql(extra))
	if err != nil {
		return err
	}
    {{end}}
	return nil
}

{{range .Body1}}{{$n := .}}
func (tbl {{$.Prefix}}{{$.Type}}Table) create{{$.Prefix}}{{$n}}IndexSql(ifNotExist string) string {
	indexPrefix := tbl.Prefix
	if strings.HasSuffix(indexPrefix, ".") {
		indexPrefix = tbl.Prefix[0:len(indexPrefix)-1]
	}
	return fmt.Sprintf(sqlCreate{{$.Prefix}}{{$n}}Index, ifNotExist, indexPrefix, tbl.Prefix, tbl.Name)
}
{{end}}
`

var tCreateIndex = template.Must(template.New("CreateIndex").Funcs(funcMap).Parse(sCreateIndex))

//-------------------------------------------------------------------------------------------------
//TODO
//func (builder UpdateBuilder) SetFields(fields FieldList) UpdateBuilder {
//	for i := 0; i < len(fields); i++ {
//		builder.Upd = builder.Upd.Set(fields[i].Name, fields[i].Value)
//	}
//	return builder
//}
//
//func (builder UpdateBuilder) Set(what string, v interface{}) UpdateBuilder {
//	builder.Upd = builder.Upd.Set(what, v)
//	return builder
//}
//
//func (builder UpdateBuilder) Unset(what string) UpdateBuilder {
//	builder.Upd = builder.Upd.SetSQL(what, "null")
//	return builder
//}
