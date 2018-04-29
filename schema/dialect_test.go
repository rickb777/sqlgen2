package schema

import (
	"testing"
	. "github.com/onsi/gomega"
	"bytes"
)

func TestSplitAndQuote(t *testing.T) {
	RegisterTestingT(t)

	cases := []struct {
		di       Dialect
		expected string
	}{
		{Sqlite, "`A`,`Bb`,`Ccc`"},
		{Mysql, "`A`,`Bb`,`Ccc`"},
		{Postgres, `"a","bb","ccc"`},
	}
	for _, c := range cases {
		s := c.di.SplitAndQuote("A,Bb,Ccc")
		Ω(s).Should(Equal(c.expected), c.di.String())
	}
}

func TestQuote(t *testing.T) {
	RegisterTestingT(t)

	cases := []struct {
		di       Dialect
		expected string
	}{
		{Sqlite, "`Aaaa`"},
		{Mysql, "`Aaaa`"},
		{Postgres, `"aaaa"`},
	}
	for _, c := range cases {
		s1 := c.di.Quote("Aaaa")
		Ω(s1).Should(Equal(c.expected), c.di.String())

		b2 := &bytes.Buffer{}
		c.di.QuoteW(b2, "Aaaa")
		Ω(b2.String()).Should(Equal(c.expected), c.di.String())
	}
}

func TestQuoteWithPlaceholder(t *testing.T) {
	RegisterTestingT(t)

	cases := []struct {
		di       Dialect
		expected string
	}{
		{Sqlite, "`Aaaa`=?"},
		{Mysql, "`Aaaa`=?"},
		{Postgres, `"aaaa"=$3`},
	}
	for _, c := range cases {
		b := &bytes.Buffer{}
		c.di.QuoteWithPlaceholder(b, "Aaaa", 3)
		Ω(b.String()).Should(Equal(c.expected), c.di.String())
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

	s := Mysql.ReplacePlaceholders("?,?,?,?,?,?,?,?,?,?,?", nil)
	Ω(s).Should(Equal("?,?,?,?,?,?,?,?,?,?,?"))

	s = Postgres.ReplacePlaceholders("?,?,?,?,?,?,?,?,?,?,?", nil)
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
