package dto

import "time"

// CreateUserReq represents the request body for creating a user
type CreateUserReq struct {
	FirstName   string    `json:"firstName" binding:"required"`
	LastName    string    `json:"lastName" binding:"required"`
	DateOfBirth time.Time `json:"dateOfBirth" binding:"required" time_format:"2006-01-02"`
	Gender      bool      `json:"gender"`
	IsActive    bool      `json:"isActive"`
	Country     string    `json:"country" binding:"required"`
	Email       string    `json:"email" binding:"required,email"`
}

// JoinOrganizationReq represents the request body for a user joining an organization
type JoinOrganizationReq struct {
	UserId         string `json:"userId" binding:"required"`
	OrganizationId string `json:"organizationId" binding:"required"`
}

// LeaveOrganizationReq represents the request body for a user leaving an organization
type LeaveOrganizationReq struct {
	UserId         string `json:"userId" binding:"required"`
	OrganizationId string `json:"organizationId" binding:"required"`
}

// GetUserResp represents the response body of a user
type GetUserResp struct {
	Id          string    `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Gender      string    `json:"gender"`
	IsActive    bool      `json:"isActive"`
	Country     string    `json:"country"`
	Email       string    `json:"email"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
