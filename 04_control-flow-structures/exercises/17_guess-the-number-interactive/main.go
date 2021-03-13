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

	n, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}

	rand.Seed(time.Now().UTC().UnixNano())
	myGuess := rand.Intn(n)
	fmt.Println("my guess is", myGuess)
	wrongGuessesCnt := 0
	for {
		line, err = r.ReadString('\n')
		if err != nil {
			panic(err)
		}

		guess, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			panic(err)
		}

		if guess == myGuess {
			fmt.Println("You guessed the number correctly!")
			break
		} else {
			wrongGuessesCnt++
			fmt.Println("wrong guess! Try again...")
		}
	}

	fmt.Printf("You had %d wrong guesses. ", wrongGuessesCnt)
	if wrongGuessesCnt <= n/4 {
		fmt.Println("You were quite lucky!")
	} else if wrongGuessesCnt <= n/2 {
		fmt.Println("You did alright.")
	} else {
		fmt.Println("It took you a while...")
	}
}
