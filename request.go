package facio

type Req struct {
	BaseURL string
	Method  string
	Headers map[string]string
}

// NewRequest receives method and url to create
// a base request to be made
func NewRequest(method, url string) (*Req, error) {
	var err error
	var r Req

	method, err = checkMethod(method)
	if err != nil {
		return &r, err
	}

	url, err = parseURL(url)

	if err != nil {
		return &r, err
	}

	r.BaseURL = url
	r.Method = method
	r.Headers = make(map[string]string)

	return &r, err
}

// AddHeader add a header into request
func (r *Req) AddHeader(name header, value string) *Req {
	r.Headers[name.getHeader()] = value
	return r
}

// AddCustomHeader add a custom header into request
func (r *Req) AddCustomHeader(name, value string) *Req {
	r.Headers[name] = value
	return r
}
