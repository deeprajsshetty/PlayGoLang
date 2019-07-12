package main

import "fmt"

// Using iota, create 4 constants for the NEXT 4 years.
// Print the constant values.

const (
	a6 = 2012 + iota
	b6 = 2012 + iota
	c6
	d6
)

func main() {
	fmt.Println(a6) // 2012
	fmt.Println(b6) // 2013
	fmt.Println(c6) // 2014
	fmt.Println(d6) // 2015
}
