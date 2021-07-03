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

	// Endpoint to fetch based on Client's BaseURL
	Endpoint string

	// Method type POST, GET, PATCH, PUT, etc
	// Not case sensitive
	Method string
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
