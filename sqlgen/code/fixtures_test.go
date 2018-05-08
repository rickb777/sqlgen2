package code

import (
	. "github.com/rickb777/sqlgen2/schema"
	. "github.com/rickb777/sqlgen2/sqlgen/parse"
)

var i64 = Type{Name: "int64", Base: Int64}
var boo = Type{Name: "bool", Base: Bool}
var cat = Type{Name: "Category", Base: Int32}
var str = Type{Name: "string", Base: String}
var spt = Type{Name: "string", IsPtr: true, Base: String}
var phn = Type{Name: "PhoneNumber", IsPtr: true, Base: String}
var ipt = Type{Name: "int32", IsPtr: true, Base: Int32}
var upt = Type{Name: "uint32", IsPtr: true, Base: Uint32}
var fpt = Type{Name: "float32", IsPtr: true, Base: Float32}
var sli = Type{Name: "[]string", Base: Slice}
var bgi = Type{PkgPath: "math/big", PkgName: "big", Name: "Int", Base: Struct}
var bys = Type{Name: "[]byte", Base: Slice}
var scv1 = Type{Name: "Foo", IsScanner: true, IsValuer: true, Base: String}
var scv2 = Type{Name: "Foo", IsScanner: true, IsValuer: true, Base: String, IsPtr: true}
var scv3 = Type{Name: "Foo", IsScanner: true, IsValuer: true, Base: Int32, IsPtr: true}
var bar1 = Type{Name: "Bar", Base: String}
//var bar2 = Type{Name: "Bar", Base: String, IsPtr: true}
//var bar3 = Type{Name: "Bar", Base: Int32, IsPtr: true}
var tim = Type{PkgPath: "time", PkgName: "time", Name: "Time", Base: Struct}

var nums = &Node{Name: "Numbers", Type: Type{PkgName: "pkg1", Name: "Numbers", Base: Struct}}

var id = &Field{Node{"Id", i64, nil}, "id", ENCNONE, Tag{Primary: true, Auto: true}}
var category = &Field{Node{"Cat", cat, nil}, "cat", ENCNONE, Tag{Index: "catIdx"}}
var name = &Field{Node{"Name", str, nil}, "username", ENCNONE, Tag{Size: 2048, Name: "username", Default: "anon", Unique: "nameIdx"}}
var mobile = &Field{Node{"Mobile", phn, nil}, "mobile", ENCNONE, Tag{}}
var active = &Field{Node{"Active", boo, nil}, "active", ENCNONE, Tag{}}
var qual = &Field{Node{"Qual", spt, nil}, "qual", ENCNONE, Tag{}}
var diff = &Field{Node{"Diff", ipt, nums}, "diff", ENCNONE, Tag{}}
var age = &Field{Node{"Age", upt, nums}, "age", ENCNONE, Tag{}}
var bmi = &Field{Node{"Bmi", fpt, nums}, "bmi", ENCNONE, Tag{}}
var labels = &Field{Node{"Labels", sli, nil}, "labels", ENCJSON, Tag{Encode: "json"}}
var fave = &Field{Node{"Fave", bgi, nil}, "fave", ENCJSON, Tag{Encode: "json"}}
var avatar = &Field{Node{"Avatar", bys, nil}, "avatar", ENCNONE, Tag{}}
var fooey1 = &Field{Node{"Foo1", scv1, nil}, "foo1", ENCNONE, Tag{}}
var fooey2 = &Field{Node{"Foo2", scv2, nil}, "foo2", ENCNONE, Tag{}}
var fooey3 = &Field{Node{"Foo3", scv3, nil}, "foo3", ENCNONE, Tag{}}
var barey1 = &Field{Node{"Bar1", bar1, nil}, "bar1", ENCDRIVER, Tag{Encode: "driver"}}
//var barey2 = &Field{Node{"Bar2", bar2, nil}, "bar2", ENCDRIVER, Tag{Encode: "driver"}}
//var barey3 = &Field{Node{"Bar3", bar3, nil}, "bar3", ENCDRIVER, Tag{Encode: "driver"}}
var updated = &Field{Node{"Updated", tim, nil}, "updated", ENCTEXT, Tag{Size: 100, Encode: "text"}}

var icat = &Index{"catIdx", false, FieldList{category}}
var iname = &Index{"nameIdx", true, FieldList{name}}

func fixtureTable() *TableDescription {
	return &TableDescription{
		Type: "Example",
		Name: "examples",
		Fields: FieldList{
			id,
			category,
			name,
			mobile,
			qual,
			diff,
			age,
			bmi,
			active,
			labels,
			fave,
			avatar,
			fooey1,
			fooey2,
			fooey3,
			barey1,
			//barey2,
			//barey3,
			updated,
		},
		Index: []*Index{
			icat,
			iname,
		},
		Primary: id,
	}
}
