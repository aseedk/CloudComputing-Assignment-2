package validation

import (
	"cloud-computing/organization/organization/src/restful/models/dao"
	"cloud-computing/organization/organization/src/restful/models/dto"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

// CreateOrganization validates create organization request.
func CreateOrganization(_ context.Context, _ dto.CreateOrganizationReq) error {
	// No validation required for now.
	return nil
}

// UpdateOrganization validates update organization request.
func UpdateOrganization(ctx context.Context, organizationId string) error {
	count, err := dao.CountOrganizations(ctx, bson.M{dao.FieldOrganizationId: organizationId})
	if err != nil {
		return err
	}
	if count < 1 {
		return generateCustomError(404, "Organization not found", nil)
	}
	return nil
}

// DeleteOrganization validates delete organization request.
func DeleteOrganization(ctx context.Context, organizationId string) error {
	count, err := dao.CountOrganizations(ctx, bson.M{dao.FieldOrganizationId: organizationId})
	if err != nil {
		return err
	}
	if count < 1 {
		return generateCustomError(404, "Organization not found", nil)
	}
	return nil
}

// GetOrganization validates get organization request.
func GetOrganization(ctx context.Context, organizationId string) error {
	count, err := dao.CountOrganizations(ctx, bson.M{dao.FieldOrganizationId: organizationId})
	if err != nil {
		return err
	}
	if count < 1 {
		return generateCustomError(404, "Organization not found", nil)
	}
	return nil
}

// QueryOrganizations validates query organization request.
func QueryOrganizations(_ context.Context, req *dto.QueryOrganizationReq) error {
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
