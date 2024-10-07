// Page 47
// Listing 3-2: Using the PostForm() function instead of Post().
package main

import (
	"log"
	"net/http"
	"net/url"
)

func main() {
	r1, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatalln("Get error:", err)
	}
	// Read response body. Not shown.
	defer r1.Body.Close()

	r2, err := http.Head("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatalln("Head error:", err)
	}
	// Read response body. Not shown.
	defer r2.Body.Close()

	form := url.Values{}
	form.Add("foo", "bar")

	r3, err := http.PostForm("https://www.google.com/robots.txt", form)
	if err != nil {
		log.Fatalln("Post error:", err)
	}
	// Read response body. Not shown.
	defer r3.Body.Close()
}
