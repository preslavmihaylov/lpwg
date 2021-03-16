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

	elements1 := strings.Split(strings.TrimSpace(line), " ")

	line, err = r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	elements2 := strings.Split(strings.TrimSpace(line), " ")

	var union []string
	for _, elem := range elements1 {
		exists := false
		for _, other := range union {
			if elem == other {
				exists = true
				break
			}
		}

		if !exists {
			union = append(union, elem)
		}
	}

	for _, elem := range elements2 {
		exists := false
		for _, other := range union {
			if elem == other {
				exists = true
				break
			}
		}

		if !exists {
			union = append(union, elem)
		}
	}

	fmt.Println(strings.Join(union, " "))
}
