package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	dayOfWeek := strings.TrimSpace(line)

	// If the day of week is Mon-Fri, print "It's a workday"
	// In any other case, print "It's a weekend"
	switch dayOfWeek {
	case "Monday":
		fallthrough
	case "Tuesday":
		fallthrough
	case "Wednesday":
		fmt.Println("It's a day before Wednesday")
		fallthrough
	case "Thursday":
		fallthrough
	case "Friday":
		fmt.Println("It's a workday")
	default:
		fmt.Println("It's a weekend")
	}

	// if dayOfWeek == "Monday" {
	// 	fmt.Println("It's a workday")
	// } else if dayOfWeek == "Tuesday" {
	// 	fmt.Println("It's a workday")
	// } else if dayOfWeek == "Wednesday" {
	// 	fmt.Println("It's a workday")
	// } else if dayOfWeek == "Thursday" {
	// 	fmt.Println("It's a workday")
	// } else if dayOfWeek == "Friday" {
	// 	fmt.Println("It's a workday")
	// } else {
	// 	fmt.Println("It's a weekend")
	// }
}
