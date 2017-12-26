package code

import (
	"testing"
	"github.com/rickb777/sqlgen2/sqlgen/parse/exit"
	"bytes"
)

func TestWriteSliceFuncWithPk(t *testing.T) {
	exit.TestableExit()

	view := NewView("Example", "X", "")
	view.Table = fixtureTable()
	buf := &bytes.Buffer{}

	WriteSliceFunc(buf, view, false)

	code := buf.String()
	expected := `
func sliceXExample(v *Example) ([]interface{}, error) {

	v4, err := json.Marshal(&v.Labels)
	if err != nil {
		return nil, err
	}
	v5, err := json.Marshal(&v.Fave)
	if err != nil {
		return nil, err
	}
	v7, err := encoding.MarshalText(&v.Updated)
	if err != nil {
		return nil, err
	}

	return []interface{}{
		v.Id,
		v.Cat,
		v.Name,
		v.Active,
		v4,
		v5,
		v.Avatar,
		v7,

	}, nil
}
`
	if code != expected {
		outputDiff(expected, "expected.txt")
		outputDiff(code, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
	}
}

func TestWriteSliceFuncWithoutPk(t *testing.T) {
	exit.TestableExit()

	view := NewView("Example", "X", "")
	view.Table = fixtureTable()
	buf := &bytes.Buffer{}

	WriteSliceFunc(buf, view, true)

	code := buf.String()
	expected := `
func sliceXExampleWithoutPk(v *Example) ([]interface{}, error) {

	v4, err := json.Marshal(&v.Labels)
	if err != nil {
		return nil, err
	}
	v5, err := json.Marshal(&v.Fave)
	if err != nil {
		return nil, err
	}
	v7, err := encoding.MarshalText(&v.Updated)
	if err != nil {
		return nil, err
	}

	return []interface{}{
		v.Cat,
		v.Name,
		v.Active,
		v4,
		v5,
		v.Avatar,
		v7,

	}, nil
}
`
	if code != expected {
		outputDiff(expected, "expected.txt")
		outputDiff(code, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
	}
}
