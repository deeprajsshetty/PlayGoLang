package main

import "fmt"

func main() {
	// Create and Print a variable of type string using a raw
	// string literal.
	a5 := `" "In an age of extreme polarization and dearth of good will, music is a powerful force that brings people together….
			-Aaron Davis, letter in Billboard.com, 26 Apr. 2019`
	fmt.Println(a5)
}

/* Ouput Result:
	" "In an age of extreme polarization and dearth of good will, music is a powerful force that brings people together….
                        -Aaron Davis, letter in Billboard.com, 26 Apr. 2019
 */