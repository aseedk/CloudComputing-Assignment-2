package controller

import (
	"cloud-computing/organization/organization/src/config"
	"errors"
	"net/http"
)

// ErrorResponse represents a structured error response
type ErrorResponse struct {
	ErrorCode int         `json:"errorCode"`
	ErrorMsg  string      `json:"errorMsg"`
	ErrorData interface{} `json:"errorData,omitempty"`
}

// ParseError extracts error details from the error object
func ParseError(err error) ErrorResponse {
	var customErr config.CustomError
	if errors.As(err, &customErr) {
		return ErrorResponse{
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
