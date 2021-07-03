package facio

// Client responsible to manage Requests and Responses
type Client struct {
	// URL that all requests will be made in
	BaseURL   string
	Request   Request
	HeaderMap HeaderMap
}

// Create new client
func NewClient(base string) Client {
	return Client{
		BaseURL: base,
		Request: Request{
			Headers: make(HeaderResult),
		},
		HeaderMap: make(HeaderMap),
	}
}

// Return Request response
// With: Status Code, Body, Header, etc.
func (c Client) Send(req Request) (Response, ErrHandler) {
	return req.call()
}
