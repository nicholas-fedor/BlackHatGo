package main

import "fmt"

func main() {
	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		fmt.Println(address)
	}
}