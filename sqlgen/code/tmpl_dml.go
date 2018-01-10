package code

import "text/template"

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

const sSliceItem = `
//--------------------------------------------------------------------------------
{{range .Table.SimpleFields}}
// Slice{{.Name}} gets the {{.Name}} column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl {{$.Prefix}}{{$.Type}}{{$.Thing}}) Slice{{camel .SqlName}}(wh where.Expression, qc where.QueryConstraint) ([]{{.Type.Type}}, error) {
	return tbl.get{{.Type.Tag}}list("{{.SqlName}}", wh, qc)
}
{{end}}
{{range .Table.SimpleFields.DistinctTypes}}
func (tbl {{$.Prefix}}{{$.Type}}{{$.Thing}}) get{{.Tag}}list(sqlname string, wh where.Expression, qc where.QueryConstraint) ([]{{.Type}}, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", sqlname, tbl.prefix, tbl.name, whs, orderBy)
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

const sGetRow = `{{if .Table.Primary}}
//--------------------------------------------------------------------------------

// Get{{.Type}} gets the record with a given primary key value.
// If not found, *{{.Type}} will be nil.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Get{{.Type}}(id {{.Table.Primary.Type.Name}}) (*{{.Type}}, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s WHERE {{.Table.Primary.SqlName}}=?", {{.Prefix}}{{.Type}}ColumnNames, tbl.prefix, tbl.name)
	return tbl.QueryOne(query, id)
}

// Get{{.Types}} gets records from the table according to a list of primary keys.
// Although the list of ids can be arbitrarily long, there are practical limits;
// note that Oracle DB has a limit of 1000.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Get{{.Types}}(id ...{{.Table.Primary.Type.Name}}) (list {{.List}}, err error) {
	if len(id) > 0 {
		pl := tbl.dialect.Placeholders(len(id))
		query := fmt.Sprintf("SELECT %s FROM %s%s WHERE {{.Table.Primary.SqlName}} IN (%s)", {{.Prefix}}{{.Type}}ColumnNames, tbl.prefix, tbl.name, pl)
		args := make([]interface{}, len(id))

		for i, v := range id {
			args[i] = v
		}

		list, err = tbl.Query(query, args...)
	}

	return list, err
}
{{end -}}
`

var tGetRow = template.Must(template.New("GetRow").Funcs(funcMap).Parse(sGetRow))

//-------------------------------------------------------------------------------------------------

const sSelectRows = `
// SelectOneSA allows a single Example to be obtained from the table that match a 'where' clause
// and some limit. Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
// If not found, *Example will be nil.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) SelectOneSA(where, orderBy string, args ...interface{}) (*{{.Type}}, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s LIMIT 1", {{.Prefix}}{{.Type}}ColumnNames, tbl.prefix, tbl.name, where, orderBy)
	return tbl.QueryOne(query, args...)
}

// SelectOne allows a single {{.Type}} to be obtained from the sqlgen2.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
// If not found, *Example will be nil.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) SelectOne(wh where.Expression, qc where.QueryConstraint) (*{{.Type}}, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	return tbl.SelectOneSA(whs, orderBy, args...)
}

// SelectSA allows {{.Types}} to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) SelectSA(where, orderBy string, args ...interface{}) ({{.List}}, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", {{.Prefix}}{{.Type}}ColumnNames, tbl.prefix, tbl.name, where, orderBy)
	return tbl.Query(query, args...)
}

// Select allows {{.Types}} to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Select(wh where.Expression, qc where.QueryConstraint) ({{.List}}, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	return tbl.SelectSA(whs, orderBy, args...)
}
`

var tSelectRows = template.Must(template.New("SelectRows").Funcs(funcMap).Parse(sSelectRows))

//-------------------------------------------------------------------------------------------------

const sCountRows = `
// CountSA counts {{.Types}} in the table that match a 'where' clause.
// Use a blank string for the 'where' argument if it is not needed.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) CountSA(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s%s %s", tbl.prefix, tbl.name, where)
	tbl.logQuery(query, args...)
	row := tbl.db.QueryRowContext(tbl.ctx, query, args...)
	err = row.Scan(&count)
	return count, err
}

// Count counts the {{.Types}} in the table that match a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Count(wh where.Expression) (count int64, err error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	return tbl.CountSA(whs, args...)
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
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) UpdateFields(wh where.Expression, fields ...sql.NamedArg) (int64, error) {
	query, args := tbl.updateFields(wh, fields...)
	return tbl.Exec(query, args...)
}

func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) updateFields(wh where.Expression, fields ...sql.NamedArg) (string, []interface{}) {
	list := sqlgen2.NamedArgList(fields)
	assignments := strings.Join(list.Assignments(tbl.dialect, 1), ", ")
	whs, wargs := where.BuildExpression(wh, tbl.dialect)
	query := fmt.Sprintf("UPDATE %s%s SET %s %s", tbl.prefix, tbl.name, assignments, whs)
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
{{if .Table.Primary -}}
// Delete{{.Types}} deletes rows from the table, given some primary keys.
// The list of ids can be arbitrarily long.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Delete{{.Types}}(id ...{{.Table.Primary.Type.Name}}) (int64, error) {
	const batch = 1000 // limited by Oracle DB
	const qt = "DELETE FROM %s%s WHERE {{.Table.Primary.SqlName}} IN (%s)"

	var count, n int64
	var err error
	var max = batch
	if len(id) < batch {
		max = len(id)
	}
	args := make([]interface{}, max)

	if len(id) > batch {
		pl := tbl.dialect.Placeholders(batch)
		query := fmt.Sprintf(qt, tbl.prefix, tbl.name, pl)

		for len(id) > batch {
			for i := 0; i < batch; i++ {
				args[i] = id[i]
			}

			n, err = tbl.Exec(query, args...)
			count += n
			if err != nil {
				return count, err
			}

			id = id[batch:]
		}
	}

	if len(id) > 0 {
		pl := tbl.dialect.Placeholders(len(id))
		query := fmt.Sprintf(qt, tbl.prefix, tbl.name, pl)

		for i := 0; i < batch; i++ {
			args[i] = id[i]
		}

		n, err = tbl.Exec(query, args...)
		count += n
	}

	return count, err
}

{{end -}}
// Delete deletes one or more rows from the table, given a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Delete(wh where.Expression) (int64, error) {
	query, args := tbl.deleteRows(wh)
	return tbl.Exec(query, args...)
}

func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) deleteRows(wh where.Expression) (string, []interface{}) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	query := fmt.Sprintf("DELETE FROM %s%s %s", tbl.prefix, tbl.name, whs)
	return query, args
}
`

var tDelete = template.Must(template.New("Delete").Funcs(funcMap).Parse(sDelete))

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
