package facio

import "fmt"

// Facio is the main struct to execute all actions and config the client
type Facio struct {
	// Client stores the main config
	client *client
}

// NewFacio returns Facio with given Client
func NewFacio(client *client) *Facio {
	return &Facio{client: client}
}

// NewDefaultFacio returns a Facio
func NewDefaultFacio(baseURL string) (*Facio, ErrHandler) {
	client, err := NewClient(baseURL)
	return &Facio{
		client: client,
	}, err
}

// SetBaseURL update the URL that Facio makes calls
func (f *Facio) SetBaseURL(baseURL string) (*Facio, ErrHandler) {
	err := f.client.setBaseURL(baseURL)

	return f, err
}

// Request call into certain endpoint based on baseURL defined on client
func (f *Facio) Request(method, endpoint string, hm HeaderMap) (*request, ErrHandler) {
	var err ErrHandler

	method, err = checkMethod(method)

	if !err.isNil {
		return &request{}, err
	}

	req := &request{
		client:    f.client,
		method:    method,
		endpoint:  parseEndpoint(endpoint),
		headerMap: hm,
	}

	req.url = fmt.Sprintf("%s%s", req.client.baseURL, req.endpoint)

	return req, NewNilError()
}

// Do the given Request
func (f *Facio) Do(req *request) (*response, ErrHandler) {
	return req.call()
}
