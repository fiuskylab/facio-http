package facio

type ErrHandler struct {
	// To check is there's an error or not
	// Unexported to avoid user errors
	isNil bool
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
