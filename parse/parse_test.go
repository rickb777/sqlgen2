package parse

import (
	"bytes"
	"github.com/kortschak/utter"
	"reflect"
	"strings"
	"testing"
)

var i = 1

func TestParseOK(t *testing.T) {
	i = 1

	doTestParseOK(t, "pkg1", "Struct",
		`package pkg1

		type Struct struct {
			Id       int64 |sql:"pk: true, auto: true"|
			Number   int
			Title    string
		}`,
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
	)

	doTestParseOK(t, "pkg2", "Struct",
		`package pkg2

		type Struct struct {
			Flag  bool     |sql:"-"|
		}`,
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
	)

	doTestParseOK(t, "pkg3", "Struct",
		`package pkg3

		type Struct struct {
			Labels   []string  |sql:"encode: json"|
		}`,
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
	)

	doTestParseOK(t, "pkg4", "Struct",
		`package pkg4

		type Struct struct {
			Table    map[string]int  |sql:"encode: json"|
		}`,
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
	)

	doTestParseOK(t, "pkg5", "Struct",
		`package pkg5

		type Struct struct {
			Author    *Author
		}
		type Author struct {
			Name     string
		}`,
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
	)

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

func doTestParseOK(t *testing.T, pkg, name, isource string, want *Node) {
	// fix edges missing in the literal values
	for _, n0 := range want.Nodes {
		n0.Parent = want
		for _, n1 := range n0.Nodes {
			n1.Parent = n0
		}
	}

	// allow nested back-ticks
	source := strings.Replace(isource, "|", "`", -1)

	f1 := file{"issue.go", bytes.NewBufferString(source)}

	got, err := parseAll(pkg, name, []file{f1})

	if err != nil {
		t.Errorf("%d: Error parsing: %s", i, err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("%d: Wanted %s\nGot %s", i, utter.Sdump(want), utter.Sdump(got))
	}
	i++
}
