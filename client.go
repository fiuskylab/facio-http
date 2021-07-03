package facio

// Client responsible to manage Requests and Responses
type Client struct {
	// URL that all requests will be made in
	BaseURL string
}

// Create new client
func NewClient(base string) Client {
	return Client{
		BaseURL: base,
	}
}

// Return Request response
// With: Status Code, Body, Header, etc.
func (c Client) Send(req Request) (Response, ErrHandler) {
	return req.call()
}
