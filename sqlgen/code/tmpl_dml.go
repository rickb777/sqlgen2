package code

import "text/template"

const sExec = `
// Exec executes a query without returning any rows.
// It returns the number of rows affected (if the database driver supports this).
//
// The args are for any placeholder parameters in the query.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Exec(req require.Requirement, query string, args ...interface{}) (int64, error) {
	return support.Exec(tbl, req, query, args...)
}
`

var tExec = template.Must(template.New("Exec").Funcs(funcMap).Parse(sExec))

//-------------------------------------------------------------------------------------------------

const sQueryRows = `
// Query is the low-level access method for {{.Types}}.
//
// It places a requirement, which may be nil, on the size of the expected results: this
// controls whether an error is generated when this expectation is not met.
//
// Note that this method applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Query(req require.Requirement, query string, args ...interface{}) ({{.List}}, error) {
	query = tbl.ReplaceTableName(query)
	vv, err := tbl.doQuery(req, false, query, args...)
	return vv, err
}

// QueryOne is the low-level access method for one {{.Type}}.
// If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, *{{.Type}} will be nil.
//
// Note that this method applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) QueryOne(query string, args ...interface{}) (*{{.Type}}, error) {
	query = tbl.ReplaceTableName(query)
	return tbl.doQueryOne(nil, query, args...)
}

// MustQueryOne is the low-level access method for one {{.Type}}.
//
// It places a requirement that exactly one result must be found; an error is generated when this expectation is not met.
//
// Note that this method applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) MustQueryOne(query string, args ...interface{}) (*{{.Type}}, error) {
	query = tbl.ReplaceTableName(query)
	return tbl.doQueryOne(require.One, query, args...)
}

func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) doQueryOne(req require.Requirement, query string, args ...interface{}) (*{{.Type}}, error) {
	list, err := tbl.doQuery(req, true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) doQuery(req require.Requirement, firstOnly bool, query string, args ...interface{}) ({{.List}}, error) {
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, tbl.logError(err)
	}
	defer rows.Close()

	vv, n, err := scan{{.Prefix}}{{.Types}}(rows, firstOnly)
	return vv, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(err, req, n))
}
`

var tQueryRows = template.Must(template.New("QueryRows").Funcs(funcMap).Parse(sQueryRows))

const sQueryThings = `
// QueryOneNullString is a low-level access method for one string. This can be used for function queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) QueryOneNullString(query string, args ...interface{}) (result sql.NullString, err error) {
	err = support.QueryOneNullThing(tbl, nil, &result, query, args...)
	return result, err
}

// MustQueryOneNullString is a low-level access method for one string. This can be used for function queries and
// such like.
//
// It places a requirement that exactly one result must be found; an error is generated when this expectation is not met.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) MustQueryOneNullString(query string, args ...interface{}) (result sql.NullString, err error) {
	err = support.QueryOneNullThing(tbl, require.One, &result, query, args...)
	return result, err
}

// QueryOneNullInt64 is a low-level access method for one int64. This can be used for 'COUNT(1)' queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) QueryOneNullInt64(query string, args ...interface{}) (result sql.NullInt64, err error) {
	err = support.QueryOneNullThing(tbl, nil, &result, query, args...)
	return result, err
}

// MustQueryOneNullInt64 is a low-level access method for one int64. This can be used for 'COUNT(1)' queries and
// such like.
//
// It places a requirement that exactly one result must be found; an error is generated when this expectation is not met.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) MustQueryOneNullInt64(query string, args ...interface{}) (result sql.NullInt64, err error) {
	err = support.QueryOneNullThing(tbl, require.One, &result, query, args...)
	return result, err
}

// QueryOneNullFloat64 is a low-level access method for one float64. This can be used for 'AVG(...)' queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) QueryOneNullFloat64(query string, args ...interface{}) (result sql.NullFloat64, err error) {
	err = support.QueryOneNullThing(tbl, nil, &result, query, args...)
	return result, err
}

// MustQueryOneNullFloat64 is a low-level access method for one float64. This can be used for 'AVG(...)' queries and
// such like.
//
// It places a requirement that exactly one result must be found; an error is generated when this expectation is not met.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) MustQueryOneNullFloat64(query string, args ...interface{}) (result sql.NullFloat64, err error) {
	err = support.QueryOneNullThing(tbl, require.One, &result, query, args...)
	return result, err
}

// ReplaceTableName replaces all occurrences of "{TABLE}" with the table's name.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) ReplaceTableName(query string) string {
	return strings.Replace(query, "{TABLE}", tbl.name.String(), -1)
}
`

