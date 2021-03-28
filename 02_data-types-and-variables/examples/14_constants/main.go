package main

import (
	"fmt"
)

func main() {
	// Constants, like variables, are a placeholder for a value.
	// However, some limitations when using them is that
	// they can never change once they're created
	const pi = 3.14
	fmt.Println(pi)

	// And you can only store primitive data types inside a constant
	// const now = time.Now() // compilation error

	// A typical use-case for constants is to provide a name to some value
	// in your code.
	// For example, instead of calculating r * r * 3.14, it is more convenient
	// to write & read r * r * pi, which makes it obvious what that value represents
}
