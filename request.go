package facio

type Req struct {
	baseURL string
	method  string
	headers map[string]string
	formURL map[string][]string
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

	r = Req{
		baseURL: url,
		method:  method,
		headers: make(map[string]string),
		formURL: make(map[string][]string),
	}

	return &r, err
}

// AddHeader add a header into request
func (r *Req) AddHeader(name header, value string) *Req {
	r.headers[name.getHeader()] = value
	return r
}

// AddCustomHeader add a custom header into request
func (r *Req) AddCustomHeader(name, value string) *Req {
	r.headers[name] = value
	return r
}

// SetURLForm receives a map with form-value variable
func (r *Req) SetURLForm(vals map[string][]string) *Req {
	r.formURL = vals

	return r
}
