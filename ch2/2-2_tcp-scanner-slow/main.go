// Page 25
// Listing 2-2: Scanning 1024 ports of scanme.nmap.org
package main

import "fmt"

func main() {
	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		fmt.Println(address)
	}
}