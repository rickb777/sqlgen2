// generated code - do not edit

package parse

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const kindEnumStrings = "InvalidBoolIntInt8Int16Int32Int64UintUint8Uint16Uint32Uint64Float32Float64Complex64Complex128InterfaceBytesMapPtrStringSliceStruct"

var kindEnumIndex = [...]uint16{0, 7, 11, 14, 18, 23, 28, 33, 37, 42, 48, 54, 60, 67, 74, 83, 93, 102, 107, 110, 113, 119, 124, 130}

var AllKinds = []Kind{Invalid, Bool, Int, Int8, Int16, Int32, Int64, Uint, Uint8, Uint16, Uint32, Uint64, Float32, Float64, Complex64, Complex128, Interface, Bytes, Map, Ptr, String, Slice, Struct}

// String returns the string representation of a Kind
func (i Kind) String() string {
	o := i.Ordinal()
	if o < 0 || o >= len(AllKinds) {
		return fmt.Sprintf("Kind(%d)", i)
	}
	return kindEnumStrings[kindEnumIndex[o]:kindEnumIndex[o+1]]
}

// Ordinal returns the ordinal number of a Kind
func (i Kind) Ordinal() int {
	switch i {
	case Invalid:
		return 0
	case Bool:
		return 1
	case Int:
		return 2
	case Int8:
		return 3
	case Int16:
		return 4
	case Int32:
		return 5
	case Int64:
		return 6
	case Uint:
		return 7
	case Uint8:
		return 8
	case Uint16:
		return 9
	case Uint32:
		return 10
	case Uint64:
		return 11
	case Float32:
		return 12
	case Float64:
		return 13
	case Complex64:
		return 14
	case Complex128:
		return 15
	case Interface:
		return 16
	case Bytes:
		return 17
	case Map:
		return 18
	case Ptr:
		return 19
	case String:
		return 20
	case Slice:
		return 21
	case Struct:
		return 22
	}
	return -1
}

// Parse parses a string to find the corresponding Kind, accepting either one of the string
// values or an ordinal number.
func (v *Kind) Parse(s string) error {
	ord, err := strconv.Atoi(s)
	if err == nil && 0 <= ord && ord < len(AllKinds) {
		*v = AllKinds[ord]
		return nil
	}
	var i0 uint16 = 0
	for j := 1; j < len(kindEnumIndex); j++ {
		i1 := kindEnumIndex[j]
		p := kindEnumStrings[i0:i1]
		if s == p {
			*v = AllKinds[j-1]
			return nil
		}
		i0 = i1
	}
	return errors.New(s + ": unrecognised Kind")
}

// AsKind parses a string to find the corresponding Kind, accepting either one of the string
// values or an ordinal number.
func AsKind(s string) (Kind, error) {
	var i = new(Kind)
	err := i.Parse(s)
	return *i, err
}

// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
func (i Kind) MarshalText() (text []byte, err error) {
	return []byte(i.String()), nil
}

// UnmarshalText converts transmitted values to ordinary values.
func (i *Kind) UnmarshalText(text []byte) error {
	return i.Parse(string(text))
}

// KindMarshalJSONUsingString controls whether generated JSON uses ordinals or strings. By default,
// it is false and ordinals are used. Set it true to cause quoted strings to be used instead,
// these being easier to read but taking more resources.
var KindMarshalJSONUsingString = false

// MarshalJSON converts values to bytes suitable for transmission via JSON. By default, the
// ordinal integer is emitted, but a quoted string is emitted instead if
// KindMarshalJSONUsingString is true.
func (i Kind) MarshalJSON() ([]byte, error) {
	if KindMarshalJSONUsingString {
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
func (i *Kind) UnmarshalJSON(text []byte) error {
	// Ignore null, like in the main JSON package.
	if string(text) == "null" {
		return nil
	}
    s := strings.Trim(string(text), "\"")
	return i.Parse(s)
}
