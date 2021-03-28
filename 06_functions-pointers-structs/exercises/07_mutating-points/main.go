package main

import "fmt"

type Point struct {
	row, col int
}

func main() {
	p1 := Point{row: 1, col: 1}
	p2 := Point{row: 1, col: 1}
	p3 := Point{row: 1, col: 1}

	incPtr(&p1)
	incNoPtr(p2)
	incNewPtr(&p3)
	fmt.Println(p1, p2, p3)
}

func incPtr(p *Point) {
	p.row++
	p.col++
}

func incNoPtr(p Point) {
	p.row++
	p.col++
}

func incNewPtr(p *Point) {
	p = &Point{row: 2, col: 2}
	p.row++
	p.col++
}
