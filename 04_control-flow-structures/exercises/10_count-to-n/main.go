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

	num, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}

	var result string
	for i := 0; i < num; i++ {
		result += fmt.Sprintf("%d ", i)
	}

	fmt.Println(strings.TrimSpace(result))
}
