package main

import "fmt"

// At the package level scope, assign the following values to the
// three variables

var x3 = 42
var y3 = "Test User"
var z3 = true

func main() {

	// Use fmt.Sprintf to print all of the values to one single string.
	// Assign the returned value of TYPE string using the short declaration
	// operator to the variable with the IDENTIFIER "s3".
	s3 := fmt.Sprintf("%v\t%v\t%v", x3, y3, z3)

	// Print out the value stored by variable "s3"
	fmt.Println(s3) // Result: 42	Test User	true
}
