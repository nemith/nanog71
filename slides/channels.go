package main

import (
	"fmt"
	"time"
)

// START OMIT

func main() {
	data := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		data <- 1
		fmt.Println("Done sending data")
	}()

	fmt.Println("Waiting for data")
	val := <-data
	fmt.Printf("got: %d\n", val)
}

// END OMIT