var tQueryThings = template.Must(template.New("QueryThings").Funcs(funcMap).Parse(sQueryThings))

//-------------------------------------------------------------------------------------------------

const sGetRow = `{{if .Table.Primary}}
//--------------------------------------------------------------------------------

// Get{{.Type}} gets the record with a given primary key value.
// If not found, *{{.Type}} will be nil.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Get{{.Type}}(id {{.Table.Primary.Type.Name}}) (*{{.Type}}, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE {{.Table.Primary.SqlName}}=?", {{.CamelName}}ColumnNames, tbl.name)
	v, err := tbl.doQueryOne(nil, query, id)
	return v, err
}

// MustGet{{.Type}} gets the record with a given primary key value.
//
// It places a requirement that exactly one result must be found; an error is generated when this expectation is not met.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) MustGet{{.Type}}(id {{.Table.Primary.Type.Name}}) (*{{.Type}}, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE {{.Table.Primary.SqlName}}=?", {{.CamelName}}ColumnNames, tbl.name)
	v, err := tbl.doQueryOne(require.One, query, id)
	return v, err
}

// Get{{.Types}} gets records from the table according to a list of primary keys.
// Although the list of ids can be arbitrarily long, there are practical limits;
// note that Oracle DB has a limit of 1000.
//
// It places a requirement, which may be nil, on the size of the expected results: in particular, require.All
// controls whether an error is generated not all the ids produce a result.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Get{{.Types}}(req require.Requirement, id ...{{.Table.Primary.Type.Name}}) (list {{.List}}, err error) {
	if len(id) > 0 {
		if req == require.All {
			req = require.Exactly(len(id))
		}
		pl := tbl.dialect.Placeholders(len(id))
		query := fmt.Sprintf("SELECT %s FROM %s WHERE {{.Table.Primary.SqlName}} IN (%s)", {{.CamelName}}ColumnNames, tbl.name, pl)
		args := make([]interface{}, len(id))

		for i, v := range id {
			args[i] = v
		}

		list, err = tbl.doQuery(req, false, query, args...)
	}

	return list, err
}
{{end -}}
`

var tGetRow = template.Must(template.New("GetRow").Funcs(funcMap).Parse(sGetRow))

//-------------------------------------------------------------------------------------------------

const sSelectRows = `
// SelectOneWhere allows a single Example to be obtained from the table that match a 'where' clause
// and some limit. Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
// If not found, *Example will be nil.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.One
// controls whether an error is generated when no result is found.
//
// The args are for any placeholder parameters in the query.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) SelectOneWhere(req require.Requirement, where, orderBy string, args ...interface{}) (*{{.Type}}, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT 1", {{.CamelName}}ColumnNames, tbl.name, where, orderBy)
	v, err := tbl.doQueryOne(req, query, args...)
	return v, err
}

// SelectOne allows a single {{.Type}} to be obtained from the sqlgen2.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
// If not found, *Example will be nil.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.One
// controls whether an error is generated when no result is found.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) SelectOne(req require.Requirement, wh where.Expression, qc where.QueryConstraint) (*{{.Type}}, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	return tbl.SelectOneWhere(req, whs, orderBy, args...)
}

// SelectWhere allows {{.Types}} to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
//
// The args are for any placeholder parameters in the query.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) SelectWhere(req require.Requirement, where, orderBy string, args ...interface{}) ({{.List}}, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", {{.CamelName}}ColumnNames, tbl.name, where, orderBy)
	vv, err := tbl.doQuery(req, false, query, args...)
	return vv, err
}

// Select allows {{.Types}} to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Select(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ({{.List}}, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	return tbl.SelectWhere(req, whs, orderBy, args...)
}
`

var tSelectRows = template.Must(template.New("SelectRows").Funcs(funcMap).Parse(sSelectRows))

