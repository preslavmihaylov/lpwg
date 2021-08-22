package main

import "fmt"

// You use type <name> struct { ... } to define your own data type
// The code inside the brackets is not normal executable code.
// Instead, it is the "blueprints" of this data type - how the data type is defined
// Here, you specify pairs of <name> <data-type>.
// E.g.
// type Person struct {
//   name string
//   age int
//   address string
//   ...
// }
//
// You can also use the short-hand <name1>, <name2> <data-type>
// if your fields share the same data type.
// E.g.
// type Person struct {
//   name, address string
//   age int
// }
type Point struct {
	row, col int
}

func main() {
	var p Point
	// To instantiate a value of your new data type, you use this syntax.
	// If you don't instantiate all fields,
	// those will be left with their zero values
	p = Point{
		row: 0,
		col: 5,
	}

	// Although this variable uses the same data type, the values stored
	// inside are different
	p1 := Point{
		row: 1,
		col: 2,
	}

	fmt.Printf("row: %d, col: %d\n", p.row, p.col)   // row: 0, col: 5
	fmt.Printf("row: %d, col: %d\n", p1.row, p1.col) // row: 1, col: 2

	// You can also use pointers with your custom data types
	// and they behave the same way as they do with primitive data types
	p2 := &p1
	p2.row = 3
	p2.col = 4
	fmt.Printf("row: %d, col: %d\n", p1.row, p1.col) // row: 3, col: 4

	// The rules for passing variables to functions applies for custom data types as well.

	// If you pass a variable as a value,
	// any changes in the function don't take effect beyond the function scope.
	incPoint(p1)
	fmt.Println(p1) // {3 4}

	// But if you pass a reference to the value instead,
	// changes in the function take effect beyond its scope
	incPointPtr(&p1)
	fmt.Println(p1) // {4 5}
}

func incPoint(p Point) {
	p.row += 1
	p.col += 1
}

func incPointPtr(p *Point) {
	// When working with pointers to custom data types,
	// you don't need to dereference the value of the pointer explicitly.
	//
	// What's written below can also be written as:
	//   (*p).row += 1
	//   (*p).col += 1
	//
	// But you can get away with not writing the dereference operator explicitly
	p.row += 1
	p.col += 1
}
