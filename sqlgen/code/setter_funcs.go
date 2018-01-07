package code

import (
	"io"
	"fmt"
)

func WriteSetters(w io.Writer, view View, genSetters string) {
	fmt.Fprintln(w, sectionBreak)

	opt := setterOptionMap[genSetters]
	if opt == none {
		return
	}

	for _, field := range view.Table.Fields {
		if field.Tags.Skip && opt < all {
			continue
		} else if !isExported(field.Name) && opt < all {
			continue
		} else if !field.Type.IsPtr && opt < exported {
			continue
		} else {
			view.Setter = field
			must(tSetter.Execute(w, view))
		}
	}
}

func isExported(name string) bool {
	return len(name) > 0 && 'A' <= name[0] && name[0] <= 'Z'
}

type setterOption int

const (
	none     setterOption = iota
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
