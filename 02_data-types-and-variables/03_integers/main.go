package main

import "fmt"

func main() {
	// Integer data types share the property that they can only store whole numbers
	// But the different integer data types vary in size & sign-ness

	// sign-ness -> whether the data type stores negative values or not
	// unsigned data types can store twice as many positive values than their signed counterparts

	// These are all signed integer types
	var n1 int8 = 120
	var n2 int16 = 16000
	var n3 int32 = 2000000000
	var n4 int64 = 9000000000000000000

	// These are all unsigned integer types
	var n1u uint8 = 250
	var n2u uint16 = 65000
	var n3u uint32 = 4000000000
	var n4u uint64 = 15000000000000000000

	// These types are aliases for other ones

	// int is an alias for int32 or int64 based on your computer type.
	// Most likely int64 nowadays
	var a1 int = 9000000000000000000

	// byte is an alias for uin8
	var a2 byte = 255

	fmt.Println(n1, n2, n3, n4, n1u, n2u, n3u, n4u, a1, a2)
}
