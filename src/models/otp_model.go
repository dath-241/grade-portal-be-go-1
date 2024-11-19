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
	ExpiredAt time.Time     `bson:"expiredAt"`
}

func OtpModel() *mongo.Collection {
	InitModel("project", "otp")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "expiredAt", Value: 1}},
		Options: options.Index().SetExpireAfterSeconds(300),
	}
	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Fatalf("Failed to create TTL index: %v", err)
	}
	return collection
}
