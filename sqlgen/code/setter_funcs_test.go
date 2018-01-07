package code

import (
	"testing"
	"github.com/rickb777/sqlgen2/sqlgen/parse/exit"
	"bytes"
)

func TestWriteSetters_all(t *testing.T) {
	exit.TestableExit()

	view := NewView("Example", "X", "")
	view.Table = fixtureTable()
	buf := &bytes.Buffer{}

	WriteSetters(buf, view, view.FilterSetters("all"))

	code := buf.String()
	expected := `
//--------------------------------------------------------------------------------

// SetId sets the Id field and returns the modified Example.
func (v *Example) SetId(x int64) *Example {
	v.Id = x
	return v
}

// SetCat sets the Cat field and returns the modified Example.
func (v *Example) SetCat(x Category) *Example {
	v.Cat = x
	return v
}

// SetName sets the Name field and returns the modified Example.
func (v *Example) SetName(x string) *Example {
	v.Name = x
	return v
}

// SetQual sets the Qual field and returns the modified Example.
func (v *Example) SetQual(x string) *Example {
	v.Qual = &x
	return v
}

// SetAge sets the Age field and returns the modified Example.
func (v *Example) SetAge(x int32) *Example {
	v.Age = &x
	return v
}

// SetBmi sets the Bmi field and returns the modified Example.
func (v *Example) SetBmi(x float32) *Example {
	v.Bmi = &x
	return v
}

// SetActive sets the Active field and returns the modified Example.
func (v *Example) SetActive(x bool) *Example {
	v.Active = x
	return v
}

// SetLabels sets the Labels field and returns the modified Example.
func (v *Example) SetLabels(x []string) *Example {
	v.Labels = x
	return v
}

// SetFave sets the Fave field and returns the modified Example.
func (v *Example) SetFave(x big.Int) *Example {
	v.Fave = x
	return v
}

// SetAvatar sets the Avatar field and returns the modified Example.
func (v *Example) SetAvatar(x []byte) *Example {
	v.Avatar = x
	return v
}

// SetUpdated sets the Updated field and returns the modified Example.
func (v *Example) SetUpdated(x time.Time) *Example {
	v.Updated = x
	return v
}
`
	if code != expected {
		outputDiff(expected, "expected.txt")
		outputDiff(code, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
	}
}

func TestWriteSetters_exported(t *testing.T) {
	exit.TestableExit()

	view := NewView("Example", "X", "")
	view.Table = fixtureTable()
	buf := &bytes.Buffer{}

	WriteSetters(buf, view, view.FilterSetters("exported"))

	code := buf.String()
	expected := `
//--------------------------------------------------------------------------------

// SetId sets the Id field and returns the modified Example.
func (v *Example) SetId(x int64) *Example {
	v.Id = x
	return v
}

// SetCat sets the Cat field and returns the modified Example.
func (v *Example) SetCat(x Category) *Example {
	v.Cat = x
	return v
}

// SetName sets the Name field and returns the modified Example.
func (v *Example) SetName(x string) *Example {
	v.Name = x
	return v
}

// SetQual sets the Qual field and returns the modified Example.
func (v *Example) SetQual(x string) *Example {
	v.Qual = &x
	return v
}

// SetAge sets the Age field and returns the modified Example.
func (v *Example) SetAge(x int32) *Example {
	v.Age = &x
	return v
}

// SetBmi sets the Bmi field and returns the modified Example.
func (v *Example) SetBmi(x float32) *Example {
	v.Bmi = &x
	return v
}

// SetActive sets the Active field and returns the modified Example.
func (v *Example) SetActive(x bool) *Example {
	v.Active = x
	return v
}

// SetLabels sets the Labels field and returns the modified Example.
func (v *Example) SetLabels(x []string) *Example {
	v.Labels = x
	return v
}

// SetFave sets the Fave field and returns the modified Example.
func (v *Example) SetFave(x big.Int) *Example {
	v.Fave = x
	return v
}

// SetAvatar sets the Avatar field and returns the modified Example.
func (v *Example) SetAvatar(x []byte) *Example {
	v.Avatar = x
	return v
}

// SetUpdated sets the Updated field and returns the modified Example.
func (v *Example) SetUpdated(x time.Time) *Example {
	v.Updated = x
	return v
}
`
	if code != expected {
		outputDiff(expected, "expected.txt")
		outputDiff(code, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
	}
}

func TestWriteSetters_optional(t *testing.T) {
	exit.TestableExit()

	view := NewView("Example", "X", "")
	view.Table = fixtureTable()
	buf := &bytes.Buffer{}

	WriteSetters(buf, view, view.FilterSetters("optional"))

	code := buf.String()
	expected := `
//--------------------------------------------------------------------------------

// SetQual sets the Qual field and returns the modified Example.
func (v *Example) SetQual(x string) *Example {
	v.Qual = &x
	return v
}

// SetAge sets the Age field and returns the modified Example.
func (v *Example) SetAge(x int32) *Example {
	v.Age = &x
	return v
}

// SetBmi sets the Bmi field and returns the modified Example.
func (v *Example) SetBmi(x float32) *Example {
	v.Bmi = &x
	return v
}
`
	if code != expected {
		outputDiff(expected, "expected.txt")
		outputDiff(code, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
	}
}

func TestWriteSetters_none(t *testing.T) {
	exit.TestableExit()

	view := NewView("Example", "X", "")
	view.Table = fixtureTable()
	buf := &bytes.Buffer{}

	WriteSetters(buf, view, view.FilterSetters("none"))

	code := buf.String()
	expected := ""
	if code != expected {
		outputDiff(expected, "expected.txt")
		outputDiff(code, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
	}
}
