// A simple type derived from []string
// Not thread-safe.
//
// Generated from simple/list.tpl with Type=string
// options: Comparable:true Numeric:false Ordered:true Stringer:<no value>

package util

import (

	"math/rand"
	"sort"
)

// StringList is a slice of type string. Use it where you would use []string.
// To add items to the list, simply use the normal built-in append function.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type StringList []string

//-------------------------------------------------------------------------------------------------

func newStringList(len, cap int) StringList {
	return make(StringList, len, cap)
}

// NewStringList constructs a new list containing the supplied values, if any.
func NewStringList(values ...string) StringList {
	result := newStringList(len(values), len(values))
	copy(result, values)
	return result
}

// ConvertStringList constructs a new list containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned list will contain all the values that were correctly converted.
func ConvertStringList(values ...interface{}) (StringList, bool) {
	result := newStringList(0, len(values))

	for _, i := range values {
		v, ok := i.(string)
		if ok {
			result = append(result, v)
		}
	}

	return result, len(result) == len(values)
}

// BuildStringListFromChan constructs a new StringList from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildStringListFromChan(source <-chan string) StringList {
	result := newStringList(0, 0)
	for v := range source {
		result = append(result, v)
	}
	return result
}

// ToInterfaceSlice returns the elements of the current list as a slice of arbitrary type.
func (list StringList) ToInterfaceSlice() []interface{} {
	var s []interface{}
	for _, v := range list {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (list StringList) Clone() StringList {
	return NewStringList(list...)
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the list.
// Panics if the index is out of range.
// The simple list is a dressed-up slice and normal slice operations will also work.
func (list StringList) Get(i int) string {
	return list[i]
}

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty
func (list StringList) Head() string {
	return list[0]
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty
func (list StringList) Last() string {
	return list[len(list)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty
func (list StringList) Tail() StringList {
	return StringList(list[1:])
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty
func (list StringList) Init() StringList {
	return StringList(list[:len(list)-1])
}

// IsEmpty tests whether StringList is empty.
func (list StringList) IsEmpty() bool {
	return list.Len() == 0
}

// NonEmpty tests whether StringList is empty.
func (list StringList) NonEmpty() bool {
	return list.Len() > 0
}

// IsSequence returns true for lists.
func (list StringList) IsSequence() bool {
	return true
}

// IsSet returns false for lists.
func (list StringList) IsSet() bool {
	return false
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list - an alias of Len().
func (list StringList) Size() int {
	return len(list)
}

// Len returns the number of items in the list - an alias of Size().
// This is one of the three methods in the standard sort.Interface.
func (list StringList) Len() int {
	return len(list)
}

// Swap exchanges two elements, which is necessary during sorting etc.
// This is one of the three methods in the standard sort.Interface.
func (list StringList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

//-------------------------------------------------------------------------------------------------


// Contains determines if a given item is already in the list.
func (list StringList) Contains(v string) bool {
	return list.Exists(func (x string) bool {
		return x == v
	})
}

// ContainsAll determines if the given items are all in the list.
// This is potentially a slow method and should only be used rarely.
func (list StringList) ContainsAll(i ...string) bool {
	for _, v := range i {
		if !list.Contains(v) {
			return false
		}
	}
	return true
}

// Exists verifies that one or more elements of StringList return true for the passed func.
func (list StringList) Exists(fn func(string) bool) bool {
	for _, v := range list {
		if fn(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of StringList return true for the passed func.
func (list StringList) Forall(fn func(string) bool) bool {
	for _, v := range list {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over StringList and executes the passed func against each element.
func (list StringList) Foreach(fn func(string)) {
	for _, v := range list {
		fn(v)
	}
}

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (list StringList) Send() <-chan string {
	ch := make(chan string)
	go func() {
		for _, v := range list {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// Reverse returns a copy of StringList with all elements in the reverse order.
func (list StringList) Reverse() StringList {
	numItems := len(list)
	result := newStringList(numItems, numItems)
	last := numItems - 1
	for i, v := range list {
		result[last-i] = v
	}
	return result
}

// Shuffle returns a shuffled copy of StringList, using a version of the Fisher-Yates shuffle.
func (list StringList) Shuffle() StringList {
	result := list.Clone()
	numItems := len(list)
	for i := 0; i < numItems; i++ {
		r := i + rand.Intn(numItems-i)
		result[i], result[r] = result[r], result[i]
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of StringList containing the leading n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list StringList) Take(n int) StringList {
	if n > len(list) {
		return list
	}
	return list[0:n]
}

// Drop returns a slice of StringList without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list StringList) Drop(n int) StringList {
	if n == 0 {
		return list
	}

	l := len(list)
	if n < l {
		return list[n:]
	}
	return list[l:]
}

// TakeLast returns a slice of StringList containing the trailing n elements of the source list.
// If n is greater than the size of the list, the whole original list is returned.
func (list StringList) TakeLast(n int) StringList {
	l := len(list)
	if n > l {
		return list
	}
	return list[l-n:]
}

// DropLast returns a slice of StringList without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list StringList) DropLast(n int) StringList {
	if n == 0 {
		return list
	}

	l := len(list)
	if n > l {
		return list[l:]
	} else {
		return list[0 : l-n]
	}
}

// TakeWhile returns a new StringList containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elemense are excluded.
func (list StringList) TakeWhile(p func(string) bool) StringList {
	result := newStringList(0, 0)
	for _, v := range list {
		if p(v) {
			result = append(result, v)
		} else {
			return result
		}
	}
	return result
}

// DropWhile returns a new StringList containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elemense are added.
func (list StringList) DropWhile(p func(string) bool) StringList {
	result := newStringList(0, 0)
	adding := false

	for _, v := range list {
		if !p(v) || adding {
			adding = true
			result = append(result, v)
		}
	}

	return result
}

//-------------------------------------------------------------------------------------------------

// Find returns the first string that returns true for some function.
// False is returned if none match.
func (list StringList) Find(fn func(string) bool) (string, bool) {

	for _, v := range list {
		if fn(v) {
			return v, true
		}
	}


	var empty string
	return empty, false

}

// Filter returns a new StringList whose elements return true for func.
// The original list is not modified
func (list StringList) Filter(fn func(string) bool) StringList {
	result := newStringList(0, len(list)/2)

	for _, v := range list {
		if fn(v) {
			result = append(result, v)
		}
	}

	return result
}

// Partition returns two new stringLists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
// The original list is not modified
func (list StringList) Partition(p func(string) bool) (StringList, StringList) {
	matching := newStringList(0, len(list)/2)
	others := newStringList(0, len(list)/2)

	for _, v := range list {
		if p(v) {
			matching = append(matching, v)
		} else {
			others = append(others, v)
		}
	}

	return matching, others
}

// Map returns a new StringList by transforming every element with a function fn.
// The resulting list is the same size as the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list StringList) Map(fn func(string) string) StringList {
	result := newStringList(0, len(list))

	for _, v := range list {
		result = append(result, fn(v))
	}

	return result
}

// FlatMap returns a new StringList by transforming every element with a function fn that
// returns zero or more items in a slice. The resulting list may have a different size to the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list StringList) FlatMap(fn func(string) []string) StringList {
	result := newStringList(0, len(list))

	for _, v := range list {
		result = append(result, fn(v)...)
	}

	return result
}

// CountBy gives the number elements of StringList that return true for the passed predicate.
func (list StringList) CountBy(predicate func(string) bool) (result int) {
	for _, v := range list {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of StringList containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list StringList) MinBy(less func(string, string) bool) string {
	l := len(list)
	if l == 0 {
		panic("Cannot determine the minimum of an empty list.")
	}

	m := 0
	for i := 1; i < l; i++ {
		if less(list[i], list[m]) {
			m = i
		}
	}

	return list[m]
}

// MaxBy returns an element of StringList containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list StringList) MaxBy(less func(string, string) bool) string {
	l := len(list)
	if l == 0 {
		panic("Cannot determine the maximum of an empty list.")
	}

	m := 0
	for i := 1; i < l; i++ {
		if less(list[m], list[i]) {
			m = i
		}
	}

	return list[m]
}

// DistinctBy returns a new StringList whose elements are unique, where equality is defined by a passed func.
func (list StringList) DistinctBy(equal func(string, string) bool) StringList {
	result := newStringList(0, len(list))
Outer:
	for _, v := range list {
		for _, r := range result {
			if equal(v, r) {
				continue Outer
			}
		}
		result = append(result, v)
	}
	return result
}

// IndexWhere finds the index of the first element satisfying some predicate. If none exists, -1 is returned.
func (list StringList) IndexWhere(p func(string) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying some predicate at or after some start index.
// If none exists, -1 is returned.
func (list StringList) IndexWhere2(p func(string) bool, from int) int {
	for i, v := range list {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying some predicate.
// If none exists, -1 is returned.
func (list StringList) LastIndexWhere(p func(string) bool) int {
	return list.LastIndexWhere2(p, len(list))
}

// LastIndexWhere2 finds the index of the last element satisfying some predicate at or before some start index.
// If none exists, -1 is returned.
func (list StringList) LastIndexWhere2(p func(string) bool, before int) int {
	if before < 0 {
		before = len(list)
	}
	for i := len(list) - 1; i >= 0; i-- {
		v := list[i]
		if i <= before && p(v) {
			return i
		}
	}
	return -1
}


//-------------------------------------------------------------------------------------------------
// These methods are included when string is comparable.

// Equals determines if two lists are equal to each other.
// If they both are the same size and have the same items in the same order, they are considered equal.
// Order of items is not relevent for sets to be equal.
func (list StringList) Equals(other StringList) bool {
	if list.Size() != other.Size() {
		return false
	}

	for i, v := range list {
		if v != other[i] {
			return false
		}
	}

	return true
}

//-------------------------------------------------------------------------------------------------

type sortableStringList struct {
	less func(i, j string) bool
	m []string
}

func (sl sortableStringList) Less(i, j int) bool {
	return sl.less(sl.m[i], sl.m[j])
}

func (sl sortableStringList) Len() int {
	return len(sl.m)
}

func (sl sortableStringList) Swap(i, j int) {
	sl.m[i], sl.m[j] = sl.m[j], sl.m[i]
}

// SortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
func (list StringList) SortBy(less func(i, j string) bool) StringList {

	sort.Sort(sortableStringList{less, list})
	return list
}

// StableSortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
// The algorithm keeps the original order of equal elements.
func (list StringList) StableSortBy(less func(i, j string) bool) StringList {

	sort.Stable(sortableStringList{less, list})
	return list
}


//-------------------------------------------------------------------------------------------------
// These methods are included when string is ordered.

// Sorted alters the list so that the elements are sorted by their natural ordering.
func (list StringList) Sorted() StringList {
	return list.SortBy(func(a, b string) bool {
		return a < b
	})
}

// StableSorted alters the list so that the elements are sorted by their natural ordering.
func (list StringList) StableSorted() StringList {
	return list.StableSortBy(func(a, b string) bool {
		return a < b
	})
}

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (list StringList) Min() string {
	m := list.MinBy(func(a string, b string) bool {
		return a < b
	})
	return m
}

// Max returns the first element containing the maximum value, when compared to other elements.
// Panics if the collection is empty.
func (list StringList) Max() (result string) {
	m := list.MaxBy(func(a string, b string) bool {
		return a < b
	})
	return m
}


