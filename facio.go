package facio

import "strings"

// Request struct
type Client struct {
	// URL that all requests will be made in
	BaseURL string
	Request Request
}

type Request struct {
	// Headers sent to request
	// ["Header Name": ["Header Value 1", "Header Value 2"]]
	headers map[string]string
}

// Response struct
type Response struct{}

// Create new client
func NewClient(base string) Client {
	return Client{
		BaseURL: base,
		Request: Request{},
	}
}

func (r Request) AddHeader(header string, values ...string) Request {
	value := strings.Join(values, ", ")
	r.headers[header] = value

	return r
}
