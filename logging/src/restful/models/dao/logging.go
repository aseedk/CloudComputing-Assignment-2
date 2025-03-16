package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Constants for field names
const (
	FieldId             = "_id"
	FieldURL            = "url"
	FieldMethod         = "method"
	FieldUserId         = "userId"
	FieldOrganizationId = "organizationId"
	FieldRequestBody    = "requestBody"
	FieldRequestQuery   = "requestQuery"
	FieldResponseBody   = "responseBody"
	FieldTimestamp      = "timestamp"
)

// Log represents the log model
type Log struct {
	Id             primitive.ObjectID `bson:"_id,omitempty"`
	URL            string             `bson:"url"`
	Method         string             `bson:"method"`
	UserId         string             `bson:"userId"`
	OrganizationId string             `bson:"organizationId"`
	RequestBody    string             `bson:"requestBody"`
	RequestQuery   string             `bson:"requestQuery"`
	ResponseBody   string             `bson:"responseBody"`
	Timestamp      time.Time          `bson:"timestamp"`
}

// CreateLog inserts a new log entry into the database
func CreateLog(ctx context.Context, logEntry Log) error {
	logEntry.Timestamp = time.Now()
	_, err := LogCollection.InsertOne(ctx, logEntry)
	return err
}

// QueryLogs retrieves logs based on filters
func QueryLogs(ctx context.Context, userId *string, organizationId *string, dateFrom *time.Time, dateTo *time.Time, page int, limit int) ([]Log, int64, error) {
	filter := bson.M{}
	if userId != nil {
		filter[FieldUserId] = *userId
	}
	if organizationId != nil {
		filter[FieldOrganizationId] = *organizationId
	}
	if dateFrom != nil && dateTo != nil {
		filter[FieldTimestamp] = bson.M{"$gte": *dateFrom, "$lte": *dateTo}
	}

	total, err := LogCollection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	queryOptions := options.Find().SetSkip(int64((page - 1) * limit)).SetLimit(int64(limit))

	cursor, err := LogCollection.Find(ctx, filter, queryOptions)
	if err != nil {
		return nil, 0, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			log.Println(err.Error())
		}
	}(cursor, ctx)

	var logs []Log
	if err = cursor.All(ctx, &logs); err != nil {
		return nil, 0, err
	}

	return logs, total, nil
}
