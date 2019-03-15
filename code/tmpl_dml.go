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
// Query is the low-level request method for this table. The query is logged using whatever logger is
// configured. If an error arises, this too is logged.
//
// If you need a context other than the background, use WithContext before calling Query.
//
// The args are for any placeholder parameters in the query.
//
// The caller must call rows.Close() on the result.
//
// Wrap the result in *sqlapi.Rows if you need to access its data as a map.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Query(query string, args ...interface{}) (sqlapi.SqlRows, error) {
	return support.Query(tbl, query, args...)
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
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) QueryOneNullString(req require.Requirement, query string, args ...interface{}) (result sql.NullString, err error) {
	err = support.QueryOneNullThing(tbl, req, &result, query, args...)
	return result, err
}

// QueryOneNullInt64 is a low-level access method for one int64. This can be used for 'COUNT(1)' queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) QueryOneNullInt64(req require.Requirement, query string, args ...interface{}) (result sql.NullInt64, err error) {
	err = support.QueryOneNullThing(tbl, req, &result, query, args...)
	return result, err
}

// QueryOneNullFloat64 is a low-level access method for one float64. This can be used for 'AVG(...)' queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) QueryOneNullFloat64(req require.Requirement, query string, args ...interface{}) (result sql.NullFloat64, err error) {
	err = support.QueryOneNullThing(tbl, req, &result, query, args...)
	return result, err
}
`

var tQueryThings = template.Must(template.New("QueryThings").Funcs(funcMap).Parse(sQueryThings))

//-------------------------------------------------------------------------------------------------

const sGetRow = `
//--------------------------------------------------------------------------------

func all{{.CamelName}}ColumnNamesQuoted(q dialect.Quoter) string {
	return strings.Join(q.QuoteN(listOf{{$.CamelName}}{{.Thing}}ColumnNames), ",")
}
{{- if .Table.Primary}}

//--------------------------------------------------------------------------------

// Get{{.Types}}By{{.Table.Primary.Name}} gets records from the table according to a list of primary keys.
// Although the list of ids can be arbitrarily long, there are practical limits;
// note that Oracle DB has a limit of 1000.
//
// It places a requirement, which may be nil, on the size of the expected results: in particular, require.All
// controls whether an error is generated not all the ids produce a result.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Get{{.Types}}By{{.Table.Primary.Name}}(req require.Requirement, id ...{{.Table.Primary.Type.Type}}) (list {{.List}}, err error) {
	if len(id) > 0 {
		if req == require.All {
			req = require.Exactly(len(id))
		}
		args := make([]interface{}, len(id))

		for i, v := range id {
			args[i] = v
		}

		list, err = tbl.get{{.Types}}(req, tbl.pk, args...)
	}

	return list, err
}

// Get{{.Type}}By{{.Table.Primary.Name}} gets the record with a given primary key value.
// If not found, *{{.Type}} will be nil.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Get{{.Type}}By{{.Table.Primary.Name}}(req require.Requirement, id {{.Table.Primary.Type.Type}}) (*{{.Type}}, error) {
	return tbl.get{{.Type}}(req, tbl.pk, id)
}
{{- end}}
{{- range .Table.Index}}
{{- if .Unique}}

// Get{{$.Type}}By{{.JoinedNames "And"}} gets the record with{{if .Single}} a{{end}} given {{.Fields.SqlNames.MkString "+"}} value{{if not .Single}}s{{end}}.
// If not found, *{{$.Type}} will be nil.
func (tbl {{$.Prefix}}{{$.Type}}{{$.Thing}}) Get{{$.Type}}By{{.JoinedNames "And"}}(req require.Requirement, {{.Fields.FormalParams.MkString ", "}}) (*{{$.Type}}, error) {
	return tbl.SelectOne(req, where.And({{.Fields.WhereClauses.MkString ", "}}), nil)
}
{{- else }}

// Get{{$.Types}}By{{.JoinedNames "And"}} gets the records with{{if .Single}} a{{end}} given {{.Fields.SqlNames.MkString "+"}} value{{if not .Single}}s{{end}}.
// If not found, the resulting slice will be empty (nil).
func (tbl {{$.Prefix}}{{$.Type}}{{$.Thing}}) Get{{$.Types}}By{{.JoinedNames "And"}}(req require.Requirement, {{.Fields.FormalParams.MkString ", "}}) ({{$.List}}, error) {
	return tbl.Select(req, where.And({{.Fields.WhereClauses.MkString ", "}}), nil)
}
{{- end}}
{{- end}}

