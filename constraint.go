package sqlgen2

import (
	"fmt"
)

// Constraint represents data that augments the data-definition SQL statements such as CREATE TABLE.
type Constraint interface {
	// ConstraintSql constructs the CONSTRAINT clause to be included in the CREATE TABLE.
	ConstraintSql(name TableName, index int) string
}

// Constraints holds constraints.
type Constraints []Constraint

// ConstraintSql constructs a list of statements to be included in the CREATE TABLE.
func (cc Constraints) ConstraintSql(name TableName) (statements []string) {
	for i, c := range cc {
		statements = append(statements, c.ConstraintSql(name, i+1))
	}
	return statements
}

//-------------------------------------------------------------------------------------------------

// CheckConstraint holds an expression that refers to table columns and is applied as a precondition
// whenever a table insert, update or delete is attempted. The CheckConstraint expression is in SQL.
type CheckConstraint struct {
	Expression string
}

// ConstraintSql constructs the CONSTRAINT clause to be included in the CREATE TABLE.
func (c CheckConstraint) ConstraintSql(name TableName, index int) string {
	return fmt.Sprintf("CONSTRAINT %s_c%d CHECK (%s)", name, index, c.Expression)
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
// Column is the 'owner' of the constraint.
type FkConstraint struct {
	Column         string // only one column is supported
	Parent         Reference
	Update, Delete Consequence
}

// FkConstraintOn constructs a foreign key constraint in a fluent style.
func FkConstraintOn(column string) FkConstraint {
	return FkConstraint{Column: column}
}

// RefersTo sets the parent reference.
func (c FkConstraint) RefersTo(tableName, column string) FkConstraint {
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
func (c FkConstraint) ConstraintSql(name TableName, index int) string {
	return fmt.Sprintf("CONSTRAINT %s_c%d %s", name, index, c.Sql(name.Prefix))
}

// Sql constructs the foreign key clause needed to configure the database.
func (c FkConstraint) Sql(prefix string) string {
	return fmt.Sprintf("foreign key (%s) references %s%s (%s)%s%s",
		c.Column, prefix, c.Parent.TableName, c.Parent.Column,
		c.Update.Apply(" ", "update"),
		c.Delete.Apply(" ", "delete"))
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

//type AlterTable struct {
//	fullName TableName
//	stmt     string
//}
//
//func (a AlterTable) Sql() string {
//	return fmt.Sprintf("ALTER TABLE %s\n  ADD %s;", a.fullName, a.stmt)
//}
