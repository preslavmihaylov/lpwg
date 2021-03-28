package main

import (
	"time"

	tm "github.com/buger/goterm"
)

func main() {
	tm.Clear()
	PrintAt(5, 5, "(")
	PrintAt(5, 6, "*")
	PrintAt(5, 7, ")")

	time.Sleep(5 * time.Second)
}

func PrintAt(row, col int, msg string) {
	tm.MoveCursor(col, row)
	tm.Print(msg)
	tm.Flush()
}
