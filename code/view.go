package code

import (
	"fmt"
	"strings"
	"text/template"

	"bitbucket.org/pkg/inflect"
	"github.com/acsellers/inflections"
	"github.com/rickb777/sqlapi/constraint"
	"github.com/rickb777/sqlapi/dialect"
	"github.com/rickb777/sqlapi/schema"
)

type View struct {
	TypePkg    string
	TablePkg   string
	Prefix     string
	Type       string
	Types      string
	DbName     string
	Thing      string
	Thinger    string
	Interface1 string
	List       string
	Suffix     string
	Sql        string
	Sqlapi     string
	Scan       string
	Body1      []string
	Body2      []string
	Body3      []string
	Dialects   []dialect.Dialect
	Table      *schema.TableDescription
	Setter     *schema.Field
}

func NewView(typePkg, tablePkg, name, prefix, tableName, list, sql, api string) View {
	if list == "" {
		list = fmt.Sprintf("[]*%s%s", typePkg, name)
	}
	pl := inflections.Pluralize(name)
	tn := strings.ToLower(pl)
	if tableName != "" {
		tn = tableName
	}
	return View{
		TypePkg:    typePkg,
		TablePkg:   tablePkg,
		Prefix:     prefix,
		Type:       name,
		Types:      pl,
		DbName:     tn,
		Thing:      "Table",
		Thinger:    "Tabler",
		Interface1: api + ".Table",
		List:       list,
		Dialects:   dialect.AllDialects,
		Sql:        sql,
		Sqlapi:     api,
		Scan:       "scan",
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
