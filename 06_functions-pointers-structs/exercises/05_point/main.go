package main

import "fmt"

type Point struct {
	row, col int
}

func main() {
	p1 := Point{row: 1, col: 5}
	p2 := Point{row: 3, col: 4}
	p3 := &Point{row: 2, col: 6}
	fmt.Println(p1, p2, p3)
}
