package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	// You can take parts of a slice using the syntax numbers[start:end]
	// This results in the elements at indices [start,end)
	subslice1 := numbers[0:2]
	fmt.Println(subslice1)   // [1 2]
	fmt.Println(numbers[:2]) // a shorthand for the same thing

	subslice := numbers[2:len(numbers)]
	fmt.Println(subslice)    // [3 4 5 6]
	fmt.Println(numbers[2:]) // a shorthand for the same thing

	// Note that taking subslices doesn't
	// make copies of the original elements.
	// It just creates a reference to them
	subslice[0] = 42
	fmt.Println(numbers)  // [1 2 42 4 5 6]
	fmt.Println(subslice) // [42 4 5 6]
}
