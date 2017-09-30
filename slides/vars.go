package main

import (
	"fmt"
)

// START OMIT

func main() {
	var a int // Must define type
	b := 30   // type infered (is int)

	var x string
	y := "MPLS Rulez"

	fmt.Printf("a: %d  \nb: %d\n", a, b)
	fmt.Printf("x: %q  \ny: %q\n", x, y)
	fmt.Printf("Type of b: %T\n", b)
	fmt.Printf("Type of y: %T\n", y)
}

// END OMIT
