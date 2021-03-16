package main

import "fmt"

func main() {
	// The way you initialize an array is very similar to initializing a slice
	// You simply specify a fixed size between the square brackets
	zombies := [2]string{"Paul", "George"}

	// Arrays can't grow or shrink the way slices can
	// zombies = append(zombies, "Katya")

	// But you can still change existing elements
	fmt.Println(zombies)
	zombies[1] = "Katya"
	fmt.Println(zombies)
}
