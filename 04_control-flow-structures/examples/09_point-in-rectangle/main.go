package main

import "fmt"

func main() {
	// Check if the point is inside the rectangle
	pointX, pointY := 6, 3
	topLeftX, topLeftY := 1, 5
	botRightX, botRightY := 5, 1

	aboveBotLine := botRightY < pointY
	belowTopLine := topLeftY > pointY
	leftOfRightLine := botRightX > pointX
	rightOfLeftLine := topLeftX < pointX

	if aboveBotLine && belowTopLine && leftOfRightLine && rightOfLeftLine {
		fmt.Println("The point is inside the rectangle")
	} else {
		fmt.Println("The point is outside the rectangle")
	}
}
