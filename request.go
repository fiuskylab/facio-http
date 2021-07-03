package facio

// Store every info about a request to be made
type Request struct {
	// Endpoint to fetch based on Client's BaseURL
	endpoint string

	// Method type POST, GET, PATCH, PUT, etc
	// Not case sensitive
	method string

	// Headers sent to request
	// ["Header Name": ["Header Value 1", "Header Value 2"]]
	headers HeaderResult

	headerMap HeaderMap
}

// Return Request with method and endpoints
func NewRequest(method, endpoint string) (Request, ErrHandler) {
	upperMethod, err := checkMethod(method)

	if !err.IsNil() {
		return Request{}, err
	}

	return Request{
		method:    upperMethod,
		endpoint:  endpoint,
		headers:   make(HeaderResult),
		headerMap: make(HeaderMap),
	}, NewNilError()
}

// Add a single header with multiple values
func (r *Request) AddHeader(h header, values ...string) ErrHandler {
	if len(values) == 0 {
		return NewError(msgInvalidHeaderValue, h.getHeader(), values)
	}

	var err ErrHandler

	r.headerMap, err = r.headerMap.add(h, values...)

	if !err.IsNil() {
		return err
	}

	return NewNilError()
}

// Add multiple headers following the structure o HeaderMap
func (r *Request) AddHeaders(hr HeaderMap) ErrHandler {
	r.headerMap = hr

	return NewNilError()
}
