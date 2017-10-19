package schema

import (
	"testing"
	"github.com/rickb777/sqlgen/sqlgen/parse/exit"
	. "github.com/rickb777/sqlgen/sqlgen/parse"
	"reflect"
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
			&Index{Name: "BibbleIdx", Fields:[]*Field{ef3}})

	ef4 := &Field{"Bobble", "", VARCHAR, false, false, 0}
	doCompare(t, &Node{
		Name: "Bobble",
		Type: Type{"", "string", String},
		Tags: &Tag{Unique: "BobbleIdx"},
	},
		true, ef4,
			&Index{Name: "BobbleIdx", Fields:[]*Field{ef4}, Unique: true})
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

	if ! ok {
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
