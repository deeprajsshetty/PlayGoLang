package main

import "fmt"

// Slice built on Array So looks similar and Normally we won't use Array
func main() {
	// Created Slice with 10 Values and Type of Int
	a2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Using Range printing Index and Value with TYPE
	for i, v := range a2 {
		fmt.Printf("Index: %v\tValue: %v\n", i, v)
	}
	//Printing SLICE TYPE
	fmt.Printf("%T\n", a2)
}

/* Output:
Index: 0        Value: 1
Index: 1        Value: 2
Index: 2        Value: 3
Index: 3        Value: 4
Index: 4        Value: 5
Index: 5        Value: 6
Index: 6        Value: 7
Index: 7        Value: 8
Index: 8        Value: 9
[]int

*/
