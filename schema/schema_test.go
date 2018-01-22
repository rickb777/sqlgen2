package schema

import (
	"testing"
	"github.com/rickb777/sqlgen2/model"
)

func TestDistinctTypes(t *testing.T) {
	cases := []struct {
		list     FieldList
		expected model.TypeSet
	}{
		{FieldList{id}, model.NewTypeSet(i64)},
		{FieldList{id, id, id}, model.NewTypeSet(i64)},
		{FieldList{id, category}, model.NewTypeSet(i64, cat)},
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
			updated}, model.NewTypeSet(i64, boo, cat, str, spt, ipt, upt, fpt, bgi, sli, bys, tim)},
	}
	for _, c := range cases {
		s := c.list.DistinctTypes()
		if !model.NewTypeSet(s...).Equals(c.expected) {
			t.Errorf("expected %d::%+v but got %d::%+v", len(c.expected), c.expected, len(s), s)
		}
	}
}
