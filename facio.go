package facio

import "strings"

// Response struct
type Response struct{}

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
