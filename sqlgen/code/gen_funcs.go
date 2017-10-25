package code

import (
	"io"
	"github.com/rickb777/sqlgen2/schema"
	"github.com/rickb777/sqlgen2/sqlgen/output"
	"bitbucket.org/pkg/inflect"
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
		must(tInsertSimple.Execute(w, view))
	}
}

func WriteUpdateFunc(w io.Writer, view View, table *schema.Table) {
	if table.HasPrimaryKey() {
		must(tUpdate.Execute(w, view))
	}
}

func WriteExecFunc(w io.Writer, view View, table *schema.Table) {
	must(tExec.Execute(w, view))
}

func WriteCreateTableFunc(w io.Writer, view View, table *schema.Table) {
	must(tCreateTable.Execute(w, view))
}

func WriteCreateIndexFunc(w io.Writer, view View, table *schema.Table) {
	for _, ix := range table.Index {
		view.Body1 = append(view.Body1, inflect.Camelize(ix.Name))
	}
	must(tCreateIndex.Execute(w, view))
}

// join is a helper function that joins nodes
// together by name using the seperator.
func must(err error) {
	output.Require(err == nil, "%v\n", err)
}
