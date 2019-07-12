package main

import (
	"fmt"
	"runtime"
)

func main() {

	// Boolean Type
	boolType()

	// Numeric Types (int and float64)
	numericTypes()

	// runtime Packages
	runtimePackage()

	// String Types
	stringTypes()

	// Constant Types
	constantTypes()

	// iota Type
	iotaType()

	// Bit Shifting
	bitShifting()
}

// A bool is a binary TYPE having two possible values of either
// “true” and “false.”
func boolType() {
	a := 4
	b := 8
	c := 4
	fmt.Println(a == c) // Result: true
	fmt.Println(b <= a) // Result: false
}

// Numeric Types used to represent integer and decimal Values
func numericTypes() {
	x := 99    // Integer
	y := 11.22 // floating point

	fmt.Printf("%T\t x = %v\n", x, x) // Result: int	x = 99
	fmt.Printf("%T\t y = %v\n", y, y) // Result: float64	y = 11.22
}

// Runtime Package used to identify the System Configuration like
// 64bit/32bit etc etc
func runtimePackage() {
	fmt.Println(runtime.GOOS)   // windows
	fmt.Println(runtime.GOARCH) // amd64
}

// A String value is a (Possibly empty) sequence of bytes,
// Strings are immutable: Once created, it is impossible to change
// the contents of a string
func stringTypes() {
	s := "Hello, 世界"

	// Printing the uninterpreted bytes of the string or slice
	fmt.Printf("%s\n", s)

	// Printing double-quoted string safely escaped with Go syntax
	fmt.Printf("%q\n", s)

	// Printing Hexa Decimal for the String "s"
	fmt.Printf("%x\n", s)
	fmt.Printf("---%x\n", "世")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x \n", s[i])
	}

	// Playing with Actual String
	sentence := "Go Language is future"

	// Printing String "sentence"
	fmt.Println(sentence)

	// Printing the uninterpreted bytes of the string or slice
	fmt.Printf("%s\n", sentence)

	// Printing double-quoted string safely escaped with Go syntax
	fmt.Printf("%q\n", sentence)

	// Printing base 16, with lower-case letters for a-f
	fmt.Printf("%x\n", sentence)

	// Printing base 16, with upper-case letters for a-f
	fmt.Printf("%X\n", sentence)

	// Printing Unicode format: U+1234; same as "U+%04X"
	fmt.Printf("%U\n", sentence)

	// Printing base 10 by Character by Character
	for i := 0; i < len(sentence); i++ {
		fmt.Printf("%x ", sentence[i])
	}
	fmt.Println("")

	// Printing base 10 by Character by Character
	for i := 0; i < len(sentence); i++ {
		fmt.Printf("%d ", sentence[i])
	}
	fmt.Println("")

	// Printing Unicode format: U+1234; same as "U+%04X"
	for i, v := range sentence {
		fmt.Printf("%#U \t %d", v, i)
	}

	// Printing Bytes
	bs := []byte(sentence)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i, v := range sentence {
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}

}

// Constants is a Simple and unchanging value, only exist at compile
// time. there are TYPED and UNTYPED constants
// Eg. const hello = "Hello, World" - UNTYPED
// const typedHello string = "Hello, World" - TYPED
const (
	a int     = 50
	b float64 = 87.99
	c string  = "Test User"
)

const x = 46
const y = 78.09

type intValue int
type floatValue float64

func constantTypes() {

	// Printing a, b and c Constant Values
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

	// Printing x and y values by converting Typed Variable
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	fmt.Printf("%T\n", intValue(x))
	fmt.Printf("%T\n", floatValue(y))
}

// Within a constant declaration, the predeclared identifier iota
// represents successive untyped integer constants. It is reset
// to 0 whenever the reserved word const appears in the source.
// It can be used to construct a set of related constants:

const (
	a0 = iota
	a1 = iota
	a2 = iota
	a3
	a4
	a5
)

const (
	a6 = iota
	a7
	a8
	a9
)

func iotaType() {
	fmt.Println(a0)
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	fmt.Println(a5)

	// Resetting Values from here as it declared in Seperate Const
	fmt.Println(a6)
	fmt.Println(a7)
	fmt.Println(a8)
	fmt.Println(a9)
}

// Bit shifting is when you shift binary digits to the left or right.
// We can use bit shifting, along with iota, to build some creative constants.

const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
	TB = 1 << (iota * 10) // 1 << (4 * 10)
)

func bitShifting() {
	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

	// Using iota shifting 10 digits Left
	fmt.Println("binary\t\t\t\tdecimal")
	fmt.Printf("%b\t\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
}

/* Output Result:

true
false
int      x = 99
float64  y = 11.22
windows
amd64
Hello, 世界
"Hello, 世界"
48656c6c6f2c20e4b896e7958c
---e4b896
48
65
6c
6c
6f
2c
20
e4
b8
96
e7
95
8c
Go Language is future
Go Language is future
"Go Language is future"
476f204c616e677561676520697320667574757265
476F204C616E677561676520697320667574757265
%!U(string=Go Language is future)
47 6f 20 4c 61 6e 67 75 61 67 65 20 69 73 20 66 75 74 75 72 65
71 111 32 76 97 110 103 117 97 103 101 32 105 115 32 102 117 116 117 114 101
U+0047 'G'       0U+006F 'o'     1U+0020 ' '     2U+004C 'L'     3U+0061 'a'     4U+006E 'n'     5U+0067 'g'     6U+0075 'u'     7U+0061 'a'     8U+0067 'g'     9U+0065 'e'     10U+0020 ' '    11U+0069 'i'    12U+0073 's'      13U+0020 ' '    14U+0066 'f'    15U+0075 'u'    16U+0074 't'    17U+0075 'u'    18U+0072 'r'    19U+0065 'e'    20[71 111 32 76 97 110 103 117 97 103 101 32 105 115 32 102 117 116 117 114 101]
[]uint8
at index position 0 we have hex 0x47
at index position 1 we have hex 0x6f
at index position 2 we have hex 0x20
at index position 3 we have hex 0x4c
at index position 4 we have hex 0x61
at index position 5 we have hex 0x6e
at index position 6 we have hex 0x67
at index position 7 we have hex 0x75
at index position 8 we have hex 0x61
at index position 9 we have hex 0x67
at index position 10 we have hex 0x65
at index position 11 we have hex 0x20
at index position 12 we have hex 0x69
at index position 13 we have hex 0x73
at index position 14 we have hex 0x20
at index position 15 we have hex 0x66
at index position 16 we have hex 0x75
at index position 17 we have hex 0x74
at index position 18 we have hex 0x75
at index position 19 we have hex 0x72
at index position 20 we have hex 0x65
50
87.99
Test User
int
float64
string
46
78.09
int
float64
main.intValue
main.floatValue
0
1
2
3
4
5
0
1
2
3
1024                    10000000000
1048576                 100000000000000000000
1073741824              1000000000000000000000000000000
binary                          decimal
10000000000                     1024
100000000000000000000           1048576
1000000000000000000000000000000 1073741824

*/
