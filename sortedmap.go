package sortedmap

import (
	"sort"
)

type SortedMap struct {
	keys   []interface{}
	values map[interface{}]interface{}
	less   func(i, j interface{}) bool
}

// NewSortedMap creates a new SortedMap and returns it.
// It requires a comparison function to specify the sorting order.
func NewSortedMap(less func(i, j interface{}) bool) *SortedMap {
	return &SortedMap{
		keys:   make([]interface{}, 0),
		values: make(map[interface{}]interface{}),
		less:   less,
	}
}

// Add adds a new key-value pair to the SortedMap.
// If the key already exists, it updates the value.
func (sm *SortedMap) Add(key, value interface{}) {
	index := sort.Search(len(sm.keys), func(i int) bool { return !sm.less(sm.keys[i], key) })
	if index < len(sm.keys) && !sm.less(sm.keys[index], key) && !sm.less(key, sm.keys[index]) {
		sm.values[key] = value
	} else {
		sm.keys = append(sm.keys, nil)
		copy(sm.keys[index+1:], sm.keys[index:])
		sm.keys[index] = key
		sm.values[key] = value
	}
}

// Remove removes the specified key from the SortedMap.
func (sm *SortedMap) Remove(key interface{}) {
	index := sort.Search(len(sm.keys), func(i int) bool { return !sm.less(sm.keys[i], key) })
	if index < len(sm.keys) && !sm.less(sm.keys[index], key) && !sm.less(key, sm.keys[index]) {
		copy(sm.keys[index:], sm.keys[index+1:])
		sm.keys = sm.keys[:len(sm.keys)-1]
		delete(sm.values, key)
	}
}

// Get returns the value associated with the specified key.
func (sm *SortedMap) Get(key interface{}) interface{} {
	return sm.values[key]
}

// Keys returns all keys in the SortedMap in sorted order.
func (sm *SortedMap) Keys() []interface{} {
	return sm.keys
}

// Values returns all values in the SortedMap in sorted order of keys.
func (sm *SortedMap) Values() []interface{} {
	values := make([]interface{}, len(sm.keys))
	for i, key := range sm.keys {
		values[i] = sm.values[key]
	}
	return values
}
