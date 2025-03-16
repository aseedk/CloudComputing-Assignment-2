package dao

import (
	"cloud-computing/users/database"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var (
	UserCollection              *mongo.Collection = nil
	UserOrganizationsCollection *mongo.Collection = nil
	IdGeneratorCollection       *mongo.Collection = nil
)

// PageReq represents pagination query parameters
type PageReq struct {
	Page  *int
	Limit *int
	Skip  *int
}

func InitMongoDB() (err error) {
	if UserCollection = database.GetCollection("users", "user"); UserCollection == nil {
		err = errors.New("OrganizationCollection not found")
		return
	} else {
		log.Println("OrganizationCollection created")
	}

	if UserOrganizationsCollection = database.GetCollection("users", "user_organizations"); UserOrganizationsCollection == nil {
		err = errors.New("UserOrganizationsCollection not found")
		return
	} else {
		log.Println("UserOrganizationsCollection created")
	}

	if IdGeneratorCollection = database.GetCollection("idGenerator", "user"); IdGeneratorCollection == nil {
		err = errors.New("IdGeneratorCollection not found")
		return
	} else {
		log.Println("IdGeneratorCollection created")
	}
	return
}
