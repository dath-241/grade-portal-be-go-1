package models // Giữ nguyên gói này để dễ dàng nhập khẩu

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type InterfaceAdmin struct {
	ID        bson.ObjectID `bson:"_id,omitempty"` // Tag cho ID
	Email     string        `bson:"email"`         // Tag cho email
	Name      string        `bson:"name"`          // Tag cho role
	Ms        string        `bson:"ms"`
	Faculty   string        `bson:"faculty"`
	CreatedBy bson.ObjectID `bson:"createdBy"` // Tag cho thời gian tạo
}

func AdminModel() *mongo.Collection {
	InitModel("project", "admin")
	return collection
}
