package facio

// Client responsible to manage Requests and Responses
type client struct {
	// URL that all requests will be made in
	baseURL string
}

// NewClient returns new client
func NewClient(base string) (*client, ErrHandler) {
	var err ErrHandler

	c := &client{}

	c.baseURL, err = parseURL(base)

	return c, err
}

// setBaseURL update the baseURL for Facio calls
func (c *client) setBaseURL(baseURL string) ErrHandler {
	var err ErrHandler

	c.baseURL, err = parseURL(baseURL)

	return err
}

// Send returns Request response
// With: Status Code, Body, Header, etc.
func (c *client) Send(req *request) (*response, ErrHandler) {
	return req.call()
}
