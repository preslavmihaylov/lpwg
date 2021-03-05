package main

import (
	"fmt"
)

func main() {
	// Whenever a variable is declared but no initial value is assigned to it, that variable will get its corresponding zero value
	// This is the default value for the given data type
	var message string // ""
	var isTrue bool    // false
	var num int        // 0

	fmt.Println(message)
	fmt.Println(isTrue)
	fmt.Println(num)
}
