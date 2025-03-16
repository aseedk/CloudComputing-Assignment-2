package dao

import (
	"cloud-computing/logging/database"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var (
	LogCollection *mongo.Collection = nil
)

// PageReq represents pagination query parameters
type PageReq struct {
	Page  *int
	Limit *int
	Skip  *int
}

func InitMongoDB() (err error) {
	if LogCollection = database.GetCollection("logging", "log"); LogCollection == nil {
		err = errors.New("OrganizationCollection not found")
		return
	} else {
		log.Println("OrganizationCollection created")
	}

	return
}
