package code

import (
	"bytes"
	. "github.com/rickb777/sqlapi/schema"
	. "github.com/rickb777/sqlapi/types"
	"github.com/rickb777/sqlgen2/parse/exit"
	"strings"
	"testing"
)

func simpleFixtureTable() *TableDescription {
	id := &Field{Node{"Id", Type{Name: "int64", Base: Int64}, nil}, "id", ENCNONE, &Tag{Primary: true, Auto: true}}
	name := &Field{Node{"Name", Type{Name: "string", Base: String}, nil}, "name", ENCNONE, nil}
	age := &Field{Node{"Age", Type{Name: "int", Base: Int}, nil}, "age", ENCNONE, nil}

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
	name := &Field{Node{"Name", Type{Name: "string", Base: String}, nil}, "name", ENCNONE, nil}
	age := &Field{Node{"Age", Type{Name: "int", Base: Int}, nil}, "age", ENCNONE, nil}

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

func TestWriteQueryRows(t *testing.T) {
	exit.TestableExit()

	view := NewView("", "", "Example", "X", "", "", "sql", "sqlapi")
	view.Table = simpleFixtureTable()

	buf1 := &bytes.Buffer{}
	buf2 := &bytes.Buffer{}

	WriteQueryRows(buf1, buf2, view)

	code := buf2.String()
	expected := strings.Replace(`
//--------------------------------------------------------------------------------

// Query is the low-level request method for this table. The SQL query must return all the columns necessary for
// Example values. Placeholders should be vanilla '?' marks, which will be replaced if necessary according to
// the chosen dialect.
//
// The query is logged using whatever logger is configured. If an error arises, this too is logged.
//
// If you need a context other than the background, use WithContext before calling Query.
//
// The args are for any placeholder parameters in the query.
//
// The support API provides a core 'support.Query' function, on which this method depends. If appropriate,
// use that function directly; wrap the result in *sqlapi.Rows if you need to access its data as a map.
func (tbl XExampleTable) Query(req require.Requirement, query string, args ...interface{}) ([]*Example, error) {
	return doXExampleTableQueryAndScan(tbl, req, false, query, args)
}

func doXExampleTableQueryAndScan(tbl XExampleTabler, req require.Requirement, firstOnly bool, query string, args ...interface{}) ([]*Example, error) {
	rows, err := support.Query(tbl, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vv, n, err := scanXExamples(query, rows, firstOnly)
	return vv, tbl.Logger().LogIfError(require.ChainErrorIfQueryNotSatisfiedBy(err, req, n))
}
`, "¬", "`", -1)
	if code != expected {
		outputDiff(expected, "expected.txt")
		outputDiff(code, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
	}
}

//-------------------------------------------------------------------------------------------------

func TestWriteQueryThings(t *testing.T) {
	exit.TestableExit()

	view := NewView("", "", "Example", "X", "", "", "sql", "sqlapi")
	view.Table = simpleFixtureTable()

	buf1 := &bytes.Buffer{}
	buf2 := &bytes.Buffer{}

	WriteQueryThings(buf1, buf2, view)

	code := buf2.String()
	expected := strings.Replace(`
//--------------------------------------------------------------------------------

// QueryOneNullString is a low-level access method for one string. This can be used for function queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl XExampleTable) QueryOneNullString(req require.Requirement, query string, args ...interface{}) (result sql.NullString, err error) {
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
func (tbl XExampleTable) QueryOneNullInt64(req require.Requirement, query string, args ...interface{}) (result sql.NullInt64, err error) {
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
func (tbl XExampleTable) QueryOneNullFloat64(req require.Requirement, query string, args ...interface{}) (result sql.NullFloat64, err error) {
	err = support.QueryOneNullThing(tbl, req, &result, query, args...)
	return result, err
}
`, "¬", "`", -1)
	if code != expected {
		outputDiff(expected, "expected.txt")
		outputDiff(code, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
	}
}

//-------------------------------------------------------------------------------------------------

func TestWriteUpdateFunc_noPK(t *testing.T) {
	exit.TestableExit()

	view := NewView("", "", "Example", "X", "", "", "sql", "sqlapi")
	view.Table = simpleNoPKTable()

	buf1 := &bytes.Buffer{}
	buf2 := &bytes.Buffer{}

	WriteUpdateFunc(buf1, buf2, view)

	code := buf2.String()
	expected := strings.Replace(`
// UpdateFields updates one or more columns, given a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl XExampleTable) UpdateFields(req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error) {
	return support.UpdateFields(tbl, req, wh, fields...)
}
`, "¬", "`", -1)
	if code != expected {
		outputDiff(expected, "expected.txt")
		outputDiff(code, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
	}
}

func TestWriteExecFunc(t *testing.T) {
	exit.TestableExit()

	view := NewView("", "", "Example", "X", "", "", "sql", "sqlapi")
	view.Table = simpleFixtureTable()

	buf := &bytes.Buffer{}

	WriteExecFunc(buf, view)

	code := buf.String()
	expected := strings.Replace(`
//--------------------------------------------------------------------------------

// Exec executes a query without returning any rows.
// It returns the number of rows affected (if the database driver supports this).
//
// The args are for any placeholder parameters in the query.
func (tbl XExampleTable) Exec(req require.Requirement, query string, args ...interface{}) (int64, error) {
	return support.Exec(tbl, req, query, args...)
}
`, "¬", "`", -1)
	if code != expected {
		outputDiff(expected, "expected.txt")
		outputDiff(code, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
	}
}
