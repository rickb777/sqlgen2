package main

import (
	"bytes"
	"fmt"
	"github.com/kortschak/utter"
	. "github.com/rickb777/sqlgen2/schema"
	. "github.com/rickb777/sqlgen2/sqlgen/parse"
	"github.com/rickb777/sqlgen2/sqlgen/parse/exit"
	"go/token"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestParseAndLoad_typesWithAllFieldsUnexported(t *testing.T) {
	exit.TestableExit()
	Debug = true

	template := `package pkg1

type Example struct {
	Event    Date    %s
}

type Date struct {
	day int32
}
`

	cases := []struct {
		tag  string
		date *Field
	}{
		{
			"",
			&Field{Node{"Event", Type{"", "", "Date", Struct}, nil}, "event", ENCNONE, Tag{}},
		},
		{
			"`sql:\"encode: json\"`",
			&Field{Node{"Event", Type{"", "", "Date", Struct}, nil}, "event", ENCJSON, Tag{Encode: "json"}},
		},
		{
			"`sql:\"type: integer\"`",
			&Field{Node{"Event", Type{"", "", "Date", Int}, nil}, "event", ENCNONE, Tag{Type: "integer"}},
		},
	}

	for i, c := range cases {
		code := fmt.Sprintf(template, c.tag)
		source := Source{"issue.go", bytes.NewBufferString(code)}

		pkgStore, err := ParseGroups(token.NewFileSet(), Group{"pkg1", []Source{source}})
		if err != nil {
			t.Fatalf("Error parsing: %s", err)
		}

		table, err := load(pkgStore, "pkg1", "Example")
		if err != nil {
			t.Fatalf("Error loading: %s", err)
		}

		expected := &TableDescription{
			Type: "Example",
			Name: "examples",
			Fields: []*Field{
				c.date,
			},
		}

		if !reflect.DeepEqual(table, expected) {
			ex := utter.Sdump(expected)
			ac := utter.Sdump(table)
			outputDiff(ex, "expected.txt")
			outputDiff(ac, "got.txt")
			t.Errorf("%d: expected | got\n%s\n", i, sideBySideDiff(ex, ac))
		}
	}
}

func TestParseAndLoad_nestingWithPointers(t *testing.T) {
	exit.TestableExit()
	Debug = true
	code := strings.Replace(`package pkg1

type Example struct {
	Commit   *Commit
}

type Commit struct {
	Author   *Author
}

type Author struct {
	Name     string
}

`, "|", "`", -1)

	source := Source{"issue.go", bytes.NewBufferString(code)}

	pkgStore, err := ParseGroups(token.NewFileSet(), Group{"pkg1", []Source{source}})
	if err != nil {
		t.Fatalf("Error parsing: %s", err)
	}

	table, err := load(pkgStore, "pkg1", "Example")
	if err != nil {
		t.Fatalf("Error loading: %s", err)
	}

	p1 := &Node{Name: "Commit", Type: Type{Name: "Commit", Base: Struct}}
	p2 := &Node{Name: "Author", Type: Type{Name: "Author", Base: Struct}, Parent: p1}
	author := &Field{Node{"Name", Type{Name: "string", Base: String}, p2}, "name", ENCNONE, Tag{}}

	expected := &TableDescription{
		Type: "Example",
		Name: "examples",
		Fields: []*Field{
			author,
		},
	}

	if !reflect.DeepEqual(table, expected) {
		ex := utter.Sdump(expected)
		ac := utter.Sdump(table)
		outputDiff(ex, "expected.txt")
		outputDiff(ac, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(ex, ac))
	}
}

func TestParseAndLoad_slices(t *testing.T) {
	exit.TestableExit()
	Debug = true
	code := strings.Replace(`package pkg1

type Example struct {
	Labels      []string   |sql:"encode: json"|
	Categories  []Category |sql:"encode: json"|
}

type Category int32
`, "|", "`", -1)

	source := Source{"issue.go", bytes.NewBufferString(code)}

	pkgStore, err := ParseGroups(token.NewFileSet(), Group{"pkg1", []Source{source}})
	if err != nil {
		t.Fatalf("Error parsing: %s", err)
	}

	table, err := load(pkgStore, "pkg1", "Example")
	if err != nil {
		t.Fatalf("Error loading: %s", err)
	}

	labels := &Field{Node{"Labels", Type{"", "", "[]string", Slice}, nil}, "labels", ENCJSON, Tag{Encode: "json"}}
	categories := &Field{Node{"Categories", Type{"", "", "Category", Slice}, nil}, "categories", ENCJSON, Tag{Encode: "json"}}

	expected := &TableDescription{
		Type: "Example",
		Name: "examples",
		Fields: []*Field{
			labels,
			categories,
		},
	}

	if !reflect.DeepEqual(table, expected) {
		ex := utter.Sdump(expected)
		ac := utter.Sdump(table)
		outputDiff(ex, "expected.txt")
		outputDiff(ac, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(ex, ac))
	}
}

func TestParseAndLoad_multipleNamesWithTags(t *testing.T) {
	exit.TestableExit()
	Debug = true
	code := strings.Replace(`package pkg1

type Example struct {
	Aaa, Bbb   string  |sql:"size: 32, index: foo"|
}
`, "|", "`", -1)

	source := Source{"issue.go", bytes.NewBufferString(code)}

	pkgStore, err := ParseGroups(token.NewFileSet(), Group{"pkg1", []Source{source}})
	if err != nil {
		t.Fatalf("Error parsing: %s", err)
	}

	table, err := load(pkgStore, "pkg1", "Example")
	if err != nil {
		t.Fatalf("Error loading: %s", err)
	}

	aaa := &Field{Node{"Aaa", Type{"", "", "string", String}, nil}, "aaa", ENCNONE, Tag{Size: 32, Index: "foo"}}
	bbb := &Field{Node{"Bbb", Type{"", "", "string", String}, nil}, "bbb", ENCNONE, Tag{Size: 32, Index: "foo"}}

	idx := &Index{"foo", false, []*Field{aaa, bbb}}

	expected := &TableDescription{
		Type: "Example",
		Name: "examples",
		Fields: []*Field{
			aaa,
			bbb,
		},
		Index: []*Index{
			idx,
		},
	}

	if !reflect.DeepEqual(table, expected) {
		ex := utter.Sdump(expected)
		ac := utter.Sdump(table)
		outputDiff(ex, "expected.txt")
		outputDiff(ac, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(ex, ac))
	}
}

//-------------------------------------------------------------------------------------------------

func TestParseAndLoad_embeddedTypes(t *testing.T) {
	exit.TestableExit()
	Debug = true
	code := strings.Replace(`package pkg1

import "go/token"

type Example struct {
	Cat    Category
	Commit
	token.Position // a convenient concrete type with exported fields
}

type Category int32

type Commit struct {
	Author
	Message   string
}

type Author struct {
	Name     string
	Email    string
}

`, "|", "`", -1)

	source := Source{"issue.go", bytes.NewBufferString(code)}

	pkgStore, err := ParseGroups(token.NewFileSet(), Group{"pkg1", []Source{source}})
	if err != nil {
		t.Fatalf("Error parsing: %s", err)
	}

	table, err := load(pkgStore, "pkg1", "Example")
	if err != nil {
		t.Fatalf("Error loading: %s", err)
	}

	p1 := &Node{Name: "Commit"}
	p2 := &Node{Name: "Author", Parent: p1}
	p3 := &Node{Name: "Position"}

	category := &Field{Node{"Cat", Type{"", "", "Category", Int32}, nil}, "cat", ENCNONE, Tag{}}
	name := &Field{Node{"Name", Type{"", "", "string", String}, p2}, "name", ENCNONE, Tag{}}
	email := &Field{Node{"Email", Type{"", "", "string", String}, p2}, "email", ENCNONE, Tag{}}
	message := &Field{Node{"Message", Type{"", "", "string", String}, p1}, "message", ENCNONE, Tag{}}
	filename := &Field{Node{"Filename", Type{"", "", "string", String}, p3}, "filename", ENCNONE, Tag{}}
	offset := &Field{Node{"Offset", Type{"", "", "int", Int}, p3}, "offset", ENCNONE, Tag{}}
	line := &Field{Node{"Line", Type{"", "", "int", Int}, p3}, "line", ENCNONE, Tag{}}
	column := &Field{Node{"Column", Type{"", "", "int", Int}, p3}, "column", ENCNONE, Tag{}}

	expected := &TableDescription{
		Type: "Example",
		Name: "examples",
		Fields: []*Field{
			category,
			name,
			email,
			message,
			filename,
			offset,
			line,
			column,
		},
	}

	if !reflect.DeepEqual(table, expected) {
		ex := utter.Sdump(expected)
		ac := utter.Sdump(table)
		outputDiff(ex, "expected.txt")
		outputDiff(ac, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(ex, ac))
	}
}

func TestParseAndLoad_embeddedTypes_inDifferentPackages(t *testing.T) {
	exit.TestableExit()
	Debug = true
	code := strings.Replace(`package pkg1

import "github.com/rickb777/sqlgen2/demo"

type Example struct {
	Cat      demo.Category
	demo.Dates
	Name     string
}
`, "|", "`", -1)

	source := Source{"issue1.go", bytes.NewBufferString(code)}

	pkgStore, err := ParseGroups(token.NewFileSet(),
		Group{"pkg1", []Source{source}},
	)
	if err != nil {
		t.Fatalf("Error parsing: %s", err)
	}

	table, err := load(pkgStore, "pkg1", "Example")
	if err != nil {
		t.Fatalf("Error loading: %s", err)
	}

	p1 := &Node{Name: "Dates"}

	category := &Field{Node{"Cat", Type{"github.com/rickb777/sqlgen2/demo", "demo", "Category", Uint8}, nil}, "cat", ENCNONE, Tag{}}
	after := &Field{Node{"After", Type{"", "", "string", String}, p1}, "after", ENCNONE, Tag{Size: 20}}
	before := &Field{Node{"Before", Type{"", "", "string", String}, p1}, "before", ENCNONE, Tag{Size: 20}}
	name := &Field{Node{"Name", Type{"", "", "string", String}, nil}, "name", ENCNONE, Tag{}}

	expected := &TableDescription{
		Type: "Example",
		Name: "examples",
		Fields: []*Field{
			category,
			after,
			before,
			name,
		},
	}

	if !reflect.DeepEqual(table, expected) {
		ex := utter.Sdump(expected)
		ac := utter.Sdump(table)
		outputDiff(ex, "expected.txt")
		outputDiff(ac, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(ex, ac))
	}
}

func TestParseAndLoad_multiplePackagesWithPrimaryAndIndexes(t *testing.T) {
	exit.TestableExit()
	Debug = true
	code := strings.Replace(`package pkg1

import (
	"github.com/rickb777/sqlgen2/demo"
	"math/big"
	"time"
)

type Example struct {
	Id         int64         |sql:"pk: true, auto: true"|
	Number     int
	Category   demo.Category
	Foo        int           |sql:"-"|
	Commit     *Commit
	Title      string        |sql:"index: titleIdx"|
	Hobby      string        |sql:"size: 2048"|
	Labels     []string      |sql:"encode: json"|
	Active     bool
	Avatar     []byte
	Fave       big.Int       |sql:"encode: json"|
	Updated    time.Time     |sql:"encode: text"|
}

type Commit struct {
	Message   string        |sql:"size: 2048, name: text"|
	Timestamp time.Time     |sql:"-"|
	Author    *demo.Author
}

`, "|", "`", -1)

	source := Source{"issue1.go", bytes.NewBufferString(code)}

	pkgStore, err := ParseGroups(token.NewFileSet(), Group{"pkg1", []Source{source}})
	if err != nil {
		t.Fatalf("Error parsing: %s", err)
	}

	table, err := load(pkgStore, "pkg1", "Example")
	if err != nil {
		t.Fatalf("Error loading: %s", err)
	}

	p1 := &Node{Name: "Commit", Type: Type{Name: "Commit", Base: Struct}}
	p2 := &Node{Name: "Author", Type: Type{PkgPath: "github.com/rickb777/sqlgen2/demo", PkgName: "demo", Name: "Author", Base: Struct}, Parent: p1}

	id := &Field{Node{"Id", Type{"", "", "int64", Int64}, nil}, "id", ENCNONE, Tag{Primary: true, Auto: true}}
	number := &Field{Node{"Number", Type{"", "", "int", Int}, nil}, "number", ENCNONE, Tag{}}
	category := &Field{Node{"Category", Type{"github.com/rickb777/sqlgen2/demo", "demo", "Category", Uint8}, nil}, "category", ENCNONE, Tag{}}
	commitMessage := &Field{Node{"Message", Type{"", "", "string", String}, p1}, "text", ENCNONE, Tag{Size: 2048, Name: "text"}}
	//commitTimestamp := &Field{Node{"Timestamp", Type{"time", "Time", String}, p1}, "commit_timestamp", VARCHAR, ENCNONE, Tag{}}
	authorName := &Field{Node{"Name", Type{"", "", "string", String}, p2}, "commit_author_name", ENCNONE, Tag{Prefixed: true}}
	authorEmail := &Field{Node{"Email", Type{"github.com/rickb777/sqlgen2/demo", "demo", "Email", String}, p2}, "commit_author_email", ENCNONE, Tag{Prefixed: true}}
	authorUser := &Field{Node{"Username", Type{"", "", "string", String}, p2}, "commit_author_username", ENCNONE, Tag{Prefixed: true}}
	title := &Field{Node{"Title", Type{"", "", "string", String}, nil}, "title", ENCNONE, Tag{Index: "titleIdx"}}
	////owner := &Field{"Owner", "team_owner", VARCHAR, Tag{}}
	hobby := &Field{Node{"Hobby", Type{"", "", "string", String}, nil}, "hobby", ENCNONE, Tag{Size: 2048}}
	labels := &Field{Node{"Labels", Type{"", "", "[]string", Slice}, nil}, "labels", ENCJSON, Tag{Encode: "json"}}
	active := &Field{Node{"Active", Type{"", "", "bool", Bool}, nil}, "active", ENCNONE, Tag{}}
	avatar := &Field{Node{"Avatar", Type{"", "", "[]byte", Slice}, nil}, "avatar", ENCNONE, Tag{}}
	fave := &Field{Node{"Fave", Type{"math/big", "big", "Int", Struct}, nil}, "fave", ENCJSON, Tag{Encode: "json"}}
	updated := &Field{Node{"Updated", Type{"time", "time", "Time", Struct}, nil}, "updated", ENCTEXT, Tag{Encode: "text"}}

	ititle := &Index{"titleIdx", false, []*Field{title}}

	expected := &TableDescription{
		Type: "Example",
		Name: "examples",
		Fields: []*Field{
			id,
			number,
			category,
			commitMessage,
			//commitTimestamp,
			authorName,
			authorEmail,
			authorUser,
			title,
			////owner,
			hobby,
			labels,
			active,
			avatar,
			fave,
			updated,
		},
		Index: []*Index{
			ititle,
		},
		Primary: id,
	}

	if !reflect.DeepEqual(table, expected) {
		ex := utter.Sdump(expected)
		ac := utter.Sdump(table)
		outputDiff(ex, "expected.txt")
		outputDiff(ac, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(ex, ac))
	} else {
		t.Log("OK\n")
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
		buf.WriteString(fmt.Sprintf("%-50s", truncate(ea, 50)))
		if i < len(bb) {
			eb := strings.Replace(bb[i], "\t", "    ", -1)
			if ea != eb {
				buf.WriteString(" <> ")
			} else {
				buf.WriteString("    ")
			}
			buf.WriteString(truncate(eb, 50))
		}
		buf.WriteByte('\n')
		i++
	}

	for ; i < len(bb); i++ {
		buf.WriteString(fmt.Sprintf("%50s <> ", ""))
		eb := strings.Replace(bb[i], "\t", "    ", -1)
		buf.WriteString(truncate(eb, 50))
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
