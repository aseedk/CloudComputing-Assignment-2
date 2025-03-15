package dao

import (
	"cloud-computing/organization/organization/src/database"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var (
	OrganizationCollection *mongo.Collection = nil
	IdGeneratorCollection  *mongo.Collection = nil
)

// PageReq represents pagination query parameters
type PageReq struct {
	Page  *int
	Limit *int
	Skip  *int
}

func InitMongoDB() (err error) {
	if OrganizationCollection = database.GetCollection("organizations", "organization"); OrganizationCollection == nil {
		err = errors.New("OrganizationCollection not found")
		return
	} else {
		log.Println("OrganizationCollection created")
	}

	if IdGeneratorCollection = database.GetCollection("idGenerator", "organization"); IdGeneratorCollection == nil {
		err = errors.New("IdGeneratorCollection not found")
		return
	} else {
		log.Println("IdGeneratorCollection created")
	}
	return
}
