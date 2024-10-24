// config/mongo.go
package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var MongoClient *mongo.Client

func ConnectMongoDB(uri string) {
	var err error
	MongoClient, err = mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Error creating MongoDB client: ", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err != nil {
		log.Fatal("Error connecting to MongoDB: ", err)
	}
	err = MongoClient.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Could not ping MongoDB: ", err)
	}

	log.Println("Connected to MongoDB")
}
