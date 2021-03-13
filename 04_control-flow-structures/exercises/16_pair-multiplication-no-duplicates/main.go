package main

import (
	"bufio"
	"fmt"
	"math"
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

	line, err = r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	x, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}

	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		for j := 1; j <= n; j++ {
			if i*j == x {
				fmt.Printf("%d * %d = %d\n", i, j, x)
			}
		}
	}
}
