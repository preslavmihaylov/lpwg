package main

import "fmt"

func main() {
	var sum int
	num := 123456

	for num > 0 { // while num is greater than 0
		digit := num % 10
		num = num / 10
		sum = sum + digit
	}

	fmt.Println(sum)
}
