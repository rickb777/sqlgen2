package schema

import (
	. "github.com/rickb777/sqlgen/sqlgen/parse"
	"github.com/rickb777/sqlgen/sqlgen/parse/exit"
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func TestConvertLeafNodeToField(t *testing.T) {
	exit.TestableExit()

	ef1 := &Field{"Id", "", INTEGER, true, false, 0}
	doCompare(t, &Node{
		Name: "Id",
		Type: Type{"", "uint32", Uint32},
		Tags: &Tag{Primary: true},
	},
		true, ef1, nil)

	ef2 := &Field{"Wibble", "", VARCHAR, false, false, 0}
	doCompare(t, &Node{
		Name: "Wibble",
		Type: Type{"stringy", "Thingy", String},
		Tags: &Tag{},
	},
		true, ef2, nil)

	ef3 := &Field{"Bibble", "", VARCHAR, false, false, 0}
	doCompare(t, &Node{
		Name: "Bibble",
		Type: Type{"", "string", String},
		Tags: &Tag{Index: "BibbleIdx"},
	},
		true, ef3,
		&Index{Name: "BibbleIdx", Fields: []*Field{ef3}})

	ef4 := &Field{"Bobble", "", VARCHAR, false, false, 0}
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
	code := strings.Replace(`package pkg1

type Struct struct {
	Id       int64 |sql:"pk: true, auto: true"|
	Number   int
	Foo      int |sql:"-"|
	Title, Description, Owner string
}`, "|", "`", -1)

	source := Source{"issue.go", bytes.NewBufferString(code)}

	_, err := DoParse("pkg1", "Struct", []Source{source})
	if err != nil {
		t.Errorf("Error parsing: %s", err)
	}
	//TODO finish
}
