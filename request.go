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

	// Headers sent to request
	// ["Header Name": ["Header Value 1", "Header Value 2"]]
	headers HeaderResult

	headerMap HeaderMap

	req http.Request
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
		headers:   make(HeaderResult),
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

	if !err.IsNil() {
		return err
	}

	return r.prepareHeaders()
}

// Add multiple headers following the structure o HeaderMap
func (r *request) AddHeaders(hr HeaderMap) ErrHandler {
	r.headerMap = hr

	return r.prepareHeaders()
}

// Build the headers that will be sent to Endpoint
func (r *request) prepareHeaders() ErrHandler {
	headers, err := r.headerMap.build()

	r.headers = headers

	return err
}

func (r request) call(client *client) (*response, ErrHandler) {
	return &response{}, NewNilError()
}
