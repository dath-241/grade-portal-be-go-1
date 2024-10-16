package models // Giữ nguyên gói này để dễ dàng nhập khẩu

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
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
	return collection
}
