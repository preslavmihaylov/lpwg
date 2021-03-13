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

	n, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}

	for curr := 2; curr <= n; curr++ {
		isPrime := true
		for dividedBy := 2; dividedBy < curr; dividedBy++ {
			if curr%dividedBy == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			fmt.Printf("%d ", curr)
		}
	}
}
