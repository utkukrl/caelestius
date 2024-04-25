package sortedmap

import (
	"fmt"
	"testing"
)

func TestSortedMap(t *testing.T) {
	t.Run("IntType", func(t *testing.T) {
		sm := NewSortedMap(func(i, j interface{}) bool {
			return i.(int) < j.(int)
		})

		sm.Add(3, 30.5)
		sm.Add(1, 10.5)
		sm.Add(2, 20.5)

		if sm.Get(1) != 10.5 {
			t.Errorf("Expected value for key 1 is 10.5, but got %f", sm.Get(1))
		}

		sm.Remove(2)
		if sm.Get(2) != nil {
			t.Errorf("Expected value for removed key 2 is nil, but got %v", sm.Get(2))
		}

		expectedKeys := []interface{}{1, 3}
		keys := sm.Keys()
		for i, key := range keys {
			if key != expectedKeys[i] {
				t.Errorf("Expected key %v but got %v", expectedKeys[i], key)
			}
		}

		expectedValues := []interface{}{10.5, 30.5}
		values := sm.Values()
		for i, value := range values {
			if value != expectedValues[i] {
				t.Errorf("Expected value %v but got %v", expectedValues[i], value)
			}
		}
		t.Run("IterationOrder", func(t *testing.T) {
			for i, key := range sm.Keys() {
				if key != expectedKeys[i] {
					t.Errorf("Expected key %v at index %d but got %v", expectedKeys[i], i, key)
				}
			}
		})
	})

	t.Run("FloatType", func(t *testing.T) {
		sm := NewSortedMap(func(i, j interface{}) bool {
			return i.(float64) < j.(float64)
		})

		sm.Add(3.5, "three point five")
		sm.Add(1.5, "one point five")
		sm.Add(2.5, "two point five")

		if sm.Get(1.5) != "one point five" {
			t.Errorf("Expected value for key 1.5 is 'one point five', but got %s", sm.Get(1.5))
		}

		sm.Remove(2.5)
		if sm.Get(2.5) != nil {
			t.Errorf("Expected value for removed key 2.5 is nil, but got %v", sm.Get(2.5))
		}

		expectedKeys := []interface{}{1.5, 3.5}
		keys := sm.Keys()
		for i, key := range keys {
			if key != expectedKeys[i] {
				t.Errorf("Expected key %v but got %v", expectedKeys[i], key)
			}
		}

		expectedValues := []interface{}{"one point five", "three point five"}
		values := sm.Values()
		for i, value := range values {
			if value != expectedValues[i] {
				t.Errorf("Expected value %v but got %v", expectedValues[i], value)
			}
		}
		t.Run("IterationOrder", func(t *testing.T) {
			for i, key := range sm.Keys() {
				if key != expectedKeys[i] {
					t.Errorf("Expected key %v at index %d but got %v", expectedKeys[i], i, key)
				}
			}
		})
	})

	t.Run("StringType", func(t *testing.T) {
		sm := NewSortedMap(func(i, j interface{}) bool {
			return i.(string) < j.(string)
		})

		sm.Add("c", 30)
		sm.Add("a", 10)
		sm.Add("b", 20)
		sm.Add("d", 40)

		if sm.Get("a") != 10 {
			t.Errorf("Expected value for key 'a' is 10, but got %d", sm.Get("a"))
		}

		sm.Remove("b")
		if sm.Get("b") != nil {
			t.Errorf("Expected value for removed key 'b' is nil, but got %v", sm.Get("b"))
		}

		expectedKeys := []interface{}{"a", "c", "d"}
		keys := sm.Keys()
		for i, key := range keys {
			if key != expectedKeys[i] {
				t.Errorf("Expected key %v but got %v", expectedKeys[i], key)
			}
		}

		expectedValues := []interface{}{10, 30, 40}
		values := sm.Values()
		for i, value := range values {
			if value != expectedValues[i] {
				t.Errorf("Expected value %v but got %v", expectedValues[i], value)
			}
		}
		t.Run("IterationOrder", func(t *testing.T) {
			for i, key := range sm.Keys() {
				if key != expectedKeys[i] {
					t.Errorf("Expected key %v at index %d but got %v", expectedKeys[i], i, key)
				}
			}
		})
		fmt.Println(sm.Keys())
		fmt.Println(sm.Values())
	})
}
