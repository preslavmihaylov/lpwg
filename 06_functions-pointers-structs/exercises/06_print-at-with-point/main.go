package main

import tm "github.com/buger/goterm"

type Point struct {
	row, col int
}

func main() {
	tm.Clear()
	PrintAt(Point{5, 5}, "Hello")
}

func PrintAt(p Point, msg string) {
	tm.MoveCursor(p.col, p.row)
	tm.Print(msg)
	tm.Flush()
}
