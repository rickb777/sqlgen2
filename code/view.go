package code

import (
	"fmt"
	"github.com/rickb777/sqlapi/dialect"
	"strings"
	"text/template"

	"bitbucket.org/pkg/inflect"
	"github.com/acsellers/inflections"
	"github.com/rickb777/sqlapi/constraint"
	"github.com/rickb777/sqlapi/schema"
)

type View struct {
	Prefix     string
	Type       string
	Types      string
	DbName     string
	Thing      string
	Interface1 string
	Interface2 string
	List       string
	Suffix     string
	Body1      []string
	Body2      []string
	Body3      []string
	Dialects   []dialect.Dialect
	Table      *schema.TableDescription
	Setter     *schema.Field
}

func NewView(name, prefix, tableName, list string) View {
	if list == "" {
		list = fmt.Sprintf("[]*%s", name)
	}
	pl := inflections.Pluralize(name)
	tn := strings.ToLower(pl)
	if tableName != "" {
		tn = tableName
	}
	return View{
		Prefix:     prefix,
		Type:       name,
		Types:      pl,
		DbName:     tn,
		Thing:      "Table",
		Interface1: "sqlapi.Table",
		Interface2: "sqlapi.Table",
		List:       list,
		Dialects:   dialect.AllDialects,
	}
}

func (v View) CamelName() string {
	return v.Prefix + inflect.Camelize(v.Table.Type)
}

func (v View) Constraints() (list constraint.Constraints) {
	for _, f := range v.Table.Fields {
		c := constraint.FkConstraintOfField(f)
		if c.Parent.TableName != "" {
			list = append(list, c)
		}
	}
	return list
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
	"lc": func(s interface{}) string {
		return strings.ToLower(fmt.Sprintf("%s", s))
	},
	"title": func(s interface{}) string {
		return strings.Title(fmt.Sprintf("%s", s))
	},
}
