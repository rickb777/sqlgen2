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
	Db           *sql.DB
	Dialect      dialect.Dialect
}

// New{{.Prefix}}{{.Type}}Table returns a new table instance.
func New{{.Prefix}}{{.Type}}Table(prefix, name string, db *sql.DB, dialect dialect.Dialect) {{.Prefix}}{{.Type}}Table {
	if name == "" {
		name = {{.Prefix}}{{.Type}}TableName
	}
	return {{.Prefix}}{{.Type}}Table{prefix, name, db, dialect}
}
`

var tTable = template.Must(template.New("Table").Funcs(funcMap).Parse(sTable))

//-------------------------------------------------------------------------------------------------

// function template to scan a single row.
const sScanRow = `
// Scan{{.Prefix}}{{.Type}} reads a database record into a single value.
func Scan{{.Prefix}}{{.Type}}(row *sql.Row) (*{{.Type}}, error) {
{{range .Body1}}{{.}}{{end}}
	err := row.Scan(
{{range .Body2}}{{.}}{{end}}
	)
	if err != nil {
		return nil, err
	}

	v := &{{.Type}}{}
{{range .Body3}}{{.}}{{end}}
	return v, nil
}
`

var tScanRow = template.Must(template.New("ScanRow").Funcs(funcMap).Parse(sScanRow))

//-------------------------------------------------------------------------------------------------

// function template to scan multiple rows.
const sScanRows = `
// Scan{{.Prefix}}{{.Types}} reads database records into a slice of values.
func Scan{{.Prefix}}{{.Types}}(rows *sql.Rows) ([]*{{.Type}}, error) {
	var err error
	var vv []*{{.Type}}

{{range .Body1}}{{.}}{{end}}
	for rows.Next() {
		err = rows.Scan(
{{range .Body2}}{{.}}{{end}}
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
func Slice{{.Prefix}}{{.Type}}{{.Suffix}}(v *{{.Type}}) []interface{} {
{{range .Body1}}{{.}}{{end}}
{{range .Body2}}{{.}}{{end}}
	return []interface{}{
{{range .Body3}}{{.}}{{end}}
	}
}
`

var tSliceRow = template.Must(template.New("SliceRow").Funcs(funcMap).Parse(sSliceRow))

//-------------------------------------------------------------------------------------------------

const sSelectRow = `
// QueryOne is the low-level access function for one {{.Type}}.
func (tbl {{.Prefix}}{{.Type}}Table) QueryOne(query string, args ...interface{}) (*{{.Type}}, error) {
	row := tbl.Db.QueryRow(query, args...)
	return Scan{{.Prefix}}{{.Type}}(row)
}

// SelectOneSA allows a single {{.Type}} to be obtained from the database using supplied dialect-specific parameters.
func (tbl {{.Prefix}}{{.Type}}Table) SelectOneSA(where, limitClause string, args ...interface{}) (*{{.Type}}, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", {{.Prefix}}{{.Type}}ColumnNames, tbl.Prefix, tbl.Name, where, limitClause)
	return tbl.QueryOne(query, args...)
}

// SelectOne allows a single {{.Type}} to be obtained from the database.
func (tbl {{.Prefix}}{{.Type}}Table) SelectOne(where where.Expression, dialect dialect.Dialect) (*{{.Type}}, error) {
	wh, args := where.Build(dialect)
	return tbl.SelectOneSA(wh, "LIMIT 1", args)
}
`

var tSelectRow = template.Must(template.New("SelectRow").Funcs(funcMap).Parse(sSelectRow))

//-------------------------------------------------------------------------------------------------

// function template to select multiple rows.
const sSelectRows = `
func (tbl {{.Prefix}}{{.Type}}Table) Query(query string, args ...interface{}) ([]*{{.Type}}, error) {
	rows, err := tbl.Db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return Scan{{.Prefix}}{{.Types}}(rows)
}

// SelectSA allows {{.Types}} to be obtained from the database using supplied dialect-specific parameters.
func (tbl {{.Prefix}}{{.Type}}Table) SelectSA(where string, args ...interface{}) ([]*{{.Type}}, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s", {{.Prefix}}{{.Type}}ColumnNames, tbl.Prefix, tbl.Name, where)
	return tbl.Query(query, args...)
}

// Select allows {{.Types}} to be obtained from the database that match a 'where' clause.
func (tbl {{.Prefix}}{{.Type}}Table) Select(where where.Expression, dialect dialect.Dialect) ([]*{{.Type}}, error) {
	return tbl.SelectSA(where.Build(dialect))
}
`

var tSelectRows = template.Must(template.New("SelectRows").Funcs(funcMap).Parse(sSelectRows))

//-------------------------------------------------------------------------------------------------

const sCountRows = `
// CountSA counts {{.Types}} in the database using supplied dialect-specific parameters.
func (tbl {{.Prefix}}{{.Type}}Table) CountSA(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s%s %s", tbl.Prefix, tbl.Name, where)
	row := tbl.Db.QueryRow(query, args)
	err = row.Scan(&count)
	return count, err
}

// Count counts the {{.Types}} in the database that match a 'where' clause.
func (tbl {{.Prefix}}{{.Type}}Table) Count(where where.Expression, dialect dialect.Dialect) (count int64, err error) {
	return tbl.CountSA(where.Build(dialect))
}
`

var tCountRows = template.Must(template.New("CountRows").Funcs(funcMap).Parse(sCountRows))

//-------------------------------------------------------------------------------------------------

// function template to insert a single row, updating the primary key in the struct.
const sInsertAndGetLastId = `
// Insert adds new records for the {{.Types}}. The {{.Types}} have their primary key fields
// set to the new record identifiers.
func (tbl {{.Prefix}}{{.Type}}Table) Insert(vv ...*{{.Type}}) error {
	var stmt, params string
	switch tbl.Dialect {
	case dialect.Postgres:
		stmt = sqlInsert{{$.Prefix}}{{$.Type}}Postgres
		params = s{{$.Prefix}}{{$.Type}}DataColumnParamsPostgres
	default:
		stmt = sqlInsert{{$.Prefix}}{{$.Type}}Simple
		params = s{{$.Prefix}}{{$.Type}}DataColumnParamsSimple
	}
	st, err := tbl.Db.Prepare(fmt.Sprintf(stmt, tbl.Prefix, tbl.Name, params))
	if err != nil {
		return err
	}
	defer st.Close()

	for _, v := range vv {
		res, err := st.Exec(Slice{{.Prefix}}{{.Type}}WithoutPk(v)...)
		if err != nil {
			return err
		}

		v.{{.Table.Primary.Name}}, err = res.LastInsertId()
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
func (tbl {{.Prefix}}{{.Type}}Table) Insert(vv ...*{{.Type}}) error {
	var stmt, params string
	switch tbl.Dialect {
	case dialect.Postgres:
		stmt = sqlInsert{{$.Prefix}}{{$.Type}}Postgres
		params = ""
	default:
		stmt = sqlInsert{{$.Prefix}}{{$.Type}}Simple
		params = ""
	}
	st, err := tbl.Db.Prepare(fmt.Sprintf(stmt, tbl.Prefix, tbl.Name, params))
	if err != nil {
		return err
	}
	defer st.Close()

	for _, v := range vv {
		res, err := st.Exec(Slice{{.Prefix}}{{.Type}}Stmt(v)...)
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
// Update updates a record. It returns the number of rows affected.
// Not every database or database driver may support this.
func (tbl {{.Prefix}}{{.Type}}Table) Update(v *{{.Type}}) (int64, error) {
	var stmt string
	switch tbl.Dialect {
	case dialect.Postgres:
		stmt = sqlUpdate{{$.Prefix}}{{$.Type}}ByPkPostgres
	default:
		stmt = sqlUpdate{{$.Prefix}}{{$.Type}}ByPkSimple
	}
	query := fmt.Sprintf(stmt, tbl.Prefix, tbl.Name)
	args := Slice{{.Prefix}}{{.Type}}WithoutPk(v)
	args = append(args, v.{{.Table.Primary.Name}})
	return tbl.Exec(query, args...)
}
`

var tUpdate = template.Must(template.New("Update").Funcs(funcMap).Parse(sUpdate))

//-------------------------------------------------------------------------------------------------

// function template to call sql exec
const sExec = `
// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// It returns the number of rows affected.
// Not every database or database driver may support this.
func (tbl {{.Prefix}}{{.Type}}Table) Exec(query string, args ...interface{}) (int64, error) {
	res, err := tbl.Db.Exec(query, args...)
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
// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// It returns the number of rows affected.
// Not every database or database driver may support this.
func (tbl {{.Prefix}}{{.Type}}Table) CreateTable(ifNotExist bool) (int64, error) {
	return tbl.Exec(tbl.createTableSql(ifNotExist))
}

func (tbl {{.Prefix}}{{.Type}}Table) createTableSql(ifNotExist bool) string {
	var stmt string
	switch tbl.Dialect {
	{{range .Dialects -}}
	case dialect.{{.}}: stmt = sqlCreate{{$.Prefix}}{{$.Type}}Table{{.}}
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
	var stmt string
	switch tbl.Dialect {
	{{range $.Dialects -}}
	case dialect.{{.}}: stmt = sqlCreate{{$.Prefix}}{{$n}}Index{{.}}
    {{end -}}
	}
	return fmt.Sprintf(stmt, ifNotExist, tbl.Prefix, tbl.Name)
}
{{end}}

`

var tCreateIndex = template.Must(template.New("CreateIndex").Funcs(funcMap).Parse(sCreateIndex))
