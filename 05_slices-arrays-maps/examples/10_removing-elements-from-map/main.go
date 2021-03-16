package main

import "fmt"

func main() {
	playerHPs := map[string]int{
		"coolBoy123": 10,
		"Aldarien":   50,
		"badBoyzz":   30,
	}

	hp, ok := playerHPs["coolBoy123"]
	fmt.Println(hp, ok) // 10 true

	// To delete elements from a map, use the built-in delete function
	delete(playerHPs, "coolBoy123")
	hp, ok = playerHPs["coolBoy123"]
	fmt.Println(hp, ok) // 0 false

	// Assigning the value 0 is not the same as deleting the entry
	playerHPs["Aldarien"] = 0
	hp, ok = playerHPs["Aldarien"]
	fmt.Println(hp, ok) // 0 true

	hp, ok = playerHPs["qweqwrwerwerwerwe"]
	fmt.Println(hp, ok) // 0 false

	// Deleting non-existent entries doesn't do anything
	delete(playerHPs, "qweqwrwerwerwerwe")
	hp, ok = playerHPs["qweqwrwerwerwerwe"]
	fmt.Println(hp, ok) // 0 false
}
