package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type InterfaceClass struct {
	ID            bson.ObjectID `bson:"_id,omitempty"`
	Semester      string        `bson:"semester"`
	Name          string        `bson:"name"` // nhom lop
	CourseId      bson.ObjectID `bson:"course_id"`
	ListStudentMs []string      `bson:"listStudent_ms"`
	TeacherId     bson.ObjectID `bson:"teacher_id"`
	CreatedBy     bson.ObjectID `bson:"createdBy"` // Tag cho thời gian tạo
	UpdatedBy     bson.ObjectID `bson:"updatedBy"`
}
type InterfaceClassStudent struct {
	ID        bson.ObjectID `bson:"_id,omitempty"`
	Semester  string        `bson:"semester"`
	Name      string        `bson:"name"` // nhom lop
	CourseId  bson.ObjectID `bson:"course_id"`
	TeacherId bson.ObjectID `bson:"teacher_id"`
}

func ClassModel() *mongo.Collection {
	InitModel("project", "class")
	return collection
}
