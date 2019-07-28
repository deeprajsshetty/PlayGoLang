package main

import "fmt"

// Declare STRUCT which will Hold Values of SLICES Also
type person struct {
	firstName   string
	lastName    string
	favFlavours []string
}

func main() {

	// Creating Values of STRUCT person with Slice DataStructure
	p1 := person{
		firstName: "James",
		lastName:  "Bond",
		favFlavours: []string{
			"chocolate",
			"martini",
			"rum and coke",
		},
	}

	p2 := person{
		firstName: "Miss",
		lastName:  "Moneypenny",
		favFlavours: []string{
			"strawberry",
			"vanilla",
			"capuccino",
		},
	}

	// Printing Values of STRUCT person which having p1 AND P2.
	// Slices Printing using RANG
	fmt.Println(p1.firstName)
	fmt.Println(p1.lastName)
	for i, v := range p1.favFlavours {
		fmt.Println(i, v)
	}

	fmt.Println(p2.firstName)
	fmt.Println(p2.lastName)
	for i, v := range p2.favFlavours {
		fmt.Println(i, v)
	}
}

/* OUTPUT:
James
Bond
0 chocolate
1 martini
2 rum and coke
Miss
Moneypenny
0 strawberry
1 vanilla
2 capuccino

 */
