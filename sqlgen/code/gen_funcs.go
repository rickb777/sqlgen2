package code

import (
	"io"
	"github.com/rickb777/sqlgen2"
	"github.com/rickb777/sqlgen2/schema"
	"github.com/rickb777/sqlgen2/sqlgen/output"
	"fmt"
	"strings"
)

func WriteType(w io.Writer, view View) {
	must(tTable.Execute(w, view))
}

func WriteQueryFuncs(w io.Writer, view View) {
	fmt.Fprintln(w, sectionBreak)

	must(tQueryRow.Execute(w, view))
	must(tQueryRows.Execute(w, view))
}

func WriteSelectRow(w io.Writer, view View) {
	fmt.Fprintln(w, sectionBreak)

	must(tSelectRow.Execute(w, view))
	must(tSelectRows.Execute(w, view))
	must(tCountRows.Execute(w, view))

	tableName := view.Prefix + view.Table.Type
	fmt.Fprintf(w, constStringQ,
		identifier("", tableName, "ColumnNames"), strings.Join(view.Table.ColumnNames(true), ", "))
}

func WriteInsertFunc(w io.Writer, view View) {
	fmt.Fprintln(w, sectionBreak)

	if view.Table.HasLastInsertId() {
		must(tInsertAndGetLastId.Execute(w, view))
	} else {
		must(tInsertSimple.Execute(w, view))
	}

	tableName := view.Prefix + view.Table.Type
	fmt.Fprintf(w, constStringWithTicks,
		identifier("sqlInsert", tableName, "Simple"), schema.New(schema.Sqlite).Insert(view.Table))

	fmt.Fprintf(w, constStringWithTicks,
		identifier("sqlInsert", tableName, "Postgres"), schema.New(schema.Postgres).Insert(view.Table))

	//fmt.Fprintf(w, constStringQ,
	//	identifier("s", tableName, "ColumnParamsSimple"), sqlgen2.Sqlite.Placeholders(table.NumColumnNames(true)))

	fmt.Fprintf(w, constStringQ,
		identifier("s", tableName, "DataColumnParamsSimple"), sqlgen2.Sqlite.Placeholders(view.Table.NumColumnNames(false)))

	//fmt.Fprintf(w, constStringQ,
	//	identifier("s", tableName, "ColumnParamsPostgres"), sqlgen2.Postgres.Placeholders(table.NumColumnNames(true)))

	fmt.Fprintf(w, constStringQ,
		identifier("s", tableName, "DataColumnParamsPostgres"), sqlgen2.Postgres.Placeholders(view.Table.NumColumnNames(false)))

}

func WriteUpdateFunc(w io.Writer, view View) {
	if view.Table.HasPrimaryKey() {
		fmt.Fprintln(w, sectionBreak)

		must(tUpdate.Execute(w, view))

		tableName := view.Prefix + view.Table.Type
		fmt.Fprintf(w, constStringWithTicks,
			identifier("sqlUpdate", tableName, "ByPkSimple"), schema.New(schema.Sqlite).Update(view.Table, []*schema.Field{view.Table.Primary}))

		fmt.Fprintf(w, constStringWithTicks,
			identifier("sqlUpdate", tableName, "ByPkPostgres"), schema.New(schema.Postgres).Update(view.Table, []*schema.Field{view.Table.Primary}))
	}
}

func WriteDeleteFunc(w io.Writer, view View) {
	fmt.Fprintln(w, sectionBreak)
	must(tDelete.Execute(w, view))
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
