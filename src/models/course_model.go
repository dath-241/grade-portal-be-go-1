package models // Cùng package với user_model.go

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type InterfaceCourse struct {
	ID        bson.ObjectID `bson:"_id,omitempty"`
	MS        string        `bson:"ms"`
	Credit    int           `bson:"credit"`
	Name      string        `bson:"name"`
	Desc      string        `bson:"desc"`
	HS        [5]int        `bson:"hs"`
	CreatedBy bson.ObjectID `bson:"createdby"`
	UpdatedBy bson.ObjectID `bson:"updatedby"`
}

func CourseModel() *mongo.Collection {
	InitModel("project", "course")
	return collection
}
