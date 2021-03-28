package main

import "fmt"

func main() {
	start, end := 10, 20
	iterate(&start, &end)
	fmt.Println(start, end) // 20 20

	start, end = 10, 20
	iterateNoPtr(start, end)
	fmt.Println(start, end) // 10 20
}

// Passing values as pointers allows you to
// modify the original variables passed into the function
func iterate(start, end *int) {
	for *start < *end {
		fmt.Println(*start)
		(*start)++
	}
}

// If you're not using pointers, you're copying values instead.
// Hence, any changes you make inside the function only change
// the function parameters, but not the original variables
func iterateNoPtr(start, end int) {
	for start < end {
		fmt.Println(start)
		start++
	}
}
