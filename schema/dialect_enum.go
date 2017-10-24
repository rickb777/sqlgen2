// generated code - do not edit

package schema

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const dialectidEnumStrings = "SqlitePostgresMysql"

var dialectidEnumIndex = [...]uint16{0, 6, 14, 19}

var AllDialectIds = []DialectId{Sqlite, Postgres, Mysql}

// String returns the string representation of a DialectId
func (i DialectId) String() string {
	o := i.Ordinal()
	if o < 0 || o >= len(AllDialectIds) {
		return fmt.Sprintf("DialectId(%d)", i)
	}
	return dialectidEnumStrings[dialectidEnumIndex[o]:dialectidEnumIndex[o+1]]
}

// Ordinal returns the ordinal number of a DialectId
func (i DialectId) Ordinal() int {
	switch i {
	case Sqlite:
		return 0
	case Postgres:
		return 1
	case Mysql:
		return 2
	}
	return -1
}

// Parse parses a string to find the corresponding DialectId, accepting either one of the string
// values or an ordinal number.
func (v *DialectId) Parse(s string) error {
	ord, err := strconv.Atoi(s)
	if err == nil && 0 <= ord && ord < len(AllDialectIds) {
		*v = AllDialectIds[ord]
		return nil
	}
	var i0 uint16 = 0
	for j := 1; j < len(dialectidEnumIndex); j++ {
		i1 := dialectidEnumIndex[j]
		p := dialectidEnumStrings[i0:i1]
		if s == p {
			*v = AllDialectIds[j-1]
			return nil
		}
		i0 = i1
	}
	return errors.New(s + ": unrecognised DialectId")
}

// AsDialectId parses a string to find the corresponding DialectId, accepting either one of the string
// values or an ordinal number.
func AsDialectId(s string) (DialectId, error) {
	var i = new(DialectId)
	err := i.Parse(s)
	return *i, err
}

// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
func (i DialectId) MarshalText() (text []byte, err error) {
	return []byte(i.String()), nil
}

// UnmarshalText converts transmitted values to ordinary values.
func (i *DialectId) UnmarshalText(text []byte) error {
	return i.Parse(string(text))
}

// DialectIdMarshalJSONUsingString controls whether generated JSON uses ordinals or strings. By default,
// it is false and ordinals are used. Set it true to cause quoted strings to be used instead,
// these being easier to read but taking more resources.
var DialectIdMarshalJSONUsingString = false

// MarshalJSON converts values to bytes suitable for transmission via JSON. By default, the
// ordinal integer is emitted, but a quoted string is emitted instead if
// DialectIdMarshalJSONUsingString is true.
func (i DialectId) MarshalJSON() ([]byte, error) {
	if DialectIdMarshalJSONUsingString {
		s := []byte(i.String())
		b := make([]byte, len(s)+2)
		b[0] = '"'
		copy(b[1:], s)
		b[len(s)+1] = '"'
		return b, nil
	}
	// else use the ordinal
	s := strconv.Itoa(i.Ordinal())
	return []byte(s), nil
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (i *DialectId) UnmarshalJSON(text []byte) error {
	// Ignore null, like in the main JSON package.
	if string(text) == "null" {
		return nil
	}
    s := strings.Trim(string(text), "\"")
	return i.Parse(s)
}
