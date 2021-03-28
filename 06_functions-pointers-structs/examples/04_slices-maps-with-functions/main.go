package main

import "fmt"

func main() {
	// Maps behave as if you're always using a pointer
	// This is so because most Go programs always use maps as pointers
	// so the Go maintainers have made this change so that it's more convenient
	mp := map[string]int{"foo": 1}
	updateMap(mp)
	fmt.Println(mp) // "foo": 42, "bar": 56

	// When you pass a slice as a function argument,
	// you're passing a view of the original slice, not a copy
	// changing any of the slice's values inside the function
	// will change the original slice as well
	vals := []int{1, 2, 3, 4, 5}
	updateSlice(vals)
	fmt.Println(vals) // 2, 3, 4, 5, 6

	// However, if you add/remove elements from a slice
	// in a function, you're creating a new slice,
	// not modifying the original one
	vals = []int{1, 2, 3}
	addElement(vals, 4) // doesn't do anything to vals
	fmt.Println(vals)   // 1, 2, 3

	// If you want to add/remove elements from a slice,
	// consider returning it as a result from the function instead
	vals = addElementV2(vals, 4)
	fmt.Println(vals) // 1, 2, 3, 4
}

func updateSlice(vals []int) {
	for i := range vals {
		vals[i] += 1
	}
}

func addElement(vals []int, elem int) {
	vals = append(vals, elem)
}

func addElementV2(vals []int, elem int) []int {
	return append(vals, elem)
}

func updateMap(mp map[string]int) {
	mp["foo"] = 42
	mp["bar"] = 56
}
