package service

import (
	"cloud-computing/organization/organization/src/restful/models/dao"
	"cloud-computing/organization/organization/src/restful/models/dto"
	"context"
)

// CreateOrganization handles business logic for creating an organization
func CreateOrganization(ctx context.Context, req dto.CreateOrganizationReq) (*dto.CommonResponse, error) {
	org := dao.Organization{
		Name: req.Name,
	}
	err := dao.CreateOrganization(ctx, org)
	if err != nil {
		return nil, err
	}
	return &dto.CommonResponse{Code: 0, Message: "Organization created successfully", Data: nil}, nil
}

// UpdateOrganization handles business logic for updating an organization
func UpdateOrganization(ctx context.Context, organizationId string, req dto.UpdateOrganizationReq) (*dto.CommonResponse, error) {
	err := dao.UpdateOrganization(ctx, organizationId, req.Name)
	if err != nil {
		return nil, err
	}
	return &dto.CommonResponse{Code: 0, Message: "Organization updated successfully", Data: nil}, nil
}

// DeleteOrganization handles business logic for deleting an organization
func DeleteOrganization(ctx context.Context, organizationId string) (*dto.CommonResponse, error) {
	err := dao.DeleteOrganization(ctx, organizationId)
	if err != nil {
		return nil, err
	}
	return &dto.CommonResponse{Code: 0, Message: "Organization deleted successfully", Data: nil}, nil
}

// GetOrganization retrieves an organization by its ID
func GetOrganization(ctx context.Context, organizationId string) (*dto.CommonResponse, error) {
	org, err := dao.GetOrganization(ctx, organizationId)
	if err != nil {
		return nil, err
	}
	return &dto.CommonResponse{Code: 0, Message: "Success", Data: org}, nil
}

// QueryOrganizations retrieves multiple organizations with pagination and filters
func QueryOrganizations(ctx context.Context, req dto.QueryOrganizationReq) (*dto.CommonResponse, error) {
	organizationsDao, total, err := dao.QueryOrganizations(ctx, req.OrganizationIds, req.Name, req.DateFrom, req.DateTo,
		dao.PageReq{
			Page:  req.Page,
			Limit: req.Limit,
			Skip:  req.Skip,
		})
	if err != nil {
		return nil, err
	}

	var organizationsDto = make([]dto.GetOrganizationResp, len(organizationsDao))

	for i, org := range organizationsDao {
		organizationsDto[i] = dto.GetOrganizationResp{
			Id:             org.Id.Hex(),
			OrganizationId: org.OrganizationId,
			Name:           org.Name,
			CreatedAt:      org.CreatedAt,
			UpdatedAt:      org.UpdatedAt,
		}
	}

	// Create the response object
	resp := dto.QueryOrganizationResp{
		Total: total,
		Page:  *req.Page,
		Size:  len(organizationsDao),
		Data:  organizationsDto,
	}
	return &dto.CommonResponse{Code: 0, Message: "Success", Data: resp}, nil
}
