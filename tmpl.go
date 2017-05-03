package main

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
type {{.Prefix}}{{.Type}}Table struct {
	Name string
	Db   *sql.DB
}
`

var tTable = template.Must(template.New("Table").Funcs(funcMap).Parse(sTable))

//-------------------------------------------------------------------------------------------------

const sConst = `
const s{{.Name}} = {{ticked .Body}}
`

const sTableName = `
func {{.Name}}(tableName string) string {
	return fmt.Sprintf(s{{.Name}}, tableName)
}
`

var tConst = template.Must(template.New("Const").Funcs(funcMap).Parse(sConst))
var tConstWithTableName = template.Must(template.New("Const").Funcs(funcMap).Parse(sConst + sTableName))

//-------------------------------------------------------------------------------------------------

// function template to scan a single row.
const sScanRow = `
// Scan{{.Type}} reads a database record into a single value.
func Scan{{.Type}}(row *sql.Row) (*{{.Type}}, error) {
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
// Scan{{.Types}} reads database records into a slice of values.
func Scan{{.Types}}(rows *sql.Rows) ([]*{{.Type}}, error) {
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
func Slice{{.Type}}{{.Suffix}}(v *{{.Type}}) []interface{} {
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
func (tbl {{.Prefix}}{{.Type}}Table) QueryOne(query string, args ...interface{}) (*{{.Type}}, error) {
	row := tbl.Db.QueryRow(query, args...)
	return Scan{{.Type}}(row)
}

func (tbl {{.Prefix}}{{.Type}}Table) SelectOne(where string, args ...interface{}) (*{{.Type}}, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s LIMIT 1", s{{.Type}}ColumnNames, tbl.Name, where)
	return tbl.QueryOne(query, args...)
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
	return Scan{{.Types}}(rows)
}

func (tbl {{.Prefix}}{{.Type}}Table) Select(where string, args ...interface{}) ([]*{{.Type}}, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s", s{{.Type}}ColumnNames, tbl.Name, where)
	return tbl.Query(query, args...)
}
`

var tSelectRows = template.Must(template.New("SelectRows").Funcs(funcMap).Parse(sSelectRows))

//-------------------------------------------------------------------------------------------------

const sCountRows = `
func (tbl {{.Prefix}}{{.Type}}Table) Count(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s %s", tbl.Name, where)
	row := tbl.Db.QueryRow(query, args)
	err = row.Scan(&count)
	return count, err
}
`

var tCountRows = template.Must(template.New("CountRows").Funcs(funcMap).Parse(sCountRows))

//-------------------------------------------------------------------------------------------------

// function template to insert a single row, updating the primary key in the struct.
const sInsertAndGetLastId = `
func (tbl {{.Prefix}}{{.Type}}Table) Insert(v *{{.Type}}) error {
	query := fmt.Sprintf(sInsert{{.Type}}Stmt, tbl.Name)
	res, err := tbl.Db.Exec(query, Slice{{.Type}}WithoutPk(v)...)
	if err != nil {
		return err
	}

	v.{{.Table.Primary.Name}}, err = res.LastInsertId()
	return err
}
`

var tInsertAndGetLastId = template.Must(template.New("InsertAndGetLastId").Funcs(funcMap).Parse(sInsertAndGetLastId))

//-------------------------------------------------------------------------------------------------

// function template to insert a single row.
const sInsert = `
func (tbl {{.Prefix}}{{.Type}}Table) Insert(v *{{.Type}}) error {
	query := fmt.Sprintf(sInsert{{.Type}}Stmt, tbl.Name)
	_, err := tbl.Db.Exec(query, Slice{{.Type}}(v)...)
	return err
}
`

var tInsert = template.Must(template.New("Insert").Funcs(funcMap).Parse(sInsert))

//-------------------------------------------------------------------------------------------------

// function template to update a single row.
const sUpdate = `
// Update updates a record. It returns the number of rows affected.
// Not every database or database driver may support this.
func (tbl {{.Prefix}}{{.Type}}Table) Update(v *{{.Type}}) (int64, error) {
	query := fmt.Sprintf(sUpdate{{.Type}}ByPkStmt, tbl.Name)
	args := Slice{{.Type}}WithoutPk(v)
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
