package main

import "fmt"

func main() {

	// DELETE from SLICE
	a5 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(append(a5))

	a5 = append(a5[:3], a5[6:]...) // [42, 43, 44, 48, 49, 50, 51]
	fmt.Println(a5)
}

/*	Output:
	[42 43 44 45 46 47 48 49 50 51]
	[42 43 44 48 49 50 51]
*/
