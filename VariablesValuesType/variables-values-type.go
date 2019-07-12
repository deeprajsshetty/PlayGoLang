package main

import "fmt"

// GO is Static Programming Language and Go is all about TYPE

var a = "DEEP"

var b = `Hello Deep,

		How are you? "Great"`

var s = 43

// DECLARE there is a VARIABLE with the IDENTIFIER "z" and that the VARIABLE with the IDENTIFIER "Z"
// is of Type int ASSIGNS the ZERO VALUE of TYPE int to "z", false for booleans, 0 for integers,
// 0.0 for floats, "" for strings, and nil for pointers, functions, interfaces, slices, channels, and maps.
var z int

func main() {
	// Starting with printing Hello World
	printHelloWorld()

	// Printing Numbers in Loops
	printNumberLoops()

	// Declare Variable and Assigning Values to Variable, Using Var Keyword
	fmt.Println(s)
	variables()

	// Exploring Type i.e. Data Types
	printValuesTypes()

	// Exit Here
	exitHere()
}

// Starting with printing Hello World
func printHelloWorld() {
	fmt.Println("Hello World")
}

// Printing Numbers in Loops
func printNumberLoops() {
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}
}

func variables() {
	x := 10 // Declare Variable and Assign Values.
	fmt.Println(x)
	x = 30 // Assigning Values to x variable
	fmt.Println(x)
	y := 20 + 30
	fmt.Println(y)

	// Var Keyword Scope - One time declare in the file and Use anywhere e.g. inside functions or Main.
	fmt.Println(s) // Value Assigned Here as 43
	fmt.Println(z) // Default value is zero
}

// To find out Type of Variable then Use "%T" and To Wrire ParaGraph the use `` Symbol
func printValuesTypes() {
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(b)
	fmt.Printf("%T\n", b)
}

func exitHere() {
	fmt.Println("Exit code here")
}
