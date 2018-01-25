package code

import (
	"bytes"
	"testing"
	"github.com/rickb777/sqlgen2/sqlgen/parse/exit"
	. "github.com/rickb777/sqlgen2/model"
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

func TestWriteQueryRows(t *testing.T) {
	exit.TestableExit()

	view := NewView("Example", "X", "")
	view.Table = simpleFixtureTable()

	buf := &bytes.Buffer{}

	WriteQueryRows(buf, view)

	code := buf.String()
	expected := strings.Replace(`
//--------------------------------------------------------------------------------

// Query is the low-level access method for Examples.
//
// It places a requirement, which may be nil, on the size of the expected results: this
// controls whether an error is generated when this expectation is not met.
//
// Note that this method applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl XExampleTable) Query(req require.Requirement, query string, args ...interface{}) ([]*Example, error) {
	query = tbl.ReplaceTableName(query)
	vv, err := tbl.doQuery(req, false, query, args...)
	return vv, err
}

// QueryOne is the low-level access method for one Example.
// If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, *Example will be nil.
//
// Note that this method applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl XExampleTable) QueryOne(query string, args ...interface{}) (*Example, error) {
	query = tbl.ReplaceTableName(query)
	return tbl.doQueryOne(nil, query, args...)
}

// MustQueryOne is the low-level access method for one Example.
//
// It places a requirement that exactly one result must be found; an error is generated when this expectation is not met.
//
// Note that this method applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl XExampleTable) MustQueryOne(query string, args ...interface{}) (*Example, error) {
	query = tbl.ReplaceTableName(query)
	return tbl.doQueryOne(require.One, query, args...)
}

func (tbl XExampleTable) doQueryOne(req require.Requirement, query string, args ...interface{}) (*Example, error) {
	list, err := tbl.doQuery(req, true, query, args...)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

func (tbl XExampleTable) doQuery(req require.Requirement, firstOnly bool, query string, args ...interface{}) ([]*Example, error) {
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, tbl.logError(err)
	}
	defer rows.Close()

	vv, n, err := scanXExamples(rows, firstOnly)
	return vv, tbl.logIfError(require.ChainErrorIfQueryNotSatisfiedBy(err, req, n))
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

	view := NewView("Example", "X", "")
	view.Table = simpleFixtureTable()

	buf := &bytes.Buffer{}

	WriteQueryThings(buf, view)

	code := buf.String()
	expected := strings.Replace(`
//--------------------------------------------------------------------------------

// QueryOneNullString is a low-level access method for one string. This can be used for function queries and
// such like. If the query selected many rows, only the first is returned; the rest are discarded.
// If not found, the result will be invalid.
//
// Note that this applies ReplaceTableName to the query string.
//
// The args are for any placeholder parameters in the query.
func (tbl XExampleTable) QueryOneNullString(query string, args ...interface{}) (result sql.NullString, err error) {
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
func (tbl XExampleTable) MustQueryOneNullString(query string, args ...interface{}) (result sql.NullString, err error) {
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
func (tbl XExampleTable) QueryOneNullInt64(query string, args ...interface{}) (result sql.NullInt64, err error) {
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
func (tbl XExampleTable) MustQueryOneNullInt64(query string, args ...interface{}) (result sql.NullInt64, err error) {
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
func (tbl XExampleTable) QueryOneNullFloat64(query string, args ...interface{}) (result sql.NullFloat64, err error) {
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
func (tbl XExampleTable) MustQueryOneNullFloat64(query string, args ...interface{}) (result sql.NullFloat64, err error) {
	err = support.QueryOneNullThing(tbl, require.One, &result, query, args...)
	return result, err
}

// ReplaceTableName replaces all occurrences of "{TABLE}" with the table's name.
func (tbl XExampleTable) ReplaceTableName(query string) string {
	return strings.Replace(query, "{TABLE}", tbl.name.String(), -1)
}
`, "¬", "`", -1)
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
	expected := strings.Replace(`
//--------------------------------------------------------------------------------

var allXExampleQuotedColumnNames = []string{
	schema.Sqlite.SplitAndQuote(XExampleColumnNames),
	schema.Mysql.SplitAndQuote(XExampleColumnNames),
	schema.Postgres.SplitAndQuote(XExampleColumnNames),
}

//--------------------------------------------------------------------------------

// GetExample gets the record with a given primary key value.
// If not found, *Example will be nil.
func (tbl XExampleTable) GetExample(id int64) (*Example, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s=?",
		allXExampleQuotedColumnNames[tbl.dialect.Index()], tbl.name, tbl.dialect.Quote("id"))
	v, err := tbl.doQueryOne(nil, query, id)
	return v, err
}

// MustGetExample gets the record with a given primary key value.
//
// It places a requirement that exactly one result must be found; an error is generated when this expectation is not met.
func (tbl XExampleTable) MustGetExample(id int64) (*Example, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s=?",
		allXExampleQuotedColumnNames[tbl.dialect.Index()], tbl.name, tbl.dialect.Quote("id"))
	v, err := tbl.doQueryOne(require.One, query, id)
	return v, err
}

// GetExamples gets records from the table according to a list of primary keys.
// Although the list of ids can be arbitrarily long, there are practical limits;
// note that Oracle DB has a limit of 1000.
//
// It places a requirement, which may be nil, on the size of the expected results: in particular, require.All
// controls whether an error is generated not all the ids produce a result.
func (tbl XExampleTable) GetExamples(req require.Requirement, id ...int64) (list []*Example, err error) {
	if len(id) > 0 {
		if req == require.All {
			req = require.Exactly(len(id))
		}
		pl := tbl.dialect.Placeholders(len(id))
		query := fmt.Sprintf("SELECT %s FROM %s WHERE %s IN (%s)",
			allXExampleQuotedColumnNames[tbl.dialect.Index()], tbl.name, tbl.dialect.Quote("id"), pl)
		args := make([]interface{}, len(id))

		for i, v := range id {
			args[i] = v
		}

		list, err = tbl.doQuery(req, false, query, args...)
	}

	return list, err
}
`, "¬", "`", -1)
	if code != expected {
		outputDiff(expected, "expected.txt")
		outputDiff(code, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
	}
}

//-------------------------------------------------------------------------------------------------

func TestWriteSelectItem(t *testing.T) {
	exit.TestableExit()

	view := NewView("Example", "X", "")
	view.Table = simpleFixtureTable()

	buf := &bytes.Buffer{}

	WriteSliceColumn(buf, view)

	code := buf.String()
	expected := strings.Replace(`
//--------------------------------------------------------------------------------

// SliceId gets the Id column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl XExampleTable) SliceId(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return tbl.getint64list(req, "id", wh, qc)
}

// SliceName gets the Name column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl XExampleTable) SliceName(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return tbl.getstringlist(req, "name", wh, qc)
}

// SliceAge gets the Age column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl XExampleTable) SliceAge(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]int, error) {
	return tbl.getintlist(req, "age", wh, qc)
}


func (tbl XExampleTable) getintlist(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]int, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", tbl.dialect.Quote(sqlname), tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, tbl.logError(err)
	}
	defer rows.Close()

	var v int
	list := make([]int, 0, 10)

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

func (tbl XExampleTable) getint64list(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", tbl.dialect.Quote(sqlname), tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, tbl.logError(err)
	}
	defer rows.Close()

	var v int64
	list := make([]int64, 0, 10)

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

func (tbl XExampleTable) getstringlist(req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", tbl.dialect.Quote(sqlname), tbl.name, whs, orderBy)
	tbl.logQuery(query, args...)
	rows, err := tbl.db.QueryContext(tbl.ctx, query, args...)
	if err != nil {
		return nil, tbl.logError(err)
	}
	defer rows.Close()

	var v string
	list := make([]string, 0, 10)

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

`, "¬", "`", -1)
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
	expected := strings.Replace(`
//--------------------------------------------------------------------------------

// SelectOneWhere allows a single Example to be obtained from the table that match a 'where' clause
// and some limit. Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
// If not found, *Example will be nil.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.One
// controls whether an error is generated when no result is found.
//
// The args are for any placeholder parameters in the query.
func (tbl XExampleTable) SelectOneWhere(req require.Requirement, where, orderBy string, args ...interface{}) (*Example, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT 1", XExampleColumnNames, tbl.name, where, orderBy)
	v, err := tbl.doQueryOne(req, query, args...)
	return v, err
}

// SelectOne allows a single Example to be obtained from the sqlgen2.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
// If not found, *Example will be nil.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.One
// controls whether an error is generated when no result is found.
func (tbl XExampleTable) SelectOne(req require.Requirement, wh where.Expression, qc where.QueryConstraint) (*Example, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	return tbl.SelectOneWhere(req, whs, orderBy, args...)
}

// SelectWhere allows Examples to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in 'orderBy'.
// Use blank strings for the 'where' and/or 'orderBy' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
//
// The args are for any placeholder parameters in the query.
func (tbl XExampleTable) SelectWhere(req require.Requirement, where, orderBy string, args ...interface{}) ([]*Example, error) {
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", XExampleColumnNames, tbl.name, where, orderBy)
	vv, err := tbl.doQuery(req, false, query, args...)
	return vv, err
}

// Select allows Examples to be obtained from the table that match a 'where' clause.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
//
// It places a requirement, which may be nil, on the size of the expected results: for example require.AtLeastOne
// controls whether an error is generated when no result is found.
func (tbl XExampleTable) Select(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]*Example, error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	orderBy := where.BuildQueryConstraint(qc, tbl.dialect)
	return tbl.SelectWhere(req, whs, orderBy, args...)
}

// CountWhere counts Examples in the table that match a 'where' clause.
// Use a blank string for the 'where' argument if it is not needed.
//
// The args are for any placeholder parameters in the query.
func (tbl XExampleTable) CountWhere(where string, args ...interface{}) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s %s", tbl.name, where)
	tbl.logQuery(query, args...)
	row := tbl.db.QueryRowContext(tbl.ctx, query, args...)
	err = row.Scan(&count)
	return count, tbl.logIfError(err)
}

// Count counts the Examples in the table that match a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed.
func (tbl XExampleTable) Count(wh where.Expression) (count int64, err error) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	return tbl.CountWhere(whs, args...)
}
`, "¬", "`", -1)
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

var allXExampleQuotedInserts = []string{
	// Sqlite
	"(¬name¬,¬age¬) VALUES (?,?)",
	// Mysql
	"(¬name¬,¬age¬) VALUES (?,?)",
	// Postgres
	¬("name","age") VALUES ($1,$2)¬,
}

//--------------------------------------------------------------------------------

// Insert adds new records for the Examples.

// The Example.PreInsert() method will be called, if it exists.
func (tbl XExampleTable) Insert(req require.Requirement, vv ...*Example) error {
	if req == require.All {
		req = require.Exactly(len(vv))
	}

	var count int64
	columns := allXExampleQuotedInserts[tbl.dialect.Index()]
	query := fmt.Sprintf("INSERT INTO %s %s", tbl.name, columns)
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

		fields, err := sliceXExample(v)
		if err != nil {
			return tbl.logError(err)
		}

		tbl.logQuery(query, fields...)
		res, err := st.ExecContext(tbl.ctx, fields...)
		if err != nil {
			return tbl.logError(err)
		}

		
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
`, "¬", "`", -1)
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

var allXExampleQuotedInserts = []string{
	// Sqlite
	"(¬name¬,¬age¬) VALUES (?,?)",
	// Mysql
	"(¬name¬,¬age¬) VALUES (?,?)",
	// Postgres
	¬("name","age") VALUES ($1,$2) returning "id"¬,
}

//--------------------------------------------------------------------------------

// Insert adds new records for the Examples.
// The Examples have their primary key fields set to the new record identifiers.
// The Example.PreInsert() method will be called, if it exists.
func (tbl XExampleTable) Insert(req require.Requirement, vv ...*Example) error {
	if req == require.All {
		req = require.Exactly(len(vv))
	}

	var count int64
	columns := allXExampleQuotedInserts[tbl.dialect.Index()]
	query := fmt.Sprintf("INSERT INTO %s %s", tbl.name, columns)
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

		fields, err := sliceXExampleWithoutPk(v)
		if err != nil {
			return tbl.logError(err)
		}

		tbl.logQuery(query, fields...)
		res, err := st.ExecContext(tbl.ctx, fields...)
		if err != nil {
			return tbl.logError(err)
		}

		v.Id, err = res.LastInsertId()
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

	view := NewView("Example", "X", "")
	view.Table = simpleNoPKTable()

	buf := &bytes.Buffer{}

	WriteUpdateFunc(buf, view)

	code := buf.String()
	expected := strings.Replace(`
// UpdateFields updates one or more columns, given a 'where' clause.
//
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

func TestWriteUpdateFunc_withPK(t *testing.T) {
	exit.TestableExit()

	view := NewView("Example", "X", "")
	view.Table = simpleFixtureTable()

	buf := &bytes.Buffer{}

	WriteUpdateFunc(buf, view)

	code := buf.String()
	expected := strings.Replace(`
// UpdateFields updates one or more columns, given a 'where' clause.
//
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl XExampleTable) UpdateFields(req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error) {
	return support.UpdateFields(tbl, req, wh, fields...)
}

//--------------------------------------------------------------------------------

var allXExampleQuotedUpdates = []string{
	// Sqlite
	"¬name¬=?,¬age¬=? WHERE ¬id¬=?",
	// Mysql
	"¬name¬=?,¬age¬=? WHERE ¬id¬=?",
	// Postgres
	¬"name"=$2,"age"=$3 WHERE "id"=$1¬,
}

//--------------------------------------------------------------------------------

// Update updates records, matching them by primary key. It returns the number of rows affected.
// The Example.PreUpdate(Execer) method will be called, if it exists.
func (tbl XExampleTable) Update(req require.Requirement, vv ...*Example) (int64, error) {
	if req == require.All {
		req = require.Exactly(len(vv))
	}

	var count int64
	columns := allXExampleQuotedUpdates[tbl.dialect.Index()]
	query := fmt.Sprintf("UPDATE %s SET %s", tbl.name, columns)

	for _, v := range vv {
		var iv interface{} = v
		if hook, ok := iv.(sqlgen2.CanPreUpdate); ok {
			err := hook.PreUpdate()
			if err != nil {
				return count, tbl.logError(err)
			}
		}

		args, err := sliceXExampleWithoutPk(v)
		args = append(args, v.Id)
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
`, "¬", "`", -1)
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
func (tbl XExampleTable) DeleteExamples(req require.Requirement, id ...int64) (int64, error) {
	const batch = 1000 // limited by Oracle DB
	const qt = "DELETE FROM %s WHERE %s IN (%s)"

	if req == require.All {
		req = require.Exactly(len(id))
	}

	var count, n int64
	var err error
	var max = batch
	if len(id) < batch {
		max = len(id)
	}
	col := tbl.dialect.Quote("id")
	args := make([]interface{}, max)

	if len(id) > batch {
		pl := tbl.dialect.Placeholders(batch)
		query := fmt.Sprintf(qt, tbl.name, col, pl)

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
		query := fmt.Sprintf(qt, tbl.name, col, pl)

		for i := 0; i < len(id); i++ {
			args[i] = id[i]
		}

		n, err = tbl.Exec(nil, query, args...)
		count += n
	}

	return count, tbl.logIfError(require.ChainErrorIfExecNotSatisfiedBy(err, req, n))
}

// Delete deletes one or more rows from the table, given a 'where' clause.
// Use a nil value for the 'wh' argument if it is not needed (very risky!).
func (tbl XExampleTable) Delete(req require.Requirement, wh where.Expression) (int64, error) {
	query, args := tbl.deleteRows(wh)
	return tbl.Exec(req, query, args...)
}

func (tbl XExampleTable) deleteRows(wh where.Expression) (string, []interface{}) {
	whs, args := where.BuildExpression(wh, tbl.dialect)
	query := fmt.Sprintf("DELETE FROM %s %s", tbl.name, whs)
	return query, args
}

//--------------------------------------------------------------------------------
`, "¬", "`", -1)
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
