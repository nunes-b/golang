package main

import "fmt"

// Today we gona learn about types.
// Boolean - True or False
// String - Sequence of bytes
// Integer
// Float(float64/float32) - Decimal
func main() {

	fmt.Printf("Type %T - Value %v\n", true, true)
	fmt.Printf("Type %T Value %v\n", "Rômulo", "Rômulo")
	fmt.Printf("Type %T Value %v\n", 42, 42)
	fmt.Printf("Type %T Value %v\n", 3.4, 3.4)
}
