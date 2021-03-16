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

	elementsStr := strings.Split(strings.TrimSpace(line), " ")
	var elements []int
	for _, elem := range elementsStr {
		num, err := strconv.Atoi(elem)
		if err != nil {
			panic(err)
		}

		elements = append(elements, num)
	}

	// repeat until you get to the end
	for startIndex := 0; startIndex < len(elements); startIndex++ {
		// go through each element which is not sorted yet
		minIndex := startIndex
		for i := startIndex; i < len(elements); i++ {
			// if the current element is minimum, record it as minimum
			if elements[i] < elements[minIndex] {
				minIndex = i
			}
		}

		// swap the first element and the minimum one
		elements[startIndex], elements[minIndex] =
			elements[minIndex], elements[startIndex]
	}

	for _, elem := range elements {
		fmt.Print(elem, " ")
	}
}
