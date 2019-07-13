package main

import "fmt"

func main() {

	// Loop - init, condition, post
	initConditionPostLoop()

	// Loop - Nested Loops
	nestedLoops()

	// Loop - for Statement: Three ways we can do loops in Go
	threeWayLoopOnlyByFor()

	// Loop - break & continue
	breakContinueLoop()

	// Loop - Printing ASCII from 33 to 122
	printASCII()

	// Conditional - if Statement
	conditionalIfStatement()

	// Conditional - if, else if, else
	conditionalIfElseifElse()

	// Conditional - switch statement
	switchStatement()

	// Conditional - switch FallThrough Statement
	switchFallthroughStatement()

}

// for loop - initialization, condition and post
func initConditionPostLoop() {
	/* SYNTAX:
	for init; condition; post {
		fmt.Println("Hello, playgroud")
	}
	*/
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
}

// for loop - Nested Loop
func nestedLoops() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("The outer loop: %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\t The inner loop: %d\n", j)
		}
	}
}

// There are three ways you can do loops in Go - they all just use the "for"
// keyword (1)
func threeWayLoopOnlyByFor() {
	// Using init; condition; post
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	// Using FOR condition
	i := 0
	for i < 8 {
		fmt.Println("The variable i is less than 8")
		i++
	}
	x := 0
	// Using only FOR
	for {
		if x > 10 {
			break
		}
		fmt.Println(x)
		x++
	}
}

//Loop - break & continue
func breakContinueLoop() {
	// Declare Variable and Assign Value
	x := 0

	for {
		x++
		if x > 100 {
			break
		}

		if x%2 != 0 {
			continue
		}

		fmt.Println(x)
	}
}

// Loop - Printing ASCII from 33 to 122
func printASCII() {
	x := 33
	for {
		if x > 122 {
			break
		}
		fmt.Printf("%v corresponds to %+q in ASCII \n", x, x)
		x++
	}
}

// Conditional - if Statement
func conditionalIfStatement() {
	if true {
		fmt.Println("001")
	}

	if false {
		fmt.Println("002")
	}

	if !true {
		fmt.Println("003")
	}

	if !false {
		fmt.Println("004")
	}

	if 2 == 2 {
		fmt.Println("005")
	}

	// Execution flow start from Left to Right
	if x := 42; x == 42 {
		fmt.Println("001")
	}
}

// Conditional - if, else if, else
func conditionalIfElseifElse() {
	// IF ELSE
	x := 42
	if x == 40 {
		fmt.Println("our value was 40")
	} else {
		fmt.Println("our value was not 40")
	}

	// IF, ELSE IF, ELSE
	y := 124
	if y == 40 {
		fmt.Println("our value was 40")
	} else if y == 41 {
		fmt.Println("our value was 41")
	} else {
		fmt.Println("our value was not 40 or 41")
	}

	// IF, ELSE IF, ELSE IF, ELSE IF, ELSE
	z := 99
	if z == 40 {
		fmt.Println("our value was 40")
	} else if z == 99 {
		fmt.Println("our value was 99")
	} else if z == 42 {
		fmt.Println("our value was 42")
	} else if z == 43 {
		fmt.Println("our value was 43")
	} else {
		fmt.Println("our value was not 40, 99, 42, or 43")
	}
}

// Conditional - switch statement using default Keyword
func switchStatement() {
	n := "Deep"
	switch n {
	case "Deepraj", "Deep", "Do No":
		fmt.Println("miss money or bond or dr no")
	case "M":
		fmt.Println("m")
	case "Q":
		fmt.Println("this is q ")
	default:
		fmt.Println("this is default")
	}

}

// Conditional - switch statement using fallthrough
// Normally we should not use this.
func switchFallthroughStatement() {
	n := "Deep"
	switch n {
	case "Deepraj", "Deep", "Do No":
		fmt.Println("miss money or bond or dr no \n")
		fallthrough
	case "M":
		fmt.Println("m \n")
		fallthrough
	case "Q":
		fmt.Println("this is q \n")
		fallthrough
	default:
		fmt.Println("this is default \n")
	}
}

