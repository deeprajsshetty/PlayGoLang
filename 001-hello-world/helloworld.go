package main

import "fmt"

func main() {
	// Printing Messages
	printHelloWorld()

	// Printing Numbers in Loops
	printNumberLoops()

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
