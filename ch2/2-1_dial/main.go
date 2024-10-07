// Pages 24-25
// Listing 2-1: A basic port scanner that scans only one port
package main

import (
	"fmt"
	"net"
)

func main() {
	_, err := net.Dial("tcp", "scanme.nmap.org:80")
	if err == nil {
		fmt.Println("Connection Successful")
	}
}
