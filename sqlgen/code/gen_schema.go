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
func WriteSchema(w io.Writer, view View, table *schema.TableDescription) {
	sqlite := schema.New(schema.Sqlite)

	fmt.Fprintln(w, sectionBreak)

	tableName := view.Prefix + table.Type

	for _, did := range schema.AllDialectIds {
		d := schema.New(did)
		ds := did.String()

		fmt.Fprintf(w, constStringWithTicks,
			identifier("sqlCreate", tableName, "Table"+ds),
			"CREATE TABLE %s%s%s ("+d.Table(table, did)+"\n)"+d.CreateTableSettings())
	}

	if len(table.Index) > 0 {
		fmt.Fprintln(w, sectionBreak)
	}

	for _, ix := range table.Index {
		fmt.Fprintf(w, constStringWithTicks,
			identifier("sqlCreate"+view.Prefix, ix.Name, "Index"), sqlite.Index(table, ix))
	}

	fmt.Fprintln(w, sectionBreak)

	fmt.Fprintf(w, "\nconst %s = %d\n",
		identifier("Num", tableName, "Columns"), table.NumColumnNames(true))

	fmt.Fprintf(w, "\nconst %s = %d\n",
		identifier("Num", tableName, "DataColumns"), table.NumColumnNames(false))

	if table.HasPrimaryKey() {
		fmt.Fprintf(w, constStringQ,
			identifier("", tableName, "Pk"), table.Primary.Name)
	}

	fmt.Fprintf(w, constStringQ,
		identifier("", tableName, "DataColumnNames"), Join(table.ColumnNames(false), ", "))

	fmt.Fprintln(w, sectionBreak)
}

func WriteCreateTableAndIndexesFunc(w io.Writer, view View, table *schema.TableDescription) {
	fmt.Fprintln(w, sectionBreak)

	for _, ix := range table.Index {
		view.Body1 = append(view.Body1, inflect.Camelize(ix.Name))
	}

	must(tCreateTableAndIndexes.Execute(w, view))
}

func identifier(prefix, id, suffix string) string {
	return prefix + inflect.Camelize(id) + suffix
}
