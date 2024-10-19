package models // Cùng package với user_model.go

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type InterfaceCourse struct {
	MS       string `bson:"ms"`
	Credit   int    `bson:"credit"`
	Name     string `bson:"name"`
	Desc     string `bson:"desc"`
	CreateBy string `bson:"createby"`
}

func Course() *mongo.Collection {
	InitModel("project", "admin")
	return collection
}
