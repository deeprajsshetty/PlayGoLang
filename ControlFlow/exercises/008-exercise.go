package main

import "fmt"

func main() {
	// Example to demonstrate SWITCH statement without Expression
	switch {
	case false:
		fmt.Println("Should not Print")
	case true:
		fmt.Println("Should Print") //Result: Should Print
	}
}
