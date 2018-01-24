package code

import (
	"fmt"
	"io"
	. "strings"

	"github.com/rickb777/sqlgen2/schema"
	"bitbucket.org/pkg/inflect"
)

const sectionBreak = "\n//--------------------------------------------------------------------------------"

type ConstView struct {
	Name string
	Body interface{}
}

//const constString = "\nconst %s = %s\n"
const constStringD = "\nconst %s = %d\n"
const constStringQ = "\nconst %s = %q\n"
const constStringWithTicks = "\nconst %s = `\n%s`\n"

// WritePackage writes the Go package header to
// writer w with the given package name.
func WritePackage(w io.Writer, name string) {
	fmt.Fprintf(w, sPackage, name)
}

// writeSchema writes SQL statements to CREATE, INSERT,
// UPDATE and DELETE values from Table t.
func WriteSchema(w io.Writer, view View) {
	tableName := view.CamelName()

	fmt.Fprintln(w, sectionBreak)

	fmt.Fprintf(w, constStringD, "Num"+tableName+"Columns", view.Table.NumColumnNames(true))
	fmt.Fprintf(w, constStringD, "Num"+tableName+"DataColumns", view.Table.NumColumnNames(false))

	if view.Table.HasPrimaryKey() {
		fmt.Fprintf(w, constStringQ, tableName+"Pk", view.Table.Primary.Name)
	}

	fmt.Fprintf(w, constStringQ, tableName+"DataColumnNames", Join(view.Table.ColumnNames(false), ", "))

	fmt.Fprintln(w, sectionBreak)

	must(tCreateTableFunc.Execute(w, view))

	for _, d := range schema.AllDialects {
		ds := d.String()
		fmt.Fprintf(w, constStringWithTicks, "sqlCreateColumns"+tableName+view.Thing+ds, d.TableDDL(view.Table))
		fmt.Fprintf(w, constStringQ, "sqlCreateSettings"+tableName+view.Thing+ds, d.CreateTableSettings())
	}

	fmt.Fprintf(w, "\nconst sqlConstrain%s%s = `", tableName, view.Thing)
	for i, f := range view.Table.Fields {
		if f.Tags.ForeignKey != "" {
			slice := Split(f.Tags.ForeignKey, ".")
			fmt.Fprintf(w, "\n CONSTRAINT %sc%d foreign key (%s) references %%s%s (%s)", tableName, i, f.SqlName, slice[0], slice[1])
			if f.Tags.OnUpdate != "" {
				fmt.Fprintf(w, " on update %s", f.Tags.OnUpdate)
			}
			if f.Tags.OnDelete != "" {
				fmt.Fprintf(w, " on delete %s", f.Tags.OnDelete)
			}
		}
	}
	fmt.Fprintf(w, "\n`\n")

	if len(view.Table.Index) > 0 {
		fmt.Fprintln(w, sectionBreak)

		must(tCreateIndexesFunc.Execute(w, view))

		fmt.Fprintln(w, sectionBreak)

		for _, ix := range view.Table.Index {
			fmt.Fprintf(w, constStringQ,
				"sql"+view.Prefix+inflect.Camelize(ix.Name)+"IndexColumns", ix.Columns())
		}
	}

	fmt.Fprintln(w, sectionBreak)

	must(tTruncate.Execute(w, view))
}
