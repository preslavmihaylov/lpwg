package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	animalType := strings.TrimSpace(line)

	// given a type of animal, print an example of an animal of that type.
	// The known animal types are "Mammal", "Amphibian", "Birds", "Fish"
	// If you get any other type of input, print "I don't know that animal type"
	if animalType == "Mammal" {
		fmt.Println("An example is a hooman")
	} else if animalType == "Amphibian" {
		fmt.Println("An example is the frog")
	} else if animalType == "Birds" {
		fmt.Println("An example is the eagle")
	} else if animalType == "Fish" {
		fmt.Println("An example is the Salmon")
	} else {
		fmt.Println("I don't know that animal type")
	}
}
