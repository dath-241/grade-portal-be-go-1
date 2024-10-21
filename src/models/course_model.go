package models // Cùng package với user_model.go

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type InterfaceCourse struct {
	MS        string `bson:"ms"`
	Credit    int    `bson:"credit"`
	Name      string `bson:"name"`
	Desc      string `bson:"desc"`
	CreatedBy string `bson:"createdby"`
}

func Course() *mongo.Collection {
	InitModel("project", "course")
	return collection
}
