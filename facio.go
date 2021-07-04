package facio

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
func NewDefaultFacio(baseURL string) *Facio {
	client := NewClient(baseURL)
	return &Facio{
		client: client,
	}
}

// SetBaseURL update the URL that Facio makes calls
func (f *Facio) SetBaseURL(baseURL string) (*Facio, ErrHandler) {
	err := f.client.setBaseURL(baseURL)

	return f, err
}

// Request call into certain endpoint based on baseURL defined on client
func (f *Facio) Request(method, endpoint string, hm HeaderMap) (*response, ErrHandler) {
	req := &request{
		client:    f.client,
		method:    method,
		endpoint:  endpoint,
		headerMap: hm,
	}

	return req.call()
}
