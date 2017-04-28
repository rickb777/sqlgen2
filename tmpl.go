package main

// template to declare the package name.
var sPackage = `// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

package %s
`

const sTable = `
// %sTableName is the default name for this table.
const %sTableName = %q

// %sTable holds a given table name with the database reference, providing access methods below.
type %sTable struct {
	Name string
	Db   *sql.DB
}
`

const sConst = `
const s%s = %s

func %s(tableName string) string {
	return fmt.Sprintf(s%s, tableName)
}
`

// function template to scan a single row.
const sScanRow = `
// Scan%s reads a database record into a single value.
func Scan%s(row *sql.Row) (*%s, error) {
%s
	err := row.Scan(
%s
	)
	if err != nil {
		return nil, err
	}

	v := &%s{}
%s

	return v, nil
}
`

// function template to scan multiple rows.
const sScanRows = `
// Scan%s reads database records into a slice of values.
func Scan%s(rows *sql.Rows) ([]*%s, error) {
	var err error
	var vv []*%s

%s
	for rows.Next() {
		err = rows.Scan(
%s
		)
		if err != nil {
			return vv, err
		}

		v := &%s{}
%s
		vv = append(vv, v)
	}
	return vv, rows.Err()
}
`

const sSliceRow = `
func Slice%s%s(v *%s) []interface{} {
%s
%s
	return []interface{}{
%s
	}
}
`

const sSelectRow = `
func (tbl %sTable) SelectOne(query string, args ...interface{}) (*%s, error) {
	row := tbl.Db.QueryRow(query, args...)
	return Scan%s(row)
}
`

// function template to select multiple rows.
const sSelectRows = `
func (tbl %sTable) Select(query string, args ...interface{}) ([]*%s, error) {
	rows, err := tbl.Db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return Scan%s(rows)
}
`

// function template to insert a single row, updating the primary key in the struct.
const sInsertAndGetLastId = `
func (tbl %sTable) Insert(v *%s) error {
	query := fmt.Sprintf(sInsert%sStmt, tbl.Name)
	res, err := tbl.Db.Exec(query, Slice%sWithoutPk(v)...)
	if err != nil {
		return err
	}

	v.%s, err = res.LastInsertId()
	return err
}
`

// function template to insert a single row.
const sInsert = `
func (tbl %sTable) Insert(v *%s) error {
	query := fmt.Sprintf(sInsert%sStmt, tbl.Name)
	_, err := tbl.Db.Exec(query, Slice%s(v)...)
	return err
}
`

// function template to update a single row.
const sUpdate = `
// Update updates a record. It returns the number of rows affected.
// Not every database or database driver may support this.
func (tbl %sTable) Update(v *%s) (int64, error) {
	query := fmt.Sprintf(sUpdate%sByPkStmt, tbl.Name)
	args := Slice%sWithoutPk(v)
	args = append(args, v.%s)
	return tbl.Exec(query, args...)
}
`

// function template to call sql exec
const sExec = `
// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// It returns the number of rows affected.
// Not every database or database driver may support this.
func (tbl %sTable) Exec(query string, args ...interface{}) (int64, error) {
	res, err := tbl.Db.Exec(query, args...)
	if err != nil {
		return 0, nil
	}
	return res.RowsAffected()
}
`
