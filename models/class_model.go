package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type InterfaceClass struct {
	ID            bson.ObjectID `bson:"_id,omitempty"`
	Semester      string        `bson:"semester"`
	Name          string        `bson:"name"` // nhom lop
	CourseId      string        `bson:"course_id"`
	ListStudentId []string      `bson:"listStudent_id"`
	TeacherId     string        `bson:"teacher_id"`
	CreatedBy     time.Time     `bson:"createdBy"` // Tag cho thời gian tạo
	UpdatedAt     string        `bson:"updatedBy"`
}

func ClassModel() *mongo.Collection {
	InitModel("project", "class")
	return collection
}
