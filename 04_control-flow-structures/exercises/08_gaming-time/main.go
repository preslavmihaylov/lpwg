package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	currentTime, err := time.Parse("03:04:05 PM", strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}

	startTime1, _ := time.Parse("03:04:05 PM", "07:00:00 PM")
	endTime1, _ := time.Parse("03:04:05 PM", "08:59:59 PM")

	startTime2, _ := time.Parse("03:04:05 PM", "08:00:00 AM")
	endTime2, _ := time.Parse("03:04:05 PM", "09:59:59 AM")

	startTime3a, _ := time.Parse("03:04:05 PM", "11:00:00 PM")
	endTime3a, _ := time.Parse("03:04:05 PM", "11:59:59 PM")

	startTime3b, _ := time.Parse("03:04:05 PM", "12:00:00 AM")
	endTime3b, _ := time.Parse("03:04:05 PM", "05:59:59 AM")

	if isInTimeRange(currentTime, startTime1, endTime1) ||
		isInTimeRange(currentTime, startTime2, endTime2) ||
		isInTimeRange(currentTime, startTime3a, endTime3a) ||
		isInTimeRange(currentTime, startTime3b, endTime3b) {
		fmt.Println("It's gaming time!")
	} else {
		fmt.Println("It's not time for games yet...")
	}
}

func isInTimeRange(currentTime, startTime, endTime time.Time) bool {
	return currentTime.Equal(startTime) ||
		currentTime.Equal(endTime) ||
		(currentTime.After(startTime) && currentTime.Before(endTime))
}
