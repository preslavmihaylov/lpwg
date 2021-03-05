package main

import "fmt"

func main() {
	// Booleans are used to represent two values - "true" or "false"
	// These can be used to express conditions in our code - e.g. If the person is named "John", print their name
	var isCool bool = false

	// Runes represent symbols which have a corresponding numeric representation in-memory
	// The mapping between a number and a symbol is defined by the Unicode table
	var ch rune = '😼'

	// Strigns represent text, which is essentially a set of characters, which are attached to one-another
	var message string = "Go is amazing! 😼"
	fmt.Println(isCool, ch, message)
}
