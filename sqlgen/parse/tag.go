package parse

import (
	"fmt"
	"reflect"
	"strings"

	"gopkg.in/yaml.v2"
	"os"
	"io/ioutil"
	"bytes"
)

const TagKey = "sql"

// Tag stores the parsed data from the tag string in
// a struct field.
type Tag struct {
	Name       string `yaml:"name"`
	Type       string `yaml:"type"`
	Default    string `yaml:"default"`
	Prefixed   bool   `yaml:"prefixed"`
	Primary    bool   `yaml:"pk"`
	Natural    bool   `yaml:"nk"`
	Auto       bool   `yaml:"auto"`
	Index      string `yaml:"index"`
	Unique     string `yaml:"unique"`
	//Check      string `yaml:"check"` TODO
	ForeignKey string `yaml:"fk"`
	OnUpdate   string `yaml:"onupdate"`
	OnDelete   string `yaml:"ondelete"`
	Size       int    `yaml:"size"`
	Skip       bool   `yaml:"skip"`
	Encode     string `yaml:"encode"`
}

func (tag Tag) ParentTable() string {
	if tag.ForeignKey == "" {
		return ""
	}
	slice := strings.Split(tag.ForeignKey, ".")
	return slice[0]
}

func (tag Tag) ParentPK() string {
	if tag.ForeignKey == "" {
		return ""
	}
	slice := strings.Split(tag.ForeignKey, ".")
	return slice[1]
}

//-------------------------------------------------------------------------------------------------

func inSet(s string, set ...string) bool {
	for _, x := range set {
		if s == x {
			return true
		}
	}
	return false
}

func normalizeConsequence(consequence string) string {
	if consequence == "setnull" {
		return "set null"
	}
	if consequence == "setdefault" {
		return "set default"
	}
	return consequence
}

func normalize(tag *Tag) *Tag {
	tag.OnUpdate = normalizeConsequence(tag.OnUpdate)
	tag.OnDelete = normalizeConsequence(tag.OnDelete)
	return tag
}

func validate(tag *Tag) error {
	sep := ""
	buf := &bytes.Buffer{}

	if tag.Primary && tag.Natural {
		fmt.Fprintf(buf, "%sprimary key cannot also be a natural key", sep)
		sep = "; "
	}

	if tag.Auto && tag.Natural {
		fmt.Fprintf(buf, "%snatural key cannot use auto-increment", sep)
		sep = "; "
	}

	if tag.Auto && !tag.Primary {
		fmt.Fprintf(buf, "%sauto-increment can only be used on primary keys", sep)
		sep = "; "
	}

	if tag.Natural && tag.Index != "" {
		fmt.Fprintf(buf, "%snatural key cannot be used with index", sep)
		sep = "; "
	}

	if tag.Natural && tag.Unique != "" {
		fmt.Fprintf(buf, "%snatural key should not be used with unique", sep)
		sep = "; "
	}

	if tag.Size < 0 {
		fmt.Fprintf(buf, "%ssize cannot be negative (%d)", sep, tag.Size)
		sep = "; "
	}

	if !inSet(tag.Encode, "", "json", "text", "driver") {
		fmt.Fprintf(buf, "%sunrecognised encode value %q", sep, tag.Encode)
		sep = "; "
	}

	if tag.ForeignKey != "" {
		if tag.Primary || tag.Natural {
			fmt.Fprintf(buf, "%sforeign key cannot also be a primary key nor a natural key", sep)
			sep = "; "
		}

		slice := strings.Split(tag.ForeignKey, ".")
		if len(slice) != 2 {
			fmt.Fprintf(buf, "%sfk value (%q) must be in 'tablename.column' form", sep, tag.ForeignKey)
			sep = "; "
		}
	}

	if !inSet(tag.OnUpdate, "", "cascade", "delete", "restrict", "set null", "set default") {
		fmt.Fprintf(buf, "%sunrecognised onupdate value %q", sep, tag.OnUpdate)
		sep = "; "
	}

	if !inSet(tag.OnDelete, "", "cascade", "delete", "restrict", "set null", "set default") {
		fmt.Fprintf(buf, "%sunrecognised ondelete value %q", sep, tag.OnDelete)
		sep = "; "
	}

	if buf.Len() > 0 {
		return fmt.Errorf(buf.String())
	}
	return nil
}

// ParseTag parses a tag string from the struct
// field and unmarshals into a Tag struct.
func ParseTag(raw string) (*Tag, error) {
	var tag = new(Tag)

	raw = strings.TrimSpace(topAndTail(strings.TrimSpace(raw)))
	structTag := reflect.StructTag(raw)
	value := strings.TrimSpace(structTag.Get(TagKey))

	if value == "-" {
		tag.Skip = true
		return tag, nil
	}

	// wrap the string in curly braces so that we can use the Yaml parser.
	yamlValue := fmt.Sprintf("{ %s }", value)

	// unmarshals the Yaml formatted string into the Tag structure.
	var err = yaml.Unmarshal([]byte(yamlValue), tag)
	if err != nil {
		return tag, err
	}

	normalize(tag)
	return tag, validate(tag)
}

type Tags map[string]Tag

func (tags Tags) String() string {
	b := &bytes.Buffer{}
	for n, t := range tags {
		fmt.Fprintf(b, "%-10s: %+v\n", n, t)
	}
	return b.String()
}

func topAndTail(s string) string {
	last := len(s) - 1
	if len(s) >= 2 && s[0] == s[last] {
		return s[1:last]
	}
	return s
}

func ReadTagsFile(file string) (Tags, error) {
	if file == "" {
		return nil, nil
	}

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	tags := make(Tags)

	err = yaml.Unmarshal(b, tags)
	if err != nil {
		return nil, err
	}

	DevInfo("tags from %s\n%s\n", file, tags)
	return tags, nil
}
