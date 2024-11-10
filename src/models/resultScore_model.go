package models

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type InterfaceScore struct {
	BT  []float32 `bson:"bt"`
	TN  []float32 `bson:"tn"`
	BTL []float32 `bson:"btl"`
	GK  float32   `bson:"gk"`
	CK  float32   `bson:"ck"`
}

type InterfaceResultScore struct {
	ID       bson.ObjectID `bson:"_id,omitempty"`
	Semester string        `bson:"semester"`
	SCORE    []struct {
		MSSV string         `bson:"mssv"`
		Data InterfaceScore `bson:"data"`
	} `bson:"score"`
	ClassID   bson.ObjectID `bson:"class_id"`
	CourseID  bson.ObjectID `bson:"course_id"`
	ExpiredAt time.Time     `bson:"expiredAt"`
	CreatedBy bson.ObjectID `bson:"createdBy"`
	UpdatedBy bson.ObjectID `bson:"updatedBy"`
}

func ResultScoreModel() *mongo.Collection {
	InitModel("project", "resultscore")
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "expiredAt", Value: 1}},     // Create index on expiredAt field
		Options: options.Index().SetExpireAfterSeconds(0), // TTL of 0 means it expires as soon as the timestamp is reached
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Fatalf("Failed to create TTL index: %v", err)
	}
	return collection
}
