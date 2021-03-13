package main

import "fmt"

func main() {
	// This version of the for loop does exactly what this code does:
	// cnt := 0
	// for cnt < 10 {
	//   fmt.Println(cnt)
	//   cnt++
	// }
	//
	// It is simply a more ergonomic variant of the above
	for cnt := 0; cnt < 10; cnt++ {
		fmt.Println(cnt)
	}
}
