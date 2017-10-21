package main

import (
	"fmt"
	"io"
	. "strings"

	"bitbucket.org/pkg/inflect"
	"github.com/rickb777/sqlgen/schema"
	"github.com/rickb777/sqlgen/sqlgen/parse"
)

const sectionBreak = "\n//--------------------------------------------------------------------------------"

type ConstView struct {
	Name string
	Body interface{}
}

// writeSchema writes SQL statements to CREATE, INSERT,
// UPDATE and DELETE values from Table t.
func writeSchema(w io.Writer, d schema.Dialect, tree *parse.Node, t *schema.Table) {

	fmt.Fprintln(w, sectionBreak)

	tableName := t.Type

	must(tConst.Execute(w, ConstView{
		identifier("Num", tableName, "Columns"),
		t.NumColumnNames(true),
	}))

	must(tConstStr.Execute(w, ConstView{
		identifier("", tableName, "ColumnNames"),
		Join(t.ColumnNames(true), ", "),
	}))

	must(tConstStr.Execute(w, ConstView{
		identifier("", tableName, "DataColumnNames"),
		Join(t.ColumnNames(false), ", "),
	}))

	must(tConstStr.Execute(w, ConstView{
		identifier("", tableName, "ColumnParams"),
		d.ColumnParams(t, true),
	}))

	must(tConstStr.Execute(w, ConstView{
		identifier("", tableName, "DataColumnParams"),
		d.ColumnParams(t, false),
	}))

	must(tConstWithTableName.Execute(w, ConstView{
		identifier("Create", tableName, "Stmt"),
		d.Table(t),
	}))

	must(tConstWithTableName.Execute(w, ConstView{
		identifier("Insert", tableName, "Stmt"),
		d.Insert(t),
	}))

	if t.HasPrimaryKey() {
		must(tConstWithTableName.Execute(w, ConstView{
			identifier("Update", tableName, "ByPkStmt"),
			d.Update(t, []*schema.Field{t.Primary}),
		}))

		must(tConstWithTableName.Execute(w, ConstView{
			identifier("Delete", tableName, "ByPkStmt"),
			d.Delete(t, []*schema.Field{t.Primary}),
		}))
	}

	fmt.Fprintln(w, sectionBreak)

	for _, ix := range t.Index {

		must(tConstWithTableName.Execute(w, ConstView{
			identifier("Create", ix.Name, "Stmt"),
			d.Index(t, ix),
		}))

		if !ix.Unique {

			//must(tConstWithTableName.Execute(w, ConstView{
			//	identifier("Select", ix.Name, "RangeStmt"),
			//	d.SelectRange(t, ix.Fields),
			//}))
			//
			//must(tConstWithTableName.Execute(w, ConstView{
			//	identifier("Select", ix.Name, "CountStmt"),
			//	d.SelectCount(t, ix.Fields),
			//}))

		} else {

			must(tConstWithTableName.Execute(w, ConstView{
				identifier("Update", ix.Name, "Stmt"),
				d.Update(t, ix.Fields),
			}))

			must(tConstWithTableName.Execute(w, ConstView{
				identifier("Delete", ix.Name, "Stmt"),
				d.Delete(t, ix.Fields),
			}))
		}
	}
}

// WritePackage writes the Go package header to
// writer w with the given package name.
func writePackage(w io.Writer, name string) {
	fmt.Fprintf(w, sPackage, name)
}

func identifier(prefix, id, suffix string) string {
	return prefix + inflect.Camelize(id) + suffix
}
