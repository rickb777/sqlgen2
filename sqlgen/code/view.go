package code

import (
	"fmt"
	. "strings"
	"text/template"

	. "github.com/acsellers/inflections"
	"github.com/rickb777/sqlgen2/schema"
	"bitbucket.org/pkg/inflect"
)

type View struct {
	Prefix   string
	Type     string
	Types    string
	List     string
	Suffix   string
	Body1    []string
	Body2    []string
	Body3    []string
	Dialects []string
	Table    *schema.TableDescription
}

func NewView(name, prefix, list string) View {
	if list == "" {
		list = fmt.Sprintf("[]*%s", name)
	}
	return View{
		Prefix:   prefix,
		Type:     name,
		Types:    Pluralize(name),
		List:     list,
		Dialects: schema.Dialects,
	}
}

func (v View) DbName() string {
	return ToLower(v.Types)
}

var funcMap = template.FuncMap{
	"q": func(s interface{}) string {
		return fmt.Sprintf("%q", s)
	},
	"camel": func(s interface{}) string {
		return inflect.Camelize(fmt.Sprintf("%s", s))
	},
	"ticked": func(s interface{}) string {
		return fmt.Sprintf("`\n%s\n`", s)
	},
}
