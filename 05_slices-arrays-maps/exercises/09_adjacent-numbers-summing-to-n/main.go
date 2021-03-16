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

	tokens := strings.Split(strings.TrimSpace(line), " ")
	var numbers []int
	for _, token := range tokens {
		num, err := strconv.Atoi(token)
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, num)
	}

	line, err = r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	target, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}

	solutionFound := false
	for subsetSize := 2; subsetSize <= len(numbers); subsetSize++ {
		for i := 0; i < len(numbers)-subsetSize+1; i++ {
			var subset []int
			for j := i; j < i+subsetSize; j++ {
				subset = append(subset, numbers[j])
			}

			sum := 0
			for _, num := range subset {
				sum += num
			}

			if sum == target {
				subsetTokens := strings.Split(
					strings.Trim(fmt.Sprint(subset), "[]"), " ")
				subsetStr := strings.Join(subsetTokens, "+")
				fmt.Printf("%s = %d\n", subsetStr, target)
				solutionFound = true
			}
		}
	}

	if !solutionFound {
		fmt.Println("no solution exists")
	}
}
