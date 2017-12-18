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
	sqlite := schema.New(schema.Sqlite)

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

	writeCreateTableFunc(w, view)

	for _, did := range schema.AllDialectIds {
		d := schema.New(did)
		ds := did.String()

		fmt.Fprintf(w, constStringWithTicks,
			identifier("sqlCreate", tableName, "Table"+ds),
			"CREATE TABLE %s%s%s ("+d.Table(view.Table, did)+"\n)"+d.CreateTableSettings())
	}

	if len(view.Table.Index) > 0 {
		writeCreateIndexesFunc(w, view)

		fmt.Fprintln(w, sectionBreak)

		for _, ix := range view.Table.Index {
			fmt.Fprintf(w, constStringWithTicks,
				identifier("sqlCreate"+view.Prefix, ix.Name, "Index"), sqlite.Index(view.Table, ix))
		}
	}
}

func writeCreateTableFunc(w io.Writer, view View) {
	fmt.Fprintln(w, sectionBreak)

	must(tCreateTableFunc.Execute(w, view))
}

func writeCreateIndexesFunc(w io.Writer, view View) {
	fmt.Fprintln(w, sectionBreak)

	for _, ix := range view.Table.Index {
		view.Body1 = append(view.Body1, inflect.Camelize(ix.Name))
	}

	must(tCreateIndexesFunc.Execute(w, view))
}

func identifier(prefix, id, suffix string) string {
	return prefix + inflect.Camelize(id) + suffix
}
