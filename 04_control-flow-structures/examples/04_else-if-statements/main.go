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

	// if number < 10 => the number is very small
	// if number is > 10 => the number is small
	// if number is > 100 => the number is medium
	// if number is > 500 => the number is large
	if num < 10 {
		fmt.Println("the number is very small")
	}

	if num > 500 {
		fmt.Println("the number is large")
	} else if num > 100 {
		fmt.Println("the number is medium")
	} else if num > 10 {
		fmt.Println("the number is small")
	}
}
