package code

import (
	"bytes"
	"fmt"
	. "github.com/rickb777/sqlapi/schema"
	. "github.com/rickb777/sqlapi/types"
	"github.com/rickb777/sqlgen2/parse/exit"
	"os"
	"strings"
	"testing"
)

func TestWriteRowsFunc1(t *testing.T) {
	exit.TestableExit()

	p1 := &Node{Name: "Commit"}
	p2 := &Node{Name: "Author", Parent: p1}

	category2 := &Field{Node{"Cat", Type{Name: "Category", Base: Int32}, nil}, "cat", ENCNONE, nil}
	name2 := &Field{Node{"Name", Type{Name: "string", Base: String}, p2}, "commit_author_name", ENCNONE, nil}
	email2 := &Field{Node{"Email", Type{Name: "string", Base: String}, p2}, "commit_author_email", ENCNONE, nil}
	message2 := &Field{Node{"Message", Type{Name: "string", Base: String}, p1}, "commit_message", ENCNONE, nil}

	table := &TableDescription{
		Type: "Example",
		Name: "examples",
		Fields: FieldList{
			category2,
			name2,
			email2,
			message2,
		},
	}

	view := NewView("", "", "Example", "X", "", "", "sql", "sqlapi")
	view.Table = table

	buf := &bytes.Buffer{}

	WriteScanRows(buf, view)

	code := buf.String()
	expected := `
// scanXExamples reads rows from the database and returns a slice of corresponding values.
// It also returns a number indicating how many rows were read; this will be larger than the length of the
// slice if reading stopped after the first row.
func scanXExamples(query string, rows sqlapi.SqlRows, firstOnly bool) (vv []*Example, n int64, err error) {
	for rows.Next() {
		n++

		var v0 Category
		var v1 string
		var v2 string
		var v3 string

		err = rows.Scan(
			&v0,
			&v1,
			&v2,
			&v3,
		)
		if err != nil {
			return vv, n, errors.Wrap(err, query)
		}

		v := &Example{}
		v.Cat = v0
		v.Commit.Author.Name = v1
		v.Commit.Author.Email = v2
		v.Commit.Message = v3

		var iv interface{} = v
		if hook, ok := iv.(sqlapi.CanPostGet); ok {
			err = hook.PostGet()
			if err != nil {
				return vv, n, errors.Wrap(err, query)
			}
		}

		vv = append(vv, v)

		if firstOnly {
			if rows.Next() {
				n++
			}
			return vv, n, errors.Wrap(rows.Err(), query)
		}
	}

	return vv, n, errors.Wrap(rows.Err(), query)
}
`
	expectCodeEqual(t, code, expected)
	disallowTrailingWhitespace(t, code)
}

func TestWriteRowFunc2(t *testing.T) {
	exit.TestableExit()

	view := NewView("", "", "Example", "X", "", "", "sql", "sqlapi")
	view.Scan = "Scan"
	view.Table = fixtureTable()
	buf := &bytes.Buffer{}

	WriteScanRows(buf, view)

	code := buf.String()
	expected := `
// ScanXExamples reads rows from the database and returns a slice of corresponding values.
// It also returns a number indicating how many rows were read; this will be larger than the length of the
// slice if reading stopped after the first row.
func ScanXExamples(query string, rows sqlapi.SqlRows, firstOnly bool) (vv []*Example, n int64, err error) {
	for rows.Next() {
		n++

		var v0 int64
		var v1 Category
		var v2 string
		var v3 sql.NullString
		var v4 sql.NullString
		var v5 sql.NullInt64
		var v6 sql.NullInt64
		var v7 sql.NullFloat64
		var v8 bool
		var v9 []byte
		var v10 []byte
		var v11 []byte
		var v12 Foo
		var v13 sql.NullString
		var v14 sql.NullInt64
		var v15 Bar
		var v16 []byte

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
			&v11,
			&v12,
			&v13,
			&v14,
			&v15,
			&v16,
		)
		if err != nil {
			return vv, n, errors.Wrap(err, query)
		}

		v := &Example{}
		v.Id = v0
		v.Cat = v1
		v.Name = v2
		if v3.Valid {
			a := PhoneNumber(v3.String)
			v.Mobile = &a
		}
		if v4.Valid {
			a := v4.String
			v.Qual = &a
		}
		if v5.Valid {
			a := int32(v5.Int64)
			v.Numbers.Diff = &a
		}
		if v6.Valid {
			a := uint32(v6.Int64)
			v.Numbers.Age = &a
		}
		if v7.Valid {
			a := float32(v7.Float64)
			v.Numbers.Bmi = &a
		}
		v.Active = v8
		err = json.Unmarshal(v9, &v.Labels)
		if err != nil {
			return nil, n, errors.Wrap(err, query)
		}
		err = json.Unmarshal(v10, &v.Fave)
		if err != nil {
			return nil, n, errors.Wrap(err, query)
		}
		v.Avatar = v11
		v.Foo1 = v12
		if v13.Valid {
			v.Foo2 = new(Foo)
			err = v.Foo2.Scan(v13.String)
			if err != nil {
				return nil, n, errors.Wrap(err, query)
			}
		}
		if v14.Valid {
			v.Foo3 = new(Foo)
			err = v.Foo3.Scan(v14.Int64)
			if err != nil {
				return nil, n, errors.Wrap(err, query)
			}
		}
		v.Bar1 = v15
		err = encoding.UnmarshalText(v16, &v.Updated)
		if err != nil {
			return nil, n, errors.Wrap(err, query)
		}

		var iv interface{} = v
		if hook, ok := iv.(sqlapi.CanPostGet); ok {
			err = hook.PostGet()
			if err != nil {
				return vv, n, errors.Wrap(err, query)
			}
		}

		vv = append(vv, v)

		if firstOnly {
			if rows.Next() {
				n++
			}
			return vv, n, errors.Wrap(rows.Err(), query)
		}
	}

	return vv, n, errors.Wrap(rows.Err(), query)
}
`
	expectCodeEqual(t, code, expected)
	disallowTrailingWhitespace(t, code)
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
