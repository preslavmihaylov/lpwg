package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	num := 1

	// The for statement works similarly to the if-statement:
	// it checks the condition and if it's true, it executes the code block
	// The difference is that the for loop will not exit the statement after
	// the execution is done. Instead, it will check the condition again and
	// continue repeating the code block while the condition holds true
	for num < 10 {
		fmt.Println(num)
		num++
	}

	gameOver := false
	rand.Seed(time.Now().UTC().UnixNano())

	// You're not prohibited from using for loops
	// meant for things other than counting numbers
	for !gameOver {
		fmt.Println("The game is on-going")
		num := rand.Intn(100)
		if num < 10 {
			gameOver = true
		}
	}
}
