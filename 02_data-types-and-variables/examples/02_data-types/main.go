package main

import (
	"fmt"
	"time"
)

func main() {
	// text - Hello World
	var message string = "Go is amazing!"
	fmt.Println(message)

	// whole numbers - 14
	var number int = 14
	fmt.Println(number)

	// real numbers - 1.43
	var realNum float32 = 1.43
	fmt.Println(realNum)

	// characters - S, !
	var ch rune = 'S'
	fmt.Println(ch)

	// boolean - true or false
	var isTrue bool = true
	fmt.Println(isTrue)

	// composite data type - a complex type which is constructed from several primitive types
	var result time.Time = time.Now()
	fmt.Println(result)
}
