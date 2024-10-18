package models // Giữ nguyên gói này để dễ dàng nhập khẩu

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type InterfaceUser struct {
	ID        bson.ObjectID `bson:"_id,omitempty"` // Tag cho ID
	Email     string        `bson:"email"`         // Tag cho email
	Password  string        `bson:"password"`      // Tag cho password
	Name      string        `bson:"name"`          // Tag cho role
	Ms        string        `bson:"ms"`
	Faculty   string        `bson:"faculty"`
	Role      string        `bson:"role"`
	CreatedAt time.Time     `bson:"createdAt"` // Tag cho thời gian tạo
	ExpiredAt time.Time     `bson:"expiredAt"` // Tag cho thời gian hết hạn
}

func UserModel() *mongo.Collection {
	InitModel("project", "account")
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
