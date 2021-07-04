package facio

import "fmt"

type ErrHandler struct {
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

func NewError(msgTemplate string, values ...interface{}) ErrHandler {
	return ErrHandler{
		isNil: false,
		Error: fmt.Sprintf(msgTemplate, values...),
	}
}

// Create a new nil error
func NewNilError() ErrHandler {
	return ErrHandler{
		isNil: true,
	}
}

// Check if there's an error or not
func (e ErrHandler) IsNil() bool {
	return e.isNil
}
