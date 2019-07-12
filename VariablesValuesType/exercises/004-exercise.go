package main

import "fmt"

// Create your own type. Have the underlying type be an int.
type myVariable int

// Create a VARIABLE of your new TYPE with the IDENTIFIER "x4"
// using the "VAR" keyword
var x4 myVariable

// At the Package level scope, using the var Keyword, create a
// Variable with the IDENTIFIER "y4". The variable should be of
// the UNDERLYING TYPE of your custom TYPE "x4"
var y4 int

func main() {
	// Print out the VALUE of the variable "x4"
	fmt.Println(x4) // Result: 0

	// Print out the TYPE of the variable "x4"
	fmt.Printf("%T\n", x4) // Result: main.myVariable

	//Assign 100 to the VARIABLE "x4" using the "=" OPERATOR
	x4 = 100

	// Print out the VALUE of the variable "x4"
	fmt.Println(x4) // Result: 100

	// Now use CONVERSION to convert the TYPE of the VALUE stored in
	// "x4" to the UNDERLYING TYPE
	y4 = int(x4)

	// Print out the value stored in "y4" and TYPE of "y4"
	fmt.Println(y4)      // Result: 100
	fmt.Printf("%T", y4) // Result: int
}
