package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type InterfaceClass struct {
	ID            bson.ObjectID `bson:"_id,omitempty"`
	Semester      string        `bson:"semester"`
	Name          string        `bson:"name"` // nhom lop
	CourseId      string        `bson:"course_id"`
	ListStudentId []string      `bson:"listStudent_id"`
	TeacherId     bson.ObjectID `bson:"teacher_id"`
	CreatedBy     bson.ObjectID `bson:"createdBy"` // Tag cho thời gian tạo
	UpdatedBy     bson.ObjectID `bson:"updatedBy"`
}

func ClassModel() *mongo.Collection {
	InitModel("project", "class")
	return collection
}