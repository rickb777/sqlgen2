package code

import (
	"io"
	"github.com/rickb777/sqlgen2/dialect"
	"github.com/rickb777/sqlgen2/schema"
	"github.com/rickb777/sqlgen2/sqlgen/output"
	"fmt"
	"strings"
)

func WriteType(w io.Writer, view View) {
	must(tTable.Execute(w, view))
}

func WriteSelectRow(w io.Writer, view View, table *schema.Table) {
	must(tSelectRow.Execute(w, view))
	must(tSelectRows.Execute(w, view))
	must(tCountRows.Execute(w, view))

	tableName := view.Prefix + table.Type
	fmt.Fprintf(w, constStringQ,
		identifier("", tableName, "ColumnNames"), strings.Join(table.ColumnNames(true), ", "))
}

func WriteInsertFunc(w io.Writer, view View, table *schema.Table) {
	if table.HasLastInsertId() {
		must(tInsertAndGetLastId.Execute(w, view))
	} else {
		must(tInsertSimple.Execute(w, view))
	}

	tableName := view.Prefix + table.Type
	fmt.Fprintf(w, constStringWithTicks,
		identifier("sqlInsert", tableName, "Simple"), schema.New(schema.Sqlite).Insert(table))

	fmt.Fprintf(w, constStringWithTicks,
		identifier("sqlInsert", tableName, "Postgres"), schema.New(schema.Postgres).Insert(table))

	//fmt.Fprintf(w, constStringQ,
	//	identifier("s", tableName, "ColumnParamsSimple"), dialect.Sqlite.Placeholders(table.NumColumnNames(true)))

	fmt.Fprintf(w, constStringQ,
		identifier("s", tableName, "DataColumnParamsSimple"), dialect.Sqlite.Placeholders(table.NumColumnNames(false)))

	//fmt.Fprintf(w, constStringQ,
	//	identifier("s", tableName, "ColumnParamsPostgres"), dialect.Postgres.Placeholders(table.NumColumnNames(true)))

	fmt.Fprintf(w, constStringQ,
		identifier("s", tableName, "DataColumnParamsPostgres"), dialect.Postgres.Placeholders(table.NumColumnNames(false)))

}

func WriteUpdateFunc(w io.Writer, view View, table *schema.Table) {
	if table.HasPrimaryKey() {
		must(tUpdate.Execute(w, view))

		tableName := view.Prefix + table.Type
		fmt.Fprintf(w, constStringWithTicks,
			identifier("sqlUpdate", tableName, "ByPkSimple"), schema.New(schema.Sqlite).Update(table, []*schema.Field{table.Primary}))

		fmt.Fprintf(w, constStringWithTicks,
			identifier("sqlUpdate", tableName, "ByPkPostgres"), schema.New(schema.Postgres).Update(table, []*schema.Field{table.Primary}))
	}
}

func WriteExecFunc(w io.Writer, view View, table *schema.Table) {
	must(tExec.Execute(w, view))
}

// join is a helper function that joins nodes
// together by name using the seperator.
func must(err error) {
	output.Require(err == nil, "%v\n", err)
}
