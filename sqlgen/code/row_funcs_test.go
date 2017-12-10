package code

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
	"github.com/rickb777/sqlgen2/sqlgen/parse/exit"
	. "github.com/rickb777/sqlgen2/schema"
	. "github.com/rickb777/sqlgen2/sqlgen/parse"
)

//TODO unfinished work based on this test
func xTestWriteRowFunc1(t *testing.T) {
	exit.TestableExit()

	p1 := &Node{Name: "Commit"}
	p2 := &Node{Name: "Author", Parent: p1}

	category := &Field{Node{"Cat", Type{"", "", "Category", Int32}, nil}, "cat", INTEGER, ENCNONE, Tag{}}
	name := &Field{Node{"Name", Type{"", "", "string", String}, p2}, "commit_author_name", VARCHAR, ENCNONE, Tag{}}
	email := &Field{Node{"Email", Type{"", "", "string", String}, p2}, "commit_author_email", VARCHAR, ENCNONE, Tag{}}
	message := &Field{Node{"Message", Type{"", "", "string", String}, p1}, "commit_message", VARCHAR, ENCNONE, Tag{}}

	table := &TableDescription{
		Type: "Example",
		Name: "examples",
		Fields: []*Field{
			category,
			name,
			email,
			message,
		},
	}

	//table := schema.Load(node)
	view := NewView("Example", "X", "")
	//view.Table = table

	buf := &bytes.Buffer{}

	WriteRowFunc(buf, view, table)

	code := buf.String()
	expected := `
// ScanXExample reads a table record into a single value.
func ScanXExample(row *sql.Row) (*Example, error) {
	var v0 int64
	var v1 Category
	var v2 string
	var v3 string
	var v4 string

	err := row.Scan(
		&v0,
		&v1,
		&v2,
		&v3,
		&v4,

	)
	if err != nil {
		return nil, err
	}

	v := &Example{}
	v.Cat = v0
	v.Id = v1
	v.Commit = &Commit{}
	v.Commit.Author = &Author{}
	v.Commit.Author.Name = v2
	v.Commit.Author.Email = v3
	v.Commit.Author.Message = v4

	return v, nil
}
`
	if code != expected {
		outputDiff(expected, "expected.txt")
		outputDiff(code, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
	}
}

func fixtureTable() *TableDescription {
	id := &Field{Node{"Id", Type{"", "", "int64", Int64}, nil}, "id", INTEGER, ENCNONE, Tag{Primary: true, Auto: true}}
	category := &Field{Node{"Cat", Type{"", "", "Category", Int32}, nil}, "cat", INTEGER, ENCNONE, Tag{}}
	name := &Field{Node{"Name", Type{"", "", "string", String}, nil}, "name", VARCHAR, ENCNONE, Tag{Size: 2048}}
	active := &Field{Node{"Active", Type{"", "", "bool", Bool}, nil}, "active", BOOLEAN, ENCNONE, Tag{}}
	labels := &Field{Node{"Labels", Type{"", "", "[]string", Slice}, nil}, "labels", BLOB, ENCJSON, Tag{Encode: "json"}}
	fave := &Field{Node{"Fave", Type{"math/big", "big", "Int", Struct}, nil}, "fave", BLOB, ENCJSON, Tag{Encode: "json"}}
	updated := &Field{Node{"Updated", Type{"time", "time", "Time", Struct}, nil}, "updated", BLOB, ENCTEXT, Tag{Encode: "text"}}

	return &TableDescription{
		Type: "Example",
		Name: "examples",
		Fields: []*Field{
			id,
			category,
			name,
			active,
			labels,
			fave,
			updated,
		},
	}
}

func TestWriteRowFunc2(t *testing.T) {
	exit.TestableExit()

	view := NewView("Example", "X", "")
	buf := &bytes.Buffer{}

	WriteRowFunc(buf, view, fixtureTable())

	code := buf.String()
	expected := `
// ScanXExample reads a table record into a single value.
func ScanXExample(row *sql.Row) (*Example, error) {
	var v0 int64
	var v1 Category
	var v2 string
	var v3 bool
	var v4 []byte
	var v5 []byte
	var v6 []byte

	err := row.Scan(
		&v0,
		&v1,
		&v2,
		&v3,
		&v4,
		&v5,
		&v6,

	)
	if err != nil {
		return nil, err
	}

	v := &Example{}
	v.Id = v0
	v.Cat = v1
	v.Name = v2
	v.Active = v3
	err = json.Unmarshal(v4, &v.Labels)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(v5, &v.Fave)
	if err != nil {
		return nil, err
	}
	err = encoding.UnmarshalText(v6, &v.Updated)
	if err != nil {
		return nil, err
	}

	return v, nil
}
`
	if code != expected {
		outputDiff(expected, "expected.txt")
		outputDiff(code, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
	}
}

func TestWriteRowsFunc2(t *testing.T) {
	exit.TestableExit()

	view := NewView("Example", "X", "")
	buf := &bytes.Buffer{}

	WriteRowsFunc(buf, view, fixtureTable())

	code := buf.String()
	expected := `
// ScanXExamples reads table records into a slice of values.
func ScanXExamples(rows *sql.Rows) ([]*Example, error) {
	var err error
	var vv []*Example

	var v0 int64
	var v1 Category
	var v2 string
	var v3 bool
	var v4 []byte
	var v5 []byte
	var v6 []byte

	for rows.Next() {
		err = rows.Scan(
			&v0,
			&v1,
			&v2,
			&v3,
			&v4,
			&v5,
			&v6,

		)
		if err != nil {
			return vv, err
		}

		v := &Example{}
		v.Id = v0
		v.Cat = v1
		v.Name = v2
		v.Active = v3
		err = json.Unmarshal(v4, &v.Labels)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(v5, &v.Fave)
		if err != nil {
			return nil, err
		}
		err = encoding.UnmarshalText(v6, &v.Updated)
		if err != nil {
			return nil, err
		}

		vv = append(vv, v)
	}
	return vv, rows.Err()
}
`
	if code != expected {
		outputDiff(expected, "expected.txt")
		outputDiff(code, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
	}
}


func outputDiff(a, name string) {
	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	f.WriteString(a)
	f.WriteString("\n")
	err = f.Close()
	if err != nil {
		panic(err)
	}
}

func sideBySideDiff(a, b string) string {
	aa := strings.Split(a, "\n")
	bb := strings.Split(b, "\n")

	buf := &bytes.Buffer{}
	i := 0

	for _, ea := range aa {
		ea := strings.Replace(ea, "\t", "    ", -1)
		buf.WriteString(fmt.Sprintf("%-60s", truncate(ea, 60)))
		if i < len(bb) {
			eb := strings.Replace(bb[i], "\t", "    ", -1)
			if ea != eb {
				buf.WriteString(" <> ")
			} else {
				buf.WriteString("    ")
			}
			buf.WriteString(truncate(eb, 60))
		}
		buf.WriteByte('\n')
		i++
	}

	for ; i < len(bb); i++ {
		buf.WriteString(fmt.Sprintf("%60s    ", ""))
		eb := strings.Replace(bb[i], "\t", "    ", -1)
		buf.WriteString(truncate(eb, 60))
		buf.WriteByte('\n')
	}

	return buf.String()
}

func truncate(s string, n int) string {
	if len(s) > n {
		return s[:n-3] + "..."
	}
	return s
}
