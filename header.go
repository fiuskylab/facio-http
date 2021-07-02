package facio

import "strings"

// Type to store a key with an array of strings
// HeaderMap{
// 		"Authorization": { "Bearer EEEYY.TOKEN.AZZ" },
// 		"Accept-Encoding": { "gzip", "deflate" },
// }

//
const (
	Authorization = iota
)

// Headers written into the Client
type HeaderMap map[string][]string

// Headers "compiled"
type HeaderResult map[string]string

// Building Headers
func (hm HeaderMap) build() (HeaderResult, ErrHandler) {
	hr := make(HeaderResult, len(hm))

	for key, val := range hm {
		if len(val) == 0 {
			return HeaderResult{}, NewError(msgInvalidHeaderValue, key, val)
		}
		hr[key] = strings.Join(val, ", ")
	}

	return hr, NewNilError()
}

// Add header value to map
func (hm HeaderMap) add(header string, values ...string) (HeaderMap, ErrHandler) {
	if len(values) == 0 {
		return hm, NewError(msgInvalidHeaderValue, header, values)
	}

	hm[header] = values

	return hm, NewNilError()
}
