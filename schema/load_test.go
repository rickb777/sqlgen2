package schema

import (
	. "github.com/rickb777/sqlgen2/sqlgen/parse"
	"github.com/rickb777/sqlgen2/sqlgen/parse/exit"
	"bytes"
	"reflect"
	"strings"
	"testing"
	"github.com/kortschak/utter"
)

func TestConvertLeafNodeToField(t *testing.T) {
	exit.TestableExit()

	ef1 := &Field{"Id", nil, "", INTEGER, ENCNONE, true, false, 0}
	doCompare(t, &Node{
		Name: "Id",
		Type: Type{"", "uint32", Uint32},
		Tags: &Tag{Primary: true},
	},
		true, ef1, nil)

	ef2 := &Field{"Wibble", nil, "", VARCHAR, ENCNONE, false, false, 0}
	doCompare(t, &Node{
		Name: "Wibble",
		Type: Type{"stringy", "Thingy", String},
		Tags: &Tag{},
	},
		true, ef2, nil)

	ef3 := &Field{"Bibble", nil, "", VARCHAR, ENCNONE, false, false, 0}
	doCompare(t, &Node{
		Name: "Bibble",
		Type: Type{"", "string", String},
		Tags: &Tag{Index: "BibbleIdx"},
	},
		true, ef3,
		&Index{Name: "BibbleIdx", Fields: []*Field{ef3}})

	ef4 := &Field{"Bobble", nil, "", VARCHAR, ENCNONE, false, false, 0}
	doCompare(t, &Node{
		Name: "Bobble",
		Type: Type{"", "string", String},
		Tags: &Tag{Unique: "BobbleIdx"},
	},
		true, ef4,
		&Index{Name: "BobbleIdx", Fields: []*Field{ef4}, Unique: true})
}

func doCompare(t *testing.T, leaf *Node, wantOk bool, expectedField *Field, expectedIndex *Index) {
	t.Helper()

	table := new(Table)
	indices := map[string]*Index{}

	field, ok := convertLeafNodeToField(leaf, indices, table)
	if !wantOk {
		if ok {
			t.Errorf("Should be not OK -> expected %+v", expectedField)
		}
		return
	}

	if !ok {
		t.Errorf("NOT OK -> expected %+v", expectedField)
	}

	if !reflect.DeepEqual(field, expectedField) {
		t.Errorf("\nexpected %+v\n"+
			"but got  %+v", expectedField, field)
	}

	if expectedIndex != nil {
		if len(indices) != 1 {
			t.Errorf("\nexpected 1 index %+v", expectedIndex)
		} else {
			index := indices[expectedField.Name+"Idx"]
			if !reflect.DeepEqual(index, expectedIndex) {
				t.Errorf("\nexpected %+v\n"+
					"but got  %+v", expectedIndex, index)
			}
		}
	}
}

func TestParseAndLoad(t *testing.T) {
	exit.TestableExit()
	code := strings.Replace(`package pkg1

type Example struct {
	Id         int64    |sql:"pk: true, auto: true"|
	Number     int
	Category   Category
	Foo        int      |sql:"-"|
	Commit     *Commit
	Title      string   |sql:"index: titleIdx"|
	//TODO Owner      string   |sql:"name: team_owner"|
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

	source := Source{"issue.go", bytes.NewBufferString(code)}

	node, err := DoParse("pkg1", "Example", []Source{source})
	if err != nil {
		t.Errorf("Error parsing: %s", err)
	}

	table := Load(node)

	id := &Field{"Id", []string{"Id"}, "id", INTEGER, ENCNONE, true, true, 0}
	number := &Field{"Number", []string{"Number"}, "number", INTEGER, ENCNONE, false, false, 0}
	category := &Field{"Category", []string{"Category"}, "category", INTEGER, ENCNONE, false, false, 0}
	commitTitle := &Field{"Message", []string{"Commit", "Message"}, "commit_message", VARCHAR, ENCNONE, false, false, 0}
	timestamp := &Field{"Timestamp", []string{"Commit", "Timestamp"}, "commit_timestamp", INTEGER, ENCNONE, false, false, 0}
	author := &Field{"Name", []string{"Commit", "Author", "Name"}, "commit_author_name", VARCHAR, ENCNONE, false, false, 0}
	email := &Field{"Email", []string{"Commit", "Author", "Email"}, "commit_author_email", VARCHAR, ENCNONE, false, false, 0}
	title := &Field{"Title", []string{"Title"}, "title", VARCHAR, ENCNONE, false, false, 0}
	//owner := &Field{"Owner", "team_owner", VARCHAR, false, false, 0}
	hobby := &Field{"Hobby", []string{"Hobby"}, "hobby", VARCHAR, ENCNONE, false, false, 2048}
	labels := &Field{"Labels", []string{"Labels"}, "labels", BLOB, ENCJSON, false, false, 0}
	active := &Field{"Active", []string{"Active"}, "active", BOOLEAN, ENCNONE, false, false, 0}

	ititle := &Index{"titleIdx", false, []*Field{title}}

	expected := &Table{
		Type: "Example",
		Name: "examples",
		Fields: []*Field{
			id,
			number,
			category,
			commitTitle,
			timestamp,
			author,
			email,
			title,
			//owner,
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
		t.Errorf("\nexpected %s\n"+
			"but got  %s", utter.Sdump(expected), utter.Sdump(table))
	}
}
