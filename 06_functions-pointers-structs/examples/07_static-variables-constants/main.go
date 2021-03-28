package main

import "fmt"

// When a variable is declared outside a function it is called a "stati variable"
// They are accessible from any function and all changes to it are reflected everywhere.
// They're also referred to as "global variables"
var number = 42

func main() {
	updateNumber()
	fmt.Println(number) // 14
}

func updateNumber() {
	// The change to number in this function will be reflected in all other functions
	number = 14
	fmt.Println(number) // 14
}
