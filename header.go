package facio

import (
	"net/http"
	"strings"
)

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
func (hm HeaderMap) build() (HeaderResult, errHandler) {
	hr := make(HeaderResult, len(hm))

	for key, val := range hm {
		if len(val) == 0 {
			return HeaderResult{}, newError(msgInvalidHeaderValue, key.getHeader(), val)
		}
		hr[key.getHeader()] = strings.Join(val, ",")
	}

	return hr, newNilError()
}

// Add header value to map
func (hm HeaderMap) add(h header, values ...string) (HeaderMap, errHandler) {
	if len(values) == 0 {
		return hm, newError(msgInvalidHeaderValue, h.getHeader(), values)
	}

	hm[h] = values

	return hm, newNilError()
}

// toHTTPHeader return the map in http.Header type format
func (hm HeaderMap) toHTTPHeader() http.Header {
	var httpHeader http.Header

	for key, val := range hm {
		httpHeader[key.getHeader()] = val
	}

	return httpHeader
}
