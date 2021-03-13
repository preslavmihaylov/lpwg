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

	first, second := 1, 1
	fmt.Printf("%d %d ", first, second)
	for i := 3; i <= n; i++ {
		curr := first + second
		first = second
		second = curr

		fmt.Printf("%d ", curr)
	}
}
