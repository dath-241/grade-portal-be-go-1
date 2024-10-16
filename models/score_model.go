package models // Cùng package với user_model.go

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type InterfaceCoure struct {
}

func Course() *mongo.Collection {
	InitModel("project", "admin")
	return collection
}
