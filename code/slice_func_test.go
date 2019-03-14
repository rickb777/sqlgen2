package code

import (
	"bytes"
	"github.com/rickb777/sqlgen2/parse/exit"
	"testing"
)

func TestWriteConstructInsert(t *testing.T) {
	exit.TestableExit()

	view := NewView("Example", "X", "", "")
	view.Table = fixtureTable()
	buf := &bytes.Buffer{}

	WriteConstructInsert(buf, view)

	code := buf.String()
	expected := `
func (tbl XExampleTable) constructXExampleInsert(w dialect.StringWriter, v *Example, withPk bool) (s []interface{}, err error) {
	q := tbl.Dialect().Quoter()
	s = make([]interface{}, 0, 17)

	comma := ""
	w.WriteString(" (")

	if withPk {
		q.QuoteW(w, "id")
		comma = ","
		s = append(s, v.Id)
	}

	w.WriteString(comma)

	q.QuoteW(w, "cat")
	s = append(s, v.Cat)
	comma = ","
	w.WriteString(comma)

	q.QuoteW(w, "username")
	s = append(s, v.Name)
	if v.Mobile != nil {
		w.WriteString(comma)

		q.QuoteW(w, "mobile")
		s = append(s, v.Mobile)
	}
	if v.Qual != nil {
		w.WriteString(comma)

		q.QuoteW(w, "qual")
		s = append(s, v.Qual)
	}
	if v.Numbers.Diff != nil {
		w.WriteString(comma)

		q.QuoteW(w, "diff")
		s = append(s, v.Numbers.Diff)
	}
	if v.Numbers.Age != nil {
		w.WriteString(comma)

		q.QuoteW(w, "age")
		s = append(s, v.Numbers.Age)
	}
	if v.Numbers.Bmi != nil {
		w.WriteString(comma)

		q.QuoteW(w, "bmi")
		s = append(s, v.Numbers.Bmi)
	}
	w.WriteString(comma)

	q.QuoteW(w, "active")
	s = append(s, v.Active)
	w.WriteString(comma)

	q.QuoteW(w, "labels")
	x, err := json.Marshal(&v.Labels)
	if err != nil {
		return nil, tbl.database.LogError(errors.WithStack(err))
	}
	s = append(s, x)
	w.WriteString(comma)

	q.QuoteW(w, "fave")
	x, err := json.Marshal(&v.Fave)
	if err != nil {
		return nil, tbl.database.LogError(errors.WithStack(err))
	}
	s = append(s, x)
	w.WriteString(comma)

	q.QuoteW(w, "avatar")
	s = append(s, v.Avatar)
	w.WriteString(comma)

	q.QuoteW(w, "foo1")
	s = append(s, v.Foo1)
	if v.Foo2 != nil {
		w.WriteString(comma)

		q.QuoteW(w, "foo2")
		s = append(s, v.Foo2)
	}
	if v.Foo3 != nil {
		w.WriteString(comma)

		q.QuoteW(w, "foo3")
		s = append(s, v.Foo3)
	}
	w.WriteString(comma)

	q.QuoteW(w, "bar1")
	s = append(s, v.Bar1)
	w.WriteString(comma)

	q.QuoteW(w, "updated")
	x, err := encoding.MarshalText(&v.Updated)
	if err != nil {
		return nil, tbl.database.LogError(errors.WithStack(err))
	}
	s = append(s, x)
	w.WriteString(")")
	return s, nil
}
`
	if code != expected {
		outputDiff(expected, "expected.txt")
		outputDiff(code, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
	}
}

func TestWriteConstructUpdate(t *testing.T) {
	exit.TestableExit()

	view := NewView("Example", "X", "", "")
	view.Table = fixtureTable()
	buf := &bytes.Buffer{}

	WriteConstructUpdate(buf, view)

	code := buf.String()
	expected := `
func (tbl XExampleTable) constructXExampleUpdate(w dialect.StringWriter, v *Example) (s []interface{}, err error) {
	q := tbl.Dialect().Quoter()
	j := 1
	s = make([]interface{}, 0, 16)

	comma := ""

	w.WriteString(comma)
	q.QuoteW(w, "cat")
	w.WriteString("=?")
	s = append(s, v.Cat)
	comma = ", "
	j++

	w.WriteString(comma)
	q.QuoteW(w, "username")
	w.WriteString("=?")
	s = append(s, v.Name)
	j++

	w.WriteString(comma)
	if v.Mobile != nil {
		q.QuoteW(w, "mobile")
		w.WriteString("=?")
		s = append(s, v.Mobile)
		j++
	} else {
		q.QuoteW(w, "mobile")
		w.WriteString("=NULL")
	}

	w.WriteString(comma)
	if v.Qual != nil {
		q.QuoteW(w, "qual")
		w.WriteString("=?")
		s = append(s, v.Qual)
		j++
	} else {
		q.QuoteW(w, "qual")
		w.WriteString("=NULL")
	}

	w.WriteString(comma)
	if v.Numbers.Diff != nil {
		q.QuoteW(w, "diff")
		w.WriteString("=?")
		s = append(s, v.Numbers.Diff)
		j++
	} else {
		q.QuoteW(w, "diff")
		w.WriteString("=NULL")
	}

	w.WriteString(comma)
	if v.Numbers.Age != nil {
		q.QuoteW(w, "age")
		w.WriteString("=?")
		s = append(s, v.Numbers.Age)
		j++
	} else {
		q.QuoteW(w, "age")
		w.WriteString("=NULL")
	}

	w.WriteString(comma)
	if v.Numbers.Bmi != nil {
		q.QuoteW(w, "bmi")
		w.WriteString("=?")
		s = append(s, v.Numbers.Bmi)
		j++
	} else {
		q.QuoteW(w, "bmi")
		w.WriteString("=NULL")
	}

	w.WriteString(comma)
	q.QuoteW(w, "active")
	w.WriteString("=?")
	s = append(s, v.Active)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "labels")
	w.WriteString("=?")
	j++
	x, err := json.Marshal(&v.Labels)
	if err != nil {
		return nil, tbl.database.LogError(errors.WithStack(err))
	}
	s = append(s, x)

	w.WriteString(comma)
	q.QuoteW(w, "fave")
	w.WriteString("=?")
	j++
	x, err := json.Marshal(&v.Fave)
	if err != nil {
		return nil, tbl.database.LogError(errors.WithStack(err))
	}
	s = append(s, x)

	w.WriteString(comma)
	q.QuoteW(w, "avatar")
	w.WriteString("=?")
	s = append(s, v.Avatar)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "foo1")
	w.WriteString("=?")
	s = append(s, v.Foo1)
	j++

	w.WriteString(comma)
	if v.Foo2 != nil {
		q.QuoteW(w, "foo2")
		w.WriteString("=?")
		s = append(s, v.Foo2)
		j++
	} else {
		q.QuoteW(w, "foo2")
		w.WriteString("=NULL")
	}

	w.WriteString(comma)
	if v.Foo3 != nil {
		q.QuoteW(w, "foo3")
		w.WriteString("=?")
		s = append(s, v.Foo3)
		j++
	} else {
		q.QuoteW(w, "foo3")
		w.WriteString("=NULL")
	}

	w.WriteString(comma)
	q.QuoteW(w, "bar1")
	w.WriteString("=?")
	s = append(s, v.Bar1)
	j++

	w.WriteString(comma)
	q.QuoteW(w, "updated")
	w.WriteString("=?")
	j++
	x, err := encoding.MarshalText(&v.Updated)
	if err != nil {
		return nil, tbl.database.LogError(errors.WithStack(err))
	}
	s = append(s, x)

	return s, nil
}
`
	if code != expected {
		outputDiff(expected, "expected.txt")
		outputDiff(code, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
	}
}
