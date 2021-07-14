package main

import "fmt"

func main() {
	// Check if the point is inside the rectangle
	pointX, pointY := 3, 3
	topLeftX, topLeftY := 1, 5
	botRightX, botRightY := 5, 1

	betweenLeftAndRight := pointX > topLeftX && pointX < botRightX
	betweenBotAndTop := pointY > botRightY && pointY < topLeftY
	if betweenLeftAndRight && betweenBotAndTop {
		fmt.Println("The point is inside the rectangle")
	} else {
		fmt.Println("The point is outside the rectangle")
	}
}
