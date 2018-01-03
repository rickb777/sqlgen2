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
	id := &Field{Node{"Id", Type{"", "", "int64", Int64}, nil}, "id", ENCNONE, Tag{}}
	name := &Field{Node{"Name", Type{"", "", "string", String}, nil}, "name", ENCNONE, Tag{}}

	return &TableDescription{
		Type: "Example",
		Name: "examples",
		Fields: []*Field{
			id,
			name,
		},
		Primary: id,
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

const XExampleColumnNames = "id, name"
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
	view.Table = simpleFixtureTable()
	view.Table.Primary = nil

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
	id,
	name
) VALUES (%s)
|

const sqlInsertXExamplePostgres = |
INSERT INTO %s%s (
	id,
	name
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
	id,
	name
) VALUES (%s)
|

const sqlInsertXExamplePostgres = |
INSERT INTO %s%s (
	id,
	name
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
