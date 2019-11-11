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

func WriteSliceColumn(w1, w2 io.Writer, view View) {
	must(tSliceItemDecl.Execute(w1, view))
	must(tSliceItemFunc.Execute(w2, view))
}

func WriteInsertFunc(w1, w2 io.Writer, view View) {
	must(tInsertDecl.Execute(w1, view))
	must(tInsertFunc.Execute(w2, view))
}

func WriteUpsertFunc(w1, w2 io.Writer, view View) {
	if view.Table.HasPrimaryKey() {
		must(tUpsertDecl.Execute(w1, view))
		must(tUpsertFunc.Execute(w2, view))
	}
}

func WriteUpdateFunc(w1, w2 io.Writer, view View) {
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

func WriteDeleteFunc(w1, w2 io.Writer, view View) {
	fmt.Fprintln(w2, sectionBreak)
	must(tDeleteDecl.Execute(w1, view))
	must(tDeleteFunc.Execute(w2, view))
	fmt.Fprintln(w2, sectionBreak)
}

func WriteExecFunc(w1, w2 io.Writer, view View) {
	fmt.Fprintln(w2, sectionBreak)
	must(tExecDecl.Execute(w1, view))
	must(tExecFunc.Execute(w2, view))
}

// join is a helper function that joins nodes
// together by name using the seperator.
func must(err error) {
	output.Require(err == nil, "%v\n", err)
}
