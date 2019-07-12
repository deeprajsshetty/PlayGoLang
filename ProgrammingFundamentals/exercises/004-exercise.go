package main

import "fmt"

func main() {

	// Assign an int to a VARIABLE "a4"
	a4 := 99

	// Printing "a4" in Decimal, Binary and Hex
	fmt.Printf("%d\t%b\t%#x\n", a4, a4, a4) // Result: 99	1100011	0x63

	// shifts the bits of that int over 1 position to the left,
	// and assigns that to a variable
	a4 = a4 << 1

	// Now Print "a4" in Decimal, Binary and Hex
	fmt.Printf("%d\t%b\t%#x", a4, a4, a4) // Result: 198	11000110	0xc6
}
