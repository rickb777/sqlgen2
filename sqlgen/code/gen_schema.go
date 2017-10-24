package code

import (
	"fmt"
	"io"
	. "strings"

	"bitbucket.org/pkg/inflect"
	"github.com/rickb777/sqlgen2/schema"
)

const sectionBreak = "\n//--------------------------------------------------------------------------------"

type ConstView struct {
	Name string
	Body interface{}
}

// writeSchema writes SQL statements to CREATE, INSERT,
// UPDATE and DELETE values from Table t.
func WriteSchema(w io.Writer, table *schema.Table, dids ...schema.DialectId) {

	fmt.Fprintln(w, sectionBreak)

	tableName := table.Type

	must(tConst.Execute(w, ConstView{
		identifier("Num", tableName, "Columns"),
		table.NumColumnNames(true),
	}))

	must(tConstStr.Execute(w, ConstView{
		identifier("", tableName, "ColumnNames"),
		Join(table.ColumnNames(true), ", "),
	}))

	must(tConstStr.Execute(w, ConstView{
		identifier("", tableName, "DataColumnNames"),
		Join(table.ColumnNames(false), ", "),
	}))

	for _, did := range dids {
		d := schema.New(did)
		ds := did.String()

		must(tConstStr.Execute(w, ConstView{
			identifier("", tableName, "ColumnParams"+ds),
			d.ColumnParams(table, true),
		}))

		must(tConstStr.Execute(w, ConstView{
			identifier("", tableName, "DataColumnParams"+ds),
			d.ColumnParams(table, false),
		}))

		must(tConstStr.Execute(w, ConstView{
			identifier("Create", tableName, "Stmt"+ds),
			d.Table(table) + d.CreateTableSettings(),
		}))

		must(tConstStr.Execute(w, ConstView{
			identifier("Insert", tableName, "Stmt"+ds),
			d.Insert(table),
		}))

		if table.HasPrimaryKey() {
			must(tConstStr.Execute(w, ConstView{
				identifier("Update", tableName, "ByPkStmt"+ds),
				d.Update(table, []*schema.Field{table.Primary}),
			}))

			must(tConstStr.Execute(w, ConstView{
				identifier("Delete", tableName, "ByPkStmt"+ds),
				d.Delete(table, []*schema.Field{table.Primary}),
			}))
		}

		fmt.Fprintln(w, sectionBreak)

		for _, ix := range table.Index {

			must(tConstStr.Execute(w, ConstView{
				identifier("Create", ix.Name, "Stmt"+ds),
				d.Index(table, ix),
			}))

			if !ix.Unique {

				//must(tConstWithTableName.Execute(w, ConstView{
				//	identifier("Select", ix.Name, "RangeStmt"),
				//	d.SelectRange(table, ix.Fields),
				//}))
				//
				//must(tConstWithTableName.Execute(w, ConstView{
				//	identifier("Select", ix.Name, "CountStmt"),
				//	d.SelectCount(table, ix.Fields),
				//}))

			} else {

				must(tConstStr.Execute(w, ConstView{
					identifier("Update", ix.Name, "Stmt"+ds),
					d.Update(table, ix.Fields),
				}))

				must(tConstStr.Execute(w, ConstView{
					identifier("Delete", ix.Name, "Stmt"+ds),
					d.Delete(table, ix.Fields),
				}))
			}
		}
	}
}

// WritePackage writes the Go package header to
// writer w with the given package name.
func WritePackage(w io.Writer, name string) {
	fmt.Fprintf(w, sPackage, name)
}

func identifier(prefix, id, suffix string) string {
	return prefix + inflect.Camelize(id) + suffix
}
