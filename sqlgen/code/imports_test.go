package code

//import (
//	"bytes"
//	"testing"
//	"github.com/rickb777/sqlgen2/schema"
//	"github.com/rickb777/sqlgen2/sqlgen/parse/exit"
//	"strings"
//	"github.com/rickb777/sqlgen2/sqlgen/parse"
//)
//
//func TestWriteImports_withoutAny(t *testing.T) {
//	buf := &bytes.Buffer{}
//
//	WriteImports(buf, &schema.TableDescription{}, NewStringSet())
//
//	code := buf.String()
//	if code != `` {
//		t.Errorf("got\n%s\n", code)
//	}
//}
//
//func TestWriteImports_withoutExtras(t *testing.T) {
//	buf := &bytes.Buffer{}
//
//	WriteImports(buf, &schema.TableDescription{}, NewStringSet("foo", "bar"))
//
//	code := buf.String()
//	expected := `
//import (
//	"bar"
//	"foo"
//)
//`
//	if code != expected {
//		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
//	}
//}
//
//func TestWriteImports_withExtras(t *testing.T) {
//	exit.TestableExit()
//	literal := strings.Replace(`package pkg1
//
//type Example struct {
//	Id         int64    |sql:"pk: true, auto: true"|
//	Labels     []string |sql:"encode: json"|
//}
//
//`, "|", "`", -1)
//
//	source := parse.Source{"issue.go", bytes.NewBufferString(literal)}
//
//	node, err := parse.DoParse("pkg1", "Example", []parse.Source{source})
//	if err != nil {
//		t.Errorf("Error parsing: %s", err)
//	}
//
//	table := schema.Load(node)
//
//	buf := &bytes.Buffer{}
//
//	WriteImports(buf, table, NewStringSet("foo", "bar"))
//
//	code := buf.String()
//	expected := `
//import (
//	"bar"
//	"encoding/json"
//	"foo"
//)
//`
//	if code != expected {
//		t.Errorf("expected | got\n%s\n", sideBySideDiff(expected, code))
//	}
//}
