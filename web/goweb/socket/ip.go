package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage:%s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	addr := net.ParseIP(os.Args[1])
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Printf("The address is %v\n", addr)
	}

}
