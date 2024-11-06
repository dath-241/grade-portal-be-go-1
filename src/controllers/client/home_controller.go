package controller_client

import (
	"LearnGo/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func HomeController(c *gin.Context) {
	CollectionUser := models.AccountModel()
	cursor, _ := CollectionUser.Find(context.TODO(), bson.M{})
	defer cursor.Close(context.TODO())
	var results []models.InterfaceAccount
	cursor.All(context.TODO(), &results)
	id, _ := bson.ObjectIDFromHex("6709d21a3cfbb91af03b9492")
	CollectionUser.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"email": "lyvinhthai321@gmail.com"})

	ids := []string{"6709d21a3cfbb91af03b9492", "670d30c81ab7133c418bd4d8"}
	var objectIDS []bson.ObjectID
	for _, v := range ids {
		ID, _ := bson.ObjectIDFromHex(v)
		objectIDS = append(objectIDS, ID)
	}
	CollectionUser.UpdateMany(context.TODO(), bson.M{"_id": bson.M{"$in": objectIDS}}, bson.M{"$set": bson.M{"name": "Thaideptrai"}})

	c.JSON(http.StatusOK, gin.H{
		"user": results,
	})
}
