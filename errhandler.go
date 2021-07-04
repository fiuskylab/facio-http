package facio

import "fmt"

type errHandler struct {
	// To check is there's an error or not
	// Unexported to avoid user errors
	isNil bool

	// Error message
	Error string
}

const (
	msgInvalidHeaderValue = `The header %s received is invalid: %s`
	msgInvalidMethod      = `The method %s is invalid`
	msgInvalidURL         = `The url "%s" is invalid`
)

func newError(msgTemplate string, values ...interface{}) errHandler {
	return errHandler{
		isNil: false,
		Error: fmt.Sprintf(msgTemplate, values...),
	}
}

// Create a new nil error
func newNilError() errHandler {
	return errHandler{
		isNil: true,
	}
}

// Check if there's an error or not
func (e errHandler) IsNil() bool {
	return e.isNil
}
