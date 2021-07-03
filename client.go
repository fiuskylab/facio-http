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

// Shortcut for NewRequest("POST", "/foo")
func (c Client) Post(endpoint string) (Request, ErrHandler) {
	return c.NewRequest("POST", endpoint)
}

// Shortcut for NewRequest("GET", "/foo")
func (c Client) Get(endpoint string) (Request, ErrHandler) {
	return c.NewRequest("GET", endpoint)
}

// Shortcut for NewRequest("PUT", "/foo")
func (c Client) Put(endpoint string) (Request, ErrHandler) {
	return c.NewRequest("PUT", endpoint)
}

// Shortcut for NewRequest("PATCH", "/foo")
func (c Client) Patch(endpoint string) (Request, ErrHandler) {
	return c.NewRequest("PATCH", endpoint)
}

// Shortcut for NewRequest("DELETE", "/foo")
func (c Client) Delete(endpoint string) (Request, ErrHandler) {
	return c.NewRequest("DELETE", endpoint)
}

// Shortcut for NewRequest("HEAD", "/foo")
func (c Client) Head(endpoint string) (Request, ErrHandler) {
	return c.NewRequest("HEAD", endpoint)
}

// Shortcut for NewRequest("CONNECT", "/foo")
func (c Client) Connect(endpoint string) (Request, ErrHandler) {
	return c.NewRequest("CONNECT", endpoint)
}

// Shortcut for NewRequest("OPTIONS", "/foo")
func (c Client) Options(endpoint string) (Request, ErrHandler) {
	return c.NewRequest("OPTIONS", endpoint)
}

// Shortcut for NewRequest("TRACE", "/foo")
func (c Client) Trace(endpoint string) (Request, ErrHandler) {
	return c.NewRequest("TRACE", endpoint)
}
