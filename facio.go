package facio

import "strings"

// Request struct
type Client struct {
	// URL that all requests will be made in
	BaseURL string
	Request Request
}

type Request struct {
	// Headers sent to request
	// ["Header Name": ["Header Value 1", "Header Value 2"]]
	Headers map[string]string
}

// Response struct
type Response struct{}

// Create new client
func NewClient(base string) Client {
	return Client{
		BaseURL: base,
		Request: Request{
			Headers: make(map[string]string),
		},
	}
}

// Add a single header with multiple values
func (c Client) AddHeader(header string, values ...string) (Request, ErrHandler) {
	if len(values) == 0 {
		return c.Request, NewError(msgInvalidHeaderValue, header, values)
	}

	value := strings.Join(values, ", ")
	c.Request.Headers[header] = value

	return c.Request, NewNilError()
}

// Type to store a key with an array of strings
// HeaderMap{
// 		"Authorization": { "Bearer EEEYY.TOKEN.AZZ" },
// 		"Accept-Encoding": { "gzip", "deflate" },
// }
type HeaderMap map[string][]string

// Add multiple headers following the structure o HeaderMap
func (c Client) AddHeaders(hr HeaderMap) (Request, ErrHandler) {
	for headerName, headerValues := range hr {
		if len(headerValues) == 0 {
			return c.Request, NewError(msgInvalidHeaderValue, headerName, headerValues)
		}
		value := strings.Join(headerValues, ", ")
		c.Request.Headers[headerName] = value
	}

	return c.Request, NewNilError()
}
