package code

import (
	"testing"
	"github.com/rickb777/sqlgen2/sqlgen/parse/exit"
	"bytes"
)

func TestWriteConstructInsert(t *testing.T) {
	exit.TestableExit()

	view := NewView("Example", "X", "", "")
	view.Table = fixtureTable()
	buf := &bytes.Buffer{}

	WriteConstructInsert(buf, view)

	code := buf.String()
	expected := `
func constructXExampleInsert(w io.Writer, v *Example, dialect schema.Dialect, withPk bool) (s []interface{}, err error) {
	s = make([]interface{}, 0, 17)

	comma := ""
	io.WriteString(w, " (")

	if withPk {
		dialect.QuoteW(w, "id")
		comma = ","
		s = append(s, v.Id)
	}

	io.WriteString(w, comma)

	dialect.QuoteW(w, "cat")
	s = append(s, v.Cat)
	comma = ","
	io.WriteString(w, comma)

	dialect.QuoteW(w, "username")
	s = append(s, v.Name)
	if v.Mobile != nil {
		io.WriteString(w, comma)

		dialect.QuoteW(w, "mobile")
		s = append(s, v.Mobile)
	}
	if v.Qual != nil {
		io.WriteString(w, comma)

		dialect.QuoteW(w, "qual")
		s = append(s, v.Qual)
	}
	if v.Numbers.Diff != nil {
		io.WriteString(w, comma)

		dialect.QuoteW(w, "diff")
		s = append(s, v.Numbers.Diff)
	}
	if v.Numbers.Age != nil {
		io.WriteString(w, comma)

		dialect.QuoteW(w, "age")
		s = append(s, v.Numbers.Age)
	}
	if v.Numbers.Bmi != nil {
		io.WriteString(w, comma)

		dialect.QuoteW(w, "bmi")
		s = append(s, v.Numbers.Bmi)
	}
	io.WriteString(w, comma)

	dialect.QuoteW(w, "active")
	s = append(s, v.Active)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "labels")
	x, err := json.Marshal(&v.Labels)
	if err != nil {
		return nil, err
	}
	s = append(s, x)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "fave")
	x, err := json.Marshal(&v.Fave)
	if err != nil {
		return nil, err
	}
	s = append(s, x)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "avatar")
	s = append(s, v.Avatar)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "foo1")
	s = append(s, v.Foo1)
	if v.Foo2 != nil {
		io.WriteString(w, comma)

		dialect.QuoteW(w, "foo2")
		s = append(s, v.Foo2)
	}
	if v.Foo3 != nil {
		io.WriteString(w, comma)

		dialect.QuoteW(w, "foo3")
		s = append(s, v.Foo3)
	}
	io.WriteString(w, comma)

	dialect.QuoteW(w, "bar1")
	s = append(s, v.Bar1)
	io.WriteString(w, comma)

	dialect.QuoteW(w, "updated")
	x, err := encoding.MarshalText(&v.Updated)
	if err != nil {
		return nil, err
	}
	s = append(s, x)
	io.WriteString(w, ")")
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
func constructXExampleUpdate(w io.Writer, v *Example, dialect schema.Dialect) (s []interface{}, err error) {
	j := 1
	s = make([]interface{}, 0, 16)

	comma := ""

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "cat", j)
	s = append(s, v.Cat)
	comma = ", "
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "username", j)
	s = append(s, v.Name)
		j++

	io.WriteString(w, comma)
	if v.Mobile != nil {
		dialect.QuoteWithPlaceholder(w, "mobile", j)
		s = append(s, v.Mobile)
		j++
	} else {
		dialect.QuoteW(w, "mobile")
		io.WriteString(w, "=NULL")
	}

	io.WriteString(w, comma)
	if v.Qual != nil {
		dialect.QuoteWithPlaceholder(w, "qual", j)
		s = append(s, v.Qual)
		j++
	} else {
		dialect.QuoteW(w, "qual")
		io.WriteString(w, "=NULL")
	}

	io.WriteString(w, comma)
	if v.Numbers.Diff != nil {
		dialect.QuoteWithPlaceholder(w, "diff", j)
		s = append(s, v.Numbers.Diff)
		j++
	} else {
		dialect.QuoteW(w, "diff")
		io.WriteString(w, "=NULL")
	}

	io.WriteString(w, comma)
	if v.Numbers.Age != nil {
		dialect.QuoteWithPlaceholder(w, "age", j)
		s = append(s, v.Numbers.Age)
		j++
	} else {
		dialect.QuoteW(w, "age")
		io.WriteString(w, "=NULL")
	}

	io.WriteString(w, comma)
	if v.Numbers.Bmi != nil {
		dialect.QuoteWithPlaceholder(w, "bmi", j)
		s = append(s, v.Numbers.Bmi)
		j++
	} else {
		dialect.QuoteW(w, "bmi")
		io.WriteString(w, "=NULL")
	}

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "active", j)
	s = append(s, v.Active)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "labels", j)
		j++
	x, err := json.Marshal(&v.Labels)
	if err != nil {
		return nil, err
	}
	s = append(s, x)

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "fave", j)
		j++
	x, err := json.Marshal(&v.Fave)
	if err != nil {
		return nil, err
	}
	s = append(s, x)

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "avatar", j)
	s = append(s, v.Avatar)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "foo1", j)
	s = append(s, v.Foo1)
		j++

	io.WriteString(w, comma)
	if v.Foo2 != nil {
		dialect.QuoteWithPlaceholder(w, "foo2", j)
		s = append(s, v.Foo2)
		j++
	} else {
		dialect.QuoteW(w, "foo2")
		io.WriteString(w, "=NULL")
	}

	io.WriteString(w, comma)
	if v.Foo3 != nil {
		dialect.QuoteWithPlaceholder(w, "foo3", j)
		s = append(s, v.Foo3)
		j++
	} else {
		dialect.QuoteW(w, "foo3")
		io.WriteString(w, "=NULL")
	}

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "bar1", j)
	s = append(s, v.Bar1)
		j++

	io.WriteString(w, comma)
	dialect.QuoteWithPlaceholder(w, "updated", j)
		j++
	x, err := encoding.MarshalText(&v.Updated)
	if err != nil {
		return nil, err
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
