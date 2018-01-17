package code

import (
	"bytes"
	"testing"
	"github.com/rickb777/sqlgen2/sqlgen/parse/exit"
	. "github.com/rickb777/sqlgen2/schema"
	. "github.com/rickb777/sqlgen2/sqlgen/parse"
	"strings"
)

func simpleFixtureTable() *TableDescription {
	id := &Field{Node{"Id", Type{Name: "int64", Base: Int64}, nil}, "id", ENCNONE, Tag{Primary: true, Auto: true}}
	name := &Field{Node{"Name", Type{Name: "string", Base: String}, nil}, "name", ENCNONE, Tag{}}
	age := &Field{Node{"Age", Type{Name: "int", Base: Int}, nil}, "age", ENCNONE, Tag{}}

	return &TableDescription{
		Type: "Example",
		Name: "examples",
		Fields: FieldList{
			id,
			name,
			age,
		},
		Primary: id,
	}
}

func simpleNoPKTable() *TableDescription {
	name := &Field{Node{"Name", Type{Name: "string", Base: String}, nil}, "name", ENCNONE, Tag{}}
	age := &Field{Node{"Age", Type{Name: "int", Base: Int}, nil}, "age", ENCNONE, Tag{}}

	return &TableDescription{
		Type: "Example",
		Name: "examples",
		Fields: FieldList{
			name,
			age,
		},
	}
}

//-------------------------------------------------------------------------------------------------

