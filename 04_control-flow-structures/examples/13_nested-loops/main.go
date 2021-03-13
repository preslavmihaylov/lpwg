package main

import "fmt"

func main() {
	for cnt1 := 0; cnt1 < 10; cnt1++ {
		for cnt2 := 0; cnt2 < 10; cnt2++ {
			fmt.Println(cnt1, cnt2)
		}
	}
}
