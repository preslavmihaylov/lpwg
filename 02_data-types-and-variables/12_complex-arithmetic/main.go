package main

import (
	"fmt"
	"math"
)

func main() {
	// Complex arithmetic operations can be found in the "math" package
	// Explore it at https://golang.org/pkg/math/#Abs

	// Calculate Square-Root of a number
	var num int = 21
	result := math.Sqrt(float64(num))
	fmt.Println(result)

	// Calculate the Absolute value
	num = -14
	fmt.Println(math.Abs(float64(num)))

	// Raise x to the power of y
	num = 9
	fmt.Println(math.Pow(float64(num), 3))
}
