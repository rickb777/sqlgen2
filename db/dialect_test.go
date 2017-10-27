package db

import (
	"testing"
)

func TestPlaceholders(t *testing.T) {
	cases := []struct{
		di     Dialect
		n int
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
		if s != c.expected {
			t.Errorf("expected %q but got %q", c.expected, s)
		}
	}
}

func TestReplacePlaceholders(t *testing.T) {
	s := Mysql.ReplacePlaceholders("?,?,?,?,?,?,?,?,?,?,?")
	if s != "?,?,?,?,?,?,?,?,?,?,?" {
		t.Errorf("expected ?,?,?,?,?,?,?,?,?,?,? but got %q", s)
	}

	s = Postgres.ReplacePlaceholders("?,?,?,?,?,?,?,?,?,?,?")
	if s != "$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11" {
		t.Errorf("expected $1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11 but got %q", s)
	}

}
