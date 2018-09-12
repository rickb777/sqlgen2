package parse

import (
	"os"
)

//func TestStructWith3FieldsAndTags(t *testing.T) {
//	doTestParseOK(t,
//		&Node{
//			Name: "Struct",
//			Type: Type{"pkg1", "Struct", Struct},
//			Nodes: []*Node{
//				{
//					Name: "Id",
//					Type: Type{"", "int64", Int64},
//					Tags: &Tag{Primary: true, Auto: true},
//				}, {
//					Name: "Number",
//					Type: Type{"", "int", Int},
//					Tags: &Tag{},
//				}, {
//					Name: "Title",
//					Type: Type{"", "string", String},
//					Tags: &Tag{},
//				}, {
//					Name: "Description",
//					Type: Type{"", "string", String},
//					Tags: &Tag{},
//				}, {
//					Name: "Owner",
//					Type: Type{"", "string", String},
//					Tags: &Tag{},
//				},
//			},
//		},
//		"pkg1", "Struct",
//		`package pkg1
//
//		type Struct struct {
//			Id       int64 |sql:"pk: true, auto: true"|
//			Number   int
//			Title, Description, Owner    string // must find all three fields
//		}`,
//	)
//}
//
//func TestStructWith1BoolFieldAndIgnoreTag(t *testing.T) {
//	doTestParseOK(t,
//		&Node{
//			Name: "Struct",
//			Type: Type{"pkg2", "Struct", Struct},
//			Nodes: []*Node{
//				{
//					Name: "Flag",
//					Type: Type{"", "bool", Bool},
//					Tags: &Tag{Skip: true},
//				},
//			},
//		},
//		"pkg2", "Struct",
//		`package pkg2
//
//		type Struct struct {
//			Flag  bool     |sql:"-"|
//		}`,
//	)
//}
//
//func TestStructWith1SliceFieldAndJsonTag(t *testing.T) {
//	doTestParseOK(t,
//		&Node{
//			Name: "Struct",
//			Type: Type{"pkg3", "Struct", Struct},
//			Nodes: []*Node{
//				{
//					Name: "Labels",
//					Type: Type{"", "[]string", Slice},
//					Tags: &Tag{Encode: "json"},
//				},
//			},
//		},
//		"pkg3", "Struct",
//		`package pkg3
//
//		type Struct struct {
//			Labels   []string  |sql:"encode: json"|
//		}`,
//	)
//}
//
//func TestStructWith1MapFieldAndJsonTag(t *testing.T) {
//	doTestParseOK(t,
//		&Node{
//			Name: "Struct",
//			Type: Type{"pkg4", "Struct", Struct},
//			Nodes: []*Node{
//				{
//					Name: "Table",
//					Type: Type{"", "map[string]int", Map},
//					Tags: &Tag{Encode: "json"},
//				},
//			},
//		},
//		"pkg4", "Struct",
//		`package pkg4
//
//		type Struct struct {
//			Table    map[string]int  |sql:"encode: json"|
//		}`,
//	)
//}

//func TestStructWithNestedStructType(t *testing.T) {
//	doTestParseOK(t,
//		&Node{
//			Name: "Struct",
//			Type: Type{"pkg5", "Struct", Struct},
//			Nodes: []*Node{
//				{
//					Name: "Author",
//					Type: Type{"", "Author", Ptr},
//					Tags: &Tag{},
//					Nodes: []*Node{
//						{
//							Name: "Name",
//							Type: Type{"", "string", String},
//							Tags: &Tag{},
//						},
//					},
//				},
//			},
//		},
//		"pkg5", "Struct",
//		`package pkg5
//
//		type Struct struct {
//			Author    *Author
//		}
//		type Author struct {
//			Name     string
//		}`,
//	)
//}

//func TestStructWithNestedSimpleType(t *testing.T) {
//	doTestParseOK(t,
//		&Node{
//			Name: "Struct",
//			Type: Type{"pkg6", "Struct", Struct},
//			Nodes: []*Node{
//				{
//					Name: "Cat",
//					Type: Type{"pkg6", "Category", Int32},
//					Tags: &Tag{},
//				},
//			},
//		},
//		"pkg6", "Struct",
//		`package pkg6
//
//		type Category int32
//
//		type Struct struct {
//			Cat      Category
//		}`,
//	)
//}

//func TestStructWithNestedSimpleTypeInOtherPackageOrder1(t *testing.T) {
//	doTestParseOK(t,
//		&Node{
//			Name: "Struct",
//			Type: Type{"pkg7", "Struct", Struct},
//			Nodes: []*Node{
//				{
//					Name: "Cat",
//					Type: Type{"other", "Category", Int32},
//					Tags: &Tag{},
//				},
//			},
//		},
//		"pkg7", "Struct",
//		`package pkg7
//
//		type Struct struct {
//			Cat      other.Category
//		}`,
//		`package other
//
//		type Category int32
//		`,
//	)
//}

//func TestStructWithNestedSimpleTypeInOtherPackageOrder2(t *testing.T) {
//	doTestParseOK(t,
//		&Node{
//			Name: "Struct",
//			Type: Type{"pkg8", "Struct", Struct},
//			Nodes: []*Node{
//				{
//					Name: "Cat",
//					Type: Type{"other", "Category", Int32},
//					Tags: &Tag{},
//				},
//			},
//		},
//		"pkg8", "Struct",
//		`package other
//
//		type Category int32
//		`,
//		`package pkg8
//
//		type Struct struct {
//			Cat      other.Category
//		}`,
//	)
//}

