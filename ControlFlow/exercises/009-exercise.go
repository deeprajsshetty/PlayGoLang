package main

import "fmt"

func main() {
	// Example that uses the SWITCH Statement using the SWITCH EXPRESSION
	a9 := "Cricket"

	switch a9 {
	case "Tennis":
		fmt.Printf("Selected Sports is %v\n", a9)
	case "Cricket":
		fmt.Printf("Selected Sports is %v\n", a9)
	default:
		fmt.Println("Nothing Selected")
	}
}

/* Output:
Selected Sports is Cricket
*/
