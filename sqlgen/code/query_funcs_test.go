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
	id := &Field{Node{"Id", Type{"", "", "int64", Int64}, nil}, "id", ENCNONE, Tag{Primary: true, Auto: true}}
	name := &Field{Node{"Name", Type{"", "", "string", String}, nil}, "name", ENCNONE, Tag{}}
	age := &Field{Node{"Age", Type{"", "", "int", Int}, nil}, "age", ENCNONE, Tag{}}

	return &TableDescription{
		Type: "Example",
		Name: "examples",
		Fields: []*Field{
			id,
			name,
			age,
		},
		Primary: id,
	}
}

func simpleNoPKTable() *TableDescription {
	name := &Field{Node{"Name", Type{"", "", "string", String}, nil}, "name", ENCNONE, Tag{}}
	age := &Field{Node{"Age", Type{"", "", "int", Int}, nil}, "age", ENCNONE, Tag{}}

	return &TableDescription{
		Type: "Example",
		Name: "examples",
		Fields: []*Field{
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
func (tbl XExampleTable) QueryOne(query string, args ...interface{}) (*Example, error) {
	tbl.logQuery(query, args...)
	row := tbl.Db.QueryRowContext(tbl.Ctx, query, args...)
	return scanXExample(row)
}

// Query is the low-level access function for Examples.
func (tbl XExampleTable) Query(query string, args ...interface{}) ([]*Example, error) {
	tbl.logQuery(query, args...)
	rows, err := tbl.Db.QueryContext(tbl.Ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanXExamples(rows)
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

// Get gets the record with a given primary key value.
func (tbl XExampleTable) Get(id int64) (*Example, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s WHERE id=?", XExampleColumnNames, tbl.Prefix, tbl.Name)
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

	WriteSelectRow(buf, view)

	code := buf.String()
	expected := `
//--------------------------------------------------------------------------------

// SelectOneSA allows a single Example to be obtained from the table that match a 'where' clause and some limit.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
func (tbl XExampleTable) SelectOneSA(where, orderBy string, args ...interface{}) (*Example, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s LIMIT 1", XExampleColumnNames, tbl.Prefix, tbl.Name, where, orderBy)
	return tbl.QueryOne(query, args...)
}

// SelectOne allows a single Example to be obtained from the sqlgen2.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
func (tbl XExampleTable) SelectOne(where where.Expression, orderBy string) (*Example, error) {
	wh, args := where.Build(tbl.Dialect)
	return tbl.SelectOneSA(wh, orderBy, args...)
}

// SelectSA allows Examples to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
func (tbl XExampleTable) SelectSA(where, orderBy string, args ...interface{}) ([]*Example, error) {
	query := fmt.Sprintf("SELECT %s FROM %s%s %s %s", XExampleColumnNames, tbl.Prefix, tbl.Name, where, orderBy)
	return tbl.Query(query, args...)
}

// Select allows Examples to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
func (tbl XExampleTable) Select(where where.Expression, orderBy string) ([]*Example, error) {
	wh, args := where.Build(tbl.Dialect)
	return tbl.SelectSA(wh, orderBy, args...)
}

// CountSA counts Examples in the table that match a 'where' clause.
func (tbl XExampleTable) CountSA(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s%s %s", tbl.Prefix, tbl.Name, where)
	tbl.logQuery(query, args...)
	row := tbl.Db.QueryRowContext(tbl.Ctx, query, args...)
	err = row.Scan(&count)
	return count, err
}

// Count counts the Examples in the table that match a 'where' clause.
func (tbl XExampleTable) Count(where where.Expression) (count int64, err error) {
	wh, args := where.Build(tbl.Dialect)
	return tbl.CountSA(wh, args...)
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
	var stmt, params string
	switch tbl.Dialect {
	case schema.Postgres:
		stmt = sqlInsertXExamplePostgres
		params = sXExampleDataColumnParamsPostgres
	default:
		stmt = sqlInsertXExampleSimple
		params = sXExampleDataColumnParamsSimple
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

const sqlInsertXExampleSimple = |
INSERT INTO %s%s (
	name,
	age
) VALUES (%s)
|

const sqlInsertXExamplePostgres = |
INSERT INTO %s%s (
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
	var stmt, params string
	switch tbl.Dialect {
	case schema.Postgres:
		stmt = sqlInsertXExamplePostgres
		params = sXExampleDataColumnParamsPostgres
	default:
		stmt = sqlInsertXExampleSimple
		params = sXExampleDataColumnParamsSimple
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

const sqlInsertXExampleSimple = |
INSERT INTO %s%s (
	name,
	age
) VALUES (%s)
|

const sqlInsertXExamplePostgres = |
INSERT INTO %s%s (
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
func (tbl XExampleTable) UpdateFields(where where.Expression, fields ...sql.NamedArg) (int64, error) {
	query, args := tbl.updateFields(where, fields...)
	return tbl.Exec(query, args...)
}

func (tbl XExampleTable) updateFields(where where.Expression, fields ...sql.NamedArg) (string, []interface{}) {
	list := sqlgen2.NamedArgList(fields)
	assignments := strings.Join(list.Assignments(tbl.Dialect, 1), ", ")
	whereClause, wargs := where.Build(tbl.Dialect)
	query := fmt.Sprintf("UPDATE %s%s SET %s %s", tbl.Prefix, tbl.Name, assignments, whereClause)
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
func (tbl XExampleTable) UpdateFields(where where.Expression, fields ...sql.NamedArg) (int64, error) {
	query, args := tbl.updateFields(where, fields...)
	return tbl.Exec(query, args...)
}

func (tbl XExampleTable) updateFields(where where.Expression, fields ...sql.NamedArg) (string, []interface{}) {
	list := sqlgen2.NamedArgList(fields)
	assignments := strings.Join(list.Assignments(tbl.Dialect, 1), ", ")
	whereClause, wargs := where.Build(tbl.Dialect)
	query := fmt.Sprintf("UPDATE %s%s SET %s %s", tbl.Prefix, tbl.Name, assignments, whereClause)
	args := append(list.Values(), wargs...)
	return query, args
}

//--------------------------------------------------------------------------------

// Update updates records, matching them by primary key. It returns the number of rows affected.
// The Example.PreUpdate(Execer) method will be called, if it exists.
func (tbl XExampleTable) Update(vv ...*Example) (int64, error) {
	var stmt string
	switch tbl.Dialect {
	case schema.Postgres:
		stmt = sqlUpdateXExampleByPkPostgres
	default:
		stmt = sqlUpdateXExampleByPkSimple
	}

	var count int64
	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPreUpdate); ok {
			hook.PreUpdate(tbl.Db)
		}

		query := fmt.Sprintf(stmt, tbl.Prefix, tbl.Name)

		args, err := sliceXExampleWithoutPk(v)
		if err != nil {
			return count, err
		}

		args = append(args, v.Id)
		tbl.logQuery(query, args...)
		n, err := tbl.Exec(query, args...)
		if err != nil {
			return count, err
		}
		count += n
	}
	return count, nil
}

// Upsert updates a record, matching it by primary key, or it inserts a new record if necessary.
func (tbl XExampleTable) Upsert(v *Example) (isnew bool, err error) {
	n, err := tbl.Update(v)
	if err != nil {
		return false, err
	}

	if n == 0 {
		isnew = true
		err = tbl.Insert(v)
		if err != nil {
			return
		}
	}

	return
}

const sqlUpdateXExampleByPkSimple = |
UPDATE %%s%%s SET
	name=?,
	age=?
WHERE id=?
|

const sqlUpdateXExampleByPkPostgres = |
UPDATE %%s%%s SET
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