//-------------------------------------------------------------------------------------------------

const sCountRows = `
// CountWhere counts {{.Types}} in the table that match a 'where' clause.
// Use a blank string for the 'where' argument if it is not needed.
//
// The args are for any placeholder parameters in the query.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) CountWhere(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s %s", tbl.name, where)
	tbl.logQuery(query, args...)
	row := tbl.db.QueryRowContext(tbl.ctx, query, args...)
	err = row.Scan(&count)
	return count, tbl.logIfError(err)
}

// Count counts the {{.Types}} in the table that match a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Count(wh where.Expression) (count int64, err error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	return tbl.CountWhere(whs, args...)
}
`

var tCountRows = template.Must(template.New("CountRows").Funcs(funcMap).Parse(sCountRows))

//-------------------------------------------------------------------------------------------------

const sSliceItem = `
//--------------------------------------------------------------------------------
{{range .Table.SimpleFields}}
// Slice{{.Name}} gets the {{.Name}} column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl {{$.Prefix}}{{$.Type}}{{$.Thing}}) Slice{{camel .SqlName}}(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]{{.Type.Type}}, error) {
	return tbl.get{{.Type.Tag}}list(req, "{{.SqlName}}", wh, qc)
}
{{end}}
{{range .Table.SimpleFields.DistinctTypes}}
func (tbl {{$.Prefix}}{{$.Type}}{{$.Thing}}) get{{.Tag}}list(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]{{.Type}}, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", sqlname, tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, tbl.logError(err)
	}
	defer rows.Close()

	var v {{.Type}}
	list := make([]{{.Type}}, 0, 10)

	for rows.Next() {
		err = rows.Scan(&v)
		if err == sql.ErrNoRows {
			return list, tbl.logIfError(require.ErrorIfQueryNotSatisfiedBy(req, int64(len(list))))
		} else {
			list = append(list, v)
		}
	}
	return list, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, int64(len(list))))
}
{{end}}
`

var tSliceItem = template.Must(template.New("SliceItem").Funcs(funcMap).Parse(sSliceItem))

//-------------------------------------------------------------------------------------------------

// function template to insert a single row, updating the primary key in the struct.
const sInsert = `
// Insert adds new records for the {{.Types}}.
{{if .Table.HasLastInsertId}}// The {{.Types}} have their primary key fields set to the new record identifiers.{{end}}
// The {{.Type}}.PreInsert() method will be called, if it exists.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Insert(req require.Requirement, vv ...*{{.Type}}) error {
	var params string
	switch tbl.dialect {
	case schema.Postgres:
		params = s{{$.Prefix}}{{$.Type}}DataColumnParamsPostgres
	default:
		params = s{{$.Prefix}}{{$.Type}}DataColumnParamsSimple
	}

	if req == require.All {
		req = require.Exactly(len(vv))
	}

	var count int64
	query := fmt.Sprintf(sqlInsert{{$.Prefix}}{{$.Type}}, tbl.name, params)
	st, err := tbl.db.PrepareContext(tbl.ctx, query)
	if err != nil {
		return err
	}
	defer st.Close()

	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPreInsert); ok {
			err := hook.PreInsert()
			if err != nil {
				return tbl.logError(err)
			}
		}

		fields, err := slice{{.Prefix}}{{.Type}}{{if .Table.HasLastInsertId}}WithoutPk{{end}}(v)
		if err != nil {
			return tbl.logError(err)
		}

		tbl.logQuery(query, fields...)
		res, err := st.ExecContext(tbl.ctx, fields...)
		if err != nil {
			return tbl.logError(err)
		}

		{{if .Table.HasLastInsertId -}}
		{{if eq .Table.Primary.Type.Name "int64" -}}
		v.{{.Table.Primary.Name}}, err = res.LastInsertId()
		{{- else -}}
		_i64, err := res.LastInsertId()
		v.{{.Table.Primary.Name}} = {{.Table.Primary.Type.Name}}(_i64)
		{{end}}
		{{- end}}
		if err != nil {
			return tbl.logError(err)
		}

		n, err := res.RowsAffected()
		if err != nil {
			return tbl.logError(err)
		}
		count += n
	}

	return tbl.logIfError(require.ErrorIfExecNotSatisfiedBy(req, count))
}
`

