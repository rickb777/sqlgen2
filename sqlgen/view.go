package main

import (
	"fmt"
	. "strings"
	"text/template"

	. "github.com/acsellers/inflections"
	"github.com/rickb777/sqlgen2/schema"
	"github.com/rickb777/sqlgen2/sqlgen/parse"
)

type View struct {
	Prefix string
	Type   string
	Types  string
	Suffix string
	Body1  []string
	Body2  []string
	Body3  []string
	Table  *schema.Table
}

func newView(tree *parse.Node, prefix string) View {
	return View{
		Prefix: prefix,
		Type:   tree.Type.Name,
		Types:  Pluralize(tree.Type.Name),
	}
}

func (v View) DbName() string {
	return ToLower(v.Types)
}

var funcMap = template.FuncMap{
	"q": func(s interface{}) string {
		return fmt.Sprintf("%q", s)
	},
	"ticked": func(s interface{}) string {
		return fmt.Sprintf("`\n%s\n`", s)
	},
}
