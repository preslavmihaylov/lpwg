package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	line1, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	line2, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	words := strings.Split(strings.TrimSpace(line1), " ")
	excludedWords := strings.Split(strings.TrimSpace(line2), " ")

	var filteredWords []string
	for _, word := range words {
		excluded := false
		for _, excludedWord := range excludedWords {
			if word == excludedWord {
				excluded = true
				break
			}
		}

		if !excluded {
			filteredWords = append(filteredWords, word)
		}
	}

	fmt.Println(strings.Join(filteredWords, " "))
}
