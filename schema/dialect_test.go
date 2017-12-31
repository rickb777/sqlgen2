package schema

import (
	"testing"
)

func TestPlaceholders(t *testing.T) {
	cases := []struct{
		di     SDialect
		n int
		expected string
	}{
		{MysqlDialect, 0, ""},
		{MysqlDialect, 1, "?"},
		{MysqlDialect, 3, "?,?,?"},
		{MysqlDialect, 11, "?,?,?,?,?,?,?,?,?,?,?"},

		{PostgresDialect, 0, ""},
		{PostgresDialect, 1, "$1"},
		{PostgresDialect, 3, "$1,$2,$3"},
		{PostgresDialect, 11, "$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11"},
	}
	for _, c := range cases {
		s := c.di.Placeholders(c.n)
		if s != c.expected {
			t.Errorf("expected %q but got %q", c.expected, s)
		}
	}
}

func TestReplacePlaceholders(t *testing.T) {
	s := MysqlDialect.ReplacePlaceholders("?,?,?,?,?,?,?,?,?,?,?")
	if s != "?,?,?,?,?,?,?,?,?,?,?" {
		t.Errorf("expected ?,?,?,?,?,?,?,?,?,?,? but got %q", s)
	}

	s = PostgresDialect.ReplacePlaceholders("?,?,?,?,?,?,?,?,?,?,?")
	if s != "$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11" {
		t.Errorf("expected $1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11 but got %q", s)
	}

}
