package main

import "fmt"

func main() {
	var aFloat float32 = 14.3

	fmt.Println(aFloat + 3)
	fmt.Println(aFloat - 3)
	fmt.Println(aFloat * 3)
	fmt.Println(aFloat / 3)

	// modulus is not defined for floats
	// fmt.Println(aFloat % 3)

	var bFloat float32 = aFloat
	aFloat++
	bFloat--
	fmt.Println(aFloat, bFloat)

	// Adding 0.1 to 3.0 8 times should result in 3.8, right?
	// Well, due to some pecularities when dealing with floating point numbers, you'd get 3.8000000000000007 instead.
	// Be wary of this drawback. If you're writing a program which demands very accurate results, don't use floats!
	var num1 float64 = 3.0
	num1 = num1 + 0.1
	num1 = num1 + 0.1
	num1 = num1 + 0.1
	num1 = num1 + 0.1
	num1 = num1 + 0.1
	num1 = num1 + 0.1
	num1 = num1 + 0.1
	num1 = num1 + 0.1
	fmt.Println(num1)

	// Interestingly enough, division by zero is defined for floats.
	// The reasoning for this is not important.
	// But the take away is that you can divide floats by zero and you shouldn't be surprised if you get results like these in your programs
	var toDivByZero = 3.0
	fmt.Println(toDivByZero / 0.0)
	fmt.Println(toDivByZero / -0.0)
	fmt.Println(-toDivByZero / 0.0)
	fmt.Println(-toDivByZero / -0.0)
}
