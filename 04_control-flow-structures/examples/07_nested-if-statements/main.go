package main

import (
	"fmt"
)

func main() {
	num1, num2, num3 := 15, 15, 14

	// You're given three numbers.
	// Check which number is greatest or if they are equal
	if num1 == num2 {
		if num2 == num3 {
			fmt.Println("The numbers are equal")
		} else {
			if num3 > num2 {
				fmt.Println("The third number is greatest")
			} else {
				fmt.Println("The first and second number are greatest")
			}
		}
	}

	// TODO: Finish this challenge yourself
}
