package main

import (
	"errors"
	"fmt"
)

func main() {
	// This is a void function - it doesn't take any arguments and doesn't
	// return any results.
	// You just invoke it and it does its thing
	printGameInfo()

	// This is a void function which accepts two arguments of type string
	printPlayerMessage("Asmongold", "Lowko")

	// The arguments can be of different types if you need them to
	printPlayerNameAndAge("Asmongold", 25)

	// You can't invoke a function with arguments of different types
	// than what's expected
	// printPlayerNameAndAge("Asmongold", "25")

	// This is a function which returns a value of type int
	result := calculateSumOfDigits(12345)
	fmt.Println(result) // 15

	// A function can also return multiple values.
	// In Go, it's common to see functions return
	// a second 'error' return type
	// which indicates whether any error occurred inside the function
	result, err := calculateSumOfDigitsWithErr(12345)
	if err != nil {
		panic(err)
	}

	fmt.Println(result) // 15

	// A variadic function can accept
	// a variable amount of arguments
	printPlayerInfos("Asmongold", "Lowko", "Winter")
	printPlayerInfos("Asmongold", "Lowko")

	// Effectively, it works as if you pass a slice of values
	// to the function
	// printPlayerInfos([]string{"Asmongold", "Lowko", "Winter"})
}

func printGameInfo() {
	fmt.Println("The game is called Quake")
	fmt.Println("The battle is about to begin, get ready!")
}

// (player1, player2 string) is the same as (player1 string, player2 string)
// It's just a syntax sugar - made to help writing code easier
func printPlayerMessage(player1, player2 string) {
	fmt.Printf("Player %s and %s are on the starting line...\n", player1, player2)
	fmt.Println("The battle is about to begin!")
}

func printPlayerNameAndAge(name string, age int) {
	fmt.Printf("Player's name is %s and their age is %d\n", name, age)
}

func calculateSumOfDigits(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}

	// The 'return' keyword stops the function execution
	// and passes the value you specify as a result from the
	// function call to the invoker
	return sum
}

func calculateSumOfDigitsWithErr(num int) (int, error) {
	if num < 0 {
		return 0, errors.New("negative values are not supported")
	}

	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}

	return sum, nil
}

func printPlayerInfos(players ...string) {
	// At this point, players' type is a slice of strings
	for _, player := range players {
		fmt.Printf("Player %s is cool!\n", player)
	}
}
