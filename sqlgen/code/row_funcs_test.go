package code

import (
	"bytes"
	"testing"
	"github.com/rickb777/sqlgen2/sqlgen/parse/exit"
	"strings"
	"github.com/rickb777/sqlgen2/sqlgen/parse"
	"fmt"
	"os"
)

var literal = strings.Replace(`package pkg1

type Party struct {
	Id         int64    |sql:"pk: true, auto: true"|
	Number     int
	Category   Category
	Foo        int      |sql:"-"|
	Commit     *Commit
	Title      string   |sql:"index: titleIdx"|
	Hobby      string   |sql:"size: 2048"|
	Labels     []string |sql:"encode: json"|
	Active     bool
}

type Category int32

type Commit struct {
	Message   string
	Timestamp int64 // TODO should be able to support time.Time
	Author    *Author
}

type Author struct {
	Name     string
	Email    string
}

`, "|", "`", -1)

func TestWriteRowFunc(t *testing.T) {
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

	WriteRowFunc(buf, tree, view)

	code := buf.String()
	expected := `
// ScanParty reads a database record into a single value.
func ScanParty(row *sql.Row) (*Party, error) {
	var v0 int64
	var v1 int
	var v2 pkg1.Category
	var v3 string
	var v4 int64
	var v5 string
	var v6 string
	var v7 string
	var v8 string
	var v9 []byte
	var v10 bool

	err := row.Scan(
		&v0,
		&v1,
		&v2,
		&v3,
		&v4,
		&v5,
		&v6,
		&v7,
		&v8,
		&v9,
		&v10,

	)
	if err != nil {
		return nil, err
	}

	v := &Party{}
	v.Id = v0
	v.Number = v1
	v.Category = v2
	v.Commit = &Commit{}
	v.Commit.Message = v3
	v.Commit.Timestamp = v4
	v.Commit.Author = &Author{}
	v.Commit.Author.Name = v5
	v.Commit.Author.Email = v6
	v.Title = v7
	v.Hobby = v8
	json.Unmarshal(v9, &v.Labels)
	v.Active = v10

	return v, nil
}
`
	if code != expected {
		outputDiff(expected, "expected.txt")
		outputDiff(code, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
	}
}

//TODO plural records names are not handled correctly - fix this
func TestWriteRowsFunc(t *testing.T) {
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

	WriteRowsFunc(buf, tree, view)

	code := buf.String()
	expected := `
// ScanPartys reads database records into a slice of values.
func ScanPartys(rows *sql.Rows) ([]*Party, error) {
	var err error
	var vv []*Party

	var v0 int64
	var v1 int
	var v2 pkg1.Category
	var v3 string
	var v4 int64
	var v5 string
	var v6 string
	var v7 string
	var v8 string
	var v9 []byte
	var v10 bool

	for rows.Next() {
		err = rows.Scan(
			&v0,
			&v1,
			&v2,
			&v3,
			&v4,
			&v5,
			&v6,
			&v7,
			&v8,
			&v9,
			&v10,

		)
		if err != nil {
			return vv, err
		}

		v := &Party{}
		v.Id = v0
		v.Number = v1
		v.Category = v2
		v.Commit = &Commit{}
		v.Commit.Message = v3
		v.Commit.Timestamp = v4
		v.Commit.Author = &Author{}
		v.Commit.Author.Name = v5
		v.Commit.Author.Email = v6
		v.Title = v7
		v.Hobby = v8
		json.Unmarshal(v9, &v.Labels)
		v.Active = v10

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
