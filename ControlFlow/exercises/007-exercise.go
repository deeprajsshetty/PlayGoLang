package main

import "fmt"

func main() {
	// Simple ELSE IF example
	a7 := "Test User"

	if a7 == "Not Test User" {
		fmt.Println(a7)

	} else if a7 == "Test User" {
		fmt.Println(a7) // Test User
	} else {
		fmt.Println("Neither")
	}
}
