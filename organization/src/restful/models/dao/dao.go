package dao

import (
	"cloud-computing/organization/organization/src/database"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var (
	OrganizationCollection *mongo.Collection = nil
)

func init() {
	if err := InitMongoDB(); err != nil {
		panic(err)
	}
}

func InitMongoDB() (err error) {
	if OrganizationCollection = database.GetCollection("organizations", "organization"); OrganizationCollection == nil {
		err = errors.New("OrganizationCollection not found")
		return
	} else {
		log.Println("OrganizationCollection created")
	}
	return
}
