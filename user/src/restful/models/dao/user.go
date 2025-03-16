package dao

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// Constants for field names
const (
	FieldId             = "_id"
	FieldUserId         = "userId"
	FieldFirstName      = "firstName"
	FieldLastName       = "lastName"
	FieldDateOfBirth    = "dateOfBirth"
	FieldGender         = "gender"
	FieldIsActive       = "isActive"
	FieldCountry        = "country"
	FieldEmail          = "email"
	FieldCreatedAt      = "createdAt"
	FieldUpdatedAt      = "updatedAt"
	FieldOrganizationId = "organizationId"
)

// User represents the user model
type User struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	UserId      string             `bson:"userId"`
	FirstName   string             `bson:"firstName"`
	LastName    string             `bson:"lastName"`
	DateOfBirth time.Time          `bson:"dateOfBirth"`
	Gender      bool               `bson:"gender"`
	IsActive    bool               `bson:"isActive"`
	Country     string             `bson:"country"`
	Email       string             `bson:"email"`
	CreatedAt   time.Time          `bson:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt"`
}

type UserOrganization struct {
	UserId         string `bson:"userId"`
	OrganizationId string `bson:"organizationId"`
}

// GenerateUserId generates a unique user ID
func GenerateUserId(ctx context.Context) (string, error) {
	filter := bson.M{"_id": "userId"}
	update := bson.M{"$inc": bson.M{"sequence_value": 1}}

	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)

	var result struct {
		SequenceValue int `bson:"sequence_value"`
	}
	err := IdGeneratorCollection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&result)
	if err != nil {
		log.Println("Error generating user ID:", err)
		return "", err
	}

	return fmt.Sprintf("USR-%06d", result.SequenceValue), nil
}

// CreateUser inserts a new user into the database
func CreateUser(ctx context.Context, user User) error {
	user.Id = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	_, err := UserCollection.InsertOne(ctx, user)
	return err
}

// GetUser fetches a single user by ID
func GetUser(ctx context.Context, userId string) (*User, error) {
	var user User
	err := UserCollection.FindOne(ctx, bson.M{FieldUserId: userId}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// JoinOrganization allows a user to join an organization
func JoinOrganization(ctx context.Context, userId string, organizationId string) error {
	userOrganization := UserOrganization{
		UserId:         userId,
		OrganizationId: organizationId,
	}
	_, err := UserOrganizationsCollection.InsertOne(ctx, userOrganization)
	return err
}

// LeaveOrganization removes the user from an organization
func LeaveOrganization(ctx context.Context, userId string, organizationId string) error {
	filter := bson.M{FieldUserId: userId, FieldOrganizationId: organizationId}

	_, err := UserOrganizationsCollection.DeleteMany(ctx, filter)
	return err
}

// CountUsers returns the count of users matching the filter
func CountUsers(ctx context.Context, filter bson.M) (int, error) {
	total, err := UserCollection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return int(total), nil
}

func QueryUsers(ctx context.Context, organizationId *string, isActive *bool, pageReq PageReq) ([]User, int, error) {
	filter := bson.M{}
	if organizationId != nil {
		filter[FieldOrganizationId] = *organizationId
	}
	if isActive != nil {
		filter[FieldIsActive] = *isActive
	}

	total, err := CountUsers(ctx, filter)
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

	cursor, err := UserCollection.Find(ctx, filter, queryOption)
	if err != nil {
		return nil, 0, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			log.Println(err.Error())
		}
	}(cursor, ctx)

	var users []User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func CreateUserOrganization(ctx context.Context, userOrganization UserOrganization) error {
	_, err := UserOrganizationsCollection.InsertOne(ctx, userOrganization)
	return err
}

func DeleteUserOrganization(ctx context.Context, userId string, organizationId string) error {
	_, err := UserOrganizationsCollection.DeleteOne(ctx, bson.M{FieldUserId: userId, FieldOrganizationId: organizationId})
	return err
}

func ExistUserOrganization(ctx context.Context, userId string, organizationId string) (bool, error) {
	count, err := UserOrganizationsCollection.CountDocuments(ctx, bson.M{FieldUserId: userId, FieldOrganizationId: organizationId})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
