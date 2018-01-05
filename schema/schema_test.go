package schema

import (
	"testing"
	. "github.com/rickb777/sqlgen2/sqlgen/parse"
)

func TestDistinctTypes(t *testing.T) {
	i64 := Type{"", "", "int64", false, Int64}
	str := Type{"", "", "string", false, String}
	spt := Type{"", "", "string", true, String}
	cat := Type{"", "", "Category", false, Int32}
	bgi := Type{"math/big", "big", "Int", false, Struct}
	bys := Type{"", "", "[]byte", false, Slice}
	tim := Type{"time", "time", "Time", false, Struct}

	id := &Field{Node{"Id", i64, nil}, "id", ENCNONE, Tag{}}
	category := &Field{Node{"Cat", cat, nil}, "cat", ENCNONE, Tag{}}
	name := &Field{Node{"Name", str, nil}, "name", ENCNONE, Tag{}}
	qual := &Field{Node{"Qual", spt, nil}, "qual", ENCNONE, Tag{}}
	fave := &Field{Node{"Fave", bgi, nil}, "fave", ENCJSON, Tag{}}
	avatar := &Field{Node{"Avatar", bys, nil}, "avatar", ENCNONE, Tag{}}
	updated := &Field{Node{"Updated", tim, nil}, "updated", ENCTEXT, Tag{}}

	cases := []struct {
		list     FieldList
		expected TypeSet
	}{
		{FieldList{id}, NewTypeSet(i64)},
		{FieldList{id, id, id}, NewTypeSet(i64)},
		{FieldList{id, category}, NewTypeSet(i64, cat)},
		{FieldList{id, name, qual, category, fave, avatar, updated}, NewTypeSet(i64, str, spt, cat, bgi, bys, tim)},
	}
	for _, c := range cases {
		s := c.list.DistinctTypes()
		if !NewTypeSet(s...).Equals(c.expected) {
			t.Errorf("expected %d::%+v but got %d::%+v", len(c.expected), c.expected, len(s), s)
		}
	}

}
