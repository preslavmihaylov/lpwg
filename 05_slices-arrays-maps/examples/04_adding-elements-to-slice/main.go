package main

import "fmt"

func main() {
	zombies := []string{}

	// You can use append() to add elements to a slice
	// It takes a slice as a first argument and
	// a variable number of elements as its later arguments
	fmt.Println(zombies)
	zombies = append(zombies, "Paul", "George", "Katya")
	fmt.Println(zombies)

	// You can also use append to not modify an existing slice
	// but instead create a new one which copies the existing slice &
	// adds elements to it
	zombiesCopy := append(zombies, "Lucy")
	fmt.Println(zombiesCopy)

	odds := []int{1, 3, 5, 7}
	evens := []int{2, 4, 6, 8}
	result := append(odds, evens...) // same as invoking append(odds, 2, 4, 6, 8)
	fmt.Println(result)
}
