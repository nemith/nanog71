package main

import (
	"encoding/json"
	"log"
	"nanog71/pkg/bmp"
	"os"
)

func handler(msg *bmp.Message) {
	if err := json.NewEncoder(os.Stdout).Encode(msg); err != nil {
		log.Printf("Failed to encode json: %v", err)
	}
}

func main() {
	err := bmp.ListenAndServe(":11019", bmp.HandlerFunc(handler))
	if err != nil {
		log.Fatal(err)
	}
}
