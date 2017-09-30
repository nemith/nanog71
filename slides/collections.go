package main

import (
	"fmt"
)

// START OMIT

func main() {
	ports := []int{22, 179, 80, 443}
	fmt.Println(ports[2])
	fmt.Println(ports[1:3])

	mapNames := map[int]string{
		22:  "ssh",
		179: "bgp",
	}
	fmt.Println(mapNames[22])
	fmt.Println(mapNames[179])
}

// END OMIT
