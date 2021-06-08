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
	Headers map[string]string
}

// Response struct
type Response struct{}

// Create new client
func NewClient(base string) Client {
	return Client{
		BaseURL: base,
		Request: Request{
			Headers: make(map[string]string),
		},
	}
}

func (c Client) AddHeader(header string, values ...string) (Request, ErrHandler) {
	if len(values) == 0 {
		return c.Request, NewError(msgInvalidValue, header, values)
	}

	value := strings.Join(values, ", ")
	c.Request.Headers[header] = value

	return c.Request, NewNilError()
}
