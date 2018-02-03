// A simple type derived from map[string]struct{}
// Not thread-safe.
//
// Generated from simple/set.tpl with Type=string
// options: Numeric:false Stringer:<no value> Mutable:always

package util

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

// BuildStringSetFromChan constructs a new StringSet from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildStringSetFromChan(source <-chan string) StringSet {
	set := make(StringSet)
	for v := range source {
		set[v] = struct{}{}
	}
	return set
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
	for v, _ := range set {
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

// IsSequence returns true for lists.
func (set StringSet) IsSequence() bool {
	return false
}

// IsSet returns false for lists.
func (set StringSet) IsSet() bool {
	return true
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set StringSet) Size() int {
	return len(set)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set StringSet) Cardinality() int {
	return set.Size()
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

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set StringSet) Send() <-chan string {
	ch := make(chan string)
	go func() {
		for v := range set {
			ch <- v
		}
		close(ch)
	}()

	return ch
}

//-------------------------------------------------------------------------------------------------

// Forall applies a predicate function to every element in the set. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (set StringSet) Forall(fn func(string) bool) bool {
	for v := range set {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Exists applies a predicate function to every element in the set. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (set StringSet) Exists(fn func(string) bool) bool {
	for v := range set {
		if fn(v) {
			return true
		}
	}
	return false
}

// Foreach iterates over stringSet and executes the passed func against each element.
func (set StringSet) Foreach(fn func(string)) {
	for v := range set {
		fn(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Filter returns a new StringSet whose elements return true for func.
// The original set is not modified
func (set StringSet) Filter(fn func(string) bool) StringSet {
	result := NewStringSet()
	for v := range set {
		if fn(v) {
			result[v] = struct{}{}
		}
	}
	return result
}

// Partition returns two new stringSets whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
// The original set is not modified
func (set StringSet) Partition(p func(string) bool) (StringSet, StringSet) {
	matching := NewStringSet()
	others := NewStringSet()
	for v := range set {
		if p(v) {
			matching[v] = struct{}{}
		} else {
			others[v] = struct{}{}
		}
	}
	return matching, others
}

// Map returns a new StringSet by transforming every element with a function fn.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set StringSet) Map(fn func(string) string) StringSet {
	result := NewStringSet()

	for v := range set {
        result[fn(v)] = struct{}{}
	}

	return result
}

// FlatMap returns a new StringSet by transforming every element with a function fn that
// returns zero or more items in a slice. The resulting list may have a different size to the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set StringSet) FlatMap(fn func(string) []string) StringSet {
	result := NewStringSet()

	for v, _ := range set {
	    for _, x := range fn(v) {
            result[x] = struct{}{}
	    }
	}

	return result
}

// CountBy gives the number elements of StringSet that return true for the passed predicate.
func (set StringSet) CountBy(predicate func(string) bool) (result int) {
	for v := range set {
		if predicate(v) {
			result++
		}
	}
	return
}


//-------------------------------------------------------------------------------------------------
// These methods are included when string is ordered.

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (list StringSet) Min() string {
	return list.MinBy(func(a string, b string) bool {
		return a < b
	})
}

// Max returns the first element containing the maximum value, when compared to other elements.
// Panics if the collection is empty.
func (list StringSet) Max() (result string) {
	return list.MaxBy(func(a string, b string) bool {
		return a < b
	})
}

// MinBy returns an element of StringSet containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set StringSet) MinBy(less func(string, string) bool) string {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty list.")
	}
	var m string
	first := true
	for v := range set {
		if first {
			m = v
			first = false
		} else if less(v, m) {
			m = v
		}
	}
	return m
}

// MaxBy returns an element of StringSet containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (set StringSet) MaxBy(less func(string, string) bool) string {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty list.")
	}
	var m string
	first := true
	for v := range set {
		if first {
			m = v
			first = false
		} else if less(m, v) {
			m = v
		}
	}
	return m
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


