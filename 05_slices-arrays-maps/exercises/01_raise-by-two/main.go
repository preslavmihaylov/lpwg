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
	for _, element := range elementsStr {
		num, err := strconv.Atoi(element)
		if err != nil {
			panic(err)
		}

		elements = append(elements, num)
	}

	for i := 0; i < len(elements); i++ {
		elements[i] *= elements[i]
	}

	for _, elem := range elements {
		fmt.Print(elem, " ")
	}
}
