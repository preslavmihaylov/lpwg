package main

import "fmt"

func main() {
	var input1 float32 = 3.14
	var input2 float64 = 12.3456789
	var input3 int16 = 1234
	var input4 int16 = 1234

	var output1 int = int(input1)
	var output2 float32 = float32(input2)
	var output3 int8 = int8(input3)
	var output4 uint8 = uint8(input4)
	fmt.Println(output1)
	fmt.Println(output2)
	fmt.Println(output3)
	fmt.Println(output4)
}
