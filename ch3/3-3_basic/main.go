package main

import (
	"log"
	"net/http"
)

func main() {
	req, err := http.NewRequest("DELETE", "https://www.google.com/robots.txt", nil)
	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	// Read response body and close.
	defer resp.Body.Close()
}
