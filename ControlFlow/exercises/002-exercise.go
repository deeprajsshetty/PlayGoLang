package main

import "fmt"

func main() {

	// Print every rune code point of the uppercase alphabet three times.
	for i2 := 65; i2 <= 90; i2++ {
		fmt.Println(i2)
		for j2 := 0; j2 < 3; j2++ {
			fmt.Printf("\t%#U\n", j2)
		}

	}
}

/* OutPut:-
65
        U+0000
        U+0001
        U+0002
66
        U+0000
        U+0001
        U+0002
67
        U+0000
        U+0001
        U+0002
68
        U+0000
        U+0001
        U+0002
69
        U+0000
        U+0001
        U+0002
70
        U+0000
        U+0001
        U+0002
71
        U+0000
        U+0001
        U+0002
72
        U+0000
        U+0001
        U+0002
73
        U+0000
        U+0001
        U+0002
74
        U+0000
        U+0001
        U+0002
75
        U+0000
        U+0001
        U+0002
76
        U+0000
        U+0001
        U+0002
77
        U+0000
        U+0001
        U+0002
78
        U+0000
        U+0001
        U+0002
79
        U+0000
        U+0001
        U+0002
80
        U+0000
        U+0001
        U+0002
81
        U+0000
        U+0001
        U+0002
82
        U+0000
        U+0001
        U+0002
83
        U+0000
        U+0001
        U+0002
84
        U+0000
        U+0001
        U+0002
85
        U+0000
        U+0001
        U+0002
86
        U+0000
        U+0001
        U+0002
87
        U+0000
        U+0001
        U+0002
88
        U+0000
        U+0001
        U+0002
89
        U+0000
        U+0001
        U+0002
90
        U+0000
        U+0001
        U+0002

*/