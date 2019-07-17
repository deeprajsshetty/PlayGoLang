package main

import "fmt"

func main() {
	a3 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(a3)

	// Slice of Slice
	fmt.Println(a3[:5])
	fmt.Println(a3[5:])
	fmt.Println(a3[2:7])
	fmt.Println(a3[1:6])
}

/*	Output:
	[42 43 44 45 46 47 48 49 50 51]
	[42 43 44 45 46]
	[47 48 49 50 51]
	[44 45 46 47 48]
	[43 44 45 46 47]
*/
