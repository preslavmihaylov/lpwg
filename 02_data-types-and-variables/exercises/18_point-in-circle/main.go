package main

import (
	"fmt"
	"math"
)

func main() {
	// sqrt(x^2 + y^2) <= radius

	radius := 4.0
	x1, y1 := -2.0, 1.0
	isInside := math.Sqrt((x1*x1)+(y1*y1)) <= radius
	fmt.Println(x1, ",", y1, "->", isInside)

	x2, y2 := 1.0, 4.0
	isInside = math.Sqrt((x2*x2)+(y2*y2)) <= radius
	fmt.Println(x2, ",", y2, "->", isInside)

	x3, y3 := 2.0, 2.0
	isInside = math.Sqrt((x3*x3)+(y3*y3)) <= radius
	fmt.Println(x3, ",", y3, "->", isInside)

	x4, y4 := -2.0, -2.0
	isInside = math.Sqrt((x4*x4)+(y4*y4)) <= radius
	fmt.Println(x4, ",", y4, "->", isInside)

	x5, y5 := 4.0, 0.0
	isInside = math.Sqrt((x5*x5)+(y5*y5)) <= radius
	fmt.Println(x5, ",", y5, "->", isInside)

	x6, y6 := 0.0, 4.0
	isInside = math.Sqrt((x6*x6)+(y6*y6)) <= radius
	fmt.Println(x6, ",", y6, "->", isInside)

	x7, y7 := 6.0, -2.0
	isInside = math.Sqrt((x7*x7)+(y7*y7)) <= radius
	fmt.Println(x7, ",", y7, "->", isInside)

	x8, y8 := -2.0, -2.0
	isInside = math.Sqrt((x8*x8)+(y8*y8)) <= radius
	fmt.Println(x8, ",", y8, "->", isInside)

	x9, y9 := -2.96, -2.69
	isInside = math.Sqrt((x9*x9)+(y9*y9)) <= radius
	fmt.Println(x9, ",", y9, "->", isInside)

	x10, y10 := -2.96, -2.9
	isInside = math.Sqrt((x10*x10)+(y10*y10)) <= radius
	fmt.Println(x10, ",", y10, "->", isInside)
}
