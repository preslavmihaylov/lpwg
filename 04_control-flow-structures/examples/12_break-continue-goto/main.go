package main

import "fmt"

func main() {
	// break allows you to unconditionally stop the execution of the for loop
	for cnt := 0; cnt < 10; cnt++ {
		if cnt == 7 {
			break
		}

		fmt.Println(cnt)
	}

	// continue doesn't stop the whole execution of the for loop
	// but it allows you to stop the execution of a particular cycle
	// in the for loop
	fmt.Println("continue demonstration...")
	for cnt := 0; cnt < 10; cnt++ {
		if cnt%2 == 1 {
			continue
		}

		fmt.Println(cnt)
	}

	// goto allows you to jump anywhere in the function
	// based on a label (similar to bookmark) which you place somewhere
	// and later you goto <that-label>
	//
	// NOTE: It's a bad practice using goto and you should avoid doing that
	// in most circumstances
	fmt.Println("goto demonstration...")
	for cnt := 0; cnt < 10; cnt++ {
		if cnt == 5 {
			goto exit
		}

		fmt.Println(cnt)
	}

	fmt.Println("Exiting program 1...")
exit:

	fmt.Println("Exiting program 2...")
}
