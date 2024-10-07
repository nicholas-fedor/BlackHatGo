// Page 48
// Listing 3-4: Sending a PUT request
package main

import (
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	form := url.Values{}
	form.Add("foo", "bar")
	var client http.Client
	req, err := http.NewRequest(
		"PUT",
		"https://www.google.com/robots.txt",
		strings.NewReader(form.Encode()),
	)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	// Read response body and close.
	defer resp.Body.Close()
}