func TestWriteQueryFuncs(t *testing.T) {
	exit.TestableExit()

	view := NewView("Example", "X", "")
	view.Table = simpleFixtureTable()

	buf := &bytes.Buffer{}

	WriteQueryFuncs(buf, view)

	code := buf.String()
	expected := `
//--------------------------------------------------------------------------------

// QueryOne is the low-level access function for one Example.
// If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, *Example will be nil.
func (tbl XExampleTable) QueryOne(query string, args ...interface{}) (*Example, error) {
	list, err := tbl.doQuery(true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

// Query is the low-level access function for Examples.
func (tbl XExampleTable) Query(query string, args ...interface{}) ([]*Example, error) {
	return tbl.doQuery(false, query, args...)
}

func (tbl XExampleTable) doQuery(firstOnly bool, query string, args ...interface{}) ([]*Example, error) {
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanXExamples(rows, firstOnly)
}
`
	if code != expected {
		outputDiff(expected, "expected.txt")
		outputDiff(code, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
	}
}

//-------------------------------------------------------------------------------------------------

func TestWriteGetRow(t *testing.T) {
	exit.TestableExit()

	view := NewView("Example", "X", "")
	view.Table = simpleFixtureTable()

	buf := &bytes.Buffer{}

	WriteGetRow(buf, view)

	code := buf.String()
	expected := `
//--------------------------------------------------------------------------------

// GetExample gets the record with a given primary key value.
// If not found, *Example will be nil.
func (tbl XExampleTable) GetExample(id int64) (*Example, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE id=?", XExampleColumnNames, tbl.name)
	return tbl.QueryOne(query, id)
}

// GetExamples gets records from the table according to a list of primary keys.
// Although the list of ids can be arbitrarily long, there are practical limits;
// note that Oracle DB has a limit of 1000.
func (tbl XExampleTable) GetExamples(id ...int64) (list []*Example, err error) {
	if len(id) > 0 {
		pl := tbl.dialect.Placeholders(len(id))
		query := fmt.Sprintf("SELECT %s FROM %s WHERE id IN (%s)", XExampleColumnNames, tbl.name, pl)
		args := make([]interface{}, len(id))

		for i, v := range id {
			args[i] = v
		}

		list, err = tbl.Query(query, args...)
	}

	return list, err
}
`
	if code != expected {
		outputDiff(expected, "expected.txt")
		outputDiff(code, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
	}
}

//-------------------------------------------------------------------------------------------------

func xTestWriteSelectItem(t *testing.T) {
	exit.TestableExit()

	view := NewView("Example", "X", "")
	view.Table = simpleFixtureTable()

	buf := &bytes.Buffer{}

	WriteSliceColumn(buf, view)

	code := buf.String()
	expected := `
//--------------------------------------------------------------------------------

// Get gets the record with a given primary key value.
func (tbl XExampleTable) Get(id int64) (*Example, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE id=?", XExampleColumnNames, tbl.name)
	return tbl.QueryOne(query, id)
}
`
	if code != expected {
		outputDiff(expected, "expected.txt")
		outputDiff(code, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
	}
}

func TestWriteSelectRow(t *testing.T) {
	exit.TestableExit()

	view := NewView("Example", "X", "")
	view.Table = simpleFixtureTable()

	buf := &bytes.Buffer{}

	WriteSelectRowsFuncs(buf, view)

	code := buf.String()
	expected := `
//--------------------------------------------------------------------------------

// SelectOneWhere allows a single Example to be obtained from the table that match a 'where' clause
// and some limit. Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
// If not found, *Example will be nil.
func (tbl XExampleTable) SelectOneWhere(where, orderBy string, args ...interface{}) (*Example, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT 1", XExampleColumnNames, tbl.name, where, orderBy)
	return tbl.QueryOne(query, args...)
}

// SelectOne allows a single Example to be obtained from the sqlgen2.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
// If not found, *Example will be nil.
func (tbl XExampleTable) SelectOne(wh where.Expression, qc where.QueryConstraint) (*Example, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	return tbl.SelectOneWhere(whs, orderBy, args...)
}

// SelectWhere allows Examples to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
func (tbl XExampleTable) SelectWhere(where, orderBy string, args ...interface{}) ([]*Example, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", XExampleColumnNames, tbl.name, where, orderBy)
	return tbl.Query(query, args...)
}

// Select allows Examples to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl XExampleTable) Select(wh where.Expression, qc where.QueryConstraint) ([]*Example, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	return tbl.SelectWhere(whs, orderBy, args...)
}

// CountWhere counts Examples in the table that match a 'where' clause.
// Use a blank string for the 'where' argument if it is not needed.
func (tbl XExampleTable) CountWhere(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s %s", tbl.name, where)
	tbl.logQuery(query, args...)
	row := tbl.db.QueryRowContext(tbl.ctx, query, args...)
	err = row.Scan(&count)
	return count, err
}

// Count counts the Examples in the table that match a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed.
func (tbl XExampleTable) Count(wh where.Expression) (count int64, err error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	return tbl.CountWhere(whs, args...)
}

const XExampleColumnNames = "id, name, age"
`
	if code != expected {
		outputDiff(expected, "expected.txt")
		outputDiff(code, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
	}
}

//-------------------------------------------------------------------------------------------------

func TestWriteInsertFunc_noPK(t *testing.T) {
	exit.TestableExit()

	view := NewView("Example", "X", "")
	view.Table = simpleNoPKTable()

	buf := &bytes.Buffer{}

	WriteInsertFunc(buf, view)

	code := buf.String()
	expected := strings.Replace(`
//--------------------------------------------------------------------------------

// Insert adds new records for the Examples.
// The Example.PreInsert(Execer) method will be called, if it exists.
func (tbl XExampleTable) Insert(vv ...*Example) error {
	var params string
	switch tbl.dialect {
	case schema.Postgres:
		params = sXExampleDataColumnParamsPostgres
	default:
		params = sXExampleDataColumnParamsSimple
	}

	query := fmt.Sprintf(sqlInsertXExample, tbl.name, params)
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

		fields, err := sliceXExample(v)
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

const sqlInsertXExample = |
INSERT INTO %s (
	name,
	age
) VALUES (%s)
|

const sXExampleDataColumnParamsSimple = "?,?"

const sXExampleDataColumnParamsPostgres = "$1,$2"
`, "|", "`", -1)
	if code != expected {
		outputDiff(expected, "expected.txt")
		outputDiff(code, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
	}
}

func TestWriteInsertFunc_withPK(t *testing.T) {
	exit.TestableExit()

	view := NewView("Example", "X", "")
	view.Table = simpleFixtureTable()

	buf := &bytes.Buffer{}

	WriteInsertFunc(buf, view)

	code := buf.String()
	expected := strings.Replace(`
//--------------------------------------------------------------------------------

// Insert adds new records for the Examples. The Examples have their primary key fields
// set to the new record identifiers.
// The Example.PreInsert(Execer) method will be called, if it exists.
func (tbl XExampleTable) Insert(vv ...*Example) error {
	var params string
	switch tbl.dialect {
	case schema.Postgres:
		params = sXExampleDataColumnParamsPostgres
	default:
		params = sXExampleDataColumnParamsSimple
	}

	query := fmt.Sprintf(sqlInsertXExample, tbl.name, params)
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

		fields, err := sliceXExampleWithoutPk(v)
		if err != nil {
			return err
		}

		tbl.logQuery(query, fields...)
		res, err := st.Exec(fields...)
		if err != nil {
			return err
		}

		v.Id, err = res.LastInsertId()
		if err != nil {
			return err
		}
	}

	return nil
}

const sqlInsertXExample = |
INSERT INTO %s (
	name,
	age
) VALUES (%s)
|

const sXExampleDataColumnParamsSimple = "?,?"

const sXExampleDataColumnParamsPostgres = "$1,$2"
`, "|", "`", -1)
	if code != expected {
		outputDiff(expected, "expected.txt")
		outputDiff(code, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
	}
}

//-------------------------------------------------------------------------------------------------

func TestWriteUpdateFunc_noPK(t *testing.T) {
	exit.TestableExit()

	view := NewView("Example", "X", "")
	view.Table = simpleNoPKTable()

	buf := &bytes.Buffer{}

	WriteUpdateFunc(buf, view)

	code := buf.String()
	expected := strings.Replace(`
//--------------------------------------------------------------------------------

// UpdateFields updates one or more columns, given a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl XExampleTable) UpdateFields(wh where.Expression, fields ...sql.NamedArg) (int64, error) {
	query, args := tbl.updateFields(wh, fields...)
	return tbl.Exec(query, args...)
}

func (tbl XExampleTable) updateFields(wh where.Expression, fields ...sql.NamedArg) (string, []interface{}) {
	list := sqlgen2.NamedArgList(fields)
	assignments := strings.Join(list.Assignments(tbl.dialect, 1), ", ")
	whs, wargs := where.BuildExpression(wh, tbl.dialect)
	query := fmt.Sprintf("UPDATE %s SET %s %s", tbl.name, assignments, whs)
	args := append(list.Values(), wargs...)
	return query, args
}
`, "|", "`", -1)
	if code != expected {
		outputDiff(expected, "expected.txt")
		outputDiff(code, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
	}
}

func TestWriteUpdateFunc_withPK(t *testing.T) {
	exit.TestableExit()

	view := NewView("Example", "X", "")
	view.Table = simpleFixtureTable()

	buf := &bytes.Buffer{}

	WriteUpdateFunc(buf, view)

	code := buf.String()
	expected := strings.Replace(`
//--------------------------------------------------------------------------------

// UpdateFields updates one or more columns, given a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl XExampleTable) UpdateFields(wh where.Expression, fields ...sql.NamedArg) (int64, error) {
	query, args := tbl.updateFields(wh, fields...)
	return tbl.Exec(query, args...)
}

func (tbl XExampleTable) updateFields(wh where.Expression, fields ...sql.NamedArg) (string, []interface{}) {
	list := sqlgen2.NamedArgList(fields)
	assignments := strings.Join(list.Assignments(tbl.dialect, 1), ", ")
	whs, wargs := where.BuildExpression(wh, tbl.dialect)
	query := fmt.Sprintf("UPDATE %s SET %s %s", tbl.name, assignments, whs)
	args := append(list.Values(), wargs...)
	return query, args
}

//--------------------------------------------------------------------------------

// Update updates records, matching them by primary key. It returns the number of rows affected.
// The Example.PreUpdate(Execer) method will be called, if it exists.
func (tbl XExampleTable) Update(vv ...*Example) (int64, error) {
	var stmt string
	switch tbl.dialect {
	case schema.Postgres:
		stmt = sqlUpdateXExampleByPkPostgres
	default:
		stmt = sqlUpdateXExampleByPkSimple
	}

	query := fmt.Sprintf(stmt, tbl.name)

	var count int64
	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPreUpdate); ok {
			hook.PreUpdate()
		}

		args, err := sliceXExampleWithoutPk(v)
		if err != nil {
			return count, err
		}

		args = append(args, v.Id)
		n, err := tbl.Exec(query, args...)
		if err != nil {
			return count, err
		}
		count += n
	}
	return count, nil
}

const sqlUpdateXExampleByPkSimple = |
UPDATE %s SET
	name=?,
	age=?
WHERE id=?
|

const sqlUpdateXExampleByPkPostgres = |
UPDATE %s SET
	name=$2,
	age=$3
WHERE id=$1
|
`, "|", "`", -1)
	if code != expected {
		outputDiff(expected, "expected.txt")
		outputDiff(code, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
	}
}

func TestWriteDeleteFunc(t *testing.T) {
	exit.TestableExit()

	view := NewView("Example", "X", "")
	view.Table = simpleFixtureTable()

	buf := &bytes.Buffer{}

	WriteDeleteFunc(buf, view)

	code := buf.String()
	expected := strings.Replace(`
//--------------------------------------------------------------------------------

// DeleteExamples deletes rows from the table, given some primary keys.
// The list of ids can be arbitrarily long.
func (tbl XExampleTable) DeleteExamples(id ...int64) (int64, error) {
	const batch = 1000 // limited by Oracle DB
	const qt = "DELETE FROM %s WHERE id IN (%s)"

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
		query := fmt.Sprintf(qt, tbl.name, pl)

		for i := 0; i < len(id); i++ {
			args[i] = id[i]
		}

		n, err = tbl.Exec(query, args...)
		count += n
	}

	return count, err
}

// Delete deletes one or more rows from the table, given a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl XExampleTable) Delete(wh where.Expression) (int64, error) {
	query, args := tbl.deleteRows(wh)
	return tbl.Exec(query, args...)
}

func (tbl XExampleTable) deleteRows(wh where.Expression) (string, []interface{}) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	query := fmt.Sprintf("DELETE FROM %s %s", tbl.name, whs)
	return query, args
}

//--------------------------------------------------------------------------------
`, "|", "`", -1)
	if code != expected {
		outputDiff(expected, "expected.txt")
		outputDiff(code, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
	}
}

func TestWriteExecFunc(t *testing.T) {
	exit.TestableExit()

	view := NewView("Example", "X", "")
	view.Table = simpleFixtureTable()

	buf := &bytes.Buffer{}

	WriteExecFunc(buf, view)

	code := buf.String()
	expected := strings.Replace(`
//--------------------------------------------------------------------------------

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// It returns the number of rows affected (of the database drive supports this).
func (tbl XExampleTable) Exec(query string, args ...interface{}) (int64, error) {
	tbl.logQuery(query, args...)
	res, err := tbl.db.ExecContext(tbl.ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
`, "|", "`", -1)
	if code != expected {
		outputDiff(expected, "expected.txt")
		outputDiff(code, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
	}
}
