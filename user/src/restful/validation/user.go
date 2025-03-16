package validation

import (
	"cloud-computing/users/restful/models/dao"
	"cloud-computing/users/restful/models/dto"
	"cloud-computing/users/xutil"
	"context"
	"errors"
)

// CreateUser validates create user request.
func CreateUser(_ context.Context, _ dto.CreateUserReq) error {
	// No validation required for now.
	return nil
}

// JoinOrganization validates join organization request.
func JoinOrganization(ctx context.Context, req dto.JoinOrganizationReq) error {
	if !xutil.CheckOrganizationExists(ctx, req.OrganizationId) {
		return errors.New("organization not found")
	}
	return nil
}

// LeaveOrganization validates leave organization request.
func LeaveOrganization(ctx context.Context, req dto.LeaveOrganizationReq) error {
	if !xutil.CheckOrganizationExists(ctx, req.OrganizationId) {
		return errors.New("organization not found")
	}

	if exist, err := dao.ExistUserOrganization(ctx, req.UserId, req.OrganizationId); err != nil {
		return errors.New("user not in organization")
	} else {
		if !exist {
			return errors.New("user not in organization")
		} else {
			return nil
		}
	}
}
