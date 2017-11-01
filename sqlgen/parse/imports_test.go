package parse

import (
	"bytes"
	"github.com/rickb777/sqlgen2/sqlgen/parse/exit"
	"reflect"
	"testing"
)

func TestBasicParsing(t *testing.T) {
	exit.TestableExit()

	source1 := `package pkg1

		import (
			. "bytes"
			spew "github.com/kortschak/utter"
		)

		type Number1 int32
		type NumberAlias = Number1

		const One Number1 = 1
		`

	source2 := `package pkg2

		import (
			"github.com/rickb777/sqlgen2/pkg1"
		)

		type Irrelevant interface {
			Foo()
		}

		type Interesting struct {
			I1   int
			Num1 pkg1.Number1
		}
		`

	source3 := `package pkg3

		import (
			"go/token"
			"github.com/rickb777/sqlgen2/pkg2"
		)

		type More struct {
			S1   string
			Num1 pkg1.Number1
		}
		`

	var sources []Source
	sources = append(sources, Source{"a/pkg1/source1.go", bytes.NewBufferString(source1)})
	sources = append(sources, Source{"a/pkg2/source2.go", bytes.NewBufferString(source2)})
	sources = append(sources, Source{"a/pkg3/source3.go", bytes.NewBufferString(source3)})

	err := parseAllFiles(sources)
	if err != nil {
		t.Errorf("Error parsing: %s", err)
	}

	if len(sources) != 3 {
		t.Errorf("expected 3 but got %d", len(sources))
	}

	cases := []struct {
		expSource string
		expPkg    string
		expImp    ImportList
	}{
		{"a/pkg1/source1.go", "pkg1", ImportList{Import{".", "bytes"}, Import{"spew", "github.com/kortschak/utter"}}},
		{"a/pkg2/source2.go", "pkg2", ImportList{Import{"", "github.com/rickb777/sqlgen2/pkg1"}}},
		{"a/pkg3/source3.go", "pkg3", ImportList{Import{"", "go/token"}, Import{"", "github.com/rickb777/sqlgen2/pkg2"}}},
	}

	for i, c := range cases {
		if files[i].FilePath != c.expSource {
			t.Errorf("%d: expected %q but got %q", i, c.expSource, files[i].FilePath)
		}
		if files[i].Pkg != c.expPkg {
			t.Errorf("%d: expected %q but got %q", i, c.expPkg, files[i].Pkg)
		}
		if !reflect.DeepEqual(files[i].ImportList, c.expImp) {
			t.Errorf("%d: expected %+v but got %+v", i, c.expImp, files[i].ImportList)
		}
	}
}
