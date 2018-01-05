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

func fixtureTable() *TableDescription {
	i64 := Type{"", "", "int64", false, Int64}
	boo := Type{"", "", "bool", false, Bool}
	cat := Type{"", "", "Category", false, Int32}
	str := Type{"", "", "string", false, String}
	spt := Type{"", "", "string", true, String}
	ipt := Type{"", "", "int32", true, Int32}
	fpt := Type{"", "", "float32", true, Float32}
	sli := Type{"", "", "[]string", false, Slice}
	bgi := Type{"math/big", "big", "Int", false, Struct}
	bys := Type{"", "", "[]byte", false, Slice}
	tim := Type{"time", "time", "Time", false, Struct}

	id := &Field{Node{"Id", i64, nil}, "id", ENCNONE, Tag{Primary: true, Auto: true}}
	category := &Field{Node{"Cat", cat, nil}, "cat", ENCNONE, Tag{Index: "catIdx"}}
	name := &Field{Node{"Name", str, nil}, "username", ENCNONE, Tag{Size: 2048, Name: "username", Unique: "nameIdx"}}
	active := &Field{Node{"Active", boo, nil}, "active", ENCNONE, Tag{}}
	qual := &Field{Node{"Qual", spt, nil}, "qual", ENCNONE, Tag{}}
	age := &Field{Node{"Age", ipt, nil}, "age", ENCNONE, Tag{}}
	bmi := &Field{Node{"Bmi", fpt, nil}, "bmi", ENCNONE, Tag{}}
	labels := &Field{Node{"Labels", sli, nil}, "labels", ENCJSON, Tag{Encode: "json"}}
	fave := &Field{Node{"Fave", bgi, nil}, "fave", ENCJSON, Tag{Encode: "json"}}
	avatar := &Field{Node{"Avatar", bys, nil}, "avatar", ENCNONE, Tag{}}
	updated := &Field{Node{"Updated", tim, nil}, "updated", ENCTEXT, Tag{Size: 100, Encode: "text"}}

	icat := &Index{"catIdx", false, FieldList{category}}
	iname := &Index{"nameIdx", true, FieldList{name}}

	return &TableDescription{
		Type: "Example",
		Name: "examples",
		Fields: FieldList{
			id,
			category,
			name,
			qual,
			age,
			bmi,
			active,
			labels,
			fave,
			avatar,
			updated,
		},
		Index: []*Index{
			icat,
			iname,
		},
		Primary: id,
	}
}

func TestWriteRowsFunc1(t *testing.T) {
	exit.TestableExit()

	p1 := &Node{Name: "Commit"}
	p2 := &Node{Name: "Author", Parent: p1}

	category := &Field{Node{"Cat", Type{"", "", "Category", false, Int32}, nil}, "cat", ENCNONE, Tag{}}
	name := &Field{Node{"Name", Type{"", "", "string", false, String}, p2}, "commit_author_name", ENCNONE, Tag{}}
	email := &Field{Node{"Email", Type{"", "", "string", false, String}, p2}, "commit_author_email", ENCNONE, Tag{}}
	message := &Field{Node{"Message", Type{"", "", "string", false, String}, p1}, "commit_message", ENCNONE, Tag{}}

	table := &TableDescription{
		Type: "Example",
		Name: "examples",
		Fields: FieldList{
			category,
			name,
			email,
			message,
		},
	}

	view := NewView("Example", "X", "")
	view.Table = table

	buf := &bytes.Buffer{}

	WriteRowsFunc(buf, view)

	code := buf.String()
	expected := `
// scanXExamples reads table records into a slice of values.
func scanXExamples(rows *sql.Rows, firstOnly bool) ([]*Example, error) {
	var err error
	var vv []*Example

	var v0 Category
	var v1 string
	var v2 string
	var v3 string

	for rows.Next() {
		err = rows.Scan(
			&v0,
			&v1,
			&v2,
			&v3,
		)
		if err != nil {
			return vv, err
		}

		v := &Example{}
		v.Cat = v0
		v.Commit.Author.Name = v1
		v.Commit.Author.Email = v2
		v.Commit.Message = v3

		vv = append(vv, v)

		if firstOnly {
			return vv, rows.Err()
		}
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

func TestWriteRowFunc2(t *testing.T) {
	exit.TestableExit()

	view := NewView("Example", "X", "")
	view.Table = fixtureTable()
	buf := &bytes.Buffer{}

	WriteRowsFunc(buf, view)

	code := buf.String()
	expected := `
// scanXExamples reads table records into a slice of values.
func scanXExamples(rows *sql.Rows, firstOnly bool) ([]*Example, error) {
	var err error
	var vv []*Example

	var v0 int64
	var v1 Category
	var v2 string
	var v3 sql.NullString
	var v4 sql.NullInt64
	var v5 sql.NullFloat64
	var v6 bool
	var v7 []byte
	var v8 []byte
	var v9 []byte
	var v10 []byte

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

		v := &Example{}
		v.Id = v0
		v.Cat = v1
		v.Name = v2
		if v3.Valid {
			v.Qual = &(string(v3.String))
		}
		if v4.Valid {
			v.Age = &(int32(v4.Int64))
		}
		if v5.Valid {
			v.Bmi = &(float32(v5.Float64))
		}
		v.Active = v6
		err = json.Unmarshal(v7, &v.Labels)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(v8, &v.Fave)
		if err != nil {
			return nil, err
		}
		v.Avatar = v9
		err = encoding.UnmarshalText(v10, &v.Updated)
		if err != nil {
			return nil, err
		}

		vv = append(vv, v)

		if firstOnly {
			return vv, rows.Err()
		}
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
		eax := strings.Replace(ea, "    ", "˙˙˙˙", -1)
		eax = strings.Replace(eax, "\t", "————", -1)
		buf.WriteString(fmt.Sprintf("%-60s", truncate(eax, 60)))
		if i < len(bb) {
			ebx := strings.Replace(bb[i], "    ", "˙˙˙˙", -1)
			ebx = strings.Replace(ebx, "\t", "————", -1)
			if ea != bb[i] {
				buf.WriteString(" <> ")
			} else {
				buf.WriteString("    ")
			}
			buf.WriteString(truncate(ebx, 60))
		}
		buf.WriteByte('\n')
		i++
	}

	for ; i < len(bb); i++ {
		buf.WriteString(fmt.Sprintf("%60s    ", ""))
		ebx := strings.Replace(bb[i], "    ", "˙˙˙˙", -1)
		ebx = strings.Replace(ebx, "\t", "————", -1)
		buf.WriteString(truncate(ebx, 60))
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
