package main

import "fmt"

func main() {
	num := 123
	num = num / 10
	secondToLastDigit := num % 10

	fmt.Println(secondToLastDigit)
}
