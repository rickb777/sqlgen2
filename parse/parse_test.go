package parse

import (
	"bytes"
	"github.com/kortschak/utter"
	"reflect"
	"strings"
	"testing"
)

func TestParseOK(t *testing.T) {
	parseTests := []struct {
		pkg, name, source string
		want              *Node
	}{
		{
			pkg: "demo", name: "Issue1",
			source: `
package demo

type Issue1 struct {
	Id       int64 |sql:"pk: true, auto: true"|
	Number   int
	Title    string
}
`,
			want: &Node{
				Pkg:  "demo",
				Name: "Issue1",
				Type: "Issue1",
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
		},
		{
			pkg: "demo", name: "Issue2",
			source: `
package demo

type Category int32

type Issue2 struct {
	//Cat      Category
	Labels   []string |sql:"encode: json"|
	Locked   bool     |sql:"-"|
}
`,
			want: &Node{
				Pkg:  "demo",
				Name: "Issue2",
				Type: "Issue2",
				Nodes: []*Node{
					{
						//Name: "Cat",
						//Kind: Int32,
						//Type: "Category",
						//Tags: &Tag{},
						//},{
						Name: "Labels",
						Kind: Slice,
						Type: "[]string",
						Tags: &Tag{Encode: "json"},
					}, {
						Name: "Locked",
						Kind: Bool,
						Type: "bool",
						Tags: &Tag{Skip: true},
					},
				},
			},
		},
	}

	for i, test := range parseTests {
		for _, n := range test.want.Nodes {
			n.Parent = test.want
		}
		source := strings.Replace(test.source, "|", "`", -1)
		f1 := file{"issue.go", bytes.NewBufferString(source)}
		got, err := parseAll(test.pkg, test.name, []file{f1})

		if err != nil {
			t.Errorf("%d: Error parsing: %s", i, err)
		}

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Wanted %s\nGot %s", utter.Sdump(test.want), utter.Sdump(got))
		}
	}
}
