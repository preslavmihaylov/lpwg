package main

import (
	"fmt"
	"strings"
)

func main() {
	// == checks if the two objects are equal
	fmt.Println(42 == 32) // false

	// != checks if the objects are different
	fmt.Println("Hello There" != "Hello There ") // true

	// < and > check that the given number is strictly greater than/less than what's given
	fmt.Println(32 > 13) // true
	fmt.Println(32 < 13)

	// <= and >= check that the numbers are greater than/less than or equal
	fmt.Println(13 >= 13) // true
	fmt.Println(13 <= 42) // true

	// You can get boolean results from various functions as well
	message := "Hello there Mr. Smith!"
	fmt.Println(strings.Contains(message, "Hello there Ms."))
}
