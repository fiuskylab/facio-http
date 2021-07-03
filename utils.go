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
