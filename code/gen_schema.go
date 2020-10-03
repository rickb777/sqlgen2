package code

import (
	"fmt"
	"github.com/markbates/inflect"
	"github.com/rickb777/sqlapi"
	"github.com/rickb777/sqlapi/dialect"
	"io"
)

const sectionBreak = "\n//-------------------------------------------------------------------------------------------------"

type ConstView struct {
	Name string
	Body interface{}
}

func WritePackageHeader(w io.Writer, name, version string) {
	fmt.Fprintf(w, sPackage, sqlapi.Version, version, name)
}

func WritePrimaryDeclarations(w io.Writer, view View) {
	tableName := view.CamelName()
	fullName := tableName + view.Thing

	fmt.Fprintln(w, sectionBreak)

	fmt.Fprintf(w, "\n// Num%sColumns is the total number of columns in %s.\n", fullName, fullName)
	fmt.Fprintf(w, "const Num%sColumns = %d\n", fullName, view.Table.NumColumnNames(true))

	fmt.Fprintf(w, "\n// Num%sDataColumns is the number of columns in %s not including the auto-increment key.\n", fullName, fullName)
	fmt.Fprintf(w, "const Num%sDataColumns = %d\n", fullName, view.Table.NumColumnNames(false))

	fmt.Fprintf(w, "\n// %sColumnNames is the list of columns in %s.\n", fullName, fullName)
	fmt.Fprintf(w, "const %sColumnNames = %q\n", fullName, view.Table.ColumnNames(true).MkString(","))
	if view.Table.HasPrimaryKey() {
		fmt.Fprintf(w, "\n// %sDataColumnNames is the list of data columns in %s.\n", fullName, fullName)
		fmt.Fprintf(w, "const %sDataColumnNames = %q\n", fullName, view.Table.ColumnNames(false).MkString(","))
	}

	fmt.Fprintf(w, "\nvar listOf%sColumnNames = strings.Split(%sColumnNames, \",\")\n", fullName, fullName)
}

func WriteSchemaDeclarations(w io.Writer, view View) {
	tableName := view.CamelName()
	fullName := tableName + view.Thing

	fmt.Fprintln(w, sectionBreak)

	for _, d := range dialect.AllDialects {
		ds := d.Name()
		fmt.Fprintf(w, "\nvar sql%sCreateColumns%s = []string{\n", fullName, ds)
		for _, field := range view.Table.Fields {
			fmt.Fprintf(w, "\t%q,\n", d.FieldAsColumn(field))
		}
		fmt.Fprintln(w, "}")
	}

	if len(view.Table.Index) > 0 {
		// write constants that specify the columns needed in each index

		fmt.Fprintln(w, sectionBreak+"\n")
		before := ""
		for _, ix := range view.Table.Index {
			cols := ix.Fields.SqlNames()
			colsStr := cols.MkString(",")
			name := view.Prefix + inflect.Camelize(ix.Name)
			fmt.Fprintf(w, "%sconst sql%sIndexColumns = %q\n", before, name, colsStr)
			if len(cols) == 1 {
				fmt.Fprintf(w, "\nvar listOf%sIndexColumns = []string{%q}\n", name, cols[0])
			} else {
				fmt.Fprintf(w, "\nvar listOf%sIndexColumns = strings.Split(sql%sIndexColumns, \",\")\n", name, name)
			}
			before = "\n"
		}
	}
}

func WriteSchemaFunctions(w1, w2 io.Writer, view View) {
	fmt.Fprintln(w2, sectionBreak)

	must(tCreateTableDecl.Execute(w1, view))
	must(tCreateTableFunc.Execute(w2, view))

	if len(view.Table.Index) > 0 {
		fmt.Fprintln(w2, sectionBreak)
		must(tCreateIndexesDecl.Execute(w1, view))
		must(tCreateIndexesFunc.Execute(w2, view))
	}

	fmt.Fprintln(w2, sectionBreak)
	must(tTruncateDecl.Execute(w1, view))
	must(tTruncateFunc.Execute(w2, view))
}
