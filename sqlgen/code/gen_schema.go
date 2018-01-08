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
	tableName := view.Prefix + view.Table.Type

	fmt.Fprintln(w, sectionBreak)

	fmt.Fprintf(w, "\nconst %s = %d\n",
		identifier("Num", tableName, "Columns"), view.Table.NumColumnNames(true))

	fmt.Fprintf(w, "\nconst %s = %d\n",
		identifier("Num", tableName, "DataColumns"), view.Table.NumColumnNames(false))

	if view.Table.HasPrimaryKey() {
		fmt.Fprintf(w, constStringQ,
			identifier("", tableName, "Pk"), view.Table.Primary.Name)
	}

	fmt.Fprintf(w, constStringQ,
		identifier("", tableName, "DataColumnNames"), Join(view.Table.ColumnNames(false), ", "))

	fmt.Fprintln(w, sectionBreak)

	must(tCreateTableFunc.Execute(w, view))

	for _, d := range schema.AllDialects {
		ds := d.String()

		fmt.Fprintf(w, constStringWithTicks,
			identifier("sqlCreate", tableName, view.Thing+ds),
			"CREATE TABLE %s%s%s ("+d.TableDDL(view.Table)+"\n)"+d.CreateTableSettings())
	}

	if len(view.Table.Index) > 0 {
		fmt.Fprintln(w, sectionBreak)

		must(tCreateIndexesFunc.Execute(w, view))

		fmt.Fprintln(w, sectionBreak)

		for _, ix := range view.Table.Index {
			fmt.Fprintf(w, constStringQ,
				identifier("sql"+view.Prefix, ix.Name, "IndexColumns"), ix.Columns())
		}
	}
}

func identifier(prefix, id, suffix string) string {
	return prefix + inflect.Camelize(id) + suffix
}
