package dto

import "time"

// CreateOrganizationReq represents the request body for creating an organization
type CreateOrganizationReq struct {
	Name string `json:"name" binding:"required"`
}

// UpdateOrganizationReq represents the request body for updating an organization
type UpdateOrganizationReq struct {
	OrganizationId string `uri:"organizationId" binding:"required"`
	Name           string `json:"name" binding:"required"`
}

// DeleteOrganizationReq represents the request body for deleting an organization
type DeleteOrganizationReq struct {
	OrganizationId string `uri:"organizationId" binding:"required"`
}

// GetOrganizationReq represents the request for fetching a single organization
type GetOrganizationReq struct {
	OrganizationId string `uri:"organizationId" binding:"required"`
}

// QueryOrganizationReq represents the request for querying multiple organizations with pagination and filters
type QueryOrganizationReq struct {
	Name      *string    `form:"name"`
	CreatedAt *time.Time `form:"createdAt"`
	Page      *int       `form:"page" binding:"omitempty,min=1"`
	Limit     *int       `form:"limit" binding:"omitempty,min=1,max=100"`
	Skip      *int       `form:"skip" binding:"omitempty,min=0"`
	DateFrom  *time.Time `form:"dateFrom" time_format:"2006-01-02"`
	DateTo    *time.Time `form:"dateTo" time_format:"2006-01-02"`
	OrderBy   *string    `form:"orderBy"`
	SortBy    *string    `form:"sortBy"`
}

// GetOrganizationResp represents the complete response body of an organization
type GetOrganizationResp struct {
	Id             string    `json:"id"`
	OrganizationId string    `json:"organizationId"`
	Name           string    `json:"name"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

// QueryOrganizationResp represents the response for querying multiple organizations with pagination details
type QueryOrganizationResp struct {
	Total int                   `json:"total"`
	Page  int                   `json:"page"`
	Size  int                   `json:"size"`
	Data  []GetOrganizationResp `json:"data"`
}
