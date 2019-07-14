package main

import "fmt"

func main() {
	// Print out the remainder (modulus) which is found for
	// each number between 10 and 100 when it is divided by 4.
	for i := 10; i <= 100; i++ {
		fmt.Printf("When %v is divided by 4, the remainder (aka modulus) is %v\n", i, i%4)
	}
}
/* Result:

When 10 is divided by 4, the remainder (aka modulus) is 2
When 11 is divided by 4, the remainder (aka modulus) is 3
When 12 is divided by 4, the remainder (aka modulus) is 0
When 13 is divided by 4, the remainder (aka modulus) is 1
When 14 is divided by 4, the remainder (aka modulus) is 2
When 15 is divided by 4, the remainder (aka modulus) is 3
When 16 is divided by 4, the remainder (aka modulus) is 0
When 17 is divided by 4, the remainder (aka modulus) is 1
When 18 is divided by 4, the remainder (aka modulus) is 2
When 19 is divided by 4, the remainder (aka modulus) is 3
When 20 is divided by 4, the remainder (aka modulus) is 0
When 21 is divided by 4, the remainder (aka modulus) is 1
When 22 is divided by 4, the remainder (aka modulus) is 2
When 23 is divided by 4, the remainder (aka modulus) is 3
When 24 is divided by 4, the remainder (aka modulus) is 0
When 25 is divided by 4, the remainder (aka modulus) is 1
When 26 is divided by 4, the remainder (aka modulus) is 2
When 27 is divided by 4, the remainder (aka modulus) is 3
When 28 is divided by 4, the remainder (aka modulus) is 0
When 29 is divided by 4, the remainder (aka modulus) is 1
When 30 is divided by 4, the remainder (aka modulus) is 2
When 31 is divided by 4, the remainder (aka modulus) is 3
When 32 is divided by 4, the remainder (aka modulus) is 0
When 33 is divided by 4, the remainder (aka modulus) is 1
When 34 is divided by 4, the remainder (aka modulus) is 2
When 35 is divided by 4, the remainder (aka modulus) is 3
When 36 is divided by 4, the remainder (aka modulus) is 0
When 37 is divided by 4, the remainder (aka modulus) is 1
When 38 is divided by 4, the remainder (aka modulus) is 2
When 39 is divided by 4, the remainder (aka modulus) is 3
When 40 is divided by 4, the remainder (aka modulus) is 0
When 41 is divided by 4, the remainder (aka modulus) is 1
When 42 is divided by 4, the remainder (aka modulus) is 2
When 43 is divided by 4, the remainder (aka modulus) is 3
When 44 is divided by 4, the remainder (aka modulus) is 0
When 45 is divided by 4, the remainder (aka modulus) is 1
When 46 is divided by 4, the remainder (aka modulus) is 2
When 47 is divided by 4, the remainder (aka modulus) is 3
When 48 is divided by 4, the remainder (aka modulus) is 0
When 49 is divided by 4, the remainder (aka modulus) is 1
When 50 is divided by 4, the remainder (aka modulus) is 2
When 51 is divided by 4, the remainder (aka modulus) is 3
When 52 is divided by 4, the remainder (aka modulus) is 0
When 53 is divided by 4, the remainder (aka modulus) is 1
When 54 is divided by 4, the remainder (aka modulus) is 2
When 55 is divided by 4, the remainder (aka modulus) is 3
When 56 is divided by 4, the remainder (aka modulus) is 0
When 57 is divided by 4, the remainder (aka modulus) is 1
When 58 is divided by 4, the remainder (aka modulus) is 2
When 59 is divided by 4, the remainder (aka modulus) is 3
When 60 is divided by 4, the remainder (aka modulus) is 0
When 61 is divided by 4, the remainder (aka modulus) is 1
When 62 is divided by 4, the remainder (aka modulus) is 2
When 63 is divided by 4, the remainder (aka modulus) is 3
When 64 is divided by 4, the remainder (aka modulus) is 0
When 65 is divided by 4, the remainder (aka modulus) is 1
When 66 is divided by 4, the remainder (aka modulus) is 2
When 67 is divided by 4, the remainder (aka modulus) is 3
When 68 is divided by 4, the remainder (aka modulus) is 0
When 69 is divided by 4, the remainder (aka modulus) is 1
When 70 is divided by 4, the remainder (aka modulus) is 2
When 71 is divided by 4, the remainder (aka modulus) is 3
When 72 is divided by 4, the remainder (aka modulus) is 0
When 73 is divided by 4, the remainder (aka modulus) is 1
When 74 is divided by 4, the remainder (aka modulus) is 2
When 75 is divided by 4, the remainder (aka modulus) is 3
When 76 is divided by 4, the remainder (aka modulus) is 0
When 77 is divided by 4, the remainder (aka modulus) is 1
When 78 is divided by 4, the remainder (aka modulus) is 2
When 79 is divided by 4, the remainder (aka modulus) is 3
When 80 is divided by 4, the remainder (aka modulus) is 0
When 81 is divided by 4, the remainder (aka modulus) is 1
When 82 is divided by 4, the remainder (aka modulus) is 2
When 83 is divided by 4, the remainder (aka modulus) is 3
When 84 is divided by 4, the remainder (aka modulus) is 0
When 85 is divided by 4, the remainder (aka modulus) is 1
When 86 is divided by 4, the remainder (aka modulus) is 2
When 87 is divided by 4, the remainder (aka modulus) is 3
When 88 is divided by 4, the remainder (aka modulus) is 0
When 89 is divided by 4, the remainder (aka modulus) is 1
When 90 is divided by 4, the remainder (aka modulus) is 2
When 91 is divided by 4, the remainder (aka modulus) is 3
When 92 is divided by 4, the remainder (aka modulus) is 0
When 93 is divided by 4, the remainder (aka modulus) is 1
When 94 is divided by 4, the remainder (aka modulus) is 2
When 95 is divided by 4, the remainder (aka modulus) is 3
When 96 is divided by 4, the remainder (aka modulus) is 0
When 97 is divided by 4, the remainder (aka modulus) is 1
When 98 is divided by 4, the remainder (aka modulus) is 2
When 99 is divided by 4, the remainder (aka modulus) is 3
When 100 is divided by 4, the remainder (aka modulus) is 0

 */