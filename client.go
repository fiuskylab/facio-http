package facio

// Client responsible to manage Requests and Responses
type Client struct {
	// URL that all requests will be made in
	BaseURL string
	Request Request
}

// Create new client
func NewClient(base string) Client {
	return Client{
		BaseURL: base,
		Request: Request{
			Headers: make(map[string]string),
		},
	}
}
