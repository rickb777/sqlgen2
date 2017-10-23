package code

import (
	"bytes"
	"testing"
	"github.com/rickb777/sqlgen2/sqlgen/parse/exit"
	"github.com/rickb777/sqlgen2/sqlgen/parse"
)

func TestWriteRowFunc_withoutPK(t *testing.T) {
	exit.TestableExit()

	source := parse.Source{"issue.go", bytes.NewBufferString(literal)}

	tree, err := parse.DoParse("pkg1", "Party", []parse.Source{source})
	if err != nil {
		t.Errorf("Error parsing: %s", err)
	}

	//table := schema.Load(node)
	view := NewView(tree, "X")
	//view.Table = table

	buf := &bytes.Buffer{}

	WriteSliceFunc(buf, tree, view, false)

	code := buf.String()
	expected := `
func SliceParty(v *Party) []interface{} {
	var v0 int64
	var v1 int
	var v2 pkg1.Category
	var v4 string
	var v5 int64
	var v6 string
	var v7 string
	var v8 string
	var v9 string
	var v10 []byte
	var v11 bool

	v0 = v.Id
	v1 = v.Number
	v2 = v.Category
	if v.Commit != nil {
		v4 = v.Commit.Message
		v5 = v.Commit.Timestamp
		if v.Commit.Author != nil {
			v6 = v.Commit.Author.Name
			v7 = v.Commit.Author.Email
			v8 = v.Title
			v9 = v.Hobby
			v10, _ = json.Marshal(&v.Labels)
			v11 = v.Active
		}
	}

	return []interface{}{
		v0,
		v1,
		v2,
		v4,
		v5,
		v6,
		v7,
		v8,
		v9,
		v10,
		v11,

	}
}
`
	if code != expected {
		outputDiff(expected, "expected.txt")
		outputDiff(code, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
	}
}
