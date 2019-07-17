package main

import "fmt"

func main() {
	// Created MAP of String and TYPE of Slice
	m8 := map[string][]string{
		"USER1": []string{`TEST1`, `TEST2`, `TEST3`},
		"USER2": []string{`TEST4`, `TEST5`, `TEST6`},
		"USER3": []string{`TEST7`, `TEST8`, `TEST9`},
	}

	// Added New Map to Existing Map
	m8["USER4"] = []string{`TEST10`, `TEST11`, `TEST12`}

	// Printing MAP which contains Type of Slice
	for k8, v8 := range m8 {
		fmt.Println("This is the record for", k8)
		for i8, v8 := range v8 {
			fmt.Println("\t", i8, v8)
		}
	}
}

/* Output:
This is the record for USER1
         0 TEST1
         1 TEST2
         2 TEST3
This is the record for USER2
         0 TEST4
         1 TEST5
         2 TEST6
This is the record for USER3
         0 TEST7
         1 TEST8
         2 TEST9
This is the record for USER4
         0 TEST10
         1 TEST11
         2 TEST12

*/
