package constraint

import (
	"fmt"
	"github.com/rickb777/sqlgen2"
	"github.com/rickb777/sqlgen2/util"
)

// FkConstraints holds foreign key constraints.
type FkConstraints []FkConstraint

// Find finds the FkConstraint on a specified column.
func (cc FkConstraints) Find(fkColumn string) (FkConstraint, bool) {
	for _, fkc := range cc {
		if fkc.ForeignKeyColumn == fkColumn {
			return fkc, true
		}
	}
	return FkConstraint{}, false
}

//-------------------------------------------------------------------------------------------------

// Reference holds a table + column reference used by constraints.
// The table name should not include any schema or other prefix.
type Reference struct {
	TableName string
	Column    string // only one column is supported
}

//-------------------------------------------------------------------------------------------------

// FkConstraint holds a pair of references and their update/delete consequences.
// ForeignKeyColumn is the 'owner' of the constraint.
type FkConstraint struct {
	ForeignKeyColumn string // only one column is supported
	Parent           Reference
	Update, Delete   Consequence
}

// FkConstraintOn constructs a foreign key constraint in a fluent style.
func FkConstraintOn(column string) FkConstraint {
	return FkConstraint{ForeignKeyColumn: column}
}

// RefersTo sets the parent reference.
func (c FkConstraint) RefersTo(tableName string, column string) FkConstraint {
	c.Parent = Reference{tableName, column}
	return c
}

// OnUpdate sets the update consequence.
func (c FkConstraint) OnUpdate(consequence Consequence) FkConstraint {
	c.Update = consequence
	return c
}

// OnDelete sets the delete consequence.
func (c FkConstraint) OnDelete(consequence Consequence) FkConstraint {
	c.Delete = consequence
	return c
}

// ConstraintSql constructs the CONSTRAINT clause to be included in the CREATE TABLE.
func (c FkConstraint) ConstraintSql(dialect Dialect, name sqlgen2.TableName, index int) string {
	return fmt.Sprintf("CONSTRAINT %s_c%d %s", name, index, c.Sql(dialect, name.Prefix))
}

// Column constructs the foreign key clause needed to configure the database.
func (c FkConstraint) Sql(dialect Dialect, prefix string) string {
	return fmt.Sprintf("foreign key (%s) references %s%s (%s)%s%s",
		dialect.Quote(c.ForeignKeyColumn), prefix, c.Parent.TableName, dialect.Quote(c.Parent.Column),
		c.Update.Apply(" ", "update"),
		c.Delete.Apply(" ", "delete"))
}

func (c FkConstraint) GoString() string {
	return fmt.Sprintf(`constraint.FkConstraint{"%s", constraint.Reference{"%s", "%s"}, "%s", "%s"}`,
		c.ForeignKeyColumn, c.Parent.TableName, c.Parent.Column, c.Update, c.Delete)
}

//func (c FkConstraint) AlterTable() AlterTable {
//	return AlterTable{c.Child.TableName, c.ConstraintSql(0)}
//}

// Disabled changes both the Update and Delete consequences to NoAction.
func (c FkConstraint) Disabled() FkConstraint {
	c.Update = NoAction
	c.Delete = NoAction
	return c
}

// RelationshipWith constructs the Relationship that is expressed by the parent reference in
// the FkConstraint and the child's foreign key.
//
// The table names do not include any prefix.
func (c FkConstraint) RelationshipWith(child sqlgen2.TableName) Relationship {
	return Relationship{
		Parent: c.Parent,
		Child:  Reference{child.Name, c.ForeignKeyColumn},
	}
}

//-------------------------------------------------------------------------------------------------

// Relationship represents a parent-child relationship.
// Only simple keys are supported (compound keys are not supported).
type Relationship struct {
	Parent, Child Reference
}

// IdsUnusedAsForeignKeys finds all the primary keys in the parent table that have no foreign key
// in the dependent (child) table. The table tbl provides the database or transaction handle; either
// the parent or the child table can be used for thi purpose.
func (rel Relationship) IdsUnusedAsForeignKeys(tbl sqlgen2.Table) (util.Int64Set, error) {
	// TODO benchmark two candidates and choose the better
	// http://stackoverflow.com/questions/3427353/sql-statement-question-how-to-retrieve-records-of-a-table-where-the-primary-ke?rq=1
	//	s := fmt.Sprintf(
	//		`SELECT a.%s
	//			FROM %s a
	//			WHERE NOT EXISTS (
	//   				SELECT 1 FROM %s b
	//   				WHERE %s.%s = %s.%s
	//			)`,
	//		primary.ForeignKeyColumn, primary.TableName, foreign.TableName, primary.TableName, primary.ForeignKeyColumn, foreign.TableName, foreign.ForeignKeyColumn)

	// http://stackoverflow.com/questions/13108587/selecting-primary-keys-that-does-not-has-foreign-keys-in-another-table
	pfx := tbl.Name().Prefix
	s := fmt.Sprintf(
		`SELECT a.%s
			FROM %s%s a
			LEFT OUTER JOIN %s%s b ON a.%s = b.%s
			WHERE b.%s IS null`,
		rel.Parent.Column, pfx, rel.Parent.TableName, pfx, rel.Child.TableName, rel.Parent.Column, rel.Child.Column, rel.Child.Column)
	return fetchIds(tbl, s)
}

// IdsUsedAsForeignKeys finds all the primary keys in the parent table that have at least one foreign key
// in the dependent (child) table.
func (rel Relationship) IdsUsedAsForeignKeys(tbl sqlgen2.Table) (util.Int64Set, error) {
	pfx := tbl.Name().Prefix
	s := fmt.Sprintf(
		`SELECT DISTINCT a.%s AS Id
			FROM %s%s a
			INNER JOIN %s%s b ON a.%s = b.%s`,
		rel.Parent.Column, pfx, rel.Parent.TableName, pfx, rel.Child.TableName, rel.Parent.Column, rel.Child.Column)
	return fetchIds(tbl, s)
}

func fetchIds(tbl sqlgen2.Table, query string) (util.Int64Set, error) {
	rows, err := tbl.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	set := util.NewInt64Set()
	for rows.Next() {
		var id int64
		rows.Scan(&id)
		set.Add(id)
	}
	return set, tbl.Database().LogIfError(rows.Err())
}
