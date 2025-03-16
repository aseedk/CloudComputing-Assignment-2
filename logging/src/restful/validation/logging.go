package validation

import (
	"cloud-computing/logging/restful/models/dto"
	"context"
)

// CreateLog validates the create log request.
func CreateLog(_ context.Context, _ dto.CreateLogReq) error {
	// No validation required for now.
	return nil
}

// QueryLogs validates the query logs request.
func QueryLogs(_ context.Context, req *dto.QueryLogsReq) error {
	if req.Page == nil {
		defaultPage := 0
		req.Page = &defaultPage
	}
	if req.Limit == nil {
		defaultLimit := 10
		req.Limit = &defaultLimit
	}
	return nil
}
