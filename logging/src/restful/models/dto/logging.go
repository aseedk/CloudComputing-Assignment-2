package dto

import "time"

// CreateLogReq represents the request body for creating a log entry
type CreateLogReq struct {
	URL            string `json:"url" binding:"required"`
	Method         string `json:"method" binding:"required"`
	UserId         string `json:"userId" binding:"required"`
	OrganizationId string `json:"organizationId" binding:"required"`
	RequestBody    string `json:"requestBody"`
	RequestQuery   string `json:"requestQuery"`
	ResponseBody   string `json:"responseBody"`
}

// QueryLogsReq represents the request for querying logs with filters
type QueryLogsReq struct {
	UserId         *string    `form:"userId"`
	OrganizationId *string    `form:"organizationId"`
	DateFrom       *time.Time `form:"dateFrom" time_format:"2006-01-02"`
	DateTo         *time.Time `form:"dateTo" time_format:"2006-01-02"`
	Page           *int       `form:"page" binding:"omitempty"`
	Limit          *int       `form:"limit" binding:"omitempty,max=100"`
	Skip           *int       `form:"skip" binding:"omitempty"`
	OrderBy        *string    `form:"orderBy"`
	SortBy         *string    `form:"sortBy"`
}

// GetLogResp represents the response body for a single log entry
type GetLogResp struct {
	Id             string    `json:"_id"`
	UserId         string    `json:"userId"`
	OrganizationId string    `json:"organizationId"`
	RequestBody    string    `json:"requestBody"`
	RequestQuery   string    `json:"requestQuery"`
	ResponseBody   string    `json:"responseBody"`
	Timestamp      time.Time `json:"timestamp"`
}

// QueryLogsResp represents the response for querying multiple Log with pagination details
type QueryLogsResp struct {
	Total int          `json:"total"`
	Page  int          `json:"page"`
	Size  int          `json:"size"`
	Data  []GetLogResp `json:"data"`
}
