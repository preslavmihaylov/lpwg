package main

import (
	"fmt"
	"math"
)

func main() {
	delta := 0.00005
	a, b := 3.100005, 3.100003
	absDiff := math.Abs(a - b)
	if absDiff < delta {
		fmt.Println(a, "is equal to", b)
	} else {
		fmt.Println(a, "is not equal to", b)
	}

	a, b = 1.12003, 1.12009
	absDiff = math.Abs(a - b)
	if absDiff < delta {
		fmt.Println(a, "is equal to", b)
	} else {
		fmt.Println(a, "is not equal to", b)
	}
}
