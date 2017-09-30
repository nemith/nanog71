package main

import (
	"fmt"
	"time"
)

// START OMIT
func dostuff() {
	fmt.Println("Start dostuff()")
	time.Sleep(1 * time.Second)
	fmt.Println("End dostuff()")
}

func main() {
	go dostuff() // Don't block, run in background
	fmt.Println("From main()")
	time.Sleep(2 * time.Second)
}

// END OMIT
