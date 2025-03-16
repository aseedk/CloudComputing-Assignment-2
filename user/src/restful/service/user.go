package service

import (
	"cloud-computing/users/restful/models/dao"
	"cloud-computing/users/restful/models/dto"
	"context"
	"fmt"
)

// CreateUser handles the business logic for creating a user
func CreateUser(ctx context.Context, req dto.CreateUserReq) (*dto.CommonResponse, error) {
	userId, err := dao.GenerateUserId(ctx)
	if err != nil {
		return nil, err
	}

	user := dao.User{
		UserId:      userId,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		DateOfBirth: req.DateOfBirth,
		Gender:      req.Gender,
		Country:     req.Country,
		IsActive:    req.IsActive,
	}

	err = dao.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &dto.CommonResponse{
		Success: true,
		Code:    0,
		Message: "User created successfully",
		Data:    userId,
	}, nil
}

// JoinOrganization handles the business logic for a user joining an organization
func JoinOrganization(ctx context.Context, req dto.JoinOrganizationReq) (*dto.CommonResponse, error) {
	userOrganization := dao.UserOrganization{
		UserId:         req.UserId,
		OrganizationId: req.OrganizationId,
	}

	err := dao.JoinOrganization(ctx, userOrganization.UserId, userOrganization.OrganizationId)
	if err != nil {
		return nil, err
	}

	return &dto.CommonResponse{
		Success: true,
		Code:    0,
		Message: "User joined organization successfully",
	}, nil
}

// LeaveOrganization handles the business logic for a user leaving an organization
func LeaveOrganization(ctx context.Context, req dto.LeaveOrganizationReq) (*dto.CommonResponse, error) {
	userOrganization := dao.UserOrganization{
		UserId:         req.UserId,
		OrganizationId: req.OrganizationId,
	}

	err := dao.LeaveOrganization(ctx, userOrganization.UserId, userOrganization.OrganizationId)
	if err != nil {
		return nil, err
	}

	return &dto.CommonResponse{
		Success: true,
		Code:    0,
		Message: fmt.Sprintf("user %s left organization %s successfully", req.UserId, req.OrganizationId),
	}, nil
}
