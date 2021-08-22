package main

import "fmt"

func main() {
	num := 42

	// the normal way you assign values makes a copy
	numCpy := num

	// making a copy allows you to independently modify values in memory
	numCpy = 56
	fmt.Println(num, numCpy) // 42 56

	// pointers, on the other hand, allow you to copy the address
	// of a value, not the value directly
	var numPtr *int

	// this operation takes the reference/address
	// of the value stored in num
	numPtr = &num

	// * in front of a variable means "dereferencing"
	// That is, taking the value which is being pointed to and
	// changing it
	*numPtr = 56
	fmt.Println(num, *numPtr) // 56, 56
}
