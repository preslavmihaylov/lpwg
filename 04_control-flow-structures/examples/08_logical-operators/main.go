package main

import "fmt"

func main() {
	cond1 := true
	cond2 := true
	cond3 := true
	cond4 := true

	// Logical AND returns true only if ALL of the sub-conditions are true
	if cond1 && cond2 && cond3 && cond4 {
		fmt.Println("All conditions match")
	} else {
		fmt.Println("All conditions don't match")
	}

	// Logical OR returns true if AT LEAST ONE of the sub-conditions is true
	if cond1 || cond2 || cond3 || cond4 {
		fmt.Println("At least one condition matches")
	} else {
		fmt.Println("Neither of the conditions match")
	}

	carType := "Automatic" // or Manual
	gearsCnt := 0

	// Logical AND and logical OR can be combined to result in
	// more complicated conditional expressions
	if (carType == "Automatic" && gearsCnt == 0) ||
		(carType == "Manual" && gearsCnt > 0) {
		fmt.Println("Correct configuration")
	} else {
		fmt.Println("The configuration is incorrect")
	}

	// The NOT operator can be used to invert
	// the result of any boolean expression
	if !(cond1 && cond2 && cond3 && cond4) {
		fmt.Println("At least one condition doesn't match")
	} else {
		fmt.Println("All conditions match")
	}
}
