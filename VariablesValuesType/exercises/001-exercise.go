package main

import "fmt"

func main() {
	// Using the short declaration operator, ASSIGN these VALUES to VARIABLES
	// with the IDENTIFIERS "x1" and "y1" and "z1"
	x1 := 42
	y1 := "Test User"
	z1 := true

	// A single print statements
	fmt.Println(x1, y1, z1) // Result: 42 Test User true

	// Multiple print statements
	fmt.Println(x1) // Result: 42
	fmt.Println(y1) // Result: Test User
	fmt.Println(z1) // Result: true
}
