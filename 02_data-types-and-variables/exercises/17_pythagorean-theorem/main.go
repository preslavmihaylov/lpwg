package main

import (
	"fmt"
	"math"
)

func main() {
	// c = sqrt(a^2 + b^2)
	a, b := 3.0, 4.0
	c := math.Sqrt(a*a + b*b)
	fmt.Println(c)

	a2, b2 := 6.0, 7.0
	c2 := math.Sqrt(a2*a2 + b2*b2)
	fmt.Println(c2)
}
