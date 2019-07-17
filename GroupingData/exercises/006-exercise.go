package main

import "fmt"

func main() {
	// Using Make define len and Capacity for Slice and
	// Store all United States
	a6 := make([]string, 50, 500)
	fmt.Println(len(a6))
	fmt.Println(cap(a6))

	// put a value into each of the 50 positions in the length of the slice
	// we are putting values into the 50 positions that are the length
	// of the slice
	for i := 0; i < 50; i++ {
		a6[i] = fmt.Sprintf("Position %d", i)
	}

	fmt.Println(a6)
	fmt.Println(len(a6))
	fmt.Println(cap(a6))

	// append values to the slice grows the length of the slice
	// the underlying array "capacity" of 500 is the same
	a6 = append(a6, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`,
		` California`, ` Colorado`, ` Connecticut`, ` Delaware`,
		` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`,
		` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`,
		` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`,
		` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`,
		` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`,
		` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`,
		` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`,
		` Rhode Island`, ` South Carolina`, ` South Dakota`,
		` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`,
		` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
	fmt.Println(a6)
	fmt.Println(len(a6))
	fmt.Println(cap(a6))

	for i := 0; i < len(a6); i++ {
		fmt.Println(i, a6[i])
	}
}

/* Output:-

50
500
[Position 0 Position 1 Position 2 Position 3 Position 4 Position 5 Position 6 Position 7 Position 8 Position 9 Position 10 Position 11 Position 12 Position 13 Position 14 Position 15 Position 16 Position 17 Position 18 Position 19 Position 20 Position 21 Position 22 Position 23 Position 24 Position 25 Position 26 Position 27 Position 28 Position 29 Position 30 Position 31 Position 32 Position 33 Position 34 Position 35 Position 36 Position 37 Position 38 Position 39 Position 40 Position 41 Position 42 Position 43 Position 44 Position 45 Position 46 Position 47 Position 48 Position 49]
50
500
[Position 0 Position 1 Position 2 Position 3 Position 4 Position 5 Position 6 Position 7 Position 8 Position 9 Position 10 Position 11 Position 12 Position 13 Position 14 Position 15 Position 16 Position 17 Position 18 Position 19 Position 20 Position 21 Position 22 Position 23 Position 24 Position 25 Position 26 Position 27 Position 28 Position 29 Position 30 Position 31 Position 32 Position 33 Position 34 Position 35 Position 36 Position 37 Position 38 Position 39 Position 40 Position 41 Position 42 Position 43 Position 44 Position 45 Position 46 Position 47 Position 48 Position 49  Alabama  Alaska  Arizona  Arkansas  California  Colorado  Connecticut  Delaware  Florida  Georgia  Hawaii  Idaho  Illinois  Indiana  Iowa  Kansas  Kentucky  Louisiana  Maine  Maryland  Massachusetts  Michigan  Minnesota  Mississippi  Missouri  Montana  Nebraska  Nevada  New Hampshire  New Jersey  New Mexico  New York  North Carolina  North Dakota  Ohio  Oklahoma  Oregon  Pennsylvania  Rhode Island  South Carolina  South Dakota  Tennessee  Texas  Utah  Vermont  Virginia  Washington  West Virginia  Wisconsin  Wyoming]
100
500
0 Position 0
1 Position 1
2 Position 2
3 Position 3
4 Position 4
5 Position 5
6 Position 6
7 Position 7
8 Position 8
9 Position 9
10 Position 10
11 Position 11
12 Position 12
13 Position 13
14 Position 14
15 Position 15
16 Position 16
17 Position 17
18 Position 18
19 Position 19
20 Position 20
21 Position 21
22 Position 22
23 Position 23
24 Position 24
25 Position 25
26 Position 26
27 Position 27
28 Position 28
29 Position 29
30 Position 30
31 Position 31
32 Position 32
33 Position 33
34 Position 34
35 Position 35
36 Position 36
37 Position 37
38 Position 38
39 Position 39
40 Position 40
41 Position 41
42 Position 42
43 Position 43
44 Position 44
45 Position 45
46 Position 46
47 Position 47
48 Position 48
49 Position 49
50  Alabama
51  Alaska
52  Arizona
53  Arkansas
54  California
55  Colorado
56  Connecticut
57  Delaware
58  Florida
59  Georgia
60  Hawaii
61  Idaho
62  Illinois
63  Indiana
64  Iowa
65  Kansas
66  Kentucky
67  Louisiana
68  Maine
69  Maryland
70  Massachusetts
71  Michigan
72  Minnesota
73  Mississippi
74  Missouri
75  Montana
76  Nebraska
77  Nevada
78  New Hampshire
79  New Jersey
80  New Mexico
81  New York
82  North Carolina
83  North Dakota
84  Ohio
85  Oklahoma
86  Oregon
87  Pennsylvania
88  Rhode Island
89  South Carolina
90  South Dakota
91  Tennessee
92  Texas
93  Utah
94  Vermont
95  Virginia
96  Washington
97  West Virginia
98  Wisconsin
99  Wyoming

*/
