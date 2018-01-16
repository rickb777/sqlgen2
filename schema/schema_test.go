package schema

import (
	"testing"
	. "github.com/rickb777/sqlgen2/sqlgen/parse"
)

func TestDistinctTypes(t *testing.T) {
	i64 := Type{Name: "int64", Base: Int64}
	boo := Type{Name: "bool", Base: Bool}
	cat := Type{Name: "Category", Base: Int32}
	str := Type{Name: "string", Base: String}
	spt := Type{Name: "string", IsPtr: true, Base: String}
	ipt := Type{Name: "int32", IsPtr: true, Base: Int32}
	upt := Type{Name: "uint32", IsPtr: true, Base: Uint32}
	fpt := Type{Name: "float32", IsPtr: true, Base: Float32}
	sli := Type{Name: "[]string", Base: Slice}
	bgi := Type{PkgPath: "math/big", PkgName: "big", Name: "Int", Base: Struct}
	bys := Type{Name: "[]byte", Base: Slice}
	tim := Type{PkgPath: "time", PkgName: "time", Name: "Time", Base: Struct}

	id := &Field{Node{"Id", i64, nil}, "id", ENCNONE, Tag{Primary: true, Auto: true}}
	category := &Field{Node{"Cat", cat, nil}, "cat", ENCNONE, Tag{Index: "catIdx"}}
	name := &Field{Node{"Name", str, nil}, "username", ENCNONE, Tag{Size: 2048, Name: "username", Unique: "nameIdx"}}
	active := &Field{Node{"Active", boo, nil}, "active", ENCNONE, Tag{}}
	qual := &Field{Node{"Qual", spt, nil}, "qual", ENCNONE, Tag{}}
	diff := &Field{Node{"Diff", ipt, nil}, "diff", ENCNONE, Tag{}}
	age := &Field{Node{"Age", upt, nil}, "age", ENCNONE, Tag{}}
	bmi := &Field{Node{"Bmi", fpt, nil}, "bmi", ENCNONE, Tag{}}
	labels := &Field{Node{"Labels", sli, nil}, "labels", ENCJSON, Tag{Encode: "json"}}
	fave := &Field{Node{"Fave", bgi, nil}, "fave", ENCJSON, Tag{Encode: "json"}}
	avatar := &Field{Node{"Avatar", bys, nil}, "avatar", ENCNONE, Tag{}}
	updated := &Field{Node{"Updated", tim, nil}, "updated", ENCTEXT, Tag{Size: 100, Encode: "text"}}

	//icat := &Index{"catIdx", false, FieldList{category}}
	//iname := &Index{"nameIdx", true, FieldList{name}}

	cases := []struct {
		list     FieldList
		expected TypeSet
	}{
		{FieldList{id}, NewTypeSet(i64)},
		{FieldList{id, id, id}, NewTypeSet(i64)},
		{FieldList{id, category}, NewTypeSet(i64, cat)},
		{FieldList{id,
			category,
			name,
			qual,
			diff,
			age,
			bmi,
			active,
			labels,
			fave,
			avatar,
			updated}, NewTypeSet(i64, boo, cat, str, spt, ipt, upt, fpt, bgi, sli, bys, tim)},
	}
	for _, c := range cases {
		s := c.list.DistinctTypes()
		if !NewTypeSet(s...).Equals(c.expected) {
			t.Errorf("expected %d::%+v but got %d::%+v", len(c.expected), c.expected, len(s), s)
		}
	}
}
