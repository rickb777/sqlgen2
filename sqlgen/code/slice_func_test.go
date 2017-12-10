package code

import (
	"testing"
	"github.com/rickb777/sqlgen2/sqlgen/parse/exit"
	"bytes"
)

func TestWriteSliceFuncWithPk(t *testing.T) {
	exit.TestableExit()

	view := NewView("Example", "X", "")
	buf := &bytes.Buffer{}

	WriteSliceFunc(buf, view, fixtureTable(), false)

	code := buf.String()
	expected := `
func SliceXExample(v *Example) ([]interface{}, error) {
	var err error

	var v0 int64
	var v1 Category
	var v2 string
	var v3 bool
	var v4 []byte
	var v5 []byte
	var v6 []byte

	v0 = v.Id
	v1 = v.Cat
	v2 = v.Name
	v3 = v.Active
	v4, err = json.Marshal(&v.Labels)
	if err != nil {
		return nil, err
	}
	v5, err = json.Marshal(&v.Fave)
	if err != nil {
		return nil, err
	}
	v6, err = encoding.MarshalText(&v.Updated)
	if err != nil {
		return nil, err
	}

	return []interface{}{
		v0,
		v1,
		v2,
		v3,
		v4,
		v5,
		v6,

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
	buf := &bytes.Buffer{}

	WriteSliceFunc(buf, view, fixtureTable(), true)

	code := buf.String()
	expected := `
func SliceXExampleWithoutPk(v *Example) ([]interface{}, error) {
	var err error

	var v1 Category
	var v2 string
	var v3 bool
	var v4 []byte
	var v5 []byte
	var v6 []byte

	v1 = v.Cat
	v2 = v.Name
	v3 = v.Active
	v4, err = json.Marshal(&v.Labels)
	if err != nil {
		return nil, err
	}
	v5, err = json.Marshal(&v.Fave)
	if err != nil {
		return nil, err
	}
	v6, err = encoding.MarshalText(&v.Updated)
	if err != nil {
		return nil, err
	}

	return []interface{}{
		v1,
		v2,
		v3,
		v4,
		v5,
		v6,

	}, nil
}
`
	if code != expected {
		outputDiff(expected, "expected.txt")
		outputDiff(code, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
	}
}
