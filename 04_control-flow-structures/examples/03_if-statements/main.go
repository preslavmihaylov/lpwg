package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	message := strings.TrimSpace(line)

	// if the user provides Mr. ..., then print "Hello Sir!"
	if message == "Mr." {
		fmt.Println("Hello Sir!")
	}

	// if the user provides Ms. ..., then print "Hello Madam!"
	if message == "Ms." {
		fmt.Println("Hello Madam!")
	}
}
