package main

import "fmt"

func main() {
	num1 := 14
	num2 := 24

	sum := num1 + num2
	subtracted := num1 - num2
	multiplied := num1 * num2
	divided := 56 / 8
	modulus := 15 % 2 // this results in the remainder from dividing 15 by 2
	fmt.Println(sum, subtracted, multiplied, divided, modulus)

	// the values haven't changed from any of the previous operations
	fmt.Println(num1, num2)

	// these operations change the values stored inside the variables
	num1++ // 14 + 1
	num2-- // 24 - 1
	fmt.Println(num1, num2)

	// When you add a number to a variable which results in a value which can't be stored in it, you get an overflow
	// When that happens, you get the minimum value which can be stored in that number instead
	var smallInt int8 = 100
	smallInt = smallInt + 30 // example of overflow
	fmt.Println(smallInt)

	// Underflow is the opposite of overflow and the same properties hold
	var verySmallInt int8 = -127
	verySmallInt = verySmallInt - 2 // example of underflow
	fmt.Println(verySmallInt)

	// Oveflow can happen with multiplication as well
	var toMultiply int8 = 60
	toMultiply = toMultiply * 3 // example of overflow via multiplication
	fmt.Println(toMultiply)

	// When dividing whole numbers, you get whole numbers as a result
	// Any fractional part is lost in the process
	var toDivide int8 = 15
	toDivide = toDivide / 2 // example of whole number division
	fmt.Println(toDivide)   // results in 7

	// The modulus operator gives you the remainder of a division
	var toMod int8 = 15
	toMod = toMod % 2
	fmt.Println(toMod) // results in 1. 15 / 2 = 7 with remainder of 1

	var toModNegative int8 = -15
	toModNegative = toModNegative % -2
	fmt.Println(toModNegative) // results in -1. -15 / -2 = 7 with remainder of -1

	// Dividing a number by zero is prohibited and results in a compilation error
	// var divByZero int8 = 10
	// divByZero = divByZero / 0
	// fmt.Println(divByZero)
}
