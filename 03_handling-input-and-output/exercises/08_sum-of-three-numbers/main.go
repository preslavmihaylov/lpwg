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

	num1, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}

	line, err = r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	num2, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}

	line, err = r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	num3, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}

	fmt.Println("Result:", num1+num2+num3)

	line, err = r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	num1, err = strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}

	line, err = r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	num2, err = strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}

	line, err = r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	num3, err = strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}

	fmt.Println("Result:", num1+num2+num3)
}
