package main

import "fmt"

type Router struct {
	Hostname  string
	IPAddress string
	ASN       int
}

// START OMIT
func (r *Router) String() string {
	return fmt.Sprintf("%s (AS%d) - %s", r.Hostname, r.ASN, r.IPAddress)
}

func (r *Router) ISPrivateASN() bool {
	return r.ASN >= 65000 && r.ASN <= 65535
}

func main() {
	r := Router{
		Hostname:  "edge01.sjc1",
		IPAddress: "2001:db8::/32",
		ASN:       65000,
	}
	fmt.Println(r.String())
	fmt.Println(r.ISPrivateASN())
}

// END OMIT
