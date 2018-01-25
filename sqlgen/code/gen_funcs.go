package code

import (
	"io"
	"github.com/rickb777/sqlgen2/sqlgen/output"
	"fmt"
)

func WriteType(w io.Writer, view View) {
	must(tTable.Execute(w, view))
}

func WriteQueryRows(w io.Writer, view View) {
	fmt.Fprintln(w, sectionBreak)

	must(tQueryRows.Execute(w, view))
}

func WriteQueryThings(w io.Writer, view View) {
	fmt.Fprintln(w, sectionBreak)

	must(tQueryThings.Execute(w, view))
}

func WriteGetRow(w io.Writer, view View) {
	must(tGetRow.Execute(w, view))
}

func WriteSelectRowsFuncs(w io.Writer, view View) {
	fmt.Fprintln(w, sectionBreak)

	must(tSelectRows.Execute(w, view))
	must(tCountRows.Execute(w, view))
}

func WriteSliceColumn(w io.Writer, view View) {
	must(tSliceItem.Execute(w, view))
}

func WriteInsertFunc(w io.Writer, view View) {
	must(tInsert.Execute(w, view))
}

func WriteUpdateFunc(w io.Writer, view View) {
	must(tUpdateFields.Execute(w, view))

	if view.Table.HasPrimaryKey() {
		must(tUpdate.Execute(w, view))

		//tableName := view.CamelName()
		//fmt.Fprintf(w, constStringWithTicks,
		//	"sqlUpdate"+tableName+"ByPkSimple", schema.Sqlite.UpdateDML(view.Table))
		//
		//fmt.Fprintf(w, constStringWithTicks,
		//	"sqlUpdate"+tableName+"ByPkPostgres", schema.Postgres.UpdateDML(view.Table))
	}
}

func WriteDeleteFunc(w io.Writer, view View) {
	fmt.Fprintln(w, sectionBreak)
	must(tDelete.Execute(w, view))
	fmt.Fprintln(w, sectionBreak)
}

func WriteExecFunc(w io.Writer, view View) {
	fmt.Fprintln(w, sectionBreak)
	must(tExec.Execute(w, view))
}

// join is a helper function that joins nodes
// together by name using the seperator.
func must(err error) {
	output.Require(err == nil, "%v\n", err)
}
