package main

import "fmt"

func main() {
	n1, n2 := 1.0, 78.5
	kgToLbsCoefficient := 2.2046226218

	res1, res2 := n1*kgToLbsCoefficient, n2*kgToLbsCoefficient
	fmt.Println(res1, res2)
}
