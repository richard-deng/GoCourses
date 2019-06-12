package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s address", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	cname, err := net.LookupCNAME(name)
	if err != nil {
		fmt.Println("LookupCNAME error: ", err.Error())
		os.Exit(1)
	}
	fmt.Println("The host canonical name is: ", cname)
	os.Exit(0)
}
