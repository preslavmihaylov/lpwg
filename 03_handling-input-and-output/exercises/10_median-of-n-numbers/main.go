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

	cnt, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}

	nums := []int{}
	i := 0
	for i < cnt {
		line, err = r.ReadString('\n')
		if err != nil {
			panic(err)
		}

		num, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			panic(err)
		}

		nums = append(nums, num)
		i++
	}

	median := nums[cnt/2]
	fmt.Printf("Median=%d\n", median)
}
