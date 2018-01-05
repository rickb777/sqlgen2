// A simple type derived from map[Type]struct{}
// Not thread-safe.
//
// Generated from simple/set.tpl with Type=Type
// options: Numeric:<no value> Stringer:<no value> Mutable:always

package schema

// TypeSet is the primary type that represents a set
type TypeSet map[Type]struct{}

// NewTypeSet creates and returns a reference to an empty set.
func NewTypeSet(values ...Type) TypeSet {
	set := make(TypeSet)
	for _, i := range values {
		set[i] = struct{}{}
	}
	return set
}

// ConvertTypeSet constructs a new set containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
func ConvertTypeSet(values ...interface{}) (TypeSet, bool) {
	set := make(TypeSet)

	for _, i := range values {
		v, ok := i.(Type)
		if ok {
		    set[v] = struct{}{}
		}
	}

	return set, len(set) == len(values)
}

// BuildTypeSetFromChan constructs a new TypeSet from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildTypeSetFromChan(source <-chan Type) TypeSet {
	set := make(TypeSet)
	for v := range source {
		set[v] = struct{}{}
	}
	return set
}

// ToSlice returns the elements of the current set as a slice.
func (set TypeSet) ToSlice() []Type {
	var s []Type
	for v := range set {
		s = append(s, v)
	}
	return s
}

// ToInterfaceSlice returns the elements of the current set as a slice of arbitrary type.
func (set TypeSet) ToInterfaceSlice() []interface{} {
	var s []interface{}
	for v := range set {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (set TypeSet) Clone() TypeSet {
	clonedSet := NewTypeSet()
	for v := range set {
		clonedSet.doAdd(v)
	}
	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set TypeSet) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set TypeSet) NonEmpty() bool {
	return set.Size() > 0
}

// IsSequence returns true for lists.
func (set TypeSet) IsSequence() bool {
	return false
}

// IsSet returns false for lists.
func (set TypeSet) IsSet() bool {
	return true
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set TypeSet) Size() int {
	return len(set)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set TypeSet) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

// Add adds items to the current set, returning the modified set.
func (set TypeSet) Add(i ...Type) TypeSet {
	for _, v := range i {
		set.doAdd(v)
	}
	return set
}

func (set TypeSet) doAdd(i Type) {
	set[i] = struct{}{}
}

// Contains determines if a given item is already in the set.
func (set TypeSet) Contains(i Type) bool {
	_, found := set[i]
	return found
}

// ContainsAll determines if the given items are all in the set
func (set TypeSet) ContainsAll(i ...Type) bool {
	for _, v := range i {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

//-------------------------------------------------------------------------------------------------

// IsSubset determines if every item in the other set is in this set.
func (set TypeSet) IsSubset(other TypeSet) bool {
	for v := range set {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

// IsSuperset determines if every item of this set is in the other set.
func (set TypeSet) IsSuperset(other TypeSet) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set TypeSet) Append(more ...Type) TypeSet {
	unionedSet := set.Clone()
	for _, v := range more {
		unionedSet.doAdd(v)
	}
	return unionedSet
}

// Union returns a new set with all items in both sets.
func (set TypeSet) Union(other TypeSet) TypeSet {
	unionedSet := set.Clone()
	for v := range other {
		unionedSet.doAdd(v)
	}
	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set TypeSet) Intersect(other TypeSet) TypeSet {
	intersection := NewTypeSet()
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
func (set TypeSet) Difference(other TypeSet) TypeSet {
	differencedSet := NewTypeSet()
	for v := range set {
		if !other.Contains(v) {
			differencedSet.doAdd(v)
		}
	}
	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set TypeSet) SymmetricDifference(other TypeSet) TypeSet {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Clear clears the entire set to be the empty set.
func (set *TypeSet) Clear() {
	*set = NewTypeSet()
}

// Remove allows the removal of a single item from the set.
func (set TypeSet) Remove(i Type) {
	delete(set, i)
}

//-------------------------------------------------------------------------------------------------

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set TypeSet) Send() <-chan Type {
	ch := make(chan Type)
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
func (set TypeSet) Forall(fn func(Type) bool) bool {
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
func (set TypeSet) Exists(fn func(Type) bool) bool {
	for v := range set {
		if fn(v) {
			return true
		}
	}
	return false
}

// Foreach iterates over TypeSet and executes the passed func against each element.
func (set TypeSet) Foreach(fn func(Type)) {
	for v := range set {
		fn(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Filter returns a new TypeSet whose elements return true for func.
func (set TypeSet) Filter(fn func(Type) bool) TypeSet {
	result := NewTypeSet()
	for v := range set {
		if fn(v) {
			result[v] = struct{}{}
		}
	}
	return result
}

// Partition returns two new TypeLists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (set TypeSet) Partition(p func(Type) bool) (TypeSet, TypeSet) {
	matching := NewTypeSet()
	others := NewTypeSet()
	for v := range set {
		if p(v) {
			matching[v] = struct{}{}
		} else {
			others[v] = struct{}{}
		}
	}
	return matching, others
}

// CountBy gives the number elements of TypeSet that return true for the passed predicate.
func (set TypeSet) CountBy(predicate func(Type) bool) (result int) {
	for v := range set {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of TypeSet containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set TypeSet) MinBy(less func(Type, Type) bool) Type {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty list.")
	}
	var m Type
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

// MaxBy returns an element of TypeSet containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (set TypeSet) MaxBy(less func(Type, Type) bool) Type {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty list.")
	}
	var m Type
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
func (set TypeSet) Equals(other TypeSet) bool {
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


