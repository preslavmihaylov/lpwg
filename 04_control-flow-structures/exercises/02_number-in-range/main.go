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

	x, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}

	line, err = r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	a, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}

	line, err = r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	b, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}

	if x >= a && x < b {
		fmt.Printf("%d is in the range [%d; %d)", x, a, b)
	} else {
		fmt.Printf("%d is not in the range [%d; %d)", x, a, b)
	}
}
