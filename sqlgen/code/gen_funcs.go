package code

import (
	"io"
	"github.com/rickb777/sqlgen2/schema"
	"github.com/rickb777/sqlgen2/sqlgen/output"
)

func WriteType(w io.Writer, view View) {
	must(tTable.Execute(w, view))
}

func WriteSelectRow(w io.Writer, view View) {
	must(tSelectRow.Execute(w, view))
}

func WriteSelectRows(w io.Writer, view View) {
	must(tSelectRows.Execute(w, view))
}

func WriteCountRows(w io.Writer, view View) {
	must(tCountRows.Execute(w, view))
}

func WriteInsertFunc(w io.Writer, view View, table *schema.Table) {
	if table.HasLastInsertId() {
		must(tInsertAndGetLastId.Execute(w, view))
	} else {
		must(tInsert.Execute(w, view))
	}
}

func WriteUpdateFunc(w io.Writer, view View, table *schema.Table) {
	if table.HasPrimaryKey() {
		must(tUpdate.Execute(w, view))
	}
}

func WriteExecFunc(w io.Writer, view View, table *schema.Table) {
	if table.HasPrimaryKey() {
		must(tExec.Execute(w, view))
	}
}

func WriteCreateTableFunc(w io.Writer, view View, table *schema.Table) {
	if table.HasPrimaryKey() {
		must(tCreateTable.Execute(w, view))
	}
}

// join is a helper function that joins nodes
// together by name using the seperator.
func must(err error) {
	output.Require(err == nil, "%v\n", err)
}
