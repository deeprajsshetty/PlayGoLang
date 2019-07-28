package main

import (
	"fmt"
)

func main() {
	// Declaring Anonymous Struct and Printing Values
	// Used Embedded Struct and MAP and SLICES
	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Moneypenny": 555,
			"Q":          777,
			"M":          888,
		},
		favDrinks: []string{
			"Martini",
			"Water",
		},
	}
	fmt.Println(s.first)
	fmt.Println(s.friends)
	fmt.Println(s.favDrinks)

	for k, v := range s.friends {
		fmt.Println(k, v)
	}

	for i, val := range s.favDrinks {
		fmt.Println(i, val)
	}
}

/* OUTPUT:
James
map[M:888 Moneypenny:555 Q:777]
[Martini Water]
Moneypenny 555
Q 777
M 888
0 Martini
1 Water
*/
