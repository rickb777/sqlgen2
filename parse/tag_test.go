package parse

import (
	"reflect"
	"testing"
)

var tagTests = []struct {
	raw string
	tag *Tag
}{
	{
		TagKey + `:"-"`,
		&Tag{Skip: true},
	},
	{
		TagKey + `:"pk: true, auto: true"`,
		&Tag{Primary: true, Auto: true},
	},
	{
		TagKey + `:"name: foo"`,
		&Tag{Name: "foo"},
	},
	{
		TagKey + `:"type: varchar"`,
		&Tag{Type: "varchar"},
	},
	{
		TagKey + `:"size: 2048"`,
		&Tag{Size: 2048},
	},
	{
		TagKey + `:"index: fake_index"`,
		&Tag{Index: "fake_index"},
	},
	{
		TagKey + `:"unique: fake_unique_index"`,
		&Tag{Unique: "fake_unique_index"},
	},
}

func TestParseTag(t *testing.T) {
	for _, test := range tagTests {

		var want = test.tag
		var got, err = parseTag(test.raw)

		if err != nil {
			t.Errorf("Got Error parsing Tag %s. %s", test.raw, err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Wanted Tag %+v, got Tag %+v", want, got)
		}
	}
}
