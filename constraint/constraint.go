package constraint

import (
	"fmt"
	"github.com/rickb777/sqlgen2"
	"github.com/rickb777/sqlgen2/util"
)

type Dialect interface {
	Quote(column string) string
}

// Constraint represents data that augments the data-definition SQL statements such as CREATE TABLE.
type Constraint interface {
	// ConstraintSql constructs the CONSTRAINT clause to be included in the CREATE TABLE.
	ConstraintSql(dialect Dialect, name sqlgen2.TableName, index int) string
	GoString() string
}

// Constraints holds constraints.
type Constraints []Constraint

//-------------------------------------------------------------------------------------------------

// CheckConstraint holds an expression that refers to table columns and is applied as a precondition
// whenever a table insert, update or delete is attempted. The CheckConstraint expression is in SQL.
type CheckConstraint struct {
	Expression string
}

// ConstraintSql constructs the CONSTRAINT clause to be included in the CREATE TABLE.
func (c CheckConstraint) ConstraintSql(dialect Dialect, name sqlgen2.TableName, index int) string {
	return fmt.Sprintf("CONSTRAINT %s_c%d CHECK (%s)", name, index, c.Expression)
}

func (c CheckConstraint) GoString() string {
	return "TODO"
}

//-------------------------------------------------------------------------------------------------

// Consequence is the action to be performed after updating or deleting a record constrained by foreign key.
type Consequence string

const (
	// unspecified option is available but its semantics vary by DB vendor, so it's not included here.
	NoAction   Consequence = "no action"
	Restrict   Consequence = "restrict"
	Cascade    Consequence = "cascade"
	SetNull    Consequence = "set null"
	SetDefault Consequence = "set default"
	Delete     Consequence = "delete" // not MySQL
)

// Apply constructs the SQL sub-clause for a consequence of a specified action.
// The prefix is typically arbitrary whitespace.
func (c Consequence) Apply(pfx, action string) string {
	if c == "" {
		return "" // implicitly equivalent to NoAction
	}
	return fmt.Sprintf("%son %s %s", pfx, action, c)
}

//-------------------------------------------------------------------------------------------------

// Reference holds a table + column reference used by constraints.
type Reference struct {
	TableName string // without schema or other prefix
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
	return fmt.Sprintf("CONSTRAINT %s_c%d %s", name, index, c.Sql(dialect , name.Prefix))
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

func (c FkConstraint) Disabled() FkConstraint {
	c.Update = NoAction
	c.Delete = NoAction
	return c
}

//-------------------------------------------------------------------------------------------------

// IdsUnusedAsForeignKeys finds all the primary keys in the parent table that have no foreign key
// in the dependent (child) table.
func (c FkConstraint) IdsUnusedAsForeignKeys(tbl sqlgen2.Table) (util.Int64Set, error) {
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
	s := fmt.Sprintf(
		`SELECT a.%s
			FROM %s%s a
			LEFT OUTER JOIN %s b ON a.%s = b.%s
			WHERE b.%s IS null`,
		c.Parent.Column, tbl.Name().Prefix, c.Parent.TableName, tbl.Name(), c.Parent.Column, c.ForeignKeyColumn, c.ForeignKeyColumn)
	return fetchIds(tbl, s)
}

// IdsUsedAsForeignKeys finds all the primary keys in the parent table that have at least one foreign key
// in the dependent (child) table.
func (c FkConstraint) IdsUsedAsForeignKeys(tbl sqlgen2.Table) (util.Int64Set, error) {
	s := fmt.Sprintf(
		`SELECT DISTINCT a.%s AS Id
			FROM %s%s a
			INNER JOIN %s b ON a.%s = b.%s`,
		c.Parent.Column, tbl.Name().Prefix, c.Parent.TableName, tbl.Name(), c.Parent.Column, c.ForeignKeyColumn)
	return fetchIds(tbl, s)
}

func fetchIds(tbl sqlgen2.Table, query string) (util.Int64Set, error) {
	database := tbl.Database()
	database.LogQuery(query)
	rows, err := tbl.Execer().QueryContext(tbl.Ctx(), query)
	if err != nil {
		database.LogError(err)
		return nil, err
	}
	defer rows.Close()

	set := util.NewInt64Set()
	for rows.Next() {
		var id int64
		rows.Scan(&id)
		set.Add(id)
	}
	return set, database.LogIfError(rows.Err())
}
