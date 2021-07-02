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

// Add a single header with multiple values
func (c *Client) AddHeader(header string, values ...string) ErrHandler {
	if len(values) == 0 {
		return NewError(msgInvalidHeaderValue, header, values)
	}

	var err ErrHandler

	c.HeaderMap, err = c.HeaderMap.add(header, values...)

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

	var err ErrHandler

	c.Request.Headers, err = hr.build()

	if !err.IsNil() {
		return c.Request, err
	}

	return c.Request, NewNilError()
}
