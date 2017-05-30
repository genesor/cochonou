package http

// Error represents an error encoded in JSON given as
// response for HTTP requests when an error occurs.
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details"`
}

// NewJSONError creates a JSON serializable error.
func NewJSONError(code, message string, err error) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Details: err.Error(),
	}
}