func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) get{{.Type}}(req require.Requirement, column string, arg interface{}) (*{{.Type}}, error) {
	d := tbl.Dialect()
	q := d.Quoter()
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s=?",
		all{{.CamelName}}ColumnNamesQuoted(q), tbl.quotedName(), q.Quote(column))
	v, err := tbl.doQueryAndScanOne(req, query, arg)
	return v, err
}

func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) get{{.Types}}(req require.Requirement, column string, args ...interface{}) (list {{.List}}, err error) {
	if len(args) > 0 {
		if req == require.All {
			req = require.Exactly(len(args))
		}
		d := tbl.Dialect()
		q := d.Quoter()
		pl := d.Placeholders(len(args))
		query := fmt.Sprintf("SELECT %s FROM %s WHERE %s IN (%s)",
			all{{.CamelName}}ColumnNamesQuoted(q), tbl.quotedName(), q.Quote(column), pl)
		list, err = tbl.doQueryAndScan(req, false, query, args...)
	}

	return list, err
}

func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) doQueryAndScanOne(req require.Requirement, query string, args ...interface{}) (*{{.Type}}, error) {
	list, err := tbl.doQueryAndScan(req, true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) doQueryAndScan(req require.Requirement, firstOnly bool, query string, args ...interface{}) ({{.List}}, error) {
	rows, err := support.Query(tbl, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vv, n, err := scan{{.Prefix}}{{.Types}}(query, rows, firstOnly)
	return vv, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(err, req, n))
}

// Fetch fetches a list of {{.Type}} based on a supplied query. This is mostly used for join queries that map its
// result columns to the fields of {{.Type}}. Other queries might be better handled by GetXxx or Select methods.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Fetch(req require.Requirement, query string, args ...interface{}) ({{.List}}, error) {
	return tbl.doQueryAndScan(req, false, query, args...)
}
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
	query := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT 1",
		all{{.CamelName}}ColumnNamesQuoted(tbl.Dialect().Quoter()), tbl.quotedName(), where, orderBy)
	v, err := tbl.doQueryAndScanOne(req, query, args...)
	return v, err
}

// SelectOne allows a single {{.Type}} to be obtained from the database.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
// If not found, *Example will be nil.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.One
// controls whether an error is generated when no result is found.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) SelectOne(req require.Requirement, wh where.Expression, qc where.QueryConstraint) (*{{.Type}}, error) {
	q := tbl.Dialect().Quoter()
	whs, args := where.Where(wh, q)
	orderBy := where.BuildQueryConstraint(qc, q)
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
	query := fmt.Sprintf("SELECT %s FROM %s %s %s",
		all{{.CamelName}}ColumnNamesQuoted(tbl.Dialect().Quoter()), tbl.quotedName(), where, orderBy)
	vv, err := tbl.doQueryAndScan(req, false, query, args...)
	return vv, err
}

// Select allows {{.Types}} to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Select(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ({{.List}}, error) {
	q := tbl.Dialect().Quoter()
	whs, args := where.Where(wh, q)
	orderBy := where.BuildQueryConstraint(qc, q)
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
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s %s", tbl.quotedName(), where)
	rows, err := support.Query(tbl, query, args...)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&count)
	}
	return count, tbl.logIfError(err)
}

// Count counts the {{.Types}} in the table that match a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Count(wh where.Expression) (count int64, err error) {
	whs, args := where.Where(wh, tbl.Dialect().Quoter())
	return tbl.CountWhere(whs, args...)
}
`

var tCountRows = template.Must(template.New("CountRows").Funcs(funcMap).Parse(sCountRows))

//-------------------------------------------------------------------------------------------------

const sSliceItem = `
//--------------------------------------------------------------------------------
{{- if .Table.HasPrimaryKey}}

// Slice{{camel .Table.Primary.SqlName}} gets the {{.Table.Primary.SqlName}} column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Slice{{camel .Table.Primary.SqlName}}(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]{{.Table.Primary.Type.Type}}, error) {
	return support.Slice{{camel .Table.Primary.Type.Tag}}List(tbl, req, tbl.pk, wh, qc)
}
{{- end}}
{{- range .Table.SimpleFields.NoSkips.NoPrimary.BasicType}}

// Slice{{camel .SqlName}} gets the {{.SqlName}} column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl {{$.Prefix}}{{$.Type}}{{$.Thing}}) Slice{{camel .SqlName}}(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]{{.Type.Type}}, error) {
	return support.Slice{{camel .Type.Tag}}List(tbl, req, "{{.SqlName}}", wh, qc)
}
{{- end}}
{{- range .Table.SimpleFields.NoSkips.NoPrimary.DerivedType}}

