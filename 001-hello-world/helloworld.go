package main

import "fmt"

func main() {
	// Printing Messages
	printHelloWorld()

	// Printing Numbers in Loops
	printNumberLoops()

	// Declare Variable and Assigning Values to Variable
	declareAssignVariable()

	// Exit Here
	exitHere()
}

func printHelloWorld() {
	fmt.Println("Hello World")
}

func printNumberLoops() {
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
}

func exitHere() {
	fmt.Println("Exit code here")
}

func declareAssignVariable() {
	x := 10
	fmt.Println(x)
	x = 30
	fmt.Println(x)
	y := 20 + 30
	fmt.Println(y)
}
