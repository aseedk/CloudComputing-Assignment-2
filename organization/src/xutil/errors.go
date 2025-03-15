package xutil

import "cloud-computing/organization/config"

// GenerateCustomError creates a new custom error instance
func GenerateCustomError(code int, message string, data interface{}) error {
	return config.CustomError{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
