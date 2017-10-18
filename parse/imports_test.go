package parse

import (
	"bytes"
	"testing"
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
		tp := Type{Pkg: c.shortName}
		i := FindImport(tp)
		if i != c.expected {
			t.Errorf("%s -> expected %q but got %q", c.shortName, c.expected, i)
		}
	}
}
