package models

import (
	"context"
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

// AddStudentsToCourse adds one or many students to a course
func AddStudentsToCourse(courseID string, students []string) error {
	collection := ClassModel() // Assuming ClassModel returns the MongoDB collection for courses

	filter := bson.M{"courseID": courseID}
	update := bson.M{
			"$addToSet": bson.M{
					"students": bson.M{
							"$each": students,
					},
			},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}