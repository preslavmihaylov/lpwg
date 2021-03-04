package main

import "fmt"

func main() {
	// shorthand for
	// var message string = "I love Go!"
	message := "I love Go!"

	// another way would be:
	// var message string
	// message = "I love Go!"

	fmt.Println(message)

	// Variables are meant to store values which can change over time
	// Here, I'm changing the text which is stored in "message"
	message = "because it is a great programming language"

	// Although I'm executing the exact same instruction
	// as the one on line 14, the output is different because
	// the value stored in "message" is different
	fmt.Println(message)

	// This won't work because we're redeclaring a variable with the same name we've used before:
	// var message string
	// message = "another variable"
	// fmt.Println(message)

	// variable names should be descriptive.
	// This makes it easier for you and your collaborators to read the code in the future

	// Example of a non-descriptive variable name:
	// a := "Peter Jackson"

	// Example of a descriptive variable name:
	firstAndLastName := "Peter Jackson"

	fmt.Println(firstAndLastName)

	// What makes a valid variable name?
	//
	// characters from a-z (capital case as well)
	// numbers (as long as it's not in the beginning)
	// underscores

	// Valid variable names: num, num1, num2, n1n2n3, _n1, _n2, n_2
	// Invalid variable names: 1num, 2num, 3num, &*^&%num

	// Conventions for writing variables:
	//
	// camelCase -> firstName, lastName ...
	// PascalCase -> FirstName, LastName ...
	// snake_case -> first_name, last_name ...
	// UPPER_CASE -> FIRST_NAME, LAST_NAME ...
}