// Slice{{camel .SqlName}} gets the {{.SqlName}} column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl {{$.Prefix}}{{$.Type}}{{$.Thing}}) Slice{{camel .SqlName}}(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]{{.Type.Type}}, error) {
	return tbl.slice{{camel .Type.Tag}}List(req, "{{.SqlName}}", wh, qc)
}
{{- end}}
{{- range .Table.SimpleFields.NoSkips.NoPrimary.DerivedType.DistinctTypes}}

func (tbl {{$.Prefix}}{{$.Type}}{{$.Thing}}) slice{{camel .Tag}}List(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]{{.Type}}, error) {
	q := tbl.Dialect().Quoter()
	whs, args := where.Where(wh, q)
	orderBy := where.BuildQueryConstraint(qc, q)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", q.Quote(sqlname), tbl.quotedName(), whs, orderBy)
	rows, err := support.Query(tbl, query, args...)
	if err != nil {
		return nil, err
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
{{- end}}
`

var tSliceItem = template.Must(template.New("SliceItem").Funcs(funcMap).Parse(sSliceItem))

//-------------------------------------------------------------------------------------------------

const sConstructInsert = `
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) construct{{.Prefix}}{{.Type}}Insert(w dialect.StringWriter, v *{{.Type}}, withPk bool) (s []interface{}, err error) {
{{range .Body1}}{{.}}{{end}}
{{range .Body2}}{{.}}{{end}}
	return s, nil
}
`

var tConstructInsert = template.Must(template.New("ConstructInsert").Funcs(funcMap).Parse(sConstructInsert))

//-------------------------------------------------------------------------------------------------

const sConstructUpdate = `
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) construct{{.Prefix}}{{.Type}}Update(w dialect.StringWriter, v *{{.Type}}) (s []interface{}, err error) {
{{range .Body1}}{{.}}{{end}}
{{range .Body2}}{{.}}{{end}}
	return s, nil
}
`

var tConstructUpdate = template.Must(template.New("ConstructUpdate").Funcs(funcMap).Parse(sConstructUpdate))

//-------------------------------------------------------------------------------------------------

const sInsert = `
//--------------------------------------------------------------------------------

// Insert adds new records for the {{.Types}}.
{{if .Table.HasLastInsertId}}// The {{.Types}} have their primary key fields set to the new record identifiers.{{end}}
// The {{.Type}}.PreInsert() method will be called, if it exists.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Insert(req require.Requirement, vv ...*{{.Type}}) error {
	if req == require.All {
		req = require.Exactly(len(vv))
	}

	var count int64
	insertHasReturningPhrase := {{if .Table.HasLastInsertId -}}tbl.Dialect().InsertHasReturningPhrase(){{else}}false{{end}}
	returning := ""
	{{if .Table.Primary -}}
	if tbl.Dialect().InsertHasReturningPhrase() {
		returning = fmt.Sprintf(" returning %q", tbl.pk)
	}

	{{end -}}
	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlapi.CanPreInsert); ok {
			err := hook.PreInsert()
			if err != nil {
				return tbl.logError(err)
			}
		}

		b := dialect.Adapt(&bytes.Buffer{})
		b.WriteString("INSERT INTO ")
		tbl.quotedNameW(b)

		fields, err := tbl.construct{{.Prefix}}{{.Type}}Insert(b, v, {{not .Table.HasLastInsertId}})
		if err != nil {
			return tbl.logError(err)
		}

		b.WriteString(" VALUES (")
		b.WriteString(tbl.Dialect().Placeholders(len(fields)))
		b.WriteString(")")
		b.WriteString(returning)

		query := b.String()
		tbl.logQuery(query, fields...)

		var n int64 = 1
		if insertHasReturningPhrase {
			row := tbl.db.QueryRowContext(tbl.ctx, query, fields...)
			{{- if .Table.HasLastInsertId}}
			{{- if eq .Table.Primary.Type.Name "int64"}}
			err = row.Scan(&v.{{.Table.Primary.Name}})
			{{- else}}
			var i64 int64
			err = row.Scan(&i64)
			v.{{.Table.Primary.Name}} = {{.Table.Primary.Type.Name}}(i64)
			{{- end}}
			{{- else}}
			var i64 int64
			err = row.Scan(&i64)
			{{- end}}

		} else {
			res, e2 := tbl.db.ExecContext(tbl.ctx, query, fields...)
			if e2 != nil {
				return tbl.logError(e2)
			}

			{{if .Table.HasLastInsertId -}}
			{{if eq .Table.Primary.Type.Name "int64" -}}
			v.{{.Table.Primary.Name}}, err = res.LastInsertId()
			{{else -}}
			i64, e2 := res.LastInsertId()
			v.{{.Table.Primary.Name}} = {{.Table.Primary.Type.Name}}(i64)
			{{end -}}
			{{end -}}
			if e2 != nil {
				return tbl.logError(e2)
			}

			n, err = res.RowsAffected()
		}

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
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) UpdateFields(req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error) {
	return support.UpdateFields(tbl, req, wh, fields...)
}
`

var tUpdateFields = template.Must(template.New("UpdateFields").Funcs(funcMap).Parse(sUpdateFields))

//-------------------------------------------------------------------------------------------------

// function template to update rows.
const sUpdate = `
{{- if .Table.Primary}}
//--------------------------------------------------------------------------------

