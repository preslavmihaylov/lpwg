package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello World")
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	a, err := strconv.ParseFloat(strings.TrimSpace(line), 64)
	if err != nil {
		panic(err)
	}

	line, err = r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	b, err := strconv.ParseFloat(strings.TrimSpace(line), 64)
	if err != nil {
		panic(err)
	}

	line, err = r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	c, err := strconv.ParseFloat(strings.TrimSpace(line), 64)
	if err != nil {
		panic(err)
	}

	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		fmt.Println("no real roots exist!")
	} else if discriminant == 0 {
		fmt.Printf("x = %.2f\n", (-b)/(2*a))
	} else {
		x1 := (-b + math.Sqrt(discriminant)) / (2 * a)
		x2 := (-b - math.Sqrt(discriminant)) / (2 * a)
		fmt.Printf("x1 = %.2f\nx2 = %.2f\n", x1, x2)
	}
}
