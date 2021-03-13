package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	num, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}

	gamesBought := []string{}

	if num >= 30 {
		gamesBought = append(gamesBought, "StarCraft 2")
		num -= 30
	}

	if num >= 130 {
		gamesBought = append(gamesBought, "Cyberpunk 2077")
		num -= 130
	}

	if num >= 60 {
		gamesBought = append(gamesBought, "Witcher 3")
		num -= 60
	}

	if len(gamesBought) == 0 {
		fmt.Println("I couldn't buy anything!")
	} else {
		fmt.Println("Here's what I bought:")
		fmt.Println(strings.Join(gamesBought, ", "))
	}

}
