package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type InterfaceHallOfFame struct {
	ID       bson.ObjectID `bson:"_id,omitempty"`
	Semester string        `bson:"semester"`
	Tier     []struct {
		CourseId bson.ObjectID `bson:"course_id"`
		Data     [10]struct {
			MSSV string  `bson:"mssv"`
			DTB  float32 `bson:"dtb"`
		} `bson:"data"`
	} `bson:"tier"`
}

func HallOfFameModel() *mongo.Collection {
	InitModel("project", "hall-of-fame")
	return collection
}
