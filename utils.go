package facio

import "strings"

// Check if defined method exists
func checkMethod(method string) (string, errHandler) {
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
		return methodUpper, newNilError()
	}

	return "", newError(msgInvalidMethod, method)
}

// parseURL removes the "/" from the end of the string
func parseURL(baseURL string) (string, errHandler) {
	baseBytes := []byte(baseURL)
	lenBase := len(baseBytes)

	if lenBase == 0 {
		return "", newError(msgInvalidURL, baseURL)
	}

	// if base url is http://foo.bar/ it returns http://foo.bar
	if baseBytes[lenBase-1] == byte('/') {
		baseURL = string(baseBytes[:lenBase-1])
	}

	return baseURL, newNilError()
}

// parseEndpoint add "/" at the beginning
func parseEndpoint(endpoint string) string {
	epBytes := []byte(endpoint)
	breaker := []byte("/")
	if epBytes[0] == breaker[0] {
		return endpoint
	}

	return string(append(breaker, epBytes...))
}
