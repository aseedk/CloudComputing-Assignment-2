package controller

import (
	"cloud-computing/logging/config"
	"errors"
	"net/http"
	"time"
)

const ApiContextTimeout = 5 * time.Second

// ErrorResponse represents a structured error response
type ErrorResponse struct {
	Success   bool        `json:"success"`
	ErrorCode int         `json:"code"`
	ErrorMsg  string      `json:"message"`
	ErrorData interface{} `json:"data,omitempty"`
}

// ParseError extracts error details from the error object
func ParseError(err error) ErrorResponse {
	var customErr config.CustomError
	if errors.As(err, &customErr) {
		return ErrorResponse{
			Success:   false,
			ErrorCode: customErr.Code,
			ErrorMsg:  customErr.Message,
			ErrorData: customErr.Data,
		}
	}
	return ErrorResponse{
		ErrorCode: http.StatusInternalServerError,
		ErrorMsg:  err.Error(),
	}
}
