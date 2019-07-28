package main

import "fmt"

// Create New Type vehicle which contains fields doors and color
type vehicle struct {
	doors int
	color string
}

// Create New Types truck and sedan and Embedded to vehicle Struct
type truck struct {
	vehicle
	fourwheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	// EMBEDDED STRUCT vehicle to truck and sedan and Printing
	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "white",
		},
		fourwheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: false,
	}
	fmt.Println(t)
	fmt.Println(s)
	fmt.Println(t.doors)
	fmt.Println(s.doors)
}

/* OUTPUT:
{{2 white} true}
{{4 black} false}
2
4
*/
