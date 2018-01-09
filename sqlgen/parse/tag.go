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
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	Prefixed bool   `yaml:"prefixed"`
	Primary  bool   `yaml:"pk"`
	Auto     bool   `yaml:"auto"`
	Index    string `yaml:"index"`
	Unique   string `yaml:"unique"`
	Size     int    `yaml:"size"`
	Skip     bool   `yaml:"skip"`
	Encode   string `yaml:"encode"`
	//JSONAttr string `yaml:"json"`
}

// ParseTag parses a tag string from the struct
// field and unmarshals into a Tag struct.
func ParseTag(raw string) (*Tag, error) {
	var tag = new(Tag)

	raw = strings.TrimSpace(topAndTail(strings.TrimSpace(raw)))
	structTag := reflect.StructTag(raw)
	value := strings.TrimSpace(structTag.Get(TagKey))

	// if the tag indicates the field should
	// be skipped we can exit right away.
	if value == "-" {
		tag.Skip = true
		return tag, nil
	}

	// otherwise wrap the string in curly braces
	// so that we can use the Yaml parser.
	yamlValue := fmt.Sprintf("{ %s }", value)

	// unmarshals the Yaml formatted string into
	// the Tag structure.
	var err = yaml.Unmarshal([]byte(yamlValue), tag)

	//if tag.JSONAttr == "" {
	//	tag.JSONAttr = structTag.Get("json")
	//}

	return tag, err
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
