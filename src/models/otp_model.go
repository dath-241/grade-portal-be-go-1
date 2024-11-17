package models

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type InterfaceOtp struct {
	ID        bson.ObjectID `bson:"_id,omitempty"`
	Email     string        `bson:"email"`
	Ms        string        `bson:"ms"`
	Otp       string        `bson:"otb"`
	CreatedAt time.Time     `bson:"created_at"`
}

func OtpModel() *mongo.Collection {
	InitModel("project", "otp")
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "expiredAt", Value: 1}},       // Create index on expiredAt field
		Options: options.Index().SetExpireAfterSeconds(300), // TTL of 0 means it expires as soon as the timestamp is reached
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Fatalf("Failed to create TTL index: %v", err)
	}
	return collection
}
