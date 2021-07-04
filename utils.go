package facio

import "strings"

// Check if defined method exists
func checkMethod(method string) (string, ErrHandler) {
	methodUpper := strings.ToUpper(method)

	// Not the prettiest solution
	if methodUpper == "GET" ||
		methodUpper == "HEAD" ||
		methodUpper == "POST" ||
		methodUpper == "PUT" ||
		methodUpper == "PATCH" ||
		methodUpper == "DELETE" ||
		methodUpper == "CONNECT" ||
		methodUpper == "OPTIONS" ||
		methodUpper == "TRACE" {
		return methodUpper, NewNilError()
	}

	return "", NewError(msgInvalidMethod, method)
}

func parseURL(baseURL string) (string, ErrHandler) {
	baseBytes := []byte(baseURL)
	lenBase := len(baseBytes)

	if lenBase == 0 {
		return "", NewError(msgInvalidURL, baseURL)
	}

	// if base url is http://foo.bar/ it returns http://foo.bar
	if baseBytes[lenBase-1] == byte('/') {
		baseURL = string(baseBytes[:lenBase-1])
	}

	return baseURL, NewNilError()
}
