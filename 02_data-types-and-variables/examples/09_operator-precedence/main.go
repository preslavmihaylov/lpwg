package main

import "fmt"

func main() {
	// operations have precedence. For example, + is of lower precedence than multiplication (*) or division (/)
	// To override the precedence, use brackets
	// Additionally, when writing complicated expressions like the one below, it makes sense to add brackets, even if redundant
	// in order to make it easier for the reader to understand what is going on and how the expression will evaluate
	var num int8 = (14+10)/2 - 1 + (10 / 3) - 19
	fmt.Println(num)
}
