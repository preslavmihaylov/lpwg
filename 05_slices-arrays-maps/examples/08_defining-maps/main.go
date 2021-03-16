package main

func main() {
	// Maps allow you to use arbitrary primitive data types as keys
	// in order to associate eg. a player's nickname to health points.
	// With slices, it's harder because you will have to associate
	// the nicknames with integers you can use to index into the slice

	// string - the data type of the key
	// int - the data type of the value
	playerHPs := map[string]int{
		// In order to initialize values in a map, use the syntax:
		// key: value,
		"coolBoy123": 10,
		"Aldarien":   50,
		"badBoyzz":   30,
	}

	_ = playerHPs
}
