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

	numbers := strings.Split(strings.TrimSpace(line), " ")
	sum := 0
	for _, numStr := range numbers {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}

		sum = sum + num
	}

	avg := float64(sum) / float64(len(numbers))
	fmt.Printf("Average=%.2f\n", avg)
}
