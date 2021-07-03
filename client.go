package facio

import "strings"

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

// Add a single header with multiple values
func (c *Client) AddHeader(h header, values ...string) ErrHandler {
	if len(values) == 0 {
		return NewError(msgInvalidHeaderValue, h.getHeader(), values)
	}

	var err ErrHandler

	c.HeaderMap, err = c.HeaderMap.add(h, values...)

	if !err.IsNil() {
		return err
	}

	return NewNilError()
}

// Add multiple headers following the structure o HeaderMap
func (c *Client) AddHeaders(hr HeaderMap) ErrHandler {
	c.HeaderMap = hr

	return NewNilError()
}

// Return Request with built Headers, etc.
func (c Client) NewRequest(method, endpoint string) (Request, ErrHandler) {
	headers, err := c.HeaderMap.build()

	if !err.IsNil() {
		return Request{}, err
	}

	req := Request{
		Headers:  headers,
		Endpoint: endpoint,
		Method:   strings.ToUpper(method),
	}

	c.Request = req

	return c.Request, NewNilError()
}
