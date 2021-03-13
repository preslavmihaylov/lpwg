package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	rangeOfNumbers, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}

	line, err = r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	userGuess, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}

	rand.Seed(time.Now().UTC().UnixNano())
	myGuess := rand.Intn(rangeOfNumbers)
	if userGuess == myGuess {
		fmt.Println("You guessed the number correctly!")
	} else {
		fmt.Println("That's not the number I had in mind.")
		fmt.Printf("You guessed %d but my number was %d\n", userGuess, myGuess)
	}
}
