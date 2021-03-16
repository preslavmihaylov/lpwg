package main

import "fmt"

func main() {
	zombies := []string{"Paul", "Gooey", "George", "Lucy"}

	// You can get values from a slice based on the elements' indices
	// The first element is at index 0
	// and the last one is at index length - 1
	paul := zombies[0]
	george := zombies[2]
	fmt.Println(paul, george) // Paul George

	// You can also modify elements in a slice by specifying the index
	// and a new value for the selected element
	fmt.Println(zombies) // [Paul Gooey George Lucy]
	zombies[1] = "Katya"
	fmt.Println(zombies) // [Paul Katya George Lucy]

	// The built in function len gives you the length of a given data structure
	// In this case, with a slice, it will return the count of elements in it
	fmt.Println(len(zombies)) // 4
}
