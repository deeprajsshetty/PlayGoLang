package main

import "fmt"

// Use VAR to DECLARE three VARIABLES. The Variable should have package
// level Scope
var x2 int
var y2 string
var z2 bool

func main() {
	// Print out the values for each identifier.
	// Compiler assigned values to the variable.
	// What are these values called?

	fmt.Println(x2) // Result: 0
	fmt.Println(y2) // Result: nil (Empty)
	fmt.Println(z2) // Result: false
}
