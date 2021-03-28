package main

import (
	"fmt"
	"time"
)

func main() {
	// Other than primitive data types, you also have composite types
	// They consist of multiple primitive data types and/or composite types
	// This is one example - the time.Time data type
	var t1 time.Time
	t1 = time.Now()

	// These types have their own operations defined for them
	// and you have to consult the documentation in order to
	// understand what operations these data types support
	year, month, day := t1.Date()
	fmt.Println(year, month, day)
}
