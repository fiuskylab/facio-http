package facio

// Client responsible to manage Requests and Responses
type Client struct {
	// URL that all requests will be made in
	BaseURL string
}

// NewClient returns new client
func NewClient(base string) Client {
	baseBytes := []byte(base)
	lenBase := len(baseBytes)

	// if base url is http://foo.bar/ it returns http://foo.bar
	if baseBytes[lenBase-1] == byte('/') {
		base = string(baseBytes[:lenBase-1])
	}

	return Client{
		BaseURL: base,
	}
}

// Send returns Request response
// With: Status Code, Body, Header, etc.
func (c Client) Send(req Request) (Response, ErrHandler) {
	return req.call(c)
}
