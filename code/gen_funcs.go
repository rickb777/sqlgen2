package code

import (
	"fmt"
	"github.com/rickb777/sqlgen2/output"
	"io"
)

func WriteType(w1, w2 io.Writer, view View) {
	must(tTabler.Execute(w1, view))
	must(tTable.Execute(w2, view))
}

func WriteQueryRows(w1, w2 io.Writer, view View) {
	fmt.Fprintln(w2, sectionBreak)

	must(tQueryRowsDecl.Execute(w1, view))
	must(tQueryRowsFunc.Execute(w2, view))
}

func WriteQueryThings(w1, w2 io.Writer, view View) {
	fmt.Fprintln(w2, sectionBreak)

	must(tQueryThingsDecl.Execute(w1, view))
	must(tQueryThingsFunc.Execute(w2, view))
}

func WriteGetRow(w1, w2 io.Writer, view View) {
	must(tGetRowDecl.Execute(w1, view))
	must(tGetRowFunc.Execute(w2, view))
}

func WriteSelectRowsFuncs(w1, w2 io.Writer, view View) {
	fmt.Fprintln(w2, sectionBreak)
	must(tSelectRowsDecl.Execute(w1, view))
	must(tSelectRowsFunc.Execute(w2, view))
}

func WriteCountRowsFuncs(w1, w2 io.Writer, view View) {
	fmt.Fprintln(w2, sectionBreak)
	must(tCountRowsDecl.Execute(w1, view))
	must(tCountRowsFunc.Execute(w2, view))
}

func WriteSliceColumn(w io.Writer, view View) {
	must(tSliceItem.Execute(w, view))
}

func WriteInsertFunc(w1, w2 io.Writer, view View) {
	must(tInsertDecl.Execute(w1, view))
	must(tInsertFunc.Execute(w2, view))
}

func WriteUpsertFunc(w io.Writer, view View) {
	if view.Table.HasPrimaryKey() {
		must(tUpsert.Execute(w, view))
	}
}

func WriteUpdateFunc(w1, w2 io.Writer, view View) {
	must(tUpdateFieldsDecl.Execute(w1, view))
	must(tUpdateFieldsFunc.Execute(w2, view))

	if view.Table.HasPrimaryKey() {
		must(tUpdateDecl.Execute(w1, view))
		must(tUpdateFunc.Execute(w2, view))

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
