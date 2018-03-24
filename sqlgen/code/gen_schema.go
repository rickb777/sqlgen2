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
const constString = "\nconst %s ="
const constStringD = "\nconst %s = %d\n"
const constStringQ = "\nconst %s = %q\n"
const constStringT = "\nconst %s = `%s`\n"
const constStringWithTicks = "\nconst %s = `\n%s`\n"

func WritePackageHeader(w io.Writer, name string) {
	fmt.Fprintf(w, sPackage, name)
}

func WritePrimaryDeclarations(w io.Writer, view View) {
	tableName := view.CamelName()

	fmt.Fprintln(w, sectionBreak)

	fmt.Fprintf(w, constStringD, "Num"+tableName+"Columns", view.Table.NumColumnNames(true))
	fmt.Fprintf(w, constStringD, "Num"+tableName+"DataColumns", view.Table.NumColumnNames(false))

	fmt.Fprintf(w, constStringQ, tableName+"ColumnNames", view.Table.ColumnNames(true).MkString(","))
	if view.Table.HasPrimaryKey() {
		fmt.Fprintf(w, constStringQ, tableName+"DataColumnNames", view.Table.ColumnNames(false).MkString(","))
	}
}

func WriteSchemaDeclarations(w io.Writer, view View) {
	tableName := view.CamelName()

	fmt.Fprintln(w, sectionBreak)

	for _, d := range schema.AllDialects {
		ds := d.String()
		fmt.Fprintf(w, constString, "sql"+tableName+view.Thing+"CreateColumns"+ds)
		fmt.Fprintln(w, d.TableDDL(view.Table))
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

		for _, ix := range view.Table.Index {
			cols := ix.Columns()
			fmt.Fprintf(w, constStringQ, "sql"+view.Prefix+inflect.Camelize(ix.Name)+"IndexColumns", cols)
		}
	}
}

func WriteSchemaFunctions(w io.Writer, view View) {
	fmt.Fprintln(w, sectionBreak)

	must(tCreateTableFunc.Execute(w, view))

	if len(view.Table.Index) > 0 {
		fmt.Fprintln(w, sectionBreak)
		must(tCreateIndexesFunc.Execute(w, view))
	}

	fmt.Fprintln(w, sectionBreak)
	must(tTruncate.Execute(w, view))
}
