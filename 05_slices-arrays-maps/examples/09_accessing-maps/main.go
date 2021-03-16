package main

import "fmt"

func main() {
	playerHPs := map[string]int{
		"coolBoy123": 10,
		"Aldarien":   50,
		"badBoyzz":   30,
	}

	// Values in a map can be accessed similarly
	// to how you access values in a slice.
	// Only the data type you use to access the value is based
	// on the data type of the key in the map
	hp := playerHPs["badBoyzz"]
	fmt.Println(hp) // 30

	// Modifying values in a map is not too difficult
	playerHPs["badBoyzz"] = 25
	fmt.Println(playerHPs["badBoyzz"]) // 25

	// If you access a value which doesn't exist
	// you won't get an error contrary to how slices work.
	// Instead, you can optionally get a second boolean value
	// which indicates whether the value exists or not
	hp, ok := playerHPs["prettyBoy"]
	fmt.Println(hp, ok) // 0 false
	hp, ok = playerHPs["badBoyzz"]
	fmt.Println(hp, ok) // 25 true

	var nonInitMap map[string]int
	// assigning/adding values in a non-initialized map will result
	// in a panic
	// nonInitMap["somekey"] = 42 // this panics
	_ = nonInitMap

	// Try initializing empty maps like this when you can
	// to avoid that nasty issue
	initMap := map[string]int{}

	// This is equivalent to:
	// _, ok = nonInitMap["somekey"]
	// if !ok { ... }
	if _, ok = initMap["somekey"]; !ok {
		// Adding new values to a map is the same as
		// changing existing ones
		initMap["somekey"] = 42
	}
}
