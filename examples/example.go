package main

import (
	"fmt"

	sortedmap "github.com/utkukrl/caelestius"
)

func main() {
	sm := sortedmap.NewSortedMap()
	sm.Add(3, "three")
	sm.Add(1, "one")
	sm.Add(4, "four")
	sm.Add(2, "two")

	fmt.Println("Keys:", sm.Keys())
	fmt.Println("Values:", sm.Values())

	sm.Remove(3)
	key := 2
	value := sm.Get(key)
	fmt.Printf("Anahtar %d için değer: %v\n", key, value)
}
