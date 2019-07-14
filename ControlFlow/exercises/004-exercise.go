package main

import "fmt"

func main() {
	// Loop - Using only for - We don't have While loop in go
	// Alternative for While Loop
	a4 := 1
	for {
		if a4 > 9 {
			break
		}

		fmt.Println(a4) // 1 2 3 4 5 6 7 8 9
		a4++
	}
}
