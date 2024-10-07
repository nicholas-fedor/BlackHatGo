// Pages 30-31
// Listing 2-8: Port scanning with multiple channels.
package main

import (
	"fmt"
	"net"
	"sort"
)

// The worker function has been modified to accept two channels.
func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// If the port is closed, send a zero.
			results <- 0
			continue
		}
		conn.Close()
		// If the port is open, send the port.
		results <- p
	}
}

func main() {
	// Creates a channel for assigning ports
	ports := make(chan int, 100)
	// A separate channel for communicating results from the worker to the main thread.
	results := make(chan int)
	// Using a slice to store the results so they can be sorted later.
	var openPorts []int

	// Creates worker goroutine for each port.
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	// Separate goroutine that pipes ports for the worker goroutines to handle.
	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	// Result-gathering loop
	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}

	// Ensures channels are closed.
	close(ports)
	close(results)

	// Sorts the slice of open ports
	sort.Ints(openPorts)

	// Prints the results to the screen.
	for _, port := range openPorts {
		fmt.Printf("%d open\n", port)
	}
}
