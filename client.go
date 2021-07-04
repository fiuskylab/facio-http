package facio

// Client responsible to manage Requests and Responses
type client struct {
	// URL that all requests will be made in
	baseURL string
}

// NewClient returns new client
func NewClient(base string) *client {
	baseBytes := []byte(base)
	lenBase := len(baseBytes)

	// if base url is http://foo.bar/ it returns http://foo.bar
	if baseBytes[lenBase-1] == byte('/') {
		base = string(baseBytes[:lenBase-1])
	}

	return &client{
		baseURL: base,
	}
}

// Send returns Request response
// With: Status Code, Body, Header, etc.
func (c *client) Send(req *request) (*response, ErrHandler) {
	return req.call(c)
}
