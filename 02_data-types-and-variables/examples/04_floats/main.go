package main

import "fmt"

func main() {
	// floating-point data types are used to represent real numbers
	// Since real numbers can be infinitely big and computers have a finite memory,
	// we can't represent all real numbers from real life
	//
	// But we can represent them up to a precision which is sufficient

	// The difference between float32 and float64 is in the preciseness with which
	// a real number can be represented. In simpler words - the # of digits in the number
	var num1 float32 = 3.1475634353453453453545454354353453453453453453453354
	fmt.Println(num1) // prints 3.1475635

	var num2 float64 = 3.1475634353453453453545454354353453453453453453453354
	fmt.Println(num2) // prints 3.1475634353453454
}
