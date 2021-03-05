package main

import "fmt"

func main() {
	var a int = 32
	var b int = 14

	// First way - using a third variable:
	// oldA := a
	// a = b
	// b = oldA

	// Second way - using Go's multiple assignment:
	// a, b = b, a

	// Third way - using the sum of two numbers:
	a = b + a // (a + b)
	b = a - b // (a + b) - b = a
	a = a - b // (a + b) - a = b

	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
