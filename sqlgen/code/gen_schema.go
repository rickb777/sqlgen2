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
const constStringQ = "\nconst %s = %q\n"
const constStringWithTicks = "\nconst %s = `\n%s\n`\n"

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

	fmt.Fprintf(w, "\nconst %s = %d\n",
		"Num"+tableName+"Columns", view.Table.NumColumnNames(true))

	fmt.Fprintf(w, "\nconst %s = %d\n",
		"Num"+tableName+"DataColumns", view.Table.NumColumnNames(false))

	if view.Table.HasPrimaryKey() {
		fmt.Fprintf(w, constStringQ,
			tableName+"Pk", view.Table.Primary.Name)
	}

	fmt.Fprintf(w, constStringQ,
		tableName+"DataColumnNames", Join(view.Table.ColumnNames(false), ", "))

	fmt.Fprintln(w, sectionBreak)

	must(tCreateTableFunc.Execute(w, view))

	for _, d := range schema.AllDialects {
		ds := d.String()

		fmt.Fprintf(w, constStringWithTicks,
			"sqlCreate"+tableName+view.Thing+ds,
			"CREATE TABLE %s%s ("+d.TableDDL(view.Table)+"\n%s)"+d.CreateTableSettings())
	}

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
