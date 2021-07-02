package facio

import "strings"

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
			Headers: make(HeaderResult),
		},
	}
}

// Add a single header with multiple values
func (c Client) AddHeader(header string, values ...string) (Request, ErrHandler) {
	if len(values) == 0 {
		return c.Request, NewError(msgInvalidHeaderValue, header, values)
	}

	value := strings.Join(values, ", ")
	c.Request.Headers[header] = value

	return c.Request, NewNilError()
}

// Add multiple headers following the structure o HeaderMap
func (c Client) AddHeaders(hr HeaderMap) (Request, ErrHandler) {
	var err ErrHandler

	c.Request.Headers, err = hr.build()

	if !err.IsNil() {
		return c.Request, err
	}

	return c.Request, NewNilError()
}
