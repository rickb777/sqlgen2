package code

import (
	"fmt"
	"github.com/rickb777/sqlapi/schema"
	"io"
)

type SetterView struct {
	TypePkg string
	Type    string
	Setter  *schema.Field
}

func (v View) FilterSetters(genSetters string) schema.FieldList {
	opt := setterOptionMap[genSetters]
	if opt == none {
		return nil
	}

	var list schema.FieldList

	for _, field := range v.Table.Fields {
		if field.GetTags().Skip && opt < all {
			continue
		} else if !isExported(field.Name) && opt < all {
			continue
		} else if !field.Type.IsPtr && opt < exported {
			continue
		} else {
			list = append(list, field)
		}
	}

	return list
}

func WriteSetters(w1, w2 io.Writer, view View, fields schema.FieldList) {
	if len(fields) > 0 {
		fmt.Fprintln(w2, sectionBreak)

		vm := SetterView{
			TypePkg: view.TypePkg,
			Type:    view.Type,
		}

		for _, field := range fields {
			vm.Setter = field
			must(tSetterDecl.Execute(w1, vm))
			must(tSetterFunc.Execute(w2, vm))
		}
	}
}

func isExported(name string) bool {
	return len(name) > 0 && 'A' <= name[0] && name[0] <= 'Z'
}

//-------------------------------------------------------------------------------------------------

type setterOption int

const (
	none setterOption = iota
	optional
	exported
	all
)

var setterOptionMap = map[string]setterOption{
	"none":     none,
	"option":   optional,
	"optional": optional,
	"exported": exported,
	"all":      all,
}
