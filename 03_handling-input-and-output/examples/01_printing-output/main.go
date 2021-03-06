package main

import "fmt"

func main() {
	// Println prints the value you provide to standard output and it appends a newline in the end
	fmt.Println("Hello World!")

	// If multiple arguments are specified, they will all be printed on-screen with whitespace in-between them
	fmt.Println("Hello", "World", 42, "!")

	// Print works the same way as Println but it doesn't add whitespace between arguments
	// and it doesn't append a newline character at the end of the line
	fmt.Print("Hello", "World", 42, "!")

	name := "Pres"
	var initial rune = 'P'
	age := 42
	height := 180
	hoursWatched := 1.5

	// You can use Printf with various types of placeholders to specify how you want to format the output.
	// %s - format a string
	// %c - format a character
	// %d - format a whole number
	// %f - format a number with a fractional part
	fmt.Printf("Hello Mr. %c and Welcome to the Course!\nYour full name is %v\nYour age is %d and your height is %d.\nGood for you!\n",
		initial, name, age, height)

	// You can use various options in addition to the type of output to specify different properties:
	// {number} - width the value should take. Defaults to as much as it needs
	// .{number} - how many decimal places after the dot should be shown
	fmt.Println("Lecture Breakdown:")
	fmt.Printf("Introduction to Programming | %10.2f |\n", 0.67)
	fmt.Printf("Data Types & Variables      | %10.2f |\n", 0.39)
	fmt.Printf("Handing Input & Output      | %10.2f |\n", 123.39)

	// Sprintf uses the exact same arguments & options to format output, but instead of rendering that in
	// standard output, it returns the output as a string
	result := fmt.Sprintf("You've watched %10.2f hours from the course.", hoursWatched)
	fmt.Println(result)
}
