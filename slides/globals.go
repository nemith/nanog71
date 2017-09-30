package main

import "fmt"

const MaxLatency = 0.15

var FavoriteProtcol = "BGP"

const (
	Cisco = iota
	Juniper
	Arista
)

func main() {
	fmt.Printf("Max Latency: %f,  Favorite Protocol: %s\n", MaxLatency, FavoriteProtcol)
	fmt.Println(Cisco, Juniper, Arista)
}
