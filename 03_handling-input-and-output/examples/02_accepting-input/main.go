package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// bufio.Reader allows you to read text line by line from standard input and from any other source as well
	r := bufio.NewReader(os.Stdin)

	// ReadString with a delimiter '\n' will read text from stdin until a newline is encountered.
	// The newline will be included in the text which was read
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	// strconv.Atoi converts a strings to an integer and returns an error if anything goes wrong.
	// We use strings.TrimSpace to remove the extra newline at the end of the line
	num, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}

	fmt.Println("My number is:", num)
	fmt.Println("Let me increment it by 5:", num+5)

	line, err = r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	// strconv.ParseFloat converts a string to a floating-point number
	// The second argument specifies the precision you want to use. 32 = float32, 64 = float64
	floatNum, err := strconv.ParseFloat(strings.TrimSpace(line), 64)
	if err != nil {
		panic(err)
	}

	fmt.Println("My float number is:", floatNum)
}
