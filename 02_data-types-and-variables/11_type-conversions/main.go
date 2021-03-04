package main

import (
	"fmt"
	"strconv"
)

func main() {
	// In Go, if you want to convert a value from one data type to another,
	// you have to use type conversions
	var intNum int8 = 42
	var biggerInt int16 = int16(intNum)
	fmt.Println(biggerInt)

	// Converting from some types to other is not possible
	// For example, converting from a string to an integer is not possible using type conversion
	// var message string = "Hello"
	// var aNum int = int(message)
	// fmt.Println(aNum)

	// When converting from a bigger type to a smaller one, it is possible to lose some data in the process
	var bigNum int16 = 300
	var smallNum int8 = int8(bigNum)
	fmt.Println(smallNum) // results in 44

	// Integers can be safely converted to real numbers without losing any data
	var anInt int16 = 42
	var aFloat float32 = float32(anInt)
	fmt.Println(aFloat)

	// However, when converting from a float to an integer
	// you can lose some data in the process. In this case, you'd lose the fractional 0.5 from 3.5
	aFloat = 3.5
	var anotherInt int = int(aFloat)
	fmt.Println(anotherInt)

	// var numAsString string = "42"
	// var num int = int(numAsString)
	// fmt.Println(num)

	num, err := strconv.Atoi("42")
	fmt.Println(num, err)

	var numAsString string = strconv.Itoa(num)
	fmt.Println(numAsString)
}
