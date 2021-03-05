package main

import "fmt"

func main() {
	var ch rune = '5'
	fmt.Println(ch)

	// Characters have the same properties as whole numbers because they are whole numbers in memory
	// All operations on characters work the same way they do on whole numbers
	ch = ch + 1
	fmt.Println(ch)
	fmt.Println(ch / 2)
	fmt.Println(ch * 2)
	fmt.Println(ch % 2)

	// The only operation for strings available is called "concatenation"
	// It allows us to "glue" two strings together
	var message string = "Go is"
	message = message + " awesome!"
	fmt.Println(message)

	// Concatenation only works for strings.
	// If you want to append other types to a string, you'll need to convert them to a string first
	// fmt.Println("The symbol is " + ch)

	// If you want to specify special symbols in your strings, you have to escape them using a backslash - "\"
	fmt.Println("Go is\n\t \\awesome\\")
	fmt.Println("Those kids called me \"cool\". What does that mean?")
}
