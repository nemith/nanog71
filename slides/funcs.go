package main

import (
	"fmt"
)

// START OMIT
func add(a int, b int) int {
	return a + b
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("Juniper", "Cisco")
	x := add(40, 2)
	fmt.Println(a, b, x)
}

// END OMIT
