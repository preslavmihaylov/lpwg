package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	line = strings.TrimSpace(line)
	occurrences := map[rune]int{}
	for _, ch := range line {
		if ch == ' ' || !unicode.IsLetter(ch) {
			continue
		}

		occurrences[unicode.ToLower(ch)] += 1
	}

	keys := make([]int, 0, len(occurrences))
	for k := range occurrences {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Printf("%c: %d\n", k, occurrences[rune(k)])
	}
}
