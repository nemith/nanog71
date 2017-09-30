package main

import (
	"fmt"
	"time"
)

// START OMIT

func main() {
	// C Style
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Loop over slice or map
	protocols := []string{"BGP", "OSPF", "RIP"}
	for _, p := range protocols {
		fmt.Println(p)
	}

	// Loop forever!
	for {
		fmt.Println("LEEEEETTTS GO!")
		time.Sleep(1 * time.Second)
	}
}

// END OMIT