//func TestStructWithNestingAcross2Packages(t *testing.T) {
//	doTestParseOK(t,
//		&Node{
//			Name: "Struct",
//			Type: Type{"pkg9", "Struct", Struct},
//			Nodes: []*Node{
//				{
//					Name: "Id",
//					Type: Type{"", "uint32", Uint32},
//					Tags: &Tag{},
//				},
//				{
//					Name: "Wibble",
//					Type: Type{"stringy", "Thingy", String},
//					Tags: &Tag{},
//				},
//				{
//					Name: "Bobble",
//					Type: Type{"", "string", String},
//					Tags: &Tag{},
//				},
//			},
//		},
//		"pkg9", "Struct",
//		`package stringy
//
//		type Thingy string
//		`,
//		`package froob
//
//		type Inner1 struct {
//			Wibble stringy.Thingy
//		}
//		`,
//
//		`package pkg9
//
//		type Struct struct {
//			Id uint32
//			froob.Inner1
//			Bobble string
//		}
//		`,
//	)
//}

//func TestStructWithNestingInTheSamePackage(t *testing.T) {
//	doTestParseOK(t,
//		&Node{
//			Name: "Struct",
//			Type: Type{"pkg10", "Struct", Struct},
//			Nodes: []*Node{
//				{
//					Name: "Id",
//					Type: Type{"", "uint32", Uint32},
//					Tags: &Tag{},
//				},
//				{
//					Name: "Uid",
//					Type: Type{"", "uint32", Uint32},
//					Tags: &Tag{},
//				},
//				{
//					Name: "Name",
//					Type: Type{"pkg10", "Username", String},
//					Tags: &Tag{},
//				},
//				{
//					Name: "Wibble",
//					Type: Type{"pkg10", "Thingy", String},
//					Tags: &Tag{},
//				},
//				{
//					Name: "Bobble",
//					Type: Type{"pkg10", "Username", String},
//					Tags: &Tag{},
//				},
//			},
//		},
//		"pkg10", "Struct",
//		`package pkg10
//
//		type Thingy string
//
//		type Username string
//
//		type User struct {
//			Uid  uint32
//			Name Username
//		}
//
//		type UserWithThingy struct {
//			User
//			Wibble Thingy
//		}
//
//		type Struct struct {
//			Id uint32
//			UserWithThingy
//			Bobble Username
//		}
//		`,
//	)
//}

//func TestStructWithNestingAcross3Packages(t *testing.T) {
//	doTestParseOK(t,
//		&Node{
//			Name: "Struct",
//			Type: Type{"pkg11", "Struct", Struct},
//			Nodes: []*Node{
//				{
//					Name: "Id",
//					Type: Type{"", "uint32", Uint32},
//					Tags: &Tag{},
//				},
//				{
//					Name: "Uid",
//					Type: Type{"", "uint32", Uint32},
//					Tags: &Tag{},
//				},
//				{
//					Name: "Name",
//					Type: Type{"userindex", "Username", String},
//					Tags: &Tag{},
//				},
//				{
//					Name: "Wibble",
//					Type: Type{"stringy", "Thingy", String},
//					Tags: &Tag{},
//				},
//				{
//					Name: "Bobble",
//					Type: Type{"userindex", "Username", String},
//					Tags: &Tag{},
//				},
//			},
//		},
//		"pkg11", "Struct",
//		`package stringy
//
//		type Thingy string
//		`,
//
//		`package userindex
//
//		type Username string
//		`,
//
//		`package userindex
//
//		type User struct {
//			Uid  uint32
//			Name Username
//		}
//
//		type UserWithThingy struct {
//			User
//			Wibble stringy.Thingy
//		}
//		`,
//
//		`package pkg11
//
//		type Struct struct {
//			Id uint32
//			userindex.UserWithThingy
//			Bobble userindex.Username
//		}
//		`,
//	)
//}

//func doTestParseOK(t *testing.T, want *Node, pkg, name string, isource ...string) {
//	t.Helper()
//	exit.TestableExit()
//	Debug = true
//
//	// fix edges missing in the literal values
//	for _, n0 := range want.Nodes {
//		n0.Parent = want
//		for _, n1 := range n0.Nodes {
//			n1.Parent = n0
//		}
//	}
//
//	files := make([]Source, len(isource))
//
//	for i, s := range isource {
//		// allow nested back-ticks
//		source := strings.Replace(s, "|", "`", -1)
//		files[i] = Source{fmt.Sprintf("issue%d.go", i), bytes.NewBufferString(source)}
//	}
//
//	err := ParseGroups(files)
//	if err != nil {
//		t.Errorf("Error parsing: %s", err)
//	}
//
//	got, err := findMatchingNodes(pkg, name)
//	if err != nil {
//		t.Errorf("Error parsing: %s", err)
//	}
//
//	if !reflect.DeepEqual(got, want) {
//		ex := utter.Sdump(want)
//		ac := utter.Sdump(got)
//		outputDiff(ex, pkg+"-expected.txt")
//		outputDiff(ac, pkg+"-got.txt")
//		t.Errorf("Wanted %s\nGot %s", ex, ac)
//	}
//}

//-------------------------------------------------------------------------------------------------

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
