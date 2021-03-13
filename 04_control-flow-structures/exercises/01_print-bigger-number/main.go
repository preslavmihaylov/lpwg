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

	if num1 > num2 {
		fmt.Println("The bigger number is", num1)
	} else if num1 < num2 {
		fmt.Println("The bigger number is", num2)
	} else {
		fmt.Println("The numbers are equal")
	}
}
