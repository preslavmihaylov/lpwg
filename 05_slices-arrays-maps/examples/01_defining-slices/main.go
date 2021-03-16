package main

import "fmt"

func main() {
	// A slice allows you to store a
	// set of values in a single variables

	// This syntax means:
	// [] - create a slice
	// string - which has elements of type string
	// {} - and has these elements in it from the get-go
	zombies := []string{}

	// Inside the curly brackets, you can specify some initial values to
	// store in the slice
	otherZombies := []string{"Paul", "George", "Gooey", "Lucy"}

	// You can't store values of a different type in the slice
	// if it's defined as a "slice of strings"
	// nums := []string{1, 2, 3, 4}

	// To store values of a different type, you'll need to
	// initialize the slice with that data type in mind
	nums := []int{1, 2, 3, 4}

	fmt.Println(zombies)      // []
	fmt.Println(otherZombies) // [Paul George Gooey Lucy]
	fmt.Println(nums)         // [1 2 3 4]
}
