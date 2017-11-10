package schema

import (
	"bytes"
	"fmt"
	"github.com/kortschak/utter"
	. "github.com/rickb777/sqlgen2/sqlgen/parse"
	"github.com/rickb777/sqlgen2/sqlgen/parse/exit"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"
)

//func TestConvertLeafNodeToField(t *testing.T) {
//	exit.TestableExit()
//
//	ef1 := &Field{"Id", Type{"", "uint32", Uint32}, nil, "id", INTEGER, ENCNONE, Tag{Primary:true}}
//	doCompare(t, &Node{
//		Name: "Id",
//		Type: Type{"", "uint32", Uint32},
//		Tags: &Tag{Primary: true},
//	},
//		true, ef1, nil)
//
//	ef2 := &Field{"Wibble", Type{"stringy", "Thingy", String}, nil, "", VARCHAR, ENCNONE, Tag{}}
//	doCompare(t, &Node{
//		Name: "Wibble",
//		Type: Type{"stringy", "Thingy", String},
//		Tags: &Tag{},
//	},
//		true, ef2, nil)
//
//	ef3 := &Field{"Bibble", Type{"", "string", String}, nil, "", VARCHAR, ENCNONE, Tag{}}
//	doCompare(t, &Node{
//		Name: "Bibble",
//		Type: Type{"", "string", String},
//		Tags: &Tag{Index: "BibbleIdx"},
//	},
//		true, ef3,
//		&Index{Name: "BibbleIdx", Fields: []*Field{ef3}})
//
//	ef4 := &Field{"Bobble", Type{"", "string", String}, nil, "", VARCHAR, ENCNONE, Tag{}}
//	doCompare(t, &Node{
//		Name: "Bobble",
//		Type: Type{"", "string", String},
//		Tags: &Tag{Unique: "BobbleIdx"},
//	},
//		true, ef4,
//		&Index{Name: "BobbleIdx", Fields: []*Field{ef4}, Unique: true})
//}
//
//func doCompare(t *testing.T, leaf *Node, wantOk bool, expectedField *Field, expectedIndex *Index) {
//	t.Helper()
//
//	table := new(Table)
//	indices := map[string]*Index{}
//	tags := make(map[string]Tag)
//
//	field, ok := convertLeafNodeToField(nil, "", "", tags, indices, table)
//	if !wantOk {
//		if ok {
//			t.Errorf("Should be not OK -> expected %+v", expectedField)
//		}
//		return
//	}
//
//	if !ok {
//		t.Errorf("NOT OK -> expected %+v", expectedField)
//	}
//
//	if !reflect.DeepEqual(field, expectedField) {
//		t.Errorf("\nexpected %+v\n"+
//			"but got  %+v", expectedField, field)
//	}
//
//	if expectedIndex != nil {
//		if len(indices) != 1 {
//			t.Errorf("\nexpected 1 index %+v", expectedIndex)
//		} else {
//			index := indices[expectedIndex.Name]
//			if !reflect.DeepEqual(index, expectedIndex) {
//				t.Errorf("\nexpected %+v\n"+
//					"but got  %+v", expectedIndex, index)
//			}
//		}
//	}
//}

func TestParseAndLoad(t *testing.T) {
	exit.TestableExit()
	Debug = true
	code := strings.Replace(`package pkg1

type Example struct {
	Id         int64    |sql:"pk: true, auto: true"|
	Number     int
	//Category   Category
	Foo        int      |sql:"-"|
	//Commit     *Commit
	Title      string   |sql:"index: titleIdx"|
	////TODO Owner      string   |sql:"name: team_owner"|
	Hobby      string   |sql:"size: 2048"|
	Labels     []string |sql:"encode: json"|
	Active     bool
}

//type Category int32
//
//type Commit struct {
//	Message   string
//	Timestamp int64 // TODO should be able to support time.Time
//	Author    *Author
//}
//
//type Author struct {
//	Name     string
//	Email    string
//}

`, "|", "`", -1)

	source := Source{"issue.go", bytes.NewBufferString(code)}

	pkgStore, err := ParseGroups(Group{"Example", []Source{source}})
	if err != nil {
		t.Fatalf("Error parsing: %s", err)
	}

	table, err := Load(pkgStore, "pkg1", "Example")
	if err != nil {
		t.Fatalf("Error loading: %s", err)
	}

	id := &Field{"Id", Type{"", "int64", Int64}, PathOf("Id"), "id", INTEGER, ENCNONE, Tag{Primary: true, Auto: true}}
	number := &Field{"Number", Type{"", "int", Int}, PathOf("Number"), "number", INTEGER, ENCNONE, Tag{}}
	//category := &Field{"Category", Type{"pkg1", "Category", Int32}, PathOf("Category"), "category", INTEGER, ENCNONE, Tag{}}
	//commitTitle := &Field{"Message", Type{"", "string", String}, PathOf("Commit", "Message"), "commit_message", VARCHAR, ENCNONE, Tag{}}
	//timestamp := &Field{"Timestamp", Type{"", "int64", Int64}, PathOf("Commit", "Timestamp"), "commit_timestamp", INTEGER, ENCNONE, Tag{}}
	//author := &Field{"Name", Type{"", "string", String}, PathOf("Commit", "Author", "Name"), "commit_author_name", VARCHAR, ENCNONE, Tag{}}
	//email := &Field{"Email", Type{"", "string", String}, PathOf("Commit", "Author", "Email"), "commit_author_email", VARCHAR, ENCNONE, Tag{}}
	title := &Field{"Title", Type{"", "string", String}, PathOf("Title"), "title", VARCHAR, ENCNONE, Tag{Index:"titleIdx"}}
	////owner := &Field{"Owner", "team_owner", VARCHAR, Tag{}}
	hobby := &Field{"Hobby", Type{"", "string", String}, PathOf("Hobby"), "hobby", VARCHAR, ENCNONE, Tag{Size: 2048}}
	labels := &Field{"Labels", Type{"", "[]string", Slice}, PathOf("Labels"), "labels", BLOB, ENCJSON, Tag{Encode:"json"}}
	active := &Field{"Active", Type{"", "bool", Bool}, PathOf("Active"), "active", BOOLEAN, ENCNONE, Tag{}}

	ititle := &Index{"titleIdx", false, []*Field{title}}

	expected := &Table{
		Type: "Example",
		Name: "examples",
		Fields: []*Field{
			id,
			number,
			//category,
			//commitTitle,
			//timestamp,
			//author,
			//email,
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
		os.Stderr.Sync()
		time.Sleep(100 * time.Millisecond)
		ex := utter.Sdump(expected)
		ac := utter.Sdump(table)
		outputDiff(ex, "expected.txt")
		outputDiff(ac, "got.txt")
		t.Errorf("expected | got\n%s\n", sideBySideDiff(ex, ac))
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
