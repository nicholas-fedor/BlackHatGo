// Page 53
// Listing 3-7: Shodan Client definition
package shodan

const BaseURL = "https://api.shodan.io"

type Client struct {
	apiKey string
}

func New(apiKey string) *Client {
	return &Client{apiKey: apiKey}
}
