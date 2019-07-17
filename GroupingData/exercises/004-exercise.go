package main

import "fmt"

func main() {
	// Created Slice with Values
	a4 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// Append 52 value to Initial Slice
	a4 = append(a4, 52)

	// Print Slice
	fmt.Println(a4)

	// In ONE STATEMENT append 53, 54 and 55 to a4
	a4 = append(a4, 53, 54, 55)
	fmt.Println(a4)

	// Add One More slice to a4 and Printing
	b4 := []int{56, 57, 58, 59, 60}
	a4 = append(a4, b4...)
	fmt.Println(a4)
}

/*	Output:
	[42 43 44 45 46 47 48 49 50 51 52]
	[42 43 44 45 46 47 48 49 50 51 52 53 54 55]
	[42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60]
*/
