package main

import "fmt"

// START OMIT
type Router struct {
	Hostname  string
	IPAddress string
	ASN       int
}

func main() {
	r := Router{
		Hostname:  "edge01.sjc1",
		IPAddress: "2001:db8::/32",
	}
	r.ASN = 32934

	fmt.Printf("%s:\n\tIP: %s\n\tASN: %d", r.Hostname, r.IPAddress, r.ASN)
}

// END OMIT
