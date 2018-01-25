package schema

import (
	"testing"
	. "github.com/onsi/gomega"
	"strings"
)

func TestInsertDML(t *testing.T) {
	RegisterTestingT(t)

	table := &TableDescription{"Foo", "Foo", FieldList{id, category, name}, nil, id}

	cases := []struct {
		di       Dialect
		expected string
	}{
		{Sqlite, "\"(`cat`, `username`) VALUES (?,?)\""},
		{Mysql, "\"(`cat`, `username`) VALUES (?,?)\""},
		{Postgres, "`(\"cat\", \"username\") VALUES ($1,$2) returning \"id\"`"},
	}
	for _, c := range cases {
		s := c.di.InsertDML(table)
		Ω(s).Should(Equal(c.expected), c.di.String())
	}
}

func TestUpdateDML(t *testing.T) {
	RegisterTestingT(t)

	table := &TableDescription{"Foo", "Foo", FieldList{id, category, name}, nil, id}

	cases := []struct {
		di       Dialect
		expected string
	}{
		{Sqlite, strings.Replace(`"¬cat¬=?,¬username¬=? WHERE ¬id¬=?"`, "¬", "`", -1)},

		{Mysql, strings.Replace(`"¬cat¬=?,¬username¬=? WHERE ¬id¬=?"`, "¬", "`", -1)},

		{Postgres, strings.Replace(`¬"cat"=$2,"username"=$3 WHERE "id"=$1¬`, "¬", "`", -1)},
	}
	for _, c := range cases {
		s := c.di.UpdateDML(table)
		Ω(s).Should(Equal(c.expected), c.di.String())
	}
}

func TestSplitAndQuote(t *testing.T) {
	RegisterTestingT(t)

	cases := []struct {
		di       Dialect
		expected string
	}{
		{Sqlite, "`a`,`bb`,`ccc`"},
		{Mysql, "`a`,`bb`,`ccc`"},
		{Postgres, `"a","bb","ccc"`},
	}
	for _, c := range cases {
		s := c.di.SplitAndQuote("a,bb,ccc")
		Ω(s).Should(Equal(c.expected), c.di.String())
	}
}

func TestPlaceholders(t *testing.T) {
	RegisterTestingT(t)

	cases := []struct {
		di       Dialect
		n        int
		expected string
	}{
		{Mysql, 0, ""},
		{Mysql, 1, "?"},
		{Mysql, 3, "?,?,?"},
		{Mysql, 11, "?,?,?,?,?,?,?,?,?,?,?"},

		{Postgres, 0, ""},
		{Postgres, 1, "$1"},
		{Postgres, 3, "$1,$2,$3"},
		{Postgres, 11, "$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11"},
	}
	for _, c := range cases {
		s := c.di.Placeholders(c.n)
		Ω(s).Should(Equal(c.expected))
	}
}

func TestReplacePlaceholders(t *testing.T) {
	RegisterTestingT(t)

	s := Mysql.ReplacePlaceholders("?,?,?,?,?,?,?,?,?,?,?")
	Ω(s).Should(Equal("?,?,?,?,?,?,?,?,?,?,?"))

	s = Postgres.ReplacePlaceholders("?,?,?,?,?,?,?,?,?,?,?")
	Ω(s).Should(Equal("$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11"))
}

func TestPickDialect(t *testing.T) {
	RegisterTestingT(t)

	cases := []struct {
		di   Dialect
		name string
	}{
		{Mysql, "MySQL"},
		{Postgres, "Postgres"},
		{Postgres, "PostgreSQL"},
		{Sqlite, "SQLite"},
		{Sqlite, "sqlite3"},
	}
	for _, c := range cases {
		s := PickDialect(c.name)
		Ω(s).Should(Equal(c.di))
	}
}
