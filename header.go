package facio

import "strings"

// Type to store a key with an array of strings
// HeaderMap{
// 		"Authorization": { "Bearer EEEYY.TOKEN.AZZ" },
// 		"Accept-Encoding": { "gzip", "deflate" },
// }

// Headers written into the Client
type HeaderMap map[header][]string

// "Compiled" headers
type HeaderResult map[string]string

// Building Headers
func (hm HeaderMap) build() (HeaderResult, ErrHandler) {
	hr := make(HeaderResult, len(hm))

	for key, val := range hm {
		if len(val) == 0 {
			return HeaderResult{}, NewError(msgInvalidHeaderValue, key.getHeader(), val)
		}
		hr[key.getHeader()] = strings.Join(val, ", ")
	}

	return hr, NewNilError()
}

// Add header value to map
func (hm HeaderMap) add(h header, values ...string) (HeaderMap, ErrHandler) {
	if len(values) == 0 {
		return hm, NewError(msgInvalidHeaderValue, h.getHeader(), values)
	}

	hm[h] = values

	return hm, NewNilError()
}
