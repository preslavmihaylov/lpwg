package main

import "fmt"

func main() {
	// You can use any kind of primitive data type for the keys and values
	// when initializing a map
	primes := map[int]bool{
		2: true,
		3: false,
		4: false,
		5: true,
	}

	// Notice that a map of int to bool is not the same as a slice of bools.
	// With slices, you still have the elements ordered and you still get
	// errors when you access something which doesn't exist.
	// With maps, that's not the case.
	// The example below won't result in an error:
	fmt.Println(primes[50]) // 0

	// Iterating over a map is similar to how you iterate over a slice.
	// The only difference is that in a map, the elements' order is not
	// guaranteed.
	// In this case, on my machine I got the elements printed in this order:
	// 5 -> true
	// 2 -> true
	// 3 -> false
	// 4 -> false
	for key, value := range primes {
		fmt.Println(key, "->", value)
	}
}
