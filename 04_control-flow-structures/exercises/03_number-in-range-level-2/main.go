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

	line, err = r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	isStartInclusive := strings.TrimSpace(line) == "in"

	line, err = r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	isEndInclusive := strings.TrimSpace(line) == "in"

	var numberIsInRangeStart bool
	if isStartInclusive {
		numberIsInRangeStart = x >= a
	} else {
		numberIsInRangeStart = x > a
	}

	var numberIsInRangeEnd bool
	if isEndInclusive {
		numberIsInRangeEnd = x <= b
	} else {
		numberIsInRangeEnd = x < b
	}

	var rangeStartChar, rangeEndChar rune
	if isStartInclusive {
		rangeStartChar = '['
	} else {
		rangeStartChar = '('
	}

	if isEndInclusive {
		rangeEndChar = ']'
	} else {
		rangeEndChar = ')'
	}

	if numberIsInRangeStart && numberIsInRangeEnd {
		fmt.Printf("%d is in the range %c%d; %d%c", x, rangeStartChar, a, b, rangeEndChar)
	} else {
		fmt.Printf("%d is not in the range %c%d; %d%c", x, rangeStartChar, a, b, rangeEndChar)
	}
}
