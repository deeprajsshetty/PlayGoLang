package main

import "fmt"

func main() {
	// Created Array with Static 5 Values and Type of Int
	a1 := []int{1, 2, 3, 4, 5}

	// Using Range printing Index and Value with TYPE
	for i, v := range a1 {
		fmt.Printf("Index: %v\tValue: %v\n", i, v)
	}
	//Printing ARRAY TYPE
	fmt.Printf("%T\n", a1)
}

/* Output:
Index: 0        Value: 1
Index: 1        Value: 2
Index: 2        Value: 3
Index: 3        Value: 4
Index: 4        Value: 5
[]int
*/
