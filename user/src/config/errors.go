package config

// CustomError represents a structured error with an error code, message, and optional data.
type CustomError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Error implements the error interface.
func (e CustomError) Error() string {
	return e.Message
}

// NewCustomError creates a new CustomError instance.
func NewCustomError(code int, message string, data interface{}) error {
	return CustomError{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
