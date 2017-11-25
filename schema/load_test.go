package schema

import (
	"bytes"
	"fmt"
	"github.com/kortschak/utter"
	. "github.com/rickb777/sqlgen2/sqlgen/parse"
	"github.com/rickb777/sqlgen2/sqlgen/parse/exit"
	"go/token"
	"os"
	"reflect"
	"strings"
	"testing"
)

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

	table, err := Load(pkgStore, "pkg1", "Example")
	if err != nil {
		t.Fatalf("Error loading: %s", err)
	}

	p1 := &Node{Name: "Commit"}
	p2 := &Node{Name: "Author", Parent: p1}
	author := &Field{Node{"Name", Type{"", "string", String}, p2}, "commit_author_name", VARCHAR, ENCNONE, Tag{}}

	expected := &Table{
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

	table, err := Load(pkgStore, "pkg1", "Example")
	if err != nil {
		t.Fatalf("Error loading: %s", err)
	}

	labels := &Field{Node{"Labels", Type{"", "[]string", Slice}, nil}, "labels", BLOB, ENCJSON, Tag{Encode: "json"}}
	categories := &Field{Node{"Categories", Type{"", "Category", Slice}, nil}, "categories", BLOB, ENCJSON, Tag{Encode: "json"}}

	expected := &Table{
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

	table, err := Load(pkgStore, "pkg1", "Example")
	if err != nil {
		t.Fatalf("Error loading: %s", err)
	}

	aaa := &Field{Node{"Aaa", Type{"", "string", String}, nil}, "aaa", VARCHAR, ENCNONE, Tag{Size: 32, Index: "foo"}}
	bbb := &Field{Node{"Bbb", Type{"", "string", String}, nil}, "bbb", VARCHAR, ENCNONE, Tag{Size: 32, Index: "foo"}}

	idx := &Index{"foo", false, []*Field{aaa, bbb}}

	expected := &Table{
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

func TestParseAndLoad_embeddedTypes(t *testing.T) {
	exit.TestableExit()
	Debug = true
	code := strings.Replace(`package pkg1

type Example struct {
	Cat    Category
	Commit
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

	table, err := Load(pkgStore, "pkg1", "Example")
	if err != nil {
		t.Fatalf("Error loading: %s", err)
	}

	p1 := &Node{Name: "Commit"}
	p2 := &Node{Name: "Author", Parent: p1}

	category := &Field{Node{"Cat", Type{"", "Category", Int32}, nil}, "cat", INTEGER, ENCNONE, Tag{}}
	name := &Field{Node{"Name", Type{"", "string", String}, p2}, "commit_author_name", VARCHAR, ENCNONE, Tag{}}
	email := &Field{Node{"Email", Type{"", "string", String}, p2}, "commit_author_email", VARCHAR, ENCNONE, Tag{}}
	message := &Field{Node{"Message", Type{"", "string", String}, p1}, "commit_message", VARCHAR, ENCNONE, Tag{}}

	expected := &Table{
		Type: "Example",
		Name: "examples",
		Fields: []*Field{
			category,
			name,
			email,
			message,
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
	"time"
)

type Example struct {
	Id         int64    |sql:"pk: true, auto: true"|
	Number     int
	Category   demo.Category
	Foo        int      |sql:"-"|
	Commit     *Commit
	Title      string   |sql:"index: titleIdx"|
	////TODO Owner      string   |sql:"name: team_owner"|
	Hobby      string   |sql:"size: 2048"|
	Labels     []string |sql:"encode: json"|
	Active     bool
}

type Commit struct {
	Message   string        |sql:"size: 2048"|
	//Timestamp time.Time     |sql:"-"|
	Author    *demo.Author
}

`, "|", "`", -1)

	source := Source{"issue1.go", bytes.NewBufferString(code)}

	pkgStore, err := ParseGroups(token.NewFileSet(), Group{"pkg1", []Source{source}})
	if err != nil {
		t.Fatalf("Error parsing: %s", err)
	}

	table, err := Load(pkgStore, "pkg1", "Example")
	if err != nil {
		t.Fatalf("Error loading: %s", err)
	}

	p1 := &Node{Name: "Commit"}
	p2 := &Node{Name: "Author", Parent: p1}

	id := &Field{Node{"Id", Type{"", "int64", Int64}, nil}, "id", INTEGER, ENCNONE, Tag{Primary: true, Auto: true}}
	number := &Field{Node{"Number", Type{"", "int", Int}, nil}, "number", INTEGER, ENCNONE, Tag{}}
	category := &Field{Node{"Category", Type{"demo", "Category", Uint8}, nil}, "category", INTEGER, ENCNONE, Tag{}}
	commitTitle := &Field{Node{"Message", Type{"", "string", String}, p1}, "commit_message", VARCHAR, ENCNONE, Tag{Size: 2048}}
	//commitTimestamp := &Field{Node{"Timestamp", Type{"time", "Time", String}, p1}, "commit_timestamp", VARCHAR, ENCNONE, Tag{}}
	authorName := &Field{Node{"Name", Type{"", "string", String}, p2}, "commit_author_name", VARCHAR, ENCNONE, Tag{}}
	authorEmail := &Field{Node{"Email", Type{"", "string", String}, p2}, "commit_author_email", VARCHAR, ENCNONE, Tag{}}
	authorUser := &Field{Node{"Username", Type{"", "string", String}, p2}, "user", VARCHAR, ENCNONE, Tag{Name: "user"}}
	title := &Field{Node{"Title", Type{"", "string", String}, nil}, "title", VARCHAR, ENCNONE, Tag{Index: "titleIdx"}}
	////owner := &Field{"Owner", "team_owner", VARCHAR, Tag{}}
	hobby := &Field{Node{"Hobby", Type{"", "string", String}, nil}, "hobby", VARCHAR, ENCNONE, Tag{Size: 2048}}
	labels := &Field{Node{"Labels", Type{"", "[]string", Slice}, nil}, "labels", BLOB, ENCJSON, Tag{Encode: "json"}}
	active := &Field{Node{"Active", Type{"", "bool", Bool}, nil}, "active", BOOLEAN, ENCNONE, Tag{}}

	ititle := &Index{"titleIdx", false, []*Field{title}}

	expected := &Table{
		Type: "Example",
		Name: "examples",
		Fields: []*Field{
			id,
			number,
			category,
			commitTitle,
			//commitTimestamp,
			authorName,
			authorEmail,
			authorUser,
			title,
			////owner,
			hobby,
			labels,
			active,
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
