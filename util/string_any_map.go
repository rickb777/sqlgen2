// A simple type derived from map[string]interface{}.
// Not thread-safe.
//
// Generated from simple/map.tpl with Key=string Type=interface{}
// options: Comparable:<no value> Stringer:<no value> KeyList:<no value> ValueList:<no value> Mutable:always

package util

// StringAnyMap is the primary type that represents a map
type StringAnyMap map[string]interface{}

// StringAnyTuple represents a key/value pair.
type StringAnyTuple struct {
	Key string
	Val interface{}
}

// StringAnyTuples can be used as a builder for unmodifiable maps.
type StringAnyTuples []StringAnyTuple

func (ts StringAnyTuples) Append1(k string, v interface{}) StringAnyTuples {
	return append(ts, StringAnyTuple{k, v})
}

func (ts StringAnyTuples) Append2(k1 string, v1 interface{}, k2 string, v2 interface{}) StringAnyTuples {
	return append(ts, StringAnyTuple{k1, v1}, StringAnyTuple{k2, v2})
}

func (ts StringAnyTuples) Append3(k1 string, v1 interface{}, k2 string, v2 interface{}, k3 string, v3 interface{}) StringAnyTuples {
	return append(ts, StringAnyTuple{k1, v1}, StringAnyTuple{k2, v2}, StringAnyTuple{k3, v3})
}

//-------------------------------------------------------------------------------------------------

func newStringAnyMap() StringAnyMap {
	return StringAnyMap(make(map[string]interface{}))
}

// NewStringAnyMap creates and returns a reference to a map containing one item.
func NewStringAnyMap1(k string, v interface{}) StringAnyMap {
	mm := newStringAnyMap()
	mm[k] = v
	return mm
}

// NewStringAnyMap creates and returns a reference to a map, optionally containing some items.
func NewStringAnyMap(kv ...StringAnyTuple) StringAnyMap {
	mm := newStringAnyMap()
	for _, t := range kv {
		mm[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm StringAnyMap) Keys() []string {
	var s []string
	for k, _ := range mm {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm StringAnyMap) Values() []interface{} {
	var s []interface{}
	for _, v := range mm {
		s = append(s, v)
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm StringAnyMap) ToSlice() []StringAnyTuple {
	var s []StringAnyTuple
	for k, v := range mm {
		s = append(s, StringAnyTuple{k, v})
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm StringAnyMap) Get(k string) (interface{}, bool) {
	v, found := mm[k]
	return v, found
}

// Put adds an item to the current map, replacing interface{} prior value.
func (mm StringAnyMap) Put(k string, v interface{}) bool {
	_, found := mm[k]
	mm[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm StringAnyMap) ContainsKey(k string) bool {
	_, found := mm[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm StringAnyMap) ContainsAllKeys(kk ...string) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *StringAnyMap) Clear() {
	*mm = make(map[string]interface{})
}

// Remove a single item from the map.
func (mm StringAnyMap) Remove(k string) {
	delete(mm, k)
}

// Pop removes a single item from the map, returning the value present until removal.
func (mm StringAnyMap) Pop(k string) (interface{}, bool) {
	v, found := mm[k]
	delete(mm, k)
	return v, found
}

// Size returns how minterface{} items are currently in the map. This is a synonym for Len.
func (mm StringAnyMap) Size() int {
	return len(mm)
}

// IsEmpty returns true if the map is empty.
func (mm StringAnyMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm StringAnyMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
func (mm StringAnyMap) DropWhere(fn func(string, interface{}) bool) StringAnyTuples {
	removed := make(StringAnyTuples, 0)
	for k, v := range mm {
		if fn(k, v) {
			removed = append(removed, StringAnyTuple{k, v})
			delete(mm, k)
		}
	}
	return removed
}

// Foreach applies a function to every element in the map.
// The function can safely alter the values via side-effects.
func (mm StringAnyMap) Foreach(fn func(string, interface{})) {
	for k, v := range mm {
		fn(k, v)
	}
}

// Forall applies a predicate function to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm StringAnyMap) Forall(fn func(string, interface{}) bool) bool {
	for k, v := range mm {
		if !fn(k, v) {
			return false
		}
	}
	return true
}

// Exists applies a predicate function to every element in the map. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (mm StringAnyMap) Exists(fn func(string, interface{}) bool) bool {
	for k, v := range mm {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Find returns the first interface{} that returns true for some function.
// False is returned if none match.
// The original map is not modified
func (mm StringAnyMap) Find(fn func(string, interface{}) bool) (StringAnyTuple, bool) {
	for k, v := range mm {
		if fn(k, v) {
			return StringAnyTuple{k, v}, true
		}
	}

	return StringAnyTuple{}, false
}

// Filter applies a predicate function to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified
func (mm StringAnyMap) Filter(fn func(string, interface{}) bool) StringAnyMap {
	result := NewStringAnyMap()
	for k, v := range mm {
		if fn(k, v) {
			result[k] = v
		}
	}
	return result
}

// Partition applies a predicate function to every element in the map. It divides the map into two copied maps,
// the first containing all the elements for which the predicate returned true, and the second containing all
// the others.
// The original map is not modified
func (mm StringAnyMap) Partition(fn func(string, interface{}) bool) (matching StringAnyMap, others StringAnyMap) {
	matching = NewStringAnyMap()
	others = NewStringAnyMap()
	for k, v := range mm {
		if fn(k, v) {
			matching[k] = v
		} else {
			others[k] = v
		}
	}
	return
}

// Map returns a new AnyMap by transforming every element with a function fn.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm StringAnyMap) Map(fn func(string, interface{}) (string, interface{})) StringAnyMap {
	result := NewStringAnyMap()

	for k1, v1 := range mm {
	    k2, v2 := fn(k1, v1)
	    result[k2] = v2
	}

	return result
}

// FlatMap returns a new AnyMap by transforming every element with a function fn that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm StringAnyMap) FlatMap(fn func(string, interface{}) []StringAnyTuple) StringAnyMap {
	result := NewStringAnyMap()

	for k1, v1 := range mm {
	    ts := fn(k1, v1)
	    for _, t := range ts {
            result[t.Key] = t.Val
	    }
	}

	return result
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm StringAnyMap) Clone() StringAnyMap {
	result := NewStringAnyMap()
	for k, v := range mm {
		result[k] = v
	}
	return result
}


