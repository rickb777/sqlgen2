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
func WriteSchema(w io.Writer, table *schema.Table) {

	fmt.Fprintln(w, sectionBreak)

	tableName := table.Type

	must(tConst.Execute(w, ConstView{
		identifier("Num", tableName, "Columns"), table.NumColumnNames(true),
	}))

	must(tConstStr.Execute(w, ConstView{
		identifier("s", tableName, "ColumnNames"), Join(table.ColumnNames(true), ", "),
	}))

	must(tConstStr.Execute(w, ConstView{
		identifier("s", tableName, "DataColumnNames"), Join(table.ColumnNames(false), ", "),
	}))

	for _, did := range schema.AllDialectIds {
		d := schema.New(did)
		ds := did.String()

		//must(tConstStr.Execute(w, ConstView{identifier("", tableName, "ColumnParams"+ds), d.ColumnParams(table, true)}))

		//must(tConstStr.Execute(w, ConstView{identifier("", tableName, "DataColumnParams"+ds), d.ColumnParams(table, false)}))

		must(tConstStr.Execute(w, ConstView{
			identifier("sCreate", tableName, "Stmt"+ds), d.Table(table) + d.CreateTableSettings(),
		}))
	}

	for _, did := range schema.AllDialectIds {
		d := schema.New(did)
		ds := did.String()

		fmt.Fprintln(w, sectionBreak)

		if did == schema.Sqlite {
			did2 := schema.Mysql
			//d2 := schema.New(did2)
			ds2 := did2.String()

			must(tConst.Execute(w, ConstView{
				identifier("sInsert", tableName, "Stmt"+ds), identifier("sInsert", tableName, "Stmt"+ds2),
			}))

			if table.HasPrimaryKey() {
				must(tConst.Execute(w, ConstView{
					identifier("sUpdate", tableName, "ByPkStmt"+ds), identifier("sUpdate", tableName, "ByPkStmt"+ds2),
				}))

				must(tConst.Execute(w, ConstView{
					identifier("sDelete", tableName, "ByPkStmt"+ds), identifier("sDelete", tableName, "ByPkStmt"+ds2),
				}))
			}

		} else {
			must(tConstStr.Execute(w, ConstView{
				identifier("sInsert", tableName, "Stmt"+ds), d.Insert(table),
			}))

			if table.HasPrimaryKey() {
				must(tConstStr.Execute(w, ConstView{
					identifier("sUpdate", tableName, "ByPkStmt"+ds), d.Update(table, []*schema.Field{table.Primary}),
				}))

				must(tConstStr.Execute(w, ConstView{
					identifier("sDelete", tableName, "ByPkStmt"+ds), d.Delete(table, []*schema.Field{table.Primary}),
				}))
			}
		}
	}

	for _, did := range schema.AllDialectIds {
		d := schema.New(did)
		ds := did.String()

		for _, ix := range table.Index {

			fmt.Fprintln(w, sectionBreak)

			must(tConstStr.Execute(w, ConstView{
				identifier("sCreate", ix.Name, "Stmt"+ds), d.Index(table, ix),
			}))

			if !ix.Unique {

				//must(tConstWithTableName.Execute(w, ConstView{
				//	identifier("sSelect", ix.Name, "RangeStmt"), d.SelectRange(table, ix.Fields),
				//}))
				//
				//must(tConstWithTableName.Execute(w, ConstView{
				//	identifier("sSelect", ix.Name, "CountStmt"), d.SelectCount(table, ix.Fields),
				//}))

			} else {

				must(tConstStr.Execute(w, ConstView{
					identifier("sUpdate", ix.Name, "Stmt"+ds), d.Update(table, ix.Fields),
				}))

				must(tConstStr.Execute(w, ConstView{
					identifier("sDelete", ix.Name, "Stmt"+ds), d.Delete(table, ix.Fields),
				}))
			}
		}
	}

	fmt.Fprintln(w, sectionBreak)
}

// WritePackage writes the Go package header to
// writer w with the given package name.
func WritePackage(w io.Writer, name string) {
	fmt.Fprintf(w, sPackage, name)
}

func identifier(prefix, id, suffix string) string {
	return prefix + inflect.Camelize(id) + suffix
}