// Update updates records, matching them by primary key. It returns the number of rows affected.
// The {{.Type}}.PreUpdate(Execer) method will be called, if it exists.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Update(req require.Requirement, vv ...*{{.Type}}) (int64, error) {
	if req == require.All {
		req = require.Exactly(len(vv))
	}

	var count int64
	d := tbl.Dialect()
	q := d.Quoter()

	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlapi.CanPreUpdate); ok {
			err := hook.PreUpdate()
			if err != nil {
				return count, tbl.logError(err)
			}
		}

		b := dialect.Adapt(&bytes.Buffer{})
		b.WriteString("UPDATE ")
		tbl.quotedNameW(b)
		b.WriteString(" SET ")

		args, err := tbl.construct{{.Prefix}}{{.Type}}Update(b, v)
		if err != nil {
			return count, err
		}
		args = append(args, v.{{.Table.Primary.Name}})

		b.WriteString(" WHERE ")
		q.QuoteW(b, tbl.pk)
		b.WriteString("=?")

		query := b.String()
		n, err := tbl.Exec(nil, query, args...)
		if err != nil {
			return count, err
		}
		count += n
	}

	return count, tbl.logIfError(require.ErrorIfExecNotSatisfiedBy(req, count))
}
{{- end}}
`

var tUpdate = template.Must(template.New("Update").Funcs(funcMap).Parse(sUpdate))

//-------------------------------------------------------------------------------------------------

const sDelete = `
{{- if .Table.Primary}}
// Delete{{.Types}} deletes rows from the table, given some primary keys.
// The list of ids can be arbitrarily long.
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Delete{{.Types}}(req require.Requirement, id ...{{.Table.Primary.Type.Type}}) (int64, error) {
	const batch = 1000 // limited by Oracle DB
	const qt = "DELETE FROM %s WHERE %s IN (%s)"
	qName := tbl.quotedName()

	if req == require.All {
		req = require.Exactly(len(id))
	}

	var count, n int64
	var err error
	var max = batch
	if len(id) < batch {
		max = len(id)
	}
	d := tbl.Dialect()
	col := d.Quoter().Quote(tbl.pk)
	args := make([]interface{}, max)

	if len(id) > batch {
		pl := d.Placeholders(batch)
		query := fmt.Sprintf(qt, qName, col, pl)

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
		pl := d.Placeholders(len(id))
		query := fmt.Sprintf(qt, qName, col, pl)

		for i := 0; i < len(id); i++ {
			args[i] = id[i]
		}

		n, err = tbl.Exec(nil, query, args...)
		count += n
	}

	return count, tbl.logIfError(require.ChainErrorIfExecNotSatisfiedBy(err, req, n))
}
{{- end}}

// Delete deletes one or more rows from the table, given a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) Delete(req require.Requirement, wh where.Expression) (int64, error) {
	query, args := tbl.deleteRows(wh)
	return tbl.Exec(req, query, args...)
}

func (tbl {{.Prefix}}{{.Type}}{{.Thing}}) deleteRows(wh where.Expression) (string, []interface{}) {
	whs, args := where.Build(" WHERE ", wh, tbl.Dialect().Quoter())
	query := fmt.Sprintf("DELETE FROM %s%s", tbl.quotedName(), whs)
	return query, args
}
`

var tDelete = template.Must(template.New("Delete").Funcs(funcMap).Parse(sDelete))
