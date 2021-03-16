package main

import "fmt"

func main() {
	zombies := []string{"Paul", "Katya", "George", "Lucy"}

	// You can remove an element from a slice using
	// a combination of what we already know
	// zombies[:toRemove] - take a subslice with elems [0,toRemove)
	// zombies[toRemove+1:] - take a subslice with elems [toRemove+1,len)
	// append(slice1, slice2...) - combine two slices
	toRemove := 2
	fmt.Println(zombies)
	zombies = append(zombies[:toRemove], zombies[toRemove+1:]...)
	fmt.Println(zombies)

	// You should be careful with this though
	// If you don't reuse the original slice
	// you might get nasty side-effects
	zombies1 := []string{"Paul", "Katya", "Lucy"}
	zombies2 := append(zombies1[:toRemove], zombies1[toRemove+1:]...)
	fmt.Println(zombies1) // [Paul Lucy Lucy]
	fmt.Println(zombies2) // [Paul Lucy]

	// This happens because of the fact that subslices create
	// a view on the original slice, not a copy.
	// If you want to achieve something like that,
	// you need to explicitly create a copy of the original slice
	toRemove = 1
	zombiesCopy := make([]string, len(zombies[:toRemove]))
	copy(zombiesCopy, zombies[:toRemove])
	zombiesWithoutKatya := append(zombiesCopy, zombies[toRemove+1:]...)
	fmt.Println(zombies)             // [Paul Katya Lucy]
	fmt.Println(zombiesWithoutKatya) // [Paul Lucy]

}
