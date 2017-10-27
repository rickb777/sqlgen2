// A simple type derived from map[string]struct{}
// Not thread-safe.
//
// Generated from simple/set.tpl with Type=string
// options: Numeric:<no value> Stringer:true Mutable:always

package code

import (
	"bytes"
	"fmt"
)

// StringSet is the primary type that represents a set
type StringSet map[string]struct{}

// NewStringSet creates and returns a reference to an empty set.
func NewStringSet(values ...string) StringSet {
	set := make(StringSet)
	for _, i := range values {
		set[i] = struct{}{}
	}
	return set
}

// ConvertStringSet constructs a new set containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
func ConvertStringSet(values ...interface{}) (StringSet, bool) {
	set := make(StringSet)

	for _, i := range values {
		v, ok := i.(string)
		if ok {
			set[v] = struct{}{}
		}
	}

	return set, len(set) == len(values)
}

// ToSlice returns the elements of the current set as a slice.
func (set StringSet) ToSlice() []string {
	var s []string
	for v := range set {
		s = append(s, v)
	}
	return s
}

// ToInterfaceSlice returns the elements of the current set as a slice of arbitrary type.
func (set StringSet) ToInterfaceSlice() []interface{} {
	var s []interface{}
	for v := range set {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (set StringSet) Clone() StringSet {
	clonedSet := NewStringSet()
	for v := range set {
		clonedSet.doAdd(v)
	}
	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set StringSet) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set StringSet) NonEmpty() bool {
	return set.Size() > 0
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set StringSet) Size() int {
	return len(set)
}

//-------------------------------------------------------------------------------------------------

// Add adds items to the current set, returning the modified set.
func (set StringSet) Add(i ...string) StringSet {
	for _, v := range i {
		set.doAdd(v)
	}
	return set
}

func (set StringSet) doAdd(i string) {
	set[i] = struct{}{}
}

// Contains determines if a given item is already in the set.
func (set StringSet) Contains(i string) bool {
	_, found := set[i]
	return found
}

// ContainsAll determines if the given items are all in the set
func (set StringSet) ContainsAll(i ...string) bool {
	for _, v := range i {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

//-------------------------------------------------------------------------------------------------

// IsSubset determines if every item in the other set is in this set.
func (set StringSet) IsSubset(other StringSet) bool {
	for v := range set {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

// IsSuperset determines if every item of this set is in the other set.
func (set StringSet) IsSuperset(other StringSet) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set StringSet) Append(more ...string) StringSet {
	unionedSet := set.Clone()
	for _, v := range more {
		unionedSet.doAdd(v)
	}
	return unionedSet
}

// Union returns a new set with all items in both sets.
func (set StringSet) Union(other StringSet) StringSet {
	unionedSet := set.Clone()
	for v := range other {
		unionedSet.doAdd(v)
	}
	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set StringSet) Intersect(other StringSet) StringSet {
	intersection := NewStringSet()
	// loop over smaller set
	if set.Size() < other.Size() {
		for v := range set {
			if other.Contains(v) {
				intersection.doAdd(v)
			}
		}
	} else {
		for v := range other {
			if set.Contains(v) {
				intersection.doAdd(v)
			}
		}
	}
	return intersection
}

// Difference returns a new set with items in the current set but not in the other set
func (set StringSet) Difference(other StringSet) StringSet {
	differencedSet := NewStringSet()
	for v := range set {
		if !other.Contains(v) {
			differencedSet.doAdd(v)
		}
	}
	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set StringSet) SymmetricDifference(other StringSet) StringSet {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Clear clears the entire set to be the empty set.
func (set *StringSet) Clear() {
	*set = NewStringSet()
}

// Remove allows the removal of a single item from the set.
func (set StringSet) Remove(i string) {
	delete(set, i)
}

//-------------------------------------------------------------------------------------------------

// Equals determines if two sets are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (set StringSet) Equals(other StringSet) bool {
	if set.Size() != other.Size() {
		return false
	}
	for v := range set {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

//-------------------------------------------------------------------------------------------------

func (set StringSet) StringList() []string {
	strings := make([]string, 0)
	for v := range set {
		strings = append(strings, fmt.Sprintf("%v", v))
	}
	return strings
}

func (set StringSet) String() string {
	return set.mkString3Bytes("", ", ", "").String()
}

// implements encoding.Marshaler interface {
func (set StringSet) MarshalJSON() ([]byte, error) {
	return set.mkString3Bytes("[\"", "\", \"", "\"").Bytes(), nil
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (set StringSet) MkString(sep string) string {
	return set.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (set StringSet) MkString3(pfx, mid, sfx string) string {
	return set.mkString3Bytes(pfx, mid, sfx).String()
}

func (set StringSet) mkString3Bytes(pfx, mid, sfx string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(pfx)
	sep := ""
	for v := range set {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v", v))
		sep = mid
	}
	b.WriteString(sfx)
	return b
}

// StringMap renders the set as a map of strings. The value of each item in the set becomes stringified as a key in the
// resulting map.
func (set StringSet) StringMap() map[string]bool {
	strings := make(map[string]bool)
	for v := range set {
		strings[fmt.Sprintf("%v", v)] = true
	}
	return strings
}
