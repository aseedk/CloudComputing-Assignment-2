package service

import (
	"cloud-computing/logging/restful/models/dao"
	"cloud-computing/logging/restful/models/dto"
	"context"
)

// CreateLog handles business logic for creating a log entry
func CreateLog(ctx context.Context, req dto.CreateLogReq) (*dto.CommonResponse, error) {
	logEntry := dao.Log{
		UserId:         req.UserId,
		OrganizationId: req.OrganizationId,
		Method:         req.Method,
		URL:            req.URL,
		RequestBody:    req.RequestBody,
		RequestQuery:   req.RequestQuery,
		ResponseBody:   req.ResponseBody,
	}

	err := dao.CreateLog(ctx, logEntry)
	if err != nil {
		return nil, err
	}

	return &dto.CommonResponse{Success: true, Code: 0, Message: "Log created successfully"}, nil
}

// QueryLogs retrieves multiple logs with pagination and filters
func QueryLogs(ctx context.Context, req dto.QueryLogsReq) (*dto.CommonResponse, error) {
	logsDao, total, err := dao.QueryLogs(ctx, req.UserId, req.OrganizationId, req.DateFrom, req.DateTo,
		*req.Page, *req.Limit)
	if err != nil {
		return nil, err
	}

	var logsDto = make([]dto.GetLogResp, len(logsDao))

	for i, log := range logsDao {
		logsDto[i] = dto.GetLogResp{
			Id:             log.Id.Hex(),
			UserId:         log.UserId,
			OrganizationId: log.OrganizationId,
			RequestBody:    log.RequestBody,
			RequestQuery:   log.RequestQuery,
			ResponseBody:   log.ResponseBody,
			Timestamp:      log.Timestamp,
		}
	}

	// Create the response object
	resp := dto.QueryLogsResp{
		Total: int(total),
		Page:  *req.Page,
		Size:  len(logsDao),
		Data:  logsDto,
	}
	return &dto.CommonResponse{Success: true, Code: 0, Message: "Success", Data: resp}, nil
}
