package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}

	// You can use a normal for-loop with three variables
	// to iterate a slice
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Accessing", numbers[i])
	}

	// But using the for-range loop is more intuitive.
	// It returns two variables - the index and the value at that index
	for i, element := range numbers {
		fmt.Println("Incrementing", element)

		// if you want to modify the elements in the slice
		// using for-range, you should use the index to access
		// the element you need
		numbers[i] += 1

		// changing the value returned from for-range directly
		// won't work as this is a copy of the original value.
		// This operation will change the variable itself
		// but it won't affect the value stored in the slice
		element += 1
	}

	fmt.Println(numbers)
}
