package main

import "fmt"

func main() {
	// Creating a Slice of a Slice of string ([][]string)
	// Store that in Multi-Dimensional Slice
	a7 := []string{"User1", "User2", "User3"}
	b7 := []string{"User4", "User5", "User6"}

	c7 := [][]string{a7, b7}
	fmt.Println(c7)

	// Range over the records, then range over the data in each record.
	for i, d7 := range c7 {
		fmt.Println("record: ", i)
		for j, val := range d7 {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}
	}
}

/*	Output:
	[[User1 User2 User3] [User4 User5 User6]]
	record:  0
         index position: 0       value:          User1
         index position: 1       value:          User2
         index position: 2       value:          User3
	record:  1
         index position: 0       value:          User4
         index position: 1       value:          User5
         index position: 2       value:          User6

*/
