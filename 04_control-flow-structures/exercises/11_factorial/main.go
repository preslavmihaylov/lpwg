package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 5! == 1 * 2 * 3 * 4 * 5
	//       1 + 2 + 3 + 4 + 5
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	num, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}

	factorial := 1
	for i := 1; i <= num; i++ {
		factorial *= i
	}

	fmt.Println(factorial)
}