/* Result:-

0
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
28
29
30
31
32
33
34
35
36
37
38
39
40
41
42
43
44
45
46
47
48
49
50
51
52
53
54
55
56
57
58
59
60
61
62
63
64
65
66
67
68
69
70
71
72
73
74
75
76
77
78
79
80
81
82
83
84
85
86
87
88
89
90
91
92
93
94
95
96
97
98
99
100
The outer loop: 0
                 The inner loop: 0
                 The inner loop: 1
                 The inner loop: 2
The outer loop: 1
                 The inner loop: 0
                 The inner loop: 1
                 The inner loop: 2
The outer loop: 2
                 The inner loop: 0
                 The inner loop: 1
                 The inner loop: 2
The outer loop: 3
                 The inner loop: 0
                 The inner loop: 1
                 The inner loop: 2
The outer loop: 4
                 The inner loop: 0
                 The inner loop: 1
                 The inner loop: 2
The outer loop: 5
                 The inner loop: 0
                 The inner loop: 1
                 The inner loop: 2
The outer loop: 6
                 The inner loop: 0
                 The inner loop: 1
                 The inner loop: 2
The outer loop: 7
                 The inner loop: 0
                 The inner loop: 1
                 The inner loop: 2
The outer loop: 8
                 The inner loop: 0
                 The inner loop: 1
                 The inner loop: 2
The outer loop: 9
                 The inner loop: 0
                 The inner loop: 1
                 The inner loop: 2
The outer loop: 10
                 The inner loop: 0
                 The inner loop: 1
                 The inner loop: 2
0
1
2
3
4
5
The variable i is less than 8
The variable i is less than 8
The variable i is less than 8
The variable i is less than 8
The variable i is less than 8
The variable i is less than 8
The variable i is less than 8
The variable i is less than 8
0
1
2
3
4
5
6
7
8
9
10
2
4
6
8
10
12
14
16
18
20
22
24
26
28
30
32
34
36
38
40
42
44
46
48
50
52
54
56
58
60
62
64
66
68
70
72
74
76
78
80
82
84
86
88
90
92
94
96
98
100
33 corresponds to '!' in ASCII
34 corresponds to '"' in ASCII
35 corresponds to '#' in ASCII
36 corresponds to '$' in ASCII
37 corresponds to '%' in ASCII
38 corresponds to '&' in ASCII
39 corresponds to '\'' in ASCII
40 corresponds to '(' in ASCII
41 corresponds to ')' in ASCII
42 corresponds to '*' in ASCII
43 corresponds to '+' in ASCII
44 corresponds to ',' in ASCII
45 corresponds to '-' in ASCII
46 corresponds to '.' in ASCII
47 corresponds to '/' in ASCII
48 corresponds to '0' in ASCII
49 corresponds to '1' in ASCII
50 corresponds to '2' in ASCII
51 corresponds to '3' in ASCII
52 corresponds to '4' in ASCII
53 corresponds to '5' in ASCII
54 corresponds to '6' in ASCII
55 corresponds to '7' in ASCII
56 corresponds to '8' in ASCII
57 corresponds to '9' in ASCII
58 corresponds to ':' in ASCII
59 corresponds to ';' in ASCII
60 corresponds to '<' in ASCII
61 corresponds to '=' in ASCII
62 corresponds to '>' in ASCII
63 corresponds to '?' in ASCII
64 corresponds to '@' in ASCII
65 corresponds to 'A' in ASCII
66 corresponds to 'B' in ASCII
67 corresponds to 'C' in ASCII
68 corresponds to 'D' in ASCII
69 corresponds to 'E' in ASCII
70 corresponds to 'F' in ASCII
71 corresponds to 'G' in ASCII
72 corresponds to 'H' in ASCII
73 corresponds to 'I' in ASCII
74 corresponds to 'J' in ASCII
75 corresponds to 'K' in ASCII
76 corresponds to 'L' in ASCII
77 corresponds to 'M' in ASCII
78 corresponds to 'N' in ASCII
79 corresponds to 'O' in ASCII
80 corresponds to 'P' in ASCII
81 corresponds to 'Q' in ASCII
82 corresponds to 'R' in ASCII
83 corresponds to 'S' in ASCII
84 corresponds to 'T' in ASCII
85 corresponds to 'U' in ASCII
86 corresponds to 'V' in ASCII
87 corresponds to 'W' in ASCII
88 corresponds to 'X' in ASCII
89 corresponds to 'Y' in ASCII
90 corresponds to 'Z' in ASCII
91 corresponds to '[' in ASCII
92 corresponds to '\\' in ASCII
93 corresponds to ']' in ASCII
94 corresponds to '^' in ASCII
95 corresponds to '_' in ASCII
96 corresponds to '`' in ASCII
97 corresponds to 'a' in ASCII
98 corresponds to 'b' in ASCII
99 corresponds to 'c' in ASCII
100 corresponds to 'd' in ASCII
101 corresponds to 'e' in ASCII
102 corresponds to 'f' in ASCII
103 corresponds to 'g' in ASCII
104 corresponds to 'h' in ASCII
105 corresponds to 'i' in ASCII
106 corresponds to 'j' in ASCII
107 corresponds to 'k' in ASCII
108 corresponds to 'l' in ASCII
109 corresponds to 'm' in ASCII
110 corresponds to 'n' in ASCII
111 corresponds to 'o' in ASCII
112 corresponds to 'p' in ASCII
113 corresponds to 'q' in ASCII
114 corresponds to 'r' in ASCII
115 corresponds to 's' in ASCII
116 corresponds to 't' in ASCII
117 corresponds to 'u' in ASCII
118 corresponds to 'v' in ASCII
119 corresponds to 'w' in ASCII
120 corresponds to 'x' in ASCII
121 corresponds to 'y' in ASCII
122 corresponds to 'z' in ASCII
001
004
005
001
our value was not 40
our value was not 40 or 41
our value was 99
miss money or bond or dr no
miss money or bond or dr no

m

this is q

this is default

*/
