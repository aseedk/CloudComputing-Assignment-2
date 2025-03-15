package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// Constants for field names
const (
	FieldId               = "_id"
	FieldOrganizationId   = "organizationId"
	FieldOrganizationName = "name"
	FieldCreatedAt        = "createdAt"
	FieldUpdatedAt        = "updatedAt"
)

// Organization represents the organization model
type Organization struct {
	Id             primitive.ObjectID `bson:"_id,omitempty"`
	OrganizationId string             `bson:"organizationId"`
	Name           string             `bson:"name"`
	CreatedAt      time.Time          `bson:"createdAt"`
	UpdatedAt      time.Time          `bson:"updatedAt"`
}

// PageReq represents pagination query parameters
type PageReq struct {
	Page  *int
	Limit *int
	Skip  *int
}

// CreateOrganization inserts a new organization into the database
func CreateOrganization(ctx context.Context, org Organization) error {
	org.Id = primitive.NewObjectID()
	org.CreatedAt = time.Now()
	org.UpdatedAt = time.Now()
	_, err := OrganizationCollection.InsertOne(ctx, org)
	return err
}

// UpdateOrganization updates an existing organization
func UpdateOrganization(ctx context.Context, orgId string, name string) error {
	filter := bson.M{FieldOrganizationId: orgId}
	update := bson.M{"$set": bson.M{FieldOrganizationName: name, FieldUpdatedAt: time.Now()}}
	_, err := OrganizationCollection.UpdateOne(ctx, filter, update)
	return err
}

// DeleteOrganization removes an organization from the database
func DeleteOrganization(ctx context.Context, orgId string) error {
	filter := bson.M{FieldOrganizationId: orgId}
	_, err := OrganizationCollection.DeleteOne(ctx, filter)
	return err
}

// GetOrganization fetches a single organization by ID
func GetOrganization(ctx context.Context, orgId string) (*Organization, error) {
	var org Organization
	err := OrganizationCollection.FindOne(ctx, bson.M{FieldOrganizationId: orgId}).Decode(&org)
	if err != nil {
		return nil, err
	}
	return &org, nil
}

// CountOrganizations returns the count of organizations matching the filter
func CountOrganizations(ctx context.Context, filter bson.M) (int, error) {
	total, err := OrganizationCollection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return int(total), nil
}

// QueryOrganizations retrieves multiple organizations with filtering and pagination
func QueryOrganizations(ctx context.Context, organizationIds []string, name *string, dateFrom *time.Time, dateTo *time.Time, pageReq PageReq) ([]Organization, int, error) {
	filter := bson.M{}
	if len(organizationIds) > 0 {
		filter[FieldOrganizationId] = bson.M{"$in": organizationIds}
	}
	if name != nil {
		filter[FieldOrganizationName] = *name
	}
	if dateFrom != nil && dateTo != nil {
		filter[FieldCreatedAt] = bson.M{"$gte": *dateFrom, "$lte": *dateTo}
	}

	total, err := CountOrganizations(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	queryOption := options.Find()
	if pageReq.Skip != nil {
		queryOption.SetSkip(int64(*pageReq.Skip))
	}
	if pageReq.Limit != nil {
		queryOption.SetLimit(int64(*pageReq.Limit))
	}

	cursor, err := OrganizationCollection.Find(ctx, filter, queryOption)
	if err != nil {
		return nil, 0, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			log.Println(err.Error())
		}
	}(cursor, ctx)

	var organizations []Organization
	if err = cursor.All(ctx, &organizations); err != nil {
		return nil, 0, err
	}

	return organizations, total, nil
}
