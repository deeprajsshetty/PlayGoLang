package main

import "fmt"

// Declaration of Structs
type person struct {
	firstName string
	lastName  string
	age       int
}

// Declaration of one more Structs to Demostrate Embedded Struct
type doctorPerson struct {
	person
	isDoctor bool
}

func main() {
	// Creating STRUCT person TYPE of VALUES
	fullName1 := person{
		firstName: "FIRST1",
		lastName:  "LAST1",
		age:       32,
	}

	fullName2 := person{
		firstName: "SECOND2",
		lastName:  "LAST2",
		age:       38,
	}

	// Printing Values assigned to Struct Person
	fmt.Println(fullName1)
	fmt.Println(fullName2)

	fmt.Println(fullName1.firstName, fullName1.lastName, fullName1.age)
	fmt.Println(fullName2.firstName, fullName2.lastName, fullName2.age)

	// Embedded Structs
	doctorName := doctorPerson{
		person: person{
			firstName: "DEEP",
			lastName:  "S",
			age:       33,
		},
		isDoctor: true,
	}

	fmt.Println(doctorName)
	fmt.Println(doctorName.firstName, doctorName.lastName, doctorName.age, doctorName.isDoctor)

}

/*	OutPut:
{FIRST1 LAST1 32}
{SECOND2 LAST2 38}
FIRST1 LAST1 32
SECOND2 LAST2 38
{{DEEP S 33} true}
DEEP S 33 true
 */
