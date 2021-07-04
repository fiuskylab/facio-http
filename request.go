package facio

import "net/http"

// Store every info about a request to be made
type request struct {
	client *client
	// Endpoint to fetch based on Client's BaseURL
	endpoint string

	// Method type POST, GET, PATCH, PUT, etc
	// Not case sensitive
	method string

	// URL
	url string

	// Headers sent to request
	// ["Header Name": ["Header Value 1", "Header Value 2"]]
	headerMap HeaderMap

	req *http.Request
}

// Return Request with method and endpoints
func NewRequest(method, endpoint string) (*request, ErrHandler) {
	upperMethod, err := checkMethod(method)

	if !err.IsNil() {
		return &request{}, err
	}

	return &request{
		method:    upperMethod,
		endpoint:  endpoint,
		headerMap: make(HeaderMap),
	}, NewNilError()
}

// Add a single header with multiple values
func (r *request) AddHeader(h header, values ...string) ErrHandler {
	if len(values) == 0 {
		return NewError(msgInvalidHeaderValue, h.getHeader(), values)
	}

	var err ErrHandler

	r.headerMap, err = r.headerMap.add(h, values...)

	return err
}

// Add multiple headers following the structure o HeaderMap
func (r *request) AddHeaders(hr HeaderMap) {
	r.headerMap = hr
}

// Call makes the request and return a response
func (r *request) call() (*response, ErrHandler) {
	req, err := http.NewRequest(r.method, r.url, nil)

	if err != nil {
		errH := ErrHandler{
			Error: err.Error(),
		}
		return &response{}, errH
	}

	r.req = req
	r.req.Header = r.headerMap.toHTTPHeader()

	return &response{}, NewNilError()
}
