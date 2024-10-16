package models

import (
	"LearnGo/config"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

var collection *mongo.Collection

func InitModel(database string, col string) {
	if config.MongoClient == nil {
		log.Fatal("MongoDB client is not initialized")
	}
	collection = config.MongoClient.Database(database).Collection(col)
}
