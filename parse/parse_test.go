package parse

import (
	"bytes"
	"github.com/kortschak/utter"
	"reflect"
	"strings"
	"testing"
	"fmt"
)


func TestStructWith3FieldsAndTags(t *testing.T) {
	doTestParseOK(t,
		&Node{
			Pkg:  "pkg1",
			Name: "Struct",
			Type: "Struct",
			Nodes: []*Node{
				{
					Name: "Id",
					Kind: Int64,
					Type: "int64",
					Tags: &Tag{Primary: true, Auto: true},
				}, {
					Name: "Number",
					Kind: Int,
					Type: "int",
					Tags: &Tag{},
				}, {
					Name: "Title",
					Kind: String,
					Type: "string",
					Tags: &Tag{},
				},
			},
		},
		"pkg1", "Struct",
		`package pkg1

		type Struct struct {
			Id       int64 |sql:"pk: true, auto: true"|
			Number   int
			Title    string
		}`,
	)
}

func TestStructWith1FieldAndIgnoreTag(t *testing.T) {
	doTestParseOK(t,
		&Node{
			Pkg:  "pkg2",
			Name: "Struct",
			Type: "Struct",
			Nodes: []*Node{
				{
					Name: "Flag",
					Kind: Bool,
					Type: "bool",
					Tags: &Tag{Skip: true},
				},
			},
		},
		"pkg2", "Struct",
		`package pkg2

		type Struct struct {
			Flag  bool     |sql:"-"|
		}`,
	)
}

func TestStructWith1FieldAndJsonTag(t *testing.T) {
	doTestParseOK(t,
		&Node{
			Pkg:  "pkg3",
			Name: "Struct",
			Type: "Struct",
			Nodes: []*Node{
				{
					Name: "Labels",
					Kind: Slice,
					Type: "[]string",
					Tags: &Tag{Encode: "json"},
				},
			},
		},
		"pkg3", "Struct",
		`package pkg3

		type Struct struct {
			Labels   []string  |sql:"encode: json"|
		}`,
	)
}

func TestStructWith1MapFieldAndJsonTag(t *testing.T) {
	doTestParseOK(t,
		&Node{
			Pkg:  "pkg4",
			Name: "Struct",
			Type: "Struct",
			Nodes: []*Node{
				{
					Name: "Table",
					Kind: Map,
					Type: "map[string]int",
					Tags: &Tag{Encode: "json"},
				},
			},
		},
		"pkg4", "Struct",
		`package pkg4

		type Struct struct {
			Table    map[string]int  |sql:"encode: json"|
		}`,
	)
}

func TestStructWithNesting(t *testing.T) {
	doTestParseOK(t,
		&Node{
			Pkg:  "pkg5",
			Name: "Struct",
			Type: "Struct",
			Nodes: []*Node{
				{
					Name: "Author",
					Kind: Ptr,
					Type: "Author",
					Tags: &Tag{},
					Nodes: []*Node{
						{
							Name: "Name",
							Kind: String,
							Type: "string",
							Tags: &Tag{},
						},
					},
				},
			},
		},
		"pkg5", "Struct",
		`package pkg5

		type Struct struct {
			Author    *Author
		}
		type Author struct {
			Name     string
		}`,
	)
}

func TestStructWithNestedSimpleType(t *testing.T) {
	// *** known bug ***
	//doTestParseOK(t, "pkg6", "Struct",
	//	`package pkg6
	//
	//	type Category int32
	//
	//	type Struct struct {
	//		Cat      Category
	//	}`,
	//	&Node{
	//		Pkg:  "pkg6",
	//		Name: "Struct",
	//		Type: "Struct",
	//		Nodes: []*Node{
	//			{
	//				Name: "Cat",
	//				Kind: Int32,
	//				Type: "Category",
	//				Tags: &Tag{},
	//			},
	//		},
	//	},
	//)
}

func TestStructWithNestingAcrossPackages(t *testing.T) {
	doTestParseOK(t,
		&Node{
			Pkg:  "pkg7a",
			Name: "Struct",
			Type: "Struct",
			Nodes: []*Node{
				{
					Name: "Id",
					Kind: Uint32,
					Type: "uint32",
					Tags: &Tag{},
				},
				{
					Name: "Wibble",
					Kind: String,
					Type: "string",
					Tags: &Tag{},
				},
				{
					Name: "Bibble",
					Kind: String,
					Type: "string",
					Tags: &Tag{},
				},
				{
					Name: "Bobble",
					Kind: String,
					Type: "string",
					Tags: &Tag{},
				},
			},
		},
		"pkg7a", "Struct",
		`package pkg7b

		type Inner2 struct {
			Wibble string
		}

		type Inner1 struct {
			Inner2
			Bibble string
		}
		`,

		`package pkg7a

		type Struct struct {
			Id uint32
			pkg7b.Inner1
			Bobble string
		}
		`,
	)
}

func doTestParseOK(t *testing.T, want *Node, pkg, name string, isource ...string) {
	t.Helper()

	// fix edges missing in the literal values
	for _, n0 := range want.Nodes {
		n0.Parent = want
		for _, n1 := range n0.Nodes {
			n1.Parent = n0
		}
	}

	files := make([]file, len(isource))

	for i, s := range isource {
		// allow nested back-ticks
		source := strings.Replace(s, "|", "`", -1)
		files[i] = file{fmt.Sprintf("issue%d.go", i), bytes.NewBufferString(source)}
	}

	got, err := parseAllFiles(pkg, name, files)

	if err != nil {
		t.Errorf("Error parsing: %s", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Wanted %s\nGot %s", utter.Sdump(want), utter.Sdump(got))
	}
}
