package main

import "fmt"

func main() {

	// Array - In some instances we might use Array. Mostly the recommendation
	// use Slices instead Array
	playwithArrays()

	// Slice - It is built Top of Array. Slice allow you to group together
	// VALUES OF SAME TYPE
	playwithSlices()

	// Slice - Append to a Slice
	appendSlices()

	// Slice - Make
	// Make used to define static Size and Capacity to Slice so it behaves
	// similar to Array but Dynamically Size will grow like Slice
	// Memory Allocation will happen based on Capacity i.e when Slices Grow
	// more than defined Capacity then it will assign double the Capacity
	makeSlices()

	// Slice - multi-dimensional slice
	multiDimentionalSlice()

	// “throw away” the variable and the same thing happens
	throwAwaySlices()

	// Map - Declare, Assign Value, Add Element and Range, Delete Map
	playwithMap()
}

// Array with Code Samples
func playwithArrays() {
	// Declare Array
	var x [5]int
	fmt.Println(x) // [0 0 0 0 0]

	// Assigning Value to index 3 and Print
	x[3] = 26
	fmt.Println(x) //[0 0 0 26 0]

	// Find Length of Array
	fmt.Println(len(x)) // 5
}

// Slice - Declare Slice, Append Slice, Delete and Make slice
// Multi-dimensional Slice
func playwithSlices() {

	// Declare and Assign Values to Simple Slice
	x := []int{4, 5, 6, 8}
	fmt.Println(x)      // [4 5 6 8]
	fmt.Println(len(x)) // 4

	// Slice - for range
	for i, v := range x {
		fmt.Println(i, v)
	}

	// Slice - Slicing a Slice
	// We can cut part of the Slice
	fmt.Println(x[1:])  // [5 6 8]
	fmt.Println(x[1:3]) // [5 6]
}

// Slice - Append to a Slice
func appendSlices() {

	// Declare and Assign Value to Slice
	x := []int{2, 4, 6, 8}
	fmt.Println(x) // [2 4 6 8]

	// Append Value to Slice x
	x = append(x, 10, 12, 14, 16)
	fmt.Println(x) // [2 4 6 8 10 12 14 16]

	// Create New Slice called y and append to x
	y := []int{22, 24, 26, 28}
	x = append(x, y...)
	fmt.Println(x) // [2 4 6 8 10 12 14 16 22 24 26 28]

	// Deleting from Slices using Append
	x = append(x[:2], x[4:]...) // [2 4 10 12 14 16 22 24 26 28]
	fmt.Println(x)
}

// Slice - Make
func makeSlices() {

	// Using Make define len and Capacity to Slice and Print Result
	x := make([]int, 10, 12)
	fmt.Println(x)      // 	[0 0 0 0 0 0 0 0 0 0]
	fmt.Println(len(x)) // 10
	fmt.Println(cap(x)) // 12
	x[0] = 42
	x[9] = 999
	// Assign value to 0 and 9 index in Slice
	fmt.Println(x)      // [42 0 0 0 0 0 0 0 0 999]
	fmt.Println(len(x)) // 10
	fmt.Println(cap(x)) // 12

	// Append one more value to Slice
	x = append(x, 3423)

	fmt.Println(x)      // [42 0 0 0 0 0 0 0 0 999 3423]
	fmt.Println(len(x)) // 10
	fmt.Println(cap(x)) // 12

	// Append one more value to Slice
	x = append(x, 3424)

	fmt.Println(x)      // [42 0 0 0 0 0 0 0 0 999 3423 3424]
	fmt.Println(len(x)) // 10
	fmt.Println(cap(x)) // 12

	// Append one more value to Slice
	x = append(x, 3425)

	fmt.Println(x)      // [42 0 0 0 0 0 0 0 0 999 3423 3424 3425]
	fmt.Println(len(x)) // 13
	// Capacity Doubled here as we utilized all the Capacity alocated
	// first time.
	fmt.Println(cap(x)) // 24
}

// Slice - multi-dimensional slice
func multiDimentionalSlice() {
	// Declaring Seperate Slices of TYPE String
	slice1 := []string{"User 1", "User 2", "User 3"}
	fmt.Println(slice1) // [User 1 User 2 User 3]
	slice2 := []string{"User 4", "User 5", "User 6"}
	fmt.Println(slice2) // [User 4 User 5 User 6]

	// Assigning Above mentioned Two arrays to Multidemsional Slice
	multiDems := [][]string{slice1, slice2}
	fmt.Println(multiDems) // [[User 1 User 2 User 3][User 4 User 5 User 6]]
}

// “throw away” the variable and the same thing happens
func throwAwaySlices() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(x) // [1 2 3 4 5 6 7 8 9]

	// throw Away
	_ = append(x[:2], x[5:]...)
	fmt.Println(x) // [1 2 3 6 7 8 9 8 9]
}

// Map - Declare, Assign Value, Add Element and Range, Delete Map
func playwithMap() {
	// Declaring Map and Assigning Values and Print
	m := map[string]int{
		"Test User1": 34,
		"Test User2": 99,
	}
	fmt.Println(m)               // map[Test User1:34 Test User2:99]
	fmt.Println(m["Test User1"]) // Test User1
	fmt.Println(m["Test User2"]) // Test User2

	// In case Key not available then we can check.
	fmt.Println(m["Not User"]) // 0
	v, ok := m["Not User"]
	fmt.Println(v)  // 0
	fmt.Println(ok) // false

	if v, ok := m["Not User"]; ok {
		fmt.Println(v)
	}

	// Adding Element to Map and Using Range to Print
	m["Test User3"] = 22
	fmt.Println(m["Test User3"]) // Test User3

	for k, v := range m {
		fmt.Println(k, v) // Output Pasted Below
	}

	// Deleting Map
	if v, ok := m["Test User3"]; ok {
		fmt.Println("value:", v)
		delete(m, "Test User3")
	}
	fmt.Println(m) // map[Test User1:34 Test User2:99]
}

/* Output Result:

[0 0 0 0 0]
[0 0 0 26 0]
5
[4 5 6 8]
4
0 4
1 5
2 6
3 8
[5 6 8]
[5 6]
[2 4 6 8]
[2 4 6 8 10 12 14 16]
[2 4 6 8 10 12 14 16 22 24 26 28]
[2 4 10 12 14 16 22 24 26 28]
[0 0 0 0 0 0 0 0 0 0]
10
12
[42 0 0 0 0 0 0 0 0 999]
10
12
[42 0 0 0 0 0 0 0 0 999 3423]
11
12
[42 0 0 0 0 0 0 0 0 999 3423 3424]
12
12
[42 0 0 0 0 0 0 0 0 999 3423 3424 3425]
13
24
[User 1 User 2 User 3]
[User 4 User 5 User 6]
[[User 1 User 2 User 3] [User 4 User 5 User 6]]
[1 2 3 4 5 6 7 8 9]
[1 2 6 7 8 9 7 8 9]
map[Test User1:34 Test User2:99]
34
99
0
0
false
22
Test User1 34
Test User2 99
Test User3 22
value: 22
map[Test User1:34 Test User2:99]

 */
