package parse

import (
	"bytes"
	"github.com/kortschak/utter"
	"reflect"
	"strings"
	"testing"
	"fmt"
	"github.com/rickb777/sqlgen/parse/exit"
)

func TestFindImport(t *testing.T) {
	exit.TestableExit()

	source1 := `package pkg7b

		import (
			"bytes"
			"github.com/kortschak/utter"
		)
		`

	source2 := `package pkg7a

		import (
			"bytes"
			"github.com/rickb777/sqlgen/schema"
		)
		`
	source3 := `package thingy

		import (
			"go/token"
			"github.com/rickb777/sqlgen/parse"
		)
		`

	files := make([]file, 0)

	files = append(files, file{"issue1.go", bytes.NewBufferString(source1)})
	files = append(files, file{"issue2.go", bytes.NewBufferString(source2)})
	files = append(files, file{"issue3.go", bytes.NewBufferString(source3)})

	err := parseAllFiles(files)
	if err != nil {
		t.Errorf("Error parsing: %s", err)
	}

	cases := []struct {
		shortName, expected string
	}{
		{"bytes", "bytes"},
		{"utter", "github.com/kortschak/utter"},
		{"schema", "github.com/rickb777/sqlgen/schema"},
		{"parse", "github.com/rickb777/sqlgen/parse"},
		{"token", "go/token"},
	}

	for _, c := range cases {
		i := FindImport(c.shortName)
		if i != c.expected {
			t.Errorf("%s -> expected %q but got %q", c.shortName, c.expected, i)
		}
	}
}

//-------------------------------------------------------------------------------------------------

func TestStructWith3FieldsAndTags(t *testing.T) {
	doTestParseOK(t,
		&Node{
			Name: "Struct",
			Type: Type{"pkg1", "Struct", Struct},
			Nodes: []*Node{
				{
					Name: "Id",
					Type: Type{"", "int64", Int64},
					Tags: &Tag{Primary: true, Auto: true},
				}, {
					Name: "Number",
					Type: Type{"", "int", Int},
					Tags: &Tag{},
				}, {
					Name: "Title",
					Type: Type{"", "string", String},
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
			Name: "Struct",
			Type: Type{"pkg2", "Struct", Struct},
			Nodes: []*Node{
				{
					Name: "Flag",
					Type: Type{"", "bool", Bool},
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
			Name: "Struct",
			Type: Type{"pkg3", "Struct", Struct},
			Nodes: []*Node{
				{
					Name: "Labels",
					Type: Type{"", "[]string", Slice},
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
			Name: "Struct",
			Type: Type{"pkg4", "Struct", Struct},
			Nodes: []*Node{
				{
					Name: "Table",
					Type: Type{"", "map[string]int", Map},
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
			Name: "Struct",
			Type: Type{"pkg5", "Struct", Struct},
			Nodes: []*Node{
				{
					Name: "Author",
					Type: Type{"", "Author", Ptr},
					Tags: &Tag{},
					Nodes: []*Node{
						{
							Name: "Name",
							Type: Type{"", "string", String},
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
	doTestParseOK(t,
		&Node{
			Name: "Struct",
			Type: Type{"pkg6", "Struct", Struct},
			Nodes: []*Node{
				{
					Name: "Cat",
					Type: Type{"pkg6", "Category", Int32},
					Tags: &Tag{},
				},
			},
		},
		"pkg6", "Struct",
		`package pkg6

		type Category int32

		type Struct struct {
			Cat      Category
		}`,
	)
}

func TestStructWithNestedSimpleTypeInOtherPackageOrder1(t *testing.T) {
	doTestParseOK(t,
		&Node{
			Name: "Struct",
			Type: Type{"pkg6", "Struct", Struct},
			Nodes: []*Node{
				{
					Name: "Cat",
					Type: Type{"other", "Category", Int32},
					Tags: &Tag{},
				},
			},
		},
		"pkg6", "Struct",
		`package pkg6

		type Struct struct {
			Cat      other.Category
		}`,
		`package other

		type Category int32
		`,
	)
}

func TestStructWithNestedSimpleTypeInOtherPackageOrder2(t *testing.T) {
	doTestParseOK(t,
		&Node{
			Name: "Struct",
			Type: Type{"pkg6", "Struct", Struct},
			Nodes: []*Node{
				{
					Name: "Cat",
					Type: Type{"other", "Category", Int32},
					Tags: &Tag{},
				},
			},
		},
		"pkg6", "Struct",
		`package other

		type Category int32
		`,
		`package pkg6

		type Struct struct {
			Cat      other.Category
		}`,
	)
}

func TestStructWithNestingAcrossPackages(t *testing.T) {
	doTestParseOK(t,
		&Node{
			Name: "Struct",
			Type: Type{"pkg7a", "Struct", Struct},
			Nodes: []*Node{
				{
					Name: "Id",
					Type: Type{"", "uint32", Uint32},
					Tags: &Tag{},
				},
				{
					Name: "Wibble",
					Type: Type{"stringy", "Thingy", String},
					Tags: &Tag{},
				},
				{
					Name: "Bibble",
					Type: Type{"", "string", String},
					Tags: &Tag{},
				},
				{
					Name: "Bobble",
					Type: Type{"", "string", String},
					Tags: &Tag{},
				},
			},
		},
		"pkg7a", "Struct",
		`package stringy

		type Thingy string
		`,
		`package pkg7b

		type Inner2 struct {
			Wibble stringy.Thingy
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

	err := parseAllFiles(files)
	if err != nil {
		t.Errorf("Error parsing: %s", err)
	}

	got, err := findMatchingNodes(pkg, name)
	if err != nil {
		t.Errorf("Error parsing: %s", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Wanted %s\nGot %s", utter.Sdump(want), utter.Sdump(got))
	}
}
