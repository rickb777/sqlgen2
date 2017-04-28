package main

import (
	"fmt"
	"io"

	"bitbucket.org/pkg/inflect"
	"github.com/rickb777/sqlgen/parse"
	"github.com/rickb777/sqlgen/schema"
)

const sectionBreak = "\n//--------------------------------------------------------------------------------"

// writeSchema writes SQL statements to CREATE, INSERT,
// UPDATE and DELETE values from Table t.
func writeSchema(w io.Writer, d schema.Dialect, tree *parse.Node, t *schema.Table) {

	fmt.Fprintln(w, sectionBreak)

	writeConst(w,
		d.Table(t),
		identifier("Create", inflect.Singularize(t.Name), "Stmt"),
	)

	writeConst(w,
		d.Insert(t),
		identifier("Insert", inflect.Singularize(t.Name), "Stmt"),
	)

	writeConst(w,
		d.Select(t, nil),
		identifier("Select", inflect.Singularize(t.Name), "Stmt"),
	)

	writeConst(w,
		d.SelectRange(t, nil),
		identifier("Select", inflect.Singularize(t.Name), "RangeStmt"),
	)

	writeConst(w,
		d.SelectCount(t, nil),
		identifier("Select", inflect.Singularize(t.Name), "CountStmt"),
	)

	if t.HasPrimaryKey() {
		writeConst(w,
			d.Select(t, []*schema.Field{t.Primary}),
			identifier("Select", inflect.Singularize(t.Name), "ByPkStmt"),
		)

		writeConst(w,
			d.Update(t, []*schema.Field{t.Primary}),
			identifier("Update", inflect.Singularize(t.Name), "ByPkStmt"),
		)

		//fmt.Fprintf(w, "var %s = map[string]string{\n",
		//	inflect.Typeify(fmt.Sprintf("update_%s_json_map", inflect.Singularize(t.Name))))
		//for i, node := range tree.Edges() {
		//	if i < len(t.Fields) {
		//		columnName := t.Fields[i].SqlName
		//		//if field.Patchable
		//		jsonAttr := strings.Split(node.Tags.JSONAttr, ",")[0]
		//		if jsonAttr != "-" {
		//			fmt.Fprintf(w, "\"%s\": \"%s\",\n", jsonAttr, columnName)
		//		}
		//	}
		//}
		//fmt.Fprintf(w, "}")

		writeConst(w,
			d.Delete(t, []*schema.Field{t.Primary}),
			identifier("Delete", inflect.Singularize(t.Name), "ByPkeyStmt"),
		)
	}

	fmt.Fprintln(w, sectionBreak)

	for _, ix := range t.Index {

		writeConst(w,
			d.Index(t, ix),
			identifier("Create", ix.Name, "Stmt"),
		)

		writeConst(w,
			d.Select(t, ix.Fields),
			identifier("Select", ix.Name, "Stmt"),
		)

		if !ix.Unique {

			writeConst(w,
				d.SelectRange(t, ix.Fields),
				identifier("Select", ix.Name, "RangeStmt"),
			)

			writeConst(w,
				d.SelectCount(t, ix.Fields),
				identifier("Select", ix.Name, "CountStmt"),
			)

		} else {

			writeConst(w,
				d.Update(t, ix.Fields),
				identifier("Update", ix.Name, "Stmt"),
			)

			writeConst(w,
				d.Delete(t, ix.Fields),
				identifier("Delete", ix.Name, "Stmt"),
			)
		}
	}
}

// WritePackage writes the Go package header to
// writer w with the given package name.
func writePackage(w io.Writer, name string) {
	fmt.Fprintf(w, sPackage, name)
}

// writeConst is a helper function that writes the
// body string to a Go const.
func writeConst(w io.Writer, body string, name string) {
	// quote the body using multi-line quotes
	body = fmt.Sprintf("`\n%s\n`", body)

	fmt.Fprintf(w, sConst, name, body, name, name)
}

func identifier(prefix, id, suffix string) string {
	return prefix + inflect.Camelize(id) + suffix
}