var tInsert = template.Must(template.New("Insert").Funcs(funcMap).Parse(sInsert))

//-------------------------------------------------------------------------------------------------

// function template to update a single row.
const sUpdateFields = `
// UpdateFields updates one or more columns, given a 'where' clause.
//
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) UpdateFields(req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error) {
	return support.UpdateFields(tbl, req, wh, fields...)
}
`

var tUpdateFields = template.Must(template.New("UpdateFields").Funcs(funcMap).Parse(sUpdateFields))

//-------------------------------------------------------------------------------------------------

// function template to update rows.
const sUpdate = `{{if .Table.Primary}}
//--------------------------------------------------------------------------------

// Update updates records, matching them by primary key. It returns the number of rows affected.
// The {{.Type}}.PreUpdate(Execer) method will be called, if it exists.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Update(req require.Requirement, vv ...*{{.Type}}) (int64, error) {
	var stmt string
	switch tbl.dialect {
	case schema.Postgres:
		stmt = sqlUpdate{{$.Prefix}}{{$.Type}}ByPkPostgres
	default:
		stmt = sqlUpdate{{$.Prefix}}{{$.Type}}ByPkSimple
	}

	if req == require.All {
		req = require.Exactly(len(vv))
	}

	var count int64
	query := fmt.Sprintf(stmt, tbl.name)

	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPreUpdate); ok {
			err := hook.PreUpdate()
			if err != nil {
				return count, tbl.logError(err)
			}
		}

		args, err := slice{{.Prefix}}{{.Type}}WithoutPk(v)
		args = append(args, v.{{.Table.Primary.Name}})
		if err != nil {
			return count, tbl.logError(err)
		}

		n, err := tbl.Exec(nil, query, args...)
		if err != nil {
			return count, err
		}
		count += n
	}

	return count, tbl.logIfError(require.ErrorIfExecNotSatisfiedBy(req, count))
}
{{end -}}
`

var tUpdate = template.Must(template.New("Update").Funcs(funcMap).Parse(sUpdate))

//-------------------------------------------------------------------------------------------------

const sDelete = `
{{if .Table.Primary -}}
// Delete{{.Types}} deletes rows from the table, given some primary keys.
// The list of ids can be arbitrarily long.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Delete{{.Types}}(req require.Requirement, id ...{{.Table.Primary.Type.Name}}) (int64, error) {
	const batch = 1000 // limited by Oracle DB
	const qt = "DELETE FROM %s WHERE {{.Table.Primary.SqlName}} IN (%s)"

	if req == require.All {
		req = require.Exactly(len(id))
	}

	var count, n int64
	var err error
	var max = batch
	if len(id) < batch {
		max = len(id)
	}
	args := make([]interface{}, max)

	if len(id) > batch {
		pl := tbl.dialect.Placeholders(batch)
		query := fmt.Sprintf(qt, tbl.name, pl)

		for len(id) > batch {
			for i := 0; i < batch; i++ {
				args[i] = id[i]
			}

			n, err = tbl.Exec(nil, query, args...)
			count += n
			if err != nil {
				return count, err
			}

			id = id[batch:]
		}
	}

	if len(id) > 0 {
		pl := tbl.dialect.Placeholders(len(id))
		query := fmt.Sprintf(qt, tbl.name, pl)

		for i := 0; i < len(id); i++ {
			args[i] = id[i]
		}

		n, err = tbl.Exec(nil, query, args...)
		count += n
	}

	return count, tbl.logIfError(require.ChainErrorIfExecNotSatisfiedBy(err, req, n))
}

{{end -}}
// Delete deletes one or more rows from the table, given a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Delete(req require.Requirement, wh where.Expression) (int64, error) {
	query, args := tbl.deleteRows(wh)
	return tbl.Exec(req, query, args...)
}

func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) deleteRows(wh where.Expression) (string, []interface{}) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	query := fmt.Sprintf("DELETE FROM %s %s", tbl.name, whs)
	return query, args
}
`

var tDelete = template.Must(template.New("Delete").Funcs(funcMap).Parse(sDelete))
