package code

import (
	"bitbucket.org/pkg/inflect"
	"fmt"
	"github.com/rickb777/sqlapi"
	"github.com/rickb777/sqlapi/dialect"
	"io"
)

const sectionBreak = "\n//--------------------------------------------------------------------------------"

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
		ds := d.String()
		fmt.Fprintf(w, "\nvar sql%sCreateColumns%s = []string{\n", fullName, ds)
		for _, field := range view.Table.Fields {
			fmt.Fprintf(w, "\t%q,\n", d.FieldAsColumn(field))
		}
		fmt.Fprintln(w, "}")
	}

	//fmt.Fprintf(w, "\nconst sqlConstrain%s%s = `", tableName, view.Thing)
	//for i, f := range view.Table.Fields {
	//	if f.GetTags().ForeignKey != "" {
	//		slice := Split(f.Tags.ForeignKey, ".")
	//		fmt.Fprintf(w, "\n CONSTRAINT %sc%d foreign key (%s) references %%s%s (%s)", tableName, i, f.SqlName, slice[0], slice[1])
	//		if f.Tags.OnUpdate != "" {
	//			fmt.Fprintf(w, " on update %s", f.Tags.OnUpdate)
	//		}
	//		if f.Tags.OnDelete != "" {
	//			fmt.Fprintf(w, " on delete %s", f.Tags.OnDelete)
	//		}
	//	}
	//}
	//fmt.Fprintf(w, "\n`\n")

	if len(view.Table.Index) > 0 {
		// write constants that specify the columns needed in each index

		fmt.Fprintln(w, sectionBreak)
		for _, ix := range view.Table.Index {
			cols := ix.Columns()
			fmt.Fprintf(w, "const sql%sIndexColumns = %q\n", view.Prefix+inflect.Camelize(ix.Name), cols)
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
