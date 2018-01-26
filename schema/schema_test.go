package schema

import (
	"testing"
)

func TestDistinctTypes(t *testing.T) {
	cases := []struct {
		list     FieldList
		expected TypeSet
	}{
		{FieldList{id}, NewTypeSet(i64)},
		{FieldList{id, id, id}, NewTypeSet(i64)},
		{FieldList{id, category}, NewTypeSet(i64, cat)},
		{FieldList{id,
			category,
			name,
			qual,
			diff,
			age,
			bmi,
			active,
			labels,
			fave,
			avatar,
			updated}, NewTypeSet(i64, boo, cat, str, spt, ipt, upt, fpt, bgi, sli, bys, tim)},
	}
	for _, c := range cases {
		s := c.list.DistinctTypes()
		if !NewTypeSet(s...).Equals(c.expected) {
			t.Errorf("expected %d::%+v but got %d::%+v", len(c.expected), c.expected, len(s), s)
		}
	}
}
