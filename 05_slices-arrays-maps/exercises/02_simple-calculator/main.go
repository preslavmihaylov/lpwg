package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var numbers []int
	var cmd string
	r := bufio.NewReader(os.Stdin)
	for cmd != "exit" {
		line, err := r.ReadString('\n')
		if err != nil {
			panic(err)
		}

		cmdElements := strings.Split(strings.TrimSpace(line), " ")
		cmd = cmdElements[0]
		args := []int{}
		for _, elem := range cmdElements[1:] {
			num, err := strconv.Atoi(elem)
			if err != nil {
				panic(err)
			}

			args = append(args, num)
		}

		msg := ""
		switch cmd {
		case "add":
			numbers = append(numbers, args[0])
			msg = "OK"
		case "inc":
			if args[0] >= len(numbers) {
				msg = "invalid index"
				break
			}

			numbers[args[0]]++
			msg = "OK"
		case "sub":
			if args[0] >= len(numbers) {
				msg = "invalid index"
				break
			}

			numbers[args[0]] -= args[1]
			msg = "OK"
		case "mul":
			if args[0] >= len(numbers) {
				msg = "invalid index"
				break
			}

			numbers[args[0]] *= args[1]
			msg = "OK"
		case "acc":
			if args[0] >= len(numbers) {
				msg = "invalid index"
				break
			}

			numbers[args[0]] += args[1]
			msg = "OK"
		case "show":
			fmt.Println(strings.Trim(fmt.Sprint(numbers), "[]"))
		}

		if msg != "" {
			fmt.Println(msg)
		}
	}

}
