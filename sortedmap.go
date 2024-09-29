package sortedmap

// SortedMap struct
type SortedMap struct {
	keys   []interface{}
	values map[interface{}]interface{}
}

// NewSortedMap creates a new SortedMap.
func NewSortedMap() *SortedMap {
	return &SortedMap{
		keys:   make([]interface{}, 0),
		values: make(map[interface{}]interface{}),
	}
}

// Add adds a new key-value pair to the SortedMap.
// If the key already exists, it updates the value.
func (sm *SortedMap) Add(key, value interface{}) {
	if _, exists := sm.values[key]; !exists {
		sm.keys = append(sm.keys, key)
	}
	sm.values[key] = value
}

// Remove removes the specified key from the SortedMap.
func (sm *SortedMap) Remove(key interface{}) {
	if _, exists := sm.values[key]; exists {
		// Remove the key from the values map
		delete(sm.values, key)
		// Find and remove the key from the keys slice
		for i, k := range sm.keys {
			if k == key {
				sm.keys = append(sm.keys[:i], sm.keys[i+1:]...)
				break
			}
		}
	}
}

// Get returns the value associated with the specified key.
func (sm *SortedMap) Get(key interface{}) interface{} {
	return sm.values[key]
}

// Keys returns all keys in the SortedMap in the order they were added.
func (sm *SortedMap) Keys() []interface{} {
	return sm.keys
}

// Values returns all values in the SortedMap in the order of their keys.
func (sm *SortedMap) Values() []interface{} {
	values := make([]interface{}, len(sm.keys))
	for i, key := range sm.keys {
		values[i] = sm.values[key]
	}
	return values
}
